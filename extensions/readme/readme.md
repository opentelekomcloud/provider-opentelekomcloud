# Provider T Cloud Public (formerly Open Telekom Cloud)

[provider-opentelekomcloud](https://github.com/opentelekomcloud/provider-opentelekomcloud) is a Crossplane provider that is built using [Upjet](https://github.com/crossplane/upjet) code generation tools and exposes XRM-conformant managed resources for the [T Cloud Public](https://public.t-cloud.com/en) API. The provider already introduced the v2 crossplane runtime. Cluster scoped `ManagedResources` are now legacy APIs, thus we recommend using the modern `opentelekomcloud.m.crossplane.io` namespaced APIs instead.

## Provider Resources Overview

`provider-opentelekomcloud` is built on top of the [Terraform T Cloud Public provider](https://github.com/opentelekomcloud/terraform-provider-opentelekomcloud)
. This means that all resources supported by the Terraform provider are also configurable through our Crossplane provider.

Please note that some services are not yet fully configured. While you can still provision and manage these services, dynamic value assignment is not configured for them. In such cases, cross-resource identifiers must be configured manually. You can check [this](https://github.com/opentelekomcloud/provider-opentelekomcloud/issues/7) issue tracker to see the status of the services.

## Getting Started

For an up to date and detailed installation guide please follow the [getting started](https://github.com/opentelekomcloud/provider-opentelekomcloud#getting-started) guide.

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please open an [issue](https://github.com/opentelekomcloud/provider-opentelekomcloud/issues).
