package lts

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/opentelekomcloud/provider-opentelekomcloud/config/common"
)

const (
	tfCceAccessV3            = "opentelekomcloud_lts_cce_access_v3"
	tfComputeInstanceV2      = "opentelekomcloud_compute_instance_v2"
	tfCrossAccountAccessV2   = "opentelekomcloud_lts_cross_account_access_v2"
	tfGroupV2                = "opentelekomcloud_lts_group_v2"
	tfHostAccessV3           = "opentelekomcloud_lts_host_access_v3"
	tfHostGroupV3            = "opentelekomcloud_lts_host_group_v3"
	tfIdentityAgencyV3       = "opentelekomcloud_identity_agency_v3"
	tfKeywordsAlarmRuleV2    = "opentelekomcloud_lts_keywords_alarm_rule_v2"
	tfKmsKeyV1               = "opentelekomcloud_kms_key_v1"
	tfNotificationTemplateV2 = "opentelekomcloud_lts_notification_template_v2"
	tfObsBucket              = "opentelekomcloud_obs_bucket"
	tfQuickSearchCriteriaV1  = "opentelekomcloud_lts_quick_search_criteria_v1"
	tfSmnTopicV2             = "opentelekomcloud_smn_topic_v2"
	tfStreamV2               = "opentelekomcloud_lts_stream_v2"
	tfStructuringTemplateV3  = "opentelekomcloud_lts_structuring_template_v3"
	tfTransferV2             = "opentelekomcloud_lts_transfer_v2"

	selectorFieldNameAgency    = "AgencySelector"
	refFieldNameAgency         = "AgencyRef"
	selectorFieldNameHost      = "HostSelector"
	refFieldNameHost           = "HostRef"
	selectorFieldNameHostGroup = "HostGroupSelector"
	refFieldNameHostGroup      = "HostGroupRef"
	selectorFieldNameKMS       = "KMSSelector"
	refFieldNameKMS            = "KMSRef"
	selectorFieldNameLogGroup  = "LogGroupSelector"
	refFieldNameLogGroup       = "LogGroupRef"
	selectorFieldNameLogStream = "LogStreamSelector"
	refFieldNameLogStream      = "LogStreamRef"
	selectorFieldNameObsBucket = "ObsBucketSelector"
	refFieldNameObsBucket      = "ObsBucketRef"
	selectorFieldNameTemplate  = "TemplateSelector"
	refFieldNameTemplate       = "TemplateRef"
	selectorFieldNameTopicUrn  = "TopicUrnSelector"
	refFieldNameTopicUrn       = "TopicUrnRef"
)

func Configure(p *config.Provider) {
	p.AddResourceConfigurator(tfStreamV2, func(r *config.Resource) {
		r.References["group_id"] = config.Reference{
			TerraformName:     tfGroupV2,
			SelectorFieldName: selectorFieldNameLogGroup,
			RefFieldName:      refFieldNameLogGroup,
		}
	})

	p.AddResourceConfigurator(tfCceAccessV3, func(r *config.Resource) {
		// NOTE: cluster_id intentionally has no cross-resource reference.
		// A reference from the LTS API package to the CCE API package would create
		// the import cycle cce -> vpc -> lts -> cce, because VPC FlowLogV1
		// references LTS group/stream and CCE cluster/node resources reference VPC.
		// Users must provide the exact CCE cluster ID instead of using selectors.
		r.MetaResource.ArgumentDocs["cluster_id"] = "Specifies the CCE cluster ID. This field does not support Crossplane selectors or references because a reference would create an import cycle between the CCE, VPC and LTS API packages. Provide the exact CCE cluster ID instead of selecting it by labels."
		r.References["log_group_id"] = config.Reference{
			TerraformName:     tfGroupV2,
			SelectorFieldName: selectorFieldNameLogGroup,
			RefFieldName:      refFieldNameLogGroup,
		}
		r.References["log_stream_id"] = config.Reference{
			TerraformName:     tfStreamV2,
			SelectorFieldName: selectorFieldNameLogStream,
			RefFieldName:      refFieldNameLogStream,
		}
		r.References["host_group_ids"] = config.Reference{
			TerraformName:     tfHostGroupV3,
			SelectorFieldName: selectorFieldNameHostGroup,
			RefFieldName:      refFieldNameHostGroup,
		}
	})

	p.AddResourceConfigurator(tfCrossAccountAccessV2, func(r *config.Resource) {
		r.References["log_group_id"] = config.Reference{
			TerraformName:     tfGroupV2,
			SelectorFieldName: selectorFieldNameLogGroup,
			RefFieldName:      refFieldNameLogGroup,
		}
		r.References["log_stream_id"] = config.Reference{
			TerraformName:     tfStreamV2,
			SelectorFieldName: selectorFieldNameLogStream,
			RefFieldName:      refFieldNameLogStream,
		}
	})

	p.AddResourceConfigurator(tfHostAccessV3, func(r *config.Resource) {
		r.References["log_group_id"] = config.Reference{
			TerraformName:     tfGroupV2,
			SelectorFieldName: selectorFieldNameLogGroup,
			RefFieldName:      refFieldNameLogGroup,
		}
		r.References["log_stream_id"] = config.Reference{
			TerraformName:     tfStreamV2,
			SelectorFieldName: selectorFieldNameLogStream,
			RefFieldName:      refFieldNameLogStream,
		}
		r.References["host_group_ids"] = config.Reference{
			TerraformName:     tfHostGroupV3,
			SelectorFieldName: selectorFieldNameHostGroup,
			RefFieldName:      refFieldNameHostGroup,
		}
	})

	p.AddResourceConfigurator(tfHostGroupV3, func(r *config.Resource) {
		r.UseAsync = true
		r.References["host_ids"] = config.Reference{
			TerraformName:     tfComputeInstanceV2,
			SelectorFieldName: selectorFieldNameHost,
			RefFieldName:      refFieldNameHost,
		}
	})

	p.AddResourceConfigurator(tfKeywordsAlarmRuleV2, func(r *config.Resource) {
		r.References["keywords_requests.log_group_id"] = config.Reference{
			TerraformName:     tfGroupV2,
			SelectorFieldName: selectorFieldNameLogGroup,
			RefFieldName:      refFieldNameLogGroup,
		}
		r.References["keywords_requests.log_stream_id"] = config.Reference{
			TerraformName:     tfStreamV2,
			SelectorFieldName: selectorFieldNameLogStream,
			RefFieldName:      refFieldNameLogStream,
		}
		r.References["notification_rule.template_name"] = config.Reference{
			TerraformName:     tfNotificationTemplateV2,
			SelectorFieldName: selectorFieldNameTemplate,
			RefFieldName:      refFieldNameTemplate,
		}
		r.References["notification_rule.topics.topic_urn"] = config.Reference{
			TerraformName:     tfSmnTopicV2,
			SelectorFieldName: selectorFieldNameTopicUrn,
			RefFieldName:      refFieldNameTopicUrn,
		}
	})

	p.AddResourceConfigurator(tfQuickSearchCriteriaV1, func(r *config.Resource) {
		r.References["log_group_id"] = config.Reference{
			TerraformName:     tfGroupV2,
			SelectorFieldName: selectorFieldNameLogGroup,
			RefFieldName:      refFieldNameLogGroup,
		}
		r.References["log_stream_id"] = config.Reference{
			TerraformName:     tfStreamV2,
			SelectorFieldName: selectorFieldNameLogStream,
			RefFieldName:      refFieldNameLogStream,
		}
	})

	p.AddResourceConfigurator(tfStructuringTemplateV3, func(r *config.Resource) {
		r.References["log_group_id"] = config.Reference{
			TerraformName:     tfGroupV2,
			SelectorFieldName: selectorFieldNameLogGroup,
			RefFieldName:      refFieldNameLogGroup,
		}
		r.References["log_stream_id"] = config.Reference{
			TerraformName:     tfStreamV2,
			SelectorFieldName: selectorFieldNameLogStream,
			RefFieldName:      refFieldNameLogStream,
		}
	})

	p.AddResourceConfigurator(tfTransferV2, func(r *config.Resource) {
		r.References["log_group_id"] = config.Reference{
			TerraformName:     tfGroupV2,
			SelectorFieldName: selectorFieldNameLogGroup,
			RefFieldName:      refFieldNameLogGroup,
		}
		r.References["log_streams.log_stream_id"] = config.Reference{
			TerraformName:     tfStreamV2,
			SelectorFieldName: selectorFieldNameLogStream,
			RefFieldName:      refFieldNameLogStream,
		}
		r.References["log_transfer_info.log_transfer_detail.obs_bucket_name"] = config.Reference{
			TerraformName:     tfObsBucket,
			Extractor:         common.ObsBucketExtractor,
			SelectorFieldName: selectorFieldNameObsBucket,
			RefFieldName:      refFieldNameObsBucket,
		}
		r.References["log_transfer_info.log_transfer_detail.obs_encrypted_id"] = config.Reference{
			TerraformName:     tfKmsKeyV1,
			SelectorFieldName: selectorFieldNameKMS,
			RefFieldName:      refFieldNameKMS,
		}
		r.References["log_transfer_info.log_agency_transfer.agency_name"] = config.Reference{
			TerraformName:     tfIdentityAgencyV3,
			Extractor:         common.AgencyNameExtractor,
			SelectorFieldName: selectorFieldNameAgency,
			RefFieldName:      refFieldNameAgency,
		}
	})
}
