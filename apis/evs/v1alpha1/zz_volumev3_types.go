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

type AttachmentInitParameters struct {
}

type AttachmentObservation struct {
	Device *string `json:"device,omitempty" tf:"device,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`
}

type AttachmentParameters struct {
}

type VolumeV3InitParameters struct {

	// The availability zone for the volume.
	// Changing this creates a new volume.
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// The backup ID from which to create the volume.
	// Changing this creates a new volume.
	BackupID *string `json:"backupId,omitempty" tf:"backup_id,omitempty"`

	// Specifies to delete all snapshots associated with the EVS disk. Default is false.
	Cascade *bool `json:"cascade,omitempty" tf:"cascade,omitempty"`

	// A description of the volume. Changing this updates the volume's description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The device type of volume to create. Valid options are VBD and SCSI.
	// Defaults to VBD. Changing this creates a new volume.
	DeviceType *string `json:"deviceType,omitempty" tf:"device_type,omitempty"`

	// The image ID from which to create the volume.
	// Changing this creates a new volume.
	ImageID *string `json:"imageId,omitempty" tf:"image_id,omitempty"`

	// The Encryption KMS ID to create the volume.
	// Changing this creates a new volume.
	KMSID *string `json:"kmsId,omitempty" tf:"kms_id,omitempty"`

	// Specifies whether the disk is shareable. The default value is false.
	// Changing this creates a new volume.
	Multiattach *bool `json:"multiattach,omitempty" tf:"multiattach,omitempty"`

	// A unique name for the volume. Changing this updates the volume's name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The size of the volume to create (in gigabytes). This parameter is mandatory when
	// you create an empty EVS disk or use an image or a snapshot to create an EVS disk.
	// Decreasing this value creates a new volume.
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// The snapshot ID from which to create the volume.
	// Changing this creates a new volume.
	SnapshotID *string `json:"snapshotId,omitempty" tf:"snapshot_id,omitempty"`

	// Tags key/value pairs to associate with the volume.
	// Changing this updates the existing volume tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The type of volume to create.
	// Currently, the value can be SSD, SAS, SATA, co-p1, or uh-l1.
	// Changing this creates a new volume.
	VolumeType *string `json:"volumeType,omitempty" tf:"volume_type,omitempty"`
}

type VolumeV3Observation struct {

	// If a volume is attached to an instance, this attribute will
	// display the Attachment ID, Instance ID, and the Device as the Instance sees it.
	Attachment []AttachmentObservation `json:"attachment,omitempty" tf:"attachment,omitempty"`

	// The availability zone for the volume.
	// Changing this creates a new volume.
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// The backup ID from which to create the volume.
	// Changing this creates a new volume.
	BackupID *string `json:"backupId,omitempty" tf:"backup_id,omitempty"`

	// Specifies to delete all snapshots associated with the EVS disk. Default is false.
	Cascade *bool `json:"cascade,omitempty" tf:"cascade,omitempty"`

	// A description of the volume. Changing this updates the volume's description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The device type of volume to create. Valid options are VBD and SCSI.
	// Defaults to VBD. Changing this creates a new volume.
	DeviceType *string `json:"deviceType,omitempty" tf:"device_type,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The image ID from which to create the volume.
	// Changing this creates a new volume.
	ImageID *string `json:"imageId,omitempty" tf:"image_id,omitempty"`

	// The Encryption KMS ID to create the volume.
	// Changing this creates a new volume.
	KMSID *string `json:"kmsId,omitempty" tf:"kms_id,omitempty"`

	// Specifies whether the disk is shareable. The default value is false.
	// Changing this creates a new volume.
	Multiattach *bool `json:"multiattach,omitempty" tf:"multiattach,omitempty"`

	// A unique name for the volume. Changing this updates the volume's name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The size of the volume to create (in gigabytes). This parameter is mandatory when
	// you create an empty EVS disk or use an image or a snapshot to create an EVS disk.
	// Decreasing this value creates a new volume.
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// The snapshot ID from which to create the volume.
	// Changing this creates a new volume.
	SnapshotID *string `json:"snapshotId,omitempty" tf:"snapshot_id,omitempty"`

	// Tags key/value pairs to associate with the volume.
	// Changing this updates the existing volume tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The type of volume to create.
	// Currently, the value can be SSD, SAS, SATA, co-p1, or uh-l1.
	// Changing this creates a new volume.
	VolumeType *string `json:"volumeType,omitempty" tf:"volume_type,omitempty"`

	// Specifies the unique identifier used for mounting the EVS disk.
	Wwn *string `json:"wwn,omitempty" tf:"wwn,omitempty"`
}

type VolumeV3Parameters struct {

	// The availability zone for the volume.
	// Changing this creates a new volume.
	// +kubebuilder:validation:Optional
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// The backup ID from which to create the volume.
	// Changing this creates a new volume.
	// +kubebuilder:validation:Optional
	BackupID *string `json:"backupId,omitempty" tf:"backup_id,omitempty"`

	// Specifies to delete all snapshots associated with the EVS disk. Default is false.
	// +kubebuilder:validation:Optional
	Cascade *bool `json:"cascade,omitempty" tf:"cascade,omitempty"`

	// A description of the volume. Changing this updates the volume's description.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The device type of volume to create. Valid options are VBD and SCSI.
	// Defaults to VBD. Changing this creates a new volume.
	// +kubebuilder:validation:Optional
	DeviceType *string `json:"deviceType,omitempty" tf:"device_type,omitempty"`

	// The image ID from which to create the volume.
	// Changing this creates a new volume.
	// +kubebuilder:validation:Optional
	ImageID *string `json:"imageId,omitempty" tf:"image_id,omitempty"`

	// The Encryption KMS ID to create the volume.
	// Changing this creates a new volume.
	// +kubebuilder:validation:Optional
	KMSID *string `json:"kmsId,omitempty" tf:"kms_id,omitempty"`

	// Specifies whether the disk is shareable. The default value is false.
	// Changing this creates a new volume.
	// +kubebuilder:validation:Optional
	Multiattach *bool `json:"multiattach,omitempty" tf:"multiattach,omitempty"`

	// A unique name for the volume. Changing this updates the volume's name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The size of the volume to create (in gigabytes). This parameter is mandatory when
	// you create an empty EVS disk or use an image or a snapshot to create an EVS disk.
	// Decreasing this value creates a new volume.
	// +kubebuilder:validation:Optional
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// The snapshot ID from which to create the volume.
	// Changing this creates a new volume.
	// +kubebuilder:validation:Optional
	SnapshotID *string `json:"snapshotId,omitempty" tf:"snapshot_id,omitempty"`

	// Tags key/value pairs to associate with the volume.
	// Changing this updates the existing volume tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The type of volume to create.
	// Currently, the value can be SSD, SAS, SATA, co-p1, or uh-l1.
	// Changing this creates a new volume.
	// +kubebuilder:validation:Optional
	VolumeType *string `json:"volumeType,omitempty" tf:"volume_type,omitempty"`
}

// VolumeV3Spec defines the desired state of VolumeV3
type VolumeV3Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VolumeV3Parameters `json:"forProvider"`
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
	InitProvider VolumeV3InitParameters `json:"initProvider,omitempty"`
}

// VolumeV3Status defines the observed state of VolumeV3.
type VolumeV3Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VolumeV3Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VolumeV3 is the Schema for the VolumeV3s API. Manages an EVS resource within OpenTelekomCloud.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,opentelekomcloud}
type VolumeV3 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.availabilityZone) || (has(self.initProvider) && has(self.initProvider.availabilityZone))",message="spec.forProvider.availabilityZone is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.volumeType) || (has(self.initProvider) && has(self.initProvider.volumeType))",message="spec.forProvider.volumeType is a required parameter"
	Spec   VolumeV3Spec   `json:"spec"`
	Status VolumeV3Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VolumeV3List contains a list of VolumeV3s
type VolumeV3List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VolumeV3 `json:"items"`
}

// Repository type metadata.
var (
	VolumeV3_Kind             = "VolumeV3"
	VolumeV3_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VolumeV3_Kind}.String()
	VolumeV3_KindAPIVersion   = VolumeV3_Kind + "." + CRDGroupVersion.String()
	VolumeV3_GroupVersionKind = CRDGroupVersion.WithKind(VolumeV3_Kind)
)

func init() {
	SchemeBuilder.Register(&VolumeV3{}, &VolumeV3List{})
}