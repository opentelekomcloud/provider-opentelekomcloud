# Web Identity (workload identity) authentication

> [!NOTE]
> Web Identity is available on the **namespaced** API group
> `opentelekomcloud.m.crossplane.io` only. The legacy cluster-scoped
> `ProviderConfig`/`ClusterProviderConfig` does not support
> `spec.credentials.source: WebIdentity`.

This guide shows how to authenticate the provider to T Cloud Public **without
storing a long-lived password or AK/SK in a Kubernetes secret**. Instead of static
credentials, the provider reads a short-lived **OIDC identity token** from the
provider pod, exchanges it with T Cloud IAM for an OTC token (and optionally for
short-lived AK/SK), and uses those to manage resources.

Web Identity enables secretless, least-privilege authentication:

1. An OIDC identity provider (IdP) issues the token the provider pod presents
   (the CCE cluster API server for OTC CCE, or an external IdP).
2. T Cloud IAM trusts that IdP and maps its identity to a virtual IAM user in a
   local IAM group; the group holds the roles needed to manage resources in a
   target project or domain.
3. The provider reads the token, calls T Cloud IAM to exchange it, and uses the
   result to configure the Terraform OTC provider.

The token source is configurable per-`ProviderConfig` via the `tokenConfig` API:
it can be read from a **filesystem path** (a projected Kubernetes
service-account token — recommended) or from a **Kubernetes Secret**.

## How Web Identity differs from the other credential sources

| | `Secret` (default) | `WebIdentity` |
|---|---|---|
| Credential stored | Long-lived AK/SK (or password) in a Secret | Nothing long-lived; only a short-lived OIDC token |
| Permission boundary | The AK/SK owner | The mapped IAM group + its roles |
| Rotation | Manual (edit the Secret) | Automatic (projected token) or per-token refresh |
| Scope | whatever the AK/SK can do | the project/domain set in `scope` |

`WebIdentity` is useful when you want to avoid storing T Cloud credentials in
Kubernetes at all, scope the provider to a single project, and rotate credentials
without touching a Secret.

## Prerequisites

- A T Cloud account with **Security Administrator** privileges (required to
  create federated identity providers and identity mappings, and to manage IAM
  groups/roles). See the
  [User Management Model](https://docs.otc.t-systems.com/en-us/usermanual/iam/iam_01_0034.html).
- `kubectl` configured for the target cluster.
- Crossplane installed, and the T Cloud provider installed.
- Knowledge of the target **region**, **project** and (optionally) **domain**
  the provider should operate in.
- For the recommended CCE path: an **OTC CCE cluster of version v1.19.16 or
  later** (the cluster API server acts as the OIDC IdP).

## Overview of the moving parts

```
   +---------------------------------------------------------------+
   |  Kubernetes cluster                                           |
   |                                                               |
   |  Provider pod                                                 |
   |   - reads OIDC ID token from a file or Secret                 |
   |   - POST {authURL}/v3.0/OS-AUTH/id-token/tokens (X-Idp-Id)    |
   |   - POST {authURL}/v3.0/OS-CREDENTIAL/securitytokens (opt.)   |
   +--------|------------------------------------------------------+
            v
    T Cloud IAM (federated identity)
    - OIDC identity provider (trusts the CCE API server or an external IdP)
    - identity mapping (federated identity -> virtual user + group)
    - identity protocol (binds provider + mapping; Terraform/Crossplane only)
             |
             v
    IAM group with the IAM roles to manage resources in the
    target project (or domain) scope
```

Three things must be set up and **agree with each other**:

1. An **OIDC IdP** that mints the ID token the pod presents.
2. **T Cloud IAM federation** that trusts the IdP and maps its identities to a
   virtual IAM user in a group with the required permissions.
3. **Kubernetes** objects that (a) get the token into the provider pod and
   (b) tell the provider how to use it (`ProviderConfig`).

They are wired together by two values that must match exactly:

| Value | Where it appears | Must match |
|-------|------------------|-----------|
| IdP name / `idpID` | IAM identity provider `name`; `ProviderConfig.spec.credentials.webIdentity.idpID`; the `X-Idp-Id` header | the provider's `idpID` |
| Issuer (`iss`) | the token's `iss` claim; IAM `access_config.provider_url` | the IdP's issuer URL |

## Step 1: Choose the OIDC identity provider (IdP)

### Option A: the CCE cluster API server (recommended for OTC CCE)

For an OTC CCE cluster (**v1.19.16 or later**), the cluster API server *is* the
OIDC identity provider. It exposes an OIDC issuer and a signing-key endpoint
(JWKS), and Kubernetes projects a short-lived service-account token into the
provider pod (Step 3). No external IdP, no CronJob, and no Secret to keep fresh
are needed.

#### 1.1 Obtain the cluster signing public key (JWKS)

Fetch the cluster's OIDC signing key with `kubectl`. This is the **Signing
Key** you will register with T Cloud IAM in Step 2.

```console
kubectl get --raw /openid/v1/jwks
# {"keys":[{"use":"sig","kty":"RSA","kid":"***","alg":"RS256","n":"***","e":"AQAB"}]}
```

The returned JSON is the JWKS. The fields (`alg`, `kty`, `kid`, `use`, `n`, `e`)
go into `signing_key` in T Cloud IAM (Step 2).

#### 1.2 Note the issuer URL and client ID

- **Issuer** — the default is
  `https://kubernetes.default.svc.cluster.local`. If the cluster has
  **OIDC Provider** enabled under *Settings > Kubernetes*, use the
  **Service Account Issuer** (`service-account-issuer`) shown in *Settings >
  Kubernetes* instead. This is the `iss` claim of every token the cluster
  issues and it **must** be recorded as `provider_url` in T Cloud IAM (Step 2).
- **Client ID** — choose a value (e.g. `crossplane-provider`). This is the token
  **audience** you will set when projecting the service-account token (Step 3).
  A client ID must not consist only of digits (if it does, quote it in the
  workload YAML).

#### 1.3 Prepare the identity mapping key

The mapping (created in T Cloud IAM, Step 2.2) keys off the token's `sub`
claim. For a projected service-account token, `sub` is
`system:serviceaccount:<namespace>:<serviceaccount>`. For the recommended setup
(a ServiceAccount `oidc-token` in `crossplane-system`), the `sub` is:

```
system:serviceaccount:crossplane-system:oidc-token
```

The mapping's `remote` rule must match this `sub` (Step 2.2).

#### 1.4 Enable OIDC Provider / Service Account Token Volume Projection on the cluster

The CCE API server can act as the OIDC IdP. Before
the projected service-account token (Step 3) carries the issuer/JWKS the provider
needs, you must check if **OIDC Provider** (a.k.a. *Service Account Token Volume
Projection*) is enabled on the cluster. This is a **cluster-level** setting available on
CCE **v1.19.16 or later**.

- **Recommended / verified path — the CCE console.** In the CCE console, check
  **Service Account Token Volume Projection** under *Settings >
  Kubernetes*. The cluster's **Service Account Issuer**
  (`service-account-issuer`) is shown there — this is the value you record as
  `provider_url` in T Cloud IAM (see 1.2). If you leave it at the default, the
  issuer is `https://kubernetes.default.svc.cluster.local`.
- **Terraform.** The `opentelekomcloud_cce_cluster_v3` resource has no dedicated
  attribute for this. For a brand-new cluster it may be passable via the
  creation-only `extend_param` map, but the exact key is **not verified**.

> **Note:** toggling this setting is cluster-wide and restarts the CCE API
> server. It changes the token issuer/JWKS exposed by the API server, requires
> no node changes, and does not affect workloads that do not request projected
> service-account tokens.

### Option B: an external IdP (alternative)

If you are not on OTC CCE, or prefer an external identity system, you can use
any OIDC IdP you control (e.g. **Keycloak**, dex, ...). In that case the token
is typically minted into a Kubernetes Secret by an external controller, and the
provider reads it with `tokenConfig.source: Secret` (Step 4, Option B).

The key requirements are the same: the IdP must expose a stable issuer (`iss`)
and a signing key (JWKS) you can register with T Cloud IAM, and it must be able
to issue an **OIDC ID token** (not an access token) whose claims (`iss`, `sub`,
`aud`) line up with the identity provider and mapping you register in Step 2.
See "External IdP notes" under Step 4 for the specifics that bite in practice.

## Step 2: Configure T Cloud IAM federation

T Cloud IAM must be able to accept the IdP token and map it into a usable
IAM group. In the **console** that takes two federated-identity objects — an
identity **provider** and an identity **mapping** — plus the **IAM roles**
granted to the mapped group. (Terraform and Crossplane additionally need an
identity **protocol** that binds the two.) All of this can be set up three
ways:

- **T Cloud IAM console** — *Identity Providers* (and the role-assignment pages
  for the roles).
- **Terraform** — `opentelekomcloud_identity_provider`,
  `opentelekomcloud_identity_mapping_v3`,
  `opentelekomcloud_identity_protocol_v3`, and
  `opentelekomcloud_identity_role_assignment_v3`.
- **Crossplane** — if you already have an AK/SK-based `ProviderConfig` with
  sufficient permissions, you can manage them as managed resources
  (`ProviderV3`, `MappingV3`, `ProtocolV3`, `RoleAssignmentV3` under the
  `identity.opentelekomcloud.m.crossplane.io` API group).

### 2.1 Identity provider

Register the IdP. The fields to set, and the values they must take:

- **Name** — unique, up to 64 chars of letters/digits/`_`/`-`, e.g.
  `crossplane-workload-identity`. **This name is the `idpID`** referenced
  everywhere else (it becomes the `X-Idp-Id` header).
- **Protocol** — `oidc` (the provider's protocol type; not to be confused with
  the identity *protocol* object used by Terraform/Crossplane).
- **Enabled** — `true`. A disabled provider rejects all token exchanges.
- **Access type** — `program` (programmatic access).
- **Issuer URL** (`provider_url`) — the IdP's `iss`; for CCE the cluster's
  `service-account-issuer` (or `https://kubernetes.default.svc.cluster.local`
  as default). Must equal the token's `iss` claim.
- **Client ID** — the IdP's client id; for CCE this is the token **audience**
  you project in Step 3 (e.g. `crossplane-provider`).
- **Signing key** (`signing_key`) — the IdP's public JWKS; for CCE the keys from
  `kubectl get --raw /openid/v1/jwks`.

### 2.2 Identity mapping

The mapping maps a federated identity to a **virtual IAM user** (a
federated/virtual user created on demand — you do **not** need a pre-existing
local IAM user) in a **local IAM group**. The group is what carries the
permissions (Step 2.3).

- **Match the token's `sub`.** For the CCE path, `sub` is
  `system:serviceaccount:<namespace>:<serviceaccount>` (e.g.
  `system:serviceaccount:crossplane-system:oidc-token` — see Step 1.3). Match it
  with a `sub` rule. For an external IdP, match whatever claim you want to key
  off.

The mapping rules are a list of `{ local, remote }` rules: `remote` matches
token claims (`type` plus `any_one_of` / `not_any_of`), `local` names the target
IAM user (`user.name`, a virtual user) and the group(s) (`group.name` /
`groups`) to attach it to. See the
[Syntax of Identity Conversion Rules](https://docs.otc.t-systems.com/en-us/usermanual/iam/en-us_topic_0079620340.html)
for the full rule grammar.

Example (CCE path — maps the projected service-account token to the virtual
user `fed-crossplane-provider-user` in the `crossplane-managers` group):

```json
[
  {
    "local": [
      { "user": { "name": "fed-crossplane-provider-user" } },
      { "group": { "name": "crossplane-managers" } }
    ],
    "remote": [
      {
        "type": "sub",
        "any_one_of": ["system:serviceaccount:crossplane-system:oidc-token"]
      }
    ]
  }
]
```

The `user` is virtual and need not exist beforehand; the `group`
(`crossplane-managers`) **must** already exist.

### 2.3 Grant permissions to the group

Attach the IAM roles the provider needs to the mapped group for the
target project/domain. This is the permission boundary for Web Identity — the
provider can do exactly what this group can. Prefer a **least-privilege** custom
role over `admin`.

## Step 3: Get the token into the provider pod

The provider reads the OIDC token from a path in its own filesystem (Option A)
or from a Secret (Option B). For Option A you control this with a
`DeploymentRuntimeConfig` that attaches a `runtimeConfigRef` to the `Provider`.

### Option A: projected service-account token (recommended)

Create a ServiceAccount (the one whose `sub` your IAM mapping matches) and a
`DeploymentRuntimeConfig` that sets the provider pod's `serviceAccountName` and
mounts the projected token:

```yaml
apiVersion: v1
kind: ServiceAccount
metadata:
  name: oidc-token
  namespace: crossplane-system
---
apiVersion: pkg.crossplane.io/v1beta1
kind: DeploymentRuntimeConfig
metadata:
  name: cce-workload-identity
spec:
  deploymentTemplate:
    spec:
      selector: {}
      template:
        spec:
          serviceAccountName: oidc-token   # the SA whose sub the IAM mapping matches
          containers:
            - name: package-runtime
              volumeMounts:
                - name: cce-oidc-token
                  mountPath: /var/run/secrets/otc/tokens
                  readOnly: true
          volumes:
            - name: cce-oidc-token
              projected:
                sources:
                  - serviceAccountToken:
                      path: oidc-token
                      audience: crossplane-provider   # must equal the IAM client ID
                      expirationSeconds: 7200
```

Attach it to the provider:

```yaml
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-opentelekomcloud
spec:
  package: xpkg.upbound.io/opentelekomcloud/provider-opentelekomcloud:<version>
  runtimeConfigRef:
    name: cce-workload-identity
```

The provider pod must restart to pick up the new volume + service account.
Verify the volume is mounted:

```console
POD_NAME=$(kubectl get pods -n crossplane-system -l pkg.crossplane.io/revision \
  -o jsonpath='{.items[0].metadata.name}')
kubectl get pod "$POD_NAME" -n crossplane-system -o jsonpath='{.spec.volumes[*].name}'
# should list cce-oidc-token
```

### Option B: token from a Secret (external IdP)

For an external IdP, a controller (e.g. a CronJob) writes the ID token to a
Secret; no projected volume is required. To create it manually:

```console
kubectl create secret generic otc-oidc-token \
  --from-literal=token="<oidc-id-token>" \
  -n <provider-config-namespace>
```

## Step 4: Create the ProviderConfig

With the IdP, IAM federation and token all in place, create a namespaced
`ProviderConfig`.

### Option A: token from a projected file (recommended)

This is the primary path for OTC CCE: the token comes from the projected
service-account volume (Step 3, Option A).

```yaml
apiVersion: opentelekomcloud.m.crossplane.io/v1beta1
kind: ProviderConfig
metadata:
  name: webidentity
  namespace: test
spec:
  credentials:
    source: WebIdentity
    webIdentity:
      authURL: https://iam.eu-de.otc.t-systems.com   # IAM endpoint for your region
      idpID: crossplane-workload-identity            # the IAM provider name from Step 2.1
      region: eu-de
      scope:
        project:
          id: <project-id>                            # the project to act in
      tokenConfig:
        source: Filesystem
        fs:
          path: /var/run/secrets/otc/tokens/oidc-token
      temporaryCredentials:
        enabled: true
        durationSeconds: 3600
```

### Option B: token from a Secret (external IdP)

```yaml
apiVersion: opentelekomcloud.m.crossplane.io/v1beta1
kind: ProviderConfig
metadata:
  name: webidentity
  namespace: test
spec:
  credentials:
    source: WebIdentity
    webIdentity:
      authURL: https://iam.eu-de.otc.t-systems.com
      idpID: crossplane-workload-identity
      region: eu-de
      scope:
        project:
          id: <project-id>
      tokenConfig:
        source: Secret
        secretRef:
          name: otc-oidc-token
          key: token
      temporaryCredentials:
        enabled: true
        durationSeconds: 3600
```

#### External IdP notes

A few requirements are easy to miss when the token is minted by an external
IdP (Keycloak, dex, ...):

- **It must be an OIDC ID token, not an access token.** T Cloud IAM's
  `id-token` endpoint rejects an access token (`typ: Bearer`, which a
  client-credentials grant returns by default) with "id token is invalid".
  Request `scope=openid` so the response includes an `id_token` (`typ: ID`).
- **The token's `iss` must match the registered `provider_url`.** Mint the token
  over the IdP URL that you registered with T Cloud IAM (Step 2.1). If you mint
  over a different (e.g. in-cluster) URL, `iss` differs and IAM rejects the
  token.
- **The token's `aud` must match the `client_id`**, and the `sub` must match the
  mapping's `remote` rule (Step 2.2).

For an external IdP you typically need a controller to keep the Secret fresh
because the tokens are short-lived. A CronJob that performs the IdP token grant
and upserts the `id_token` into the Secret every minute is a common pattern.

### Field reference

| Field | Required | Description |
|-------|----------|-------------|
| `authURL` | yes | The base URL of the T Cloud IAM identity service for your region, e.g. `https://iam.eu-de.otc.t-systems.com`. A trailing `/v3` or `/v3.0` is optional and normalised. |
| `idpID` | yes | The name of the federated identity provider registered with T Cloud IAM (Step 2.1). Sent as the `X-Idp-Id` header. |
| `region` | yes | The T Cloud region to operate in, e.g. `eu-de`. |
| `scope` | no | Limits the resulting OTC token to a project or a domain. At most one of `project`/`domain`; within it exactly one of `id`/`name`. Omit for an unscoped token. **Prefer project scope** (see Limitations). |
| `tokenConfig.source` | yes | `Filesystem` (projected token, recommended) or `Secret`. |
| `tokenConfig.fs.path` | if `Filesystem` | Path to the OIDC ID token file inside the provider pod, e.g. `/var/run/secrets/otc/tokens/oidc-token`. |
| `tokenConfig.secretRef` | if `Secret` | The Secret holding the token. For a namespaced `ProviderConfig` the Secret is resolved in the managed resource's namespace, so set `name`/`key` (the `namespace` is filled in automatically). |
| `temporaryCredentials` | no | Optional. When present, `enabled` defaults to `true` and `durationSeconds` defaults to `3600` (range `900`–`86400`). Omit the object entirely for token-only auth (no temporary AK/SK). |
| `temporaryCredentials.enabled` | no | `true` (default when the object is present) or `false`. |
| `temporaryCredentials.durationSeconds` | no | Validity of the temporary AK/SK in seconds. `900`–`86400`; default `3600`. |

> [!NOTE]
> **When to enable `temporaryCredentials`.** The OTC token alone authenticates
> most services. Services that sign requests with AK/SK (e.g. **OBS**) require
> temporary credentials. Enable `temporaryCredentials` if you manage such
> services; otherwise you can omit it.

Reference the `ProviderConfig` from managed resources:

```yaml
spec:
  forProvider: { ... }
  providerConfigRef:
    kind: ProviderConfig
    name: webidentity
```

## Step 5: Verify the configuration

```console
# Is the ProviderConfig accepted?
kubectl get providerconfig.opentelekomcloud.m <provider-config-namespace>/webidentity -o yaml

# Create a trivial resource and watch it reconcile.
kubectl get vpcv1 <name> -n <ns> -o yaml
# Look for status.conditions Ready == True
```

A minimal smoke test for OTC CCE (a `VpcV1` referencing the `ProviderConfig`):

```yaml
apiVersion: vpc.opentelekomcloud.m.crossplane.io/v1alpha1
kind: VpcV1
metadata:
  name: vpc-cce-workload-identity
  namespace: test
spec:
  forProvider:
    cidr: 192.168.0.0/16
    name: vpc-cce-workload-identity
  providerConfigRef:
    name: webidentity
    kind: ProviderConfig
```

## Troubleshooting

Check the provider logs for the raw IAM error response:

```console
kubectl logs -n crossplane-system <provider-pod> -c package-runtime
```

| Symptom | Likely cause / fix |
|---------|--------------------|
| `401` / `InvalidIdentityToken` from IAM | `idpID` does not match the registered IAM provider name; or the token's `iss` does not match `provider_url`; or the `signing_key` is stale (key rotated). |
| `400` / `IAM.0007` / "Request parameter 'id token' is invalid" | The token was received but IAM could not verify it. Decode it (`cut -d'.' -f2 \| base64 -d`) and check: `iss` == `provider_url`, `aud` == `client_id`, `sub` matches the mapping, `exp` is in the future, and the `signing_key` JWKS matches `kubectl get --raw /openid/v1/jwks`. |
| `Incorrect token audience` | The token audience is not what the federation expects; fix the `serviceAccountToken.audience` (Option A) or the IdP client (Option B). |
| `403` after a successful token exchange | The mapped group lacks the required role in the target project/domain (Step 2.3), or the `scope` in the `ProviderConfig` does not match where the permission is granted. |
| Mapping does not match | The token's `sub`/claims do not match the `remote` rule in the mapping (Step 2.2). |
| `No suitable endpoint could be found in the service catalog` | A domain-scoped (or unscoped) token was used for a project-scoped resource. Use a project-scoped token (see Limitations). |
| Token file not found | The projected volume is not mounted, or `fs.path` is wrong. Verify the volume (Step 3) and the path. |
| Secret not found | `secretRef` name/key/namespace mismatch, or the token controller has not written it yet. For a namespaced `ProviderConfig`, the Secret is resolved in the managed resource's namespace. |
| Temp AK/SK missing for OBS | `temporaryCredentials` is omitted or `enabled: false`; add it. |
| Provider works but a service (e.g. OBS) fails | You configured token-only auth but the service needs AK/SK signing; enable `temporaryCredentials`. |

## Security best practices

- **No stored long-lived credentials.** The pod only holds a short-lived OIDC
  token; the OTC token and (optional) AK/SK are short-lived and re-issued on
  every reconcile. The token is re-read fresh each time and is never cached, so
  token rotation is always picked up.
- **Least privilege.** Scope the IAM roles granted to the mapped group to only
  what the provider needs, and prefer project scope over domain scope.
- **Rotate IdP keys carefully.** If you rotate the IdP signing key, update the
  IAM `signing_key` promptly, or token verification fails.
- **Prefer project scope** over domain scope for the token to limit blast radius.

## Limitations

- **Namespaced API group only.** `WebIdentity` is not supported on the legacy
  cluster-scoped `ClusterProviderConfig`.
- **Domain-scoped tokens cannot create project-scoped resources.** A
  domain-scoped token only carries domain-level endpoints in its service
  catalog, so creating a project-scoped resource (VPC, OBS bucket, ...) fails
  with `No suitable endpoint could be found in the service catalog`. **Use a
  project-scoped token** (`scope.project` with `id` or `name`) for deployment.
  A project-scoped token can still manage global services (e.g. OBS).
- **Temporary credentials are required for AK/SK-signing services** such as
  OBS. Enable `temporaryCredentials` if you manage those.
- **No credential caching.** Every reconcile of every managed resource re-exchanges
  the OIDC token with T Cloud IAM. This keeps the flow simple and always uses a
  fresh token, at the cost of extra IAM calls as the number of managed resources
  grows.

