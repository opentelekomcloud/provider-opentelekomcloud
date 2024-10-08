/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	v1alpha11 "github.com/opentelekomcloud/provider-opentelekomcloud/apis/compute/v1alpha1"
	v1alpha1 "github.com/opentelekomcloud/provider-opentelekomcloud/apis/vpc/v1alpha1"
	common "github.com/opentelekomcloud/provider-opentelekomcloud/config/common"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this BackupV3.
func (mg *BackupV3) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InstanceID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.InstanceIDRef,
		Selector:     mg.Spec.ForProvider.InstanceIDSelector,
		To: reference.To{
			List:    &InstanceV3List{},
			Managed: &InstanceV3{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InstanceID")
	}
	mg.Spec.ForProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.InstanceID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.InstanceIDRef,
		Selector:     mg.Spec.InitProvider.InstanceIDSelector,
		To: reference.To{
			List:    &InstanceV3List{},
			Managed: &InstanceV3{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.InstanceID")
	}
	mg.Spec.InitProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.InstanceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this InstanceV3.
func (mg *InstanceV3) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ParamGroupID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ParamGroupIDRef,
		Selector:     mg.Spec.ForProvider.ParamGroupIDSelector,
		To: reference.To{
			List:    &ParametergroupV3List{},
			Managed: &ParametergroupV3{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ParamGroupID")
	}
	mg.Spec.ForProvider.ParamGroupID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ParamGroupIDRef = rsp.ResolvedReference

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.PublicIps),
		Extract:       common.ExtractEipAddress(),
		References:    mg.Spec.ForProvider.PublicIpsRefs,
		Selector:      mg.Spec.ForProvider.PublicIpsSelector,
		To: reference.To{
			List:    &v1alpha1.EIPV1List{},
			Managed: &v1alpha1.EIPV1{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PublicIps")
	}
	mg.Spec.ForProvider.PublicIps = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.PublicIpsRefs = mrsp.ResolvedReferences

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SecurityGroupID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ComputeSecurityGroupIDRefs,
		Selector:     mg.Spec.ForProvider.ComputeSecurityGroupIDSelector,
		To: reference.To{
			List:    &v1alpha11.SecgroupV2List{},
			Managed: &v1alpha11.SecgroupV2{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SecurityGroupID")
	}
	mg.Spec.ForProvider.SecurityGroupID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ComputeSecurityGroupIDRefs = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SubnetID),
		Extract:      common.ExtractNetworkID(),
		Reference:    mg.Spec.ForProvider.SubnetIDRef,
		Selector:     mg.Spec.ForProvider.SubnetIDSelector,
		To: reference.To{
			List:    &v1alpha1.SubnetV1List{},
			Managed: &v1alpha1.SubnetV1{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SubnetID")
	}
	mg.Spec.ForProvider.SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SubnetIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VPCID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.VPCIDRef,
		Selector:     mg.Spec.ForProvider.VPCIDSelector,
		To: reference.To{
			List:    &v1alpha1.VpcV1List{},
			Managed: &v1alpha1.VpcV1{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VPCID")
	}
	mg.Spec.ForProvider.VPCID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VPCIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ParamGroupID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.ParamGroupIDRef,
		Selector:     mg.Spec.InitProvider.ParamGroupIDSelector,
		To: reference.To{
			List:    &ParametergroupV3List{},
			Managed: &ParametergroupV3{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ParamGroupID")
	}
	mg.Spec.InitProvider.ParamGroupID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ParamGroupIDRef = rsp.ResolvedReference

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.PublicIps),
		Extract:       common.ExtractEipAddress(),
		References:    mg.Spec.InitProvider.PublicIpsRefs,
		Selector:      mg.Spec.InitProvider.PublicIpsSelector,
		To: reference.To{
			List:    &v1alpha1.EIPV1List{},
			Managed: &v1alpha1.EIPV1{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.PublicIps")
	}
	mg.Spec.InitProvider.PublicIps = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.InitProvider.PublicIpsRefs = mrsp.ResolvedReferences

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.SecurityGroupID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.ComputeSecurityGroupIDRefs,
		Selector:     mg.Spec.InitProvider.ComputeSecurityGroupIDSelector,
		To: reference.To{
			List:    &v1alpha11.SecgroupV2List{},
			Managed: &v1alpha11.SecgroupV2{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.SecurityGroupID")
	}
	mg.Spec.InitProvider.SecurityGroupID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ComputeSecurityGroupIDRefs = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.SubnetID),
		Extract:      common.ExtractNetworkID(),
		Reference:    mg.Spec.InitProvider.SubnetIDRef,
		Selector:     mg.Spec.InitProvider.SubnetIDSelector,
		To: reference.To{
			List:    &v1alpha1.SubnetV1List{},
			Managed: &v1alpha1.SubnetV1{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.SubnetID")
	}
	mg.Spec.InitProvider.SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.SubnetIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.VPCID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.VPCIDRef,
		Selector:     mg.Spec.InitProvider.VPCIDSelector,
		To: reference.To{
			List:    &v1alpha1.VpcV1List{},
			Managed: &v1alpha1.VpcV1{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.VPCID")
	}
	mg.Spec.InitProvider.VPCID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.VPCIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ReadReplicaV3.
func (mg *ReadReplicaV3) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.PublicIps),
		Extract:       common.ExtractEipAddress(),
		References:    mg.Spec.ForProvider.PublicIpsRefs,
		Selector:      mg.Spec.ForProvider.PublicIpsSelector,
		To: reference.To{
			List:    &v1alpha1.EIPV1List{},
			Managed: &v1alpha1.EIPV1{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PublicIps")
	}
	mg.Spec.ForProvider.PublicIps = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.PublicIpsRefs = mrsp.ResolvedReferences

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ReplicaOfID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ReplicaOfIDRef,
		Selector:     mg.Spec.ForProvider.ReplicaOfIDSelector,
		To: reference.To{
			List:    &InstanceV3List{},
			Managed: &InstanceV3{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ReplicaOfID")
	}
	mg.Spec.ForProvider.ReplicaOfID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ReplicaOfIDRef = rsp.ResolvedReference

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.PublicIps),
		Extract:       common.ExtractEipAddress(),
		References:    mg.Spec.InitProvider.PublicIpsRefs,
		Selector:      mg.Spec.InitProvider.PublicIpsSelector,
		To: reference.To{
			List:    &v1alpha1.EIPV1List{},
			Managed: &v1alpha1.EIPV1{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.PublicIps")
	}
	mg.Spec.InitProvider.PublicIps = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.InitProvider.PublicIpsRefs = mrsp.ResolvedReferences

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ReplicaOfID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.ReplicaOfIDRef,
		Selector:     mg.Spec.InitProvider.ReplicaOfIDSelector,
		To: reference.To{
			List:    &InstanceV3List{},
			Managed: &InstanceV3{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ReplicaOfID")
	}
	mg.Spec.InitProvider.ReplicaOfID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ReplicaOfIDRef = rsp.ResolvedReference

	return nil
}
