package networking

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/opentelekomcloud/provider-opentelekomcloud/config/common"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("opentelekomcloud_networking_subnet_v2", func(r *config.Resource) {
		r.References["network_id"] = config.Reference{
			TerraformName: "opentelekomcloud_networking_network_v2",
		}
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"allocation_pools"},
		}
	})
	p.AddResourceConfigurator("opentelekomcloud_networking_router_interface_v2", func(r *config.Resource) {
		r.References["router_id"] = config.Reference{
			TerraformName: "opentelekomcloud_networking_router_v2",
		}
		r.References["subnet_id"] = config.Reference{
			TerraformName: "opentelekomcloud_networking_subnet_v2",
		}
	})
	p.AddResourceConfigurator("opentelekomcloud_networking_router_v2", func(r *config.Resource) {
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"external_gateway"},
		}
	})
	p.AddResourceConfigurator("opentelekomcloud_networking_floatingip_associate_v2", func(r *config.Resource) {
		r.UseAsync = true
		r.References["port_id"] = config.Reference{
			TerraformName: "opentelekomcloud_networking_port_v2",
		}
		r.References["floating_ip"] = config.Reference{
			TerraformName: "opentelekomcloud_networking_floatingip_v2",
			Extractor:     common.FipAddressExtractor,
		}
	})
	p.AddResourceConfigurator("opentelekomcloud_networking_port_v2", func(r *config.Resource) {
		r.UseAsync = true
		r.References["network_id"] = config.Reference{
			TerraformName: "opentelekomcloud_networking_network_v2",
		}
		r.References["fixed_ip.subnet_id"] = config.Reference{
			TerraformName: "opentelekomcloud_networking_subnet_v2",
		}
		r.References["security_group_ids"] = config.Reference{
			TerraformName: "opentelekomcloud_networking_secgroup_v2",
		}
	})
	p.AddResourceConfigurator("opentelekomcloud_networking_port_secgroup_associate_v2", func(r *config.Resource) {
		r.UseAsync = true
		r.References["port_id"] = config.Reference{
			TerraformName: "opentelekomcloud_networking_port_v2",
		}
		r.References["security_group_ids"] = config.Reference{
			TerraformName: "opentelekomcloud_networking_secgroup_v2",
		}
	})
	p.AddResourceConfigurator("opentelekomcloud_networking_router_interface_v2", func(r *config.Resource) {
		r.UseAsync = true
		r.References["router_id"] = config.Reference{
			TerraformName: "opentelekomcloud_networking_router_v2",
		}
		r.References["subnet_id"] = config.Reference{
			TerraformName: "opentelekomcloud_networking_subnet_v2",
		}
	})
	p.AddResourceConfigurator("opentelekomcloud_networking_router_route_v2", func(r *config.Resource) {
		r.UseAsync = true
		r.References["router_id"] = config.Reference{
			TerraformName: "opentelekomcloud_networking_router_v2",
		}
	})
	p.AddResourceConfigurator("opentelekomcloud_networking_secgroup_rule_v2", func(r *config.Resource) {
		r.UseAsync = true
		r.References["security_group_id"] = config.Reference{
			TerraformName: "opentelekomcloud_networking_secgroup_v2",
		}
	})
	p.AddResourceConfigurator("opentelekomcloud_networking_vip_v2", func(r *config.Resource) {
		r.UseAsync = true
		r.References["network_id"] = config.Reference{
			TerraformName: "opentelekomcloud_networking_network_v2",
		}
		r.References["subnet_id"] = config.Reference{
			TerraformName: "opentelekomcloud_networking_subnet_v2",
		}
	})
	p.AddResourceConfigurator("opentelekomcloud_networking_vip_associate_v2", func(r *config.Resource) {
		r.UseAsync = true
		r.References["vip_id"] = config.Reference{
			TerraformName: "opentelekomcloud_networking_vip_v2",
		}
		r.References["port_ids"] = config.Reference{
			TerraformName: "opentelekomcloud_networking_port_v2",
		}
	})
}
