/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	v1alpha13 "github.com/opentelekomcloud/provider-opentelekomcloud/apis/compute/v1alpha1"
	v1alpha11 "github.com/opentelekomcloud/provider-opentelekomcloud/apis/identity/v1alpha1"
	v1alpha12 "github.com/opentelekomcloud/provider-opentelekomcloud/apis/kms/v1alpha1"
	v1alpha1 "github.com/opentelekomcloud/provider-opentelekomcloud/apis/vpc/v1alpha1"
	common "github.com/opentelekomcloud/provider-opentelekomcloud/config/common"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this AddonV3.
func (mg *AddonV3) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ClusterID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ClusterIDRef,
		Selector:     mg.Spec.ForProvider.ClusterIDSelector,
		To: reference.To{
			List:    &ClusterV3List{},
			Managed: &ClusterV3{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ClusterID")
	}
	mg.Spec.ForProvider.ClusterID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClusterIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ClusterID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.ClusterIDRef,
		Selector:     mg.Spec.InitProvider.ClusterIDSelector,
		To: reference.To{
			List:    &ClusterV3List{},
			Managed: &ClusterV3{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ClusterID")
	}
	mg.Spec.InitProvider.ClusterID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ClusterIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ClusterV3.
func (mg *ClusterV3) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EIP),
		Extract:      common.ExtractEipAddress(),
		Reference:    mg.Spec.ForProvider.EIPRef,
		Selector:     mg.Spec.ForProvider.EIPSelector,
		To: reference.To{
			List:    &v1alpha1.EIPV1List{},
			Managed: &v1alpha1.EIPV1{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.EIP")
	}
	mg.Spec.ForProvider.EIP = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.EIPRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EniSubnetCidr),
		Extract:      common.ExtractSubnetCIDR(),
		Reference:    mg.Spec.ForProvider.EniSubnetCidrRef,
		Selector:     mg.Spec.ForProvider.EniSubnetCidrSelector,
		To: reference.To{
			List:    &v1alpha1.SubnetV1List{},
			Managed: &v1alpha1.SubnetV1{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.EniSubnetCidr")
	}
	mg.Spec.ForProvider.EniSubnetCidr = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.EniSubnetCidrRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EniSubnetID),
		Extract:      common.ExtractSubnetID(),
		Reference:    mg.Spec.ForProvider.EniSubnetIDRef,
		Selector:     mg.Spec.ForProvider.EniSubnetIDSelector,
		To: reference.To{
			List:    &v1alpha1.SubnetV1List{},
			Managed: &v1alpha1.SubnetV1{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.EniSubnetID")
	}
	mg.Spec.ForProvider.EniSubnetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.EniSubnetIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.HighwaySubnetID),
		Extract:      common.ExtractNetworkID(),
		Reference:    mg.Spec.ForProvider.HighwaySubnetIDRef,
		Selector:     mg.Spec.ForProvider.HighwaySubnetIDSelector,
		To: reference.To{
			List:    &v1alpha1.SubnetV1List{},
			Managed: &v1alpha1.SubnetV1{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.HighwaySubnetID")
	}
	mg.Spec.ForProvider.HighwaySubnetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.HighwaySubnetIDRef = rsp.ResolvedReference

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
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.EIP),
		Extract:      common.ExtractEipAddress(),
		Reference:    mg.Spec.InitProvider.EIPRef,
		Selector:     mg.Spec.InitProvider.EIPSelector,
		To: reference.To{
			List:    &v1alpha1.EIPV1List{},
			Managed: &v1alpha1.EIPV1{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.EIP")
	}
	mg.Spec.InitProvider.EIP = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.EIPRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.EniSubnetCidr),
		Extract:      common.ExtractSubnetCIDR(),
		Reference:    mg.Spec.InitProvider.EniSubnetCidrRef,
		Selector:     mg.Spec.InitProvider.EniSubnetCidrSelector,
		To: reference.To{
			List:    &v1alpha1.SubnetV1List{},
			Managed: &v1alpha1.SubnetV1{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.EniSubnetCidr")
	}
	mg.Spec.InitProvider.EniSubnetCidr = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.EniSubnetCidrRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.EniSubnetID),
		Extract:      common.ExtractSubnetID(),
		Reference:    mg.Spec.InitProvider.EniSubnetIDRef,
		Selector:     mg.Spec.InitProvider.EniSubnetIDSelector,
		To: reference.To{
			List:    &v1alpha1.SubnetV1List{},
			Managed: &v1alpha1.SubnetV1{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.EniSubnetID")
	}
	mg.Spec.InitProvider.EniSubnetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.EniSubnetIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.HighwaySubnetID),
		Extract:      common.ExtractNetworkID(),
		Reference:    mg.Spec.InitProvider.HighwaySubnetIDRef,
		Selector:     mg.Spec.InitProvider.HighwaySubnetIDSelector,
		To: reference.To{
			List:    &v1alpha1.SubnetV1List{},
			Managed: &v1alpha1.SubnetV1{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.HighwaySubnetID")
	}
	mg.Spec.InitProvider.HighwaySubnetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.HighwaySubnetIDRef = rsp.ResolvedReference

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

// ResolveReferences of this NodePoolV3.
func (mg *NodePoolV3) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AgencyName),
		Extract:      common.ExtractAgencyName(),
		Reference:    mg.Spec.ForProvider.AgencyNameRef,
		Selector:     mg.Spec.ForProvider.AgencyNameSelector,
		To: reference.To{
			List:    &v1alpha11.AgencyV3List{},
			Managed: &v1alpha11.AgencyV3{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.AgencyName")
	}
	mg.Spec.ForProvider.AgencyName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.AgencyNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ClusterID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ClusterIDRef,
		Selector:     mg.Spec.ForProvider.ClusterIDSelector,
		To: reference.To{
			List:    &ClusterV3List{},
			Managed: &ClusterV3{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ClusterID")
	}
	mg.Spec.ForProvider.ClusterID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClusterIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.DataVolumes); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DataVolumes[i3].KMSID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.DataVolumes[i3].KMSIDRef,
			Selector:     mg.Spec.ForProvider.DataVolumes[i3].KMSIDSelector,
			To: reference.To{
				List:    &v1alpha12.KeyV1List{},
				Managed: &v1alpha12.KeyV1{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.DataVolumes[i3].KMSID")
		}
		mg.Spec.ForProvider.DataVolumes[i3].KMSID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.DataVolumes[i3].KMSIDRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KeyPair),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.KeyPairRef,
		Selector:     mg.Spec.ForProvider.KeyPairSelector,
		To: reference.To{
			List:    &v1alpha13.KeypairV2List{},
			Managed: &v1alpha13.KeypairV2{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.KeyPair")
	}
	mg.Spec.ForProvider.KeyPair = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.KeyPairRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.RootVolume); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RootVolume[i3].KMSID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.RootVolume[i3].KMSIDRef,
			Selector:     mg.Spec.ForProvider.RootVolume[i3].KMSIDSelector,
			To: reference.To{
				List:    &v1alpha12.KeyV1List{},
				Managed: &v1alpha12.KeyV1{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.RootVolume[i3].KMSID")
		}
		mg.Spec.ForProvider.RootVolume[i3].KMSID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.RootVolume[i3].KMSIDRef = rsp.ResolvedReference

	}
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
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.AgencyName),
		Extract:      common.ExtractAgencyName(),
		Reference:    mg.Spec.InitProvider.AgencyNameRef,
		Selector:     mg.Spec.InitProvider.AgencyNameSelector,
		To: reference.To{
			List:    &v1alpha11.AgencyV3List{},
			Managed: &v1alpha11.AgencyV3{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.AgencyName")
	}
	mg.Spec.InitProvider.AgencyName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.AgencyNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ClusterID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.ClusterIDRef,
		Selector:     mg.Spec.InitProvider.ClusterIDSelector,
		To: reference.To{
			List:    &ClusterV3List{},
			Managed: &ClusterV3{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ClusterID")
	}
	mg.Spec.InitProvider.ClusterID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ClusterIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.DataVolumes); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.DataVolumes[i3].KMSID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.DataVolumes[i3].KMSIDRef,
			Selector:     mg.Spec.InitProvider.DataVolumes[i3].KMSIDSelector,
			To: reference.To{
				List:    &v1alpha12.KeyV1List{},
				Managed: &v1alpha12.KeyV1{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.DataVolumes[i3].KMSID")
		}
		mg.Spec.InitProvider.DataVolumes[i3].KMSID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.DataVolumes[i3].KMSIDRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.KeyPair),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.KeyPairRef,
		Selector:     mg.Spec.InitProvider.KeyPairSelector,
		To: reference.To{
			List:    &v1alpha13.KeypairV2List{},
			Managed: &v1alpha13.KeypairV2{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.KeyPair")
	}
	mg.Spec.InitProvider.KeyPair = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.KeyPairRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.RootVolume); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RootVolume[i3].KMSID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.RootVolume[i3].KMSIDRef,
			Selector:     mg.Spec.InitProvider.RootVolume[i3].KMSIDSelector,
			To: reference.To{
				List:    &v1alpha12.KeyV1List{},
				Managed: &v1alpha12.KeyV1{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.RootVolume[i3].KMSID")
		}
		mg.Spec.InitProvider.RootVolume[i3].KMSID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.RootVolume[i3].KMSIDRef = rsp.ResolvedReference

	}
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

	return nil
}

// ResolveReferences of this NodeV3.
func (mg *NodeV3) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AgencyName),
		Extract:      common.ExtractAgencyName(),
		Reference:    mg.Spec.ForProvider.AgencyNameRef,
		Selector:     mg.Spec.ForProvider.AgencyNameSelector,
		To: reference.To{
			List:    &v1alpha11.AgencyV3List{},
			Managed: &v1alpha11.AgencyV3{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.AgencyName")
	}
	mg.Spec.ForProvider.AgencyName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.AgencyNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ClusterID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ClusterIDRef,
		Selector:     mg.Spec.ForProvider.ClusterIDSelector,
		To: reference.To{
			List:    &ClusterV3List{},
			Managed: &ClusterV3{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ClusterID")
	}
	mg.Spec.ForProvider.ClusterID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClusterIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.DataVolumes); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DataVolumes[i3].KMSID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.DataVolumes[i3].KMSIDRef,
			Selector:     mg.Spec.ForProvider.DataVolumes[i3].KMSIDSelector,
			To: reference.To{
				List:    &v1alpha12.KeyV1List{},
				Managed: &v1alpha12.KeyV1{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.DataVolumes[i3].KMSID")
		}
		mg.Spec.ForProvider.DataVolumes[i3].KMSID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.DataVolumes[i3].KMSIDRef = rsp.ResolvedReference

	}
	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.EIPIds),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.EIPIdsRefs,
		Selector:      mg.Spec.ForProvider.EIPIdsSelector,
		To: reference.To{
			List:    &v1alpha1.EIPV1List{},
			Managed: &v1alpha1.EIPV1{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.EIPIds")
	}
	mg.Spec.ForProvider.EIPIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.EIPIdsRefs = mrsp.ResolvedReferences

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KeyPair),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.KeyPairRef,
		Selector:     mg.Spec.ForProvider.KeyPairSelector,
		To: reference.To{
			List:    &v1alpha13.KeypairV2List{},
			Managed: &v1alpha13.KeypairV2{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.KeyPair")
	}
	mg.Spec.ForProvider.KeyPair = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.KeyPairRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.RootVolume); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RootVolume[i3].KMSID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.RootVolume[i3].KMSIDRef,
			Selector:     mg.Spec.ForProvider.RootVolume[i3].KMSIDSelector,
			To: reference.To{
				List:    &v1alpha12.KeyV1List{},
				Managed: &v1alpha12.KeyV1{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.RootVolume[i3].KMSID")
		}
		mg.Spec.ForProvider.RootVolume[i3].KMSID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.RootVolume[i3].KMSIDRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.AgencyName),
		Extract:      common.ExtractAgencyName(),
		Reference:    mg.Spec.InitProvider.AgencyNameRef,
		Selector:     mg.Spec.InitProvider.AgencyNameSelector,
		To: reference.To{
			List:    &v1alpha11.AgencyV3List{},
			Managed: &v1alpha11.AgencyV3{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.AgencyName")
	}
	mg.Spec.InitProvider.AgencyName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.AgencyNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ClusterID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.ClusterIDRef,
		Selector:     mg.Spec.InitProvider.ClusterIDSelector,
		To: reference.To{
			List:    &ClusterV3List{},
			Managed: &ClusterV3{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ClusterID")
	}
	mg.Spec.InitProvider.ClusterID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ClusterIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.DataVolumes); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.DataVolumes[i3].KMSID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.DataVolumes[i3].KMSIDRef,
			Selector:     mg.Spec.InitProvider.DataVolumes[i3].KMSIDSelector,
			To: reference.To{
				List:    &v1alpha12.KeyV1List{},
				Managed: &v1alpha12.KeyV1{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.DataVolumes[i3].KMSID")
		}
		mg.Spec.InitProvider.DataVolumes[i3].KMSID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.DataVolumes[i3].KMSIDRef = rsp.ResolvedReference

	}
	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.EIPIds),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.InitProvider.EIPIdsRefs,
		Selector:      mg.Spec.InitProvider.EIPIdsSelector,
		To: reference.To{
			List:    &v1alpha1.EIPV1List{},
			Managed: &v1alpha1.EIPV1{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.EIPIds")
	}
	mg.Spec.InitProvider.EIPIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.InitProvider.EIPIdsRefs = mrsp.ResolvedReferences

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.KeyPair),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.KeyPairRef,
		Selector:     mg.Spec.InitProvider.KeyPairSelector,
		To: reference.To{
			List:    &v1alpha13.KeypairV2List{},
			Managed: &v1alpha13.KeypairV2{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.KeyPair")
	}
	mg.Spec.InitProvider.KeyPair = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.KeyPairRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.RootVolume); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RootVolume[i3].KMSID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.RootVolume[i3].KMSIDRef,
			Selector:     mg.Spec.InitProvider.RootVolume[i3].KMSIDSelector,
			To: reference.To{
				List:    &v1alpha12.KeyV1List{},
				Managed: &v1alpha12.KeyV1{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.RootVolume[i3].KMSID")
		}
		mg.Spec.InitProvider.RootVolume[i3].KMSID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.RootVolume[i3].KMSIDRef = rsp.ResolvedReference

	}

	return nil
}