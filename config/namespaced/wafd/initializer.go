package wafd

import (
	"context"
	"encoding/json"
	"reflect"

	"github.com/crossplane/crossplane-runtime/v2/pkg/fieldpath"
	"github.com/crossplane/crossplane-runtime/v2/pkg/reconciler/managed"
	xpresource "github.com/crossplane/crossplane-runtime/v2/pkg/resource"
	"github.com/crossplane/upjet/v2/pkg/config"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// newPreciseProtectionContentsInitializer keeps spec and state consistent for
// precise protection conditions that use value_list_id. On every reconcile it
// reads spec.forProvider.conditions and status.atProvider.conditions, matches
// conditions by valueListId and category, and copies the observed contents
// into the matching spec condition. It patches the resource only when the
// contents differ and leaves conditions with explicit contents untouched. This
// removes the ForceNew diff that would otherwise trigger a replacement blocked
// by prevent_destroy.
var newPreciseProtectionContentsInitializer config.NewInitializerFn = func(kube client.Client) managed.Initializer {
	return managed.InitializerFn(func(ctx context.Context, mg xpresource.Managed) error {
		obj, ok := mg.(client.Object)
		if !ok {
			return nil
		}
		original := obj.DeepCopyObject().(client.Object)
		paved, err := fieldpath.PaveObject(obj)
		if err != nil {
			return err
		}
		specRaw, err := paved.GetValue("spec.forProvider.conditions")
		if err != nil {
			return nil
		}
		observedRaw, err := paved.GetValue("status.atProvider.conditions")
		if err != nil {
			return nil
		}
		specConditions, ok := specRaw.([]any)
		if !ok {
			return nil
		}
		observedConditions, ok := observedRaw.([]any)
		if !ok {
			return nil
		}

		changed := false
		for _, specItem := range specConditions {
			specCondition, ok := specItem.(map[string]any)
			if !ok {
				continue
			}
			valueListID, _ := specCondition["valueListId"].(string)
			if valueListID == "" {
				continue
			}
			for _, observedItem := range observedConditions {
				observedCondition, ok := observedItem.(map[string]any)
				if !ok {
					continue
				}
				if observedValueListID, _ := observedCondition["valueListId"].(string); observedValueListID != valueListID {
					continue
				}
				if specCategory, ok := specCondition["category"].(string); ok {
					if observedCategory, ok := observedCondition["category"].(string); ok && specCategory != observedCategory {
						continue
					}
				}
				if !reflect.DeepEqual(specCondition["contents"], observedCondition["contents"]) {
					specCondition["contents"] = observedCondition["contents"]
					changed = true
				}
				break
			}
		}
		if !changed {
			return nil
		}
		if err := paved.SetValue("spec.forProvider.conditions", specConditions); err != nil {
			return err
		}
		pavedBytes, err := paved.MarshalJSON()
		if err != nil {
			return err
		}
		if err := json.Unmarshal(pavedBytes, obj); err != nil {
			return err
		}
		return kube.Patch(ctx, obj, client.MergeFrom(original))
	})
}
