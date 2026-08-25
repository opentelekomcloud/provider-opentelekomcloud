package sfs

import "github.com/crossplane/upjet/v2/pkg/config"

const (
	tfSfsFileSystemV2       = "opentelekomcloud_sfs_file_system_v2"
	tfVpcV1                 = "opentelekomcloud_vpc_v1"
	tfVpcSubnetV1           = "opentelekomcloud_vpc_subnet_v1"
	tfVpcSecgroupV3         = "opentelekomcloud_vpc_secgroup_v3"
	tfSfsShareAccessRulesV2 = "opentelekomcloud_sfs_share_access_rules_v2"
	tfSfsTurboShareV1       = "opentelekomcloud_sfs_turbo_share_v1"
	tfKmsKeyV1              = "opentelekomcloud_kms_key_v1"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator(tfSfsFileSystemV2, func(r *config.Resource) {
		config.MoveToStatus(r.TerraformResource, "access_level")
		config.MoveToStatus(r.TerraformResource, "access_type")
		config.MoveToStatus(r.TerraformResource, "access_to")
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"access_level", "access_type", "access_to"},
		}
	})
	p.AddResourceConfigurator(tfSfsShareAccessRulesV2, func(r *config.Resource) {
		r.References["share_id"] = config.Reference{
			TerraformName:     tfSfsFileSystemV2,
			SelectorFieldName: "ShareSelector",
			RefFieldName:      "ShareSelectorRef",
		}
		r.References["access_rule.access_to"] = config.Reference{
			TerraformName: tfVpcV1,
		}
	})
	p.AddResourceConfigurator(tfSfsTurboShareV1, func(r *config.Resource) {
		r.References["vpc_id"] = config.Reference{
			TerraformName:     tfVpcV1,
			SelectorFieldName: "VPCSelector",
			RefFieldName:      "VPCSelectorRef",
		}
		r.References["subnet_id"] = config.Reference{
			TerraformName:     tfVpcSubnetV1,
			SelectorFieldName: "SubnetSelector",
			RefFieldName:      "SubnetSelectorRef",
		}
		r.References["security_group_id"] = config.Reference{
			TerraformName:     tfVpcSecgroupV3,
			SelectorFieldName: "SecgroupSelector",
			RefFieldName:      "SecgroupSelectorRef",
		}
		r.References["crypt_key_id"] = config.Reference{
			TerraformName:     tfKmsKeyV1,
			SelectorFieldName: "KMSSelector",
			RefFieldName:      "KMSSelectorRef",
		}
	})
}
