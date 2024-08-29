package dds

import (
	"github.com/crossplane/upjet/pkg/config"
	"github.com/opentelekomcloud/provider-opentelekomcloud/config/common"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("opentelekomcloud_dds_instance_v3", func(r *config.Resource) {
		r.UseAsync = true
		r.References["security_group_id"] = config.Reference{
			TerraformName: "opentelekomcloud_compute_secgroup_v2",
		}
		r.References["vpc_id"] = config.Reference{
			TerraformName: "opentelekomcloud_vpc_v1",
		}
		r.References["subnet_id"] = config.Reference{
			TerraformName: "opentelekomcloud_vpc_subnet_v1",
			Extractor:     common.NetworkIDExtractor,
		}
	})
	p.AddResourceConfigurator("opentelekomcloud_dds_backup_v3", func(r *config.Resource) {
		r.UseAsync = true
		r.References["instance_id"] = config.Reference{
			TerraformName: "opentelekomcloud_dds_instance_v3",
		}
	})
}
