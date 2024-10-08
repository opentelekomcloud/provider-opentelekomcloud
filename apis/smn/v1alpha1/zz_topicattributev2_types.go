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

type TopicAttributeV2InitParameters struct {

	// Attribute name. Valid value is access_policy.
	AttributeName *string `json:"attributeName,omitempty" tf:"attribute_name,omitempty"`

	// Topic attribute value. The value cannot exceed 30 KB.
	TopicAttribute *string `json:"topicAttribute,omitempty" tf:"topic_attribute,omitempty"`

	// Resource identifier of a topic, which is unique.
	TopicUrn *string `json:"topicUrn,omitempty" tf:"topic_urn,omitempty"`
}

type TopicAttributeV2Observation struct {

	// Attribute name. Valid value is access_policy.
	AttributeName *string `json:"attributeName,omitempty" tf:"attribute_name,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Topic attribute value. The value cannot exceed 30 KB.
	TopicAttribute *string `json:"topicAttribute,omitempty" tf:"topic_attribute,omitempty"`

	// Resource identifier of a topic, which is unique.
	TopicUrn *string `json:"topicUrn,omitempty" tf:"topic_urn,omitempty"`
}

type TopicAttributeV2Parameters struct {

	// Attribute name. Valid value is access_policy.
	// +kubebuilder:validation:Optional
	AttributeName *string `json:"attributeName,omitempty" tf:"attribute_name,omitempty"`

	// Topic attribute value. The value cannot exceed 30 KB.
	// +kubebuilder:validation:Optional
	TopicAttribute *string `json:"topicAttribute,omitempty" tf:"topic_attribute,omitempty"`

	// Resource identifier of a topic, which is unique.
	// +kubebuilder:validation:Optional
	TopicUrn *string `json:"topicUrn,omitempty" tf:"topic_urn,omitempty"`
}

// TopicAttributeV2Spec defines the desired state of TopicAttributeV2
type TopicAttributeV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TopicAttributeV2Parameters `json:"forProvider"`
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
	InitProvider TopicAttributeV2InitParameters `json:"initProvider,omitempty"`
}

// TopicAttributeV2Status defines the observed state of TopicAttributeV2.
type TopicAttributeV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TopicAttributeV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// TopicAttributeV2 is the Schema for the TopicAttributeV2s API. Manages an SMN Topic Attribute resource within OpenTelekomCloud.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,opentelekomcloud}
type TopicAttributeV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.attributeName) || (has(self.initProvider) && has(self.initProvider.attributeName))",message="spec.forProvider.attributeName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.topicAttribute) || (has(self.initProvider) && has(self.initProvider.topicAttribute))",message="spec.forProvider.topicAttribute is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.topicUrn) || (has(self.initProvider) && has(self.initProvider.topicUrn))",message="spec.forProvider.topicUrn is a required parameter"
	Spec   TopicAttributeV2Spec   `json:"spec"`
	Status TopicAttributeV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TopicAttributeV2List contains a list of TopicAttributeV2s
type TopicAttributeV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TopicAttributeV2 `json:"items"`
}

// Repository type metadata.
var (
	TopicAttributeV2_Kind             = "TopicAttributeV2"
	TopicAttributeV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TopicAttributeV2_Kind}.String()
	TopicAttributeV2_KindAPIVersion   = TopicAttributeV2_Kind + "." + CRDGroupVersion.String()
	TopicAttributeV2_GroupVersionKind = CRDGroupVersion.WithKind(TopicAttributeV2_Kind)
)

func init() {
	SchemeBuilder.Register(&TopicAttributeV2{}, &TopicAttributeV2List{})
}
