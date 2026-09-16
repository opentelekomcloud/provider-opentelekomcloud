# ProviderConfig test cases

## providerconfig.yaml

Test case: `providerconfig.yaml` - creates the `ClusterProviderConfig` used by all
namespaced managed resources. It reads the OpenTelekomCloud credentials from the
`provider-secret` Kubernetes secret (key `credentials`) in the `crossplane-system`
namespace.

### APIs used

| Kind | Full API name |
|------|---------------|
| ClusterProviderConfig | `clusterproviderconfigs.opentelekomcloud.m.crossplane.io` |
| Secret | `secrets` (core/v1) |

### The credentials secret

Create `provider-secret` in `crossplane-system` (see `secret.yaml.tmpl`). Its
`credentials` key must contain a JSON document with the provider settings.
#### Required fields:

| Key | What to set |
|-----|-------------|
| `access_key` | Your OTC AK (access key) - create one under IAM > My Credentials > API Keys |
| `secret_key` | Your OTC SK (secret key) - shown once when the AK is created |
| `auth_url` | The IAM endpoint of your region, e.g. `https://iam.eu-de.otc.t-systems.com/v3` |
| `domain_id` or `domain_name` | The ID or name of the domain (account/project) the AK belongs to - one of the two is required |
| `tenant_id` or `tenant_name` | The ID or name of the project (tenant) to act in - one of the two is required |

Alternative auth options (instead of AK/SK): `user_name` + `password`, or
`user_id` + `password`, or `token`. The provider passes all of these through to
the Terraform OTC provider
(https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/latest/docs),
so any field supported there can be used: `region`, `security_token`, `passcode`,
`insecure`, `endpoint_type`, `cacert_file`, `cert`, `key`, `agency_name`,
`agency_domain_name`, `delegated_project` (delegated/agency auth).

#### Optional settings (the values in the template are sane defaults):

| Key | What to set |
|-----|-------------|
| `swauth` | `"false"` unless you use SWS (software) authentication |
| `allow_reauth` | `"true"` to automatically re-authenticate when the token expires |
| `max_retries` | number of retries on API errors (default `2`) |
| `max_backoff_retries` | number of backoff retries (default `6`) |
| `backoff_retry_timeout` | backoff timeout in seconds (default `60`) |
| `insecure` | `"false"` - set to `"true"` only for private endpoints with self-signed certs |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `xxxxxxxxxxx` placeholders | `secret.yaml.tmpl` `stringData.credentials` | Replace with your real AK/SK, `auth_url`, `domain_id`/`domain_name`, `tenant_id`/`tenant_name` |
| `auth_url` | `secret.yaml.tmpl` | Change the region in the URL (`eu-de` -> your region, e.g. `https://iam.eu-de.otc.t-systems.com/v3`) |
| secret name/namespace | `providerconfig.yaml` `spec.credentials.secretRef` | Defaults to `provider-secret` in `crossplane-system` - change if you put the secret elsewhere |
| provider config name | `providerconfig.yaml` `metadata.name` | Defaults to `default` - some examples reference it explicitly (e.g. `default-tlp`, `default-nl` in `obs/bucketreplication.yaml`) |
