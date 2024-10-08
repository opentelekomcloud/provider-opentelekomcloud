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
)

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

// ExtractEipAddress extracts the value of `spec.forProvider.address`
// from a Terraformed resource. If mr is not a Terraformed
// resource, returns an empty string.
func ExtractEipAddress() xpref.ExtractValueFn {
	return func(mr xpresource.Managed) string {
		tr, ok := mr.(resource.Terraformed)
		if !ok {
			return ""
		}
		o, err := tr.GetParameters()
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

// ExtractFipAddress extracts the value of `spec.forProvider.address`
// from a Terraformed resource. If mr is not a Terraformed
// resource, returns an empty string.
func ExtractFipAddress() xpref.ExtractValueFn {
	return func(mr xpresource.Managed) string {
		tr, ok := mr.(resource.Terraformed)
		if !ok {
			return ""
		}
		o, err := tr.GetParameters()
		if err != nil {
			return ""
		}
		if k := o["address"]; k != nil {
			return k.(string)
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
