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

type DataDisksInitParameters struct {

	// The Encryption KMS ID of the data disk. Changing this
	// creates a new server.
	KMSID *string `json:"kmsId,omitempty" tf:"kms_id,omitempty"`

	// The size of the data disk in GB. The value range is 10 to 32768.
	// Changing this creates a new server.
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// Specifies the snapshot ID or ID of the original data disk contained in the full-ECS image.
	// Changing this creates a new server.
	SnapshotID *string `json:"snapshotId,omitempty" tf:"snapshot_id,omitempty"`

	// The data disk type of the server. For HANA, HL1, and HL2 ECSs use co-p1 and uh-l1 disks.
	// Changing this creates a new server. Options are limited depending on AZ. Available options are:
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type DataDisksObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Encryption KMS ID of the data disk. Changing this
	// creates a new server.
	KMSID *string `json:"kmsId,omitempty" tf:"kms_id,omitempty"`

	// The size of the data disk in GB. The value range is 10 to 32768.
	// Changing this creates a new server.
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// Specifies the snapshot ID or ID of the original data disk contained in the full-ECS image.
	// Changing this creates a new server.
	SnapshotID *string `json:"snapshotId,omitempty" tf:"snapshot_id,omitempty"`

	// The data disk type of the server. For HANA, HL1, and HL2 ECSs use co-p1 and uh-l1 disks.
	// Changing this creates a new server. Options are limited depending on AZ. Available options are:
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type DataDisksParameters struct {

	// The Encryption KMS ID of the data disk. Changing this
	// creates a new server.
	// +kubebuilder:validation:Optional
	KMSID *string `json:"kmsId,omitempty" tf:"kms_id,omitempty"`

	// The size of the data disk in GB. The value range is 10 to 32768.
	// Changing this creates a new server.
	// +kubebuilder:validation:Optional
	Size *float64 `json:"size" tf:"size,omitempty"`

	// Specifies the snapshot ID or ID of the original data disk contained in the full-ECS image.
	// Changing this creates a new server.
	// +kubebuilder:validation:Optional
	SnapshotID *string `json:"snapshotId,omitempty" tf:"snapshot_id,omitempty"`

	// The data disk type of the server. For HANA, HL1, and HL2 ECSs use co-p1 and uh-l1 disks.
	// Changing this creates a new server. Options are limited depending on AZ. Available options are:
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type InstanceV1InitParameters struct {

	// Whether configure automatic recovery of an instance.
	AutoRecovery *bool `json:"autoRecovery,omitempty" tf:"auto_recovery,omitempty"`

	// The availability zone in which to create the server.
	// Changing this creates a new server.
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// An array of one or more data disks to attach to the
	// instance. The data_disks object structure is documented below. Changing this
	// creates a new server.
	DataDisks []DataDisksInitParameters `json:"dataDisks,omitempty" tf:"data_disks,omitempty"`

	// Delete the data disks upon termination of the instance.
	// Defaults to false. Changing this creates a new server.
	DeleteDisksOnTermination *bool `json:"deleteDisksOnTermination,omitempty" tf:"delete_disks_on_termination,omitempty"`

	// The name of the desired flavor for the server.
	Flavor *string `json:"flavor,omitempty" tf:"flavor,omitempty"`

	// The ID of the desired image for the server. Changing this creates a new server.
	ImageID *string `json:"imageId,omitempty" tf:"image_id,omitempty"`

	// A unique name for the instance.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// An array of one or more networks to attach to the
	// instance. The nics object structure is documented below. Changing this
	// creates a new server.
	Nics []NicsInitParameters `json:"nics,omitempty" tf:"nics,omitempty"`

	// The Encryption KMS ID of the system disk. Changing this
	// creates a new server.
	SystemDiskKMSID *string `json:"systemDiskKmsId,omitempty" tf:"system_disk_kms_id,omitempty"`

	// The system disk size in GB, The value range is 1 to 1024.
	// Changing this creates a new server.
	SystemDiskSize *float64 `json:"systemDiskSize,omitempty" tf:"system_disk_size,omitempty"`

	// The system disk type of the server. For HANA, HL1, and HL2 ECSs use co-p1 and uh-l1 disks.
	// Changing this creates a new server. Options are limited depending on AZ. Available options are:
	SystemDiskType *string `json:"systemDiskType,omitempty" tf:"system_disk_type,omitempty"`

	// Tags key/value pairs to associate with the instance.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The user data to provide when launching the instance.
	// Changing this creates a new server.
	UserData *string `json:"userData,omitempty" tf:"user_data,omitempty"`
}

type InstanceV1Observation struct {

	// Whether configure automatic recovery of an instance.
	AutoRecovery *bool `json:"autoRecovery,omitempty" tf:"auto_recovery,omitempty"`

	// The availability zone in which to create the server.
	// Changing this creates a new server.
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// An array of one or more data disks to attach to the
	// instance. The data_disks object structure is documented below. Changing this
	// creates a new server.
	DataDisks []DataDisksObservation `json:"dataDisks,omitempty" tf:"data_disks,omitempty"`

	// Delete the data disks upon termination of the instance.
	// Defaults to false. Changing this creates a new server.
	DeleteDisksOnTermination *bool `json:"deleteDisksOnTermination,omitempty" tf:"delete_disks_on_termination,omitempty"`

	// The name of the desired flavor for the server.
	Flavor *string `json:"flavor,omitempty" tf:"flavor,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the desired image for the server. Changing this creates a new server.
	ImageID *string `json:"imageId,omitempty" tf:"image_id,omitempty"`

	// The name of a key pair to put on the server. The key
	// pair must already be created and associated with the tenant's account.
	// Changing this creates a new server.
	KeyName *string `json:"keyName,omitempty" tf:"key_name,omitempty"`

	// A unique name for the instance.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// An array of one or more networks to attach to the
	// instance. The nics object structure is documented below. Changing this
	// creates a new server.
	Nics []NicsObservation `json:"nics,omitempty" tf:"nics,omitempty"`

	// An array of one or more security group IDs
	// to associate with the server. If this parameter is left blank, the default
	// security group is bound to the ECS by default.
	SecurityGroups []*string `json:"securityGroups,omitempty" tf:"security_groups,omitempty"`

	// The ID of the system disk.
	SystemDiskID *string `json:"systemDiskId,omitempty" tf:"system_disk_id,omitempty"`

	// The Encryption KMS ID of the system disk. Changing this
	// creates a new server.
	SystemDiskKMSID *string `json:"systemDiskKmsId,omitempty" tf:"system_disk_kms_id,omitempty"`

	// The system disk size in GB, The value range is 1 to 1024.
	// Changing this creates a new server.
	SystemDiskSize *float64 `json:"systemDiskSize,omitempty" tf:"system_disk_size,omitempty"`

	// The system disk type of the server. For HANA, HL1, and HL2 ECSs use co-p1 and uh-l1 disks.
	// Changing this creates a new server. Options are limited depending on AZ. Available options are:
	SystemDiskType *string `json:"systemDiskType,omitempty" tf:"system_disk_type,omitempty"`

	// Tags key/value pairs to associate with the instance.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The user data to provide when launching the instance.
	// Changing this creates a new server.
	UserData *string `json:"userData,omitempty" tf:"user_data,omitempty"`

	// The ID of the desired VPC for the server. Changing this creates a new server.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	VolumesAttached []VolumesAttachedObservation `json:"volumesAttached,omitempty" tf:"volumes_attached,omitempty"`
}

type InstanceV1Parameters struct {

	// Whether configure automatic recovery of an instance.
	// +kubebuilder:validation:Optional
	AutoRecovery *bool `json:"autoRecovery,omitempty" tf:"auto_recovery,omitempty"`

	// The availability zone in which to create the server.
	// Changing this creates a new server.
	// +kubebuilder:validation:Optional
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// References to SecgroupV2 in compute to populate securityGroups.
	// +kubebuilder:validation:Optional
	ComputeSecurityGroupIDRefs []v1.Reference `json:"computeSecurityGroupIdRefs,omitempty" tf:"-"`

	// Selector for a list of SecgroupV2 in compute to populate securityGroups.
	// +kubebuilder:validation:Optional
	ComputeSecurityGroupIDSelector *v1.Selector `json:"computeSecurityGroupIdSelector,omitempty" tf:"-"`

	// An array of one or more data disks to attach to the
	// instance. The data_disks object structure is documented below. Changing this
	// creates a new server.
	// +kubebuilder:validation:Optional
	DataDisks []DataDisksParameters `json:"dataDisks,omitempty" tf:"data_disks,omitempty"`

	// Delete the data disks upon termination of the instance.
	// Defaults to false. Changing this creates a new server.
	// +kubebuilder:validation:Optional
	DeleteDisksOnTermination *bool `json:"deleteDisksOnTermination,omitempty" tf:"delete_disks_on_termination,omitempty"`

	// The name of the desired flavor for the server.
	// +kubebuilder:validation:Optional
	Flavor *string `json:"flavor,omitempty" tf:"flavor,omitempty"`

	// The ID of the desired image for the server. Changing this creates a new server.
	// +kubebuilder:validation:Optional
	ImageID *string `json:"imageId,omitempty" tf:"image_id,omitempty"`

	// The name of a key pair to put on the server. The key
	// pair must already be created and associated with the tenant's account.
	// Changing this creates a new server.
	// +crossplane:generate:reference:type=github.com/opentelekomcloud/provider-opentelekomcloud/apis/compute/v1alpha1.KeypairV2
	// +kubebuilder:validation:Optional
	KeyName *string `json:"keyName,omitempty" tf:"key_name,omitempty"`

	// Reference to a KeypairV2 in compute to populate keyName.
	// +kubebuilder:validation:Optional
	KeyNameRef *v1.Reference `json:"keyNameRef,omitempty" tf:"-"`

	// Selector for a KeypairV2 in compute to populate keyName.
	// +kubebuilder:validation:Optional
	KeyNameSelector *v1.Selector `json:"keyNameSelector,omitempty" tf:"-"`

	// A unique name for the instance.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// An array of one or more networks to attach to the
	// instance. The nics object structure is documented below. Changing this
	// creates a new server.
	// +kubebuilder:validation:Optional
	Nics []NicsParameters `json:"nics,omitempty" tf:"nics,omitempty"`

	// The administrative password to assign to the server.
	// Changing this creates a new server.
	// +kubebuilder:validation:Optional
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// An array of one or more security group IDs
	// to associate with the server. If this parameter is left blank, the default
	// security group is bound to the ECS by default.
	// +crossplane:generate:reference:type=github.com/opentelekomcloud/provider-opentelekomcloud/apis/compute/v1alpha1.SecgroupV2
	// +crossplane:generate:reference:refFieldName=ComputeSecurityGroupIDRefs
	// +crossplane:generate:reference:selectorFieldName=ComputeSecurityGroupIDSelector
	// +kubebuilder:validation:Optional
	SecurityGroups []*string `json:"securityGroups,omitempty" tf:"security_groups,omitempty"`

	// The Encryption KMS ID of the system disk. Changing this
	// creates a new server.
	// +kubebuilder:validation:Optional
	SystemDiskKMSID *string `json:"systemDiskKmsId,omitempty" tf:"system_disk_kms_id,omitempty"`

	// The system disk size in GB, The value range is 1 to 1024.
	// Changing this creates a new server.
	// +kubebuilder:validation:Optional
	SystemDiskSize *float64 `json:"systemDiskSize,omitempty" tf:"system_disk_size,omitempty"`

	// The system disk type of the server. For HANA, HL1, and HL2 ECSs use co-p1 and uh-l1 disks.
	// Changing this creates a new server. Options are limited depending on AZ. Available options are:
	// +kubebuilder:validation:Optional
	SystemDiskType *string `json:"systemDiskType,omitempty" tf:"system_disk_type,omitempty"`

	// Tags key/value pairs to associate with the instance.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The user data to provide when launching the instance.
	// Changing this creates a new server.
	// +kubebuilder:validation:Optional
	UserData *string `json:"userData,omitempty" tf:"user_data,omitempty"`

	// The ID of the desired VPC for the server. Changing this creates a new server.
	// +crossplane:generate:reference:type=github.com/opentelekomcloud/provider-opentelekomcloud/apis/vpc/v1alpha1.V1
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a V1 in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a V1 in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`
}

type NicsInitParameters struct {

	// Specifies a fixed IPv4 address to be used on this
	// network. Changing this creates a new server.
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`
}

type NicsObservation struct {

	// Specifies a fixed IPv4 address to be used on this
	// network. Changing this creates a new server.
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	MacAddress *string `json:"macAddress,omitempty" tf:"mac_address,omitempty"`

	// The network UUID to attach to the server. Changing this creates a new server.
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	PortID *string `json:"portId,omitempty" tf:"port_id,omitempty"`

	// The data disk type of the server. For HANA, HL1, and HL2 ECSs use co-p1 and uh-l1 disks.
	// Changing this creates a new server. Options are limited depending on AZ. Available options are:
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type NicsParameters struct {

	// Specifies a fixed IPv4 address to be used on this
	// network. Changing this creates a new server.
	// +kubebuilder:validation:Optional
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// The network UUID to attach to the server. Changing this creates a new server.
	// +crossplane:generate:reference:type=github.com/opentelekomcloud/provider-opentelekomcloud/apis/vpc/v1alpha1.SubnetV1
	// +crossplane:generate:reference:extractor=github.com/opentelekomcloud/provider-opentelekomcloud/config/compute.ExtractNetworkId()
	// +kubebuilder:validation:Optional
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// Reference to a SubnetV1 in vpc to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDRef *v1.Reference `json:"networkIdRef,omitempty" tf:"-"`

	// Selector for a SubnetV1 in vpc to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDSelector *v1.Selector `json:"networkIdSelector,omitempty" tf:"-"`
}

type VolumesAttachedInitParameters struct {
}

type VolumesAttachedObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Encryption KMS ID of the data disk. Changing this
	// creates a new server.
	KMSID *string `json:"kmsId,omitempty" tf:"kms_id,omitempty"`

	// The size of the data disk in GB. The value range is 10 to 32768.
	// Changing this creates a new server.
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// Specifies the snapshot ID or ID of the original data disk contained in the full-ECS image.
	// Changing this creates a new server.
	SnapshotID *string `json:"snapshotId,omitempty" tf:"snapshot_id,omitempty"`

	// The data disk type of the server. For HANA, HL1, and HL2 ECSs use co-p1 and uh-l1 disks.
	// Changing this creates a new server. Options are limited depending on AZ. Available options are:
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type VolumesAttachedParameters struct {
}

// InstanceV1Spec defines the desired state of InstanceV1
type InstanceV1Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InstanceV1Parameters `json:"forProvider"`
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
	InitProvider InstanceV1InitParameters `json:"initProvider,omitempty"`
}

// InstanceV1Status defines the observed state of InstanceV1.
type InstanceV1Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InstanceV1Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceV1 is the Schema for the InstanceV1s API. Manages a ECS Instance resource within OpenTelekomCloud.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,opentelekomcloud}
type InstanceV1 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.availabilityZone) || (has(self.initProvider) && has(self.initProvider.availabilityZone))",message="spec.forProvider.availabilityZone is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.flavor) || (has(self.initProvider) && has(self.initProvider.flavor))",message="spec.forProvider.flavor is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.imageId) || (has(self.initProvider) && has(self.initProvider.imageId))",message="spec.forProvider.imageId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.nics) || (has(self.initProvider) && has(self.initProvider.nics))",message="spec.forProvider.nics is a required parameter"
	Spec   InstanceV1Spec   `json:"spec"`
	Status InstanceV1Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceV1List contains a list of InstanceV1s
type InstanceV1List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InstanceV1 `json:"items"`
}

// Repository type metadata.
var (
	InstanceV1_Kind             = "InstanceV1"
	InstanceV1_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: InstanceV1_Kind}.String()
	InstanceV1_KindAPIVersion   = InstanceV1_Kind + "." + CRDGroupVersion.String()
	InstanceV1_GroupVersionKind = CRDGroupVersion.WithKind(InstanceV1_Kind)
)

func init() {
	SchemeBuilder.Register(&InstanceV1{}, &InstanceV1List{})
}
