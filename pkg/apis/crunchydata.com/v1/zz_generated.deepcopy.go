// +build !ignore_autogenerated

/*
Copyright 2020 - 2021 Crunchy Data Solutions, Inc.
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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	corev1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterAnnotations) DeepCopyInto(out *ClusterAnnotations) {
	*out = *in
	if in.Backrest != nil {
		in, out := &in.Backrest, &out.Backrest
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Global != nil {
		in, out := &in.Global, &out.Global
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.PgBouncer != nil {
		in, out := &in.PgBouncer, &out.PgBouncer
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Postgres != nil {
		in, out := &in.Postgres, &out.Postgres
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterAnnotations.
func (in *ClusterAnnotations) DeepCopy() *ClusterAnnotations {
	if in == nil {
		return nil
	}
	out := new(ClusterAnnotations)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeAffinitySpec) DeepCopyInto(out *NodeAffinitySpec) {
	*out = *in
	if in.Default != nil {
		in, out := &in.Default, &out.Default
		*out = new(corev1.NodeAffinity)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeAffinitySpec.
func (in *NodeAffinitySpec) DeepCopy() *NodeAffinitySpec {
	if in == nil {
		return nil
	}
	out := new(NodeAffinitySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PGDataSourceSpec) DeepCopyInto(out *PGDataSourceSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PGDataSourceSpec.
func (in *PGDataSourceSpec) DeepCopy() *PGDataSourceSpec {
	if in == nil {
		return nil
	}
	out := new(PGDataSourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PgBouncerSpec) DeepCopyInto(out *PgBouncerSpec) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make(corev1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	if in.Limits != nil {
		in, out := &in.Limits, &out.Limits
		*out = make(corev1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PgBouncerSpec.
func (in *PgBouncerSpec) DeepCopy() *PgBouncerSpec {
	if in == nil {
		return nil
	}
	out := new(PgBouncerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PgStorageSpec) DeepCopyInto(out *PgStorageSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PgStorageSpec.
func (in *PgStorageSpec) DeepCopy() *PgStorageSpec {
	if in == nil {
		return nil
	}
	out := new(PgStorageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Pgcluster) DeepCopyInto(out *Pgcluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Pgcluster.
func (in *Pgcluster) DeepCopy() *Pgcluster {
	if in == nil {
		return nil
	}
	out := new(Pgcluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Pgcluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PgclusterList) DeepCopyInto(out *PgclusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Pgcluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PgclusterList.
func (in *PgclusterList) DeepCopy() *PgclusterList {
	if in == nil {
		return nil
	}
	out := new(PgclusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PgclusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PgclusterSpec) DeepCopyInto(out *PgclusterSpec) {
	*out = *in
	out.PrimaryStorage = in.PrimaryStorage
	out.WALStorage = in.WALStorage
	out.ReplicaStorage = in.ReplicaStorage
	out.BackrestStorage = in.BackrestStorage
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make(corev1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	if in.Limits != nil {
		in, out := &in.Limits, &out.Limits
		*out = make(corev1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	if in.BackrestResources != nil {
		in, out := &in.BackrestResources, &out.BackrestResources
		*out = make(corev1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	if in.BackrestLimits != nil {
		in, out := &in.BackrestLimits, &out.BackrestLimits
		*out = make(corev1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	if in.ExporterResources != nil {
		in, out := &in.ExporterResources, &out.ExporterResources
		*out = make(corev1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	if in.ExporterLimits != nil {
		in, out := &in.ExporterLimits, &out.ExporterLimits
		*out = make(corev1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	in.PgBouncer.DeepCopyInto(&out.PgBouncer)
	if in.UserLabels != nil {
		in, out := &in.UserLabels, &out.UserLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.NodeAffinity.DeepCopyInto(&out.NodeAffinity)
	out.PodAntiAffinity = in.PodAntiAffinity
	if in.SyncReplication != nil {
		in, out := &in.SyncReplication, &out.SyncReplication
		*out = new(bool)
		**out = **in
	}
	if in.BackrestConfig != nil {
		in, out := &in.BackrestConfig, &out.BackrestConfig
		*out = make([]corev1.VolumeProjection, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.BackrestStorageTypes != nil {
		in, out := &in.BackrestStorageTypes, &out.BackrestStorageTypes
		*out = make([]BackrestStorageType, len(*in))
		copy(*out, *in)
	}
	if in.TablespaceMounts != nil {
		in, out := &in.TablespaceMounts, &out.TablespaceMounts
		*out = make(map[string]PgStorageSpec, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.TLS = in.TLS
	out.PGDataSource = in.PGDataSource
	in.Annotations.DeepCopyInto(&out.Annotations)
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]corev1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PgclusterSpec.
func (in *PgclusterSpec) DeepCopy() *PgclusterSpec {
	if in == nil {
		return nil
	}
	out := new(PgclusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PgclusterStatus) DeepCopyInto(out *PgclusterStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PgclusterStatus.
func (in *PgclusterStatus) DeepCopy() *PgclusterStatus {
	if in == nil {
		return nil
	}
	out := new(PgclusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Pgpolicy) DeepCopyInto(out *Pgpolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Pgpolicy.
func (in *Pgpolicy) DeepCopy() *Pgpolicy {
	if in == nil {
		return nil
	}
	out := new(Pgpolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Pgpolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PgpolicyList) DeepCopyInto(out *PgpolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Pgpolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PgpolicyList.
func (in *PgpolicyList) DeepCopy() *PgpolicyList {
	if in == nil {
		return nil
	}
	out := new(PgpolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PgpolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PgpolicySpec) DeepCopyInto(out *PgpolicySpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PgpolicySpec.
func (in *PgpolicySpec) DeepCopy() *PgpolicySpec {
	if in == nil {
		return nil
	}
	out := new(PgpolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PgpolicyStatus) DeepCopyInto(out *PgpolicyStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PgpolicyStatus.
func (in *PgpolicyStatus) DeepCopy() *PgpolicyStatus {
	if in == nil {
		return nil
	}
	out := new(PgpolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Pgreplica) DeepCopyInto(out *Pgreplica) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Pgreplica.
func (in *Pgreplica) DeepCopy() *Pgreplica {
	if in == nil {
		return nil
	}
	out := new(Pgreplica)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Pgreplica) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PgreplicaList) DeepCopyInto(out *PgreplicaList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Pgreplica, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PgreplicaList.
func (in *PgreplicaList) DeepCopy() *PgreplicaList {
	if in == nil {
		return nil
	}
	out := new(PgreplicaList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PgreplicaList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PgreplicaSpec) DeepCopyInto(out *PgreplicaSpec) {
	*out = *in
	out.ReplicaStorage = in.ReplicaStorage
	if in.UserLabels != nil {
		in, out := &in.UserLabels, &out.UserLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.NodeAffinity != nil {
		in, out := &in.NodeAffinity, &out.NodeAffinity
		*out = new(corev1.NodeAffinity)
		(*in).DeepCopyInto(*out)
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]corev1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PgreplicaSpec.
func (in *PgreplicaSpec) DeepCopy() *PgreplicaSpec {
	if in == nil {
		return nil
	}
	out := new(PgreplicaSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PgreplicaStatus) DeepCopyInto(out *PgreplicaStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PgreplicaStatus.
func (in *PgreplicaStatus) DeepCopy() *PgreplicaStatus {
	if in == nil {
		return nil
	}
	out := new(PgreplicaStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Pgtask) DeepCopyInto(out *Pgtask) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Pgtask.
func (in *Pgtask) DeepCopy() *Pgtask {
	if in == nil {
		return nil
	}
	out := new(Pgtask)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Pgtask) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PgtaskList) DeepCopyInto(out *PgtaskList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Pgtask, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PgtaskList.
func (in *PgtaskList) DeepCopy() *PgtaskList {
	if in == nil {
		return nil
	}
	out := new(PgtaskList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PgtaskList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PgtaskSpec) DeepCopyInto(out *PgtaskSpec) {
	*out = *in
	out.StorageSpec = in.StorageSpec
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PgtaskSpec.
func (in *PgtaskSpec) DeepCopy() *PgtaskSpec {
	if in == nil {
		return nil
	}
	out := new(PgtaskSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PgtaskStatus) DeepCopyInto(out *PgtaskStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PgtaskStatus.
func (in *PgtaskStatus) DeepCopy() *PgtaskStatus {
	if in == nil {
		return nil
	}
	out := new(PgtaskStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodAntiAffinitySpec) DeepCopyInto(out *PodAntiAffinitySpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodAntiAffinitySpec.
func (in *PodAntiAffinitySpec) DeepCopy() *PodAntiAffinitySpec {
	if in == nil {
		return nil
	}
	out := new(PodAntiAffinitySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLSSpec) DeepCopyInto(out *TLSSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLSSpec.
func (in *TLSSpec) DeepCopy() *TLSSpec {
	if in == nil {
		return nil
	}
	out := new(TLSSpec)
	in.DeepCopyInto(out)
	return out
}
