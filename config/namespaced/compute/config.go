package compute

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/opentelekomcloud/provider-opentelekomcloud/config/common"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("opentelekomcloud_compute_instance_v2", func(r *config.Resource) {
		r.UseAsync = true
		r.References["key_pair"] = config.Reference{
			TerraformName: "opentelekomcloud_compute_keypair_v2",
		}
		r.References["security_groups"] = config.Reference{
			TerraformName: "opentelekomcloud_compute_secgroup_v2",
		}
		r.References["block_device.uuid"] = config.Reference{
			TerraformName: "opentelekomcloud_blockstorage_volume_v2",
		}
		r.References["network.uuid"] = config.Reference{
			TerraformName: "opentelekomcloud_networking_network_v2",
		}
		r.References["network.name"] = config.Reference{
			TerraformName: "opentelekomcloud_networking_network_v2",
			Extractor:     common.NetworkingNetworkNameExtractor,
		}
	})
	p.AddResourceConfigurator("opentelekomcloud_compute_keypair_v2", func(r *config.Resource) {
		r.UseAsync = true
	})
	p.AddResourceConfigurator("opentelekomcloud_ecs_instance_v1", func(r *config.Resource) {
		r.UseAsync = true
		r.References["key_name"] = config.Reference{
			TerraformName: "opentelekomcloud_compute_keypair_v2",
		}
		r.References["security_groups"] = config.Reference{
			TerraformName:     "opentelekomcloud_compute_secgroup_v2",
			RefFieldName:      "ComputeSecurityGroupRefs",
			SelectorFieldName: "ComputeSecurityGroupSelector",
		}
		r.References["vpc_id"] = config.Reference{
			TerraformName:     "opentelekomcloud_vpc_v1",
			SelectorFieldName: "VPCSelector",
			RefFieldName:      "VPCRef",
		}
		r.References["nics.network_id"] = config.Reference{
			TerraformName:     "opentelekomcloud_vpc_subnet_v1",
			Extractor:         common.NetworkIDExtractor,
			SelectorFieldName: "NetworkSelector",
			RefFieldName:      "NetworkRef",
		}
		r.References["data_disks.kms_id"] = config.Reference{
			TerraformName:     "opentelekomcloud_kms_key_v1",
			SelectorFieldName: "KMSSelector",
			RefFieldName:      "KMSRef",
		}
		r.References["system_disk_kms_id"] = config.Reference{
			TerraformName:     "opentelekomcloud_kms_key_v1",
			SelectorFieldName: "SystemDiskKMSSelector",
			RefFieldName:      "SystemDiskKMSRef",
		}
	})
	p.AddResourceConfigurator("opentelekomcloud_compute_volume_attach_v2", func(r *config.Resource) {
		r.UseAsync = true
		r.References["instance_id"] = config.Reference{
			TerraformName:     "opentelekomcloud_compute_instance_v2",
			SelectorFieldName: "InstanceSelector",
			RefFieldName:      "InstanceRef",
		}
		r.References["volume_id"] = config.Reference{
			TerraformName:     "opentelekomcloud_blockstorage_volume_v2",
			SelectorFieldName: "VolumeSelector",
			RefFieldName:      "VolumeRef",
		}
	})
	p.AddResourceConfigurator("opentelekomcloud_compute_floatingip_associate_v2", func(r *config.Resource) {
		r.UseAsync = true
		r.References["instance_id"] = config.Reference{
			TerraformName:     "opentelekomcloud_compute_instance_v2",
			SelectorFieldName: "InstanceSelector",
			RefFieldName:      "InstanceRef",
		}
		r.References["floating_ip"] = config.Reference{
			TerraformName: "opentelekomcloud_networking_floatingip_v2",
			Extractor:     common.FipAddressExtractor,
		}
	})
}
