/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	v1alpha1 "github.com/opentelekomcloud/provider-opentelekomcloud/apis/compute/v1alpha1"
	v1alpha11 "github.com/opentelekomcloud/provider-opentelekomcloud/apis/vpc/v1alpha1"
	compute "github.com/opentelekomcloud/provider-opentelekomcloud/config/compute"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this InstanceV1.
func (mg *InstanceV1) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KeyName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.KeyNameRef,
		Selector:     mg.Spec.ForProvider.KeyNameSelector,
		To: reference.To{
			List:    &v1alpha1.KeypairV2List{},
			Managed: &v1alpha1.KeypairV2{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.KeyName")
	}
	mg.Spec.ForProvider.KeyName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.KeyNameRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Nics); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Nics[i3].NetworkID),
			Extract:      compute.ExtractNetworkID(),
			Reference:    mg.Spec.ForProvider.Nics[i3].NetworkIDRef,
			Selector:     mg.Spec.ForProvider.Nics[i3].NetworkIDSelector,
			To: reference.To{
				List:    &v1alpha11.SubnetV1List{},
				Managed: &v1alpha11.SubnetV1{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Nics[i3].NetworkID")
		}
		mg.Spec.ForProvider.Nics[i3].NetworkID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Nics[i3].NetworkIDRef = rsp.ResolvedReference

	}
	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.SecurityGroups),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.ComputeSecurityGroupIDRefs,
		Selector:      mg.Spec.ForProvider.ComputeSecurityGroupIDSelector,
		To: reference.To{
			List:    &v1alpha1.SecgroupV2List{},
			Managed: &v1alpha1.SecgroupV2{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SecurityGroups")
	}
	mg.Spec.ForProvider.SecurityGroups = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.ComputeSecurityGroupIDRefs = mrsp.ResolvedReferences

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VPCID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.VPCIDRef,
		Selector:     mg.Spec.ForProvider.VPCIDSelector,
		To: reference.To{
			List:    &v1alpha11.VpcV1List{},
			Managed: &v1alpha11.VpcV1{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VPCID")
	}
	mg.Spec.ForProvider.VPCID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VPCIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.KeyName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.KeyNameRef,
		Selector:     mg.Spec.InitProvider.KeyNameSelector,
		To: reference.To{
			List:    &v1alpha1.KeypairV2List{},
			Managed: &v1alpha1.KeypairV2{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.KeyName")
	}
	mg.Spec.InitProvider.KeyName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.KeyNameRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.Nics); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Nics[i3].NetworkID),
			Extract:      compute.ExtractNetworkID(),
			Reference:    mg.Spec.InitProvider.Nics[i3].NetworkIDRef,
			Selector:     mg.Spec.InitProvider.Nics[i3].NetworkIDSelector,
			To: reference.To{
				List:    &v1alpha11.SubnetV1List{},
				Managed: &v1alpha11.SubnetV1{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.Nics[i3].NetworkID")
		}
		mg.Spec.InitProvider.Nics[i3].NetworkID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.Nics[i3].NetworkIDRef = rsp.ResolvedReference

	}
	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.SecurityGroups),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.InitProvider.ComputeSecurityGroupIDRefs,
		Selector:      mg.Spec.InitProvider.ComputeSecurityGroupIDSelector,
		To: reference.To{
			List:    &v1alpha1.SecgroupV2List{},
			Managed: &v1alpha1.SecgroupV2{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.SecurityGroups")
	}
	mg.Spec.InitProvider.SecurityGroups = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.InitProvider.ComputeSecurityGroupIDRefs = mrsp.ResolvedReferences

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.VPCID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.VPCIDRef,
		Selector:     mg.Spec.InitProvider.VPCIDSelector,
		To: reference.To{
			List:    &v1alpha11.VpcV1List{},
			Managed: &v1alpha11.VpcV1{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.VPCID")
	}
	mg.Spec.InitProvider.VPCID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.VPCIDRef = rsp.ResolvedReference

	return nil
}
