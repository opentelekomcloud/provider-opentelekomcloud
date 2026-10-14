/*
Copyright 2022 Upbound Inc.
*/

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	xpv1 "github.com/crossplane/crossplane-runtime/v2/apis/common/v1"
	xpv2 "github.com/crossplane/crossplane-runtime/v2/apis/common/v2"
)

// A ProviderConfigSpec defines the desired state of a ProviderConfig.
type ProviderConfigSpec struct {
	// Credentials required to authenticate to this provider.
	Credentials ProviderCredentials `json:"credentials"`
}

// ProviderCredentials required to authenticate.
type ProviderCredentials struct {
	// Source of the provider credentials.
	// +kubebuilder:validation:Enum=None;Secret;InjectedIdentity;Environment;Filesystem;WebIdentity
	Source xpv1.CredentialsSource `json:"source"`

	// WebIdentity configures authenticating to the cloud using a workload
	// identity. It is only used when Source is WebIdentity.
	// +optional
	WebIdentity *WebIdentity `json:"webIdentity,omitempty"`

	xpv1.CommonCredentialSelectors `json:",inline"`
}

// WebIdentity configures authenticating to the cloud using an OIDC identity
// token (workload identity). The token is read from a filesystem path or a
// Kubernetes secret, exchanged with the T Cloud IAM identity endpoint for an
// OTC token, and optionally swapped for short-lived AK/SK credentials.
type WebIdentity struct {
	// AuthURL is the base URL of the T Cloud IAM identity service, e.g.
	// https://iam.eu-de.otc.t-systems.com.
	AuthURL string `json:"authURL"`

	// IdpID is the ID of the federated identity provider (IdP) registered with
	// T Cloud IAM.
	IdpID string `json:"idpID"`

	// Region is the T Cloud region to operate in, e.g. eu-de.
	Region string `json:"region"`

	// Scope limits the resulting OTC token to a project or a domain. At most
	// one of Project and Domain may be set, and within the selected scope
	// exactly one of Id and Name must be set. When no scope is set the token is
	// unscoped.
	// +optional
	Scope *WebIdentityScope `json:"scope,omitempty"`

	// TokenConfig defines where the OIDC identity token is read from.
	// +kubebuilder:validation:Required
	TokenConfig WebIdentityTokenConfig `json:"tokenConfig"`

	// TemporaryCredentials configures exchanging the OTC token for short-lived
	// AK/SK credentials, required by services that sign requests with AK/SK
	// (e.g. OBS). When omitted, no temporary credentials are issued.
	// +optional
	TemporaryCredentials *WebIdentityTemporaryCredentials `json:"temporaryCredentials,omitempty"`
}

// WebIdentityScope is either a project scope or a domain scope. At most one of
// the two may be set.
type WebIdentityScope struct {
	// Project scope, either by Id or Name (exactly one).
	// +optional
	Project *WebIdentityScopeResource `json:"project,omitempty"`

	// Domain scope, either by Id or Name (exactly one).
	// +optional
	Domain *WebIdentityScopeResource `json:"domain,omitempty"`
}

// WebIdentityScopeResource identifies a project or domain by exactly one of Id
// or Name.
type WebIdentityScopeResource struct {
	// +optional
	Id string `json:"id,omitempty"`

	// +optional
	Name string `json:"name,omitempty"`
}

// WebIdentityTokenConfig defines the source of the OIDC identity token.
type WebIdentityTokenConfig struct {
	// Source of the OIDC identity token.
	// +kubebuilder:validation:Enum=Secret;Filesystem
	// +kubebuilder:validation:Required
	Source xpv1.CredentialsSource `json:"source"`

	// SecretRef is a reference to a secret key holding the OIDC identity token,
	// used when Source is Secret.
	// +optional
	SecretRef *xpv1.SecretKeySelector `json:"secretRef,omitempty"`

	// Fs is a reference to a filesystem path holding the OIDC identity token,
	// used when Source is Filesystem.
	// +optional
	Fs *xpv1.FsSelector `json:"fs,omitempty"`
}

// WebIdentityTemporaryCredentials configures obtaining short-lived AK/SK
// credentials from the OTC token.
type WebIdentityTemporaryCredentials struct {
	// Enabled controls whether temporary credentials are obtained. Defaults to
	// true when this object is present.
	// +optional
	Enabled *bool `json:"enabled,omitempty"`

	// DurationSeconds is the validity period of the temporary credentials in
	// seconds. It must be between 900 (15 minutes) and 86400 (24 hours).
	// Defaults to 3600.
	// +optional
	// +kubebuilder:validation:Minimum=900
	// +kubebuilder:validation:Maximum=86400
	DurationSeconds *int `json:"durationSeconds,omitempty"`
}

// A ProviderConfigStatus reflects the observed state of a ProviderConfig.
type ProviderConfigStatus struct {
	xpv1.ProviderConfigStatus `json:",inline"`
}

// +kubebuilder:object:root=true

// A ProviderConfig configures a OpenTelekomCloud provider.
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="SECRET-NAME",type="string",JSONPath=".spec.credentials.secretRef.name",priority=1
// +kubebuilder:resource:scope=Namespaced
// +kubebuilder:resource:scope=Namespaced,categories={crossplane,provider,opentelekomcloud}
// +kubebuilder:storageversion
type ProviderConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ProviderConfigSpec   `json:"spec"`
	Status ProviderConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProviderConfigList contains a list of ProviderConfig.
type ProviderConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProviderConfig `json:"items"`
}

// +kubebuilder:object:root=true

// A ProviderConfigUsage indicates that a resource is using a ProviderConfig.
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="CONFIG-NAME",type="string",JSONPath=".providerConfigRef.name"
// +kubebuilder:printcolumn:name="RESOURCE-KIND",type="string",JSONPath=".resourceRef.kind"
// +kubebuilder:printcolumn:name="RESOURCE-NAME",type="string",JSONPath=".resourceRef.name"
// +kubebuilder:resource:scope=Namespaced,categories={crossplane,provider,opentelekomcloud}
// +kubebuilder:storageversion
type ProviderConfigUsage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	xpv2.TypedProviderConfigUsage `json:",inline"`
}

// +kubebuilder:object:root=true

// ProviderConfigUsageList contains a list of ProviderConfigUsage
type ProviderConfigUsageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProviderConfigUsage `json:"items"`
}

// +kubebuilder:object:root=true

// A ClusterProviderConfig configures the OpenTelekomCloud provider.
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="SECRET-NAME",type="string",JSONPath=".spec.credentials.secretRef.name",priority=1
// +kubebuilder:resource:scope=Cluster,categories={crossplane,provider,opentelekomcloud}
// +kubebuilder:storageversion
type ClusterProviderConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ProviderConfigSpec   `json:"spec"`
	Status ProviderConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterProviderConfigList contains a list of ProviderConfig.
type ClusterProviderConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterProviderConfig `json:"items"`
}
