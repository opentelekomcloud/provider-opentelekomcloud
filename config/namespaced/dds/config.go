package dds

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/opentelekomcloud/provider-opentelekomcloud/config/common"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("opentelekomcloud_dds_instance_v3", func(r *config.Resource) {
		r.UseAsync = true
		r.References["security_group_id"] = config.Reference{
			TerraformName:     "opentelekomcloud_compute_secgroup_v2",
			SelectorFieldName: "ComputeSecurityGroupSelector",
			RefFieldName:      "ComputeSecurityGroupRefs",
		}
		r.References["vpc_id"] = config.Reference{
			TerraformName:     "opentelekomcloud_vpc_v1",
			SelectorFieldName: "VPCSelector",
			RefFieldName:      "VPCRef",
		}
		r.References["subnet_id"] = config.Reference{
			TerraformName:     "opentelekomcloud_vpc_subnet_v1",
			Extractor:         common.NetworkIDExtractor,
			SelectorFieldName: "SubnetSelector",
			RefFieldName:      "SubnetRef",
		}
	})
	p.AddResourceConfigurator("opentelekomcloud_dds_backup_v3", func(r *config.Resource) {
		r.UseAsync = true
		r.References["instance_id"] = config.Reference{
			TerraformName:     "opentelekomcloud_dds_instance_v3",
			SelectorFieldName: "InstanceSelector",
			RefFieldName:      "InstanceRef",
		}
	})
	p.AddResourceConfigurator("opentelekomcloud_dds_public_ip_associate_v3", func(r *config.Resource) {
		r.UseAsync = true
		r.References["public_ip"] = config.Reference{
			TerraformName:     "opentelekomcloud_vpc_eip_v1",
			SelectorFieldName: "PublicIPSelector",
			RefFieldName:      "PublicIPSelectorRef",
			Extractor:         common.EipAddressExtractor,
		}
		r.References["public_ip_id"] = config.Reference{
			TerraformName:     "opentelekomcloud_vpc_eip_v1",
			SelectorFieldName: "PublicIPSelectorByID",
			RefFieldName:      "PublicIPSelectorByIDRef",
		}
	})
}
