package vpc

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("opentelekomcloud_vpc_subnet_v1", func(r *config.Resource) {
		r.References["vpc_id"] = config.Reference{
			TerraformName: "opentelekomcloud_vpc_v1",
		}
	})
}
