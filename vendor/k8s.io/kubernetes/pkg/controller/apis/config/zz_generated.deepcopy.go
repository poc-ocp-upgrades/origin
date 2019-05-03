package config

import (
 runtime "k8s.io/apimachinery/pkg/runtime"
)

func (in *AttachDetachControllerConfiguration) DeepCopyInto(out *AttachDetachControllerConfiguration) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.ReconcilerSyncLoopPeriod = in.ReconcilerSyncLoopPeriod
 return
}
func (in *AttachDetachControllerConfiguration) DeepCopy() *AttachDetachControllerConfiguration {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(AttachDetachControllerConfiguration)
 in.DeepCopyInto(out)
 return out
}
func (in *CSRSigningControllerConfiguration) DeepCopyInto(out *CSRSigningControllerConfiguration) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.ClusterSigningDuration = in.ClusterSigningDuration
 return
}
func (in *CSRSigningControllerConfiguration) DeepCopy() *CSRSigningControllerConfiguration {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(CSRSigningControllerConfiguration)
 in.DeepCopyInto(out)
 return out
}
func (in *CloudProviderConfiguration) DeepCopyInto(out *CloudProviderConfiguration) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 return
}
func (in *CloudProviderConfiguration) DeepCopy() *CloudProviderConfiguration {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(CloudProviderConfiguration)
 in.DeepCopyInto(out)
 return out
}
func (in *DaemonSetControllerConfiguration) DeepCopyInto(out *DaemonSetControllerConfiguration) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 return
}
func (in *DaemonSetControllerConfiguration) DeepCopy() *DaemonSetControllerConfiguration {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(DaemonSetControllerConfiguration)
 in.DeepCopyInto(out)
 return out
}
func (in *DeploymentControllerConfiguration) DeepCopyInto(out *DeploymentControllerConfiguration) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.DeploymentControllerSyncPeriod = in.DeploymentControllerSyncPeriod
 return
}
func (in *DeploymentControllerConfiguration) DeepCopy() *DeploymentControllerConfiguration {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(DeploymentControllerConfiguration)
 in.DeepCopyInto(out)
 return out
}
func (in *DeprecatedControllerConfiguration) DeepCopyInto(out *DeprecatedControllerConfiguration) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 return
}
func (in *DeprecatedControllerConfiguration) DeepCopy() *DeprecatedControllerConfiguration {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(DeprecatedControllerConfiguration)
 in.DeepCopyInto(out)
 return out
}
func (in *EndpointControllerConfiguration) DeepCopyInto(out *EndpointControllerConfiguration) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 return
}
func (in *EndpointControllerConfiguration) DeepCopy() *EndpointControllerConfiguration {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(EndpointControllerConfiguration)
 in.DeepCopyInto(out)
 return out
}
func (in *GarbageCollectorControllerConfiguration) DeepCopyInto(out *GarbageCollectorControllerConfiguration) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.GCIgnoredResources != nil {
  in, out := &in.GCIgnoredResources, &out.GCIgnoredResources
  *out = make([]GroupResource, len(*in))
  copy(*out, *in)
 }
 return
}
func (in *GarbageCollectorControllerConfiguration) DeepCopy() *GarbageCollectorControllerConfiguration {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(GarbageCollectorControllerConfiguration)
 in.DeepCopyInto(out)
 return out
}
func (in *GenericControllerManagerConfiguration) DeepCopyInto(out *GenericControllerManagerConfiguration) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.MinResyncPeriod = in.MinResyncPeriod
 out.ClientConnection = in.ClientConnection
 out.ControllerStartInterval = in.ControllerStartInterval
 out.LeaderElection = in.LeaderElection
 if in.Controllers != nil {
  in, out := &in.Controllers, &out.Controllers
  *out = make([]string, len(*in))
  copy(*out, *in)
 }
 out.Debugging = in.Debugging
 return
}
func (in *GenericControllerManagerConfiguration) DeepCopy() *GenericControllerManagerConfiguration {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(GenericControllerManagerConfiguration)
 in.DeepCopyInto(out)
 return out
}
func (in *GroupResource) DeepCopyInto(out *GroupResource) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 return
}
func (in *GroupResource) DeepCopy() *GroupResource {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(GroupResource)
 in.DeepCopyInto(out)
 return out
}
func (in *HPAControllerConfiguration) DeepCopyInto(out *HPAControllerConfiguration) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.HorizontalPodAutoscalerSyncPeriod = in.HorizontalPodAutoscalerSyncPeriod
 out.HorizontalPodAutoscalerUpscaleForbiddenWindow = in.HorizontalPodAutoscalerUpscaleForbiddenWindow
 out.HorizontalPodAutoscalerDownscaleForbiddenWindow = in.HorizontalPodAutoscalerDownscaleForbiddenWindow
 out.HorizontalPodAutoscalerDownscaleStabilizationWindow = in.HorizontalPodAutoscalerDownscaleStabilizationWindow
 out.HorizontalPodAutoscalerCPUInitializationPeriod = in.HorizontalPodAutoscalerCPUInitializationPeriod
 out.HorizontalPodAutoscalerInitialReadinessDelay = in.HorizontalPodAutoscalerInitialReadinessDelay
 return
}
func (in *HPAControllerConfiguration) DeepCopy() *HPAControllerConfiguration {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(HPAControllerConfiguration)
 in.DeepCopyInto(out)
 return out
}
func (in *JobControllerConfiguration) DeepCopyInto(out *JobControllerConfiguration) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 return
}
func (in *JobControllerConfiguration) DeepCopy() *JobControllerConfiguration {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(JobControllerConfiguration)
 in.DeepCopyInto(out)
 return out
}
func (in *KubeCloudSharedConfiguration) DeepCopyInto(out *KubeCloudSharedConfiguration) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.CloudProvider = in.CloudProvider
 out.RouteReconciliationPeriod = in.RouteReconciliationPeriod
 out.NodeMonitorPeriod = in.NodeMonitorPeriod
 out.NodeSyncPeriod = in.NodeSyncPeriod
 return
}
func (in *KubeCloudSharedConfiguration) DeepCopy() *KubeCloudSharedConfiguration {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(KubeCloudSharedConfiguration)
 in.DeepCopyInto(out)
 return out
}
func (in *KubeControllerManagerConfiguration) DeepCopyInto(out *KubeControllerManagerConfiguration) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 in.Generic.DeepCopyInto(&out.Generic)
 out.KubeCloudShared = in.KubeCloudShared
 out.AttachDetachController = in.AttachDetachController
 out.CSRSigningController = in.CSRSigningController
 out.DaemonSetController = in.DaemonSetController
 out.DeploymentController = in.DeploymentController
 out.DeprecatedController = in.DeprecatedController
 out.EndpointController = in.EndpointController
 in.GarbageCollectorController.DeepCopyInto(&out.GarbageCollectorController)
 out.HPAController = in.HPAController
 out.JobController = in.JobController
 out.NamespaceController = in.NamespaceController
 out.NodeIPAMController = in.NodeIPAMController
 out.NodeLifecycleController = in.NodeLifecycleController
 out.PersistentVolumeBinderController = in.PersistentVolumeBinderController
 out.PodGCController = in.PodGCController
 out.ReplicaSetController = in.ReplicaSetController
 out.ReplicationController = in.ReplicationController
 out.ResourceQuotaController = in.ResourceQuotaController
 out.SAController = in.SAController
 out.ServiceController = in.ServiceController
 out.TTLAfterFinishedController = in.TTLAfterFinishedController
 return
}
func (in *KubeControllerManagerConfiguration) DeepCopy() *KubeControllerManagerConfiguration {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(KubeControllerManagerConfiguration)
 in.DeepCopyInto(out)
 return out
}
func (in *KubeControllerManagerConfiguration) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *NamespaceControllerConfiguration) DeepCopyInto(out *NamespaceControllerConfiguration) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.NamespaceSyncPeriod = in.NamespaceSyncPeriod
 return
}
func (in *NamespaceControllerConfiguration) DeepCopy() *NamespaceControllerConfiguration {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(NamespaceControllerConfiguration)
 in.DeepCopyInto(out)
 return out
}
func (in *NodeIPAMControllerConfiguration) DeepCopyInto(out *NodeIPAMControllerConfiguration) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 return
}
func (in *NodeIPAMControllerConfiguration) DeepCopy() *NodeIPAMControllerConfiguration {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(NodeIPAMControllerConfiguration)
 in.DeepCopyInto(out)
 return out
}
func (in *NodeLifecycleControllerConfiguration) DeepCopyInto(out *NodeLifecycleControllerConfiguration) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.NodeStartupGracePeriod = in.NodeStartupGracePeriod
 out.NodeMonitorGracePeriod = in.NodeMonitorGracePeriod
 out.PodEvictionTimeout = in.PodEvictionTimeout
 return
}
func (in *NodeLifecycleControllerConfiguration) DeepCopy() *NodeLifecycleControllerConfiguration {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(NodeLifecycleControllerConfiguration)
 in.DeepCopyInto(out)
 return out
}
func (in *PersistentVolumeBinderControllerConfiguration) DeepCopyInto(out *PersistentVolumeBinderControllerConfiguration) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.PVClaimBinderSyncPeriod = in.PVClaimBinderSyncPeriod
 out.VolumeConfiguration = in.VolumeConfiguration
 return
}
func (in *PersistentVolumeBinderControllerConfiguration) DeepCopy() *PersistentVolumeBinderControllerConfiguration {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(PersistentVolumeBinderControllerConfiguration)
 in.DeepCopyInto(out)
 return out
}
func (in *PersistentVolumeRecyclerConfiguration) DeepCopyInto(out *PersistentVolumeRecyclerConfiguration) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 return
}
func (in *PersistentVolumeRecyclerConfiguration) DeepCopy() *PersistentVolumeRecyclerConfiguration {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(PersistentVolumeRecyclerConfiguration)
 in.DeepCopyInto(out)
 return out
}
func (in *PodGCControllerConfiguration) DeepCopyInto(out *PodGCControllerConfiguration) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 return
}
func (in *PodGCControllerConfiguration) DeepCopy() *PodGCControllerConfiguration {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(PodGCControllerConfiguration)
 in.DeepCopyInto(out)
 return out
}
func (in *ReplicaSetControllerConfiguration) DeepCopyInto(out *ReplicaSetControllerConfiguration) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 return
}
func (in *ReplicaSetControllerConfiguration) DeepCopy() *ReplicaSetControllerConfiguration {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ReplicaSetControllerConfiguration)
 in.DeepCopyInto(out)
 return out
}
func (in *ReplicationControllerConfiguration) DeepCopyInto(out *ReplicationControllerConfiguration) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 return
}
func (in *ReplicationControllerConfiguration) DeepCopy() *ReplicationControllerConfiguration {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ReplicationControllerConfiguration)
 in.DeepCopyInto(out)
 return out
}
func (in *ResourceQuotaControllerConfiguration) DeepCopyInto(out *ResourceQuotaControllerConfiguration) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.ResourceQuotaSyncPeriod = in.ResourceQuotaSyncPeriod
 return
}
func (in *ResourceQuotaControllerConfiguration) DeepCopy() *ResourceQuotaControllerConfiguration {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ResourceQuotaControllerConfiguration)
 in.DeepCopyInto(out)
 return out
}
func (in *SAControllerConfiguration) DeepCopyInto(out *SAControllerConfiguration) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 return
}
func (in *SAControllerConfiguration) DeepCopy() *SAControllerConfiguration {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(SAControllerConfiguration)
 in.DeepCopyInto(out)
 return out
}
func (in *ServiceControllerConfiguration) DeepCopyInto(out *ServiceControllerConfiguration) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 return
}
func (in *ServiceControllerConfiguration) DeepCopy() *ServiceControllerConfiguration {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ServiceControllerConfiguration)
 in.DeepCopyInto(out)
 return out
}
func (in *TTLAfterFinishedControllerConfiguration) DeepCopyInto(out *TTLAfterFinishedControllerConfiguration) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 return
}
func (in *TTLAfterFinishedControllerConfiguration) DeepCopy() *TTLAfterFinishedControllerConfiguration {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(TTLAfterFinishedControllerConfiguration)
 in.DeepCopyInto(out)
 return out
}
func (in *VolumeConfiguration) DeepCopyInto(out *VolumeConfiguration) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.PersistentVolumeRecyclerConfiguration = in.PersistentVolumeRecyclerConfiguration
 return
}
func (in *VolumeConfiguration) DeepCopy() *VolumeConfiguration {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(VolumeConfiguration)
 in.DeepCopyInto(out)
 return out
}
