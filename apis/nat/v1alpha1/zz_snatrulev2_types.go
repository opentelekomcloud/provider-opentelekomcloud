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

type SnatRuleV2InitParameters struct {

	// Specifies CIDR, which can be in the format of a network segment or
	// a host IP address. This parameter and network_id are alternative. If the value of
	// source_type is 0, the CIDR block must be a subset of the VPC subnet CIDR block. If
	// the value of source_type is 1, the CIDR block must be a CIDR block of Direct Connect
	// and cannot conflict with the VPC CIDR blocks. Changing this creates a new snat rule.
	Cidr *string `json:"cidr,omitempty" tf:"cidr,omitempty"`

	// ID of the floating ip this snat rule connects to.
	// Changing this creates a new snat rule.
	// +crossplane:generate:reference:type=github.com/opentelekomcloud/provider-opentelekomcloud/apis/vpc/v1alpha1.EIPV1
	FloatingIPID *string `json:"floatingIpId,omitempty" tf:"floating_ip_id,omitempty"`

	// Reference to a EIPV1 in vpc to populate floatingIpId.
	// +kubebuilder:validation:Optional
	FloatingIPIDRef *v1.Reference `json:"floatingIpIdRef,omitempty" tf:"-"`

	// Selector for a EIPV1 in vpc to populate floatingIpId.
	// +kubebuilder:validation:Optional
	FloatingIPIDSelector *v1.Selector `json:"floatingIpIdSelector,omitempty" tf:"-"`

	// ID of the nat gateway this snat rule belongs to.
	// Changing this creates a new snat rule.
	// +crossplane:generate:reference:type=github.com/opentelekomcloud/provider-opentelekomcloud/apis/nat/v1alpha1.GatewayV2
	NATGatewayID *string `json:"natGatewayId,omitempty" tf:"nat_gateway_id,omitempty"`

	// Reference to a GatewayV2 in nat to populate natGatewayId.
	// +kubebuilder:validation:Optional
	NATGatewayIDRef *v1.Reference `json:"natGatewayIdRef,omitempty" tf:"-"`

	// Selector for a GatewayV2 in nat to populate natGatewayId.
	// +kubebuilder:validation:Optional
	NATGatewayIDSelector *v1.Selector `json:"natGatewayIdSelector,omitempty" tf:"-"`

	// ID of the network this snat rule connects to. This parameter
	// and cidr are alternative. Changing this creates a new snat rule.
	// +crossplane:generate:reference:type=github.com/opentelekomcloud/provider-opentelekomcloud/apis/vpc/v1alpha1.SubnetV1
	// +crossplane:generate:reference:extractor=github.com/opentelekomcloud/provider-opentelekomcloud/config/common.ExtractNetworkID()
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// Reference to a SubnetV1 in vpc to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDRef *v1.Reference `json:"networkIdRef,omitempty" tf:"-"`

	// Selector for a SubnetV1 in vpc to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDSelector *v1.Selector `json:"networkIdSelector,omitempty" tf:"-"`

	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// 0: Either network_id or cidr can be specified in a VPC. 1:
	// Only cidr can be specified over a dedicated network. Changing this creates a new snat rule.
	SourceType *float64 `json:"sourceType,omitempty" tf:"source_type,omitempty"`
}

type SnatRuleV2Observation struct {

	// Specifies CIDR, which can be in the format of a network segment or
	// a host IP address. This parameter and network_id are alternative. If the value of
	// source_type is 0, the CIDR block must be a subset of the VPC subnet CIDR block. If
	// the value of source_type is 1, the CIDR block must be a CIDR block of Direct Connect
	// and cannot conflict with the VPC CIDR blocks. Changing this creates a new snat rule.
	Cidr *string `json:"cidr,omitempty" tf:"cidr,omitempty"`

	// ID of the floating ip this snat rule connects to.
	// Changing this creates a new snat rule.
	FloatingIPID *string `json:"floatingIpId,omitempty" tf:"floating_ip_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// ID of the nat gateway this snat rule belongs to.
	// Changing this creates a new snat rule.
	NATGatewayID *string `json:"natGatewayId,omitempty" tf:"nat_gateway_id,omitempty"`

	// ID of the network this snat rule connects to. This parameter
	// and cidr are alternative. Changing this creates a new snat rule.
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// 0: Either network_id or cidr can be specified in a VPC. 1:
	// Only cidr can be specified over a dedicated network. Changing this creates a new snat rule.
	SourceType *float64 `json:"sourceType,omitempty" tf:"source_type,omitempty"`
}

type SnatRuleV2Parameters struct {

	// Specifies CIDR, which can be in the format of a network segment or
	// a host IP address. This parameter and network_id are alternative. If the value of
	// source_type is 0, the CIDR block must be a subset of the VPC subnet CIDR block. If
	// the value of source_type is 1, the CIDR block must be a CIDR block of Direct Connect
	// and cannot conflict with the VPC CIDR blocks. Changing this creates a new snat rule.
	// +kubebuilder:validation:Optional
	Cidr *string `json:"cidr,omitempty" tf:"cidr,omitempty"`

	// ID of the floating ip this snat rule connects to.
	// Changing this creates a new snat rule.
	// +crossplane:generate:reference:type=github.com/opentelekomcloud/provider-opentelekomcloud/apis/vpc/v1alpha1.EIPV1
	// +kubebuilder:validation:Optional
	FloatingIPID *string `json:"floatingIpId,omitempty" tf:"floating_ip_id,omitempty"`

	// Reference to a EIPV1 in vpc to populate floatingIpId.
	// +kubebuilder:validation:Optional
	FloatingIPIDRef *v1.Reference `json:"floatingIpIdRef,omitempty" tf:"-"`

	// Selector for a EIPV1 in vpc to populate floatingIpId.
	// +kubebuilder:validation:Optional
	FloatingIPIDSelector *v1.Selector `json:"floatingIpIdSelector,omitempty" tf:"-"`

	// ID of the nat gateway this snat rule belongs to.
	// Changing this creates a new snat rule.
	// +crossplane:generate:reference:type=github.com/opentelekomcloud/provider-opentelekomcloud/apis/nat/v1alpha1.GatewayV2
	// +kubebuilder:validation:Optional
	NATGatewayID *string `json:"natGatewayId,omitempty" tf:"nat_gateway_id,omitempty"`

	// Reference to a GatewayV2 in nat to populate natGatewayId.
	// +kubebuilder:validation:Optional
	NATGatewayIDRef *v1.Reference `json:"natGatewayIdRef,omitempty" tf:"-"`

	// Selector for a GatewayV2 in nat to populate natGatewayId.
	// +kubebuilder:validation:Optional
	NATGatewayIDSelector *v1.Selector `json:"natGatewayIdSelector,omitempty" tf:"-"`

	// ID of the network this snat rule connects to. This parameter
	// and cidr are alternative. Changing this creates a new snat rule.
	// +crossplane:generate:reference:type=github.com/opentelekomcloud/provider-opentelekomcloud/apis/vpc/v1alpha1.SubnetV1
	// +crossplane:generate:reference:extractor=github.com/opentelekomcloud/provider-opentelekomcloud/config/common.ExtractNetworkID()
	// +kubebuilder:validation:Optional
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// Reference to a SubnetV1 in vpc to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDRef *v1.Reference `json:"networkIdRef,omitempty" tf:"-"`

	// Selector for a SubnetV1 in vpc to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDSelector *v1.Selector `json:"networkIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// 0: Either network_id or cidr can be specified in a VPC. 1:
	// Only cidr can be specified over a dedicated network. Changing this creates a new snat rule.
	// +kubebuilder:validation:Optional
	SourceType *float64 `json:"sourceType,omitempty" tf:"source_type,omitempty"`
}

// SnatRuleV2Spec defines the desired state of SnatRuleV2
type SnatRuleV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SnatRuleV2Parameters `json:"forProvider"`
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
	InitProvider SnatRuleV2InitParameters `json:"initProvider,omitempty"`
}

// SnatRuleV2Status defines the observed state of SnatRuleV2.
type SnatRuleV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SnatRuleV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SnatRuleV2 is the Schema for the SnatRuleV2s API. Manages a NAT SNAT Rule resource within OpenTelekomCloud.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,opentelekomcloud}
type SnatRuleV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SnatRuleV2Spec   `json:"spec"`
	Status            SnatRuleV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SnatRuleV2List contains a list of SnatRuleV2s
type SnatRuleV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SnatRuleV2 `json:"items"`
}

// Repository type metadata.
var (
	SnatRuleV2_Kind             = "SnatRuleV2"
	SnatRuleV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SnatRuleV2_Kind}.String()
	SnatRuleV2_KindAPIVersion   = SnatRuleV2_Kind + "." + CRDGroupVersion.String()
	SnatRuleV2_GroupVersionKind = CRDGroupVersion.WithKind(SnatRuleV2_Kind)
)

func init() {
	SchemeBuilder.Register(&SnatRuleV2{}, &SnatRuleV2List{})
}
