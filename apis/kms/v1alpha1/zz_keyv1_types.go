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

type KeyV1InitParameters struct {

	// Specifies whether the key is enabled from Pending Deletion state. The value true indicates
	// that the key state Pending Deletion will be cancelled.
	AllowCancelDeletion *bool `json:"allowCancelDeletion,omitempty" tf:"allow_cancel_deletion,omitempty"`

	// Specifies whether the key is enabled. Defaults to true.
	// Changing this updates the state of existing key.
	IsEnabled *bool `json:"isEnabled,omitempty" tf:"is_enabled,omitempty"`

	// The alias in which to create the key. It is required when
	// we create a new key. Changing this updates the alias of key.
	KeyAlias *string `json:"keyAlias,omitempty" tf:"key_alias,omitempty"`

	// The description of the key as viewed in OpenTelekomCloud console.
	// Changing this updates the description of key.
	KeyDescription *string `json:"keyDescription,omitempty" tf:"key_description,omitempty"`

	// Duration in days after which the key is deleted
	// after destruction of the resource, must be between 7 and 1096 days. Defaults to 7.
	// It only is used when delete a key.
	PendingDays *string `json:"pendingDays,omitempty" tf:"pending_days,omitempty"`

	// Region where a key resides. Changing this creates a new key.
	Realm *string `json:"realm,omitempty" tf:"realm,omitempty"`

	// Specifies whether the key is enabled for rotation.
	RotationEnabled *bool `json:"rotationEnabled,omitempty" tf:"rotation_enabled,omitempty"`

	// Rotation interval. The value is an integer ranging from 30 to 365.
	// Set the interval based on how often a CMK is used.
	// If it is frequently used, set a short interval; otherwise, set a long one.
	RotationInterval *float64 `json:"rotationInterval,omitempty" tf:"rotation_interval,omitempty"`

	// Tags key/value pairs to associate with the AutoScaling Group.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type KeyV1Observation struct {

	// Specifies whether the key is enabled from Pending Deletion state. The value true indicates
	// that the key state Pending Deletion will be cancelled.
	AllowCancelDeletion *bool `json:"allowCancelDeletion,omitempty" tf:"allow_cancel_deletion,omitempty"`

	// Creation time (time stamp) of a key.
	CreationDate *string `json:"creationDate,omitempty" tf:"creation_date,omitempty"`

	// Identification of a Master Key. The value 1 indicates a Default
	// Master Key, and the value 0 indicates a key.
	DefaultKeyFlag *string `json:"defaultKeyFlag,omitempty" tf:"default_key_flag,omitempty"`

	// ID of a user domain for the key.
	DomainID *string `json:"domainId,omitempty" tf:"domain_id,omitempty"`

	// Expiration time.
	ExpirationTime *string `json:"expirationTime,omitempty" tf:"expiration_time,omitempty"`

	// The globally unique identifier for the key.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies whether the key is enabled. Defaults to true.
	// Changing this updates the state of existing key.
	IsEnabled *bool `json:"isEnabled,omitempty" tf:"is_enabled,omitempty"`

	// The alias in which to create the key. It is required when
	// we create a new key. Changing this updates the alias of key.
	KeyAlias *string `json:"keyAlias,omitempty" tf:"key_alias,omitempty"`

	// The description of the key as viewed in OpenTelekomCloud console.
	// Changing this updates the description of key.
	KeyDescription *string `json:"keyDescription,omitempty" tf:"key_description,omitempty"`

	// Origin of a key. The default value is kms.
	Origin *string `json:"origin,omitempty" tf:"origin,omitempty"`

	// Duration in days after which the key is deleted
	// after destruction of the resource, must be between 7 and 1096 days. Defaults to 7.
	// It only is used when delete a key.
	PendingDays *string `json:"pendingDays,omitempty" tf:"pending_days,omitempty"`

	// Region where a key resides. Changing this creates a new key.
	Realm *string `json:"realm,omitempty" tf:"realm,omitempty"`

	// Specifies whether the key is enabled for rotation.
	RotationEnabled *bool `json:"rotationEnabled,omitempty" tf:"rotation_enabled,omitempty"`

	// Rotation interval. The value is an integer ranging from 30 to 365.
	// Set the interval based on how often a CMK is used.
	// If it is frequently used, set a short interval; otherwise, set a long one.
	RotationInterval *float64 `json:"rotationInterval,omitempty" tf:"rotation_interval,omitempty"`

	// Number of key rotations.
	RotationNumber *float64 `json:"rotationNumber,omitempty" tf:"rotation_number,omitempty"`

	// Scheduled deletion time (time stamp) of a key.
	ScheduledDeletionDate *string `json:"scheduledDeletionDate,omitempty" tf:"scheduled_deletion_date,omitempty"`

	// Tags key/value pairs to associate with the AutoScaling Group.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type KeyV1Parameters struct {

	// Specifies whether the key is enabled from Pending Deletion state. The value true indicates
	// that the key state Pending Deletion will be cancelled.
	// +kubebuilder:validation:Optional
	AllowCancelDeletion *bool `json:"allowCancelDeletion,omitempty" tf:"allow_cancel_deletion,omitempty"`

	// Specifies whether the key is enabled. Defaults to true.
	// Changing this updates the state of existing key.
	// +kubebuilder:validation:Optional
	IsEnabled *bool `json:"isEnabled,omitempty" tf:"is_enabled,omitempty"`

	// The alias in which to create the key. It is required when
	// we create a new key. Changing this updates the alias of key.
	// +kubebuilder:validation:Optional
	KeyAlias *string `json:"keyAlias,omitempty" tf:"key_alias,omitempty"`

	// The description of the key as viewed in OpenTelekomCloud console.
	// Changing this updates the description of key.
	// +kubebuilder:validation:Optional
	KeyDescription *string `json:"keyDescription,omitempty" tf:"key_description,omitempty"`

	// Duration in days after which the key is deleted
	// after destruction of the resource, must be between 7 and 1096 days. Defaults to 7.
	// It only is used when delete a key.
	// +kubebuilder:validation:Optional
	PendingDays *string `json:"pendingDays,omitempty" tf:"pending_days,omitempty"`

	// Region where a key resides. Changing this creates a new key.
	// +kubebuilder:validation:Optional
	Realm *string `json:"realm,omitempty" tf:"realm,omitempty"`

	// Specifies whether the key is enabled for rotation.
	// +kubebuilder:validation:Optional
	RotationEnabled *bool `json:"rotationEnabled,omitempty" tf:"rotation_enabled,omitempty"`

	// Rotation interval. The value is an integer ranging from 30 to 365.
	// Set the interval based on how often a CMK is used.
	// If it is frequently used, set a short interval; otherwise, set a long one.
	// +kubebuilder:validation:Optional
	RotationInterval *float64 `json:"rotationInterval,omitempty" tf:"rotation_interval,omitempty"`

	// Tags key/value pairs to associate with the AutoScaling Group.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// KeyV1Spec defines the desired state of KeyV1
type KeyV1Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     KeyV1Parameters `json:"forProvider"`
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
	InitProvider KeyV1InitParameters `json:"initProvider,omitempty"`
}

// KeyV1Status defines the observed state of KeyV1.
type KeyV1Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        KeyV1Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// KeyV1 is the Schema for the KeyV1s API. Manages a KMS Key resource within OpenTelekomCloud.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,opentelekomcloud}
type KeyV1 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.keyAlias) || (has(self.initProvider) && has(self.initProvider.keyAlias))",message="spec.forProvider.keyAlias is a required parameter"
	Spec   KeyV1Spec   `json:"spec"`
	Status KeyV1Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KeyV1List contains a list of KeyV1s
type KeyV1List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KeyV1 `json:"items"`
}

// Repository type metadata.
var (
	KeyV1_Kind             = "KeyV1"
	KeyV1_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: KeyV1_Kind}.String()
	KeyV1_KindAPIVersion   = KeyV1_Kind + "." + CRDGroupVersion.String()
	KeyV1_GroupVersionKind = CRDGroupVersion.WithKind(KeyV1_Kind)
)

func init() {
	SchemeBuilder.Register(&KeyV1{}, &KeyV1List{})
}
