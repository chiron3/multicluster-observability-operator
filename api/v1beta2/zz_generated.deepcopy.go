// +build !ignore_autogenerated

/*
Copyright 2021.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1beta2

import (
	"github.com/open-cluster-management/multicluster-observability-operator/api/v1beta1"
	"k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Condition) DeepCopyInto(out *Condition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Condition.
func (in *Condition) DeepCopy() *Condition {
	if in == nil {
		return nil
	}
	out := new(Condition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultiClusterObservability) DeepCopyInto(out *MultiClusterObservability) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultiClusterObservability.
func (in *MultiClusterObservability) DeepCopy() *MultiClusterObservability {
	if in == nil {
		return nil
	}
	out := new(MultiClusterObservability)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MultiClusterObservability) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultiClusterObservabilityList) DeepCopyInto(out *MultiClusterObservabilityList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MultiClusterObservability, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultiClusterObservabilityList.
func (in *MultiClusterObservabilityList) DeepCopy() *MultiClusterObservabilityList {
	if in == nil {
		return nil
	}
	out := new(MultiClusterObservabilityList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MultiClusterObservabilityList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultiClusterObservabilitySpec) DeepCopyInto(out *MultiClusterObservabilitySpec) {
	*out = *in
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RetentionConfig != nil {
		in, out := &in.RetentionConfig, &out.RetentionConfig
		*out = new(RetentionConfig)
		**out = **in
	}
	if in.StorageConfig != nil {
		in, out := &in.StorageConfig, &out.StorageConfig
		*out = new(StorageConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.ObservabilityAddonSpec != nil {
		in, out := &in.ObservabilityAddonSpec, &out.ObservabilityAddonSpec
		*out = new(v1beta1.ObservabilityAddonSpec)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultiClusterObservabilitySpec.
func (in *MultiClusterObservabilitySpec) DeepCopy() *MultiClusterObservabilitySpec {
	if in == nil {
		return nil
	}
	out := new(MultiClusterObservabilitySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultiClusterObservabilityStatus) DeepCopyInto(out *MultiClusterObservabilityStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultiClusterObservabilityStatus.
func (in *MultiClusterObservabilityStatus) DeepCopy() *MultiClusterObservabilityStatus {
	if in == nil {
		return nil
	}
	out := new(MultiClusterObservabilityStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PreConfiguredStorage) DeepCopyInto(out *PreConfiguredStorage) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PreConfiguredStorage.
func (in *PreConfiguredStorage) DeepCopy() *PreConfiguredStorage {
	if in == nil {
		return nil
	}
	out := new(PreConfiguredStorage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RetentionConfig) DeepCopyInto(out *RetentionConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RetentionConfig.
func (in *RetentionConfig) DeepCopy() *RetentionConfig {
	if in == nil {
		return nil
	}
	out := new(RetentionConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageConfig) DeepCopyInto(out *StorageConfig) {
	*out = *in
	if in.MetricObjectStorage != nil {
		in, out := &in.MetricObjectStorage, &out.MetricObjectStorage
		*out = new(PreConfiguredStorage)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageConfig.
func (in *StorageConfig) DeepCopy() *StorageConfig {
	if in == nil {
		return nil
	}
	out := new(StorageConfig)
	in.DeepCopyInto(out)
	return out
}
