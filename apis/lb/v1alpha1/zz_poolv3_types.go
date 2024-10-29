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

type PoolV3InitParameters struct {

	// Provides supplementary information about the backend server group.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies the load balancing algorithm used by the load balancer to route requests to backend servers.
	LBAlgorithm *string `json:"lbAlgorithm,omitempty" tf:"lb_algorithm,omitempty"`

	// Specifies the ID of the listener associated with the backend server group.
	// +crossplane:generate:reference:type=github.com/opentelekomcloud/provider-opentelekomcloud/apis/lb/v1alpha1.ListenerV3
	ListenerID *string `json:"listenerId,omitempty" tf:"listener_id,omitempty"`

	// Reference to a ListenerV3 in lb to populate listenerId.
	// +kubebuilder:validation:Optional
	ListenerIDRef *v1.Reference `json:"listenerIdRef,omitempty" tf:"-"`

	// Selector for a ListenerV3 in lb to populate listenerId.
	// +kubebuilder:validation:Optional
	ListenerIDSelector *v1.Selector `json:"listenerIdSelector,omitempty" tf:"-"`

	// Specifies the ID of the associated load balancer.
	// +crossplane:generate:reference:type=github.com/opentelekomcloud/provider-opentelekomcloud/apis/lb/v1alpha1.LoadbalancerV3
	LoadbalancerID *string `json:"loadbalancerId,omitempty" tf:"loadbalancer_id,omitempty"`

	// Reference to a LoadbalancerV3 in lb to populate loadbalancerId.
	// +kubebuilder:validation:Optional
	LoadbalancerIDRef *v1.Reference `json:"loadbalancerIdRef,omitempty" tf:"-"`

	// Selector for a LoadbalancerV3 in lb to populate loadbalancerId.
	// +kubebuilder:validation:Optional
	LoadbalancerIDSelector *v1.Selector `json:"loadbalancerIdSelector,omitempty" tf:"-"`

	// Specifies whether to enable removal protection for the pool members.
	// true: Enable removal protection.
	// false (default): Disable removal protection.
	MemberDeletionProtection *bool `json:"memberDeletionProtection,omitempty" tf:"member_deletion_protection,omitempty"`

	// Specifies the backend server group name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the project ID.
	// +crossplane:generate:reference:type=github.com/opentelekomcloud/provider-opentelekomcloud/apis/identity/v1alpha1.ProjectV3
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Reference to a ProjectV3 in identity to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDRef *v1.Reference `json:"projectIdRef,omitempty" tf:"-"`

	// Selector for a ProjectV3 in identity to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDSelector *v1.Selector `json:"projectIdSelector,omitempty" tf:"-"`

	// Specifies the protocol used by the backend server group to receive requests.
	// TCP, UDP, HTTP, HTTPS, and QUIC are supported.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Specifies whether to enable sticky sessions.
	SessionPersistence []SessionPersistenceInitParameters `json:"sessionPersistence,omitempty" tf:"session_persistence,omitempty"`

	// Specifies the sticky session type. The value can be SOURCE_IP, HTTP_COOKIE, or APP_COOKIE.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Specifies the ID of the VPC where the backend server group works.
	// +crossplane:generate:reference:type=github.com/opentelekomcloud/provider-opentelekomcloud/apis/vpc/v1alpha1.VpcV1
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VpcV1 in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VpcV1 in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`
}

type PoolV3Observation struct {

	// Provides supplementary information about the backend server group.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies the backend server group ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the IP version supported by the backend server group.
	IPVersion *string `json:"ipVersion,omitempty" tf:"ip_version,omitempty"`

	// Specifies the load balancing algorithm used by the load balancer to route requests to backend servers.
	LBAlgorithm *string `json:"lbAlgorithm,omitempty" tf:"lb_algorithm,omitempty"`

	// Specifies the ID of the listener associated with the backend server group.
	ListenerID *string `json:"listenerId,omitempty" tf:"listener_id,omitempty"`

	// Specifies the ID of the associated load balancer.
	LoadbalancerID *string `json:"loadbalancerId,omitempty" tf:"loadbalancer_id,omitempty"`

	// Specifies whether to enable removal protection for the pool members.
	// true: Enable removal protection.
	// false (default): Disable removal protection.
	MemberDeletionProtection *bool `json:"memberDeletionProtection,omitempty" tf:"member_deletion_protection,omitempty"`

	// Specifies the backend server group name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the project ID.
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Specifies the protocol used by the backend server group to receive requests.
	// TCP, UDP, HTTP, HTTPS, and QUIC are supported.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Specifies whether to enable sticky sessions.
	SessionPersistence []SessionPersistenceObservation `json:"sessionPersistence,omitempty" tf:"session_persistence,omitempty"`

	// Specifies the sticky session type. The value can be SOURCE_IP, HTTP_COOKIE, or APP_COOKIE.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Specifies the ID of the VPC where the backend server group works.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

type PoolV3Parameters struct {

	// Provides supplementary information about the backend server group.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies the load balancing algorithm used by the load balancer to route requests to backend servers.
	// +kubebuilder:validation:Optional
	LBAlgorithm *string `json:"lbAlgorithm,omitempty" tf:"lb_algorithm,omitempty"`

	// Specifies the ID of the listener associated with the backend server group.
	// +crossplane:generate:reference:type=github.com/opentelekomcloud/provider-opentelekomcloud/apis/lb/v1alpha1.ListenerV3
	// +kubebuilder:validation:Optional
	ListenerID *string `json:"listenerId,omitempty" tf:"listener_id,omitempty"`

	// Reference to a ListenerV3 in lb to populate listenerId.
	// +kubebuilder:validation:Optional
	ListenerIDRef *v1.Reference `json:"listenerIdRef,omitempty" tf:"-"`

	// Selector for a ListenerV3 in lb to populate listenerId.
	// +kubebuilder:validation:Optional
	ListenerIDSelector *v1.Selector `json:"listenerIdSelector,omitempty" tf:"-"`

	// Specifies the ID of the associated load balancer.
	// +crossplane:generate:reference:type=github.com/opentelekomcloud/provider-opentelekomcloud/apis/lb/v1alpha1.LoadbalancerV3
	// +kubebuilder:validation:Optional
	LoadbalancerID *string `json:"loadbalancerId,omitempty" tf:"loadbalancer_id,omitempty"`

	// Reference to a LoadbalancerV3 in lb to populate loadbalancerId.
	// +kubebuilder:validation:Optional
	LoadbalancerIDRef *v1.Reference `json:"loadbalancerIdRef,omitempty" tf:"-"`

	// Selector for a LoadbalancerV3 in lb to populate loadbalancerId.
	// +kubebuilder:validation:Optional
	LoadbalancerIDSelector *v1.Selector `json:"loadbalancerIdSelector,omitempty" tf:"-"`

	// Specifies whether to enable removal protection for the pool members.
	// true: Enable removal protection.
	// false (default): Disable removal protection.
	// +kubebuilder:validation:Optional
	MemberDeletionProtection *bool `json:"memberDeletionProtection,omitempty" tf:"member_deletion_protection,omitempty"`

	// Specifies the backend server group name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the project ID.
	// +crossplane:generate:reference:type=github.com/opentelekomcloud/provider-opentelekomcloud/apis/identity/v1alpha1.ProjectV3
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Reference to a ProjectV3 in identity to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDRef *v1.Reference `json:"projectIdRef,omitempty" tf:"-"`

	// Selector for a ProjectV3 in identity to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDSelector *v1.Selector `json:"projectIdSelector,omitempty" tf:"-"`

	// Specifies the protocol used by the backend server group to receive requests.
	// TCP, UDP, HTTP, HTTPS, and QUIC are supported.
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Specifies whether to enable sticky sessions.
	// +kubebuilder:validation:Optional
	SessionPersistence []SessionPersistenceParameters `json:"sessionPersistence,omitempty" tf:"session_persistence,omitempty"`

	// Specifies the sticky session type. The value can be SOURCE_IP, HTTP_COOKIE, or APP_COOKIE.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Specifies the ID of the VPC where the backend server group works.
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

type SessionPersistenceInitParameters struct {

	// Specifies the cookie name. This parameter will take effect only when type is set to APP_COOKIE.
	// The value can contain only letters, digits, hyphens (-), underscores (_), and periods (.).
	CookieName *string `json:"cookieName,omitempty" tf:"cookie_name,omitempty"`

	// Specifies the stickiness duration, in minutes.
	// This parameter will not take effect when type is set to APP_COOKIE.
	PersistenceTimeout *float64 `json:"persistenceTimeout,omitempty" tf:"persistence_timeout,omitempty"`

	// Specifies the sticky session type. The value can be SOURCE_IP, HTTP_COOKIE, or APP_COOKIE.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type SessionPersistenceObservation struct {

	// Specifies the cookie name. This parameter will take effect only when type is set to APP_COOKIE.
	// The value can contain only letters, digits, hyphens (-), underscores (_), and periods (.).
	CookieName *string `json:"cookieName,omitempty" tf:"cookie_name,omitempty"`

	// Specifies the stickiness duration, in minutes.
	// This parameter will not take effect when type is set to APP_COOKIE.
	PersistenceTimeout *float64 `json:"persistenceTimeout,omitempty" tf:"persistence_timeout,omitempty"`

	// Specifies the sticky session type. The value can be SOURCE_IP, HTTP_COOKIE, or APP_COOKIE.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type SessionPersistenceParameters struct {

	// Specifies the cookie name. This parameter will take effect only when type is set to APP_COOKIE.
	// The value can contain only letters, digits, hyphens (-), underscores (_), and periods (.).
	// +kubebuilder:validation:Optional
	CookieName *string `json:"cookieName,omitempty" tf:"cookie_name,omitempty"`

	// Specifies the stickiness duration, in minutes.
	// This parameter will not take effect when type is set to APP_COOKIE.
	// +kubebuilder:validation:Optional
	PersistenceTimeout *float64 `json:"persistenceTimeout,omitempty" tf:"persistence_timeout,omitempty"`

	// Specifies the sticky session type. The value can be SOURCE_IP, HTTP_COOKIE, or APP_COOKIE.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

// PoolV3Spec defines the desired state of PoolV3
type PoolV3Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PoolV3Parameters `json:"forProvider"`
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
	InitProvider PoolV3InitParameters `json:"initProvider,omitempty"`
}

// PoolV3Status defines the observed state of PoolV3.
type PoolV3Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PoolV3Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// PoolV3 is the Schema for the PoolV3s API. Manages a LB Pool resource within OpenTelekomCloud.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,opentelekomcloud}
type PoolV3 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.lbAlgorithm) || (has(self.initProvider) && has(self.initProvider.lbAlgorithm))",message="spec.forProvider.lbAlgorithm is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.protocol) || (has(self.initProvider) && has(self.initProvider.protocol))",message="spec.forProvider.protocol is a required parameter"
	Spec   PoolV3Spec   `json:"spec"`
	Status PoolV3Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PoolV3List contains a list of PoolV3s
type PoolV3List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PoolV3 `json:"items"`
}

// Repository type metadata.
var (
	PoolV3_Kind             = "PoolV3"
	PoolV3_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PoolV3_Kind}.String()
	PoolV3_KindAPIVersion   = PoolV3_Kind + "." + CRDGroupVersion.String()
	PoolV3_GroupVersionKind = CRDGroupVersion.WithKind(PoolV3_Kind)
)

func init() {
	SchemeBuilder.Register(&PoolV3{}, &PoolV3List{})
}
