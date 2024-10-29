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

type IPGroupInitParameters struct {

	// Specifies whether to enable access control.
	// true: Access control will be enabled.
	// false (default): Access control will be disabled.
	Enable *bool `json:"enable,omitempty" tf:"enable,omitempty"`

	// Specifies the ID of the IP address group associated with the listener.
	// Specifies the ID of the IP address group associated with the listener.
	// If ip_list in opentelekomcloud_lb_ipgroup_v3 is set to an empty array [] and type to whitelist, no IP addresses are allowed to access the listener.
	// If ip_list in opentelekomcloud_lb_ipgroup_v3 is set to an empty array [] and type to blacklist, any IP address is allowed to access the listener.
	// +crossplane:generate:reference:type=github.com/opentelekomcloud/provider-opentelekomcloud/apis/lb/v1alpha1.IpgroupV3
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Reference to a IpgroupV3 in lb to populate id.
	// +kubebuilder:validation:Optional
	IDRef *v1.Reference `json:"idRef,omitempty" tf:"-"`

	// Selector for a IpgroupV3 in lb to populate id.
	// +kubebuilder:validation:Optional
	IDSelector *v1.Selector `json:"idSelector,omitempty" tf:"-"`

	// Specifies how access to the listener is controlled.
	// white (default): A whitelist will be configured. Only IP addresses in the whitelist can access the listener.
	// black: A blacklist will be configured. IP addresses in the blacklist are not allowed to access the listener.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type IPGroupObservation struct {

	// Specifies whether to enable access control.
	// true: Access control will be enabled.
	// false (default): Access control will be disabled.
	Enable *bool `json:"enable,omitempty" tf:"enable,omitempty"`

	// Specifies the ID of the IP address group associated with the listener.
	// Specifies the ID of the IP address group associated with the listener.
	// If ip_list in opentelekomcloud_lb_ipgroup_v3 is set to an empty array [] and type to whitelist, no IP addresses are allowed to access the listener.
	// If ip_list in opentelekomcloud_lb_ipgroup_v3 is set to an empty array [] and type to blacklist, any IP address is allowed to access the listener.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies how access to the listener is controlled.
	// white (default): A whitelist will be configured. Only IP addresses in the whitelist can access the listener.
	// black: A blacklist will be configured. IP addresses in the blacklist are not allowed to access the listener.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type IPGroupParameters struct {

	// Specifies whether to enable access control.
	// true: Access control will be enabled.
	// false (default): Access control will be disabled.
	// +kubebuilder:validation:Optional
	Enable *bool `json:"enable,omitempty" tf:"enable,omitempty"`

	// Specifies the ID of the IP address group associated with the listener.
	// Specifies the ID of the IP address group associated with the listener.
	// If ip_list in opentelekomcloud_lb_ipgroup_v3 is set to an empty array [] and type to whitelist, no IP addresses are allowed to access the listener.
	// If ip_list in opentelekomcloud_lb_ipgroup_v3 is set to an empty array [] and type to blacklist, any IP address is allowed to access the listener.
	// +crossplane:generate:reference:type=github.com/opentelekomcloud/provider-opentelekomcloud/apis/lb/v1alpha1.IpgroupV3
	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Reference to a IpgroupV3 in lb to populate id.
	// +kubebuilder:validation:Optional
	IDRef *v1.Reference `json:"idRef,omitempty" tf:"-"`

	// Selector for a IpgroupV3 in lb to populate id.
	// +kubebuilder:validation:Optional
	IDSelector *v1.Selector `json:"idSelector,omitempty" tf:"-"`

	// Specifies how access to the listener is controlled.
	// white (default): A whitelist will be configured. Only IP addresses in the whitelist can access the listener.
	// black: A blacklist will be configured. IP addresses in the blacklist are not allowed to access the listener.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type InsertHeadersInitParameters struct {

	// Specifies whether to transparently transmit the load balancer EIP
	// to backend servers. If forward_elb_ip is set to true, the load balancer EIP will be stored in
	// the HTTP header and passed to backend servers.
	ForwardELBIP *bool `json:"forwardElbIp,omitempty" tf:"forward_elb_ip,omitempty"`

	// Specifies whether to transparently transmit the source port of
	// the client to backend servers. If forwarded_for_port is set to true, the source port of the
	// client will be stored in the HTTP header and passed to backend servers.
	ForwardedForPort *bool `json:"forwardedForPort,omitempty" tf:"forwarded_for_port,omitempty"`

	// Specifies whether to rewrite the X-Forwarded-Host header.
	// If forwarded_host is set to true, X-Forwarded-Host in the request header from the clients
	// can be set to Host in the request header sent from the load balancer to backend servers.
	ForwardedHost *bool `json:"forwardedHost,omitempty" tf:"forwarded_host,omitempty"`

	// Specifies whether to transparently transmit the listening port of
	// the load balancer to backend servers. If forwarded_port is set to true, the listening port of
	// the load balancer will be stored in the HTTP header and passed to backend servers.
	ForwardedPort *bool `json:"forwardedPort,omitempty" tf:"forwarded_port,omitempty"`
}

type InsertHeadersObservation struct {

	// Specifies whether to transparently transmit the load balancer EIP
	// to backend servers. If forward_elb_ip is set to true, the load balancer EIP will be stored in
	// the HTTP header and passed to backend servers.
	ForwardELBIP *bool `json:"forwardElbIp,omitempty" tf:"forward_elb_ip,omitempty"`

	// Specifies whether to transparently transmit the source port of
	// the client to backend servers. If forwarded_for_port is set to true, the source port of the
	// client will be stored in the HTTP header and passed to backend servers.
	ForwardedForPort *bool `json:"forwardedForPort,omitempty" tf:"forwarded_for_port,omitempty"`

	// Specifies whether to rewrite the X-Forwarded-Host header.
	// If forwarded_host is set to true, X-Forwarded-Host in the request header from the clients
	// can be set to Host in the request header sent from the load balancer to backend servers.
	ForwardedHost *bool `json:"forwardedHost,omitempty" tf:"forwarded_host,omitempty"`

	// Specifies whether to transparently transmit the listening port of
	// the load balancer to backend servers. If forwarded_port is set to true, the listening port of
	// the load balancer will be stored in the HTTP header and passed to backend servers.
	ForwardedPort *bool `json:"forwardedPort,omitempty" tf:"forwarded_port,omitempty"`
}

type InsertHeadersParameters struct {

	// Specifies whether to transparently transmit the load balancer EIP
	// to backend servers. If forward_elb_ip is set to true, the load balancer EIP will be stored in
	// the HTTP header and passed to backend servers.
	// +kubebuilder:validation:Optional
	ForwardELBIP *bool `json:"forwardElbIp,omitempty" tf:"forward_elb_ip,omitempty"`

	// Specifies whether to transparently transmit the source port of
	// the client to backend servers. If forwarded_for_port is set to true, the source port of the
	// client will be stored in the HTTP header and passed to backend servers.
	// +kubebuilder:validation:Optional
	ForwardedForPort *bool `json:"forwardedForPort,omitempty" tf:"forwarded_for_port,omitempty"`

	// Specifies whether to rewrite the X-Forwarded-Host header.
	// If forwarded_host is set to true, X-Forwarded-Host in the request header from the clients
	// can be set to Host in the request header sent from the load balancer to backend servers.
	// +kubebuilder:validation:Optional
	ForwardedHost *bool `json:"forwardedHost,omitempty" tf:"forwarded_host,omitempty"`

	// Specifies whether to transparently transmit the listening port of
	// the load balancer to backend servers. If forwarded_port is set to true, the listening port of
	// the load balancer will be stored in the HTTP header and passed to backend servers.
	// +kubebuilder:validation:Optional
	ForwardedPort *bool `json:"forwardedPort,omitempty" tf:"forwarded_port,omitempty"`
}

type ListenerV3InitParameters struct {
	AdminStateUp *bool `json:"adminStateUp,omitempty" tf:"admin_state_up,omitempty"`

	// Specifies whether to enable advanced forwarding.
	// If advanced forwarding is enabled, more flexible forwarding policies and rules are supported.
	// The value can be true (enable advanced forwarding) or false (disable advanced forwarding),
	// and the default value is false. Changing this creates a new Listener.
	AdvancedForwarding *bool `json:"advancedForwarding,omitempty" tf:"advanced_forwarding,omitempty"`

	// Specifies the ID of the CA certificate used by the listener.
	ClientCATLSContainerRef *string `json:"clientCaTlsContainerRef,omitempty" tf:"client_ca_tls_container_ref,omitempty"`

	// Specifies the timeout duration for waiting for a request from a client, in seconds.
	// This parameter is available only for HTTP and HTTPS listeners. The value ranges from 1 to 300, and
	// the default value is 60. An error will be returned if you configure this parameter for TCP and UDP listeners.
	ClientTimeout *float64 `json:"clientTimeout,omitempty" tf:"client_timeout,omitempty"`

	// Specifies the ID of the default backend server group. If there is no
	// matched forwarding policy, requests are forwarded to the default backend server for processing.
	// +crossplane:generate:reference:type=github.com/opentelekomcloud/provider-opentelekomcloud/apis/lb/v1alpha1.PoolV3
	DefaultPoolID *string `json:"defaultPoolId,omitempty" tf:"default_pool_id,omitempty"`

	// Reference to a PoolV3 in lb to populate defaultPoolId.
	// +kubebuilder:validation:Optional
	DefaultPoolIDRef *v1.Reference `json:"defaultPoolIdRef,omitempty" tf:"-"`

	// Selector for a PoolV3 in lb to populate defaultPoolId.
	// +kubebuilder:validation:Optional
	DefaultPoolIDSelector *v1.Selector `json:"defaultPoolIdSelector,omitempty" tf:"-"`

	// Specifies the ID of the server certificate used by the listener.
	// +crossplane:generate:reference:type=github.com/opentelekomcloud/provider-opentelekomcloud/apis/lb/v1alpha1.CertificateV3
	DefaultTLSContainerRef *string `json:"defaultTlsContainerRef,omitempty" tf:"default_tls_container_ref,omitempty"`

	// Reference to a CertificateV3 in lb to populate defaultTlsContainerRef.
	// +kubebuilder:validation:Optional
	DefaultTLSContainerRefRef *v1.Reference `json:"defaultTlsContainerRefRef,omitempty" tf:"-"`

	// Selector for a CertificateV3 in lb to populate defaultTlsContainerRef.
	// +kubebuilder:validation:Optional
	DefaultTLSContainerRefSelector *v1.Selector `json:"defaultTlsContainerRefSelector,omitempty" tf:"-"`

	// Provides supplementary information about the listener.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies whether to use HTTP/2. This parameter is available only for HTTPS
	// listeners. If you configure this parameter for other types of listeners, it will not take effect. Enable
	// HTTP/2 if you want the clients to use HTTP/2 to communicate with the load balancer.
	// However, connections between the load balancer and backend servers use HTTP/1.x by default.
	Http2Enable *bool `json:"http2Enable,omitempty" tf:"http2_enable,omitempty"`

	// Specifies the IP address group associated with the listener.
	IPGroup []IPGroupInitParameters `json:"ipGroup,omitempty" tf:"ip_group,omitempty"`

	// Specifies the HTTP header fields.
	InsertHeaders []InsertHeadersInitParameters `json:"insertHeaders,omitempty" tf:"insert_headers,omitempty"`

	// Specifies the idle timeout duration, in seconds.
	KeepAliveTimeout *float64 `json:"keepAliveTimeout,omitempty" tf:"keep_alive_timeout,omitempty"`

	// Specifies the ID of the load balancer that the listener is added to.
	// +crossplane:generate:reference:type=github.com/opentelekomcloud/provider-opentelekomcloud/apis/lb/v1alpha1.LoadbalancerV3
	LoadbalancerID *string `json:"loadbalancerId,omitempty" tf:"loadbalancer_id,omitempty"`

	// Reference to a LoadbalancerV3 in lb to populate loadbalancerId.
	// +kubebuilder:validation:Optional
	LoadbalancerIDRef *v1.Reference `json:"loadbalancerIdRef,omitempty" tf:"-"`

	// Selector for a LoadbalancerV3 in lb to populate loadbalancerId.
	// +kubebuilder:validation:Optional
	LoadbalancerIDSelector *v1.Selector `json:"loadbalancerIdSelector,omitempty" tf:"-"`

	// Specifies whether to enable health check retries for backend servers.
	// This parameter is available only for HTTP and HTTPS listeners. An error will be returned if you configure
	// this parameter for TCP and UDP listeners.
	MemberRetryEnable *bool `json:"memberRetryEnable,omitempty" tf:"member_retry_enable,omitempty"`

	// Specifies the timeout duration for waiting for a request from a
	// backend server, in seconds. This parameter is available only for HTTP and HTTPS listeners.
	// The value ranges from 1 to 300, and the default value is 60. An error will be returned if
	// you configure this parameter for TCP and UDP listeners.
	MemberTimeout *float64 `json:"memberTimeout,omitempty" tf:"member_timeout,omitempty"`

	// Specifies the listener name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The protocol - can either be TCP, HTTP, HTTPS or UDP.
	// Changing this creates a new Listener.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Specifies the port used by the listener. Changing this creates a new Listener.
	ProtocolPort *float64 `json:"protocolPort,omitempty" tf:"protocol_port,omitempty"`

	// Specifies the ID of the custom security policy.
	// +crossplane:generate:reference:type=github.com/opentelekomcloud/provider-opentelekomcloud/apis/lb/v1alpha1.SecurityPolicyV3
	SecurityPolicyID *string `json:"securityPolicyId,omitempty" tf:"security_policy_id,omitempty"`

	// Reference to a SecurityPolicyV3 in lb to populate securityPolicyId.
	// +kubebuilder:validation:Optional
	SecurityPolicyIDRef *v1.Reference `json:"securityPolicyIdRef,omitempty" tf:"-"`

	// Selector for a SecurityPolicyV3 in lb to populate securityPolicyId.
	// +kubebuilder:validation:Optional
	SecurityPolicyIDSelector *v1.Selector `json:"securityPolicyIdSelector,omitempty" tf:"-"`

	// Lists the IDs of SNI certificates (server certificates with domain names) used by the listener.
	// Each SNI certificate can have up to 30 domain names, and each domain name in the SNI certificate must be unique.
	// This parameter will be ignored and an empty array will be returned if the listener's protocol is not HTTPS.
	// +listType=set
	SniContainerRefs []*string `json:"sniContainerRefs,omitempty" tf:"sni_container_refs,omitempty"`

	// Specifies how wildcard domain name matches with the SNI certificates
	// used by the listener.
	SniMatchAlgo *string `json:"sniMatchAlgo,omitempty" tf:"sni_match_algo,omitempty"`

	// Specifies the security policy that will be used by the listener.
	// This parameter is available only for HTTPS listeners. An error will be returned if the protocol
	// of the listener is not HTTPS. Possible values are: tls-1-0, tls-1-1, tls-1-0-inherit, tls-1-2,
	// tls-1-2-strict, tls-1-2-fs, tls-1-0-with-1-3, tls-1-2-fs-with-1-3, hybrid-policy-1-0, tls-1-2-strict-no-cbc.
	TLSCiphersPolicy *string `json:"tlsCiphersPolicy,omitempty" tf:"tls_ciphers_policy,omitempty"`

	// Tags key/value pairs to associate with the loadbalancer listener.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ListenerV3Observation struct {
	AdminStateUp *bool `json:"adminStateUp,omitempty" tf:"admin_state_up,omitempty"`

	// Specifies whether to enable advanced forwarding.
	// If advanced forwarding is enabled, more flexible forwarding policies and rules are supported.
	// The value can be true (enable advanced forwarding) or false (disable advanced forwarding),
	// and the default value is false. Changing this creates a new Listener.
	AdvancedForwarding *bool `json:"advancedForwarding,omitempty" tf:"advanced_forwarding,omitempty"`

	// Specifies the ID of the CA certificate used by the listener.
	ClientCATLSContainerRef *string `json:"clientCaTlsContainerRef,omitempty" tf:"client_ca_tls_container_ref,omitempty"`

	// Specifies the timeout duration for waiting for a request from a client, in seconds.
	// This parameter is available only for HTTP and HTTPS listeners. The value ranges from 1 to 300, and
	// the default value is 60. An error will be returned if you configure this parameter for TCP and UDP listeners.
	ClientTimeout *float64 `json:"clientTimeout,omitempty" tf:"client_timeout,omitempty"`

	// Indicates the creation time.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// Specifies the ID of the default backend server group. If there is no
	// matched forwarding policy, requests are forwarded to the default backend server for processing.
	DefaultPoolID *string `json:"defaultPoolId,omitempty" tf:"default_pool_id,omitempty"`

	// Specifies the ID of the server certificate used by the listener.
	DefaultTLSContainerRef *string `json:"defaultTlsContainerRef,omitempty" tf:"default_tls_container_ref,omitempty"`

	// Provides supplementary information about the listener.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies whether to use HTTP/2. This parameter is available only for HTTPS
	// listeners. If you configure this parameter for other types of listeners, it will not take effect. Enable
	// HTTP/2 if you want the clients to use HTTP/2 to communicate with the load balancer.
	// However, connections between the load balancer and backend servers use HTTP/1.x by default.
	Http2Enable *bool `json:"http2Enable,omitempty" tf:"http2_enable,omitempty"`

	// Specifies the ID of the IP address group associated with the listener.
	// Specifies the ID of the IP address group associated with the listener.
	// If ip_list in opentelekomcloud_lb_ipgroup_v3 is set to an empty array [] and type to whitelist, no IP addresses are allowed to access the listener.
	// If ip_list in opentelekomcloud_lb_ipgroup_v3 is set to an empty array [] and type to blacklist, any IP address is allowed to access the listener.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the IP address group associated with the listener.
	IPGroup []IPGroupObservation `json:"ipGroup,omitempty" tf:"ip_group,omitempty"`

	// Specifies the HTTP header fields.
	InsertHeaders []InsertHeadersObservation `json:"insertHeaders,omitempty" tf:"insert_headers,omitempty"`

	// Specifies the idle timeout duration, in seconds.
	KeepAliveTimeout *float64 `json:"keepAliveTimeout,omitempty" tf:"keep_alive_timeout,omitempty"`

	// Specifies the ID of the load balancer that the listener is added to.
	LoadbalancerID *string `json:"loadbalancerId,omitempty" tf:"loadbalancer_id,omitempty"`

	// Specifies whether to enable health check retries for backend servers.
	// This parameter is available only for HTTP and HTTPS listeners. An error will be returned if you configure
	// this parameter for TCP and UDP listeners.
	MemberRetryEnable *bool `json:"memberRetryEnable,omitempty" tf:"member_retry_enable,omitempty"`

	// Specifies the timeout duration for waiting for a request from a
	// backend server, in seconds. This parameter is available only for HTTP and HTTPS listeners.
	// The value ranges from 1 to 300, and the default value is 60. An error will be returned if
	// you configure this parameter for TCP and UDP listeners.
	MemberTimeout *float64 `json:"memberTimeout,omitempty" tf:"member_timeout,omitempty"`

	// Specifies the listener name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The protocol - can either be TCP, HTTP, HTTPS or UDP.
	// Changing this creates a new Listener.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Specifies the port used by the listener. Changing this creates a new Listener.
	ProtocolPort *float64 `json:"protocolPort,omitempty" tf:"protocol_port,omitempty"`

	// Specifies the ID of the custom security policy.
	SecurityPolicyID *string `json:"securityPolicyId,omitempty" tf:"security_policy_id,omitempty"`

	// Lists the IDs of SNI certificates (server certificates with domain names) used by the listener.
	// Each SNI certificate can have up to 30 domain names, and each domain name in the SNI certificate must be unique.
	// This parameter will be ignored and an empty array will be returned if the listener's protocol is not HTTPS.
	// +listType=set
	SniContainerRefs []*string `json:"sniContainerRefs,omitempty" tf:"sni_container_refs,omitempty"`

	// Specifies how wildcard domain name matches with the SNI certificates
	// used by the listener.
	SniMatchAlgo *string `json:"sniMatchAlgo,omitempty" tf:"sni_match_algo,omitempty"`

	// Specifies the security policy that will be used by the listener.
	// This parameter is available only for HTTPS listeners. An error will be returned if the protocol
	// of the listener is not HTTPS. Possible values are: tls-1-0, tls-1-1, tls-1-0-inherit, tls-1-2,
	// tls-1-2-strict, tls-1-2-fs, tls-1-0-with-1-3, tls-1-2-fs-with-1-3, hybrid-policy-1-0, tls-1-2-strict-no-cbc.
	TLSCiphersPolicy *string `json:"tlsCiphersPolicy,omitempty" tf:"tls_ciphers_policy,omitempty"`

	// Tags key/value pairs to associate with the loadbalancer listener.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Indicates the update time.
	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`
}

type ListenerV3Parameters struct {

	// +kubebuilder:validation:Optional
	AdminStateUp *bool `json:"adminStateUp,omitempty" tf:"admin_state_up,omitempty"`

	// Specifies whether to enable advanced forwarding.
	// If advanced forwarding is enabled, more flexible forwarding policies and rules are supported.
	// The value can be true (enable advanced forwarding) or false (disable advanced forwarding),
	// and the default value is false. Changing this creates a new Listener.
	// +kubebuilder:validation:Optional
	AdvancedForwarding *bool `json:"advancedForwarding,omitempty" tf:"advanced_forwarding,omitempty"`

	// Specifies the ID of the CA certificate used by the listener.
	// +kubebuilder:validation:Optional
	ClientCATLSContainerRef *string `json:"clientCaTlsContainerRef,omitempty" tf:"client_ca_tls_container_ref,omitempty"`

	// Specifies the timeout duration for waiting for a request from a client, in seconds.
	// This parameter is available only for HTTP and HTTPS listeners. The value ranges from 1 to 300, and
	// the default value is 60. An error will be returned if you configure this parameter for TCP and UDP listeners.
	// +kubebuilder:validation:Optional
	ClientTimeout *float64 `json:"clientTimeout,omitempty" tf:"client_timeout,omitempty"`

	// Specifies the ID of the default backend server group. If there is no
	// matched forwarding policy, requests are forwarded to the default backend server for processing.
	// +crossplane:generate:reference:type=github.com/opentelekomcloud/provider-opentelekomcloud/apis/lb/v1alpha1.PoolV3
	// +kubebuilder:validation:Optional
	DefaultPoolID *string `json:"defaultPoolId,omitempty" tf:"default_pool_id,omitempty"`

	// Reference to a PoolV3 in lb to populate defaultPoolId.
	// +kubebuilder:validation:Optional
	DefaultPoolIDRef *v1.Reference `json:"defaultPoolIdRef,omitempty" tf:"-"`

	// Selector for a PoolV3 in lb to populate defaultPoolId.
	// +kubebuilder:validation:Optional
	DefaultPoolIDSelector *v1.Selector `json:"defaultPoolIdSelector,omitempty" tf:"-"`

	// Specifies the ID of the server certificate used by the listener.
	// +crossplane:generate:reference:type=github.com/opentelekomcloud/provider-opentelekomcloud/apis/lb/v1alpha1.CertificateV3
	// +kubebuilder:validation:Optional
	DefaultTLSContainerRef *string `json:"defaultTlsContainerRef,omitempty" tf:"default_tls_container_ref,omitempty"`

	// Reference to a CertificateV3 in lb to populate defaultTlsContainerRef.
	// +kubebuilder:validation:Optional
	DefaultTLSContainerRefRef *v1.Reference `json:"defaultTlsContainerRefRef,omitempty" tf:"-"`

	// Selector for a CertificateV3 in lb to populate defaultTlsContainerRef.
	// +kubebuilder:validation:Optional
	DefaultTLSContainerRefSelector *v1.Selector `json:"defaultTlsContainerRefSelector,omitempty" tf:"-"`

	// Provides supplementary information about the listener.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies whether to use HTTP/2. This parameter is available only for HTTPS
	// listeners. If you configure this parameter for other types of listeners, it will not take effect. Enable
	// HTTP/2 if you want the clients to use HTTP/2 to communicate with the load balancer.
	// However, connections between the load balancer and backend servers use HTTP/1.x by default.
	// +kubebuilder:validation:Optional
	Http2Enable *bool `json:"http2Enable,omitempty" tf:"http2_enable,omitempty"`

	// Specifies the IP address group associated with the listener.
	// +kubebuilder:validation:Optional
	IPGroup []IPGroupParameters `json:"ipGroup,omitempty" tf:"ip_group,omitempty"`

	// Specifies the HTTP header fields.
	// +kubebuilder:validation:Optional
	InsertHeaders []InsertHeadersParameters `json:"insertHeaders,omitempty" tf:"insert_headers,omitempty"`

	// Specifies the idle timeout duration, in seconds.
	// +kubebuilder:validation:Optional
	KeepAliveTimeout *float64 `json:"keepAliveTimeout,omitempty" tf:"keep_alive_timeout,omitempty"`

	// Specifies the ID of the load balancer that the listener is added to.
	// +crossplane:generate:reference:type=github.com/opentelekomcloud/provider-opentelekomcloud/apis/lb/v1alpha1.LoadbalancerV3
	// +kubebuilder:validation:Optional
	LoadbalancerID *string `json:"loadbalancerId,omitempty" tf:"loadbalancer_id,omitempty"`

	// Reference to a LoadbalancerV3 in lb to populate loadbalancerId.
	// +kubebuilder:validation:Optional
	LoadbalancerIDRef *v1.Reference `json:"loadbalancerIdRef,omitempty" tf:"-"`

	// Selector for a LoadbalancerV3 in lb to populate loadbalancerId.
	// +kubebuilder:validation:Optional
	LoadbalancerIDSelector *v1.Selector `json:"loadbalancerIdSelector,omitempty" tf:"-"`

	// Specifies whether to enable health check retries for backend servers.
	// This parameter is available only for HTTP and HTTPS listeners. An error will be returned if you configure
	// this parameter for TCP and UDP listeners.
	// +kubebuilder:validation:Optional
	MemberRetryEnable *bool `json:"memberRetryEnable,omitempty" tf:"member_retry_enable,omitempty"`

	// Specifies the timeout duration for waiting for a request from a
	// backend server, in seconds. This parameter is available only for HTTP and HTTPS listeners.
	// The value ranges from 1 to 300, and the default value is 60. An error will be returned if
	// you configure this parameter for TCP and UDP listeners.
	// +kubebuilder:validation:Optional
	MemberTimeout *float64 `json:"memberTimeout,omitempty" tf:"member_timeout,omitempty"`

	// Specifies the listener name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The protocol - can either be TCP, HTTP, HTTPS or UDP.
	// Changing this creates a new Listener.
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Specifies the port used by the listener. Changing this creates a new Listener.
	// +kubebuilder:validation:Optional
	ProtocolPort *float64 `json:"protocolPort,omitempty" tf:"protocol_port,omitempty"`

	// Specifies the ID of the custom security policy.
	// +crossplane:generate:reference:type=github.com/opentelekomcloud/provider-opentelekomcloud/apis/lb/v1alpha1.SecurityPolicyV3
	// +kubebuilder:validation:Optional
	SecurityPolicyID *string `json:"securityPolicyId,omitempty" tf:"security_policy_id,omitempty"`

	// Reference to a SecurityPolicyV3 in lb to populate securityPolicyId.
	// +kubebuilder:validation:Optional
	SecurityPolicyIDRef *v1.Reference `json:"securityPolicyIdRef,omitempty" tf:"-"`

	// Selector for a SecurityPolicyV3 in lb to populate securityPolicyId.
	// +kubebuilder:validation:Optional
	SecurityPolicyIDSelector *v1.Selector `json:"securityPolicyIdSelector,omitempty" tf:"-"`

	// Lists the IDs of SNI certificates (server certificates with domain names) used by the listener.
	// Each SNI certificate can have up to 30 domain names, and each domain name in the SNI certificate must be unique.
	// This parameter will be ignored and an empty array will be returned if the listener's protocol is not HTTPS.
	// +kubebuilder:validation:Optional
	// +listType=set
	SniContainerRefs []*string `json:"sniContainerRefs,omitempty" tf:"sni_container_refs,omitempty"`

	// Specifies how wildcard domain name matches with the SNI certificates
	// used by the listener.
	// +kubebuilder:validation:Optional
	SniMatchAlgo *string `json:"sniMatchAlgo,omitempty" tf:"sni_match_algo,omitempty"`

	// Specifies the security policy that will be used by the listener.
	// This parameter is available only for HTTPS listeners. An error will be returned if the protocol
	// of the listener is not HTTPS. Possible values are: tls-1-0, tls-1-1, tls-1-0-inherit, tls-1-2,
	// tls-1-2-strict, tls-1-2-fs, tls-1-0-with-1-3, tls-1-2-fs-with-1-3, hybrid-policy-1-0, tls-1-2-strict-no-cbc.
	// +kubebuilder:validation:Optional
	TLSCiphersPolicy *string `json:"tlsCiphersPolicy,omitempty" tf:"tls_ciphers_policy,omitempty"`

	// Tags key/value pairs to associate with the loadbalancer listener.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// ListenerV3Spec defines the desired state of ListenerV3
type ListenerV3Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ListenerV3Parameters `json:"forProvider"`
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
	InitProvider ListenerV3InitParameters `json:"initProvider,omitempty"`
}

// ListenerV3Status defines the observed state of ListenerV3.
type ListenerV3Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ListenerV3Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ListenerV3 is the Schema for the ListenerV3s API. Manages a LB Listener resource within OpenTelekomCloud.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,opentelekomcloud}
type ListenerV3 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.protocol) || (has(self.initProvider) && has(self.initProvider.protocol))",message="spec.forProvider.protocol is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.protocolPort) || (has(self.initProvider) && has(self.initProvider.protocolPort))",message="spec.forProvider.protocolPort is a required parameter"
	Spec   ListenerV3Spec   `json:"spec"`
	Status ListenerV3Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ListenerV3List contains a list of ListenerV3s
type ListenerV3List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ListenerV3 `json:"items"`
}

// Repository type metadata.
var (
	ListenerV3_Kind             = "ListenerV3"
	ListenerV3_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ListenerV3_Kind}.String()
	ListenerV3_KindAPIVersion   = ListenerV3_Kind + "." + CRDGroupVersion.String()
	ListenerV3_GroupVersionKind = CRDGroupVersion.WithKind(ListenerV3_Kind)
)

func init() {
	SchemeBuilder.Register(&ListenerV3{}, &ListenerV3List{})
}
