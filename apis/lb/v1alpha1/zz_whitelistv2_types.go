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

type WhitelistV2InitParameters struct {

	// Specify whether to enable access control.
	EnableWhitelist *bool `json:"enableWhitelist,omitempty" tf:"enable_whitelist,omitempty"`

	// The Listener ID that the whitelist will be associated with. Changing this creates a new whitelist.
	ListenerID *string `json:"listenerId,omitempty" tf:"listener_id,omitempty"`

	// Required for admins. The UUID of the tenant who owns
	// the whitelist.  Only administrative users can specify a tenant UUID
	// other than their own. Changing this creates a new whitelist.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// Specifies the IP addresses in the whitelist. Use commas(,) to separate
	// the multiple IP addresses.
	Whitelist *string `json:"whitelist,omitempty" tf:"whitelist,omitempty"`
}

type WhitelistV2Observation struct {

	// Specify whether to enable access control.
	EnableWhitelist *bool `json:"enableWhitelist,omitempty" tf:"enable_whitelist,omitempty"`

	// The unique ID for the whitelist.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Listener ID that the whitelist will be associated with. Changing this creates a new whitelist.
	ListenerID *string `json:"listenerId,omitempty" tf:"listener_id,omitempty"`

	// Required for admins. The UUID of the tenant who owns
	// the whitelist.  Only administrative users can specify a tenant UUID
	// other than their own. Changing this creates a new whitelist.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// Specifies the IP addresses in the whitelist. Use commas(,) to separate
	// the multiple IP addresses.
	Whitelist *string `json:"whitelist,omitempty" tf:"whitelist,omitempty"`
}

type WhitelistV2Parameters struct {

	// Specify whether to enable access control.
	// +kubebuilder:validation:Optional
	EnableWhitelist *bool `json:"enableWhitelist,omitempty" tf:"enable_whitelist,omitempty"`

	// The Listener ID that the whitelist will be associated with. Changing this creates a new whitelist.
	// +kubebuilder:validation:Optional
	ListenerID *string `json:"listenerId,omitempty" tf:"listener_id,omitempty"`

	// Required for admins. The UUID of the tenant who owns
	// the whitelist.  Only administrative users can specify a tenant UUID
	// other than their own. Changing this creates a new whitelist.
	// +kubebuilder:validation:Optional
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// Specifies the IP addresses in the whitelist. Use commas(,) to separate
	// the multiple IP addresses.
	// +kubebuilder:validation:Optional
	Whitelist *string `json:"whitelist,omitempty" tf:"whitelist,omitempty"`
}

// WhitelistV2Spec defines the desired state of WhitelistV2
type WhitelistV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     WhitelistV2Parameters `json:"forProvider"`
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
	InitProvider WhitelistV2InitParameters `json:"initProvider,omitempty"`
}

// WhitelistV2Status defines the observed state of WhitelistV2.
type WhitelistV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        WhitelistV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// WhitelistV2 is the Schema for the WhitelistV2s API. Manages a ELB Whitelist resource within OpenTelekomCloud.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,opentelekomcloud}
type WhitelistV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.listenerId) || (has(self.initProvider) && has(self.initProvider.listenerId))",message="spec.forProvider.listenerId is a required parameter"
	Spec   WhitelistV2Spec   `json:"spec"`
	Status WhitelistV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WhitelistV2List contains a list of WhitelistV2s
type WhitelistV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WhitelistV2 `json:"items"`
}

// Repository type metadata.
var (
	WhitelistV2_Kind             = "WhitelistV2"
	WhitelistV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: WhitelistV2_Kind}.String()
	WhitelistV2_KindAPIVersion   = WhitelistV2_Kind + "." + CRDGroupVersion.String()
	WhitelistV2_GroupVersionKind = CRDGroupVersion.WithKind(WhitelistV2_Kind)
)

func init() {
	SchemeBuilder.Register(&WhitelistV2{}, &WhitelistV2List{})
}