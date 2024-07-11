package dns

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("opentelekomcloud_dns_zone_v2", func(r *config.Resource) {
		r.References["zone_id"] = config.Reference{
			Type: "ZoneV2",
		}
	})
}
