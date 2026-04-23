package swr

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("opentelekomcloud_kms_grant_v1", func(r *config.Resource) {
		r.UseAsync = true
		r.References["grantee_principal"] = config.Reference{
			TerraformName: "opentelekomcloud_identity_user_v3",
		}
		r.References["key_id"] = config.Reference{
			TerraformName: "opentelekomcloud_kms_key_v1",
		}
	})
	// TODO add support for import CMK
	// p.AddResourceConfigurator("opentelekomcloud_kms_key_material_v1", func(r *config.Resource) {
	// 	r.UseAsync = true
	// 	r.References["key_id"] = config.Reference{
	// 		TerraformName: "opentelekomcloud_kms_key_v1",
	// 	}
	// 	r.References["import_token"] = config.Reference{
	// 		TerraformName: "opentelekomcloud_kms_key_material_parameters_v1",
	// 	}
	// 	r.References["encrypted_key_material"] = config.Reference{
	// 		TerraformName: "external",
	// 	}
	// })
}
