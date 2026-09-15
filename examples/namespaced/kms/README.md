# KMS test cases

## kms.yaml

Test case: `kms.yaml` - creates an IAM user, a KMS key, and a grant giving the user
encryption operations on the key.

### APIs used

| Kind | Full API name |
|------|---------------|
| UserV3 | `userv3s.identity.opentelekomcloud.m.crossplane.io` |
| KeyV1 | `keyv1s.kms.opentelekomcloud.m.crossplane.io` |
| GrantV1 | `grantv1s.kms.opentelekomcloud.m.crossplane.io` |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${Rand.RFC1123Subdomain}` | UserV3 `spec.forProvider.name`, KeyV1 `spec.forProvider.keyAlias` | Replace with a unique RFC1123-compatible suffix (or leave for the example generator to substitute) |
| `passwordSecretRef` (secret `example-secret`, key `example-key`) | UserV3 `spec.forProvider.passwordSecretRef` | The referenced secret does not exist in this file - create it before applying, e.g. `kubectl create secret generic example-secret -n test --from-literal=example-key=<user-password>` |
| `realm: eu-de` | KeyV1 `spec.forProvider.realm` | Change to the region where the key should be created |
| `operations` | GrantV1 `spec.forProvider.operations` | Adjust to the key operations you want to grant (`describe-key`, `create-datakey`, `encrypt-datakey`) |
| `namespace: test` | all resources | All resources are created in the `test` namespace - adjust if needed |
