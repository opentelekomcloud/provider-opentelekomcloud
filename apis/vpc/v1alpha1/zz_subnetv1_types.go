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

type SubnetV1InitParameters struct {

	// Identifies the availability zone (AZ) to which the subnet belongs.
	// The value must be an existing AZ in the system. Changing this creates a new Subnet.
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// Specifies the network segment on which the subnet resides. The value must be in CIDR format.
	// The value must be within the CIDR block of the VPC. The subnet mask cannot be greater than 28.
	// Changing this creates a new Subnet.
	Cidr *string `json:"cidr,omitempty" tf:"cidr,omitempty"`

	// Specifies whether the DHCP function is enabled for the subnet. The value can
	// be true or false. If this parameter is left blank, it is set to true by default.
	DHCPEnable *bool `json:"dhcpEnable,omitempty" tf:"dhcp_enable,omitempty"`

	// Specifies the DNS server address list of a subnet. This field is required if you
	// need to use more than two DNS servers. This parameter value is the superset of both DNS server address
	// 1 and DNS server address 2.
	DNSList []*string `json:"dnsList,omitempty" tf:"dns_list,omitempty"`

	// A description of the VPC subnet.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies the gateway of the subnet. The value must be a valid IP address.
	// The value must be an IP address in the subnet segment. Changing this creates a new Subnet.
	GatewayIP *string `json:"gatewayIp,omitempty" tf:"gateway_ip,omitempty"`

	// Specifies whether IPv6 is enabled. If IPv6 is enabled, you can use IPv6 CIDR blocks. The value can
	// be true or false. If this parameter is left blank, it is set to false by default.
	IPv6Enable *bool `json:"ipv6Enable,omitempty" tf:"ipv6_enable,omitempty"`

	// The subnet name. The value is a string of 1 to 64 characters that can contain letters,
	// digits, underscores (_), and hyphens (-).
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the NTP server address configured for the subnet.
	NtpAddresses *string `json:"ntpAddresses,omitempty" tf:"ntp_addresses,omitempty"`

	// Specifies the IP address of DNS server 1 on the subnet. The value must be a
	// valid IP address. Default is 100.125.4.25, OpenTelekomCloud internal DNS server.
	PrimaryDNS *string `json:"primaryDns,omitempty" tf:"primary_dns,omitempty"`

	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the IP address of DNS server 2 on the subnet. The value must be a
	// valid IP address. Default is 100.125.129.199, OpenTelekomCloud secondary internal DNS server.
	SecondaryDNS *string `json:"secondaryDns,omitempty" tf:"secondary_dns,omitempty"`

	// The key/value pairs to associate with the subnet.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the ID of the VPC to which the subnet belongs. Changing this creates a new Subnet.
	// +crossplane:generate:reference:type=github.com/opentelekomcloud/provider-opentelekomcloud/apis/vpc/v1alpha1.VpcV1
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VpcV1 in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VpcV1 in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`
}

type SubnetV1Observation struct {

	// Identifies the availability zone (AZ) to which the subnet belongs.
	// The value must be an existing AZ in the system. Changing this creates a new Subnet.
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// Specifies the network segment on which the subnet resides. The value must be in CIDR format.
	// The value must be within the CIDR block of the VPC. The subnet mask cannot be greater than 28.
	// Changing this creates a new Subnet.
	Cidr *string `json:"cidr,omitempty" tf:"cidr,omitempty"`

	CidrIPv6 *string `json:"cidrIpv6,omitempty" tf:"cidr_ipv6,omitempty"`

	// Specifies whether the DHCP function is enabled for the subnet. The value can
	// be true or false. If this parameter is left blank, it is set to true by default.
	DHCPEnable *bool `json:"dhcpEnable,omitempty" tf:"dhcp_enable,omitempty"`

	// Specifies the DNS server address list of a subnet. This field is required if you
	// need to use more than two DNS servers. This parameter value is the superset of both DNS server address
	// 1 and DNS server address 2.
	DNSList []*string `json:"dnsList,omitempty" tf:"dns_list,omitempty"`

	// A description of the VPC subnet.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies the gateway of the subnet. The value must be a valid IP address.
	// The value must be an IP address in the subnet segment. Changing this creates a new Subnet.
	GatewayIP *string `json:"gatewayIp,omitempty" tf:"gateway_ip,omitempty"`

	GatewayIPv6 *string `json:"gatewayIpv6,omitempty" tf:"gateway_ipv6,omitempty"`

	// Specifies a resource ID in UUID format. Same as OpenStack network ID (OS_NETWORK_ID).
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies whether IPv6 is enabled. If IPv6 is enabled, you can use IPv6 CIDR blocks. The value can
	// be true or false. If this parameter is left blank, it is set to false by default.
	IPv6Enable *bool `json:"ipv6Enable,omitempty" tf:"ipv6_enable,omitempty"`

	// The subnet name. The value is a string of 1 to 64 characters that can contain letters,
	// digits, underscores (_), and hyphens (-).
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the OpenStack network ID.
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// Specifies the NTP server address configured for the subnet.
	NtpAddresses *string `json:"ntpAddresses,omitempty" tf:"ntp_addresses,omitempty"`

	// Specifies the IP address of DNS server 1 on the subnet. The value must be a
	// valid IP address. Default is 100.125.4.25, OpenTelekomCloud internal DNS server.
	PrimaryDNS *string `json:"primaryDns,omitempty" tf:"primary_dns,omitempty"`

	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the IP address of DNS server 2 on the subnet. The value must be a
	// valid IP address. Default is 100.125.129.199, OpenTelekomCloud secondary internal DNS server.
	SecondaryDNS *string `json:"secondaryDns,omitempty" tf:"secondary_dns,omitempty"`

	// Specifies the status of the subnet. The value can be ACTIVE, DOWN, UNKNOWN, or ERROR.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Specifies the OpenStack subnet ID.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// The key/value pairs to associate with the subnet.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the ID of the VPC to which the subnet belongs. Changing this creates a new Subnet.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

type SubnetV1Parameters struct {

	// Identifies the availability zone (AZ) to which the subnet belongs.
	// The value must be an existing AZ in the system. Changing this creates a new Subnet.
	// +kubebuilder:validation:Optional
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// Specifies the network segment on which the subnet resides. The value must be in CIDR format.
	// The value must be within the CIDR block of the VPC. The subnet mask cannot be greater than 28.
	// Changing this creates a new Subnet.
	// +kubebuilder:validation:Optional
	Cidr *string `json:"cidr,omitempty" tf:"cidr,omitempty"`

	// Specifies whether the DHCP function is enabled for the subnet. The value can
	// be true or false. If this parameter is left blank, it is set to true by default.
	// +kubebuilder:validation:Optional
	DHCPEnable *bool `json:"dhcpEnable,omitempty" tf:"dhcp_enable,omitempty"`

	// Specifies the DNS server address list of a subnet. This field is required if you
	// need to use more than two DNS servers. This parameter value is the superset of both DNS server address
	// 1 and DNS server address 2.
	// +kubebuilder:validation:Optional
	DNSList []*string `json:"dnsList,omitempty" tf:"dns_list,omitempty"`

	// A description of the VPC subnet.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies the gateway of the subnet. The value must be a valid IP address.
	// The value must be an IP address in the subnet segment. Changing this creates a new Subnet.
	// +kubebuilder:validation:Optional
	GatewayIP *string `json:"gatewayIp,omitempty" tf:"gateway_ip,omitempty"`

	// Specifies whether IPv6 is enabled. If IPv6 is enabled, you can use IPv6 CIDR blocks. The value can
	// be true or false. If this parameter is left blank, it is set to false by default.
	// +kubebuilder:validation:Optional
	IPv6Enable *bool `json:"ipv6Enable,omitempty" tf:"ipv6_enable,omitempty"`

	// The subnet name. The value is a string of 1 to 64 characters that can contain letters,
	// digits, underscores (_), and hyphens (-).
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the NTP server address configured for the subnet.
	// +kubebuilder:validation:Optional
	NtpAddresses *string `json:"ntpAddresses,omitempty" tf:"ntp_addresses,omitempty"`

	// Specifies the IP address of DNS server 1 on the subnet. The value must be a
	// valid IP address. Default is 100.125.4.25, OpenTelekomCloud internal DNS server.
	// +kubebuilder:validation:Optional
	PrimaryDNS *string `json:"primaryDns,omitempty" tf:"primary_dns,omitempty"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the IP address of DNS server 2 on the subnet. The value must be a
	// valid IP address. Default is 100.125.129.199, OpenTelekomCloud secondary internal DNS server.
	// +kubebuilder:validation:Optional
	SecondaryDNS *string `json:"secondaryDns,omitempty" tf:"secondary_dns,omitempty"`

	// The key/value pairs to associate with the subnet.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the ID of the VPC to which the subnet belongs. Changing this creates a new Subnet.
	// +crossplane:generate:reference:type=github.com/opentelekomcloud/provider-opentelekomcloud/apis/vpc/v1alpha1.VpcV1
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VpcV1 in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VpcV1 in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`
}

// SubnetV1Spec defines the desired state of SubnetV1
type SubnetV1Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SubnetV1Parameters `json:"forProvider"`
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
	InitProvider SubnetV1InitParameters `json:"initProvider,omitempty"`
}

// SubnetV1Status defines the observed state of SubnetV1.
type SubnetV1Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SubnetV1Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SubnetV1 is the Schema for the SubnetV1s API. Manages a VPC Subnet resource within OpenTelekomCloud.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,opentelekomcloud}
type SubnetV1 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.cidr) || (has(self.initProvider) && has(self.initProvider.cidr))",message="spec.forProvider.cidr is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.gatewayIp) || (has(self.initProvider) && has(self.initProvider.gatewayIp))",message="spec.forProvider.gatewayIp is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   SubnetV1Spec   `json:"spec"`
	Status SubnetV1Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SubnetV1List contains a list of SubnetV1s
type SubnetV1List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SubnetV1 `json:"items"`
}

// Repository type metadata.
var (
	SubnetV1_Kind             = "SubnetV1"
	SubnetV1_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SubnetV1_Kind}.String()
	SubnetV1_KindAPIVersion   = SubnetV1_Kind + "." + CRDGroupVersion.String()
	SubnetV1_GroupVersionKind = CRDGroupVersion.WithKind(SubnetV1_Kind)
)

func init() {
	SchemeBuilder.Register(&SubnetV1{}, &SubnetV1List{})
}
