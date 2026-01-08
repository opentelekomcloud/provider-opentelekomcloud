package nat

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/opentelekomcloud/provider-opentelekomcloud/config/common"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("opentelekomcloud_nat_gateway_v2", func(r *config.Resource) {
		r.UseAsync = true
		r.References["internal_network_id"] = config.Reference{
			TerraformName: "opentelekomcloud_vpc_subnet_v1",
			Extractor:     common.NetworkIDExtractor,
		}
		r.References["router_id"] = config.Reference{
			TerraformName: "opentelekomcloud_vpc_v1",
		}
	})
	p.AddResourceConfigurator("opentelekomcloud_nat_dnat_rule_v2", func(r *config.Resource) {
		r.UseAsync = true
		r.References["floating_ip_id"] = config.Reference{
			TerraformName: "opentelekomcloud_vpc_eip_v1",
		}
		r.References["port_id"] = config.Reference{
			TerraformName: "opentelekomcloud_networking_port_v2",
		}
		r.References["nat_gateway_id"] = config.Reference{
			TerraformName: "opentelekomcloud_nat_gateway_v2",
		}
	})
	p.AddResourceConfigurator("opentelekomcloud_nat_snat_rule_v2", func(r *config.Resource) {
		r.UseAsync = true
		r.References["floating_ip_id"] = config.Reference{
			TerraformName: "opentelekomcloud_vpc_eip_v1",
		}
		r.References["network_id"] = config.Reference{
			TerraformName: "opentelekomcloud_vpc_subnet_v1",
			Extractor:     common.NetworkIDExtractor,
		}
		r.References["nat_gateway_id"] = config.Reference{
			TerraformName: "opentelekomcloud_nat_gateway_v2",
		}
	})
}
