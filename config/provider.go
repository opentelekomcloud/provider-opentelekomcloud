/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	"github.com/opentelekomcloud/provider-opentelekomcloud/config/vpc"

	"github.com/opentelekomcloud/provider-opentelekomcloud/config/blockstorage"
	"github.com/opentelekomcloud/provider-opentelekomcloud/config/cce"
	"github.com/opentelekomcloud/provider-opentelekomcloud/config/compute"
	"github.com/opentelekomcloud/provider-opentelekomcloud/config/dcs"
	"github.com/opentelekomcloud/provider-opentelekomcloud/config/deh"
	"github.com/opentelekomcloud/provider-opentelekomcloud/config/dis"
	"github.com/opentelekomcloud/provider-opentelekomcloud/config/dms"
	"github.com/opentelekomcloud/provider-opentelekomcloud/config/dns"
	"github.com/opentelekomcloud/provider-opentelekomcloud/config/fg"
	"github.com/opentelekomcloud/provider-opentelekomcloud/config/fw"
	"github.com/opentelekomcloud/provider-opentelekomcloud/config/identity"
	"github.com/opentelekomcloud/provider-opentelekomcloud/config/image"
	"github.com/opentelekomcloud/provider-opentelekomcloud/config/lb"
	"github.com/opentelekomcloud/provider-opentelekomcloud/config/nat"
	"github.com/opentelekomcloud/provider-opentelekomcloud/config/networking"
	"github.com/opentelekomcloud/provider-opentelekomcloud/config/obs"
	"github.com/opentelekomcloud/provider-opentelekomcloud/config/rds"
	"github.com/opentelekomcloud/provider-opentelekomcloud/config/sfs"
	"github.com/opentelekomcloud/provider-opentelekomcloud/config/smn"
	"github.com/opentelekomcloud/provider-opentelekomcloud/config/vpcep"
	"github.com/opentelekomcloud/provider-opentelekomcloud/config/vpnaas"
	"github.com/opentelekomcloud/provider-opentelekomcloud/config/wafd"

	ujconfig "github.com/crossplane/upjet/pkg/config"
)

const (
	resourcePrefix = "opentelekomcloud"
	modulePath     = "github.com/opentelekomcloud/provider-opentelekomcloud"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("opentelekomcloud.crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		blockstorage.Configure,
		cce.Configure,
		compute.Configure,
		dcs.Configure,
		deh.Configure,
		dis.Configure,
		dms.Configure,
		dns.Configure,
		fg.Configure,
		fw.Configure,
		identity.Configure,
		image.Configure,
		lb.Configure,
		nat.Configure,
		networking.Configure,
		obs.Configure,
		rds.Configure,
		sfs.Configure,
		smn.Configure,
		vpcep.Configure,
		vpc.Configure,
		vpnaas.Configure,
		wafd.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
