package vpc

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/opentelekomcloud/provider-opentelekomcloud/config/common"
)

const (
	tfLtsGroupV2                     = "opentelekomcloud_lts_group_v2"
	tfLtsStreamV2                    = "opentelekomcloud_lts_stream_v2"
	tfNetworkingPortV2               = "opentelekomcloud_networking_port_v2"
	tfVpcBandwidthAssociateV2        = "opentelekomcloud_vpc_bandwidth_associate_v2"
	tfVpcBandwidthV2                 = "opentelekomcloud_vpc_bandwidth_v2"
	tfVpcEipV1                       = "opentelekomcloud_vpc_eip_v1"
	tfVpcFlowLogV1                   = "opentelekomcloud_vpc_flow_log_v1"
	tfVpcIpAddressGroupV3            = "opentelekomcloud_vpc_ip_address_group_v3"
	tfVpcPeeringConnectionAccepterV2 = "opentelekomcloud_vpc_peering_connection_accepter_v2"
	tfVpcPeeringConnectionV2         = "opentelekomcloud_vpc_peering_connection_v2"
	tfVpcRouteTableRouteV1           = "opentelekomcloud_vpc_route_table_route_v1"
	tfVpcRouteTableSubnetAssociateV1 = "opentelekomcloud_vpc_route_table_subnet_associate_v1"
	tfVpcRouteTableV1                = "opentelekomcloud_vpc_route_table_v1"
	tfVpcRouteV2                     = "opentelekomcloud_vpc_route_v2"
	tfVpcSecondaryCidrV3             = "opentelekomcloud_vpc_secondary_cidr_v3"
	tfVpcSecgroupRuleV3              = "opentelekomcloud_vpc_secgroup_rule_v3"
	tfVpcSecgroupV3                  = "opentelekomcloud_vpc_secgroup_v3"
	tfVpcSubnetPrivateIPV1           = "opentelekomcloud_vpc_subnet_private_ip_v1"
	tfVpcSubnetV1                    = "opentelekomcloud_vpc_subnet_v1"
	tfVpcV1                          = "opentelekomcloud_vpc_v1"

	selectorFieldNameBandwidth            = "BandwidthSelector"
	refFieldNameBandwidth                 = "BandwidthRef"
	selectorFieldNameLogGroup             = "LogGroupSelector"
	refFieldNameLogGroup                  = "LogGroupRef"
	selectorFieldNameLogStream            = "LogStreamSelector"
	refFieldNameLogStream                 = "LogSteamRef"
	selectorFieldNamePeerVPC              = "PeerVPCSelector"
	refFieldNamePeerVPC                   = "PeerVPCRef"
	selectorFieldNamePort                 = "PortSelector"
	refFieldNamePort                      = "PortRef"
	selectorFieldNameRemoteAddressGroup   = "RemoteAddressGroupSelector"
	refFieldNameRemoteAddressGroup        = "RemoteAddressGroupRef"
	selectorFieldNameRemoteSecurityGroup  = "RemoteSecurityGroupSelector"
	refFieldNameRemoteSecurityGroup       = "RemoteSecurityGroupRef"
	selectorFieldNameResource             = "ResourceSelector"
	refFieldNameResource                  = "ResourceRef"
	selectorFieldNameRouteTable           = "RouteTableSelector"
	refFieldNameRouteTable                = "RouteTableRef"
	selectorFieldNameSecurityGroup        = "SecurityGroupSelector"
	refFieldNameSecurityGroup             = "SecurityGroupRef"
	selectorFieldNameSubnet               = "SubnetSelector"
	refFieldNameSubnet                    = "SubnetRef"
	selectorFieldNameVPC                  = "VPCSelector"
	refFieldNameVPC                       = "VPCRef"
	selectorFieldNameVPCPeeringConnection = "VPCPeeringConnectionSelector"
	refFieldNameVPCPeeringConnection      = "VPCPeeringConnectionRef"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator(tfVpcSubnetV1, func(r *config.Resource) {
		r.UseAsync = true
		r.References["vpc_id"] = config.Reference{
			TerraformName:     tfVpcV1,
			SelectorFieldName: selectorFieldNameVPC,
			RefFieldName:      refFieldNameVPC,
		}
	})
	p.AddResourceConfigurator(tfVpcV1, func(r *config.Resource) {
		r.UseAsync = true
		r.Kind = "VpcV1"
	})
	p.AddResourceConfigurator(tfVpcBandwidthAssociateV2, func(r *config.Resource) {
		r.UseAsync = true
		r.References["bandwidth"] = config.Reference{
			TerraformName: tfVpcBandwidthV2,
		}
		r.References["floating_ips"] = config.Reference{
			TerraformName: tfVpcEipV1,
		}
	})
	p.AddResourceConfigurator(tfVpcFlowLogV1, func(r *config.Resource) {
		r.UseAsync = true
		r.References["resource_id"] = config.Reference{
			TerraformName:     tfVpcV1,
			SelectorFieldName: selectorFieldNameResource,
			RefFieldName:      refFieldNameResource,
		}
		r.References["log_group_id"] = config.Reference{
			TerraformName:     tfLtsGroupV2,
			SelectorFieldName: selectorFieldNameLogGroup,
			RefFieldName:      refFieldNameLogGroup,
		}
		r.References["log_topic_id"] = config.Reference{
			TerraformName:     tfLtsStreamV2,
			SelectorFieldName: selectorFieldNameLogStream,
			RefFieldName:      refFieldNameLogStream,
		}
	})
	p.AddResourceConfigurator(tfVpcPeeringConnectionV2, func(r *config.Resource) {
		r.UseAsync = true
		r.References["vpc_id"] = config.Reference{
			TerraformName:     tfVpcV1,
			SelectorFieldName: selectorFieldNameVPC,
			RefFieldName:      refFieldNameVPC,
		}
		r.References["peer_vpc_id"] = config.Reference{
			TerraformName:     tfVpcV1,
			SelectorFieldName: selectorFieldNamePeerVPC,
			RefFieldName:      refFieldNamePeerVPC,
		}
	})
	p.AddResourceConfigurator(tfVpcPeeringConnectionAccepterV2, func(r *config.Resource) {
		r.UseAsync = true
		r.References["vpc_peering_connection_id"] = config.Reference{
			TerraformName:     tfVpcPeeringConnectionV2,
			SelectorFieldName: selectorFieldNameVPCPeeringConnection,
			RefFieldName:      refFieldNameVPCPeeringConnection,
		}
	})
	p.AddResourceConfigurator(tfVpcRouteV2, func(r *config.Resource) {
		r.UseAsync = true
		r.References["vpc_id"] = config.Reference{
			TerraformName:     tfVpcV1,
			SelectorFieldName: selectorFieldNameVPC,
			RefFieldName:      refFieldNameVPC,
		}
	})
	p.AddResourceConfigurator(tfVpcRouteTableV1, func(r *config.Resource) {
		r.UseAsync = true
		r.References["vpc_id"] = config.Reference{
			TerraformName:     tfVpcV1,
			SelectorFieldName: selectorFieldNameVPC,
			RefFieldName:      refFieldNameVPC,
		}
	})
	p.AddResourceConfigurator(tfVpcRouteTableRouteV1, func(r *config.Resource) {
		r.UseAsync = true
		r.MetaResource.ArgumentDocs["type"] = `Specifies the route type. Valid values are "ecs", "eni", "vip", "nat", "peering", "vpn", "dc", "egw", "er", and "subeni". The "local" route is created by the platform and cannot be managed. The type determines which kind of resource or address nexthop must contain.`
		r.MetaResource.ArgumentDocs["nexthop"] = `Specifies the next hop. Its value depends on type: "ecs" is an ECS instance ID in the VPC, "eni" a supplementary NIC of an ECS, "vip" a virtual IP address, "nat" a NAT gateway ID, "peering" a VPC peering connection ID, "vpn" a VPN gateway ID, "dc" a Direct Connect gateway ID, "egw" a VPC endpoint ID, "er" an enterprise router ID, and "subeni" a supplementary network interface ID. Set this value explicitly because no cross-resource reference is generated for it: the target resource type varies with the route type.`
		r.References["vpc_id"] = config.Reference{
			TerraformName:     tfVpcV1,
			SelectorFieldName: selectorFieldNameVPC,
			RefFieldName:      refFieldNameVPC,
		}
		r.References["route_table_id"] = config.Reference{
			TerraformName:     tfVpcRouteTableV1,
			SelectorFieldName: selectorFieldNameRouteTable,
			RefFieldName:      refFieldNameRouteTable,
		}
	})
	p.AddResourceConfigurator(tfVpcRouteTableSubnetAssociateV1, func(r *config.Resource) {
		r.UseAsync = true
		r.References["route_table_id"] = config.Reference{
			TerraformName:     tfVpcRouteTableV1,
			SelectorFieldName: selectorFieldNameRouteTable,
			RefFieldName:      refFieldNameRouteTable,
		}
		r.References["subnet_id"] = config.Reference{
			TerraformName:     tfVpcSubnetV1,
			SelectorFieldName: selectorFieldNameSubnet,
			RefFieldName:      refFieldNameSubnet,
		}
	})
	p.AddResourceConfigurator(tfVpcSubnetPrivateIPV1, func(r *config.Resource) {
		r.UseAsync = true
		r.References["subnet_id"] = config.Reference{
			TerraformName:     tfVpcSubnetV1,
			Extractor:         common.NetworkIDExtractor,
			SelectorFieldName: selectorFieldNameSubnet,
			RefFieldName:      refFieldNameSubnet,
		}
	})
	p.AddResourceConfigurator(tfVpcEipV1, func(r *config.Resource) {
		r.UseAsync = true
		r.References["bandwidth.id"] = config.Reference{
			TerraformName:     tfVpcBandwidthV2,
			SelectorFieldName: selectorFieldNameBandwidth,
			RefFieldName:      refFieldNameBandwidth,
		}
		r.References["publicip.port_id"] = config.Reference{
			TerraformName:     tfNetworkingPortV2,
			SelectorFieldName: selectorFieldNamePort,
			RefFieldName:      refFieldNamePort,
		}
	})
	p.AddResourceConfigurator(tfVpcSecondaryCidrV3, func(r *config.Resource) {
		r.UseAsync = true
		r.References["vpc_id"] = config.Reference{
			TerraformName:     tfVpcV1,
			SelectorFieldName: selectorFieldNameVPC,
			RefFieldName:      refFieldNameVPC,
		}
	})
	p.AddResourceConfigurator(tfVpcSecgroupRuleV3, func(r *config.Resource) {
		r.UseAsync = true
		r.MetaResource.ArgumentDocs["security_group_id"] = `Configuration block defining a security_group for the rule. Only opentelekomcloud_vpc_secgroup_v3 (secgroupv3s.vpc.opentelekomcloud.m.crossplane.io) is supported for cross resource reference configuration.`
		r.References["security_group_id"] = config.Reference{
			TerraformName:     tfVpcSecgroupV3,
			SelectorFieldName: selectorFieldNameSecurityGroup,
			RefFieldName:      refFieldNameSecurityGroup,
		}
		r.References["remote_group_id"] = config.Reference{
			TerraformName:     tfVpcSecgroupV3,
			SelectorFieldName: selectorFieldNameRemoteSecurityGroup,
			RefFieldName:      refFieldNameRemoteSecurityGroup,
		}
		r.References["remote_address_group_id"] = config.Reference{
			TerraformName:     tfVpcIpAddressGroupV3,
			SelectorFieldName: selectorFieldNameRemoteAddressGroup,
			RefFieldName:      refFieldNameRemoteAddressGroup,
		}
	})
}
