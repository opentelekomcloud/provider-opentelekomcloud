package dis

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/opentelekomcloud/provider-opentelekomcloud/config/common"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("opentelekomcloud_dis_stream_v2", func(r *config.Resource) {
		r.UseAsync = true
	})
	p.AddResourceConfigurator("opentelekomcloud_dis_dump_task_v2", func(r *config.Resource) {
		r.UseAsync = true
		r.References["stream_name"] = config.Reference{
			TerraformName: "opentelekomcloud_dis_stream_v2",
			Extractor:     common.DisStreamNameExtractor,
		}
		r.References["obs_destination_descriptor.obs_bucket_path"] = config.Reference{
			TerraformName: "opentelekomcloud_obs_bucket",
			Extractor:     common.ObsBucketExtractor,
		}
	})
	p.AddResourceConfigurator("opentelekomcloud_dis_checkpoint_v2", func(r *config.Resource) {
		r.UseAsync = true
		r.References["app_name"] = config.Reference{
			TerraformName: "opentelekomcloud_dis_app_v2",
			Extractor:     common.DisAppNameExtractor,
		}
		r.References["stream_name"] = config.Reference{
			TerraformName: "opentelekomcloud_dis_stream_v2",
			Extractor:     common.DisStreamNameExtractor,
		}
	})
	p.AddResourceConfigurator("opentelekomcloud_dis_app_v2", func(r *config.Resource) {
		r.UseAsync = true
	})
}
