/*
Copyright 2021 Upbound Inc.
*/

package clients

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"

	xpv1 "github.com/crossplane/crossplane-runtime/v2/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/v2/pkg/resource"
	"github.com/pkg/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"

	namespacedv1beta1 "github.com/opentelekomcloud/provider-opentelekomcloud/apis/namespaced/v1beta1"
)

const (
	// sourceWebIdentity is the credentials source value that triggers the web
	// identity (workload identity) authentication flow.
	sourceWebIdentity xpv1.CredentialsSource = "WebIdentity"

	errWebIdentityRequired = "spec.credentials.webIdentity must be set when credentials.source is WebIdentity"
	errTokenConfigRequired = "spec.credentials.webIdentity.tokenConfig must be set"
	errTokenConfigSource   = "spec.credentials.webIdentity.tokenConfig.source must be Secret or Filesystem"
	errScopeMutuallyExcl   = "spec.credentials.webIdentity.scope: at most one of project and domain may be set"
	errScopeEmpty          = "spec.credentials.webIdentity.scope: within the selected scope exactly one of id and name must be set"
	errTempDurationInvalid = "spec.credentials.webIdentity.temporaryCredentials.durationSeconds must be between 900 and 86400"
	errEmptyToken          = "could not read a non-empty OIDC identity token"
	errExchangeToken       = "cannot exchange OIDC identity token with IAM"
	errTempCredentials     = "cannot obtain temporary credentials"
)

const defaultTempDurationSeconds = 3600

// buildWebIdentityConfig authenticates using an OIDC identity token and returns
// the Terraform provider configuration to use.
func buildWebIdentityConfig(ctx context.Context, c client.Client, pcSpec *namespacedv1beta1.ProviderConfigSpec) (map[string]any, error) {
	wi := pcSpec.Credentials.WebIdentity
	if wi == nil {
		return nil, errors.New(errWebIdentityRequired)
	}
	tc := &wi.TokenConfig
	if err := validateTokenConfig(tc); err != nil {
		return nil, err
	}
	scope, err := buildScopeMap(wi.Scope)
	if err != nil {
		return nil, err
	}

	token, err := readIdentityToken(ctx, c, tc)
	if err != nil {
		return nil, err
	}

	iamToken, err := exchangeIdToken(ctx, wi.AuthURL, wi.IdpID, token, scope)
	if err != nil {
		return nil, errors.Wrap(err, errExchangeToken)
	}

	cfg := map[string]any{
		"auth_url": normalizeAuthURLV3(wi.AuthURL),
		"region":   wi.Region,
		"token":    iamToken,
	}

	for k, v := range projectTenantFields(wi.Scope) {
		cfg[k] = v
	}

	// TODO(webidentity/domain): a domain-scoped token only carries domain-level
	// endpoints in its catalog and cannot create project-scoped resources, so it
	// currently fails with "No suitable endpoint could be found in the service
	// catalog" when a project-scoped resource is created. Decide how to handle
	// this (reject early, require a project, or resolve the target project).

	if wi.TemporaryCredentials != nil && tempCredsEnabled(wi.TemporaryCredentials.Enabled) {
		if err := addTemporaryCredentials(ctx, wi, iamToken, cfg); err != nil {
			return nil, err
		}
	}

	return cfg, nil
}

// projectTenantFields returns the Terraform tenant fields for a project-scoped
// token. The Terraform OTC provider requires a tenant to be set (see its
// validateProject check), and a project-scoped token only exposes project-scoped
// service endpoints in its catalog. Propagating the project keeps endpoint
// resolution from failing with "No suitable endpoint could be found in the
// service catalog". The project id is preferred over the name. A nil scope, an
// empty scope or a domain scope return nil.
func projectTenantFields(scope *namespacedv1beta1.WebIdentityScope) map[string]any {
	if scope == nil || scope.Project == nil {
		return nil
	}
	if scope.Project.Id != "" {
		return map[string]any{"tenant_id": scope.Project.Id}
	}
	return map[string]any{"tenant_name": scope.Project.Name}
}

// addTemporaryCredentials exchanges the OTC token for short-lived AK/SK and adds
// them to the Terraform configuration.
func addTemporaryCredentials(ctx context.Context, wi *namespacedv1beta1.WebIdentity, iamToken string, cfg map[string]any) error {
	duration := defaultTempDurationSeconds
	if wi.TemporaryCredentials.DurationSeconds != nil {
		duration = *wi.TemporaryCredentials.DurationSeconds
	}
	if duration < 900 || duration > 86400 {
		return errors.New(errTempDurationInvalid)
	}
	cred, err := createTemporaryCredentials(ctx, wi.AuthURL, iamToken, duration)
	if err != nil {
		return errors.Wrap(err, errTempCredentials)
	}
	cfg["access_key"] = cred.AccessKey
	cfg["secret_key"] = cred.SecretKey
	cfg["security_token"] = cred.SecurityToken
	return nil
}

// validateTokenConfig checks that the token source is one of the supported
// values (Secret or Filesystem).
func validateTokenConfig(tc *namespacedv1beta1.WebIdentityTokenConfig) error {
	if tc.Source == "" {
		return errors.New(errTokenConfigRequired)
	}
	if tc.Source != xpv1.CredentialsSourceSecret && tc.Source != xpv1.CredentialsSourceFilesystem {
		return errors.Errorf("%s (got %q)", errTokenConfigSource, tc.Source)
	}
	return nil
}

// readIdentityToken reads the OIDC identity token from the configured source
// for every call, so that rotation of projected tokens is observed. Leading and
// trailing whitespace is stripped.
func readIdentityToken(ctx context.Context, c client.Client, tc *namespacedv1beta1.WebIdentityTokenConfig) (string, error) {
	selectors := xpv1.CommonCredentialSelectors{
		SecretRef: tc.SecretRef,
		Fs:        tc.Fs,
	}
	data, err := resource.CommonCredentialExtractor(ctx, tc.Source, c, selectors)
	if err != nil {
		return "", errors.Wrap(err, "cannot read OIDC identity token")
	}
	token := strings.TrimSpace(string(data))
	if token == "" {
		return "", errors.New(errEmptyToken)
	}
	return token, nil
}

// buildScopeMap validates the scope configuration and returns the request
// scope object. A nil map means the token is unscoped.
func buildScopeMap(scope *namespacedv1beta1.WebIdentityScope) (map[string]any, error) {
	if scope == nil {
		return nil, nil
	}
	if scope.Project != nil && scope.Domain != nil {
		return nil, errors.New(errScopeMutuallyExcl)
	}

	var selected *namespacedv1beta1.WebIdentityScopeResource
	switch {
	case scope.Project != nil:
		selected = scope.Project
	case scope.Domain != nil:
		selected = scope.Domain
	default:
		return nil, nil
	}

	res, err := buildScopeResourceMap(selected)
	if err != nil {
		return nil, err
	}
	key := "project"
	if scope.Domain != nil {
		key = "domain"
	}
	return map[string]any{key: res}, nil
}

func buildScopeResourceMap(r *namespacedv1beta1.WebIdentityScopeResource) (map[string]any, error) {
	if (r.Id == "") == (r.Name == "") {
		return nil, errors.New(errScopeEmpty)
	}
	if r.Id != "" {
		return map[string]any{"id": r.Id}, nil
	}
	return map[string]any{"name": r.Name}, nil
}

func tempCredsEnabled(enabled *bool) bool {
	if enabled == nil {
		return true
	}
	return *enabled
}

// normalizeAuthURLV3 returns the auth URL in the form expected by the Terraform
// OTC provider (ending in /v3), stripping any existing /v3 or /v3.0 suffix.
func normalizeAuthURLV3(authURL string) string {
	u := strings.TrimRight(authURL, "/")
	u = strings.TrimSuffix(u, "/v3.0")
	u = strings.TrimSuffix(u, "/v3")
	return u + "/v3"
}

// normalizeAuthURLBase returns the base auth URL without any version suffix,
// e.g. https://iam.eu-de.otc.t-systems.com
func normalizeAuthURLBase(authURL string) string {
	u := strings.TrimRight(authURL, "/")
	u = strings.TrimSuffix(u, "/v3.0")
	u = strings.TrimSuffix(u, "/v3")
	return u
}

// buildIdTokenBody assembles the body for the OS-AUTH/id-token/tokens request:
// the ID token is carried under auth.id_token.id, with an optional auth.scope
// (a project or a domain). See the IAM "obtaining a token by using an ID token"
// API reference.
func buildIdTokenBody(idToken string, scope map[string]any) ([]byte, error) {
	auth := map[string]any{
		"id_token": map[string]any{
			"id": idToken,
		},
	}
	if scope != nil {
		auth["scope"] = scope
	}
	return json.Marshal(map[string]any{"auth": auth})
}

// exchangeIdToken exchanges an OIDC identity token for an OTC token. The signed
// token is returned in the X-Subject-Token response header.
func exchangeIdToken(ctx context.Context, authURL, idpID, idToken string, scope map[string]any) (string, error) {
	url := normalizeAuthURLBase(authURL) + "/v3.0/OS-AUTH/id-token/tokens"

	body, err := buildIdTokenBody(idToken, scope)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Idp-Id", idpID)

	resp, err := newHTTPClient().Do(req)
	if err != nil {
		return "", err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusCreated {
		return "", errors.Errorf("identity token exchange failed with status %d: %s", resp.StatusCode, readBody(resp))
	}

	if t := strings.TrimSpace(resp.Header.Get("X-Subject-Token")); t != "" {
		return t, nil
	}
	return "", errors.New("identity token exchange response did not contain an X-Subject-Token")
}

type tempCredential struct {
	AccessKey     string `json:"access"`
	SecretKey     string `json:"secret"`
	SecurityToken string `json:"securitytoken"`
	ExpiresAt     string `json:"expires_at"`
}

type tempCredentialResponse struct {
	Credential tempCredential `json:"credential"`
}

// createTemporaryCredentials exchanges an OTC token for short-lived AK/SK
// credentials.
func createTemporaryCredentials(ctx context.Context, authURL, iamToken string, durationSeconds int) (*tempCredential, error) {
	url := normalizeAuthURLBase(authURL) + "/v3.0/OS-CREDENTIAL/securitytokens"

	body := map[string]any{
		"auth": map[string]any{
			"identity": map[string]any{
				"methods": []string{"token"},
				"token": map[string]any{
					"id":               iamToken,
					"duration_seconds": durationSeconds,
				},
			},
		},
	}

	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Auth-Token", iamToken)

	resp, err := newHTTPClient().Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusCreated {
		return nil, errors.Errorf("temporary credentials request failed with status %d: %s", resp.StatusCode, readBody(resp))
	}

	var out tempCredentialResponse
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, errors.Wrap(err, "cannot decode temporary credentials response")
	}
	return &out.Credential, nil
}

func readBody(resp *http.Response) string {
	b, _ := io.ReadAll(io.LimitReader(resp.Body, 4096))
	return strings.TrimSpace(string(b))
}

func newHTTPClient() *http.Client {
	return &http.Client{Timeout: 30 * time.Second}
}
