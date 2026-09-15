# SFS test cases

## filesystemv2.yaml

Test case: `filesystemv2.yaml` - creates a basic SFS NFS share.

### APIs used

| Kind | Full API name |
|------|---------------|
| FileSystemV2 | `filesystemv2s.sfs.opentelekomcloud.m.crossplane.io` |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${Rand.RFC1123Subdomain}` | name, `spec.forProvider.name` | Replace with a unique RFC1123-compatible suffix (or leave for the example generator to substitute) |
| `size: 50` | `spec.forProvider.size` | Adjust the share size (GB) |
| `namespace: test` | all resources | All resources are created in the `test` namespace - adjust if needed |

## shareaccessrulesv2.yaml

Test case: `shareaccessrulesv2.yaml` - creates an SFS share, a VPC, and a
cert-based access rule restricting access to that VPC.

### APIs used

| Kind | Full API name |
|------|---------------|
| ShareAccessRulesV2 | `shareaccessrulesv2s.sfs.opentelekomcloud.m.crossplane.io` |
| FileSystemV2 | `filesystemv2s.sfs.opentelekomcloud.m.crossplane.io` |
| VpcV1 | `vpcv1s.vpc.opentelekomcloud.m.crossplane.io` |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${Rand.RFC1123Subdomain}` | names of all resources | Replace with a unique RFC1123-compatible suffix (or leave for the example generator to substitute) |
| `size: 50` | FileSystemV2 `spec.forProvider.size` | Adjust the share size (GB) |
| `cidr: 192.168.0.0/16` | VpcV1 `spec.forProvider.cidr` | Adjust to a CIDR that does not conflict with existing networks |
| `namespace: test` | all resources | All resources are created in the `test` namespace - adjust if needed |

## turbosharev1.yaml

Test case: `turbosharev1.yaml` - creates an SFS Turbo (high-performance) NFS
share with a VPC, subnet, security group, and security group rule.

### APIs used

| Kind | Full API name |
|------|---------------|
| VpcV1 | `vpcv1s.vpc.opentelekomcloud.m.crossplane.io` |
| SubnetV1 | `subnetv1s.vpc.opentelekomcloud.m.crossplane.io` |
| SecgroupV3 | `secgroupv3s.vpc.opentelekomcloud.m.crossplane.io` |
| SecgroupRuleV3 | `secgrouprulev3s.vpc.opentelekomcloud.m.crossplane.io` |
| TurboShareV1 | `turbosharev1s.sfs.opentelekomcloud.m.crossplane.io` |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${Rand.RFC1123Subdomain}` | names of all resources | Replace with a unique RFC1123-compatible suffix (or leave for the example generator to substitute) |
| `availabilityZone: eu-de-01` | TurboShareV1 `spec.forProvider.availabilityZone` | Change to an AZ in your region |
| `size: 500` | TurboShareV1 `spec.forProvider.size` | Adjust the share size (GB) |
| `multiPort: "1-60000"` | SecgroupRuleV3 `spec.forProvider.multiPort` | Adjust the port range for the security rule |
| `namespace: test` | all resources | All resources are created in the `test` namespace - adjust if needed |

## Prerequisites

This example requires a configured `ClusterProviderConfig` - see [providerconfig](../providerconfig/README.md).
