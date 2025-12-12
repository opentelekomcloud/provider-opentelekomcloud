package common

import (
	xpref "github.com/crossplane/crossplane-runtime/pkg/reference"
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/crossplane/upjet/pkg/resource"
)

const (
	// SelfPackagePath is the golang path for this package.
	SelfPackagePath = "github.com/opentelekomcloud/provider-opentelekomcloud/config/common"

	// NetworkIDExtractor is the golang path to ExtractNetworkID function
	// in this package.
	NetworkIDExtractor = SelfPackagePath + ".ExtractNetworkID()"

	// SubnetIDExtractor is the golang path to ExtractSubnetID function
	// in this package.
	SubnetIDExtractor = SelfPackagePath + ".ExtractSubnetID()"

	// EipAddressExtractor is the golang path to ExtractEipAddress function
	// in this package.
	EipAddressExtractor = SelfPackagePath + ".ExtractEipAddress()"

	// FipAddressExtractor is the golang path to ExtractFipAddress function
	// in this package.
	FipAddressExtractor = SelfPackagePath + ".ExtractFipAddress()"

	// SubnetCIDRExtractor is the golang path to ExtractSubnetCIDR function
	// in this package.
	SubnetCIDRExtractor = SelfPackagePath + ".ExtractSubnetCIDR()"

	// AgencyNameExtractor is the golang path to ExtractAgencyName function
	// in this package.
	AgencyNameExtractor = SelfPackagePath + ".ExtractAgencyName()" // AgencyNameExtractor is the golang path to ExtractAgencyName function

	// DisStreamNameExtractor is the golang path to ExtractDisStreamName function
	// in this package.
	DisStreamNameExtractor = SelfPackagePath + ".ExtractDisStreamName()"

	// DisAppNameExtractor is the golang path to ExtractDisAppName function
	// in this package.
	DisAppNameExtractor = SelfPackagePath + ".ExtractDisAppName()"

	// ObsBucketExtractor is the golang path to ExtractObsBucket function
	// in this package.
	ObsBucketExtractor = SelfPackagePath + ".ExtractObsBucket()"

	// UsernameExtractor is the golang path to ExtractUsername function
	// in this package.
	UsernameExtractor = SelfPackagePath + ".ExtractUsername()"
)

// getStringFromMap safely returns a non-empty string value for a key.
func getStringFromMap(o map[string]any, key string) string {
	if v, ok := o[key].(string); ok && v != "" {
		return v
	}
	return ""
}

// firstPublicIP tries to read publicip[0].ip_address from an observation map.
func firstPublicIP(o map[string]any) (string, bool) {
	raw, ok := o["publicip"].([]any)
	if !ok || len(raw) == 0 {
		return "", false
	}
	m, ok := raw[0].(map[string]any)
	if !ok {
		return "", false
	}
	v, ok := m["ip_address"].(string)
	if !ok || v == "" {
		return "", false
	}
	return v, true
}

// ExtractNetworkID extracts the value of `spec.forProvider.network_id`
// from an Observable resource. If mr is not an Observable
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

// ExtractEipAddress extracts the EIP address from an Observable resource's
// observation (status.atProvider). It supports both nested
// publicip[0].ip_address and flat ip_address/address shapes.
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

		if v, ok := firstPublicIP(o); ok {
			return v
		}
		if v := getStringFromMap(o, "ip_address"); v != "" {
			return v
		}
		if v := getStringFromMap(o, "address"); v != "" {
			return v
		}
		return ""
	}
}

// ExtractFipAddress extracts the Floating IP address from an Observable
// resource's observation (status.atProvider). It supports common shapes:
// address and ip_address.
func ExtractFipAddress() xpref.ExtractValueFn {
	return func(mr xpresource.Managed) string {
		tr, ok := mr.(resource.Observable)
		if !ok {
			return ""
		}
		o, err := tr.GetObservation()
		if err != nil {
			return ""
		}
		if v, ok := o["address"].(string); ok && v != "" {
			return v
		}
		if v, ok := o["ip_address"].(string); ok && v != "" {
			return v
		}
		return ""
	}
}

// ExtractSubnetCIDR extracts the value of `spec.forProvider.cidr`
// from an Observable resource. If mr is not an Observable
// resource, returns an empty string.
func ExtractSubnetCIDR() xpref.ExtractValueFn {
	return func(mr xpresource.Managed) string {
		tr, ok := mr.(resource.Observable)
		if !ok {
			return ""
		}
		o, err := tr.GetObservation()
		if err != nil {
			return ""
		}
		if k := o["cidr"]; k != nil {
			return k.(string)
		}
		return ""
	}
}

// ExtractSubnetID extracts the value of `spec.forProvider.subnet_id`
// from an Observable resource. If mr is not an Observable
// resource, returns an empty string.
func ExtractSubnetID() xpref.ExtractValueFn {
	return func(mr xpresource.Managed) string {
		tr, ok := mr.(resource.Observable)
		if !ok {
			return ""
		}
		o, err := tr.GetObservation()
		if err != nil {
			return ""
		}
		if k := o["subnet_id"]; k != nil {
			return k.(string)
		}
		return ""
	}
}

// ExtractAgencyName extracts the value of `spec.forProvider.name`
// from a Terraformed resource. If mr is not a Terraformed
// resource, returns an empty string.
func ExtractAgencyName() xpref.ExtractValueFn {
	return func(mr xpresource.Managed) string {
		tr, ok := mr.(resource.Terraformed)
		if !ok {
			return ""
		}
		o, err := tr.GetParameters()
		if err != nil {
			return ""
		}
		if k := o["name"]; k != nil {
			return k.(string)
		}

		return ""
	}
}

// ExtractDisStreamName extracts the value of `spec.forProvider.name`
// from a Terraformed resource. If mr is not a Terraformed
// resource, returns an empty string.
func ExtractDisStreamName() xpref.ExtractValueFn {
	return func(mr xpresource.Managed) string {
		tr, ok := mr.(resource.Terraformed)
		if !ok {
			return ""
		}
		o, err := tr.GetParameters()
		if err != nil {
			return ""
		}
		if k := o["name"]; k != nil {
			return k.(string)
		}

		return ""
	}
}

// ExtractDisAppName extracts the value of `spec.forProvider.name`
// from a Terraformed resource. If mr is not a Terraformed
// resource, returns an empty string.
func ExtractDisAppName() xpref.ExtractValueFn {
	return func(mr xpresource.Managed) string {
		tr, ok := mr.(resource.Terraformed)
		if !ok {
			return ""
		}
		o, err := tr.GetParameters()
		if err != nil {
			return ""
		}
		if k := o["name"]; k != nil {
			return k.(string)
		}

		return ""
	}
}

// ExtractObsBucket extracts the value of `spec.forProvider.bucket`
// from a Terraformed resource. If mr is not a Terraformed
// resource, returns an empty string.
func ExtractObsBucket() xpref.ExtractValueFn {
	return func(mr xpresource.Managed) string {
		tr, ok := mr.(resource.Terraformed)
		if !ok {
			return ""
		}
		o, err := tr.GetParameters()
		if err != nil {
			return ""
		}
		if k := o["bucket"]; k != nil {
			return k.(string)
		}

		return ""
	}
}

func ExtractUsername() xpref.ExtractValueFn {
	return func(mr xpresource.Managed) string {
		tr, ok := mr.(resource.Terraformed)
		if !ok {
			return ""
		}
		o, err := tr.GetParameters()
		if err != nil {
			return ""
		}
		if k := o["name"]; k != nil {
			return k.(string)
		}

		return ""
	}
}
