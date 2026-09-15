# ECS test cases

## instance.yaml

Test case: `instance.yaml` - creates an ECS (Elastic Cloud Server) VM with a
keypair, VPC, subnet, and security group.

### APIs used

| Kind | Full API name |
|------|---------------|
| KeypairV2 | `keypairv2s.compute.opentelekomcloud.m.crossplane.io` |
| VpcV1 | `vpcv1s.vpc.opentelekomcloud.m.crossplane.io` |
| SubnetV1 | `subnetv1s.vpc.opentelekomcloud.m.crossplane.io` |
| SecgroupV2 | `secgroupv2s.compute.opentelekomcloud.m.crossplane.io` |
| InstanceV1 | `instancev1s.ecs.opentelekomcloud.m.crossplane.io` |

### Configuration needed

| Item | Where | What to do |
|------|-------|-----------|
| `${Rand.RFC1123Subdomain}` | names of KeypairV2, VpcV1, SubnetV1, SecgroupV2, InstanceV1 | Replace with a unique RFC1123-compatible suffix (or leave for the example generator to substitute) |
| `publicKey` | KeypairV2 `spec.forProvider.publicKey` | Replace the hardcoded RSA public key with your own SSH public key |
| `imageId: 54c617ad-5c85-4426-b721-494a9c7719c2` | InstanceV1 `spec.forProvider.imageId` | Replace with the ID of an image that exists in your region |
| `flavor: s2.large.2` | InstanceV1 `spec.forProvider.flavor` | Adjust to a flavor available in your region |
| `availabilityZone: eu-de-01` | InstanceV1 | Change to an AZ in your region |
| `security group rules` (ports 22/8080 from 0.0.0.0/0) | SecgroupV2 `spec.forProvider.rule` | Restrict the CIDR and ports to what you actually need |
| `namespace: test` | all resources | All resources are created in the `test` namespace - adjust if needed |
