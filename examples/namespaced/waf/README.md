# WAF test cases

The dedicated WAF API is project-scoped. The domain and rule examples assume a
dedicated WAF instance is available in the target project; see
[`instancev1.yaml`](#instancev1yaml) for provisioning one.

> [!IMPORTANT]
> WAF resources require you to configure the REGION.

## instancev1.yaml

Test case: `instancev1.yaml` - creates a VPC, subnet, security group, and a
dedicated WAF instance.

### APIs used

| Kind | Full API name |
|------|---------------|
| VpcV1 | `vpcv1s.vpc.opentelekomcloud.m.crossplane.io` |
| SubnetV1 | `subnetv1s.vpc.opentelekomcloud.m.crossplane.io` |
| SecgroupV3 | `secgroupv3s.vpc.opentelekomcloud.m.crossplane.io` |
| InstanceV1 | `instancev1s.waf.opentelekomcloud.m.crossplane.io` |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${Rand.RFC1123Subdomain}` | names of all resources and cloud names | Replace with a unique RFC1123-compatible suffix (or leave for the example generator to substitute) |
| `availabilityZone: eu-de-01`, `flavor: s2.large.2`, `specification: waf.instance.professional` | InstanceV1 `spec.forProvider` | Adjust to the AZ, flavor, and edition available in your region |
| `cidr: 192.168.0.0/16` / `192.168.0.0/24` | VpcV1 / SubnetV1 | Adjust to CIDRs that do not conflict with existing networks |
| `namespace: test` | all resources | All resources are created in the `test` namespace - adjust if needed |

Creating a dedicated WAF instance can take up to 30 minutes and deletion up to
20 minutes.

## policyv1.yaml

Test case: `policyv1.yaml` - creates a standalone WAF protection policy.

### APIs used

| Kind | Full API name |
|------|---------------|
| PolicyV1 | `policyv1s.waf.opentelekomcloud.m.crossplane.io` |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${Rand.RFC1123Subdomain}` | metadata name, `spec.forProvider.name` | Replace with a unique RFC1123-compatible suffix (or leave for the example generator to substitute) |
| `protectionMode: block`, `level: 2`, `options` | PolicyV1 `spec.forProvider` | Adjust the protection mode, level, and enabled rule groups |
| `namespace: test` | all resources | All resources are created in the `test` namespace - adjust if needed |

## certificatev1.yaml

Test case: `certificatev1.yaml` - creates a WAF dedicated certificate from a
self-signed test certificate and private key supplied by the datasource values.

### APIs used

| Kind | Full API name |
|------|---------------|
| CertificateV1 | `certificatev1s.waf.opentelekomcloud.m.crossplane.io` |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${Rand.RFC1123Subdomain}` | metadata name, `spec.forProvider.name` | Replace with a unique RFC1123-compatible suffix (or leave for the example generator to substitute) |
| `${data.wafdCertificate}` | CertificateV1 `spec.forProvider.content` | PEM certificate from the datasource values (self-signed test certificate valid until 2036-11-15); replace with your own certificate |
| `${data.wafdPrivateKey}` | CertificateV1 `spec.forProvider.key` | Private key matching `${data.wafdCertificate}` from the datasource values; replace together with the certificate |
| `namespace: test` | all resources | All resources are created in the `test` namespace - adjust if needed |

## referencetablev1.yaml

Test case: `referencetablev1.yaml` - creates a WAF reference table of URLs.

### APIs used

| Kind | Full API name |
|------|---------------|
| ReferenceTableV1 | `referencetablev1s.waf.opentelekomcloud.m.crossplane.io` |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${Rand.RFC1123Subdomain}` | metadata name, `spec.forProvider.name` | Replace with a unique RFC1123-compatible suffix (or leave for the example generator to substitute) |
| `type: url`, `conditions` | ReferenceTableV1 `spec.forProvider` | Adjust the table type and entries; the type must match the category of the rule condition that uses the table |
| `namespace: test` | all resources | All resources are created in the `test` namespace - adjust if needed |

## domainv1.yaml

Test case: `domainv1.yaml` - creates a VPC, subnet, policy, and an HTTP domain
protected by the dedicated WAF.

### APIs used

| Kind | Full API name |
|------|---------------|
| VpcV1 | `vpcv1s.vpc.opentelekomcloud.m.crossplane.io` |
| SubnetV1 | `subnetv1s.vpc.opentelekomcloud.m.crossplane.io` |
| PolicyV1 | `policyv1s.waf.opentelekomcloud.m.crossplane.io` |
| DomainV1 | `domainv1s.waf.opentelekomcloud.m.crossplane.io` |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${Rand.RFC1123Subdomain}` | names of all resources and cloud names | Replace with a unique RFC1123-compatible suffix (or leave for the example generator to substitute) |
| `domain: www.wafd-...example.com` | DomainV1 `spec.forProvider.domain` | Replace with the domain name or IP address to protect |
| `address: 192.168.0.10`, `port: 8080` | DomainV1 `spec.forProvider.server` | Point at a real backend reachable from the selected VPC |
| `keepPolicy: true` | DomainV1 `spec.forProvider.keepPolicy` | Keep `true` so deleting the domain does not delete the referenced policy |
| `namespace: test` | all resources | All resources are created in the `test` namespace - adjust if needed |

## ccrulev1.yaml

Test case: `ccrulev1.yaml` - creates a policy, a URL reference table, and a CC
protection rule that uses the table in its condition.

### APIs used

| Kind | Full API name |
|------|---------------|
| PolicyV1 | `policyv1s.waf.opentelekomcloud.m.crossplane.io` |
| ReferenceTableV1 | `referencetablev1s.waf.opentelekomcloud.m.crossplane.io` |
| CcRuleV1 | `ccrulev1s.waf.opentelekomcloud.m.crossplane.io` |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${Rand.RFC1123Subdomain}` | names of all resources and cloud names | Replace with a unique RFC1123-compatible suffix (or leave for the example generator to substitute) |
| `url: /abc1`, `limitNum`, `limitPeriod`, `lockTime` | CcRuleV1 `spec.forProvider` | Adjust the protected URL and rate-limiting thresholds |
| `logicOperation: contain_any` + reference table | CcRuleV1 `spec.forProvider.conditions` | Keep the reference table type (`url`) in sync with the condition category |
| `namespace: test` | all resources | All resources are created in the `test` namespace - adjust if needed |

## anticrawlerrulev1.yaml

Test case: `anticrawlerrulev1.yaml` - creates a policy and an anti-crawler rule.

### APIs used

| Kind | Full API name |
|------|---------------|
| PolicyV1 | `policyv1s.waf.opentelekomcloud.m.crossplane.io` |
| AntiCrawlerRuleV1 | `anticrawlerrulev1s.waf.opentelekomcloud.m.crossplane.io` |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${Rand.RFC1123Subdomain}` | names of all resources and cloud names | Replace with a unique RFC1123-compatible suffix (or leave for the example generator to substitute) |
| `url: /patent/id`, `logic: 3`, `protectionMode` | AntiCrawlerRuleV1 `spec.forProvider` | Adjust the protected URL and detection settings |
| `namespace: test` | all resources | All resources are created in the `test` namespace - adjust if needed |

## datamaskingrulev1.yaml

Test case: `datamaskingrulev1.yaml` - creates a policy and a data masking rule.

### APIs used

| Kind | Full API name |
|------|---------------|
| PolicyV1 | `policyv1s.waf.opentelekomcloud.m.crossplane.io` |
| DataMaskingRuleV1 | `datamaskingrulev1s.waf.opentelekomcloud.m.crossplane.io` |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${Rand.RFC1123Subdomain}` | names of all resources and cloud names | Replace with a unique RFC1123-compatible suffix (or leave for the example generator to substitute) |
| `category: params`, `url: /login` | DataMaskingRuleV1 `spec.forProvider` | Adjust the masked field location and protected URL |
| `namespace: test` | all resources | All resources are created in the `test` namespace - adjust if needed |

## knownattacksourcerulev1.yaml

Test case: `knownattacksourcerulev1.yaml` - creates a policy and a known attack
source rule.

### APIs used

| Kind | Full API name |
|------|---------------|
| PolicyV1 | `policyv1s.waf.opentelekomcloud.m.crossplane.io` |
| KnownAttackSourceRuleV1 | `knownattacksourcerulev1s.waf.opentelekomcloud.m.crossplane.io` |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${Rand.RFC1123Subdomain}` | names of all resources and cloud names | Replace with a unique RFC1123-compatible suffix (or leave for the example generator to substitute) |
| `category: long_cookie_block`, `blockTime: 300` | KnownAttackSourceRuleV1 `spec.forProvider` | Adjust the block category and duration |
| `namespace: test` | all resources | All resources are created in the `test` namespace - adjust if needed |

## webtamperrulev1.yaml

Test case: `webtamperrulev1.yaml` - creates a VPC, subnet, policy, domain, and a
web tamper protection rule bound to the domain.

### APIs used

| Kind | Full API name |
|------|---------------|
| VpcV1 | `vpcv1s.vpc.opentelekomcloud.m.crossplane.io` |
| SubnetV1 | `subnetv1s.vpc.opentelekomcloud.m.crossplane.io` |
| PolicyV1 | `policyv1s.waf.opentelekomcloud.m.crossplane.io` |
| DomainV1 | `domainv1s.waf.opentelekomcloud.m.crossplane.io` |
| WebTamperRuleV1 | `webtamperrulev1s.waf.opentelekomcloud.m.crossplane.io` |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${Rand.RFC1123Subdomain}` | names of all resources and cloud names | Replace with a unique RFC1123-compatible suffix (or leave for the example generator to substitute) |
| `domain`, `address`, `port` | DomainV1 `spec.forProvider` | Replace with the domain and backend to protect |
| `url: /login` | WebTamperRuleV1 `spec.forProvider.url` | Adjust the URL protected against tampering |
| `namespace: test` | all resources | All resources are created in the `test` namespace - adjust if needed |

## antileakagerulev1.yaml

Test case: `antileakagerulev1.yaml` - creates a policy and an information leakage
protection rule.

### APIs used

| Kind | Full API name |
|------|---------------|
| PolicyV1 | `policyv1s.waf.opentelekomcloud.m.crossplane.io` |
| AntiLeakageRuleV1 | `antileakagerulev1s.waf.opentelekomcloud.m.crossplane.io` |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${Rand.RFC1123Subdomain}` | names of all resources and cloud names | Replace with a unique RFC1123-compatible suffix (or leave for the example generator to substitute) |
| `category: sensitive`, `contents: [id_card]`, `action: block` | AntiLeakageRuleV1 `spec.forProvider` | Adjust the sensitive data type, contents, and action |
| `namespace: test` | all resources | All resources are created in the `test` namespace - adjust if needed |

## alarmmaskingrulev1.yaml

Test case: `alarmmaskingrulev1.yaml` - creates a VPC, subnet, policy, domain, and
an alarm masking rule for the domain.

### APIs used

| Kind | Full API name |
|------|---------------|
| VpcV1 | `vpcv1s.vpc.opentelekomcloud.m.crossplane.io` |
| SubnetV1 | `subnetv1s.vpc.opentelekomcloud.m.crossplane.io` |
| PolicyV1 | `policyv1s.waf.opentelekomcloud.m.crossplane.io` |
| DomainV1 | `domainv1s.waf.opentelekomcloud.m.crossplane.io` |
| AlarmMaskingRuleV1 | `alarmmaskingrulev1s.waf.opentelekomcloud.m.crossplane.io` |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${Rand.RFC1123Subdomain}` | names of all resources and cloud names | Replace with a unique RFC1123-compatible suffix (or leave for the example generator to substitute) |
| `domain`, `address`, `port` | DomainV1 `spec.forProvider` | Replace with the domain and backend to protect |
| `rule: xss`, `conditions` | AlarmMaskingRuleV1 `spec.forProvider` | Adjust the masked rule type and matching conditions (`xss`, `webshell`, or `vuln`) |
| `advancedSettings: [{index: header, contents: [content-type]}]` | AlarmMaskingRuleV1 `spec.forProvider.advancedSettings` | Keep a valid `index`/`contents` pair: the WAFD API rejects an empty `advanced` object, which the provider sends when this block is omitted |
| `namespace: test` | all resources | All resources are created in the `test` namespace - adjust if needed |

## geoiprulev1.yaml

Test case: `geoiprulev1.yaml` - creates a policy and a geo IP access control rule.

### APIs used

| Kind | Full API name |
|------|---------------|
| PolicyV1 | `policyv1s.waf.opentelekomcloud.m.crossplane.io` |
| GeoIPRuleV1 | `geoiprulev1s.waf.opentelekomcloud.m.crossplane.io` |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${Rand.RFC1123Subdomain}` | names of all resources and cloud names | Replace with a unique RFC1123-compatible suffix (or leave for the example generator to substitute) |
| `regionCode: BR`, `action: 0` | GeoIPRuleV1 `spec.forProvider` | Change the geographic region and action |
| `namespace: test` | all resources | All resources are created in the `test` namespace - adjust if needed |

## blacklistrulev1.yaml

Test case: `blacklistrulev1.yaml` - creates a policy, a known attack source rule,
and an IP blacklist rule that references the source rule.

### APIs used

| Kind | Full API name |
|------|---------------|
| PolicyV1 | `policyv1s.waf.opentelekomcloud.m.crossplane.io` |
| KnownAttackSourceRuleV1 | `knownattacksourcerulev1s.waf.opentelekomcloud.m.crossplane.io` |
| BlacklistRuleV1 | `blacklistrulev1s.waf.opentelekomcloud.m.crossplane.io` |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${Rand.RFC1123Subdomain}` | names of all resources and cloud names | Replace with a unique RFC1123-compatible suffix (or leave for the example generator to substitute) |
| `ipAddress: 192.168.1.0/24`, `action: 0` | BlacklistRuleV1 `spec.forProvider` | Adjust the blocked address range and action |
| `namespace: test` | all resources | All resources are created in the `test` namespace - adjust if needed |

## preciseprotectionrulev1.yaml

Test case: `preciseprotectionrulev1.yaml` - creates a policy, a reference table, a
known attack source rule, and a precise protection rule that references both.

### APIs used

| Kind | Full API name |
|------|---------------|
| PolicyV1 | `policyv1s.waf.opentelekomcloud.m.crossplane.io` |
| ReferenceTableV1 | `referencetablev1s.waf.opentelekomcloud.m.crossplane.io` |
| KnownAttackSourceRuleV1 | `knownattacksourcerulev1s.waf.opentelekomcloud.m.crossplane.io` |
| PreciseProtectionRuleV1 | `preciseprotectionrulev1s.waf.opentelekomcloud.m.crossplane.io` |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${Rand.RFC1123Subdomain}` | names of all resources and cloud names | Replace with a unique RFC1123-compatible suffix (or leave for the example generator to substitute) |
| `priority: 50`, `conditions`, `action` | PreciseProtectionRuleV1 `spec.forProvider` | Adjust the rule priority, matching conditions, and action |
| `namespace: test` | all resources | All resources are created in the `test` namespace - adjust if needed |

## Prerequisites

This example requires a configured `ClusterProviderConfig` - see [providerconfig](../providerconfig/README.md).
