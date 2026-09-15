# EVS test cases

## volumev3.yaml

Test case: `volumev3.yaml` - creates an EVS (Elastic Volume Storage) volume
encrypted with a KMS key.

### APIs used

| Kind | Full API name |
|------|---------------|
| VolumeV3 | `volumev3s.evs.opentelekomcloud.m.crossplane.io` |
| KeyV1 | `keyv1s.kms.opentelekomcloud.m.crossplane.io` |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${Rand.RFC1123Subdomain}` | names of VolumeV3, KeyV1, KeyV1 `spec.forProvider.keyAlias` | Replace with a unique RFC1123-compatible suffix (or leave for the example generator to substitute) |
| `availabilityZone: eu-de-01` | VolumeV3 `spec.forProvider.availabilityZone` | Change to an AZ in your region |
| `size: 20` / `volumeType: SAS` | VolumeV3 | Adjust size and volume type to what you need |
| `realm: eu-de` | KeyV1 `spec.forProvider.realm` | Change to the region where the key should be created |
| `namespace: test` | all resources | All resources are created in the `test` namespace - adjust if needed |

## snapshotv2.yaml

Test case: `snapshotv2.yaml` - creates an EVS volume and a snapshot of it.

### APIs used

| Kind | Full API name |
|------|---------------|
| VolumeV3 | `volumev3s.evs.opentelekomcloud.m.crossplane.io` |
| SnapshotV2 | `snapshotv2s.evs.opentelekomcloud.m.crossplane.io` |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${Rand.RFC1123Subdomain}` | names of VolumeV3, SnapshotV2 | Replace with a unique RFC1123-compatible suffix (or leave for the example generator to substitute) |
| `availabilityZone: eu-de-01` | VolumeV3 `spec.forProvider.availabilityZone` | Change to an AZ in your region |
| `size: 12` / `volumeType: SAS` | VolumeV3 | Adjust size and volume type to what you need |
| `namespace: test` | all resources | All resources are created in the `test` namespace - adjust if needed |
