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

type AgencyV3InitParameters struct {

	// The name of delegated domain.
	DelegatedDomainName *string `json:"delegatedDomainName,omitempty" tf:"delegated_domain_name,omitempty"`

	// Provides supplementary information about the
	// agency. The value is a string of 0 to 255 characters.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// An array of role names which stand for the
	// permissions to be granted to agency on domain.
	DomainRoles []*string `json:"domainRoles,omitempty" tf:"domain_roles,omitempty"`

	// The name of agency. The name is a string of 1 to 64
	// characters.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// An array of roles and projects which are used to
	// grant permissions to agency on project. The structure is documented below.
	ProjectRole []ProjectRoleInitParameters `json:"projectRole,omitempty" tf:"project_role,omitempty"`
}

type AgencyV3Observation struct {

	// The time when the agency was created.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// The name of delegated domain.
	DelegatedDomainName *string `json:"delegatedDomainName,omitempty" tf:"delegated_domain_name,omitempty"`

	// Provides supplementary information about the
	// agency. The value is a string of 0 to 255 characters.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// An array of role names which stand for the
	// permissions to be granted to agency on domain.
	DomainRoles []*string `json:"domainRoles,omitempty" tf:"domain_roles,omitempty"`

	// Validity period of an agency. The default value is null,
	// indicating that the agency is permanently valid.
	Duration *string `json:"duration,omitempty" tf:"duration,omitempty"`

	// The expiration time of agency
	ExpireTime *string `json:"expireTime,omitempty" tf:"expire_time,omitempty"`

	// The agency ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of agency. The name is a string of 1 to 64
	// characters.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// An array of roles and projects which are used to
	// grant permissions to agency on project. The structure is documented below.
	ProjectRole []ProjectRoleObservation `json:"projectRole,omitempty" tf:"project_role,omitempty"`
}

type AgencyV3Parameters struct {

	// The name of delegated domain.
	// +kubebuilder:validation:Optional
	DelegatedDomainName *string `json:"delegatedDomainName,omitempty" tf:"delegated_domain_name,omitempty"`

	// Provides supplementary information about the
	// agency. The value is a string of 0 to 255 characters.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// An array of role names which stand for the
	// permissions to be granted to agency on domain.
	// +kubebuilder:validation:Optional
	DomainRoles []*string `json:"domainRoles,omitempty" tf:"domain_roles,omitempty"`

	// The name of agency. The name is a string of 1 to 64
	// characters.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// An array of roles and projects which are used to
	// grant permissions to agency on project. The structure is documented below.
	// +kubebuilder:validation:Optional
	ProjectRole []ProjectRoleParameters `json:"projectRole,omitempty" tf:"project_role,omitempty"`
}

type ProjectRoleInitParameters struct {

	// The name of project
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// An array of role names
	Roles []*string `json:"roles,omitempty" tf:"roles,omitempty"`
}

type ProjectRoleObservation struct {

	// The name of project
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// An array of role names
	Roles []*string `json:"roles,omitempty" tf:"roles,omitempty"`
}

type ProjectRoleParameters struct {

	// The name of project
	// +kubebuilder:validation:Optional
	Project *string `json:"project" tf:"project,omitempty"`

	// An array of role names
	// +kubebuilder:validation:Optional
	Roles []*string `json:"roles" tf:"roles,omitempty"`
}

// AgencyV3Spec defines the desired state of AgencyV3
type AgencyV3Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AgencyV3Parameters `json:"forProvider"`
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
	InitProvider AgencyV3InitParameters `json:"initProvider,omitempty"`
}

// AgencyV3Status defines the observed state of AgencyV3.
type AgencyV3Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AgencyV3Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AgencyV3 is the Schema for the AgencyV3s API. Manages a IAM Cgency resource within OpenTelekomCloud.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,opentelekomcloud}
type AgencyV3 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.delegatedDomainName) || (has(self.initProvider) && has(self.initProvider.delegatedDomainName))",message="spec.forProvider.delegatedDomainName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   AgencyV3Spec   `json:"spec"`
	Status AgencyV3Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AgencyV3List contains a list of AgencyV3s
type AgencyV3List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AgencyV3 `json:"items"`
}

// Repository type metadata.
var (
	AgencyV3_Kind             = "AgencyV3"
	AgencyV3_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AgencyV3_Kind}.String()
	AgencyV3_KindAPIVersion   = AgencyV3_Kind + "." + CRDGroupVersion.String()
	AgencyV3_GroupVersionKind = CRDGroupVersion.WithKind(AgencyV3_Kind)
)

func init() {
	SchemeBuilder.Register(&AgencyV3{}, &AgencyV3List{})
}