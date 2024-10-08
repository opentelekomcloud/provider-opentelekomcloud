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

type BlockDeviceInitParameters struct {

	// The boot index of the volume. It defaults to 0. Changing this creates a new server.
	BootIndex *float64 `json:"bootIndex,omitempty" tf:"boot_index,omitempty"`

	// Delete the volume / block device upon termination of the instance. Defaults to
	// false. Changing this creates a new server.
	DeleteOnTermination *bool `json:"deleteOnTermination,omitempty" tf:"delete_on_termination,omitempty"`

	// The type that gets created. Currently only support "volume". Changing this creates a
	// new server.
	DestinationType *string `json:"destinationType,omitempty" tf:"destination_type,omitempty"`

	// A unique name for the resource.
	DeviceName *string `json:"deviceName,omitempty" tf:"device_name,omitempty"`

	GuestFormat *string `json:"guestFormat,omitempty" tf:"guest_format,omitempty"`

	// The source type of the device. Must be one of
	// "blank", "image", "volume", or "snapshot". Changing this creates a new server.
	SourceType *string `json:"sourceType,omitempty" tf:"source_type,omitempty"`

	// The UUID of the image, volume, or snapshot. Changing
	// this creates a new server.
	// +crossplane:generate:reference:type=github.com/opentelekomcloud/provider-opentelekomcloud/apis/blockstorage/v1alpha1.VolumeV2
	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`

	// Reference to a VolumeV2 in blockstorage to populate uuid.
	// +kubebuilder:validation:Optional
	UUIDRef *v1.Reference `json:"uuidRef,omitempty" tf:"-"`

	// Selector for a VolumeV2 in blockstorage to populate uuid.
	// +kubebuilder:validation:Optional
	UUIDSelector *v1.Selector `json:"uuidSelector,omitempty" tf:"-"`

	// The size of the volume to create (in gigabytes). Required in the following combinations: source=image
	// and destination=volume, and source=blank and destination=volume. Changing this creates a new server.
	VolumeSize *float64 `json:"volumeSize,omitempty" tf:"volume_size,omitempty"`

	// Currently, the value can be SSD (ultra-I/O disk type),
	// SAS (high I/O disk type), or SATA (common I/O disk type)
	// OTC-API
	VolumeType *string `json:"volumeType,omitempty" tf:"volume_type,omitempty"`
}

type BlockDeviceObservation struct {

	// The boot index of the volume. It defaults to 0. Changing this creates a new server.
	BootIndex *float64 `json:"bootIndex,omitempty" tf:"boot_index,omitempty"`

	// Delete the volume / block device upon termination of the instance. Defaults to
	// false. Changing this creates a new server.
	DeleteOnTermination *bool `json:"deleteOnTermination,omitempty" tf:"delete_on_termination,omitempty"`

	// The type that gets created. Currently only support "volume". Changing this creates a
	// new server.
	DestinationType *string `json:"destinationType,omitempty" tf:"destination_type,omitempty"`

	// A unique name for the resource.
	DeviceName *string `json:"deviceName,omitempty" tf:"device_name,omitempty"`

	GuestFormat *string `json:"guestFormat,omitempty" tf:"guest_format,omitempty"`

	// The source type of the device. Must be one of
	// "blank", "image", "volume", or "snapshot". Changing this creates a new server.
	SourceType *string `json:"sourceType,omitempty" tf:"source_type,omitempty"`

	// The UUID of the image, volume, or snapshot. Changing
	// this creates a new server.
	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`

	// The size of the volume to create (in gigabytes). Required in the following combinations: source=image
	// and destination=volume, and source=blank and destination=volume. Changing this creates a new server.
	VolumeSize *float64 `json:"volumeSize,omitempty" tf:"volume_size,omitempty"`

	// Currently, the value can be SSD (ultra-I/O disk type),
	// SAS (high I/O disk type), or SATA (common I/O disk type)
	// OTC-API
	VolumeType *string `json:"volumeType,omitempty" tf:"volume_type,omitempty"`
}

type BlockDeviceParameters struct {

	// The boot index of the volume. It defaults to 0. Changing this creates a new server.
	// +kubebuilder:validation:Optional
	BootIndex *float64 `json:"bootIndex,omitempty" tf:"boot_index,omitempty"`

	// Delete the volume / block device upon termination of the instance. Defaults to
	// false. Changing this creates a new server.
	// +kubebuilder:validation:Optional
	DeleteOnTermination *bool `json:"deleteOnTermination,omitempty" tf:"delete_on_termination,omitempty"`

	// The type that gets created. Currently only support "volume". Changing this creates a
	// new server.
	// +kubebuilder:validation:Optional
	DestinationType *string `json:"destinationType,omitempty" tf:"destination_type,omitempty"`

	// A unique name for the resource.
	// +kubebuilder:validation:Optional
	DeviceName *string `json:"deviceName,omitempty" tf:"device_name,omitempty"`

	// +kubebuilder:validation:Optional
	GuestFormat *string `json:"guestFormat,omitempty" tf:"guest_format,omitempty"`

	// The source type of the device. Must be one of
	// "blank", "image", "volume", or "snapshot". Changing this creates a new server.
	// +kubebuilder:validation:Optional
	SourceType *string `json:"sourceType" tf:"source_type,omitempty"`

	// The UUID of the image, volume, or snapshot. Changing
	// this creates a new server.
	// +crossplane:generate:reference:type=github.com/opentelekomcloud/provider-opentelekomcloud/apis/blockstorage/v1alpha1.VolumeV2
	// +kubebuilder:validation:Optional
	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`

	// Reference to a VolumeV2 in blockstorage to populate uuid.
	// +kubebuilder:validation:Optional
	UUIDRef *v1.Reference `json:"uuidRef,omitempty" tf:"-"`

	// Selector for a VolumeV2 in blockstorage to populate uuid.
	// +kubebuilder:validation:Optional
	UUIDSelector *v1.Selector `json:"uuidSelector,omitempty" tf:"-"`

	// The size of the volume to create (in gigabytes). Required in the following combinations: source=image
	// and destination=volume, and source=blank and destination=volume. Changing this creates a new server.
	// +kubebuilder:validation:Optional
	VolumeSize *float64 `json:"volumeSize,omitempty" tf:"volume_size,omitempty"`

	// Currently, the value can be SSD (ultra-I/O disk type),
	// SAS (high I/O disk type), or SATA (common I/O disk type)
	// OTC-API
	// +kubebuilder:validation:Optional
	VolumeType *string `json:"volumeType,omitempty" tf:"volume_type,omitempty"`
}

type InstanceV2InitParameters struct {

	// The first detected Fixed IPv4 address or the Floating IP.
	AccessIPV4 *string `json:"accessIpV4,omitempty" tf:"access_ip_v4,omitempty"`

	// The first detected Fixed IPv6 address.
	AccessIPV6 *string `json:"accessIpV6,omitempty" tf:"access_ip_v6,omitempty"`

	// The administrative password to assign to the server. Changing this changes the root password
	// on the existing server.
	AdminPassSecretRef *v1.SecretKeySelector `json:"adminPassSecretRef,omitempty" tf:"-"`

	// Configures or deletes automatic recovery of an instance. Defaults to true.
	AutoRecovery *bool `json:"autoRecovery,omitempty" tf:"auto_recovery,omitempty"`

	// The availability zone in which to create the server. Changing this creates a new
	// server.
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// Configuration of block devices. The block_device structure is documented below. Changing
	// this creates a new server. You can specify multiple block devices which will create an instance with multiple disks.
	// This configuration is very flexible, so please see the
	// following reference
	// for more information.
	BlockDevice []BlockDeviceInitParameters `json:"blockDevice,omitempty" tf:"block_device,omitempty"`

	// Whether to use the config_drive feature to configure the instance. Changing this creates a
	// new server.
	ConfigDrive *bool `json:"configDrive,omitempty" tf:"config_drive,omitempty"`

	// Server description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The flavor ID of the desired flavor for the server.
	// Changing this resizes the existing server.
	FlavorID *string `json:"flavorId,omitempty" tf:"flavor_id,omitempty"`

	// The name of the desired flavor for the server. Changing
	// this resizes the existing server.
	FlavorName *string `json:"flavorName,omitempty" tf:"flavor_name,omitempty"`

	// The image ID of the desired image for the server. Changing this creates a new server.
	ImageID *string `json:"imageId,omitempty" tf:"image_id,omitempty"`

	// The name of the desired image for the server. Changing this creates a new server.
	ImageName *string `json:"imageName,omitempty" tf:"image_name,omitempty"`

	// The name of a key pair to put on the server. The key pair must already be created and
	// associated with the tenant's account. Changing this creates a new server.
	// +crossplane:generate:reference:type=github.com/opentelekomcloud/provider-opentelekomcloud/apis/compute/v1alpha1.KeypairV2
	KeyPair *string `json:"keyPair,omitempty" tf:"key_pair,omitempty"`

	// Reference to a KeypairV2 in compute to populate keyPair.
	// +kubebuilder:validation:Optional
	KeyPairRef *v1.Reference `json:"keyPairRef,omitempty" tf:"-"`

	// Selector for a KeypairV2 in compute to populate keyPair.
	// +kubebuilder:validation:Optional
	KeyPairSelector *v1.Selector `json:"keyPairSelector,omitempty" tf:"-"`

	// Metadata key/value pairs to make available from within the instance. Changing this updates the
	// existing server metadata.
	// +mapType=granular
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// A unique name for the resource.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// An array of one or more networks to attach to the instance. Required when there are multiple
	// networks defined for the tenant. The network object structure is documented below. Changing this creates a new server.
	Network []NetworkInitParameters `json:"network,omitempty" tf:"network,omitempty"`

	// Provide the VM state. Only active and shutoff are supported values.
	PowerState *string `json:"powerState,omitempty" tf:"power_state,omitempty"`

	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The path to the private key to use for SSH access. Required only if you want to
	// get the password from the windows instance.
	SSHPrivateKeyPathSecretRef *v1.SecretKeySelector `json:"sshPrivateKeyPathSecretRef,omitempty" tf:"-"`

	// Provide the Nova scheduler with hints on how the instance should be launched. The
	// available hints are described below.
	SchedulerHints []SchedulerHintsInitParameters `json:"schedulerHints,omitempty" tf:"scheduler_hints,omitempty"`

	// An array of one or more security group names to associate with the server. Changing
	// this results in adding/removing security groups from the existing server.
	// +crossplane:generate:reference:type=github.com/opentelekomcloud/provider-opentelekomcloud/apis/compute/v1alpha1.SecgroupV2
	// +listType=set
	SecurityGroups []*string `json:"securityGroups,omitempty" tf:"security_groups,omitempty"`

	// References to SecgroupV2 in compute to populate securityGroups.
	// +kubebuilder:validation:Optional
	SecurityGroupsRefs []v1.Reference `json:"securityGroupsRefs,omitempty" tf:"-"`

	// Selector for a list of SecgroupV2 in compute to populate securityGroups.
	// +kubebuilder:validation:Optional
	SecurityGroupsSelector *v1.Selector `json:"securityGroupsSelector,omitempty" tf:"-"`

	// Whether to try stop instance gracefully before destroying it, thus giving chance
	// for guest OS daemons to stop correctly. If instance doesn't stop within a timeout, it will be destroyed anyway.
	StopBeforeDestroy *bool `json:"stopBeforeDestroy,omitempty" tf:"stop_before_destroy,omitempty"`

	// Tags key/value pairs to associate with the instance.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The user data to provide when launching the instance. Changing this creates a new server.
	UserData *string `json:"userData,omitempty" tf:"user_data,omitempty"`
}

type InstanceV2Observation struct {

	// The first detected Fixed IPv4 address or the Floating IP.
	AccessIPV4 *string `json:"accessIpV4,omitempty" tf:"access_ip_v4,omitempty"`

	// The first detected Fixed IPv6 address.
	AccessIPV6 *string `json:"accessIpV6,omitempty" tf:"access_ip_v6,omitempty"`

	// +mapType=granular
	AllMetadata map[string]*string `json:"allMetadata,omitempty" tf:"all_metadata,omitempty"`

	// Configures or deletes automatic recovery of an instance. Defaults to true.
	AutoRecovery *bool `json:"autoRecovery,omitempty" tf:"auto_recovery,omitempty"`

	// The availability zone in which to create the server. Changing this creates a new
	// server.
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// Configuration of block devices. The block_device structure is documented below. Changing
	// this creates a new server. You can specify multiple block devices which will create an instance with multiple disks.
	// This configuration is very flexible, so please see the
	// following reference
	// for more information.
	BlockDevice []BlockDeviceObservation `json:"blockDevice,omitempty" tf:"block_device,omitempty"`

	// Whether to use the config_drive feature to configure the instance. Changing this creates a
	// new server.
	ConfigDrive *bool `json:"configDrive,omitempty" tf:"config_drive,omitempty"`

	// Server description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The flavor ID of the desired flavor for the server.
	// Changing this resizes the existing server.
	FlavorID *string `json:"flavorId,omitempty" tf:"flavor_id,omitempty"`

	// The name of the desired flavor for the server. Changing
	// this resizes the existing server.
	FlavorName *string `json:"flavorName,omitempty" tf:"flavor_name,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The image ID of the desired image for the server. Changing this creates a new server.
	ImageID *string `json:"imageId,omitempty" tf:"image_id,omitempty"`

	// The name of the desired image for the server. Changing this creates a new server.
	ImageName *string `json:"imageName,omitempty" tf:"image_name,omitempty"`

	// The name of a key pair to put on the server. The key pair must already be created and
	// associated with the tenant's account. Changing this creates a new server.
	KeyPair *string `json:"keyPair,omitempty" tf:"key_pair,omitempty"`

	// Metadata key/value pairs to make available from within the instance. Changing this updates the
	// existing server metadata.
	// +mapType=granular
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// A unique name for the resource.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// An array of one or more networks to attach to the instance. Required when there are multiple
	// networks defined for the tenant. The network object structure is documented below. Changing this creates a new server.
	Network []NetworkObservation `json:"network,omitempty" tf:"network,omitempty"`

	// Provide the VM state. Only active and shutoff are supported values.
	PowerState *string `json:"powerState,omitempty" tf:"power_state,omitempty"`

	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Provide the Nova scheduler with hints on how the instance should be launched. The
	// available hints are described below.
	SchedulerHints []SchedulerHintsObservation `json:"schedulerHints,omitempty" tf:"scheduler_hints,omitempty"`

	// An array of one or more security group names to associate with the server. Changing
	// this results in adding/removing security groups from the existing server.
	// +listType=set
	SecurityGroups []*string `json:"securityGroups,omitempty" tf:"security_groups,omitempty"`

	// Whether to try stop instance gracefully before destroying it, thus giving chance
	// for guest OS daemons to stop correctly. If instance doesn't stop within a timeout, it will be destroyed anyway.
	StopBeforeDestroy *bool `json:"stopBeforeDestroy,omitempty" tf:"stop_before_destroy,omitempty"`

	// Tags key/value pairs to associate with the instance.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The user data to provide when launching the instance. Changing this creates a new server.
	UserData *string `json:"userData,omitempty" tf:"user_data,omitempty"`

	VolumeAttached []VolumeAttachedObservation `json:"volumeAttached,omitempty" tf:"volume_attached,omitempty"`
}

type InstanceV2Parameters struct {

	// The first detected Fixed IPv4 address or the Floating IP.
	// +kubebuilder:validation:Optional
	AccessIPV4 *string `json:"accessIpV4,omitempty" tf:"access_ip_v4,omitempty"`

	// The first detected Fixed IPv6 address.
	// +kubebuilder:validation:Optional
	AccessIPV6 *string `json:"accessIpV6,omitempty" tf:"access_ip_v6,omitempty"`

	// The administrative password to assign to the server. Changing this changes the root password
	// on the existing server.
	// +kubebuilder:validation:Optional
	AdminPassSecretRef *v1.SecretKeySelector `json:"adminPassSecretRef,omitempty" tf:"-"`

	// Configures or deletes automatic recovery of an instance. Defaults to true.
	// +kubebuilder:validation:Optional
	AutoRecovery *bool `json:"autoRecovery,omitempty" tf:"auto_recovery,omitempty"`

	// The availability zone in which to create the server. Changing this creates a new
	// server.
	// +kubebuilder:validation:Optional
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// Configuration of block devices. The block_device structure is documented below. Changing
	// this creates a new server. You can specify multiple block devices which will create an instance with multiple disks.
	// This configuration is very flexible, so please see the
	// following reference
	// for more information.
	// +kubebuilder:validation:Optional
	BlockDevice []BlockDeviceParameters `json:"blockDevice,omitempty" tf:"block_device,omitempty"`

	// Whether to use the config_drive feature to configure the instance. Changing this creates a
	// new server.
	// +kubebuilder:validation:Optional
	ConfigDrive *bool `json:"configDrive,omitempty" tf:"config_drive,omitempty"`

	// Server description.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The flavor ID of the desired flavor for the server.
	// Changing this resizes the existing server.
	// +kubebuilder:validation:Optional
	FlavorID *string `json:"flavorId,omitempty" tf:"flavor_id,omitempty"`

	// The name of the desired flavor for the server. Changing
	// this resizes the existing server.
	// +kubebuilder:validation:Optional
	FlavorName *string `json:"flavorName,omitempty" tf:"flavor_name,omitempty"`

	// The image ID of the desired image for the server. Changing this creates a new server.
	// +kubebuilder:validation:Optional
	ImageID *string `json:"imageId,omitempty" tf:"image_id,omitempty"`

	// The name of the desired image for the server. Changing this creates a new server.
	// +kubebuilder:validation:Optional
	ImageName *string `json:"imageName,omitempty" tf:"image_name,omitempty"`

	// The name of a key pair to put on the server. The key pair must already be created and
	// associated with the tenant's account. Changing this creates a new server.
	// +crossplane:generate:reference:type=github.com/opentelekomcloud/provider-opentelekomcloud/apis/compute/v1alpha1.KeypairV2
	// +kubebuilder:validation:Optional
	KeyPair *string `json:"keyPair,omitempty" tf:"key_pair,omitempty"`

	// Reference to a KeypairV2 in compute to populate keyPair.
	// +kubebuilder:validation:Optional
	KeyPairRef *v1.Reference `json:"keyPairRef,omitempty" tf:"-"`

	// Selector for a KeypairV2 in compute to populate keyPair.
	// +kubebuilder:validation:Optional
	KeyPairSelector *v1.Selector `json:"keyPairSelector,omitempty" tf:"-"`

	// Metadata key/value pairs to make available from within the instance. Changing this updates the
	// existing server metadata.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// A unique name for the resource.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// An array of one or more networks to attach to the instance. Required when there are multiple
	// networks defined for the tenant. The network object structure is documented below. Changing this creates a new server.
	// +kubebuilder:validation:Optional
	Network []NetworkParameters `json:"network,omitempty" tf:"network,omitempty"`

	// Provide the VM state. Only active and shutoff are supported values.
	// +kubebuilder:validation:Optional
	PowerState *string `json:"powerState,omitempty" tf:"power_state,omitempty"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The path to the private key to use for SSH access. Required only if you want to
	// get the password from the windows instance.
	// +kubebuilder:validation:Optional
	SSHPrivateKeyPathSecretRef *v1.SecretKeySelector `json:"sshPrivateKeyPathSecretRef,omitempty" tf:"-"`

	// Provide the Nova scheduler with hints on how the instance should be launched. The
	// available hints are described below.
	// +kubebuilder:validation:Optional
	SchedulerHints []SchedulerHintsParameters `json:"schedulerHints,omitempty" tf:"scheduler_hints,omitempty"`

	// An array of one or more security group names to associate with the server. Changing
	// this results in adding/removing security groups from the existing server.
	// +crossplane:generate:reference:type=github.com/opentelekomcloud/provider-opentelekomcloud/apis/compute/v1alpha1.SecgroupV2
	// +kubebuilder:validation:Optional
	// +listType=set
	SecurityGroups []*string `json:"securityGroups,omitempty" tf:"security_groups,omitempty"`

	// References to SecgroupV2 in compute to populate securityGroups.
	// +kubebuilder:validation:Optional
	SecurityGroupsRefs []v1.Reference `json:"securityGroupsRefs,omitempty" tf:"-"`

	// Selector for a list of SecgroupV2 in compute to populate securityGroups.
	// +kubebuilder:validation:Optional
	SecurityGroupsSelector *v1.Selector `json:"securityGroupsSelector,omitempty" tf:"-"`

	// Whether to try stop instance gracefully before destroying it, thus giving chance
	// for guest OS daemons to stop correctly. If instance doesn't stop within a timeout, it will be destroyed anyway.
	// +kubebuilder:validation:Optional
	StopBeforeDestroy *bool `json:"stopBeforeDestroy,omitempty" tf:"stop_before_destroy,omitempty"`

	// Tags key/value pairs to associate with the instance.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The user data to provide when launching the instance. Changing this creates a new server.
	// +kubebuilder:validation:Optional
	UserData *string `json:"userData,omitempty" tf:"user_data,omitempty"`
}

type NetworkInitParameters struct {

	// Specifies if this network should be used for provisioning access. Accepts true or false.
	// Defaults to false.
	AccessNetwork *bool `json:"accessNetwork,omitempty" tf:"access_network,omitempty"`

	// Specifies a fixed IPv4 address to be used on this network. Changing this creates a new
	// server.
	FixedIPV4 *string `json:"fixedIpV4,omitempty" tf:"fixed_ip_v4,omitempty"`

	// Specifies a fixed IPv6 address to be used on this network. Changing this creates a new
	// server.
	FixedIPV6 *string `json:"fixedIpV6,omitempty" tf:"fixed_ip_v6,omitempty"`

	// The human-readable name of the network. Changing this creates
	// a new server.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The port UUID of a network to attach to the server. Changing
	// this creates a new server.
	Port *string `json:"port,omitempty" tf:"port,omitempty"`

	// The network UUID to attach to the server. Changing this
	// creates a new server.
	// +crossplane:generate:reference:type=github.com/opentelekomcloud/provider-opentelekomcloud/apis/networking/v1alpha1.NetworkV2
	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`

	// Reference to a NetworkV2 in networking to populate uuid.
	// +kubebuilder:validation:Optional
	UUIDRef *v1.Reference `json:"uuidRef,omitempty" tf:"-"`

	// Selector for a NetworkV2 in networking to populate uuid.
	// +kubebuilder:validation:Optional
	UUIDSelector *v1.Selector `json:"uuidSelector,omitempty" tf:"-"`
}

type NetworkObservation struct {

	// Specifies if this network should be used for provisioning access. Accepts true or false.
	// Defaults to false.
	AccessNetwork *bool `json:"accessNetwork,omitempty" tf:"access_network,omitempty"`

	// Specifies a fixed IPv4 address to be used on this network. Changing this creates a new
	// server.
	FixedIPV4 *string `json:"fixedIpV4,omitempty" tf:"fixed_ip_v4,omitempty"`

	// Specifies a fixed IPv6 address to be used on this network. Changing this creates a new
	// server.
	FixedIPV6 *string `json:"fixedIpV6,omitempty" tf:"fixed_ip_v6,omitempty"`

	Mac *string `json:"mac,omitempty" tf:"mac,omitempty"`

	// The human-readable name of the network. Changing this creates
	// a new server.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The port UUID of a network to attach to the server. Changing
	// this creates a new server.
	Port *string `json:"port,omitempty" tf:"port,omitempty"`

	// The network UUID to attach to the server. Changing this
	// creates a new server.
	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`
}

type NetworkParameters struct {

	// Specifies if this network should be used for provisioning access. Accepts true or false.
	// Defaults to false.
	// +kubebuilder:validation:Optional
	AccessNetwork *bool `json:"accessNetwork,omitempty" tf:"access_network,omitempty"`

	// Specifies a fixed IPv4 address to be used on this network. Changing this creates a new
	// server.
	// +kubebuilder:validation:Optional
	FixedIPV4 *string `json:"fixedIpV4,omitempty" tf:"fixed_ip_v4,omitempty"`

	// Specifies a fixed IPv6 address to be used on this network. Changing this creates a new
	// server.
	// +kubebuilder:validation:Optional
	FixedIPV6 *string `json:"fixedIpV6,omitempty" tf:"fixed_ip_v6,omitempty"`

	// The human-readable name of the network. Changing this creates
	// a new server.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The port UUID of a network to attach to the server. Changing
	// this creates a new server.
	// +kubebuilder:validation:Optional
	Port *string `json:"port,omitempty" tf:"port,omitempty"`

	// The network UUID to attach to the server. Changing this
	// creates a new server.
	// +crossplane:generate:reference:type=github.com/opentelekomcloud/provider-opentelekomcloud/apis/networking/v1alpha1.NetworkV2
	// +kubebuilder:validation:Optional
	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`

	// Reference to a NetworkV2 in networking to populate uuid.
	// +kubebuilder:validation:Optional
	UUIDRef *v1.Reference `json:"uuidRef,omitempty" tf:"-"`

	// Selector for a NetworkV2 in networking to populate uuid.
	// +kubebuilder:validation:Optional
	UUIDSelector *v1.Selector `json:"uuidSelector,omitempty" tf:"-"`
}

type SchedulerHintsInitParameters struct {

	// An IP Address in CIDR form. The instance will be placed on a compute node that is in
	// the same subnet.
	BuildNearHostIP *string `json:"buildNearHostIp,omitempty" tf:"build_near_host_ip,omitempty"`

	// The ID of DeH. This parameter takes effect only when the value of tenancy is dedicated.
	DehID *string `json:"dehId,omitempty" tf:"deh_id,omitempty"`

	// A list of instance UUIDs. The instance will be scheduled on a different host than all
	// other instances.
	DifferentHost []*string `json:"differentHost,omitempty" tf:"different_host,omitempty"`

	// A UUID of a Server Group. The instance will be placed into that group.
	Group *string `json:"group,omitempty" tf:"group,omitempty"`

	// A conditional query that a compute node must pass in order to host an instance.
	Query []*string `json:"query,omitempty" tf:"query,omitempty"`

	// A list of instance UUIDs. The instance will be scheduled on the same host of those specified.
	SameHost []*string `json:"sameHost,omitempty" tf:"same_host,omitempty"`

	// The name of a cell to host the instance.
	TargetCell *string `json:"targetCell,omitempty" tf:"target_cell,omitempty"`

	// The tenancy specifies whether the ECS is to be created on a Dedicated Host
	// (DeH) or in a shared pool.
	Tenancy *string `json:"tenancy,omitempty" tf:"tenancy,omitempty"`
}

type SchedulerHintsObservation struct {

	// An IP Address in CIDR form. The instance will be placed on a compute node that is in
	// the same subnet.
	BuildNearHostIP *string `json:"buildNearHostIp,omitempty" tf:"build_near_host_ip,omitempty"`

	// The ID of DeH. This parameter takes effect only when the value of tenancy is dedicated.
	DehID *string `json:"dehId,omitempty" tf:"deh_id,omitempty"`

	// A list of instance UUIDs. The instance will be scheduled on a different host than all
	// other instances.
	DifferentHost []*string `json:"differentHost,omitempty" tf:"different_host,omitempty"`

	// A UUID of a Server Group. The instance will be placed into that group.
	Group *string `json:"group,omitempty" tf:"group,omitempty"`

	// A conditional query that a compute node must pass in order to host an instance.
	Query []*string `json:"query,omitempty" tf:"query,omitempty"`

	// A list of instance UUIDs. The instance will be scheduled on the same host of those specified.
	SameHost []*string `json:"sameHost,omitempty" tf:"same_host,omitempty"`

	// The name of a cell to host the instance.
	TargetCell *string `json:"targetCell,omitempty" tf:"target_cell,omitempty"`

	// The tenancy specifies whether the ECS is to be created on a Dedicated Host
	// (DeH) or in a shared pool.
	Tenancy *string `json:"tenancy,omitempty" tf:"tenancy,omitempty"`
}

type SchedulerHintsParameters struct {

	// An IP Address in CIDR form. The instance will be placed on a compute node that is in
	// the same subnet.
	// +kubebuilder:validation:Optional
	BuildNearHostIP *string `json:"buildNearHostIp,omitempty" tf:"build_near_host_ip,omitempty"`

	// The ID of DeH. This parameter takes effect only when the value of tenancy is dedicated.
	// +kubebuilder:validation:Optional
	DehID *string `json:"dehId,omitempty" tf:"deh_id,omitempty"`

	// A list of instance UUIDs. The instance will be scheduled on a different host than all
	// other instances.
	// +kubebuilder:validation:Optional
	DifferentHost []*string `json:"differentHost,omitempty" tf:"different_host,omitempty"`

	// A UUID of a Server Group. The instance will be placed into that group.
	// +kubebuilder:validation:Optional
	Group *string `json:"group,omitempty" tf:"group,omitempty"`

	// A conditional query that a compute node must pass in order to host an instance.
	// +kubebuilder:validation:Optional
	Query []*string `json:"query,omitempty" tf:"query,omitempty"`

	// A list of instance UUIDs. The instance will be scheduled on the same host of those specified.
	// +kubebuilder:validation:Optional
	SameHost []*string `json:"sameHost,omitempty" tf:"same_host,omitempty"`

	// The name of a cell to host the instance.
	// +kubebuilder:validation:Optional
	TargetCell *string `json:"targetCell,omitempty" tf:"target_cell,omitempty"`

	// The tenancy specifies whether the ECS is to be created on a Dedicated Host
	// (DeH) or in a shared pool.
	// +kubebuilder:validation:Optional
	Tenancy *string `json:"tenancy,omitempty" tf:"tenancy,omitempty"`
}

type VolumeAttachedInitParameters struct {
}

type VolumeAttachedObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type VolumeAttachedParameters struct {
}

// InstanceV2Spec defines the desired state of InstanceV2
type InstanceV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InstanceV2Parameters `json:"forProvider"`
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
	InitProvider InstanceV2InitParameters `json:"initProvider,omitempty"`
}

// InstanceV2Status defines the observed state of InstanceV2.
type InstanceV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InstanceV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// InstanceV2 is the Schema for the InstanceV2s API. Manages an ECS Instance resource within OpenTelekomCloud.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,opentelekomcloud}
type InstanceV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   InstanceV2Spec   `json:"spec"`
	Status InstanceV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceV2List contains a list of InstanceV2s
type InstanceV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InstanceV2 `json:"items"`
}

// Repository type metadata.
var (
	InstanceV2_Kind             = "InstanceV2"
	InstanceV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: InstanceV2_Kind}.String()
	InstanceV2_KindAPIVersion   = InstanceV2_Kind + "." + CRDGroupVersion.String()
	InstanceV2_GroupVersionKind = CRDGroupVersion.WithKind(InstanceV2_Kind)
)

func init() {
	SchemeBuilder.Register(&InstanceV2{}, &InstanceV2List{})
}
