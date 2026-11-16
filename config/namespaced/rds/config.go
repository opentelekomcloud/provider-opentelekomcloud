package rds

import (
	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/opentelekomcloud/provider-opentelekomcloud/config/common"
)

const (
	tfBackupV3                      = "opentelekomcloud_rds_backup_v3"
	tfComputeSecgroupV2             = "opentelekomcloud_compute_secgroup_v2"
	tfInstanceMinorVersionUpgradeV3 = "opentelekomcloud_rds_instance_minor_version_upgrade_v3"
	tfInstanceV3                    = "opentelekomcloud_rds_instance_v3"
	tfKmsKeyV1                      = "opentelekomcloud_kms_key_v1"
	tfMaintenanceV3                 = "opentelekomcloud_rds_maintenance_v3"
	tfParametergroupV3              = "opentelekomcloud_rds_parametergroup_v3"
	tfPostgresExtensionV3           = "opentelekomcloud_rds_postgres_extension_v3"
	tfPublicIPAssociateV3           = "opentelekomcloud_rds_public_ip_associate_v3"
	tfReadReplicaV3                 = "opentelekomcloud_rds_read_replica_v3"
	tfVpcEipV1                      = "opentelekomcloud_vpc_eip_v1"
	tfVpcSubnetV1                   = "opentelekomcloud_vpc_subnet_v1"
	tfVpcV1                         = "opentelekomcloud_vpc_v1"

	selectorFieldNameComputeSecurityGroup = "ComputeSecurityGroupSelector"
	refFieldNameComputeSecurityGroup      = "ComputeSecurityGroupRefs"
	selectorFieldNameDiskEncryptionKey    = "DiskEncryptionKeySelector"
	refFieldNameDiskEncryptionKey         = "DiskEncryptionKeyRef"
	selectorFieldNameInstance             = "InstanceSelector"
	refFieldNameInstance                  = "InstanceRef"
	selectorFieldNameParameterGroup       = "ParameterGroupSelector"
	refFieldNameParameterGroup            = "ParameterGroupRef"
	selectorFieldNamePublicIP             = "PublicIPSelector"
	refFieldNamePublicIP                  = "PublicIPSelectorRef"
	selectorFieldNamePublicIPByID         = "PublicIPSelectorByID"
	refFieldNamePublicIPByID              = "PublicIPSelectorByIDRef"
	selectorFieldNamePublicIps            = "PublicIpsSelector"
	refFieldNamePublicIps                 = "PublicIpsRefs"
	selectorFieldNameSubnet               = "SubnetSelector"
	refFieldNameSubnet                    = "SubnetRef"
	selectorFieldNameVPC                  = "VPCSelector"
	refFieldNameVPC                       = "VPCRef"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator(tfInstanceV3, func(r *config.Resource) {
		r.UseAsync = true
		r.References["security_group_id"] = config.Reference{
			TerraformName:     tfComputeSecgroupV2,
			RefFieldName:      refFieldNameComputeSecurityGroup,
			SelectorFieldName: selectorFieldNameComputeSecurityGroup,
		}
		r.References["vpc_id"] = config.Reference{
			TerraformName:     tfVpcV1,
			SelectorFieldName: selectorFieldNameVPC,
			RefFieldName:      refFieldNameVPC,
		}
		r.References["subnet_id"] = config.Reference{
			TerraformName:     tfVpcSubnetV1,
			Extractor:         common.NetworkIDExtractor,
			SelectorFieldName: selectorFieldNameSubnet,
			RefFieldName:      refFieldNameSubnet,
		}
		r.References["public_ips"] = config.Reference{
			TerraformName:     tfVpcEipV1,
			Extractor:         common.EipAddressExtractor,
			RefFieldName:      refFieldNamePublicIps,
			SelectorFieldName: selectorFieldNamePublicIps,
		}
		r.References["param_group_id"] = config.Reference{
			TerraformName:     tfParametergroupV3,
			SelectorFieldName: selectorFieldNameParameterGroup,
			RefFieldName:      refFieldNameParameterGroup,
		}
		r.References["volume.disk_encryption_id"] = config.Reference{
			TerraformName:     tfKmsKeyV1,
			SelectorFieldName: selectorFieldNameDiskEncryptionKey,
			RefFieldName:      refFieldNameDiskEncryptionKey,
		}
		r.TerraformCustomDiff = func(diff *terraform.InstanceDiff, _ *terraform.InstanceState, _ *terraform.ResourceConfig) (*terraform.InstanceDiff, error) {
			if ipsDiff, ok := diff.Attributes["public_ips.#"]; ok && ipsDiff.New == "" {
				delete(diff.Attributes, "public_ips.#")
			}
			return diff, nil
		}
	})
	p.AddResourceConfigurator(tfBackupV3, func(r *config.Resource) {
		r.UseAsync = true
		r.References["instance_id"] = config.Reference{
			TerraformName:     tfInstanceV3,
			SelectorFieldName: selectorFieldNameInstance,
			RefFieldName:      refFieldNameInstance,
		}
	})
	p.AddResourceConfigurator(tfReadReplicaV3, func(r *config.Resource) {
		r.UseAsync = true
		r.References["replica_of_id"] = config.Reference{
			TerraformName:     tfInstanceV3,
			SelectorFieldName: selectorFieldNameInstance,
			RefFieldName:      refFieldNameInstance,
		}
		r.References["public_ips"] = config.Reference{
			TerraformName:     tfVpcEipV1,
			Extractor:         common.EipAddressExtractor,
			RefFieldName:      refFieldNamePublicIps,
			SelectorFieldName: selectorFieldNamePublicIps,
		}
		r.References["volume.disk_encryption_id"] = config.Reference{
			TerraformName:     tfKmsKeyV1,
			SelectorFieldName: selectorFieldNameDiskEncryptionKey,
			RefFieldName:      refFieldNameDiskEncryptionKey,
		}
	})
	p.AddResourceConfigurator(tfPostgresExtensionV3, func(r *config.Resource) {
		r.UseAsync = true
		r.References["instance_id"] = config.Reference{
			TerraformName:     tfInstanceV3,
			SelectorFieldName: selectorFieldNameInstance,
			RefFieldName:      refFieldNameInstance,
		}
	})
	p.AddResourceConfigurator(tfMaintenanceV3, func(r *config.Resource) {
		r.References["instance_id"] = config.Reference{
			TerraformName:     tfInstanceV3,
			SelectorFieldName: selectorFieldNameInstance,
			RefFieldName:      refFieldNameInstance,
		}
	})
	p.AddResourceConfigurator(tfInstanceMinorVersionUpgradeV3, func(r *config.Resource) {
		r.References["instance_id"] = config.Reference{
			TerraformName:     tfInstanceV3,
			SelectorFieldName: selectorFieldNameInstance,
			RefFieldName:      refFieldNameInstance,
		}
	})
	p.AddResourceConfigurator(tfPublicIPAssociateV3, func(r *config.Resource) {
		r.UseAsync = true
		r.References["instance_id"] = config.Reference{
			TerraformName:     tfInstanceV3,
			SelectorFieldName: selectorFieldNameInstance,
			RefFieldName:      refFieldNameInstance,
		}
		r.References["public_ip"] = config.Reference{
			TerraformName:     tfVpcEipV1,
			Extractor:         common.EipAddressExtractor,
			SelectorFieldName: selectorFieldNamePublicIP,
			RefFieldName:      refFieldNamePublicIP,
		}
		r.References["public_ip_id"] = config.Reference{
			TerraformName:     tfVpcEipV1,
			SelectorFieldName: selectorFieldNamePublicIPByID,
			RefFieldName:      refFieldNamePublicIPByID,
		}
	})
}
