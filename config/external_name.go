/*
Copyright 2022 Upbound Inc.
*/

package config

import "github.com/crossplane/upjet/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// APIGW
	"opentelekomcloud_apigw_acl_policy_v2":                  config.IdentifierFromProvider,
	"opentelekomcloud_apigw_acl_policy_associate_v2":        config.IdentifierFromProvider,
	"opentelekomcloud_apigw_api_v2":                         config.IdentifierFromProvider,
	"opentelekomcloud_apigw_api_publishment_v2":             config.IdentifierFromProvider,
	"opentelekomcloud_apigw_application_v2":                 config.IdentifierFromProvider,
	"opentelekomcloud_apigw_application_authorization_v2":   config.IdentifierFromProvider,
	"opentelekomcloud_apigw_appcode_v2":                     config.IdentifierFromProvider,
	"opentelekomcloud_apigw_certificate_v2":                 config.IdentifierFromProvider,
	"opentelekomcloud_apigw_custom_authorizer_v2":           config.IdentifierFromProvider,
	"opentelekomcloud_apigw_environment_v2":                 config.IdentifierFromProvider,
	"opentelekomcloud_apigw_environment_variable_v2":        config.IdentifierFromProvider,
	"opentelekomcloud_apigw_gateway_v2":                     config.IdentifierFromProvider,
	"opentelekomcloud_apigw_gateway_routes_v2":              config.IdentifierFromProvider,
	"opentelekomcloud_apigw_gateway_feature_v2":             config.IdentifierFromProvider,
	"opentelekomcloud_apigw_group_v2":                       config.IdentifierFromProvider,
	"opentelekomcloud_apigw_response_v2":                    config.IdentifierFromProvider,
	"opentelekomcloud_apigw_signature_v2":                   config.IdentifierFromProvider,
	"opentelekomcloud_apigw_signature_associate_v2":         config.IdentifierFromProvider,
	"opentelekomcloud_apigw_throttling_policy_v2":           config.IdentifierFromProvider,
	"opentelekomcloud_apigw_throttling_policy_associate_v2": config.IdentifierFromProvider,
	"opentelekomcloud_apigw_vpc_channel_v2":                 config.IdentifierFromProvider,

	// ASM
	"opentelekomcloud_asm_service_mesh_v1": config.IdentifierFromProvider,

	// BLOCK STORAGE / EVS
	"opentelekomcloud_evs_volume_v3":          config.IdentifierFromProvider,
	"opentelekomcloud_blockstorage_volume_v2": config.IdentifierFromProvider,

	// CBR
	"opentelekomcloud_cbr_policy_v3": config.IdentifierFromProvider,
	"opentelekomcloud_cbr_vault_v3":  config.IdentifierFromProvider,

	// CCE
	"opentelekomcloud_cce_cluster_v3":     config.IdentifierFromProvider,
	"opentelekomcloud_cce_node_v3":        config.IdentifierFromProvider,
	"opentelekomcloud_cce_node_pool_v3":   config.IdentifierFromProvider,
	"opentelekomcloud_cce_addon_v3":       config.IdentifierFromProvider,
	"opentelekomcloud_cce_node_attach_v3": config.IdentifierFromProvider,

	// CWF
	"opentelekomcloud_cfw_acl_rule_v1":                 config.IdentifierFromProvider,
	"opentelekomcloud_cfw_address_group_v1":            config.IdentifierFromProvider,
	"opentelekomcloud_cfw_address_group_member_v1":     config.IdentifierFromProvider,
	"opentelekomcloud_cfw_blacklist_whitelist_rule_v1": config.IdentifierFromProvider,
	"opentelekomcloud_cfw_domain_name_group_v1":        config.IdentifierFromProvider,
	"opentelekomcloud_cfw_eip_protection_v1":           config.IdentifierFromProvider,
	"opentelekomcloud_cfw_firewall_v1":                 config.IdentifierFromProvider,
	"opentelekomcloud_cfw_ips_protection_v1":           config.IdentifierFromProvider,
	"opentelekomcloud_cfw_service_group_v1":            config.IdentifierFromProvider,
	"opentelekomcloud_cfw_service_group_member_v1":     config.IdentifierFromProvider,

	// COMPUTE / ECS
	"opentelekomcloud_compute_secgroup_v2":             config.IdentifierFromProvider,
	"opentelekomcloud_compute_servergroup_v2":          config.IdentifierFromProvider,
	"opentelekomcloud_compute_floatingip_v2":           config.IdentifierFromProvider,
	"opentelekomcloud_compute_floatingip_associate_v2": config.IdentifierFromProvider,
	"opentelekomcloud_compute_instance_v2":             config.IdentifierFromProvider,
	"opentelekomcloud_compute_keypair_v2":              config.IdentifierFromProvider,
	"opentelekomcloud_compute_volume_attach_v2":        config.IdentifierFromProvider,
	"opentelekomcloud_ecs_instance_v1":                 config.IdentifierFromProvider,

	// CTS
	"opentelekomcloud_cts_event_notification_v3": config.IdentifierFromProvider,
	"opentelekomcloud_cts_tracker_v3":            config.IdentifierFromProvider,

	// CSS
	"opentelekomcloud_css_cluster_v1":                config.IdentifierFromProvider,
	"opentelekomcloud_css_cluster_restart_v1":        config.IdentifierFromProvider,
	"opentelekomcloud_css_configuration_v1":          config.IdentifierFromProvider,
	"opentelekomcloud_css_snapshot_configuration_v1": config.IdentifierFromProvider,

	// DCS
	"opentelekomcloud_dcs_instance_v2": config.IdentifierFromProvider,

	// DDM
	"opentelekomcloud_ddm_instance_v1": config.IdentifierFromProvider,
	"opentelekomcloud_ddm_schema_v1":   config.IdentifierFromProvider,

	// DDS
	"opentelekomcloud_dds_instance_v3": config.IdentifierFromProvider,
	"opentelekomcloud_dds_backup_v3":   config.IdentifierFromProvider,
	"opentelekomcloud_dds_lts_log_v3":  config.IdentifierFromProvider,

	// DEH
	"opentelekomcloud_deh_host_v1": config.IdentifierFromProvider,

	// DIS
	"opentelekomcloud_dis_stream_v2":     config.IdentifierFromProvider,
	"opentelekomcloud_dis_app_v2":        config.IdentifierFromProvider,
	"opentelekomcloud_dis_checkpoint_v2": config.IdentifierFromProvider,
	"opentelekomcloud_dis_dump_task_v2":  config.IdentifierFromProvider,

	// DNS
	"opentelekomcloud_dns_ptrrecord_v2": config.IdentifierFromProvider,
	"opentelekomcloud_dns_recordset_v2": config.IdentifierFromProvider,
	"opentelekomcloud_dns_zone_v2":      config.IdentifierFromProvider,

	// DMS
	"opentelekomcloud_dms_instance_v1":                  config.IdentifierFromProvider,
	"opentelekomcloud_dms_instance_v2":                  config.IdentifierFromProvider,
	"opentelekomcloud_dms_topic_v1":                     config.IdentifierFromProvider,
	"opentelekomcloud_dms_user_v2":                      config.IdentifierFromProvider,
	"opentelekomcloud_dms_user_permission_v1":           config.IdentifierFromProvider,
	"opentelekomcloud_dms_consumer_group_v2":            config.IdentifierFromProvider,
	"opentelekomcloud_dms_dedicated_instance_v2":        config.IdentifierFromProvider,
	"opentelekomcloud_dms_reassign_partitions_v2":       config.IdentifierFromProvider,
	"opentelekomcloud_dms_smart_connect_v2":             config.IdentifierFromProvider,
	"opentelekomcloud_dms_smart_connect_task_v2":        config.IdentifierFromProvider,
	"opentelekomcloud_dms_smart_connect_task_action_v2": config.IdentifierFromProvider,

	// DRS
	"opentelekomcloud_drs_task_v3": config.IdentifierFromProvider,

	// DWS
	"opentelekomcloud_dws_cluster_v1": config.IdentifierFromProvider,

	// ER
	"opentelekomcloud_er_association_v3":    config.IdentifierFromProvider,
	"opentelekomcloud_er_instance_v3":       config.IdentifierFromProvider,
	"opentelekomcloud_er_flow_log_v3":       config.IdentifierFromProvider,
	"opentelekomcloud_er_propagation_v3":    config.IdentifierFromProvider,
	"opentelekomcloud_er_static_route_v3":   config.IdentifierFromProvider,
	"opentelekomcloud_er_route_table_v3":    config.IdentifierFromProvider,
	"opentelekomcloud_er_vpc_attachment_v3": config.IdentifierFromProvider,

	// FG
	"opentelekomcloud_fgs_function_v2":            config.IdentifierFromProvider,
	"opentelekomcloud_fgs_async_invoke_config_v2": config.IdentifierFromProvider,
	"opentelekomcloud_fgs_event_v2":               config.IdentifierFromProvider,
	"opentelekomcloud_fgs_trigger_v2":             config.IdentifierFromProvider,
	"opentelekomcloud_fgs_dependency_version_v2":  config.IdentifierFromProvider,

	// FW
	"opentelekomcloud_fw_firewall_group_v2": config.IdentifierFromProvider,
	"opentelekomcloud_fw_policy_v2":         config.IdentifierFromProvider,
	"opentelekomcloud_fw_rule_v2":           config.IdentifierFromProvider,

	// GEMINIDB
	"opentelekomcloud_gemini_instance_v3": config.IdentifierFromProvider,

	// HSS
	"opentelekomcloud_hss_host_group_v5":      config.IdentifierFromProvider,
	"opentelekomcloud_hss_host_protection_v5": config.IdentifierFromProvider,

	// IDENTITY
	"opentelekomcloud_identity_agency_v3":                config.IdentifierFromProvider,
	"opentelekomcloud_identity_credential_v3":            config.IdentifierFromProvider,
	"opentelekomcloud_identity_group_v3":                 config.IdentifierFromProvider,
	"opentelekomcloud_identity_group_membership_v3":      config.IdentifierFromProvider,
	"opentelekomcloud_identity_mapping_v3":               config.IdentifierFromProvider,
	"opentelekomcloud_identity_project_v3":               config.IdentifierFromProvider,
	"opentelekomcloud_identity_protocol_v3":              config.IdentifierFromProvider,
	"opentelekomcloud_identity_provider_v3":              config.IdentifierFromProvider,
	"opentelekomcloud_identity_provider":                 config.IdentifierFromProvider,
	"opentelekomcloud_identity_role_v3":                  config.IdentifierFromProvider,
	"opentelekomcloud_identity_login_policy_v3":          config.IdentifierFromProvider,
	"opentelekomcloud_identity_password_policy_v3":       config.IdentifierFromProvider,
	"opentelekomcloud_identity_protection_policy_v3":     config.IdentifierFromProvider,
	"opentelekomcloud_identity_role_assignment_v3":       config.IdentifierFromProvider,
	"opentelekomcloud_identity_user_group_membership_v3": config.IdentifierFromProvider,
	"opentelekomcloud_identity_user_v3":                  config.IdentifierFromProvider,

	// IMAGE
	"opentelekomcloud_images_image_access_accept_v2": config.IdentifierFromProvider,
	"opentelekomcloud_images_image_access_v2":        config.IdentifierFromProvider,
	"opentelekomcloud_images_image_v2":               config.IdentifierFromProvider,
	"opentelekomcloud_ims_data_image_v2":             config.IdentifierFromProvider,
	"opentelekomcloud_ims_image_v2":                  config.IdentifierFromProvider,

	// KMS
	"opentelekomcloud_kms_key_v1":          config.IdentifierFromProvider,
	"opentelekomcloud_kms_grant_v1":        config.IdentifierFromProvider,
	"opentelekomcloud_kms_key_material_v1": config.IdentifierFromProvider,

	// LB
	"opentelekomcloud_lb_certificate_v3":     config.IdentifierFromProvider,
	"opentelekomcloud_lb_ipgroup_v3":         config.IdentifierFromProvider,
	"opentelekomcloud_lb_loadbalancer_v3":    config.IdentifierFromProvider,
	"opentelekomcloud_lb_listener_v3":        config.IdentifierFromProvider,
	"opentelekomcloud_lb_lts_log_v3":         config.IdentifierFromProvider,
	"opentelekomcloud_lb_member_v3":          config.IdentifierFromProvider,
	"opentelekomcloud_lb_monitor_v3":         config.IdentifierFromProvider,
	"opentelekomcloud_lb_policy_v3":          config.IdentifierFromProvider,
	"opentelekomcloud_lb_pool_v3":            config.IdentifierFromProvider,
	"opentelekomcloud_lb_rule_v3":            config.IdentifierFromProvider,
	"opentelekomcloud_lb_security_policy_v3": config.IdentifierFromProvider,

	// LTS
	"opentelekomcloud_logtank_group_v2":             config.IdentifierFromProvider,
	"opentelekomcloud_logtank_topic_v2":             config.IdentifierFromProvider,
	"opentelekomcloud_logtank_transfer_v2":          config.IdentifierFromProvider,
	"opentelekomcloud_lts_group_v2":                 config.IdentifierFromProvider,
	"opentelekomcloud_lts_cce_access_v3":            config.IdentifierFromProvider,
	"opentelekomcloud_lts_cross_account_access_v2":  config.IdentifierFromProvider,
	"opentelekomcloud_lts_host_access_v3":           config.IdentifierFromProvider,
	"opentelekomcloud_lts_host_group_v3":            config.IdentifierFromProvider,
	"opentelekomcloud_lts_keywords_alarm_rule_v2":   config.IdentifierFromProvider,
	"opentelekomcloud_lts_stream_v2":                config.IdentifierFromProvider,
	"opentelekomcloud_lts_notification_template_v2": config.IdentifierFromProvider,
	"opentelekomcloud_lts_transfer_v2":              config.IdentifierFromProvider,
	"opentelekomcloud_lts_quick_search_criteria_v1": config.IdentifierFromProvider,

	// NAT
	"opentelekomcloud_nat_gateway_v2":            config.IdentifierFromProvider,
	"opentelekomcloud_nat_dnat_rule_v2":          config.IdentifierFromProvider,
	"opentelekomcloud_nat_snat_rule_v2":          config.IdentifierFromProvider,
	"opentelekomcloud_private_nat_dnat_rule_v3":  config.IdentifierFromProvider,
	"opentelekomcloud_private_nat_gateway_v3":    config.IdentifierFromProvider,
	"opentelekomcloud_private_nat_snat_rule_v3":  config.IdentifierFromProvider,
	"opentelekomcloud_private_nat_transit_ip_v3": config.IdentifierFromProvider,

	// NETWORKING / VPC
	"opentelekomcloud_networking_floatingip_v2":              config.IdentifierFromProvider,
	"opentelekomcloud_networking_floatingip_associate_v2":    config.IdentifierFromProvider,
	"opentelekomcloud_networking_network_v2":                 config.IdentifierFromProvider,
	"opentelekomcloud_networking_port_v2":                    config.IdentifierFromProvider,
	"opentelekomcloud_networking_port_secgroup_associate_v2": config.IdentifierFromProvider,
	"opentelekomcloud_networking_router_v2":                  config.IdentifierFromProvider,
	"opentelekomcloud_networking_router_interface_v2":        config.IdentifierFromProvider,
	"opentelekomcloud_networking_router_route_v2":            config.IdentifierFromProvider,
	"opentelekomcloud_networking_secgroup_v2":                config.IdentifierFromProvider,
	"opentelekomcloud_networking_secgroup_rule_v2":           config.IdentifierFromProvider,
	"opentelekomcloud_networking_subnet_v2":                  config.IdentifierFromProvider,
	"opentelekomcloud_networking_vip_v2":                     config.IdentifierFromProvider,
	"opentelekomcloud_networking_vip_associate_v2":           config.IdentifierFromProvider,
	"opentelekomcloud_vpc_bandwidth_associate_v2":            config.IdentifierFromProvider,
	"opentelekomcloud_vpc_bandwidth_v2":                      config.IdentifierFromProvider,
	"opentelekomcloud_vpc_eip_v1":                            config.IdentifierFromProvider,
	"opentelekomcloud_vpc_v1":                                config.IdentifierFromProvider,
	"opentelekomcloud_vpc_peering_connection_v2":             config.IdentifierFromProvider,
	"opentelekomcloud_vpc_peering_connection_accepter_v2":    config.IdentifierFromProvider,
	"opentelekomcloud_vpc_route_table_v1":                    config.IdentifierFromProvider,
	"opentelekomcloud_vpc_route_v2":                          config.IdentifierFromProvider,
	"opentelekomcloud_vpc_subnet_v1":                         config.IdentifierFromProvider,
	"opentelekomcloud_vpc_secgroup_v3":                       config.IdentifierFromProvider,
	"opentelekomcloud_vpc_secgroup_rule_v3":                  config.IdentifierFromProvider,
	"opentelekomcloud_vpc_flow_log_v1":                       config.IdentifierFromProvider,

	// OBS / S3
	"opentelekomcloud_obs_bucket":             config.IdentifierFromProvider,
	"opentelekomcloud_obs_bucket_acl":         config.IdentifierFromProvider,
	"opentelekomcloud_obs_bucket_inventory":   config.IdentifierFromProvider,
	"opentelekomcloud_obs_bucket_object":      config.IdentifierFromProvider,
	"opentelekomcloud_obs_bucket_object_acl":  config.IdentifierFromProvider,
	"opentelekomcloud_obs_bucket_policy":      config.IdentifierFromProvider,
	"opentelekomcloud_obs_bucket_replication": config.IdentifierFromProvider,
	"opentelekomcloud_s3_bucket":              config.IdentifierFromProvider,
	"opentelekomcloud_s3_bucket_policy":       config.IdentifierFromProvider,
	"opentelekomcloud_s3_bucket_object":       config.IdentifierFromProvider,

	// RDS
	"opentelekomcloud_rds_backup_v3":              config.IdentifierFromProvider,
	"opentelekomcloud_rds_public_ip_associate_v3": config.IdentifierFromProvider,
	"opentelekomcloud_rds_instance_v3":            config.IdentifierFromProvider,
	"opentelekomcloud_rds_maintenance_v3":         config.IdentifierFromProvider,
	"opentelekomcloud_rds_parametergroup_v3":      config.IdentifierFromProvider,
	"opentelekomcloud_rds_read_replica_v3":        config.IdentifierFromProvider,

	// RMS
	"opentelekomcloud_rms_advanced_query_v1":             config.IdentifierFromProvider,
	"opentelekomcloud_rms_resource_recorder_v1":          config.IdentifierFromProvider,
	"opentelekomcloud_rms_policy_assignment_v1":          config.IdentifierFromProvider,
	"opentelekomcloud_rms_policy_assignment_evaluate_v1": config.IdentifierFromProvider,

	// RTS
	"opentelekomcloud_rts_software_deployment_v1": config.IdentifierFromProvider,
	"opentelekomcloud_rts_software_config_v1":     config.IdentifierFromProvider,
	"opentelekomcloud_rts_stack_v1":               config.IdentifierFromProvider,

	// SFS
	"opentelekomcloud_sfs_file_system_v2":        config.IdentifierFromProvider,
	"opentelekomcloud_sfs_share_access_rules_v2": config.IdentifierFromProvider,
	"opentelekomcloud_sfs_turbo_share_v1":        config.IdentifierFromProvider,

	// SMN
	"opentelekomcloud_smn_topic_v2":           config.IdentifierFromProvider,
	"opentelekomcloud_smn_topic_attribute_v2": config.IdentifierFromProvider,
	"opentelekomcloud_smn_subscription_v2":    config.IdentifierFromProvider,

	// SWR
	"opentelekomcloud_swr_domain_v2":                   config.IdentifierFromProvider,
	"opentelekomcloud_swr_organization_permissions_v2": config.IdentifierFromProvider,
	"opentelekomcloud_swr_organization_v2":             config.IdentifierFromProvider,
	"opentelekomcloud_swr_repository_v2":               config.IdentifierFromProvider,

	// TAURUSDB
	"opentelekomcloud_taurusdb_mysql_backup_v3":           config.IdentifierFromProvider,
	"opentelekomcloud_taurusdb_mysql_instance_v3":         config.IdentifierFromProvider,
	"opentelekomcloud_taurusdb_mysql_quota_v3":            config.IdentifierFromProvider,
	"opentelekomcloud_taurusdb_mysql_sql_control_rule_v3": config.IdentifierFromProvider,
	"opentelekomcloud_taurusdb_mysql_proxy_v3":            config.IdentifierFromProvider,

	// TMS
	"opentelekomcloud_tms_tags_v1":          config.IdentifierFromProvider,
	"opentelekomcloud_tms_resource_tags_v1": config.IdentifierFromProvider,

	// VPCEP
	"opentelekomcloud_vpcep_approval_v1": config.IdentifierFromProvider,
	"opentelekomcloud_vpcep_endpoint_v1": config.IdentifierFromProvider,
	"opentelekomcloud_vpcep_service_v1":  config.IdentifierFromProvider,

	// VPNAAS
	"opentelekomcloud_vpnaas_ipsec_policy_v2":    config.IdentifierFromProvider,
	"opentelekomcloud_vpnaas_service_v2":         config.IdentifierFromProvider,
	"opentelekomcloud_vpnaas_ike_policy_v2":      config.IdentifierFromProvider,
	"opentelekomcloud_vpnaas_endpoint_group_v2":  config.IdentifierFromProvider,
	"opentelekomcloud_vpnaas_site_connection_v2": config.IdentifierFromProvider,

	// EVPN
	"opentelekomcloud_enterprise_vpn_gateway_v5":            config.IdentifierFromProvider,
	"opentelekomcloud_enterprise_vpn_customer_gateway_v5":   config.IdentifierFromProvider,
	"opentelekomcloud_enterprise_vpn_connection_v5":         config.IdentifierFromProvider,
	"opentelekomcloud_enterprise_vpn_connection_monitor_v5": config.IdentifierFromProvider,

	// WAFD
	"opentelekomcloud_waf_dedicated_instance_v1":                 config.IdentifierFromProvider,
	"opentelekomcloud_waf_dedicated_domain_v1":                   config.IdentifierFromProvider,
	"opentelekomcloud_waf_dedicated_policy_v1":                   config.IdentifierFromProvider,
	"opentelekomcloud_waf_dedicated_certificate_v1":              config.IdentifierFromProvider,
	"opentelekomcloud_waf_dedicated_cc_rule_v1":                  config.IdentifierFromProvider,
	"opentelekomcloud_waf_dedicated_anti_crawler_rule_v1":        config.IdentifierFromProvider,
	"opentelekomcloud_waf_dedicated_data_masking_rule_v1":        config.IdentifierFromProvider,
	"opentelekomcloud_waf_dedicated_known_attack_source_rule_v1": config.IdentifierFromProvider,
	"opentelekomcloud_waf_dedicated_web_tamper_rule_v1":          config.IdentifierFromProvider,
	"opentelekomcloud_waf_dedicated_anti_leakage_rule_v1":        config.IdentifierFromProvider,
	"opentelekomcloud_waf_dedicated_alarm_masking_rule_v1":       config.IdentifierFromProvider,
	"opentelekomcloud_waf_dedicated_geo_ip_rule_v1":              config.IdentifierFromProvider,
	"opentelekomcloud_waf_dedicated_blacklist_rule_v1":           config.IdentifierFromProvider,
	"opentelekomcloud_waf_dedicated_precise_protection_rule_v1":  config.IdentifierFromProvider,
	"opentelekomcloud_waf_dedicated_reference_table_v1":          config.IdentifierFromProvider,
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
