/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type CredentialV3InitParameters struct {

	// Description of the access key.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Either a base-64 encoded PGP public key, or a keybase username in the form
	// keybase:some_person_that_exists. Changing this creates a new resource.
	PgpKey *string `json:"pgpKey,omitempty" tf:"pgp_key,omitempty"`

	// Status of the access key to be changed to. The value can be active or inactive.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// IAM user ID. If not set, will create AK/SK for yourself.
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`
}

type CredentialV3Observation struct {

	// Time of the access key creation.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// Description of the access key.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	KeyFingerprint *string `json:"keyFingerprint,omitempty" tf:"key_fingerprint,omitempty"`

	// Time of the access key last usage.
	LastUseTime *string `json:"lastUseTime,omitempty" tf:"last_use_time,omitempty"`

	// Either a base-64 encoded PGP public key, or a keybase username in the form
	// keybase:some_person_that_exists. Changing this creates a new resource.
	PgpKey *string `json:"pgpKey,omitempty" tf:"pgp_key,omitempty"`

	// Status of the access key to be changed to. The value can be active or inactive.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// IAM user ID. If not set, will create AK/SK for yourself.
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`
}

type CredentialV3Parameters struct {

	// Description of the access key.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Either a base-64 encoded PGP public key, or a keybase username in the form
	// keybase:some_person_that_exists. Changing this creates a new resource.
	// +kubebuilder:validation:Optional
	PgpKey *string `json:"pgpKey,omitempty" tf:"pgp_key,omitempty"`

	// Status of the access key to be changed to. The value can be active or inactive.
	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// IAM user ID. If not set, will create AK/SK for yourself.
	// +kubebuilder:validation:Optional
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`
}

// CredentialV3Spec defines the desired state of CredentialV3
type CredentialV3Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CredentialV3Parameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider CredentialV3InitParameters `json:"initProvider,omitempty"`
}

// CredentialV3Status defines the observed state of CredentialV3.
type CredentialV3Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CredentialV3Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// CredentialV3 is the Schema for the CredentialV3s API. Manages a IAM Credential resource within OpenTelekomCloud.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,opentelekomcloud}
type CredentialV3 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CredentialV3Spec   `json:"spec"`
	Status            CredentialV3Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CredentialV3List contains a list of CredentialV3s
type CredentialV3List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CredentialV3 `json:"items"`
}

// Repository type metadata.
var (
	CredentialV3_Kind             = "CredentialV3"
	CredentialV3_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CredentialV3_Kind}.String()
	CredentialV3_KindAPIVersion   = CredentialV3_Kind + "." + CRDGroupVersion.String()
	CredentialV3_GroupVersionKind = CRDGroupVersion.WithKind(CredentialV3_Kind)
)

func init() {
	SchemeBuilder.Register(&CredentialV3{}, &CredentialV3List{})
}
