package dms

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/opentelekomcloud/provider-opentelekomcloud/config/common"
)

const (
	instanceV1          = "opentelekomcloud_dms_instance_v1"
	dedicatedInstanceV2 = "opentelekomcloud_dms_dedicated_instance_v2"
	topicV2             = "opentelekomcloud_dms_topic_v2"
	userV2              = "opentelekomcloud_dms_user_v2"
	smartConnectTaskV2  = "opentelekomcloud_dms_smart_connect_task_v2"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator(instanceV1, func(r *config.Resource) {
		r.UseAsync = true
		networkReferences(r, "subnet_id")
	})

	p.AddResourceConfigurator("opentelekomcloud_dms_instance_v2", func(r *config.Resource) {
		r.UseAsync = true
		networkReferences(r, "subnet_id")
		publicIPReference(r)
	})

	p.AddResourceConfigurator(dedicatedInstanceV2, func(r *config.Resource) {
		r.UseAsync = true
		networkReferences(r, "network_id")
		publicIPReference(r)
		// TF requires >= 3 cross_vpc_accesses, but single-node instances report one;
		// late-initializing it breaks every refresh.
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"cross_vpc_accesses"},
		}
	})

	p.AddResourceConfigurator("opentelekomcloud_dms_topic_v1", func(r *config.Resource) {
		r.References["instance_id"] = config.Reference{
			TerraformName:     instanceV1,
			SelectorFieldName: "InstanceSelector",
			RefFieldName:      "InstanceRef",
		}
	})

	for _, name := range []string{
		topicV2,
		userV2,
		"opentelekomcloud_dms_consumer_group_v2",
		"opentelekomcloud_dms_smart_connect_v2",
	} {
		p.AddResourceConfigurator(name, func(r *config.Resource) {
			dedicatedInstanceReference(r)
		})
	}

	p.AddResourceConfigurator("opentelekomcloud_dms_user_permission_v1", func(r *config.Resource) {
		dedicatedInstanceReference(r)
		r.References["topic_name"] = config.Reference{
			TerraformName:     topicV2,
			SelectorFieldName: "TopicSelector",
			RefFieldName:      "TopicRef",
		}
		r.References["policies.username"] = config.Reference{
			TerraformName:     userV2,
			SelectorFieldName: "UserSelector",
			RefFieldName:      "UserRef",
		}
	})

	p.AddResourceConfigurator("opentelekomcloud_dms_reassign_partitions_v2", func(r *config.Resource) {
		dedicatedInstanceReference(r)
		r.References["reassignments.topic"] = config.Reference{
			TerraformName:     topicV2,
			SelectorFieldName: "TopicSelector",
			RefFieldName:      "TopicRef",
		}
	})

	p.AddResourceConfigurator(smartConnectTaskV2, func(r *config.Resource) {
		dedicatedInstanceReference(r)
		r.References["source_task.peer_instance_id"] = config.Reference{
			TerraformName:     dedicatedInstanceV2,
			SelectorFieldName: "PeerInstanceSelector",
			RefFieldName:      "PeerInstanceRef",
		}
		r.References["topics"] = config.Reference{
			TerraformName:     topicV2,
			SelectorFieldName: "TopicSelector",
			RefFieldName:      "TopicRefs",
		}
	})

	p.AddResourceConfigurator("opentelekomcloud_dms_smart_connect_task_action_v2", func(r *config.Resource) {
		dedicatedInstanceReference(r)
		r.References["task_id"] = config.Reference{
			TerraformName:     smartConnectTaskV2,
			SelectorFieldName: "TaskSelector",
			RefFieldName:      "TaskRef",
		}
	})
}

// networkReferences adds VPC, subnet and security group references.
// DMS expects the neutron network ID, so the subnet field is resolved via NetworkIDExtractor.
func networkReferences(r *config.Resource, subnetField string) {
	r.References["vpc_id"] = config.Reference{
		TerraformName:     "opentelekomcloud_vpc_v1",
		SelectorFieldName: "VPCSelector",
		RefFieldName:      "VPCRef",
	}
	r.References[subnetField] = config.Reference{
		TerraformName:     "opentelekomcloud_vpc_subnet_v1",
		Extractor:         common.NetworkIDExtractor,
		SelectorFieldName: "SubnetSelector",
		RefFieldName:      "SubnetRef",
	}
	r.References["security_group_id"] = config.Reference{
		TerraformName:     "opentelekomcloud_compute_secgroup_v2",
		SelectorFieldName: "ComputeSecurityGroupSelector",
		RefFieldName:      "ComputeSecurityGroupRefs",
	}
}

func publicIPReference(r *config.Resource) {
	r.References["publicip_id"] = config.Reference{
		TerraformName:     "opentelekomcloud_networking_floatingip_v2",
		SelectorFieldName: "PublicIPSelector",
		RefFieldName:      "PublicIPRefs",
	}
}

func dedicatedInstanceReference(r *config.Resource) {
	r.References["instance_id"] = config.Reference{
		TerraformName:     dedicatedInstanceV2,
		SelectorFieldName: "InstanceSelector",
		RefFieldName:      "InstanceRef",
	}
}
