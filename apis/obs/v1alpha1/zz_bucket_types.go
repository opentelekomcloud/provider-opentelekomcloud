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

type BucketInitParameters struct {

	// Specifies the ACL policy for a bucket. The predefined common policies are as follows:
	// private, public-read, public-read-write and log-delivery-write. Defaults to private.
	ACL *string `json:"acl,omitempty" tf:"acl,omitempty"`

	// Specifies the name of the bucket. Changing this parameter will create a new resource.
	// A bucket must be named according to the globally applied DNS naming regulations as follows:
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// A rule of Cross-Origin Resource Sharing (documented below).
	CorsRule []CorsRuleInitParameters `json:"corsRule,omitempty" tf:"cors_rule,omitempty"`

	// A configuration of bucket event notifications (documented below).
	EventNotifications []EventNotificationsInitParameters `json:"eventNotifications,omitempty" tf:"event_notifications,omitempty"`

	// A boolean that indicates all objects should be deleted from the bucket so that the
	// bucket can be destroyed without error. Default to false.
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// A configuration of object lifecycle management (documented below).
	LifecycleRule []LifecycleRuleInitParameters `json:"lifecycleRule,omitempty" tf:"lifecycle_rule,omitempty"`

	// A settings of bucket logging (documented below).
	Logging []LoggingInitParameters `json:"logging,omitempty" tf:"logging,omitempty"`

	// Whether enable a bucket as a parallel file system.
	ParallelFs *bool `json:"parallelFs,omitempty" tf:"parallel_fs,omitempty"`

	// If specified, the region this bucket should reside in. Otherwise,
	// the region used by the provider.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// A configuration of server side encryption (documented below).
	ServerSideEncryption []ServerSideEncryptionInitParameters `json:"serverSideEncryption,omitempty" tf:"server_side_encryption,omitempty"`

	// Specifies the storage class of the bucket. OBS provides three storage classes:
	// STANDARD, WARM (Infrequent Access) and COLD (Archive). Defaults to STANDARD.
	StorageClass *string `json:"storageClass,omitempty" tf:"storage_class,omitempty"`

	// A mapping of tags to assign to the bucket. Each tag is represented by one key-value pair.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the user domain names. The restriction requirements for this field
	// are as follows:
	UserDomainNames []*string `json:"userDomainNames,omitempty" tf:"user_domain_names,omitempty"`

	// Set to true to enable versioning. Once you version-enable a bucket, it can never return to an
	// unversioned state. You can, however, suspend versioning on that bucket. If omitted, during bucket
	// creation it will be in Disabled state.
	Versioning *bool `json:"versioning,omitempty" tf:"versioning,omitempty"`

	// A website object (documented below).
	Website []WebsiteInitParameters `json:"website,omitempty" tf:"website,omitempty"`

	// A settings of bucket default WORM policy and a retention period (documented below).
	// worm_policy requires versioning to be enabled.
	WormPolicy []WormPolicyInitParameters `json:"wormPolicy,omitempty" tf:"worm_policy,omitempty"`
}

type BucketObservation struct {

	// Specifies the ACL policy for a bucket. The predefined common policies are as follows:
	// private, public-read, public-read-write and log-delivery-write. Defaults to private.
	ACL *string `json:"acl,omitempty" tf:"acl,omitempty"`

	// Specifies the name of the bucket. Changing this parameter will create a new resource.
	// A bucket must be named according to the globally applied DNS naming regulations as follows:
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// The bucket domain name. Will be of format bucketname.obs.region.otc.t-systems.com.
	BucketDomainName *string `json:"bucketDomainName,omitempty" tf:"bucket_domain_name,omitempty"`

	// The OBS version of the bucket.
	BucketVersion *string `json:"bucketVersion,omitempty" tf:"bucket_version,omitempty"`

	// A rule of Cross-Origin Resource Sharing (documented below).
	CorsRule []CorsRuleObservation `json:"corsRule,omitempty" tf:"cors_rule,omitempty"`

	// A configuration of bucket event notifications (documented below).
	EventNotifications []EventNotificationsObservation `json:"eventNotifications,omitempty" tf:"event_notifications,omitempty"`

	// A boolean that indicates all objects should be deleted from the bucket so that the
	// bucket can be destroyed without error. Default to false.
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// Unique ID of the event notification. If the user does not specify an ID, the system assigns an ID automatically.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A configuration of object lifecycle management (documented below).
	LifecycleRule []LifecycleRuleObservation `json:"lifecycleRule,omitempty" tf:"lifecycle_rule,omitempty"`

	// A settings of bucket logging (documented below).
	Logging []LoggingObservation `json:"logging,omitempty" tf:"logging,omitempty"`

	// Whether enable a bucket as a parallel file system.
	ParallelFs *bool `json:"parallelFs,omitempty" tf:"parallel_fs,omitempty"`

	// If specified, the region this bucket should reside in. Otherwise,
	// the region used by the provider.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// A configuration of server side encryption (documented below).
	ServerSideEncryption []ServerSideEncryptionObservation `json:"serverSideEncryption,omitempty" tf:"server_side_encryption,omitempty"`

	// Specifies the storage class of the bucket. OBS provides three storage classes:
	// STANDARD, WARM (Infrequent Access) and COLD (Archive). Defaults to STANDARD.
	StorageClass *string `json:"storageClass,omitempty" tf:"storage_class,omitempty"`

	// A mapping of tags to assign to the bucket. Each tag is represented by one key-value pair.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the user domain names. The restriction requirements for this field
	// are as follows:
	UserDomainNames []*string `json:"userDomainNames,omitempty" tf:"user_domain_names,omitempty"`

	// Set to true to enable versioning. Once you version-enable a bucket, it can never return to an
	// unversioned state. You can, however, suspend versioning on that bucket. If omitted, during bucket
	// creation it will be in Disabled state.
	Versioning *bool `json:"versioning,omitempty" tf:"versioning,omitempty"`

	// A website object (documented below).
	Website []WebsiteObservation `json:"website,omitempty" tf:"website,omitempty"`

	// A settings of bucket default WORM policy and a retention period (documented below).
	// worm_policy requires versioning to be enabled.
	WormPolicy []WormPolicyObservation `json:"wormPolicy,omitempty" tf:"worm_policy,omitempty"`
}

type BucketParameters struct {

	// Specifies the ACL policy for a bucket. The predefined common policies are as follows:
	// private, public-read, public-read-write and log-delivery-write. Defaults to private.
	// +kubebuilder:validation:Optional
	ACL *string `json:"acl,omitempty" tf:"acl,omitempty"`

	// Specifies the name of the bucket. Changing this parameter will create a new resource.
	// A bucket must be named according to the globally applied DNS naming regulations as follows:
	// +kubebuilder:validation:Optional
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// A rule of Cross-Origin Resource Sharing (documented below).
	// +kubebuilder:validation:Optional
	CorsRule []CorsRuleParameters `json:"corsRule,omitempty" tf:"cors_rule,omitempty"`

	// A configuration of bucket event notifications (documented below).
	// +kubebuilder:validation:Optional
	EventNotifications []EventNotificationsParameters `json:"eventNotifications,omitempty" tf:"event_notifications,omitempty"`

	// A boolean that indicates all objects should be deleted from the bucket so that the
	// bucket can be destroyed without error. Default to false.
	// +kubebuilder:validation:Optional
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// A configuration of object lifecycle management (documented below).
	// +kubebuilder:validation:Optional
	LifecycleRule []LifecycleRuleParameters `json:"lifecycleRule,omitempty" tf:"lifecycle_rule,omitempty"`

	// A settings of bucket logging (documented below).
	// +kubebuilder:validation:Optional
	Logging []LoggingParameters `json:"logging,omitempty" tf:"logging,omitempty"`

	// Whether enable a bucket as a parallel file system.
	// +kubebuilder:validation:Optional
	ParallelFs *bool `json:"parallelFs,omitempty" tf:"parallel_fs,omitempty"`

	// If specified, the region this bucket should reside in. Otherwise,
	// the region used by the provider.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// A configuration of server side encryption (documented below).
	// +kubebuilder:validation:Optional
	ServerSideEncryption []ServerSideEncryptionParameters `json:"serverSideEncryption,omitempty" tf:"server_side_encryption,omitempty"`

	// Specifies the storage class of the bucket. OBS provides three storage classes:
	// STANDARD, WARM (Infrequent Access) and COLD (Archive). Defaults to STANDARD.
	// +kubebuilder:validation:Optional
	StorageClass *string `json:"storageClass,omitempty" tf:"storage_class,omitempty"`

	// A mapping of tags to assign to the bucket. Each tag is represented by one key-value pair.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the user domain names. The restriction requirements for this field
	// are as follows:
	// +kubebuilder:validation:Optional
	UserDomainNames []*string `json:"userDomainNames,omitempty" tf:"user_domain_names,omitempty"`

	// Set to true to enable versioning. Once you version-enable a bucket, it can never return to an
	// unversioned state. You can, however, suspend versioning on that bucket. If omitted, during bucket
	// creation it will be in Disabled state.
	// +kubebuilder:validation:Optional
	Versioning *bool `json:"versioning,omitempty" tf:"versioning,omitempty"`

	// A website object (documented below).
	// +kubebuilder:validation:Optional
	Website []WebsiteParameters `json:"website,omitempty" tf:"website,omitempty"`

	// A settings of bucket default WORM policy and a retention period (documented below).
	// worm_policy requires versioning to be enabled.
	// +kubebuilder:validation:Optional
	WormPolicy []WormPolicyParameters `json:"wormPolicy,omitempty" tf:"worm_policy,omitempty"`
}

type CorsRuleInitParameters struct {

	// Specifies the allowed header of cross-origin requests.
	// Only CORS requests matching the allowed header are valid.
	AllowedHeaders []*string `json:"allowedHeaders,omitempty" tf:"allowed_headers,omitempty"`

	// Specifies the acceptable operation type of buckets and objects.
	// The methods include GET, PUT, POST, DELETE or HEAD.
	AllowedMethods []*string `json:"allowedMethods,omitempty" tf:"allowed_methods,omitempty"`

	// Requests from this origin can access the bucket. Multiple matching rules are allowed.
	// One rule occupies one line, and allows one wildcard character (*) at most.
	AllowedOrigins []*string `json:"allowedOrigins,omitempty" tf:"allowed_origins,omitempty"`

	// Specifies the exposed header in CORS responses, providing additional
	// information for clients.
	ExposeHeaders []*string `json:"exposeHeaders,omitempty" tf:"expose_headers,omitempty"`

	// Specifies the duration that your browser can cache CORS responses,
	// expressed in seconds. The default value is 100.
	MaxAgeSeconds *float64 `json:"maxAgeSeconds,omitempty" tf:"max_age_seconds,omitempty"`
}

type CorsRuleObservation struct {

	// Specifies the allowed header of cross-origin requests.
	// Only CORS requests matching the allowed header are valid.
	AllowedHeaders []*string `json:"allowedHeaders,omitempty" tf:"allowed_headers,omitempty"`

	// Specifies the acceptable operation type of buckets and objects.
	// The methods include GET, PUT, POST, DELETE or HEAD.
	AllowedMethods []*string `json:"allowedMethods,omitempty" tf:"allowed_methods,omitempty"`

	// Requests from this origin can access the bucket. Multiple matching rules are allowed.
	// One rule occupies one line, and allows one wildcard character (*) at most.
	AllowedOrigins []*string `json:"allowedOrigins,omitempty" tf:"allowed_origins,omitempty"`

	// Specifies the exposed header in CORS responses, providing additional
	// information for clients.
	ExposeHeaders []*string `json:"exposeHeaders,omitempty" tf:"expose_headers,omitempty"`

	// Specifies the duration that your browser can cache CORS responses,
	// expressed in seconds. The default value is 100.
	MaxAgeSeconds *float64 `json:"maxAgeSeconds,omitempty" tf:"max_age_seconds,omitempty"`
}

type CorsRuleParameters struct {

	// Specifies the allowed header of cross-origin requests.
	// Only CORS requests matching the allowed header are valid.
	// +kubebuilder:validation:Optional
	AllowedHeaders []*string `json:"allowedHeaders,omitempty" tf:"allowed_headers,omitempty"`

	// Specifies the acceptable operation type of buckets and objects.
	// The methods include GET, PUT, POST, DELETE or HEAD.
	// +kubebuilder:validation:Optional
	AllowedMethods []*string `json:"allowedMethods" tf:"allowed_methods,omitempty"`

	// Requests from this origin can access the bucket. Multiple matching rules are allowed.
	// One rule occupies one line, and allows one wildcard character (*) at most.
	// +kubebuilder:validation:Optional
	AllowedOrigins []*string `json:"allowedOrigins" tf:"allowed_origins,omitempty"`

	// Specifies the exposed header in CORS responses, providing additional
	// information for clients.
	// +kubebuilder:validation:Optional
	ExposeHeaders []*string `json:"exposeHeaders,omitempty" tf:"expose_headers,omitempty"`

	// Specifies the duration that your browser can cache CORS responses,
	// expressed in seconds. The default value is 100.
	// +kubebuilder:validation:Optional
	MaxAgeSeconds *float64 `json:"maxAgeSeconds,omitempty" tf:"max_age_seconds,omitempty"`
}

type EventNotificationsInitParameters struct {

	// Type of events that need to be notified.
	Events []*string `json:"events,omitempty" tf:"events,omitempty"`

	// Filtering rules. The rules filter objects based on the prefixes and suffixes of object names.
	FilterRule []FilterRuleInitParameters `json:"filterRule,omitempty" tf:"filter_rule,omitempty"`

	// Unique ID of the event notification. If the user does not specify an ID, the system assigns an ID automatically.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// URN of the event notification topic. After detecting a specific event, OBS sends a message to the topic.
	Topic *string `json:"topic,omitempty" tf:"topic,omitempty"`
}

type EventNotificationsObservation struct {

	// Type of events that need to be notified.
	Events []*string `json:"events,omitempty" tf:"events,omitempty"`

	// Filtering rules. The rules filter objects based on the prefixes and suffixes of object names.
	FilterRule []FilterRuleObservation `json:"filterRule,omitempty" tf:"filter_rule,omitempty"`

	// Unique ID of the event notification. If the user does not specify an ID, the system assigns an ID automatically.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// URN of the event notification topic. After detecting a specific event, OBS sends a message to the topic.
	Topic *string `json:"topic,omitempty" tf:"topic,omitempty"`
}

type EventNotificationsParameters struct {

	// Type of events that need to be notified.
	// +kubebuilder:validation:Optional
	Events []*string `json:"events" tf:"events,omitempty"`

	// Filtering rules. The rules filter objects based on the prefixes and suffixes of object names.
	// +kubebuilder:validation:Optional
	FilterRule []FilterRuleParameters `json:"filterRule,omitempty" tf:"filter_rule,omitempty"`

	// Unique ID of the event notification. If the user does not specify an ID, the system assigns an ID automatically.
	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// URN of the event notification topic. After detecting a specific event, OBS sends a message to the topic.
	// +kubebuilder:validation:Optional
	Topic *string `json:"topic" tf:"topic,omitempty"`
}

type ExpirationInitParameters struct {

	// Default protection period, in days.
	// The value is from 1 to 36500.
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`
}

type ExpirationObservation struct {

	// Default protection period, in days.
	// The value is from 1 to 36500.
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`
}

type ExpirationParameters struct {

	// Default protection period, in days.
	// The value is from 1 to 36500.
	// +kubebuilder:validation:Optional
	Days *float64 `json:"days" tf:"days,omitempty"`
}

type FilterRuleInitParameters struct {

	// Unique identifier for lifecycle rules. The Rule Name contains a maximum of 255 characters.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies keywords of object names so that objects can be filtered based on the prefixes or suffixes.
	// The value contains a maximum of 1024 characters.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type FilterRuleObservation struct {

	// Unique identifier for lifecycle rules. The Rule Name contains a maximum of 255 characters.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies keywords of object names so that objects can be filtered based on the prefixes or suffixes.
	// The value contains a maximum of 1024 characters.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type FilterRuleParameters struct {

	// Unique identifier for lifecycle rules. The Rule Name contains a maximum of 255 characters.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies keywords of object names so that objects can be filtered based on the prefixes or suffixes.
	// The value contains a maximum of 1024 characters.
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type LifecycleRuleInitParameters struct {

	// Specifies lifecycle rule status.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Specifies a period when objects that have been last updated are automatically
	// deleted. (documented below).
	Expiration []ExpirationInitParameters `json:"expiration,omitempty" tf:"expiration,omitempty"`

	// Unique identifier for lifecycle rules. The Rule Name contains a maximum of 255 characters.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies a period when noncurrent object versions are
	// automatically deleted. (documented below).
	NoncurrentVersionExpiration []NoncurrentVersionExpirationInitParameters `json:"noncurrentVersionExpiration,omitempty" tf:"noncurrent_version_expiration,omitempty"`

	// Specifies a period when noncurrent object versions are
	// automatically transitioned to WARM or COLD storage class (documented below).
	NoncurrentVersionTransition []NoncurrentVersionTransitionInitParameters `json:"noncurrentVersionTransition,omitempty" tf:"noncurrent_version_transition,omitempty"`

	// Object key prefix identifying one or more objects to which the rule applies.
	// If omitted, all objects in the bucket will be managed by the lifecycle rule. The prefix cannot start
	// or end with a slash (/), cannot have consecutive slashes (/), and cannot contain the following
	// special characters: :*?"<>|.
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// Specifies a period when objects that have been last updated are automatically
	// transitioned to WARM or COLD storage class (documented below).
	Transition []TransitionInitParameters `json:"transition,omitempty" tf:"transition,omitempty"`
}

type LifecycleRuleObservation struct {

	// Specifies lifecycle rule status.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Specifies a period when objects that have been last updated are automatically
	// deleted. (documented below).
	Expiration []ExpirationObservation `json:"expiration,omitempty" tf:"expiration,omitempty"`

	// Unique identifier for lifecycle rules. The Rule Name contains a maximum of 255 characters.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies a period when noncurrent object versions are
	// automatically deleted. (documented below).
	NoncurrentVersionExpiration []NoncurrentVersionExpirationObservation `json:"noncurrentVersionExpiration,omitempty" tf:"noncurrent_version_expiration,omitempty"`

	// Specifies a period when noncurrent object versions are
	// automatically transitioned to WARM or COLD storage class (documented below).
	NoncurrentVersionTransition []NoncurrentVersionTransitionObservation `json:"noncurrentVersionTransition,omitempty" tf:"noncurrent_version_transition,omitempty"`

	// Object key prefix identifying one or more objects to which the rule applies.
	// If omitted, all objects in the bucket will be managed by the lifecycle rule. The prefix cannot start
	// or end with a slash (/), cannot have consecutive slashes (/), and cannot contain the following
	// special characters: :*?"<>|.
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// Specifies a period when objects that have been last updated are automatically
	// transitioned to WARM or COLD storage class (documented below).
	Transition []TransitionObservation `json:"transition,omitempty" tf:"transition,omitempty"`
}

type LifecycleRuleParameters struct {

	// Specifies lifecycle rule status.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`

	// Specifies a period when objects that have been last updated are automatically
	// deleted. (documented below).
	// +kubebuilder:validation:Optional
	Expiration []ExpirationParameters `json:"expiration,omitempty" tf:"expiration,omitempty"`

	// Unique identifier for lifecycle rules. The Rule Name contains a maximum of 255 characters.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// Specifies a period when noncurrent object versions are
	// automatically deleted. (documented below).
	// +kubebuilder:validation:Optional
	NoncurrentVersionExpiration []NoncurrentVersionExpirationParameters `json:"noncurrentVersionExpiration,omitempty" tf:"noncurrent_version_expiration,omitempty"`

	// Specifies a period when noncurrent object versions are
	// automatically transitioned to WARM or COLD storage class (documented below).
	// +kubebuilder:validation:Optional
	NoncurrentVersionTransition []NoncurrentVersionTransitionParameters `json:"noncurrentVersionTransition,omitempty" tf:"noncurrent_version_transition,omitempty"`

	// Object key prefix identifying one or more objects to which the rule applies.
	// If omitted, all objects in the bucket will be managed by the lifecycle rule. The prefix cannot start
	// or end with a slash (/), cannot have consecutive slashes (/), and cannot contain the following
	// special characters: :*?"<>|.
	// +kubebuilder:validation:Optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// Specifies a period when objects that have been last updated are automatically
	// transitioned to WARM or COLD storage class (documented below).
	// +kubebuilder:validation:Optional
	Transition []TransitionParameters `json:"transition,omitempty" tf:"transition,omitempty"`
}

type LoggingInitParameters struct {

	// The name of the bucket that will receive the log objects.
	// The acl policy of the target bucket should be log-delivery-write.
	TargetBucket *string `json:"targetBucket,omitempty" tf:"target_bucket,omitempty"`

	// To specify a key prefix for log objects.
	TargetPrefix *string `json:"targetPrefix,omitempty" tf:"target_prefix,omitempty"`
}

type LoggingObservation struct {

	// The name of the bucket that will receive the log objects.
	// The acl policy of the target bucket should be log-delivery-write.
	TargetBucket *string `json:"targetBucket,omitempty" tf:"target_bucket,omitempty"`

	// To specify a key prefix for log objects.
	TargetPrefix *string `json:"targetPrefix,omitempty" tf:"target_prefix,omitempty"`
}

type LoggingParameters struct {

	// The name of the bucket that will receive the log objects.
	// The acl policy of the target bucket should be log-delivery-write.
	// +kubebuilder:validation:Optional
	TargetBucket *string `json:"targetBucket" tf:"target_bucket,omitempty"`

	// To specify a key prefix for log objects.
	// +kubebuilder:validation:Optional
	TargetPrefix *string `json:"targetPrefix,omitempty" tf:"target_prefix,omitempty"`
}

type NoncurrentVersionExpirationInitParameters struct {

	// Default protection period, in days.
	// The value is from 1 to 36500.
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`
}

type NoncurrentVersionExpirationObservation struct {

	// Default protection period, in days.
	// The value is from 1 to 36500.
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`
}

type NoncurrentVersionExpirationParameters struct {

	// Default protection period, in days.
	// The value is from 1 to 36500.
	// +kubebuilder:validation:Optional
	Days *float64 `json:"days" tf:"days,omitempty"`
}

type NoncurrentVersionTransitionInitParameters struct {

	// Default protection period, in days.
	// The value is from 1 to 36500.
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`

	// Specifies the storage class of the bucket. OBS provides three storage classes:
	// STANDARD, WARM (Infrequent Access) and COLD (Archive). Defaults to STANDARD.
	StorageClass *string `json:"storageClass,omitempty" tf:"storage_class,omitempty"`
}

type NoncurrentVersionTransitionObservation struct {

	// Default protection period, in days.
	// The value is from 1 to 36500.
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`

	// Specifies the storage class of the bucket. OBS provides three storage classes:
	// STANDARD, WARM (Infrequent Access) and COLD (Archive). Defaults to STANDARD.
	StorageClass *string `json:"storageClass,omitempty" tf:"storage_class,omitempty"`
}

type NoncurrentVersionTransitionParameters struct {

	// Default protection period, in days.
	// The value is from 1 to 36500.
	// +kubebuilder:validation:Optional
	Days *float64 `json:"days" tf:"days,omitempty"`

	// Specifies the storage class of the bucket. OBS provides three storage classes:
	// STANDARD, WARM (Infrequent Access) and COLD (Archive). Defaults to STANDARD.
	// +kubebuilder:validation:Optional
	StorageClass *string `json:"storageClass" tf:"storage_class,omitempty"`
}

type ServerSideEncryptionInitParameters struct {

	// The algorithm used for SSE. Only kms is supported.
	Algorithm *string `json:"algorithm,omitempty" tf:"algorithm,omitempty"`

	// The ID of KMS key used for the encryption.
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`
}

type ServerSideEncryptionObservation struct {

	// The algorithm used for SSE. Only kms is supported.
	Algorithm *string `json:"algorithm,omitempty" tf:"algorithm,omitempty"`

	// The ID of KMS key used for the encryption.
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`
}

type ServerSideEncryptionParameters struct {

	// The algorithm used for SSE. Only kms is supported.
	// +kubebuilder:validation:Optional
	Algorithm *string `json:"algorithm" tf:"algorithm,omitempty"`

	// The ID of KMS key used for the encryption.
	// +kubebuilder:validation:Optional
	KMSKeyID *string `json:"kmsKeyId" tf:"kms_key_id,omitempty"`
}

type TransitionInitParameters struct {

	// Default protection period, in days.
	// The value is from 1 to 36500.
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`

	// Specifies the storage class of the bucket. OBS provides three storage classes:
	// STANDARD, WARM (Infrequent Access) and COLD (Archive). Defaults to STANDARD.
	StorageClass *string `json:"storageClass,omitempty" tf:"storage_class,omitempty"`
}

type TransitionObservation struct {

	// Default protection period, in days.
	// The value is from 1 to 36500.
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`

	// Specifies the storage class of the bucket. OBS provides three storage classes:
	// STANDARD, WARM (Infrequent Access) and COLD (Archive). Defaults to STANDARD.
	StorageClass *string `json:"storageClass,omitempty" tf:"storage_class,omitempty"`
}

type TransitionParameters struct {

	// Default protection period, in days.
	// The value is from 1 to 36500.
	// +kubebuilder:validation:Optional
	Days *float64 `json:"days" tf:"days,omitempty"`

	// Specifies the storage class of the bucket. OBS provides three storage classes:
	// STANDARD, WARM (Infrequent Access) and COLD (Archive). Defaults to STANDARD.
	// +kubebuilder:validation:Optional
	StorageClass *string `json:"storageClass" tf:"storage_class,omitempty"`
}

type WebsiteInitParameters struct {

	// Specifies the error page returned when an error occurs during static website access.
	// Only HTML, JPG, PNG, BMP, and WEBP files under the root directory are supported.
	ErrorDocument *string `json:"errorDocument,omitempty" tf:"error_document,omitempty"`

	// Specifies the default homepage of the
	// static website, only HTML web pages are supported. OBS only allows files such as index.html in the root
	// directory of a bucket to function as the default homepage. That is to say, do not set the default homepage
	// with a multi-level directory structure (for example, /page/index.html).
	IndexDocument *string `json:"indexDocument,omitempty" tf:"index_document,omitempty"`

	// A hostname to redirect all website requests for this bucket to.
	// Hostname can optionally be prefixed with a protocol (http:// or https://) to use when redirecting
	// requests. The default is the protocol that is used in the original request.
	RedirectAllRequestsTo *string `json:"redirectAllRequestsTo,omitempty" tf:"redirect_all_requests_to,omitempty"`

	// A JSON or XML format containing routing rules describing redirect
	// behavior and when redirects are applied. Each rule contains a Condition and a Redirect
	// as shown in the following table:
	RoutingRules *string `json:"routingRules,omitempty" tf:"routing_rules,omitempty"`
}

type WebsiteObservation struct {

	// Specifies the error page returned when an error occurs during static website access.
	// Only HTML, JPG, PNG, BMP, and WEBP files under the root directory are supported.
	ErrorDocument *string `json:"errorDocument,omitempty" tf:"error_document,omitempty"`

	// Specifies the default homepage of the
	// static website, only HTML web pages are supported. OBS only allows files such as index.html in the root
	// directory of a bucket to function as the default homepage. That is to say, do not set the default homepage
	// with a multi-level directory structure (for example, /page/index.html).
	IndexDocument *string `json:"indexDocument,omitempty" tf:"index_document,omitempty"`

	// A hostname to redirect all website requests for this bucket to.
	// Hostname can optionally be prefixed with a protocol (http:// or https://) to use when redirecting
	// requests. The default is the protocol that is used in the original request.
	RedirectAllRequestsTo *string `json:"redirectAllRequestsTo,omitempty" tf:"redirect_all_requests_to,omitempty"`

	// A JSON or XML format containing routing rules describing redirect
	// behavior and when redirects are applied. Each rule contains a Condition and a Redirect
	// as shown in the following table:
	RoutingRules *string `json:"routingRules,omitempty" tf:"routing_rules,omitempty"`
}

type WebsiteParameters struct {

	// Specifies the error page returned when an error occurs during static website access.
	// Only HTML, JPG, PNG, BMP, and WEBP files under the root directory are supported.
	// +kubebuilder:validation:Optional
	ErrorDocument *string `json:"errorDocument,omitempty" tf:"error_document,omitempty"`

	// Specifies the default homepage of the
	// static website, only HTML web pages are supported. OBS only allows files such as index.html in the root
	// directory of a bucket to function as the default homepage. That is to say, do not set the default homepage
	// with a multi-level directory structure (for example, /page/index.html).
	// +kubebuilder:validation:Optional
	IndexDocument *string `json:"indexDocument,omitempty" tf:"index_document,omitempty"`

	// A hostname to redirect all website requests for this bucket to.
	// Hostname can optionally be prefixed with a protocol (http:// or https://) to use when redirecting
	// requests. The default is the protocol that is used in the original request.
	// +kubebuilder:validation:Optional
	RedirectAllRequestsTo *string `json:"redirectAllRequestsTo,omitempty" tf:"redirect_all_requests_to,omitempty"`

	// A JSON or XML format containing routing rules describing redirect
	// behavior and when redirects are applied. Each rule contains a Condition and a Redirect
	// as shown in the following table:
	// +kubebuilder:validation:Optional
	RoutingRules *string `json:"routingRules,omitempty" tf:"routing_rules,omitempty"`
}

type WormPolicyInitParameters struct {

	// Default protection period, in days.
	// The value is from 1 to 36500.
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`

	// Default protection period, in years. In a leap year, only 365 days are calculated.
	// The value is from 1 to 100.
	Years *float64 `json:"years,omitempty" tf:"years,omitempty"`
}

type WormPolicyObservation struct {

	// Default protection period, in days.
	// The value is from 1 to 36500.
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`

	// Default protection period, in years. In a leap year, only 365 days are calculated.
	// The value is from 1 to 100.
	Years *float64 `json:"years,omitempty" tf:"years,omitempty"`
}

type WormPolicyParameters struct {

	// Default protection period, in days.
	// The value is from 1 to 36500.
	// +kubebuilder:validation:Optional
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`

	// Default protection period, in years. In a leap year, only 365 days are calculated.
	// The value is from 1 to 100.
	// +kubebuilder:validation:Optional
	Years *float64 `json:"years,omitempty" tf:"years,omitempty"`
}

// BucketSpec defines the desired state of Bucket
type BucketSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BucketParameters `json:"forProvider"`
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
	InitProvider BucketInitParameters `json:"initProvider,omitempty"`
}

// BucketStatus defines the observed state of Bucket.
type BucketStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BucketObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Bucket is the Schema for the Buckets API. Manages a OBS Bucket resource within OpenTelekomCloud.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,opentelekomcloud}
type Bucket struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.bucket) || (has(self.initProvider) && has(self.initProvider.bucket))",message="spec.forProvider.bucket is a required parameter"
	Spec   BucketSpec   `json:"spec"`
	Status BucketStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BucketList contains a list of Buckets
type BucketList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Bucket `json:"items"`
}

// Repository type metadata.
var (
	Bucket_Kind             = "Bucket"
	Bucket_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Bucket_Kind}.String()
	Bucket_KindAPIVersion   = Bucket_Kind + "." + CRDGroupVersion.String()
	Bucket_GroupVersionKind = CRDGroupVersion.WithKind(Bucket_Kind)
)

func init() {
	SchemeBuilder.Register(&Bucket{}, &BucketList{})
}