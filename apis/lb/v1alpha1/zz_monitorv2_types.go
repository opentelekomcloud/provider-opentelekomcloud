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

type MonitorV2InitParameters struct {

	// The administrative state of the monitor.
	// A valid value is true (UP) or false (DOWN).
	AdminStateUp *bool `json:"adminStateUp,omitempty" tf:"admin_state_up,omitempty"`

	// The time, in seconds, between sending probes to members.
	Delay *float64 `json:"delay,omitempty" tf:"delay,omitempty"`

	// The domain_name of the HTTP request during the health check.
	DomainName *string `json:"domainName,omitempty" tf:"domain_name,omitempty"`

	// Required for HTTP types. Expected HTTP codes
	// for a passing HTTP monitor. You can either specify a single status like
	// "200", or a list like "200,202".
	ExpectedCodes *string `json:"expectedCodes,omitempty" tf:"expected_codes,omitempty"`

	// Required for HTTP types. The HTTP method used
	// for requests by the monitor. If this attribute is not specified, it
	// defaults to GET. The value can be GET, HEAD, POST, PUT, DELETE,
	// TRACE, OPTIONS, CONNECT, and PATCH.
	HTTPMethod *string `json:"httpMethod,omitempty" tf:"http_method,omitempty"`

	// Number of permissible ping failures before
	// changing the member's status to INACTIVE. Must be a number between 1 and 10.
	MaxRetries *float64 `json:"maxRetries,omitempty" tf:"max_retries,omitempty"`

	// Specifies the health check port. The port number
	// ranges from 1 to 65535. The value is left blank by default, indicating that
	// the port of the backend server is used as the health check port.
	MonitorPort *float64 `json:"monitorPort,omitempty" tf:"monitor_port,omitempty"`

	// The Name of the Monitor.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The id of the pool that this monitor will be assigned to.
	PoolID *string `json:"poolId,omitempty" tf:"pool_id,omitempty"`

	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Required for admins. The UUID of the tenant who owns
	// the monitor. Only administrative users can specify a tenant UUID
	// other than their own. Changing this creates a new monitor.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// Maximum number of seconds for a monitor to wait for a
	// ping reply before it times out. The value must be less than the delay value.
	Timeout *float64 `json:"timeout,omitempty" tf:"timeout,omitempty"`

	// The type of probe, which is TCP, UDP_CONNECT, or HTTP,
	// that is sent by the load balancer to verify the member state. Changing this
	// creates a new monitor.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Required for HTTP types. URI path that will be
	// accessed if monitor type is HTTP.
	URLPath *string `json:"urlPath,omitempty" tf:"url_path,omitempty"`
}

type MonitorV2Observation struct {

	// The administrative state of the monitor.
	// A valid value is true (UP) or false (DOWN).
	AdminStateUp *bool `json:"adminStateUp,omitempty" tf:"admin_state_up,omitempty"`

	// The time, in seconds, between sending probes to members.
	Delay *float64 `json:"delay,omitempty" tf:"delay,omitempty"`

	// The domain_name of the HTTP request during the health check.
	DomainName *string `json:"domainName,omitempty" tf:"domain_name,omitempty"`

	// Required for HTTP types. Expected HTTP codes
	// for a passing HTTP monitor. You can either specify a single status like
	// "200", or a list like "200,202".
	ExpectedCodes *string `json:"expectedCodes,omitempty" tf:"expected_codes,omitempty"`

	// Required for HTTP types. The HTTP method used
	// for requests by the monitor. If this attribute is not specified, it
	// defaults to GET. The value can be GET, HEAD, POST, PUT, DELETE,
	// TRACE, OPTIONS, CONNECT, and PATCH.
	HTTPMethod *string `json:"httpMethod,omitempty" tf:"http_method,omitempty"`

	// The unique ID for the monitor.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Number of permissible ping failures before
	// changing the member's status to INACTIVE. Must be a number between 1 and 10.
	MaxRetries *float64 `json:"maxRetries,omitempty" tf:"max_retries,omitempty"`

	// Specifies the health check port. The port number
	// ranges from 1 to 65535. The value is left blank by default, indicating that
	// the port of the backend server is used as the health check port.
	MonitorPort *float64 `json:"monitorPort,omitempty" tf:"monitor_port,omitempty"`

	// The Name of the Monitor.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The id of the pool that this monitor will be assigned to.
	PoolID *string `json:"poolId,omitempty" tf:"pool_id,omitempty"`

	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Required for admins. The UUID of the tenant who owns
	// the monitor. Only administrative users can specify a tenant UUID
	// other than their own. Changing this creates a new monitor.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// Maximum number of seconds for a monitor to wait for a
	// ping reply before it times out. The value must be less than the delay value.
	Timeout *float64 `json:"timeout,omitempty" tf:"timeout,omitempty"`

	// The type of probe, which is TCP, UDP_CONNECT, or HTTP,
	// that is sent by the load balancer to verify the member state. Changing this
	// creates a new monitor.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Required for HTTP types. URI path that will be
	// accessed if monitor type is HTTP.
	URLPath *string `json:"urlPath,omitempty" tf:"url_path,omitempty"`
}

type MonitorV2Parameters struct {

	// The administrative state of the monitor.
	// A valid value is true (UP) or false (DOWN).
	// +kubebuilder:validation:Optional
	AdminStateUp *bool `json:"adminStateUp,omitempty" tf:"admin_state_up,omitempty"`

	// The time, in seconds, between sending probes to members.
	// +kubebuilder:validation:Optional
	Delay *float64 `json:"delay,omitempty" tf:"delay,omitempty"`

	// The domain_name of the HTTP request during the health check.
	// +kubebuilder:validation:Optional
	DomainName *string `json:"domainName,omitempty" tf:"domain_name,omitempty"`

	// Required for HTTP types. Expected HTTP codes
	// for a passing HTTP monitor. You can either specify a single status like
	// "200", or a list like "200,202".
	// +kubebuilder:validation:Optional
	ExpectedCodes *string `json:"expectedCodes,omitempty" tf:"expected_codes,omitempty"`

	// Required for HTTP types. The HTTP method used
	// for requests by the monitor. If this attribute is not specified, it
	// defaults to GET. The value can be GET, HEAD, POST, PUT, DELETE,
	// TRACE, OPTIONS, CONNECT, and PATCH.
	// +kubebuilder:validation:Optional
	HTTPMethod *string `json:"httpMethod,omitempty" tf:"http_method,omitempty"`

	// Number of permissible ping failures before
	// changing the member's status to INACTIVE. Must be a number between 1 and 10.
	// +kubebuilder:validation:Optional
	MaxRetries *float64 `json:"maxRetries,omitempty" tf:"max_retries,omitempty"`

	// Specifies the health check port. The port number
	// ranges from 1 to 65535. The value is left blank by default, indicating that
	// the port of the backend server is used as the health check port.
	// +kubebuilder:validation:Optional
	MonitorPort *float64 `json:"monitorPort,omitempty" tf:"monitor_port,omitempty"`

	// The Name of the Monitor.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The id of the pool that this monitor will be assigned to.
	// +kubebuilder:validation:Optional
	PoolID *string `json:"poolId,omitempty" tf:"pool_id,omitempty"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Required for admins. The UUID of the tenant who owns
	// the monitor. Only administrative users can specify a tenant UUID
	// other than their own. Changing this creates a new monitor.
	// +kubebuilder:validation:Optional
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// Maximum number of seconds for a monitor to wait for a
	// ping reply before it times out. The value must be less than the delay value.
	// +kubebuilder:validation:Optional
	Timeout *float64 `json:"timeout,omitempty" tf:"timeout,omitempty"`

	// The type of probe, which is TCP, UDP_CONNECT, or HTTP,
	// that is sent by the load balancer to verify the member state. Changing this
	// creates a new monitor.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Required for HTTP types. URI path that will be
	// accessed if monitor type is HTTP.
	// +kubebuilder:validation:Optional
	URLPath *string `json:"urlPath,omitempty" tf:"url_path,omitempty"`
}

// MonitorV2Spec defines the desired state of MonitorV2
type MonitorV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MonitorV2Parameters `json:"forProvider"`
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
	InitProvider MonitorV2InitParameters `json:"initProvider,omitempty"`
}

// MonitorV2Status defines the observed state of MonitorV2.
type MonitorV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MonitorV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// MonitorV2 is the Schema for the MonitorV2s API. Manages a ELB Monitor resource within OpenTelekomCloud.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,opentelekomcloud}
type MonitorV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.delay) || (has(self.initProvider) && has(self.initProvider.delay))",message="spec.forProvider.delay is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.maxRetries) || (has(self.initProvider) && has(self.initProvider.maxRetries))",message="spec.forProvider.maxRetries is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.poolId) || (has(self.initProvider) && has(self.initProvider.poolId))",message="spec.forProvider.poolId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.timeout) || (has(self.initProvider) && has(self.initProvider.timeout))",message="spec.forProvider.timeout is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type) || (has(self.initProvider) && has(self.initProvider.type))",message="spec.forProvider.type is a required parameter"
	Spec   MonitorV2Spec   `json:"spec"`
	Status MonitorV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MonitorV2List contains a list of MonitorV2s
type MonitorV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MonitorV2 `json:"items"`
}

// Repository type metadata.
var (
	MonitorV2_Kind             = "MonitorV2"
	MonitorV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MonitorV2_Kind}.String()
	MonitorV2_KindAPIVersion   = MonitorV2_Kind + "." + CRDGroupVersion.String()
	MonitorV2_GroupVersionKind = CRDGroupVersion.WithKind(MonitorV2_Kind)
)

func init() {
	SchemeBuilder.Register(&MonitorV2{}, &MonitorV2List{})
}
