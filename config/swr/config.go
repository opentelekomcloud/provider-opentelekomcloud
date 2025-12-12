package swr

import (
	"github.com/crossplane/upjet/pkg/config"
	"github.com/opentelekomcloud/provider-opentelekomcloud/config/common"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("opentelekomcloud_swr_repository_v2", func(r *config.Resource) {
		r.UseAsync = true
		r.References["organization"] = config.Reference{
			TerraformName: "opentelekomcloud_swr_organization_v2",
		}
	})
	p.AddResourceConfigurator("opentelekomcloud_swr_domain_v2", func(r *config.Resource) {
		r.UseAsync = true
		r.References["organization"] = config.Reference{
			TerraformName: "opentelekomcloud_swr_organization_v2",
		}
		r.References["repository"] = config.Reference{
			TerraformName: "opentelekomcloud_swr_repository_v2",
		}
	})
	p.AddResourceConfigurator("opentelekomcloud_swr_organization_permissions_v2", func(r *config.Resource) {
		r.UseAsync = true
		r.References["organization"] = config.Reference{
			TerraformName: "opentelekomcloud_swr_organization_v2",
		}
		r.References["user_id"] = config.Reference{
			TerraformName: "opentelekomcloud_identity_user_v3",
		}
		r.References["username"] = config.Reference{
			TerraformName: "opentelekomcloud_identity_user_v3",
			Extractor:     common.UsernameExtractor,
		}
	})
}
