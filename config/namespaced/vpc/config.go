package vpc

import "github.com/crossplane/upjet/v2/pkg/config"

const (
	tfVpcV1 = "opentelekomcloud_vpc_v1"

	VPCSelector = "VPCSelector"
	VPCRef      = "VPCRef"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("opentelekomcloud_vpc_subnet_v1", func(r *config.Resource) {
		r.UseAsync = true
		r.References["vpc_id"] = config.Reference{
			TerraformName:     tfVpcV1,
			SelectorFieldName: VPCSelector,
			RefFieldName:      VPCRef,
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
			TerraformName:     tfVpcV1,
			SelectorFieldName: "ResourceSelector",
			RefFieldName:      "ResourceRef",
		}
		r.References["log_group_id"] = config.Reference{
			TerraformName:     "opentelekomcloud_logtank_group_v2",
			SelectorFieldName: "LogGroupSelector",
			RefFieldName:      "LogGroupRef",
		}
		r.References["log_topic_id"] = config.Reference{
			TerraformName:     "opentelekomcloud_logtank_topic_v2",
			SelectorFieldName: "LogTopicSelector",
			RefFieldName:      "LogTopicRef",
		}
	})
	p.AddResourceConfigurator("opentelekomcloud_vpc_peering_connection_v2", func(r *config.Resource) {
		r.UseAsync = true
		r.References["vpc_id"] = config.Reference{
			TerraformName:     tfVpcV1,
			SelectorFieldName: VPCSelector,
			RefFieldName:      VPCRef,
		}
		r.References["peer_vpc_id"] = config.Reference{
			TerraformName:     tfVpcV1,
			SelectorFieldName: "PeerVPCSelector",
			RefFieldName:      "PeerVPCRef",
		}
	})
	p.AddResourceConfigurator("opentelekomcloud_vpc_peering_connection_accepter_v2", func(r *config.Resource) {
		r.UseAsync = true
		r.References["vpc_peering_connection_id"] = config.Reference{
			TerraformName:     "opentelekomcloud_vpc_peering_connection_v2",
			SelectorFieldName: "VPCPeeringConnectionSelector",
			RefFieldName:      "VPCPeeringConnectionRef",
		}
	})
	p.AddResourceConfigurator("opentelekomcloud_vpc_route_v2", func(r *config.Resource) {
		r.UseAsync = true
		r.References["vpc_id"] = config.Reference{
			TerraformName:     tfVpcV1,
			SelectorFieldName: VPCSelector,
			RefFieldName:      VPCRef,
		}
	})
	p.AddResourceConfigurator("opentelekomcloud_vpc_route_table_v1", func(r *config.Resource) {
		r.UseAsync = true
		r.References["vpc_id"] = config.Reference{
			TerraformName:     tfVpcV1,
			SelectorFieldName: VPCSelector,
			RefFieldName:      VPCRef,
		}
	})
	p.AddResourceConfigurator("opentelekomcloud_vpc_secondary_cidr_v3", func(r *config.Resource) {
		r.UseAsync = true
		r.References["vpc_id"] = config.Reference{
			TerraformName:     tfVpcV1,
			SelectorFieldName: VPCSelector,
			RefFieldName:      VPCRef,
		}
	})
	p.AddResourceConfigurator("opentelekomcloud_vpc_secgroup_rule_v3", func(r *config.Resource) {
		r.UseAsync = true
		r.MetaResource.ArgumentDocs["security_group_id"] = `Configuration block defining a security_group for the rule. Only opentelekomcloud_vpc_secgroup_v3 (secgroupv3s.vpc.opentelekomcloud.m.crossplane.io) is supported for cross resource reference configuration.`
		r.References["security_group_id"] = config.Reference{
			TerraformName:     "opentelekomcloud_vpc_secgroup_v3",
			SelectorFieldName: "SecurityGroupSelector",
			RefFieldName:      "SecurityGroupRef",
		}
	})
}
