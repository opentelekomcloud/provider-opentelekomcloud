# Upgrading the provider

To upgrade an existing Provider edit the installed `Provider Package` by either applying a new Provider manifest or with `kubectl edit providers.pkg.crossplane.io`.

```console
kubectl get providers.pkg.crossplane.io
NAME                        INSTALLED   HEALTHY   PACKAGE                                                             AGE
provider-opentelekomcloud   True        True      xpkg.upbound.io/opentelekomcloud/provider-opentelekomcloud:v0.7.0   42d
```


Update the version number in the Provider’s `spec.package`` and apply the change. Crossplane installs the new image and creates a new `ProviderRevision`.
```diff

apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-opentelekomcloud
spec:
- package: xpkg.upbound.io/opentelekomcloud/provider-opentelekomcloud:v0.7.0
+ package: xpkg.upbound.io/opentelekomcloud/provider-opentelekomcloud:v0.8.0
```

```console
kubectl apply -f provider.yaml
```


The `ProviderRevision` allows Crossplane to store deprecated Provider CRDs without removing them until you decide.

```console
kubectl get providerrevisions
NAME                                     HEALTHY   RUNTIME   IMAGE                                                               STATE      AGE
provider-opentelekomcloud-98495c1871db             False     xpkg.upbound.io/opentelekomcloud/provider-opentelekomcloud:v0.8.0   Active     25s
provider-opentelekomcloud-e329557ed634   True      True      xpkg.upbound.io/opentelekomcloud/provider-opentelekomcloud:v0.7.0   Inactive   42d
```

## Rolling back

To roll back to an earlier version of the provider you can update the provider manifest with an earlier version and apply it.

```diff
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-opentelekomcloud
spec:
+ package: xpkg.upbound.io/opentelekomcloud/provider-opentelekomcloud:v0.7.0
- package: xpkg.upbound.io/opentelekomcloud/provider-opentelekomcloud:v0.8.0
```

```console
kubectl apply -f provider.yaml
```

```console
kubectl get providerrevisions
NAME                                     HEALTHY   RUNTIME   IMAGE                                                               STATE      AGE
provider-opentelekomcloud-98495c1871db   True      True      xpkg.upbound.io/opentelekomcloud/provider-opentelekomcloud:v0.8.0   Inactive   99s
provider-opentelekomcloud-e329557ed634   True      True      xpkg.upbound.io/opentelekomcloud/provider-opentelekomcloud:v0.7.0   Active     42d
```
