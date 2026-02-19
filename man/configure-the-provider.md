# Configure the provider

## Changing the default values of the provider

You can customize the provider’s default behavior using `DeploymentRuntimeConfig`.

For example, the provider performs a drift check every 10 minutes by default. If you need faster reconciliation, use the `--poll` flag to reduce the interval(will put more load on APIs be careful).

You can review the default settings and feature flags in `cmd/provider/main.go`.

For more details about runtime and provider configuration, refer to the official Crossplane [documentation](https://docs.crossplane.io/latest/packages/providers/#configure-a-provider)

### EG: enable logging and faster reconcile

```console

apiVersion: pkg.crossplane.io/v1beta1
kind: DeploymentRuntimeConfig
metadata:
  name: fast-reconcile
spec:
  deploymentTemplate:
    spec:
      selector: {}
      template:
        spec:
          containers:
            - name: package-runtime
              args:
                - --poll=3m
                - --debug
---

apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-opentelekomcloud
spec:
  package: xpkg.upbound.io/opentelekomcloud/provider-opentelekomcloud:v0.7.0
  runtimeConfigRef:
    name: fast-reconcile
```
