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

type TransferV2InitParameters struct {

	// Indicates a custom path to store the log files.
	DirPrefixName *string `json:"dirPrefixName,omitempty" tf:"dir_prefix_name,omitempty"`

	// Specifies the ID of a log transfer.
	LogGroupID *string `json:"logGroupId,omitempty" tf:"log_group_id,omitempty"`

	// Specifies the log topics(streams) ids.
	// +listType=set
	LogStreamIds []*string `json:"logStreamIds,omitempty" tf:"log_stream_ids,omitempty"`

	// Specifies the name of an OBS bucket.
	ObsBucketName *string `json:"obsBucketName,omitempty" tf:"obs_bucket_name,omitempty"`

	// Indicates the length of the log transfer interval.
	// Possible values: 1, 2, 3, 5, 6, 12, and 30.
	Period *float64 `json:"period,omitempty" tf:"period,omitempty"`

	// Indicates the unit of the log transfer interval.
	// Possible values: min, hour.
	PeriodUnit *string `json:"periodUnit,omitempty" tf:"period_unit,omitempty"`

	// Indicates the file name prefix of the log files transferred to an OBS bucket.
	PrefixName *string `json:"prefixName,omitempty" tf:"prefix_name,omitempty"`

	// Indicates storage format for logs. Possible values are: RAW, JSON.
	StorageFormat *string `json:"storageFormat,omitempty" tf:"storage_format,omitempty"`

	// Indicates whether the log transfer is enabled. Default: true.
	SwitchOn *bool `json:"switchOn,omitempty" tf:"switch_on,omitempty"`
}

type TransferV2Observation struct {

	// Specifies the time when a log transfer was created.
	CreateTime *float64 `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// Indicates a custom path to store the log files.
	DirPrefixName *string `json:"dirPrefixName,omitempty" tf:"dir_prefix_name,omitempty"`

	// The log transfer ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the ID of a log transfer.
	LogGroupID *string `json:"logGroupId,omitempty" tf:"log_group_id,omitempty"`

	// The name of log group.
	LogGroupName *string `json:"logGroupName,omitempty" tf:"log_group_name,omitempty"`

	// Specifies the log topics(streams) ids.
	// +listType=set
	LogStreamIds []*string `json:"logStreamIds,omitempty" tf:"log_stream_ids,omitempty"`

	// The log transfer mode. cycle indicates periodical transfer.
	LogTransferMode *string `json:"logTransferMode,omitempty" tf:"log_transfer_mode,omitempty"`

	// The log transfer type.
	LogTransferType *string `json:"logTransferType,omitempty" tf:"log_transfer_type,omitempty"`

	// Specifies the name of an OBS bucket.
	ObsBucketName *string `json:"obsBucketName,omitempty" tf:"obs_bucket_name,omitempty"`

	// Specifies whether OBS bucket encryption is enabled.
	ObsEncryptionEnable *bool `json:"obsEncryptionEnable,omitempty" tf:"obs_encryption_enable,omitempty"`

	// Specifies the KMS key ID for an OBS transfer task.
	ObsEncryptionID *string `json:"obsEncryptionId,omitempty" tf:"obs_encryption_id,omitempty"`

	// Indicates the length of the log transfer interval.
	// Possible values: 1, 2, 3, 5, 6, 12, and 30.
	Period *float64 `json:"period,omitempty" tf:"period,omitempty"`

	// Indicates the unit of the log transfer interval.
	// Possible values: min, hour.
	PeriodUnit *string `json:"periodUnit,omitempty" tf:"period_unit,omitempty"`

	// Indicates the file name prefix of the log files transferred to an OBS bucket.
	PrefixName *string `json:"prefixName,omitempty" tf:"prefix_name,omitempty"`

	// The log transfer status.
	// ENABLE/DISABLE indicates that log transfer is enabled/disabled.
	// EXCEPTION indicates that log transfer is abnormal.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Indicates storage format for logs. Possible values are: RAW, JSON.
	StorageFormat *string `json:"storageFormat,omitempty" tf:"storage_format,omitempty"`

	// Indicates whether the log transfer is enabled. Default: true.
	SwitchOn *bool `json:"switchOn,omitempty" tf:"switch_on,omitempty"`
}

type TransferV2Parameters struct {

	// Indicates a custom path to store the log files.
	// +kubebuilder:validation:Optional
	DirPrefixName *string `json:"dirPrefixName,omitempty" tf:"dir_prefix_name,omitempty"`

	// Specifies the ID of a log transfer.
	// +kubebuilder:validation:Optional
	LogGroupID *string `json:"logGroupId,omitempty" tf:"log_group_id,omitempty"`

	// Specifies the log topics(streams) ids.
	// +kubebuilder:validation:Optional
	// +listType=set
	LogStreamIds []*string `json:"logStreamIds,omitempty" tf:"log_stream_ids,omitempty"`

	// Specifies the name of an OBS bucket.
	// +kubebuilder:validation:Optional
	ObsBucketName *string `json:"obsBucketName,omitempty" tf:"obs_bucket_name,omitempty"`

	// Indicates the length of the log transfer interval.
	// Possible values: 1, 2, 3, 5, 6, 12, and 30.
	// +kubebuilder:validation:Optional
	Period *float64 `json:"period,omitempty" tf:"period,omitempty"`

	// Indicates the unit of the log transfer interval.
	// Possible values: min, hour.
	// +kubebuilder:validation:Optional
	PeriodUnit *string `json:"periodUnit,omitempty" tf:"period_unit,omitempty"`

	// Indicates the file name prefix of the log files transferred to an OBS bucket.
	// +kubebuilder:validation:Optional
	PrefixName *string `json:"prefixName,omitempty" tf:"prefix_name,omitempty"`

	// Indicates storage format for logs. Possible values are: RAW, JSON.
	// +kubebuilder:validation:Optional
	StorageFormat *string `json:"storageFormat,omitempty" tf:"storage_format,omitempty"`

	// Indicates whether the log transfer is enabled. Default: true.
	// +kubebuilder:validation:Optional
	SwitchOn *bool `json:"switchOn,omitempty" tf:"switch_on,omitempty"`
}

// TransferV2Spec defines the desired state of TransferV2
type TransferV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TransferV2Parameters `json:"forProvider"`
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
	InitProvider TransferV2InitParameters `json:"initProvider,omitempty"`
}

// TransferV2Status defines the observed state of TransferV2.
type TransferV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TransferV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// TransferV2 is the Schema for the TransferV2s API. Manages a LTS Log Transfer resource within OpenTelekomCloud.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,opentelekomcloud}
type TransferV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.logGroupId) || (has(self.initProvider) && has(self.initProvider.logGroupId))",message="spec.forProvider.logGroupId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.logStreamIds) || (has(self.initProvider) && has(self.initProvider.logStreamIds))",message="spec.forProvider.logStreamIds is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.obsBucketName) || (has(self.initProvider) && has(self.initProvider.obsBucketName))",message="spec.forProvider.obsBucketName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.period) || (has(self.initProvider) && has(self.initProvider.period))",message="spec.forProvider.period is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.periodUnit) || (has(self.initProvider) && has(self.initProvider.periodUnit))",message="spec.forProvider.periodUnit is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.storageFormat) || (has(self.initProvider) && has(self.initProvider.storageFormat))",message="spec.forProvider.storageFormat is a required parameter"
	Spec   TransferV2Spec   `json:"spec"`
	Status TransferV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TransferV2List contains a list of TransferV2s
type TransferV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TransferV2 `json:"items"`
}

// Repository type metadata.
var (
	TransferV2_Kind             = "TransferV2"
	TransferV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TransferV2_Kind}.String()
	TransferV2_KindAPIVersion   = TransferV2_Kind + "." + CRDGroupVersion.String()
	TransferV2_GroupVersionKind = CRDGroupVersion.WithKind(TransferV2_Kind)
)

func init() {
	SchemeBuilder.Register(&TransferV2{}, &TransferV2List{})
}
