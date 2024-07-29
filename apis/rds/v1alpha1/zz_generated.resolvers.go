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
	rds "github.com/opentelekomcloud/provider-opentelekomcloud/config/rds"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

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
		Extract:       rds.ExtractEipAddress(),
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
		Extract:      rds.ExtractNetworkID(),
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

	return nil
}