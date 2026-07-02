package identity

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("opentelekomcloud_identity_credential_v3", func(r *config.Resource) {
		r.References["user_id"] = config.Reference{
			TerraformName: "opentelekomcloud_identity_user_v3",
		}
	})
}
