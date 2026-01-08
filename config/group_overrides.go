package config

import "github.com/crossplane/upjet/v2/pkg/config"

// GroupOverrides maps TF resource names to the desired ShortGroup.
var GroupOverrides = map[string]string{
	"opentelekomcloud_enterprise_vpn_gateway_v5":            "enterprisevpn",
	"opentelekomcloud_enterprise_vpn_customer_gateway_v5":   "enterprisevpn",
	"opentelekomcloud_enterprise_vpn_connection_v5":         "enterprisevpn",
	"opentelekomcloud_enterprise_vpn_connection_monitor_v5": "enterprisevpn",
	"opentelekomcloud_private_nat_dnat_rule_v3":             "privatenat",
	"opentelekomcloud_private_nat_gateway_v3":               "privatenat",
	"opentelekomcloud_private_nat_snat_rule_v3":             "privatenat",
	"opentelekomcloud_private_nat_transit_ip_v3":            "privatenat",
}

// ApplyGroupOverrides sets ShortGroup for the listed resources.
func ApplyGroupOverrides() config.ResourceOption {
	return func(r *config.Resource) {
		if g, ok := GroupOverrides[r.Name]; ok {
			r.ShortGroup = g
		}
	}
}
