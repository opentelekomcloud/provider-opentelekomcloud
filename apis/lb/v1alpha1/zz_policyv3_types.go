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

type FixedResponseConfigInitParameters struct {

	// - Specifies the format of the response body.
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// - Specifies the content of the response message body.
	MessageBody *string `json:"messageBody,omitempty" tf:"message_body,omitempty"`

	// Specifies the fixed HTTP status code configured in the forwarding rule.
	// The value can be any integer in the range of 200-299, 400-499, or 500-599.
	StatusCode *string `json:"statusCode,omitempty" tf:"status_code,omitempty"`
}

type FixedResponseConfigObservation struct {

	// - Specifies the format of the response body.
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// - Specifies the content of the response message body.
	MessageBody *string `json:"messageBody,omitempty" tf:"message_body,omitempty"`

	// Specifies the fixed HTTP status code configured in the forwarding rule.
	// The value can be any integer in the range of 200-299, 400-499, or 500-599.
	StatusCode *string `json:"statusCode,omitempty" tf:"status_code,omitempty"`
}

type FixedResponseConfigParameters struct {

	// - Specifies the format of the response body.
	// +kubebuilder:validation:Optional
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// - Specifies the content of the response message body.
	// +kubebuilder:validation:Optional
	MessageBody *string `json:"messageBody,omitempty" tf:"message_body,omitempty"`

	// Specifies the fixed HTTP status code configured in the forwarding rule.
	// The value can be any integer in the range of 200-299, 400-499, or 500-599.
	// +kubebuilder:validation:Optional
	StatusCode *string `json:"statusCode" tf:"status_code,omitempty"`
}

type PolicyV3InitParameters struct {

	// The Policy action - can either be REDIRECT_TO_POOL,
	// or REDIRECT_TO_LISTENER. Changing this creates a new Policy.
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// Provides supplementary information about the forwarding policy.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies the configuration of the page that will be returned.
	// This parameter will take effect when advanced_forwarding is set to true.
	// If this parameter is passed and advanced_forwarding is set to false, an error will be returned.
	// Not available in eu-nl.
	FixedResponseConfig []FixedResponseConfigInitParameters `json:"fixedResponseConfig,omitempty" tf:"fixed_response_config,omitempty"`

	// The Listener on which the Policy will be associated with.
	// Changing this creates a new Policy.
	ListenerID *string `json:"listenerId,omitempty" tf:"listener_id,omitempty"`

	// Specifies the forwarding policy name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The position of this policy on the listener. Positions start at 1.
	// Changing this creates a new Policy.
	Position *float64 `json:"position,omitempty" tf:"position,omitempty"`

	// Specifies the forwarding policy priority.
	// A smaller value indicates a higher priority. The value must be unique for forwarding policies of the same listener.
	// This parameter will take effect only when advanced_forwarding is set to true.
	// If this parameter is passed and advanced_forwarding is set to false, an error will be returned.
	// This parameter is unsupported for shared load balancers and not available in eu-nl.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// Required for admins. The UUID of the tenant who owns
	// the Policy. Only administrative users can specify a tenant UUID other than
	// their own. Changing this creates a new Policy.
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Requests matching this policy will be redirected to the listener with this ID.
	// Only valid if action is REDIRECT_TO_LISTENER.
	RedirectListenerID *string `json:"redirectListenerId,omitempty" tf:"redirect_listener_id,omitempty"`

	// Requests matching this policy will be redirected to the pool with this ID.
	// Only valid if action is REDIRECT_TO_POOL.
	RedirectPoolID *string `json:"redirectPoolId,omitempty" tf:"redirect_pool_id,omitempty"`

	// Specifies the configuration of the backend server group that the requests
	// are forwarded to. This parameter is valid only when action is set to REDIRECT_TO_POOL.
	RedirectPoolsConfig []RedirectPoolsConfigInitParameters `json:"redirectPoolsConfig,omitempty" tf:"redirect_pools_config,omitempty"`

	// Specifies the URL to which requests are forwarded.
	RedirectURL *string `json:"redirectUrl,omitempty" tf:"redirect_url,omitempty"`

	// Specifies the URL to which requests are forwarded.
	// For dedicated load balancers, This parameter will take effect when advanced_forwarding is set to true.
	// If it is passed when advanced_forwarding is set to false, an error will be returned. Not available in eu-nl.
	RedirectURLConfig []RedirectURLConfigInitParameters `json:"redirectUrlConfig,omitempty" tf:"redirect_url_config,omitempty"`

	// Lists the forwarding rules in the forwarding policy.
	Rules []RulesInitParameters `json:"rules,omitempty" tf:"rules,omitempty"`
}

type PolicyV3Observation struct {

	// The Policy action - can either be REDIRECT_TO_POOL,
	// or REDIRECT_TO_LISTENER. Changing this creates a new Policy.
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// Provides supplementary information about the forwarding policy.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies the configuration of the page that will be returned.
	// This parameter will take effect when advanced_forwarding is set to true.
	// If this parameter is passed and advanced_forwarding is set to false, an error will be returned.
	// Not available in eu-nl.
	FixedResponseConfig []FixedResponseConfigObservation `json:"fixedResponseConfig,omitempty" tf:"fixed_response_config,omitempty"`

	// The unique ID for the policy.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Listener on which the Policy will be associated with.
	// Changing this creates a new Policy.
	ListenerID *string `json:"listenerId,omitempty" tf:"listener_id,omitempty"`

	// Specifies the forwarding policy name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The position of this policy on the listener. Positions start at 1.
	// Changing this creates a new Policy.
	Position *float64 `json:"position,omitempty" tf:"position,omitempty"`

	// Specifies the forwarding policy priority.
	// A smaller value indicates a higher priority. The value must be unique for forwarding policies of the same listener.
	// This parameter will take effect only when advanced_forwarding is set to true.
	// If this parameter is passed and advanced_forwarding is set to false, an error will be returned.
	// This parameter is unsupported for shared load balancers and not available in eu-nl.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// Required for admins. The UUID of the tenant who owns
	// the Policy. Only administrative users can specify a tenant UUID other than
	// their own. Changing this creates a new Policy.
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Requests matching this policy will be redirected to the listener with this ID.
	// Only valid if action is REDIRECT_TO_LISTENER.
	RedirectListenerID *string `json:"redirectListenerId,omitempty" tf:"redirect_listener_id,omitempty"`

	// Requests matching this policy will be redirected to the pool with this ID.
	// Only valid if action is REDIRECT_TO_POOL.
	RedirectPoolID *string `json:"redirectPoolId,omitempty" tf:"redirect_pool_id,omitempty"`

	// Specifies the configuration of the backend server group that the requests
	// are forwarded to. This parameter is valid only when action is set to REDIRECT_TO_POOL.
	RedirectPoolsConfig []RedirectPoolsConfigObservation `json:"redirectPoolsConfig,omitempty" tf:"redirect_pools_config,omitempty"`

	// Specifies the URL to which requests are forwarded.
	RedirectURL *string `json:"redirectUrl,omitempty" tf:"redirect_url,omitempty"`

	// Specifies the URL to which requests are forwarded.
	// For dedicated load balancers, This parameter will take effect when advanced_forwarding is set to true.
	// If it is passed when advanced_forwarding is set to false, an error will be returned. Not available in eu-nl.
	RedirectURLConfig []RedirectURLConfigObservation `json:"redirectUrlConfig,omitempty" tf:"redirect_url_config,omitempty"`

	// Lists the forwarding rules in the forwarding policy.
	Rules []RulesObservation `json:"rules,omitempty" tf:"rules,omitempty"`

	// Specifies the provisioning status of the forwarding policy.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type PolicyV3Parameters struct {

	// The Policy action - can either be REDIRECT_TO_POOL,
	// or REDIRECT_TO_LISTENER. Changing this creates a new Policy.
	// +kubebuilder:validation:Optional
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// Provides supplementary information about the forwarding policy.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies the configuration of the page that will be returned.
	// This parameter will take effect when advanced_forwarding is set to true.
	// If this parameter is passed and advanced_forwarding is set to false, an error will be returned.
	// Not available in eu-nl.
	// +kubebuilder:validation:Optional
	FixedResponseConfig []FixedResponseConfigParameters `json:"fixedResponseConfig,omitempty" tf:"fixed_response_config,omitempty"`

	// The Listener on which the Policy will be associated with.
	// Changing this creates a new Policy.
	// +kubebuilder:validation:Optional
	ListenerID *string `json:"listenerId,omitempty" tf:"listener_id,omitempty"`

	// Specifies the forwarding policy name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The position of this policy on the listener. Positions start at 1.
	// Changing this creates a new Policy.
	// +kubebuilder:validation:Optional
	Position *float64 `json:"position,omitempty" tf:"position,omitempty"`

	// Specifies the forwarding policy priority.
	// A smaller value indicates a higher priority. The value must be unique for forwarding policies of the same listener.
	// This parameter will take effect only when advanced_forwarding is set to true.
	// If this parameter is passed and advanced_forwarding is set to false, an error will be returned.
	// This parameter is unsupported for shared load balancers and not available in eu-nl.
	// +kubebuilder:validation:Optional
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// Required for admins. The UUID of the tenant who owns
	// the Policy. Only administrative users can specify a tenant UUID other than
	// their own. Changing this creates a new Policy.
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Requests matching this policy will be redirected to the listener with this ID.
	// Only valid if action is REDIRECT_TO_LISTENER.
	// +kubebuilder:validation:Optional
	RedirectListenerID *string `json:"redirectListenerId,omitempty" tf:"redirect_listener_id,omitempty"`

	// Requests matching this policy will be redirected to the pool with this ID.
	// Only valid if action is REDIRECT_TO_POOL.
	// +kubebuilder:validation:Optional
	RedirectPoolID *string `json:"redirectPoolId,omitempty" tf:"redirect_pool_id,omitempty"`

	// Specifies the configuration of the backend server group that the requests
	// are forwarded to. This parameter is valid only when action is set to REDIRECT_TO_POOL.
	// +kubebuilder:validation:Optional
	RedirectPoolsConfig []RedirectPoolsConfigParameters `json:"redirectPoolsConfig,omitempty" tf:"redirect_pools_config,omitempty"`

	// Specifies the URL to which requests are forwarded.
	// +kubebuilder:validation:Optional
	RedirectURL *string `json:"redirectUrl,omitempty" tf:"redirect_url,omitempty"`

	// Specifies the URL to which requests are forwarded.
	// For dedicated load balancers, This parameter will take effect when advanced_forwarding is set to true.
	// If it is passed when advanced_forwarding is set to false, an error will be returned. Not available in eu-nl.
	// +kubebuilder:validation:Optional
	RedirectURLConfig []RedirectURLConfigParameters `json:"redirectUrlConfig,omitempty" tf:"redirect_url_config,omitempty"`

	// Lists the forwarding rules in the forwarding policy.
	// +kubebuilder:validation:Optional
	Rules []RulesParameters `json:"rules,omitempty" tf:"rules,omitempty"`
}

type RedirectPoolsConfigInitParameters struct {

	// - Specifies the ID of the backend server group.
	PoolID *string `json:"poolId,omitempty" tf:"pool_id,omitempty"`

	// - Specifies the weight of the backend server group. The value ranges from 0 to 100.
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type RedirectPoolsConfigObservation struct {

	// - Specifies the ID of the backend server group.
	PoolID *string `json:"poolId,omitempty" tf:"pool_id,omitempty"`

	// - Specifies the weight of the backend server group. The value ranges from 0 to 100.
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type RedirectPoolsConfigParameters struct {

	// - Specifies the ID of the backend server group.
	// +kubebuilder:validation:Optional
	PoolID *string `json:"poolId" tf:"pool_id,omitempty"`

	// - Specifies the weight of the backend server group. The value ranges from 0 to 100.
	// +kubebuilder:validation:Optional
	Weight *float64 `json:"weight" tf:"weight,omitempty"`
}

type RedirectURLConfigInitParameters struct {

	// - Specifies the host name that requests are redirected to.
	// The value can contain only letters, digits, hyphens (-), and periods (.) and must start with a letter or digit.
	// The default value is ${host}, indicating that the host of the request will be used.
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	// - Specifies the path that requests are redirected to.
	// The default value is ${path}, indicating that the path of the request will be used.
	// The value can contain only letters, digits, and special characters _~';@^- %#&$.*+?,=!:|/()[]{}
	// and must start with a slash (/).
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// - Specifies the port that requests are redirected to. The default value is ${port},
	// indicating that the port of the request will be used.
	Port *string `json:"port,omitempty" tf:"port,omitempty"`

	// - Specifies the protocol for redirection. The value can be HTTP, HTTPS,
	// or ${protocol}.
	// The default value is ${protocol}, indicating that the protocol of the request will be used.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// - Specifies the query string set in the URL for redirection.
	// The default value is ${query}, indicating that the query string of the request will be used.
	Query *string `json:"query,omitempty" tf:"query,omitempty"`

	// Specifies the fixed HTTP status code configured in the forwarding rule.
	// The value can be any integer in the range of 200-299, 400-499, or 500-599.
	StatusCode *string `json:"statusCode,omitempty" tf:"status_code,omitempty"`
}

type RedirectURLConfigObservation struct {

	// - Specifies the host name that requests are redirected to.
	// The value can contain only letters, digits, hyphens (-), and periods (.) and must start with a letter or digit.
	// The default value is ${host}, indicating that the host of the request will be used.
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	// - Specifies the path that requests are redirected to.
	// The default value is ${path}, indicating that the path of the request will be used.
	// The value can contain only letters, digits, and special characters _~';@^- %#&$.*+?,=!:|/()[]{}
	// and must start with a slash (/).
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// - Specifies the port that requests are redirected to. The default value is ${port},
	// indicating that the port of the request will be used.
	Port *string `json:"port,omitempty" tf:"port,omitempty"`

	// - Specifies the protocol for redirection. The value can be HTTP, HTTPS,
	// or ${protocol}.
	// The default value is ${protocol}, indicating that the protocol of the request will be used.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// - Specifies the query string set in the URL for redirection.
	// The default value is ${query}, indicating that the query string of the request will be used.
	Query *string `json:"query,omitempty" tf:"query,omitempty"`

	// Specifies the fixed HTTP status code configured in the forwarding rule.
	// The value can be any integer in the range of 200-299, 400-499, or 500-599.
	StatusCode *string `json:"statusCode,omitempty" tf:"status_code,omitempty"`
}

type RedirectURLConfigParameters struct {

	// - Specifies the host name that requests are redirected to.
	// The value can contain only letters, digits, hyphens (-), and periods (.) and must start with a letter or digit.
	// The default value is ${host}, indicating that the host of the request will be used.
	// +kubebuilder:validation:Optional
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	// - Specifies the path that requests are redirected to.
	// The default value is ${path}, indicating that the path of the request will be used.
	// The value can contain only letters, digits, and special characters _~';@^- %#&$.*+?,=!:|/()[]{}
	// and must start with a slash (/).
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// - Specifies the port that requests are redirected to. The default value is ${port},
	// indicating that the port of the request will be used.
	// +kubebuilder:validation:Optional
	Port *string `json:"port,omitempty" tf:"port,omitempty"`

	// - Specifies the protocol for redirection. The value can be HTTP, HTTPS,
	// or ${protocol}.
	// The default value is ${protocol}, indicating that the protocol of the request will be used.
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// - Specifies the query string set in the URL for redirection.
	// The default value is ${query}, indicating that the query string of the request will be used.
	// +kubebuilder:validation:Optional
	Query *string `json:"query,omitempty" tf:"query,omitempty"`

	// Specifies the fixed HTTP status code configured in the forwarding rule.
	// The value can be any integer in the range of 200-299, 400-499, or 500-599.
	// +kubebuilder:validation:Optional
	StatusCode *string `json:"statusCode" tf:"status_code,omitempty"`
}

type RulesInitParameters struct {

	// - Specifies how requests are matched with the domain name or URL.
	// The values can be: EQUAL_TO, REGEX, STARTS_WITH.
	CompareType *string `json:"compareType,omitempty" tf:"compare_type,omitempty"`

	// Specifies the match content. The value can be one of the following: HOST_NAME, PATH.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Specifies the value of the match item. For example, if a domain name is
	// used for matching, value is the domain name.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type RulesObservation struct {

	// - Specifies how requests are matched with the domain name or URL.
	// The values can be: EQUAL_TO, REGEX, STARTS_WITH.
	CompareType *string `json:"compareType,omitempty" tf:"compare_type,omitempty"`

	// Specifies the match content. The value can be one of the following: HOST_NAME, PATH.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Specifies the value of the match item. For example, if a domain name is
	// used for matching, value is the domain name.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type RulesParameters struct {

	// - Specifies how requests are matched with the domain name or URL.
	// The values can be: EQUAL_TO, REGEX, STARTS_WITH.
	// +kubebuilder:validation:Optional
	CompareType *string `json:"compareType" tf:"compare_type,omitempty"`

	// Specifies the match content. The value can be one of the following: HOST_NAME, PATH.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`

	// Specifies the value of the match item. For example, if a domain name is
	// used for matching, value is the domain name.
	// +kubebuilder:validation:Optional
	Value *string `json:"value" tf:"value,omitempty"`
}

// PolicyV3Spec defines the desired state of PolicyV3
type PolicyV3Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PolicyV3Parameters `json:"forProvider"`
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
	InitProvider PolicyV3InitParameters `json:"initProvider,omitempty"`
}

// PolicyV3Status defines the observed state of PolicyV3.
type PolicyV3Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PolicyV3Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PolicyV3 is the Schema for the PolicyV3s API. Manages a LB Policy resource within OpenTelekomCloud.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,opentelekomcloud}
type PolicyV3 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.action) || (has(self.initProvider) && has(self.initProvider.action))",message="spec.forProvider.action is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.listenerId) || (has(self.initProvider) && has(self.initProvider.listenerId))",message="spec.forProvider.listenerId is a required parameter"
	Spec   PolicyV3Spec   `json:"spec"`
	Status PolicyV3Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PolicyV3List contains a list of PolicyV3s
type PolicyV3List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PolicyV3 `json:"items"`
}

// Repository type metadata.
var (
	PolicyV3_Kind             = "PolicyV3"
	PolicyV3_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PolicyV3_Kind}.String()
	PolicyV3_KindAPIVersion   = PolicyV3_Kind + "." + CRDGroupVersion.String()
	PolicyV3_GroupVersionKind = CRDGroupVersion.WithKind(PolicyV3_Kind)
)

func init() {
	SchemeBuilder.Register(&PolicyV3{}, &PolicyV3List{})
}