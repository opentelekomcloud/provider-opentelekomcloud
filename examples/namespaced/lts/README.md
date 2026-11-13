# LTS test cases

## cceaccessv3.yaml

Test case: `cceaccessv3.yaml` - creates an LTS log group and stream and a CCE
access configuration that collects container stdout logs. The CCE cluster must
already exist; the cluster ID cannot be selected by labels because a reference
from the LTS API package to the CCE API package would create an import cycle.

### APIs used

| Kind | Full API name |
|------|---------------|
| GroupV2 | `groupv2s.lts.opentelekomcloud.m.crossplane.io` |
| StreamV2 | `streamv2s.lts.opentelekomcloud.m.crossplane.io` |
| CceAccessV3 | `cceaccessv3s.lts.opentelekomcloud.m.crossplane.io` |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${Rand.RFC1123Subdomain}` | names of all resources and cloud names | Replace with a unique RFC1123-compatible suffix (or leave for the example generator to substitute) |
| `REPLACE_WITH_CCE_CLUSTER_ID` | CceAccessV3 `spec.forProvider.clusterId` | Replace with the exact ID of an existing CCE cluster. Selectors are not supported for this field |
| `namespace: test` | all resources | All resources are created in the `test` namespace - adjust if needed |

## crossaccountaccessv2.yaml

Test case: `crossaccountaccessv2.yaml` - creates a local LTS log group and stream
and a cross-account access rule. The delegator account details must be supplied
manually; this example cannot be fully self-contained.

### APIs used

| Kind | Full API name |
|------|---------------|
| GroupV2 | `groupv2s.lts.opentelekomcloud.m.crossplane.io` |
| StreamV2 | `streamv2s.lts.opentelekomcloud.m.crossplane.io` |
| CrossAccountAccessV2 | `crossaccountaccessv2s.lts.opentelekomcloud.m.crossplane.io` |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${Rand.RFC1123Subdomain}` | names of all resources and cloud names | Replace with a unique RFC1123-compatible suffix (or leave for the example generator to substitute) |
| `REPLACE_WITH_AGENCY_DOMAIN_NAME` | CrossAccountAccessV2 `spec.forProvider.agencyDomainName` | Replace with the delegator account domain name |
| `REPLACE_WITH_AGENCY_NAME` | CrossAccountAccessV2 `spec.forProvider.agencyName` | Replace with the IAM agency name created by the delegator |
| `REPLACE_WITH_AGENCY_PROJECT_ID` | CrossAccountAccessV2 `spec.forProvider.agencyProjectId` | Replace with the delegator project ID |
| `REPLACE_WITH_AGENCY_GROUP_ID` / `_NAME` | CrossAccountAccessV2 `spec.forProvider.logAgencyGroupId` / `logAgencyGroupName` | Replace with the log group ID and name in the delegated account |
| `REPLACE_WITH_AGENCY_STREAM_ID` / `_NAME` | CrossAccountAccessV2 `spec.forProvider.logAgencyStreamId` / `logAgencyStreamName` | Replace with the log stream ID and name in the delegated account |
| `namespace: test` | all resources | All resources are created in the `test` namespace - adjust if needed |

## groupv2.yaml

Test case: `groupv2.yaml` - creates a standalone LTS log group.

### APIs used

| Kind | Full API name |
|------|---------------|
| GroupV2 | `groupv2s.lts.opentelekomcloud.m.crossplane.io` |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${Rand.RFC1123Subdomain}` | metadata name, `spec.forProvider.groupName` | Replace with a unique RFC1123-compatible suffix (or leave for the example generator to substitute) |
| `ttlInDays: 7` | GroupV2 `spec.forProvider.ttlInDays` | Adjust the log retention period (1-365 days) |
| `namespace: test` | all resources | All resources are created in the `test` namespace - adjust if needed |

## hostaccessv3.yaml

Test case: `hostaccessv3.yaml` - creates an LTS log group, log stream, host group,
and a host access configuration collecting `/var/log/*`.

### APIs used

| Kind | Full API name |
|------|---------------|
| GroupV2 | `groupv2s.lts.opentelekomcloud.m.crossplane.io` |
| StreamV2 | `streamv2s.lts.opentelekomcloud.m.crossplane.io` |
| HostGroupV3 | `hostgroupv3s.lts.opentelekomcloud.m.crossplane.io` |
| HostAccessV3 | `hostaccessv3s.lts.opentelekomcloud.m.crossplane.io` |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${Rand.RFC1123Subdomain}` | names of all resources and cloud names | Replace with a unique RFC1123-compatible suffix (or leave for the example generator to substitute) |
| `paths: [/var/log/*]` | HostAccessV3 `spec.forProvider.accessConfig` | Adjust the collection paths to the logs you want to ship |
| `namespace: test` | all resources | All resources are created in the `test` namespace - adjust if needed |

The host group is created without hosts. Install the LTS ICAgent on your hosts
for actual log collection.

## hostgroupv3.yaml

Test case: `hostgroupv3.yaml` - creates an LTS host group.

### APIs used

| Kind | Full API name |
|------|---------------|
| HostGroupV3 | `hostgroupv3s.lts.opentelekomcloud.m.crossplane.io` |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${Rand.RFC1123Subdomain}` | metadata name, `spec.forProvider.name` | Replace with a unique RFC1123-compatible suffix (or leave for the example generator to substitute) |
| `type: linux` | HostGroupV3 `spec.forProvider.type` | Change to `windows` for Windows hosts |
| `namespace: test` | all resources | All resources are created in the `test` namespace - adjust if needed |

## keywordsalarmrulev2.yaml

Test case: `keywordsalarmrulev2.yaml` - creates an LTS log group and stream, an
SMN topic, and a keyword alarm rule that sends notifications to the topic.

### APIs used

| Kind | Full API name |
|------|---------------|
| GroupV2 | `groupv2s.lts.opentelekomcloud.m.crossplane.io` |
| StreamV2 | `streamv2s.lts.opentelekomcloud.m.crossplane.io` |
| TopicV2 | `topicv2s.smn.opentelekomcloud.m.crossplane.io` |
| KeywordsAlarmRuleV2 | `keywordsalarmrulev2s.lts.opentelekomcloud.m.crossplane.io` |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${Rand.RFC1123Subdomain}` | names of all resources and cloud names | Replace with a unique RFC1123-compatible suffix (or leave for the example generator to substitute) |
| `keyword` / `condition` / `number` | KeywordsAlarmRuleV2 `spec.forProvider.keywordsRequests` | Adjust the keyword, comparison condition, and threshold |
| `severity: CRITICAL` | KeywordsAlarmRuleV2 `spec.forProvider.severity` | Change to `INFO`, `MINOR`, `MAJOR`, or `CRITICAL` |
| `name: lts-topic-...` | KeywordsAlarmRuleV2 `notificationRule.topics` | Keep in sync with the TopicV2 `spec.forProvider.name` |
| `namespace: test` | all resources | All resources are created in the `test` namespace - adjust if needed |

## lts.yaml

Test case: `lts.yaml` - creates a VPC, an LTS log group, an LTS log stream, and a VPC flow log that sends VPC traffic logs to LTS.

### APIs used

| Kind | Full API name |
|------|---------------|
| VpcV1 | `vpcv1s.vpc.opentelekomcloud.m.crossplane.io` |
| GroupV2 | `groupv2s.lts.opentelekomcloud.m.crossplane.io` |
| StreamV2 | `streamv2s.lts.opentelekomcloud.m.crossplane.io` |
| FlowLogV1 | `flowlogv1s.vpc.opentelekomcloud.m.crossplane.io` |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${Rand.RFC1123Subdomain}` | names of VpcV1, GroupV2, StreamV2, FlowLogV1 | Replace with a unique RFC1123-compatible suffix (or leave for the example generator to substitute) |
| `cidr: 192.168.0.0/16` | VpcV1 `spec.forProvider.cidr` | Adjust to a CIDR that does not conflict with existing networks |
| `ttlInDays: 7` | GroupV2 `spec.forProvider.ttlInDays` | Adjust the log retention period (1-365 days) |
| `namespace: test` | all resources | All resources are created in the `test` namespace - adjust if needed |

## notificationtemplatev2.yaml

Test case: `notificationtemplatev2.yaml` - creates an LTS notification template
with an SMS sub-template.

### APIs used

| Kind | Full API name |
|------|---------------|
| NotificationTemplateV2 | `notificationtemplatev2s.lts.opentelekomcloud.m.crossplane.io` |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${Rand.RFC1123Subdomain}` | metadata name, `spec.forProvider.name` | Replace with a unique RFC1123-compatible suffix (or leave for the example generator to substitute) |
| `language: en-us` | NotificationTemplateV2 `spec.forProvider.language` | Change to the language supported by your region |
| `content: "Alarm received;"` | NotificationTemplateV2 `spec.forProvider.templates` | Replace with your notification content |
| `namespace: test` | all resources | All resources are created in the `test` namespace - adjust if needed |

## quicksearchcriteriav1.yaml

Test case: `quicksearchcriteriav1.yaml` - creates an LTS quick search criteria for
a log group and stream.

### APIs used

| Kind | Full API name |
|------|---------------|
| GroupV2 | `groupv2s.lts.opentelekomcloud.m.crossplane.io` |
| StreamV2 | `streamv2s.lts.opentelekomcloud.m.crossplane.io` |
| QuickSearchCriteriaV1 | `quicksearchcriteriav1s.lts.opentelekomcloud.m.crossplane.io` |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${Rand.RFC1123Subdomain}` | names of all resources and cloud names | Replace with a unique RFC1123-compatible suffix (or leave for the example generator to substitute) |
| `criteria: content:test` | QuickSearchCriteriaV1 `spec.forProvider.criteria` | Replace with the search criteria you need |
| `type: ORIGINALLOG` | QuickSearchCriteriaV1 `spec.forProvider.type` | Keep `ORIGINALLOG` for raw logs |
| `namespace: test` | all resources | All resources are created in the `test` namespace - adjust if needed |

## streamv2.yaml

Test case: `streamv2.yaml` - creates an LTS log group and a log stream in it.

### APIs used

| Kind | Full API name |
|------|---------------|
| GroupV2 | `groupv2s.lts.opentelekomcloud.m.crossplane.io` |
| StreamV2 | `streamv2s.lts.opentelekomcloud.m.crossplane.io` |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${Rand.RFC1123Subdomain}` | names of both resources and cloud names | Replace with a unique RFC1123-compatible suffix (or leave for the example generator to substitute) |
| `ttlInDays: 7` | GroupV2 `spec.forProvider.ttlInDays` | Adjust the log retention period (1-365 days) |
| `namespace: test` | all resources | All resources are created in the `test` namespace - adjust if needed |

## structuringtemplatev3.yaml

Test case: `structuringtemplatev3.yaml` - creates an LTS log group and stream and
applies the built-in `CTS` structuring template.

### APIs used

| Kind | Full API name |
|------|---------------|
| GroupV2 | `groupv2s.lts.opentelekomcloud.m.crossplane.io` |
| StreamV2 | `streamv2s.lts.opentelekomcloud.m.crossplane.io` |
| StructuringTemplateV3 | `structuringtemplatev3s.lts.opentelekomcloud.m.crossplane.io` |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${Rand.RFC1123Subdomain}` | names of all resources and cloud names | Replace with a unique RFC1123-compatible suffix (or leave for the example generator to substitute) |
| `templateName: CTS` | StructuringTemplateV3 `spec.forProvider.templateName` | Change to another built-in template if needed |
| `namespace: test` | all resources | All resources are created in the `test` namespace - adjust if needed |

## transferv2.yaml

Test case: `transferv2.yaml` - creates an LTS log group and stream, an OBS bucket,
and a transfer task that exports logs to the bucket.

### APIs used

| Kind | Full API name |
|------|---------------|
| GroupV2 | `groupv2s.lts.opentelekomcloud.m.crossplane.io` |
| StreamV2 | `streamv2s.lts.opentelekomcloud.m.crossplane.io` |
| Bucket | `buckets.obs.opentelekomcloud.m.crossplane.io` |
| TransferV2 | `transferv2s.lts.opentelekomcloud.m.crossplane.io` |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${Rand.RFC1123Subdomain}` | names of all resources and cloud names | Replace with a unique RFC1123-compatible suffix (or leave for the example generator to substitute) |
| `bucket: lts-bkt-...` | Bucket `spec.forProvider.bucket` | Use a globally unique bucket name (lowercase, 3-63 chars) |
| `obsPeriod: 3` / `obsPeriodUnit: hour` | TransferV2 `logTransferDetail` | Adjust the transfer interval as required by LTS |
| `namespace: test` | all resources | All resources are created in the `test` namespace - adjust if needed |

## Prerequisites

This example requires a configured `ClusterProviderConfig` - see [providerconfig](../providerconfig/README.md).
