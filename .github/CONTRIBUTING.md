# How to contribute to the project
For a general guide, please read the [Upjet Contribution Guide](https://github.com/crossplane/upjet/blob/main/CONTRIBUTING.md).

## Fixing bugs
If you’d like to help by fixing open bugs, we’d greatly appreciate it! Bug fixing is a vital part of keeping the project healthy. The best place to start is our [bug tracker](https://github.com/opentelekomcloud/provider-opentelekomcloud/issues?q=is%3Aissue%20state%3Aopen%20label%3Abug), where you can browse currently open issues and pick one that isn’t actively being worked on. Before diving in, it’s often a good idea to leave a comment on the issue to confirm its current status and check whether anyone has already made progress or discovered anything useful.

## Adding new resources
The [OpenTelekomCloud Terraform provider](https://github.com/opentelekomcloud/terraform-provider-opentelekomcloud) exposes many resources, but not all of them are properly configured in our Crossplane provider. You can track the implementation status of OpenTelekomCloud services in this [issue](https://github.com/opentelekomcloud/provider-opentelekomcloud/issues/7).  

If you have a use case for a service that is missing or only partially implemented, feel free to open an issue describing your needs and explaining why the service is important to you. If you’d like to contribute the implementation yourself, we recommend reviewing [Adding a new resource](https://github.com/crossplane/upjet/blob/main/docs/adding-new-resource.md), [Configure a resource](https://github.com/crossplane/upjet/blob/main/docs/configuring-a-resource.md), and this example [PR](https://github.com/opentelekomcloud/provider-opentelekomcloud/pull/33) to help you get started.


## Building and running the provider

### After cloning the repository fetch the [crossplane/build](https://github.com/crossplane/build/) submodule
```console
make submodules
```

### Make your changes in the code...

### Run the code-generation pipeline

```console
make generate
```

### Apply generated CRDS
```yaml
kubectl apply -f package/crds
```

### Set up credentials secret and ProviderConfig for AUTH
> [!NOTE]
> This is one way of creating a `ProviderConfig`, please check [provider-configuration](https://docs.crossplane.io/latest/packages/providers/#provider-configuration) for more options.
```console
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Secret
metadata:
  name: provider-opentelekomcloud-creds
  namespace: app
type: Opaque
stringData:
  credentials: |
    {
      "user_name": "admin",
      "password": "t0ps3cr3t11",
      "auth_url": "https://iam.eu-de.otc.t-systems.com/v3",
      "domain_name": "OTCxxxx",
      "tenant_name": "eu-de_project",
      "swauth": "false",
      "allow_reauth": "true",
      "max_retries": "2",
      "max_backoff_retries": "6",
      "backoff_retry_timeout": "60",
      "insecure": "false"
    }
---
apiVersion: opentelekomcloud.m.crossplane.io/v1beta1
kind: ProviderConfig
metadata:
  name: default
  namespace: app
spec:
  credentials:
    source: Secret
    secretRef:
      name: provider-opentelekomcloud-creds
      namespace: crossplane-system
      key: credentials
EOF
```

### Run against a Kubernetes cluster
> [!NOTE]
> For this to work you need to have crossplane core already [deployed](https://docs.crossplane.io/latest/get-started/install/) in your cluster. Crossplane will use your locally running provider as it's OpenTelekomCloud operator.
```console
make run
```


### Build binary

```console
make build
```

### Troubleshooting
To enable debug logging mode you have to use [DeploymentRuntimeConfig](https://docs.crossplane.io/latest/packages/providers/#runtime-configuration), which is a new feature in v2 Crossplane. In legacy versions you have to use `ControllerConfig`. You can check the available flags in `cmd/provider/main.go`.
```console
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1beta1
kind: DeploymentRuntimeConfig
metadata:
  name: enable-debug
spec:
  deploymentTemplate:
    spec:
      selector: {}
      template:
        spec:
          containers:
            - name: package-runtime
              args:
                - --debug
EOF
```

### Review changes
> [!NOTE]
> Small code changes can lead to large diffs. `scripts/diff_helper.sh` can help you ignoring generated files.

## Integration testing
`/examples/` are used for automated tests utilizing [uptest](https://github.com/crossplane/upjet/blob/main/docs/testing-with-uptest.md) framework. If you contribute new resources, please make sure to include corresponding test cases. These tests create real resources and can be expensive so they are not started on each PR. Maintainers can trigger these tests by posting specific comments in the PR:
- `/test-examples="examples/namespaced/swr/repositoryv2.yaml"` (single resource)
- `/test-examples="examples/namespaced/swr/repositoryv2.yaml, examples/namespaced/swr/organizationsv2.yaml"` (multiple resources)
- `/test-examples="examples/namespaced/swr/"` (all resources in directory)

### Local testing
For faster feedback loop you can use the e2e make target to run the tests locally. 

1. Create credentails
```console
export UPTEST_CLOUD_CREDENTIALS={"user_name":"some_user","password":"some_password","auth_url":"https://iam.eu-de.otc.t-systems.com/v3","domain_name":"OTCXXXX","tenant_name":"eu-de_project","swauth":"false","allow_reauth":"true","max_retries":"2","max_backoff_retries":"6","backoff_retry_timeout":"60","insecure":"false"}
```

2. Trigger testing for a specific directory
> [!NOTE]
> Uptest default timeout value is 20m, unless you are testing resources which take a long time to create you should use lower timeout.
```console
UPTEST_DEFAULT_TIMEOUT=5m UPTEST_EXAMPLE_LIST=$(find examples/namespaced/swr/*.yaml | tr '\n' ',') make e2e
```

This will do the following:
- Builds provider from source
- Creates local cluster using kind
- Deploys crossplane core
- Deploys provider
- Configures provider credentials
- Configures ClusterProviderConfig
- Renders chainsaw manifests
- Runs tests
- Cleans up
