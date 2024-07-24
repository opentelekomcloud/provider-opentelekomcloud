package compute

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("opentelekomcloud_compute_instance_v2", func(r *config.Resource) {
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
		r.References["key_name"] = config.Reference{
			TerraformName: "opentelekomcloud_compute_keypair_v2",
		}
		r.References["security_groups"] = config.Reference{
			TerraformName:     "opentelekomcloud_compute_secgroup_v2",
			RefFieldName:      "ComputeSecurityGroupIDRefs",
			SelectorFieldName: "ComputeSecurityGroupIDSelector",
		}
	})
}
