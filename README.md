
# Provider opentelekomcloud

`provider-opentelekomcloud` is a [Crossplane](https://crossplane.io/) provider that is built using [Upjet](https://github.com/crossplane/upjet) code generation tools and exposes XRM-conformant managed resources for the opentelekomcloud API. The provider has been upgraded to support Crossplane v2, which introduced a lot of changes and new features like namespaced [ManagedResouces](https://docs.crossplane.io/latest/managed-resources/managed-resources/). Cluster scoped MRs are now legacy APIs, thus we recommend using the modern `opentelekomcloud.m.crossplane.io` namespaced APIs instead. For more information please check [What’s New in v2?](https://docs.crossplane.io/latest/whats-new/)

## Getting Started

### Install provider-opentelekomcloud

#### Install Crossplane

Before you begin, make sure you have the following tools installed:
- [kubectl](https://kubernetes.io/docs/reference/kubectl/kubectl/)
- [helm](https://helm.sh/docs/intro/install/)
- [up](https://github.com/upbound/up)

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

```console
helm install crossplane --namespace crossplane-system crossplane-stable/crossplane 
```

After installation, verify that Crossplane is running correctly:

```console
helm list -n crossplane-system
```

```console
kubectl get all -n crossplane-system
```

#### Install the Provider
> [!NOTE]
> The provider ships hundreds of ManagedResouces by default, which will increase the load on `kube-apiserver`. Please consider using [MRAP](https://docs.crossplane.io/latest/managed-resources/managed-resource-activation-policies/) to only activate needed resources.

Install the provider by using the following command after changing the image tag to the [latest release](https://marketplace.upbound.io/providers/opentelekomcloud/provider-opentelekomcloud):

```console
up ctp provider install opentelekomcloud/provider-opentelekomcloud:${version}
```

Alternatively, you can use declarative installation:

```console
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-opentelekomcloud
spec:
  package: xpkg.upbound.io/opentelekomcloud/provider-opentelekomcloud:${version}
EOF
```

Notice that in this example Provider resource is referencing ControllerConfig with debug disabled.

You can see the API reference [here](https://marketplace.upbound.io/providers/opentelekomcloud/provider-opentelekomcloud/latest).

### Configure provider-opentelekomcloud

ProviderConfig setup with secret:

```console
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Secret
metadata:
  name: provider-opentelekomcloud-creds
  namespace: app
type: Opaque
stringData:
  credentials: |
    {
      "user_name": "admin",
      "password": "t0ps3cr3t11",
      "auth_url": "https://iam.eu-de.otc.t-systems.com/v3",
      "domain_name": "...",
      "tenant_name": "...",
      "swauth": "false",
      "allow_reauth": "true",
      "max_retries": "2",
      "max_backoff_retries": "6",
      "backoff_retry_timeout": "60",
      "insecure": "false"
    }
---
apiVersion: opentelekomcloud.m.crossplane.io/v1beta1
kind: ProviderConfig
metadata:
  name: namespaced-providerconf
  namespace: app
spec:
  credentials:
    source: Secret
    secretRef:
      name: provider-opentelekomcloud-creds
      namespace: crossplane-system
      key: credentials
EOF
```

Reference the `ProviderConfig` in the MR:

```console
apiVersion: obs.opentelekomcloud.m.crossplane.io/v1alpha1
kind: Bucket
metadata:
  annotations:
    meta.upbound.io/example-id: obs/v1alpha1/bucket
  labels:
    testing.upbound.io/example-name: b
  name: b
  namespace: app
spec:
  forProvider:
    acl: private
    bucket: crossplane-test
    tags:
      Env: Test
      foo: bar
      managed: xplane
  providerConfigRef:
    name: namespaced-providerconf
    kind: ProviderConfig
```

### Running provider in debug mode

To enable debug mode, create the following ControllerConfig and reference it in your Provider resource:

```console
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1alpha1
kind: ControllerConfig
metadata:
  name: debug-config
spec:
  args:
    - --debug
---
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-opentelekomcloud
spec:
  package: xpkg.upbound.io/opentelekomcloud/provider-opentelekomcloud:v0.3.0
  controllerConfigRef:
    name: debug-config
EOF
```

## Developing

Run the code-generation pipeline:

```console
go run cmd/generator/main.go "$PWD"
```

Apply generated CRDS
```yaml
kubectl apply -f package/crds
```

Run against a Kubernetes cluster:

```console
make run
```

Build, push, and install:

```console
make all
```

Build binary:

```console
make build
```

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please open an [issue](https://github.com/opentelekomcloud/provider-opentelekomcloud/issues).
