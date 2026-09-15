# Identity test cases

## aclv3.yaml

Test case: `aclv3.yaml` - creates an IAM console access control list (ACL) with IP and IPv6 ranges.

### APIs used

| Kind | Full API name |
|------|---------------|
| ACLV3 | `aclv3s.identity.opentelekomcloud.m.crossplane.io` |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${Rand.RFC1123Subdomain}` | metadata name | Replace with a unique RFC1123-compatible suffix (or leave for the example generator to substitute) |
| `range: 0.0.0.0-255.255.255.255` | ipRanges | Replace with the actual IP range you want to allow for console access |
| `range: 0000:0000:0000:0000:0000:0000:0000:0000-FFFF:FFFF:FFFF:FFFF:FFFF:FFFF:FFFF:FFFF` | ipv6Ranges | Replace with the actual IPv6 range you want to allow |
| `namespace: test` | all resources | All resources are created in the `test` namespace - adjust if needed |

## agencyv3.yaml

Test case: `agencyv3.yaml` - creates an IAM agency (delegation) with domain and project roles.

### APIs used

| Kind | Full API name |
|------|---------------|
| AgencyV3 | `agencyv3s.identity.opentelekomcloud.m.crossplane.io` |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${Rand.RFC1123Subdomain}` | metadata name | Replace with a unique RFC1123-compatible suffix (or leave for the example generator to substitute) |
| `${data.delegatedDomainName}` | AgencyV3 `spec.forProvider.delegatedDomainName` | Replace with the actual delegated domain name (the domain that delegates permissions) |
| `name: test_agency` | AgencyV3 `spec.forProvider.name` | Replace with a meaningful agency name |
| `project: eu-de` | projectRole | Change to your project region |
| `roles: [...]` | domainRoles, projectRole | Adjust the IAM roles to the ones you need to delegate |
| `namespace: test` | all resources | All resources are created in the `test` namespace - adjust if needed |

## credentialv3.yaml

Test case: `credentialv3.yaml` - creates an IAM user with a password secret and a credential for that user.

### APIs used

| Kind | Full API name |
|------|---------------|
| UserV3 | `userv3s.identity.opentelekomcloud.m.crossplane.io` |
| CredentialV3 | `credentialv3s.identity.opentelekomcloud.m.crossplane.io` |
| Secret | `secrets` (core/v1) |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${Rand.RFC1123Subdomain}` | names of UserV3, CredentialV3 | Replace with a unique RFC1123-compatible suffix (or leave for the example generator to substitute) |
| `${data.userKey}` | Secret `example-secret`, key `example-key` | Replace with the base64-encoded user password (e.g. `base64 <<< '<password>'`) |
| `namespace: test` | all resources | All resources are created in the `test` namespace - adjust if needed |

## groupmembershipv3.yaml

Test case: `groupmembershipv3.yaml` - creates an IAM group and two users, then adds both users to the group via GroupMembershipV3.

### APIs used

| Kind | Full API name |
|------|---------------|
| GroupV3 | `groupv3s.identity.opentelekomcloud.m.crossplane.io` |
| UserV3 | `userv3s.identity.opentelekomcloud.m.crossplane.io` |
| GroupMembershipV3 | `groupmembershipv3s.identity.opentelekomcloud.m.crossplane.io` |
| Secret | `secrets` (core/v1) |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${Rand.RFC1123Subdomain}` | names of GroupV3, UserV3 (x2), GroupMembershipV3 | Replace with a unique RFC1123-compatible suffix (or leave for the example generator to substitute) |
| `${data.userKey}` | Secret `example-secret`, key `example-key` | Replace with the base64-encoded user password (e.g. `base64 <<< '<password>'`) |
| `namespace: test` | all resources | All resources are created in the `test` namespace - adjust if needed |

## groupv3.yaml

Test case: `groupv3.yaml` - creates an IAM group.

### APIs used

| Kind | Full API name |
|------|---------------|
| GroupV3 | `groupv3s.identity.opentelekomcloud.m.crossplane.io` |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${Rand.RFC1123Subdomain}` | name, `spec.forProvider.name` | Replace with a unique RFC1123-compatible suffix (or leave for the example generator to substitute) |
| `namespace: test` | all resources | All resources are created in the `test` namespace - adjust if needed |

## mappingv3.yaml

Test case: `mappingv3.yaml` - creates an IAM identity provider mapping rule.

### APIs used

| Kind | Full API name |
|------|---------------|
| MappingV3 | `mappingv3s.identity.opentelekomcloud.m.crossplane.io` |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${Rand.RFC1123Subdomain}` | metadata name | Replace with a unique RFC1123-compatible suffix (or leave for the example generator to substitute) |
| `mappingId: ACME` | `spec.forProvider.mappingId` | Replace with your identity provider's mapping ID |
| `rules: [...]` | `spec.forProvider.rules` | Adjust the SAML attribute mapping rules to match your IdP |
| `namespace: test` | all resources | All resources are created in the `test` namespace - adjust if needed |

## passwordpolicyv3.yaml

Test case: `passwordpolicyv3.yaml` - configures the IAM password policy.

### APIs used

| Kind | Full API name |
|------|---------------|
| PasswordPolicyV3 | `passwordpolicyv3s.identity.opentelekomcloud.m.crossplane.io` |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${Rand.RFC1123Subdomain}` | metadata name | Replace with a unique RFC1123-compatible suffix (or leave for the example generator to substitute) |
| `minimumPasswordLength: 8` / other policy fields | `spec.forProvider` | Adjust to your organization's password policy requirements |
| `namespace: test` | all resources | All resources are created in the `test` namespace - adjust if needed |

## projectv3.yaml

Test case: `projectv3.yaml` - creates an IAM project (region-scoped).

### APIs used

| Kind | Full API name |
|------|---------------|
| ProjectV3 | `projectv3s.identity.opentelekomcloud.m.crossplane.io` |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${Rand.RFC1123Subdomain}` | name, `spec.forProvider.name` | Replace with a unique RFC1123-compatible suffix (or leave for the example generator to substitute) |
| `name: eu-de_iam-proj-...` | `spec.forProvider.name` | The prefix (`eu-de_`) must match your region - change to your region's prefix |
| `namespace: test` | all resources | All resources are created in the `test` namespace - adjust if needed |

## protectionpolicyv3.yaml

Test case: `protectionpolicyv3.yaml` - configures the IAM self-management protection policy.

### APIs used

| Kind | Full API name |
|------|---------------|
| ProtectionPolicyV3 | `protectionpolicyv3s.identity.opentelekomcloud.m.crossplane.io` |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${Rand.RFC1123Subdomain}` | metadata name | Replace with a unique RFC1123-compatible suffix (or leave for the example generator to substitute) |
| `selfManagement: [...]` | `spec.forProvider.selfManagement` | Adjust which authentication methods (accessKey, email, mobile, password) are enabled |
| `namespace: test` | all resources | All resources are created in the `test` namespace - adjust if needed |

## protocolv3.yaml

Test case: `protocolv3.yaml` - creates a SAML identity provider (ProviderV3), a mapping rule (MappingV3), and binds them together with a SAML protocol (ProtocolV3).

### APIs used

| Kind | Full API name |
|------|---------------|
| ProviderV3 | `providerv3s.identity.opentelekomcloud.m.crossplane.io` |
| MappingV3 | `mappingv3s.identity.opentelekomcloud.m.crossplane.io` |
| ProtocolV3 | `protocolv3s.identity.opentelekomcloud.m.crossplane.io` |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${Rand.RFC1123Subdomain}` | names of ProviderV3, MappingV3, ProtocolV3 | Replace with a unique RFC1123-compatible suffix (or leave for the example generator to substitute) |
| `name: ACME-...` | ProviderV3 `spec.forProvider.name` | Replace with your identity provider's name |
| `mappingId: ACME` | MappingV3 `spec.forProvider.mappingId` | Replace with your identity provider's mapping ID |
| `rules: [...]` | MappingV3 `spec.forProvider.rules` | Adjust the SAML attribute mapping rules to match your IdP |
| `namespace: test` | all resources | All resources are created in the `test` namespace - adjust if needed |

## providerv3.yaml

Test case: `providerv3.yaml` - creates an IAM identity provider (IdP).

### APIs used

| Kind | Full API name |
|------|---------------|
| ProviderV3 | `providerv3s.identity.opentelekomcloud.m.crossplane.io` |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${Rand.RFC1123Subdomain}` | name, `spec.forProvider.name` | Replace with a unique RFC1123-compatible suffix (or leave for the example generator to substitute) |
| `name: ACME-...` | `spec.forProvider.name` | Replace with your identity provider's name |
| `namespace: test` | all resources | All resources are created in the `test` namespace - adjust if needed |

## roleassignmentv3.yaml

Test case: `roleassignmentv3.yaml` - creates an IAM group, project, and custom role, then assigns the role to the group on the project.

### APIs used

| Kind | Full API name |
|------|---------------|
| RoleAssignmentV3 | `roleassignmentv3s.identity.opentelekomcloud.m.crossplane.io` |
| GroupV3 | `groupv3s.identity.opentelekomcloud.m.crossplane.io` |
| ProjectV3 | `projectv3s.identity.opentelekomcloud.m.crossplane.io` |
| RoleV3 | `rolev3s.identity.opentelekomcloud.m.crossplane.io` |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${Rand.RFC1123Subdomain}` | names of RoleAssignmentV3, GroupV3, ProjectV3, RoleV3 | Replace with a unique RFC1123-compatible suffix (or leave for the example generator to substitute) |
| `name: eu-de_iam-proj-...` | ProjectV3 `spec.forProvider.name` | The prefix (`eu-de_`) must match your region - change to your region's prefix |
| `statement: [...]` | RoleV3 `spec.forProvider.statement` | Adjust the IAM policy statement to the permissions you need |
| `resource: OBS:*:*:bucket:test-bucket` | RoleV3 statement | Replace `test-bucket` with the actual bucket name |
| `namespace: test` | all resources | All resources are created in the `test` namespace - adjust if needed |

## rolev3.yaml

Test case: `rolev3.yaml` - creates a custom IAM role with an OBS bucket policy.

### APIs used

| Kind | Full API name |
|------|---------------|
| RoleV3 | `rolev3s.identity.opentelekomcloud.m.crossplane.io` |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${Rand.RFC1123Subdomain}` | name, `spec.forProvider.displayName` | Replace with a unique RFC1123-compatible suffix (or leave for the example generator to substitute) |
| `resource: OBS:*:*:bucket:test-bucket` | `spec.forProvider.statement[].resource` | Replace `test-bucket` with the actual OBS bucket name |
| `g:ProjectName: [eu-de]` | statement condition | Change to your region |
| `namespace: test` | all resources | All resources are created in the `test` namespace - adjust if needed |

## usergroupmembershipv3.yaml

Test case: `usergroupmembershipv3.yaml` - creates an IAM user with a password secret and two groups, then adds the user to the groups via UserGroupMembershipV3.

### APIs used

| Kind | Full API name |
|------|---------------|
| UserGroupMembershipV3 | `usergroupmembershipv3s.identity.opentelekomcloud.m.crossplane.io` |
| GroupV3 | `groupv3s.identity.opentelekomcloud.m.crossplane.io` |
| UserV3 | `userv3s.identity.opentelekomcloud.m.crossplane.io` |
| Secret | `secrets` (core/v1) |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${Rand.RFC1123Subdomain}` | names of UserGroupMembershipV3, GroupV3 (x2), UserV3 | Replace with a unique RFC1123-compatible suffix (or leave for the example generator to substitute) |
| `${data.userKey}` | Secret `example-secret`, key `example-key` | Replace with the base64-encoded user password (e.g. `base64 <<< '<password>'`) |
| `namespace: test` | all resources | All resources are created in the `test` namespace - adjust if needed |

## userv3.yaml

Test case: `userv3.yaml` - creates an IAM user with a password from a Kubernetes secret.

### APIs used

| Kind | Full API name |
|------|---------------|
| UserV3 | `userv3s.identity.opentelekomcloud.m.crossplane.io` |
| Secret | `secrets` (core/v1) |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${Rand.RFC1123Subdomain}` | name, `spec.forProvider.name` | Replace with a unique RFC1123-compatible suffix (or leave for the example generator to substitute) |
| `${data.userKey}` | Secret `example-secret`, key `example-key` | Replace with the base64-encoded user password (e.g. `base64 <<< '<password>'`) |
| `namespace: test` | all resources | All resources are created in the `test` namespace - adjust if needed |

## Prerequisites

This example requires a configured `ClusterProviderConfig` - see [providerconfig](../providerconfig/README.md).
