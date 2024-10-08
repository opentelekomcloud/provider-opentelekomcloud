//go:build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessRuleInitParameters) DeepCopyInto(out *AccessRuleInitParameters) {
	*out = *in
	if in.AccessLevel != nil {
		in, out := &in.AccessLevel, &out.AccessLevel
		*out = new(string)
		**out = **in
	}
	if in.AccessTo != nil {
		in, out := &in.AccessTo, &out.AccessTo
		*out = new(string)
		**out = **in
	}
	if in.AccessType != nil {
		in, out := &in.AccessType, &out.AccessType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessRuleInitParameters.
func (in *AccessRuleInitParameters) DeepCopy() *AccessRuleInitParameters {
	if in == nil {
		return nil
	}
	out := new(AccessRuleInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessRuleObservation) DeepCopyInto(out *AccessRuleObservation) {
	*out = *in
	if in.AccessLevel != nil {
		in, out := &in.AccessLevel, &out.AccessLevel
		*out = new(string)
		**out = **in
	}
	if in.AccessRuleStatus != nil {
		in, out := &in.AccessRuleStatus, &out.AccessRuleStatus
		*out = new(string)
		**out = **in
	}
	if in.AccessTo != nil {
		in, out := &in.AccessTo, &out.AccessTo
		*out = new(string)
		**out = **in
	}
	if in.AccessType != nil {
		in, out := &in.AccessType, &out.AccessType
		*out = new(string)
		**out = **in
	}
	if in.ShareAccessID != nil {
		in, out := &in.ShareAccessID, &out.ShareAccessID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessRuleObservation.
func (in *AccessRuleObservation) DeepCopy() *AccessRuleObservation {
	if in == nil {
		return nil
	}
	out := new(AccessRuleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessRuleParameters) DeepCopyInto(out *AccessRuleParameters) {
	*out = *in
	if in.AccessLevel != nil {
		in, out := &in.AccessLevel, &out.AccessLevel
		*out = new(string)
		**out = **in
	}
	if in.AccessTo != nil {
		in, out := &in.AccessTo, &out.AccessTo
		*out = new(string)
		**out = **in
	}
	if in.AccessType != nil {
		in, out := &in.AccessType, &out.AccessType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessRuleParameters.
func (in *AccessRuleParameters) DeepCopy() *AccessRuleParameters {
	if in == nil {
		return nil
	}
	out := new(AccessRuleParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileSystemV2) DeepCopyInto(out *FileSystemV2) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileSystemV2.
func (in *FileSystemV2) DeepCopy() *FileSystemV2 {
	if in == nil {
		return nil
	}
	out := new(FileSystemV2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FileSystemV2) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileSystemV2InitParameters) DeepCopyInto(out *FileSystemV2InitParameters) {
	*out = *in
	if in.AccessLevel != nil {
		in, out := &in.AccessLevel, &out.AccessLevel
		*out = new(string)
		**out = **in
	}
	if in.AccessTo != nil {
		in, out := &in.AccessTo, &out.AccessTo
		*out = new(string)
		**out = **in
	}
	if in.AccessType != nil {
		in, out := &in.AccessType, &out.AccessType
		*out = new(string)
		**out = **in
	}
	if in.AvailabilityZone != nil {
		in, out := &in.AvailabilityZone, &out.AvailabilityZone
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.IsPublic != nil {
		in, out := &in.IsPublic, &out.IsPublic
		*out = new(bool)
		**out = **in
	}
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.ShareProto != nil {
		in, out := &in.ShareProto, &out.ShareProto
		*out = new(string)
		**out = **in
	}
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		*out = new(float64)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileSystemV2InitParameters.
func (in *FileSystemV2InitParameters) DeepCopy() *FileSystemV2InitParameters {
	if in == nil {
		return nil
	}
	out := new(FileSystemV2InitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileSystemV2List) DeepCopyInto(out *FileSystemV2List) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FileSystemV2, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileSystemV2List.
func (in *FileSystemV2List) DeepCopy() *FileSystemV2List {
	if in == nil {
		return nil
	}
	out := new(FileSystemV2List)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FileSystemV2List) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileSystemV2Observation) DeepCopyInto(out *FileSystemV2Observation) {
	*out = *in
	if in.AccessLevel != nil {
		in, out := &in.AccessLevel, &out.AccessLevel
		*out = new(string)
		**out = **in
	}
	if in.AccessRuleStatus != nil {
		in, out := &in.AccessRuleStatus, &out.AccessRuleStatus
		*out = new(string)
		**out = **in
	}
	if in.AccessTo != nil {
		in, out := &in.AccessTo, &out.AccessTo
		*out = new(string)
		**out = **in
	}
	if in.AccessType != nil {
		in, out := &in.AccessType, &out.AccessType
		*out = new(string)
		**out = **in
	}
	if in.AvailabilityZone != nil {
		in, out := &in.AvailabilityZone, &out.AvailabilityZone
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.ExportLocation != nil {
		in, out := &in.ExportLocation, &out.ExportLocation
		*out = new(string)
		**out = **in
	}
	if in.Host != nil {
		in, out := &in.Host, &out.Host
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.IsPublic != nil {
		in, out := &in.IsPublic, &out.IsPublic
		*out = new(bool)
		**out = **in
	}
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.ShareAccessID != nil {
		in, out := &in.ShareAccessID, &out.ShareAccessID
		*out = new(string)
		**out = **in
	}
	if in.ShareProto != nil {
		in, out := &in.ShareProto, &out.ShareProto
		*out = new(string)
		**out = **in
	}
	if in.ShareType != nil {
		in, out := &in.ShareType, &out.ShareType
		*out = new(string)
		**out = **in
	}
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		*out = new(float64)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.VolumeType != nil {
		in, out := &in.VolumeType, &out.VolumeType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileSystemV2Observation.
func (in *FileSystemV2Observation) DeepCopy() *FileSystemV2Observation {
	if in == nil {
		return nil
	}
	out := new(FileSystemV2Observation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileSystemV2Parameters) DeepCopyInto(out *FileSystemV2Parameters) {
	*out = *in
	if in.AccessLevel != nil {
		in, out := &in.AccessLevel, &out.AccessLevel
		*out = new(string)
		**out = **in
	}
	if in.AccessTo != nil {
		in, out := &in.AccessTo, &out.AccessTo
		*out = new(string)
		**out = **in
	}
	if in.AccessType != nil {
		in, out := &in.AccessType, &out.AccessType
		*out = new(string)
		**out = **in
	}
	if in.AvailabilityZone != nil {
		in, out := &in.AvailabilityZone, &out.AvailabilityZone
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.IsPublic != nil {
		in, out := &in.IsPublic, &out.IsPublic
		*out = new(bool)
		**out = **in
	}
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.ShareProto != nil {
		in, out := &in.ShareProto, &out.ShareProto
		*out = new(string)
		**out = **in
	}
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		*out = new(float64)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileSystemV2Parameters.
func (in *FileSystemV2Parameters) DeepCopy() *FileSystemV2Parameters {
	if in == nil {
		return nil
	}
	out := new(FileSystemV2Parameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileSystemV2Spec) DeepCopyInto(out *FileSystemV2Spec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileSystemV2Spec.
func (in *FileSystemV2Spec) DeepCopy() *FileSystemV2Spec {
	if in == nil {
		return nil
	}
	out := new(FileSystemV2Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileSystemV2Status) DeepCopyInto(out *FileSystemV2Status) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileSystemV2Status.
func (in *FileSystemV2Status) DeepCopy() *FileSystemV2Status {
	if in == nil {
		return nil
	}
	out := new(FileSystemV2Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShareAccessRulesV2) DeepCopyInto(out *ShareAccessRulesV2) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShareAccessRulesV2.
func (in *ShareAccessRulesV2) DeepCopy() *ShareAccessRulesV2 {
	if in == nil {
		return nil
	}
	out := new(ShareAccessRulesV2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ShareAccessRulesV2) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShareAccessRulesV2InitParameters) DeepCopyInto(out *ShareAccessRulesV2InitParameters) {
	*out = *in
	if in.AccessRule != nil {
		in, out := &in.AccessRule, &out.AccessRule
		*out = make([]AccessRuleInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ShareID != nil {
		in, out := &in.ShareID, &out.ShareID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShareAccessRulesV2InitParameters.
func (in *ShareAccessRulesV2InitParameters) DeepCopy() *ShareAccessRulesV2InitParameters {
	if in == nil {
		return nil
	}
	out := new(ShareAccessRulesV2InitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShareAccessRulesV2List) DeepCopyInto(out *ShareAccessRulesV2List) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ShareAccessRulesV2, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShareAccessRulesV2List.
func (in *ShareAccessRulesV2List) DeepCopy() *ShareAccessRulesV2List {
	if in == nil {
		return nil
	}
	out := new(ShareAccessRulesV2List)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ShareAccessRulesV2List) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShareAccessRulesV2Observation) DeepCopyInto(out *ShareAccessRulesV2Observation) {
	*out = *in
	if in.AccessRule != nil {
		in, out := &in.AccessRule, &out.AccessRule
		*out = make([]AccessRuleObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.ShareID != nil {
		in, out := &in.ShareID, &out.ShareID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShareAccessRulesV2Observation.
func (in *ShareAccessRulesV2Observation) DeepCopy() *ShareAccessRulesV2Observation {
	if in == nil {
		return nil
	}
	out := new(ShareAccessRulesV2Observation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShareAccessRulesV2Parameters) DeepCopyInto(out *ShareAccessRulesV2Parameters) {
	*out = *in
	if in.AccessRule != nil {
		in, out := &in.AccessRule, &out.AccessRule
		*out = make([]AccessRuleParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ShareID != nil {
		in, out := &in.ShareID, &out.ShareID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShareAccessRulesV2Parameters.
func (in *ShareAccessRulesV2Parameters) DeepCopy() *ShareAccessRulesV2Parameters {
	if in == nil {
		return nil
	}
	out := new(ShareAccessRulesV2Parameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShareAccessRulesV2Spec) DeepCopyInto(out *ShareAccessRulesV2Spec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShareAccessRulesV2Spec.
func (in *ShareAccessRulesV2Spec) DeepCopy() *ShareAccessRulesV2Spec {
	if in == nil {
		return nil
	}
	out := new(ShareAccessRulesV2Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShareAccessRulesV2Status) DeepCopyInto(out *ShareAccessRulesV2Status) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShareAccessRulesV2Status.
func (in *ShareAccessRulesV2Status) DeepCopy() *ShareAccessRulesV2Status {
	if in == nil {
		return nil
	}
	out := new(ShareAccessRulesV2Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TurboShareV1) DeepCopyInto(out *TurboShareV1) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TurboShareV1.
func (in *TurboShareV1) DeepCopy() *TurboShareV1 {
	if in == nil {
		return nil
	}
	out := new(TurboShareV1)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TurboShareV1) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TurboShareV1InitParameters) DeepCopyInto(out *TurboShareV1InitParameters) {
	*out = *in
	if in.AvailabilityZone != nil {
		in, out := &in.AvailabilityZone, &out.AvailabilityZone
		*out = new(string)
		**out = **in
	}
	if in.CryptKeyID != nil {
		in, out := &in.CryptKeyID, &out.CryptKeyID
		*out = new(string)
		**out = **in
	}
	if in.Enhanced != nil {
		in, out := &in.Enhanced, &out.Enhanced
		*out = new(bool)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.SecurityGroupID != nil {
		in, out := &in.SecurityGroupID, &out.SecurityGroupID
		*out = new(string)
		**out = **in
	}
	if in.ShareProto != nil {
		in, out := &in.ShareProto, &out.ShareProto
		*out = new(string)
		**out = **in
	}
	if in.ShareType != nil {
		in, out := &in.ShareType, &out.ShareType
		*out = new(string)
		**out = **in
	}
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		*out = new(float64)
		**out = **in
	}
	if in.SubnetID != nil {
		in, out := &in.SubnetID, &out.SubnetID
		*out = new(string)
		**out = **in
	}
	if in.VPCID != nil {
		in, out := &in.VPCID, &out.VPCID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TurboShareV1InitParameters.
func (in *TurboShareV1InitParameters) DeepCopy() *TurboShareV1InitParameters {
	if in == nil {
		return nil
	}
	out := new(TurboShareV1InitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TurboShareV1List) DeepCopyInto(out *TurboShareV1List) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TurboShareV1, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TurboShareV1List.
func (in *TurboShareV1List) DeepCopy() *TurboShareV1List {
	if in == nil {
		return nil
	}
	out := new(TurboShareV1List)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TurboShareV1List) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TurboShareV1Observation) DeepCopyInto(out *TurboShareV1Observation) {
	*out = *in
	if in.AvailabilityZone != nil {
		in, out := &in.AvailabilityZone, &out.AvailabilityZone
		*out = new(string)
		**out = **in
	}
	if in.AvailableCapacity != nil {
		in, out := &in.AvailableCapacity, &out.AvailableCapacity
		*out = new(string)
		**out = **in
	}
	if in.CryptKeyID != nil {
		in, out := &in.CryptKeyID, &out.CryptKeyID
		*out = new(string)
		**out = **in
	}
	if in.Enhanced != nil {
		in, out := &in.Enhanced, &out.Enhanced
		*out = new(bool)
		**out = **in
	}
	if in.ExpandType != nil {
		in, out := &in.ExpandType, &out.ExpandType
		*out = new(string)
		**out = **in
	}
	if in.ExportLocation != nil {
		in, out := &in.ExportLocation, &out.ExportLocation
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.SecurityGroupID != nil {
		in, out := &in.SecurityGroupID, &out.SecurityGroupID
		*out = new(string)
		**out = **in
	}
	if in.ShareProto != nil {
		in, out := &in.ShareProto, &out.ShareProto
		*out = new(string)
		**out = **in
	}
	if in.ShareType != nil {
		in, out := &in.ShareType, &out.ShareType
		*out = new(string)
		**out = **in
	}
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		*out = new(float64)
		**out = **in
	}
	if in.SubnetID != nil {
		in, out := &in.SubnetID, &out.SubnetID
		*out = new(string)
		**out = **in
	}
	if in.VPCID != nil {
		in, out := &in.VPCID, &out.VPCID
		*out = new(string)
		**out = **in
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TurboShareV1Observation.
func (in *TurboShareV1Observation) DeepCopy() *TurboShareV1Observation {
	if in == nil {
		return nil
	}
	out := new(TurboShareV1Observation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TurboShareV1Parameters) DeepCopyInto(out *TurboShareV1Parameters) {
	*out = *in
	if in.AvailabilityZone != nil {
		in, out := &in.AvailabilityZone, &out.AvailabilityZone
		*out = new(string)
		**out = **in
	}
	if in.CryptKeyID != nil {
		in, out := &in.CryptKeyID, &out.CryptKeyID
		*out = new(string)
		**out = **in
	}
	if in.Enhanced != nil {
		in, out := &in.Enhanced, &out.Enhanced
		*out = new(bool)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.SecurityGroupID != nil {
		in, out := &in.SecurityGroupID, &out.SecurityGroupID
		*out = new(string)
		**out = **in
	}
	if in.ShareProto != nil {
		in, out := &in.ShareProto, &out.ShareProto
		*out = new(string)
		**out = **in
	}
	if in.ShareType != nil {
		in, out := &in.ShareType, &out.ShareType
		*out = new(string)
		**out = **in
	}
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		*out = new(float64)
		**out = **in
	}
	if in.SubnetID != nil {
		in, out := &in.SubnetID, &out.SubnetID
		*out = new(string)
		**out = **in
	}
	if in.VPCID != nil {
		in, out := &in.VPCID, &out.VPCID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TurboShareV1Parameters.
func (in *TurboShareV1Parameters) DeepCopy() *TurboShareV1Parameters {
	if in == nil {
		return nil
	}
	out := new(TurboShareV1Parameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TurboShareV1Spec) DeepCopyInto(out *TurboShareV1Spec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TurboShareV1Spec.
func (in *TurboShareV1Spec) DeepCopy() *TurboShareV1Spec {
	if in == nil {
		return nil
	}
	out := new(TurboShareV1Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TurboShareV1Status) DeepCopyInto(out *TurboShareV1Status) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TurboShareV1Status.
func (in *TurboShareV1Status) DeepCopy() *TurboShareV1Status {
	if in == nil {
		return nil
	}
	out := new(TurboShareV1Status)
	in.DeepCopyInto(out)
	return out
}
