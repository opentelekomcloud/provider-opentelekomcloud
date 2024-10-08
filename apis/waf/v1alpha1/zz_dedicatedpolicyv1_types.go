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

type DedicatedPolicyV1InitParameters struct {

	// The deep inspection in basic web protection.
	DeepInspection *bool `json:"deepInspection,omitempty" tf:"deep_inspection,omitempty"`

	// Specifies the detection mode in Precise Protection.
	FullDetection *bool `json:"fullDetection,omitempty" tf:"full_detection,omitempty"`

	// The header inspection in basic web protection.
	HeaderInspection *bool `json:"headerInspection,omitempty" tf:"header_inspection,omitempty"`

	// Specifies the protection level.
	// Values are:
	Level *float64 `json:"level,omitempty" tf:"level,omitempty"`

	// The policy name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the protection switches.
	// The options block supports:
	Options []OptionsInitParameters `json:"options,omitempty" tf:"options,omitempty"`

	// Specifies the protective action after a rule is matched.
	// Values are:
	ProtectionMode *string `json:"protectionMode,omitempty" tf:"protection_mode,omitempty"`

	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The shiro decryption check in basic web protection.
	ShiroDecryptionCheck *bool `json:"shiroDecryptionCheck,omitempty" tf:"shiro_decryption_check,omitempty"`
}

type DedicatedPolicyV1Observation struct {

	// Time the policy is created. The value is a 13-digit timestamp, in ms.
	CreatedAt *float64 `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// The deep inspection in basic web protection.
	DeepInspection *bool `json:"deepInspection,omitempty" tf:"deep_inspection,omitempty"`

	// Specifies the domain IDs.
	Domains []*string `json:"domains,omitempty" tf:"domains,omitempty"`

	// Specifies the detection mode in Precise Protection.
	FullDetection *bool `json:"fullDetection,omitempty" tf:"full_detection,omitempty"`

	// The header inspection in basic web protection.
	HeaderInspection *bool `json:"headerInspection,omitempty" tf:"header_inspection,omitempty"`

	// ID of the policy.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the protection level.
	// Values are:
	Level *float64 `json:"level,omitempty" tf:"level,omitempty"`

	// The policy name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the protection switches.
	// The options block supports:
	Options []OptionsObservation `json:"options,omitempty" tf:"options,omitempty"`

	// Specifies the protective action after a rule is matched.
	// Values are:
	ProtectionMode *string `json:"protectionMode,omitempty" tf:"protection_mode,omitempty"`

	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The shiro decryption check in basic web protection.
	ShiroDecryptionCheck *bool `json:"shiroDecryptionCheck,omitempty" tf:"shiro_decryption_check,omitempty"`
}

type DedicatedPolicyV1Parameters struct {

	// The deep inspection in basic web protection.
	// +kubebuilder:validation:Optional
	DeepInspection *bool `json:"deepInspection,omitempty" tf:"deep_inspection,omitempty"`

	// Specifies the detection mode in Precise Protection.
	// +kubebuilder:validation:Optional
	FullDetection *bool `json:"fullDetection,omitempty" tf:"full_detection,omitempty"`

	// The header inspection in basic web protection.
	// +kubebuilder:validation:Optional
	HeaderInspection *bool `json:"headerInspection,omitempty" tf:"header_inspection,omitempty"`

	// Specifies the protection level.
	// Values are:
	// +kubebuilder:validation:Optional
	Level *float64 `json:"level,omitempty" tf:"level,omitempty"`

	// The policy name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the protection switches.
	// The options block supports:
	// +kubebuilder:validation:Optional
	Options []OptionsParameters `json:"options,omitempty" tf:"options,omitempty"`

	// Specifies the protective action after a rule is matched.
	// Values are:
	// +kubebuilder:validation:Optional
	ProtectionMode *string `json:"protectionMode,omitempty" tf:"protection_mode,omitempty"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The shiro decryption check in basic web protection.
	// +kubebuilder:validation:Optional
	ShiroDecryptionCheck *bool `json:"shiroDecryptionCheck,omitempty" tf:"shiro_decryption_check,omitempty"`
}

type OptionsInitParameters struct {

	// JavaScript anti-crawler function.
	AntiCrawler *bool `json:"antiCrawler,omitempty" tf:"anti_crawler,omitempty"`

	// Whether the information leakage prevention is enabled.
	AntiLeakage *bool `json:"antiLeakage,omitempty" tf:"anti_leakage,omitempty"`

	// Specifies whether Web Tamper Protection is enabled.
	AntiTamper *bool `json:"antiTamper,omitempty" tf:"anti_tamper,omitempty"`

	// Specifies whether Blacklist and Whitelist is enabled.
	Blacklist *bool `json:"blacklist,omitempty" tf:"blacklist,omitempty"`

	// Specifies whether CC Attack Protection is enabled.
	Cc *bool `json:"cc,omitempty" tf:"cc,omitempty"`

	// Specifies whether General Check in Basic Web Protection is enabled.
	Common *bool `json:"common,omitempty" tf:"common,omitempty"`

	// Specifies whether the master crawler detection switch in Basic Web Protection is enabled.
	Crawler *bool `json:"crawler,omitempty" tf:"crawler,omitempty"`

	// Specifies whether the Search Engine switch in Basic Web Protection is enabled.
	CrawlerEngine *bool `json:"crawlerEngine,omitempty" tf:"crawler_engine,omitempty"`

	// Specifies whether detection of other crawlers in Basic Web Protection is enabled.
	CrawlerOther *bool `json:"crawlerOther,omitempty" tf:"crawler_other,omitempty"`

	// Specifies whether the Scanner switch in Basic Web Protection is enabled.
	CrawlerScanner *bool `json:"crawlerScanner,omitempty" tf:"crawler_scanner,omitempty"`

	// Specifies whether the Script Tool switch in Basic Web Protection is enabled.
	CrawlerScript *bool `json:"crawlerScript,omitempty" tf:"crawler_script,omitempty"`

	// Specifies whether Precise Protection is enabled.
	Custom *bool `json:"custom,omitempty" tf:"custom,omitempty"`

	// Whether the Known Attack Source protection is enabled.
	FollowedAction *bool `json:"followedAction,omitempty" tf:"followed_action,omitempty"`

	// Whether geolocation access control is enabled.
	GeolocationAccessControl *bool `json:"geolocationAccessControl,omitempty" tf:"geolocation_access_control,omitempty"`

	// Whether false alarm masking is enabled.
	Ignore *bool `json:"ignore,omitempty" tf:"ignore,omitempty"`

	// Specifies whether Data Masking is enabled.
	Privacy *bool `json:"privacy,omitempty" tf:"privacy,omitempty"`

	// Specifies whether Basic Web Protection is enabled.
	WebAttack *bool `json:"webAttack,omitempty" tf:"web_attack,omitempty"`

	// Specifies whether webshell detection in Basic Web Protection is enabled.
	WebShell *bool `json:"webShell,omitempty" tf:"web_shell,omitempty"`
}

type OptionsObservation struct {

	// JavaScript anti-crawler function.
	AntiCrawler *bool `json:"antiCrawler,omitempty" tf:"anti_crawler,omitempty"`

	// Whether the information leakage prevention is enabled.
	AntiLeakage *bool `json:"antiLeakage,omitempty" tf:"anti_leakage,omitempty"`

	// Specifies whether Web Tamper Protection is enabled.
	AntiTamper *bool `json:"antiTamper,omitempty" tf:"anti_tamper,omitempty"`

	// Specifies whether Blacklist and Whitelist is enabled.
	Blacklist *bool `json:"blacklist,omitempty" tf:"blacklist,omitempty"`

	BotEnable *bool `json:"botEnable,omitempty" tf:"bot_enable,omitempty"`

	// Specifies whether CC Attack Protection is enabled.
	Cc *bool `json:"cc,omitempty" tf:"cc,omitempty"`

	// Specifies whether General Check in Basic Web Protection is enabled.
	Common *bool `json:"common,omitempty" tf:"common,omitempty"`

	// Specifies whether the master crawler detection switch in Basic Web Protection is enabled.
	Crawler *bool `json:"crawler,omitempty" tf:"crawler,omitempty"`

	// Specifies whether the Search Engine switch in Basic Web Protection is enabled.
	CrawlerEngine *bool `json:"crawlerEngine,omitempty" tf:"crawler_engine,omitempty"`

	// Specifies whether detection of other crawlers in Basic Web Protection is enabled.
	CrawlerOther *bool `json:"crawlerOther,omitempty" tf:"crawler_other,omitempty"`

	// Specifies whether the Scanner switch in Basic Web Protection is enabled.
	CrawlerScanner *bool `json:"crawlerScanner,omitempty" tf:"crawler_scanner,omitempty"`

	// Specifies whether the Script Tool switch in Basic Web Protection is enabled.
	CrawlerScript *bool `json:"crawlerScript,omitempty" tf:"crawler_script,omitempty"`

	// Specifies whether Precise Protection is enabled.
	Custom *bool `json:"custom,omitempty" tf:"custom,omitempty"`

	// Whether the Known Attack Source protection is enabled.
	FollowedAction *bool `json:"followedAction,omitempty" tf:"followed_action,omitempty"`

	// Whether geolocation access control is enabled.
	GeolocationAccessControl *bool `json:"geolocationAccessControl,omitempty" tf:"geolocation_access_control,omitempty"`

	// Whether false alarm masking is enabled.
	Ignore *bool `json:"ignore,omitempty" tf:"ignore,omitempty"`

	Precise *bool `json:"precise,omitempty" tf:"precise,omitempty"`

	// Specifies whether Data Masking is enabled.
	Privacy *bool `json:"privacy,omitempty" tf:"privacy,omitempty"`

	// Specifies whether Basic Web Protection is enabled.
	WebAttack *bool `json:"webAttack,omitempty" tf:"web_attack,omitempty"`

	// Specifies whether webshell detection in Basic Web Protection is enabled.
	WebShell *bool `json:"webShell,omitempty" tf:"web_shell,omitempty"`
}

type OptionsParameters struct {

	// JavaScript anti-crawler function.
	// +kubebuilder:validation:Optional
	AntiCrawler *bool `json:"antiCrawler,omitempty" tf:"anti_crawler,omitempty"`

	// Whether the information leakage prevention is enabled.
	// +kubebuilder:validation:Optional
	AntiLeakage *bool `json:"antiLeakage,omitempty" tf:"anti_leakage,omitempty"`

	// Specifies whether Web Tamper Protection is enabled.
	// +kubebuilder:validation:Optional
	AntiTamper *bool `json:"antiTamper,omitempty" tf:"anti_tamper,omitempty"`

	// Specifies whether Blacklist and Whitelist is enabled.
	// +kubebuilder:validation:Optional
	Blacklist *bool `json:"blacklist,omitempty" tf:"blacklist,omitempty"`

	// Specifies whether CC Attack Protection is enabled.
	// +kubebuilder:validation:Optional
	Cc *bool `json:"cc,omitempty" tf:"cc,omitempty"`

	// Specifies whether General Check in Basic Web Protection is enabled.
	// +kubebuilder:validation:Optional
	Common *bool `json:"common,omitempty" tf:"common,omitempty"`

	// Specifies whether the master crawler detection switch in Basic Web Protection is enabled.
	// +kubebuilder:validation:Optional
	Crawler *bool `json:"crawler,omitempty" tf:"crawler,omitempty"`

	// Specifies whether the Search Engine switch in Basic Web Protection is enabled.
	// +kubebuilder:validation:Optional
	CrawlerEngine *bool `json:"crawlerEngine,omitempty" tf:"crawler_engine,omitempty"`

	// Specifies whether detection of other crawlers in Basic Web Protection is enabled.
	// +kubebuilder:validation:Optional
	CrawlerOther *bool `json:"crawlerOther,omitempty" tf:"crawler_other,omitempty"`

	// Specifies whether the Scanner switch in Basic Web Protection is enabled.
	// +kubebuilder:validation:Optional
	CrawlerScanner *bool `json:"crawlerScanner,omitempty" tf:"crawler_scanner,omitempty"`

	// Specifies whether the Script Tool switch in Basic Web Protection is enabled.
	// +kubebuilder:validation:Optional
	CrawlerScript *bool `json:"crawlerScript,omitempty" tf:"crawler_script,omitempty"`

	// Specifies whether Precise Protection is enabled.
	// +kubebuilder:validation:Optional
	Custom *bool `json:"custom,omitempty" tf:"custom,omitempty"`

	// Whether the Known Attack Source protection is enabled.
	// +kubebuilder:validation:Optional
	FollowedAction *bool `json:"followedAction,omitempty" tf:"followed_action,omitempty"`

	// Whether geolocation access control is enabled.
	// +kubebuilder:validation:Optional
	GeolocationAccessControl *bool `json:"geolocationAccessControl,omitempty" tf:"geolocation_access_control,omitempty"`

	// Whether false alarm masking is enabled.
	// +kubebuilder:validation:Optional
	Ignore *bool `json:"ignore,omitempty" tf:"ignore,omitempty"`

	// Specifies whether Data Masking is enabled.
	// +kubebuilder:validation:Optional
	Privacy *bool `json:"privacy,omitempty" tf:"privacy,omitempty"`

	// Specifies whether Basic Web Protection is enabled.
	// +kubebuilder:validation:Optional
	WebAttack *bool `json:"webAttack,omitempty" tf:"web_attack,omitempty"`

	// Specifies whether webshell detection in Basic Web Protection is enabled.
	// +kubebuilder:validation:Optional
	WebShell *bool `json:"webShell,omitempty" tf:"web_shell,omitempty"`
}

// DedicatedPolicyV1Spec defines the desired state of DedicatedPolicyV1
type DedicatedPolicyV1Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DedicatedPolicyV1Parameters `json:"forProvider"`
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
	InitProvider DedicatedPolicyV1InitParameters `json:"initProvider,omitempty"`
}

// DedicatedPolicyV1Status defines the observed state of DedicatedPolicyV1.
type DedicatedPolicyV1Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DedicatedPolicyV1Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// DedicatedPolicyV1 is the Schema for the DedicatedPolicyV1s API. Manages a WAF Dedicated Policy resource within OpenTelekomCloud.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,opentelekomcloud}
type DedicatedPolicyV1 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   DedicatedPolicyV1Spec   `json:"spec"`
	Status DedicatedPolicyV1Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DedicatedPolicyV1List contains a list of DedicatedPolicyV1s
type DedicatedPolicyV1List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DedicatedPolicyV1 `json:"items"`
}

// Repository type metadata.
var (
	DedicatedPolicyV1_Kind             = "DedicatedPolicyV1"
	DedicatedPolicyV1_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DedicatedPolicyV1_Kind}.String()
	DedicatedPolicyV1_KindAPIVersion   = DedicatedPolicyV1_Kind + "." + CRDGroupVersion.String()
	DedicatedPolicyV1_GroupVersionKind = CRDGroupVersion.WithKind(DedicatedPolicyV1_Kind)
)

func init() {
	SchemeBuilder.Register(&DedicatedPolicyV1{}, &DedicatedPolicyV1List{})
}
