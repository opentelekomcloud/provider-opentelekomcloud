package wafd

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/opentelekomcloud/provider-opentelekomcloud/config/common"
)

const (
	tfAlarmMaskingRuleV1      = "opentelekomcloud_waf_dedicated_alarm_masking_rule_v1"
	tfAntiCrawlerRuleV1       = "opentelekomcloud_waf_dedicated_anti_crawler_rule_v1"
	tfAntiLeakageRuleV1       = "opentelekomcloud_waf_dedicated_anti_leakage_rule_v1"
	tfBlacklistRuleV1         = "opentelekomcloud_waf_dedicated_blacklist_rule_v1"
	tfCcRuleV1                = "opentelekomcloud_waf_dedicated_cc_rule_v1"
	tfCertificateV1           = "opentelekomcloud_waf_dedicated_certificate_v1"
	tfDataMaskingRuleV1       = "opentelekomcloud_waf_dedicated_data_masking_rule_v1"
	tfDomainV1                = "opentelekomcloud_waf_dedicated_domain_v1"
	tfGeoIpRuleV1             = "opentelekomcloud_waf_dedicated_geo_ip_rule_v1"
	tfInstanceV1              = "opentelekomcloud_waf_dedicated_instance_v1"
	tfKnownAttackSourceRuleV1 = "opentelekomcloud_waf_dedicated_known_attack_source_rule_v1"
	tfPolicyV1                = "opentelekomcloud_waf_dedicated_policy_v1"
	tfPreciseProtectionRuleV1 = "opentelekomcloud_waf_dedicated_precise_protection_rule_v1"
	tfReferenceTableV1        = "opentelekomcloud_waf_dedicated_reference_table_v1"
	tfWebTamperRuleV1         = "opentelekomcloud_waf_dedicated_web_tamper_rule_v1"

	tfVpcSecgroupV3 = "opentelekomcloud_vpc_secgroup_v3"
	tfVpcSubnetV1   = "opentelekomcloud_vpc_subnet_v1"
	tfVpcV1         = "opentelekomcloud_vpc_v1"

	selectorFieldNameCertificate       = "CertificateSelector"
	refFieldNameCertificate            = "CertificateRef"
	selectorFieldNameDomain            = "DomainSelector"
	refFieldNameDomain                 = "DomainRef"
	selectorFieldNameKnownAttackSource = "KnownAttackSourceSelector"
	refFieldNameKnownAttackSource      = "KnownAttackSourceRef"
	selectorFieldNamePolicy            = "PolicySelector"
	refFieldNamePolicy                 = "PolicyRef"
	selectorFieldNameReferenceTable    = "ReferenceTableSelector"
	refFieldNameReferenceTable         = "ReferenceTableRef"
	selectorFieldNameSecurityGroup     = "SecurityGroupSelector"
	refFieldNameSecurityGroup          = "SecurityGroupRef"
	selectorFieldNameSubnet            = "SubnetSelector"
	refFieldNameSubnet                 = "SubnetRef"
	selectorFieldNameVPC               = "VPCSelector"
	refFieldNameVPC                    = "VPCRef"
)

func Configure(p *config.Provider) {
	// Classic WAF (opentelekomcloud_waf_*) is legacy and not planned for
	// implementation: its default kinds (DomainV1, PolicyV1, CertificateV1)
	// would collide with the shortened WAFD kind names.
	p.AddResourceConfigurator(tfInstanceV1, func(r *config.Resource) {
		r.UseAsync = true
		r.Kind = "InstanceV1"
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
		r.References["security_group"] = config.Reference{
			TerraformName:     tfVpcSecgroupV3,
			SelectorFieldName: selectorFieldNameSecurityGroup,
			RefFieldName:      refFieldNameSecurityGroup,
		}
	})

	p.AddResourceConfigurator(tfDomainV1, func(r *config.Resource) {
		r.Kind = "DomainV1"
		r.References["server.vpc_id"] = config.Reference{
			TerraformName:     tfVpcV1,
			SelectorFieldName: selectorFieldNameVPC,
			RefFieldName:      refFieldNameVPC,
		}
		r.References["certificate_id"] = config.Reference{
			TerraformName:     tfCertificateV1,
			SelectorFieldName: selectorFieldNameCertificate,
			RefFieldName:      refFieldNameCertificate,
		}
		r.References["policy_id"] = config.Reference{
			TerraformName:     tfPolicyV1,
			SelectorFieldName: selectorFieldNamePolicy,
			RefFieldName:      refFieldNamePolicy,
		}
	})

	p.AddResourceConfigurator(tfPolicyV1, func(r *config.Resource) {
		r.Kind = "PolicyV1"
	})

	p.AddResourceConfigurator(tfCertificateV1, func(r *config.Resource) {
		r.Kind = "CertificateV1"
		r.TerraformResource.Schema["key"].Sensitive = true
	})

	p.AddResourceConfigurator(tfCcRuleV1, func(r *config.Resource) {
		r.Kind = "CcRuleV1"
		r.References["policy_id"] = config.Reference{
			TerraformName:     tfPolicyV1,
			SelectorFieldName: selectorFieldNamePolicy,
			RefFieldName:      refFieldNamePolicy,
		}
		r.References["conditions.value_list_id"] = config.Reference{
			TerraformName:     tfReferenceTableV1,
			SelectorFieldName: selectorFieldNameReferenceTable,
			RefFieldName:      refFieldNameReferenceTable,
		}
	})

	p.AddResourceConfigurator(tfAntiCrawlerRuleV1, func(r *config.Resource) {
		r.Kind = "AntiCrawlerRuleV1"
		r.References["policy_id"] = config.Reference{
			TerraformName:     tfPolicyV1,
			SelectorFieldName: selectorFieldNamePolicy,
			RefFieldName:      refFieldNamePolicy,
		}
	})

	p.AddResourceConfigurator(tfDataMaskingRuleV1, func(r *config.Resource) {
		r.Kind = "DataMaskingRuleV1"
		r.References["policy_id"] = config.Reference{
			TerraformName:     tfPolicyV1,
			SelectorFieldName: selectorFieldNamePolicy,
			RefFieldName:      refFieldNamePolicy,
		}
	})

	p.AddResourceConfigurator(tfKnownAttackSourceRuleV1, func(r *config.Resource) {
		r.Kind = "KnownAttackSourceRuleV1"
		r.References["policy_id"] = config.Reference{
			TerraformName:     tfPolicyV1,
			SelectorFieldName: selectorFieldNamePolicy,
			RefFieldName:      refFieldNamePolicy,
		}
	})

	p.AddResourceConfigurator(tfWebTamperRuleV1, func(r *config.Resource) {
		r.Kind = "WebTamperRuleV1"
		r.References["policy_id"] = config.Reference{
			TerraformName:     tfPolicyV1,
			SelectorFieldName: selectorFieldNamePolicy,
			RefFieldName:      refFieldNamePolicy,
		}
		r.References["hostname"] = config.Reference{
			TerraformName:     tfDomainV1,
			Extractor:         common.WafDomainNameExtractor,
			SelectorFieldName: selectorFieldNameDomain,
			RefFieldName:      refFieldNameDomain,
		}
	})

	p.AddResourceConfigurator(tfAntiLeakageRuleV1, func(r *config.Resource) {
		r.Kind = "AntiLeakageRuleV1"
		r.References["policy_id"] = config.Reference{
			TerraformName:     tfPolicyV1,
			SelectorFieldName: selectorFieldNamePolicy,
			RefFieldName:      refFieldNamePolicy,
		}
	})

	p.AddResourceConfigurator(tfAlarmMaskingRuleV1, func(r *config.Resource) {
		r.Kind = "AlarmMaskingRuleV1"
		r.References["policy_id"] = config.Reference{
			TerraformName:     tfPolicyV1,
			SelectorFieldName: selectorFieldNamePolicy,
			RefFieldName:      refFieldNamePolicy,
		}
		r.References["domains"] = config.Reference{
			TerraformName:     tfDomainV1,
			Extractor:         common.WafDomainNameExtractor,
			SelectorFieldName: selectorFieldNameDomain,
			RefFieldName:      refFieldNameDomain,
		}
	})

	p.AddResourceConfigurator(tfGeoIpRuleV1, func(r *config.Resource) {
		r.Kind = "GeoIPRuleV1"
		r.References["policy_id"] = config.Reference{
			TerraformName:     tfPolicyV1,
			SelectorFieldName: selectorFieldNamePolicy,
			RefFieldName:      refFieldNamePolicy,
		}
	})

	p.AddResourceConfigurator(tfBlacklistRuleV1, func(r *config.Resource) {
		r.Kind = "BlacklistRuleV1"
		r.MetaResource.ArgumentDocs["followed_action_id"] = "ID of a known attack source rule. Changing this creates a new rule. The WAFD API requires a session tag on the protected domain. The traffic_identifier field is computed only in this provider, so configure the session tag out of band (console or traffic_mark.cookie API) before using this field."
		r.References["policy_id"] = config.Reference{
			TerraformName:     tfPolicyV1,
			SelectorFieldName: selectorFieldNamePolicy,
			RefFieldName:      refFieldNamePolicy,
		}
		r.References["followed_action_id"] = config.Reference{
			TerraformName:     tfKnownAttackSourceRuleV1,
			SelectorFieldName: selectorFieldNameKnownAttackSource,
			RefFieldName:      refFieldNameKnownAttackSource,
		}
	})

	p.AddResourceConfigurator(tfPreciseProtectionRuleV1, func(r *config.Resource) {
		r.Kind = "PreciseProtectionRuleV1"
		r.References["policy_id"] = config.Reference{
			TerraformName:     tfPolicyV1,
			SelectorFieldName: selectorFieldNamePolicy,
			RefFieldName:      refFieldNamePolicy,
		}
		r.References["conditions.value_list_id"] = config.Reference{
			TerraformName:     tfReferenceTableV1,
			SelectorFieldName: selectorFieldNameReferenceTable,
			RefFieldName:      refFieldNameReferenceTable,
		}
		r.References["action.followed_action_id"] = config.Reference{
			TerraformName:     tfKnownAttackSourceRuleV1,
			SelectorFieldName: selectorFieldNameKnownAttackSource,
			RefFieldName:      refFieldNameKnownAttackSource,
		}
	})

	p.AddResourceConfigurator(tfReferenceTableV1, func(r *config.Resource) {
		r.Kind = "ReferenceTableV1"
	})
}
