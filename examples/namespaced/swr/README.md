# SWR test cases

## organizationsv2.yaml

Test case: `organizationsv2.yaml` - creates an SWR (SoftWare Repository) organization.

### APIs used

| Kind | Full API name |
|------|---------------|
| OrganizationV2 | `organizationv2s.swr.opentelekomcloud.m.crossplane.io` |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${Rand.RFC1123Subdomain}` | name, `spec.forProvider.name` | Replace with a unique RFC1123-compatible suffix (or leave for the example generator to substitute) |
| `namespace: test` | all resources | All resources are created in the `test` namespace - adjust if needed |

## repositoryv2.yaml

Test case: `repositoryv2.yaml` - creates an SWR organization and a private
repository within it.

### APIs used

| Kind | Full API name |
|------|---------------|
| RepositoryV2 | `repositoryv2s.swr.opentelekomcloud.m.crossplane.io` |
| OrganizationV2 | `organizationv2s.swr.opentelekomcloud.m.crossplane.io` |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${Rand.RFC1123Subdomain}` | names of RepositoryV2, OrganizationV2, `spec.forProvider.name` fields | Replace with a unique RFC1123-compatible suffix (or leave for the example generator to substitute) |
| `category: linux` | RepositoryV2 `spec.forProvider.category` | Change to `windows` if you need a Windows image repository |
| `namespace: test` | all resources | All resources are created in the `test` namespace - adjust if needed |

## domainv2.yaml

Test case: `domainv2.yaml` - creates an SWR organization, a repository, and a
read-only access domain for that repository.

### APIs used

| Kind | Full API name |
|------|---------------|
| DomainV2 | `domainv2s.swr.opentelekomcloud.m.crossplane.io` |
| RepositoryV2 | `repositoryv2s.swr.opentelekomcloud.m.crossplane.io` |
| OrganizationV2 | `organizationv2s.swr.opentelekomcloud.m.crossplane.io` |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${Rand.RFC1123Subdomain}` | names of all resources, `spec.forProvider.name` fields | Replace with a unique RFC1123-compatible suffix (or leave for the example generator to substitute) |
| `${data.swrAccessDomain}` | DomainV2 `spec.forProvider.accessDomain` | Replace with the actual access domain (endpoint) you want to use for pulling images |
| `permission: read` | DomainV2 `spec.forProvider.permission` | Change to `readwrite` if you also need push access |
| `deadline: forever` | DomainV2 `spec.forProvider.deadline` | Change to a specific date/time if you want the access to expire |
| `namespace: test` | all resources | All resources are created in the `test` namespace - adjust if needed |

## organizationpermissionsv2.yaml

Test case: `organizationpermissionsv2.yaml` - creates an SWR organization, an
IAM user, and grants the user admin permissions (auth: 3) on the organization.

### APIs used

| Kind | Full API name |
|------|---------------|
| OrganizationPermissionsV2 | `organizationpermissionsv2s.swr.opentelekomcloud.m.crossplane.io` |
| OrganizationV2 | `organizationv2s.swr.opentelekomcloud.m.crossplane.io` |
| UserV3 | `userv3s.identity.opentelekomcloud.m.crossplane.io` |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${Rand.RFC1123Subdomain}` | names of all resources, `spec.forProvider.name` fields | Replace with a unique RFC1123-compatible suffix (or leave for the example generator to substitute) |
| `auth: 3` | OrganizationPermissionsV2 `spec.forProvider.auth` | Change to `1` (read-only) or `2` (read-write) depending on the access level you need |
| `namespace: test` | all resources | All resources are created in the `test` namespace - adjust if needed |

## Prerequisites

This example requires a configured `ClusterProviderConfig` - see [providerconfig](../providerconfig/README.md).
