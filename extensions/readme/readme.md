# Upjet-based Crossplane provider for opentelekomcloud

`provider-opentelekomcloud` is a [Crossplane](https://crossplane.io/) provider that is built using [Upjet](https://github.com/crossplane/upjet) code generation tools and exposes XRM-conformant managed resources for the opentelekomcloud API.

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
up ctp provider install opentelekomcloud/provider-opentelekomcloud:v0.3.0
```

Alternatively, you can use declarative installation:

```console
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-opentelekomcloud
spec:
  package: xpkg.upbound.io/opentelekomcloud/provider-opentelekomcloud:v0.3.0
EOF
```

Notice that in this example Provider resource is referencing ControllerConfig with debug disabled.

You can see the API reference [here](https://marketplace.upbound.io/providers/opentelekomcloud/provider-opentelekomcloud/latest).

## Contributing

To contribute bug fixes or features to our provider:

 - Communicate your intent.
 - Make your changes.
 - Test your changes.
 - Update documentation and examples where appropriate.
 - [Open a Pull Request (PR)](https://github.com/opentelekomcloud/provider-opentelekomcloud).

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please open an [issue](https://github.com/opentelekomcloud/provider-opentelekomcloud/issues).

## License
The provider is released under the Apache 2.0 license.