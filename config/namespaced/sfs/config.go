package sfs

import "github.com/crossplane/upjet/v2/pkg/config"

const (
	tfVpcV1           = "opentelekomcloud_vpc_v1"
	tfVpcSubnetV1     = "opentelekomcloud_vpc_subnet_v1"
	tfVpcSecgroupV3   = "opentelekomcloud_vpc_secgroup_v3"
	tfSfsTurboShareV1 = "opentelekomcloud_sfs_turbo_share_v1"
	tfKmsKeyV1        = "opentelekomcloud_kms_key_v1"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
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
