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

type BackupStrategyInitParameters struct {

	// to 0, this parameter is no need to set.
	KeepDays *float64 `json:"keepDays,omitempty" tf:"keep_days,omitempty"`

	// Specifies the backup cycle. Data will be automatically backed up on the
	// selected days every week.
	Period *string `json:"period,omitempty" tf:"period,omitempty"`

	// Specifies the backup time window. Automated backups will be triggered
	// during the backup time window. The value cannot be empty. It must be a valid value in the
	// "hh:mm-HH:MM" format. The current time is in the UTC format.
	StartTime *string `json:"startTime,omitempty" tf:"start_time,omitempty"`
}

type BackupStrategyObservation struct {

	// to 0, this parameter is no need to set.
	KeepDays *float64 `json:"keepDays,omitempty" tf:"keep_days,omitempty"`

	// Specifies the backup cycle. Data will be automatically backed up on the
	// selected days every week.
	Period *string `json:"period,omitempty" tf:"period,omitempty"`

	// Specifies the backup time window. Automated backups will be triggered
	// during the backup time window. The value cannot be empty. It must be a valid value in the
	// "hh:mm-HH:MM" format. The current time is in the UTC format.
	StartTime *string `json:"startTime,omitempty" tf:"start_time,omitempty"`
}

type BackupStrategyParameters struct {

	// to 0, this parameter is no need to set.
	// +kubebuilder:validation:Optional
	KeepDays *float64 `json:"keepDays" tf:"keep_days,omitempty"`

	// Specifies the backup cycle. Data will be automatically backed up on the
	// selected days every week.
	// +kubebuilder:validation:Optional
	Period *string `json:"period,omitempty" tf:"period,omitempty"`

	// Specifies the backup time window. Automated backups will be triggered
	// during the backup time window. The value cannot be empty. It must be a valid value in the
	// "hh:mm-HH:MM" format. The current time is in the UTC format.
	// +kubebuilder:validation:Optional
	StartTime *string `json:"startTime" tf:"start_time,omitempty"`
}

type FlavorInitParameters struct {

	// Specifies the node quantity. Valid value:
	Num *float64 `json:"num,omitempty" tf:"num,omitempty"`

	// Specifies the disk size. The value must be a multiple of 10. The unit is GB.
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// Specifies the resource specification code.
	SpecCode *string `json:"specCode,omitempty" tf:"spec_code,omitempty"`

	// Specifies the disk type. Valid value: ULTRAHIGH which indicates the type SSD.
	Storage *string `json:"storage,omitempty" tf:"storage,omitempty"`

	// Specifies the node type. Valid value:
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type FlavorObservation struct {

	// Specifies the node quantity. Valid value:
	Num *float64 `json:"num,omitempty" tf:"num,omitempty"`

	// Specifies the disk size. The value must be a multiple of 10. The unit is GB.
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// Specifies the resource specification code.
	SpecCode *string `json:"specCode,omitempty" tf:"spec_code,omitempty"`

	// Specifies the disk type. Valid value: ULTRAHIGH which indicates the type SSD.
	Storage *string `json:"storage,omitempty" tf:"storage,omitempty"`

	// Specifies the node type. Valid value:
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type FlavorParameters struct {

	// Specifies the node quantity. Valid value:
	// +kubebuilder:validation:Optional
	Num *float64 `json:"num" tf:"num,omitempty"`

	// Specifies the disk size. The value must be a multiple of 10. The unit is GB.
	// +kubebuilder:validation:Optional
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// Specifies the resource specification code.
	// +kubebuilder:validation:Optional
	SpecCode *string `json:"specCode" tf:"spec_code,omitempty"`

	// Specifies the disk type. Valid value: ULTRAHIGH which indicates the type SSD.
	// +kubebuilder:validation:Optional
	Storage *string `json:"storage,omitempty" tf:"storage,omitempty"`

	// Specifies the node type. Valid value:
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type InstanceV3DatastoreInitParameters struct {

	// Specifies the storage engine. Currently, DDS supports the WiredTiger and RocksDB
	// storage engine. The values are wiredTiger, rocksDB.
	// WiredTiger engine supports versions 3.2, 3.4, 4.0 while RocksDB supports versions 4.2, 4.4
	StorageEngine *string `json:"storageEngine,omitempty" tf:"storage_engine,omitempty"`

	// Specifies the database type. DDS Community Edition is supported.
	// The value is DDS-Community.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Specifies the database version.
	// The values are 3.2, 3.4, 4.0, 4.2, 4.4.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type InstanceV3DatastoreObservation struct {

	// Specifies the storage engine. Currently, DDS supports the WiredTiger and RocksDB
	// storage engine. The values are wiredTiger, rocksDB.
	// WiredTiger engine supports versions 3.2, 3.4, 4.0 while RocksDB supports versions 4.2, 4.4
	StorageEngine *string `json:"storageEngine,omitempty" tf:"storage_engine,omitempty"`

	// Specifies the database type. DDS Community Edition is supported.
	// The value is DDS-Community.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Specifies the database version.
	// The values are 3.2, 3.4, 4.0, 4.2, 4.4.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type InstanceV3DatastoreParameters struct {

	// Specifies the storage engine. Currently, DDS supports the WiredTiger and RocksDB
	// storage engine. The values are wiredTiger, rocksDB.
	// WiredTiger engine supports versions 3.2, 3.4, 4.0 while RocksDB supports versions 4.2, 4.4
	// +kubebuilder:validation:Optional
	StorageEngine *string `json:"storageEngine,omitempty" tf:"storage_engine,omitempty"`

	// Specifies the database type. DDS Community Edition is supported.
	// The value is DDS-Community.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`

	// Specifies the database version.
	// The values are 3.2, 3.4, 4.0, 4.2, 4.4.
	// +kubebuilder:validation:Optional
	Version *string `json:"version" tf:"version,omitempty"`
}

type InstanceV3InitParameters struct {

	// Specifies the ID of the availability zone.
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// Specifies the advanced backup policy. The structure is
	// described below.
	BackupStrategy []BackupStrategyInitParameters `json:"backupStrategy,omitempty" tf:"backup_strategy,omitempty"`

	// Specifies database information. The structure is described
	// below.
	Datastore []InstanceV3DatastoreInitParameters `json:"datastore,omitempty" tf:"datastore,omitempty"`

	// Specifies the disk encryption ID of the instance.
	DiskEncryptionID *string `json:"diskEncryptionId,omitempty" tf:"disk_encryption_id,omitempty"`

	// Specifies the flavor information. The structure is described below.
	// Changing this creates a new instance.
	Flavor []FlavorInitParameters `json:"flavor,omitempty" tf:"flavor,omitempty"`

	// Specifies the mode of the database instance.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// Specifies the DB instance name. The DB instance name of the same
	// type is unique in the same tenant.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the Administrator password of the database instance.
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// Specifies the database access port. The valid values are range from 2100 to 9500 and
	// 27017, 27018, 27019. Defaults to 8635.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Specifies the region of the DDS instance.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies whether to enable or disable SSL. Defaults to true.
	// -> The instance will be restarted in the background when switching SSL. Please operate with caution.
	SSL *bool `json:"ssl,omitempty" tf:"ssl,omitempty"`

	// Specifies the security group ID of the DDS instance.
	// +crossplane:generate:reference:type=github.com/opentelekomcloud/provider-opentelekomcloud/apis/compute/v1alpha1.SecgroupV2
	SecurityGroupID *string `json:"securityGroupId,omitempty" tf:"security_group_id,omitempty"`

	// Reference to a SecgroupV2 in compute to populate securityGroupId.
	// +kubebuilder:validation:Optional
	SecurityGroupIDRef *v1.Reference `json:"securityGroupIdRef,omitempty" tf:"-"`

	// Selector for a SecgroupV2 in compute to populate securityGroupId.
	// +kubebuilder:validation:Optional
	SecurityGroupIDSelector *v1.Selector `json:"securityGroupIdSelector,omitempty" tf:"-"`

	// Specifies the subnet Network ID.
	// +crossplane:generate:reference:type=github.com/opentelekomcloud/provider-opentelekomcloud/apis/vpc/v1alpha1.SubnetV1
	// +crossplane:generate:reference:extractor=github.com/opentelekomcloud/provider-opentelekomcloud/config/common.ExtractNetworkID()
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a SubnetV1 in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a SubnetV1 in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// Tags key/value pairs to associate with the volume.
	// Changing this updates the existing volume tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the VPC ID.
	// +crossplane:generate:reference:type=github.com/opentelekomcloud/provider-opentelekomcloud/apis/vpc/v1alpha1.VpcV1
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VpcV1 in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VpcV1 in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`
}

type InstanceV3Observation struct {

	// Specifies the ID of the availability zone.
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// Specifies the advanced backup policy. The structure is
	// described below.
	BackupStrategy []BackupStrategyObservation `json:"backupStrategy,omitempty" tf:"backup_strategy,omitempty"`

	// Indicates the creation time.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// Indicates the DB Administator name.
	DBUsername *string `json:"dbUsername,omitempty" tf:"db_username,omitempty"`

	// Specifies database information. The structure is described
	// below.
	Datastore []InstanceV3DatastoreObservation `json:"datastore,omitempty" tf:"datastore,omitempty"`

	// Specifies the disk encryption ID of the instance.
	DiskEncryptionID *string `json:"diskEncryptionId,omitempty" tf:"disk_encryption_id,omitempty"`

	// Specifies the flavor information. The structure is described below.
	// Changing this creates a new instance.
	Flavor []FlavorObservation `json:"flavor,omitempty" tf:"flavor,omitempty"`

	// Indicates the node ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the mode of the database instance.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// Specifies the DB instance name. The DB instance name of the same
	// type is unique in the same tenant.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Indicates the instance nodes information. Structure is documented below.
	Nodes []NodesObservation `json:"nodes,omitempty" tf:"nodes,omitempty"`

	// Indicates the billing mode. 0: indicates the pay-per-use billing mode.
	PayMode *string `json:"payMode,omitempty" tf:"pay_mode,omitempty"`

	// Specifies the database access port. The valid values are range from 2100 to 9500 and
	// 27017, 27018, 27019. Defaults to 8635.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Specifies the region of the DDS instance.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies whether to enable or disable SSL. Defaults to true.
	// -> The instance will be restarted in the background when switching SSL. Please operate with caution.
	SSL *bool `json:"ssl,omitempty" tf:"ssl,omitempty"`

	// Specifies the security group ID of the DDS instance.
	SecurityGroupID *string `json:"securityGroupId,omitempty" tf:"security_group_id,omitempty"`

	// Indicates the DB instance status.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Specifies the subnet Network ID.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Tags key/value pairs to associate with the volume.
	// Changing this updates the existing volume tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Indicates the time zone.
	TimeZone *string `json:"timeZone,omitempty" tf:"time_zone,omitempty"`

	// Indicates the update time.
	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`

	// Specifies the VPC ID.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

type InstanceV3Parameters struct {

	// Specifies the ID of the availability zone.
	// +kubebuilder:validation:Optional
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// Specifies the advanced backup policy. The structure is
	// described below.
	// +kubebuilder:validation:Optional
	BackupStrategy []BackupStrategyParameters `json:"backupStrategy,omitempty" tf:"backup_strategy,omitempty"`

	// Specifies database information. The structure is described
	// below.
	// +kubebuilder:validation:Optional
	Datastore []InstanceV3DatastoreParameters `json:"datastore,omitempty" tf:"datastore,omitempty"`

	// Specifies the disk encryption ID of the instance.
	// +kubebuilder:validation:Optional
	DiskEncryptionID *string `json:"diskEncryptionId,omitempty" tf:"disk_encryption_id,omitempty"`

	// Specifies the flavor information. The structure is described below.
	// Changing this creates a new instance.
	// +kubebuilder:validation:Optional
	Flavor []FlavorParameters `json:"flavor,omitempty" tf:"flavor,omitempty"`

	// Specifies the mode of the database instance.
	// +kubebuilder:validation:Optional
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// Specifies the DB instance name. The DB instance name of the same
	// type is unique in the same tenant.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the Administrator password of the database instance.
	// +kubebuilder:validation:Optional
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// Specifies the database access port. The valid values are range from 2100 to 9500 and
	// 27017, 27018, 27019. Defaults to 8635.
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Specifies the region of the DDS instance.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies whether to enable or disable SSL. Defaults to true.
	// -> The instance will be restarted in the background when switching SSL. Please operate with caution.
	// +kubebuilder:validation:Optional
	SSL *bool `json:"ssl,omitempty" tf:"ssl,omitempty"`

	// Specifies the security group ID of the DDS instance.
	// +crossplane:generate:reference:type=github.com/opentelekomcloud/provider-opentelekomcloud/apis/compute/v1alpha1.SecgroupV2
	// +kubebuilder:validation:Optional
	SecurityGroupID *string `json:"securityGroupId,omitempty" tf:"security_group_id,omitempty"`

	// Reference to a SecgroupV2 in compute to populate securityGroupId.
	// +kubebuilder:validation:Optional
	SecurityGroupIDRef *v1.Reference `json:"securityGroupIdRef,omitempty" tf:"-"`

	// Selector for a SecgroupV2 in compute to populate securityGroupId.
	// +kubebuilder:validation:Optional
	SecurityGroupIDSelector *v1.Selector `json:"securityGroupIdSelector,omitempty" tf:"-"`

	// Specifies the subnet Network ID.
	// +crossplane:generate:reference:type=github.com/opentelekomcloud/provider-opentelekomcloud/apis/vpc/v1alpha1.SubnetV1
	// +crossplane:generate:reference:extractor=github.com/opentelekomcloud/provider-opentelekomcloud/config/common.ExtractNetworkID()
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a SubnetV1 in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a SubnetV1 in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// Tags key/value pairs to associate with the volume.
	// Changing this updates the existing volume tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the VPC ID.
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

type NodesInitParameters struct {
}

type NodesObservation struct {

	// Indicates the node ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Indicates the node name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Indicates the private IP address of a node. This parameter is valid only for
	// mongos nodes, replica set instances.
	PrivateIP *string `json:"privateIp,omitempty" tf:"private_ip,omitempty"`

	// Indicates the EIP that has been bound on a node. This parameter is valid only for
	// mongos nodes of cluster instances, primary nodes and secondary nodes of replica set instances.
	PublicIP *string `json:"publicIp,omitempty" tf:"public_ip,omitempty"`

	// Indicates the node role.
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// Indicates the node status.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Indicates the node type.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type NodesParameters struct {
}

// InstanceV3Spec defines the desired state of InstanceV3
type InstanceV3Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InstanceV3Parameters `json:"forProvider"`
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
	InitProvider InstanceV3InitParameters `json:"initProvider,omitempty"`
}

// InstanceV3Status defines the observed state of InstanceV3.
type InstanceV3Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InstanceV3Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// InstanceV3 is the Schema for the InstanceV3s API. Manages a DDS Instance resource within OpenTelekomCloud.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,opentelekomcloud}
type InstanceV3 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.availabilityZone) || (has(self.initProvider) && has(self.initProvider.availabilityZone))",message="spec.forProvider.availabilityZone is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.datastore) || (has(self.initProvider) && has(self.initProvider.datastore))",message="spec.forProvider.datastore is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.flavor) || (has(self.initProvider) && has(self.initProvider.flavor))",message="spec.forProvider.flavor is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.mode) || (has(self.initProvider) && has(self.initProvider.mode))",message="spec.forProvider.mode is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.passwordSecretRef)",message="spec.forProvider.passwordSecretRef is a required parameter"
	Spec   InstanceV3Spec   `json:"spec"`
	Status InstanceV3Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceV3List contains a list of InstanceV3s
type InstanceV3List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InstanceV3 `json:"items"`
}

// Repository type metadata.
var (
	InstanceV3_Kind             = "InstanceV3"
	InstanceV3_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: InstanceV3_Kind}.String()
	InstanceV3_KindAPIVersion   = InstanceV3_Kind + "." + CRDGroupVersion.String()
	InstanceV3_GroupVersionKind = CRDGroupVersion.WithKind(InstanceV3_Kind)
)

func init() {
	SchemeBuilder.Register(&InstanceV3{}, &InstanceV3List{})
}
