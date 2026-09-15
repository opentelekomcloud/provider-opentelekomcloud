# DDS test cases

## instance.yaml

Test case: `instance.yaml` - creates a DDS replica set instance with a
VPC, subnet, security group, password secret, and a backup.

### APIs used

| Kind | Full API name |
|------|---------------|
| Secret | `secrets` (core/v1) |
| VpcV1 | `vpcv1s.vpc.opentelekomcloud.m.crossplane.io` |
| SubnetV1 | `subnetv1s.vpc.opentelekomcloud.m.crossplane.io` |
| SecgroupV2 | `secgroupv2s.compute.opentelekomcloud.m.crossplane.io` |
| InstanceV3 | `instancev3s.dds.opentelekomcloud.m.crossplane.io` |
| BackupV3 | `backupv3s.dds.opentelekomcloud.m.crossplane.io` |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${Rand.RFC1123Subdomain}` | names of VpcV1, SubnetV1, SecgroupV2, InstanceV3, BackupV3 | Replace with a unique RFC1123-compatible suffix (or leave for the example generator to substitute) |
| `${data.ddsKey}` | Secret `example-dds-password`, key `example-dds-key` | Replace with the base64-encoded DDS instance password (e.g. `base64 <<< '<password>'`) |
| `specCode: dds.mongodb.s2.medium.4.repset` | InstanceV3 `spec.forProvider.flavor` | Adjust to a DDS spec available in your region |
| `version: "4.0"` / `type: DDS-Community` | InstanceV3 `spec.forProvider.datastore` | Change to the MongoDB version/type you need |
| `availabilityZone: eu-de-01` | InstanceV3 | Change to an AZ in your region |
| `security group rules` (ports 22/8080/27017 from 0.0.0.0/0) | SecgroupV2 `spec.forProvider.rule` | Restrict the CIDR and ports to what you actually need |
| `namespace: test` | all resources | All resources are created in the `test` namespace - adjust if needed |

## Prerequisites

This example requires a configured `ClusterProviderConfig` - see [providerconfig](../providerconfig/README.md).
