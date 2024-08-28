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

type LoadbalancerV3InitParameters struct {

	// The administrative state of the LoadBalancer. A valid value is only true (UP).
	AdminStateUp *bool `json:"adminStateUp,omitempty" tf:"admin_state_up,omitempty"`

	// Specifies the availability zones where the LoadBalancer will be located.
	// Changing this creates a new LoadBalancer.
	// +listType=set
	AvailabilityZones []*string `json:"availabilityZones,omitempty" tf:"availability_zones,omitempty"`

	// Specifies whether to enable deletion protection for the load balancer.
	// true: Enable deletion protection.
	// false (default): Disable deletion protection.
	DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

	// Provides supplementary information about the load balancer.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The value can be true (enabled) or false (disabled).
	IPTargetEnable *bool `json:"ipTargetEnable,omitempty" tf:"ip_target_enable,omitempty"`

	// The ID of the Layer-4 flavor.
	L4Flavor *string `json:"l4Flavor,omitempty" tf:"l4_flavor,omitempty"`

	// The ID of the Layer-7 flavor.
	L7Flavor *string `json:"l7Flavor,omitempty" tf:"l7_flavor,omitempty"`

	// The LoadBalancer name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the subnet Network ID.
	// +listType=set
	NetworkIds []*string `json:"networkIds,omitempty" tf:"network_ids,omitempty"`

	// The elastic IP address of the instance. The public_ip structure
	// is described below. Changing this creates a new LoadBalancer.
	PublicIP []PublicIPInitParameters `json:"publicIp,omitempty" tf:"public_ip,omitempty"`

	// ID of the router (or VPC) this LoadBalancer belongs to. Changing
	// this creates a new LoadBalancer.
	RouterID *string `json:"routerId,omitempty" tf:"router_id,omitempty"`

	// The ID of the subnet to which the LoadBalancer belongs. Required when using vip_address.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Tags key/value pairs to associate with the load balancer.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The ip address of the LoadBalancer. Changing this creates a new LoadBalancer.
	VipAddress *string `json:"vipAddress,omitempty" tf:"vip_address,omitempty"`
}

type LoadbalancerV3Observation struct {

	// The administrative state of the LoadBalancer. A valid value is only true (UP).
	AdminStateUp *bool `json:"adminStateUp,omitempty" tf:"admin_state_up,omitempty"`

	// Specifies the availability zones where the LoadBalancer will be located.
	// Changing this creates a new LoadBalancer.
	// +listType=set
	AvailabilityZones []*string `json:"availabilityZones,omitempty" tf:"availability_zones,omitempty"`

	// The time the LoadBalancer was created.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// Specifies whether to enable deletion protection for the load balancer.
	// true: Enable deletion protection.
	// false (default): Disable deletion protection.
	DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

	// Provides supplementary information about the load balancer.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// ID of an existing elastic IP. Required when using existing EIP.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The value can be true (enabled) or false (disabled).
	IPTargetEnable *bool `json:"ipTargetEnable,omitempty" tf:"ip_target_enable,omitempty"`

	// The ID of the Layer-4 flavor.
	L4Flavor *string `json:"l4Flavor,omitempty" tf:"l4_flavor,omitempty"`

	// The ID of the Layer-7 flavor.
	L7Flavor *string `json:"l7Flavor,omitempty" tf:"l7_flavor,omitempty"`

	// The LoadBalancer name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the subnet Network ID.
	// +listType=set
	NetworkIds []*string `json:"networkIds,omitempty" tf:"network_ids,omitempty"`

	// The elastic IP address of the instance. The public_ip structure
	// is described below. Changing this creates a new LoadBalancer.
	PublicIP []PublicIPObservation `json:"publicIp,omitempty" tf:"public_ip,omitempty"`

	// ID of the router (or VPC) this LoadBalancer belongs to. Changing
	// this creates a new LoadBalancer.
	RouterID *string `json:"routerId,omitempty" tf:"router_id,omitempty"`

	// The ID of the subnet to which the LoadBalancer belongs. Required when using vip_address.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Tags key/value pairs to associate with the load balancer.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The time the LoadBalancer was last updated.
	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`

	// The ip address of the LoadBalancer. Changing this creates a new LoadBalancer.
	VipAddress *string `json:"vipAddress,omitempty" tf:"vip_address,omitempty"`

	// The Port ID of the Load Balancer IP.
	VipPortID *string `json:"vipPortId,omitempty" tf:"vip_port_id,omitempty"`
}

type LoadbalancerV3Parameters struct {

	// The administrative state of the LoadBalancer. A valid value is only true (UP).
	// +kubebuilder:validation:Optional
	AdminStateUp *bool `json:"adminStateUp,omitempty" tf:"admin_state_up,omitempty"`

	// Specifies the availability zones where the LoadBalancer will be located.
	// Changing this creates a new LoadBalancer.
	// +kubebuilder:validation:Optional
	// +listType=set
	AvailabilityZones []*string `json:"availabilityZones,omitempty" tf:"availability_zones,omitempty"`

	// Specifies whether to enable deletion protection for the load balancer.
	// true: Enable deletion protection.
	// false (default): Disable deletion protection.
	// +kubebuilder:validation:Optional
	DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

	// Provides supplementary information about the load balancer.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The value can be true (enabled) or false (disabled).
	// +kubebuilder:validation:Optional
	IPTargetEnable *bool `json:"ipTargetEnable,omitempty" tf:"ip_target_enable,omitempty"`

	// The ID of the Layer-4 flavor.
	// +kubebuilder:validation:Optional
	L4Flavor *string `json:"l4Flavor,omitempty" tf:"l4_flavor,omitempty"`

	// The ID of the Layer-7 flavor.
	// +kubebuilder:validation:Optional
	L7Flavor *string `json:"l7Flavor,omitempty" tf:"l7_flavor,omitempty"`

	// The LoadBalancer name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the subnet Network ID.
	// +kubebuilder:validation:Optional
	// +listType=set
	NetworkIds []*string `json:"networkIds,omitempty" tf:"network_ids,omitempty"`

	// The elastic IP address of the instance. The public_ip structure
	// is described below. Changing this creates a new LoadBalancer.
	// +kubebuilder:validation:Optional
	PublicIP []PublicIPParameters `json:"publicIp,omitempty" tf:"public_ip,omitempty"`

	// ID of the router (or VPC) this LoadBalancer belongs to. Changing
	// this creates a new LoadBalancer.
	// +kubebuilder:validation:Optional
	RouterID *string `json:"routerId,omitempty" tf:"router_id,omitempty"`

	// The ID of the subnet to which the LoadBalancer belongs. Required when using vip_address.
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Tags key/value pairs to associate with the load balancer.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The ip address of the LoadBalancer. Changing this creates a new LoadBalancer.
	// +kubebuilder:validation:Optional
	VipAddress *string `json:"vipAddress,omitempty" tf:"vip_address,omitempty"`
}

type PublicIPInitParameters struct {

	// Bandwidth billing type. Possible value is traffic.
	BandwidthChargeMode *string `json:"bandwidthChargeMode,omitempty" tf:"bandwidth_charge_mode,omitempty"`

	// Bandwidth name. Required when creating a new EIP.
	BandwidthName *string `json:"bandwidthName,omitempty" tf:"bandwidth_name,omitempty"`

	// Bandwidth sharing type. Possible values are: PER, WHOLE.
	// Required when creating a new EIP.
	BandwidthShareType *string `json:"bandwidthShareType,omitempty" tf:"bandwidth_share_type,omitempty"`

	// Bandwidth size. Required when creating a new EIP.
	BandwidthSize *float64 `json:"bandwidthSize,omitempty" tf:"bandwidth_size,omitempty"`

	// ID of an existing elastic IP. Required when using existing EIP.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Elastic IP type. The value can be 5_bgp or 5_mailbgp.
	// Required when creating a new EIP.
	IPType *string `json:"ipType,omitempty" tf:"ip_type,omitempty"`
}

type PublicIPObservation struct {
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// Bandwidth billing type. Possible value is traffic.
	BandwidthChargeMode *string `json:"bandwidthChargeMode,omitempty" tf:"bandwidth_charge_mode,omitempty"`

	// Bandwidth name. Required when creating a new EIP.
	BandwidthName *string `json:"bandwidthName,omitempty" tf:"bandwidth_name,omitempty"`

	// Bandwidth sharing type. Possible values are: PER, WHOLE.
	// Required when creating a new EIP.
	BandwidthShareType *string `json:"bandwidthShareType,omitempty" tf:"bandwidth_share_type,omitempty"`

	// Bandwidth size. Required when creating a new EIP.
	BandwidthSize *float64 `json:"bandwidthSize,omitempty" tf:"bandwidth_size,omitempty"`

	// ID of an existing elastic IP. Required when using existing EIP.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Elastic IP type. The value can be 5_bgp or 5_mailbgp.
	// Required when creating a new EIP.
	IPType *string `json:"ipType,omitempty" tf:"ip_type,omitempty"`

	Managed *bool `json:"Managed,omitempty" tf:"_managed,omitempty"`
}

type PublicIPParameters struct {

	// Bandwidth billing type. Possible value is traffic.
	// +kubebuilder:validation:Optional
	BandwidthChargeMode *string `json:"bandwidthChargeMode,omitempty" tf:"bandwidth_charge_mode,omitempty"`

	// Bandwidth name. Required when creating a new EIP.
	// +kubebuilder:validation:Optional
	BandwidthName *string `json:"bandwidthName,omitempty" tf:"bandwidth_name,omitempty"`

	// Bandwidth sharing type. Possible values are: PER, WHOLE.
	// Required when creating a new EIP.
	// +kubebuilder:validation:Optional
	BandwidthShareType *string `json:"bandwidthShareType,omitempty" tf:"bandwidth_share_type,omitempty"`

	// Bandwidth size. Required when creating a new EIP.
	// +kubebuilder:validation:Optional
	BandwidthSize *float64 `json:"bandwidthSize,omitempty" tf:"bandwidth_size,omitempty"`

	// ID of an existing elastic IP. Required when using existing EIP.
	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Elastic IP type. The value can be 5_bgp or 5_mailbgp.
	// Required when creating a new EIP.
	// +kubebuilder:validation:Optional
	IPType *string `json:"ipType,omitempty" tf:"ip_type,omitempty"`
}

// LoadbalancerV3Spec defines the desired state of LoadbalancerV3
type LoadbalancerV3Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LoadbalancerV3Parameters `json:"forProvider"`
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
	InitProvider LoadbalancerV3InitParameters `json:"initProvider,omitempty"`
}

// LoadbalancerV3Status defines the observed state of LoadbalancerV3.
type LoadbalancerV3Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LoadbalancerV3Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// LoadbalancerV3 is the Schema for the LoadbalancerV3s API. Manages a LB Loadbalancer resource within OpenTelekomCloud.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,opentelekomcloud}
type LoadbalancerV3 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.availabilityZones) || (has(self.initProvider) && has(self.initProvider.availabilityZones))",message="spec.forProvider.availabilityZones is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.networkIds) || (has(self.initProvider) && has(self.initProvider.networkIds))",message="spec.forProvider.networkIds is a required parameter"
	Spec   LoadbalancerV3Spec   `json:"spec"`
	Status LoadbalancerV3Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LoadbalancerV3List contains a list of LoadbalancerV3s
type LoadbalancerV3List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LoadbalancerV3 `json:"items"`
}

// Repository type metadata.
var (
	LoadbalancerV3_Kind             = "LoadbalancerV3"
	LoadbalancerV3_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LoadbalancerV3_Kind}.String()
	LoadbalancerV3_KindAPIVersion   = LoadbalancerV3_Kind + "." + CRDGroupVersion.String()
	LoadbalancerV3_GroupVersionKind = CRDGroupVersion.WithKind(LoadbalancerV3_Kind)
)

func init() {
	SchemeBuilder.Register(&LoadbalancerV3{}, &LoadbalancerV3List{})
}
