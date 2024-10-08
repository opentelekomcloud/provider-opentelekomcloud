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

type BandwidthV2InitParameters struct {

	// Specifies the bandwidth name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the bandwidth size.
	// The value ranges from 5 Mbit/s to 1000 Mbit/s by default.
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`
}

type BandwidthV2Observation struct {

	// Specifies the bandwidth ID, which uniquely identifies the bandwidth.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the bandwidth name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the bandwidth size.
	// The value ranges from 5 Mbit/s to 1000 Mbit/s by default.
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// Specifies the bandwidth status.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type BandwidthV2Parameters struct {

	// Specifies the bandwidth name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the bandwidth size.
	// The value ranges from 5 Mbit/s to 1000 Mbit/s by default.
	// +kubebuilder:validation:Optional
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`
}

// BandwidthV2Spec defines the desired state of BandwidthV2
type BandwidthV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BandwidthV2Parameters `json:"forProvider"`
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
	InitProvider BandwidthV2InitParameters `json:"initProvider,omitempty"`
}

// BandwidthV2Status defines the observed state of BandwidthV2.
type BandwidthV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BandwidthV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// BandwidthV2 is the Schema for the BandwidthV2s API. Manages a VPC Bandwidth resource within OpenTelekomCloud.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,opentelekomcloud}
type BandwidthV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.size) || (has(self.initProvider) && has(self.initProvider.size))",message="spec.forProvider.size is a required parameter"
	Spec   BandwidthV2Spec   `json:"spec"`
	Status BandwidthV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BandwidthV2List contains a list of BandwidthV2s
type BandwidthV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BandwidthV2 `json:"items"`
}

// Repository type metadata.
var (
	BandwidthV2_Kind             = "BandwidthV2"
	BandwidthV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BandwidthV2_Kind}.String()
	BandwidthV2_KindAPIVersion   = BandwidthV2_Kind + "." + CRDGroupVersion.String()
	BandwidthV2_GroupVersionKind = CRDGroupVersion.WithKind(BandwidthV2_Kind)
)

func init() {
	SchemeBuilder.Register(&BandwidthV2{}, &BandwidthV2List{})
}
