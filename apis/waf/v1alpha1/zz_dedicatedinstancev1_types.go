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

type DedicatedInstanceV1InitParameters struct {

	// Dedicated engine CPU architecture. Default value is x86.
	// Changing this will create a new instance.
	Architecture *string `json:"architecture,omitempty" tf:"architecture,omitempty"`

	// AZ where the dedicated engine is to be created. Changing this will create a new instance.
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// ID of the specifications of the ECS hosting the dedicated engine.
	// You can go to the management console and confirm supported specifications. Changing this will create a new instance.
	Flavor *string `json:"flavor,omitempty" tf:"flavor,omitempty"`

	// The name of WAF dedicated instance. Duplicate names are allowed, we suggest to keeping the
	// name unique.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region where a dedicated engine is to be created. If omitted, the
	// provider-level region will be used. Changing this setting will create a new instance.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Whether the dedicated WAF instance is network interface type.
	// Default value is true. Changing this will create a new instance.
	// -> Note: This type of instance is not available in eu-ch2 region, must set res_tenant value to false there.
	ResTenant *bool `json:"resTenant,omitempty" tf:"res_tenant,omitempty"`

	// ID of the security group where the dedicated engine is located.
	// Changing this will create a new instance.
	SecurityGroup []*string `json:"securityGroup,omitempty" tf:"security_group,omitempty"`

	// Specifications of the dedicated engine version. Values are:
	Specification *string `json:"specification,omitempty" tf:"specification,omitempty"`

	// ID of the VPC subnet where the dedicated engine is located.
	// Subnet_id has the same value as network_id obtained by calling the OpenStack APIs. Changing this will create a
	// new instance.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// ID of the VPC where the dedicated engine is located. Changing this will create a new
	// instance.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

type DedicatedInstanceV1Observation struct {

	// The access status of the instance.
	AccessStatus *float64 `json:"accessStatus,omitempty" tf:"access_status,omitempty"`

	// Dedicated engine CPU architecture. Default value is x86.
	// Changing this will create a new instance.
	Architecture *string `json:"architecture,omitempty" tf:"architecture,omitempty"`

	// AZ where the dedicated engine is to be created. Changing this will create a new instance.
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// Billing status of dedicated WAF engine. The value can be 0, 1, or 2.
	BillingStatus *float64 `json:"billingStatus,omitempty" tf:"billing_status,omitempty"`

	// Timestamp when the dedicated WAF engine was created.
	CreatedAt *float64 `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// ID of the specifications of the ECS hosting the dedicated engine.
	// You can go to the management console and confirm supported specifications. Changing this will create a new instance.
	Flavor *string `json:"flavor,omitempty" tf:"flavor,omitempty"`

	// The id of the instance.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of WAF dedicated instance. Duplicate names are allowed, we suggest to keeping the
	// name unique.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region where a dedicated engine is to be created. If omitted, the
	// provider-level region will be used. Changing this setting will create a new instance.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Whether the dedicated WAF instance is network interface type.
	// Default value is true. Changing this will create a new instance.
	// -> Note: This type of instance is not available in eu-ch2 region, must set res_tenant value to false there.
	ResTenant *bool `json:"resTenant,omitempty" tf:"res_tenant,omitempty"`

	// ID of the security group where the dedicated engine is located.
	// Changing this will create a new instance.
	SecurityGroup []*string `json:"securityGroup,omitempty" tf:"security_group,omitempty"`

	// The id of the instance server.
	ServerID *string `json:"serverId,omitempty" tf:"server_id,omitempty"`

	// The ip of the instance service.
	ServiceIP *string `json:"serviceIp,omitempty" tf:"service_ip,omitempty"`

	// Specifications of the dedicated engine version. Values are:
	Specification *string `json:"specification,omitempty" tf:"specification,omitempty"`

	// Running status of the dedicated engine.
	// The value can be:
	Status *float64 `json:"status,omitempty" tf:"status,omitempty"`

	// ID of the VPC subnet where the dedicated engine is located.
	// Subnet_id has the same value as network_id obtained by calling the OpenStack APIs. Changing this will create a
	// new instance.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// The instance is to support upgrades. false: Cannot be upgraded, true: Can be upgraded.
	Upgradable *bool `json:"upgradable,omitempty" tf:"upgradable,omitempty"`

	// ID of the VPC where the dedicated engine is located. Changing this will create a new
	// instance.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

type DedicatedInstanceV1Parameters struct {

	// Dedicated engine CPU architecture. Default value is x86.
	// Changing this will create a new instance.
	// +kubebuilder:validation:Optional
	Architecture *string `json:"architecture,omitempty" tf:"architecture,omitempty"`

	// AZ where the dedicated engine is to be created. Changing this will create a new instance.
	// +kubebuilder:validation:Optional
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// ID of the specifications of the ECS hosting the dedicated engine.
	// You can go to the management console and confirm supported specifications. Changing this will create a new instance.
	// +kubebuilder:validation:Optional
	Flavor *string `json:"flavor,omitempty" tf:"flavor,omitempty"`

	// The name of WAF dedicated instance. Duplicate names are allowed, we suggest to keeping the
	// name unique.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region where a dedicated engine is to be created. If omitted, the
	// provider-level region will be used. Changing this setting will create a new instance.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Whether the dedicated WAF instance is network interface type.
	// Default value is true. Changing this will create a new instance.
	// -> Note: This type of instance is not available in eu-ch2 region, must set res_tenant value to false there.
	// +kubebuilder:validation:Optional
	ResTenant *bool `json:"resTenant,omitempty" tf:"res_tenant,omitempty"`

	// ID of the security group where the dedicated engine is located.
	// Changing this will create a new instance.
	// +kubebuilder:validation:Optional
	SecurityGroup []*string `json:"securityGroup,omitempty" tf:"security_group,omitempty"`

	// Specifications of the dedicated engine version. Values are:
	// +kubebuilder:validation:Optional
	Specification *string `json:"specification,omitempty" tf:"specification,omitempty"`

	// ID of the VPC subnet where the dedicated engine is located.
	// Subnet_id has the same value as network_id obtained by calling the OpenStack APIs. Changing this will create a
	// new instance.
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// ID of the VPC where the dedicated engine is located. Changing this will create a new
	// instance.
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

// DedicatedInstanceV1Spec defines the desired state of DedicatedInstanceV1
type DedicatedInstanceV1Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DedicatedInstanceV1Parameters `json:"forProvider"`
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
	InitProvider DedicatedInstanceV1InitParameters `json:"initProvider,omitempty"`
}

// DedicatedInstanceV1Status defines the observed state of DedicatedInstanceV1.
type DedicatedInstanceV1Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DedicatedInstanceV1Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// DedicatedInstanceV1 is the Schema for the DedicatedInstanceV1s API. Manages a WAF Dedicated Instance resource within OpenTelekomCloud.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,opentelekomcloud}
type DedicatedInstanceV1 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.availabilityZone) || (has(self.initProvider) && has(self.initProvider.availabilityZone))",message="spec.forProvider.availabilityZone is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.flavor) || (has(self.initProvider) && has(self.initProvider.flavor))",message="spec.forProvider.flavor is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.securityGroup) || (has(self.initProvider) && has(self.initProvider.securityGroup))",message="spec.forProvider.securityGroup is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.specification) || (has(self.initProvider) && has(self.initProvider.specification))",message="spec.forProvider.specification is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.subnetId) || (has(self.initProvider) && has(self.initProvider.subnetId))",message="spec.forProvider.subnetId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.vpcId) || (has(self.initProvider) && has(self.initProvider.vpcId))",message="spec.forProvider.vpcId is a required parameter"
	Spec   DedicatedInstanceV1Spec   `json:"spec"`
	Status DedicatedInstanceV1Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DedicatedInstanceV1List contains a list of DedicatedInstanceV1s
type DedicatedInstanceV1List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DedicatedInstanceV1 `json:"items"`
}

// Repository type metadata.
var (
	DedicatedInstanceV1_Kind             = "DedicatedInstanceV1"
	DedicatedInstanceV1_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DedicatedInstanceV1_Kind}.String()
	DedicatedInstanceV1_KindAPIVersion   = DedicatedInstanceV1_Kind + "." + CRDGroupVersion.String()
	DedicatedInstanceV1_GroupVersionKind = CRDGroupVersion.WithKind(DedicatedInstanceV1_Kind)
)

func init() {
	SchemeBuilder.Register(&DedicatedInstanceV1{}, &DedicatedInstanceV1List{})
}
