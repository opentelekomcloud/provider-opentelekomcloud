package lb

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/opentelekomcloud/provider-opentelekomcloud/config/common"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("opentelekomcloud_lb_loadbalancer_v3", func(r *config.Resource) {
		r.UseAsync = true
		r.References["router_id"] = config.Reference{
			TerraformName: "opentelekomcloud_vpc_v1",
		}
		r.References["network_ids"] = config.Reference{
			TerraformName: "opentelekomcloud_vpc_subnet_v1",
			Extractor:     common.NetworkIDExtractor,
		}
		r.References["subnet_id"] = config.Reference{
			TerraformName: "opentelekomcloud_vpc_subnet_v1",
		}
		r.References["public_ip.id"] = config.Reference{
			TerraformName: "opentelekomcloud_vpc_eip_v1",
		}
	})
	p.AddResourceConfigurator("opentelekomcloud_lb_listener_v3", func(r *config.Resource) {
		r.UseAsync = true
		r.References["loadbalancer_id"] = config.Reference{
			TerraformName: "opentelekomcloud_lb_loadbalancer_v3",
		}
		r.References["default_tls_container_ref"] = config.Reference{
			TerraformName: "opentelekomcloud_lb_certificate_v3",
		}
		r.References["default_pool_id"] = config.Reference{
			TerraformName: "opentelekomcloud_lb_pool_v3",
		}
		r.References["ip_group.id"] = config.Reference{
			TerraformName: "opentelekomcloud_lb_ipgroup_v3",
		}
		r.References["security_policy_id"] = config.Reference{
			TerraformName: "opentelekomcloud_lb_security_policy_v3",
		}
	})
	p.AddResourceConfigurator("opentelekomcloud_lb_pool_v3", func(r *config.Resource) {
		r.UseAsync = true
		r.References["project_id"] = config.Reference{
			TerraformName: "opentelekomcloud_identity_project_v3",
		}
		r.References["listener_id"] = config.Reference{
			TerraformName: "opentelekomcloud_lb_listener_v3",
		}
		r.References["loadbalancer_id"] = config.Reference{
			TerraformName: "opentelekomcloud_lb_loadbalancer_v3",
		}
		r.References["vpc_id"] = config.Reference{
			TerraformName: "opentelekomcloud_vpc_v1",
		}
	})
	p.AddResourceConfigurator("opentelekomcloud_lb_member_v3", func(r *config.Resource) {
		r.UseAsync = true
		r.References["project_id"] = config.Reference{
			TerraformName: "opentelekomcloud_identity_project_v3",
		}
		r.References["pool_id"] = config.Reference{
			TerraformName: "opentelekomcloud_lb_pool_v3",
		}
		r.References["subnet_id"] = config.Reference{
			TerraformName: "opentelekomcloud_vpc_subnet_v1",
		}
	})
	p.AddResourceConfigurator("opentelekomcloud_lb_ipgroup_v3", func(r *config.Resource) {
		r.UseAsync = true
		r.References["project_id"] = config.Reference{
			TerraformName: "opentelekomcloud_identity_project_v3",
		}
		r.References["ip_list.ip"] = config.Reference{
			TerraformName: "opentelekomcloud_vpc_eip_v1",
			Extractor:     common.EipAddressExtractor,
		}
	})
	p.AddResourceConfigurator("opentelekomcloud_lb_monitor_v3", func(r *config.Resource) {
		r.UseAsync = true
		r.References["pool_id"] = config.Reference{
			TerraformName: "opentelekomcloud_lb_pool_v3",
		}
		r.References["project_id"] = config.Reference{
			TerraformName: "opentelekomcloud_identity_project_v3",
		}
	})
	p.AddResourceConfigurator("opentelekomcloud_lb_policy_v3", func(r *config.Resource) {
		r.UseAsync = true
		r.References["project_id"] = config.Reference{
			TerraformName: "opentelekomcloud_identity_project_v3",
		}
		r.References["listener_id"] = config.Reference{
			TerraformName: "opentelekomcloud_lb_listener_v3",
		}
		r.References["redirect_listener_id"] = config.Reference{
			TerraformName: "opentelekomcloud_lb_listener_v3",
		}
		r.References["redirect_pool_id"] = config.Reference{
			TerraformName: "opentelekomcloud_lb_pool_v3",
		}
		r.References["redirect_pools_config.pool_id"] = config.Reference{
			TerraformName: "opentelekomcloud_lb_pool_v3",
		}
	})
	p.AddResourceConfigurator("opentelekomcloud_lb_rule_v3", func(r *config.Resource) {
		r.UseAsync = true
		r.References["policy_id"] = config.Reference{
			TerraformName: "opentelekomcloud_lb_policy_v3",
		}
		r.References["project_id"] = config.Reference{
			TerraformName: "opentelekomcloud_identity_project_v3",
		}
	})
	p.AddResourceConfigurator("opentelekomcloud_lb_security_policy_v3", func(r *config.Resource) {
		r.UseAsync = true
		r.References["project_id"] = config.Reference{
			TerraformName: "opentelekomcloud_identity_project_v3",
		}
		r.References["listener_ids"] = config.Reference{
			TerraformName: "opentelekomcloud_lb_listener_v3",
		}
	})
}
