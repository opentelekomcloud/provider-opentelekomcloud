package rds

import (
	"github.com/crossplane/upjet/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/opentelekomcloud/provider-opentelekomcloud/config/common"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("opentelekomcloud_rds_instance_v3", func(r *config.Resource) {
		r.UseAsync = true
		r.References["security_group_id"] = config.Reference{
			TerraformName:     "opentelekomcloud_compute_secgroup_v2",
			RefFieldName:      "ComputeSecurityGroupIDRefs",
			SelectorFieldName: "ComputeSecurityGroupIDSelector",
		}
		r.References["vpc_id"] = config.Reference{
			TerraformName: "opentelekomcloud_vpc_v1",
		}
		r.References["subnet_id"] = config.Reference{
			TerraformName: "opentelekomcloud_vpc_subnet_v1",
			Extractor:     common.NetworkIDExtractor,
		}
		r.References["public_ips"] = config.Reference{
			TerraformName:     "opentelekomcloud_vpc_eip_v1",
			Extractor:         common.EipAddressExtractor,
			RefFieldName:      "PublicIpsRefs",
			SelectorFieldName: "PublicIpsSelector",
		}
		r.References["param_group_id"] = config.Reference{
			TerraformName: "opentelekomcloud_rds_parametergroup_v3",
		}
		r.TerraformCustomDiff = func(diff *terraform.InstanceDiff, _ *terraform.InstanceState, _ *terraform.ResourceConfig) (*terraform.InstanceDiff, error) {
			if ipsDiff, ok := diff.Attributes["public_ips.#"]; ok && ipsDiff.New == "" {
				delete(diff.Attributes, "public_ips.#")
			}
			return diff, nil
		}
	})
	p.AddResourceConfigurator("opentelekomcloud_rds_backup_v3", func(r *config.Resource) {
		r.UseAsync = true
		r.References["instance_id"] = config.Reference{
			TerraformName: "opentelekomcloud_rds_instance_v3",
		}
	})
	p.AddResourceConfigurator("opentelekomcloud_rds_read_replica_v3", func(r *config.Resource) {
		r.UseAsync = true
		r.References["replica_of_id"] = config.Reference{
			TerraformName: "opentelekomcloud_rds_instance_v3",
		}
		r.References["public_ips"] = config.Reference{
			TerraformName:     "opentelekomcloud_vpc_eip_v1",
			Extractor:         common.EipAddressExtractor,
			RefFieldName:      "PublicIpsRefs",
			SelectorFieldName: "PublicIpsSelector",
		}
	})
}
