/*
Copyright 2021 Upbound Inc.
*/

package clients

import (
	"context"
	"encoding/json"

	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane/upjet/pkg/terraform"

	"github.com/opentelekomcloud/provider-opentelekomcloud/apis/v1beta1"
)

const (
	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal opentelekomcloud credentials as JSON"
)

// TerraformSetupBuilder builds Terraform a terraform.SetupFn function which
// returns Terraform provider setup configuration
func TerraformSetupBuilder(version, providerSource, providerVersion string) terraform.SetupFn {
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{
			Version: version,
			Requirement: terraform.ProviderRequirement{
				Source:  providerSource,
				Version: providerVersion,
			},
		}

		configRef := mg.GetProviderConfigReference()
		if configRef == nil {
			return ps, errors.New(errNoProviderConfig)
		}
		pc := &v1beta1.ProviderConfig{}
		if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
			return ps, errors.Wrap(err, errGetProviderConfig)
		}

		t := resource.NewProviderConfigUsageTracker(client, &v1beta1.ProviderConfigUsage{})
		if err := t.Track(ctx, mg); err != nil {
			return ps, errors.Wrap(err, errTrackUsage)
		}

		data, err := resource.CommonCredentialExtractor(ctx, pc.Spec.Credentials.Source, client, pc.Spec.Credentials.CommonCredentialSelectors)
		if err != nil {
			return ps, errors.Wrap(err, errExtractCredentials)
		}
		creds := map[string]string{}
		if err := json.Unmarshal(data, &creds); err != nil {
			return ps, errors.Wrap(err, errUnmarshalCredentials)
		}

		// Set credentials in Terraform provider configuration.
		// Keep in sync with configuration options of Terraform OpenStack provider:
		// https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/latest/docs
		// Exceptions:
		// "cloud" is not used, because we don't support giving a clouds.yaml directly to the provider.
		ps.Configuration = map[string]any{
			"access_key":            creds["access_key"],
			"secret_key":            creds["secret_key"],
			"auth_url":              creds["auth_url"],
			"region":                creds["region"],
			"user_name":             creds["user_name"],
			"user_id":               creds["user_id"],
			"tenant_id":             creds["tenant_id"],
			"tenant_name":           creds["tenant_name"],
			"password":              creds["password"],
			"token":                 creds["token"],
			"security_token":        creds["security_token"],
			"passcode":              creds["passcode"],
			"domain_id":             creds["domain_id"],
			"domain_name":           creds["domain_name"],
			"insecure":              creds["insecure"],
			"endpoint_type":         creds["endpoint_type"],
			"cacert_file":           creds["cacert_file"],
			"cert":                  creds["cert"],
			"key":                   creds["key"],
			"swauth":                creds["swauth"],
			"agency_name":           creds["agency_name"],
			"agency_domain_name":    creds["agency_domain_name"],
			"delegated_project":     creds["delegated_project"],
			"allow_reauth":          creds["allow_reauth"],
			"max_retries":           creds["max_retries"],
			"max_backoff_retries":   creds["max_backoff_retries"],
			"backoff_retry_timeout": creds["backoff_retry_timeout"],
		}
		return ps, nil
	}
}
