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

## routetableroutev1.yaml

Test case: `routetableroutev1.yaml` - creates a VPC and a route table and adds a
peering route to the route table.

### APIs used

| Kind | Full API name |
|------|---------------|
| VpcV1 | `vpcv1s.vpc.opentelekomcloud.m.crossplane.io` |
| RouteTableV1 | `routetablev1s.vpc.opentelekomcloud.m.crossplane.io` |
| RouteTableRouteV1 | `routetableroutev1s.vpc.opentelekomcloud.m.crossplane.io` |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${Rand.RFC1123Subdomain}` | names of VpcV1, RouteTableV1, RouteTableRouteV1 | Replace with a unique RFC1123-compatible suffix (or leave for the example generator to substitute) |
| `${data.peeringConnectionId}` | RouteTableRouteV1 `spec.forProvider.nexthop` | Replace with the ID of a peering connection whose VPC matches the VPC created by this example |
| `destination: 172.16.0.0/16` | RouteTableRouteV1 `spec.forProvider.destination` | Change to the peer VPC CIDR of your peering connection |
| `type: peering` | RouteTableRouteV1 `spec.forProvider.type` | Change to the route type you need (`ecs`, `eni`, `vip`, `nat`, `vpn`, `dc`, `egw`, `er`, `subeni`) |
| `cidr: 192.168.0.0/16` | VpcV1 `spec.forProvider.cidr` | Adjust to a CIDR that does not conflict with existing networks |
| `namespace: test` | all resources | All resources are created in the `test` namespace - adjust if needed |

## routetablesubnetassociatev1.yaml

Test case: `routetablesubnetassociatev1.yaml` - creates a VPC, a subnet and a
route table and associates the subnet with the route table.

### APIs used

| Kind | Full API name |
|------|---------------|
| VpcV1 | `vpcv1s.vpc.opentelekomcloud.m.crossplane.io` |
| SubnetV1 | `subnetv1s.vpc.opentelekomcloud.m.crossplane.io` |
| RouteTableV1 | `routetablev1s.vpc.opentelekomcloud.m.crossplane.io` |
| RouteTableSubnetAssociateV1 | `routetablesubnetassociatev1s.vpc.opentelekomcloud.m.crossplane.io` |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${Rand.RFC1123Subdomain}` | names of VpcV1, SubnetV1, RouteTableV1, RouteTableSubnetAssociateV1 | Replace with a unique RFC1123-compatible suffix (or leave for the example generator to substitute) |
| `cidr: 192.168.0.0/16` | VpcV1 `spec.forProvider.cidr` | Adjust to a CIDR that does not conflict with existing networks |
| `cidr: 192.168.0.0/24` / `gatewayIp: 192.168.0.1` | SubnetV1 `spec.forProvider` | Adjust to a subnet CIDR and gateway inside the VPC CIDR |
| `namespace: test` | all resources | All resources are created in the `test` namespace - adjust if needed |

## subnetprivateipv1.yaml

Test case: `subnetprivateipv1.yaml` - creates a VPC and a subnet and assigns a
private IP address in the subnet.

### APIs used

| Kind | Full API name |
|------|---------------|
| VpcV1 | `vpcv1s.vpc.opentelekomcloud.m.crossplane.io` |
| SubnetV1 | `subnetv1s.vpc.opentelekomcloud.m.crossplane.io` |
| SubnetPrivateIPV1 | `subnetprivateipv1s.vpc.opentelekomcloud.m.crossplane.io` |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${Rand.RFC1123Subdomain}` | names of VpcV1, SubnetV1, SubnetPrivateIPV1 | Replace with a unique RFC1123-compatible suffix (or leave for the example generator to substitute) |
| `ipAddress: 192.168.0.10` | SubnetPrivateIPV1 `spec.forProvider.ipAddress` | Replace with a free IPv4 address in the subnet, or remove it to let the platform assign one |
| `cidr: 192.168.0.0/16` | VpcV1 `spec.forProvider.cidr` | Adjust to a CIDR that does not conflict with existing networks |
| `cidr: 192.168.0.0/24` / `gatewayIp: 192.168.0.1` | SubnetV1 `spec.forProvider` | Adjust to a subnet CIDR and gateway inside the VPC CIDR |
| `namespace: test` | all resources | All resources are created in the `test` namespace - adjust if needed |

## Prerequisites

This example requires a configured `ClusterProviderConfig` - see [providerconfig](../providerconfig/README.md).
