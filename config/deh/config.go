package deh

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("opentelekomcloud_deh_host_v1", func(r *config.Resource) {
		r.UseAsync = true
	})
}
