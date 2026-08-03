package obs

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/opentelekomcloud/provider-opentelekomcloud/config/common"
)

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("opentelekomcloud_obs_bucket", func(r *config.Resource) {
		// ACL in terraform provider - The acl argument manages supported canned ACLs only: private, public-read, public-read-write, and log-delivery-write. Drift for those canned ACLs is detected on refresh. If the bucket is managed with a custom ACL grant set, use opentelekomcloud_obs_bucket_acl instead of the inline acl argument.
		config.MoveToStatus(r.TerraformResource, "acl")
		r.MetaResource.ArgumentDocs["acl"] = `Deprecated, defaults to ACL=private. Use bucketacls.obs.opentelekomcloud.m.crossplane.io for ACL management. buckets.obs.opentelekomcloud.m.crossplane.io can only observe ACL state, but cannot change it.`
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"acl"},
		}
		r.References["logging.target_bucket"] = config.Reference{
			TerraformName: "opentelekomcloud_obs_bucket",
			Extractor:     common.ObsBucketExtractor,
		}
		r.References["logging.agency"] = config.Reference{
			TerraformName: "opentelekomcloud_identity_agency_v3",
			Extractor:     common.AgencyNameExtractor,
		}
		r.References["server_side_encryption.kms_key_id"] = config.Reference{
			TerraformName:     "opentelekomcloud_kms_key_v1",
			RefFieldName:      "KMSSelectorRef",
			SelectorFieldName: "KMSSelector",
		}
		r.References["event_notifications.topic"] = config.Reference{
			TerraformName: "opentelekomcloud_smn_topic_v2",
		}
	})

	p.AddResourceConfigurator("opentelekomcloud_obs_bucket_acl", func(r *config.Resource) {
		r.MetaResource.Description = `Manages an OBS bucket acl resource within OpenTelekomCloud. For proper functionality need to configure domain_id in credentials.`
		r.References["bucket"] = config.Reference{
			TerraformName: "opentelekomcloud_obs_bucket",
			Extractor:     common.ObsBucketExtractor,
		}
	})

	p.AddResourceConfigurator("opentelekomcloud_obs_bucket_inventory", func(r *config.Resource) {
		r.References["bucket"] = config.Reference{
			TerraformName: "opentelekomcloud_obs_bucket",
			Extractor:     common.ObsBucketExtractor,
		}
		r.References["destination.bucket"] = config.Reference{
			TerraformName: "opentelekomcloud_obs_bucket",
			Extractor:     common.ObsBucketExtractor,
		}
	})

	p.AddResourceConfigurator("opentelekomcloud_obs_bucket_object", func(r *config.Resource) {
		config.MoveToStatus(r.TerraformResource, "acl")
		r.MetaResource.ArgumentDocs["acl"] = `Defaults to ACL=private. Use bucketobjectsacl.obs.opentelekomcloud.m.crossplane.io for ACL management. bucketobjet.obs.opentelekomcloud.m.crossplane.io can only observe ACL state, but cannot change it.`
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"acl"},
		}

		r.References["bucket"] = config.Reference{
			TerraformName: "opentelekomcloud_obs_bucket",
			Extractor:     common.ObsBucketExtractor,
		}
		r.References["sse_kms_key_id"] = config.Reference{
			TerraformName:     "opentelekomcloud_kms_key_v1",
			RefFieldName:      "KMSSelectorRef",
			SelectorFieldName: "KMSSelector",
		}
	})

	p.AddResourceConfigurator("opentelekomcloud_obs_bucket_object_acl", func(r *config.Resource) {
		r.References["bucket"] = config.Reference{
			TerraformName: "opentelekomcloud_obs_bucket",
			Extractor:     common.ObsBucketExtractor,
		}
		r.References["key"] = config.Reference{
			TerraformName: "opentelekomcloud_obs_bucket_object",
		}
	})

	p.AddResourceConfigurator("opentelekomcloud_obs_bucket_policy", func(r *config.Resource) {
		r.References["bucket"] = config.Reference{
			TerraformName: "opentelekomcloud_obs_bucket",
			Extractor:     common.ObsBucketExtractor,
		}
		r.MetaResource.ArgumentDocs["policy"] = `The text of the policy. Supports dynamic value substitution using the placeholder ${this.bucket} which will be replaced with the resolved bucket name from spec.forProvider.bucketSelector. Example: "Resource": ["${this.bucket}/*"]. Alternatively if you choose to configure a different bucket you need to provide bucket/resource as a string. Example: "Resource": ["mybucket/*"]`

		// common.PolicyResourceReplacer replaces ${this.bucket} placeholder in policy.resource with the resolved bucket name from bucketSelector.
		// Policy:  {"Statement":[{"Effect":"Allow","Resource":["${this.bucket}/*"]}]}
		// BECOMES
		// Policy:  {"Statement":[{"Effect":"Allow","Resource":["xp-test-55555/*"]}]}
		r.TerraformConversions = append(r.TerraformConversions, common.PolicyResourceReplacer{})
	})
}
