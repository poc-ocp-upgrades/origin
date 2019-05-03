package core

import (
 v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 runtime "k8s.io/apimachinery/pkg/runtime"
 types "k8s.io/apimachinery/pkg/types"
)

func (in *AWSElasticBlockStoreVolumeSource) DeepCopyInto(out *AWSElasticBlockStoreVolumeSource) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 return
}
func (in *AWSElasticBlockStoreVolumeSource) DeepCopy() *AWSElasticBlockStoreVolumeSource {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(AWSElasticBlockStoreVolumeSource)
 in.DeepCopyInto(out)
 return out
}
func (in *Affinity) DeepCopyInto(out *Affinity) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.NodeAffinity != nil {
  in, out := &in.NodeAffinity, &out.NodeAffinity
  *out = new(NodeAffinity)
  (*in).DeepCopyInto(*out)
 }
 if in.PodAffinity != nil {
  in, out := &in.PodAffinity, &out.PodAffinity
  *out = new(PodAffinity)
  (*in).DeepCopyInto(*out)
 }
 if in.PodAntiAffinity != nil {
  in, out := &in.PodAntiAffinity, &out.PodAntiAffinity
  *out = new(PodAntiAffinity)
  (*in).DeepCopyInto(*out)
 }
 return
}
func (in *Affinity) DeepCopy() *Affinity {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(Affinity)
 in.DeepCopyInto(out)
 return out
}
func (in *AttachedVolume) DeepCopyInto(out *AttachedVolume) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 return
}
func (in *AttachedVolume) DeepCopy() *AttachedVolume {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(AttachedVolume)
 in.DeepCopyInto(out)
 return out
}
func (in *AvoidPods) DeepCopyInto(out *AvoidPods) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.PreferAvoidPods != nil {
  in, out := &in.PreferAvoidPods, &out.PreferAvoidPods
  *out = make([]PreferAvoidPodsEntry, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 return
}
func (in *AvoidPods) DeepCopy() *AvoidPods {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(AvoidPods)
 in.DeepCopyInto(out)
 return out
}
func (in *AzureDiskVolumeSource) DeepCopyInto(out *AzureDiskVolumeSource) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.CachingMode != nil {
  in, out := &in.CachingMode, &out.CachingMode
  *out = new(AzureDataDiskCachingMode)
  **out = **in
 }
 if in.FSType != nil {
  in, out := &in.FSType, &out.FSType
  *out = new(string)
  **out = **in
 }
 if in.ReadOnly != nil {
  in, out := &in.ReadOnly, &out.ReadOnly
  *out = new(bool)
  **out = **in
 }
 if in.Kind != nil {
  in, out := &in.Kind, &out.Kind
  *out = new(AzureDataDiskKind)
  **out = **in
 }
 return
}
func (in *AzureDiskVolumeSource) DeepCopy() *AzureDiskVolumeSource {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(AzureDiskVolumeSource)
 in.DeepCopyInto(out)
 return out
}
func (in *AzureFilePersistentVolumeSource) DeepCopyInto(out *AzureFilePersistentVolumeSource) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.SecretNamespace != nil {
  in, out := &in.SecretNamespace, &out.SecretNamespace
  *out = new(string)
  **out = **in
 }
 return
}
func (in *AzureFilePersistentVolumeSource) DeepCopy() *AzureFilePersistentVolumeSource {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(AzureFilePersistentVolumeSource)
 in.DeepCopyInto(out)
 return out
}
func (in *AzureFileVolumeSource) DeepCopyInto(out *AzureFileVolumeSource) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 return
}
func (in *AzureFileVolumeSource) DeepCopy() *AzureFileVolumeSource {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(AzureFileVolumeSource)
 in.DeepCopyInto(out)
 return out
}
func (in *Binding) DeepCopyInto(out *Binding) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
 out.Target = in.Target
 return
}
func (in *Binding) DeepCopy() *Binding {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(Binding)
 in.DeepCopyInto(out)
 return out
}
func (in *Binding) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *CSIPersistentVolumeSource) DeepCopyInto(out *CSIPersistentVolumeSource) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.VolumeAttributes != nil {
  in, out := &in.VolumeAttributes, &out.VolumeAttributes
  *out = make(map[string]string, len(*in))
  for key, val := range *in {
   (*out)[key] = val
  }
 }
 if in.ControllerPublishSecretRef != nil {
  in, out := &in.ControllerPublishSecretRef, &out.ControllerPublishSecretRef
  *out = new(SecretReference)
  **out = **in
 }
 if in.NodeStageSecretRef != nil {
  in, out := &in.NodeStageSecretRef, &out.NodeStageSecretRef
  *out = new(SecretReference)
  **out = **in
 }
 if in.NodePublishSecretRef != nil {
  in, out := &in.NodePublishSecretRef, &out.NodePublishSecretRef
  *out = new(SecretReference)
  **out = **in
 }
 return
}
func (in *CSIPersistentVolumeSource) DeepCopy() *CSIPersistentVolumeSource {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(CSIPersistentVolumeSource)
 in.DeepCopyInto(out)
 return out
}
func (in *Capabilities) DeepCopyInto(out *Capabilities) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Add != nil {
  in, out := &in.Add, &out.Add
  *out = make([]Capability, len(*in))
  copy(*out, *in)
 }
 if in.Drop != nil {
  in, out := &in.Drop, &out.Drop
  *out = make([]Capability, len(*in))
  copy(*out, *in)
 }
 return
}
func (in *Capabilities) DeepCopy() *Capabilities {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(Capabilities)
 in.DeepCopyInto(out)
 return out
}
func (in *CephFSPersistentVolumeSource) DeepCopyInto(out *CephFSPersistentVolumeSource) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Monitors != nil {
  in, out := &in.Monitors, &out.Monitors
  *out = make([]string, len(*in))
  copy(*out, *in)
 }
 if in.SecretRef != nil {
  in, out := &in.SecretRef, &out.SecretRef
  *out = new(SecretReference)
  **out = **in
 }
 return
}
func (in *CephFSPersistentVolumeSource) DeepCopy() *CephFSPersistentVolumeSource {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(CephFSPersistentVolumeSource)
 in.DeepCopyInto(out)
 return out
}
func (in *CephFSVolumeSource) DeepCopyInto(out *CephFSVolumeSource) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Monitors != nil {
  in, out := &in.Monitors, &out.Monitors
  *out = make([]string, len(*in))
  copy(*out, *in)
 }
 if in.SecretRef != nil {
  in, out := &in.SecretRef, &out.SecretRef
  *out = new(LocalObjectReference)
  **out = **in
 }
 return
}
func (in *CephFSVolumeSource) DeepCopy() *CephFSVolumeSource {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(CephFSVolumeSource)
 in.DeepCopyInto(out)
 return out
}
func (in *CinderPersistentVolumeSource) DeepCopyInto(out *CinderPersistentVolumeSource) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.SecretRef != nil {
  in, out := &in.SecretRef, &out.SecretRef
  *out = new(SecretReference)
  **out = **in
 }
 return
}
func (in *CinderPersistentVolumeSource) DeepCopy() *CinderPersistentVolumeSource {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(CinderPersistentVolumeSource)
 in.DeepCopyInto(out)
 return out
}
func (in *CinderVolumeSource) DeepCopyInto(out *CinderVolumeSource) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.SecretRef != nil {
  in, out := &in.SecretRef, &out.SecretRef
  *out = new(LocalObjectReference)
  **out = **in
 }
 return
}
func (in *CinderVolumeSource) DeepCopy() *CinderVolumeSource {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(CinderVolumeSource)
 in.DeepCopyInto(out)
 return out
}
func (in *ClientIPConfig) DeepCopyInto(out *ClientIPConfig) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.TimeoutSeconds != nil {
  in, out := &in.TimeoutSeconds, &out.TimeoutSeconds
  *out = new(int32)
  **out = **in
 }
 return
}
func (in *ClientIPConfig) DeepCopy() *ClientIPConfig {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ClientIPConfig)
 in.DeepCopyInto(out)
 return out
}
func (in *ComponentCondition) DeepCopyInto(out *ComponentCondition) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 return
}
func (in *ComponentCondition) DeepCopy() *ComponentCondition {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ComponentCondition)
 in.DeepCopyInto(out)
 return out
}
func (in *ComponentStatus) DeepCopyInto(out *ComponentStatus) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
 if in.Conditions != nil {
  in, out := &in.Conditions, &out.Conditions
  *out = make([]ComponentCondition, len(*in))
  copy(*out, *in)
 }
 return
}
func (in *ComponentStatus) DeepCopy() *ComponentStatus {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ComponentStatus)
 in.DeepCopyInto(out)
 return out
}
func (in *ComponentStatus) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *ComponentStatusList) DeepCopyInto(out *ComponentStatusList) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]ComponentStatus, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 return
}
func (in *ComponentStatusList) DeepCopy() *ComponentStatusList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ComponentStatusList)
 in.DeepCopyInto(out)
 return out
}
func (in *ComponentStatusList) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *ConfigMap) DeepCopyInto(out *ConfigMap) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
 if in.Data != nil {
  in, out := &in.Data, &out.Data
  *out = make(map[string]string, len(*in))
  for key, val := range *in {
   (*out)[key] = val
  }
 }
 if in.BinaryData != nil {
  in, out := &in.BinaryData, &out.BinaryData
  *out = make(map[string][]byte, len(*in))
  for key, val := range *in {
   var outVal []byte
   if val == nil {
    (*out)[key] = nil
   } else {
    in, out := &val, &outVal
    *out = make([]byte, len(*in))
    copy(*out, *in)
   }
   (*out)[key] = outVal
  }
 }
 return
}
func (in *ConfigMap) DeepCopy() *ConfigMap {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ConfigMap)
 in.DeepCopyInto(out)
 return out
}
func (in *ConfigMap) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *ConfigMapEnvSource) DeepCopyInto(out *ConfigMapEnvSource) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.LocalObjectReference = in.LocalObjectReference
 if in.Optional != nil {
  in, out := &in.Optional, &out.Optional
  *out = new(bool)
  **out = **in
 }
 return
}
func (in *ConfigMapEnvSource) DeepCopy() *ConfigMapEnvSource {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ConfigMapEnvSource)
 in.DeepCopyInto(out)
 return out
}
func (in *ConfigMapKeySelector) DeepCopyInto(out *ConfigMapKeySelector) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.LocalObjectReference = in.LocalObjectReference
 if in.Optional != nil {
  in, out := &in.Optional, &out.Optional
  *out = new(bool)
  **out = **in
 }
 return
}
func (in *ConfigMapKeySelector) DeepCopy() *ConfigMapKeySelector {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ConfigMapKeySelector)
 in.DeepCopyInto(out)
 return out
}
func (in *ConfigMapList) DeepCopyInto(out *ConfigMapList) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]ConfigMap, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 return
}
func (in *ConfigMapList) DeepCopy() *ConfigMapList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ConfigMapList)
 in.DeepCopyInto(out)
 return out
}
func (in *ConfigMapList) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *ConfigMapNodeConfigSource) DeepCopyInto(out *ConfigMapNodeConfigSource) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 return
}
func (in *ConfigMapNodeConfigSource) DeepCopy() *ConfigMapNodeConfigSource {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ConfigMapNodeConfigSource)
 in.DeepCopyInto(out)
 return out
}
func (in *ConfigMapProjection) DeepCopyInto(out *ConfigMapProjection) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.LocalObjectReference = in.LocalObjectReference
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]KeyToPath, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 if in.Optional != nil {
  in, out := &in.Optional, &out.Optional
  *out = new(bool)
  **out = **in
 }
 return
}
func (in *ConfigMapProjection) DeepCopy() *ConfigMapProjection {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ConfigMapProjection)
 in.DeepCopyInto(out)
 return out
}
func (in *ConfigMapVolumeSource) DeepCopyInto(out *ConfigMapVolumeSource) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.LocalObjectReference = in.LocalObjectReference
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]KeyToPath, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 if in.DefaultMode != nil {
  in, out := &in.DefaultMode, &out.DefaultMode
  *out = new(int32)
  **out = **in
 }
 if in.Optional != nil {
  in, out := &in.Optional, &out.Optional
  *out = new(bool)
  **out = **in
 }
 return
}
func (in *ConfigMapVolumeSource) DeepCopy() *ConfigMapVolumeSource {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ConfigMapVolumeSource)
 in.DeepCopyInto(out)
 return out
}
func (in *Container) DeepCopyInto(out *Container) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Command != nil {
  in, out := &in.Command, &out.Command
  *out = make([]string, len(*in))
  copy(*out, *in)
 }
 if in.Args != nil {
  in, out := &in.Args, &out.Args
  *out = make([]string, len(*in))
  copy(*out, *in)
 }
 if in.Ports != nil {
  in, out := &in.Ports, &out.Ports
  *out = make([]ContainerPort, len(*in))
  copy(*out, *in)
 }
 if in.EnvFrom != nil {
  in, out := &in.EnvFrom, &out.EnvFrom
  *out = make([]EnvFromSource, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 if in.Env != nil {
  in, out := &in.Env, &out.Env
  *out = make([]EnvVar, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 in.Resources.DeepCopyInto(&out.Resources)
 if in.VolumeMounts != nil {
  in, out := &in.VolumeMounts, &out.VolumeMounts
  *out = make([]VolumeMount, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 if in.VolumeDevices != nil {
  in, out := &in.VolumeDevices, &out.VolumeDevices
  *out = make([]VolumeDevice, len(*in))
  copy(*out, *in)
 }
 if in.LivenessProbe != nil {
  in, out := &in.LivenessProbe, &out.LivenessProbe
  *out = new(Probe)
  (*in).DeepCopyInto(*out)
 }
 if in.ReadinessProbe != nil {
  in, out := &in.ReadinessProbe, &out.ReadinessProbe
  *out = new(Probe)
  (*in).DeepCopyInto(*out)
 }
 if in.Lifecycle != nil {
  in, out := &in.Lifecycle, &out.Lifecycle
  *out = new(Lifecycle)
  (*in).DeepCopyInto(*out)
 }
 if in.SecurityContext != nil {
  in, out := &in.SecurityContext, &out.SecurityContext
  *out = new(SecurityContext)
  (*in).DeepCopyInto(*out)
 }
 return
}
func (in *Container) DeepCopy() *Container {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(Container)
 in.DeepCopyInto(out)
 return out
}
func (in *ContainerImage) DeepCopyInto(out *ContainerImage) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Names != nil {
  in, out := &in.Names, &out.Names
  *out = make([]string, len(*in))
  copy(*out, *in)
 }
 return
}
func (in *ContainerImage) DeepCopy() *ContainerImage {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ContainerImage)
 in.DeepCopyInto(out)
 return out
}
func (in *ContainerPort) DeepCopyInto(out *ContainerPort) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 return
}
func (in *ContainerPort) DeepCopy() *ContainerPort {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ContainerPort)
 in.DeepCopyInto(out)
 return out
}
func (in *ContainerState) DeepCopyInto(out *ContainerState) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Waiting != nil {
  in, out := &in.Waiting, &out.Waiting
  *out = new(ContainerStateWaiting)
  **out = **in
 }
 if in.Running != nil {
  in, out := &in.Running, &out.Running
  *out = new(ContainerStateRunning)
  (*in).DeepCopyInto(*out)
 }
 if in.Terminated != nil {
  in, out := &in.Terminated, &out.Terminated
  *out = new(ContainerStateTerminated)
  (*in).DeepCopyInto(*out)
 }
 return
}
func (in *ContainerState) DeepCopy() *ContainerState {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ContainerState)
 in.DeepCopyInto(out)
 return out
}
func (in *ContainerStateRunning) DeepCopyInto(out *ContainerStateRunning) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 in.StartedAt.DeepCopyInto(&out.StartedAt)
 return
}
func (in *ContainerStateRunning) DeepCopy() *ContainerStateRunning {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ContainerStateRunning)
 in.DeepCopyInto(out)
 return out
}
func (in *ContainerStateTerminated) DeepCopyInto(out *ContainerStateTerminated) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 in.StartedAt.DeepCopyInto(&out.StartedAt)
 in.FinishedAt.DeepCopyInto(&out.FinishedAt)
 return
}
func (in *ContainerStateTerminated) DeepCopy() *ContainerStateTerminated {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ContainerStateTerminated)
 in.DeepCopyInto(out)
 return out
}
func (in *ContainerStateWaiting) DeepCopyInto(out *ContainerStateWaiting) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 return
}
func (in *ContainerStateWaiting) DeepCopy() *ContainerStateWaiting {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ContainerStateWaiting)
 in.DeepCopyInto(out)
 return out
}
func (in *ContainerStatus) DeepCopyInto(out *ContainerStatus) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 in.State.DeepCopyInto(&out.State)
 in.LastTerminationState.DeepCopyInto(&out.LastTerminationState)
 return
}
func (in *ContainerStatus) DeepCopy() *ContainerStatus {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ContainerStatus)
 in.DeepCopyInto(out)
 return out
}
func (in *DaemonEndpoint) DeepCopyInto(out *DaemonEndpoint) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 return
}
func (in *DaemonEndpoint) DeepCopy() *DaemonEndpoint {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(DaemonEndpoint)
 in.DeepCopyInto(out)
 return out
}
func (in *DownwardAPIProjection) DeepCopyInto(out *DownwardAPIProjection) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]DownwardAPIVolumeFile, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 return
}
func (in *DownwardAPIProjection) DeepCopy() *DownwardAPIProjection {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(DownwardAPIProjection)
 in.DeepCopyInto(out)
 return out
}
func (in *DownwardAPIVolumeFile) DeepCopyInto(out *DownwardAPIVolumeFile) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.FieldRef != nil {
  in, out := &in.FieldRef, &out.FieldRef
  *out = new(ObjectFieldSelector)
  **out = **in
 }
 if in.ResourceFieldRef != nil {
  in, out := &in.ResourceFieldRef, &out.ResourceFieldRef
  *out = new(ResourceFieldSelector)
  (*in).DeepCopyInto(*out)
 }
 if in.Mode != nil {
  in, out := &in.Mode, &out.Mode
  *out = new(int32)
  **out = **in
 }
 return
}
func (in *DownwardAPIVolumeFile) DeepCopy() *DownwardAPIVolumeFile {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(DownwardAPIVolumeFile)
 in.DeepCopyInto(out)
 return out
}
func (in *DownwardAPIVolumeSource) DeepCopyInto(out *DownwardAPIVolumeSource) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]DownwardAPIVolumeFile, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 if in.DefaultMode != nil {
  in, out := &in.DefaultMode, &out.DefaultMode
  *out = new(int32)
  **out = **in
 }
 return
}
func (in *DownwardAPIVolumeSource) DeepCopy() *DownwardAPIVolumeSource {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(DownwardAPIVolumeSource)
 in.DeepCopyInto(out)
 return out
}
func (in *EmptyDirVolumeSource) DeepCopyInto(out *EmptyDirVolumeSource) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.SizeLimit != nil {
  in, out := &in.SizeLimit, &out.SizeLimit
  x := (*in).DeepCopy()
  *out = &x
 }
 return
}
func (in *EmptyDirVolumeSource) DeepCopy() *EmptyDirVolumeSource {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(EmptyDirVolumeSource)
 in.DeepCopyInto(out)
 return out
}
func (in *EndpointAddress) DeepCopyInto(out *EndpointAddress) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.NodeName != nil {
  in, out := &in.NodeName, &out.NodeName
  *out = new(string)
  **out = **in
 }
 if in.TargetRef != nil {
  in, out := &in.TargetRef, &out.TargetRef
  *out = new(ObjectReference)
  **out = **in
 }
 return
}
func (in *EndpointAddress) DeepCopy() *EndpointAddress {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(EndpointAddress)
 in.DeepCopyInto(out)
 return out
}
func (in *EndpointPort) DeepCopyInto(out *EndpointPort) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 return
}
func (in *EndpointPort) DeepCopy() *EndpointPort {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(EndpointPort)
 in.DeepCopyInto(out)
 return out
}
func (in *EndpointSubset) DeepCopyInto(out *EndpointSubset) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Addresses != nil {
  in, out := &in.Addresses, &out.Addresses
  *out = make([]EndpointAddress, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 if in.NotReadyAddresses != nil {
  in, out := &in.NotReadyAddresses, &out.NotReadyAddresses
  *out = make([]EndpointAddress, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 if in.Ports != nil {
  in, out := &in.Ports, &out.Ports
  *out = make([]EndpointPort, len(*in))
  copy(*out, *in)
 }
 return
}
func (in *EndpointSubset) DeepCopy() *EndpointSubset {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(EndpointSubset)
 in.DeepCopyInto(out)
 return out
}
func (in *Endpoints) DeepCopyInto(out *Endpoints) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
 if in.Subsets != nil {
  in, out := &in.Subsets, &out.Subsets
  *out = make([]EndpointSubset, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 return
}
func (in *Endpoints) DeepCopy() *Endpoints {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(Endpoints)
 in.DeepCopyInto(out)
 return out
}
func (in *Endpoints) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *EndpointsList) DeepCopyInto(out *EndpointsList) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]Endpoints, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 return
}
func (in *EndpointsList) DeepCopy() *EndpointsList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(EndpointsList)
 in.DeepCopyInto(out)
 return out
}
func (in *EndpointsList) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *EnvFromSource) DeepCopyInto(out *EnvFromSource) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.ConfigMapRef != nil {
  in, out := &in.ConfigMapRef, &out.ConfigMapRef
  *out = new(ConfigMapEnvSource)
  (*in).DeepCopyInto(*out)
 }
 if in.SecretRef != nil {
  in, out := &in.SecretRef, &out.SecretRef
  *out = new(SecretEnvSource)
  (*in).DeepCopyInto(*out)
 }
 return
}
func (in *EnvFromSource) DeepCopy() *EnvFromSource {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(EnvFromSource)
 in.DeepCopyInto(out)
 return out
}
func (in *EnvVar) DeepCopyInto(out *EnvVar) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.ValueFrom != nil {
  in, out := &in.ValueFrom, &out.ValueFrom
  *out = new(EnvVarSource)
  (*in).DeepCopyInto(*out)
 }
 return
}
func (in *EnvVar) DeepCopy() *EnvVar {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(EnvVar)
 in.DeepCopyInto(out)
 return out
}
func (in *EnvVarSource) DeepCopyInto(out *EnvVarSource) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.FieldRef != nil {
  in, out := &in.FieldRef, &out.FieldRef
  *out = new(ObjectFieldSelector)
  **out = **in
 }
 if in.ResourceFieldRef != nil {
  in, out := &in.ResourceFieldRef, &out.ResourceFieldRef
  *out = new(ResourceFieldSelector)
  (*in).DeepCopyInto(*out)
 }
 if in.ConfigMapKeyRef != nil {
  in, out := &in.ConfigMapKeyRef, &out.ConfigMapKeyRef
  *out = new(ConfigMapKeySelector)
  (*in).DeepCopyInto(*out)
 }
 if in.SecretKeyRef != nil {
  in, out := &in.SecretKeyRef, &out.SecretKeyRef
  *out = new(SecretKeySelector)
  (*in).DeepCopyInto(*out)
 }
 return
}
func (in *EnvVarSource) DeepCopy() *EnvVarSource {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(EnvVarSource)
 in.DeepCopyInto(out)
 return out
}
func (in *Event) DeepCopyInto(out *Event) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
 out.InvolvedObject = in.InvolvedObject
 out.Source = in.Source
 in.FirstTimestamp.DeepCopyInto(&out.FirstTimestamp)
 in.LastTimestamp.DeepCopyInto(&out.LastTimestamp)
 in.EventTime.DeepCopyInto(&out.EventTime)
 if in.Series != nil {
  in, out := &in.Series, &out.Series
  *out = new(EventSeries)
  (*in).DeepCopyInto(*out)
 }
 if in.Related != nil {
  in, out := &in.Related, &out.Related
  *out = new(ObjectReference)
  **out = **in
 }
 return
}
func (in *Event) DeepCopy() *Event {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(Event)
 in.DeepCopyInto(out)
 return out
}
func (in *Event) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *EventList) DeepCopyInto(out *EventList) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]Event, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 return
}
func (in *EventList) DeepCopy() *EventList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(EventList)
 in.DeepCopyInto(out)
 return out
}
func (in *EventList) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *EventSeries) DeepCopyInto(out *EventSeries) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 in.LastObservedTime.DeepCopyInto(&out.LastObservedTime)
 return
}
func (in *EventSeries) DeepCopy() *EventSeries {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(EventSeries)
 in.DeepCopyInto(out)
 return out
}
func (in *EventSource) DeepCopyInto(out *EventSource) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 return
}
func (in *EventSource) DeepCopy() *EventSource {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(EventSource)
 in.DeepCopyInto(out)
 return out
}
func (in *ExecAction) DeepCopyInto(out *ExecAction) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Command != nil {
  in, out := &in.Command, &out.Command
  *out = make([]string, len(*in))
  copy(*out, *in)
 }
 return
}
func (in *ExecAction) DeepCopy() *ExecAction {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ExecAction)
 in.DeepCopyInto(out)
 return out
}
func (in *FCVolumeSource) DeepCopyInto(out *FCVolumeSource) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.TargetWWNs != nil {
  in, out := &in.TargetWWNs, &out.TargetWWNs
  *out = make([]string, len(*in))
  copy(*out, *in)
 }
 if in.Lun != nil {
  in, out := &in.Lun, &out.Lun
  *out = new(int32)
  **out = **in
 }
 if in.WWIDs != nil {
  in, out := &in.WWIDs, &out.WWIDs
  *out = make([]string, len(*in))
  copy(*out, *in)
 }
 return
}
func (in *FCVolumeSource) DeepCopy() *FCVolumeSource {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(FCVolumeSource)
 in.DeepCopyInto(out)
 return out
}
func (in *FlexPersistentVolumeSource) DeepCopyInto(out *FlexPersistentVolumeSource) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.SecretRef != nil {
  in, out := &in.SecretRef, &out.SecretRef
  *out = new(SecretReference)
  **out = **in
 }
 if in.Options != nil {
  in, out := &in.Options, &out.Options
  *out = make(map[string]string, len(*in))
  for key, val := range *in {
   (*out)[key] = val
  }
 }
 return
}
func (in *FlexPersistentVolumeSource) DeepCopy() *FlexPersistentVolumeSource {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(FlexPersistentVolumeSource)
 in.DeepCopyInto(out)
 return out
}
func (in *FlexVolumeSource) DeepCopyInto(out *FlexVolumeSource) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.SecretRef != nil {
  in, out := &in.SecretRef, &out.SecretRef
  *out = new(LocalObjectReference)
  **out = **in
 }
 if in.Options != nil {
  in, out := &in.Options, &out.Options
  *out = make(map[string]string, len(*in))
  for key, val := range *in {
   (*out)[key] = val
  }
 }
 return
}
func (in *FlexVolumeSource) DeepCopy() *FlexVolumeSource {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(FlexVolumeSource)
 in.DeepCopyInto(out)
 return out
}
func (in *FlockerVolumeSource) DeepCopyInto(out *FlockerVolumeSource) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 return
}
func (in *FlockerVolumeSource) DeepCopy() *FlockerVolumeSource {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(FlockerVolumeSource)
 in.DeepCopyInto(out)
 return out
}
func (in *GCEPersistentDiskVolumeSource) DeepCopyInto(out *GCEPersistentDiskVolumeSource) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 return
}
func (in *GCEPersistentDiskVolumeSource) DeepCopy() *GCEPersistentDiskVolumeSource {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(GCEPersistentDiskVolumeSource)
 in.DeepCopyInto(out)
 return out
}
func (in *GitRepoVolumeSource) DeepCopyInto(out *GitRepoVolumeSource) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 return
}
func (in *GitRepoVolumeSource) DeepCopy() *GitRepoVolumeSource {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(GitRepoVolumeSource)
 in.DeepCopyInto(out)
 return out
}
func (in *GlusterfsPersistentVolumeSource) DeepCopyInto(out *GlusterfsPersistentVolumeSource) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.EndpointsNamespace != nil {
  in, out := &in.EndpointsNamespace, &out.EndpointsNamespace
  *out = new(string)
  **out = **in
 }
 return
}
func (in *GlusterfsPersistentVolumeSource) DeepCopy() *GlusterfsPersistentVolumeSource {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(GlusterfsPersistentVolumeSource)
 in.DeepCopyInto(out)
 return out
}
func (in *GlusterfsVolumeSource) DeepCopyInto(out *GlusterfsVolumeSource) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 return
}
func (in *GlusterfsVolumeSource) DeepCopy() *GlusterfsVolumeSource {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(GlusterfsVolumeSource)
 in.DeepCopyInto(out)
 return out
}
func (in *HTTPGetAction) DeepCopyInto(out *HTTPGetAction) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.Port = in.Port
 if in.HTTPHeaders != nil {
  in, out := &in.HTTPHeaders, &out.HTTPHeaders
  *out = make([]HTTPHeader, len(*in))
  copy(*out, *in)
 }
 return
}
func (in *HTTPGetAction) DeepCopy() *HTTPGetAction {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(HTTPGetAction)
 in.DeepCopyInto(out)
 return out
}
func (in *HTTPHeader) DeepCopyInto(out *HTTPHeader) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 return
}
func (in *HTTPHeader) DeepCopy() *HTTPHeader {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(HTTPHeader)
 in.DeepCopyInto(out)
 return out
}
func (in *Handler) DeepCopyInto(out *Handler) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Exec != nil {
  in, out := &in.Exec, &out.Exec
  *out = new(ExecAction)
  (*in).DeepCopyInto(*out)
 }
 if in.HTTPGet != nil {
  in, out := &in.HTTPGet, &out.HTTPGet
  *out = new(HTTPGetAction)
  (*in).DeepCopyInto(*out)
 }
 if in.TCPSocket != nil {
  in, out := &in.TCPSocket, &out.TCPSocket
  *out = new(TCPSocketAction)
  **out = **in
 }
 return
}
func (in *Handler) DeepCopy() *Handler {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(Handler)
 in.DeepCopyInto(out)
 return out
}
func (in *HostAlias) DeepCopyInto(out *HostAlias) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Hostnames != nil {
  in, out := &in.Hostnames, &out.Hostnames
  *out = make([]string, len(*in))
  copy(*out, *in)
 }
 return
}
func (in *HostAlias) DeepCopy() *HostAlias {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(HostAlias)
 in.DeepCopyInto(out)
 return out
}
func (in *HostPathVolumeSource) DeepCopyInto(out *HostPathVolumeSource) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Type != nil {
  in, out := &in.Type, &out.Type
  *out = new(HostPathType)
  **out = **in
 }
 return
}
func (in *HostPathVolumeSource) DeepCopy() *HostPathVolumeSource {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(HostPathVolumeSource)
 in.DeepCopyInto(out)
 return out
}
func (in *ISCSIPersistentVolumeSource) DeepCopyInto(out *ISCSIPersistentVolumeSource) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Portals != nil {
  in, out := &in.Portals, &out.Portals
  *out = make([]string, len(*in))
  copy(*out, *in)
 }
 if in.SecretRef != nil {
  in, out := &in.SecretRef, &out.SecretRef
  *out = new(SecretReference)
  **out = **in
 }
 if in.InitiatorName != nil {
  in, out := &in.InitiatorName, &out.InitiatorName
  *out = new(string)
  **out = **in
 }
 return
}
func (in *ISCSIPersistentVolumeSource) DeepCopy() *ISCSIPersistentVolumeSource {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ISCSIPersistentVolumeSource)
 in.DeepCopyInto(out)
 return out
}
func (in *ISCSIVolumeSource) DeepCopyInto(out *ISCSIVolumeSource) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Portals != nil {
  in, out := &in.Portals, &out.Portals
  *out = make([]string, len(*in))
  copy(*out, *in)
 }
 if in.SecretRef != nil {
  in, out := &in.SecretRef, &out.SecretRef
  *out = new(LocalObjectReference)
  **out = **in
 }
 if in.InitiatorName != nil {
  in, out := &in.InitiatorName, &out.InitiatorName
  *out = new(string)
  **out = **in
 }
 return
}
func (in *ISCSIVolumeSource) DeepCopy() *ISCSIVolumeSource {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ISCSIVolumeSource)
 in.DeepCopyInto(out)
 return out
}
func (in *KeyToPath) DeepCopyInto(out *KeyToPath) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Mode != nil {
  in, out := &in.Mode, &out.Mode
  *out = new(int32)
  **out = **in
 }
 return
}
func (in *KeyToPath) DeepCopy() *KeyToPath {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(KeyToPath)
 in.DeepCopyInto(out)
 return out
}
func (in *Lifecycle) DeepCopyInto(out *Lifecycle) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.PostStart != nil {
  in, out := &in.PostStart, &out.PostStart
  *out = new(Handler)
  (*in).DeepCopyInto(*out)
 }
 if in.PreStop != nil {
  in, out := &in.PreStop, &out.PreStop
  *out = new(Handler)
  (*in).DeepCopyInto(*out)
 }
 return
}
func (in *Lifecycle) DeepCopy() *Lifecycle {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(Lifecycle)
 in.DeepCopyInto(out)
 return out
}
func (in *LimitRange) DeepCopyInto(out *LimitRange) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
 in.Spec.DeepCopyInto(&out.Spec)
 return
}
func (in *LimitRange) DeepCopy() *LimitRange {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(LimitRange)
 in.DeepCopyInto(out)
 return out
}
func (in *LimitRange) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *LimitRangeItem) DeepCopyInto(out *LimitRangeItem) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Max != nil {
  in, out := &in.Max, &out.Max
  *out = make(ResourceList, len(*in))
  for key, val := range *in {
   (*out)[key] = val.DeepCopy()
  }
 }
 if in.Min != nil {
  in, out := &in.Min, &out.Min
  *out = make(ResourceList, len(*in))
  for key, val := range *in {
   (*out)[key] = val.DeepCopy()
  }
 }
 if in.Default != nil {
  in, out := &in.Default, &out.Default
  *out = make(ResourceList, len(*in))
  for key, val := range *in {
   (*out)[key] = val.DeepCopy()
  }
 }
 if in.DefaultRequest != nil {
  in, out := &in.DefaultRequest, &out.DefaultRequest
  *out = make(ResourceList, len(*in))
  for key, val := range *in {
   (*out)[key] = val.DeepCopy()
  }
 }
 if in.MaxLimitRequestRatio != nil {
  in, out := &in.MaxLimitRequestRatio, &out.MaxLimitRequestRatio
  *out = make(ResourceList, len(*in))
  for key, val := range *in {
   (*out)[key] = val.DeepCopy()
  }
 }
 return
}
func (in *LimitRangeItem) DeepCopy() *LimitRangeItem {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(LimitRangeItem)
 in.DeepCopyInto(out)
 return out
}
func (in *LimitRangeList) DeepCopyInto(out *LimitRangeList) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]LimitRange, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 return
}
func (in *LimitRangeList) DeepCopy() *LimitRangeList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(LimitRangeList)
 in.DeepCopyInto(out)
 return out
}
func (in *LimitRangeList) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *LimitRangeSpec) DeepCopyInto(out *LimitRangeSpec) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Limits != nil {
  in, out := &in.Limits, &out.Limits
  *out = make([]LimitRangeItem, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 return
}
func (in *LimitRangeSpec) DeepCopy() *LimitRangeSpec {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(LimitRangeSpec)
 in.DeepCopyInto(out)
 return out
}
func (in *List) DeepCopyInto(out *List) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]runtime.Object, len(*in))
  for i := range *in {
   if (*in)[i] != nil {
    (*out)[i] = (*in)[i].DeepCopyObject()
   }
  }
 }
 return
}
func (in *List) DeepCopy() *List {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(List)
 in.DeepCopyInto(out)
 return out
}
func (in *List) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *LoadBalancerIngress) DeepCopyInto(out *LoadBalancerIngress) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 return
}
func (in *LoadBalancerIngress) DeepCopy() *LoadBalancerIngress {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(LoadBalancerIngress)
 in.DeepCopyInto(out)
 return out
}
func (in *LoadBalancerStatus) DeepCopyInto(out *LoadBalancerStatus) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Ingress != nil {
  in, out := &in.Ingress, &out.Ingress
  *out = make([]LoadBalancerIngress, len(*in))
  copy(*out, *in)
 }
 return
}
func (in *LoadBalancerStatus) DeepCopy() *LoadBalancerStatus {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(LoadBalancerStatus)
 in.DeepCopyInto(out)
 return out
}
func (in *LocalObjectReference) DeepCopyInto(out *LocalObjectReference) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 return
}
func (in *LocalObjectReference) DeepCopy() *LocalObjectReference {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(LocalObjectReference)
 in.DeepCopyInto(out)
 return out
}
func (in *LocalVolumeSource) DeepCopyInto(out *LocalVolumeSource) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.FSType != nil {
  in, out := &in.FSType, &out.FSType
  *out = new(string)
  **out = **in
 }
 return
}
func (in *LocalVolumeSource) DeepCopy() *LocalVolumeSource {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(LocalVolumeSource)
 in.DeepCopyInto(out)
 return out
}
func (in *NFSVolumeSource) DeepCopyInto(out *NFSVolumeSource) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 return
}
func (in *NFSVolumeSource) DeepCopy() *NFSVolumeSource {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(NFSVolumeSource)
 in.DeepCopyInto(out)
 return out
}
func (in *Namespace) DeepCopyInto(out *Namespace) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
 in.Spec.DeepCopyInto(&out.Spec)
 out.Status = in.Status
 return
}
func (in *Namespace) DeepCopy() *Namespace {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(Namespace)
 in.DeepCopyInto(out)
 return out
}
func (in *Namespace) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *NamespaceList) DeepCopyInto(out *NamespaceList) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]Namespace, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 return
}
func (in *NamespaceList) DeepCopy() *NamespaceList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(NamespaceList)
 in.DeepCopyInto(out)
 return out
}
func (in *NamespaceList) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *NamespaceSpec) DeepCopyInto(out *NamespaceSpec) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Finalizers != nil {
  in, out := &in.Finalizers, &out.Finalizers
  *out = make([]FinalizerName, len(*in))
  copy(*out, *in)
 }
 return
}
func (in *NamespaceSpec) DeepCopy() *NamespaceSpec {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(NamespaceSpec)
 in.DeepCopyInto(out)
 return out
}
func (in *NamespaceStatus) DeepCopyInto(out *NamespaceStatus) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 return
}
func (in *NamespaceStatus) DeepCopy() *NamespaceStatus {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(NamespaceStatus)
 in.DeepCopyInto(out)
 return out
}
func (in *Node) DeepCopyInto(out *Node) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
 in.Spec.DeepCopyInto(&out.Spec)
 in.Status.DeepCopyInto(&out.Status)
 return
}
func (in *Node) DeepCopy() *Node {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(Node)
 in.DeepCopyInto(out)
 return out
}
func (in *Node) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *NodeAddress) DeepCopyInto(out *NodeAddress) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 return
}
func (in *NodeAddress) DeepCopy() *NodeAddress {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(NodeAddress)
 in.DeepCopyInto(out)
 return out
}
func (in *NodeAffinity) DeepCopyInto(out *NodeAffinity) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.RequiredDuringSchedulingIgnoredDuringExecution != nil {
  in, out := &in.RequiredDuringSchedulingIgnoredDuringExecution, &out.RequiredDuringSchedulingIgnoredDuringExecution
  *out = new(NodeSelector)
  (*in).DeepCopyInto(*out)
 }
 if in.PreferredDuringSchedulingIgnoredDuringExecution != nil {
  in, out := &in.PreferredDuringSchedulingIgnoredDuringExecution, &out.PreferredDuringSchedulingIgnoredDuringExecution
  *out = make([]PreferredSchedulingTerm, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 return
}
func (in *NodeAffinity) DeepCopy() *NodeAffinity {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(NodeAffinity)
 in.DeepCopyInto(out)
 return out
}
func (in *NodeCondition) DeepCopyInto(out *NodeCondition) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 in.LastHeartbeatTime.DeepCopyInto(&out.LastHeartbeatTime)
 in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
 return
}
func (in *NodeCondition) DeepCopy() *NodeCondition {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(NodeCondition)
 in.DeepCopyInto(out)
 return out
}
func (in *NodeConfigSource) DeepCopyInto(out *NodeConfigSource) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.ConfigMap != nil {
  in, out := &in.ConfigMap, &out.ConfigMap
  *out = new(ConfigMapNodeConfigSource)
  **out = **in
 }
 return
}
func (in *NodeConfigSource) DeepCopy() *NodeConfigSource {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(NodeConfigSource)
 in.DeepCopyInto(out)
 return out
}
func (in *NodeConfigStatus) DeepCopyInto(out *NodeConfigStatus) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Assigned != nil {
  in, out := &in.Assigned, &out.Assigned
  *out = new(NodeConfigSource)
  (*in).DeepCopyInto(*out)
 }
 if in.Active != nil {
  in, out := &in.Active, &out.Active
  *out = new(NodeConfigSource)
  (*in).DeepCopyInto(*out)
 }
 if in.LastKnownGood != nil {
  in, out := &in.LastKnownGood, &out.LastKnownGood
  *out = new(NodeConfigSource)
  (*in).DeepCopyInto(*out)
 }
 return
}
func (in *NodeConfigStatus) DeepCopy() *NodeConfigStatus {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(NodeConfigStatus)
 in.DeepCopyInto(out)
 return out
}
func (in *NodeDaemonEndpoints) DeepCopyInto(out *NodeDaemonEndpoints) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.KubeletEndpoint = in.KubeletEndpoint
 return
}
func (in *NodeDaemonEndpoints) DeepCopy() *NodeDaemonEndpoints {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(NodeDaemonEndpoints)
 in.DeepCopyInto(out)
 return out
}
func (in *NodeList) DeepCopyInto(out *NodeList) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]Node, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 return
}
func (in *NodeList) DeepCopy() *NodeList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(NodeList)
 in.DeepCopyInto(out)
 return out
}
func (in *NodeList) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *NodeProxyOptions) DeepCopyInto(out *NodeProxyOptions) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 return
}
func (in *NodeProxyOptions) DeepCopy() *NodeProxyOptions {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(NodeProxyOptions)
 in.DeepCopyInto(out)
 return out
}
func (in *NodeProxyOptions) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *NodeResources) DeepCopyInto(out *NodeResources) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Capacity != nil {
  in, out := &in.Capacity, &out.Capacity
  *out = make(ResourceList, len(*in))
  for key, val := range *in {
   (*out)[key] = val.DeepCopy()
  }
 }
 return
}
func (in *NodeResources) DeepCopy() *NodeResources {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(NodeResources)
 in.DeepCopyInto(out)
 return out
}
func (in *NodeSelector) DeepCopyInto(out *NodeSelector) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.NodeSelectorTerms != nil {
  in, out := &in.NodeSelectorTerms, &out.NodeSelectorTerms
  *out = make([]NodeSelectorTerm, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 return
}
func (in *NodeSelector) DeepCopy() *NodeSelector {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(NodeSelector)
 in.DeepCopyInto(out)
 return out
}
func (in *NodeSelectorRequirement) DeepCopyInto(out *NodeSelectorRequirement) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Values != nil {
  in, out := &in.Values, &out.Values
  *out = make([]string, len(*in))
  copy(*out, *in)
 }
 return
}
func (in *NodeSelectorRequirement) DeepCopy() *NodeSelectorRequirement {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(NodeSelectorRequirement)
 in.DeepCopyInto(out)
 return out
}
func (in *NodeSelectorTerm) DeepCopyInto(out *NodeSelectorTerm) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.MatchExpressions != nil {
  in, out := &in.MatchExpressions, &out.MatchExpressions
  *out = make([]NodeSelectorRequirement, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 if in.MatchFields != nil {
  in, out := &in.MatchFields, &out.MatchFields
  *out = make([]NodeSelectorRequirement, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 return
}
func (in *NodeSelectorTerm) DeepCopy() *NodeSelectorTerm {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(NodeSelectorTerm)
 in.DeepCopyInto(out)
 return out
}
func (in *NodeSpec) DeepCopyInto(out *NodeSpec) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Taints != nil {
  in, out := &in.Taints, &out.Taints
  *out = make([]Taint, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 if in.ConfigSource != nil {
  in, out := &in.ConfigSource, &out.ConfigSource
  *out = new(NodeConfigSource)
  (*in).DeepCopyInto(*out)
 }
 return
}
func (in *NodeSpec) DeepCopy() *NodeSpec {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(NodeSpec)
 in.DeepCopyInto(out)
 return out
}
func (in *NodeStatus) DeepCopyInto(out *NodeStatus) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Capacity != nil {
  in, out := &in.Capacity, &out.Capacity
  *out = make(ResourceList, len(*in))
  for key, val := range *in {
   (*out)[key] = val.DeepCopy()
  }
 }
 if in.Allocatable != nil {
  in, out := &in.Allocatable, &out.Allocatable
  *out = make(ResourceList, len(*in))
  for key, val := range *in {
   (*out)[key] = val.DeepCopy()
  }
 }
 if in.Conditions != nil {
  in, out := &in.Conditions, &out.Conditions
  *out = make([]NodeCondition, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 if in.Addresses != nil {
  in, out := &in.Addresses, &out.Addresses
  *out = make([]NodeAddress, len(*in))
  copy(*out, *in)
 }
 out.DaemonEndpoints = in.DaemonEndpoints
 out.NodeInfo = in.NodeInfo
 if in.Images != nil {
  in, out := &in.Images, &out.Images
  *out = make([]ContainerImage, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 if in.VolumesInUse != nil {
  in, out := &in.VolumesInUse, &out.VolumesInUse
  *out = make([]UniqueVolumeName, len(*in))
  copy(*out, *in)
 }
 if in.VolumesAttached != nil {
  in, out := &in.VolumesAttached, &out.VolumesAttached
  *out = make([]AttachedVolume, len(*in))
  copy(*out, *in)
 }
 if in.Config != nil {
  in, out := &in.Config, &out.Config
  *out = new(NodeConfigStatus)
  (*in).DeepCopyInto(*out)
 }
 return
}
func (in *NodeStatus) DeepCopy() *NodeStatus {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(NodeStatus)
 in.DeepCopyInto(out)
 return out
}
func (in *NodeSystemInfo) DeepCopyInto(out *NodeSystemInfo) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 return
}
func (in *NodeSystemInfo) DeepCopy() *NodeSystemInfo {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(NodeSystemInfo)
 in.DeepCopyInto(out)
 return out
}
func (in *ObjectFieldSelector) DeepCopyInto(out *ObjectFieldSelector) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 return
}
func (in *ObjectFieldSelector) DeepCopy() *ObjectFieldSelector {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ObjectFieldSelector)
 in.DeepCopyInto(out)
 return out
}
func (in *ObjectReference) DeepCopyInto(out *ObjectReference) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 return
}
func (in *ObjectReference) DeepCopy() *ObjectReference {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ObjectReference)
 in.DeepCopyInto(out)
 return out
}
func (in *ObjectReference) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *PersistentVolume) DeepCopyInto(out *PersistentVolume) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
 in.Spec.DeepCopyInto(&out.Spec)
 out.Status = in.Status
 return
}
func (in *PersistentVolume) DeepCopy() *PersistentVolume {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(PersistentVolume)
 in.DeepCopyInto(out)
 return out
}
func (in *PersistentVolume) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *PersistentVolumeClaim) DeepCopyInto(out *PersistentVolumeClaim) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
 in.Spec.DeepCopyInto(&out.Spec)
 in.Status.DeepCopyInto(&out.Status)
 return
}
func (in *PersistentVolumeClaim) DeepCopy() *PersistentVolumeClaim {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(PersistentVolumeClaim)
 in.DeepCopyInto(out)
 return out
}
func (in *PersistentVolumeClaim) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *PersistentVolumeClaimCondition) DeepCopyInto(out *PersistentVolumeClaimCondition) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 in.LastProbeTime.DeepCopyInto(&out.LastProbeTime)
 in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
 return
}
func (in *PersistentVolumeClaimCondition) DeepCopy() *PersistentVolumeClaimCondition {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(PersistentVolumeClaimCondition)
 in.DeepCopyInto(out)
 return out
}
func (in *PersistentVolumeClaimList) DeepCopyInto(out *PersistentVolumeClaimList) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]PersistentVolumeClaim, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 return
}
func (in *PersistentVolumeClaimList) DeepCopy() *PersistentVolumeClaimList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(PersistentVolumeClaimList)
 in.DeepCopyInto(out)
 return out
}
func (in *PersistentVolumeClaimList) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *PersistentVolumeClaimSpec) DeepCopyInto(out *PersistentVolumeClaimSpec) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.AccessModes != nil {
  in, out := &in.AccessModes, &out.AccessModes
  *out = make([]PersistentVolumeAccessMode, len(*in))
  copy(*out, *in)
 }
 if in.Selector != nil {
  in, out := &in.Selector, &out.Selector
  *out = new(v1.LabelSelector)
  (*in).DeepCopyInto(*out)
 }
 in.Resources.DeepCopyInto(&out.Resources)
 if in.StorageClassName != nil {
  in, out := &in.StorageClassName, &out.StorageClassName
  *out = new(string)
  **out = **in
 }
 if in.VolumeMode != nil {
  in, out := &in.VolumeMode, &out.VolumeMode
  *out = new(PersistentVolumeMode)
  **out = **in
 }
 if in.DataSource != nil {
  in, out := &in.DataSource, &out.DataSource
  *out = new(TypedLocalObjectReference)
  (*in).DeepCopyInto(*out)
 }
 return
}
func (in *PersistentVolumeClaimSpec) DeepCopy() *PersistentVolumeClaimSpec {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(PersistentVolumeClaimSpec)
 in.DeepCopyInto(out)
 return out
}
func (in *PersistentVolumeClaimStatus) DeepCopyInto(out *PersistentVolumeClaimStatus) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.AccessModes != nil {
  in, out := &in.AccessModes, &out.AccessModes
  *out = make([]PersistentVolumeAccessMode, len(*in))
  copy(*out, *in)
 }
 if in.Capacity != nil {
  in, out := &in.Capacity, &out.Capacity
  *out = make(ResourceList, len(*in))
  for key, val := range *in {
   (*out)[key] = val.DeepCopy()
  }
 }
 if in.Conditions != nil {
  in, out := &in.Conditions, &out.Conditions
  *out = make([]PersistentVolumeClaimCondition, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 return
}
func (in *PersistentVolumeClaimStatus) DeepCopy() *PersistentVolumeClaimStatus {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(PersistentVolumeClaimStatus)
 in.DeepCopyInto(out)
 return out
}
func (in *PersistentVolumeClaimVolumeSource) DeepCopyInto(out *PersistentVolumeClaimVolumeSource) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 return
}
func (in *PersistentVolumeClaimVolumeSource) DeepCopy() *PersistentVolumeClaimVolumeSource {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(PersistentVolumeClaimVolumeSource)
 in.DeepCopyInto(out)
 return out
}
func (in *PersistentVolumeList) DeepCopyInto(out *PersistentVolumeList) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]PersistentVolume, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 return
}
func (in *PersistentVolumeList) DeepCopy() *PersistentVolumeList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(PersistentVolumeList)
 in.DeepCopyInto(out)
 return out
}
func (in *PersistentVolumeList) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *PersistentVolumeSource) DeepCopyInto(out *PersistentVolumeSource) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.GCEPersistentDisk != nil {
  in, out := &in.GCEPersistentDisk, &out.GCEPersistentDisk
  *out = new(GCEPersistentDiskVolumeSource)
  **out = **in
 }
 if in.AWSElasticBlockStore != nil {
  in, out := &in.AWSElasticBlockStore, &out.AWSElasticBlockStore
  *out = new(AWSElasticBlockStoreVolumeSource)
  **out = **in
 }
 if in.HostPath != nil {
  in, out := &in.HostPath, &out.HostPath
  *out = new(HostPathVolumeSource)
  (*in).DeepCopyInto(*out)
 }
 if in.Glusterfs != nil {
  in, out := &in.Glusterfs, &out.Glusterfs
  *out = new(GlusterfsPersistentVolumeSource)
  (*in).DeepCopyInto(*out)
 }
 if in.NFS != nil {
  in, out := &in.NFS, &out.NFS
  *out = new(NFSVolumeSource)
  **out = **in
 }
 if in.RBD != nil {
  in, out := &in.RBD, &out.RBD
  *out = new(RBDPersistentVolumeSource)
  (*in).DeepCopyInto(*out)
 }
 if in.Quobyte != nil {
  in, out := &in.Quobyte, &out.Quobyte
  *out = new(QuobyteVolumeSource)
  **out = **in
 }
 if in.ISCSI != nil {
  in, out := &in.ISCSI, &out.ISCSI
  *out = new(ISCSIPersistentVolumeSource)
  (*in).DeepCopyInto(*out)
 }
 if in.FlexVolume != nil {
  in, out := &in.FlexVolume, &out.FlexVolume
  *out = new(FlexPersistentVolumeSource)
  (*in).DeepCopyInto(*out)
 }
 if in.Cinder != nil {
  in, out := &in.Cinder, &out.Cinder
  *out = new(CinderPersistentVolumeSource)
  (*in).DeepCopyInto(*out)
 }
 if in.CephFS != nil {
  in, out := &in.CephFS, &out.CephFS
  *out = new(CephFSPersistentVolumeSource)
  (*in).DeepCopyInto(*out)
 }
 if in.FC != nil {
  in, out := &in.FC, &out.FC
  *out = new(FCVolumeSource)
  (*in).DeepCopyInto(*out)
 }
 if in.Flocker != nil {
  in, out := &in.Flocker, &out.Flocker
  *out = new(FlockerVolumeSource)
  **out = **in
 }
 if in.AzureFile != nil {
  in, out := &in.AzureFile, &out.AzureFile
  *out = new(AzureFilePersistentVolumeSource)
  (*in).DeepCopyInto(*out)
 }
 if in.VsphereVolume != nil {
  in, out := &in.VsphereVolume, &out.VsphereVolume
  *out = new(VsphereVirtualDiskVolumeSource)
  **out = **in
 }
 if in.AzureDisk != nil {
  in, out := &in.AzureDisk, &out.AzureDisk
  *out = new(AzureDiskVolumeSource)
  (*in).DeepCopyInto(*out)
 }
 if in.PhotonPersistentDisk != nil {
  in, out := &in.PhotonPersistentDisk, &out.PhotonPersistentDisk
  *out = new(PhotonPersistentDiskVolumeSource)
  **out = **in
 }
 if in.PortworxVolume != nil {
  in, out := &in.PortworxVolume, &out.PortworxVolume
  *out = new(PortworxVolumeSource)
  **out = **in
 }
 if in.ScaleIO != nil {
  in, out := &in.ScaleIO, &out.ScaleIO
  *out = new(ScaleIOPersistentVolumeSource)
  (*in).DeepCopyInto(*out)
 }
 if in.Local != nil {
  in, out := &in.Local, &out.Local
  *out = new(LocalVolumeSource)
  (*in).DeepCopyInto(*out)
 }
 if in.StorageOS != nil {
  in, out := &in.StorageOS, &out.StorageOS
  *out = new(StorageOSPersistentVolumeSource)
  (*in).DeepCopyInto(*out)
 }
 if in.CSI != nil {
  in, out := &in.CSI, &out.CSI
  *out = new(CSIPersistentVolumeSource)
  (*in).DeepCopyInto(*out)
 }
 return
}
func (in *PersistentVolumeSource) DeepCopy() *PersistentVolumeSource {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(PersistentVolumeSource)
 in.DeepCopyInto(out)
 return out
}
func (in *PersistentVolumeSpec) DeepCopyInto(out *PersistentVolumeSpec) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Capacity != nil {
  in, out := &in.Capacity, &out.Capacity
  *out = make(ResourceList, len(*in))
  for key, val := range *in {
   (*out)[key] = val.DeepCopy()
  }
 }
 in.PersistentVolumeSource.DeepCopyInto(&out.PersistentVolumeSource)
 if in.AccessModes != nil {
  in, out := &in.AccessModes, &out.AccessModes
  *out = make([]PersistentVolumeAccessMode, len(*in))
  copy(*out, *in)
 }
 if in.ClaimRef != nil {
  in, out := &in.ClaimRef, &out.ClaimRef
  *out = new(ObjectReference)
  **out = **in
 }
 if in.MountOptions != nil {
  in, out := &in.MountOptions, &out.MountOptions
  *out = make([]string, len(*in))
  copy(*out, *in)
 }
 if in.VolumeMode != nil {
  in, out := &in.VolumeMode, &out.VolumeMode
  *out = new(PersistentVolumeMode)
  **out = **in
 }
 if in.NodeAffinity != nil {
  in, out := &in.NodeAffinity, &out.NodeAffinity
  *out = new(VolumeNodeAffinity)
  (*in).DeepCopyInto(*out)
 }
 return
}
func (in *PersistentVolumeSpec) DeepCopy() *PersistentVolumeSpec {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(PersistentVolumeSpec)
 in.DeepCopyInto(out)
 return out
}
func (in *PersistentVolumeStatus) DeepCopyInto(out *PersistentVolumeStatus) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 return
}
func (in *PersistentVolumeStatus) DeepCopy() *PersistentVolumeStatus {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(PersistentVolumeStatus)
 in.DeepCopyInto(out)
 return out
}
func (in *PhotonPersistentDiskVolumeSource) DeepCopyInto(out *PhotonPersistentDiskVolumeSource) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 return
}
func (in *PhotonPersistentDiskVolumeSource) DeepCopy() *PhotonPersistentDiskVolumeSource {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(PhotonPersistentDiskVolumeSource)
 in.DeepCopyInto(out)
 return out
}
func (in *Pod) DeepCopyInto(out *Pod) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
 in.Spec.DeepCopyInto(&out.Spec)
 in.Status.DeepCopyInto(&out.Status)
 return
}
func (in *Pod) DeepCopy() *Pod {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(Pod)
 in.DeepCopyInto(out)
 return out
}
func (in *Pod) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *PodAffinity) DeepCopyInto(out *PodAffinity) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.RequiredDuringSchedulingIgnoredDuringExecution != nil {
  in, out := &in.RequiredDuringSchedulingIgnoredDuringExecution, &out.RequiredDuringSchedulingIgnoredDuringExecution
  *out = make([]PodAffinityTerm, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 if in.PreferredDuringSchedulingIgnoredDuringExecution != nil {
  in, out := &in.PreferredDuringSchedulingIgnoredDuringExecution, &out.PreferredDuringSchedulingIgnoredDuringExecution
  *out = make([]WeightedPodAffinityTerm, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 return
}
func (in *PodAffinity) DeepCopy() *PodAffinity {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(PodAffinity)
 in.DeepCopyInto(out)
 return out
}
func (in *PodAffinityTerm) DeepCopyInto(out *PodAffinityTerm) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.LabelSelector != nil {
  in, out := &in.LabelSelector, &out.LabelSelector
  *out = new(v1.LabelSelector)
  (*in).DeepCopyInto(*out)
 }
 if in.Namespaces != nil {
  in, out := &in.Namespaces, &out.Namespaces
  *out = make([]string, len(*in))
  copy(*out, *in)
 }
 return
}
func (in *PodAffinityTerm) DeepCopy() *PodAffinityTerm {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(PodAffinityTerm)
 in.DeepCopyInto(out)
 return out
}
func (in *PodAntiAffinity) DeepCopyInto(out *PodAntiAffinity) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.RequiredDuringSchedulingIgnoredDuringExecution != nil {
  in, out := &in.RequiredDuringSchedulingIgnoredDuringExecution, &out.RequiredDuringSchedulingIgnoredDuringExecution
  *out = make([]PodAffinityTerm, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 if in.PreferredDuringSchedulingIgnoredDuringExecution != nil {
  in, out := &in.PreferredDuringSchedulingIgnoredDuringExecution, &out.PreferredDuringSchedulingIgnoredDuringExecution
  *out = make([]WeightedPodAffinityTerm, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 return
}
func (in *PodAntiAffinity) DeepCopy() *PodAntiAffinity {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(PodAntiAffinity)
 in.DeepCopyInto(out)
 return out
}
func (in *PodAttachOptions) DeepCopyInto(out *PodAttachOptions) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 return
}
func (in *PodAttachOptions) DeepCopy() *PodAttachOptions {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(PodAttachOptions)
 in.DeepCopyInto(out)
 return out
}
func (in *PodAttachOptions) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *PodCondition) DeepCopyInto(out *PodCondition) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 in.LastProbeTime.DeepCopyInto(&out.LastProbeTime)
 in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
 return
}
func (in *PodCondition) DeepCopy() *PodCondition {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(PodCondition)
 in.DeepCopyInto(out)
 return out
}
func (in *PodDNSConfig) DeepCopyInto(out *PodDNSConfig) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Nameservers != nil {
  in, out := &in.Nameservers, &out.Nameservers
  *out = make([]string, len(*in))
  copy(*out, *in)
 }
 if in.Searches != nil {
  in, out := &in.Searches, &out.Searches
  *out = make([]string, len(*in))
  copy(*out, *in)
 }
 if in.Options != nil {
  in, out := &in.Options, &out.Options
  *out = make([]PodDNSConfigOption, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 return
}
func (in *PodDNSConfig) DeepCopy() *PodDNSConfig {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(PodDNSConfig)
 in.DeepCopyInto(out)
 return out
}
func (in *PodDNSConfigOption) DeepCopyInto(out *PodDNSConfigOption) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Value != nil {
  in, out := &in.Value, &out.Value
  *out = new(string)
  **out = **in
 }
 return
}
func (in *PodDNSConfigOption) DeepCopy() *PodDNSConfigOption {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(PodDNSConfigOption)
 in.DeepCopyInto(out)
 return out
}
func (in *PodExecOptions) DeepCopyInto(out *PodExecOptions) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 if in.Command != nil {
  in, out := &in.Command, &out.Command
  *out = make([]string, len(*in))
  copy(*out, *in)
 }
 return
}
func (in *PodExecOptions) DeepCopy() *PodExecOptions {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(PodExecOptions)
 in.DeepCopyInto(out)
 return out
}
func (in *PodExecOptions) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *PodList) DeepCopyInto(out *PodList) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]Pod, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 return
}
func (in *PodList) DeepCopy() *PodList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(PodList)
 in.DeepCopyInto(out)
 return out
}
func (in *PodList) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *PodLogOptions) DeepCopyInto(out *PodLogOptions) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 if in.SinceSeconds != nil {
  in, out := &in.SinceSeconds, &out.SinceSeconds
  *out = new(int64)
  **out = **in
 }
 if in.SinceTime != nil {
  in, out := &in.SinceTime, &out.SinceTime
  *out = (*in).DeepCopy()
 }
 if in.TailLines != nil {
  in, out := &in.TailLines, &out.TailLines
  *out = new(int64)
  **out = **in
 }
 if in.LimitBytes != nil {
  in, out := &in.LimitBytes, &out.LimitBytes
  *out = new(int64)
  **out = **in
 }
 return
}
func (in *PodLogOptions) DeepCopy() *PodLogOptions {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(PodLogOptions)
 in.DeepCopyInto(out)
 return out
}
func (in *PodLogOptions) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *PodPortForwardOptions) DeepCopyInto(out *PodPortForwardOptions) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 if in.Ports != nil {
  in, out := &in.Ports, &out.Ports
  *out = make([]int32, len(*in))
  copy(*out, *in)
 }
 return
}
func (in *PodPortForwardOptions) DeepCopy() *PodPortForwardOptions {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(PodPortForwardOptions)
 in.DeepCopyInto(out)
 return out
}
func (in *PodPortForwardOptions) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *PodProxyOptions) DeepCopyInto(out *PodProxyOptions) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 return
}
func (in *PodProxyOptions) DeepCopy() *PodProxyOptions {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(PodProxyOptions)
 in.DeepCopyInto(out)
 return out
}
func (in *PodProxyOptions) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *PodReadinessGate) DeepCopyInto(out *PodReadinessGate) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 return
}
func (in *PodReadinessGate) DeepCopy() *PodReadinessGate {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(PodReadinessGate)
 in.DeepCopyInto(out)
 return out
}
func (in *PodSecurityContext) DeepCopyInto(out *PodSecurityContext) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.ShareProcessNamespace != nil {
  in, out := &in.ShareProcessNamespace, &out.ShareProcessNamespace
  *out = new(bool)
  **out = **in
 }
 if in.SELinuxOptions != nil {
  in, out := &in.SELinuxOptions, &out.SELinuxOptions
  *out = new(SELinuxOptions)
  **out = **in
 }
 if in.RunAsUser != nil {
  in, out := &in.RunAsUser, &out.RunAsUser
  *out = new(int64)
  **out = **in
 }
 if in.RunAsGroup != nil {
  in, out := &in.RunAsGroup, &out.RunAsGroup
  *out = new(int64)
  **out = **in
 }
 if in.RunAsNonRoot != nil {
  in, out := &in.RunAsNonRoot, &out.RunAsNonRoot
  *out = new(bool)
  **out = **in
 }
 if in.SupplementalGroups != nil {
  in, out := &in.SupplementalGroups, &out.SupplementalGroups
  *out = make([]int64, len(*in))
  copy(*out, *in)
 }
 if in.FSGroup != nil {
  in, out := &in.FSGroup, &out.FSGroup
  *out = new(int64)
  **out = **in
 }
 if in.Sysctls != nil {
  in, out := &in.Sysctls, &out.Sysctls
  *out = make([]Sysctl, len(*in))
  copy(*out, *in)
 }
 return
}
func (in *PodSecurityContext) DeepCopy() *PodSecurityContext {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(PodSecurityContext)
 in.DeepCopyInto(out)
 return out
}
func (in *PodSignature) DeepCopyInto(out *PodSignature) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.PodController != nil {
  in, out := &in.PodController, &out.PodController
  *out = new(v1.OwnerReference)
  (*in).DeepCopyInto(*out)
 }
 return
}
func (in *PodSignature) DeepCopy() *PodSignature {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(PodSignature)
 in.DeepCopyInto(out)
 return out
}
func (in *PodSpec) DeepCopyInto(out *PodSpec) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Volumes != nil {
  in, out := &in.Volumes, &out.Volumes
  *out = make([]Volume, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 if in.InitContainers != nil {
  in, out := &in.InitContainers, &out.InitContainers
  *out = make([]Container, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 if in.Containers != nil {
  in, out := &in.Containers, &out.Containers
  *out = make([]Container, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 if in.TerminationGracePeriodSeconds != nil {
  in, out := &in.TerminationGracePeriodSeconds, &out.TerminationGracePeriodSeconds
  *out = new(int64)
  **out = **in
 }
 if in.ActiveDeadlineSeconds != nil {
  in, out := &in.ActiveDeadlineSeconds, &out.ActiveDeadlineSeconds
  *out = new(int64)
  **out = **in
 }
 if in.NodeSelector != nil {
  in, out := &in.NodeSelector, &out.NodeSelector
  *out = make(map[string]string, len(*in))
  for key, val := range *in {
   (*out)[key] = val
  }
 }
 if in.AutomountServiceAccountToken != nil {
  in, out := &in.AutomountServiceAccountToken, &out.AutomountServiceAccountToken
  *out = new(bool)
  **out = **in
 }
 if in.SecurityContext != nil {
  in, out := &in.SecurityContext, &out.SecurityContext
  *out = new(PodSecurityContext)
  (*in).DeepCopyInto(*out)
 }
 if in.ImagePullSecrets != nil {
  in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
  *out = make([]LocalObjectReference, len(*in))
  copy(*out, *in)
 }
 if in.Affinity != nil {
  in, out := &in.Affinity, &out.Affinity
  *out = new(Affinity)
  (*in).DeepCopyInto(*out)
 }
 if in.Tolerations != nil {
  in, out := &in.Tolerations, &out.Tolerations
  *out = make([]Toleration, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 if in.HostAliases != nil {
  in, out := &in.HostAliases, &out.HostAliases
  *out = make([]HostAlias, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 if in.Priority != nil {
  in, out := &in.Priority, &out.Priority
  *out = new(int32)
  **out = **in
 }
 if in.DNSConfig != nil {
  in, out := &in.DNSConfig, &out.DNSConfig
  *out = new(PodDNSConfig)
  (*in).DeepCopyInto(*out)
 }
 if in.ReadinessGates != nil {
  in, out := &in.ReadinessGates, &out.ReadinessGates
  *out = make([]PodReadinessGate, len(*in))
  copy(*out, *in)
 }
 if in.RuntimeClassName != nil {
  in, out := &in.RuntimeClassName, &out.RuntimeClassName
  *out = new(string)
  **out = **in
 }
 if in.EnableServiceLinks != nil {
  in, out := &in.EnableServiceLinks, &out.EnableServiceLinks
  *out = new(bool)
  **out = **in
 }
 return
}
func (in *PodSpec) DeepCopy() *PodSpec {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(PodSpec)
 in.DeepCopyInto(out)
 return out
}
func (in *PodStatus) DeepCopyInto(out *PodStatus) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Conditions != nil {
  in, out := &in.Conditions, &out.Conditions
  *out = make([]PodCondition, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 if in.StartTime != nil {
  in, out := &in.StartTime, &out.StartTime
  *out = (*in).DeepCopy()
 }
 if in.InitContainerStatuses != nil {
  in, out := &in.InitContainerStatuses, &out.InitContainerStatuses
  *out = make([]ContainerStatus, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 if in.ContainerStatuses != nil {
  in, out := &in.ContainerStatuses, &out.ContainerStatuses
  *out = make([]ContainerStatus, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 return
}
func (in *PodStatus) DeepCopy() *PodStatus {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(PodStatus)
 in.DeepCopyInto(out)
 return out
}
func (in *PodStatusResult) DeepCopyInto(out *PodStatusResult) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
 in.Status.DeepCopyInto(&out.Status)
 return
}
func (in *PodStatusResult) DeepCopy() *PodStatusResult {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(PodStatusResult)
 in.DeepCopyInto(out)
 return out
}
func (in *PodStatusResult) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *PodTemplate) DeepCopyInto(out *PodTemplate) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
 in.Template.DeepCopyInto(&out.Template)
 return
}
func (in *PodTemplate) DeepCopy() *PodTemplate {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(PodTemplate)
 in.DeepCopyInto(out)
 return out
}
func (in *PodTemplate) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *PodTemplateList) DeepCopyInto(out *PodTemplateList) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]PodTemplate, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 return
}
func (in *PodTemplateList) DeepCopy() *PodTemplateList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(PodTemplateList)
 in.DeepCopyInto(out)
 return out
}
func (in *PodTemplateList) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *PodTemplateSpec) DeepCopyInto(out *PodTemplateSpec) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
 in.Spec.DeepCopyInto(&out.Spec)
 return
}
func (in *PodTemplateSpec) DeepCopy() *PodTemplateSpec {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(PodTemplateSpec)
 in.DeepCopyInto(out)
 return out
}
func (in *PortworxVolumeSource) DeepCopyInto(out *PortworxVolumeSource) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 return
}
func (in *PortworxVolumeSource) DeepCopy() *PortworxVolumeSource {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(PortworxVolumeSource)
 in.DeepCopyInto(out)
 return out
}
func (in *Preconditions) DeepCopyInto(out *Preconditions) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.UID != nil {
  in, out := &in.UID, &out.UID
  *out = new(types.UID)
  **out = **in
 }
 return
}
func (in *Preconditions) DeepCopy() *Preconditions {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(Preconditions)
 in.DeepCopyInto(out)
 return out
}
func (in *PreferAvoidPodsEntry) DeepCopyInto(out *PreferAvoidPodsEntry) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 in.PodSignature.DeepCopyInto(&out.PodSignature)
 in.EvictionTime.DeepCopyInto(&out.EvictionTime)
 return
}
func (in *PreferAvoidPodsEntry) DeepCopy() *PreferAvoidPodsEntry {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(PreferAvoidPodsEntry)
 in.DeepCopyInto(out)
 return out
}
func (in *PreferredSchedulingTerm) DeepCopyInto(out *PreferredSchedulingTerm) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 in.Preference.DeepCopyInto(&out.Preference)
 return
}
func (in *PreferredSchedulingTerm) DeepCopy() *PreferredSchedulingTerm {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(PreferredSchedulingTerm)
 in.DeepCopyInto(out)
 return out
}
func (in *Probe) DeepCopyInto(out *Probe) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 in.Handler.DeepCopyInto(&out.Handler)
 return
}
func (in *Probe) DeepCopy() *Probe {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(Probe)
 in.DeepCopyInto(out)
 return out
}
func (in *ProjectedVolumeSource) DeepCopyInto(out *ProjectedVolumeSource) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Sources != nil {
  in, out := &in.Sources, &out.Sources
  *out = make([]VolumeProjection, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 if in.DefaultMode != nil {
  in, out := &in.DefaultMode, &out.DefaultMode
  *out = new(int32)
  **out = **in
 }
 return
}
func (in *ProjectedVolumeSource) DeepCopy() *ProjectedVolumeSource {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ProjectedVolumeSource)
 in.DeepCopyInto(out)
 return out
}
func (in *QuobyteVolumeSource) DeepCopyInto(out *QuobyteVolumeSource) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 return
}
func (in *QuobyteVolumeSource) DeepCopy() *QuobyteVolumeSource {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(QuobyteVolumeSource)
 in.DeepCopyInto(out)
 return out
}
func (in *RBDPersistentVolumeSource) DeepCopyInto(out *RBDPersistentVolumeSource) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.CephMonitors != nil {
  in, out := &in.CephMonitors, &out.CephMonitors
  *out = make([]string, len(*in))
  copy(*out, *in)
 }
 if in.SecretRef != nil {
  in, out := &in.SecretRef, &out.SecretRef
  *out = new(SecretReference)
  **out = **in
 }
 return
}
func (in *RBDPersistentVolumeSource) DeepCopy() *RBDPersistentVolumeSource {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(RBDPersistentVolumeSource)
 in.DeepCopyInto(out)
 return out
}
func (in *RBDVolumeSource) DeepCopyInto(out *RBDVolumeSource) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.CephMonitors != nil {
  in, out := &in.CephMonitors, &out.CephMonitors
  *out = make([]string, len(*in))
  copy(*out, *in)
 }
 if in.SecretRef != nil {
  in, out := &in.SecretRef, &out.SecretRef
  *out = new(LocalObjectReference)
  **out = **in
 }
 return
}
func (in *RBDVolumeSource) DeepCopy() *RBDVolumeSource {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(RBDVolumeSource)
 in.DeepCopyInto(out)
 return out
}
func (in *RangeAllocation) DeepCopyInto(out *RangeAllocation) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
 if in.Data != nil {
  in, out := &in.Data, &out.Data
  *out = make([]byte, len(*in))
  copy(*out, *in)
 }
 return
}
func (in *RangeAllocation) DeepCopy() *RangeAllocation {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(RangeAllocation)
 in.DeepCopyInto(out)
 return out
}
func (in *RangeAllocation) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *ReplicationController) DeepCopyInto(out *ReplicationController) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
 in.Spec.DeepCopyInto(&out.Spec)
 in.Status.DeepCopyInto(&out.Status)
 return
}
func (in *ReplicationController) DeepCopy() *ReplicationController {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ReplicationController)
 in.DeepCopyInto(out)
 return out
}
func (in *ReplicationController) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *ReplicationControllerCondition) DeepCopyInto(out *ReplicationControllerCondition) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
 return
}
func (in *ReplicationControllerCondition) DeepCopy() *ReplicationControllerCondition {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ReplicationControllerCondition)
 in.DeepCopyInto(out)
 return out
}
func (in *ReplicationControllerList) DeepCopyInto(out *ReplicationControllerList) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]ReplicationController, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 return
}
func (in *ReplicationControllerList) DeepCopy() *ReplicationControllerList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ReplicationControllerList)
 in.DeepCopyInto(out)
 return out
}
func (in *ReplicationControllerList) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *ReplicationControllerSpec) DeepCopyInto(out *ReplicationControllerSpec) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Selector != nil {
  in, out := &in.Selector, &out.Selector
  *out = make(map[string]string, len(*in))
  for key, val := range *in {
   (*out)[key] = val
  }
 }
 if in.Template != nil {
  in, out := &in.Template, &out.Template
  *out = new(PodTemplateSpec)
  (*in).DeepCopyInto(*out)
 }
 return
}
func (in *ReplicationControllerSpec) DeepCopy() *ReplicationControllerSpec {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ReplicationControllerSpec)
 in.DeepCopyInto(out)
 return out
}
func (in *ReplicationControllerStatus) DeepCopyInto(out *ReplicationControllerStatus) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Conditions != nil {
  in, out := &in.Conditions, &out.Conditions
  *out = make([]ReplicationControllerCondition, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 return
}
func (in *ReplicationControllerStatus) DeepCopy() *ReplicationControllerStatus {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ReplicationControllerStatus)
 in.DeepCopyInto(out)
 return out
}
func (in *ResourceFieldSelector) DeepCopyInto(out *ResourceFieldSelector) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.Divisor = in.Divisor.DeepCopy()
 return
}
func (in *ResourceFieldSelector) DeepCopy() *ResourceFieldSelector {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ResourceFieldSelector)
 in.DeepCopyInto(out)
 return out
}
func (in ResourceList) DeepCopyInto(out *ResourceList) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 {
  in := &in
  *out = make(ResourceList, len(*in))
  for key, val := range *in {
   (*out)[key] = val.DeepCopy()
  }
  return
 }
}
func (in ResourceList) DeepCopy() ResourceList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ResourceList)
 in.DeepCopyInto(out)
 return *out
}
func (in *ResourceQuota) DeepCopyInto(out *ResourceQuota) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
 in.Spec.DeepCopyInto(&out.Spec)
 in.Status.DeepCopyInto(&out.Status)
 return
}
func (in *ResourceQuota) DeepCopy() *ResourceQuota {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ResourceQuota)
 in.DeepCopyInto(out)
 return out
}
func (in *ResourceQuota) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *ResourceQuotaList) DeepCopyInto(out *ResourceQuotaList) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]ResourceQuota, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 return
}
func (in *ResourceQuotaList) DeepCopy() *ResourceQuotaList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ResourceQuotaList)
 in.DeepCopyInto(out)
 return out
}
func (in *ResourceQuotaList) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *ResourceQuotaSpec) DeepCopyInto(out *ResourceQuotaSpec) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Hard != nil {
  in, out := &in.Hard, &out.Hard
  *out = make(ResourceList, len(*in))
  for key, val := range *in {
   (*out)[key] = val.DeepCopy()
  }
 }
 if in.Scopes != nil {
  in, out := &in.Scopes, &out.Scopes
  *out = make([]ResourceQuotaScope, len(*in))
  copy(*out, *in)
 }
 if in.ScopeSelector != nil {
  in, out := &in.ScopeSelector, &out.ScopeSelector
  *out = new(ScopeSelector)
  (*in).DeepCopyInto(*out)
 }
 return
}
func (in *ResourceQuotaSpec) DeepCopy() *ResourceQuotaSpec {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ResourceQuotaSpec)
 in.DeepCopyInto(out)
 return out
}
func (in *ResourceQuotaStatus) DeepCopyInto(out *ResourceQuotaStatus) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Hard != nil {
  in, out := &in.Hard, &out.Hard
  *out = make(ResourceList, len(*in))
  for key, val := range *in {
   (*out)[key] = val.DeepCopy()
  }
 }
 if in.Used != nil {
  in, out := &in.Used, &out.Used
  *out = make(ResourceList, len(*in))
  for key, val := range *in {
   (*out)[key] = val.DeepCopy()
  }
 }
 return
}
func (in *ResourceQuotaStatus) DeepCopy() *ResourceQuotaStatus {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ResourceQuotaStatus)
 in.DeepCopyInto(out)
 return out
}
func (in *ResourceRequirements) DeepCopyInto(out *ResourceRequirements) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Limits != nil {
  in, out := &in.Limits, &out.Limits
  *out = make(ResourceList, len(*in))
  for key, val := range *in {
   (*out)[key] = val.DeepCopy()
  }
 }
 if in.Requests != nil {
  in, out := &in.Requests, &out.Requests
  *out = make(ResourceList, len(*in))
  for key, val := range *in {
   (*out)[key] = val.DeepCopy()
  }
 }
 return
}
func (in *ResourceRequirements) DeepCopy() *ResourceRequirements {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ResourceRequirements)
 in.DeepCopyInto(out)
 return out
}
func (in *SELinuxOptions) DeepCopyInto(out *SELinuxOptions) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 return
}
func (in *SELinuxOptions) DeepCopy() *SELinuxOptions {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(SELinuxOptions)
 in.DeepCopyInto(out)
 return out
}
func (in *ScaleIOPersistentVolumeSource) DeepCopyInto(out *ScaleIOPersistentVolumeSource) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.SecretRef != nil {
  in, out := &in.SecretRef, &out.SecretRef
  *out = new(SecretReference)
  **out = **in
 }
 return
}
func (in *ScaleIOPersistentVolumeSource) DeepCopy() *ScaleIOPersistentVolumeSource {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ScaleIOPersistentVolumeSource)
 in.DeepCopyInto(out)
 return out
}
func (in *ScaleIOVolumeSource) DeepCopyInto(out *ScaleIOVolumeSource) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.SecretRef != nil {
  in, out := &in.SecretRef, &out.SecretRef
  *out = new(LocalObjectReference)
  **out = **in
 }
 return
}
func (in *ScaleIOVolumeSource) DeepCopy() *ScaleIOVolumeSource {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ScaleIOVolumeSource)
 in.DeepCopyInto(out)
 return out
}
func (in *ScopeSelector) DeepCopyInto(out *ScopeSelector) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.MatchExpressions != nil {
  in, out := &in.MatchExpressions, &out.MatchExpressions
  *out = make([]ScopedResourceSelectorRequirement, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 return
}
func (in *ScopeSelector) DeepCopy() *ScopeSelector {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ScopeSelector)
 in.DeepCopyInto(out)
 return out
}
func (in *ScopedResourceSelectorRequirement) DeepCopyInto(out *ScopedResourceSelectorRequirement) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Values != nil {
  in, out := &in.Values, &out.Values
  *out = make([]string, len(*in))
  copy(*out, *in)
 }
 return
}
func (in *ScopedResourceSelectorRequirement) DeepCopy() *ScopedResourceSelectorRequirement {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ScopedResourceSelectorRequirement)
 in.DeepCopyInto(out)
 return out
}
func (in *Secret) DeepCopyInto(out *Secret) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
 if in.Data != nil {
  in, out := &in.Data, &out.Data
  *out = make(map[string][]byte, len(*in))
  for key, val := range *in {
   var outVal []byte
   if val == nil {
    (*out)[key] = nil
   } else {
    in, out := &val, &outVal
    *out = make([]byte, len(*in))
    copy(*out, *in)
   }
   (*out)[key] = outVal
  }
 }
 return
}
func (in *Secret) DeepCopy() *Secret {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(Secret)
 in.DeepCopyInto(out)
 return out
}
func (in *Secret) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *SecretEnvSource) DeepCopyInto(out *SecretEnvSource) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.LocalObjectReference = in.LocalObjectReference
 if in.Optional != nil {
  in, out := &in.Optional, &out.Optional
  *out = new(bool)
  **out = **in
 }
 return
}
func (in *SecretEnvSource) DeepCopy() *SecretEnvSource {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(SecretEnvSource)
 in.DeepCopyInto(out)
 return out
}
func (in *SecretKeySelector) DeepCopyInto(out *SecretKeySelector) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.LocalObjectReference = in.LocalObjectReference
 if in.Optional != nil {
  in, out := &in.Optional, &out.Optional
  *out = new(bool)
  **out = **in
 }
 return
}
func (in *SecretKeySelector) DeepCopy() *SecretKeySelector {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(SecretKeySelector)
 in.DeepCopyInto(out)
 return out
}
func (in *SecretList) DeepCopyInto(out *SecretList) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]Secret, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 return
}
func (in *SecretList) DeepCopy() *SecretList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(SecretList)
 in.DeepCopyInto(out)
 return out
}
func (in *SecretList) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *SecretProjection) DeepCopyInto(out *SecretProjection) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.LocalObjectReference = in.LocalObjectReference
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]KeyToPath, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 if in.Optional != nil {
  in, out := &in.Optional, &out.Optional
  *out = new(bool)
  **out = **in
 }
 return
}
func (in *SecretProjection) DeepCopy() *SecretProjection {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(SecretProjection)
 in.DeepCopyInto(out)
 return out
}
func (in *SecretReference) DeepCopyInto(out *SecretReference) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 return
}
func (in *SecretReference) DeepCopy() *SecretReference {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(SecretReference)
 in.DeepCopyInto(out)
 return out
}
func (in *SecretVolumeSource) DeepCopyInto(out *SecretVolumeSource) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]KeyToPath, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 if in.DefaultMode != nil {
  in, out := &in.DefaultMode, &out.DefaultMode
  *out = new(int32)
  **out = **in
 }
 if in.Optional != nil {
  in, out := &in.Optional, &out.Optional
  *out = new(bool)
  **out = **in
 }
 return
}
func (in *SecretVolumeSource) DeepCopy() *SecretVolumeSource {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(SecretVolumeSource)
 in.DeepCopyInto(out)
 return out
}
func (in *SecurityContext) DeepCopyInto(out *SecurityContext) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Capabilities != nil {
  in, out := &in.Capabilities, &out.Capabilities
  *out = new(Capabilities)
  (*in).DeepCopyInto(*out)
 }
 if in.Privileged != nil {
  in, out := &in.Privileged, &out.Privileged
  *out = new(bool)
  **out = **in
 }
 if in.SELinuxOptions != nil {
  in, out := &in.SELinuxOptions, &out.SELinuxOptions
  *out = new(SELinuxOptions)
  **out = **in
 }
 if in.RunAsUser != nil {
  in, out := &in.RunAsUser, &out.RunAsUser
  *out = new(int64)
  **out = **in
 }
 if in.RunAsGroup != nil {
  in, out := &in.RunAsGroup, &out.RunAsGroup
  *out = new(int64)
  **out = **in
 }
 if in.RunAsNonRoot != nil {
  in, out := &in.RunAsNonRoot, &out.RunAsNonRoot
  *out = new(bool)
  **out = **in
 }
 if in.ReadOnlyRootFilesystem != nil {
  in, out := &in.ReadOnlyRootFilesystem, &out.ReadOnlyRootFilesystem
  *out = new(bool)
  **out = **in
 }
 if in.AllowPrivilegeEscalation != nil {
  in, out := &in.AllowPrivilegeEscalation, &out.AllowPrivilegeEscalation
  *out = new(bool)
  **out = **in
 }
 if in.ProcMount != nil {
  in, out := &in.ProcMount, &out.ProcMount
  *out = new(ProcMountType)
  **out = **in
 }
 return
}
func (in *SecurityContext) DeepCopy() *SecurityContext {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(SecurityContext)
 in.DeepCopyInto(out)
 return out
}
func (in *SerializedReference) DeepCopyInto(out *SerializedReference) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 out.Reference = in.Reference
 return
}
func (in *SerializedReference) DeepCopy() *SerializedReference {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(SerializedReference)
 in.DeepCopyInto(out)
 return out
}
func (in *SerializedReference) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *Service) DeepCopyInto(out *Service) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
 in.Spec.DeepCopyInto(&out.Spec)
 in.Status.DeepCopyInto(&out.Status)
 return
}
func (in *Service) DeepCopy() *Service {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(Service)
 in.DeepCopyInto(out)
 return out
}
func (in *Service) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *ServiceAccount) DeepCopyInto(out *ServiceAccount) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
 if in.Secrets != nil {
  in, out := &in.Secrets, &out.Secrets
  *out = make([]ObjectReference, len(*in))
  copy(*out, *in)
 }
 if in.ImagePullSecrets != nil {
  in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
  *out = make([]LocalObjectReference, len(*in))
  copy(*out, *in)
 }
 if in.AutomountServiceAccountToken != nil {
  in, out := &in.AutomountServiceAccountToken, &out.AutomountServiceAccountToken
  *out = new(bool)
  **out = **in
 }
 return
}
func (in *ServiceAccount) DeepCopy() *ServiceAccount {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ServiceAccount)
 in.DeepCopyInto(out)
 return out
}
func (in *ServiceAccount) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *ServiceAccountList) DeepCopyInto(out *ServiceAccountList) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]ServiceAccount, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 return
}
func (in *ServiceAccountList) DeepCopy() *ServiceAccountList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ServiceAccountList)
 in.DeepCopyInto(out)
 return out
}
func (in *ServiceAccountList) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *ServiceAccountTokenProjection) DeepCopyInto(out *ServiceAccountTokenProjection) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 return
}
func (in *ServiceAccountTokenProjection) DeepCopy() *ServiceAccountTokenProjection {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ServiceAccountTokenProjection)
 in.DeepCopyInto(out)
 return out
}
func (in *ServiceList) DeepCopyInto(out *ServiceList) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]Service, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 return
}
func (in *ServiceList) DeepCopy() *ServiceList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ServiceList)
 in.DeepCopyInto(out)
 return out
}
func (in *ServiceList) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *ServicePort) DeepCopyInto(out *ServicePort) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TargetPort = in.TargetPort
 return
}
func (in *ServicePort) DeepCopy() *ServicePort {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ServicePort)
 in.DeepCopyInto(out)
 return out
}
func (in *ServiceProxyOptions) DeepCopyInto(out *ServiceProxyOptions) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 return
}
func (in *ServiceProxyOptions) DeepCopy() *ServiceProxyOptions {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ServiceProxyOptions)
 in.DeepCopyInto(out)
 return out
}
func (in *ServiceProxyOptions) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *ServiceSpec) DeepCopyInto(out *ServiceSpec) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Ports != nil {
  in, out := &in.Ports, &out.Ports
  *out = make([]ServicePort, len(*in))
  copy(*out, *in)
 }
 if in.Selector != nil {
  in, out := &in.Selector, &out.Selector
  *out = make(map[string]string, len(*in))
  for key, val := range *in {
   (*out)[key] = val
  }
 }
 if in.ExternalIPs != nil {
  in, out := &in.ExternalIPs, &out.ExternalIPs
  *out = make([]string, len(*in))
  copy(*out, *in)
 }
 if in.SessionAffinityConfig != nil {
  in, out := &in.SessionAffinityConfig, &out.SessionAffinityConfig
  *out = new(SessionAffinityConfig)
  (*in).DeepCopyInto(*out)
 }
 if in.LoadBalancerSourceRanges != nil {
  in, out := &in.LoadBalancerSourceRanges, &out.LoadBalancerSourceRanges
  *out = make([]string, len(*in))
  copy(*out, *in)
 }
 return
}
func (in *ServiceSpec) DeepCopy() *ServiceSpec {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ServiceSpec)
 in.DeepCopyInto(out)
 return out
}
func (in *ServiceStatus) DeepCopyInto(out *ServiceStatus) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 in.LoadBalancer.DeepCopyInto(&out.LoadBalancer)
 return
}
func (in *ServiceStatus) DeepCopy() *ServiceStatus {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ServiceStatus)
 in.DeepCopyInto(out)
 return out
}
func (in *SessionAffinityConfig) DeepCopyInto(out *SessionAffinityConfig) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.ClientIP != nil {
  in, out := &in.ClientIP, &out.ClientIP
  *out = new(ClientIPConfig)
  (*in).DeepCopyInto(*out)
 }
 return
}
func (in *SessionAffinityConfig) DeepCopy() *SessionAffinityConfig {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(SessionAffinityConfig)
 in.DeepCopyInto(out)
 return out
}
func (in *StorageOSPersistentVolumeSource) DeepCopyInto(out *StorageOSPersistentVolumeSource) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.SecretRef != nil {
  in, out := &in.SecretRef, &out.SecretRef
  *out = new(ObjectReference)
  **out = **in
 }
 return
}
func (in *StorageOSPersistentVolumeSource) DeepCopy() *StorageOSPersistentVolumeSource {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(StorageOSPersistentVolumeSource)
 in.DeepCopyInto(out)
 return out
}
func (in *StorageOSVolumeSource) DeepCopyInto(out *StorageOSVolumeSource) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.SecretRef != nil {
  in, out := &in.SecretRef, &out.SecretRef
  *out = new(LocalObjectReference)
  **out = **in
 }
 return
}
func (in *StorageOSVolumeSource) DeepCopy() *StorageOSVolumeSource {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(StorageOSVolumeSource)
 in.DeepCopyInto(out)
 return out
}
func (in *Sysctl) DeepCopyInto(out *Sysctl) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 return
}
func (in *Sysctl) DeepCopy() *Sysctl {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(Sysctl)
 in.DeepCopyInto(out)
 return out
}
func (in *TCPSocketAction) DeepCopyInto(out *TCPSocketAction) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.Port = in.Port
 return
}
func (in *TCPSocketAction) DeepCopy() *TCPSocketAction {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(TCPSocketAction)
 in.DeepCopyInto(out)
 return out
}
func (in *Taint) DeepCopyInto(out *Taint) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.TimeAdded != nil {
  in, out := &in.TimeAdded, &out.TimeAdded
  *out = (*in).DeepCopy()
 }
 return
}
func (in *Taint) DeepCopy() *Taint {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(Taint)
 in.DeepCopyInto(out)
 return out
}
func (in *Toleration) DeepCopyInto(out *Toleration) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.TolerationSeconds != nil {
  in, out := &in.TolerationSeconds, &out.TolerationSeconds
  *out = new(int64)
  **out = **in
 }
 return
}
func (in *Toleration) DeepCopy() *Toleration {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(Toleration)
 in.DeepCopyInto(out)
 return out
}
func (in *TopologySelectorLabelRequirement) DeepCopyInto(out *TopologySelectorLabelRequirement) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Values != nil {
  in, out := &in.Values, &out.Values
  *out = make([]string, len(*in))
  copy(*out, *in)
 }
 return
}
func (in *TopologySelectorLabelRequirement) DeepCopy() *TopologySelectorLabelRequirement {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(TopologySelectorLabelRequirement)
 in.DeepCopyInto(out)
 return out
}
func (in *TopologySelectorTerm) DeepCopyInto(out *TopologySelectorTerm) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.MatchLabelExpressions != nil {
  in, out := &in.MatchLabelExpressions, &out.MatchLabelExpressions
  *out = make([]TopologySelectorLabelRequirement, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 return
}
func (in *TopologySelectorTerm) DeepCopy() *TopologySelectorTerm {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(TopologySelectorTerm)
 in.DeepCopyInto(out)
 return out
}
func (in *TypedLocalObjectReference) DeepCopyInto(out *TypedLocalObjectReference) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.APIGroup != nil {
  in, out := &in.APIGroup, &out.APIGroup
  *out = new(string)
  **out = **in
 }
 return
}
func (in *TypedLocalObjectReference) DeepCopy() *TypedLocalObjectReference {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(TypedLocalObjectReference)
 in.DeepCopyInto(out)
 return out
}
func (in *Volume) DeepCopyInto(out *Volume) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 in.VolumeSource.DeepCopyInto(&out.VolumeSource)
 return
}
func (in *Volume) DeepCopy() *Volume {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(Volume)
 in.DeepCopyInto(out)
 return out
}
func (in *VolumeDevice) DeepCopyInto(out *VolumeDevice) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 return
}
func (in *VolumeDevice) DeepCopy() *VolumeDevice {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(VolumeDevice)
 in.DeepCopyInto(out)
 return out
}
func (in *VolumeMount) DeepCopyInto(out *VolumeMount) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.MountPropagation != nil {
  in, out := &in.MountPropagation, &out.MountPropagation
  *out = new(MountPropagationMode)
  **out = **in
 }
 return
}
func (in *VolumeMount) DeepCopy() *VolumeMount {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(VolumeMount)
 in.DeepCopyInto(out)
 return out
}
func (in *VolumeNodeAffinity) DeepCopyInto(out *VolumeNodeAffinity) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Required != nil {
  in, out := &in.Required, &out.Required
  *out = new(NodeSelector)
  (*in).DeepCopyInto(*out)
 }
 return
}
func (in *VolumeNodeAffinity) DeepCopy() *VolumeNodeAffinity {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(VolumeNodeAffinity)
 in.DeepCopyInto(out)
 return out
}
func (in *VolumeProjection) DeepCopyInto(out *VolumeProjection) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Secret != nil {
  in, out := &in.Secret, &out.Secret
  *out = new(SecretProjection)
  (*in).DeepCopyInto(*out)
 }
 if in.DownwardAPI != nil {
  in, out := &in.DownwardAPI, &out.DownwardAPI
  *out = new(DownwardAPIProjection)
  (*in).DeepCopyInto(*out)
 }
 if in.ConfigMap != nil {
  in, out := &in.ConfigMap, &out.ConfigMap
  *out = new(ConfigMapProjection)
  (*in).DeepCopyInto(*out)
 }
 if in.ServiceAccountToken != nil {
  in, out := &in.ServiceAccountToken, &out.ServiceAccountToken
  *out = new(ServiceAccountTokenProjection)
  **out = **in
 }
 return
}
func (in *VolumeProjection) DeepCopy() *VolumeProjection {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(VolumeProjection)
 in.DeepCopyInto(out)
 return out
}
func (in *VolumeSource) DeepCopyInto(out *VolumeSource) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.HostPath != nil {
  in, out := &in.HostPath, &out.HostPath
  *out = new(HostPathVolumeSource)
  (*in).DeepCopyInto(*out)
 }
 if in.EmptyDir != nil {
  in, out := &in.EmptyDir, &out.EmptyDir
  *out = new(EmptyDirVolumeSource)
  (*in).DeepCopyInto(*out)
 }
 if in.GCEPersistentDisk != nil {
  in, out := &in.GCEPersistentDisk, &out.GCEPersistentDisk
  *out = new(GCEPersistentDiskVolumeSource)
  **out = **in
 }
 if in.AWSElasticBlockStore != nil {
  in, out := &in.AWSElasticBlockStore, &out.AWSElasticBlockStore
  *out = new(AWSElasticBlockStoreVolumeSource)
  **out = **in
 }
 if in.GitRepo != nil {
  in, out := &in.GitRepo, &out.GitRepo
  *out = new(GitRepoVolumeSource)
  **out = **in
 }
 if in.Secret != nil {
  in, out := &in.Secret, &out.Secret
  *out = new(SecretVolumeSource)
  (*in).DeepCopyInto(*out)
 }
 if in.NFS != nil {
  in, out := &in.NFS, &out.NFS
  *out = new(NFSVolumeSource)
  **out = **in
 }
 if in.ISCSI != nil {
  in, out := &in.ISCSI, &out.ISCSI
  *out = new(ISCSIVolumeSource)
  (*in).DeepCopyInto(*out)
 }
 if in.Glusterfs != nil {
  in, out := &in.Glusterfs, &out.Glusterfs
  *out = new(GlusterfsVolumeSource)
  **out = **in
 }
 if in.PersistentVolumeClaim != nil {
  in, out := &in.PersistentVolumeClaim, &out.PersistentVolumeClaim
  *out = new(PersistentVolumeClaimVolumeSource)
  **out = **in
 }
 if in.RBD != nil {
  in, out := &in.RBD, &out.RBD
  *out = new(RBDVolumeSource)
  (*in).DeepCopyInto(*out)
 }
 if in.Quobyte != nil {
  in, out := &in.Quobyte, &out.Quobyte
  *out = new(QuobyteVolumeSource)
  **out = **in
 }
 if in.FlexVolume != nil {
  in, out := &in.FlexVolume, &out.FlexVolume
  *out = new(FlexVolumeSource)
  (*in).DeepCopyInto(*out)
 }
 if in.Cinder != nil {
  in, out := &in.Cinder, &out.Cinder
  *out = new(CinderVolumeSource)
  (*in).DeepCopyInto(*out)
 }
 if in.CephFS != nil {
  in, out := &in.CephFS, &out.CephFS
  *out = new(CephFSVolumeSource)
  (*in).DeepCopyInto(*out)
 }
 if in.Flocker != nil {
  in, out := &in.Flocker, &out.Flocker
  *out = new(FlockerVolumeSource)
  **out = **in
 }
 if in.DownwardAPI != nil {
  in, out := &in.DownwardAPI, &out.DownwardAPI
  *out = new(DownwardAPIVolumeSource)
  (*in).DeepCopyInto(*out)
 }
 if in.FC != nil {
  in, out := &in.FC, &out.FC
  *out = new(FCVolumeSource)
  (*in).DeepCopyInto(*out)
 }
 if in.AzureFile != nil {
  in, out := &in.AzureFile, &out.AzureFile
  *out = new(AzureFileVolumeSource)
  **out = **in
 }
 if in.ConfigMap != nil {
  in, out := &in.ConfigMap, &out.ConfigMap
  *out = new(ConfigMapVolumeSource)
  (*in).DeepCopyInto(*out)
 }
 if in.VsphereVolume != nil {
  in, out := &in.VsphereVolume, &out.VsphereVolume
  *out = new(VsphereVirtualDiskVolumeSource)
  **out = **in
 }
 if in.AzureDisk != nil {
  in, out := &in.AzureDisk, &out.AzureDisk
  *out = new(AzureDiskVolumeSource)
  (*in).DeepCopyInto(*out)
 }
 if in.PhotonPersistentDisk != nil {
  in, out := &in.PhotonPersistentDisk, &out.PhotonPersistentDisk
  *out = new(PhotonPersistentDiskVolumeSource)
  **out = **in
 }
 if in.Projected != nil {
  in, out := &in.Projected, &out.Projected
  *out = new(ProjectedVolumeSource)
  (*in).DeepCopyInto(*out)
 }
 if in.PortworxVolume != nil {
  in, out := &in.PortworxVolume, &out.PortworxVolume
  *out = new(PortworxVolumeSource)
  **out = **in
 }
 if in.ScaleIO != nil {
  in, out := &in.ScaleIO, &out.ScaleIO
  *out = new(ScaleIOVolumeSource)
  (*in).DeepCopyInto(*out)
 }
 if in.StorageOS != nil {
  in, out := &in.StorageOS, &out.StorageOS
  *out = new(StorageOSVolumeSource)
  (*in).DeepCopyInto(*out)
 }
 return
}
func (in *VolumeSource) DeepCopy() *VolumeSource {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(VolumeSource)
 in.DeepCopyInto(out)
 return out
}
func (in *VsphereVirtualDiskVolumeSource) DeepCopyInto(out *VsphereVirtualDiskVolumeSource) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 return
}
func (in *VsphereVirtualDiskVolumeSource) DeepCopy() *VsphereVirtualDiskVolumeSource {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(VsphereVirtualDiskVolumeSource)
 in.DeepCopyInto(out)
 return out
}
func (in *WeightedPodAffinityTerm) DeepCopyInto(out *WeightedPodAffinityTerm) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 in.PodAffinityTerm.DeepCopyInto(&out.PodAffinityTerm)
 return
}
func (in *WeightedPodAffinityTerm) DeepCopy() *WeightedPodAffinityTerm {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(WeightedPodAffinityTerm)
 in.DeepCopyInto(out)
 return out
}
