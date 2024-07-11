package blockstorage

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("opentelekomcloud_blockstorage_volume_v2", func(r *config.Resource) {
		// Nothing for now
	})
}
