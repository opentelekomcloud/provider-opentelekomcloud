package rds

import (
	xpref "github.com/crossplane/crossplane-runtime/pkg/reference"
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/crossplane/upjet/pkg/config"
	"github.com/crossplane/upjet/pkg/resource"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("opentelekomcloud_rds_instance_v3", func(r *config.Resource) {
		r.UseAsync = true
		r.References["security_group_id"] = config.Reference{
			TerraformName:     "opentelekomcloud_compute_secgroup_v2",
			RefFieldName:      "ComputeSecurityGroupIDRefs",
			SelectorFieldName: "ComputeSecurityGroupIDSelector",
		}
		r.References["vpc_id"] = config.Reference{
			TerraformName: "opentelekomcloud_vpc_v1",
		}
		r.References["subnet_id"] = config.Reference{
			TerraformName: "opentelekomcloud_vpc_subnet_v1",
			Extractor:     ExtractNetworkIDFunc,
		}
		r.References["public_ips"] = config.Reference{
			TerraformName:     "opentelekomcloud_vpc_eip_v1",
			Extractor:         ExtractEipAddressFunc,
			RefFieldName:      "PublicIpsRefs",
			SelectorFieldName: "PublicIpsSelector",
		}
		r.References["param_group_id"] = config.Reference{
			TerraformName: "opentelekomcloud_rds_parametergroup_v3",
		}
	})
	p.AddResourceConfigurator("opentelekomcloud_rds_backup_v3", func(r *config.Resource) {
		r.UseAsync = true
		r.References["instance_id"] = config.Reference{
			TerraformName: "opentelekomcloud_rds_instance_v3",
		}
	})
	p.AddResourceConfigurator("opentelekomcloud_rds_read_replica_v3", func(r *config.Resource) {
		r.UseAsync = true
		r.References["replica_of_id"] = config.Reference{
			TerraformName: "opentelekomcloud_rds_instance_v3",
		}
		r.References["public_ips"] = config.Reference{
			TerraformName:     "opentelekomcloud_vpc_eip_v1",
			Extractor:         ExtractEipAddressFunc,
			RefFieldName:      "PublicIpsRefs",
			SelectorFieldName: "PublicIpsSelector",
		}
	})
}

const (
	// APISPackagePath is the package path for generated APIs root package
	APISPackagePath = "github.com/opentelekomcloud/provider-opentelekomcloud/config/rds"

	// ExtractNetworkIDFunc extracts network_id from subnet resource
	ExtractNetworkIDFunc = APISPackagePath + ".ExtractNetworkID()"
	// ExtractEipAddressFunc extracts ip_address from eip resource
	ExtractEipAddressFunc = APISPackagePath + ".ExtractEipAddress()"
)

// ExtractNetworkID extracts the value of `spec.forProvider.network_id`
// from an Observable resource. If mr is not a Observable
// resource, returns an empty string.
func ExtractNetworkID() xpref.ExtractValueFn {
	return func(mr xpresource.Managed) string {
		tr, ok := mr.(resource.Observable)
		if !ok {
			return ""
		}
		o, err := tr.GetObservation()
		if err != nil {
			return ""
		}

		if k := o["network_id"]; k != nil {
			return k.(string)
		}
		return ""
	}
}

// ExtractEipAddress extracts the value of `spec.forProvider.address`
// from an Observable resource. If mr is not a Observable
// resource, returns an empty string.
func ExtractEipAddress() xpref.ExtractValueFn {
	return func(mr xpresource.Managed) string {
		tr, ok := mr.(resource.Observable)
		if !ok {
			return ""
		}
		o, err := tr.GetObservation()
		if err != nil {
			return ""
		}
		publicIPList := o["publicip"].([]any)
		if len(publicIPList) > 0 {
			publicIP := publicIPList[0].(map[string]any)
			if k := publicIP["ip_address"]; k != nil {
				return k.(string)
			}
		}

		return ""
	}
}
