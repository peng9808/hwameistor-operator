//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022.

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

package v1alpha1

import (
	"k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdmissionControllerSpec) DeepCopyInto(out *AdmissionControllerSpec) {
	*out = *in
	if in.Common != nil {
		in, out := &in.Common, &out.Common
		*out = new(PodCommonSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Controller != nil {
		in, out := &in.Controller, &out.Controller
		*out = new(ContainerCommonSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdmissionControllerSpec.
func (in *AdmissionControllerSpec) DeepCopy() *AdmissionControllerSpec {
	if in == nil {
		return nil
	}
	out := new(AdmissionControllerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdmissionControllerStatus) DeepCopyInto(out *AdmissionControllerStatus) {
	*out = *in
	if in.Instances != nil {
		in, out := &in.Instances, &out.Instances
		*out = new(DeployStatus)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdmissionControllerStatus.
func (in *AdmissionControllerStatus) DeepCopy() *AdmissionControllerStatus {
	if in == nil {
		return nil
	}
	out := new(AdmissionControllerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApiServerSpec) DeepCopyInto(out *ApiServerSpec) {
	*out = *in
	if in.Common != nil {
		in, out := &in.Common, &out.Common
		*out = new(PodCommonSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Server != nil {
		in, out := &in.Server, &out.Server
		*out = new(ContainerCommonSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Authentication != nil {
		in, out := &in.Authentication, &out.Authentication
		*out = new(Authentication)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApiServerSpec.
func (in *ApiServerSpec) DeepCopy() *ApiServerSpec {
	if in == nil {
		return nil
	}
	out := new(ApiServerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApiServerStatus) DeepCopyInto(out *ApiServerStatus) {
	*out = *in
	if in.Instances != nil {
		in, out := &in.Instances, &out.Instances
		*out = new(DeployStatus)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApiServerStatus.
func (in *ApiServerStatus) DeepCopy() *ApiServerStatus {
	if in == nil {
		return nil
	}
	out := new(ApiServerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuditorSpec) DeepCopyInto(out *AuditorSpec) {
	*out = *in
	if in.Common != nil {
		in, out := &in.Common, &out.Common
		*out = new(PodCommonSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Auditor != nil {
		in, out := &in.Auditor, &out.Auditor
		*out = new(ContainerCommonSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuditorSpec.
func (in *AuditorSpec) DeepCopy() *AuditorSpec {
	if in == nil {
		return nil
	}
	out := new(AuditorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuditorStatus) DeepCopyInto(out *AuditorStatus) {
	*out = *in
	if in.Instances != nil {
		in, out := &in.Instances, &out.Instances
		*out = new(DeployStatus)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuditorStatus.
func (in *AuditorStatus) DeepCopy() *AuditorStatus {
	if in == nil {
		return nil
	}
	out := new(AuditorStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Authentication) DeepCopyInto(out *Authentication) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Authentication.
func (in *Authentication) DeepCopy() *Authentication {
	if in == nil {
		return nil
	}
	out := new(Authentication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CSIControllerSpec) DeepCopyInto(out *CSIControllerSpec) {
	*out = *in
	if in.Common != nil {
		in, out := &in.Common, &out.Common
		*out = new(PodCommonSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Provisioner != nil {
		in, out := &in.Provisioner, &out.Provisioner
		*out = new(ContainerCommonSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Attacher != nil {
		in, out := &in.Attacher, &out.Attacher
		*out = new(ContainerCommonSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Monitor != nil {
		in, out := &in.Monitor, &out.Monitor
		*out = new(ContainerCommonSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Resizer != nil {
		in, out := &in.Resizer, &out.Resizer
		*out = new(ContainerCommonSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.SnapshotController != nil {
		in, out := &in.SnapshotController, &out.SnapshotController
		*out = new(ContainerCommonSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Snapshotter != nil {
		in, out := &in.Snapshotter, &out.Snapshotter
		*out = new(ContainerCommonSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CSIControllerSpec.
func (in *CSIControllerSpec) DeepCopy() *CSIControllerSpec {
	if in == nil {
		return nil
	}
	out := new(CSIControllerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CSISpec) DeepCopyInto(out *CSISpec) {
	*out = *in
	if in.Registrar != nil {
		in, out := &in.Registrar, &out.Registrar
		*out = new(ContainerCommonSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Controller != nil {
		in, out := &in.Controller, &out.Controller
		*out = new(CSIControllerSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CSISpec.
func (in *CSISpec) DeepCopy() *CSISpec {
	if in == nil {
		return nil
	}
	out := new(CSISpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Cluster) DeepCopyInto(out *Cluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Cluster.
func (in *Cluster) DeepCopy() *Cluster {
	if in == nil {
		return nil
	}
	out := new(Cluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Cluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterList) DeepCopyInto(out *ClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Cluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterList.
func (in *ClusterList) DeepCopy() *ClusterList {
	if in == nil {
		return nil
	}
	out := new(ClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpec) DeepCopyInto(out *ClusterSpec) {
	*out = *in
	if in.DiskReserveConfigurations != nil {
		in, out := &in.DiskReserveConfigurations, &out.DiskReserveConfigurations
		*out = make([]DiskReserveConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LocalDiskManager != nil {
		in, out := &in.LocalDiskManager, &out.LocalDiskManager
		*out = new(LocalDiskManagerSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.LocalStorage != nil {
		in, out := &in.LocalStorage, &out.LocalStorage
		*out = new(LocalStorageSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Scheduler != nil {
		in, out := &in.Scheduler, &out.Scheduler
		*out = new(SchedulerSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Evictor != nil {
		in, out := &in.Evictor, &out.Evictor
		*out = new(EvictorSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.AdmissionController != nil {
		in, out := &in.AdmissionController, &out.AdmissionController
		*out = new(AdmissionControllerSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.ApiServer != nil {
		in, out := &in.ApiServer, &out.ApiServer
		*out = new(ApiServerSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Exporter != nil {
		in, out := &in.Exporter, &out.Exporter
		*out = new(ExporterSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Auditor != nil {
		in, out := &in.Auditor, &out.Auditor
		*out = new(AuditorSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.FailoverAssistant != nil {
		in, out := &in.FailoverAssistant, &out.FailoverAssistant
		*out = new(FailoverAssistantSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.PVCAutoResizer != nil {
		in, out := &in.PVCAutoResizer, &out.PVCAutoResizer
		*out = new(PVCAutoResizerSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.LocalDiskActionController != nil {
		in, out := &in.LocalDiskActionController, &out.LocalDiskActionController
		*out = new(LocalDiskActionControllerSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.UI != nil {
		in, out := &in.UI, &out.UI
		*out = new(UISpec)
		(*in).DeepCopyInto(*out)
	}
	if in.DRBD != nil {
		in, out := &in.DRBD, &out.DRBD
		*out = new(DRBDSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.RBAC != nil {
		in, out := &in.RBAC, &out.RBAC
		*out = new(RBACSpec)
		**out = **in
	}
	if in.StorageClass != nil {
		in, out := &in.StorageClass, &out.StorageClass
		*out = new(StorageClassSpec)
		**out = **in
	}
	if in.DataLoadManager != nil {
		in, out := &in.DataLoadManager, &out.DataLoadManager
		*out = new(DataLoadManagerSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.DataSetManager != nil {
		in, out := &in.DataSetManager, &out.DataSetManager
		*out = new(DataSetManagerSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpec.
func (in *ClusterSpec) DeepCopy() *ClusterSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterStatus) DeepCopyInto(out *ClusterStatus) {
	*out = *in
	in.ComponentStatus.DeepCopyInto(&out.ComponentStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterStatus.
func (in *ClusterStatus) DeepCopy() *ClusterStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentStatus) DeepCopyInto(out *ComponentStatus) {
	*out = *in
	if in.LocalDiskManager != nil {
		in, out := &in.LocalDiskManager, &out.LocalDiskManager
		*out = new(LocalDiskManagerStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.LocalStorage != nil {
		in, out := &in.LocalStorage, &out.LocalStorage
		*out = new(LocalStorageStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.Scheduler != nil {
		in, out := &in.Scheduler, &out.Scheduler
		*out = new(SchedulerStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.Evictor != nil {
		in, out := &in.Evictor, &out.Evictor
		*out = new(EvictorStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.AdmissionController != nil {
		in, out := &in.AdmissionController, &out.AdmissionController
		*out = new(AdmissionControllerStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.ApiServer != nil {
		in, out := &in.ApiServer, &out.ApiServer
		*out = new(ApiServerStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.Auditor != nil {
		in, out := &in.Auditor, &out.Auditor
		*out = new(AuditorStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.FailoverAssistant != nil {
		in, out := &in.FailoverAssistant, &out.FailoverAssistant
		*out = new(FailoverAssistantStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.PVCAutoResizer != nil {
		in, out := &in.PVCAutoResizer, &out.PVCAutoResizer
		*out = new(PVCAutoResizerStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.LocalDiskActionController != nil {
		in, out := &in.LocalDiskActionController, &out.LocalDiskActionController
		*out = new(LocalDiskActionControllerStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.Exporter != nil {
		in, out := &in.Exporter, &out.Exporter
		*out = new(ExporterStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.DataLoadManager != nil {
		in, out := &in.DataLoadManager, &out.DataLoadManager
		*out = new(DataLoadManagerStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.DataSetManager != nil {
		in, out := &in.DataSetManager, &out.DataSetManager
		*out = new(DataSetManagerStatus)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentStatus.
func (in *ComponentStatus) DeepCopy() *ComponentStatus {
	if in == nil {
		return nil
	}
	out := new(ComponentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerCommonSpec) DeepCopyInto(out *ContainerCommonSpec) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(ImageSpec)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerCommonSpec.
func (in *ContainerCommonSpec) DeepCopy() *ContainerCommonSpec {
	if in == nil {
		return nil
	}
	out := new(ContainerCommonSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DRBDSpec) DeepCopyInto(out *DRBDSpec) {
	*out = *in
	if in.Shipper != nil {
		in, out := &in.Shipper, &out.Shipper
		*out = new(ImageSpec)
		**out = **in
	}
	if in.NodeAffinity != nil {
		in, out := &in.NodeAffinity, &out.NodeAffinity
		*out = new(v1.NodeAffinity)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DRBDSpec.
func (in *DRBDSpec) DeepCopy() *DRBDSpec {
	if in == nil {
		return nil
	}
	out := new(DRBDSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataLoadManagerSpec) DeepCopyInto(out *DataLoadManagerSpec) {
	*out = *in
	if in.Common != nil {
		in, out := &in.Common, &out.Common
		*out = new(PodCommonSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.DataLoadManagerContainer != nil {
		in, out := &in.DataLoadManagerContainer, &out.DataLoadManagerContainer
		*out = new(ContainerCommonSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataLoadManagerSpec.
func (in *DataLoadManagerSpec) DeepCopy() *DataLoadManagerSpec {
	if in == nil {
		return nil
	}
	out := new(DataLoadManagerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataLoadManagerStatus) DeepCopyInto(out *DataLoadManagerStatus) {
	*out = *in
	if in.Instances != nil {
		in, out := &in.Instances, &out.Instances
		*out = new(DeployStatus)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataLoadManagerStatus.
func (in *DataLoadManagerStatus) DeepCopy() *DataLoadManagerStatus {
	if in == nil {
		return nil
	}
	out := new(DataLoadManagerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataSetManagerSpec) DeepCopyInto(out *DataSetManagerSpec) {
	*out = *in
	if in.Common != nil {
		in, out := &in.Common, &out.Common
		*out = new(PodCommonSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.DataSetManagerContainer != nil {
		in, out := &in.DataSetManagerContainer, &out.DataSetManagerContainer
		*out = new(ContainerCommonSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataSetManagerSpec.
func (in *DataSetManagerSpec) DeepCopy() *DataSetManagerSpec {
	if in == nil {
		return nil
	}
	out := new(DataSetManagerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataSetManagerStatus) DeepCopyInto(out *DataSetManagerStatus) {
	*out = *in
	if in.Instances != nil {
		in, out := &in.Instances, &out.Instances
		*out = new(DeployStatus)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataSetManagerStatus.
func (in *DataSetManagerStatus) DeepCopy() *DataSetManagerStatus {
	if in == nil {
		return nil
	}
	out := new(DataSetManagerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeployStatus) DeepCopyInto(out *DeployStatus) {
	*out = *in
	if in.Pods != nil {
		in, out := &in.Pods, &out.Pods
		*out = make([]PodStatus, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeployStatus.
func (in *DeployStatus) DeepCopy() *DeployStatus {
	if in == nil {
		return nil
	}
	out := new(DeployStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiskReserveConfiguration) DeepCopyInto(out *DiskReserveConfiguration) {
	*out = *in
	if in.Devices != nil {
		in, out := &in.Devices, &out.Devices
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiskReserveConfiguration.
func (in *DiskReserveConfiguration) DeepCopy() *DiskReserveConfiguration {
	if in == nil {
		return nil
	}
	out := new(DiskReserveConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EvictorSpec) DeepCopyInto(out *EvictorSpec) {
	*out = *in
	if in.Common != nil {
		in, out := &in.Common, &out.Common
		*out = new(PodCommonSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Evictor != nil {
		in, out := &in.Evictor, &out.Evictor
		*out = new(ContainerCommonSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EvictorSpec.
func (in *EvictorSpec) DeepCopy() *EvictorSpec {
	if in == nil {
		return nil
	}
	out := new(EvictorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EvictorStatus) DeepCopyInto(out *EvictorStatus) {
	*out = *in
	if in.Instances != nil {
		in, out := &in.Instances, &out.Instances
		*out = new(DeployStatus)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EvictorStatus.
func (in *EvictorStatus) DeepCopy() *EvictorStatus {
	if in == nil {
		return nil
	}
	out := new(EvictorStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExporterSpec) DeepCopyInto(out *ExporterSpec) {
	*out = *in
	if in.Common != nil {
		in, out := &in.Common, &out.Common
		*out = new(PodCommonSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Collector != nil {
		in, out := &in.Collector, &out.Collector
		*out = new(ContainerCommonSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExporterSpec.
func (in *ExporterSpec) DeepCopy() *ExporterSpec {
	if in == nil {
		return nil
	}
	out := new(ExporterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExporterStatus) DeepCopyInto(out *ExporterStatus) {
	*out = *in
	if in.Instances != nil {
		in, out := &in.Instances, &out.Instances
		*out = new(DeployStatus)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExporterStatus.
func (in *ExporterStatus) DeepCopy() *ExporterStatus {
	if in == nil {
		return nil
	}
	out := new(ExporterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FailoverAssistantSpec) DeepCopyInto(out *FailoverAssistantSpec) {
	*out = *in
	if in.Common != nil {
		in, out := &in.Common, &out.Common
		*out = new(PodCommonSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Assistant != nil {
		in, out := &in.Assistant, &out.Assistant
		*out = new(ContainerCommonSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FailoverAssistantSpec.
func (in *FailoverAssistantSpec) DeepCopy() *FailoverAssistantSpec {
	if in == nil {
		return nil
	}
	out := new(FailoverAssistantSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FailoverAssistantStatus) DeepCopyInto(out *FailoverAssistantStatus) {
	*out = *in
	if in.Instances != nil {
		in, out := &in.Instances, &out.Instances
		*out = new(DeployStatus)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FailoverAssistantStatus.
func (in *FailoverAssistantStatus) DeepCopy() *FailoverAssistantStatus {
	if in == nil {
		return nil
	}
	out := new(FailoverAssistantStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageSpec) DeepCopyInto(out *ImageSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageSpec.
func (in *ImageSpec) DeepCopy() *ImageSpec {
	if in == nil {
		return nil
	}
	out := new(ImageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalDiskActionControllerSpec) DeepCopyInto(out *LocalDiskActionControllerSpec) {
	*out = *in
	if in.Common != nil {
		in, out := &in.Common, &out.Common
		*out = new(PodCommonSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Controller != nil {
		in, out := &in.Controller, &out.Controller
		*out = new(ContainerCommonSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalDiskActionControllerSpec.
func (in *LocalDiskActionControllerSpec) DeepCopy() *LocalDiskActionControllerSpec {
	if in == nil {
		return nil
	}
	out := new(LocalDiskActionControllerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalDiskActionControllerStatus) DeepCopyInto(out *LocalDiskActionControllerStatus) {
	*out = *in
	if in.Instances != nil {
		in, out := &in.Instances, &out.Instances
		*out = new(DeployStatus)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalDiskActionControllerStatus.
func (in *LocalDiskActionControllerStatus) DeepCopy() *LocalDiskActionControllerStatus {
	if in == nil {
		return nil
	}
	out := new(LocalDiskActionControllerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalDiskManagerSpec) DeepCopyInto(out *LocalDiskManagerSpec) {
	*out = *in
	if in.CSI != nil {
		in, out := &in.CSI, &out.CSI
		*out = new(CSISpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Common != nil {
		in, out := &in.Common, &out.Common
		*out = new(PodCommonSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Manager != nil {
		in, out := &in.Manager, &out.Manager
		*out = new(ContainerCommonSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalDiskManagerSpec.
func (in *LocalDiskManagerSpec) DeepCopy() *LocalDiskManagerSpec {
	if in == nil {
		return nil
	}
	out := new(LocalDiskManagerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalDiskManagerStatus) DeepCopyInto(out *LocalDiskManagerStatus) {
	*out = *in
	if in.Instances != nil {
		in, out := &in.Instances, &out.Instances
		*out = new(DeployStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.CSI != nil {
		in, out := &in.CSI, &out.CSI
		*out = new(DeployStatus)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalDiskManagerStatus.
func (in *LocalDiskManagerStatus) DeepCopy() *LocalDiskManagerStatus {
	if in == nil {
		return nil
	}
	out := new(LocalDiskManagerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalStorageSpec) DeepCopyInto(out *LocalStorageSpec) {
	*out = *in
	if in.CSI != nil {
		in, out := &in.CSI, &out.CSI
		*out = new(CSISpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Member != nil {
		in, out := &in.Member, &out.Member
		*out = new(MemberSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Common != nil {
		in, out := &in.Common, &out.Common
		*out = new(PodCommonSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Snapshot != nil {
		in, out := &in.Snapshot, &out.Snapshot
		*out = new(SnapshotSpec)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalStorageSpec.
func (in *LocalStorageSpec) DeepCopy() *LocalStorageSpec {
	if in == nil {
		return nil
	}
	out := new(LocalStorageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalStorageStatus) DeepCopyInto(out *LocalStorageStatus) {
	*out = *in
	if in.Instances != nil {
		in, out := &in.Instances, &out.Instances
		*out = new(DeployStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.CSI != nil {
		in, out := &in.CSI, &out.CSI
		*out = new(DeployStatus)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalStorageStatus.
func (in *LocalStorageStatus) DeepCopy() *LocalStorageStatus {
	if in == nil {
		return nil
	}
	out := new(LocalStorageStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemberSpec) DeepCopyInto(out *MemberSpec) {
	*out = *in
	if in.JuicesyncImage != nil {
		in, out := &in.JuicesyncImage, &out.JuicesyncImage
		*out = new(ImageSpec)
		**out = **in
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(ImageSpec)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemberSpec.
func (in *MemberSpec) DeepCopy() *MemberSpec {
	if in == nil {
		return nil
	}
	out := new(MemberSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PVCAutoResizerSpec) DeepCopyInto(out *PVCAutoResizerSpec) {
	*out = *in
	if in.Common != nil {
		in, out := &in.Common, &out.Common
		*out = new(PodCommonSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.AutoResizer != nil {
		in, out := &in.AutoResizer, &out.AutoResizer
		*out = new(ContainerCommonSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PVCAutoResizerSpec.
func (in *PVCAutoResizerSpec) DeepCopy() *PVCAutoResizerSpec {
	if in == nil {
		return nil
	}
	out := new(PVCAutoResizerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PVCAutoResizerStatus) DeepCopyInto(out *PVCAutoResizerStatus) {
	*out = *in
	if in.Instances != nil {
		in, out := &in.Instances, &out.Instances
		*out = new(DeployStatus)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PVCAutoResizerStatus.
func (in *PVCAutoResizerStatus) DeepCopy() *PVCAutoResizerStatus {
	if in == nil {
		return nil
	}
	out := new(PVCAutoResizerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodCommonSpec) DeepCopyInto(out *PodCommonSpec) {
	*out = *in
	if in.NodeAffinity != nil {
		in, out := &in.NodeAffinity, &out.NodeAffinity
		*out = new(v1.NodeAffinity)
		(*in).DeepCopyInto(*out)
	}
	if in.PodAffinty != nil {
		in, out := &in.PodAffinty, &out.PodAffinty
		*out = new(v1.PodAffinity)
		(*in).DeepCopyInto(*out)
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = new([]v1.Toleration)
		if **in != nil {
			in, out := *in, *out
			*out = make([]v1.Toleration, len(*in))
			for i := range *in {
				(*in)[i].DeepCopyInto(&(*out)[i])
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodCommonSpec.
func (in *PodCommonSpec) DeepCopy() *PodCommonSpec {
	if in == nil {
		return nil
	}
	out := new(PodCommonSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodStatus) DeepCopyInto(out *PodStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodStatus.
func (in *PodStatus) DeepCopy() *PodStatus {
	if in == nil {
		return nil
	}
	out := new(PodStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RBACSpec) DeepCopyInto(out *RBACSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RBACSpec.
func (in *RBACSpec) DeepCopy() *RBACSpec {
	if in == nil {
		return nil
	}
	out := new(RBACSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchedulerSpec) DeepCopyInto(out *SchedulerSpec) {
	*out = *in
	if in.Common != nil {
		in, out := &in.Common, &out.Common
		*out = new(PodCommonSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Scheduler != nil {
		in, out := &in.Scheduler, &out.Scheduler
		*out = new(ContainerCommonSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchedulerSpec.
func (in *SchedulerSpec) DeepCopy() *SchedulerSpec {
	if in == nil {
		return nil
	}
	out := new(SchedulerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchedulerStatus) DeepCopyInto(out *SchedulerStatus) {
	*out = *in
	if in.Instances != nil {
		in, out := &in.Instances, &out.Instances
		*out = new(DeployStatus)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchedulerStatus.
func (in *SchedulerStatus) DeepCopy() *SchedulerStatus {
	if in == nil {
		return nil
	}
	out := new(SchedulerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnapshotSpec) DeepCopyInto(out *SnapshotSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnapshotSpec.
func (in *SnapshotSpec) DeepCopy() *SnapshotSpec {
	if in == nil {
		return nil
	}
	out := new(SnapshotSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageClassSpec) DeepCopyInto(out *StorageClassSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageClassSpec.
func (in *StorageClassSpec) DeepCopy() *StorageClassSpec {
	if in == nil {
		return nil
	}
	out := new(StorageClassSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UISpec) DeepCopyInto(out *UISpec) {
	*out = *in
	if in.Common != nil {
		in, out := &in.Common, &out.Common
		*out = new(PodCommonSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.UI != nil {
		in, out := &in.UI, &out.UI
		*out = new(ContainerCommonSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UISpec.
func (in *UISpec) DeepCopy() *UISpec {
	if in == nil {
		return nil
	}
	out := new(UISpec)
	in.DeepCopyInto(out)
	return out
}
