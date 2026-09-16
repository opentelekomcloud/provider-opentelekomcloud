# LB test cases

## loadbalancer.yaml

Test case: `loadbalancer.yaml` - creates a VPC, subnet, EIP, and a load balancer
with a public IP attached.

### APIs used

| Kind | Full API name |
|------|---------------|
| VpcV1 | `vpcv1s.vpc.opentelekomcloud.m.crossplane.io` |
| SubnetV1 | `subnetv1s.vpc.opentelekomcloud.m.crossplane.io` |
| EIPV1 | `eipv1s.vpc.opentelekomcloud.m.crossplane.io` |
| LoadbalancerV3 | `loadbalancerv3s.lb.opentelekomcloud.m.crossplane.io` |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${Rand.RFC1123Subdomain}` | names of VpcV1, SubnetV1, EIPV1, LoadbalancerV3 | Replace with a unique RFC1123-compatible suffix (or leave for the example generator to substitute) |
| `availabilityZones: [eu-de-01]` | LoadbalancerV3 | Change to an AZ in your region |
| `bandwidth: size: 8` | EIPV1 `spec.forProvider.bandwidth` | Adjust the bandwidth size if needed |
| `namespace: test` | all resources | All resources are created in the `test` namespace - adjust if needed |

## listener_and_ipgroup.yaml

Test case: `listener_and_ipgroup.yaml` - creates a VPC, subnet, load balancer,
IP group, and a listener pointing at the IP group.

### APIs used

| Kind | Full API name |
|------|---------------|
| VpcV1 | `vpcv1s.vpc.opentelekomcloud.m.crossplane.io` |
| SubnetV1 | `subnetv1s.vpc.opentelekomcloud.m.crossplane.io` |
| LoadbalancerV3 | `loadbalancerv3s.lb.opentelekomcloud.m.crossplane.io` |
| IpgroupV3 | `ipgroupv3s.lb.opentelekomcloud.m.crossplane.io` |
| ListenerV3 | `listenerv3s.lb.opentelekomcloud.m.crossplane.io` |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${Rand.RFC1123Subdomain}` | names of all resources | Replace with a unique RFC1123-compatible suffix (or leave for the example generator to substitute) |
| `ip: "192.168.10.11"` | IpgroupV3 `spec.forProvider.ipList` | Replace with the actual backend IP address |
| `protocolPort: 8080` | ListenerV3 | Change to the port you want to listen on |
| `availabilityZones: [eu-de-01]` | LoadbalancerV3 | Change to an AZ in your region |
| `namespace: test` | all resources | All resources are created in the `test` namespace - adjust if needed |

## certificatev3.yaml

Test case: `certificatev3.yaml` - creates an LB certificate (for TLS listeners) with a private key from a Kubernetes secret.

### APIs used

| Kind | Full API name |
|------|---------------|
| CertificateV3 | `certificatev3s.lb.opentelekomcloud.m.crossplane.io` |
| Secret | `secrets` (core/v1) |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `certificate: ...` | CertificateV3 `spec.forProvider.certificate` | Replace the placeholder certificate with your actual certificate (PEM format) |
| `privateKeySecretRef` (secret `example-secret`, key `example-key`) | CertificateV3 `spec.forProvider.privateKeySecretRef` | The secret contains the base64-encoded private key - replace with your actual key (`base64 <<< '<your-private-key.pem>'`) |
| `domain: www.elb.com` | `spec.forProvider.domain` | Replace with your domain |
| `name: certificate_1` | `spec.forProvider.name` | Replace with a meaningful certificate name |
| `namespace: test` | all resources | All resources are created in the `test` namespace - adjust if needed |

## pool_and_monitor.yaml

Test case: `pool_and_monitor.yaml` - creates a VPC, subnet, load balancer,
member pool with a health monitor, and a pool member.

### APIs used

| Kind | Full API name |
|------|---------------|
| VpcV1 | `vpcv1s.vpc.opentelekomcloud.m.crossplane.io` |
| SubnetV1 | `subnetv1s.vpc.opentelekomcloud.m.crossplane.io` |
| LoadbalancerV3 | `loadbalancerv3s.lb.opentelekomcloud.m.crossplane.io` |
| PoolV3 | `poolv3s.lb.opentelekomcloud.m.crossplane.io` |
| MonitorV3 | `monitorv3s.lb.opentelekomcloud.m.crossplane.io` |
| MemberV3 | `memberv3s.lb.opentelekomcloud.m.crossplane.io` |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${Rand.RFC1123Subdomain}` | names of all resources | Replace with a unique RFC1123-compatible suffix (or leave for the example generator to substitute) |
| `address: "192.168.0.100"` | MemberV3 `spec.forProvider.address` | Replace with the actual backend server IP |
| `protocolPort: 8080` | MemberV3 | Change to the backend service port |
| `monitorPort: 8080` | MonitorV3 | Change to the health-check port |
| `type: HTTP` | MonitorV3 | Change to `TCP` if you prefer TCP health checks |
| `availabilityZones: [eu-de-01]` | LoadbalancerV3 | Change to an AZ in your region |
| `namespace: test` | all resources | All resources are created in the `test` namespace - adjust if needed |

## policy_and_rule.yaml

Test case: `policy_and_rule.yaml` - creates a VPC, subnet, load balancer,
listener, pool, L7 policy with a redirect-to-pool action, and a regex rule.

### APIs used

| Kind | Full API name |
|------|---------------|
| VpcV1 | `vpcv1s.vpc.opentelekomcloud.m.crossplane.io` |
| SubnetV1 | `subnetv1s.vpc.opentelekomcloud.m.crossplane.io` |
| LoadbalancerV3 | `loadbalancerv3s.lb.opentelekomcloud.m.crossplane.io` |
| ListenerV3 | `listenerv3s.lb.opentelekomcloud.m.crossplane.io` |
| PoolV3 | `poolv3s.lb.opentelekomcloud.m.crossplane.io` |
| PolicyV3 | `policyv3s.lb.opentelekomcloud.m.crossplane.io` |
| RuleV3 | `rulev3s.lb.opentelekomcloud.m.crossplane.io` |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${Rand.RFC1123Subdomain}` | names of VpcV1, SubnetV1, LoadbalancerV3, ListenerV3, PoolV3 | Replace with a unique RFC1123-compatible suffix (or leave for the example generator to substitute) |
| `value: ^.+$` | RuleV3 `spec.forProvider.value` | Replace with the actual regex pattern for your routing rule |
| `position: 37` | PolicyV3 `spec.forProvider.position` | Adjust the policy evaluation order if needed |
| `availabilityZones: [eu-de-01]` | LoadbalancerV3 | Change to an AZ in your region |
| `namespace: test` | all resources | All resources are created in the `test` namespace - adjust if needed |

## Prerequisites

This example requires a configured `ClusterProviderConfig` - see [providerconfig](../providerconfig/README.md).
