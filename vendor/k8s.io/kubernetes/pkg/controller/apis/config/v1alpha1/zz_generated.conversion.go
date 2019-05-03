package v1alpha1

import (
 unsafe "unsafe"
 apisconfigv1alpha1 "k8s.io/apimachinery/pkg/apis/config/v1alpha1"
 v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 conversion "k8s.io/apimachinery/pkg/conversion"
 runtime "k8s.io/apimachinery/pkg/runtime"
 pkgapisconfigv1alpha1 "k8s.io/apiserver/pkg/apis/config/v1alpha1"
 v1alpha1 "k8s.io/kube-controller-manager/config/v1alpha1"
 config "k8s.io/kubernetes/pkg/controller/apis/config"
)

func init() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 localSchemeBuilder.Register(RegisterConversions)
}
func RegisterConversions(s *runtime.Scheme) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := s.AddGeneratedConversionFunc((*v1alpha1.AttachDetachControllerConfiguration)(nil), (*config.AttachDetachControllerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1alpha1_AttachDetachControllerConfiguration_To_config_AttachDetachControllerConfiguration(a.(*v1alpha1.AttachDetachControllerConfiguration), b.(*config.AttachDetachControllerConfiguration), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*config.AttachDetachControllerConfiguration)(nil), (*v1alpha1.AttachDetachControllerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_config_AttachDetachControllerConfiguration_To_v1alpha1_AttachDetachControllerConfiguration(a.(*config.AttachDetachControllerConfiguration), b.(*v1alpha1.AttachDetachControllerConfiguration), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1alpha1.CSRSigningControllerConfiguration)(nil), (*config.CSRSigningControllerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1alpha1_CSRSigningControllerConfiguration_To_config_CSRSigningControllerConfiguration(a.(*v1alpha1.CSRSigningControllerConfiguration), b.(*config.CSRSigningControllerConfiguration), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*config.CSRSigningControllerConfiguration)(nil), (*v1alpha1.CSRSigningControllerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_config_CSRSigningControllerConfiguration_To_v1alpha1_CSRSigningControllerConfiguration(a.(*config.CSRSigningControllerConfiguration), b.(*v1alpha1.CSRSigningControllerConfiguration), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1alpha1.CloudProviderConfiguration)(nil), (*config.CloudProviderConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1alpha1_CloudProviderConfiguration_To_config_CloudProviderConfiguration(a.(*v1alpha1.CloudProviderConfiguration), b.(*config.CloudProviderConfiguration), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*config.CloudProviderConfiguration)(nil), (*v1alpha1.CloudProviderConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_config_CloudProviderConfiguration_To_v1alpha1_CloudProviderConfiguration(a.(*config.CloudProviderConfiguration), b.(*v1alpha1.CloudProviderConfiguration), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1alpha1.DaemonSetControllerConfiguration)(nil), (*config.DaemonSetControllerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1alpha1_DaemonSetControllerConfiguration_To_config_DaemonSetControllerConfiguration(a.(*v1alpha1.DaemonSetControllerConfiguration), b.(*config.DaemonSetControllerConfiguration), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*config.DaemonSetControllerConfiguration)(nil), (*v1alpha1.DaemonSetControllerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_config_DaemonSetControllerConfiguration_To_v1alpha1_DaemonSetControllerConfiguration(a.(*config.DaemonSetControllerConfiguration), b.(*v1alpha1.DaemonSetControllerConfiguration), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1alpha1.DeploymentControllerConfiguration)(nil), (*config.DeploymentControllerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1alpha1_DeploymentControllerConfiguration_To_config_DeploymentControllerConfiguration(a.(*v1alpha1.DeploymentControllerConfiguration), b.(*config.DeploymentControllerConfiguration), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*config.DeploymentControllerConfiguration)(nil), (*v1alpha1.DeploymentControllerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_config_DeploymentControllerConfiguration_To_v1alpha1_DeploymentControllerConfiguration(a.(*config.DeploymentControllerConfiguration), b.(*v1alpha1.DeploymentControllerConfiguration), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1alpha1.DeprecatedControllerConfiguration)(nil), (*config.DeprecatedControllerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1alpha1_DeprecatedControllerConfiguration_To_config_DeprecatedControllerConfiguration(a.(*v1alpha1.DeprecatedControllerConfiguration), b.(*config.DeprecatedControllerConfiguration), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*config.DeprecatedControllerConfiguration)(nil), (*v1alpha1.DeprecatedControllerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_config_DeprecatedControllerConfiguration_To_v1alpha1_DeprecatedControllerConfiguration(a.(*config.DeprecatedControllerConfiguration), b.(*v1alpha1.DeprecatedControllerConfiguration), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1alpha1.EndpointControllerConfiguration)(nil), (*config.EndpointControllerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1alpha1_EndpointControllerConfiguration_To_config_EndpointControllerConfiguration(a.(*v1alpha1.EndpointControllerConfiguration), b.(*config.EndpointControllerConfiguration), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*config.EndpointControllerConfiguration)(nil), (*v1alpha1.EndpointControllerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_config_EndpointControllerConfiguration_To_v1alpha1_EndpointControllerConfiguration(a.(*config.EndpointControllerConfiguration), b.(*v1alpha1.EndpointControllerConfiguration), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1alpha1.GarbageCollectorControllerConfiguration)(nil), (*config.GarbageCollectorControllerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1alpha1_GarbageCollectorControllerConfiguration_To_config_GarbageCollectorControllerConfiguration(a.(*v1alpha1.GarbageCollectorControllerConfiguration), b.(*config.GarbageCollectorControllerConfiguration), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*config.GarbageCollectorControllerConfiguration)(nil), (*v1alpha1.GarbageCollectorControllerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_config_GarbageCollectorControllerConfiguration_To_v1alpha1_GarbageCollectorControllerConfiguration(a.(*config.GarbageCollectorControllerConfiguration), b.(*v1alpha1.GarbageCollectorControllerConfiguration), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1alpha1.GenericControllerManagerConfiguration)(nil), (*config.GenericControllerManagerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1alpha1_GenericControllerManagerConfiguration_To_config_GenericControllerManagerConfiguration(a.(*v1alpha1.GenericControllerManagerConfiguration), b.(*config.GenericControllerManagerConfiguration), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*config.GenericControllerManagerConfiguration)(nil), (*v1alpha1.GenericControllerManagerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_config_GenericControllerManagerConfiguration_To_v1alpha1_GenericControllerManagerConfiguration(a.(*config.GenericControllerManagerConfiguration), b.(*v1alpha1.GenericControllerManagerConfiguration), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1alpha1.GroupResource)(nil), (*config.GroupResource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1alpha1_GroupResource_To_config_GroupResource(a.(*v1alpha1.GroupResource), b.(*config.GroupResource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*config.GroupResource)(nil), (*v1alpha1.GroupResource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_config_GroupResource_To_v1alpha1_GroupResource(a.(*config.GroupResource), b.(*v1alpha1.GroupResource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1alpha1.HPAControllerConfiguration)(nil), (*config.HPAControllerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1alpha1_HPAControllerConfiguration_To_config_HPAControllerConfiguration(a.(*v1alpha1.HPAControllerConfiguration), b.(*config.HPAControllerConfiguration), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*config.HPAControllerConfiguration)(nil), (*v1alpha1.HPAControllerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_config_HPAControllerConfiguration_To_v1alpha1_HPAControllerConfiguration(a.(*config.HPAControllerConfiguration), b.(*v1alpha1.HPAControllerConfiguration), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1alpha1.JobControllerConfiguration)(nil), (*config.JobControllerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1alpha1_JobControllerConfiguration_To_config_JobControllerConfiguration(a.(*v1alpha1.JobControllerConfiguration), b.(*config.JobControllerConfiguration), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*config.JobControllerConfiguration)(nil), (*v1alpha1.JobControllerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_config_JobControllerConfiguration_To_v1alpha1_JobControllerConfiguration(a.(*config.JobControllerConfiguration), b.(*v1alpha1.JobControllerConfiguration), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1alpha1.KubeCloudSharedConfiguration)(nil), (*config.KubeCloudSharedConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1alpha1_KubeCloudSharedConfiguration_To_config_KubeCloudSharedConfiguration(a.(*v1alpha1.KubeCloudSharedConfiguration), b.(*config.KubeCloudSharedConfiguration), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*config.KubeCloudSharedConfiguration)(nil), (*v1alpha1.KubeCloudSharedConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_config_KubeCloudSharedConfiguration_To_v1alpha1_KubeCloudSharedConfiguration(a.(*config.KubeCloudSharedConfiguration), b.(*v1alpha1.KubeCloudSharedConfiguration), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1alpha1.KubeControllerManagerConfiguration)(nil), (*config.KubeControllerManagerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1alpha1_KubeControllerManagerConfiguration_To_config_KubeControllerManagerConfiguration(a.(*v1alpha1.KubeControllerManagerConfiguration), b.(*config.KubeControllerManagerConfiguration), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*config.KubeControllerManagerConfiguration)(nil), (*v1alpha1.KubeControllerManagerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_config_KubeControllerManagerConfiguration_To_v1alpha1_KubeControllerManagerConfiguration(a.(*config.KubeControllerManagerConfiguration), b.(*v1alpha1.KubeControllerManagerConfiguration), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1alpha1.NamespaceControllerConfiguration)(nil), (*config.NamespaceControllerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1alpha1_NamespaceControllerConfiguration_To_config_NamespaceControllerConfiguration(a.(*v1alpha1.NamespaceControllerConfiguration), b.(*config.NamespaceControllerConfiguration), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*config.NamespaceControllerConfiguration)(nil), (*v1alpha1.NamespaceControllerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_config_NamespaceControllerConfiguration_To_v1alpha1_NamespaceControllerConfiguration(a.(*config.NamespaceControllerConfiguration), b.(*v1alpha1.NamespaceControllerConfiguration), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1alpha1.NodeIPAMControllerConfiguration)(nil), (*config.NodeIPAMControllerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1alpha1_NodeIPAMControllerConfiguration_To_config_NodeIPAMControllerConfiguration(a.(*v1alpha1.NodeIPAMControllerConfiguration), b.(*config.NodeIPAMControllerConfiguration), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*config.NodeIPAMControllerConfiguration)(nil), (*v1alpha1.NodeIPAMControllerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_config_NodeIPAMControllerConfiguration_To_v1alpha1_NodeIPAMControllerConfiguration(a.(*config.NodeIPAMControllerConfiguration), b.(*v1alpha1.NodeIPAMControllerConfiguration), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1alpha1.NodeLifecycleControllerConfiguration)(nil), (*config.NodeLifecycleControllerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1alpha1_NodeLifecycleControllerConfiguration_To_config_NodeLifecycleControllerConfiguration(a.(*v1alpha1.NodeLifecycleControllerConfiguration), b.(*config.NodeLifecycleControllerConfiguration), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*config.NodeLifecycleControllerConfiguration)(nil), (*v1alpha1.NodeLifecycleControllerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_config_NodeLifecycleControllerConfiguration_To_v1alpha1_NodeLifecycleControllerConfiguration(a.(*config.NodeLifecycleControllerConfiguration), b.(*v1alpha1.NodeLifecycleControllerConfiguration), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1alpha1.PersistentVolumeBinderControllerConfiguration)(nil), (*config.PersistentVolumeBinderControllerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1alpha1_PersistentVolumeBinderControllerConfiguration_To_config_PersistentVolumeBinderControllerConfiguration(a.(*v1alpha1.PersistentVolumeBinderControllerConfiguration), b.(*config.PersistentVolumeBinderControllerConfiguration), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*config.PersistentVolumeBinderControllerConfiguration)(nil), (*v1alpha1.PersistentVolumeBinderControllerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_config_PersistentVolumeBinderControllerConfiguration_To_v1alpha1_PersistentVolumeBinderControllerConfiguration(a.(*config.PersistentVolumeBinderControllerConfiguration), b.(*v1alpha1.PersistentVolumeBinderControllerConfiguration), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1alpha1.PersistentVolumeRecyclerConfiguration)(nil), (*config.PersistentVolumeRecyclerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1alpha1_PersistentVolumeRecyclerConfiguration_To_config_PersistentVolumeRecyclerConfiguration(a.(*v1alpha1.PersistentVolumeRecyclerConfiguration), b.(*config.PersistentVolumeRecyclerConfiguration), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*config.PersistentVolumeRecyclerConfiguration)(nil), (*v1alpha1.PersistentVolumeRecyclerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_config_PersistentVolumeRecyclerConfiguration_To_v1alpha1_PersistentVolumeRecyclerConfiguration(a.(*config.PersistentVolumeRecyclerConfiguration), b.(*v1alpha1.PersistentVolumeRecyclerConfiguration), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1alpha1.PodGCControllerConfiguration)(nil), (*config.PodGCControllerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1alpha1_PodGCControllerConfiguration_To_config_PodGCControllerConfiguration(a.(*v1alpha1.PodGCControllerConfiguration), b.(*config.PodGCControllerConfiguration), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*config.PodGCControllerConfiguration)(nil), (*v1alpha1.PodGCControllerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_config_PodGCControllerConfiguration_To_v1alpha1_PodGCControllerConfiguration(a.(*config.PodGCControllerConfiguration), b.(*v1alpha1.PodGCControllerConfiguration), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1alpha1.ReplicaSetControllerConfiguration)(nil), (*config.ReplicaSetControllerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1alpha1_ReplicaSetControllerConfiguration_To_config_ReplicaSetControllerConfiguration(a.(*v1alpha1.ReplicaSetControllerConfiguration), b.(*config.ReplicaSetControllerConfiguration), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*config.ReplicaSetControllerConfiguration)(nil), (*v1alpha1.ReplicaSetControllerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_config_ReplicaSetControllerConfiguration_To_v1alpha1_ReplicaSetControllerConfiguration(a.(*config.ReplicaSetControllerConfiguration), b.(*v1alpha1.ReplicaSetControllerConfiguration), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1alpha1.ReplicationControllerConfiguration)(nil), (*config.ReplicationControllerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1alpha1_ReplicationControllerConfiguration_To_config_ReplicationControllerConfiguration(a.(*v1alpha1.ReplicationControllerConfiguration), b.(*config.ReplicationControllerConfiguration), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*config.ReplicationControllerConfiguration)(nil), (*v1alpha1.ReplicationControllerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_config_ReplicationControllerConfiguration_To_v1alpha1_ReplicationControllerConfiguration(a.(*config.ReplicationControllerConfiguration), b.(*v1alpha1.ReplicationControllerConfiguration), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1alpha1.ResourceQuotaControllerConfiguration)(nil), (*config.ResourceQuotaControllerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1alpha1_ResourceQuotaControllerConfiguration_To_config_ResourceQuotaControllerConfiguration(a.(*v1alpha1.ResourceQuotaControllerConfiguration), b.(*config.ResourceQuotaControllerConfiguration), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*config.ResourceQuotaControllerConfiguration)(nil), (*v1alpha1.ResourceQuotaControllerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_config_ResourceQuotaControllerConfiguration_To_v1alpha1_ResourceQuotaControllerConfiguration(a.(*config.ResourceQuotaControllerConfiguration), b.(*v1alpha1.ResourceQuotaControllerConfiguration), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1alpha1.SAControllerConfiguration)(nil), (*config.SAControllerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1alpha1_SAControllerConfiguration_To_config_SAControllerConfiguration(a.(*v1alpha1.SAControllerConfiguration), b.(*config.SAControllerConfiguration), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*config.SAControllerConfiguration)(nil), (*v1alpha1.SAControllerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_config_SAControllerConfiguration_To_v1alpha1_SAControllerConfiguration(a.(*config.SAControllerConfiguration), b.(*v1alpha1.SAControllerConfiguration), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1alpha1.ServiceControllerConfiguration)(nil), (*config.ServiceControllerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1alpha1_ServiceControllerConfiguration_To_config_ServiceControllerConfiguration(a.(*v1alpha1.ServiceControllerConfiguration), b.(*config.ServiceControllerConfiguration), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*config.ServiceControllerConfiguration)(nil), (*v1alpha1.ServiceControllerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_config_ServiceControllerConfiguration_To_v1alpha1_ServiceControllerConfiguration(a.(*config.ServiceControllerConfiguration), b.(*v1alpha1.ServiceControllerConfiguration), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1alpha1.TTLAfterFinishedControllerConfiguration)(nil), (*config.TTLAfterFinishedControllerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1alpha1_TTLAfterFinishedControllerConfiguration_To_config_TTLAfterFinishedControllerConfiguration(a.(*v1alpha1.TTLAfterFinishedControllerConfiguration), b.(*config.TTLAfterFinishedControllerConfiguration), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*config.TTLAfterFinishedControllerConfiguration)(nil), (*v1alpha1.TTLAfterFinishedControllerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_config_TTLAfterFinishedControllerConfiguration_To_v1alpha1_TTLAfterFinishedControllerConfiguration(a.(*config.TTLAfterFinishedControllerConfiguration), b.(*v1alpha1.TTLAfterFinishedControllerConfiguration), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1alpha1.VolumeConfiguration)(nil), (*config.VolumeConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1alpha1_VolumeConfiguration_To_config_VolumeConfiguration(a.(*v1alpha1.VolumeConfiguration), b.(*config.VolumeConfiguration), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*config.VolumeConfiguration)(nil), (*v1alpha1.VolumeConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_config_VolumeConfiguration_To_v1alpha1_VolumeConfiguration(a.(*config.VolumeConfiguration), b.(*v1alpha1.VolumeConfiguration), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*config.GenericControllerManagerConfiguration)(nil), (*v1alpha1.GenericControllerManagerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_config_GenericControllerManagerConfiguration_To_v1alpha1_GenericControllerManagerConfiguration(a.(*config.GenericControllerManagerConfiguration), b.(*v1alpha1.GenericControllerManagerConfiguration), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*config.KubeCloudSharedConfiguration)(nil), (*v1alpha1.KubeCloudSharedConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_config_KubeCloudSharedConfiguration_To_v1alpha1_KubeCloudSharedConfiguration(a.(*config.KubeCloudSharedConfiguration), b.(*v1alpha1.KubeCloudSharedConfiguration), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*config.ServiceControllerConfiguration)(nil), (*v1alpha1.ServiceControllerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_config_ServiceControllerConfiguration_To_v1alpha1_ServiceControllerConfiguration(a.(*config.ServiceControllerConfiguration), b.(*v1alpha1.ServiceControllerConfiguration), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*v1alpha1.GenericControllerManagerConfiguration)(nil), (*config.GenericControllerManagerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1alpha1_GenericControllerManagerConfiguration_To_config_GenericControllerManagerConfiguration(a.(*v1alpha1.GenericControllerManagerConfiguration), b.(*config.GenericControllerManagerConfiguration), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*v1alpha1.KubeCloudSharedConfiguration)(nil), (*config.KubeCloudSharedConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1alpha1_KubeCloudSharedConfiguration_To_config_KubeCloudSharedConfiguration(a.(*v1alpha1.KubeCloudSharedConfiguration), b.(*config.KubeCloudSharedConfiguration), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*v1alpha1.ServiceControllerConfiguration)(nil), (*config.ServiceControllerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1alpha1_ServiceControllerConfiguration_To_config_ServiceControllerConfiguration(a.(*v1alpha1.ServiceControllerConfiguration), b.(*config.ServiceControllerConfiguration), scope)
 }); err != nil {
  return err
 }
 return nil
}
func autoConvert_v1alpha1_AttachDetachControllerConfiguration_To_config_AttachDetachControllerConfiguration(in *v1alpha1.AttachDetachControllerConfiguration, out *config.AttachDetachControllerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.DisableAttachDetachReconcilerSync = in.DisableAttachDetachReconcilerSync
 out.ReconcilerSyncLoopPeriod = in.ReconcilerSyncLoopPeriod
 return nil
}
func Convert_v1alpha1_AttachDetachControllerConfiguration_To_config_AttachDetachControllerConfiguration(in *v1alpha1.AttachDetachControllerConfiguration, out *config.AttachDetachControllerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1alpha1_AttachDetachControllerConfiguration_To_config_AttachDetachControllerConfiguration(in, out, s)
}
func autoConvert_config_AttachDetachControllerConfiguration_To_v1alpha1_AttachDetachControllerConfiguration(in *config.AttachDetachControllerConfiguration, out *v1alpha1.AttachDetachControllerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.DisableAttachDetachReconcilerSync = in.DisableAttachDetachReconcilerSync
 out.ReconcilerSyncLoopPeriod = in.ReconcilerSyncLoopPeriod
 return nil
}
func Convert_config_AttachDetachControllerConfiguration_To_v1alpha1_AttachDetachControllerConfiguration(in *config.AttachDetachControllerConfiguration, out *v1alpha1.AttachDetachControllerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_config_AttachDetachControllerConfiguration_To_v1alpha1_AttachDetachControllerConfiguration(in, out, s)
}
func autoConvert_v1alpha1_CSRSigningControllerConfiguration_To_config_CSRSigningControllerConfiguration(in *v1alpha1.CSRSigningControllerConfiguration, out *config.CSRSigningControllerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ClusterSigningCertFile = in.ClusterSigningCertFile
 out.ClusterSigningKeyFile = in.ClusterSigningKeyFile
 out.ClusterSigningDuration = in.ClusterSigningDuration
 return nil
}
func Convert_v1alpha1_CSRSigningControllerConfiguration_To_config_CSRSigningControllerConfiguration(in *v1alpha1.CSRSigningControllerConfiguration, out *config.CSRSigningControllerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1alpha1_CSRSigningControllerConfiguration_To_config_CSRSigningControllerConfiguration(in, out, s)
}
func autoConvert_config_CSRSigningControllerConfiguration_To_v1alpha1_CSRSigningControllerConfiguration(in *config.CSRSigningControllerConfiguration, out *v1alpha1.CSRSigningControllerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ClusterSigningCertFile = in.ClusterSigningCertFile
 out.ClusterSigningKeyFile = in.ClusterSigningKeyFile
 out.ClusterSigningDuration = in.ClusterSigningDuration
 return nil
}
func Convert_config_CSRSigningControllerConfiguration_To_v1alpha1_CSRSigningControllerConfiguration(in *config.CSRSigningControllerConfiguration, out *v1alpha1.CSRSigningControllerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_config_CSRSigningControllerConfiguration_To_v1alpha1_CSRSigningControllerConfiguration(in, out, s)
}
func autoConvert_v1alpha1_CloudProviderConfiguration_To_config_CloudProviderConfiguration(in *v1alpha1.CloudProviderConfiguration, out *config.CloudProviderConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Name = in.Name
 out.CloudConfigFile = in.CloudConfigFile
 return nil
}
func Convert_v1alpha1_CloudProviderConfiguration_To_config_CloudProviderConfiguration(in *v1alpha1.CloudProviderConfiguration, out *config.CloudProviderConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1alpha1_CloudProviderConfiguration_To_config_CloudProviderConfiguration(in, out, s)
}
func autoConvert_config_CloudProviderConfiguration_To_v1alpha1_CloudProviderConfiguration(in *config.CloudProviderConfiguration, out *v1alpha1.CloudProviderConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Name = in.Name
 out.CloudConfigFile = in.CloudConfigFile
 return nil
}
func Convert_config_CloudProviderConfiguration_To_v1alpha1_CloudProviderConfiguration(in *config.CloudProviderConfiguration, out *v1alpha1.CloudProviderConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_config_CloudProviderConfiguration_To_v1alpha1_CloudProviderConfiguration(in, out, s)
}
func autoConvert_v1alpha1_DaemonSetControllerConfiguration_To_config_DaemonSetControllerConfiguration(in *v1alpha1.DaemonSetControllerConfiguration, out *config.DaemonSetControllerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ConcurrentDaemonSetSyncs = in.ConcurrentDaemonSetSyncs
 return nil
}
func Convert_v1alpha1_DaemonSetControllerConfiguration_To_config_DaemonSetControllerConfiguration(in *v1alpha1.DaemonSetControllerConfiguration, out *config.DaemonSetControllerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1alpha1_DaemonSetControllerConfiguration_To_config_DaemonSetControllerConfiguration(in, out, s)
}
func autoConvert_config_DaemonSetControllerConfiguration_To_v1alpha1_DaemonSetControllerConfiguration(in *config.DaemonSetControllerConfiguration, out *v1alpha1.DaemonSetControllerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ConcurrentDaemonSetSyncs = in.ConcurrentDaemonSetSyncs
 return nil
}
func Convert_config_DaemonSetControllerConfiguration_To_v1alpha1_DaemonSetControllerConfiguration(in *config.DaemonSetControllerConfiguration, out *v1alpha1.DaemonSetControllerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_config_DaemonSetControllerConfiguration_To_v1alpha1_DaemonSetControllerConfiguration(in, out, s)
}
func autoConvert_v1alpha1_DeploymentControllerConfiguration_To_config_DeploymentControllerConfiguration(in *v1alpha1.DeploymentControllerConfiguration, out *config.DeploymentControllerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ConcurrentDeploymentSyncs = in.ConcurrentDeploymentSyncs
 out.DeploymentControllerSyncPeriod = in.DeploymentControllerSyncPeriod
 return nil
}
func Convert_v1alpha1_DeploymentControllerConfiguration_To_config_DeploymentControllerConfiguration(in *v1alpha1.DeploymentControllerConfiguration, out *config.DeploymentControllerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1alpha1_DeploymentControllerConfiguration_To_config_DeploymentControllerConfiguration(in, out, s)
}
func autoConvert_config_DeploymentControllerConfiguration_To_v1alpha1_DeploymentControllerConfiguration(in *config.DeploymentControllerConfiguration, out *v1alpha1.DeploymentControllerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ConcurrentDeploymentSyncs = in.ConcurrentDeploymentSyncs
 out.DeploymentControllerSyncPeriod = in.DeploymentControllerSyncPeriod
 return nil
}
func Convert_config_DeploymentControllerConfiguration_To_v1alpha1_DeploymentControllerConfiguration(in *config.DeploymentControllerConfiguration, out *v1alpha1.DeploymentControllerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_config_DeploymentControllerConfiguration_To_v1alpha1_DeploymentControllerConfiguration(in, out, s)
}
func autoConvert_v1alpha1_DeprecatedControllerConfiguration_To_config_DeprecatedControllerConfiguration(in *v1alpha1.DeprecatedControllerConfiguration, out *config.DeprecatedControllerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.DeletingPodsQPS = in.DeletingPodsQPS
 out.DeletingPodsBurst = in.DeletingPodsBurst
 out.RegisterRetryCount = in.RegisterRetryCount
 return nil
}
func Convert_v1alpha1_DeprecatedControllerConfiguration_To_config_DeprecatedControllerConfiguration(in *v1alpha1.DeprecatedControllerConfiguration, out *config.DeprecatedControllerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1alpha1_DeprecatedControllerConfiguration_To_config_DeprecatedControllerConfiguration(in, out, s)
}
func autoConvert_config_DeprecatedControllerConfiguration_To_v1alpha1_DeprecatedControllerConfiguration(in *config.DeprecatedControllerConfiguration, out *v1alpha1.DeprecatedControllerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.DeletingPodsQPS = in.DeletingPodsQPS
 out.DeletingPodsBurst = in.DeletingPodsBurst
 out.RegisterRetryCount = in.RegisterRetryCount
 return nil
}
func Convert_config_DeprecatedControllerConfiguration_To_v1alpha1_DeprecatedControllerConfiguration(in *config.DeprecatedControllerConfiguration, out *v1alpha1.DeprecatedControllerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_config_DeprecatedControllerConfiguration_To_v1alpha1_DeprecatedControllerConfiguration(in, out, s)
}
func autoConvert_v1alpha1_EndpointControllerConfiguration_To_config_EndpointControllerConfiguration(in *v1alpha1.EndpointControllerConfiguration, out *config.EndpointControllerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ConcurrentEndpointSyncs = in.ConcurrentEndpointSyncs
 return nil
}
func Convert_v1alpha1_EndpointControllerConfiguration_To_config_EndpointControllerConfiguration(in *v1alpha1.EndpointControllerConfiguration, out *config.EndpointControllerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1alpha1_EndpointControllerConfiguration_To_config_EndpointControllerConfiguration(in, out, s)
}
func autoConvert_config_EndpointControllerConfiguration_To_v1alpha1_EndpointControllerConfiguration(in *config.EndpointControllerConfiguration, out *v1alpha1.EndpointControllerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ConcurrentEndpointSyncs = in.ConcurrentEndpointSyncs
 return nil
}
func Convert_config_EndpointControllerConfiguration_To_v1alpha1_EndpointControllerConfiguration(in *config.EndpointControllerConfiguration, out *v1alpha1.EndpointControllerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_config_EndpointControllerConfiguration_To_v1alpha1_EndpointControllerConfiguration(in, out, s)
}
func autoConvert_v1alpha1_GarbageCollectorControllerConfiguration_To_config_GarbageCollectorControllerConfiguration(in *v1alpha1.GarbageCollectorControllerConfiguration, out *config.GarbageCollectorControllerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := v1.Convert_Pointer_bool_To_bool(&in.EnableGarbageCollector, &out.EnableGarbageCollector, s); err != nil {
  return err
 }
 out.ConcurrentGCSyncs = in.ConcurrentGCSyncs
 out.GCIgnoredResources = *(*[]config.GroupResource)(unsafe.Pointer(&in.GCIgnoredResources))
 return nil
}
func Convert_v1alpha1_GarbageCollectorControllerConfiguration_To_config_GarbageCollectorControllerConfiguration(in *v1alpha1.GarbageCollectorControllerConfiguration, out *config.GarbageCollectorControllerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1alpha1_GarbageCollectorControllerConfiguration_To_config_GarbageCollectorControllerConfiguration(in, out, s)
}
func autoConvert_config_GarbageCollectorControllerConfiguration_To_v1alpha1_GarbageCollectorControllerConfiguration(in *config.GarbageCollectorControllerConfiguration, out *v1alpha1.GarbageCollectorControllerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := v1.Convert_bool_To_Pointer_bool(&in.EnableGarbageCollector, &out.EnableGarbageCollector, s); err != nil {
  return err
 }
 out.ConcurrentGCSyncs = in.ConcurrentGCSyncs
 out.GCIgnoredResources = *(*[]v1alpha1.GroupResource)(unsafe.Pointer(&in.GCIgnoredResources))
 return nil
}
func Convert_config_GarbageCollectorControllerConfiguration_To_v1alpha1_GarbageCollectorControllerConfiguration(in *config.GarbageCollectorControllerConfiguration, out *v1alpha1.GarbageCollectorControllerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_config_GarbageCollectorControllerConfiguration_To_v1alpha1_GarbageCollectorControllerConfiguration(in, out, s)
}
func autoConvert_v1alpha1_GenericControllerManagerConfiguration_To_config_GenericControllerManagerConfiguration(in *v1alpha1.GenericControllerManagerConfiguration, out *config.GenericControllerManagerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Port = in.Port
 out.Address = in.Address
 out.MinResyncPeriod = in.MinResyncPeriod
 if err := apisconfigv1alpha1.Convert_v1alpha1_ClientConnectionConfiguration_To_config_ClientConnectionConfiguration(&in.ClientConnection, &out.ClientConnection, s); err != nil {
  return err
 }
 out.ControllerStartInterval = in.ControllerStartInterval
 if err := pkgapisconfigv1alpha1.Convert_v1alpha1_LeaderElectionConfiguration_To_config_LeaderElectionConfiguration(&in.LeaderElection, &out.LeaderElection, s); err != nil {
  return err
 }
 out.Controllers = *(*[]string)(unsafe.Pointer(&in.Controllers))
 if err := pkgapisconfigv1alpha1.Convert_v1alpha1_DebuggingConfiguration_To_config_DebuggingConfiguration(&in.Debugging, &out.Debugging, s); err != nil {
  return err
 }
 return nil
}
func autoConvert_config_GenericControllerManagerConfiguration_To_v1alpha1_GenericControllerManagerConfiguration(in *config.GenericControllerManagerConfiguration, out *v1alpha1.GenericControllerManagerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Port = in.Port
 out.Address = in.Address
 out.MinResyncPeriod = in.MinResyncPeriod
 if err := apisconfigv1alpha1.Convert_config_ClientConnectionConfiguration_To_v1alpha1_ClientConnectionConfiguration(&in.ClientConnection, &out.ClientConnection, s); err != nil {
  return err
 }
 out.ControllerStartInterval = in.ControllerStartInterval
 if err := pkgapisconfigv1alpha1.Convert_config_LeaderElectionConfiguration_To_v1alpha1_LeaderElectionConfiguration(&in.LeaderElection, &out.LeaderElection, s); err != nil {
  return err
 }
 out.Controllers = *(*[]string)(unsafe.Pointer(&in.Controllers))
 if err := pkgapisconfigv1alpha1.Convert_config_DebuggingConfiguration_To_v1alpha1_DebuggingConfiguration(&in.Debugging, &out.Debugging, s); err != nil {
  return err
 }
 return nil
}
func autoConvert_v1alpha1_GroupResource_To_config_GroupResource(in *v1alpha1.GroupResource, out *config.GroupResource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Group = in.Group
 out.Resource = in.Resource
 return nil
}
func Convert_v1alpha1_GroupResource_To_config_GroupResource(in *v1alpha1.GroupResource, out *config.GroupResource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1alpha1_GroupResource_To_config_GroupResource(in, out, s)
}
func autoConvert_config_GroupResource_To_v1alpha1_GroupResource(in *config.GroupResource, out *v1alpha1.GroupResource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Group = in.Group
 out.Resource = in.Resource
 return nil
}
func Convert_config_GroupResource_To_v1alpha1_GroupResource(in *config.GroupResource, out *v1alpha1.GroupResource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_config_GroupResource_To_v1alpha1_GroupResource(in, out, s)
}
func autoConvert_v1alpha1_HPAControllerConfiguration_To_config_HPAControllerConfiguration(in *v1alpha1.HPAControllerConfiguration, out *config.HPAControllerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.HorizontalPodAutoscalerSyncPeriod = in.HorizontalPodAutoscalerSyncPeriod
 out.HorizontalPodAutoscalerUpscaleForbiddenWindow = in.HorizontalPodAutoscalerUpscaleForbiddenWindow
 out.HorizontalPodAutoscalerDownscaleStabilizationWindow = in.HorizontalPodAutoscalerDownscaleStabilizationWindow
 out.HorizontalPodAutoscalerDownscaleForbiddenWindow = in.HorizontalPodAutoscalerDownscaleForbiddenWindow
 out.HorizontalPodAutoscalerTolerance = in.HorizontalPodAutoscalerTolerance
 if err := v1.Convert_Pointer_bool_To_bool(&in.HorizontalPodAutoscalerUseRESTClients, &out.HorizontalPodAutoscalerUseRESTClients, s); err != nil {
  return err
 }
 out.HorizontalPodAutoscalerCPUInitializationPeriod = in.HorizontalPodAutoscalerCPUInitializationPeriod
 out.HorizontalPodAutoscalerInitialReadinessDelay = in.HorizontalPodAutoscalerInitialReadinessDelay
 return nil
}
func Convert_v1alpha1_HPAControllerConfiguration_To_config_HPAControllerConfiguration(in *v1alpha1.HPAControllerConfiguration, out *config.HPAControllerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1alpha1_HPAControllerConfiguration_To_config_HPAControllerConfiguration(in, out, s)
}
func autoConvert_config_HPAControllerConfiguration_To_v1alpha1_HPAControllerConfiguration(in *config.HPAControllerConfiguration, out *v1alpha1.HPAControllerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.HorizontalPodAutoscalerSyncPeriod = in.HorizontalPodAutoscalerSyncPeriod
 out.HorizontalPodAutoscalerUpscaleForbiddenWindow = in.HorizontalPodAutoscalerUpscaleForbiddenWindow
 out.HorizontalPodAutoscalerDownscaleForbiddenWindow = in.HorizontalPodAutoscalerDownscaleForbiddenWindow
 out.HorizontalPodAutoscalerDownscaleStabilizationWindow = in.HorizontalPodAutoscalerDownscaleStabilizationWindow
 out.HorizontalPodAutoscalerTolerance = in.HorizontalPodAutoscalerTolerance
 if err := v1.Convert_bool_To_Pointer_bool(&in.HorizontalPodAutoscalerUseRESTClients, &out.HorizontalPodAutoscalerUseRESTClients, s); err != nil {
  return err
 }
 out.HorizontalPodAutoscalerCPUInitializationPeriod = in.HorizontalPodAutoscalerCPUInitializationPeriod
 out.HorizontalPodAutoscalerInitialReadinessDelay = in.HorizontalPodAutoscalerInitialReadinessDelay
 return nil
}
func Convert_config_HPAControllerConfiguration_To_v1alpha1_HPAControllerConfiguration(in *config.HPAControllerConfiguration, out *v1alpha1.HPAControllerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_config_HPAControllerConfiguration_To_v1alpha1_HPAControllerConfiguration(in, out, s)
}
func autoConvert_v1alpha1_JobControllerConfiguration_To_config_JobControllerConfiguration(in *v1alpha1.JobControllerConfiguration, out *config.JobControllerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ConcurrentJobSyncs = in.ConcurrentJobSyncs
 return nil
}
func Convert_v1alpha1_JobControllerConfiguration_To_config_JobControllerConfiguration(in *v1alpha1.JobControllerConfiguration, out *config.JobControllerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1alpha1_JobControllerConfiguration_To_config_JobControllerConfiguration(in, out, s)
}
func autoConvert_config_JobControllerConfiguration_To_v1alpha1_JobControllerConfiguration(in *config.JobControllerConfiguration, out *v1alpha1.JobControllerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ConcurrentJobSyncs = in.ConcurrentJobSyncs
 return nil
}
func Convert_config_JobControllerConfiguration_To_v1alpha1_JobControllerConfiguration(in *config.JobControllerConfiguration, out *v1alpha1.JobControllerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_config_JobControllerConfiguration_To_v1alpha1_JobControllerConfiguration(in, out, s)
}
func autoConvert_v1alpha1_KubeCloudSharedConfiguration_To_config_KubeCloudSharedConfiguration(in *v1alpha1.KubeCloudSharedConfiguration, out *config.KubeCloudSharedConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := Convert_v1alpha1_CloudProviderConfiguration_To_config_CloudProviderConfiguration(&in.CloudProvider, &out.CloudProvider, s); err != nil {
  return err
 }
 out.ExternalCloudVolumePlugin = in.ExternalCloudVolumePlugin
 out.UseServiceAccountCredentials = in.UseServiceAccountCredentials
 out.AllowUntaggedCloud = in.AllowUntaggedCloud
 out.RouteReconciliationPeriod = in.RouteReconciliationPeriod
 out.NodeMonitorPeriod = in.NodeMonitorPeriod
 out.ClusterName = in.ClusterName
 out.ClusterCIDR = in.ClusterCIDR
 out.AllocateNodeCIDRs = in.AllocateNodeCIDRs
 out.CIDRAllocatorType = in.CIDRAllocatorType
 if err := v1.Convert_Pointer_bool_To_bool(&in.ConfigureCloudRoutes, &out.ConfigureCloudRoutes, s); err != nil {
  return err
 }
 out.NodeSyncPeriod = in.NodeSyncPeriod
 return nil
}
func autoConvert_config_KubeCloudSharedConfiguration_To_v1alpha1_KubeCloudSharedConfiguration(in *config.KubeCloudSharedConfiguration, out *v1alpha1.KubeCloudSharedConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := Convert_config_CloudProviderConfiguration_To_v1alpha1_CloudProviderConfiguration(&in.CloudProvider, &out.CloudProvider, s); err != nil {
  return err
 }
 out.ExternalCloudVolumePlugin = in.ExternalCloudVolumePlugin
 out.UseServiceAccountCredentials = in.UseServiceAccountCredentials
 out.AllowUntaggedCloud = in.AllowUntaggedCloud
 out.RouteReconciliationPeriod = in.RouteReconciliationPeriod
 out.NodeMonitorPeriod = in.NodeMonitorPeriod
 out.ClusterName = in.ClusterName
 out.ClusterCIDR = in.ClusterCIDR
 out.AllocateNodeCIDRs = in.AllocateNodeCIDRs
 out.CIDRAllocatorType = in.CIDRAllocatorType
 if err := v1.Convert_bool_To_Pointer_bool(&in.ConfigureCloudRoutes, &out.ConfigureCloudRoutes, s); err != nil {
  return err
 }
 out.NodeSyncPeriod = in.NodeSyncPeriod
 return nil
}
func autoConvert_v1alpha1_KubeControllerManagerConfiguration_To_config_KubeControllerManagerConfiguration(in *v1alpha1.KubeControllerManagerConfiguration, out *config.KubeControllerManagerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := Convert_v1alpha1_GenericControllerManagerConfiguration_To_config_GenericControllerManagerConfiguration(&in.Generic, &out.Generic, s); err != nil {
  return err
 }
 if err := Convert_v1alpha1_KubeCloudSharedConfiguration_To_config_KubeCloudSharedConfiguration(&in.KubeCloudShared, &out.KubeCloudShared, s); err != nil {
  return err
 }
 if err := Convert_v1alpha1_AttachDetachControllerConfiguration_To_config_AttachDetachControllerConfiguration(&in.AttachDetachController, &out.AttachDetachController, s); err != nil {
  return err
 }
 if err := Convert_v1alpha1_CSRSigningControllerConfiguration_To_config_CSRSigningControllerConfiguration(&in.CSRSigningController, &out.CSRSigningController, s); err != nil {
  return err
 }
 if err := Convert_v1alpha1_DaemonSetControllerConfiguration_To_config_DaemonSetControllerConfiguration(&in.DaemonSetController, &out.DaemonSetController, s); err != nil {
  return err
 }
 if err := Convert_v1alpha1_DeploymentControllerConfiguration_To_config_DeploymentControllerConfiguration(&in.DeploymentController, &out.DeploymentController, s); err != nil {
  return err
 }
 if err := Convert_v1alpha1_DeprecatedControllerConfiguration_To_config_DeprecatedControllerConfiguration(&in.DeprecatedController, &out.DeprecatedController, s); err != nil {
  return err
 }
 if err := Convert_v1alpha1_EndpointControllerConfiguration_To_config_EndpointControllerConfiguration(&in.EndpointController, &out.EndpointController, s); err != nil {
  return err
 }
 if err := Convert_v1alpha1_GarbageCollectorControllerConfiguration_To_config_GarbageCollectorControllerConfiguration(&in.GarbageCollectorController, &out.GarbageCollectorController, s); err != nil {
  return err
 }
 if err := Convert_v1alpha1_HPAControllerConfiguration_To_config_HPAControllerConfiguration(&in.HPAController, &out.HPAController, s); err != nil {
  return err
 }
 if err := Convert_v1alpha1_JobControllerConfiguration_To_config_JobControllerConfiguration(&in.JobController, &out.JobController, s); err != nil {
  return err
 }
 if err := Convert_v1alpha1_NamespaceControllerConfiguration_To_config_NamespaceControllerConfiguration(&in.NamespaceController, &out.NamespaceController, s); err != nil {
  return err
 }
 if err := Convert_v1alpha1_NodeIPAMControllerConfiguration_To_config_NodeIPAMControllerConfiguration(&in.NodeIPAMController, &out.NodeIPAMController, s); err != nil {
  return err
 }
 if err := Convert_v1alpha1_NodeLifecycleControllerConfiguration_To_config_NodeLifecycleControllerConfiguration(&in.NodeLifecycleController, &out.NodeLifecycleController, s); err != nil {
  return err
 }
 if err := Convert_v1alpha1_PersistentVolumeBinderControllerConfiguration_To_config_PersistentVolumeBinderControllerConfiguration(&in.PersistentVolumeBinderController, &out.PersistentVolumeBinderController, s); err != nil {
  return err
 }
 if err := Convert_v1alpha1_PodGCControllerConfiguration_To_config_PodGCControllerConfiguration(&in.PodGCController, &out.PodGCController, s); err != nil {
  return err
 }
 if err := Convert_v1alpha1_ReplicaSetControllerConfiguration_To_config_ReplicaSetControllerConfiguration(&in.ReplicaSetController, &out.ReplicaSetController, s); err != nil {
  return err
 }
 if err := Convert_v1alpha1_ReplicationControllerConfiguration_To_config_ReplicationControllerConfiguration(&in.ReplicationController, &out.ReplicationController, s); err != nil {
  return err
 }
 if err := Convert_v1alpha1_ResourceQuotaControllerConfiguration_To_config_ResourceQuotaControllerConfiguration(&in.ResourceQuotaController, &out.ResourceQuotaController, s); err != nil {
  return err
 }
 if err := Convert_v1alpha1_SAControllerConfiguration_To_config_SAControllerConfiguration(&in.SAController, &out.SAController, s); err != nil {
  return err
 }
 if err := Convert_v1alpha1_ServiceControllerConfiguration_To_config_ServiceControllerConfiguration(&in.ServiceController, &out.ServiceController, s); err != nil {
  return err
 }
 if err := Convert_v1alpha1_TTLAfterFinishedControllerConfiguration_To_config_TTLAfterFinishedControllerConfiguration(&in.TTLAfterFinishedController, &out.TTLAfterFinishedController, s); err != nil {
  return err
 }
 return nil
}
func Convert_v1alpha1_KubeControllerManagerConfiguration_To_config_KubeControllerManagerConfiguration(in *v1alpha1.KubeControllerManagerConfiguration, out *config.KubeControllerManagerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1alpha1_KubeControllerManagerConfiguration_To_config_KubeControllerManagerConfiguration(in, out, s)
}
func autoConvert_config_KubeControllerManagerConfiguration_To_v1alpha1_KubeControllerManagerConfiguration(in *config.KubeControllerManagerConfiguration, out *v1alpha1.KubeControllerManagerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := Convert_config_GenericControllerManagerConfiguration_To_v1alpha1_GenericControllerManagerConfiguration(&in.Generic, &out.Generic, s); err != nil {
  return err
 }
 if err := Convert_config_KubeCloudSharedConfiguration_To_v1alpha1_KubeCloudSharedConfiguration(&in.KubeCloudShared, &out.KubeCloudShared, s); err != nil {
  return err
 }
 if err := Convert_config_AttachDetachControllerConfiguration_To_v1alpha1_AttachDetachControllerConfiguration(&in.AttachDetachController, &out.AttachDetachController, s); err != nil {
  return err
 }
 if err := Convert_config_CSRSigningControllerConfiguration_To_v1alpha1_CSRSigningControllerConfiguration(&in.CSRSigningController, &out.CSRSigningController, s); err != nil {
  return err
 }
 if err := Convert_config_DaemonSetControllerConfiguration_To_v1alpha1_DaemonSetControllerConfiguration(&in.DaemonSetController, &out.DaemonSetController, s); err != nil {
  return err
 }
 if err := Convert_config_DeploymentControllerConfiguration_To_v1alpha1_DeploymentControllerConfiguration(&in.DeploymentController, &out.DeploymentController, s); err != nil {
  return err
 }
 if err := Convert_config_DeprecatedControllerConfiguration_To_v1alpha1_DeprecatedControllerConfiguration(&in.DeprecatedController, &out.DeprecatedController, s); err != nil {
  return err
 }
 if err := Convert_config_EndpointControllerConfiguration_To_v1alpha1_EndpointControllerConfiguration(&in.EndpointController, &out.EndpointController, s); err != nil {
  return err
 }
 if err := Convert_config_GarbageCollectorControllerConfiguration_To_v1alpha1_GarbageCollectorControllerConfiguration(&in.GarbageCollectorController, &out.GarbageCollectorController, s); err != nil {
  return err
 }
 if err := Convert_config_HPAControllerConfiguration_To_v1alpha1_HPAControllerConfiguration(&in.HPAController, &out.HPAController, s); err != nil {
  return err
 }
 if err := Convert_config_JobControllerConfiguration_To_v1alpha1_JobControllerConfiguration(&in.JobController, &out.JobController, s); err != nil {
  return err
 }
 if err := Convert_config_NamespaceControllerConfiguration_To_v1alpha1_NamespaceControllerConfiguration(&in.NamespaceController, &out.NamespaceController, s); err != nil {
  return err
 }
 if err := Convert_config_NodeIPAMControllerConfiguration_To_v1alpha1_NodeIPAMControllerConfiguration(&in.NodeIPAMController, &out.NodeIPAMController, s); err != nil {
  return err
 }
 if err := Convert_config_NodeLifecycleControllerConfiguration_To_v1alpha1_NodeLifecycleControllerConfiguration(&in.NodeLifecycleController, &out.NodeLifecycleController, s); err != nil {
  return err
 }
 if err := Convert_config_PersistentVolumeBinderControllerConfiguration_To_v1alpha1_PersistentVolumeBinderControllerConfiguration(&in.PersistentVolumeBinderController, &out.PersistentVolumeBinderController, s); err != nil {
  return err
 }
 if err := Convert_config_PodGCControllerConfiguration_To_v1alpha1_PodGCControllerConfiguration(&in.PodGCController, &out.PodGCController, s); err != nil {
  return err
 }
 if err := Convert_config_ReplicaSetControllerConfiguration_To_v1alpha1_ReplicaSetControllerConfiguration(&in.ReplicaSetController, &out.ReplicaSetController, s); err != nil {
  return err
 }
 if err := Convert_config_ReplicationControllerConfiguration_To_v1alpha1_ReplicationControllerConfiguration(&in.ReplicationController, &out.ReplicationController, s); err != nil {
  return err
 }
 if err := Convert_config_ResourceQuotaControllerConfiguration_To_v1alpha1_ResourceQuotaControllerConfiguration(&in.ResourceQuotaController, &out.ResourceQuotaController, s); err != nil {
  return err
 }
 if err := Convert_config_SAControllerConfiguration_To_v1alpha1_SAControllerConfiguration(&in.SAController, &out.SAController, s); err != nil {
  return err
 }
 if err := Convert_config_ServiceControllerConfiguration_To_v1alpha1_ServiceControllerConfiguration(&in.ServiceController, &out.ServiceController, s); err != nil {
  return err
 }
 if err := Convert_config_TTLAfterFinishedControllerConfiguration_To_v1alpha1_TTLAfterFinishedControllerConfiguration(&in.TTLAfterFinishedController, &out.TTLAfterFinishedController, s); err != nil {
  return err
 }
 return nil
}
func Convert_config_KubeControllerManagerConfiguration_To_v1alpha1_KubeControllerManagerConfiguration(in *config.KubeControllerManagerConfiguration, out *v1alpha1.KubeControllerManagerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_config_KubeControllerManagerConfiguration_To_v1alpha1_KubeControllerManagerConfiguration(in, out, s)
}
func autoConvert_v1alpha1_NamespaceControllerConfiguration_To_config_NamespaceControllerConfiguration(in *v1alpha1.NamespaceControllerConfiguration, out *config.NamespaceControllerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.NamespaceSyncPeriod = in.NamespaceSyncPeriod
 out.ConcurrentNamespaceSyncs = in.ConcurrentNamespaceSyncs
 return nil
}
func Convert_v1alpha1_NamespaceControllerConfiguration_To_config_NamespaceControllerConfiguration(in *v1alpha1.NamespaceControllerConfiguration, out *config.NamespaceControllerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1alpha1_NamespaceControllerConfiguration_To_config_NamespaceControllerConfiguration(in, out, s)
}
func autoConvert_config_NamespaceControllerConfiguration_To_v1alpha1_NamespaceControllerConfiguration(in *config.NamespaceControllerConfiguration, out *v1alpha1.NamespaceControllerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.NamespaceSyncPeriod = in.NamespaceSyncPeriod
 out.ConcurrentNamespaceSyncs = in.ConcurrentNamespaceSyncs
 return nil
}
func Convert_config_NamespaceControllerConfiguration_To_v1alpha1_NamespaceControllerConfiguration(in *config.NamespaceControllerConfiguration, out *v1alpha1.NamespaceControllerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_config_NamespaceControllerConfiguration_To_v1alpha1_NamespaceControllerConfiguration(in, out, s)
}
func autoConvert_v1alpha1_NodeIPAMControllerConfiguration_To_config_NodeIPAMControllerConfiguration(in *v1alpha1.NodeIPAMControllerConfiguration, out *config.NodeIPAMControllerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ServiceCIDR = in.ServiceCIDR
 out.NodeCIDRMaskSize = in.NodeCIDRMaskSize
 return nil
}
func Convert_v1alpha1_NodeIPAMControllerConfiguration_To_config_NodeIPAMControllerConfiguration(in *v1alpha1.NodeIPAMControllerConfiguration, out *config.NodeIPAMControllerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1alpha1_NodeIPAMControllerConfiguration_To_config_NodeIPAMControllerConfiguration(in, out, s)
}
func autoConvert_config_NodeIPAMControllerConfiguration_To_v1alpha1_NodeIPAMControllerConfiguration(in *config.NodeIPAMControllerConfiguration, out *v1alpha1.NodeIPAMControllerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ServiceCIDR = in.ServiceCIDR
 out.NodeCIDRMaskSize = in.NodeCIDRMaskSize
 return nil
}
func Convert_config_NodeIPAMControllerConfiguration_To_v1alpha1_NodeIPAMControllerConfiguration(in *config.NodeIPAMControllerConfiguration, out *v1alpha1.NodeIPAMControllerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_config_NodeIPAMControllerConfiguration_To_v1alpha1_NodeIPAMControllerConfiguration(in, out, s)
}
func autoConvert_v1alpha1_NodeLifecycleControllerConfiguration_To_config_NodeLifecycleControllerConfiguration(in *v1alpha1.NodeLifecycleControllerConfiguration, out *config.NodeLifecycleControllerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := v1.Convert_Pointer_bool_To_bool(&in.EnableTaintManager, &out.EnableTaintManager, s); err != nil {
  return err
 }
 out.NodeEvictionRate = in.NodeEvictionRate
 out.SecondaryNodeEvictionRate = in.SecondaryNodeEvictionRate
 out.NodeStartupGracePeriod = in.NodeStartupGracePeriod
 out.NodeMonitorGracePeriod = in.NodeMonitorGracePeriod
 out.PodEvictionTimeout = in.PodEvictionTimeout
 out.LargeClusterSizeThreshold = in.LargeClusterSizeThreshold
 out.UnhealthyZoneThreshold = in.UnhealthyZoneThreshold
 return nil
}
func Convert_v1alpha1_NodeLifecycleControllerConfiguration_To_config_NodeLifecycleControllerConfiguration(in *v1alpha1.NodeLifecycleControllerConfiguration, out *config.NodeLifecycleControllerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1alpha1_NodeLifecycleControllerConfiguration_To_config_NodeLifecycleControllerConfiguration(in, out, s)
}
func autoConvert_config_NodeLifecycleControllerConfiguration_To_v1alpha1_NodeLifecycleControllerConfiguration(in *config.NodeLifecycleControllerConfiguration, out *v1alpha1.NodeLifecycleControllerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := v1.Convert_bool_To_Pointer_bool(&in.EnableTaintManager, &out.EnableTaintManager, s); err != nil {
  return err
 }
 out.NodeEvictionRate = in.NodeEvictionRate
 out.SecondaryNodeEvictionRate = in.SecondaryNodeEvictionRate
 out.NodeStartupGracePeriod = in.NodeStartupGracePeriod
 out.NodeMonitorGracePeriod = in.NodeMonitorGracePeriod
 out.PodEvictionTimeout = in.PodEvictionTimeout
 out.LargeClusterSizeThreshold = in.LargeClusterSizeThreshold
 out.UnhealthyZoneThreshold = in.UnhealthyZoneThreshold
 return nil
}
func Convert_config_NodeLifecycleControllerConfiguration_To_v1alpha1_NodeLifecycleControllerConfiguration(in *config.NodeLifecycleControllerConfiguration, out *v1alpha1.NodeLifecycleControllerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_config_NodeLifecycleControllerConfiguration_To_v1alpha1_NodeLifecycleControllerConfiguration(in, out, s)
}
func autoConvert_v1alpha1_PersistentVolumeBinderControllerConfiguration_To_config_PersistentVolumeBinderControllerConfiguration(in *v1alpha1.PersistentVolumeBinderControllerConfiguration, out *config.PersistentVolumeBinderControllerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.PVClaimBinderSyncPeriod = in.PVClaimBinderSyncPeriod
 if err := Convert_v1alpha1_VolumeConfiguration_To_config_VolumeConfiguration(&in.VolumeConfiguration, &out.VolumeConfiguration, s); err != nil {
  return err
 }
 return nil
}
func Convert_v1alpha1_PersistentVolumeBinderControllerConfiguration_To_config_PersistentVolumeBinderControllerConfiguration(in *v1alpha1.PersistentVolumeBinderControllerConfiguration, out *config.PersistentVolumeBinderControllerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1alpha1_PersistentVolumeBinderControllerConfiguration_To_config_PersistentVolumeBinderControllerConfiguration(in, out, s)
}
func autoConvert_config_PersistentVolumeBinderControllerConfiguration_To_v1alpha1_PersistentVolumeBinderControllerConfiguration(in *config.PersistentVolumeBinderControllerConfiguration, out *v1alpha1.PersistentVolumeBinderControllerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.PVClaimBinderSyncPeriod = in.PVClaimBinderSyncPeriod
 if err := Convert_config_VolumeConfiguration_To_v1alpha1_VolumeConfiguration(&in.VolumeConfiguration, &out.VolumeConfiguration, s); err != nil {
  return err
 }
 return nil
}
func Convert_config_PersistentVolumeBinderControllerConfiguration_To_v1alpha1_PersistentVolumeBinderControllerConfiguration(in *config.PersistentVolumeBinderControllerConfiguration, out *v1alpha1.PersistentVolumeBinderControllerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_config_PersistentVolumeBinderControllerConfiguration_To_v1alpha1_PersistentVolumeBinderControllerConfiguration(in, out, s)
}
func autoConvert_v1alpha1_PersistentVolumeRecyclerConfiguration_To_config_PersistentVolumeRecyclerConfiguration(in *v1alpha1.PersistentVolumeRecyclerConfiguration, out *config.PersistentVolumeRecyclerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.MaximumRetry = in.MaximumRetry
 out.MinimumTimeoutNFS = in.MinimumTimeoutNFS
 out.PodTemplateFilePathNFS = in.PodTemplateFilePathNFS
 out.IncrementTimeoutNFS = in.IncrementTimeoutNFS
 out.PodTemplateFilePathHostPath = in.PodTemplateFilePathHostPath
 out.MinimumTimeoutHostPath = in.MinimumTimeoutHostPath
 out.IncrementTimeoutHostPath = in.IncrementTimeoutHostPath
 return nil
}
func Convert_v1alpha1_PersistentVolumeRecyclerConfiguration_To_config_PersistentVolumeRecyclerConfiguration(in *v1alpha1.PersistentVolumeRecyclerConfiguration, out *config.PersistentVolumeRecyclerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1alpha1_PersistentVolumeRecyclerConfiguration_To_config_PersistentVolumeRecyclerConfiguration(in, out, s)
}
func autoConvert_config_PersistentVolumeRecyclerConfiguration_To_v1alpha1_PersistentVolumeRecyclerConfiguration(in *config.PersistentVolumeRecyclerConfiguration, out *v1alpha1.PersistentVolumeRecyclerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.MaximumRetry = in.MaximumRetry
 out.MinimumTimeoutNFS = in.MinimumTimeoutNFS
 out.PodTemplateFilePathNFS = in.PodTemplateFilePathNFS
 out.IncrementTimeoutNFS = in.IncrementTimeoutNFS
 out.PodTemplateFilePathHostPath = in.PodTemplateFilePathHostPath
 out.MinimumTimeoutHostPath = in.MinimumTimeoutHostPath
 out.IncrementTimeoutHostPath = in.IncrementTimeoutHostPath
 return nil
}
func Convert_config_PersistentVolumeRecyclerConfiguration_To_v1alpha1_PersistentVolumeRecyclerConfiguration(in *config.PersistentVolumeRecyclerConfiguration, out *v1alpha1.PersistentVolumeRecyclerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_config_PersistentVolumeRecyclerConfiguration_To_v1alpha1_PersistentVolumeRecyclerConfiguration(in, out, s)
}
func autoConvert_v1alpha1_PodGCControllerConfiguration_To_config_PodGCControllerConfiguration(in *v1alpha1.PodGCControllerConfiguration, out *config.PodGCControllerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.TerminatedPodGCThreshold = in.TerminatedPodGCThreshold
 return nil
}
func Convert_v1alpha1_PodGCControllerConfiguration_To_config_PodGCControllerConfiguration(in *v1alpha1.PodGCControllerConfiguration, out *config.PodGCControllerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1alpha1_PodGCControllerConfiguration_To_config_PodGCControllerConfiguration(in, out, s)
}
func autoConvert_config_PodGCControllerConfiguration_To_v1alpha1_PodGCControllerConfiguration(in *config.PodGCControllerConfiguration, out *v1alpha1.PodGCControllerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.TerminatedPodGCThreshold = in.TerminatedPodGCThreshold
 return nil
}
func Convert_config_PodGCControllerConfiguration_To_v1alpha1_PodGCControllerConfiguration(in *config.PodGCControllerConfiguration, out *v1alpha1.PodGCControllerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_config_PodGCControllerConfiguration_To_v1alpha1_PodGCControllerConfiguration(in, out, s)
}
func autoConvert_v1alpha1_ReplicaSetControllerConfiguration_To_config_ReplicaSetControllerConfiguration(in *v1alpha1.ReplicaSetControllerConfiguration, out *config.ReplicaSetControllerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ConcurrentRSSyncs = in.ConcurrentRSSyncs
 return nil
}
func Convert_v1alpha1_ReplicaSetControllerConfiguration_To_config_ReplicaSetControllerConfiguration(in *v1alpha1.ReplicaSetControllerConfiguration, out *config.ReplicaSetControllerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1alpha1_ReplicaSetControllerConfiguration_To_config_ReplicaSetControllerConfiguration(in, out, s)
}
func autoConvert_config_ReplicaSetControllerConfiguration_To_v1alpha1_ReplicaSetControllerConfiguration(in *config.ReplicaSetControllerConfiguration, out *v1alpha1.ReplicaSetControllerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ConcurrentRSSyncs = in.ConcurrentRSSyncs
 return nil
}
func Convert_config_ReplicaSetControllerConfiguration_To_v1alpha1_ReplicaSetControllerConfiguration(in *config.ReplicaSetControllerConfiguration, out *v1alpha1.ReplicaSetControllerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_config_ReplicaSetControllerConfiguration_To_v1alpha1_ReplicaSetControllerConfiguration(in, out, s)
}
func autoConvert_v1alpha1_ReplicationControllerConfiguration_To_config_ReplicationControllerConfiguration(in *v1alpha1.ReplicationControllerConfiguration, out *config.ReplicationControllerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ConcurrentRCSyncs = in.ConcurrentRCSyncs
 return nil
}
func Convert_v1alpha1_ReplicationControllerConfiguration_To_config_ReplicationControllerConfiguration(in *v1alpha1.ReplicationControllerConfiguration, out *config.ReplicationControllerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1alpha1_ReplicationControllerConfiguration_To_config_ReplicationControllerConfiguration(in, out, s)
}
func autoConvert_config_ReplicationControllerConfiguration_To_v1alpha1_ReplicationControllerConfiguration(in *config.ReplicationControllerConfiguration, out *v1alpha1.ReplicationControllerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ConcurrentRCSyncs = in.ConcurrentRCSyncs
 return nil
}
func Convert_config_ReplicationControllerConfiguration_To_v1alpha1_ReplicationControllerConfiguration(in *config.ReplicationControllerConfiguration, out *v1alpha1.ReplicationControllerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_config_ReplicationControllerConfiguration_To_v1alpha1_ReplicationControllerConfiguration(in, out, s)
}
func autoConvert_v1alpha1_ResourceQuotaControllerConfiguration_To_config_ResourceQuotaControllerConfiguration(in *v1alpha1.ResourceQuotaControllerConfiguration, out *config.ResourceQuotaControllerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ResourceQuotaSyncPeriod = in.ResourceQuotaSyncPeriod
 out.ConcurrentResourceQuotaSyncs = in.ConcurrentResourceQuotaSyncs
 return nil
}
func Convert_v1alpha1_ResourceQuotaControllerConfiguration_To_config_ResourceQuotaControllerConfiguration(in *v1alpha1.ResourceQuotaControllerConfiguration, out *config.ResourceQuotaControllerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1alpha1_ResourceQuotaControllerConfiguration_To_config_ResourceQuotaControllerConfiguration(in, out, s)
}
func autoConvert_config_ResourceQuotaControllerConfiguration_To_v1alpha1_ResourceQuotaControllerConfiguration(in *config.ResourceQuotaControllerConfiguration, out *v1alpha1.ResourceQuotaControllerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ResourceQuotaSyncPeriod = in.ResourceQuotaSyncPeriod
 out.ConcurrentResourceQuotaSyncs = in.ConcurrentResourceQuotaSyncs
 return nil
}
func Convert_config_ResourceQuotaControllerConfiguration_To_v1alpha1_ResourceQuotaControllerConfiguration(in *config.ResourceQuotaControllerConfiguration, out *v1alpha1.ResourceQuotaControllerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_config_ResourceQuotaControllerConfiguration_To_v1alpha1_ResourceQuotaControllerConfiguration(in, out, s)
}
func autoConvert_v1alpha1_SAControllerConfiguration_To_config_SAControllerConfiguration(in *v1alpha1.SAControllerConfiguration, out *config.SAControllerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ServiceAccountKeyFile = in.ServiceAccountKeyFile
 out.ConcurrentSATokenSyncs = in.ConcurrentSATokenSyncs
 out.RootCAFile = in.RootCAFile
 return nil
}
func Convert_v1alpha1_SAControllerConfiguration_To_config_SAControllerConfiguration(in *v1alpha1.SAControllerConfiguration, out *config.SAControllerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1alpha1_SAControllerConfiguration_To_config_SAControllerConfiguration(in, out, s)
}
func autoConvert_config_SAControllerConfiguration_To_v1alpha1_SAControllerConfiguration(in *config.SAControllerConfiguration, out *v1alpha1.SAControllerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ServiceAccountKeyFile = in.ServiceAccountKeyFile
 out.ConcurrentSATokenSyncs = in.ConcurrentSATokenSyncs
 out.RootCAFile = in.RootCAFile
 return nil
}
func Convert_config_SAControllerConfiguration_To_v1alpha1_SAControllerConfiguration(in *config.SAControllerConfiguration, out *v1alpha1.SAControllerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_config_SAControllerConfiguration_To_v1alpha1_SAControllerConfiguration(in, out, s)
}
func autoConvert_v1alpha1_ServiceControllerConfiguration_To_config_ServiceControllerConfiguration(in *v1alpha1.ServiceControllerConfiguration, out *config.ServiceControllerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ConcurrentServiceSyncs = in.ConcurrentServiceSyncs
 return nil
}
func autoConvert_config_ServiceControllerConfiguration_To_v1alpha1_ServiceControllerConfiguration(in *config.ServiceControllerConfiguration, out *v1alpha1.ServiceControllerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ConcurrentServiceSyncs = in.ConcurrentServiceSyncs
 return nil
}
func autoConvert_v1alpha1_TTLAfterFinishedControllerConfiguration_To_config_TTLAfterFinishedControllerConfiguration(in *v1alpha1.TTLAfterFinishedControllerConfiguration, out *config.TTLAfterFinishedControllerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ConcurrentTTLSyncs = in.ConcurrentTTLSyncs
 return nil
}
func Convert_v1alpha1_TTLAfterFinishedControllerConfiguration_To_config_TTLAfterFinishedControllerConfiguration(in *v1alpha1.TTLAfterFinishedControllerConfiguration, out *config.TTLAfterFinishedControllerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1alpha1_TTLAfterFinishedControllerConfiguration_To_config_TTLAfterFinishedControllerConfiguration(in, out, s)
}
func autoConvert_config_TTLAfterFinishedControllerConfiguration_To_v1alpha1_TTLAfterFinishedControllerConfiguration(in *config.TTLAfterFinishedControllerConfiguration, out *v1alpha1.TTLAfterFinishedControllerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ConcurrentTTLSyncs = in.ConcurrentTTLSyncs
 return nil
}
func Convert_config_TTLAfterFinishedControllerConfiguration_To_v1alpha1_TTLAfterFinishedControllerConfiguration(in *config.TTLAfterFinishedControllerConfiguration, out *v1alpha1.TTLAfterFinishedControllerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_config_TTLAfterFinishedControllerConfiguration_To_v1alpha1_TTLAfterFinishedControllerConfiguration(in, out, s)
}
func autoConvert_v1alpha1_VolumeConfiguration_To_config_VolumeConfiguration(in *v1alpha1.VolumeConfiguration, out *config.VolumeConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := v1.Convert_Pointer_bool_To_bool(&in.EnableHostPathProvisioning, &out.EnableHostPathProvisioning, s); err != nil {
  return err
 }
 if err := v1.Convert_Pointer_bool_To_bool(&in.EnableDynamicProvisioning, &out.EnableDynamicProvisioning, s); err != nil {
  return err
 }
 if err := Convert_v1alpha1_PersistentVolumeRecyclerConfiguration_To_config_PersistentVolumeRecyclerConfiguration(&in.PersistentVolumeRecyclerConfiguration, &out.PersistentVolumeRecyclerConfiguration, s); err != nil {
  return err
 }
 out.FlexVolumePluginDir = in.FlexVolumePluginDir
 return nil
}
func Convert_v1alpha1_VolumeConfiguration_To_config_VolumeConfiguration(in *v1alpha1.VolumeConfiguration, out *config.VolumeConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1alpha1_VolumeConfiguration_To_config_VolumeConfiguration(in, out, s)
}
func autoConvert_config_VolumeConfiguration_To_v1alpha1_VolumeConfiguration(in *config.VolumeConfiguration, out *v1alpha1.VolumeConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := v1.Convert_bool_To_Pointer_bool(&in.EnableHostPathProvisioning, &out.EnableHostPathProvisioning, s); err != nil {
  return err
 }
 if err := v1.Convert_bool_To_Pointer_bool(&in.EnableDynamicProvisioning, &out.EnableDynamicProvisioning, s); err != nil {
  return err
 }
 if err := Convert_config_PersistentVolumeRecyclerConfiguration_To_v1alpha1_PersistentVolumeRecyclerConfiguration(&in.PersistentVolumeRecyclerConfiguration, &out.PersistentVolumeRecyclerConfiguration, s); err != nil {
  return err
 }
 out.FlexVolumePluginDir = in.FlexVolumePluginDir
 return nil
}
func Convert_config_VolumeConfiguration_To_v1alpha1_VolumeConfiguration(in *config.VolumeConfiguration, out *v1alpha1.VolumeConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_config_VolumeConfiguration_To_v1alpha1_VolumeConfiguration(in, out, s)
}
