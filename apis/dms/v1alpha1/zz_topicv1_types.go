// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

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

type TopicV1InitParameters struct {

	// Indicates the ID of primary DMS instance.
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Total partitions number.
	MaxPartitions *float64 `json:"maxPartitions,omitempty" tf:"max_partitions,omitempty"`

	// Indicates the name of a topic.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Indicates the number of topic partitions,
	// which is used to set the number of concurrently consumed messages.
	// Value range: 1–20. Default value: 3.
	Partition *float64 `json:"partition,omitempty" tf:"partition,omitempty"`

	// Number of remaining partitions.
	RemainPartitions *float64 `json:"remainPartitions,omitempty" tf:"remain_partitions,omitempty"`

	// Indicates the number of replicas,
	// which is configured to ensure data reliability.
	// Value range: 1–3. Default value: 3.
	Replication *float64 `json:"replication,omitempty" tf:"replication,omitempty"`

	// Indicates the retention period of a message. Its default value is 72.
	// Value range: 1–720. Default value: 72. Unit: hour.
	RetentionTime *float64 `json:"retentionTime,omitempty" tf:"retention_time,omitempty"`

	// The partition size of the topic.
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// Indicates whether to enable synchronous flushing.
	// Default value: false. Synchronous flushing compromises performance.
	SyncMessageFlush *bool `json:"syncMessageFlush,omitempty" tf:"sync_message_flush,omitempty"`

	// Indicates whether to enable synchronous replication.
	// After this function is enabled, the acks parameter on the producer client must be set to –1.
	// Otherwise, this parameter does not take effect.
	SyncReplication *bool `json:"syncReplication,omitempty" tf:"sync_replication,omitempty"`
}

type TopicV1Observation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Indicates the ID of primary DMS instance.
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Total partitions number.
	MaxPartitions *float64 `json:"maxPartitions,omitempty" tf:"max_partitions,omitempty"`

	// Indicates the name of a topic.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Indicates the number of topic partitions,
	// which is used to set the number of concurrently consumed messages.
	// Value range: 1–20. Default value: 3.
	Partition *float64 `json:"partition,omitempty" tf:"partition,omitempty"`

	// Number of remaining partitions.
	RemainPartitions *float64 `json:"remainPartitions,omitempty" tf:"remain_partitions,omitempty"`

	// Indicates the number of replicas,
	// which is configured to ensure data reliability.
	// Value range: 1–3. Default value: 3.
	Replication *float64 `json:"replication,omitempty" tf:"replication,omitempty"`

	// Indicates the retention period of a message. Its default value is 72.
	// Value range: 1–720. Default value: 72. Unit: hour.
	RetentionTime *float64 `json:"retentionTime,omitempty" tf:"retention_time,omitempty"`

	// The partition size of the topic.
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// Indicates whether to enable synchronous flushing.
	// Default value: false. Synchronous flushing compromises performance.
	SyncMessageFlush *bool `json:"syncMessageFlush,omitempty" tf:"sync_message_flush,omitempty"`

	// Indicates whether to enable synchronous replication.
	// After this function is enabled, the acks parameter on the producer client must be set to –1.
	// Otherwise, this parameter does not take effect.
	SyncReplication *bool `json:"syncReplication,omitempty" tf:"sync_replication,omitempty"`
}

type TopicV1Parameters struct {

	// Indicates the ID of primary DMS instance.
	// +kubebuilder:validation:Optional
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Total partitions number.
	// +kubebuilder:validation:Optional
	MaxPartitions *float64 `json:"maxPartitions,omitempty" tf:"max_partitions,omitempty"`

	// Indicates the name of a topic.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Indicates the number of topic partitions,
	// which is used to set the number of concurrently consumed messages.
	// Value range: 1–20. Default value: 3.
	// +kubebuilder:validation:Optional
	Partition *float64 `json:"partition,omitempty" tf:"partition,omitempty"`

	// Number of remaining partitions.
	// +kubebuilder:validation:Optional
	RemainPartitions *float64 `json:"remainPartitions,omitempty" tf:"remain_partitions,omitempty"`

	// Indicates the number of replicas,
	// which is configured to ensure data reliability.
	// Value range: 1–3. Default value: 3.
	// +kubebuilder:validation:Optional
	Replication *float64 `json:"replication,omitempty" tf:"replication,omitempty"`

	// Indicates the retention period of a message. Its default value is 72.
	// Value range: 1–720. Default value: 72. Unit: hour.
	// +kubebuilder:validation:Optional
	RetentionTime *float64 `json:"retentionTime,omitempty" tf:"retention_time,omitempty"`

	// The partition size of the topic.
	// +kubebuilder:validation:Optional
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// Indicates whether to enable synchronous flushing.
	// Default value: false. Synchronous flushing compromises performance.
	// +kubebuilder:validation:Optional
	SyncMessageFlush *bool `json:"syncMessageFlush,omitempty" tf:"sync_message_flush,omitempty"`

	// Indicates whether to enable synchronous replication.
	// After this function is enabled, the acks parameter on the producer client must be set to –1.
	// Otherwise, this parameter does not take effect.
	// +kubebuilder:validation:Optional
	SyncReplication *bool `json:"syncReplication,omitempty" tf:"sync_replication,omitempty"`
}

// TopicV1Spec defines the desired state of TopicV1
type TopicV1Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TopicV1Parameters `json:"forProvider"`
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
	InitProvider TopicV1InitParameters `json:"initProvider,omitempty"`
}

// TopicV1Status defines the observed state of TopicV1.
type TopicV1Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TopicV1Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TopicV1 is the Schema for the TopicV1s API. Manages a DMS Topic resource within OpenTelekomCloud.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,opentelekomcloud}
type TopicV1 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.instanceId) || (has(self.initProvider) && has(self.initProvider.instanceId))",message="spec.forProvider.instanceId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   TopicV1Spec   `json:"spec"`
	Status TopicV1Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TopicV1List contains a list of TopicV1s
type TopicV1List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TopicV1 `json:"items"`
}

// Repository type metadata.
var (
	TopicV1_Kind             = "TopicV1"
	TopicV1_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TopicV1_Kind}.String()
	TopicV1_KindAPIVersion   = TopicV1_Kind + "." + CRDGroupVersion.String()
	TopicV1_GroupVersionKind = CRDGroupVersion.WithKind(TopicV1_Kind)
)

func init() {
	SchemeBuilder.Register(&TopicV1{}, &TopicV1List{})
}