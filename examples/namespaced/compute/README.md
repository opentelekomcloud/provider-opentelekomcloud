# Compute test cases

## instance.yaml

Test case: `instance.yaml` - creates a VM (InstanceV2) with a network, router,
subnet, router interface, and a floating IP associated to the instance.

### APIs used

| Kind | Full API name |
|------|---------------|
| NetworkV2 | `networkv2s.networking.opentelekomcloud.m.crossplane.io` |
| RouterV2 | `routerv2s.networking.opentelekomcloud.m.crossplane.io` |
| SubnetV2 | `subnetv2s.networking.opentelekomcloud.m.crossplane.io` |
| RouterInterfaceV2 | `routerinterfacev2s.networking.opentelekomcloud.m.crossplane.io` |
| InstanceV2 | `instancev2s.compute.opentelekomcloud.m.crossplane.io` |
| FloatingipV2 | `floatingipv2s.networking.opentelekomcloud.m.crossplane.io` |
| FloatingipAssociateV2 | `floatingipassociatev2s.compute.opentelekomcloud.m.crossplane.io` |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${Rand.RFC1123Subdomain}` | names of NetworkV2, RouterV2, SubnetV2, InstanceV2, FloatingipV2 | Replace with a unique RFC1123-compatible suffix (or leave for the example generator to substitute) |
| `imageId: 765aeae4-7c4a-4bc7-93ff-ad1daa9c2fda` | InstanceV2 `spec.forProvider.imageId` | Replace with the ID of an image that exists in your region |
| `flavorId: s2.large.2` | InstanceV2 `spec.forProvider.flavorId` | Adjust to a flavor available in your region |
| `pool: admin_external_net` | FloatingipV2 `spec.forProvider.pool` | Replace with the external network pool of your project |
| `cidr: 192.168.199.0/24` | SubnetV2 | Adjust to a CIDR that does not conflict with existing networks |
| `namespace: test` | all resources | All resources are created in the `test` namespace - adjust if needed |

## Prerequisites

This example requires a configured `ClusterProviderConfig` - see [providerconfig](../providerconfig/README.md).
