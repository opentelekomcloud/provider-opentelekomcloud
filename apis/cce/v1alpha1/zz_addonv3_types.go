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

type AddonV3InitParameters struct {

	// ID of cluster to install the add-on on.
	// +crossplane:generate:reference:type=github.com/opentelekomcloud/provider-opentelekomcloud/apis/cce/v1alpha1.ClusterV3
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Reference to a ClusterV3 in cce to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDRef *v1.Reference `json:"clusterIdRef,omitempty" tf:"-"`

	// Selector for a ClusterV3 in cce to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDSelector *v1.Selector `json:"clusterIdSelector,omitempty" tf:"-"`

	// Name of the add-on template to be installed, for example, coredns.
	TemplateName *string `json:"templateName,omitempty" tf:"template_name,omitempty"`

	// Version number of the add-on to be installed or upgraded, for example, v1.0.0.
	TemplateVersion *string `json:"templateVersion,omitempty" tf:"template_version,omitempty"`

	// Parameters of the template to be installed or upgraded.
	Values []ValuesInitParameters `json:"values,omitempty" tf:"values,omitempty"`
}

type AddonV3Observation struct {

	// ID of cluster to install the add-on on.
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Installed add-on description
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Installed add-on name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Name of the add-on template to be installed, for example, coredns.
	TemplateName *string `json:"templateName,omitempty" tf:"template_name,omitempty"`

	// Version number of the add-on to be installed or upgraded, for example, v1.0.0.
	TemplateVersion *string `json:"templateVersion,omitempty" tf:"template_version,omitempty"`

	// Parameters of the template to be installed or upgraded.
	Values []ValuesObservation `json:"values,omitempty" tf:"values,omitempty"`
}

type AddonV3Parameters struct {

	// ID of cluster to install the add-on on.
	// +crossplane:generate:reference:type=github.com/opentelekomcloud/provider-opentelekomcloud/apis/cce/v1alpha1.ClusterV3
	// +kubebuilder:validation:Optional
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Reference to a ClusterV3 in cce to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDRef *v1.Reference `json:"clusterIdRef,omitempty" tf:"-"`

	// Selector for a ClusterV3 in cce to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDSelector *v1.Selector `json:"clusterIdSelector,omitempty" tf:"-"`

	// Name of the add-on template to be installed, for example, coredns.
	// +kubebuilder:validation:Optional
	TemplateName *string `json:"templateName,omitempty" tf:"template_name,omitempty"`

	// Version number of the add-on to be installed or upgraded, for example, v1.0.0.
	// +kubebuilder:validation:Optional
	TemplateVersion *string `json:"templateVersion,omitempty" tf:"template_version,omitempty"`

	// Parameters of the template to be installed or upgraded.
	// +kubebuilder:validation:Optional
	Values []ValuesParameters `json:"values,omitempty" tf:"values,omitempty"`
}

type ValuesInitParameters struct {

	// Basic add-on information.
	// +mapType=granular
	Basic map[string]*string `json:"basic,omitempty" tf:"basic,omitempty"`

	// Custom parameters of the add-on.
	// +mapType=granular
	Custom map[string]*string `json:"custom,omitempty" tf:"custom,omitempty"`

	// Specifies the json string vary depending on the add-on.
	Flavor *string `json:"flavor,omitempty" tf:"flavor,omitempty"`
}

type ValuesObservation struct {

	// Basic add-on information.
	// +mapType=granular
	Basic map[string]*string `json:"basic,omitempty" tf:"basic,omitempty"`

	// Custom parameters of the add-on.
	// +mapType=granular
	Custom map[string]*string `json:"custom,omitempty" tf:"custom,omitempty"`

	// Specifies the json string vary depending on the add-on.
	Flavor *string `json:"flavor,omitempty" tf:"flavor,omitempty"`
}

type ValuesParameters struct {

	// Basic add-on information.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Basic map[string]*string `json:"basic" tf:"basic,omitempty"`

	// Custom parameters of the add-on.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Custom map[string]*string `json:"custom" tf:"custom,omitempty"`

	// Specifies the json string vary depending on the add-on.
	// +kubebuilder:validation:Optional
	Flavor *string `json:"flavor,omitempty" tf:"flavor,omitempty"`
}

// AddonV3Spec defines the desired state of AddonV3
type AddonV3Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AddonV3Parameters `json:"forProvider"`
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
	InitProvider AddonV3InitParameters `json:"initProvider,omitempty"`
}

// AddonV3Status defines the observed state of AddonV3.
type AddonV3Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AddonV3Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// AddonV3 is the Schema for the AddonV3s API. Manages a CCE Addon resource within OpenTelekomCloud.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,opentelekomcloud}
type AddonV3 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.templateName) || (has(self.initProvider) && has(self.initProvider.templateName))",message="spec.forProvider.templateName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.templateVersion) || (has(self.initProvider) && has(self.initProvider.templateVersion))",message="spec.forProvider.templateVersion is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.values) || (has(self.initProvider) && has(self.initProvider.values))",message="spec.forProvider.values is a required parameter"
	Spec   AddonV3Spec   `json:"spec"`
	Status AddonV3Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AddonV3List contains a list of AddonV3s
type AddonV3List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AddonV3 `json:"items"`
}

// Repository type metadata.
var (
	AddonV3_Kind             = "AddonV3"
	AddonV3_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AddonV3_Kind}.String()
	AddonV3_KindAPIVersion   = AddonV3_Kind + "." + CRDGroupVersion.String()
	AddonV3_GroupVersionKind = CRDGroupVersion.WithKind(AddonV3_Kind)
)

func init() {
	SchemeBuilder.Register(&AddonV3{}, &AddonV3List{})
}
