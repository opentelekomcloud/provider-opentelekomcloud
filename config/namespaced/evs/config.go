package evs

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("opentelekomcloud_evs_volume_v3", func(r *config.Resource) {
		r.UseAsync = true
		r.References["kms_id"] = config.Reference{
			TerraformName:     "opentelekomcloud_kms_key_v1",
			SelectorFieldName: "KMSSelector",
			RefFieldName:      "KMSRef",
		}
	})
	p.AddResourceConfigurator("opentelekomcloud_evs_snapshot_v2", func(r *config.Resource) {
		r.UseAsync = true
		r.References["volume_id"] = config.Reference{
			TerraformName:     "opentelekomcloud_evs_volume_v3",
			SelectorFieldName: "VolumeSelector",
			RefFieldName:      "VolumeRef",
		}
	})
}
