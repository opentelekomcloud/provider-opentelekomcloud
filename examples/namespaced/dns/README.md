# DNS test cases

## zone.yaml

Test case: `zone.yaml` - creates a private DNS zone bound to a VPC, a public DNS
zone, and A/TXT record sets in the public zone.

### APIs used

| Kind | Full API name |
|------|---------------|
| VpcV1 | `vpcv1s.vpc.opentelekomcloud.m.crossplane.io` |
| ZoneV2 | `zonev2s.dns.opentelekomcloud.m.crossplane.io` |
| RecordsetV2 | `recordsetv2s.dns.opentelekomcloud.m.crossplane.io` |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${Rand.RFC1123Subdomain}` | metadata names of VpcV1, ZoneV2, RecordsetV2 | Replace with a unique RFC1123-compatible suffix (or leave for the example generator to substitute) |
| `name: crossplane.priv-example.com.` / `crossplane.pub-example.com.` | ZoneV2 `spec.forProvider.name` | Replace with DNS zone names you actually own/control (private zones can be arbitrary, public zones must be registered) |
| `email: crossplane@priv-example.com` / `crossplane@pub-example.com` | ZoneV2 `spec.forProvider.email` | Replace with a valid admin email for the zone |
| `routerRegion: eu-de` | private ZoneV2 `spec.forProvider.router` | Change to the region of the VPC the zone is bound to |
| `records: [10.0.0.1]` | A record set | Replace with the record value you want |
| `records: ["v=spf1 ..."]` | TXT record set | Replace with your SPF record content |
| `namespace: test` | all resources | All resources are created in the `test` namespace - adjust if needed |

## ptrrecord.yaml

Test case: `ptrrecord.yaml` - creates an EIP and a PTR record pointing at it.

### APIs used

| Kind | Full API name |
|------|---------------|
| EIPV1 | `eipv1s.vpc.opentelekomcloud.m.crossplane.io` |
| PtrrecordV2 | `ptrrecordv2s.dns.opentelekomcloud.m.crossplane.io` |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${Rand.RFC1123Subdomain}` | names of EIPV1, PtrrecordV2 | Replace with a unique RFC1123-compatible suffix (or leave for the example generator to substitute) |
| `name: ptr.crossplane.com.` | PtrrecordV2 `spec.forProvider.name` | Replace with the PTR domain name you want (must be in a domain you control) |
| `bandwidth: size: 8` | EIPV1 `spec.forProvider.bandwidth` | Adjust the bandwidth size if needed |
| `namespace: test` | all resources | All resources are created in the `test` namespace - adjust if needed |

## Prerequisites

This example requires a configured `ClusterProviderConfig` - see [providerconfig](../providerconfig/README.md).
