# DCS test cases

## instance.yaml

Test case: `instance.yaml` - creates a DCS (Redis) instance in a VPC/subnet with
a password from a Kubernetes secret.

### APIs used

| Kind | Full API name |
|------|---------------|
| Secret | `secrets` (core/v1) |
| VpcV1 | `vpcv1s.vpc.opentelekomcloud.m.crossplane.io` |
| SubnetV1 | `subnetv1s.vpc.opentelekomcloud.m.crossplane.io` |
| InstanceV2 | `instancev2s.dcs.opentelekomcloud.m.crossplane.io` |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${Rand.RFC1123Subdomain}` | names of VpcV1, SubnetV1, InstanceV2 | Replace with a unique RFC1123-compatible suffix (or leave for the example generator to substitute) |
| `${data.dcsKey}` | Secret `example-dcs-password`, key `example-dcs-key` | Replace with the base64-encoded DCS instance password (e.g. `base64 <<< '<password>'`); the secret lives in the `crossplane-system` namespace - move it to `test` or keep it there intentionally |
| `flavor: redis.ha.xu1.tiny.r2.128` | InstanceV2 `spec.forProvider.flavor` | Adjust to a DCS flavor available in your region |
| `engine: Redis` / `engineVersion: "5.0"` | InstanceV2 | Change to the engine/version you need |
| `availabilityZones: [eu-de-01]` | InstanceV2 | Change to an AZ in your region |
| `capacity: 0.125` | InstanceV2 | Adjust storage capacity (GB) |
| `namespace: test` | VpcV1, SubnetV1, InstanceV2 | All resources are created in the `test` namespace - adjust if needed |
