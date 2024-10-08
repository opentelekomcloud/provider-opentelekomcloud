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

type PoliciesInitParameters struct {

	// Permission type. Possible values:
	AccessPolicy *string `json:"accessPolicy,omitempty" tf:"access_policy,omitempty"`

	// DMS instance user name.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type PoliciesObservation struct {

	// Permission type. Possible values:
	AccessPolicy *string `json:"accessPolicy,omitempty" tf:"access_policy,omitempty"`

	// Indicates whether the user is the one selected during topic creation.
	Owner *bool `json:"owner,omitempty" tf:"owner,omitempty"`

	// DMS instance user name.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type PoliciesParameters struct {

	// Permission type. Possible values:
	// +kubebuilder:validation:Optional
	AccessPolicy *string `json:"accessPolicy" tf:"access_policy,omitempty"`

	// DMS instance user name.
	// +kubebuilder:validation:Optional
	Username *string `json:"username" tf:"username,omitempty"`
}

type UserPermissionV1InitParameters struct {

	// Indicates the ID of primary DMS instance.
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Indicates policy configuration for the topic.
	// Supported fields:
	Policies []PoliciesInitParameters `json:"policies,omitempty" tf:"policies,omitempty"`

	// Indicates the name of a topic.
	TopicName *string `json:"topicName,omitempty" tf:"topic_name,omitempty"`
}

type UserPermissionV1Observation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Indicates the ID of primary DMS instance.
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Indicates policy configuration for the topic.
	// Supported fields:
	Policies []PoliciesObservation `json:"policies,omitempty" tf:"policies,omitempty"`

	// Indicates the name of a topic.
	TopicName *string `json:"topicName,omitempty" tf:"topic_name,omitempty"`

	// Indicates topic type. 0: common topic; 1: system (internal) topic.
	TopicType *float64 `json:"topicType,omitempty" tf:"topic_type,omitempty"`
}

type UserPermissionV1Parameters struct {

	// Indicates the ID of primary DMS instance.
	// +kubebuilder:validation:Optional
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Indicates policy configuration for the topic.
	// Supported fields:
	// +kubebuilder:validation:Optional
	Policies []PoliciesParameters `json:"policies,omitempty" tf:"policies,omitempty"`

	// Indicates the name of a topic.
	// +kubebuilder:validation:Optional
	TopicName *string `json:"topicName,omitempty" tf:"topic_name,omitempty"`
}

// UserPermissionV1Spec defines the desired state of UserPermissionV1
type UserPermissionV1Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     UserPermissionV1Parameters `json:"forProvider"`
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
	InitProvider UserPermissionV1InitParameters `json:"initProvider,omitempty"`
}

// UserPermissionV1Status defines the observed state of UserPermissionV1.
type UserPermissionV1Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        UserPermissionV1Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// UserPermissionV1 is the Schema for the UserPermissionV1s API. Manages a DMS User Permissions resource within OpenTelekomCloud.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,opentelekomcloud}
type UserPermissionV1 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.instanceId) || (has(self.initProvider) && has(self.initProvider.instanceId))",message="spec.forProvider.instanceId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.policies) || (has(self.initProvider) && has(self.initProvider.policies))",message="spec.forProvider.policies is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.topicName) || (has(self.initProvider) && has(self.initProvider.topicName))",message="spec.forProvider.topicName is a required parameter"
	Spec   UserPermissionV1Spec   `json:"spec"`
	Status UserPermissionV1Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UserPermissionV1List contains a list of UserPermissionV1s
type UserPermissionV1List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UserPermissionV1 `json:"items"`
}

// Repository type metadata.
var (
	UserPermissionV1_Kind             = "UserPermissionV1"
	UserPermissionV1_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: UserPermissionV1_Kind}.String()
	UserPermissionV1_KindAPIVersion   = UserPermissionV1_Kind + "." + CRDGroupVersion.String()
	UserPermissionV1_GroupVersionKind = CRDGroupVersion.WithKind(UserPermissionV1_Kind)
)

func init() {
	SchemeBuilder.Register(&UserPermissionV1{}, &UserPermissionV1List{})
}
