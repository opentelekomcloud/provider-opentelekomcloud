# RDS test cases

## instance.yaml

Test case: `instance.yaml` - creates a PostgreSQL RDS instance with a VPC,
subnet, security group, EIP, password from a Kubernetes secret, and a
Postgres extension.

### APIs used

| Kind | Full API name |
|------|---------------|
| Secret | `secrets` (core/v1) |
| VpcV1 | `vpcv1s.vpc.opentelekomcloud.m.crossplane.io` |
| SubnetV1 | `subnetv1s.vpc.opentelekomcloud.m.crossplane.io` |
| SecgroupV2 | `secgroupv2s.compute.opentelekomcloud.m.crossplane.io` |
| EIPV1 | `eipv1s.vpc.opentelekomcloud.m.crossplane.io` |
| InstanceV3 | `instancev3s.rds.opentelekomcloud.m.crossplane.io` |
| PostgresExtensionV3 | `postgresextensionv3s.rds.opentelekomcloud.m.crossplane.io` |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${Rand.RFC1123Subdomain}` | names of VpcV1, SubnetV1, SecgroupV2, EIPV1, InstanceV3, PostgresExtensionV3 | Replace with a unique RFC1123-compatible suffix (or leave for the example generator to substitute) |
| `${data.rdsKey}` | Secret `example-rds-secret`, key `example-rds-key` | Replace with the base64-encoded RDS database password (e.g. `base64 <<< '<password>'`) |
| `flavor: rds.pg.n1.large.2` | InstanceV3 `spec.forProvider.flavor` | Adjust to an RDS flavor available in your region |
| `type: PostgreSQL` / `version: "15"` | InstanceV3 `spec.forProvider.db` | Change to the database type/version you need |
| `port: 8635` | InstanceV3 `spec.forProvider.db` | Change to the port you need (must match the security group rule) |
| `availabilityZone: [eu-de-03]` | InstanceV3 | Change to an AZ in your region |
| `volume: size: 100` / `type: CLOUDSSD` | InstanceV3 `spec.forProvider.volume` | Adjust storage size and type |
| `security group rules` (ports 8635/8080/443 from 0.0.0.0/0) | SecgroupV2 `spec.forProvider.rule` | Restrict the CIDR and ports to what you actually need |
| `bandwidth: size: 8` | EIPV1 `spec.forProvider.bandwidth` | Adjust the bandwidth size if needed |
| `extensionName: hstore` / `databaseName: postgres` | PostgresExtensionV3 `spec.forProvider` | Change to the extension and database you need |
| `namespace: test` | all resources | All resources are created in the `test` namespace - adjust if needed |

## Prerequisites

This example requires a configured `ClusterProviderConfig` - see [providerconfig](../providerconfig/README.md).
