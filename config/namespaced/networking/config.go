package networking

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/opentelekomcloud/provider-opentelekomcloud/config/common"
)

const (
	tfNetworkingFloatingIPAssociateV2   = "opentelekomcloud_networking_floatingip_associate_v2"
	tfNetworkingFloatingIPV2            = "opentelekomcloud_networking_floatingip_v2"
	tfNetworkingNetworkV2               = "opentelekomcloud_networking_network_v2"
	tfNetworkingPortSecgroupAssociateV2 = "opentelekomcloud_networking_port_secgroup_associate_v2"
	tfNetworkingPortV2                  = "opentelekomcloud_networking_port_v2"
	tfNetworkingRouterInterfaceV2       = "opentelekomcloud_networking_router_interface_v2"
	tfNetworkingRouterRouteV2           = "opentelekomcloud_networking_router_route_v2"
	tfNetworkingRouterV2                = "opentelekomcloud_networking_router_v2"
	tfNetworkingSecgroupRuleV2          = "opentelekomcloud_networking_secgroup_rule_v2"
	tfNetworkingSecgroupV2              = "opentelekomcloud_networking_secgroup_v2"
	tfNetworkingSubnetV2                = "opentelekomcloud_networking_subnet_v2"
	tfNetworkingVipAssociateV2          = "opentelekomcloud_networking_vip_associate_v2"
	tfNetworkingVipV2                   = "opentelekomcloud_networking_vip_v2"

	selectorFieldNameNetwork             = "NetworkSelector"
	refFieldNameNetwork                  = "NetworkRef"
	selectorFieldNamePort                = "PortSelector"
	refFieldNamePort                     = "PortRef"
	selectorFieldNameRemoteSecurityGroup = "RemoteSecurityGroupSelector"
	refFieldNameRemoteSecurityGroup      = "RemoteSecurityGroupRef"
	selectorFieldNameRouter              = "RouterSelector"
	refFieldNameRouter                   = "RouterRef"
	selectorFieldNameSecurityGroup       = "SecurityGroupSelector"
	refFieldNameSecurityGroup            = "SecurityGroupRef"
	selectorFieldNameSubnet              = "SubnetSelector"
	refFieldNameSubnet                   = "SubnetRef"
	selectorFieldNameVip                 = "VipSelector"
	refFieldNameVip                      = "VipRef"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator(tfNetworkingSubnetV2, func(r *config.Resource) {
		r.References["network_id"] = config.Reference{
			TerraformName:     tfNetworkingNetworkV2,
			SelectorFieldName: selectorFieldNameNetwork,
			RefFieldName:      refFieldNameNetwork,
		}
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"allocation_pools"},
		}
	})
	p.AddResourceConfigurator(tfNetworkingRouterInterfaceV2, func(r *config.Resource) {
		r.References["router_id"] = config.Reference{
			TerraformName:     tfNetworkingRouterV2,
			SelectorFieldName: selectorFieldNameRouter,
			RefFieldName:      refFieldNameRouter,
		}
		r.References["subnet_id"] = config.Reference{
			TerraformName:     tfNetworkingSubnetV2,
			SelectorFieldName: selectorFieldNameSubnet,
			RefFieldName:      refFieldNameSubnet,
		}
	})
	p.AddResourceConfigurator(tfNetworkingRouterV2, func(r *config.Resource) {
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"external_gateway"},
		}
	})
	p.AddResourceConfigurator(tfNetworkingFloatingIPAssociateV2, func(r *config.Resource) {
		r.UseAsync = true
		r.References["port_id"] = config.Reference{
			TerraformName:     tfNetworkingPortV2,
			SelectorFieldName: selectorFieldNamePort,
			RefFieldName:      refFieldNamePort,
		}
		r.References["floating_ip"] = config.Reference{
			TerraformName: tfNetworkingFloatingIPV2,
			Extractor:     common.FipAddressExtractor,
		}
	})
	p.AddResourceConfigurator(tfNetworkingPortV2, func(r *config.Resource) {
		r.UseAsync = true
		r.References["network_id"] = config.Reference{
			TerraformName:     tfNetworkingNetworkV2,
			SelectorFieldName: selectorFieldNameNetwork,
			RefFieldName:      refFieldNameNetwork,
		}
		r.References["fixed_ip.subnet_id"] = config.Reference{
			TerraformName:     tfNetworkingSubnetV2,
			SelectorFieldName: selectorFieldNameSubnet,
			RefFieldName:      refFieldNameSubnet,
		}
		r.References["security_group_ids"] = config.Reference{
			TerraformName:     tfNetworkingSecgroupV2,
			SelectorFieldName: selectorFieldNameSecurityGroup,
			RefFieldName:      refFieldNameSecurityGroup,
		}
	})
	p.AddResourceConfigurator(tfNetworkingPortSecgroupAssociateV2, func(r *config.Resource) {
		r.UseAsync = true
		r.References["port_id"] = config.Reference{
			TerraformName:     tfNetworkingPortV2,
			SelectorFieldName: selectorFieldNamePort,
			RefFieldName:      refFieldNamePort,
		}
		r.References["security_group_ids"] = config.Reference{
			TerraformName:     tfNetworkingSecgroupV2,
			SelectorFieldName: selectorFieldNameSecurityGroup,
			RefFieldName:      refFieldNameSecurityGroup,
		}
	})
	p.AddResourceConfigurator(tfNetworkingRouterInterfaceV2, func(r *config.Resource) {
		r.UseAsync = true
		r.References["router_id"] = config.Reference{
			TerraformName:     tfNetworkingRouterV2,
			SelectorFieldName: selectorFieldNameRouter,
			RefFieldName:      refFieldNameRouter,
		}
		r.References["subnet_id"] = config.Reference{
			TerraformName:     tfNetworkingSubnetV2,
			SelectorFieldName: selectorFieldNameSubnet,
			RefFieldName:      refFieldNameSubnet,
		}
	})
	p.AddResourceConfigurator(tfNetworkingRouterRouteV2, func(r *config.Resource) {
		r.UseAsync = true
		r.References["router_id"] = config.Reference{
			TerraformName:     tfNetworkingRouterV2,
			SelectorFieldName: selectorFieldNameRouter,
			RefFieldName:      refFieldNameRouter,
		}
	})
	p.AddResourceConfigurator(tfNetworkingSecgroupRuleV2, func(r *config.Resource) {
		r.UseAsync = true
		r.MetaResource.ArgumentDocs["security_group_id"] = `Configuration block defining a security_group for the rule. Only opentelekomcloud_networking_secgroup_v2 (secgroupv2s.networking.opentelekomcloud.m.crossplane.io) is supported for cross resource reference configuration.`
		r.References["security_group_id"] = config.Reference{
			TerraformName:     tfNetworkingSecgroupV2,
			SelectorFieldName: selectorFieldNameSecurityGroup,
			RefFieldName:      refFieldNameSecurityGroup,
		}
		r.References["remote_group_id"] = config.Reference{
			TerraformName:     tfNetworkingSecgroupV2,
			SelectorFieldName: selectorFieldNameRemoteSecurityGroup,
			RefFieldName:      refFieldNameRemoteSecurityGroup,
		}
	})
	p.AddResourceConfigurator(tfNetworkingVipV2, func(r *config.Resource) {
		r.UseAsync = true
		r.References["network_id"] = config.Reference{
			TerraformName:     tfNetworkingNetworkV2,
			SelectorFieldName: selectorFieldNameNetwork,
			RefFieldName:      refFieldNameNetwork,
		}
		r.References["subnet_id"] = config.Reference{
			TerraformName:     tfNetworkingSubnetV2,
			SelectorFieldName: selectorFieldNameSubnet,
			RefFieldName:      refFieldNameSubnet,
		}
	})
	p.AddResourceConfigurator(tfNetworkingVipAssociateV2, func(r *config.Resource) {
		r.UseAsync = true
		r.References["vip_id"] = config.Reference{
			TerraformName:     tfNetworkingVipV2,
			SelectorFieldName: selectorFieldNameVip,
			RefFieldName:      refFieldNameVip,
		}
		r.References["port_ids"] = config.Reference{
			TerraformName:     tfNetworkingPortV2,
			SelectorFieldName: selectorFieldNamePort,
			RefFieldName:      refFieldNamePort,
		}
	})
}
