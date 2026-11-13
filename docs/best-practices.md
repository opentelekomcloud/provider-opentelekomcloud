# Best practices

This document collects recommended practices for using `provider-opentelekomcloud`. It is a living document and will be extended as new guidance emerges.

## Prefer namespaced APIs over cluster-scoped APIs

Use the namespaced API group `*.opentelekomcloud.m.crossplane.io`. The cluster-scoped group `*.opentelekomcloud.crossplane.io` is legacy and should not be used for new deployments.

## Only load the controllers you need

Installing the provider with default settings activates every API and controller, which adds hundreds of unused CRDs and controllers to `kube-apiserver`. Restrict activation to the resources you actually use with a `ManagedResourceActivationPolicy` (MRAP). See [Managed Resource Activation Policies](mrap.md) for examples, including the recommended `provider.defaultActivations={"*.opentelekomcloud.m.crossplane.io"}` installation.

## Prefer VPC resources over Networking resources

T Cloud Public exposes two API families for network resources:

| API family | API group | Terraform prefix | Recommendation |
|------------|-----------|------------------|----------------|
| VPC | `vpc.opentelekomcloud.m.crossplane.io` | `opentelekomcloud_vpc_*` | Use for new deployments |
| Networking (Neutron) | `networking.opentelekomcloud.m.crossplane.io` | `opentelekomcloud_networking_*` | Use only for resources without a VPC equivalent |

When a VPC equivalent exists, use it:

| Recommended VPC resource | Equivalent Networking resource |
|--------------------------|--------------------------------|
| `vpcv1s.vpc.opentelekomcloud.m.crossplane.io` | `networkv2s.networking.opentelekomcloud.m.crossplane.io` |
| `subnetv1s.vpc.opentelekomcloud.m.crossplane.io` | `subnetv2s.networking.opentelekomcloud.m.crossplane.io` |
| `secgroupv3s.vpc.opentelekomcloud.m.crossplane.io` | `secgroupv2s.networking.opentelekomcloud.m.crossplane.io` |
| `secgrouprulev3s.vpc.opentelekomcloud.m.crossplane.io` | `secgrouprulev2s.networking.opentelekomcloud.m.crossplane.io` |
| `eipv1s.vpc.opentelekomcloud.m.crossplane.io` | `floatingipv2s.networking.opentelekomcloud.m.crossplane.io` |

The VPC API is the higher-level, actively developed API family. Newer capabilities are only available there, for example EIPs with bandwidth, route tables, flow logs, VPC peering, secondary CIDRs, IP address groups, and security groups v3. VPC resources also expose the OpenStack IDs that other services require, so they can usually be used even when the consuming service expects a Neutron ID.

Use the Networking family only for constructs that have no VPC equivalent:

- ports and port security-group associations
- routers, router interfaces, and router routes
- virtual IPs (VIPs)
- floating IP associations

Do not manage the same cloud object with both a Networking and a VPC managed resource. Pick one owner.

## Importing resources created with OpenStack APIs

Resources created with the native OpenStack/Neutron API, for example with the `openstack` CLI, Heat, or the Terraform `opentelekomcloud_networking_*` resources, are the same underlying resources that the VPC API manages. You can adopt them with the corresponding VPC managed resource instead of recreating them.

The import flow is the same as for any other managed resource:

1. Create the VPC managed resource and set the `crossplane.io/external-name` annotation to the ID of the existing resource.
2. Start with `managementPolicies: ["Observe"]` and wait until the resource is `Ready`.
3. Change `managementPolicies` to `["*"]` to take control.

See [Import Existing Resources](import-external-resources.md) for the full walk-through.

After import, VPC resources expose the related OpenStack IDs when they differ. For example, `SubnetV1` reports `networkId` and `subnetId` in `status.atProvider`. Use those values when another resource needs the Neutron IDs instead of creating a Networking managed resource.

## Further reading

- [Import Existing Resources](import-external-resources.md)
- [Configure the provider](configure-the-provider.md)
- [Managed Resource Activation Policies](mrap.md)
- [Upgrading the provider](upgrade-the-provider.md)
