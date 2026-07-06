# Managed Resource Activation Policies

When installing Crossplane with default values it will activate all available APIs and controllers. In most cases this will lead hundreds of unused and unnecessary resources in `kube-api`.
If you are just starting with Crossplane we recommend activating only the modern APIs when installing Crossplane:

```console
helm install crossplane crossplane-stable/crossplane \
  --set provider.defaultActivations={"*.opentelekomcloud.m.crossplane.io"} \
-n crossplane-system
```

If any of the legacy APIs are required you can activate them one-by-one

```diff
apiVersion: apiextensions.crossplane.io/v1alpha1
kind: ManagedResourceActivationPolicy
metadata:
  name: otc-modern-resources
spec:
  activate:
    - "*.opentelekomcloud.m.crossplane.io"
+    - "instancev1s.ecs.opentelekomcloud.crossplane.io"
```


