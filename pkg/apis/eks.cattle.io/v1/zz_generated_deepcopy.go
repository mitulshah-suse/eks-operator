//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2019 Wrangler Sample Controller Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by main. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EKSClusterConfig) DeepCopyInto(out *EKSClusterConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EKSClusterConfig.
func (in *EKSClusterConfig) DeepCopy() *EKSClusterConfig {
	if in == nil {
		return nil
	}
	out := new(EKSClusterConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EKSClusterConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EKSClusterConfigList) DeepCopyInto(out *EKSClusterConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EKSClusterConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EKSClusterConfigList.
func (in *EKSClusterConfigList) DeepCopy() *EKSClusterConfigList {
	if in == nil {
		return nil
	}
	out := new(EKSClusterConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EKSClusterConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EKSClusterConfigSpec) DeepCopyInto(out *EKSClusterConfigSpec) {
	*out = *in
	if in.KubernetesVersion != nil {
		in, out := &in.KubernetesVersion, &out.KubernetesVersion
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SecretsEncryption != nil {
		in, out := &in.SecretsEncryption, &out.SecretsEncryption
		*out = new(bool)
		**out = **in
	}
	if in.KmsKey != nil {
		in, out := &in.KmsKey, &out.KmsKey
		*out = new(string)
		**out = **in
	}
	if in.PublicAccess != nil {
		in, out := &in.PublicAccess, &out.PublicAccess
		*out = new(bool)
		**out = **in
	}
	if in.PrivateAccess != nil {
		in, out := &in.PrivateAccess, &out.PrivateAccess
		*out = new(bool)
		**out = **in
	}
	if in.EBSCSIDriver != nil {
		in, out := &in.EBSCSIDriver, &out.EBSCSIDriver
		*out = new(bool)
		**out = **in
	}
	if in.PublicAccessSources != nil {
		in, out := &in.PublicAccessSources, &out.PublicAccessSources
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.LoggingTypes != nil {
		in, out := &in.LoggingTypes, &out.LoggingTypes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Subnets != nil {
		in, out := &in.Subnets, &out.Subnets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SecurityGroups != nil {
		in, out := &in.SecurityGroups, &out.SecurityGroups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ServiceRole != nil {
		in, out := &in.ServiceRole, &out.ServiceRole
		*out = new(string)
		**out = **in
	}
	if in.NodeGroups != nil {
		in, out := &in.NodeGroups, &out.NodeGroups
		*out = make([]NodeGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EKSClusterConfigSpec.
func (in *EKSClusterConfigSpec) DeepCopy() *EKSClusterConfigSpec {
	if in == nil {
		return nil
	}
	out := new(EKSClusterConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EKSClusterConfigStatus) DeepCopyInto(out *EKSClusterConfigStatus) {
	*out = *in
	if in.Subnets != nil {
		in, out := &in.Subnets, &out.Subnets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SecurityGroups != nil {
		in, out := &in.SecurityGroups, &out.SecurityGroups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ManagedLaunchTemplateVersions != nil {
		in, out := &in.ManagedLaunchTemplateVersions, &out.ManagedLaunchTemplateVersions
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.TemplateVersionsToDelete != nil {
		in, out := &in.TemplateVersionsToDelete, &out.TemplateVersionsToDelete
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EKSClusterConfigStatus.
func (in *EKSClusterConfigStatus) DeepCopy() *EKSClusterConfigStatus {
	if in == nil {
		return nil
	}
	out := new(EKSClusterConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LaunchTemplate) DeepCopyInto(out *LaunchTemplate) {
	*out = *in
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
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LaunchTemplate.
func (in *LaunchTemplate) DeepCopy() *LaunchTemplate {
	if in == nil {
		return nil
	}
	out := new(LaunchTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeGroup) DeepCopyInto(out *NodeGroup) {
	*out = *in
	if in.Gpu != nil {
		in, out := &in.Gpu, &out.Gpu
		*out = new(bool)
		**out = **in
	}
	if in.Arm != nil {
		in, out := &in.Arm, &out.Arm
		*out = new(bool)
		**out = **in
	}
	if in.ImageID != nil {
		in, out := &in.ImageID, &out.ImageID
		*out = new(string)
		**out = **in
	}
	if in.NodegroupName != nil {
		in, out := &in.NodegroupName, &out.NodegroupName
		*out = new(string)
		**out = **in
	}
	if in.DiskSize != nil {
		in, out := &in.DiskSize, &out.DiskSize
		*out = new(int32)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Ec2SshKey != nil {
		in, out := &in.Ec2SshKey, &out.Ec2SshKey
		*out = new(string)
		**out = **in
	}
	if in.DesiredSize != nil {
		in, out := &in.DesiredSize, &out.DesiredSize
		*out = new(int32)
		**out = **in
	}
	if in.MaxSize != nil {
		in, out := &in.MaxSize, &out.MaxSize
		*out = new(int32)
		**out = **in
	}
	if in.MinSize != nil {
		in, out := &in.MinSize, &out.MinSize
		*out = new(int32)
		**out = **in
	}
	if in.Subnets != nil {
		in, out := &in.Subnets, &out.Subnets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.ResourceTags != nil {
		in, out := &in.ResourceTags, &out.ResourceTags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.UserData != nil {
		in, out := &in.UserData, &out.UserData
		*out = new(string)
		**out = **in
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
	if in.LaunchTemplate != nil {
		in, out := &in.LaunchTemplate, &out.LaunchTemplate
		*out = new(LaunchTemplate)
		(*in).DeepCopyInto(*out)
	}
	if in.RequestSpotInstances != nil {
		in, out := &in.RequestSpotInstances, &out.RequestSpotInstances
		*out = new(bool)
		**out = **in
	}
	if in.SpotInstanceTypes != nil {
		in, out := &in.SpotInstanceTypes, &out.SpotInstanceTypes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NodeRole != nil {
		in, out := &in.NodeRole, &out.NodeRole
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeGroup.
func (in *NodeGroup) DeepCopy() *NodeGroup {
	if in == nil {
		return nil
	}
	out := new(NodeGroup)
	in.DeepCopyInto(out)
	return out
}
