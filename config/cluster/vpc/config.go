package vpc

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("opentelekomcloud_vpc_subnet_v1", func(r *config.Resource) {
		r.UseAsync = true
		r.References["vpc_id"] = config.Reference{
			TerraformName: "opentelekomcloud_vpc_v1",
		}
	})
	p.AddResourceConfigurator("opentelekomcloud_vpc_v1", func(r *config.Resource) {
		r.UseAsync = true
		r.Kind = "VpcV1"
	})
	p.AddResourceConfigurator("opentelekomcloud_vpc_bandwidth_associate_v2", func(r *config.Resource) {
		r.UseAsync = true
		r.References["bandwidth"] = config.Reference{
			TerraformName: "opentelekomcloud_vpc_bandwidth_v2",
		}
		r.References["floating_ips"] = config.Reference{
			TerraformName: "opentelekomcloud_vpc_eip_v1",
		}
	})
	p.AddResourceConfigurator("opentelekomcloud_vpc_flow_log_v1", func(r *config.Resource) {
		r.UseAsync = true
		r.References["resource_id"] = config.Reference{
			TerraformName: "opentelekomcloud_vpc_v1",
		}
		r.References["log_group_id"] = config.Reference{
			TerraformName: "opentelekomcloud_logtank_group_v2",
		}
		r.References["log_topic_id"] = config.Reference{
			TerraformName: "opentelekomcloud_logtank_topic_v2",
		}
	})
	p.AddResourceConfigurator("opentelekomcloud_vpc_peering_connection_v2", func(r *config.Resource) {
		r.UseAsync = true
		r.References["vpc_id"] = config.Reference{
			TerraformName: "opentelekomcloud_vpc_v1",
		}
		r.References["peer_vpc_id"] = config.Reference{
			TerraformName: "opentelekomcloud_vpc_v1",
		}
	})
	p.AddResourceConfigurator("opentelekomcloud_vpc_peering_connection_accepter_v2", func(r *config.Resource) {
		r.UseAsync = true
		r.References["vpc_peering_connection_id"] = config.Reference{
			TerraformName: "opentelekomcloud_vpc_peering_connection_v2",
		}
	})
	p.AddResourceConfigurator("opentelekomcloud_vpc_route_v2", func(r *config.Resource) {
		r.UseAsync = true
		r.References["vpc_id"] = config.Reference{
			TerraformName: "opentelekomcloud_vpc_v1",
		}
	})
	p.AddResourceConfigurator("opentelekomcloud_vpc_route_table_v1", func(r *config.Resource) {
		r.UseAsync = true
		r.References["vpc_id"] = config.Reference{
			TerraformName: "opentelekomcloud_vpc_v1",
		}
	})
}
