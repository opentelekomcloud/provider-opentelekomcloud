package config

import (
	"regexp"
	"strings"
	"testing"

	"github.com/opentelekomcloud/provider-opentelekomcloud/config/namespaced/dms"
)

func TestDMSConfiguration(t *testing.T) {
	t.Run("cluster preserves existing resources without TopicV2", func(t *testing.T) {
		p := GetProvider()
		if _, ok := p.Resources[dms.TFTopicV2]; ok {
			t.Error("TopicV2 must not be generated for cluster scope")
		}
		if _, ok := p.Resources["opentelekomcloud_dms_topic_v1"]; !ok {
			t.Error("existing TopicV1 must remain available for cluster scope")
		}
	})

	p := GetProviderNamespaced()
	t.Run("TopicV2 is available in namespaced scope", func(t *testing.T) {
		if _, ok := p.Resources[dms.TFTopicV2]; !ok {
			t.Error("TopicV2 must be generated for namespaced scope")
		}
	})

	t.Run("ID references use names without ID", func(t *testing.T) {
		idSuffix := regexp.MustCompile(`(?i)id(ref|refs|selector)$`)
		for name, resource := range p.Resources {
			if !strings.HasPrefix(name, "opentelekomcloud_dms_") {
				continue
			}
			for field, ref := range resource.References {
				if !strings.HasSuffix(field, "_id") && !strings.HasSuffix(field, "_ids") {
					continue
				}
				if ref.RefFieldName == "" || ref.SelectorFieldName == "" ||
					idSuffix.MatchString(ref.RefFieldName) || idSuffix.MatchString(ref.SelectorFieldName) {
					t.Errorf("%s.%s: expected explicit ID-free reference names, got %q and %q",
						name, field, ref.RefFieldName, ref.SelectorFieldName)
				}
			}
		}
	})

	t.Run("references target the correct resources", func(t *testing.T) {
		for _, tc := range []struct {
			resource string
			field    string
			target   string
			ref      string
			selector string
		}{
			{"opentelekomcloud_dms_topic_v1", "instance_id", "opentelekomcloud_dms_instance_v1", "InstanceRef", "InstanceSelector"},
			{dms.TFTopicV2, "instance_id", "opentelekomcloud_dms_dedicated_instance_v2", "InstanceRef", "InstanceSelector"},
			{"opentelekomcloud_dms_dedicated_instance_v2", "network_id", "opentelekomcloud_vpc_subnet_v1", "SubnetRef", "SubnetSelector"},
			{"opentelekomcloud_dms_smart_connect_task_v2", "source_task.peer_instance_id", "opentelekomcloud_dms_dedicated_instance_v2", "PeerInstanceRef", "PeerInstanceSelector"},
			{"opentelekomcloud_dms_smart_connect_task_action_v2", "task_id", "opentelekomcloud_dms_smart_connect_task_v2", "TaskRef", "TaskSelector"},
		} {
			resource, ok := p.Resources[tc.resource]
			if !ok {
				t.Errorf("resource %s is missing", tc.resource)
				continue
			}
			ref, ok := resource.References[tc.field]
			if !ok {
				t.Errorf("reference %s.%s is missing", tc.resource, tc.field)
				continue
			}
			if ref.TerraformName != tc.target || ref.RefFieldName != tc.ref || ref.SelectorFieldName != tc.selector {
				t.Errorf("%s.%s: got target/ref/selector %q/%q/%q, want %q/%q/%q",
					tc.resource, tc.field, ref.TerraformName, ref.RefFieldName, ref.SelectorFieldName,
					tc.target, tc.ref, tc.selector)
			}
		}
	})
}
