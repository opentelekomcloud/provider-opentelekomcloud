# CCE test cases

## cluster.yaml

Test case: `cluster.yaml` - creates a CCE cluster with an SSH keypair, VPC, subnet,
one dedicated node, and an auto-scaling node pool.

### APIs used

| Kind | Full API name |
|------|---------------|
| KeypairV2 | `keypairv2s.compute.opentelekomcloud.m.crossplane.io` |
| VpcV1 | `vpcv1s.vpc.opentelekomcloud.m.crossplane.io` |
| SubnetV1 | `subnetv1s.vpc.opentelekomcloud.m.crossplane.io` |
| ClusterV3 | `clusterv3s.cce.opentelekomcloud.m.crossplane.io` |
| NodeV3 | `nodev3s.cce.opentelekomcloud.m.crossplane.io` |
| NodePoolV3 | `nodepoolv3s.cce.opentelekomcloud.m.crossplane.io` |
| NodePoolConfigV3 | `nodepoolconfigv3s.cce.opentelekomcloud.m.crossplane.io` |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${Rand.RFC1123Subdomain}` | names of KeypairV2, VpcV1, SubnetV1, ClusterV3, NodeV3, NodePoolV3 | Replace with a unique RFC1123-compatible suffix (or leave for the example generator to substitute) |
| `publicKey` | KeypairV2 `spec.forProvider.publicKey` | Replace the hardcoded RSA public key with your own SSH public key |
| `flavorId: cce.s1.small` | ClusterV3 `spec.forProvider.flavorId` | Adjust to the cluster flavor available in your region |
| `flavorId: s2.large.2` / `flavor: s2.xlarge.2` | NodeV3 / NodePoolV3 | Adjust node flavors to ones available in your region |
| `availabilityZone: eu-de-01` | NodeV3, NodePoolV3 | Change to an AZ in your region |
| `os: HCE OS 2.0` | NodeV3, NodePoolV3 | Change to the OS image available in your region |
| `namespace: test` | all resources | All resources are created in the `test` namespace - adjust if needed |

## nodepoolconfigv3.yaml

Test case: `nodepoolconfigv3.yaml` - creates a node pool configuration with kubelet settings for a CCE node pool.

### APIs used

| Kind | Full API name |
|------|---------------|
| NodePoolConfigV3 | `nodepoolconfigv3s.cce.opentelekomcloud.m.crossplane.io` |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `clusterSelector` / `nodePoolSelector` | NodePoolConfigV3 `spec.forProvider` | Replace the `matchLabels` selectors with selectors pointing to your actual CCE cluster and node pool |
| `name: configuration` | NodePoolConfigV3 `spec.forProvider.name` | Replace with a unique name for your node pool configuration |
| `namespace: upbound-system` | NodePoolConfigV3 `metadata.namespace` | Change to the namespace where your CCE cluster and node pool resources live |

## Prerequisites

This example requires a configured `ClusterProviderConfig` - see [providerconfig](../providerconfig/README.md).
