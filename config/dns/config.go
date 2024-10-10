package dns

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("opentelekomcloud_dns_zone_v2", func(r *config.Resource) {
		r.UseAsync = true
		r.References["router.router_id"] = config.Reference{
			TerraformName: "opentelekomcloud_vpc_v1",
		}
	})
	p.AddResourceConfigurator("opentelekomcloud_dns_recordset_v2", func(r *config.Resource) {
		r.UseAsync = true
		r.References["zone_id"] = config.Reference{
			TerraformName: "opentelekomcloud_dns_zone_v2",
		}
	})
	p.AddResourceConfigurator("opentelekomcloud_dns_ptrrecord_v2", func(r *config.Resource) {
		r.UseAsync = true
		r.References["floatingip_id"] = config.Reference{
			TerraformName: "opentelekomcloud_vpc_eip_v1",
		}
	})
}
