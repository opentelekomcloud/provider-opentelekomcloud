
# Provider opentelekomcloud

`provider-opentelekomcloud` is a [Crossplane](https://crossplane.io/) provider that is built using [Upjet](https://github.com/crossplane/upjet) code generation tools and exposes XRM-conformant managed resources for the opentelekomcloud API. The provider has been upgraded to support Crossplane v2, which introduced a lot of changes and new features like namespaced [ManagedResouces](https://docs.crossplane.io/latest/managed-resources/managed-resources/). Cluster scoped MRs are now legacy APIs, thus we recommend using the modern `opentelekomcloud.m.crossplane.io` namespaced APIs instead. For more information please check [What’s New in v2?](https://docs.crossplane.io/latest/whats-new/)

## Getting Started

You will need some flavor of kubernetes to start using Crossplane. You can use [kind](https://github.com/kubernetes-sigs/kind) for testing or any managed kubernetes service.

```console
kind create cluster --name local-dev
```

### Install provider-opentelekomcloud

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

```console
helm install crossplane --namespace crossplane-system crossplane-stable/crossplane 
```

After installation, verify that Crossplane is running correctly:

```console
helm list -n crossplane-system
```

```console
kubectl -n crossplane-system wait --for=condition=Available deployment --all --timeout=5m
```

#### Install the Provider
> [!NOTE]
> The provider ships hundreds of ManagedResouces by default, which will increase the load on `kube-apiserver`. Please consider using [MRAP](https://docs.crossplane.io/latest/managed-resources/managed-resource-activation-policies/) to only activate needed resources.

Install the provider by using the following command after changing the image tag to the [latest release](https://marketplace.upbound.io/providers/opentelekomcloud/provider-opentelekomcloud):

```console
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-opentelekomcloud
spec:
  package: xpkg.upbound.io/opentelekomcloud/provider-opentelekomcloud:v0.7.0
EOF
```

### Configure provider-opentelekomcloud

ProviderConfig setup with secret:

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
      "user_name": "admin",
      "password": "t0ps3cr3t11",
      "auth_url": "https://iam.eu-de.otc.t-systems.com/v3",
      "domain_name": "OTCxxxxx",
      "tenant_name": "eu-de_project",
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

Start deploying ManagedResources:

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


## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please open an [issue](https://github.com/opentelekomcloud/provider-opentelekomcloud/issues).
