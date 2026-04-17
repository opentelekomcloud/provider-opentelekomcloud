# Import Existing Resources

This guide demonstrates how you can import already existing resources to Crossplane. For an extensive guide please see the official [documentation](https://docs.crossplane.io/latest/guides/import-existing-resources/).

### Prepare the new Crossplane object
Create a new managed resource which matches the `apiVersion` and `Kind` of the resource which you will import under Crossplane management. We will use a simple `OBS` bucket for this which was previously created on the console with the name `crossplane-import`.

```yaml
apiVersion: obs.opentelekomcloud.m.crossplane.io/v1alpha1
kind: Bucket
metadata:
  annotations:
    meta.upbound.io/example-id: obs/v1alpha1/bucket
  labels:
    testing.upbound.io/example-name: b
  namespace: test
spec:
  forProvider:
    acl: private
    bucket: crossplane-import
```

### Import in `observe` only mode first

Add `managementPolicies: ["Observe"]` to the object and name the Kubernetes resource. Add the `crossplane.io/external-name` annotation for the resource. This name must match the name inside the Provider.

```diff
apiVersion: obs.opentelekomcloud.m.crossplane.io/v1alpha1
kind: Bucket
metadata:
  annotations:
    meta.upbound.io/example-id: obs/v1alpha1/bucket
+    crossplane.io/external-name: crossplane-import
  labels:
    testing.upbound.io/example-name: b
+  name: bucket
  namespace: test
spec:
+  managementPolicies: ["Observe"]
  forProvider:
    acl: private
    bucket: crossplane-import
```

```console
kubectl apply -f my_imported_resource.yaml
```

### Check the imported resource

```console
kubectl describe bucket.obs.opentelekomcloud.m.crossplane.io/bucket -ntest
```

```console
Name:         bucket
Namespace:    test
Labels:       testing.upbound.io/example-name=b
Annotations:  crossplane.io/external-name: crossplane-import
              meta.upbound.io/example-id: obs/v1alpha1/bucket
API Version:  obs.opentelekomcloud.m.crossplane.io/v1alpha1
Kind:         Bucket
Metadata:
  Creation Timestamp:  2026-04-16T06:56:39Z
  Finalizers:
    finalizer.managedresource.crossplane.io
  Generation:        2
  Resource Version:  29251606
  UID:               338e38ea-23b8-4a81-8191-b660d4877e44
Spec:
  For Provider:
    Acl:     private
    Bucket:  crossplane-import
  Init Provider:
  Management Policies:
    Observe
  Provider Config Ref:
    Kind:  ClusterProviderConfig
    Name:  default
Status:
  At Provider:
    Bucket:              crossplane-import
    Bucket Domain Name:  crossplane-import.obs.eu-de.otc.t-systems.com
    Bucket Version:      3.0
    Id:                  crossplane-import
    Parallel Fs:         false
    Region:              eu-de
    Storage Class:       STANDARD
    Versioning:          false
  Conditions:
    Last Transition Time:  2026-04-16T06:56:40Z
    Observed Generation:   2
    Reason:                ReconcileSuccess
    Status:                True
    Type:                  Synced
    Last Transition Time:  2026-04-16T06:57:02Z
    Reason:                Available
    Status:                True
    Type:                  Ready
Events:                    <none>
```

### Take control of the imported resource
Once Crossplane discovers the status of the resource the import has been completed and you can start managing the resource from Crossplane. Change the `managementPolicies` field to `[*]`

```diff
apiVersion: obs.opentelekomcloud.m.crossplane.io/v1alpha1
kind: Bucket
metadata:
  annotations:
    meta.upbound.io/example-id: obs/v1alpha1/bucket
    crossplane.io/external-name: crossplane-import
  labels:
    testing.upbound.io/example-name: b
  name: bucket
  namespace: test
spec:
-  managementPolicies: ["Observe"]
+  managementPolicies: ["*"]
  forProvider:
    acl: private
    bucket: crossplane-import
```
