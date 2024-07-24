package compute

import (
	xpref "github.com/crossplane/crossplane-runtime/pkg/reference"
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/crossplane/upjet/pkg/config"
	"github.com/crossplane/upjet/pkg/resource"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("opentelekomcloud_compute_instance_v2", func(r *config.Resource) {
		r.UseAsync = true
		r.References["key_pair"] = config.Reference{
			Type: "KeypairV2",
		}
		r.References["security_groups"] = config.Reference{
			Type: "SecgroupV2",
		}
	})
	p.AddResourceConfigurator("opentelekomcloud_compute_keypair_v2", func(r *config.Resource) {
		r.UseAsync = true
	})
	p.AddResourceConfigurator("opentelekomcloud_ecs_instance_v1", func(r *config.Resource) {
		r.UseAsync = true
		r.References["key_name"] = config.Reference{
			TerraformName: "opentelekomcloud_compute_keypair_v2",
		}
		r.References["security_groups"] = config.Reference{
			TerraformName:     "opentelekomcloud_compute_secgroup_v2",
			RefFieldName:      "ComputeSecurityGroupIDRefs",
			SelectorFieldName: "ComputeSecurityGroupIDSelector",
		}
		r.References["vpc_id"] = config.Reference{
			TerraformName: "opentelekomcloud_vpc_v1",
		}
		r.References["nics.network_id"] = config.Reference{
			TerraformName: "opentelekomcloud_vpc_subnet_v1",
			Extractor:     ExtractNetworkIdFunc,
		}
	})
}

const (
	// APISPackagePath is the package path for generated APIs root package
	APISPackagePath = "github.com/opentelekomcloud/provider-opentelekomcloud/config/compute"

	// ExtractNetworkIdFunc extracts network_id from subnet resource
	ExtractNetworkIdFunc = APISPackagePath + ".ExtractNetworkId()"
)

// ExtractNetworkId extracts the value of `spec.forProvider.network_id`
// from a Observable resource. If mr is not a Observable
// resource, returns an empty string.
func ExtractNetworkId() xpref.ExtractValueFn {
	return func(mr xpresource.Managed) string {
		tr, ok := mr.(resource.Observable)
		if !ok {
			return ""
		}
		o, err := tr.GetObservation()
		if err != nil {
			return ""
		}

		if k := o["network_id"]; k != nil {
			return k.(string)
		}
		return ""
	}
}
