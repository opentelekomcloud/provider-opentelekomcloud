# How to contribute to the project
For a general guide, please read the [Upjet Contribution Guide](https://github.com/crossplane/upjet/blob/main/CONTRIBUTING.md).

## Fixing bugs
If you’d like to help by fixing open bugs, we’d greatly appreciate it! Bug fixing is a vital part of keeping the project healthy. The best place to start is our [bug tracker](https://github.com/opentelekomcloud/provider-opentelekomcloud/issues?q=is%3Aissue%20state%3Aopen%20label%3Abug), where you can browse currently open issues and pick one that isn’t actively being worked on. Before diving in, it’s often a good idea to leave a comment on the issue to confirm its current status and check whether anyone has already made progress or discovered anything useful.

## Adding new resources
The [OpenTelekomCloud Terraform provider](https://github.com/opentelekomcloud/terraform-provider-opentelekomcloud) exposes many resources, but not all of them are properly configured in our Crossplane provider. You can track the implementation status of OTC services in this [issue](https://github.com/opentelekomcloud/provider-opentelekomcloud/issues/7).  

If you have a use case for a service that is missing or only partially implemented, feel free to open an issue describing your needs and explaining why the service is important to you. If you’d like to contribute the implementation yourself, we recommend reviewing [Adding a new resouce](https://github.com/crossplane/upjet/blob/main/docs/adding-new-resource.md), [configure a resource](https://github.com/crossplane/upjet/blob/main/docs/configuring-a-resource.md), and this example [PR](https://github.com/opentelekomcloud/provider-opentelekomcloud/pull/33) to help you get started.


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

### Review changes
> [!NOTE]
> Small code changes can lead to large diffs. `scripts/diff_helper.sh` can help you ignoring generated files.

## Adding test cases
[TBD](https://github.com/opentelekomcloud/provider-opentelekomcloud/issues/45)
