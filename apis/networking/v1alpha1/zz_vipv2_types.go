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

type VipV2InitParameters struct {

	// IP address desired in the subnet for this vip.
	// If you don't specify ip_address, an available IP address from
	// the specified subnet will be allocated to this vip.
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// A unique name for the vip.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the network to attach the vip to.
	// Changing this creates a new vip.
	// +crossplane:generate:reference:type=github.com/opentelekomcloud/provider-opentelekomcloud/apis/networking/v1alpha1.NetworkV2
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// Reference to a NetworkV2 in networking to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDRef *v1.Reference `json:"networkIdRef,omitempty" tf:"-"`

	// Selector for a NetworkV2 in networking to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDSelector *v1.Selector `json:"networkIdSelector,omitempty" tf:"-"`

	// Subnet in which to allocate IP address for this vip.
	// Changing this creates a new vip.
	// +crossplane:generate:reference:type=github.com/opentelekomcloud/provider-opentelekomcloud/apis/networking/v1alpha1.SubnetV2
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a SubnetV2 in networking to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a SubnetV2 in networking to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`
}

type VipV2Observation struct {

	// The device owner of the vip.
	DeviceOwner *string `json:"deviceOwner,omitempty" tf:"device_owner,omitempty"`

	// The ID of the vip.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// IP address desired in the subnet for this vip.
	// If you don't specify ip_address, an available IP address from
	// the specified subnet will be allocated to this vip.
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// A unique name for the vip.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the network to attach the vip to.
	// Changing this creates a new vip.
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// The status of vip.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Subnet in which to allocate IP address for this vip.
	// Changing this creates a new vip.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// The tenant ID of the vip.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type VipV2Parameters struct {

	// IP address desired in the subnet for this vip.
	// If you don't specify ip_address, an available IP address from
	// the specified subnet will be allocated to this vip.
	// +kubebuilder:validation:Optional
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// A unique name for the vip.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the network to attach the vip to.
	// Changing this creates a new vip.
	// +crossplane:generate:reference:type=github.com/opentelekomcloud/provider-opentelekomcloud/apis/networking/v1alpha1.NetworkV2
	// +kubebuilder:validation:Optional
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// Reference to a NetworkV2 in networking to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDRef *v1.Reference `json:"networkIdRef,omitempty" tf:"-"`

	// Selector for a NetworkV2 in networking to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDSelector *v1.Selector `json:"networkIdSelector,omitempty" tf:"-"`

	// Subnet in which to allocate IP address for this vip.
	// Changing this creates a new vip.
	// +crossplane:generate:reference:type=github.com/opentelekomcloud/provider-opentelekomcloud/apis/networking/v1alpha1.SubnetV2
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a SubnetV2 in networking to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a SubnetV2 in networking to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`
}

// VipV2Spec defines the desired state of VipV2
type VipV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VipV2Parameters `json:"forProvider"`
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
	InitProvider VipV2InitParameters `json:"initProvider,omitempty"`
}

// VipV2Status defines the observed state of VipV2.
type VipV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VipV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// VipV2 is the Schema for the VipV2s API. Manages a VPC VIP resource within OpenTelekomCloud.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,opentelekomcloud}
type VipV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VipV2Spec   `json:"spec"`
	Status            VipV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VipV2List contains a list of VipV2s
type VipV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VipV2 `json:"items"`
}

// Repository type metadata.
var (
	VipV2_Kind             = "VipV2"
	VipV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VipV2_Kind}.String()
	VipV2_KindAPIVersion   = VipV2_Kind + "." + CRDGroupVersion.String()
	VipV2_GroupVersionKind = CRDGroupVersion.WithKind(VipV2_Kind)
)

func init() {
	SchemeBuilder.Register(&VipV2{}, &VipV2List{})
}
