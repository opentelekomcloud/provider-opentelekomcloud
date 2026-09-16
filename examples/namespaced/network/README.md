# Network test cases

## networking.yaml

Test case: `networking.yaml` - creates a Neutron network, router, subnet, and
connects the router to the subnet via a router interface.

### APIs used

| Kind | Full API name |
|------|---------------|
| NetworkV2 | `networkv2s.networking.opentelekomcloud.m.crossplane.io` |
| RouterV2 | `routerv2s.networking.opentelekomcloud.m.crossplane.io` |
| SubnetV2 | `subnetv2s.networking.opentelekomcloud.m.crossplane.io` |
| RouterInterfaceV2 | `routerinterfacev2s.networking.opentelekomcloud.m.crossplane.io` |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${Rand.RFC1123Subdomain}` | `spec.forProvider.name` of NetworkV2, RouterV2, SubnetV2 | Replace with a unique RFC1123-compatible suffix (or leave for the example generator to substitute) |
| `cidr: 192.168.199.0/24` | SubnetV2 `spec.forProvider.cidr` | Adjust to a CIDR that does not conflict with existing networks |
| `namespace: test` | all resources | All resources are created in the `test` namespace - adjust if needed |

## Prerequisites

This example requires a configured `ClusterProviderConfig` - see [providerconfig](../providerconfig/README.md).
