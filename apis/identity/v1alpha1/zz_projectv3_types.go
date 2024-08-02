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

type ProjectV3InitParameters struct {

	// A description of the project.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The domain this project belongs to. Changing this
	// creates a new Project.
	DomainID *string `json:"domainId,omitempty" tf:"domain_id,omitempty"`

	// The name of the project. it must start with
	// ID of an existing region and be less than or equal to 64 characters.
	// Example: eu-de_project1.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The parent of this project. Changing this creates
	// a new Project.
	ParentID *string `json:"parentId,omitempty" tf:"parent_id,omitempty"`

	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type ProjectV3Observation struct {

	// A description of the project.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The domain this project belongs to. Changing this
	// creates a new Project.
	DomainID *string `json:"domainId,omitempty" tf:"domain_id,omitempty"`

	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	IsDomain *bool `json:"isDomain,omitempty" tf:"is_domain,omitempty"`

	// The name of the project. it must start with
	// ID of an existing region and be less than or equal to 64 characters.
	// Example: eu-de_project1.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The parent of this project. Changing this creates
	// a new Project.
	ParentID *string `json:"parentId,omitempty" tf:"parent_id,omitempty"`

	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type ProjectV3Parameters struct {

	// A description of the project.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The domain this project belongs to. Changing this
	// creates a new Project.
	// +kubebuilder:validation:Optional
	DomainID *string `json:"domainId,omitempty" tf:"domain_id,omitempty"`

	// The name of the project. it must start with
	// ID of an existing region and be less than or equal to 64 characters.
	// Example: eu-de_project1.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The parent of this project. Changing this creates
	// a new Project.
	// +kubebuilder:validation:Optional
	ParentID *string `json:"parentId,omitempty" tf:"parent_id,omitempty"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

// ProjectV3Spec defines the desired state of ProjectV3
type ProjectV3Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProjectV3Parameters `json:"forProvider"`
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
	InitProvider ProjectV3InitParameters `json:"initProvider,omitempty"`
}

// ProjectV3Status defines the observed state of ProjectV3.
type ProjectV3Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProjectV3Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ProjectV3 is the Schema for the ProjectV3s API. Manages a IAM Project resource within OpenTelekomCloud.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,opentelekomcloud}
type ProjectV3 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   ProjectV3Spec   `json:"spec"`
	Status ProjectV3Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProjectV3List contains a list of ProjectV3s
type ProjectV3List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProjectV3 `json:"items"`
}

// Repository type metadata.
var (
	ProjectV3_Kind             = "ProjectV3"
	ProjectV3_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ProjectV3_Kind}.String()
	ProjectV3_KindAPIVersion   = ProjectV3_Kind + "." + CRDGroupVersion.String()
	ProjectV3_GroupVersionKind = CRDGroupVersion.WithKind(ProjectV3_Kind)
)

func init() {
	SchemeBuilder.Register(&ProjectV3{}, &ProjectV3List{})
}
