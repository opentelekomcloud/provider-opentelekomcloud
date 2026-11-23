# T Cloud Public Provider Getting Started Guide

> [!NOTE]
> This is a quick start guide to get you up and running in minutes, for a comprehensive introduction and guide please see [Architecture Center](https://arch.otc-service.com/docs/blueprints/by-use-case/platform-engineering/crossplane/introduction-crossplane-provider).  

You will need some flavor of kubernetes to start using Crossplane. You can use [kind](https://github.com/kubernetes-sigs/kind) for testing or any managed kubernetes service.

```console
kind create cluster --name local-dev
```

### Installing the provider

#### Install Crossplane

Before you begin, make sure you have the following tools installed:
- [kubectl](https://kubernetes.io/docs/reference/kubectl/kubectl/)
- [helm](https://helm.sh/docs/intro/install/)

Start by creating a namespace for Crossplane:

```console
kubectl create namespace crossplane-system
```

Next, add the Crossplane Helm repository and update it:

```console
helm repo add crossplane-stable https://charts.crossplane.io/stable
helm repo update
```

Finally, install Crossplane using Helm:
#### Install the T Cloud Public provider
> [!NOTE]
> The provider ships hundreds of ManagedResouces by default, which will increase the load on `kube-apiserver`. Please consider using [MRAP](https://docs.crossplane.io/latest/managed-resources/managed-resource-activation-policies/) to only activate needed resources. Crossplane by [default](https://docs.crossplane.io/latest/managed-resources/managed-resource-activation-policies/#activate-everything-default-behavior) activates all controllers and APIs, even the legacy ones. If you are just started using Crossplane we recommend installing with `provider.defaultActivations={"*.opentelekomcloud.m.crossplane.io"} ` to avoid a large number of legacy resources.


```console
helm install crossplane crossplane-stable/crossplane \
  --set provider.defaultActivations={"*.opentelekomcloud.m.crossplane.io"} \
-n crossplane-system
```

After installation, verify that Crossplane is running correctly:

```console
kubectl -n crossplane-system wait --for=condition=Available deployment --all --timeout=5m
```

#### Install the T Cloud Public provider

Install the provider by using the following command after changing the image tag to the [latest release](https://marketplace.upbound.io/providers/opentelekomcloud/provider-opentelekomcloud):

```console
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-opentelekomcloud
spec:
  package: xpkg.upbound.io/opentelekomcloud/provider-opentelekomcloud:v0.9.0
EOF
```

### Configure the provider

[ClusterProviderConfig](https://github.com/opentelekomcloud/provider-opentelekomcloud/tree/main/examples/namespaced/providerconfig) setup with secret:

```console
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Secret
metadata:
  name: provider-secret
  namespace: crossplane-system
type: Opaque
stringData:
  credentials: |
    {
      "access_key": "xxxxxxxxxxx",
      "secret_key": "xxxxxxxxxxx",
      "auth_url": "https://iam.eu-de.otc.t-systems.com/v3",
      "domain_name": "OTCxxxxx",
      "domain_id": "xxxxxxxxxxx",
      "tenant_id": "xxxxxxxxxxx",
      "swauth": "false",
      "allow_reauth": "true",
      "max_retries": "2",
      "max_backoff_retries": "6",
      "backoff_retry_timeout": "60",
      "insecure": "false"
    }
---
apiVersion: opentelekomcloud.m.crossplane.io/v1beta1
kind: ClusterProviderConfig
metadata:
  name: default
spec:
  credentials:
    source: Secret
    secretRef:
      name: provider-secret
      namespace: crossplane-system
      key: credentials
EOF
```

Start deploying `ManagedResources`:

```console
apiVersion: obs.opentelekomcloud.m.crossplane.io/v1alpha1
kind: Bucket
metadata:
  annotations:
    meta.upbound.io/example-id: obs/v1alpha1/bucket
  labels:
    testing.upbound.io/example-name: b
  name: b
  namespace: default
spec:
  forProvider:
    acl: private
    bucket: crossplane-test
    tags:
      Env: Test
      foo: bar
      managed: xplane
```

Check the state of the managed resources:

```console
kubectl get bucket.obs.opentelekomcloud.m.crossplane.io
```

```console
NAME                                                                SYNCED   READY   EXTERNAL-NAME     AGE
bucket.obs.opentelekomcloud.m.crossplane.io/b   True     True    crossplane-test   3m22s
```

