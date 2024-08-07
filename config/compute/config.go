package compute

import (
	"github.com/crossplane/upjet/pkg/config"
	"github.com/opentelekomcloud/provider-opentelekomcloud/config/common"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("opentelekomcloud_compute_instance_v2", func(r *config.Resource) {
		r.UseAsync = true
		r.References["key_pair"] = config.Reference{
			TerraformName: "opentelekomcloud_compute_keypair_v2",
		}
		r.References["security_groups"] = config.Reference{
			TerraformName: "opentelekomcloud_compute_secgroup_v2",
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
			Extractor:     common.NetworkIDExtractor,
		}
	})
}
