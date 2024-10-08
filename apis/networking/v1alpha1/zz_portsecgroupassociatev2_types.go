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

type PortSecgroupAssociateV2InitParameters struct {

	// Whether to replace or append the list of security
	// groups, specified in the security_group_ids. Defaults to false.
	Force *bool `json:"force,omitempty" tf:"force,omitempty"`

	// An UUID of the port to apply security groups to.
	// +crossplane:generate:reference:type=github.com/opentelekomcloud/provider-opentelekomcloud/apis/networking/v1alpha1.PortV2
	PortID *string `json:"portId,omitempty" tf:"port_id,omitempty"`

	// Reference to a PortV2 in networking to populate portId.
	// +kubebuilder:validation:Optional
	PortIDRef *v1.Reference `json:"portIdRef,omitempty" tf:"-"`

	// Selector for a PortV2 in networking to populate portId.
	// +kubebuilder:validation:Optional
	PortIDSelector *v1.Selector `json:"portIdSelector,omitempty" tf:"-"`

	// The region in which to obtain the V2 networking client.
	// A networking client is needed to manage a port. If omitted, the
	// region argument of the provider is used. Changing this creates a new
	// resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// A list of security group IDs to apply to
	// the port. The security groups must be specified by ID and not name (as
	// opposed to how they are configured with the Compute Instance).
	// +crossplane:generate:reference:type=github.com/opentelekomcloud/provider-opentelekomcloud/apis/networking/v1alpha1.SecgroupV2
	// +listType=set
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// References to SecgroupV2 in networking to populate securityGroupIds.
	// +kubebuilder:validation:Optional
	SecurityGroupIdsRefs []v1.Reference `json:"securityGroupIdsRefs,omitempty" tf:"-"`

	// Selector for a list of SecgroupV2 in networking to populate securityGroupIds.
	// +kubebuilder:validation:Optional
	SecurityGroupIdsSelector *v1.Selector `json:"securityGroupIdsSelector,omitempty" tf:"-"`
}

type PortSecgroupAssociateV2Observation struct {

	// The collection of Security Group IDs on the port
	// which have been explicitly and implicitly added.
	// +listType=set
	AllSecurityGroupIds []*string `json:"allSecurityGroupIds,omitempty" tf:"all_security_group_ids,omitempty"`

	// Whether to replace or append the list of security
	// groups, specified in the security_group_ids. Defaults to false.
	Force *bool `json:"force,omitempty" tf:"force,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// An UUID of the port to apply security groups to.
	PortID *string `json:"portId,omitempty" tf:"port_id,omitempty"`

	// The region in which to obtain the V2 networking client.
	// A networking client is needed to manage a port. If omitted, the
	// region argument of the provider is used. Changing this creates a new
	// resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// A list of security group IDs to apply to
	// the port. The security groups must be specified by ID and not name (as
	// opposed to how they are configured with the Compute Instance).
	// +listType=set
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`
}

type PortSecgroupAssociateV2Parameters struct {

	// Whether to replace or append the list of security
	// groups, specified in the security_group_ids. Defaults to false.
	// +kubebuilder:validation:Optional
	Force *bool `json:"force,omitempty" tf:"force,omitempty"`

	// An UUID of the port to apply security groups to.
	// +crossplane:generate:reference:type=github.com/opentelekomcloud/provider-opentelekomcloud/apis/networking/v1alpha1.PortV2
	// +kubebuilder:validation:Optional
	PortID *string `json:"portId,omitempty" tf:"port_id,omitempty"`

	// Reference to a PortV2 in networking to populate portId.
	// +kubebuilder:validation:Optional
	PortIDRef *v1.Reference `json:"portIdRef,omitempty" tf:"-"`

	// Selector for a PortV2 in networking to populate portId.
	// +kubebuilder:validation:Optional
	PortIDSelector *v1.Selector `json:"portIdSelector,omitempty" tf:"-"`

	// The region in which to obtain the V2 networking client.
	// A networking client is needed to manage a port. If omitted, the
	// region argument of the provider is used. Changing this creates a new
	// resource.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// A list of security group IDs to apply to
	// the port. The security groups must be specified by ID and not name (as
	// opposed to how they are configured with the Compute Instance).
	// +crossplane:generate:reference:type=github.com/opentelekomcloud/provider-opentelekomcloud/apis/networking/v1alpha1.SecgroupV2
	// +kubebuilder:validation:Optional
	// +listType=set
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// References to SecgroupV2 in networking to populate securityGroupIds.
	// +kubebuilder:validation:Optional
	SecurityGroupIdsRefs []v1.Reference `json:"securityGroupIdsRefs,omitempty" tf:"-"`

	// Selector for a list of SecgroupV2 in networking to populate securityGroupIds.
	// +kubebuilder:validation:Optional
	SecurityGroupIdsSelector *v1.Selector `json:"securityGroupIdsSelector,omitempty" tf:"-"`
}

// PortSecgroupAssociateV2Spec defines the desired state of PortSecgroupAssociateV2
type PortSecgroupAssociateV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PortSecgroupAssociateV2Parameters `json:"forProvider"`
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
	InitProvider PortSecgroupAssociateV2InitParameters `json:"initProvider,omitempty"`
}

// PortSecgroupAssociateV2Status defines the observed state of PortSecgroupAssociateV2.
type PortSecgroupAssociateV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PortSecgroupAssociateV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// PortSecgroupAssociateV2 is the Schema for the PortSecgroupAssociateV2s API. Manages a VPC Port's Security Groups resource within OpenTelekomCloud.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,opentelekomcloud}
type PortSecgroupAssociateV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PortSecgroupAssociateV2Spec   `json:"spec"`
	Status            PortSecgroupAssociateV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PortSecgroupAssociateV2List contains a list of PortSecgroupAssociateV2s
type PortSecgroupAssociateV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PortSecgroupAssociateV2 `json:"items"`
}

// Repository type metadata.
var (
	PortSecgroupAssociateV2_Kind             = "PortSecgroupAssociateV2"
	PortSecgroupAssociateV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PortSecgroupAssociateV2_Kind}.String()
	PortSecgroupAssociateV2_KindAPIVersion   = PortSecgroupAssociateV2_Kind + "." + CRDGroupVersion.String()
	PortSecgroupAssociateV2_GroupVersionKind = CRDGroupVersion.WithKind(PortSecgroupAssociateV2_Kind)
)

func init() {
	SchemeBuilder.Register(&PortSecgroupAssociateV2{}, &PortSecgroupAssociateV2List{})
}
