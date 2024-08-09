
# Provider opentelekomcloud

`provider-opentelekomcloud` is a [Crossplane](https://crossplane.io/) provider that is built using [Upjet](https://github.com/crossplane/upjet) code generation tools and exposes XRM-conformant managed resources for the opentelekomcloud API.

## 🚨 **Note: This project is still under development and is not recommended for use in production environments.** 🚨

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

Install the provider by using the following command after changing the image tag to the [latest release](https://marketplace.upbound.io/providers/opentelekomcloud/provider-opentelekomcloud):

```console
up ctp provider install opentelekomcloud/provider-opentelekomcloud:v0.2.0
```

Alternatively, you can use declarative installation:

```console
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-opentelekomcloud
spec:
  package: xpkg.upbound.io/opentelekomcloud/provider-opentelekomcloud:v0.2.0
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
  namespace: crossplane-system
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
apiVersion: opentelekomcloud.crossplane.io/v1beta1
kind: ProviderConfig
metadata:
  name: default
spec:
  credentials:
    source: Secret
    secretRef:
      name: provider-opentelekomcloud-creds
      namespace: crossplane-system
      key: credentials
EOF
```

#### Running provider in debug mode

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
  package: xpkg.upbound.io/opentelekomcloud/provider-opentelekomcloud:v0.2.0
  controllerConfigRef:
    name: debug-config
EOF
```

## Developing

Run the code-generation pipeline:

```console
go run cmd/generator/main.go "$PWD"
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
