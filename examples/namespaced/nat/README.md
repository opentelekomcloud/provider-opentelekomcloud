# NAT test cases

## natgw.yaml

Test case: `natgw.yaml` - creates a VPC, subnet, NAT gateway, two EIPs, a DNAT
rule, and a SNAT rule.

### APIs used

| Kind | Full API name |
|------|---------------|
| VpcV1 | `vpcv1s.vpc.opentelekomcloud.m.crossplane.io` |
| SubnetV1 | `subnetv1s.vpc.opentelekomcloud.m.crossplane.io` |
| GatewayV2 | `gatewayv2s.nat.opentelekomcloud.m.crossplane.io` |
| EIPV1 | `eipv1s.vpc.opentelekomcloud.m.crossplane.io` |
| DnatRuleV2 | `dnatrulev2s.nat.opentelekomcloud.m.crossplane.io` |
| SnatRuleV2 | `snatrulev2s.nat.opentelekomcloud.m.crossplane.io` |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${Rand.RFC1123Subdomain}` | names of all resources | Replace with a unique RFC1123-compatible suffix (or leave for the example generator to substitute) |
| `privateIp: "192.168.100.10"` | DnatRuleV2 `spec.forProvider.privateIp` | Replace with the internal IP of the server you want to expose |
| `externalServicePort: 80` / `internalServicePort: 80` | DnatRuleV2 | Change to the ports you need |
| `cidr: 192.168.0.0/24` | SnatRuleV2 `spec.forProvider.cidr` | Replace with the subnet CIDR that needs outbound access |
| `bandwidth: size: 8` | EIPV1 (x2) `spec.forProvider.bandwidth` | Adjust the bandwidth size if needed |
| `namespace: test` | all resources | All resources are created in the `test` namespace - adjust if needed |
