package cce

import (
	"github.com/crossplane/upjet/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/opentelekomcloud/provider-opentelekomcloud/config/common"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("opentelekomcloud_cce_cluster_v3", func(r *config.Resource) {
		r.UseAsync = true
		r.References["vpc_id"] = config.Reference{
			TerraformName: "opentelekomcloud_vpc_v1",
		}
		r.References["subnet_id"] = config.Reference{
			TerraformName: "opentelekomcloud_vpc_subnet_v1",
			Extractor:     common.NetworkIDExtractor,
		}
		r.References["highway_subnet_id"] = config.Reference{
			TerraformName: "opentelekomcloud_vpc_subnet_v1",
			Extractor:     common.NetworkIDExtractor,
		}
		r.References["eni_subnet_id"] = config.Reference{
			TerraformName: "opentelekomcloud_vpc_subnet_v1",
			Extractor:     common.SubnetIDExtractor,
		}
		r.References["eni_subnet_cidr"] = config.Reference{
			TerraformName: "opentelekomcloud_vpc_subnet_v1",
			Extractor:     common.SubnetCIDRExtractor,
		}
		r.References["eip"] = config.Reference{
			TerraformName: "opentelekomcloud_vpc_eip_v1",
			Extractor:     common.EipAddressExtractor,
		}
		r.TerraformCustomDiff = func(diff *terraform.InstanceDiff, _ *terraform.InstanceState, _ *terraform.ResourceConfig) (*terraform.InstanceDiff, error) {
			if ipsDiff, ok := diff.Attributes["eip"]; ok && ipsDiff.New == "" {
				delete(diff.Attributes, "eip")
			}
			return diff, nil
		}
	})
	p.AddResourceConfigurator("opentelekomcloud_cce_node_v3", func(r *config.Resource) {
		r.UseAsync = true
		r.References["cluster_id"] = config.Reference{
			TerraformName: "opentelekomcloud_cce_cluster_v3",
		}
		r.References["eip_ids"] = config.Reference{
			TerraformName: "opentelekomcloud_vpc_eip_v1",
		}
		r.References["key_pair"] = config.Reference{
			TerraformName: "opentelekomcloud_compute_keypair_v2",
		}
		r.References["root_volume.kms_id"] = config.Reference{
			TerraformName: "opentelekomcloud_kms_key_v1",
		}
		r.References["data_volumes.kms_id"] = config.Reference{
			TerraformName: "opentelekomcloud_kms_key_v1",
		}
		r.References["agency_name"] = config.Reference{
			TerraformName: "opentelekomcloud_identity_agency_v3",
			Extractor:     common.AgencyNameExtractor,
		}
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{
				"iptype",
				"bandwidth_charge_mode",
				"bandwidth_size",
				"sharetype",
				"eip_count",
			},
		}
		r.TerraformCustomDiff = func(diff *terraform.InstanceDiff, _ *terraform.InstanceState, _ *terraform.ResourceConfig) (*terraform.InstanceDiff, error) {
			if ipsDiff, ok := diff.Attributes["public_ip"]; ok && ipsDiff.New == "" {
				delete(diff.Attributes, "public_ip")
			}
			return diff, nil
		}
	})
	p.AddResourceConfigurator("opentelekomcloud_cce_addon_v3", func(r *config.Resource) {
		r.UseAsync = true
		r.References["cluster_id"] = config.Reference{
			TerraformName: "opentelekomcloud_cce_cluster_v3",
		}
		r.References["custom.cluster_id"] = config.Reference{
			TerraformName: "opentelekomcloud_cce_cluster_v3",
		}
		r.References["custom.tenant_id"] = config.Reference{
			TerraformName: "opentelekomcloud_identity_project_v3",
		}
	})
	p.AddResourceConfigurator("opentelekomcloud_cce_node_pool_v3", func(r *config.Resource) {
		r.UseAsync = true
		r.References["cluster_id"] = config.Reference{
			TerraformName: "opentelekomcloud_cce_cluster_v3",
		}
		r.References["key_pair"] = config.Reference{
			TerraformName: "opentelekomcloud_compute_keypair_v2",
		}
		r.References["root_volume.kms_id"] = config.Reference{
			TerraformName: "opentelekomcloud_kms_key_v1",
		}
		r.References["data_volumes.kms_id"] = config.Reference{
			TerraformName: "opentelekomcloud_kms_key_v1",
		}
		r.References["agency_name"] = config.Reference{
			TerraformName: "opentelekomcloud_identity_agency_v3",
			Extractor:     common.AgencyNameExtractor,
		}
		r.References["subnet_id"] = config.Reference{
			TerraformName: "opentelekomcloud_vpc_subnet_v1",
			Extractor:     common.NetworkIDExtractor,
		}
	})
}
