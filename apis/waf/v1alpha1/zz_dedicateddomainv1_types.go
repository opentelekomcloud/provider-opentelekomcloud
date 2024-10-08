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

type DedicatedDomainV1InitParameters struct {

	// Specifies the certificate ID. This parameter is mandatory when client_protocol
	// is set to HTTPS.
	CertificateID *string `json:"certificateId,omitempty" tf:"certificate_id,omitempty"`

	// Specifies the cipher suite of domain.
	// Values are:
	Cipher *string `json:"cipher,omitempty" tf:"cipher,omitempty"`

	// Specifies the protected domain name or IP address (port allowed). For example,
	// www.example.com or *.example.com or www.example.com:89. Changing this creates a new domain.
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// Specifies whether to retain the policy when deleting a domain name.
	// Defaults to true.
	KeepPolicy *bool `json:"keepPolicy,omitempty" tf:"keep_policy,omitempty"`

	// Specifies the status of the PCI 3DS compliance certification check.
	// Values are: true and false. This parameter must be used together with tls and cipher.
	Pci3Ds *bool `json:"pci3Ds,omitempty" tf:"pci_3ds,omitempty"`

	// Specifies the status of the PCI DSS compliance certification check.
	// Values are: true and false. This parameter must be used together with tls and cipher.
	PciDss *bool `json:"pciDss,omitempty" tf:"pci_dss,omitempty"`

	// Specifies the policy ID associated with the domain. If not specified, a new policy
	// will be created automatically.
	PolicyID *string `json:"policyId,omitempty" tf:"policy_id,omitempty"`

	// The protection status of domain, 0: suspended, 1: enabled.
	// Default value is 1.
	ProtectStatus *float64 `json:"protectStatus,omitempty" tf:"protect_status,omitempty"`

	// Specifies whether a proxy is configured. Default value is false.
	Proxy *bool `json:"proxy,omitempty" tf:"proxy,omitempty"`

	// The region in which to create the dedicated mode domain resource. If omitted,
	// the provider-level region will be used. Changing this setting will push a new domain.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The server configuration list of the domain. A maximum of 80 can be configured.
	// The server block supports:
	Server []ServerInitParameters `json:"server,omitempty" tf:"server,omitempty"`

	// Specifies the minimum required TLS version.
	// Values are:
	TLS *string `json:"tls,omitempty" tf:"tls,omitempty"`

	// Specifies the timeout configuration.
	// The timeout_config structure is documented below.
	TimeoutConfig []TimeoutConfigInitParameters `json:"timeoutConfig,omitempty" tf:"timeout_config,omitempty"`
}

type DedicatedDomainV1Observation struct {

	// Whether a domain name is connected to WAF. Valid values are:
	AccessStatus *float64 `json:"accessStatus,omitempty" tf:"access_status,omitempty"`

	// The alarm page of domain. Valid values are:
	// +mapType=granular
	AlarmPage map[string]*string `json:"alarmPage,omitempty" tf:"alarm_page,omitempty"`

	// Specifies the certificate ID. This parameter is mandatory when client_protocol
	// is set to HTTPS.
	CertificateID *string `json:"certificateId,omitempty" tf:"certificate_id,omitempty"`

	// The name of the certificate used by the domain name.
	CertificateName *string `json:"certificateName,omitempty" tf:"certificate_name,omitempty"`

	// Specifies the cipher suite of domain.
	// Values are:
	Cipher *string `json:"cipher,omitempty" tf:"cipher,omitempty"`

	// The compliance certifications of the domain, values are:
	// +mapType=granular
	ComplianceCertification map[string]*bool `json:"complianceCertification,omitempty" tf:"compliance_certification,omitempty"`

	// Timestamp when the dedicated WAF domain was created.
	CreatedAt *float64 `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// Specifies the protected domain name or IP address (port allowed). For example,
	// www.example.com or *.example.com or www.example.com:89. Changing this creates a new domain.
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// ID of the domain.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies whether to retain the policy when deleting a domain name.
	// Defaults to true.
	KeepPolicy *bool `json:"keepPolicy,omitempty" tf:"keep_policy,omitempty"`

	// Specifies the status of the PCI 3DS compliance certification check.
	// Values are: true and false. This parameter must be used together with tls and cipher.
	Pci3Ds *bool `json:"pci3Ds,omitempty" tf:"pci_3ds,omitempty"`

	// Specifies the status of the PCI DSS compliance certification check.
	// Values are: true and false. This parameter must be used together with tls and cipher.
	PciDss *bool `json:"pciDss,omitempty" tf:"pci_dss,omitempty"`

	// Specifies the policy ID associated with the domain. If not specified, a new policy
	// will be created automatically.
	PolicyID *string `json:"policyId,omitempty" tf:"policy_id,omitempty"`

	// The protection status of domain, 0: suspended, 1: enabled.
	// Default value is 1.
	ProtectStatus *float64 `json:"protectStatus,omitempty" tf:"protect_status,omitempty"`

	// The protocol type of the client. The options are HTTP and HTTPS.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Specifies whether a proxy is configured. Default value is false.
	Proxy *bool `json:"proxy,omitempty" tf:"proxy,omitempty"`

	// The region in which to create the dedicated mode domain resource. If omitted,
	// the provider-level region will be used. Changing this setting will push a new domain.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The server configuration list of the domain. A maximum of 80 can be configured.
	// The server block supports:
	Server []ServerObservation `json:"server,omitempty" tf:"server,omitempty"`

	// Specifies the minimum required TLS version.
	// Values are:
	TLS *string `json:"tls,omitempty" tf:"tls,omitempty"`

	// Specifies the timeout configuration.
	// The timeout_config structure is documented below.
	TimeoutConfig []TimeoutConfigObservation `json:"timeoutConfig,omitempty" tf:"timeout_config,omitempty"`

	// The traffic identifier of domain. Valid values are:
	// +mapType=granular
	TrafficIdentifier map[string]*string `json:"trafficIdentifier,omitempty" tf:"traffic_identifier,omitempty"`
}

type DedicatedDomainV1Parameters struct {

	// Specifies the certificate ID. This parameter is mandatory when client_protocol
	// is set to HTTPS.
	// +kubebuilder:validation:Optional
	CertificateID *string `json:"certificateId,omitempty" tf:"certificate_id,omitempty"`

	// Specifies the cipher suite of domain.
	// Values are:
	// +kubebuilder:validation:Optional
	Cipher *string `json:"cipher,omitempty" tf:"cipher,omitempty"`

	// Specifies the protected domain name or IP address (port allowed). For example,
	// www.example.com or *.example.com or www.example.com:89. Changing this creates a new domain.
	// +kubebuilder:validation:Optional
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// Specifies whether to retain the policy when deleting a domain name.
	// Defaults to true.
	// +kubebuilder:validation:Optional
	KeepPolicy *bool `json:"keepPolicy,omitempty" tf:"keep_policy,omitempty"`

	// Specifies the status of the PCI 3DS compliance certification check.
	// Values are: true and false. This parameter must be used together with tls and cipher.
	// +kubebuilder:validation:Optional
	Pci3Ds *bool `json:"pci3Ds,omitempty" tf:"pci_3ds,omitempty"`

	// Specifies the status of the PCI DSS compliance certification check.
	// Values are: true and false. This parameter must be used together with tls and cipher.
	// +kubebuilder:validation:Optional
	PciDss *bool `json:"pciDss,omitempty" tf:"pci_dss,omitempty"`

	// Specifies the policy ID associated with the domain. If not specified, a new policy
	// will be created automatically.
	// +kubebuilder:validation:Optional
	PolicyID *string `json:"policyId,omitempty" tf:"policy_id,omitempty"`

	// The protection status of domain, 0: suspended, 1: enabled.
	// Default value is 1.
	// +kubebuilder:validation:Optional
	ProtectStatus *float64 `json:"protectStatus,omitempty" tf:"protect_status,omitempty"`

	// Specifies whether a proxy is configured. Default value is false.
	// +kubebuilder:validation:Optional
	Proxy *bool `json:"proxy,omitempty" tf:"proxy,omitempty"`

	// The region in which to create the dedicated mode domain resource. If omitted,
	// the provider-level region will be used. Changing this setting will push a new domain.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The server configuration list of the domain. A maximum of 80 can be configured.
	// The server block supports:
	// +kubebuilder:validation:Optional
	Server []ServerParameters `json:"server,omitempty" tf:"server,omitempty"`

	// Specifies the minimum required TLS version.
	// Values are:
	// +kubebuilder:validation:Optional
	TLS *string `json:"tls,omitempty" tf:"tls,omitempty"`

	// Specifies the timeout configuration.
	// The timeout_config structure is documented below.
	// +kubebuilder:validation:Optional
	TimeoutConfig []TimeoutConfigParameters `json:"timeoutConfig,omitempty" tf:"timeout_config,omitempty"`
}

type ServerInitParameters struct {

	// IP address or domain name of the web server that the client accesses. For
	// example, 192.168.1.1 or www.example.com. Changing this creates a new server.
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// Protocol type of the client. Values are HTTP and HTTPS.
	// Changing this creates a new server.
	ClientProtocol *string `json:"clientProtocol,omitempty" tf:"client_protocol,omitempty"`

	// Port number used by the web server. The value ranges from 0 to 65535. Changing this
	// creates a new server.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Protocol used by WAF to forward client requests to the server.
	// Values areHTTP and HTTPS. Changing this creates a new server.
	ServerProtocol *string `json:"serverProtocol,omitempty" tf:"server_protocol,omitempty"`

	// Server network type, IPv4 or IPv6. Valid values are: ipv4 and ipv6. Changing
	// this creates a new server.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The id of the vpc used by the server. Changing this creates a server.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

type ServerObservation struct {

	// IP address or domain name of the web server that the client accesses. For
	// example, 192.168.1.1 or www.example.com. Changing this creates a new server.
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// Protocol type of the client. Values are HTTP and HTTPS.
	// Changing this creates a new server.
	ClientProtocol *string `json:"clientProtocol,omitempty" tf:"client_protocol,omitempty"`

	// Port number used by the web server. The value ranges from 0 to 65535. Changing this
	// creates a new server.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Protocol used by WAF to forward client requests to the server.
	// Values areHTTP and HTTPS. Changing this creates a new server.
	ServerProtocol *string `json:"serverProtocol,omitempty" tf:"server_protocol,omitempty"`

	// Server network type, IPv4 or IPv6. Valid values are: ipv4 and ipv6. Changing
	// this creates a new server.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The id of the vpc used by the server. Changing this creates a server.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

type ServerParameters struct {

	// IP address or domain name of the web server that the client accesses. For
	// example, 192.168.1.1 or www.example.com. Changing this creates a new server.
	// +kubebuilder:validation:Optional
	Address *string `json:"address" tf:"address,omitempty"`

	// Protocol type of the client. Values are HTTP and HTTPS.
	// Changing this creates a new server.
	// +kubebuilder:validation:Optional
	ClientProtocol *string `json:"clientProtocol" tf:"client_protocol,omitempty"`

	// Port number used by the web server. The value ranges from 0 to 65535. Changing this
	// creates a new server.
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port" tf:"port,omitempty"`

	// Protocol used by WAF to forward client requests to the server.
	// Values areHTTP and HTTPS. Changing this creates a new server.
	// +kubebuilder:validation:Optional
	ServerProtocol *string `json:"serverProtocol" tf:"server_protocol,omitempty"`

	// Server network type, IPv4 or IPv6. Valid values are: ipv4 and ipv6. Changing
	// this creates a new server.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`

	// The id of the vpc used by the server. Changing this creates a server.
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId" tf:"vpc_id,omitempty"`
}

type TimeoutConfigInitParameters struct {

	// Specifies the timeout in seconds for WAF to connect to the origin server.
	ConnectTimeout *float64 `json:"connectTimeout,omitempty" tf:"connect_timeout,omitempty"`

	// Specifies the timeout in seconds for WAF to receive responses from the origin server.
	ReadTimeout *float64 `json:"readTimeout,omitempty" tf:"read_timeout,omitempty"`

	SendTimeout *float64 `json:"sendTimeout,omitempty" tf:"send_timeout,omitempty"`
}

type TimeoutConfigObservation struct {

	// Specifies the timeout in seconds for WAF to connect to the origin server.
	ConnectTimeout *float64 `json:"connectTimeout,omitempty" tf:"connect_timeout,omitempty"`

	// Specifies the timeout in seconds for WAF to receive responses from the origin server.
	ReadTimeout *float64 `json:"readTimeout,omitempty" tf:"read_timeout,omitempty"`

	SendTimeout *float64 `json:"sendTimeout,omitempty" tf:"send_timeout,omitempty"`
}

type TimeoutConfigParameters struct {

	// Specifies the timeout in seconds for WAF to connect to the origin server.
	// +kubebuilder:validation:Optional
	ConnectTimeout *float64 `json:"connectTimeout,omitempty" tf:"connect_timeout,omitempty"`

	// Specifies the timeout in seconds for WAF to receive responses from the origin server.
	// +kubebuilder:validation:Optional
	ReadTimeout *float64 `json:"readTimeout,omitempty" tf:"read_timeout,omitempty"`

	// +kubebuilder:validation:Optional
	SendTimeout *float64 `json:"sendTimeout,omitempty" tf:"send_timeout,omitempty"`
}

// DedicatedDomainV1Spec defines the desired state of DedicatedDomainV1
type DedicatedDomainV1Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DedicatedDomainV1Parameters `json:"forProvider"`
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
	InitProvider DedicatedDomainV1InitParameters `json:"initProvider,omitempty"`
}

// DedicatedDomainV1Status defines the observed state of DedicatedDomainV1.
type DedicatedDomainV1Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DedicatedDomainV1Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// DedicatedDomainV1 is the Schema for the DedicatedDomainV1s API. Manages a WAF Dedicated Domain resource within OpenTelekomCloud.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,opentelekomcloud}
type DedicatedDomainV1 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.domain) || (has(self.initProvider) && has(self.initProvider.domain))",message="spec.forProvider.domain is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.server) || (has(self.initProvider) && has(self.initProvider.server))",message="spec.forProvider.server is a required parameter"
	Spec   DedicatedDomainV1Spec   `json:"spec"`
	Status DedicatedDomainV1Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DedicatedDomainV1List contains a list of DedicatedDomainV1s
type DedicatedDomainV1List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DedicatedDomainV1 `json:"items"`
}

// Repository type metadata.
var (
	DedicatedDomainV1_Kind             = "DedicatedDomainV1"
	DedicatedDomainV1_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DedicatedDomainV1_Kind}.String()
	DedicatedDomainV1_KindAPIVersion   = DedicatedDomainV1_Kind + "." + CRDGroupVersion.String()
	DedicatedDomainV1_GroupVersionKind = CRDGroupVersion.WithKind(DedicatedDomainV1_Kind)
)

func init() {
	SchemeBuilder.Register(&DedicatedDomainV1{}, &DedicatedDomainV1List{})
}
