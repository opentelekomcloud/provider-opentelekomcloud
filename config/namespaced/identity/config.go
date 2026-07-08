package identity

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
// TODO TODO check SelectorFieldNames
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("opentelekomcloud_identity_credential_v3", func(r *config.Resource) {
		r.References["user_id"] = config.Reference{
			TerraformName:     "opentelekomcloud_identity_user_v3",
			SelectorFieldName: "UserSelector",
			RefFieldName:      "UserSelectorRef",
		}
	})
	p.AddResourceConfigurator("opentelekomcloud_identity_group_membership_v3", func(r *config.Resource) {
		r.References["group"] = config.Reference{
			TerraformName: "opentelekomcloud_identity_group_v3",
		}
		r.References["users"] = config.Reference{
			TerraformName: "opentelekomcloud_identity_user_v3",
		}
	})
	p.AddResourceConfigurator("opentelekomcloud_identity_user_group_membership_v3", func(r *config.Resource) {
		r.References["user"] = config.Reference{
			TerraformName: "opentelekomcloud_identity_user_v3",
		}
		r.References["groups"] = config.Reference{
			TerraformName: "opentelekomcloud_identity_group_v3",
		}
	})
	p.AddResourceConfigurator("opentelekomcloud_identity_role_assignment_v3", func(r *config.Resource) {
		r.References["group_id"] = config.Reference{
			TerraformName:     "opentelekomcloud_identity_group_v3",
			SelectorFieldName: "GroupSelector",
			RefFieldName:      "GroupSelectorRef",
		}
		r.References["project_id"] = config.Reference{
			TerraformName:     "opentelekomcloud_identity_project_v3",
			SelectorFieldName: "ProjectSelector",
			RefFieldName:      "ProjectSelectorRef",
		}
		r.References["role_id"] = config.Reference{
			TerraformName:     "opentelekomcloud_identity_role_v3",
			SelectorFieldName: "RoleSelector",
			RefFieldName:      "RoleSelectorRef",
		}
	})
	p.AddResourceConfigurator("opentelekomcloud_identity_protocol_v3", func(r *config.Resource) {
		r.References["provider_id"] = config.Reference{
			TerraformName:     "opentelekomcloud_identity_provider_v3",
			SelectorFieldName: "ProviderSelector",
			RefFieldName:      "ProviderSelectorRef",
		}
		r.References["mapping_id"] = config.Reference{
			TerraformName:     "opentelekomcloud_identity_mapping_v3",
			SelectorFieldName: "MappingSelector",
			RefFieldName:      "MappingSelectorRef",
		}
	})
}
