
# T Cloud Public Crossplane provider
> [!NOTE]
> **T Cloud Public** is formerly known as **Open Telekom Cloud**.

`provider-opentelekomcloud` is a [Crossplane](https://crossplane.io/) provider that is built using [Upjet](https://github.com/crossplane/upjet) code generation tools and exposes XRM-conformant managed resources for the **T Cloud Public** API. The provider has been upgraded to support Crossplane v2, which introduced a lot of changes and new features like namespaced [ManagedResouces](https://docs.crossplane.io/latest/managed-resources/managed-resources/). Cluster scoped MRs are now legacy APIs, thus we recommend using the modern `opentelekomcloud.m.crossplane.io` namespaced APIs instead. For more information please check [What’s New in v2?](https://docs.crossplane.io/latest/whats-new/)

## Provider Resources Overview

You can find all supported resources bundled in the **T Cloud Public** provider at the [Upbound Marketplace](https://marketplace.upbound.io/providers/opentelekomcloud/provider-opentelekomcloud/).

Please note that some services are not yet fully configured. While you can still provision and manage these services, cross resource value assignment is not configured for them. In such cases, cross-resource identifiers must be configured manually. To see the state of the individual services please see the [service coverage](docs/service-coverage.md).

## Documentation

Best practices and other important provider documentation can be found in the [`docs/`](docs/) directory.

| Document | Description |
| --- | --- |
| [Best practices](docs/best-practices.md) | Recommended practices and API families when using the provider. |
| [Getting started](docs/getting-started.md) | Quick start guide to get up and running in minutes. |
| [Configure the provider](docs/configure-the-provider.md) | Customizing the provider's default behavior. |
| [Import existing resources](docs/import-external-resources.md) | How to import already existing resources into Crossplane. |
| [Managed Resource Activation Policies](docs/mrap.md) | Limiting activated APIs and controllers to the ones you need. |
| [Service coverage](docs/service-coverage.md) | State of service onboarding and coverage in the provider. |
| [Service roadmap](docs/service-roadmap.md) | Planned service onboarding work. |
| [Upgrading the provider](docs/upgrade-the-provider.md) | How to upgrade an installed provider package. |
| [Web Identity authentication](docs/webidentity/README.md) | Workload identity(secretless) authentication for the namespaced APIs. |
| [Namespaced examples](examples/namespaced/) | Verified example manifests for the namespaced APIs, grouped by service. |

## Getting Started

To get quickly up and running with the Provider you can follow our getting started [guide](docs/getting-started.md).

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please open an [issue](https://github.com/opentelekomcloud/provider-opentelekomcloud/issues).  

---
<a href="https://www.aihonestybadge.com" target="_blank" rel="noopener"><img src="https://docs.otc.t-systems.com/_static/images/ai-assisted-magenta.svg" alt="AI Assisted Badge" style="max-width: 95px; height: 25px;" /></a>
