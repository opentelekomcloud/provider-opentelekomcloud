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

type DedicatedGeoIPRuleV1InitParameters struct {

	// Protective action.
	// The value can be:
	Action *float64 `json:"action,omitempty" tf:"action,omitempty"`

	// Rule description
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Rule name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The WAF policy ID. Changing this creates a new rule.
	PolicyID *string `json:"policyId,omitempty" tf:"policy_id,omitempty"`

	// Applicable regions. The value can be the region code. For more geographical location codes, see docs "Appendix - Geographic Location Codes."
	// Values:
	RegionCode *string `json:"regionCode,omitempty" tf:"region_code,omitempty"`
}

type DedicatedGeoIPRuleV1Observation struct {

	// Protective action.
	// The value can be:
	Action *float64 `json:"action,omitempty" tf:"action,omitempty"`

	// Timestamp the rule is created.
	CreatedAt *float64 `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// Rule description
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// ID of the rule.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Rule name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The WAF policy ID. Changing this creates a new rule.
	PolicyID *string `json:"policyId,omitempty" tf:"policy_id,omitempty"`

	// Applicable regions. The value can be the region code. For more geographical location codes, see docs "Appendix - Geographic Location Codes."
	// Values:
	RegionCode *string `json:"regionCode,omitempty" tf:"region_code,omitempty"`

	// Rule status. The value can be:
	Status *float64 `json:"status,omitempty" tf:"status,omitempty"`
}

type DedicatedGeoIPRuleV1Parameters struct {

	// Protective action.
	// The value can be:
	// +kubebuilder:validation:Optional
	Action *float64 `json:"action,omitempty" tf:"action,omitempty"`

	// Rule description
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Rule name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The WAF policy ID. Changing this creates a new rule.
	// +kubebuilder:validation:Optional
	PolicyID *string `json:"policyId,omitempty" tf:"policy_id,omitempty"`

	// Applicable regions. The value can be the region code. For more geographical location codes, see docs "Appendix - Geographic Location Codes."
	// Values:
	// +kubebuilder:validation:Optional
	RegionCode *string `json:"regionCode,omitempty" tf:"region_code,omitempty"`
}

// DedicatedGeoIPRuleV1Spec defines the desired state of DedicatedGeoIPRuleV1
type DedicatedGeoIPRuleV1Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DedicatedGeoIPRuleV1Parameters `json:"forProvider"`
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
	InitProvider DedicatedGeoIPRuleV1InitParameters `json:"initProvider,omitempty"`
}

// DedicatedGeoIPRuleV1Status defines the observed state of DedicatedGeoIPRuleV1.
type DedicatedGeoIPRuleV1Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DedicatedGeoIPRuleV1Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// DedicatedGeoIPRuleV1 is the Schema for the DedicatedGeoIPRuleV1s API. Manages a WAF Dedicated Geolocation Access Control Rule resource within OpenTelekomCloud.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,opentelekomcloud}
type DedicatedGeoIPRuleV1 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.action) || (has(self.initProvider) && has(self.initProvider.action))",message="spec.forProvider.action is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.policyId) || (has(self.initProvider) && has(self.initProvider.policyId))",message="spec.forProvider.policyId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.regionCode) || (has(self.initProvider) && has(self.initProvider.regionCode))",message="spec.forProvider.regionCode is a required parameter"
	Spec   DedicatedGeoIPRuleV1Spec   `json:"spec"`
	Status DedicatedGeoIPRuleV1Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DedicatedGeoIPRuleV1List contains a list of DedicatedGeoIPRuleV1s
type DedicatedGeoIPRuleV1List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DedicatedGeoIPRuleV1 `json:"items"`
}

// Repository type metadata.
var (
	DedicatedGeoIPRuleV1_Kind             = "DedicatedGeoIPRuleV1"
	DedicatedGeoIPRuleV1_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DedicatedGeoIPRuleV1_Kind}.String()
	DedicatedGeoIPRuleV1_KindAPIVersion   = DedicatedGeoIPRuleV1_Kind + "." + CRDGroupVersion.String()
	DedicatedGeoIPRuleV1_GroupVersionKind = CRDGroupVersion.WithKind(DedicatedGeoIPRuleV1_Kind)
)

func init() {
	SchemeBuilder.Register(&DedicatedGeoIPRuleV1{}, &DedicatedGeoIPRuleV1List{})
}
