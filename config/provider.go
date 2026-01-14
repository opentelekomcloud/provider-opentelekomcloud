/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	apigwCluster "github.com/opentelekomcloud/provider-opentelekomcloud/config/cluster/apigw"
	asmCluster "github.com/opentelekomcloud/provider-opentelekomcloud/config/cluster/asm"
	cbrCluster "github.com/opentelekomcloud/provider-opentelekomcloud/config/cluster/cbr"
	cssCluster "github.com/opentelekomcloud/provider-opentelekomcloud/config/cluster/css"
	ctsCluster "github.com/opentelekomcloud/provider-opentelekomcloud/config/cluster/cts"
	cwfCluster "github.com/opentelekomcloud/provider-opentelekomcloud/config/cluster/cwf"
	ddmCluster "github.com/opentelekomcloud/provider-opentelekomcloud/config/cluster/ddm"
	drsCluster "github.com/opentelekomcloud/provider-opentelekomcloud/config/cluster/drs"
	dwsCluster "github.com/opentelekomcloud/provider-opentelekomcloud/config/cluster/dws"
	erCluster "github.com/opentelekomcloud/provider-opentelekomcloud/config/cluster/er"
	evpnCluster "github.com/opentelekomcloud/provider-opentelekomcloud/config/cluster/evpn"
	hssCluster "github.com/opentelekomcloud/provider-opentelekomcloud/config/cluster/hss"
	rmsCluster "github.com/opentelekomcloud/provider-opentelekomcloud/config/cluster/rms"
	rtsCluster "github.com/opentelekomcloud/provider-opentelekomcloud/config/cluster/rts"
	swrCluster "github.com/opentelekomcloud/provider-opentelekomcloud/config/cluster/swr"
	taurusdbCluster "github.com/opentelekomcloud/provider-opentelekomcloud/config/cluster/taurusdb"
	tmsCluster "github.com/opentelekomcloud/provider-opentelekomcloud/config/cluster/tms"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"

	blockstorageCluster "github.com/opentelekomcloud/provider-opentelekomcloud/config/cluster/blockstorage"
	cceCluster "github.com/opentelekomcloud/provider-opentelekomcloud/config/cluster/cce"
	computeCluster "github.com/opentelekomcloud/provider-opentelekomcloud/config/cluster/compute"
	dcsCluster "github.com/opentelekomcloud/provider-opentelekomcloud/config/cluster/dcs"
	ddsCluster "github.com/opentelekomcloud/provider-opentelekomcloud/config/cluster/dds"
	dehCluster "github.com/opentelekomcloud/provider-opentelekomcloud/config/cluster/deh"
	disCluster "github.com/opentelekomcloud/provider-opentelekomcloud/config/cluster/dis"
	dmsCluster "github.com/opentelekomcloud/provider-opentelekomcloud/config/cluster/dms"
	dnsCluster "github.com/opentelekomcloud/provider-opentelekomcloud/config/cluster/dns"
	fgCluster "github.com/opentelekomcloud/provider-opentelekomcloud/config/cluster/fg"
	fwCluster "github.com/opentelekomcloud/provider-opentelekomcloud/config/cluster/fw"
	identityCluster "github.com/opentelekomcloud/provider-opentelekomcloud/config/cluster/identity"
	imageCluster "github.com/opentelekomcloud/provider-opentelekomcloud/config/cluster/image"
	lbCluster "github.com/opentelekomcloud/provider-opentelekomcloud/config/cluster/lb"
	natCluster "github.com/opentelekomcloud/provider-opentelekomcloud/config/cluster/nat"
	networkingCluster "github.com/opentelekomcloud/provider-opentelekomcloud/config/cluster/networking"
	obsCluster "github.com/opentelekomcloud/provider-opentelekomcloud/config/cluster/obs"
	rdsCluster "github.com/opentelekomcloud/provider-opentelekomcloud/config/cluster/rds"
	sfsCluster "github.com/opentelekomcloud/provider-opentelekomcloud/config/cluster/sfs"
	smnCluster "github.com/opentelekomcloud/provider-opentelekomcloud/config/cluster/smn"
	vpcCluster "github.com/opentelekomcloud/provider-opentelekomcloud/config/cluster/vpc"
	vpcepCluster "github.com/opentelekomcloud/provider-opentelekomcloud/config/cluster/vpcep"
	vpnaasCluster "github.com/opentelekomcloud/provider-opentelekomcloud/config/cluster/vpnaas"
	wafdCluster "github.com/opentelekomcloud/provider-opentelekomcloud/config/cluster/wafd"

	apigwNamespaced "github.com/opentelekomcloud/provider-opentelekomcloud/config/namespaced/apigw"
	asmNamespaced "github.com/opentelekomcloud/provider-opentelekomcloud/config/namespaced/asm"
	cbrNamespaced "github.com/opentelekomcloud/provider-opentelekomcloud/config/namespaced/cbr"
	cssNamespaced "github.com/opentelekomcloud/provider-opentelekomcloud/config/namespaced/css"
	ctsNamespaced "github.com/opentelekomcloud/provider-opentelekomcloud/config/namespaced/cts"
	cwfNamespaced "github.com/opentelekomcloud/provider-opentelekomcloud/config/namespaced/cwf"
	ddmNamespaced "github.com/opentelekomcloud/provider-opentelekomcloud/config/namespaced/ddm"
	drsNamespaced "github.com/opentelekomcloud/provider-opentelekomcloud/config/namespaced/drs"
	dwsNamespaced "github.com/opentelekomcloud/provider-opentelekomcloud/config/namespaced/dws"
	erNamespaced "github.com/opentelekomcloud/provider-opentelekomcloud/config/namespaced/er"
	evpnNamespaced "github.com/opentelekomcloud/provider-opentelekomcloud/config/namespaced/evpn"
	hssNamespaced "github.com/opentelekomcloud/provider-opentelekomcloud/config/namespaced/hss"
	rmsNamespaced "github.com/opentelekomcloud/provider-opentelekomcloud/config/namespaced/rms"
	rtsNamespaced "github.com/opentelekomcloud/provider-opentelekomcloud/config/namespaced/rts"
	swrNamespaced "github.com/opentelekomcloud/provider-opentelekomcloud/config/namespaced/swr"
	taurusdbNamespaced "github.com/opentelekomcloud/provider-opentelekomcloud/config/namespaced/taurusdb"
	tmsNamespaced "github.com/opentelekomcloud/provider-opentelekomcloud/config/namespaced/tms"

	blockstorageNamespaced "github.com/opentelekomcloud/provider-opentelekomcloud/config/namespaced/blockstorage"
	cceNamespaced "github.com/opentelekomcloud/provider-opentelekomcloud/config/namespaced/cce"
	computeNamespaced "github.com/opentelekomcloud/provider-opentelekomcloud/config/namespaced/compute"
	dcsNamespaced "github.com/opentelekomcloud/provider-opentelekomcloud/config/namespaced/dcs"
	ddsNamespaced "github.com/opentelekomcloud/provider-opentelekomcloud/config/namespaced/dds"
	dehNamespaced "github.com/opentelekomcloud/provider-opentelekomcloud/config/namespaced/deh"
	disNamespaced "github.com/opentelekomcloud/provider-opentelekomcloud/config/namespaced/dis"
	dmsNamespaced "github.com/opentelekomcloud/provider-opentelekomcloud/config/namespaced/dms"
	dnsNamespaced "github.com/opentelekomcloud/provider-opentelekomcloud/config/namespaced/dns"
	fgNamespaced "github.com/opentelekomcloud/provider-opentelekomcloud/config/namespaced/fg"
	fwNamespaced "github.com/opentelekomcloud/provider-opentelekomcloud/config/namespaced/fw"
	identityNamespaced "github.com/opentelekomcloud/provider-opentelekomcloud/config/namespaced/identity"
	imageNamespaced "github.com/opentelekomcloud/provider-opentelekomcloud/config/namespaced/image"
	lbNamespaced "github.com/opentelekomcloud/provider-opentelekomcloud/config/namespaced/lb"
	natNamespaced "github.com/opentelekomcloud/provider-opentelekomcloud/config/namespaced/nat"
	networkingNamespaced "github.com/opentelekomcloud/provider-opentelekomcloud/config/namespaced/networking"
	obsNamespaced "github.com/opentelekomcloud/provider-opentelekomcloud/config/namespaced/obs"
	rdsNamespaced "github.com/opentelekomcloud/provider-opentelekomcloud/config/namespaced/rds"
	sfsNamespaced "github.com/opentelekomcloud/provider-opentelekomcloud/config/namespaced/sfs"
	smnNamespaced "github.com/opentelekomcloud/provider-opentelekomcloud/config/namespaced/smn"
	vpcNamespaced "github.com/opentelekomcloud/provider-opentelekomcloud/config/namespaced/vpc"
	vpcepNamespaced "github.com/opentelekomcloud/provider-opentelekomcloud/config/namespaced/vpcep"
	vpnaasNamespaced "github.com/opentelekomcloud/provider-opentelekomcloud/config/namespaced/vpnaas"
	wafdNamespaced "github.com/opentelekomcloud/provider-opentelekomcloud/config/namespaced/wafd"
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
			ApplyGroupOverrides(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		apigwCluster.Configure,
		asmCluster.Configure,
		blockstorageCluster.Configure,
		cbrCluster.Configure,
		cwfCluster.Configure,
		ctsCluster.Configure,
		cceCluster.Configure,
		cssCluster.Configure,
		computeCluster.Configure,
		dcsCluster.Configure,
		ddsCluster.Configure,
		ddmCluster.Configure,
		dehCluster.Configure,
		disCluster.Configure,
		dmsCluster.Configure,
		dnsCluster.Configure,
		drsCluster.Configure,
		dwsCluster.Configure,
		erCluster.Configure,
		hssCluster.Configure,
		rmsCluster.Configure,
		rtsCluster.Configure,
		swrCluster.Configure,
		taurusdbCluster.Configure,
		tmsCluster.Configure,
		evpnCluster.Configure,
		fgCluster.Configure,
		fwCluster.Configure,
		identityCluster.Configure,
		imageCluster.Configure,
		lbCluster.Configure,
		natCluster.Configure,
		networkingCluster.Configure,
		obsCluster.Configure,
		rdsCluster.Configure,
		sfsCluster.Configure,
		smnCluster.Configure,
		vpcepCluster.Configure,
		vpcCluster.Configure,
		vpnaasCluster.Configure,
		wafdCluster.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}

func GetProviderNamespaced() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("opentelekomcloud.m.crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
			ApplyGroupOverrides(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		apigwNamespaced.Configure,
		asmNamespaced.Configure,
		blockstorageNamespaced.Configure,
		cbrNamespaced.Configure,
		cwfNamespaced.Configure,
		ctsNamespaced.Configure,
		cceNamespaced.Configure,
		cssNamespaced.Configure,
		computeNamespaced.Configure,
		dcsNamespaced.Configure,
		ddsNamespaced.Configure,
		ddmNamespaced.Configure,
		dehNamespaced.Configure,
		disNamespaced.Configure,
		dmsNamespaced.Configure,
		dnsNamespaced.Configure,
		drsNamespaced.Configure,
		dwsNamespaced.Configure,
		erNamespaced.Configure,
		hssNamespaced.Configure,
		rmsNamespaced.Configure,
		rtsNamespaced.Configure,
		swrNamespaced.Configure,
		taurusdbNamespaced.Configure,
		tmsNamespaced.Configure,
		evpnNamespaced.Configure,
		fgNamespaced.Configure,
		fwNamespaced.Configure,
		identityNamespaced.Configure,
		imageNamespaced.Configure,
		lbNamespaced.Configure,
		natNamespaced.Configure,
		networkingNamespaced.Configure,
		obsNamespaced.Configure,
		rdsNamespaced.Configure,
		sfsNamespaced.Configure,
		smnNamespaced.Configure,
		vpcepNamespaced.Configure,
		vpcNamespaced.Configure,
		vpnaasNamespaced.Configure,
		wafdNamespaced.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
