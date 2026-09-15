# VPC test cases

## vpc.yaml

Test case: `vpc.yaml` - creates a VPC and a secondary CIDR for it.

### APIs used

| Kind | Full API name |
|------|---------------|
| VpcV1 | `vpcv1s.vpc.opentelekomcloud.m.crossplane.io` |
| SecondaryCidrV3 | `secondarycidrv3s.vpc.opentelekomcloud.m.crossplane.io` |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${Rand.RFC1123Subdomain}` | names of VpcV1, SecondaryCidrV3, VpcV1 `spec.forProvider.name` | Replace with a unique RFC1123-compatible suffix (or leave for the example generator to substitute) |
| `cidr: 192.168.0.0/16` | VpcV1 `spec.forProvider.cidr` | Adjust to a CIDR that does not conflict with existing networks |
| `cidrs: [23.9.0.0/16, ...]` | SecondaryCidrV3 `spec.forProvider.cidrs` | Replace with the secondary CIDR blocks you need |
| `namespace: test` | all resources | All resources are created in the `test` namespace - adjust if needed |

## secgrouprulev3.yaml

Test case: `secgrouprulev3.yaml` - creates a VPC security group and an ingress
rule allowing TCP port 8080 from 10.10.0.0/16.

### APIs used

| Kind | Full API name |
|------|---------------|
| SecgroupV3 | `secgroupv3s.vpc.opentelekomcloud.m.crossplane.io` |
| SecgroupRuleV3 | `secgrouprulev3s.vpc.opentelekomcloud.m.crossplane.io` |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${Rand.RFC1123Subdomain}` | names of SecgroupV3, SecgroupRuleV3, SecgroupV3 `spec.forProvider.name` | Replace with a unique RFC1123-compatible suffix (or leave for the example generator to substitute) |
| `remoteIpPrefix: 10.10.0.0/16` | SecgroupRuleV3 `spec.forProvider.remoteIpPrefix` | Replace with the source CIDR you want to allow |
| `multiPort: "8080"` | SecgroupRuleV3 `spec.forProvider.multiPort` | Change to the port(s) you want to open |
| `namespace: test` | all resources | All resources are created in the `test` namespace - adjust if needed |

## ipaddressgroupv3.yaml

Test case: `ipaddressgroupv3.yaml` - creates an IP address group (used by
security group rules to reference a set of IPs).

### APIs used

| Kind | Full API name |
|------|---------------|
| IPAddressGroupV3 | `ipaddressgroupv3s.vpc.opentelekomcloud.m.crossplane.io` |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${Rand.RFC1123Subdomain}` | metadata name | Replace with a unique RFC1123-compatible suffix (or leave for the example generator to substitute) |
| `ipVersion: 4` | `spec.forProvider.ipVersion` | Change to `6` if you need an IPv6 address group |
| `namespace: test` | all resources | All resources are created in the `test` namespace - adjust if needed |
