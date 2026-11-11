# DMS (Kafka) test cases

All e2e resources live in `kafka.yaml`: one single-node Kafka instance plus its topic, user, user
permission and consumer group. uptest runs every example file separately, so dependent resources must be in
the same file. Resources find each other via `*Selector.matchLabels` on `testing.upbound.io/example-name`.

| File | e2e | Depends on |
|------|-----|------------|
| `kafka.yaml` | yes | - |
| `smart-connect.yaml` | no | `kafka.yaml` (VPC/subnet/SG and peer instance) |
| `reassign-partitions.yaml` | no | `smart-connect.yaml` (cluster instance and topic) |

```shell
UPTEST_EXAMPLE_LIST="examples/namespaced/dms/kafka.yaml"
```

Non-e2e resources carry the `upjet.upbound.io/manual-intervention` annotation.

## kafka.yaml

### Instance

Creates a single-node Kafka instance with SASL enabled in a new VPC/subnet/security group. SASL (`accessUser` + password) is required by `UserV2` and `UserPermissionV1`.

#### APIs used

| Kind | Full API name |
|------|---------------|
| Secret | `secrets` (core/v1) |
| VpcV1 | `vpcv1s.vpc.opentelekomcloud.m.crossplane.io` |
| SubnetV1 | `subnetv1s.vpc.opentelekomcloud.m.crossplane.io` |
| SecgroupV2 | `secgroupv2s.compute.opentelekomcloud.m.crossplane.io` |
| DedicatedInstanceV2 | `dedicatedinstancev2s.dms.opentelekomcloud.m.crossplane.io` |

#### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${Rand.RFC1123Subdomain}` | names of all resources | Replace with a unique RFC1123-compatible suffix (or leave for uptest to substitute) |
| `${data.dmsKey}` | Secret `example-dms-password`, key `example-dms-key` | Replace with the base64-encoded Kafka password (8-32 chars, upper/lower case, digit and special character). For e2e, add `dmsKey` to the `UPTEST_DATASOURCE` secret |
| `flavorId: s6.2u4g.single.small` | DedicatedInstanceV2 | Adjust to a single-node flavor available in your region |
| `storageSpecCode: dms.physical.storage.ultra.v2` | DedicatedInstanceV2 | Adjust to a storage spec supported by the flavor in your region |
| `availableZones: [eu-de-01]` | DedicatedInstanceV2 | Change to an AZ in your region |
| `accessUser: crossplane` | DedicatedInstanceV2 | SASL user name of the instance |
| `cidr`, `gatewayIp` | VpcV1, SubnetV1, SecgroupV2 rule | Adjust if `192.168.0.0/16` collides with existing networks |

### Topic

Creates a topic in the instance.

#### APIs used

| Kind | Full API name |
|------|---------------|
| TopicV2 | `topicv2s.dms.opentelekomcloud.m.crossplane.io` |

#### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `name: crossplane-topic` | TopicV2 | Topic name, unique within the instance |
| `replication: 1` | TopicV2 | Must not exceed the broker count of the instance |
| `partition: 3` | TopicV2 | Number of partitions |

### User

Creates a SASL user in the instance.

#### APIs used

| Kind | Full API name |
|------|---------------|
| Secret | `secrets` (core/v1) |
| UserV2 | `userv2s.dms.opentelekomcloud.m.crossplane.io` |

#### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${data.dmsKey}` | Secret `example-dms-user-password`, key `example-dms-user-key` | Replace with the base64-encoded user password |
| `username: crossplane-user` | UserV2 | 4-64 characters, must start with a letter |

### User permission

Grants the user access to the topic.

#### APIs used

| Kind | Full API name |
|------|---------------|
| UserPermissionV1 | `userpermissionv1s.dms.opentelekomcloud.m.crossplane.io` |

#### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `accessPolicy: all` | UserPermissionV1 `policies[]` | One of `all`, `pub`, `sub` |

### Consumer group

Creates a consumer group in the instance.

#### APIs used

| Kind | Full API name |
|------|---------------|
| ConsumerGroupV2 | `consumergroupv2s.dms.opentelekomcloud.m.crossplane.io` |

#### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `groupName: crossplane-consumer-group` | ConsumerGroupV2 | Consumer group name, unique within the instance |

## smart-connect.yaml

Test case: `smart-connect.yaml` - creates a 3-broker cluster instance with a topic, enables Smart Connect
on it, creates a Kafka replicator task that pushes the topic to the single-node instance from
`kafka.yaml`, and pauses the task. Not part of e2e (expensive cluster instance).

### APIs used

| Kind | Full API name |
|------|---------------|
| Secret | `secrets` (core/v1) |
| DedicatedInstanceV2 | `dedicatedinstancev2s.dms.opentelekomcloud.m.crossplane.io` |
| TopicV2 | `topicv2s.dms.opentelekomcloud.m.crossplane.io` |
| SmartConnectV2 | `smartconnectv2s.dms.opentelekomcloud.m.crossplane.io` |
| SmartConnectTaskV2 | `smartconnecttaskv2s.dms.opentelekomcloud.m.crossplane.io` |
| SmartConnectTaskActionV2 | `smartconnecttaskactionv2s.dms.opentelekomcloud.m.crossplane.io` |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `flavorId: c6.2u4g.cluster` / `brokerNum: 3` | DedicatedInstanceV2 `dms-cluster-instance` | Adjust to a cluster flavor available in your region |
| `${data.dmsKey}` | Secret `example-dms-peer-password` | Password of the peer instance SASL user (same as in `kafka.yaml`) |
| `userName: crossplane` | SmartConnectTaskV2 `sourceTask[]` | SASL user of the peer instance (`accessUser` in `kafka.yaml`) |
| `replicationFactor: 1` | SmartConnectTaskV2 `sourceTask[]` | Must not exceed the broker count of the peer instance |
| `action: pause` | SmartConnectTaskActionV2 | One of `pause`, `resume`, `restart` |

## reassign-partitions.yaml

Test case: `reassign-partitions.yaml` - reassigns the partitions of the cluster topic across 3 brokers.
Not part of e2e (requires the cluster instance from `smart-connect.yaml`).

### APIs used

| Kind | Full API name |
|------|---------------|
| ReassignPartitionsV2 | `reassignpartitionsv2s.dms.opentelekomcloud.m.crossplane.io` |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `brokers: [0, 1, 2]` / `replicationFactor: 3` | ReassignPartitionsV2 `reassignments[]` | Broker IDs and replica count; must fit the broker count of the instance |
| `throttle` | ReassignPartitionsV2 | Replication bandwidth limit, bytes/s |

## Prerequisites

This example requires a configured `ClusterProviderConfig` - see [providerconfig](../providerconfig/README.md).
