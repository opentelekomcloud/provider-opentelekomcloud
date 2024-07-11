// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	"github.com/pkg/errors"

	"github.com/crossplane/upjet/pkg/resource"
	"github.com/crossplane/upjet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this FloatingipAssociateV2
func (mg *FloatingipAssociateV2) GetTerraformResourceType() string {
	return "opentelekomcloud_compute_floatingip_associate_v2"
}

// GetConnectionDetailsMapping for this FloatingipAssociateV2
func (tr *FloatingipAssociateV2) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this FloatingipAssociateV2
func (tr *FloatingipAssociateV2) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this FloatingipAssociateV2
func (tr *FloatingipAssociateV2) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this FloatingipAssociateV2
func (tr *FloatingipAssociateV2) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this FloatingipAssociateV2
func (tr *FloatingipAssociateV2) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this FloatingipAssociateV2
func (tr *FloatingipAssociateV2) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this FloatingipAssociateV2
func (tr *FloatingipAssociateV2) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// LateInitialize this FloatingipAssociateV2 using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *FloatingipAssociateV2) LateInitialize(attrs []byte) (bool, error) {
	params := &FloatingipAssociateV2Parameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *FloatingipAssociateV2) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this FloatingipV2
func (mg *FloatingipV2) GetTerraformResourceType() string {
	return "opentelekomcloud_compute_floatingip_v2"
}

// GetConnectionDetailsMapping for this FloatingipV2
func (tr *FloatingipV2) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this FloatingipV2
func (tr *FloatingipV2) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this FloatingipV2
func (tr *FloatingipV2) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this FloatingipV2
func (tr *FloatingipV2) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this FloatingipV2
func (tr *FloatingipV2) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this FloatingipV2
func (tr *FloatingipV2) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this FloatingipV2
func (tr *FloatingipV2) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// LateInitialize this FloatingipV2 using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *FloatingipV2) LateInitialize(attrs []byte) (bool, error) {
	params := &FloatingipV2Parameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *FloatingipV2) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this InstanceV2
func (mg *InstanceV2) GetTerraformResourceType() string {
	return "opentelekomcloud_compute_instance_v2"
}

// GetConnectionDetailsMapping for this InstanceV2
func (tr *InstanceV2) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"admin_pass": "spec.forProvider.adminPassSecretRef", "encrypted_password": "status.atProvider.encryptedPassword", "password": "status.atProvider.password", "ssh_private_key_path": "spec.forProvider.sshPrivateKeyPathSecretRef"}
}

// GetObservation of this InstanceV2
func (tr *InstanceV2) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this InstanceV2
func (tr *InstanceV2) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this InstanceV2
func (tr *InstanceV2) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this InstanceV2
func (tr *InstanceV2) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this InstanceV2
func (tr *InstanceV2) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this InstanceV2
func (tr *InstanceV2) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// LateInitialize this InstanceV2 using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *InstanceV2) LateInitialize(attrs []byte) (bool, error) {
	params := &InstanceV2Parameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *InstanceV2) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this KeypairV2
func (mg *KeypairV2) GetTerraformResourceType() string {
	return "opentelekomcloud_compute_keypair_v2"
}

// GetConnectionDetailsMapping for this KeypairV2
func (tr *KeypairV2) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this KeypairV2
func (tr *KeypairV2) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this KeypairV2
func (tr *KeypairV2) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this KeypairV2
func (tr *KeypairV2) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this KeypairV2
func (tr *KeypairV2) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this KeypairV2
func (tr *KeypairV2) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this KeypairV2
func (tr *KeypairV2) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// LateInitialize this KeypairV2 using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *KeypairV2) LateInitialize(attrs []byte) (bool, error) {
	params := &KeypairV2Parameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *KeypairV2) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this SecgroupV2
func (mg *SecgroupV2) GetTerraformResourceType() string {
	return "opentelekomcloud_compute_secgroup_v2"
}

// GetConnectionDetailsMapping for this SecgroupV2
func (tr *SecgroupV2) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this SecgroupV2
func (tr *SecgroupV2) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this SecgroupV2
func (tr *SecgroupV2) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this SecgroupV2
func (tr *SecgroupV2) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this SecgroupV2
func (tr *SecgroupV2) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this SecgroupV2
func (tr *SecgroupV2) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this SecgroupV2
func (tr *SecgroupV2) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// LateInitialize this SecgroupV2 using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *SecgroupV2) LateInitialize(attrs []byte) (bool, error) {
	params := &SecgroupV2Parameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *SecgroupV2) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this ServergroupV2
func (mg *ServergroupV2) GetTerraformResourceType() string {
	return "opentelekomcloud_compute_servergroup_v2"
}

// GetConnectionDetailsMapping for this ServergroupV2
func (tr *ServergroupV2) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this ServergroupV2
func (tr *ServergroupV2) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this ServergroupV2
func (tr *ServergroupV2) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this ServergroupV2
func (tr *ServergroupV2) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this ServergroupV2
func (tr *ServergroupV2) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this ServergroupV2
func (tr *ServergroupV2) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this ServergroupV2
func (tr *ServergroupV2) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// LateInitialize this ServergroupV2 using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *ServergroupV2) LateInitialize(attrs []byte) (bool, error) {
	params := &ServergroupV2Parameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *ServergroupV2) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this VolumeAttachV2
func (mg *VolumeAttachV2) GetTerraformResourceType() string {
	return "opentelekomcloud_compute_volume_attach_v2"
}

// GetConnectionDetailsMapping for this VolumeAttachV2
func (tr *VolumeAttachV2) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this VolumeAttachV2
func (tr *VolumeAttachV2) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this VolumeAttachV2
func (tr *VolumeAttachV2) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this VolumeAttachV2
func (tr *VolumeAttachV2) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this VolumeAttachV2
func (tr *VolumeAttachV2) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this VolumeAttachV2
func (tr *VolumeAttachV2) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this VolumeAttachV2
func (tr *VolumeAttachV2) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// LateInitialize this VolumeAttachV2 using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *VolumeAttachV2) LateInitialize(attrs []byte) (bool, error) {
	params := &VolumeAttachV2Parameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *VolumeAttachV2) GetTerraformSchemaVersion() int {
	return 0
}