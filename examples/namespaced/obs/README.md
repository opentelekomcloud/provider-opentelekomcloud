# OBS test cases

## bucket.yaml

Test case: `bucket.yaml` - creates a basic OBS bucket with tags.

### APIs used

| Kind | Full API name |
|------|---------------|
| Bucket | `buckets.obs.opentelekomcloud.m.crossplane.io` |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${Rand.RFC1123Subdomain}` | name, `spec.forProvider.bucket` | Replace with a unique, globally-unique bucket name (lowercase, 3-63 chars, no underscores) |
| `namespace: test` | all resources | All resources are created in the `test` namespace - adjust if needed |

## bucket_encrypt.yaml

Test case: `bucket_encrypt.yaml` - creates a KMS key and an OBS bucket with KMS server-side encryption.

### APIs used

| Kind | Full API name |
|------|---------------|
| Bucket | `buckets.obs.opentelekomcloud.m.crossplane.io` |
| KeyV1 | `keyv1s.kms.opentelekomcloud.m.crossplane.io` |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${Rand.RFC1123Subdomain}` | names of Bucket, KeyV1, KeyV1 `spec.forProvider.keyAlias` | Replace with a unique bucket name and RFC1123-compatible suffix |
| `realm: eu-de` | KeyV1 `spec.forProvider.realm` | Change to the region where the key should be created |
| `providerConfigRef: name: default-tlp` | KeyV1 `spec.providerConfigRef` | Replace with the actual provider config name in your cluster |
| `namespace: test` | all resources | All resources are created in the `test` namespace - adjust if needed |

## bucket_logging.yaml

Test case: `bucket_logging.yaml` - creates two OBS buckets and a bucket ACL so the
source bucket can deliver access logs to the target bucket.

### APIs used

| Kind | Full API name |
|------|---------------|
| Bucket | `buckets.obs.opentelekomcloud.m.crossplane.io` |
| BucketACL | `bucketacls.obs.opentelekomcloud.m.crossplane.io` |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${Rand.RFC1123Subdomain}` | names of both Buckets, BucketACL | Replace with unique, globally-unique bucket names (lowercase, 3-63 chars) |
| `namespace: test` | all resources | All resources are created in the `test` namespace - adjust if needed |

## bucketinventory.yaml

Test case: `bucketinventory.yaml` - creates an OBS bucket and a bucket inventory
that generates weekly CSV reports of object metadata.

### APIs used

| Kind | Full API name |
|------|---------------|
| BucketInventory | `bucketinventories.obs.opentelekomcloud.m.crossplane.io` |
| Bucket | `buckets.obs.opentelekomcloud.m.crossplane.io` |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${Rand.RFC1123Subdomain}` | names of BucketInventory, Bucket | Replace with a unique, globally-unique bucket name and inventory suffix |
| `configurationId: test-id` | `spec.forProvider.configurationId` | Replace with a meaningful inventory configuration ID |
| `prefix: test-` / `filterPrefix: test-filter-prefix` | `spec.forProvider` | Adjust the output prefix and filter prefix |
| `frequency: Weekly` | `spec.forProvider.frequency` | Change to `Daily` if you need daily inventories |
| `namespace: test` | all resources | All resources are created in the `test` namespace - adjust if needed |

## bucketobject.yaml

Test case: `bucketobject.yaml` - creates an OBS bucket and uploads an encrypted
JSON object to it.

### APIs used

| Kind | Full API name |
|------|---------------|
| BucketObject | `bucketobjects.obs.opentelekomcloud.m.crossplane.io` |
| Bucket | `buckets.obs.opentelekomcloud.m.crossplane.io` |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${Rand.RFC1123Subdomain}` | names of BucketObject, Bucket | Replace with a unique, globally-unique bucket name and object suffix |
| `key: test_index_json` | `spec.forProvider.key` | Replace with the object key (path) you want to use |
| `content: ...` | `spec.forProvider.content` | Replace the JSON content with your own data |
| `contentType: application/json` | `spec.forProvider.contentType` | Change to the correct MIME type |
| `namespace: test` | all resources | All resources are created in the `test` namespace - adjust if needed |

## bucketobjectacl.yaml

Test case: `bucketobjectacl.yaml` - creates an OBS bucket, an object, and a
public-read ACL on that object.

### APIs used

| Kind | Full API name |
|------|---------------|
| BucketObjectACL | `bucketobjectacls.obs.opentelekomcloud.m.crossplane.io` |
| BucketObject | `bucketobjects.obs.opentelekomcloud.m.crossplane.io` |
| Bucket | `buckets.obs.opentelekomcloud.m.crossplane.io` |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${Rand.RFC1123Subdomain}` | names of all resources | Replace with a unique, globally-unique bucket name and suffixes |
| `key: test_index_json` | BucketObject `spec.forProvider.key` | Replace with the object key (path) you want to use |
| `content: ...` | BucketObject `spec.forProvider.content` | Replace the JSON content with your own data |
| `namespace: test` | all resources | All resources are created in the `test` namespace - adjust if needed |

## bucketpolicy-static.yaml

Test case: `bucketpolicy-static.yaml` - creates an OBS bucket with a static
bucket policy granting public Get/Put access.

### APIs used

| Kind | Full API name |
|------|---------------|
| BucketPolicy | `bucketpolicies.obs.opentelekomcloud.m.crossplane.io` |
| Bucket | `buckets.obs.opentelekomcloud.m.crossplane.io` |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${Rand.RFC1123Subdomain}` | BucketPolicy metadata name | Replace with a unique RFC1123-compatible suffix (or leave for the example generator to substitute) |
| `bucket: xp-bucket-statictest-0428` | Bucket `spec.forProvider.bucket` | Replace with a unique, globally-unique bucket name you actually want to create |
| `Resource: ["xp-bucket-statictest-0428/*"]` | BucketPolicy `spec.forProvider.policy` | Replace with the actual bucket name in the policy resource ARN |
| `namespace: test` | all resources | All resources are created in the `test` namespace - adjust if needed |

## bucketpolicy-dynamic.yaml

Test case: `bucketpolicy-dynamic.yaml` - creates an OBS bucket with a dynamic
bucket policy (uses `${this.bucket}` placeholder resolved by the provider).

### APIs used

| Kind | Full API name |
|------|---------------|
| BucketPolicy | `bucketpolicies.obs.opentelekomcloud.m.crossplane.io` |
| Bucket | `buckets.obs.opentelekomcloud.m.crossplane.io` |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${Rand.RFC1123Subdomain}` | names of BucketPolicy, Bucket | Replace with a unique, globally-unique bucket name and suffix |
| `Action: ["GetObject", "PutObject"]` | BucketPolicy `spec.forProvider.policy` | Adjust the allowed actions |
| `namespace: test` | all resources | All resources are created in the `test` namespace - adjust if needed |

## bucketreplication.yaml

Test case: `bucketreplication.yaml` - creates a source bucket (eu-de), a
destination bucket (eu-nl, using a different provider config), an IAM agency,
and a bucket replication rule between them.

### APIs used

| Kind | Full API name |
|------|---------------|
| BucketReplication | `bucketreplications.obs.opentelekomcloud.m.crossplane.io` |
| Bucket | `buckets.obs.opentelekomcloud.m.crossplane.io` |
| AgencyV3 | `agencyv3s.identity.opentelekomcloud.m.crossplane.io` |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${Rand.RFC1123Subdomain}` | names of all resources, Bucket `spec.forProvider.bucket`, AgencyV3 `spec.forProvider.name` | Replace with unique, globally-unique bucket names and RFC1123-compatible suffixes |
| `region: eu-de` | source Bucket `spec.forProvider.region` | Change to your source region |
| `region: eu-nl` | destination Bucket `spec.forProvider.region` | Change to your destination region |
| `providerConfigRef: name: default-nl` | destination Bucket `spec.providerConfigRef` | Replace with the actual provider config name for the destination region |
| `delegatedDomainName: op_svc_obs` | AgencyV3 `spec.forProvider.delegatedDomainName` | Replace with the actual delegated domain name |
| `namespace: test` | all resources | All resources are created in the `test` namespace - adjust if needed |

## bucketacl.yaml

Test case: `bucketacl.yaml` - creates an OBS bucket and a bucket ACL granting
full owner permissions.

### APIs used

| Kind | Full API name |
|------|---------------|
| BucketACL | `bucketacls.obs.opentelekomcloud.m.crossplane.io` |
| Bucket | `buckets.obs.opentelekomcloud.m.crossplane.io` |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${Rand.RFC1123Subdomain}` | names of BucketACL, Bucket | Replace with a unique, globally-unique bucket name and suffix |
| `namespace: test` | all resources | All resources are created in the `test` namespace - adjust if needed |

## Prerequisites

This example requires a configured `ClusterProviderConfig` - see [providerconfig](../providerconfig/README.md).
