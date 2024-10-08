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

type TopicV2InitParameters struct {

	// Topic display name, which is presented as the
	// name of the email sender in an email message.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The name of the topic to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The project name for the topic.
	ProjectName *string `json:"projectName,omitempty" tf:"project_name,omitempty"`

	// Tags key/value pairs to associate with the instance.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type TopicV2Observation struct {

	// Time when the topic was created.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// Topic display name, which is presented as the
	// name of the email sender in an email message.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the topic to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The project name for the topic.
	ProjectName *string `json:"projectName,omitempty" tf:"project_name,omitempty"`

	// Message pushing policy. 0 indicates that the message
	// sending fails and the message is cached in the queue. 1 indicates that the
	// failed message is discarded.
	PushPolicy *float64 `json:"pushPolicy,omitempty" tf:"push_policy,omitempty"`

	// Tags key/value pairs to associate with the instance.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Resource identifier of a topic, which is unique.
	TopicUrn *string `json:"topicUrn,omitempty" tf:"topic_urn,omitempty"`

	// Time when the topic was updated.
	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`
}

type TopicV2Parameters struct {

	// Topic display name, which is presented as the
	// name of the email sender in an email message.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The name of the topic to be created.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The project name for the topic.
	// +kubebuilder:validation:Optional
	ProjectName *string `json:"projectName,omitempty" tf:"project_name,omitempty"`

	// Tags key/value pairs to associate with the instance.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// TopicV2Spec defines the desired state of TopicV2
type TopicV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TopicV2Parameters `json:"forProvider"`
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
	InitProvider TopicV2InitParameters `json:"initProvider,omitempty"`
}

// TopicV2Status defines the observed state of TopicV2.
type TopicV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TopicV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// TopicV2 is the Schema for the TopicV2s API. Manages an SMN Topic resource within OpenTelekomCloud.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,opentelekomcloud}
type TopicV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   TopicV2Spec   `json:"spec"`
	Status TopicV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TopicV2List contains a list of TopicV2s
type TopicV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TopicV2 `json:"items"`
}

// Repository type metadata.
var (
	TopicV2_Kind             = "TopicV2"
	TopicV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TopicV2_Kind}.String()
	TopicV2_KindAPIVersion   = TopicV2_Kind + "." + CRDGroupVersion.String()
	TopicV2_GroupVersionKind = CRDGroupVersion.WithKind(TopicV2_Kind)
)

func init() {
	SchemeBuilder.Register(&TopicV2{}, &TopicV2List{})
}
