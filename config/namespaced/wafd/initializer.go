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
		return mirrorPreciseProtectionContents(ctx, kube, obj)
	})
}

func mirrorPreciseProtectionContents(ctx context.Context, kube client.Client, obj client.Object) error {
	original := obj.DeepCopyObject().(client.Object)
	paved, err := fieldpath.PaveObject(obj)
	if err != nil {
		return err
	}
	specConditions, ok := conditionsAt(paved, "spec.forProvider.conditions")
	if !ok {
		return nil
	}
	observedConditions, ok := conditionsAt(paved, "status.atProvider.conditions")
	if !ok {
		return nil
	}
	if !mirrorContents(specConditions, observedConditions) {
		return nil
	}
	return patchConditions(ctx, kube, obj, original, paved, specConditions)
}

func conditionsAt(paved *fieldpath.Paved, path string) ([]any, bool) {
	raw, err := paved.GetValue(path)
	if err != nil {
		return nil, false
	}
	conditions, ok := raw.([]any)
	return conditions, ok
}

func mirrorContents(specConditions, observedConditions []any) bool {
	changed := false
	for _, specItem := range specConditions {
		specCondition, ok := specItem.(map[string]any)
		if !ok {
			continue
		}
		observedCondition := findObservedCondition(specCondition, observedConditions)
		if observedCondition == nil {
			continue
		}
		if reflect.DeepEqual(specCondition["contents"], observedCondition["contents"]) {
			continue
		}
		specCondition["contents"] = observedCondition["contents"]
		changed = true
	}
	return changed
}

func findObservedCondition(specCondition map[string]any, observedConditions []any) map[string]any {
	valueListID := stringValue(specCondition, "valueListId")
	if valueListID == "" {
		return nil
	}
	for _, observedItem := range observedConditions {
		observedCondition, ok := observedItem.(map[string]any)
		if !ok {
			continue
		}
		if stringValue(observedCondition, "valueListId") != valueListID {
			continue
		}
		if category := stringValue(specCondition, "category"); category != "" && stringValue(observedCondition, "category") != category {
			continue
		}
		return observedCondition
	}
	return nil
}

func stringValue(m map[string]any, key string) string {
	value, _ := m[key].(string)
	return value
}

func patchConditions(ctx context.Context, kube client.Client, obj, original client.Object, paved *fieldpath.Paved, specConditions []any) error {
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
}
