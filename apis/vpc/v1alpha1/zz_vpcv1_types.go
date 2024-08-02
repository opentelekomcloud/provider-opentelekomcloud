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

type VpcV1InitParameters struct {

	// The range of available subnets in the VPC. The value ranges from
	// 10.0.0.0/8 to 10.255.255.0/24, 172.16.0.0/12 to 172.31.255.0/24,
	// or 192.168.0.0/16 to 192.168.255.0/24.
	Cidr *string `json:"cidr,omitempty" tf:"cidr,omitempty"`

	// A description of the VPC.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of the VPC. The name must be unique for a tenant. The value is a string of
	// no more than 64 characters and can contain digits, letters, underscores (_), and hyphens (-).
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Secondary CIDR block that can be added to VPCs.
	// The value cannot contain the following: 100.64.0.0/1, 214.0.0.0/7, 198.18.0.0/15, 169.254.0.0/16,
	// 0.0.0.0/8, 127.0.0.0/8, 240.0.0.0/4, 172.31.0.0/16, 192.168.0.0/16.
	// Currently, only one secondary CIDR block can be added to each VPC.
	SecondaryCidr *string `json:"secondaryCidr,omitempty" tf:"secondary_cidr,omitempty"`

	// Specifies whether the shared SNAT should be used or not. Is also
	// required for cross-tenant sharing. Shared SNAT only avadilable in eu-de region.
	// Deprecated, VPC Shared SNAT End of Life from 01.03.2024, please do not use.
	Shared *bool `json:"shared,omitempty" tf:"shared,omitempty"`

	// The key/value pairs to associate with the VPC.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type VpcV1Observation struct {

	// The range of available subnets in the VPC. The value ranges from
	// 10.0.0.0/8 to 10.255.255.0/24, 172.16.0.0/12 to 172.31.255.0/24,
	// or 192.168.0.0/16 to 192.168.255.0/24.
	Cidr *string `json:"cidr,omitempty" tf:"cidr,omitempty"`

	// A description of the VPC.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the VPC. The name must be unique for a tenant. The value is a string of
	// no more than 64 characters and can contain digits, letters, underscores (_), and hyphens (-).
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Secondary CIDR block that can be added to VPCs.
	// The value cannot contain the following: 100.64.0.0/1, 214.0.0.0/7, 198.18.0.0/15, 169.254.0.0/16,
	// 0.0.0.0/8, 127.0.0.0/8, 240.0.0.0/4, 172.31.0.0/16, 192.168.0.0/16.
	// Currently, only one secondary CIDR block can be added to each VPC.
	SecondaryCidr *string `json:"secondaryCidr,omitempty" tf:"secondary_cidr,omitempty"`

	// Specifies whether the shared SNAT should be used or not. Is also
	// required for cross-tenant sharing. Shared SNAT only avadilable in eu-de region.
	// Deprecated, VPC Shared SNAT End of Life from 01.03.2024, please do not use.
	Shared *bool `json:"shared,omitempty" tf:"shared,omitempty"`

	// The current status of the desired VPC. Can be either CREATING,
	// OK, DOWN, PENDING_UPDATE, PENDING_DELETE or ERROR.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// The key/value pairs to associate with the VPC.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type VpcV1Parameters struct {

	// The range of available subnets in the VPC. The value ranges from
	// 10.0.0.0/8 to 10.255.255.0/24, 172.16.0.0/12 to 172.31.255.0/24,
	// or 192.168.0.0/16 to 192.168.255.0/24.
	// +kubebuilder:validation:Optional
	Cidr *string `json:"cidr,omitempty" tf:"cidr,omitempty"`

	// A description of the VPC.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of the VPC. The name must be unique for a tenant. The value is a string of
	// no more than 64 characters and can contain digits, letters, underscores (_), and hyphens (-).
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Secondary CIDR block that can be added to VPCs.
	// The value cannot contain the following: 100.64.0.0/1, 214.0.0.0/7, 198.18.0.0/15, 169.254.0.0/16,
	// 0.0.0.0/8, 127.0.0.0/8, 240.0.0.0/4, 172.31.0.0/16, 192.168.0.0/16.
	// Currently, only one secondary CIDR block can be added to each VPC.
	// +kubebuilder:validation:Optional
	SecondaryCidr *string `json:"secondaryCidr,omitempty" tf:"secondary_cidr,omitempty"`

	// Specifies whether the shared SNAT should be used or not. Is also
	// required for cross-tenant sharing. Shared SNAT only avadilable in eu-de region.
	// Deprecated, VPC Shared SNAT End of Life from 01.03.2024, please do not use.
	// +kubebuilder:validation:Optional
	Shared *bool `json:"shared,omitempty" tf:"shared,omitempty"`

	// The key/value pairs to associate with the VPC.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// VpcV1Spec defines the desired state of VpcV1
type VpcV1Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VpcV1Parameters `json:"forProvider"`
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
	InitProvider VpcV1InitParameters `json:"initProvider,omitempty"`
}

// VpcV1Status defines the observed state of VpcV1.
type VpcV1Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VpcV1Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// VpcV1 is the Schema for the VpcV1s API. Manages a VPC resource within OpenTelekomCloud.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,opentelekomcloud}
type VpcV1 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.cidr) || (has(self.initProvider) && has(self.initProvider.cidr))",message="spec.forProvider.cidr is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   VpcV1Spec   `json:"spec"`
	Status VpcV1Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VpcV1List contains a list of VpcV1s
type VpcV1List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VpcV1 `json:"items"`
}

// Repository type metadata.
var (
	VpcV1_Kind             = "VpcV1"
	VpcV1_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VpcV1_Kind}.String()
	VpcV1_KindAPIVersion   = VpcV1_Kind + "." + CRDGroupVersion.String()
	VpcV1_GroupVersionKind = CRDGroupVersion.WithKind(VpcV1_Kind)
)

func init() {
	SchemeBuilder.Register(&VpcV1{}, &VpcV1List{})
}
