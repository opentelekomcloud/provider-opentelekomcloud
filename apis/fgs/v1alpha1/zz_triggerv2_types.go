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

type TriggerV2InitParameters struct {

	// Specifies the detailed configuration of the function trigger event.
	// For various types of trigger parameter configurations, please refer to the
	// documentation.
	EventData *string `json:"eventData,omitempty" tf:"event_data,omitempty"`

	// Specifies the function URN to which the function trigger belongs.
	FunctionUrn *string `json:"functionUrn,omitempty" tf:"function_urn,omitempty"`

	// Specifies the status of the function trigger.
	// The valid values are ACTIVE and DISABLED.
	// For DDS and Kafka triggers the default value is DISABLED, for other triggers= the default value is ACTIVE.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Specifies the type of the function trigger.
	// The valid values are TIMER, APIG, CTS, DDS, DEDICATEDGATEWAY, etc.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type TriggerV2Observation struct {

	// The creation time of the function trigger.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// Specifies the detailed configuration of the function trigger event.
	// For various types of trigger parameter configurations, please refer to the
	// documentation.
	EventData *string `json:"eventData,omitempty" tf:"event_data,omitempty"`

	// Specifies the function URN to which the function trigger belongs.
	FunctionUrn *string `json:"functionUrn,omitempty" tf:"function_urn,omitempty"`

	// resource ID in UUID format.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The region where the function trigger is located.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the status of the function trigger.
	// The valid values are ACTIVE and DISABLED.
	// For DDS and Kafka triggers the default value is DISABLED, for other triggers= the default value is ACTIVE.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Specifies the type of the function trigger.
	// The valid values are TIMER, APIG, CTS, DDS, DEDICATEDGATEWAY, etc.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The latest update time of the function trigger.
	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`
}

type TriggerV2Parameters struct {

	// Specifies the detailed configuration of the function trigger event.
	// For various types of trigger parameter configurations, please refer to the
	// documentation.
	// +kubebuilder:validation:Optional
	EventData *string `json:"eventData,omitempty" tf:"event_data,omitempty"`

	// Specifies the function URN to which the function trigger belongs.
	// +kubebuilder:validation:Optional
	FunctionUrn *string `json:"functionUrn,omitempty" tf:"function_urn,omitempty"`

	// Specifies the status of the function trigger.
	// The valid values are ACTIVE and DISABLED.
	// For DDS and Kafka triggers the default value is DISABLED, for other triggers= the default value is ACTIVE.
	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Specifies the type of the function trigger.
	// The valid values are TIMER, APIG, CTS, DDS, DEDICATEDGATEWAY, etc.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// TriggerV2Spec defines the desired state of TriggerV2
type TriggerV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TriggerV2Parameters `json:"forProvider"`
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
	InitProvider TriggerV2InitParameters `json:"initProvider,omitempty"`
}

// TriggerV2Status defines the observed state of TriggerV2.
type TriggerV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TriggerV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TriggerV2 is the Schema for the TriggerV2s API. Manages an FGS Trigger resource within OpenTelekomCloud.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,opentelekomcloud}
type TriggerV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.eventData) || (has(self.initProvider) && has(self.initProvider.eventData))",message="spec.forProvider.eventData is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.functionUrn) || (has(self.initProvider) && has(self.initProvider.functionUrn))",message="spec.forProvider.functionUrn is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type) || (has(self.initProvider) && has(self.initProvider.type))",message="spec.forProvider.type is a required parameter"
	Spec   TriggerV2Spec   `json:"spec"`
	Status TriggerV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TriggerV2List contains a list of TriggerV2s
type TriggerV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TriggerV2 `json:"items"`
}

// Repository type metadata.
var (
	TriggerV2_Kind             = "TriggerV2"
	TriggerV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TriggerV2_Kind}.String()
	TriggerV2_KindAPIVersion   = TriggerV2_Kind + "." + CRDGroupVersion.String()
	TriggerV2_GroupVersionKind = CRDGroupVersion.WithKind(TriggerV2_Kind)
)

func init() {
	SchemeBuilder.Register(&TriggerV2{}, &TriggerV2List{})
}