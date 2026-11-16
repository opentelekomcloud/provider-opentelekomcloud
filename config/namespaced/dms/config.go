package dms

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/opentelekomcloud/provider-opentelekomcloud/config/common"
)

const (
	tfComputeSecgroupV2        = "opentelekomcloud_compute_secgroup_v2"
	tfConsumerGroupV2          = "opentelekomcloud_dms_consumer_group_v2"
	tfDedicatedInstanceV2      = "opentelekomcloud_dms_dedicated_instance_v2"
	tfFloatingIPV2             = "opentelekomcloud_networking_floatingip_v2"
	tfInstanceV1               = "opentelekomcloud_dms_instance_v1"
	tfInstanceV2               = "opentelekomcloud_dms_instance_v2"
	tfReassignPartitionsV2     = "opentelekomcloud_dms_reassign_partitions_v2"
	tfSmartConnectTaskActionV2 = "opentelekomcloud_dms_smart_connect_task_action_v2"
	tfSmartConnectTaskV2       = "opentelekomcloud_dms_smart_connect_task_v2"
	tfSmartConnectV2           = "opentelekomcloud_dms_smart_connect_v2"
	tfSubnetV1                 = "opentelekomcloud_vpc_subnet_v1"
	tfTopicV1                  = "opentelekomcloud_dms_topic_v1"
	tfTopicV2                  = "opentelekomcloud_dms_topic_v2"
	tfUserPermissionV1         = "opentelekomcloud_dms_user_permission_v1"
	tfUserV2                   = "opentelekomcloud_dms_user_v2"
	tfVPCV1                    = "opentelekomcloud_vpc_v1"

	selectorFieldNameComputeSecurityGroup = "ComputeSecurityGroupSelector"
	refFieldNameComputeSecurityGroup      = "ComputeSecurityGroupRefs"
	selectorFieldNameInstance             = "InstanceSelector"
	refFieldNameInstance                  = "InstanceRef"
	selectorFieldNamePeerInstance         = "PeerInstanceSelector"
	refFieldNamePeerInstance              = "PeerInstanceRef"
	selectorFieldNamePublicIP             = "PublicIPSelector"
	refFieldNamePublicIP                  = "PublicIPRefs"
	selectorFieldNameSubnet               = "SubnetSelector"
	refFieldNameSubnet                    = "SubnetRef"
	selectorFieldNameTask                 = "TaskSelector"
	refFieldNameTask                      = "TaskRef"
	selectorFieldNameTopic                = "TopicSelector"
	refFieldNameTopic                     = "TopicRef"
	refFieldNameTopics                    = "TopicRefs"
	selectorFieldNameUser                 = "UserSelector"
	refFieldNameUser                      = "UserRef"
	selectorFieldNameVPC                  = "VPCSelector"
	refFieldNameVPC                       = "VPCRef"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator(tfInstanceV1, func(r *config.Resource) {
		r.UseAsync = true
		networkReferences(r, "subnet_id")
	})

	p.AddResourceConfigurator(tfInstanceV2, func(r *config.Resource) {
		r.UseAsync = true
		networkReferences(r, "subnet_id")
		publicIPReference(r)
	})

	p.AddResourceConfigurator(tfDedicatedInstanceV2, func(r *config.Resource) {
		r.UseAsync = true
		networkReferences(r, "network_id")
		publicIPReference(r)
		// TF requires >= 3 cross_vpc_accesses, but single-node instances report one;
		// late-initializing it breaks every refresh.
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"cross_vpc_accesses"},
		}
	})

	p.AddResourceConfigurator(tfTopicV1, func(r *config.Resource) {
		r.References["instance_id"] = config.Reference{
			TerraformName:     tfInstanceV1,
			SelectorFieldName: selectorFieldNameInstance,
			RefFieldName:      refFieldNameInstance,
		}
	})

	for _, name := range []string{
		tfTopicV2,
		tfUserV2,
		tfConsumerGroupV2,
		tfSmartConnectV2,
	} {
		p.AddResourceConfigurator(name, func(r *config.Resource) {
			dedicatedInstanceReference(r)
		})
	}

	p.AddResourceConfigurator(tfUserPermissionV1, func(r *config.Resource) {
		dedicatedInstanceReference(r)
		r.References["topic_name"] = config.Reference{
			TerraformName:     tfTopicV2,
			SelectorFieldName: selectorFieldNameTopic,
			RefFieldName:      refFieldNameTopic,
		}
		r.References["policies.username"] = config.Reference{
			TerraformName:     tfUserV2,
			SelectorFieldName: selectorFieldNameUser,
			RefFieldName:      refFieldNameUser,
		}
	})

	p.AddResourceConfigurator(tfReassignPartitionsV2, func(r *config.Resource) {
		dedicatedInstanceReference(r)
		r.References["reassignments.topic"] = config.Reference{
			TerraformName:     tfTopicV2,
			SelectorFieldName: selectorFieldNameTopic,
			RefFieldName:      refFieldNameTopic,
		}
	})

	p.AddResourceConfigurator(tfSmartConnectTaskV2, func(r *config.Resource) {
		dedicatedInstanceReference(r)
		r.References["source_task.peer_instance_id"] = config.Reference{
			TerraformName:     tfDedicatedInstanceV2,
			SelectorFieldName: selectorFieldNamePeerInstance,
			RefFieldName:      refFieldNamePeerInstance,
		}
		r.References["topics"] = config.Reference{
			TerraformName:     tfTopicV2,
			SelectorFieldName: selectorFieldNameTopic,
			RefFieldName:      refFieldNameTopics,
		}
	})

	p.AddResourceConfigurator(tfSmartConnectTaskActionV2, func(r *config.Resource) {
		dedicatedInstanceReference(r)
		r.References["task_id"] = config.Reference{
			TerraformName:     tfSmartConnectTaskV2,
			SelectorFieldName: selectorFieldNameTask,
			RefFieldName:      refFieldNameTask,
		}
	})
}

// networkReferences adds VPC, subnet and security group references.
// DMS expects the neutron network ID, so the subnet field is resolved via NetworkIDExtractor.
func networkReferences(r *config.Resource, subnetField string) {
	r.References["vpc_id"] = config.Reference{
		TerraformName:     tfVPCV1,
		SelectorFieldName: selectorFieldNameVPC,
		RefFieldName:      refFieldNameVPC,
	}
	r.References[subnetField] = config.Reference{
		TerraformName:     tfSubnetV1,
		Extractor:         common.NetworkIDExtractor,
		SelectorFieldName: selectorFieldNameSubnet,
		RefFieldName:      refFieldNameSubnet,
	}
	r.References["security_group_id"] = config.Reference{
		TerraformName:     tfComputeSecgroupV2,
		SelectorFieldName: selectorFieldNameComputeSecurityGroup,
		RefFieldName:      refFieldNameComputeSecurityGroup,
	}
}

func publicIPReference(r *config.Resource) {
	r.References["publicip_id"] = config.Reference{
		TerraformName:     tfFloatingIPV2,
		SelectorFieldName: selectorFieldNamePublicIP,
		RefFieldName:      refFieldNamePublicIP,
	}
}

func dedicatedInstanceReference(r *config.Resource) {
	r.References["instance_id"] = config.Reference{
		TerraformName:     tfDedicatedInstanceV2,
		SelectorFieldName: selectorFieldNameInstance,
		RefFieldName:      refFieldNameInstance,
	}
}
