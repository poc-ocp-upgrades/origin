package v1

import (
 unsafe "unsafe"
 v1 "k8s.io/api/core/v1"
 resource "k8s.io/apimachinery/pkg/api/resource"
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 conversion "k8s.io/apimachinery/pkg/conversion"
 runtime "k8s.io/apimachinery/pkg/runtime"
 types "k8s.io/apimachinery/pkg/types"
 apps "k8s.io/kubernetes/pkg/apis/apps"
 core "k8s.io/kubernetes/pkg/apis/core"
)

func init() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 localSchemeBuilder.Register(RegisterConversions)
}
func RegisterConversions(s *runtime.Scheme) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := s.AddGeneratedConversionFunc((*v1.AWSElasticBlockStoreVolumeSource)(nil), (*core.AWSElasticBlockStoreVolumeSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_AWSElasticBlockStoreVolumeSource_To_core_AWSElasticBlockStoreVolumeSource(a.(*v1.AWSElasticBlockStoreVolumeSource), b.(*core.AWSElasticBlockStoreVolumeSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.AWSElasticBlockStoreVolumeSource)(nil), (*v1.AWSElasticBlockStoreVolumeSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_AWSElasticBlockStoreVolumeSource_To_v1_AWSElasticBlockStoreVolumeSource(a.(*core.AWSElasticBlockStoreVolumeSource), b.(*v1.AWSElasticBlockStoreVolumeSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.Affinity)(nil), (*core.Affinity)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_Affinity_To_core_Affinity(a.(*v1.Affinity), b.(*core.Affinity), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.Affinity)(nil), (*v1.Affinity)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_Affinity_To_v1_Affinity(a.(*core.Affinity), b.(*v1.Affinity), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.AttachedVolume)(nil), (*core.AttachedVolume)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_AttachedVolume_To_core_AttachedVolume(a.(*v1.AttachedVolume), b.(*core.AttachedVolume), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.AttachedVolume)(nil), (*v1.AttachedVolume)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_AttachedVolume_To_v1_AttachedVolume(a.(*core.AttachedVolume), b.(*v1.AttachedVolume), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.AvoidPods)(nil), (*core.AvoidPods)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_AvoidPods_To_core_AvoidPods(a.(*v1.AvoidPods), b.(*core.AvoidPods), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.AvoidPods)(nil), (*v1.AvoidPods)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_AvoidPods_To_v1_AvoidPods(a.(*core.AvoidPods), b.(*v1.AvoidPods), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.AzureDiskVolumeSource)(nil), (*core.AzureDiskVolumeSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_AzureDiskVolumeSource_To_core_AzureDiskVolumeSource(a.(*v1.AzureDiskVolumeSource), b.(*core.AzureDiskVolumeSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.AzureDiskVolumeSource)(nil), (*v1.AzureDiskVolumeSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_AzureDiskVolumeSource_To_v1_AzureDiskVolumeSource(a.(*core.AzureDiskVolumeSource), b.(*v1.AzureDiskVolumeSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.AzureFilePersistentVolumeSource)(nil), (*core.AzureFilePersistentVolumeSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_AzureFilePersistentVolumeSource_To_core_AzureFilePersistentVolumeSource(a.(*v1.AzureFilePersistentVolumeSource), b.(*core.AzureFilePersistentVolumeSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.AzureFilePersistentVolumeSource)(nil), (*v1.AzureFilePersistentVolumeSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_AzureFilePersistentVolumeSource_To_v1_AzureFilePersistentVolumeSource(a.(*core.AzureFilePersistentVolumeSource), b.(*v1.AzureFilePersistentVolumeSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.AzureFileVolumeSource)(nil), (*core.AzureFileVolumeSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_AzureFileVolumeSource_To_core_AzureFileVolumeSource(a.(*v1.AzureFileVolumeSource), b.(*core.AzureFileVolumeSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.AzureFileVolumeSource)(nil), (*v1.AzureFileVolumeSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_AzureFileVolumeSource_To_v1_AzureFileVolumeSource(a.(*core.AzureFileVolumeSource), b.(*v1.AzureFileVolumeSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.Binding)(nil), (*core.Binding)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_Binding_To_core_Binding(a.(*v1.Binding), b.(*core.Binding), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.Binding)(nil), (*v1.Binding)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_Binding_To_v1_Binding(a.(*core.Binding), b.(*v1.Binding), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.CSIPersistentVolumeSource)(nil), (*core.CSIPersistentVolumeSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_CSIPersistentVolumeSource_To_core_CSIPersistentVolumeSource(a.(*v1.CSIPersistentVolumeSource), b.(*core.CSIPersistentVolumeSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.CSIPersistentVolumeSource)(nil), (*v1.CSIPersistentVolumeSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_CSIPersistentVolumeSource_To_v1_CSIPersistentVolumeSource(a.(*core.CSIPersistentVolumeSource), b.(*v1.CSIPersistentVolumeSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.Capabilities)(nil), (*core.Capabilities)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_Capabilities_To_core_Capabilities(a.(*v1.Capabilities), b.(*core.Capabilities), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.Capabilities)(nil), (*v1.Capabilities)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_Capabilities_To_v1_Capabilities(a.(*core.Capabilities), b.(*v1.Capabilities), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.CephFSPersistentVolumeSource)(nil), (*core.CephFSPersistentVolumeSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_CephFSPersistentVolumeSource_To_core_CephFSPersistentVolumeSource(a.(*v1.CephFSPersistentVolumeSource), b.(*core.CephFSPersistentVolumeSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.CephFSPersistentVolumeSource)(nil), (*v1.CephFSPersistentVolumeSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_CephFSPersistentVolumeSource_To_v1_CephFSPersistentVolumeSource(a.(*core.CephFSPersistentVolumeSource), b.(*v1.CephFSPersistentVolumeSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.CephFSVolumeSource)(nil), (*core.CephFSVolumeSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_CephFSVolumeSource_To_core_CephFSVolumeSource(a.(*v1.CephFSVolumeSource), b.(*core.CephFSVolumeSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.CephFSVolumeSource)(nil), (*v1.CephFSVolumeSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_CephFSVolumeSource_To_v1_CephFSVolumeSource(a.(*core.CephFSVolumeSource), b.(*v1.CephFSVolumeSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.CinderPersistentVolumeSource)(nil), (*core.CinderPersistentVolumeSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_CinderPersistentVolumeSource_To_core_CinderPersistentVolumeSource(a.(*v1.CinderPersistentVolumeSource), b.(*core.CinderPersistentVolumeSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.CinderPersistentVolumeSource)(nil), (*v1.CinderPersistentVolumeSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_CinderPersistentVolumeSource_To_v1_CinderPersistentVolumeSource(a.(*core.CinderPersistentVolumeSource), b.(*v1.CinderPersistentVolumeSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.CinderVolumeSource)(nil), (*core.CinderVolumeSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_CinderVolumeSource_To_core_CinderVolumeSource(a.(*v1.CinderVolumeSource), b.(*core.CinderVolumeSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.CinderVolumeSource)(nil), (*v1.CinderVolumeSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_CinderVolumeSource_To_v1_CinderVolumeSource(a.(*core.CinderVolumeSource), b.(*v1.CinderVolumeSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.ClientIPConfig)(nil), (*core.ClientIPConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_ClientIPConfig_To_core_ClientIPConfig(a.(*v1.ClientIPConfig), b.(*core.ClientIPConfig), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.ClientIPConfig)(nil), (*v1.ClientIPConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_ClientIPConfig_To_v1_ClientIPConfig(a.(*core.ClientIPConfig), b.(*v1.ClientIPConfig), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.ComponentCondition)(nil), (*core.ComponentCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_ComponentCondition_To_core_ComponentCondition(a.(*v1.ComponentCondition), b.(*core.ComponentCondition), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.ComponentCondition)(nil), (*v1.ComponentCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_ComponentCondition_To_v1_ComponentCondition(a.(*core.ComponentCondition), b.(*v1.ComponentCondition), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.ComponentStatus)(nil), (*core.ComponentStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_ComponentStatus_To_core_ComponentStatus(a.(*v1.ComponentStatus), b.(*core.ComponentStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.ComponentStatus)(nil), (*v1.ComponentStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_ComponentStatus_To_v1_ComponentStatus(a.(*core.ComponentStatus), b.(*v1.ComponentStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.ComponentStatusList)(nil), (*core.ComponentStatusList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_ComponentStatusList_To_core_ComponentStatusList(a.(*v1.ComponentStatusList), b.(*core.ComponentStatusList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.ComponentStatusList)(nil), (*v1.ComponentStatusList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_ComponentStatusList_To_v1_ComponentStatusList(a.(*core.ComponentStatusList), b.(*v1.ComponentStatusList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.ConfigMap)(nil), (*core.ConfigMap)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_ConfigMap_To_core_ConfigMap(a.(*v1.ConfigMap), b.(*core.ConfigMap), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.ConfigMap)(nil), (*v1.ConfigMap)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_ConfigMap_To_v1_ConfigMap(a.(*core.ConfigMap), b.(*v1.ConfigMap), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.ConfigMapEnvSource)(nil), (*core.ConfigMapEnvSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_ConfigMapEnvSource_To_core_ConfigMapEnvSource(a.(*v1.ConfigMapEnvSource), b.(*core.ConfigMapEnvSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.ConfigMapEnvSource)(nil), (*v1.ConfigMapEnvSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_ConfigMapEnvSource_To_v1_ConfigMapEnvSource(a.(*core.ConfigMapEnvSource), b.(*v1.ConfigMapEnvSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.ConfigMapKeySelector)(nil), (*core.ConfigMapKeySelector)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_ConfigMapKeySelector_To_core_ConfigMapKeySelector(a.(*v1.ConfigMapKeySelector), b.(*core.ConfigMapKeySelector), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.ConfigMapKeySelector)(nil), (*v1.ConfigMapKeySelector)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_ConfigMapKeySelector_To_v1_ConfigMapKeySelector(a.(*core.ConfigMapKeySelector), b.(*v1.ConfigMapKeySelector), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.ConfigMapList)(nil), (*core.ConfigMapList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_ConfigMapList_To_core_ConfigMapList(a.(*v1.ConfigMapList), b.(*core.ConfigMapList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.ConfigMapList)(nil), (*v1.ConfigMapList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_ConfigMapList_To_v1_ConfigMapList(a.(*core.ConfigMapList), b.(*v1.ConfigMapList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.ConfigMapNodeConfigSource)(nil), (*core.ConfigMapNodeConfigSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_ConfigMapNodeConfigSource_To_core_ConfigMapNodeConfigSource(a.(*v1.ConfigMapNodeConfigSource), b.(*core.ConfigMapNodeConfigSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.ConfigMapNodeConfigSource)(nil), (*v1.ConfigMapNodeConfigSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_ConfigMapNodeConfigSource_To_v1_ConfigMapNodeConfigSource(a.(*core.ConfigMapNodeConfigSource), b.(*v1.ConfigMapNodeConfigSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.ConfigMapProjection)(nil), (*core.ConfigMapProjection)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_ConfigMapProjection_To_core_ConfigMapProjection(a.(*v1.ConfigMapProjection), b.(*core.ConfigMapProjection), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.ConfigMapProjection)(nil), (*v1.ConfigMapProjection)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_ConfigMapProjection_To_v1_ConfigMapProjection(a.(*core.ConfigMapProjection), b.(*v1.ConfigMapProjection), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.ConfigMapVolumeSource)(nil), (*core.ConfigMapVolumeSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_ConfigMapVolumeSource_To_core_ConfigMapVolumeSource(a.(*v1.ConfigMapVolumeSource), b.(*core.ConfigMapVolumeSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.ConfigMapVolumeSource)(nil), (*v1.ConfigMapVolumeSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_ConfigMapVolumeSource_To_v1_ConfigMapVolumeSource(a.(*core.ConfigMapVolumeSource), b.(*v1.ConfigMapVolumeSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.Container)(nil), (*core.Container)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_Container_To_core_Container(a.(*v1.Container), b.(*core.Container), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.Container)(nil), (*v1.Container)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_Container_To_v1_Container(a.(*core.Container), b.(*v1.Container), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.ContainerImage)(nil), (*core.ContainerImage)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_ContainerImage_To_core_ContainerImage(a.(*v1.ContainerImage), b.(*core.ContainerImage), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.ContainerImage)(nil), (*v1.ContainerImage)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_ContainerImage_To_v1_ContainerImage(a.(*core.ContainerImage), b.(*v1.ContainerImage), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.ContainerPort)(nil), (*core.ContainerPort)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_ContainerPort_To_core_ContainerPort(a.(*v1.ContainerPort), b.(*core.ContainerPort), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.ContainerPort)(nil), (*v1.ContainerPort)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_ContainerPort_To_v1_ContainerPort(a.(*core.ContainerPort), b.(*v1.ContainerPort), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.ContainerState)(nil), (*core.ContainerState)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_ContainerState_To_core_ContainerState(a.(*v1.ContainerState), b.(*core.ContainerState), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.ContainerState)(nil), (*v1.ContainerState)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_ContainerState_To_v1_ContainerState(a.(*core.ContainerState), b.(*v1.ContainerState), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.ContainerStateRunning)(nil), (*core.ContainerStateRunning)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_ContainerStateRunning_To_core_ContainerStateRunning(a.(*v1.ContainerStateRunning), b.(*core.ContainerStateRunning), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.ContainerStateRunning)(nil), (*v1.ContainerStateRunning)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_ContainerStateRunning_To_v1_ContainerStateRunning(a.(*core.ContainerStateRunning), b.(*v1.ContainerStateRunning), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.ContainerStateTerminated)(nil), (*core.ContainerStateTerminated)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_ContainerStateTerminated_To_core_ContainerStateTerminated(a.(*v1.ContainerStateTerminated), b.(*core.ContainerStateTerminated), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.ContainerStateTerminated)(nil), (*v1.ContainerStateTerminated)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_ContainerStateTerminated_To_v1_ContainerStateTerminated(a.(*core.ContainerStateTerminated), b.(*v1.ContainerStateTerminated), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.ContainerStateWaiting)(nil), (*core.ContainerStateWaiting)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_ContainerStateWaiting_To_core_ContainerStateWaiting(a.(*v1.ContainerStateWaiting), b.(*core.ContainerStateWaiting), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.ContainerStateWaiting)(nil), (*v1.ContainerStateWaiting)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_ContainerStateWaiting_To_v1_ContainerStateWaiting(a.(*core.ContainerStateWaiting), b.(*v1.ContainerStateWaiting), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.ContainerStatus)(nil), (*core.ContainerStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_ContainerStatus_To_core_ContainerStatus(a.(*v1.ContainerStatus), b.(*core.ContainerStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.ContainerStatus)(nil), (*v1.ContainerStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_ContainerStatus_To_v1_ContainerStatus(a.(*core.ContainerStatus), b.(*v1.ContainerStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.DaemonEndpoint)(nil), (*core.DaemonEndpoint)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_DaemonEndpoint_To_core_DaemonEndpoint(a.(*v1.DaemonEndpoint), b.(*core.DaemonEndpoint), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.DaemonEndpoint)(nil), (*v1.DaemonEndpoint)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_DaemonEndpoint_To_v1_DaemonEndpoint(a.(*core.DaemonEndpoint), b.(*v1.DaemonEndpoint), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.DownwardAPIProjection)(nil), (*core.DownwardAPIProjection)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_DownwardAPIProjection_To_core_DownwardAPIProjection(a.(*v1.DownwardAPIProjection), b.(*core.DownwardAPIProjection), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.DownwardAPIProjection)(nil), (*v1.DownwardAPIProjection)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_DownwardAPIProjection_To_v1_DownwardAPIProjection(a.(*core.DownwardAPIProjection), b.(*v1.DownwardAPIProjection), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.DownwardAPIVolumeFile)(nil), (*core.DownwardAPIVolumeFile)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_DownwardAPIVolumeFile_To_core_DownwardAPIVolumeFile(a.(*v1.DownwardAPIVolumeFile), b.(*core.DownwardAPIVolumeFile), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.DownwardAPIVolumeFile)(nil), (*v1.DownwardAPIVolumeFile)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_DownwardAPIVolumeFile_To_v1_DownwardAPIVolumeFile(a.(*core.DownwardAPIVolumeFile), b.(*v1.DownwardAPIVolumeFile), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.DownwardAPIVolumeSource)(nil), (*core.DownwardAPIVolumeSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_DownwardAPIVolumeSource_To_core_DownwardAPIVolumeSource(a.(*v1.DownwardAPIVolumeSource), b.(*core.DownwardAPIVolumeSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.DownwardAPIVolumeSource)(nil), (*v1.DownwardAPIVolumeSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_DownwardAPIVolumeSource_To_v1_DownwardAPIVolumeSource(a.(*core.DownwardAPIVolumeSource), b.(*v1.DownwardAPIVolumeSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.EmptyDirVolumeSource)(nil), (*core.EmptyDirVolumeSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_EmptyDirVolumeSource_To_core_EmptyDirVolumeSource(a.(*v1.EmptyDirVolumeSource), b.(*core.EmptyDirVolumeSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.EmptyDirVolumeSource)(nil), (*v1.EmptyDirVolumeSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_EmptyDirVolumeSource_To_v1_EmptyDirVolumeSource(a.(*core.EmptyDirVolumeSource), b.(*v1.EmptyDirVolumeSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.EndpointAddress)(nil), (*core.EndpointAddress)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_EndpointAddress_To_core_EndpointAddress(a.(*v1.EndpointAddress), b.(*core.EndpointAddress), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.EndpointAddress)(nil), (*v1.EndpointAddress)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_EndpointAddress_To_v1_EndpointAddress(a.(*core.EndpointAddress), b.(*v1.EndpointAddress), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.EndpointPort)(nil), (*core.EndpointPort)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_EndpointPort_To_core_EndpointPort(a.(*v1.EndpointPort), b.(*core.EndpointPort), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.EndpointPort)(nil), (*v1.EndpointPort)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_EndpointPort_To_v1_EndpointPort(a.(*core.EndpointPort), b.(*v1.EndpointPort), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.EndpointSubset)(nil), (*core.EndpointSubset)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_EndpointSubset_To_core_EndpointSubset(a.(*v1.EndpointSubset), b.(*core.EndpointSubset), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.EndpointSubset)(nil), (*v1.EndpointSubset)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_EndpointSubset_To_v1_EndpointSubset(a.(*core.EndpointSubset), b.(*v1.EndpointSubset), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.Endpoints)(nil), (*core.Endpoints)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_Endpoints_To_core_Endpoints(a.(*v1.Endpoints), b.(*core.Endpoints), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.Endpoints)(nil), (*v1.Endpoints)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_Endpoints_To_v1_Endpoints(a.(*core.Endpoints), b.(*v1.Endpoints), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.EndpointsList)(nil), (*core.EndpointsList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_EndpointsList_To_core_EndpointsList(a.(*v1.EndpointsList), b.(*core.EndpointsList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.EndpointsList)(nil), (*v1.EndpointsList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_EndpointsList_To_v1_EndpointsList(a.(*core.EndpointsList), b.(*v1.EndpointsList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.EnvFromSource)(nil), (*core.EnvFromSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_EnvFromSource_To_core_EnvFromSource(a.(*v1.EnvFromSource), b.(*core.EnvFromSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.EnvFromSource)(nil), (*v1.EnvFromSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_EnvFromSource_To_v1_EnvFromSource(a.(*core.EnvFromSource), b.(*v1.EnvFromSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.EnvVar)(nil), (*core.EnvVar)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_EnvVar_To_core_EnvVar(a.(*v1.EnvVar), b.(*core.EnvVar), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.EnvVar)(nil), (*v1.EnvVar)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_EnvVar_To_v1_EnvVar(a.(*core.EnvVar), b.(*v1.EnvVar), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.EnvVarSource)(nil), (*core.EnvVarSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_EnvVarSource_To_core_EnvVarSource(a.(*v1.EnvVarSource), b.(*core.EnvVarSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.EnvVarSource)(nil), (*v1.EnvVarSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_EnvVarSource_To_v1_EnvVarSource(a.(*core.EnvVarSource), b.(*v1.EnvVarSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.Event)(nil), (*core.Event)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_Event_To_core_Event(a.(*v1.Event), b.(*core.Event), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.Event)(nil), (*v1.Event)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_Event_To_v1_Event(a.(*core.Event), b.(*v1.Event), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.EventList)(nil), (*core.EventList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_EventList_To_core_EventList(a.(*v1.EventList), b.(*core.EventList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.EventList)(nil), (*v1.EventList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_EventList_To_v1_EventList(a.(*core.EventList), b.(*v1.EventList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.EventSeries)(nil), (*core.EventSeries)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_EventSeries_To_core_EventSeries(a.(*v1.EventSeries), b.(*core.EventSeries), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.EventSeries)(nil), (*v1.EventSeries)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_EventSeries_To_v1_EventSeries(a.(*core.EventSeries), b.(*v1.EventSeries), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.EventSource)(nil), (*core.EventSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_EventSource_To_core_EventSource(a.(*v1.EventSource), b.(*core.EventSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.EventSource)(nil), (*v1.EventSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_EventSource_To_v1_EventSource(a.(*core.EventSource), b.(*v1.EventSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.ExecAction)(nil), (*core.ExecAction)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_ExecAction_To_core_ExecAction(a.(*v1.ExecAction), b.(*core.ExecAction), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.ExecAction)(nil), (*v1.ExecAction)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_ExecAction_To_v1_ExecAction(a.(*core.ExecAction), b.(*v1.ExecAction), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.FCVolumeSource)(nil), (*core.FCVolumeSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_FCVolumeSource_To_core_FCVolumeSource(a.(*v1.FCVolumeSource), b.(*core.FCVolumeSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.FCVolumeSource)(nil), (*v1.FCVolumeSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_FCVolumeSource_To_v1_FCVolumeSource(a.(*core.FCVolumeSource), b.(*v1.FCVolumeSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.FlexPersistentVolumeSource)(nil), (*core.FlexPersistentVolumeSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_FlexPersistentVolumeSource_To_core_FlexPersistentVolumeSource(a.(*v1.FlexPersistentVolumeSource), b.(*core.FlexPersistentVolumeSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.FlexPersistentVolumeSource)(nil), (*v1.FlexPersistentVolumeSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_FlexPersistentVolumeSource_To_v1_FlexPersistentVolumeSource(a.(*core.FlexPersistentVolumeSource), b.(*v1.FlexPersistentVolumeSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.FlexVolumeSource)(nil), (*core.FlexVolumeSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_FlexVolumeSource_To_core_FlexVolumeSource(a.(*v1.FlexVolumeSource), b.(*core.FlexVolumeSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.FlexVolumeSource)(nil), (*v1.FlexVolumeSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_FlexVolumeSource_To_v1_FlexVolumeSource(a.(*core.FlexVolumeSource), b.(*v1.FlexVolumeSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.FlockerVolumeSource)(nil), (*core.FlockerVolumeSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_FlockerVolumeSource_To_core_FlockerVolumeSource(a.(*v1.FlockerVolumeSource), b.(*core.FlockerVolumeSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.FlockerVolumeSource)(nil), (*v1.FlockerVolumeSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_FlockerVolumeSource_To_v1_FlockerVolumeSource(a.(*core.FlockerVolumeSource), b.(*v1.FlockerVolumeSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.GCEPersistentDiskVolumeSource)(nil), (*core.GCEPersistentDiskVolumeSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_GCEPersistentDiskVolumeSource_To_core_GCEPersistentDiskVolumeSource(a.(*v1.GCEPersistentDiskVolumeSource), b.(*core.GCEPersistentDiskVolumeSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.GCEPersistentDiskVolumeSource)(nil), (*v1.GCEPersistentDiskVolumeSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_GCEPersistentDiskVolumeSource_To_v1_GCEPersistentDiskVolumeSource(a.(*core.GCEPersistentDiskVolumeSource), b.(*v1.GCEPersistentDiskVolumeSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.GitRepoVolumeSource)(nil), (*core.GitRepoVolumeSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_GitRepoVolumeSource_To_core_GitRepoVolumeSource(a.(*v1.GitRepoVolumeSource), b.(*core.GitRepoVolumeSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.GitRepoVolumeSource)(nil), (*v1.GitRepoVolumeSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_GitRepoVolumeSource_To_v1_GitRepoVolumeSource(a.(*core.GitRepoVolumeSource), b.(*v1.GitRepoVolumeSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.GlusterfsPersistentVolumeSource)(nil), (*core.GlusterfsPersistentVolumeSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_GlusterfsPersistentVolumeSource_To_core_GlusterfsPersistentVolumeSource(a.(*v1.GlusterfsPersistentVolumeSource), b.(*core.GlusterfsPersistentVolumeSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.GlusterfsPersistentVolumeSource)(nil), (*v1.GlusterfsPersistentVolumeSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_GlusterfsPersistentVolumeSource_To_v1_GlusterfsPersistentVolumeSource(a.(*core.GlusterfsPersistentVolumeSource), b.(*v1.GlusterfsPersistentVolumeSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.GlusterfsVolumeSource)(nil), (*core.GlusterfsVolumeSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_GlusterfsVolumeSource_To_core_GlusterfsVolumeSource(a.(*v1.GlusterfsVolumeSource), b.(*core.GlusterfsVolumeSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.GlusterfsVolumeSource)(nil), (*v1.GlusterfsVolumeSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_GlusterfsVolumeSource_To_v1_GlusterfsVolumeSource(a.(*core.GlusterfsVolumeSource), b.(*v1.GlusterfsVolumeSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.HTTPGetAction)(nil), (*core.HTTPGetAction)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_HTTPGetAction_To_core_HTTPGetAction(a.(*v1.HTTPGetAction), b.(*core.HTTPGetAction), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.HTTPGetAction)(nil), (*v1.HTTPGetAction)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_HTTPGetAction_To_v1_HTTPGetAction(a.(*core.HTTPGetAction), b.(*v1.HTTPGetAction), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.HTTPHeader)(nil), (*core.HTTPHeader)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_HTTPHeader_To_core_HTTPHeader(a.(*v1.HTTPHeader), b.(*core.HTTPHeader), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.HTTPHeader)(nil), (*v1.HTTPHeader)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_HTTPHeader_To_v1_HTTPHeader(a.(*core.HTTPHeader), b.(*v1.HTTPHeader), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.Handler)(nil), (*core.Handler)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_Handler_To_core_Handler(a.(*v1.Handler), b.(*core.Handler), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.Handler)(nil), (*v1.Handler)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_Handler_To_v1_Handler(a.(*core.Handler), b.(*v1.Handler), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.HostAlias)(nil), (*core.HostAlias)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_HostAlias_To_core_HostAlias(a.(*v1.HostAlias), b.(*core.HostAlias), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.HostAlias)(nil), (*v1.HostAlias)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_HostAlias_To_v1_HostAlias(a.(*core.HostAlias), b.(*v1.HostAlias), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.HostPathVolumeSource)(nil), (*core.HostPathVolumeSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_HostPathVolumeSource_To_core_HostPathVolumeSource(a.(*v1.HostPathVolumeSource), b.(*core.HostPathVolumeSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.HostPathVolumeSource)(nil), (*v1.HostPathVolumeSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_HostPathVolumeSource_To_v1_HostPathVolumeSource(a.(*core.HostPathVolumeSource), b.(*v1.HostPathVolumeSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.ISCSIPersistentVolumeSource)(nil), (*core.ISCSIPersistentVolumeSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_ISCSIPersistentVolumeSource_To_core_ISCSIPersistentVolumeSource(a.(*v1.ISCSIPersistentVolumeSource), b.(*core.ISCSIPersistentVolumeSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.ISCSIPersistentVolumeSource)(nil), (*v1.ISCSIPersistentVolumeSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_ISCSIPersistentVolumeSource_To_v1_ISCSIPersistentVolumeSource(a.(*core.ISCSIPersistentVolumeSource), b.(*v1.ISCSIPersistentVolumeSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.ISCSIVolumeSource)(nil), (*core.ISCSIVolumeSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_ISCSIVolumeSource_To_core_ISCSIVolumeSource(a.(*v1.ISCSIVolumeSource), b.(*core.ISCSIVolumeSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.ISCSIVolumeSource)(nil), (*v1.ISCSIVolumeSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_ISCSIVolumeSource_To_v1_ISCSIVolumeSource(a.(*core.ISCSIVolumeSource), b.(*v1.ISCSIVolumeSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.KeyToPath)(nil), (*core.KeyToPath)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_KeyToPath_To_core_KeyToPath(a.(*v1.KeyToPath), b.(*core.KeyToPath), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.KeyToPath)(nil), (*v1.KeyToPath)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_KeyToPath_To_v1_KeyToPath(a.(*core.KeyToPath), b.(*v1.KeyToPath), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.Lifecycle)(nil), (*core.Lifecycle)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_Lifecycle_To_core_Lifecycle(a.(*v1.Lifecycle), b.(*core.Lifecycle), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.Lifecycle)(nil), (*v1.Lifecycle)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_Lifecycle_To_v1_Lifecycle(a.(*core.Lifecycle), b.(*v1.Lifecycle), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.LimitRange)(nil), (*core.LimitRange)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_LimitRange_To_core_LimitRange(a.(*v1.LimitRange), b.(*core.LimitRange), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.LimitRange)(nil), (*v1.LimitRange)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_LimitRange_To_v1_LimitRange(a.(*core.LimitRange), b.(*v1.LimitRange), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.LimitRangeItem)(nil), (*core.LimitRangeItem)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_LimitRangeItem_To_core_LimitRangeItem(a.(*v1.LimitRangeItem), b.(*core.LimitRangeItem), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.LimitRangeItem)(nil), (*v1.LimitRangeItem)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_LimitRangeItem_To_v1_LimitRangeItem(a.(*core.LimitRangeItem), b.(*v1.LimitRangeItem), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.LimitRangeList)(nil), (*core.LimitRangeList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_LimitRangeList_To_core_LimitRangeList(a.(*v1.LimitRangeList), b.(*core.LimitRangeList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.LimitRangeList)(nil), (*v1.LimitRangeList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_LimitRangeList_To_v1_LimitRangeList(a.(*core.LimitRangeList), b.(*v1.LimitRangeList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.LimitRangeSpec)(nil), (*core.LimitRangeSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_LimitRangeSpec_To_core_LimitRangeSpec(a.(*v1.LimitRangeSpec), b.(*core.LimitRangeSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.LimitRangeSpec)(nil), (*v1.LimitRangeSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_LimitRangeSpec_To_v1_LimitRangeSpec(a.(*core.LimitRangeSpec), b.(*v1.LimitRangeSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.List)(nil), (*core.List)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_List_To_core_List(a.(*v1.List), b.(*core.List), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.List)(nil), (*v1.List)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_List_To_v1_List(a.(*core.List), b.(*v1.List), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.LoadBalancerIngress)(nil), (*core.LoadBalancerIngress)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_LoadBalancerIngress_To_core_LoadBalancerIngress(a.(*v1.LoadBalancerIngress), b.(*core.LoadBalancerIngress), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.LoadBalancerIngress)(nil), (*v1.LoadBalancerIngress)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_LoadBalancerIngress_To_v1_LoadBalancerIngress(a.(*core.LoadBalancerIngress), b.(*v1.LoadBalancerIngress), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.LoadBalancerStatus)(nil), (*core.LoadBalancerStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_LoadBalancerStatus_To_core_LoadBalancerStatus(a.(*v1.LoadBalancerStatus), b.(*core.LoadBalancerStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.LoadBalancerStatus)(nil), (*v1.LoadBalancerStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_LoadBalancerStatus_To_v1_LoadBalancerStatus(a.(*core.LoadBalancerStatus), b.(*v1.LoadBalancerStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.LocalObjectReference)(nil), (*core.LocalObjectReference)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_LocalObjectReference_To_core_LocalObjectReference(a.(*v1.LocalObjectReference), b.(*core.LocalObjectReference), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.LocalObjectReference)(nil), (*v1.LocalObjectReference)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_LocalObjectReference_To_v1_LocalObjectReference(a.(*core.LocalObjectReference), b.(*v1.LocalObjectReference), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.LocalVolumeSource)(nil), (*core.LocalVolumeSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_LocalVolumeSource_To_core_LocalVolumeSource(a.(*v1.LocalVolumeSource), b.(*core.LocalVolumeSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.LocalVolumeSource)(nil), (*v1.LocalVolumeSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_LocalVolumeSource_To_v1_LocalVolumeSource(a.(*core.LocalVolumeSource), b.(*v1.LocalVolumeSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.NFSVolumeSource)(nil), (*core.NFSVolumeSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_NFSVolumeSource_To_core_NFSVolumeSource(a.(*v1.NFSVolumeSource), b.(*core.NFSVolumeSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.NFSVolumeSource)(nil), (*v1.NFSVolumeSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_NFSVolumeSource_To_v1_NFSVolumeSource(a.(*core.NFSVolumeSource), b.(*v1.NFSVolumeSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.Namespace)(nil), (*core.Namespace)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_Namespace_To_core_Namespace(a.(*v1.Namespace), b.(*core.Namespace), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.Namespace)(nil), (*v1.Namespace)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_Namespace_To_v1_Namespace(a.(*core.Namespace), b.(*v1.Namespace), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.NamespaceList)(nil), (*core.NamespaceList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_NamespaceList_To_core_NamespaceList(a.(*v1.NamespaceList), b.(*core.NamespaceList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.NamespaceList)(nil), (*v1.NamespaceList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_NamespaceList_To_v1_NamespaceList(a.(*core.NamespaceList), b.(*v1.NamespaceList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.NamespaceSpec)(nil), (*core.NamespaceSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_NamespaceSpec_To_core_NamespaceSpec(a.(*v1.NamespaceSpec), b.(*core.NamespaceSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.NamespaceSpec)(nil), (*v1.NamespaceSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_NamespaceSpec_To_v1_NamespaceSpec(a.(*core.NamespaceSpec), b.(*v1.NamespaceSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.NamespaceStatus)(nil), (*core.NamespaceStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_NamespaceStatus_To_core_NamespaceStatus(a.(*v1.NamespaceStatus), b.(*core.NamespaceStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.NamespaceStatus)(nil), (*v1.NamespaceStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_NamespaceStatus_To_v1_NamespaceStatus(a.(*core.NamespaceStatus), b.(*v1.NamespaceStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.Node)(nil), (*core.Node)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_Node_To_core_Node(a.(*v1.Node), b.(*core.Node), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.Node)(nil), (*v1.Node)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_Node_To_v1_Node(a.(*core.Node), b.(*v1.Node), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.NodeAddress)(nil), (*core.NodeAddress)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_NodeAddress_To_core_NodeAddress(a.(*v1.NodeAddress), b.(*core.NodeAddress), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.NodeAddress)(nil), (*v1.NodeAddress)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_NodeAddress_To_v1_NodeAddress(a.(*core.NodeAddress), b.(*v1.NodeAddress), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.NodeAffinity)(nil), (*core.NodeAffinity)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_NodeAffinity_To_core_NodeAffinity(a.(*v1.NodeAffinity), b.(*core.NodeAffinity), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.NodeAffinity)(nil), (*v1.NodeAffinity)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_NodeAffinity_To_v1_NodeAffinity(a.(*core.NodeAffinity), b.(*v1.NodeAffinity), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.NodeCondition)(nil), (*core.NodeCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_NodeCondition_To_core_NodeCondition(a.(*v1.NodeCondition), b.(*core.NodeCondition), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.NodeCondition)(nil), (*v1.NodeCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_NodeCondition_To_v1_NodeCondition(a.(*core.NodeCondition), b.(*v1.NodeCondition), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.NodeConfigSource)(nil), (*core.NodeConfigSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_NodeConfigSource_To_core_NodeConfigSource(a.(*v1.NodeConfigSource), b.(*core.NodeConfigSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.NodeConfigSource)(nil), (*v1.NodeConfigSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_NodeConfigSource_To_v1_NodeConfigSource(a.(*core.NodeConfigSource), b.(*v1.NodeConfigSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.NodeConfigStatus)(nil), (*core.NodeConfigStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_NodeConfigStatus_To_core_NodeConfigStatus(a.(*v1.NodeConfigStatus), b.(*core.NodeConfigStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.NodeConfigStatus)(nil), (*v1.NodeConfigStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_NodeConfigStatus_To_v1_NodeConfigStatus(a.(*core.NodeConfigStatus), b.(*v1.NodeConfigStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.NodeDaemonEndpoints)(nil), (*core.NodeDaemonEndpoints)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_NodeDaemonEndpoints_To_core_NodeDaemonEndpoints(a.(*v1.NodeDaemonEndpoints), b.(*core.NodeDaemonEndpoints), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.NodeDaemonEndpoints)(nil), (*v1.NodeDaemonEndpoints)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_NodeDaemonEndpoints_To_v1_NodeDaemonEndpoints(a.(*core.NodeDaemonEndpoints), b.(*v1.NodeDaemonEndpoints), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.NodeList)(nil), (*core.NodeList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_NodeList_To_core_NodeList(a.(*v1.NodeList), b.(*core.NodeList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.NodeList)(nil), (*v1.NodeList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_NodeList_To_v1_NodeList(a.(*core.NodeList), b.(*v1.NodeList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.NodeProxyOptions)(nil), (*core.NodeProxyOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_NodeProxyOptions_To_core_NodeProxyOptions(a.(*v1.NodeProxyOptions), b.(*core.NodeProxyOptions), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.NodeProxyOptions)(nil), (*v1.NodeProxyOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_NodeProxyOptions_To_v1_NodeProxyOptions(a.(*core.NodeProxyOptions), b.(*v1.NodeProxyOptions), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.NodeResources)(nil), (*core.NodeResources)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_NodeResources_To_core_NodeResources(a.(*v1.NodeResources), b.(*core.NodeResources), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.NodeResources)(nil), (*v1.NodeResources)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_NodeResources_To_v1_NodeResources(a.(*core.NodeResources), b.(*v1.NodeResources), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.NodeSelector)(nil), (*core.NodeSelector)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_NodeSelector_To_core_NodeSelector(a.(*v1.NodeSelector), b.(*core.NodeSelector), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.NodeSelector)(nil), (*v1.NodeSelector)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_NodeSelector_To_v1_NodeSelector(a.(*core.NodeSelector), b.(*v1.NodeSelector), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.NodeSelectorRequirement)(nil), (*core.NodeSelectorRequirement)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_NodeSelectorRequirement_To_core_NodeSelectorRequirement(a.(*v1.NodeSelectorRequirement), b.(*core.NodeSelectorRequirement), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.NodeSelectorRequirement)(nil), (*v1.NodeSelectorRequirement)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_NodeSelectorRequirement_To_v1_NodeSelectorRequirement(a.(*core.NodeSelectorRequirement), b.(*v1.NodeSelectorRequirement), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.NodeSelectorTerm)(nil), (*core.NodeSelectorTerm)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_NodeSelectorTerm_To_core_NodeSelectorTerm(a.(*v1.NodeSelectorTerm), b.(*core.NodeSelectorTerm), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.NodeSelectorTerm)(nil), (*v1.NodeSelectorTerm)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_NodeSelectorTerm_To_v1_NodeSelectorTerm(a.(*core.NodeSelectorTerm), b.(*v1.NodeSelectorTerm), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.NodeSpec)(nil), (*core.NodeSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_NodeSpec_To_core_NodeSpec(a.(*v1.NodeSpec), b.(*core.NodeSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.NodeSpec)(nil), (*v1.NodeSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_NodeSpec_To_v1_NodeSpec(a.(*core.NodeSpec), b.(*v1.NodeSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.NodeStatus)(nil), (*core.NodeStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_NodeStatus_To_core_NodeStatus(a.(*v1.NodeStatus), b.(*core.NodeStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.NodeStatus)(nil), (*v1.NodeStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_NodeStatus_To_v1_NodeStatus(a.(*core.NodeStatus), b.(*v1.NodeStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.NodeSystemInfo)(nil), (*core.NodeSystemInfo)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_NodeSystemInfo_To_core_NodeSystemInfo(a.(*v1.NodeSystemInfo), b.(*core.NodeSystemInfo), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.NodeSystemInfo)(nil), (*v1.NodeSystemInfo)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_NodeSystemInfo_To_v1_NodeSystemInfo(a.(*core.NodeSystemInfo), b.(*v1.NodeSystemInfo), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.ObjectFieldSelector)(nil), (*core.ObjectFieldSelector)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_ObjectFieldSelector_To_core_ObjectFieldSelector(a.(*v1.ObjectFieldSelector), b.(*core.ObjectFieldSelector), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.ObjectFieldSelector)(nil), (*v1.ObjectFieldSelector)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_ObjectFieldSelector_To_v1_ObjectFieldSelector(a.(*core.ObjectFieldSelector), b.(*v1.ObjectFieldSelector), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.ObjectReference)(nil), (*core.ObjectReference)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_ObjectReference_To_core_ObjectReference(a.(*v1.ObjectReference), b.(*core.ObjectReference), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.ObjectReference)(nil), (*v1.ObjectReference)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_ObjectReference_To_v1_ObjectReference(a.(*core.ObjectReference), b.(*v1.ObjectReference), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.PersistentVolume)(nil), (*core.PersistentVolume)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_PersistentVolume_To_core_PersistentVolume(a.(*v1.PersistentVolume), b.(*core.PersistentVolume), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.PersistentVolume)(nil), (*v1.PersistentVolume)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_PersistentVolume_To_v1_PersistentVolume(a.(*core.PersistentVolume), b.(*v1.PersistentVolume), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.PersistentVolumeClaim)(nil), (*core.PersistentVolumeClaim)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_PersistentVolumeClaim_To_core_PersistentVolumeClaim(a.(*v1.PersistentVolumeClaim), b.(*core.PersistentVolumeClaim), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.PersistentVolumeClaim)(nil), (*v1.PersistentVolumeClaim)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_PersistentVolumeClaim_To_v1_PersistentVolumeClaim(a.(*core.PersistentVolumeClaim), b.(*v1.PersistentVolumeClaim), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.PersistentVolumeClaimCondition)(nil), (*core.PersistentVolumeClaimCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_PersistentVolumeClaimCondition_To_core_PersistentVolumeClaimCondition(a.(*v1.PersistentVolumeClaimCondition), b.(*core.PersistentVolumeClaimCondition), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.PersistentVolumeClaimCondition)(nil), (*v1.PersistentVolumeClaimCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_PersistentVolumeClaimCondition_To_v1_PersistentVolumeClaimCondition(a.(*core.PersistentVolumeClaimCondition), b.(*v1.PersistentVolumeClaimCondition), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.PersistentVolumeClaimList)(nil), (*core.PersistentVolumeClaimList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_PersistentVolumeClaimList_To_core_PersistentVolumeClaimList(a.(*v1.PersistentVolumeClaimList), b.(*core.PersistentVolumeClaimList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.PersistentVolumeClaimList)(nil), (*v1.PersistentVolumeClaimList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_PersistentVolumeClaimList_To_v1_PersistentVolumeClaimList(a.(*core.PersistentVolumeClaimList), b.(*v1.PersistentVolumeClaimList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.PersistentVolumeClaimSpec)(nil), (*core.PersistentVolumeClaimSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_PersistentVolumeClaimSpec_To_core_PersistentVolumeClaimSpec(a.(*v1.PersistentVolumeClaimSpec), b.(*core.PersistentVolumeClaimSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.PersistentVolumeClaimSpec)(nil), (*v1.PersistentVolumeClaimSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_PersistentVolumeClaimSpec_To_v1_PersistentVolumeClaimSpec(a.(*core.PersistentVolumeClaimSpec), b.(*v1.PersistentVolumeClaimSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.PersistentVolumeClaimStatus)(nil), (*core.PersistentVolumeClaimStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_PersistentVolumeClaimStatus_To_core_PersistentVolumeClaimStatus(a.(*v1.PersistentVolumeClaimStatus), b.(*core.PersistentVolumeClaimStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.PersistentVolumeClaimStatus)(nil), (*v1.PersistentVolumeClaimStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_PersistentVolumeClaimStatus_To_v1_PersistentVolumeClaimStatus(a.(*core.PersistentVolumeClaimStatus), b.(*v1.PersistentVolumeClaimStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.PersistentVolumeClaimVolumeSource)(nil), (*core.PersistentVolumeClaimVolumeSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_PersistentVolumeClaimVolumeSource_To_core_PersistentVolumeClaimVolumeSource(a.(*v1.PersistentVolumeClaimVolumeSource), b.(*core.PersistentVolumeClaimVolumeSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.PersistentVolumeClaimVolumeSource)(nil), (*v1.PersistentVolumeClaimVolumeSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_PersistentVolumeClaimVolumeSource_To_v1_PersistentVolumeClaimVolumeSource(a.(*core.PersistentVolumeClaimVolumeSource), b.(*v1.PersistentVolumeClaimVolumeSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.PersistentVolumeList)(nil), (*core.PersistentVolumeList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_PersistentVolumeList_To_core_PersistentVolumeList(a.(*v1.PersistentVolumeList), b.(*core.PersistentVolumeList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.PersistentVolumeList)(nil), (*v1.PersistentVolumeList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_PersistentVolumeList_To_v1_PersistentVolumeList(a.(*core.PersistentVolumeList), b.(*v1.PersistentVolumeList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.PersistentVolumeSource)(nil), (*core.PersistentVolumeSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_PersistentVolumeSource_To_core_PersistentVolumeSource(a.(*v1.PersistentVolumeSource), b.(*core.PersistentVolumeSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.PersistentVolumeSource)(nil), (*v1.PersistentVolumeSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_PersistentVolumeSource_To_v1_PersistentVolumeSource(a.(*core.PersistentVolumeSource), b.(*v1.PersistentVolumeSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.PersistentVolumeSpec)(nil), (*core.PersistentVolumeSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_PersistentVolumeSpec_To_core_PersistentVolumeSpec(a.(*v1.PersistentVolumeSpec), b.(*core.PersistentVolumeSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.PersistentVolumeSpec)(nil), (*v1.PersistentVolumeSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_PersistentVolumeSpec_To_v1_PersistentVolumeSpec(a.(*core.PersistentVolumeSpec), b.(*v1.PersistentVolumeSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.PersistentVolumeStatus)(nil), (*core.PersistentVolumeStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_PersistentVolumeStatus_To_core_PersistentVolumeStatus(a.(*v1.PersistentVolumeStatus), b.(*core.PersistentVolumeStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.PersistentVolumeStatus)(nil), (*v1.PersistentVolumeStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_PersistentVolumeStatus_To_v1_PersistentVolumeStatus(a.(*core.PersistentVolumeStatus), b.(*v1.PersistentVolumeStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.PhotonPersistentDiskVolumeSource)(nil), (*core.PhotonPersistentDiskVolumeSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_PhotonPersistentDiskVolumeSource_To_core_PhotonPersistentDiskVolumeSource(a.(*v1.PhotonPersistentDiskVolumeSource), b.(*core.PhotonPersistentDiskVolumeSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.PhotonPersistentDiskVolumeSource)(nil), (*v1.PhotonPersistentDiskVolumeSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_PhotonPersistentDiskVolumeSource_To_v1_PhotonPersistentDiskVolumeSource(a.(*core.PhotonPersistentDiskVolumeSource), b.(*v1.PhotonPersistentDiskVolumeSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.Pod)(nil), (*core.Pod)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_Pod_To_core_Pod(a.(*v1.Pod), b.(*core.Pod), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.Pod)(nil), (*v1.Pod)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_Pod_To_v1_Pod(a.(*core.Pod), b.(*v1.Pod), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.PodAffinity)(nil), (*core.PodAffinity)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_PodAffinity_To_core_PodAffinity(a.(*v1.PodAffinity), b.(*core.PodAffinity), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.PodAffinity)(nil), (*v1.PodAffinity)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_PodAffinity_To_v1_PodAffinity(a.(*core.PodAffinity), b.(*v1.PodAffinity), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.PodAffinityTerm)(nil), (*core.PodAffinityTerm)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_PodAffinityTerm_To_core_PodAffinityTerm(a.(*v1.PodAffinityTerm), b.(*core.PodAffinityTerm), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.PodAffinityTerm)(nil), (*v1.PodAffinityTerm)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_PodAffinityTerm_To_v1_PodAffinityTerm(a.(*core.PodAffinityTerm), b.(*v1.PodAffinityTerm), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.PodAntiAffinity)(nil), (*core.PodAntiAffinity)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_PodAntiAffinity_To_core_PodAntiAffinity(a.(*v1.PodAntiAffinity), b.(*core.PodAntiAffinity), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.PodAntiAffinity)(nil), (*v1.PodAntiAffinity)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_PodAntiAffinity_To_v1_PodAntiAffinity(a.(*core.PodAntiAffinity), b.(*v1.PodAntiAffinity), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.PodAttachOptions)(nil), (*core.PodAttachOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_PodAttachOptions_To_core_PodAttachOptions(a.(*v1.PodAttachOptions), b.(*core.PodAttachOptions), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.PodAttachOptions)(nil), (*v1.PodAttachOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_PodAttachOptions_To_v1_PodAttachOptions(a.(*core.PodAttachOptions), b.(*v1.PodAttachOptions), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.PodCondition)(nil), (*core.PodCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_PodCondition_To_core_PodCondition(a.(*v1.PodCondition), b.(*core.PodCondition), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.PodCondition)(nil), (*v1.PodCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_PodCondition_To_v1_PodCondition(a.(*core.PodCondition), b.(*v1.PodCondition), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.PodDNSConfig)(nil), (*core.PodDNSConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_PodDNSConfig_To_core_PodDNSConfig(a.(*v1.PodDNSConfig), b.(*core.PodDNSConfig), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.PodDNSConfig)(nil), (*v1.PodDNSConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_PodDNSConfig_To_v1_PodDNSConfig(a.(*core.PodDNSConfig), b.(*v1.PodDNSConfig), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.PodDNSConfigOption)(nil), (*core.PodDNSConfigOption)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_PodDNSConfigOption_To_core_PodDNSConfigOption(a.(*v1.PodDNSConfigOption), b.(*core.PodDNSConfigOption), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.PodDNSConfigOption)(nil), (*v1.PodDNSConfigOption)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_PodDNSConfigOption_To_v1_PodDNSConfigOption(a.(*core.PodDNSConfigOption), b.(*v1.PodDNSConfigOption), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.PodExecOptions)(nil), (*core.PodExecOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_PodExecOptions_To_core_PodExecOptions(a.(*v1.PodExecOptions), b.(*core.PodExecOptions), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.PodExecOptions)(nil), (*v1.PodExecOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_PodExecOptions_To_v1_PodExecOptions(a.(*core.PodExecOptions), b.(*v1.PodExecOptions), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.PodList)(nil), (*core.PodList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_PodList_To_core_PodList(a.(*v1.PodList), b.(*core.PodList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.PodList)(nil), (*v1.PodList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_PodList_To_v1_PodList(a.(*core.PodList), b.(*v1.PodList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.PodLogOptions)(nil), (*core.PodLogOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_PodLogOptions_To_core_PodLogOptions(a.(*v1.PodLogOptions), b.(*core.PodLogOptions), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.PodLogOptions)(nil), (*v1.PodLogOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_PodLogOptions_To_v1_PodLogOptions(a.(*core.PodLogOptions), b.(*v1.PodLogOptions), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.PodPortForwardOptions)(nil), (*core.PodPortForwardOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_PodPortForwardOptions_To_core_PodPortForwardOptions(a.(*v1.PodPortForwardOptions), b.(*core.PodPortForwardOptions), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.PodPortForwardOptions)(nil), (*v1.PodPortForwardOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_PodPortForwardOptions_To_v1_PodPortForwardOptions(a.(*core.PodPortForwardOptions), b.(*v1.PodPortForwardOptions), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.PodProxyOptions)(nil), (*core.PodProxyOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_PodProxyOptions_To_core_PodProxyOptions(a.(*v1.PodProxyOptions), b.(*core.PodProxyOptions), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.PodProxyOptions)(nil), (*v1.PodProxyOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_PodProxyOptions_To_v1_PodProxyOptions(a.(*core.PodProxyOptions), b.(*v1.PodProxyOptions), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.PodReadinessGate)(nil), (*core.PodReadinessGate)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_PodReadinessGate_To_core_PodReadinessGate(a.(*v1.PodReadinessGate), b.(*core.PodReadinessGate), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.PodReadinessGate)(nil), (*v1.PodReadinessGate)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_PodReadinessGate_To_v1_PodReadinessGate(a.(*core.PodReadinessGate), b.(*v1.PodReadinessGate), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.PodSecurityContext)(nil), (*core.PodSecurityContext)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_PodSecurityContext_To_core_PodSecurityContext(a.(*v1.PodSecurityContext), b.(*core.PodSecurityContext), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.PodSecurityContext)(nil), (*v1.PodSecurityContext)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_PodSecurityContext_To_v1_PodSecurityContext(a.(*core.PodSecurityContext), b.(*v1.PodSecurityContext), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.PodSignature)(nil), (*core.PodSignature)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_PodSignature_To_core_PodSignature(a.(*v1.PodSignature), b.(*core.PodSignature), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.PodSignature)(nil), (*v1.PodSignature)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_PodSignature_To_v1_PodSignature(a.(*core.PodSignature), b.(*v1.PodSignature), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.PodSpec)(nil), (*core.PodSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_PodSpec_To_core_PodSpec(a.(*v1.PodSpec), b.(*core.PodSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.PodSpec)(nil), (*v1.PodSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_PodSpec_To_v1_PodSpec(a.(*core.PodSpec), b.(*v1.PodSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.PodStatus)(nil), (*core.PodStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_PodStatus_To_core_PodStatus(a.(*v1.PodStatus), b.(*core.PodStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.PodStatus)(nil), (*v1.PodStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_PodStatus_To_v1_PodStatus(a.(*core.PodStatus), b.(*v1.PodStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.PodStatusResult)(nil), (*core.PodStatusResult)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_PodStatusResult_To_core_PodStatusResult(a.(*v1.PodStatusResult), b.(*core.PodStatusResult), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.PodStatusResult)(nil), (*v1.PodStatusResult)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_PodStatusResult_To_v1_PodStatusResult(a.(*core.PodStatusResult), b.(*v1.PodStatusResult), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.PodTemplate)(nil), (*core.PodTemplate)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_PodTemplate_To_core_PodTemplate(a.(*v1.PodTemplate), b.(*core.PodTemplate), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.PodTemplate)(nil), (*v1.PodTemplate)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_PodTemplate_To_v1_PodTemplate(a.(*core.PodTemplate), b.(*v1.PodTemplate), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.PodTemplateList)(nil), (*core.PodTemplateList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_PodTemplateList_To_core_PodTemplateList(a.(*v1.PodTemplateList), b.(*core.PodTemplateList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.PodTemplateList)(nil), (*v1.PodTemplateList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_PodTemplateList_To_v1_PodTemplateList(a.(*core.PodTemplateList), b.(*v1.PodTemplateList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.PodTemplateSpec)(nil), (*core.PodTemplateSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_PodTemplateSpec_To_core_PodTemplateSpec(a.(*v1.PodTemplateSpec), b.(*core.PodTemplateSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.PodTemplateSpec)(nil), (*v1.PodTemplateSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_PodTemplateSpec_To_v1_PodTemplateSpec(a.(*core.PodTemplateSpec), b.(*v1.PodTemplateSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.PortworxVolumeSource)(nil), (*core.PortworxVolumeSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_PortworxVolumeSource_To_core_PortworxVolumeSource(a.(*v1.PortworxVolumeSource), b.(*core.PortworxVolumeSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.PortworxVolumeSource)(nil), (*v1.PortworxVolumeSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_PortworxVolumeSource_To_v1_PortworxVolumeSource(a.(*core.PortworxVolumeSource), b.(*v1.PortworxVolumeSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.Preconditions)(nil), (*core.Preconditions)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_Preconditions_To_core_Preconditions(a.(*v1.Preconditions), b.(*core.Preconditions), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.Preconditions)(nil), (*v1.Preconditions)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_Preconditions_To_v1_Preconditions(a.(*core.Preconditions), b.(*v1.Preconditions), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.PreferAvoidPodsEntry)(nil), (*core.PreferAvoidPodsEntry)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_PreferAvoidPodsEntry_To_core_PreferAvoidPodsEntry(a.(*v1.PreferAvoidPodsEntry), b.(*core.PreferAvoidPodsEntry), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.PreferAvoidPodsEntry)(nil), (*v1.PreferAvoidPodsEntry)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_PreferAvoidPodsEntry_To_v1_PreferAvoidPodsEntry(a.(*core.PreferAvoidPodsEntry), b.(*v1.PreferAvoidPodsEntry), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.PreferredSchedulingTerm)(nil), (*core.PreferredSchedulingTerm)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_PreferredSchedulingTerm_To_core_PreferredSchedulingTerm(a.(*v1.PreferredSchedulingTerm), b.(*core.PreferredSchedulingTerm), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.PreferredSchedulingTerm)(nil), (*v1.PreferredSchedulingTerm)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_PreferredSchedulingTerm_To_v1_PreferredSchedulingTerm(a.(*core.PreferredSchedulingTerm), b.(*v1.PreferredSchedulingTerm), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.Probe)(nil), (*core.Probe)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_Probe_To_core_Probe(a.(*v1.Probe), b.(*core.Probe), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.Probe)(nil), (*v1.Probe)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_Probe_To_v1_Probe(a.(*core.Probe), b.(*v1.Probe), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.ProjectedVolumeSource)(nil), (*core.ProjectedVolumeSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_ProjectedVolumeSource_To_core_ProjectedVolumeSource(a.(*v1.ProjectedVolumeSource), b.(*core.ProjectedVolumeSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.ProjectedVolumeSource)(nil), (*v1.ProjectedVolumeSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_ProjectedVolumeSource_To_v1_ProjectedVolumeSource(a.(*core.ProjectedVolumeSource), b.(*v1.ProjectedVolumeSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.QuobyteVolumeSource)(nil), (*core.QuobyteVolumeSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_QuobyteVolumeSource_To_core_QuobyteVolumeSource(a.(*v1.QuobyteVolumeSource), b.(*core.QuobyteVolumeSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.QuobyteVolumeSource)(nil), (*v1.QuobyteVolumeSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_QuobyteVolumeSource_To_v1_QuobyteVolumeSource(a.(*core.QuobyteVolumeSource), b.(*v1.QuobyteVolumeSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.RBDPersistentVolumeSource)(nil), (*core.RBDPersistentVolumeSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_RBDPersistentVolumeSource_To_core_RBDPersistentVolumeSource(a.(*v1.RBDPersistentVolumeSource), b.(*core.RBDPersistentVolumeSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.RBDPersistentVolumeSource)(nil), (*v1.RBDPersistentVolumeSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_RBDPersistentVolumeSource_To_v1_RBDPersistentVolumeSource(a.(*core.RBDPersistentVolumeSource), b.(*v1.RBDPersistentVolumeSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.RBDVolumeSource)(nil), (*core.RBDVolumeSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_RBDVolumeSource_To_core_RBDVolumeSource(a.(*v1.RBDVolumeSource), b.(*core.RBDVolumeSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.RBDVolumeSource)(nil), (*v1.RBDVolumeSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_RBDVolumeSource_To_v1_RBDVolumeSource(a.(*core.RBDVolumeSource), b.(*v1.RBDVolumeSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.RangeAllocation)(nil), (*core.RangeAllocation)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_RangeAllocation_To_core_RangeAllocation(a.(*v1.RangeAllocation), b.(*core.RangeAllocation), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.RangeAllocation)(nil), (*v1.RangeAllocation)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_RangeAllocation_To_v1_RangeAllocation(a.(*core.RangeAllocation), b.(*v1.RangeAllocation), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.ReplicationController)(nil), (*core.ReplicationController)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_ReplicationController_To_core_ReplicationController(a.(*v1.ReplicationController), b.(*core.ReplicationController), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.ReplicationController)(nil), (*v1.ReplicationController)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_ReplicationController_To_v1_ReplicationController(a.(*core.ReplicationController), b.(*v1.ReplicationController), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.ReplicationControllerCondition)(nil), (*core.ReplicationControllerCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_ReplicationControllerCondition_To_core_ReplicationControllerCondition(a.(*v1.ReplicationControllerCondition), b.(*core.ReplicationControllerCondition), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.ReplicationControllerCondition)(nil), (*v1.ReplicationControllerCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_ReplicationControllerCondition_To_v1_ReplicationControllerCondition(a.(*core.ReplicationControllerCondition), b.(*v1.ReplicationControllerCondition), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.ReplicationControllerList)(nil), (*core.ReplicationControllerList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_ReplicationControllerList_To_core_ReplicationControllerList(a.(*v1.ReplicationControllerList), b.(*core.ReplicationControllerList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.ReplicationControllerList)(nil), (*v1.ReplicationControllerList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_ReplicationControllerList_To_v1_ReplicationControllerList(a.(*core.ReplicationControllerList), b.(*v1.ReplicationControllerList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.ReplicationControllerSpec)(nil), (*core.ReplicationControllerSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_ReplicationControllerSpec_To_core_ReplicationControllerSpec(a.(*v1.ReplicationControllerSpec), b.(*core.ReplicationControllerSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.ReplicationControllerSpec)(nil), (*v1.ReplicationControllerSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_ReplicationControllerSpec_To_v1_ReplicationControllerSpec(a.(*core.ReplicationControllerSpec), b.(*v1.ReplicationControllerSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.ReplicationControllerStatus)(nil), (*core.ReplicationControllerStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_ReplicationControllerStatus_To_core_ReplicationControllerStatus(a.(*v1.ReplicationControllerStatus), b.(*core.ReplicationControllerStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.ReplicationControllerStatus)(nil), (*v1.ReplicationControllerStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_ReplicationControllerStatus_To_v1_ReplicationControllerStatus(a.(*core.ReplicationControllerStatus), b.(*v1.ReplicationControllerStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.ResourceFieldSelector)(nil), (*core.ResourceFieldSelector)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_ResourceFieldSelector_To_core_ResourceFieldSelector(a.(*v1.ResourceFieldSelector), b.(*core.ResourceFieldSelector), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.ResourceFieldSelector)(nil), (*v1.ResourceFieldSelector)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_ResourceFieldSelector_To_v1_ResourceFieldSelector(a.(*core.ResourceFieldSelector), b.(*v1.ResourceFieldSelector), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.ResourceQuota)(nil), (*core.ResourceQuota)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_ResourceQuota_To_core_ResourceQuota(a.(*v1.ResourceQuota), b.(*core.ResourceQuota), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.ResourceQuota)(nil), (*v1.ResourceQuota)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_ResourceQuota_To_v1_ResourceQuota(a.(*core.ResourceQuota), b.(*v1.ResourceQuota), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.ResourceQuotaList)(nil), (*core.ResourceQuotaList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_ResourceQuotaList_To_core_ResourceQuotaList(a.(*v1.ResourceQuotaList), b.(*core.ResourceQuotaList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.ResourceQuotaList)(nil), (*v1.ResourceQuotaList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_ResourceQuotaList_To_v1_ResourceQuotaList(a.(*core.ResourceQuotaList), b.(*v1.ResourceQuotaList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.ResourceQuotaSpec)(nil), (*core.ResourceQuotaSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_ResourceQuotaSpec_To_core_ResourceQuotaSpec(a.(*v1.ResourceQuotaSpec), b.(*core.ResourceQuotaSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.ResourceQuotaSpec)(nil), (*v1.ResourceQuotaSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_ResourceQuotaSpec_To_v1_ResourceQuotaSpec(a.(*core.ResourceQuotaSpec), b.(*v1.ResourceQuotaSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.ResourceQuotaStatus)(nil), (*core.ResourceQuotaStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_ResourceQuotaStatus_To_core_ResourceQuotaStatus(a.(*v1.ResourceQuotaStatus), b.(*core.ResourceQuotaStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.ResourceQuotaStatus)(nil), (*v1.ResourceQuotaStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_ResourceQuotaStatus_To_v1_ResourceQuotaStatus(a.(*core.ResourceQuotaStatus), b.(*v1.ResourceQuotaStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.ResourceRequirements)(nil), (*core.ResourceRequirements)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_ResourceRequirements_To_core_ResourceRequirements(a.(*v1.ResourceRequirements), b.(*core.ResourceRequirements), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.ResourceRequirements)(nil), (*v1.ResourceRequirements)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_ResourceRequirements_To_v1_ResourceRequirements(a.(*core.ResourceRequirements), b.(*v1.ResourceRequirements), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.SELinuxOptions)(nil), (*core.SELinuxOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_SELinuxOptions_To_core_SELinuxOptions(a.(*v1.SELinuxOptions), b.(*core.SELinuxOptions), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.SELinuxOptions)(nil), (*v1.SELinuxOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_SELinuxOptions_To_v1_SELinuxOptions(a.(*core.SELinuxOptions), b.(*v1.SELinuxOptions), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.ScaleIOPersistentVolumeSource)(nil), (*core.ScaleIOPersistentVolumeSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_ScaleIOPersistentVolumeSource_To_core_ScaleIOPersistentVolumeSource(a.(*v1.ScaleIOPersistentVolumeSource), b.(*core.ScaleIOPersistentVolumeSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.ScaleIOPersistentVolumeSource)(nil), (*v1.ScaleIOPersistentVolumeSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_ScaleIOPersistentVolumeSource_To_v1_ScaleIOPersistentVolumeSource(a.(*core.ScaleIOPersistentVolumeSource), b.(*v1.ScaleIOPersistentVolumeSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.ScaleIOVolumeSource)(nil), (*core.ScaleIOVolumeSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_ScaleIOVolumeSource_To_core_ScaleIOVolumeSource(a.(*v1.ScaleIOVolumeSource), b.(*core.ScaleIOVolumeSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.ScaleIOVolumeSource)(nil), (*v1.ScaleIOVolumeSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_ScaleIOVolumeSource_To_v1_ScaleIOVolumeSource(a.(*core.ScaleIOVolumeSource), b.(*v1.ScaleIOVolumeSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.ScopeSelector)(nil), (*core.ScopeSelector)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_ScopeSelector_To_core_ScopeSelector(a.(*v1.ScopeSelector), b.(*core.ScopeSelector), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.ScopeSelector)(nil), (*v1.ScopeSelector)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_ScopeSelector_To_v1_ScopeSelector(a.(*core.ScopeSelector), b.(*v1.ScopeSelector), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.ScopedResourceSelectorRequirement)(nil), (*core.ScopedResourceSelectorRequirement)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_ScopedResourceSelectorRequirement_To_core_ScopedResourceSelectorRequirement(a.(*v1.ScopedResourceSelectorRequirement), b.(*core.ScopedResourceSelectorRequirement), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.ScopedResourceSelectorRequirement)(nil), (*v1.ScopedResourceSelectorRequirement)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_ScopedResourceSelectorRequirement_To_v1_ScopedResourceSelectorRequirement(a.(*core.ScopedResourceSelectorRequirement), b.(*v1.ScopedResourceSelectorRequirement), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.Secret)(nil), (*core.Secret)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_Secret_To_core_Secret(a.(*v1.Secret), b.(*core.Secret), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.Secret)(nil), (*v1.Secret)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_Secret_To_v1_Secret(a.(*core.Secret), b.(*v1.Secret), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.SecretEnvSource)(nil), (*core.SecretEnvSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_SecretEnvSource_To_core_SecretEnvSource(a.(*v1.SecretEnvSource), b.(*core.SecretEnvSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.SecretEnvSource)(nil), (*v1.SecretEnvSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_SecretEnvSource_To_v1_SecretEnvSource(a.(*core.SecretEnvSource), b.(*v1.SecretEnvSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.SecretKeySelector)(nil), (*core.SecretKeySelector)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_SecretKeySelector_To_core_SecretKeySelector(a.(*v1.SecretKeySelector), b.(*core.SecretKeySelector), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.SecretKeySelector)(nil), (*v1.SecretKeySelector)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_SecretKeySelector_To_v1_SecretKeySelector(a.(*core.SecretKeySelector), b.(*v1.SecretKeySelector), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.SecretList)(nil), (*core.SecretList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_SecretList_To_core_SecretList(a.(*v1.SecretList), b.(*core.SecretList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.SecretList)(nil), (*v1.SecretList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_SecretList_To_v1_SecretList(a.(*core.SecretList), b.(*v1.SecretList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.SecretProjection)(nil), (*core.SecretProjection)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_SecretProjection_To_core_SecretProjection(a.(*v1.SecretProjection), b.(*core.SecretProjection), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.SecretProjection)(nil), (*v1.SecretProjection)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_SecretProjection_To_v1_SecretProjection(a.(*core.SecretProjection), b.(*v1.SecretProjection), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.SecretReference)(nil), (*core.SecretReference)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_SecretReference_To_core_SecretReference(a.(*v1.SecretReference), b.(*core.SecretReference), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.SecretReference)(nil), (*v1.SecretReference)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_SecretReference_To_v1_SecretReference(a.(*core.SecretReference), b.(*v1.SecretReference), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.SecretVolumeSource)(nil), (*core.SecretVolumeSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_SecretVolumeSource_To_core_SecretVolumeSource(a.(*v1.SecretVolumeSource), b.(*core.SecretVolumeSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.SecretVolumeSource)(nil), (*v1.SecretVolumeSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_SecretVolumeSource_To_v1_SecretVolumeSource(a.(*core.SecretVolumeSource), b.(*v1.SecretVolumeSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.SecurityContext)(nil), (*core.SecurityContext)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_SecurityContext_To_core_SecurityContext(a.(*v1.SecurityContext), b.(*core.SecurityContext), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.SecurityContext)(nil), (*v1.SecurityContext)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_SecurityContext_To_v1_SecurityContext(a.(*core.SecurityContext), b.(*v1.SecurityContext), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.SerializedReference)(nil), (*core.SerializedReference)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_SerializedReference_To_core_SerializedReference(a.(*v1.SerializedReference), b.(*core.SerializedReference), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.SerializedReference)(nil), (*v1.SerializedReference)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_SerializedReference_To_v1_SerializedReference(a.(*core.SerializedReference), b.(*v1.SerializedReference), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.Service)(nil), (*core.Service)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_Service_To_core_Service(a.(*v1.Service), b.(*core.Service), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.Service)(nil), (*v1.Service)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_Service_To_v1_Service(a.(*core.Service), b.(*v1.Service), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.ServiceAccount)(nil), (*core.ServiceAccount)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_ServiceAccount_To_core_ServiceAccount(a.(*v1.ServiceAccount), b.(*core.ServiceAccount), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.ServiceAccount)(nil), (*v1.ServiceAccount)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_ServiceAccount_To_v1_ServiceAccount(a.(*core.ServiceAccount), b.(*v1.ServiceAccount), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.ServiceAccountList)(nil), (*core.ServiceAccountList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_ServiceAccountList_To_core_ServiceAccountList(a.(*v1.ServiceAccountList), b.(*core.ServiceAccountList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.ServiceAccountList)(nil), (*v1.ServiceAccountList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_ServiceAccountList_To_v1_ServiceAccountList(a.(*core.ServiceAccountList), b.(*v1.ServiceAccountList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.ServiceAccountTokenProjection)(nil), (*core.ServiceAccountTokenProjection)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_ServiceAccountTokenProjection_To_core_ServiceAccountTokenProjection(a.(*v1.ServiceAccountTokenProjection), b.(*core.ServiceAccountTokenProjection), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.ServiceAccountTokenProjection)(nil), (*v1.ServiceAccountTokenProjection)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_ServiceAccountTokenProjection_To_v1_ServiceAccountTokenProjection(a.(*core.ServiceAccountTokenProjection), b.(*v1.ServiceAccountTokenProjection), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.ServiceList)(nil), (*core.ServiceList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_ServiceList_To_core_ServiceList(a.(*v1.ServiceList), b.(*core.ServiceList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.ServiceList)(nil), (*v1.ServiceList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_ServiceList_To_v1_ServiceList(a.(*core.ServiceList), b.(*v1.ServiceList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.ServicePort)(nil), (*core.ServicePort)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_ServicePort_To_core_ServicePort(a.(*v1.ServicePort), b.(*core.ServicePort), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.ServicePort)(nil), (*v1.ServicePort)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_ServicePort_To_v1_ServicePort(a.(*core.ServicePort), b.(*v1.ServicePort), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.ServiceProxyOptions)(nil), (*core.ServiceProxyOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_ServiceProxyOptions_To_core_ServiceProxyOptions(a.(*v1.ServiceProxyOptions), b.(*core.ServiceProxyOptions), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.ServiceProxyOptions)(nil), (*v1.ServiceProxyOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_ServiceProxyOptions_To_v1_ServiceProxyOptions(a.(*core.ServiceProxyOptions), b.(*v1.ServiceProxyOptions), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.ServiceSpec)(nil), (*core.ServiceSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_ServiceSpec_To_core_ServiceSpec(a.(*v1.ServiceSpec), b.(*core.ServiceSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.ServiceSpec)(nil), (*v1.ServiceSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_ServiceSpec_To_v1_ServiceSpec(a.(*core.ServiceSpec), b.(*v1.ServiceSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.ServiceStatus)(nil), (*core.ServiceStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_ServiceStatus_To_core_ServiceStatus(a.(*v1.ServiceStatus), b.(*core.ServiceStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.ServiceStatus)(nil), (*v1.ServiceStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_ServiceStatus_To_v1_ServiceStatus(a.(*core.ServiceStatus), b.(*v1.ServiceStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.SessionAffinityConfig)(nil), (*core.SessionAffinityConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_SessionAffinityConfig_To_core_SessionAffinityConfig(a.(*v1.SessionAffinityConfig), b.(*core.SessionAffinityConfig), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.SessionAffinityConfig)(nil), (*v1.SessionAffinityConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_SessionAffinityConfig_To_v1_SessionAffinityConfig(a.(*core.SessionAffinityConfig), b.(*v1.SessionAffinityConfig), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.StorageOSPersistentVolumeSource)(nil), (*core.StorageOSPersistentVolumeSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_StorageOSPersistentVolumeSource_To_core_StorageOSPersistentVolumeSource(a.(*v1.StorageOSPersistentVolumeSource), b.(*core.StorageOSPersistentVolumeSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.StorageOSPersistentVolumeSource)(nil), (*v1.StorageOSPersistentVolumeSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_StorageOSPersistentVolumeSource_To_v1_StorageOSPersistentVolumeSource(a.(*core.StorageOSPersistentVolumeSource), b.(*v1.StorageOSPersistentVolumeSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.StorageOSVolumeSource)(nil), (*core.StorageOSVolumeSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_StorageOSVolumeSource_To_core_StorageOSVolumeSource(a.(*v1.StorageOSVolumeSource), b.(*core.StorageOSVolumeSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.StorageOSVolumeSource)(nil), (*v1.StorageOSVolumeSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_StorageOSVolumeSource_To_v1_StorageOSVolumeSource(a.(*core.StorageOSVolumeSource), b.(*v1.StorageOSVolumeSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.Sysctl)(nil), (*core.Sysctl)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_Sysctl_To_core_Sysctl(a.(*v1.Sysctl), b.(*core.Sysctl), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.Sysctl)(nil), (*v1.Sysctl)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_Sysctl_To_v1_Sysctl(a.(*core.Sysctl), b.(*v1.Sysctl), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.TCPSocketAction)(nil), (*core.TCPSocketAction)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_TCPSocketAction_To_core_TCPSocketAction(a.(*v1.TCPSocketAction), b.(*core.TCPSocketAction), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.TCPSocketAction)(nil), (*v1.TCPSocketAction)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_TCPSocketAction_To_v1_TCPSocketAction(a.(*core.TCPSocketAction), b.(*v1.TCPSocketAction), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.Taint)(nil), (*core.Taint)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_Taint_To_core_Taint(a.(*v1.Taint), b.(*core.Taint), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.Taint)(nil), (*v1.Taint)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_Taint_To_v1_Taint(a.(*core.Taint), b.(*v1.Taint), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.Toleration)(nil), (*core.Toleration)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_Toleration_To_core_Toleration(a.(*v1.Toleration), b.(*core.Toleration), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.Toleration)(nil), (*v1.Toleration)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_Toleration_To_v1_Toleration(a.(*core.Toleration), b.(*v1.Toleration), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.TopologySelectorLabelRequirement)(nil), (*core.TopologySelectorLabelRequirement)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_TopologySelectorLabelRequirement_To_core_TopologySelectorLabelRequirement(a.(*v1.TopologySelectorLabelRequirement), b.(*core.TopologySelectorLabelRequirement), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.TopologySelectorLabelRequirement)(nil), (*v1.TopologySelectorLabelRequirement)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_TopologySelectorLabelRequirement_To_v1_TopologySelectorLabelRequirement(a.(*core.TopologySelectorLabelRequirement), b.(*v1.TopologySelectorLabelRequirement), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.TopologySelectorTerm)(nil), (*core.TopologySelectorTerm)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_TopologySelectorTerm_To_core_TopologySelectorTerm(a.(*v1.TopologySelectorTerm), b.(*core.TopologySelectorTerm), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.TopologySelectorTerm)(nil), (*v1.TopologySelectorTerm)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_TopologySelectorTerm_To_v1_TopologySelectorTerm(a.(*core.TopologySelectorTerm), b.(*v1.TopologySelectorTerm), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.TypedLocalObjectReference)(nil), (*core.TypedLocalObjectReference)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_TypedLocalObjectReference_To_core_TypedLocalObjectReference(a.(*v1.TypedLocalObjectReference), b.(*core.TypedLocalObjectReference), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.TypedLocalObjectReference)(nil), (*v1.TypedLocalObjectReference)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_TypedLocalObjectReference_To_v1_TypedLocalObjectReference(a.(*core.TypedLocalObjectReference), b.(*v1.TypedLocalObjectReference), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.Volume)(nil), (*core.Volume)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_Volume_To_core_Volume(a.(*v1.Volume), b.(*core.Volume), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.Volume)(nil), (*v1.Volume)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_Volume_To_v1_Volume(a.(*core.Volume), b.(*v1.Volume), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.VolumeDevice)(nil), (*core.VolumeDevice)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_VolumeDevice_To_core_VolumeDevice(a.(*v1.VolumeDevice), b.(*core.VolumeDevice), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.VolumeDevice)(nil), (*v1.VolumeDevice)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_VolumeDevice_To_v1_VolumeDevice(a.(*core.VolumeDevice), b.(*v1.VolumeDevice), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.VolumeMount)(nil), (*core.VolumeMount)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_VolumeMount_To_core_VolumeMount(a.(*v1.VolumeMount), b.(*core.VolumeMount), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.VolumeMount)(nil), (*v1.VolumeMount)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_VolumeMount_To_v1_VolumeMount(a.(*core.VolumeMount), b.(*v1.VolumeMount), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.VolumeNodeAffinity)(nil), (*core.VolumeNodeAffinity)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_VolumeNodeAffinity_To_core_VolumeNodeAffinity(a.(*v1.VolumeNodeAffinity), b.(*core.VolumeNodeAffinity), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.VolumeNodeAffinity)(nil), (*v1.VolumeNodeAffinity)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_VolumeNodeAffinity_To_v1_VolumeNodeAffinity(a.(*core.VolumeNodeAffinity), b.(*v1.VolumeNodeAffinity), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.VolumeProjection)(nil), (*core.VolumeProjection)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_VolumeProjection_To_core_VolumeProjection(a.(*v1.VolumeProjection), b.(*core.VolumeProjection), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.VolumeProjection)(nil), (*v1.VolumeProjection)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_VolumeProjection_To_v1_VolumeProjection(a.(*core.VolumeProjection), b.(*v1.VolumeProjection), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.VolumeSource)(nil), (*core.VolumeSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_VolumeSource_To_core_VolumeSource(a.(*v1.VolumeSource), b.(*core.VolumeSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.VolumeSource)(nil), (*v1.VolumeSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_VolumeSource_To_v1_VolumeSource(a.(*core.VolumeSource), b.(*v1.VolumeSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.VsphereVirtualDiskVolumeSource)(nil), (*core.VsphereVirtualDiskVolumeSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_VsphereVirtualDiskVolumeSource_To_core_VsphereVirtualDiskVolumeSource(a.(*v1.VsphereVirtualDiskVolumeSource), b.(*core.VsphereVirtualDiskVolumeSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.VsphereVirtualDiskVolumeSource)(nil), (*v1.VsphereVirtualDiskVolumeSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_VsphereVirtualDiskVolumeSource_To_v1_VsphereVirtualDiskVolumeSource(a.(*core.VsphereVirtualDiskVolumeSource), b.(*v1.VsphereVirtualDiskVolumeSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1.WeightedPodAffinityTerm)(nil), (*core.WeightedPodAffinityTerm)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_WeightedPodAffinityTerm_To_core_WeightedPodAffinityTerm(a.(*v1.WeightedPodAffinityTerm), b.(*core.WeightedPodAffinityTerm), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*core.WeightedPodAffinityTerm)(nil), (*v1.WeightedPodAffinityTerm)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_WeightedPodAffinityTerm_To_v1_WeightedPodAffinityTerm(a.(*core.WeightedPodAffinityTerm), b.(*v1.WeightedPodAffinityTerm), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*apps.ReplicaSetSpec)(nil), (*v1.ReplicationControllerSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_ReplicaSetSpec_To_v1_ReplicationControllerSpec(a.(*apps.ReplicaSetSpec), b.(*v1.ReplicationControllerSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*apps.ReplicaSetStatus)(nil), (*v1.ReplicationControllerStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_ReplicaSetStatus_To_v1_ReplicationControllerStatus(a.(*apps.ReplicaSetStatus), b.(*v1.ReplicationControllerStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*apps.ReplicaSet)(nil), (*v1.ReplicationController)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_apps_ReplicaSet_To_v1_ReplicationController(a.(*apps.ReplicaSet), b.(*v1.ReplicationController), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*core.PodSecurityContext)(nil), (*v1.PodSecurityContext)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_PodSecurityContext_To_v1_PodSecurityContext(a.(*core.PodSecurityContext), b.(*v1.PodSecurityContext), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*core.PodSpec)(nil), (*v1.PodSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_PodSpec_To_v1_PodSpec(a.(*core.PodSpec), b.(*v1.PodSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*core.PodTemplateSpec)(nil), (*v1.PodTemplateSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_PodTemplateSpec_To_v1_PodTemplateSpec(a.(*core.PodTemplateSpec), b.(*v1.PodTemplateSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*core.Pod)(nil), (*v1.Pod)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_Pod_To_v1_Pod(a.(*core.Pod), b.(*v1.Pod), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*core.ReplicationControllerSpec)(nil), (*v1.ReplicationControllerSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_ReplicationControllerSpec_To_v1_ReplicationControllerSpec(a.(*core.ReplicationControllerSpec), b.(*v1.ReplicationControllerSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*core.SecurityContext)(nil), (*v1.SecurityContext)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_core_SecurityContext_To_v1_SecurityContext(a.(*core.SecurityContext), b.(*v1.SecurityContext), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*v1.PodSecurityContext)(nil), (*core.PodSecurityContext)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_PodSecurityContext_To_core_PodSecurityContext(a.(*v1.PodSecurityContext), b.(*core.PodSecurityContext), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*v1.PodSpec)(nil), (*core.PodSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_PodSpec_To_core_PodSpec(a.(*v1.PodSpec), b.(*core.PodSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*v1.PodTemplateSpec)(nil), (*core.PodTemplateSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_PodTemplateSpec_To_core_PodTemplateSpec(a.(*v1.PodTemplateSpec), b.(*core.PodTemplateSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*v1.Pod)(nil), (*core.Pod)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_Pod_To_core_Pod(a.(*v1.Pod), b.(*core.Pod), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*v1.ReplicationControllerSpec)(nil), (*apps.ReplicaSetSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_ReplicationControllerSpec_To_apps_ReplicaSetSpec(a.(*v1.ReplicationControllerSpec), b.(*apps.ReplicaSetSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*v1.ReplicationControllerSpec)(nil), (*core.ReplicationControllerSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_ReplicationControllerSpec_To_core_ReplicationControllerSpec(a.(*v1.ReplicationControllerSpec), b.(*core.ReplicationControllerSpec), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*v1.ReplicationControllerStatus)(nil), (*apps.ReplicaSetStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_ReplicationControllerStatus_To_apps_ReplicaSetStatus(a.(*v1.ReplicationControllerStatus), b.(*apps.ReplicaSetStatus), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*v1.ReplicationController)(nil), (*apps.ReplicaSet)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_ReplicationController_To_apps_ReplicaSet(a.(*v1.ReplicationController), b.(*apps.ReplicaSet), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*v1.ResourceList)(nil), (*core.ResourceList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_ResourceList_To_core_ResourceList(a.(*v1.ResourceList), b.(*core.ResourceList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*v1.Secret)(nil), (*core.Secret)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1_Secret_To_core_Secret(a.(*v1.Secret), b.(*core.Secret), scope)
 }); err != nil {
  return err
 }
 return nil
}
func autoConvert_v1_AWSElasticBlockStoreVolumeSource_To_core_AWSElasticBlockStoreVolumeSource(in *v1.AWSElasticBlockStoreVolumeSource, out *core.AWSElasticBlockStoreVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.VolumeID = in.VolumeID
 out.FSType = in.FSType
 out.Partition = in.Partition
 out.ReadOnly = in.ReadOnly
 return nil
}
func Convert_v1_AWSElasticBlockStoreVolumeSource_To_core_AWSElasticBlockStoreVolumeSource(in *v1.AWSElasticBlockStoreVolumeSource, out *core.AWSElasticBlockStoreVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_AWSElasticBlockStoreVolumeSource_To_core_AWSElasticBlockStoreVolumeSource(in, out, s)
}
func autoConvert_core_AWSElasticBlockStoreVolumeSource_To_v1_AWSElasticBlockStoreVolumeSource(in *core.AWSElasticBlockStoreVolumeSource, out *v1.AWSElasticBlockStoreVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.VolumeID = in.VolumeID
 out.FSType = in.FSType
 out.Partition = in.Partition
 out.ReadOnly = in.ReadOnly
 return nil
}
func Convert_core_AWSElasticBlockStoreVolumeSource_To_v1_AWSElasticBlockStoreVolumeSource(in *core.AWSElasticBlockStoreVolumeSource, out *v1.AWSElasticBlockStoreVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_AWSElasticBlockStoreVolumeSource_To_v1_AWSElasticBlockStoreVolumeSource(in, out, s)
}
func autoConvert_v1_Affinity_To_core_Affinity(in *v1.Affinity, out *core.Affinity, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.NodeAffinity = (*core.NodeAffinity)(unsafe.Pointer(in.NodeAffinity))
 out.PodAffinity = (*core.PodAffinity)(unsafe.Pointer(in.PodAffinity))
 out.PodAntiAffinity = (*core.PodAntiAffinity)(unsafe.Pointer(in.PodAntiAffinity))
 return nil
}
func Convert_v1_Affinity_To_core_Affinity(in *v1.Affinity, out *core.Affinity, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_Affinity_To_core_Affinity(in, out, s)
}
func autoConvert_core_Affinity_To_v1_Affinity(in *core.Affinity, out *v1.Affinity, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.NodeAffinity = (*v1.NodeAffinity)(unsafe.Pointer(in.NodeAffinity))
 out.PodAffinity = (*v1.PodAffinity)(unsafe.Pointer(in.PodAffinity))
 out.PodAntiAffinity = (*v1.PodAntiAffinity)(unsafe.Pointer(in.PodAntiAffinity))
 return nil
}
func Convert_core_Affinity_To_v1_Affinity(in *core.Affinity, out *v1.Affinity, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_Affinity_To_v1_Affinity(in, out, s)
}
func autoConvert_v1_AttachedVolume_To_core_AttachedVolume(in *v1.AttachedVolume, out *core.AttachedVolume, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Name = core.UniqueVolumeName(in.Name)
 out.DevicePath = in.DevicePath
 return nil
}
func Convert_v1_AttachedVolume_To_core_AttachedVolume(in *v1.AttachedVolume, out *core.AttachedVolume, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_AttachedVolume_To_core_AttachedVolume(in, out, s)
}
func autoConvert_core_AttachedVolume_To_v1_AttachedVolume(in *core.AttachedVolume, out *v1.AttachedVolume, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Name = v1.UniqueVolumeName(in.Name)
 out.DevicePath = in.DevicePath
 return nil
}
func Convert_core_AttachedVolume_To_v1_AttachedVolume(in *core.AttachedVolume, out *v1.AttachedVolume, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_AttachedVolume_To_v1_AttachedVolume(in, out, s)
}
func autoConvert_v1_AvoidPods_To_core_AvoidPods(in *v1.AvoidPods, out *core.AvoidPods, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.PreferAvoidPods = *(*[]core.PreferAvoidPodsEntry)(unsafe.Pointer(&in.PreferAvoidPods))
 return nil
}
func Convert_v1_AvoidPods_To_core_AvoidPods(in *v1.AvoidPods, out *core.AvoidPods, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_AvoidPods_To_core_AvoidPods(in, out, s)
}
func autoConvert_core_AvoidPods_To_v1_AvoidPods(in *core.AvoidPods, out *v1.AvoidPods, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.PreferAvoidPods = *(*[]v1.PreferAvoidPodsEntry)(unsafe.Pointer(&in.PreferAvoidPods))
 return nil
}
func Convert_core_AvoidPods_To_v1_AvoidPods(in *core.AvoidPods, out *v1.AvoidPods, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_AvoidPods_To_v1_AvoidPods(in, out, s)
}
func autoConvert_v1_AzureDiskVolumeSource_To_core_AzureDiskVolumeSource(in *v1.AzureDiskVolumeSource, out *core.AzureDiskVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.DiskName = in.DiskName
 out.DataDiskURI = in.DataDiskURI
 out.CachingMode = (*core.AzureDataDiskCachingMode)(unsafe.Pointer(in.CachingMode))
 out.FSType = (*string)(unsafe.Pointer(in.FSType))
 out.ReadOnly = (*bool)(unsafe.Pointer(in.ReadOnly))
 out.Kind = (*core.AzureDataDiskKind)(unsafe.Pointer(in.Kind))
 return nil
}
func Convert_v1_AzureDiskVolumeSource_To_core_AzureDiskVolumeSource(in *v1.AzureDiskVolumeSource, out *core.AzureDiskVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_AzureDiskVolumeSource_To_core_AzureDiskVolumeSource(in, out, s)
}
func autoConvert_core_AzureDiskVolumeSource_To_v1_AzureDiskVolumeSource(in *core.AzureDiskVolumeSource, out *v1.AzureDiskVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.DiskName = in.DiskName
 out.DataDiskURI = in.DataDiskURI
 out.CachingMode = (*v1.AzureDataDiskCachingMode)(unsafe.Pointer(in.CachingMode))
 out.FSType = (*string)(unsafe.Pointer(in.FSType))
 out.ReadOnly = (*bool)(unsafe.Pointer(in.ReadOnly))
 out.Kind = (*v1.AzureDataDiskKind)(unsafe.Pointer(in.Kind))
 return nil
}
func Convert_core_AzureDiskVolumeSource_To_v1_AzureDiskVolumeSource(in *core.AzureDiskVolumeSource, out *v1.AzureDiskVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_AzureDiskVolumeSource_To_v1_AzureDiskVolumeSource(in, out, s)
}
func autoConvert_v1_AzureFilePersistentVolumeSource_To_core_AzureFilePersistentVolumeSource(in *v1.AzureFilePersistentVolumeSource, out *core.AzureFilePersistentVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.SecretName = in.SecretName
 out.ShareName = in.ShareName
 out.ReadOnly = in.ReadOnly
 out.SecretNamespace = (*string)(unsafe.Pointer(in.SecretNamespace))
 return nil
}
func Convert_v1_AzureFilePersistentVolumeSource_To_core_AzureFilePersistentVolumeSource(in *v1.AzureFilePersistentVolumeSource, out *core.AzureFilePersistentVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_AzureFilePersistentVolumeSource_To_core_AzureFilePersistentVolumeSource(in, out, s)
}
func autoConvert_core_AzureFilePersistentVolumeSource_To_v1_AzureFilePersistentVolumeSource(in *core.AzureFilePersistentVolumeSource, out *v1.AzureFilePersistentVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.SecretName = in.SecretName
 out.ShareName = in.ShareName
 out.ReadOnly = in.ReadOnly
 out.SecretNamespace = (*string)(unsafe.Pointer(in.SecretNamespace))
 return nil
}
func Convert_core_AzureFilePersistentVolumeSource_To_v1_AzureFilePersistentVolumeSource(in *core.AzureFilePersistentVolumeSource, out *v1.AzureFilePersistentVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_AzureFilePersistentVolumeSource_To_v1_AzureFilePersistentVolumeSource(in, out, s)
}
func autoConvert_v1_AzureFileVolumeSource_To_core_AzureFileVolumeSource(in *v1.AzureFileVolumeSource, out *core.AzureFileVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.SecretName = in.SecretName
 out.ShareName = in.ShareName
 out.ReadOnly = in.ReadOnly
 return nil
}
func Convert_v1_AzureFileVolumeSource_To_core_AzureFileVolumeSource(in *v1.AzureFileVolumeSource, out *core.AzureFileVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_AzureFileVolumeSource_To_core_AzureFileVolumeSource(in, out, s)
}
func autoConvert_core_AzureFileVolumeSource_To_v1_AzureFileVolumeSource(in *core.AzureFileVolumeSource, out *v1.AzureFileVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.SecretName = in.SecretName
 out.ShareName = in.ShareName
 out.ReadOnly = in.ReadOnly
 return nil
}
func Convert_core_AzureFileVolumeSource_To_v1_AzureFileVolumeSource(in *core.AzureFileVolumeSource, out *v1.AzureFileVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_AzureFileVolumeSource_To_v1_AzureFileVolumeSource(in, out, s)
}
func autoConvert_v1_Binding_To_core_Binding(in *v1.Binding, out *core.Binding, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_v1_ObjectReference_To_core_ObjectReference(&in.Target, &out.Target, s); err != nil {
  return err
 }
 return nil
}
func Convert_v1_Binding_To_core_Binding(in *v1.Binding, out *core.Binding, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_Binding_To_core_Binding(in, out, s)
}
func autoConvert_core_Binding_To_v1_Binding(in *core.Binding, out *v1.Binding, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_core_ObjectReference_To_v1_ObjectReference(&in.Target, &out.Target, s); err != nil {
  return err
 }
 return nil
}
func Convert_core_Binding_To_v1_Binding(in *core.Binding, out *v1.Binding, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_Binding_To_v1_Binding(in, out, s)
}
func autoConvert_v1_CSIPersistentVolumeSource_To_core_CSIPersistentVolumeSource(in *v1.CSIPersistentVolumeSource, out *core.CSIPersistentVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Driver = in.Driver
 out.VolumeHandle = in.VolumeHandle
 out.ReadOnly = in.ReadOnly
 out.FSType = in.FSType
 out.VolumeAttributes = *(*map[string]string)(unsafe.Pointer(&in.VolumeAttributes))
 out.ControllerPublishSecretRef = (*core.SecretReference)(unsafe.Pointer(in.ControllerPublishSecretRef))
 out.NodeStageSecretRef = (*core.SecretReference)(unsafe.Pointer(in.NodeStageSecretRef))
 out.NodePublishSecretRef = (*core.SecretReference)(unsafe.Pointer(in.NodePublishSecretRef))
 return nil
}
func Convert_v1_CSIPersistentVolumeSource_To_core_CSIPersistentVolumeSource(in *v1.CSIPersistentVolumeSource, out *core.CSIPersistentVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_CSIPersistentVolumeSource_To_core_CSIPersistentVolumeSource(in, out, s)
}
func autoConvert_core_CSIPersistentVolumeSource_To_v1_CSIPersistentVolumeSource(in *core.CSIPersistentVolumeSource, out *v1.CSIPersistentVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Driver = in.Driver
 out.VolumeHandle = in.VolumeHandle
 out.ReadOnly = in.ReadOnly
 out.FSType = in.FSType
 out.VolumeAttributes = *(*map[string]string)(unsafe.Pointer(&in.VolumeAttributes))
 out.ControllerPublishSecretRef = (*v1.SecretReference)(unsafe.Pointer(in.ControllerPublishSecretRef))
 out.NodeStageSecretRef = (*v1.SecretReference)(unsafe.Pointer(in.NodeStageSecretRef))
 out.NodePublishSecretRef = (*v1.SecretReference)(unsafe.Pointer(in.NodePublishSecretRef))
 return nil
}
func Convert_core_CSIPersistentVolumeSource_To_v1_CSIPersistentVolumeSource(in *core.CSIPersistentVolumeSource, out *v1.CSIPersistentVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_CSIPersistentVolumeSource_To_v1_CSIPersistentVolumeSource(in, out, s)
}
func autoConvert_v1_Capabilities_To_core_Capabilities(in *v1.Capabilities, out *core.Capabilities, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Add = *(*[]core.Capability)(unsafe.Pointer(&in.Add))
 out.Drop = *(*[]core.Capability)(unsafe.Pointer(&in.Drop))
 return nil
}
func Convert_v1_Capabilities_To_core_Capabilities(in *v1.Capabilities, out *core.Capabilities, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_Capabilities_To_core_Capabilities(in, out, s)
}
func autoConvert_core_Capabilities_To_v1_Capabilities(in *core.Capabilities, out *v1.Capabilities, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Add = *(*[]v1.Capability)(unsafe.Pointer(&in.Add))
 out.Drop = *(*[]v1.Capability)(unsafe.Pointer(&in.Drop))
 return nil
}
func Convert_core_Capabilities_To_v1_Capabilities(in *core.Capabilities, out *v1.Capabilities, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_Capabilities_To_v1_Capabilities(in, out, s)
}
func autoConvert_v1_CephFSPersistentVolumeSource_To_core_CephFSPersistentVolumeSource(in *v1.CephFSPersistentVolumeSource, out *core.CephFSPersistentVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Monitors = *(*[]string)(unsafe.Pointer(&in.Monitors))
 out.Path = in.Path
 out.User = in.User
 out.SecretFile = in.SecretFile
 out.SecretRef = (*core.SecretReference)(unsafe.Pointer(in.SecretRef))
 out.ReadOnly = in.ReadOnly
 return nil
}
func Convert_v1_CephFSPersistentVolumeSource_To_core_CephFSPersistentVolumeSource(in *v1.CephFSPersistentVolumeSource, out *core.CephFSPersistentVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_CephFSPersistentVolumeSource_To_core_CephFSPersistentVolumeSource(in, out, s)
}
func autoConvert_core_CephFSPersistentVolumeSource_To_v1_CephFSPersistentVolumeSource(in *core.CephFSPersistentVolumeSource, out *v1.CephFSPersistentVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Monitors = *(*[]string)(unsafe.Pointer(&in.Monitors))
 out.Path = in.Path
 out.User = in.User
 out.SecretFile = in.SecretFile
 out.SecretRef = (*v1.SecretReference)(unsafe.Pointer(in.SecretRef))
 out.ReadOnly = in.ReadOnly
 return nil
}
func Convert_core_CephFSPersistentVolumeSource_To_v1_CephFSPersistentVolumeSource(in *core.CephFSPersistentVolumeSource, out *v1.CephFSPersistentVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_CephFSPersistentVolumeSource_To_v1_CephFSPersistentVolumeSource(in, out, s)
}
func autoConvert_v1_CephFSVolumeSource_To_core_CephFSVolumeSource(in *v1.CephFSVolumeSource, out *core.CephFSVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Monitors = *(*[]string)(unsafe.Pointer(&in.Monitors))
 out.Path = in.Path
 out.User = in.User
 out.SecretFile = in.SecretFile
 out.SecretRef = (*core.LocalObjectReference)(unsafe.Pointer(in.SecretRef))
 out.ReadOnly = in.ReadOnly
 return nil
}
func Convert_v1_CephFSVolumeSource_To_core_CephFSVolumeSource(in *v1.CephFSVolumeSource, out *core.CephFSVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_CephFSVolumeSource_To_core_CephFSVolumeSource(in, out, s)
}
func autoConvert_core_CephFSVolumeSource_To_v1_CephFSVolumeSource(in *core.CephFSVolumeSource, out *v1.CephFSVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Monitors = *(*[]string)(unsafe.Pointer(&in.Monitors))
 out.Path = in.Path
 out.User = in.User
 out.SecretFile = in.SecretFile
 out.SecretRef = (*v1.LocalObjectReference)(unsafe.Pointer(in.SecretRef))
 out.ReadOnly = in.ReadOnly
 return nil
}
func Convert_core_CephFSVolumeSource_To_v1_CephFSVolumeSource(in *core.CephFSVolumeSource, out *v1.CephFSVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_CephFSVolumeSource_To_v1_CephFSVolumeSource(in, out, s)
}
func autoConvert_v1_CinderPersistentVolumeSource_To_core_CinderPersistentVolumeSource(in *v1.CinderPersistentVolumeSource, out *core.CinderPersistentVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.VolumeID = in.VolumeID
 out.FSType = in.FSType
 out.ReadOnly = in.ReadOnly
 out.SecretRef = (*core.SecretReference)(unsafe.Pointer(in.SecretRef))
 return nil
}
func Convert_v1_CinderPersistentVolumeSource_To_core_CinderPersistentVolumeSource(in *v1.CinderPersistentVolumeSource, out *core.CinderPersistentVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_CinderPersistentVolumeSource_To_core_CinderPersistentVolumeSource(in, out, s)
}
func autoConvert_core_CinderPersistentVolumeSource_To_v1_CinderPersistentVolumeSource(in *core.CinderPersistentVolumeSource, out *v1.CinderPersistentVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.VolumeID = in.VolumeID
 out.FSType = in.FSType
 out.ReadOnly = in.ReadOnly
 out.SecretRef = (*v1.SecretReference)(unsafe.Pointer(in.SecretRef))
 return nil
}
func Convert_core_CinderPersistentVolumeSource_To_v1_CinderPersistentVolumeSource(in *core.CinderPersistentVolumeSource, out *v1.CinderPersistentVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_CinderPersistentVolumeSource_To_v1_CinderPersistentVolumeSource(in, out, s)
}
func autoConvert_v1_CinderVolumeSource_To_core_CinderVolumeSource(in *v1.CinderVolumeSource, out *core.CinderVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.VolumeID = in.VolumeID
 out.FSType = in.FSType
 out.ReadOnly = in.ReadOnly
 out.SecretRef = (*core.LocalObjectReference)(unsafe.Pointer(in.SecretRef))
 return nil
}
func Convert_v1_CinderVolumeSource_To_core_CinderVolumeSource(in *v1.CinderVolumeSource, out *core.CinderVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_CinderVolumeSource_To_core_CinderVolumeSource(in, out, s)
}
func autoConvert_core_CinderVolumeSource_To_v1_CinderVolumeSource(in *core.CinderVolumeSource, out *v1.CinderVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.VolumeID = in.VolumeID
 out.FSType = in.FSType
 out.ReadOnly = in.ReadOnly
 out.SecretRef = (*v1.LocalObjectReference)(unsafe.Pointer(in.SecretRef))
 return nil
}
func Convert_core_CinderVolumeSource_To_v1_CinderVolumeSource(in *core.CinderVolumeSource, out *v1.CinderVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_CinderVolumeSource_To_v1_CinderVolumeSource(in, out, s)
}
func autoConvert_v1_ClientIPConfig_To_core_ClientIPConfig(in *v1.ClientIPConfig, out *core.ClientIPConfig, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.TimeoutSeconds = (*int32)(unsafe.Pointer(in.TimeoutSeconds))
 return nil
}
func Convert_v1_ClientIPConfig_To_core_ClientIPConfig(in *v1.ClientIPConfig, out *core.ClientIPConfig, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_ClientIPConfig_To_core_ClientIPConfig(in, out, s)
}
func autoConvert_core_ClientIPConfig_To_v1_ClientIPConfig(in *core.ClientIPConfig, out *v1.ClientIPConfig, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.TimeoutSeconds = (*int32)(unsafe.Pointer(in.TimeoutSeconds))
 return nil
}
func Convert_core_ClientIPConfig_To_v1_ClientIPConfig(in *core.ClientIPConfig, out *v1.ClientIPConfig, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_ClientIPConfig_To_v1_ClientIPConfig(in, out, s)
}
func autoConvert_v1_ComponentCondition_To_core_ComponentCondition(in *v1.ComponentCondition, out *core.ComponentCondition, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Type = core.ComponentConditionType(in.Type)
 out.Status = core.ConditionStatus(in.Status)
 out.Message = in.Message
 out.Error = in.Error
 return nil
}
func Convert_v1_ComponentCondition_To_core_ComponentCondition(in *v1.ComponentCondition, out *core.ComponentCondition, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_ComponentCondition_To_core_ComponentCondition(in, out, s)
}
func autoConvert_core_ComponentCondition_To_v1_ComponentCondition(in *core.ComponentCondition, out *v1.ComponentCondition, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Type = v1.ComponentConditionType(in.Type)
 out.Status = v1.ConditionStatus(in.Status)
 out.Message = in.Message
 out.Error = in.Error
 return nil
}
func Convert_core_ComponentCondition_To_v1_ComponentCondition(in *core.ComponentCondition, out *v1.ComponentCondition, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_ComponentCondition_To_v1_ComponentCondition(in, out, s)
}
func autoConvert_v1_ComponentStatus_To_core_ComponentStatus(in *v1.ComponentStatus, out *core.ComponentStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 out.Conditions = *(*[]core.ComponentCondition)(unsafe.Pointer(&in.Conditions))
 return nil
}
func Convert_v1_ComponentStatus_To_core_ComponentStatus(in *v1.ComponentStatus, out *core.ComponentStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_ComponentStatus_To_core_ComponentStatus(in, out, s)
}
func autoConvert_core_ComponentStatus_To_v1_ComponentStatus(in *core.ComponentStatus, out *v1.ComponentStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 out.Conditions = *(*[]v1.ComponentCondition)(unsafe.Pointer(&in.Conditions))
 return nil
}
func Convert_core_ComponentStatus_To_v1_ComponentStatus(in *core.ComponentStatus, out *v1.ComponentStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_ComponentStatus_To_v1_ComponentStatus(in, out, s)
}
func autoConvert_v1_ComponentStatusList_To_core_ComponentStatusList(in *v1.ComponentStatusList, out *core.ComponentStatusList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 out.Items = *(*[]core.ComponentStatus)(unsafe.Pointer(&in.Items))
 return nil
}
func Convert_v1_ComponentStatusList_To_core_ComponentStatusList(in *v1.ComponentStatusList, out *core.ComponentStatusList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_ComponentStatusList_To_core_ComponentStatusList(in, out, s)
}
func autoConvert_core_ComponentStatusList_To_v1_ComponentStatusList(in *core.ComponentStatusList, out *v1.ComponentStatusList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 out.Items = *(*[]v1.ComponentStatus)(unsafe.Pointer(&in.Items))
 return nil
}
func Convert_core_ComponentStatusList_To_v1_ComponentStatusList(in *core.ComponentStatusList, out *v1.ComponentStatusList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_ComponentStatusList_To_v1_ComponentStatusList(in, out, s)
}
func autoConvert_v1_ConfigMap_To_core_ConfigMap(in *v1.ConfigMap, out *core.ConfigMap, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 out.Data = *(*map[string]string)(unsafe.Pointer(&in.Data))
 out.BinaryData = *(*map[string][]byte)(unsafe.Pointer(&in.BinaryData))
 return nil
}
func Convert_v1_ConfigMap_To_core_ConfigMap(in *v1.ConfigMap, out *core.ConfigMap, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_ConfigMap_To_core_ConfigMap(in, out, s)
}
func autoConvert_core_ConfigMap_To_v1_ConfigMap(in *core.ConfigMap, out *v1.ConfigMap, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 out.Data = *(*map[string]string)(unsafe.Pointer(&in.Data))
 out.BinaryData = *(*map[string][]byte)(unsafe.Pointer(&in.BinaryData))
 return nil
}
func Convert_core_ConfigMap_To_v1_ConfigMap(in *core.ConfigMap, out *v1.ConfigMap, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_ConfigMap_To_v1_ConfigMap(in, out, s)
}
func autoConvert_v1_ConfigMapEnvSource_To_core_ConfigMapEnvSource(in *v1.ConfigMapEnvSource, out *core.ConfigMapEnvSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := Convert_v1_LocalObjectReference_To_core_LocalObjectReference(&in.LocalObjectReference, &out.LocalObjectReference, s); err != nil {
  return err
 }
 out.Optional = (*bool)(unsafe.Pointer(in.Optional))
 return nil
}
func Convert_v1_ConfigMapEnvSource_To_core_ConfigMapEnvSource(in *v1.ConfigMapEnvSource, out *core.ConfigMapEnvSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_ConfigMapEnvSource_To_core_ConfigMapEnvSource(in, out, s)
}
func autoConvert_core_ConfigMapEnvSource_To_v1_ConfigMapEnvSource(in *core.ConfigMapEnvSource, out *v1.ConfigMapEnvSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := Convert_core_LocalObjectReference_To_v1_LocalObjectReference(&in.LocalObjectReference, &out.LocalObjectReference, s); err != nil {
  return err
 }
 out.Optional = (*bool)(unsafe.Pointer(in.Optional))
 return nil
}
func Convert_core_ConfigMapEnvSource_To_v1_ConfigMapEnvSource(in *core.ConfigMapEnvSource, out *v1.ConfigMapEnvSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_ConfigMapEnvSource_To_v1_ConfigMapEnvSource(in, out, s)
}
func autoConvert_v1_ConfigMapKeySelector_To_core_ConfigMapKeySelector(in *v1.ConfigMapKeySelector, out *core.ConfigMapKeySelector, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := Convert_v1_LocalObjectReference_To_core_LocalObjectReference(&in.LocalObjectReference, &out.LocalObjectReference, s); err != nil {
  return err
 }
 out.Key = in.Key
 out.Optional = (*bool)(unsafe.Pointer(in.Optional))
 return nil
}
func Convert_v1_ConfigMapKeySelector_To_core_ConfigMapKeySelector(in *v1.ConfigMapKeySelector, out *core.ConfigMapKeySelector, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_ConfigMapKeySelector_To_core_ConfigMapKeySelector(in, out, s)
}
func autoConvert_core_ConfigMapKeySelector_To_v1_ConfigMapKeySelector(in *core.ConfigMapKeySelector, out *v1.ConfigMapKeySelector, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := Convert_core_LocalObjectReference_To_v1_LocalObjectReference(&in.LocalObjectReference, &out.LocalObjectReference, s); err != nil {
  return err
 }
 out.Key = in.Key
 out.Optional = (*bool)(unsafe.Pointer(in.Optional))
 return nil
}
func Convert_core_ConfigMapKeySelector_To_v1_ConfigMapKeySelector(in *core.ConfigMapKeySelector, out *v1.ConfigMapKeySelector, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_ConfigMapKeySelector_To_v1_ConfigMapKeySelector(in, out, s)
}
func autoConvert_v1_ConfigMapList_To_core_ConfigMapList(in *v1.ConfigMapList, out *core.ConfigMapList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 out.Items = *(*[]core.ConfigMap)(unsafe.Pointer(&in.Items))
 return nil
}
func Convert_v1_ConfigMapList_To_core_ConfigMapList(in *v1.ConfigMapList, out *core.ConfigMapList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_ConfigMapList_To_core_ConfigMapList(in, out, s)
}
func autoConvert_core_ConfigMapList_To_v1_ConfigMapList(in *core.ConfigMapList, out *v1.ConfigMapList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 out.Items = *(*[]v1.ConfigMap)(unsafe.Pointer(&in.Items))
 return nil
}
func Convert_core_ConfigMapList_To_v1_ConfigMapList(in *core.ConfigMapList, out *v1.ConfigMapList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_ConfigMapList_To_v1_ConfigMapList(in, out, s)
}
func autoConvert_v1_ConfigMapNodeConfigSource_To_core_ConfigMapNodeConfigSource(in *v1.ConfigMapNodeConfigSource, out *core.ConfigMapNodeConfigSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Namespace = in.Namespace
 out.Name = in.Name
 out.UID = types.UID(in.UID)
 out.ResourceVersion = in.ResourceVersion
 out.KubeletConfigKey = in.KubeletConfigKey
 return nil
}
func Convert_v1_ConfigMapNodeConfigSource_To_core_ConfigMapNodeConfigSource(in *v1.ConfigMapNodeConfigSource, out *core.ConfigMapNodeConfigSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_ConfigMapNodeConfigSource_To_core_ConfigMapNodeConfigSource(in, out, s)
}
func autoConvert_core_ConfigMapNodeConfigSource_To_v1_ConfigMapNodeConfigSource(in *core.ConfigMapNodeConfigSource, out *v1.ConfigMapNodeConfigSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Namespace = in.Namespace
 out.Name = in.Name
 out.UID = types.UID(in.UID)
 out.ResourceVersion = in.ResourceVersion
 out.KubeletConfigKey = in.KubeletConfigKey
 return nil
}
func Convert_core_ConfigMapNodeConfigSource_To_v1_ConfigMapNodeConfigSource(in *core.ConfigMapNodeConfigSource, out *v1.ConfigMapNodeConfigSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_ConfigMapNodeConfigSource_To_v1_ConfigMapNodeConfigSource(in, out, s)
}
func autoConvert_v1_ConfigMapProjection_To_core_ConfigMapProjection(in *v1.ConfigMapProjection, out *core.ConfigMapProjection, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := Convert_v1_LocalObjectReference_To_core_LocalObjectReference(&in.LocalObjectReference, &out.LocalObjectReference, s); err != nil {
  return err
 }
 out.Items = *(*[]core.KeyToPath)(unsafe.Pointer(&in.Items))
 out.Optional = (*bool)(unsafe.Pointer(in.Optional))
 return nil
}
func Convert_v1_ConfigMapProjection_To_core_ConfigMapProjection(in *v1.ConfigMapProjection, out *core.ConfigMapProjection, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_ConfigMapProjection_To_core_ConfigMapProjection(in, out, s)
}
func autoConvert_core_ConfigMapProjection_To_v1_ConfigMapProjection(in *core.ConfigMapProjection, out *v1.ConfigMapProjection, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := Convert_core_LocalObjectReference_To_v1_LocalObjectReference(&in.LocalObjectReference, &out.LocalObjectReference, s); err != nil {
  return err
 }
 out.Items = *(*[]v1.KeyToPath)(unsafe.Pointer(&in.Items))
 out.Optional = (*bool)(unsafe.Pointer(in.Optional))
 return nil
}
func Convert_core_ConfigMapProjection_To_v1_ConfigMapProjection(in *core.ConfigMapProjection, out *v1.ConfigMapProjection, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_ConfigMapProjection_To_v1_ConfigMapProjection(in, out, s)
}
func autoConvert_v1_ConfigMapVolumeSource_To_core_ConfigMapVolumeSource(in *v1.ConfigMapVolumeSource, out *core.ConfigMapVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := Convert_v1_LocalObjectReference_To_core_LocalObjectReference(&in.LocalObjectReference, &out.LocalObjectReference, s); err != nil {
  return err
 }
 out.Items = *(*[]core.KeyToPath)(unsafe.Pointer(&in.Items))
 out.DefaultMode = (*int32)(unsafe.Pointer(in.DefaultMode))
 out.Optional = (*bool)(unsafe.Pointer(in.Optional))
 return nil
}
func Convert_v1_ConfigMapVolumeSource_To_core_ConfigMapVolumeSource(in *v1.ConfigMapVolumeSource, out *core.ConfigMapVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_ConfigMapVolumeSource_To_core_ConfigMapVolumeSource(in, out, s)
}
func autoConvert_core_ConfigMapVolumeSource_To_v1_ConfigMapVolumeSource(in *core.ConfigMapVolumeSource, out *v1.ConfigMapVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := Convert_core_LocalObjectReference_To_v1_LocalObjectReference(&in.LocalObjectReference, &out.LocalObjectReference, s); err != nil {
  return err
 }
 out.Items = *(*[]v1.KeyToPath)(unsafe.Pointer(&in.Items))
 out.DefaultMode = (*int32)(unsafe.Pointer(in.DefaultMode))
 out.Optional = (*bool)(unsafe.Pointer(in.Optional))
 return nil
}
func Convert_core_ConfigMapVolumeSource_To_v1_ConfigMapVolumeSource(in *core.ConfigMapVolumeSource, out *v1.ConfigMapVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_ConfigMapVolumeSource_To_v1_ConfigMapVolumeSource(in, out, s)
}
func autoConvert_v1_Container_To_core_Container(in *v1.Container, out *core.Container, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Name = in.Name
 out.Image = in.Image
 out.Command = *(*[]string)(unsafe.Pointer(&in.Command))
 out.Args = *(*[]string)(unsafe.Pointer(&in.Args))
 out.WorkingDir = in.WorkingDir
 out.Ports = *(*[]core.ContainerPort)(unsafe.Pointer(&in.Ports))
 out.EnvFrom = *(*[]core.EnvFromSource)(unsafe.Pointer(&in.EnvFrom))
 out.Env = *(*[]core.EnvVar)(unsafe.Pointer(&in.Env))
 if err := Convert_v1_ResourceRequirements_To_core_ResourceRequirements(&in.Resources, &out.Resources, s); err != nil {
  return err
 }
 out.VolumeMounts = *(*[]core.VolumeMount)(unsafe.Pointer(&in.VolumeMounts))
 out.VolumeDevices = *(*[]core.VolumeDevice)(unsafe.Pointer(&in.VolumeDevices))
 out.LivenessProbe = (*core.Probe)(unsafe.Pointer(in.LivenessProbe))
 out.ReadinessProbe = (*core.Probe)(unsafe.Pointer(in.ReadinessProbe))
 out.Lifecycle = (*core.Lifecycle)(unsafe.Pointer(in.Lifecycle))
 out.TerminationMessagePath = in.TerminationMessagePath
 out.TerminationMessagePolicy = core.TerminationMessagePolicy(in.TerminationMessagePolicy)
 out.ImagePullPolicy = core.PullPolicy(in.ImagePullPolicy)
 if in.SecurityContext != nil {
  in, out := &in.SecurityContext, &out.SecurityContext
  *out = new(core.SecurityContext)
  if err := Convert_v1_SecurityContext_To_core_SecurityContext(*in, *out, s); err != nil {
   return err
  }
 } else {
  out.SecurityContext = nil
 }
 out.Stdin = in.Stdin
 out.StdinOnce = in.StdinOnce
 out.TTY = in.TTY
 return nil
}
func Convert_v1_Container_To_core_Container(in *v1.Container, out *core.Container, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_Container_To_core_Container(in, out, s)
}
func autoConvert_core_Container_To_v1_Container(in *core.Container, out *v1.Container, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Name = in.Name
 out.Image = in.Image
 out.Command = *(*[]string)(unsafe.Pointer(&in.Command))
 out.Args = *(*[]string)(unsafe.Pointer(&in.Args))
 out.WorkingDir = in.WorkingDir
 out.Ports = *(*[]v1.ContainerPort)(unsafe.Pointer(&in.Ports))
 out.EnvFrom = *(*[]v1.EnvFromSource)(unsafe.Pointer(&in.EnvFrom))
 out.Env = *(*[]v1.EnvVar)(unsafe.Pointer(&in.Env))
 if err := Convert_core_ResourceRequirements_To_v1_ResourceRequirements(&in.Resources, &out.Resources, s); err != nil {
  return err
 }
 out.VolumeMounts = *(*[]v1.VolumeMount)(unsafe.Pointer(&in.VolumeMounts))
 out.VolumeDevices = *(*[]v1.VolumeDevice)(unsafe.Pointer(&in.VolumeDevices))
 out.LivenessProbe = (*v1.Probe)(unsafe.Pointer(in.LivenessProbe))
 out.ReadinessProbe = (*v1.Probe)(unsafe.Pointer(in.ReadinessProbe))
 out.Lifecycle = (*v1.Lifecycle)(unsafe.Pointer(in.Lifecycle))
 out.TerminationMessagePath = in.TerminationMessagePath
 out.TerminationMessagePolicy = v1.TerminationMessagePolicy(in.TerminationMessagePolicy)
 out.ImagePullPolicy = v1.PullPolicy(in.ImagePullPolicy)
 if in.SecurityContext != nil {
  in, out := &in.SecurityContext, &out.SecurityContext
  *out = new(v1.SecurityContext)
  if err := Convert_core_SecurityContext_To_v1_SecurityContext(*in, *out, s); err != nil {
   return err
  }
 } else {
  out.SecurityContext = nil
 }
 out.Stdin = in.Stdin
 out.StdinOnce = in.StdinOnce
 out.TTY = in.TTY
 return nil
}
func Convert_core_Container_To_v1_Container(in *core.Container, out *v1.Container, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_Container_To_v1_Container(in, out, s)
}
func autoConvert_v1_ContainerImage_To_core_ContainerImage(in *v1.ContainerImage, out *core.ContainerImage, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Names = *(*[]string)(unsafe.Pointer(&in.Names))
 out.SizeBytes = in.SizeBytes
 return nil
}
func Convert_v1_ContainerImage_To_core_ContainerImage(in *v1.ContainerImage, out *core.ContainerImage, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_ContainerImage_To_core_ContainerImage(in, out, s)
}
func autoConvert_core_ContainerImage_To_v1_ContainerImage(in *core.ContainerImage, out *v1.ContainerImage, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Names = *(*[]string)(unsafe.Pointer(&in.Names))
 out.SizeBytes = in.SizeBytes
 return nil
}
func Convert_core_ContainerImage_To_v1_ContainerImage(in *core.ContainerImage, out *v1.ContainerImage, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_ContainerImage_To_v1_ContainerImage(in, out, s)
}
func autoConvert_v1_ContainerPort_To_core_ContainerPort(in *v1.ContainerPort, out *core.ContainerPort, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Name = in.Name
 out.HostPort = in.HostPort
 out.ContainerPort = in.ContainerPort
 out.Protocol = core.Protocol(in.Protocol)
 out.HostIP = in.HostIP
 return nil
}
func Convert_v1_ContainerPort_To_core_ContainerPort(in *v1.ContainerPort, out *core.ContainerPort, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_ContainerPort_To_core_ContainerPort(in, out, s)
}
func autoConvert_core_ContainerPort_To_v1_ContainerPort(in *core.ContainerPort, out *v1.ContainerPort, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Name = in.Name
 out.HostPort = in.HostPort
 out.ContainerPort = in.ContainerPort
 out.Protocol = v1.Protocol(in.Protocol)
 out.HostIP = in.HostIP
 return nil
}
func Convert_core_ContainerPort_To_v1_ContainerPort(in *core.ContainerPort, out *v1.ContainerPort, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_ContainerPort_To_v1_ContainerPort(in, out, s)
}
func autoConvert_v1_ContainerState_To_core_ContainerState(in *v1.ContainerState, out *core.ContainerState, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Waiting = (*core.ContainerStateWaiting)(unsafe.Pointer(in.Waiting))
 out.Running = (*core.ContainerStateRunning)(unsafe.Pointer(in.Running))
 out.Terminated = (*core.ContainerStateTerminated)(unsafe.Pointer(in.Terminated))
 return nil
}
func Convert_v1_ContainerState_To_core_ContainerState(in *v1.ContainerState, out *core.ContainerState, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_ContainerState_To_core_ContainerState(in, out, s)
}
func autoConvert_core_ContainerState_To_v1_ContainerState(in *core.ContainerState, out *v1.ContainerState, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Waiting = (*v1.ContainerStateWaiting)(unsafe.Pointer(in.Waiting))
 out.Running = (*v1.ContainerStateRunning)(unsafe.Pointer(in.Running))
 out.Terminated = (*v1.ContainerStateTerminated)(unsafe.Pointer(in.Terminated))
 return nil
}
func Convert_core_ContainerState_To_v1_ContainerState(in *core.ContainerState, out *v1.ContainerState, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_ContainerState_To_v1_ContainerState(in, out, s)
}
func autoConvert_v1_ContainerStateRunning_To_core_ContainerStateRunning(in *v1.ContainerStateRunning, out *core.ContainerStateRunning, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.StartedAt = in.StartedAt
 return nil
}
func Convert_v1_ContainerStateRunning_To_core_ContainerStateRunning(in *v1.ContainerStateRunning, out *core.ContainerStateRunning, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_ContainerStateRunning_To_core_ContainerStateRunning(in, out, s)
}
func autoConvert_core_ContainerStateRunning_To_v1_ContainerStateRunning(in *core.ContainerStateRunning, out *v1.ContainerStateRunning, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.StartedAt = in.StartedAt
 return nil
}
func Convert_core_ContainerStateRunning_To_v1_ContainerStateRunning(in *core.ContainerStateRunning, out *v1.ContainerStateRunning, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_ContainerStateRunning_To_v1_ContainerStateRunning(in, out, s)
}
func autoConvert_v1_ContainerStateTerminated_To_core_ContainerStateTerminated(in *v1.ContainerStateTerminated, out *core.ContainerStateTerminated, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ExitCode = in.ExitCode
 out.Signal = in.Signal
 out.Reason = in.Reason
 out.Message = in.Message
 out.StartedAt = in.StartedAt
 out.FinishedAt = in.FinishedAt
 out.ContainerID = in.ContainerID
 return nil
}
func Convert_v1_ContainerStateTerminated_To_core_ContainerStateTerminated(in *v1.ContainerStateTerminated, out *core.ContainerStateTerminated, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_ContainerStateTerminated_To_core_ContainerStateTerminated(in, out, s)
}
func autoConvert_core_ContainerStateTerminated_To_v1_ContainerStateTerminated(in *core.ContainerStateTerminated, out *v1.ContainerStateTerminated, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ExitCode = in.ExitCode
 out.Signal = in.Signal
 out.Reason = in.Reason
 out.Message = in.Message
 out.StartedAt = in.StartedAt
 out.FinishedAt = in.FinishedAt
 out.ContainerID = in.ContainerID
 return nil
}
func Convert_core_ContainerStateTerminated_To_v1_ContainerStateTerminated(in *core.ContainerStateTerminated, out *v1.ContainerStateTerminated, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_ContainerStateTerminated_To_v1_ContainerStateTerminated(in, out, s)
}
func autoConvert_v1_ContainerStateWaiting_To_core_ContainerStateWaiting(in *v1.ContainerStateWaiting, out *core.ContainerStateWaiting, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Reason = in.Reason
 out.Message = in.Message
 return nil
}
func Convert_v1_ContainerStateWaiting_To_core_ContainerStateWaiting(in *v1.ContainerStateWaiting, out *core.ContainerStateWaiting, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_ContainerStateWaiting_To_core_ContainerStateWaiting(in, out, s)
}
func autoConvert_core_ContainerStateWaiting_To_v1_ContainerStateWaiting(in *core.ContainerStateWaiting, out *v1.ContainerStateWaiting, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Reason = in.Reason
 out.Message = in.Message
 return nil
}
func Convert_core_ContainerStateWaiting_To_v1_ContainerStateWaiting(in *core.ContainerStateWaiting, out *v1.ContainerStateWaiting, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_ContainerStateWaiting_To_v1_ContainerStateWaiting(in, out, s)
}
func autoConvert_v1_ContainerStatus_To_core_ContainerStatus(in *v1.ContainerStatus, out *core.ContainerStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Name = in.Name
 if err := Convert_v1_ContainerState_To_core_ContainerState(&in.State, &out.State, s); err != nil {
  return err
 }
 if err := Convert_v1_ContainerState_To_core_ContainerState(&in.LastTerminationState, &out.LastTerminationState, s); err != nil {
  return err
 }
 out.Ready = in.Ready
 out.RestartCount = in.RestartCount
 out.Image = in.Image
 out.ImageID = in.ImageID
 out.ContainerID = in.ContainerID
 return nil
}
func Convert_v1_ContainerStatus_To_core_ContainerStatus(in *v1.ContainerStatus, out *core.ContainerStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_ContainerStatus_To_core_ContainerStatus(in, out, s)
}
func autoConvert_core_ContainerStatus_To_v1_ContainerStatus(in *core.ContainerStatus, out *v1.ContainerStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Name = in.Name
 if err := Convert_core_ContainerState_To_v1_ContainerState(&in.State, &out.State, s); err != nil {
  return err
 }
 if err := Convert_core_ContainerState_To_v1_ContainerState(&in.LastTerminationState, &out.LastTerminationState, s); err != nil {
  return err
 }
 out.Ready = in.Ready
 out.RestartCount = in.RestartCount
 out.Image = in.Image
 out.ImageID = in.ImageID
 out.ContainerID = in.ContainerID
 return nil
}
func Convert_core_ContainerStatus_To_v1_ContainerStatus(in *core.ContainerStatus, out *v1.ContainerStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_ContainerStatus_To_v1_ContainerStatus(in, out, s)
}
func autoConvert_v1_DaemonEndpoint_To_core_DaemonEndpoint(in *v1.DaemonEndpoint, out *core.DaemonEndpoint, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Port = in.Port
 return nil
}
func Convert_v1_DaemonEndpoint_To_core_DaemonEndpoint(in *v1.DaemonEndpoint, out *core.DaemonEndpoint, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_DaemonEndpoint_To_core_DaemonEndpoint(in, out, s)
}
func autoConvert_core_DaemonEndpoint_To_v1_DaemonEndpoint(in *core.DaemonEndpoint, out *v1.DaemonEndpoint, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Port = in.Port
 return nil
}
func Convert_core_DaemonEndpoint_To_v1_DaemonEndpoint(in *core.DaemonEndpoint, out *v1.DaemonEndpoint, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_DaemonEndpoint_To_v1_DaemonEndpoint(in, out, s)
}
func autoConvert_v1_DownwardAPIProjection_To_core_DownwardAPIProjection(in *v1.DownwardAPIProjection, out *core.DownwardAPIProjection, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Items = *(*[]core.DownwardAPIVolumeFile)(unsafe.Pointer(&in.Items))
 return nil
}
func Convert_v1_DownwardAPIProjection_To_core_DownwardAPIProjection(in *v1.DownwardAPIProjection, out *core.DownwardAPIProjection, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_DownwardAPIProjection_To_core_DownwardAPIProjection(in, out, s)
}
func autoConvert_core_DownwardAPIProjection_To_v1_DownwardAPIProjection(in *core.DownwardAPIProjection, out *v1.DownwardAPIProjection, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Items = *(*[]v1.DownwardAPIVolumeFile)(unsafe.Pointer(&in.Items))
 return nil
}
func Convert_core_DownwardAPIProjection_To_v1_DownwardAPIProjection(in *core.DownwardAPIProjection, out *v1.DownwardAPIProjection, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_DownwardAPIProjection_To_v1_DownwardAPIProjection(in, out, s)
}
func autoConvert_v1_DownwardAPIVolumeFile_To_core_DownwardAPIVolumeFile(in *v1.DownwardAPIVolumeFile, out *core.DownwardAPIVolumeFile, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Path = in.Path
 out.FieldRef = (*core.ObjectFieldSelector)(unsafe.Pointer(in.FieldRef))
 out.ResourceFieldRef = (*core.ResourceFieldSelector)(unsafe.Pointer(in.ResourceFieldRef))
 out.Mode = (*int32)(unsafe.Pointer(in.Mode))
 return nil
}
func Convert_v1_DownwardAPIVolumeFile_To_core_DownwardAPIVolumeFile(in *v1.DownwardAPIVolumeFile, out *core.DownwardAPIVolumeFile, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_DownwardAPIVolumeFile_To_core_DownwardAPIVolumeFile(in, out, s)
}
func autoConvert_core_DownwardAPIVolumeFile_To_v1_DownwardAPIVolumeFile(in *core.DownwardAPIVolumeFile, out *v1.DownwardAPIVolumeFile, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Path = in.Path
 out.FieldRef = (*v1.ObjectFieldSelector)(unsafe.Pointer(in.FieldRef))
 out.ResourceFieldRef = (*v1.ResourceFieldSelector)(unsafe.Pointer(in.ResourceFieldRef))
 out.Mode = (*int32)(unsafe.Pointer(in.Mode))
 return nil
}
func Convert_core_DownwardAPIVolumeFile_To_v1_DownwardAPIVolumeFile(in *core.DownwardAPIVolumeFile, out *v1.DownwardAPIVolumeFile, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_DownwardAPIVolumeFile_To_v1_DownwardAPIVolumeFile(in, out, s)
}
func autoConvert_v1_DownwardAPIVolumeSource_To_core_DownwardAPIVolumeSource(in *v1.DownwardAPIVolumeSource, out *core.DownwardAPIVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Items = *(*[]core.DownwardAPIVolumeFile)(unsafe.Pointer(&in.Items))
 out.DefaultMode = (*int32)(unsafe.Pointer(in.DefaultMode))
 return nil
}
func Convert_v1_DownwardAPIVolumeSource_To_core_DownwardAPIVolumeSource(in *v1.DownwardAPIVolumeSource, out *core.DownwardAPIVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_DownwardAPIVolumeSource_To_core_DownwardAPIVolumeSource(in, out, s)
}
func autoConvert_core_DownwardAPIVolumeSource_To_v1_DownwardAPIVolumeSource(in *core.DownwardAPIVolumeSource, out *v1.DownwardAPIVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Items = *(*[]v1.DownwardAPIVolumeFile)(unsafe.Pointer(&in.Items))
 out.DefaultMode = (*int32)(unsafe.Pointer(in.DefaultMode))
 return nil
}
func Convert_core_DownwardAPIVolumeSource_To_v1_DownwardAPIVolumeSource(in *core.DownwardAPIVolumeSource, out *v1.DownwardAPIVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_DownwardAPIVolumeSource_To_v1_DownwardAPIVolumeSource(in, out, s)
}
func autoConvert_v1_EmptyDirVolumeSource_To_core_EmptyDirVolumeSource(in *v1.EmptyDirVolumeSource, out *core.EmptyDirVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Medium = core.StorageMedium(in.Medium)
 out.SizeLimit = (*resource.Quantity)(unsafe.Pointer(in.SizeLimit))
 return nil
}
func Convert_v1_EmptyDirVolumeSource_To_core_EmptyDirVolumeSource(in *v1.EmptyDirVolumeSource, out *core.EmptyDirVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_EmptyDirVolumeSource_To_core_EmptyDirVolumeSource(in, out, s)
}
func autoConvert_core_EmptyDirVolumeSource_To_v1_EmptyDirVolumeSource(in *core.EmptyDirVolumeSource, out *v1.EmptyDirVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Medium = v1.StorageMedium(in.Medium)
 out.SizeLimit = (*resource.Quantity)(unsafe.Pointer(in.SizeLimit))
 return nil
}
func Convert_core_EmptyDirVolumeSource_To_v1_EmptyDirVolumeSource(in *core.EmptyDirVolumeSource, out *v1.EmptyDirVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_EmptyDirVolumeSource_To_v1_EmptyDirVolumeSource(in, out, s)
}
func autoConvert_v1_EndpointAddress_To_core_EndpointAddress(in *v1.EndpointAddress, out *core.EndpointAddress, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.IP = in.IP
 out.Hostname = in.Hostname
 out.NodeName = (*string)(unsafe.Pointer(in.NodeName))
 out.TargetRef = (*core.ObjectReference)(unsafe.Pointer(in.TargetRef))
 return nil
}
func Convert_v1_EndpointAddress_To_core_EndpointAddress(in *v1.EndpointAddress, out *core.EndpointAddress, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_EndpointAddress_To_core_EndpointAddress(in, out, s)
}
func autoConvert_core_EndpointAddress_To_v1_EndpointAddress(in *core.EndpointAddress, out *v1.EndpointAddress, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.IP = in.IP
 out.Hostname = in.Hostname
 out.NodeName = (*string)(unsafe.Pointer(in.NodeName))
 out.TargetRef = (*v1.ObjectReference)(unsafe.Pointer(in.TargetRef))
 return nil
}
func Convert_core_EndpointAddress_To_v1_EndpointAddress(in *core.EndpointAddress, out *v1.EndpointAddress, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_EndpointAddress_To_v1_EndpointAddress(in, out, s)
}
func autoConvert_v1_EndpointPort_To_core_EndpointPort(in *v1.EndpointPort, out *core.EndpointPort, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Name = in.Name
 out.Port = in.Port
 out.Protocol = core.Protocol(in.Protocol)
 return nil
}
func Convert_v1_EndpointPort_To_core_EndpointPort(in *v1.EndpointPort, out *core.EndpointPort, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_EndpointPort_To_core_EndpointPort(in, out, s)
}
func autoConvert_core_EndpointPort_To_v1_EndpointPort(in *core.EndpointPort, out *v1.EndpointPort, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Name = in.Name
 out.Port = in.Port
 out.Protocol = v1.Protocol(in.Protocol)
 return nil
}
func Convert_core_EndpointPort_To_v1_EndpointPort(in *core.EndpointPort, out *v1.EndpointPort, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_EndpointPort_To_v1_EndpointPort(in, out, s)
}
func autoConvert_v1_EndpointSubset_To_core_EndpointSubset(in *v1.EndpointSubset, out *core.EndpointSubset, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Addresses = *(*[]core.EndpointAddress)(unsafe.Pointer(&in.Addresses))
 out.NotReadyAddresses = *(*[]core.EndpointAddress)(unsafe.Pointer(&in.NotReadyAddresses))
 out.Ports = *(*[]core.EndpointPort)(unsafe.Pointer(&in.Ports))
 return nil
}
func Convert_v1_EndpointSubset_To_core_EndpointSubset(in *v1.EndpointSubset, out *core.EndpointSubset, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_EndpointSubset_To_core_EndpointSubset(in, out, s)
}
func autoConvert_core_EndpointSubset_To_v1_EndpointSubset(in *core.EndpointSubset, out *v1.EndpointSubset, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Addresses = *(*[]v1.EndpointAddress)(unsafe.Pointer(&in.Addresses))
 out.NotReadyAddresses = *(*[]v1.EndpointAddress)(unsafe.Pointer(&in.NotReadyAddresses))
 out.Ports = *(*[]v1.EndpointPort)(unsafe.Pointer(&in.Ports))
 return nil
}
func Convert_core_EndpointSubset_To_v1_EndpointSubset(in *core.EndpointSubset, out *v1.EndpointSubset, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_EndpointSubset_To_v1_EndpointSubset(in, out, s)
}
func autoConvert_v1_Endpoints_To_core_Endpoints(in *v1.Endpoints, out *core.Endpoints, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 out.Subsets = *(*[]core.EndpointSubset)(unsafe.Pointer(&in.Subsets))
 return nil
}
func Convert_v1_Endpoints_To_core_Endpoints(in *v1.Endpoints, out *core.Endpoints, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_Endpoints_To_core_Endpoints(in, out, s)
}
func autoConvert_core_Endpoints_To_v1_Endpoints(in *core.Endpoints, out *v1.Endpoints, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 out.Subsets = *(*[]v1.EndpointSubset)(unsafe.Pointer(&in.Subsets))
 return nil
}
func Convert_core_Endpoints_To_v1_Endpoints(in *core.Endpoints, out *v1.Endpoints, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_Endpoints_To_v1_Endpoints(in, out, s)
}
func autoConvert_v1_EndpointsList_To_core_EndpointsList(in *v1.EndpointsList, out *core.EndpointsList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 out.Items = *(*[]core.Endpoints)(unsafe.Pointer(&in.Items))
 return nil
}
func Convert_v1_EndpointsList_To_core_EndpointsList(in *v1.EndpointsList, out *core.EndpointsList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_EndpointsList_To_core_EndpointsList(in, out, s)
}
func autoConvert_core_EndpointsList_To_v1_EndpointsList(in *core.EndpointsList, out *v1.EndpointsList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 out.Items = *(*[]v1.Endpoints)(unsafe.Pointer(&in.Items))
 return nil
}
func Convert_core_EndpointsList_To_v1_EndpointsList(in *core.EndpointsList, out *v1.EndpointsList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_EndpointsList_To_v1_EndpointsList(in, out, s)
}
func autoConvert_v1_EnvFromSource_To_core_EnvFromSource(in *v1.EnvFromSource, out *core.EnvFromSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Prefix = in.Prefix
 out.ConfigMapRef = (*core.ConfigMapEnvSource)(unsafe.Pointer(in.ConfigMapRef))
 out.SecretRef = (*core.SecretEnvSource)(unsafe.Pointer(in.SecretRef))
 return nil
}
func Convert_v1_EnvFromSource_To_core_EnvFromSource(in *v1.EnvFromSource, out *core.EnvFromSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_EnvFromSource_To_core_EnvFromSource(in, out, s)
}
func autoConvert_core_EnvFromSource_To_v1_EnvFromSource(in *core.EnvFromSource, out *v1.EnvFromSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Prefix = in.Prefix
 out.ConfigMapRef = (*v1.ConfigMapEnvSource)(unsafe.Pointer(in.ConfigMapRef))
 out.SecretRef = (*v1.SecretEnvSource)(unsafe.Pointer(in.SecretRef))
 return nil
}
func Convert_core_EnvFromSource_To_v1_EnvFromSource(in *core.EnvFromSource, out *v1.EnvFromSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_EnvFromSource_To_v1_EnvFromSource(in, out, s)
}
func autoConvert_v1_EnvVar_To_core_EnvVar(in *v1.EnvVar, out *core.EnvVar, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Name = in.Name
 out.Value = in.Value
 out.ValueFrom = (*core.EnvVarSource)(unsafe.Pointer(in.ValueFrom))
 return nil
}
func Convert_v1_EnvVar_To_core_EnvVar(in *v1.EnvVar, out *core.EnvVar, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_EnvVar_To_core_EnvVar(in, out, s)
}
func autoConvert_core_EnvVar_To_v1_EnvVar(in *core.EnvVar, out *v1.EnvVar, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Name = in.Name
 out.Value = in.Value
 out.ValueFrom = (*v1.EnvVarSource)(unsafe.Pointer(in.ValueFrom))
 return nil
}
func Convert_core_EnvVar_To_v1_EnvVar(in *core.EnvVar, out *v1.EnvVar, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_EnvVar_To_v1_EnvVar(in, out, s)
}
func autoConvert_v1_EnvVarSource_To_core_EnvVarSource(in *v1.EnvVarSource, out *core.EnvVarSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.FieldRef = (*core.ObjectFieldSelector)(unsafe.Pointer(in.FieldRef))
 out.ResourceFieldRef = (*core.ResourceFieldSelector)(unsafe.Pointer(in.ResourceFieldRef))
 out.ConfigMapKeyRef = (*core.ConfigMapKeySelector)(unsafe.Pointer(in.ConfigMapKeyRef))
 out.SecretKeyRef = (*core.SecretKeySelector)(unsafe.Pointer(in.SecretKeyRef))
 return nil
}
func Convert_v1_EnvVarSource_To_core_EnvVarSource(in *v1.EnvVarSource, out *core.EnvVarSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_EnvVarSource_To_core_EnvVarSource(in, out, s)
}
func autoConvert_core_EnvVarSource_To_v1_EnvVarSource(in *core.EnvVarSource, out *v1.EnvVarSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.FieldRef = (*v1.ObjectFieldSelector)(unsafe.Pointer(in.FieldRef))
 out.ResourceFieldRef = (*v1.ResourceFieldSelector)(unsafe.Pointer(in.ResourceFieldRef))
 out.ConfigMapKeyRef = (*v1.ConfigMapKeySelector)(unsafe.Pointer(in.ConfigMapKeyRef))
 out.SecretKeyRef = (*v1.SecretKeySelector)(unsafe.Pointer(in.SecretKeyRef))
 return nil
}
func Convert_core_EnvVarSource_To_v1_EnvVarSource(in *core.EnvVarSource, out *v1.EnvVarSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_EnvVarSource_To_v1_EnvVarSource(in, out, s)
}
func autoConvert_v1_Event_To_core_Event(in *v1.Event, out *core.Event, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_v1_ObjectReference_To_core_ObjectReference(&in.InvolvedObject, &out.InvolvedObject, s); err != nil {
  return err
 }
 out.Reason = in.Reason
 out.Message = in.Message
 if err := Convert_v1_EventSource_To_core_EventSource(&in.Source, &out.Source, s); err != nil {
  return err
 }
 out.FirstTimestamp = in.FirstTimestamp
 out.LastTimestamp = in.LastTimestamp
 out.Count = in.Count
 out.Type = in.Type
 out.EventTime = in.EventTime
 out.Series = (*core.EventSeries)(unsafe.Pointer(in.Series))
 out.Action = in.Action
 out.Related = (*core.ObjectReference)(unsafe.Pointer(in.Related))
 out.ReportingController = in.ReportingController
 out.ReportingInstance = in.ReportingInstance
 return nil
}
func Convert_v1_Event_To_core_Event(in *v1.Event, out *core.Event, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_Event_To_core_Event(in, out, s)
}
func autoConvert_core_Event_To_v1_Event(in *core.Event, out *v1.Event, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_core_ObjectReference_To_v1_ObjectReference(&in.InvolvedObject, &out.InvolvedObject, s); err != nil {
  return err
 }
 out.Reason = in.Reason
 out.Message = in.Message
 if err := Convert_core_EventSource_To_v1_EventSource(&in.Source, &out.Source, s); err != nil {
  return err
 }
 out.FirstTimestamp = in.FirstTimestamp
 out.LastTimestamp = in.LastTimestamp
 out.Count = in.Count
 out.Type = in.Type
 out.EventTime = in.EventTime
 out.Series = (*v1.EventSeries)(unsafe.Pointer(in.Series))
 out.Action = in.Action
 out.Related = (*v1.ObjectReference)(unsafe.Pointer(in.Related))
 out.ReportingController = in.ReportingController
 out.ReportingInstance = in.ReportingInstance
 return nil
}
func Convert_core_Event_To_v1_Event(in *core.Event, out *v1.Event, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_Event_To_v1_Event(in, out, s)
}
func autoConvert_v1_EventList_To_core_EventList(in *v1.EventList, out *core.EventList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 out.Items = *(*[]core.Event)(unsafe.Pointer(&in.Items))
 return nil
}
func Convert_v1_EventList_To_core_EventList(in *v1.EventList, out *core.EventList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_EventList_To_core_EventList(in, out, s)
}
func autoConvert_core_EventList_To_v1_EventList(in *core.EventList, out *v1.EventList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 out.Items = *(*[]v1.Event)(unsafe.Pointer(&in.Items))
 return nil
}
func Convert_core_EventList_To_v1_EventList(in *core.EventList, out *v1.EventList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_EventList_To_v1_EventList(in, out, s)
}
func autoConvert_v1_EventSeries_To_core_EventSeries(in *v1.EventSeries, out *core.EventSeries, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Count = in.Count
 out.LastObservedTime = in.LastObservedTime
 out.State = core.EventSeriesState(in.State)
 return nil
}
func Convert_v1_EventSeries_To_core_EventSeries(in *v1.EventSeries, out *core.EventSeries, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_EventSeries_To_core_EventSeries(in, out, s)
}
func autoConvert_core_EventSeries_To_v1_EventSeries(in *core.EventSeries, out *v1.EventSeries, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Count = in.Count
 out.LastObservedTime = in.LastObservedTime
 out.State = v1.EventSeriesState(in.State)
 return nil
}
func Convert_core_EventSeries_To_v1_EventSeries(in *core.EventSeries, out *v1.EventSeries, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_EventSeries_To_v1_EventSeries(in, out, s)
}
func autoConvert_v1_EventSource_To_core_EventSource(in *v1.EventSource, out *core.EventSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Component = in.Component
 out.Host = in.Host
 return nil
}
func Convert_v1_EventSource_To_core_EventSource(in *v1.EventSource, out *core.EventSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_EventSource_To_core_EventSource(in, out, s)
}
func autoConvert_core_EventSource_To_v1_EventSource(in *core.EventSource, out *v1.EventSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Component = in.Component
 out.Host = in.Host
 return nil
}
func Convert_core_EventSource_To_v1_EventSource(in *core.EventSource, out *v1.EventSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_EventSource_To_v1_EventSource(in, out, s)
}
func autoConvert_v1_ExecAction_To_core_ExecAction(in *v1.ExecAction, out *core.ExecAction, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Command = *(*[]string)(unsafe.Pointer(&in.Command))
 return nil
}
func Convert_v1_ExecAction_To_core_ExecAction(in *v1.ExecAction, out *core.ExecAction, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_ExecAction_To_core_ExecAction(in, out, s)
}
func autoConvert_core_ExecAction_To_v1_ExecAction(in *core.ExecAction, out *v1.ExecAction, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Command = *(*[]string)(unsafe.Pointer(&in.Command))
 return nil
}
func Convert_core_ExecAction_To_v1_ExecAction(in *core.ExecAction, out *v1.ExecAction, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_ExecAction_To_v1_ExecAction(in, out, s)
}
func autoConvert_v1_FCVolumeSource_To_core_FCVolumeSource(in *v1.FCVolumeSource, out *core.FCVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.TargetWWNs = *(*[]string)(unsafe.Pointer(&in.TargetWWNs))
 out.Lun = (*int32)(unsafe.Pointer(in.Lun))
 out.FSType = in.FSType
 out.ReadOnly = in.ReadOnly
 out.WWIDs = *(*[]string)(unsafe.Pointer(&in.WWIDs))
 return nil
}
func Convert_v1_FCVolumeSource_To_core_FCVolumeSource(in *v1.FCVolumeSource, out *core.FCVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_FCVolumeSource_To_core_FCVolumeSource(in, out, s)
}
func autoConvert_core_FCVolumeSource_To_v1_FCVolumeSource(in *core.FCVolumeSource, out *v1.FCVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.TargetWWNs = *(*[]string)(unsafe.Pointer(&in.TargetWWNs))
 out.Lun = (*int32)(unsafe.Pointer(in.Lun))
 out.FSType = in.FSType
 out.ReadOnly = in.ReadOnly
 out.WWIDs = *(*[]string)(unsafe.Pointer(&in.WWIDs))
 return nil
}
func Convert_core_FCVolumeSource_To_v1_FCVolumeSource(in *core.FCVolumeSource, out *v1.FCVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_FCVolumeSource_To_v1_FCVolumeSource(in, out, s)
}
func autoConvert_v1_FlexPersistentVolumeSource_To_core_FlexPersistentVolumeSource(in *v1.FlexPersistentVolumeSource, out *core.FlexPersistentVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Driver = in.Driver
 out.FSType = in.FSType
 out.SecretRef = (*core.SecretReference)(unsafe.Pointer(in.SecretRef))
 out.ReadOnly = in.ReadOnly
 out.Options = *(*map[string]string)(unsafe.Pointer(&in.Options))
 return nil
}
func Convert_v1_FlexPersistentVolumeSource_To_core_FlexPersistentVolumeSource(in *v1.FlexPersistentVolumeSource, out *core.FlexPersistentVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_FlexPersistentVolumeSource_To_core_FlexPersistentVolumeSource(in, out, s)
}
func autoConvert_core_FlexPersistentVolumeSource_To_v1_FlexPersistentVolumeSource(in *core.FlexPersistentVolumeSource, out *v1.FlexPersistentVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Driver = in.Driver
 out.FSType = in.FSType
 out.SecretRef = (*v1.SecretReference)(unsafe.Pointer(in.SecretRef))
 out.ReadOnly = in.ReadOnly
 out.Options = *(*map[string]string)(unsafe.Pointer(&in.Options))
 return nil
}
func Convert_core_FlexPersistentVolumeSource_To_v1_FlexPersistentVolumeSource(in *core.FlexPersistentVolumeSource, out *v1.FlexPersistentVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_FlexPersistentVolumeSource_To_v1_FlexPersistentVolumeSource(in, out, s)
}
func autoConvert_v1_FlexVolumeSource_To_core_FlexVolumeSource(in *v1.FlexVolumeSource, out *core.FlexVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Driver = in.Driver
 out.FSType = in.FSType
 out.SecretRef = (*core.LocalObjectReference)(unsafe.Pointer(in.SecretRef))
 out.ReadOnly = in.ReadOnly
 out.Options = *(*map[string]string)(unsafe.Pointer(&in.Options))
 return nil
}
func Convert_v1_FlexVolumeSource_To_core_FlexVolumeSource(in *v1.FlexVolumeSource, out *core.FlexVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_FlexVolumeSource_To_core_FlexVolumeSource(in, out, s)
}
func autoConvert_core_FlexVolumeSource_To_v1_FlexVolumeSource(in *core.FlexVolumeSource, out *v1.FlexVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Driver = in.Driver
 out.FSType = in.FSType
 out.SecretRef = (*v1.LocalObjectReference)(unsafe.Pointer(in.SecretRef))
 out.ReadOnly = in.ReadOnly
 out.Options = *(*map[string]string)(unsafe.Pointer(&in.Options))
 return nil
}
func Convert_core_FlexVolumeSource_To_v1_FlexVolumeSource(in *core.FlexVolumeSource, out *v1.FlexVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_FlexVolumeSource_To_v1_FlexVolumeSource(in, out, s)
}
func autoConvert_v1_FlockerVolumeSource_To_core_FlockerVolumeSource(in *v1.FlockerVolumeSource, out *core.FlockerVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.DatasetName = in.DatasetName
 out.DatasetUUID = in.DatasetUUID
 return nil
}
func Convert_v1_FlockerVolumeSource_To_core_FlockerVolumeSource(in *v1.FlockerVolumeSource, out *core.FlockerVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_FlockerVolumeSource_To_core_FlockerVolumeSource(in, out, s)
}
func autoConvert_core_FlockerVolumeSource_To_v1_FlockerVolumeSource(in *core.FlockerVolumeSource, out *v1.FlockerVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.DatasetName = in.DatasetName
 out.DatasetUUID = in.DatasetUUID
 return nil
}
func Convert_core_FlockerVolumeSource_To_v1_FlockerVolumeSource(in *core.FlockerVolumeSource, out *v1.FlockerVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_FlockerVolumeSource_To_v1_FlockerVolumeSource(in, out, s)
}
func autoConvert_v1_GCEPersistentDiskVolumeSource_To_core_GCEPersistentDiskVolumeSource(in *v1.GCEPersistentDiskVolumeSource, out *core.GCEPersistentDiskVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.PDName = in.PDName
 out.FSType = in.FSType
 out.Partition = in.Partition
 out.ReadOnly = in.ReadOnly
 return nil
}
func Convert_v1_GCEPersistentDiskVolumeSource_To_core_GCEPersistentDiskVolumeSource(in *v1.GCEPersistentDiskVolumeSource, out *core.GCEPersistentDiskVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_GCEPersistentDiskVolumeSource_To_core_GCEPersistentDiskVolumeSource(in, out, s)
}
func autoConvert_core_GCEPersistentDiskVolumeSource_To_v1_GCEPersistentDiskVolumeSource(in *core.GCEPersistentDiskVolumeSource, out *v1.GCEPersistentDiskVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.PDName = in.PDName
 out.FSType = in.FSType
 out.Partition = in.Partition
 out.ReadOnly = in.ReadOnly
 return nil
}
func Convert_core_GCEPersistentDiskVolumeSource_To_v1_GCEPersistentDiskVolumeSource(in *core.GCEPersistentDiskVolumeSource, out *v1.GCEPersistentDiskVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_GCEPersistentDiskVolumeSource_To_v1_GCEPersistentDiskVolumeSource(in, out, s)
}
func autoConvert_v1_GitRepoVolumeSource_To_core_GitRepoVolumeSource(in *v1.GitRepoVolumeSource, out *core.GitRepoVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Repository = in.Repository
 out.Revision = in.Revision
 out.Directory = in.Directory
 return nil
}
func Convert_v1_GitRepoVolumeSource_To_core_GitRepoVolumeSource(in *v1.GitRepoVolumeSource, out *core.GitRepoVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_GitRepoVolumeSource_To_core_GitRepoVolumeSource(in, out, s)
}
func autoConvert_core_GitRepoVolumeSource_To_v1_GitRepoVolumeSource(in *core.GitRepoVolumeSource, out *v1.GitRepoVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Repository = in.Repository
 out.Revision = in.Revision
 out.Directory = in.Directory
 return nil
}
func Convert_core_GitRepoVolumeSource_To_v1_GitRepoVolumeSource(in *core.GitRepoVolumeSource, out *v1.GitRepoVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_GitRepoVolumeSource_To_v1_GitRepoVolumeSource(in, out, s)
}
func autoConvert_v1_GlusterfsPersistentVolumeSource_To_core_GlusterfsPersistentVolumeSource(in *v1.GlusterfsPersistentVolumeSource, out *core.GlusterfsPersistentVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.EndpointsName = in.EndpointsName
 out.Path = in.Path
 out.ReadOnly = in.ReadOnly
 out.EndpointsNamespace = (*string)(unsafe.Pointer(in.EndpointsNamespace))
 return nil
}
func Convert_v1_GlusterfsPersistentVolumeSource_To_core_GlusterfsPersistentVolumeSource(in *v1.GlusterfsPersistentVolumeSource, out *core.GlusterfsPersistentVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_GlusterfsPersistentVolumeSource_To_core_GlusterfsPersistentVolumeSource(in, out, s)
}
func autoConvert_core_GlusterfsPersistentVolumeSource_To_v1_GlusterfsPersistentVolumeSource(in *core.GlusterfsPersistentVolumeSource, out *v1.GlusterfsPersistentVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.EndpointsName = in.EndpointsName
 out.Path = in.Path
 out.ReadOnly = in.ReadOnly
 out.EndpointsNamespace = (*string)(unsafe.Pointer(in.EndpointsNamespace))
 return nil
}
func Convert_core_GlusterfsPersistentVolumeSource_To_v1_GlusterfsPersistentVolumeSource(in *core.GlusterfsPersistentVolumeSource, out *v1.GlusterfsPersistentVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_GlusterfsPersistentVolumeSource_To_v1_GlusterfsPersistentVolumeSource(in, out, s)
}
func autoConvert_v1_GlusterfsVolumeSource_To_core_GlusterfsVolumeSource(in *v1.GlusterfsVolumeSource, out *core.GlusterfsVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.EndpointsName = in.EndpointsName
 out.Path = in.Path
 out.ReadOnly = in.ReadOnly
 return nil
}
func Convert_v1_GlusterfsVolumeSource_To_core_GlusterfsVolumeSource(in *v1.GlusterfsVolumeSource, out *core.GlusterfsVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_GlusterfsVolumeSource_To_core_GlusterfsVolumeSource(in, out, s)
}
func autoConvert_core_GlusterfsVolumeSource_To_v1_GlusterfsVolumeSource(in *core.GlusterfsVolumeSource, out *v1.GlusterfsVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.EndpointsName = in.EndpointsName
 out.Path = in.Path
 out.ReadOnly = in.ReadOnly
 return nil
}
func Convert_core_GlusterfsVolumeSource_To_v1_GlusterfsVolumeSource(in *core.GlusterfsVolumeSource, out *v1.GlusterfsVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_GlusterfsVolumeSource_To_v1_GlusterfsVolumeSource(in, out, s)
}
func autoConvert_v1_HTTPGetAction_To_core_HTTPGetAction(in *v1.HTTPGetAction, out *core.HTTPGetAction, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Path = in.Path
 out.Port = in.Port
 out.Host = in.Host
 out.Scheme = core.URIScheme(in.Scheme)
 out.HTTPHeaders = *(*[]core.HTTPHeader)(unsafe.Pointer(&in.HTTPHeaders))
 return nil
}
func Convert_v1_HTTPGetAction_To_core_HTTPGetAction(in *v1.HTTPGetAction, out *core.HTTPGetAction, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_HTTPGetAction_To_core_HTTPGetAction(in, out, s)
}
func autoConvert_core_HTTPGetAction_To_v1_HTTPGetAction(in *core.HTTPGetAction, out *v1.HTTPGetAction, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Path = in.Path
 out.Port = in.Port
 out.Host = in.Host
 out.Scheme = v1.URIScheme(in.Scheme)
 out.HTTPHeaders = *(*[]v1.HTTPHeader)(unsafe.Pointer(&in.HTTPHeaders))
 return nil
}
func Convert_core_HTTPGetAction_To_v1_HTTPGetAction(in *core.HTTPGetAction, out *v1.HTTPGetAction, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_HTTPGetAction_To_v1_HTTPGetAction(in, out, s)
}
func autoConvert_v1_HTTPHeader_To_core_HTTPHeader(in *v1.HTTPHeader, out *core.HTTPHeader, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Name = in.Name
 out.Value = in.Value
 return nil
}
func Convert_v1_HTTPHeader_To_core_HTTPHeader(in *v1.HTTPHeader, out *core.HTTPHeader, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_HTTPHeader_To_core_HTTPHeader(in, out, s)
}
func autoConvert_core_HTTPHeader_To_v1_HTTPHeader(in *core.HTTPHeader, out *v1.HTTPHeader, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Name = in.Name
 out.Value = in.Value
 return nil
}
func Convert_core_HTTPHeader_To_v1_HTTPHeader(in *core.HTTPHeader, out *v1.HTTPHeader, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_HTTPHeader_To_v1_HTTPHeader(in, out, s)
}
func autoConvert_v1_Handler_To_core_Handler(in *v1.Handler, out *core.Handler, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Exec = (*core.ExecAction)(unsafe.Pointer(in.Exec))
 out.HTTPGet = (*core.HTTPGetAction)(unsafe.Pointer(in.HTTPGet))
 out.TCPSocket = (*core.TCPSocketAction)(unsafe.Pointer(in.TCPSocket))
 return nil
}
func Convert_v1_Handler_To_core_Handler(in *v1.Handler, out *core.Handler, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_Handler_To_core_Handler(in, out, s)
}
func autoConvert_core_Handler_To_v1_Handler(in *core.Handler, out *v1.Handler, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Exec = (*v1.ExecAction)(unsafe.Pointer(in.Exec))
 out.HTTPGet = (*v1.HTTPGetAction)(unsafe.Pointer(in.HTTPGet))
 out.TCPSocket = (*v1.TCPSocketAction)(unsafe.Pointer(in.TCPSocket))
 return nil
}
func Convert_core_Handler_To_v1_Handler(in *core.Handler, out *v1.Handler, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_Handler_To_v1_Handler(in, out, s)
}
func autoConvert_v1_HostAlias_To_core_HostAlias(in *v1.HostAlias, out *core.HostAlias, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.IP = in.IP
 out.Hostnames = *(*[]string)(unsafe.Pointer(&in.Hostnames))
 return nil
}
func Convert_v1_HostAlias_To_core_HostAlias(in *v1.HostAlias, out *core.HostAlias, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_HostAlias_To_core_HostAlias(in, out, s)
}
func autoConvert_core_HostAlias_To_v1_HostAlias(in *core.HostAlias, out *v1.HostAlias, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.IP = in.IP
 out.Hostnames = *(*[]string)(unsafe.Pointer(&in.Hostnames))
 return nil
}
func Convert_core_HostAlias_To_v1_HostAlias(in *core.HostAlias, out *v1.HostAlias, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_HostAlias_To_v1_HostAlias(in, out, s)
}
func autoConvert_v1_HostPathVolumeSource_To_core_HostPathVolumeSource(in *v1.HostPathVolumeSource, out *core.HostPathVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Path = in.Path
 out.Type = (*core.HostPathType)(unsafe.Pointer(in.Type))
 return nil
}
func Convert_v1_HostPathVolumeSource_To_core_HostPathVolumeSource(in *v1.HostPathVolumeSource, out *core.HostPathVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_HostPathVolumeSource_To_core_HostPathVolumeSource(in, out, s)
}
func autoConvert_core_HostPathVolumeSource_To_v1_HostPathVolumeSource(in *core.HostPathVolumeSource, out *v1.HostPathVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Path = in.Path
 out.Type = (*v1.HostPathType)(unsafe.Pointer(in.Type))
 return nil
}
func Convert_core_HostPathVolumeSource_To_v1_HostPathVolumeSource(in *core.HostPathVolumeSource, out *v1.HostPathVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_HostPathVolumeSource_To_v1_HostPathVolumeSource(in, out, s)
}
func autoConvert_v1_ISCSIPersistentVolumeSource_To_core_ISCSIPersistentVolumeSource(in *v1.ISCSIPersistentVolumeSource, out *core.ISCSIPersistentVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.TargetPortal = in.TargetPortal
 out.IQN = in.IQN
 out.Lun = in.Lun
 out.ISCSIInterface = in.ISCSIInterface
 out.FSType = in.FSType
 out.ReadOnly = in.ReadOnly
 out.Portals = *(*[]string)(unsafe.Pointer(&in.Portals))
 out.DiscoveryCHAPAuth = in.DiscoveryCHAPAuth
 out.SessionCHAPAuth = in.SessionCHAPAuth
 out.SecretRef = (*core.SecretReference)(unsafe.Pointer(in.SecretRef))
 out.InitiatorName = (*string)(unsafe.Pointer(in.InitiatorName))
 return nil
}
func Convert_v1_ISCSIPersistentVolumeSource_To_core_ISCSIPersistentVolumeSource(in *v1.ISCSIPersistentVolumeSource, out *core.ISCSIPersistentVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_ISCSIPersistentVolumeSource_To_core_ISCSIPersistentVolumeSource(in, out, s)
}
func autoConvert_core_ISCSIPersistentVolumeSource_To_v1_ISCSIPersistentVolumeSource(in *core.ISCSIPersistentVolumeSource, out *v1.ISCSIPersistentVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.TargetPortal = in.TargetPortal
 out.IQN = in.IQN
 out.Lun = in.Lun
 out.ISCSIInterface = in.ISCSIInterface
 out.FSType = in.FSType
 out.ReadOnly = in.ReadOnly
 out.Portals = *(*[]string)(unsafe.Pointer(&in.Portals))
 out.DiscoveryCHAPAuth = in.DiscoveryCHAPAuth
 out.SessionCHAPAuth = in.SessionCHAPAuth
 out.SecretRef = (*v1.SecretReference)(unsafe.Pointer(in.SecretRef))
 out.InitiatorName = (*string)(unsafe.Pointer(in.InitiatorName))
 return nil
}
func Convert_core_ISCSIPersistentVolumeSource_To_v1_ISCSIPersistentVolumeSource(in *core.ISCSIPersistentVolumeSource, out *v1.ISCSIPersistentVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_ISCSIPersistentVolumeSource_To_v1_ISCSIPersistentVolumeSource(in, out, s)
}
func autoConvert_v1_ISCSIVolumeSource_To_core_ISCSIVolumeSource(in *v1.ISCSIVolumeSource, out *core.ISCSIVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.TargetPortal = in.TargetPortal
 out.IQN = in.IQN
 out.Lun = in.Lun
 out.ISCSIInterface = in.ISCSIInterface
 out.FSType = in.FSType
 out.ReadOnly = in.ReadOnly
 out.Portals = *(*[]string)(unsafe.Pointer(&in.Portals))
 out.DiscoveryCHAPAuth = in.DiscoveryCHAPAuth
 out.SessionCHAPAuth = in.SessionCHAPAuth
 out.SecretRef = (*core.LocalObjectReference)(unsafe.Pointer(in.SecretRef))
 out.InitiatorName = (*string)(unsafe.Pointer(in.InitiatorName))
 return nil
}
func Convert_v1_ISCSIVolumeSource_To_core_ISCSIVolumeSource(in *v1.ISCSIVolumeSource, out *core.ISCSIVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_ISCSIVolumeSource_To_core_ISCSIVolumeSource(in, out, s)
}
func autoConvert_core_ISCSIVolumeSource_To_v1_ISCSIVolumeSource(in *core.ISCSIVolumeSource, out *v1.ISCSIVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.TargetPortal = in.TargetPortal
 out.IQN = in.IQN
 out.Lun = in.Lun
 out.ISCSIInterface = in.ISCSIInterface
 out.FSType = in.FSType
 out.ReadOnly = in.ReadOnly
 out.Portals = *(*[]string)(unsafe.Pointer(&in.Portals))
 out.DiscoveryCHAPAuth = in.DiscoveryCHAPAuth
 out.SessionCHAPAuth = in.SessionCHAPAuth
 out.SecretRef = (*v1.LocalObjectReference)(unsafe.Pointer(in.SecretRef))
 out.InitiatorName = (*string)(unsafe.Pointer(in.InitiatorName))
 return nil
}
func Convert_core_ISCSIVolumeSource_To_v1_ISCSIVolumeSource(in *core.ISCSIVolumeSource, out *v1.ISCSIVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_ISCSIVolumeSource_To_v1_ISCSIVolumeSource(in, out, s)
}
func autoConvert_v1_KeyToPath_To_core_KeyToPath(in *v1.KeyToPath, out *core.KeyToPath, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Key = in.Key
 out.Path = in.Path
 out.Mode = (*int32)(unsafe.Pointer(in.Mode))
 return nil
}
func Convert_v1_KeyToPath_To_core_KeyToPath(in *v1.KeyToPath, out *core.KeyToPath, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_KeyToPath_To_core_KeyToPath(in, out, s)
}
func autoConvert_core_KeyToPath_To_v1_KeyToPath(in *core.KeyToPath, out *v1.KeyToPath, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Key = in.Key
 out.Path = in.Path
 out.Mode = (*int32)(unsafe.Pointer(in.Mode))
 return nil
}
func Convert_core_KeyToPath_To_v1_KeyToPath(in *core.KeyToPath, out *v1.KeyToPath, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_KeyToPath_To_v1_KeyToPath(in, out, s)
}
func autoConvert_v1_Lifecycle_To_core_Lifecycle(in *v1.Lifecycle, out *core.Lifecycle, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.PostStart = (*core.Handler)(unsafe.Pointer(in.PostStart))
 out.PreStop = (*core.Handler)(unsafe.Pointer(in.PreStop))
 return nil
}
func Convert_v1_Lifecycle_To_core_Lifecycle(in *v1.Lifecycle, out *core.Lifecycle, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_Lifecycle_To_core_Lifecycle(in, out, s)
}
func autoConvert_core_Lifecycle_To_v1_Lifecycle(in *core.Lifecycle, out *v1.Lifecycle, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.PostStart = (*v1.Handler)(unsafe.Pointer(in.PostStart))
 out.PreStop = (*v1.Handler)(unsafe.Pointer(in.PreStop))
 return nil
}
func Convert_core_Lifecycle_To_v1_Lifecycle(in *core.Lifecycle, out *v1.Lifecycle, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_Lifecycle_To_v1_Lifecycle(in, out, s)
}
func autoConvert_v1_LimitRange_To_core_LimitRange(in *v1.LimitRange, out *core.LimitRange, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_v1_LimitRangeSpec_To_core_LimitRangeSpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 return nil
}
func Convert_v1_LimitRange_To_core_LimitRange(in *v1.LimitRange, out *core.LimitRange, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_LimitRange_To_core_LimitRange(in, out, s)
}
func autoConvert_core_LimitRange_To_v1_LimitRange(in *core.LimitRange, out *v1.LimitRange, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_core_LimitRangeSpec_To_v1_LimitRangeSpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 return nil
}
func Convert_core_LimitRange_To_v1_LimitRange(in *core.LimitRange, out *v1.LimitRange, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_LimitRange_To_v1_LimitRange(in, out, s)
}
func autoConvert_v1_LimitRangeItem_To_core_LimitRangeItem(in *v1.LimitRangeItem, out *core.LimitRangeItem, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Type = core.LimitType(in.Type)
 out.Max = *(*core.ResourceList)(unsafe.Pointer(&in.Max))
 out.Min = *(*core.ResourceList)(unsafe.Pointer(&in.Min))
 out.Default = *(*core.ResourceList)(unsafe.Pointer(&in.Default))
 out.DefaultRequest = *(*core.ResourceList)(unsafe.Pointer(&in.DefaultRequest))
 out.MaxLimitRequestRatio = *(*core.ResourceList)(unsafe.Pointer(&in.MaxLimitRequestRatio))
 return nil
}
func Convert_v1_LimitRangeItem_To_core_LimitRangeItem(in *v1.LimitRangeItem, out *core.LimitRangeItem, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_LimitRangeItem_To_core_LimitRangeItem(in, out, s)
}
func autoConvert_core_LimitRangeItem_To_v1_LimitRangeItem(in *core.LimitRangeItem, out *v1.LimitRangeItem, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Type = v1.LimitType(in.Type)
 out.Max = *(*v1.ResourceList)(unsafe.Pointer(&in.Max))
 out.Min = *(*v1.ResourceList)(unsafe.Pointer(&in.Min))
 out.Default = *(*v1.ResourceList)(unsafe.Pointer(&in.Default))
 out.DefaultRequest = *(*v1.ResourceList)(unsafe.Pointer(&in.DefaultRequest))
 out.MaxLimitRequestRatio = *(*v1.ResourceList)(unsafe.Pointer(&in.MaxLimitRequestRatio))
 return nil
}
func Convert_core_LimitRangeItem_To_v1_LimitRangeItem(in *core.LimitRangeItem, out *v1.LimitRangeItem, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_LimitRangeItem_To_v1_LimitRangeItem(in, out, s)
}
func autoConvert_v1_LimitRangeList_To_core_LimitRangeList(in *v1.LimitRangeList, out *core.LimitRangeList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 out.Items = *(*[]core.LimitRange)(unsafe.Pointer(&in.Items))
 return nil
}
func Convert_v1_LimitRangeList_To_core_LimitRangeList(in *v1.LimitRangeList, out *core.LimitRangeList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_LimitRangeList_To_core_LimitRangeList(in, out, s)
}
func autoConvert_core_LimitRangeList_To_v1_LimitRangeList(in *core.LimitRangeList, out *v1.LimitRangeList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 out.Items = *(*[]v1.LimitRange)(unsafe.Pointer(&in.Items))
 return nil
}
func Convert_core_LimitRangeList_To_v1_LimitRangeList(in *core.LimitRangeList, out *v1.LimitRangeList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_LimitRangeList_To_v1_LimitRangeList(in, out, s)
}
func autoConvert_v1_LimitRangeSpec_To_core_LimitRangeSpec(in *v1.LimitRangeSpec, out *core.LimitRangeSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Limits = *(*[]core.LimitRangeItem)(unsafe.Pointer(&in.Limits))
 return nil
}
func Convert_v1_LimitRangeSpec_To_core_LimitRangeSpec(in *v1.LimitRangeSpec, out *core.LimitRangeSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_LimitRangeSpec_To_core_LimitRangeSpec(in, out, s)
}
func autoConvert_core_LimitRangeSpec_To_v1_LimitRangeSpec(in *core.LimitRangeSpec, out *v1.LimitRangeSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Limits = *(*[]v1.LimitRangeItem)(unsafe.Pointer(&in.Limits))
 return nil
}
func Convert_core_LimitRangeSpec_To_v1_LimitRangeSpec(in *core.LimitRangeSpec, out *v1.LimitRangeSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_LimitRangeSpec_To_v1_LimitRangeSpec(in, out, s)
}
func autoConvert_v1_List_To_core_List(in *v1.List, out *core.List, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]runtime.Object, len(*in))
  for i := range *in {
   if err := runtime.Convert_runtime_RawExtension_To_runtime_Object(&(*in)[i], &(*out)[i], s); err != nil {
    return err
   }
  }
 } else {
  out.Items = nil
 }
 return nil
}
func Convert_v1_List_To_core_List(in *v1.List, out *core.List, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_List_To_core_List(in, out, s)
}
func autoConvert_core_List_To_v1_List(in *core.List, out *v1.List, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]runtime.RawExtension, len(*in))
  for i := range *in {
   if err := runtime.Convert_runtime_Object_To_runtime_RawExtension(&(*in)[i], &(*out)[i], s); err != nil {
    return err
   }
  }
 } else {
  out.Items = nil
 }
 return nil
}
func Convert_core_List_To_v1_List(in *core.List, out *v1.List, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_List_To_v1_List(in, out, s)
}
func autoConvert_v1_LoadBalancerIngress_To_core_LoadBalancerIngress(in *v1.LoadBalancerIngress, out *core.LoadBalancerIngress, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.IP = in.IP
 out.Hostname = in.Hostname
 return nil
}
func Convert_v1_LoadBalancerIngress_To_core_LoadBalancerIngress(in *v1.LoadBalancerIngress, out *core.LoadBalancerIngress, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_LoadBalancerIngress_To_core_LoadBalancerIngress(in, out, s)
}
func autoConvert_core_LoadBalancerIngress_To_v1_LoadBalancerIngress(in *core.LoadBalancerIngress, out *v1.LoadBalancerIngress, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.IP = in.IP
 out.Hostname = in.Hostname
 return nil
}
func Convert_core_LoadBalancerIngress_To_v1_LoadBalancerIngress(in *core.LoadBalancerIngress, out *v1.LoadBalancerIngress, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_LoadBalancerIngress_To_v1_LoadBalancerIngress(in, out, s)
}
func autoConvert_v1_LoadBalancerStatus_To_core_LoadBalancerStatus(in *v1.LoadBalancerStatus, out *core.LoadBalancerStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Ingress = *(*[]core.LoadBalancerIngress)(unsafe.Pointer(&in.Ingress))
 return nil
}
func Convert_v1_LoadBalancerStatus_To_core_LoadBalancerStatus(in *v1.LoadBalancerStatus, out *core.LoadBalancerStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_LoadBalancerStatus_To_core_LoadBalancerStatus(in, out, s)
}
func autoConvert_core_LoadBalancerStatus_To_v1_LoadBalancerStatus(in *core.LoadBalancerStatus, out *v1.LoadBalancerStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Ingress = *(*[]v1.LoadBalancerIngress)(unsafe.Pointer(&in.Ingress))
 return nil
}
func Convert_core_LoadBalancerStatus_To_v1_LoadBalancerStatus(in *core.LoadBalancerStatus, out *v1.LoadBalancerStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_LoadBalancerStatus_To_v1_LoadBalancerStatus(in, out, s)
}
func autoConvert_v1_LocalObjectReference_To_core_LocalObjectReference(in *v1.LocalObjectReference, out *core.LocalObjectReference, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Name = in.Name
 return nil
}
func Convert_v1_LocalObjectReference_To_core_LocalObjectReference(in *v1.LocalObjectReference, out *core.LocalObjectReference, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_LocalObjectReference_To_core_LocalObjectReference(in, out, s)
}
func autoConvert_core_LocalObjectReference_To_v1_LocalObjectReference(in *core.LocalObjectReference, out *v1.LocalObjectReference, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Name = in.Name
 return nil
}
func Convert_core_LocalObjectReference_To_v1_LocalObjectReference(in *core.LocalObjectReference, out *v1.LocalObjectReference, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_LocalObjectReference_To_v1_LocalObjectReference(in, out, s)
}
func autoConvert_v1_LocalVolumeSource_To_core_LocalVolumeSource(in *v1.LocalVolumeSource, out *core.LocalVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Path = in.Path
 out.FSType = (*string)(unsafe.Pointer(in.FSType))
 return nil
}
func Convert_v1_LocalVolumeSource_To_core_LocalVolumeSource(in *v1.LocalVolumeSource, out *core.LocalVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_LocalVolumeSource_To_core_LocalVolumeSource(in, out, s)
}
func autoConvert_core_LocalVolumeSource_To_v1_LocalVolumeSource(in *core.LocalVolumeSource, out *v1.LocalVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Path = in.Path
 out.FSType = (*string)(unsafe.Pointer(in.FSType))
 return nil
}
func Convert_core_LocalVolumeSource_To_v1_LocalVolumeSource(in *core.LocalVolumeSource, out *v1.LocalVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_LocalVolumeSource_To_v1_LocalVolumeSource(in, out, s)
}
func autoConvert_v1_NFSVolumeSource_To_core_NFSVolumeSource(in *v1.NFSVolumeSource, out *core.NFSVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Server = in.Server
 out.Path = in.Path
 out.ReadOnly = in.ReadOnly
 return nil
}
func Convert_v1_NFSVolumeSource_To_core_NFSVolumeSource(in *v1.NFSVolumeSource, out *core.NFSVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_NFSVolumeSource_To_core_NFSVolumeSource(in, out, s)
}
func autoConvert_core_NFSVolumeSource_To_v1_NFSVolumeSource(in *core.NFSVolumeSource, out *v1.NFSVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Server = in.Server
 out.Path = in.Path
 out.ReadOnly = in.ReadOnly
 return nil
}
func Convert_core_NFSVolumeSource_To_v1_NFSVolumeSource(in *core.NFSVolumeSource, out *v1.NFSVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_NFSVolumeSource_To_v1_NFSVolumeSource(in, out, s)
}
func autoConvert_v1_Namespace_To_core_Namespace(in *v1.Namespace, out *core.Namespace, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_v1_NamespaceSpec_To_core_NamespaceSpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 if err := Convert_v1_NamespaceStatus_To_core_NamespaceStatus(&in.Status, &out.Status, s); err != nil {
  return err
 }
 return nil
}
func Convert_v1_Namespace_To_core_Namespace(in *v1.Namespace, out *core.Namespace, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_Namespace_To_core_Namespace(in, out, s)
}
func autoConvert_core_Namespace_To_v1_Namespace(in *core.Namespace, out *v1.Namespace, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_core_NamespaceSpec_To_v1_NamespaceSpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 if err := Convert_core_NamespaceStatus_To_v1_NamespaceStatus(&in.Status, &out.Status, s); err != nil {
  return err
 }
 return nil
}
func Convert_core_Namespace_To_v1_Namespace(in *core.Namespace, out *v1.Namespace, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_Namespace_To_v1_Namespace(in, out, s)
}
func autoConvert_v1_NamespaceList_To_core_NamespaceList(in *v1.NamespaceList, out *core.NamespaceList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 out.Items = *(*[]core.Namespace)(unsafe.Pointer(&in.Items))
 return nil
}
func Convert_v1_NamespaceList_To_core_NamespaceList(in *v1.NamespaceList, out *core.NamespaceList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_NamespaceList_To_core_NamespaceList(in, out, s)
}
func autoConvert_core_NamespaceList_To_v1_NamespaceList(in *core.NamespaceList, out *v1.NamespaceList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 out.Items = *(*[]v1.Namespace)(unsafe.Pointer(&in.Items))
 return nil
}
func Convert_core_NamespaceList_To_v1_NamespaceList(in *core.NamespaceList, out *v1.NamespaceList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_NamespaceList_To_v1_NamespaceList(in, out, s)
}
func autoConvert_v1_NamespaceSpec_To_core_NamespaceSpec(in *v1.NamespaceSpec, out *core.NamespaceSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Finalizers = *(*[]core.FinalizerName)(unsafe.Pointer(&in.Finalizers))
 return nil
}
func Convert_v1_NamespaceSpec_To_core_NamespaceSpec(in *v1.NamespaceSpec, out *core.NamespaceSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_NamespaceSpec_To_core_NamespaceSpec(in, out, s)
}
func autoConvert_core_NamespaceSpec_To_v1_NamespaceSpec(in *core.NamespaceSpec, out *v1.NamespaceSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Finalizers = *(*[]v1.FinalizerName)(unsafe.Pointer(&in.Finalizers))
 return nil
}
func Convert_core_NamespaceSpec_To_v1_NamespaceSpec(in *core.NamespaceSpec, out *v1.NamespaceSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_NamespaceSpec_To_v1_NamespaceSpec(in, out, s)
}
func autoConvert_v1_NamespaceStatus_To_core_NamespaceStatus(in *v1.NamespaceStatus, out *core.NamespaceStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Phase = core.NamespacePhase(in.Phase)
 return nil
}
func Convert_v1_NamespaceStatus_To_core_NamespaceStatus(in *v1.NamespaceStatus, out *core.NamespaceStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_NamespaceStatus_To_core_NamespaceStatus(in, out, s)
}
func autoConvert_core_NamespaceStatus_To_v1_NamespaceStatus(in *core.NamespaceStatus, out *v1.NamespaceStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Phase = v1.NamespacePhase(in.Phase)
 return nil
}
func Convert_core_NamespaceStatus_To_v1_NamespaceStatus(in *core.NamespaceStatus, out *v1.NamespaceStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_NamespaceStatus_To_v1_NamespaceStatus(in, out, s)
}
func autoConvert_v1_Node_To_core_Node(in *v1.Node, out *core.Node, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_v1_NodeSpec_To_core_NodeSpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 if err := Convert_v1_NodeStatus_To_core_NodeStatus(&in.Status, &out.Status, s); err != nil {
  return err
 }
 return nil
}
func Convert_v1_Node_To_core_Node(in *v1.Node, out *core.Node, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_Node_To_core_Node(in, out, s)
}
func autoConvert_core_Node_To_v1_Node(in *core.Node, out *v1.Node, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_core_NodeSpec_To_v1_NodeSpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 if err := Convert_core_NodeStatus_To_v1_NodeStatus(&in.Status, &out.Status, s); err != nil {
  return err
 }
 return nil
}
func Convert_core_Node_To_v1_Node(in *core.Node, out *v1.Node, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_Node_To_v1_Node(in, out, s)
}
func autoConvert_v1_NodeAddress_To_core_NodeAddress(in *v1.NodeAddress, out *core.NodeAddress, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Type = core.NodeAddressType(in.Type)
 out.Address = in.Address
 return nil
}
func Convert_v1_NodeAddress_To_core_NodeAddress(in *v1.NodeAddress, out *core.NodeAddress, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_NodeAddress_To_core_NodeAddress(in, out, s)
}
func autoConvert_core_NodeAddress_To_v1_NodeAddress(in *core.NodeAddress, out *v1.NodeAddress, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Type = v1.NodeAddressType(in.Type)
 out.Address = in.Address
 return nil
}
func Convert_core_NodeAddress_To_v1_NodeAddress(in *core.NodeAddress, out *v1.NodeAddress, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_NodeAddress_To_v1_NodeAddress(in, out, s)
}
func autoConvert_v1_NodeAffinity_To_core_NodeAffinity(in *v1.NodeAffinity, out *core.NodeAffinity, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.RequiredDuringSchedulingIgnoredDuringExecution = (*core.NodeSelector)(unsafe.Pointer(in.RequiredDuringSchedulingIgnoredDuringExecution))
 out.PreferredDuringSchedulingIgnoredDuringExecution = *(*[]core.PreferredSchedulingTerm)(unsafe.Pointer(&in.PreferredDuringSchedulingIgnoredDuringExecution))
 return nil
}
func Convert_v1_NodeAffinity_To_core_NodeAffinity(in *v1.NodeAffinity, out *core.NodeAffinity, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_NodeAffinity_To_core_NodeAffinity(in, out, s)
}
func autoConvert_core_NodeAffinity_To_v1_NodeAffinity(in *core.NodeAffinity, out *v1.NodeAffinity, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.RequiredDuringSchedulingIgnoredDuringExecution = (*v1.NodeSelector)(unsafe.Pointer(in.RequiredDuringSchedulingIgnoredDuringExecution))
 out.PreferredDuringSchedulingIgnoredDuringExecution = *(*[]v1.PreferredSchedulingTerm)(unsafe.Pointer(&in.PreferredDuringSchedulingIgnoredDuringExecution))
 return nil
}
func Convert_core_NodeAffinity_To_v1_NodeAffinity(in *core.NodeAffinity, out *v1.NodeAffinity, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_NodeAffinity_To_v1_NodeAffinity(in, out, s)
}
func autoConvert_v1_NodeCondition_To_core_NodeCondition(in *v1.NodeCondition, out *core.NodeCondition, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Type = core.NodeConditionType(in.Type)
 out.Status = core.ConditionStatus(in.Status)
 out.LastHeartbeatTime = in.LastHeartbeatTime
 out.LastTransitionTime = in.LastTransitionTime
 out.Reason = in.Reason
 out.Message = in.Message
 return nil
}
func Convert_v1_NodeCondition_To_core_NodeCondition(in *v1.NodeCondition, out *core.NodeCondition, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_NodeCondition_To_core_NodeCondition(in, out, s)
}
func autoConvert_core_NodeCondition_To_v1_NodeCondition(in *core.NodeCondition, out *v1.NodeCondition, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Type = v1.NodeConditionType(in.Type)
 out.Status = v1.ConditionStatus(in.Status)
 out.LastHeartbeatTime = in.LastHeartbeatTime
 out.LastTransitionTime = in.LastTransitionTime
 out.Reason = in.Reason
 out.Message = in.Message
 return nil
}
func Convert_core_NodeCondition_To_v1_NodeCondition(in *core.NodeCondition, out *v1.NodeCondition, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_NodeCondition_To_v1_NodeCondition(in, out, s)
}
func autoConvert_v1_NodeConfigSource_To_core_NodeConfigSource(in *v1.NodeConfigSource, out *core.NodeConfigSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ConfigMap = (*core.ConfigMapNodeConfigSource)(unsafe.Pointer(in.ConfigMap))
 return nil
}
func Convert_v1_NodeConfigSource_To_core_NodeConfigSource(in *v1.NodeConfigSource, out *core.NodeConfigSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_NodeConfigSource_To_core_NodeConfigSource(in, out, s)
}
func autoConvert_core_NodeConfigSource_To_v1_NodeConfigSource(in *core.NodeConfigSource, out *v1.NodeConfigSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ConfigMap = (*v1.ConfigMapNodeConfigSource)(unsafe.Pointer(in.ConfigMap))
 return nil
}
func Convert_core_NodeConfigSource_To_v1_NodeConfigSource(in *core.NodeConfigSource, out *v1.NodeConfigSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_NodeConfigSource_To_v1_NodeConfigSource(in, out, s)
}
func autoConvert_v1_NodeConfigStatus_To_core_NodeConfigStatus(in *v1.NodeConfigStatus, out *core.NodeConfigStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Assigned = (*core.NodeConfigSource)(unsafe.Pointer(in.Assigned))
 out.Active = (*core.NodeConfigSource)(unsafe.Pointer(in.Active))
 out.LastKnownGood = (*core.NodeConfigSource)(unsafe.Pointer(in.LastKnownGood))
 out.Error = in.Error
 return nil
}
func Convert_v1_NodeConfigStatus_To_core_NodeConfigStatus(in *v1.NodeConfigStatus, out *core.NodeConfigStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_NodeConfigStatus_To_core_NodeConfigStatus(in, out, s)
}
func autoConvert_core_NodeConfigStatus_To_v1_NodeConfigStatus(in *core.NodeConfigStatus, out *v1.NodeConfigStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Assigned = (*v1.NodeConfigSource)(unsafe.Pointer(in.Assigned))
 out.Active = (*v1.NodeConfigSource)(unsafe.Pointer(in.Active))
 out.LastKnownGood = (*v1.NodeConfigSource)(unsafe.Pointer(in.LastKnownGood))
 out.Error = in.Error
 return nil
}
func Convert_core_NodeConfigStatus_To_v1_NodeConfigStatus(in *core.NodeConfigStatus, out *v1.NodeConfigStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_NodeConfigStatus_To_v1_NodeConfigStatus(in, out, s)
}
func autoConvert_v1_NodeDaemonEndpoints_To_core_NodeDaemonEndpoints(in *v1.NodeDaemonEndpoints, out *core.NodeDaemonEndpoints, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := Convert_v1_DaemonEndpoint_To_core_DaemonEndpoint(&in.KubeletEndpoint, &out.KubeletEndpoint, s); err != nil {
  return err
 }
 return nil
}
func Convert_v1_NodeDaemonEndpoints_To_core_NodeDaemonEndpoints(in *v1.NodeDaemonEndpoints, out *core.NodeDaemonEndpoints, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_NodeDaemonEndpoints_To_core_NodeDaemonEndpoints(in, out, s)
}
func autoConvert_core_NodeDaemonEndpoints_To_v1_NodeDaemonEndpoints(in *core.NodeDaemonEndpoints, out *v1.NodeDaemonEndpoints, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := Convert_core_DaemonEndpoint_To_v1_DaemonEndpoint(&in.KubeletEndpoint, &out.KubeletEndpoint, s); err != nil {
  return err
 }
 return nil
}
func Convert_core_NodeDaemonEndpoints_To_v1_NodeDaemonEndpoints(in *core.NodeDaemonEndpoints, out *v1.NodeDaemonEndpoints, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_NodeDaemonEndpoints_To_v1_NodeDaemonEndpoints(in, out, s)
}
func autoConvert_v1_NodeList_To_core_NodeList(in *v1.NodeList, out *core.NodeList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 out.Items = *(*[]core.Node)(unsafe.Pointer(&in.Items))
 return nil
}
func Convert_v1_NodeList_To_core_NodeList(in *v1.NodeList, out *core.NodeList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_NodeList_To_core_NodeList(in, out, s)
}
func autoConvert_core_NodeList_To_v1_NodeList(in *core.NodeList, out *v1.NodeList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 out.Items = *(*[]v1.Node)(unsafe.Pointer(&in.Items))
 return nil
}
func Convert_core_NodeList_To_v1_NodeList(in *core.NodeList, out *v1.NodeList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_NodeList_To_v1_NodeList(in, out, s)
}
func autoConvert_v1_NodeProxyOptions_To_core_NodeProxyOptions(in *v1.NodeProxyOptions, out *core.NodeProxyOptions, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Path = in.Path
 return nil
}
func Convert_v1_NodeProxyOptions_To_core_NodeProxyOptions(in *v1.NodeProxyOptions, out *core.NodeProxyOptions, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_NodeProxyOptions_To_core_NodeProxyOptions(in, out, s)
}
func autoConvert_core_NodeProxyOptions_To_v1_NodeProxyOptions(in *core.NodeProxyOptions, out *v1.NodeProxyOptions, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Path = in.Path
 return nil
}
func Convert_core_NodeProxyOptions_To_v1_NodeProxyOptions(in *core.NodeProxyOptions, out *v1.NodeProxyOptions, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_NodeProxyOptions_To_v1_NodeProxyOptions(in, out, s)
}
func autoConvert_v1_NodeResources_To_core_NodeResources(in *v1.NodeResources, out *core.NodeResources, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Capacity = *(*core.ResourceList)(unsafe.Pointer(&in.Capacity))
 return nil
}
func Convert_v1_NodeResources_To_core_NodeResources(in *v1.NodeResources, out *core.NodeResources, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_NodeResources_To_core_NodeResources(in, out, s)
}
func autoConvert_core_NodeResources_To_v1_NodeResources(in *core.NodeResources, out *v1.NodeResources, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Capacity = *(*v1.ResourceList)(unsafe.Pointer(&in.Capacity))
 return nil
}
func Convert_core_NodeResources_To_v1_NodeResources(in *core.NodeResources, out *v1.NodeResources, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_NodeResources_To_v1_NodeResources(in, out, s)
}
func autoConvert_v1_NodeSelector_To_core_NodeSelector(in *v1.NodeSelector, out *core.NodeSelector, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.NodeSelectorTerms = *(*[]core.NodeSelectorTerm)(unsafe.Pointer(&in.NodeSelectorTerms))
 return nil
}
func Convert_v1_NodeSelector_To_core_NodeSelector(in *v1.NodeSelector, out *core.NodeSelector, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_NodeSelector_To_core_NodeSelector(in, out, s)
}
func autoConvert_core_NodeSelector_To_v1_NodeSelector(in *core.NodeSelector, out *v1.NodeSelector, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.NodeSelectorTerms = *(*[]v1.NodeSelectorTerm)(unsafe.Pointer(&in.NodeSelectorTerms))
 return nil
}
func Convert_core_NodeSelector_To_v1_NodeSelector(in *core.NodeSelector, out *v1.NodeSelector, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_NodeSelector_To_v1_NodeSelector(in, out, s)
}
func autoConvert_v1_NodeSelectorRequirement_To_core_NodeSelectorRequirement(in *v1.NodeSelectorRequirement, out *core.NodeSelectorRequirement, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Key = in.Key
 out.Operator = core.NodeSelectorOperator(in.Operator)
 out.Values = *(*[]string)(unsafe.Pointer(&in.Values))
 return nil
}
func Convert_v1_NodeSelectorRequirement_To_core_NodeSelectorRequirement(in *v1.NodeSelectorRequirement, out *core.NodeSelectorRequirement, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_NodeSelectorRequirement_To_core_NodeSelectorRequirement(in, out, s)
}
func autoConvert_core_NodeSelectorRequirement_To_v1_NodeSelectorRequirement(in *core.NodeSelectorRequirement, out *v1.NodeSelectorRequirement, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Key = in.Key
 out.Operator = v1.NodeSelectorOperator(in.Operator)
 out.Values = *(*[]string)(unsafe.Pointer(&in.Values))
 return nil
}
func Convert_core_NodeSelectorRequirement_To_v1_NodeSelectorRequirement(in *core.NodeSelectorRequirement, out *v1.NodeSelectorRequirement, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_NodeSelectorRequirement_To_v1_NodeSelectorRequirement(in, out, s)
}
func autoConvert_v1_NodeSelectorTerm_To_core_NodeSelectorTerm(in *v1.NodeSelectorTerm, out *core.NodeSelectorTerm, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.MatchExpressions = *(*[]core.NodeSelectorRequirement)(unsafe.Pointer(&in.MatchExpressions))
 out.MatchFields = *(*[]core.NodeSelectorRequirement)(unsafe.Pointer(&in.MatchFields))
 return nil
}
func Convert_v1_NodeSelectorTerm_To_core_NodeSelectorTerm(in *v1.NodeSelectorTerm, out *core.NodeSelectorTerm, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_NodeSelectorTerm_To_core_NodeSelectorTerm(in, out, s)
}
func autoConvert_core_NodeSelectorTerm_To_v1_NodeSelectorTerm(in *core.NodeSelectorTerm, out *v1.NodeSelectorTerm, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.MatchExpressions = *(*[]v1.NodeSelectorRequirement)(unsafe.Pointer(&in.MatchExpressions))
 out.MatchFields = *(*[]v1.NodeSelectorRequirement)(unsafe.Pointer(&in.MatchFields))
 return nil
}
func Convert_core_NodeSelectorTerm_To_v1_NodeSelectorTerm(in *core.NodeSelectorTerm, out *v1.NodeSelectorTerm, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_NodeSelectorTerm_To_v1_NodeSelectorTerm(in, out, s)
}
func autoConvert_v1_NodeSpec_To_core_NodeSpec(in *v1.NodeSpec, out *core.NodeSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.PodCIDR = in.PodCIDR
 out.ProviderID = in.ProviderID
 out.Unschedulable = in.Unschedulable
 out.Taints = *(*[]core.Taint)(unsafe.Pointer(&in.Taints))
 out.ConfigSource = (*core.NodeConfigSource)(unsafe.Pointer(in.ConfigSource))
 out.DoNotUse_ExternalID = in.DoNotUse_ExternalID
 return nil
}
func Convert_v1_NodeSpec_To_core_NodeSpec(in *v1.NodeSpec, out *core.NodeSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_NodeSpec_To_core_NodeSpec(in, out, s)
}
func autoConvert_core_NodeSpec_To_v1_NodeSpec(in *core.NodeSpec, out *v1.NodeSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.PodCIDR = in.PodCIDR
 out.ProviderID = in.ProviderID
 out.Unschedulable = in.Unschedulable
 out.Taints = *(*[]v1.Taint)(unsafe.Pointer(&in.Taints))
 out.ConfigSource = (*v1.NodeConfigSource)(unsafe.Pointer(in.ConfigSource))
 out.DoNotUse_ExternalID = in.DoNotUse_ExternalID
 return nil
}
func Convert_core_NodeSpec_To_v1_NodeSpec(in *core.NodeSpec, out *v1.NodeSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_NodeSpec_To_v1_NodeSpec(in, out, s)
}
func autoConvert_v1_NodeStatus_To_core_NodeStatus(in *v1.NodeStatus, out *core.NodeStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Capacity = *(*core.ResourceList)(unsafe.Pointer(&in.Capacity))
 out.Allocatable = *(*core.ResourceList)(unsafe.Pointer(&in.Allocatable))
 out.Phase = core.NodePhase(in.Phase)
 out.Conditions = *(*[]core.NodeCondition)(unsafe.Pointer(&in.Conditions))
 out.Addresses = *(*[]core.NodeAddress)(unsafe.Pointer(&in.Addresses))
 if err := Convert_v1_NodeDaemonEndpoints_To_core_NodeDaemonEndpoints(&in.DaemonEndpoints, &out.DaemonEndpoints, s); err != nil {
  return err
 }
 if err := Convert_v1_NodeSystemInfo_To_core_NodeSystemInfo(&in.NodeInfo, &out.NodeInfo, s); err != nil {
  return err
 }
 out.Images = *(*[]core.ContainerImage)(unsafe.Pointer(&in.Images))
 out.VolumesInUse = *(*[]core.UniqueVolumeName)(unsafe.Pointer(&in.VolumesInUse))
 out.VolumesAttached = *(*[]core.AttachedVolume)(unsafe.Pointer(&in.VolumesAttached))
 out.Config = (*core.NodeConfigStatus)(unsafe.Pointer(in.Config))
 return nil
}
func Convert_v1_NodeStatus_To_core_NodeStatus(in *v1.NodeStatus, out *core.NodeStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_NodeStatus_To_core_NodeStatus(in, out, s)
}
func autoConvert_core_NodeStatus_To_v1_NodeStatus(in *core.NodeStatus, out *v1.NodeStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Capacity = *(*v1.ResourceList)(unsafe.Pointer(&in.Capacity))
 out.Allocatable = *(*v1.ResourceList)(unsafe.Pointer(&in.Allocatable))
 out.Phase = v1.NodePhase(in.Phase)
 out.Conditions = *(*[]v1.NodeCondition)(unsafe.Pointer(&in.Conditions))
 out.Addresses = *(*[]v1.NodeAddress)(unsafe.Pointer(&in.Addresses))
 if err := Convert_core_NodeDaemonEndpoints_To_v1_NodeDaemonEndpoints(&in.DaemonEndpoints, &out.DaemonEndpoints, s); err != nil {
  return err
 }
 if err := Convert_core_NodeSystemInfo_To_v1_NodeSystemInfo(&in.NodeInfo, &out.NodeInfo, s); err != nil {
  return err
 }
 out.Images = *(*[]v1.ContainerImage)(unsafe.Pointer(&in.Images))
 out.VolumesInUse = *(*[]v1.UniqueVolumeName)(unsafe.Pointer(&in.VolumesInUse))
 out.VolumesAttached = *(*[]v1.AttachedVolume)(unsafe.Pointer(&in.VolumesAttached))
 out.Config = (*v1.NodeConfigStatus)(unsafe.Pointer(in.Config))
 return nil
}
func Convert_core_NodeStatus_To_v1_NodeStatus(in *core.NodeStatus, out *v1.NodeStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_NodeStatus_To_v1_NodeStatus(in, out, s)
}
func autoConvert_v1_NodeSystemInfo_To_core_NodeSystemInfo(in *v1.NodeSystemInfo, out *core.NodeSystemInfo, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.MachineID = in.MachineID
 out.SystemUUID = in.SystemUUID
 out.BootID = in.BootID
 out.KernelVersion = in.KernelVersion
 out.OSImage = in.OSImage
 out.ContainerRuntimeVersion = in.ContainerRuntimeVersion
 out.KubeletVersion = in.KubeletVersion
 out.KubeProxyVersion = in.KubeProxyVersion
 out.OperatingSystem = in.OperatingSystem
 out.Architecture = in.Architecture
 return nil
}
func Convert_v1_NodeSystemInfo_To_core_NodeSystemInfo(in *v1.NodeSystemInfo, out *core.NodeSystemInfo, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_NodeSystemInfo_To_core_NodeSystemInfo(in, out, s)
}
func autoConvert_core_NodeSystemInfo_To_v1_NodeSystemInfo(in *core.NodeSystemInfo, out *v1.NodeSystemInfo, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.MachineID = in.MachineID
 out.SystemUUID = in.SystemUUID
 out.BootID = in.BootID
 out.KernelVersion = in.KernelVersion
 out.OSImage = in.OSImage
 out.ContainerRuntimeVersion = in.ContainerRuntimeVersion
 out.KubeletVersion = in.KubeletVersion
 out.KubeProxyVersion = in.KubeProxyVersion
 out.OperatingSystem = in.OperatingSystem
 out.Architecture = in.Architecture
 return nil
}
func Convert_core_NodeSystemInfo_To_v1_NodeSystemInfo(in *core.NodeSystemInfo, out *v1.NodeSystemInfo, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_NodeSystemInfo_To_v1_NodeSystemInfo(in, out, s)
}
func autoConvert_v1_ObjectFieldSelector_To_core_ObjectFieldSelector(in *v1.ObjectFieldSelector, out *core.ObjectFieldSelector, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.APIVersion = in.APIVersion
 out.FieldPath = in.FieldPath
 return nil
}
func Convert_v1_ObjectFieldSelector_To_core_ObjectFieldSelector(in *v1.ObjectFieldSelector, out *core.ObjectFieldSelector, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_ObjectFieldSelector_To_core_ObjectFieldSelector(in, out, s)
}
func autoConvert_core_ObjectFieldSelector_To_v1_ObjectFieldSelector(in *core.ObjectFieldSelector, out *v1.ObjectFieldSelector, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.APIVersion = in.APIVersion
 out.FieldPath = in.FieldPath
 return nil
}
func Convert_core_ObjectFieldSelector_To_v1_ObjectFieldSelector(in *core.ObjectFieldSelector, out *v1.ObjectFieldSelector, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_ObjectFieldSelector_To_v1_ObjectFieldSelector(in, out, s)
}
func autoConvert_v1_ObjectReference_To_core_ObjectReference(in *v1.ObjectReference, out *core.ObjectReference, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Kind = in.Kind
 out.Namespace = in.Namespace
 out.Name = in.Name
 out.UID = types.UID(in.UID)
 out.APIVersion = in.APIVersion
 out.ResourceVersion = in.ResourceVersion
 out.FieldPath = in.FieldPath
 return nil
}
func Convert_v1_ObjectReference_To_core_ObjectReference(in *v1.ObjectReference, out *core.ObjectReference, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_ObjectReference_To_core_ObjectReference(in, out, s)
}
func autoConvert_core_ObjectReference_To_v1_ObjectReference(in *core.ObjectReference, out *v1.ObjectReference, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Kind = in.Kind
 out.Namespace = in.Namespace
 out.Name = in.Name
 out.UID = types.UID(in.UID)
 out.APIVersion = in.APIVersion
 out.ResourceVersion = in.ResourceVersion
 out.FieldPath = in.FieldPath
 return nil
}
func Convert_core_ObjectReference_To_v1_ObjectReference(in *core.ObjectReference, out *v1.ObjectReference, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_ObjectReference_To_v1_ObjectReference(in, out, s)
}
func autoConvert_v1_PersistentVolume_To_core_PersistentVolume(in *v1.PersistentVolume, out *core.PersistentVolume, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_v1_PersistentVolumeSpec_To_core_PersistentVolumeSpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 if err := Convert_v1_PersistentVolumeStatus_To_core_PersistentVolumeStatus(&in.Status, &out.Status, s); err != nil {
  return err
 }
 return nil
}
func Convert_v1_PersistentVolume_To_core_PersistentVolume(in *v1.PersistentVolume, out *core.PersistentVolume, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_PersistentVolume_To_core_PersistentVolume(in, out, s)
}
func autoConvert_core_PersistentVolume_To_v1_PersistentVolume(in *core.PersistentVolume, out *v1.PersistentVolume, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_core_PersistentVolumeSpec_To_v1_PersistentVolumeSpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 if err := Convert_core_PersistentVolumeStatus_To_v1_PersistentVolumeStatus(&in.Status, &out.Status, s); err != nil {
  return err
 }
 return nil
}
func Convert_core_PersistentVolume_To_v1_PersistentVolume(in *core.PersistentVolume, out *v1.PersistentVolume, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_PersistentVolume_To_v1_PersistentVolume(in, out, s)
}
func autoConvert_v1_PersistentVolumeClaim_To_core_PersistentVolumeClaim(in *v1.PersistentVolumeClaim, out *core.PersistentVolumeClaim, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_v1_PersistentVolumeClaimSpec_To_core_PersistentVolumeClaimSpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 if err := Convert_v1_PersistentVolumeClaimStatus_To_core_PersistentVolumeClaimStatus(&in.Status, &out.Status, s); err != nil {
  return err
 }
 return nil
}
func Convert_v1_PersistentVolumeClaim_To_core_PersistentVolumeClaim(in *v1.PersistentVolumeClaim, out *core.PersistentVolumeClaim, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_PersistentVolumeClaim_To_core_PersistentVolumeClaim(in, out, s)
}
func autoConvert_core_PersistentVolumeClaim_To_v1_PersistentVolumeClaim(in *core.PersistentVolumeClaim, out *v1.PersistentVolumeClaim, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_core_PersistentVolumeClaimSpec_To_v1_PersistentVolumeClaimSpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 if err := Convert_core_PersistentVolumeClaimStatus_To_v1_PersistentVolumeClaimStatus(&in.Status, &out.Status, s); err != nil {
  return err
 }
 return nil
}
func Convert_core_PersistentVolumeClaim_To_v1_PersistentVolumeClaim(in *core.PersistentVolumeClaim, out *v1.PersistentVolumeClaim, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_PersistentVolumeClaim_To_v1_PersistentVolumeClaim(in, out, s)
}
func autoConvert_v1_PersistentVolumeClaimCondition_To_core_PersistentVolumeClaimCondition(in *v1.PersistentVolumeClaimCondition, out *core.PersistentVolumeClaimCondition, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Type = core.PersistentVolumeClaimConditionType(in.Type)
 out.Status = core.ConditionStatus(in.Status)
 out.LastProbeTime = in.LastProbeTime
 out.LastTransitionTime = in.LastTransitionTime
 out.Reason = in.Reason
 out.Message = in.Message
 return nil
}
func Convert_v1_PersistentVolumeClaimCondition_To_core_PersistentVolumeClaimCondition(in *v1.PersistentVolumeClaimCondition, out *core.PersistentVolumeClaimCondition, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_PersistentVolumeClaimCondition_To_core_PersistentVolumeClaimCondition(in, out, s)
}
func autoConvert_core_PersistentVolumeClaimCondition_To_v1_PersistentVolumeClaimCondition(in *core.PersistentVolumeClaimCondition, out *v1.PersistentVolumeClaimCondition, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Type = v1.PersistentVolumeClaimConditionType(in.Type)
 out.Status = v1.ConditionStatus(in.Status)
 out.LastProbeTime = in.LastProbeTime
 out.LastTransitionTime = in.LastTransitionTime
 out.Reason = in.Reason
 out.Message = in.Message
 return nil
}
func Convert_core_PersistentVolumeClaimCondition_To_v1_PersistentVolumeClaimCondition(in *core.PersistentVolumeClaimCondition, out *v1.PersistentVolumeClaimCondition, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_PersistentVolumeClaimCondition_To_v1_PersistentVolumeClaimCondition(in, out, s)
}
func autoConvert_v1_PersistentVolumeClaimList_To_core_PersistentVolumeClaimList(in *v1.PersistentVolumeClaimList, out *core.PersistentVolumeClaimList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 out.Items = *(*[]core.PersistentVolumeClaim)(unsafe.Pointer(&in.Items))
 return nil
}
func Convert_v1_PersistentVolumeClaimList_To_core_PersistentVolumeClaimList(in *v1.PersistentVolumeClaimList, out *core.PersistentVolumeClaimList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_PersistentVolumeClaimList_To_core_PersistentVolumeClaimList(in, out, s)
}
func autoConvert_core_PersistentVolumeClaimList_To_v1_PersistentVolumeClaimList(in *core.PersistentVolumeClaimList, out *v1.PersistentVolumeClaimList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 out.Items = *(*[]v1.PersistentVolumeClaim)(unsafe.Pointer(&in.Items))
 return nil
}
func Convert_core_PersistentVolumeClaimList_To_v1_PersistentVolumeClaimList(in *core.PersistentVolumeClaimList, out *v1.PersistentVolumeClaimList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_PersistentVolumeClaimList_To_v1_PersistentVolumeClaimList(in, out, s)
}
func autoConvert_v1_PersistentVolumeClaimSpec_To_core_PersistentVolumeClaimSpec(in *v1.PersistentVolumeClaimSpec, out *core.PersistentVolumeClaimSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.AccessModes = *(*[]core.PersistentVolumeAccessMode)(unsafe.Pointer(&in.AccessModes))
 out.Selector = (*metav1.LabelSelector)(unsafe.Pointer(in.Selector))
 if err := Convert_v1_ResourceRequirements_To_core_ResourceRequirements(&in.Resources, &out.Resources, s); err != nil {
  return err
 }
 out.VolumeName = in.VolumeName
 out.StorageClassName = (*string)(unsafe.Pointer(in.StorageClassName))
 out.VolumeMode = (*core.PersistentVolumeMode)(unsafe.Pointer(in.VolumeMode))
 out.DataSource = (*core.TypedLocalObjectReference)(unsafe.Pointer(in.DataSource))
 return nil
}
func Convert_v1_PersistentVolumeClaimSpec_To_core_PersistentVolumeClaimSpec(in *v1.PersistentVolumeClaimSpec, out *core.PersistentVolumeClaimSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_PersistentVolumeClaimSpec_To_core_PersistentVolumeClaimSpec(in, out, s)
}
func autoConvert_core_PersistentVolumeClaimSpec_To_v1_PersistentVolumeClaimSpec(in *core.PersistentVolumeClaimSpec, out *v1.PersistentVolumeClaimSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.AccessModes = *(*[]v1.PersistentVolumeAccessMode)(unsafe.Pointer(&in.AccessModes))
 out.Selector = (*metav1.LabelSelector)(unsafe.Pointer(in.Selector))
 if err := Convert_core_ResourceRequirements_To_v1_ResourceRequirements(&in.Resources, &out.Resources, s); err != nil {
  return err
 }
 out.VolumeName = in.VolumeName
 out.StorageClassName = (*string)(unsafe.Pointer(in.StorageClassName))
 out.VolumeMode = (*v1.PersistentVolumeMode)(unsafe.Pointer(in.VolumeMode))
 out.DataSource = (*v1.TypedLocalObjectReference)(unsafe.Pointer(in.DataSource))
 return nil
}
func Convert_core_PersistentVolumeClaimSpec_To_v1_PersistentVolumeClaimSpec(in *core.PersistentVolumeClaimSpec, out *v1.PersistentVolumeClaimSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_PersistentVolumeClaimSpec_To_v1_PersistentVolumeClaimSpec(in, out, s)
}
func autoConvert_v1_PersistentVolumeClaimStatus_To_core_PersistentVolumeClaimStatus(in *v1.PersistentVolumeClaimStatus, out *core.PersistentVolumeClaimStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Phase = core.PersistentVolumeClaimPhase(in.Phase)
 out.AccessModes = *(*[]core.PersistentVolumeAccessMode)(unsafe.Pointer(&in.AccessModes))
 out.Capacity = *(*core.ResourceList)(unsafe.Pointer(&in.Capacity))
 out.Conditions = *(*[]core.PersistentVolumeClaimCondition)(unsafe.Pointer(&in.Conditions))
 return nil
}
func Convert_v1_PersistentVolumeClaimStatus_To_core_PersistentVolumeClaimStatus(in *v1.PersistentVolumeClaimStatus, out *core.PersistentVolumeClaimStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_PersistentVolumeClaimStatus_To_core_PersistentVolumeClaimStatus(in, out, s)
}
func autoConvert_core_PersistentVolumeClaimStatus_To_v1_PersistentVolumeClaimStatus(in *core.PersistentVolumeClaimStatus, out *v1.PersistentVolumeClaimStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Phase = v1.PersistentVolumeClaimPhase(in.Phase)
 out.AccessModes = *(*[]v1.PersistentVolumeAccessMode)(unsafe.Pointer(&in.AccessModes))
 out.Capacity = *(*v1.ResourceList)(unsafe.Pointer(&in.Capacity))
 out.Conditions = *(*[]v1.PersistentVolumeClaimCondition)(unsafe.Pointer(&in.Conditions))
 return nil
}
func Convert_core_PersistentVolumeClaimStatus_To_v1_PersistentVolumeClaimStatus(in *core.PersistentVolumeClaimStatus, out *v1.PersistentVolumeClaimStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_PersistentVolumeClaimStatus_To_v1_PersistentVolumeClaimStatus(in, out, s)
}
func autoConvert_v1_PersistentVolumeClaimVolumeSource_To_core_PersistentVolumeClaimVolumeSource(in *v1.PersistentVolumeClaimVolumeSource, out *core.PersistentVolumeClaimVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ClaimName = in.ClaimName
 out.ReadOnly = in.ReadOnly
 return nil
}
func Convert_v1_PersistentVolumeClaimVolumeSource_To_core_PersistentVolumeClaimVolumeSource(in *v1.PersistentVolumeClaimVolumeSource, out *core.PersistentVolumeClaimVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_PersistentVolumeClaimVolumeSource_To_core_PersistentVolumeClaimVolumeSource(in, out, s)
}
func autoConvert_core_PersistentVolumeClaimVolumeSource_To_v1_PersistentVolumeClaimVolumeSource(in *core.PersistentVolumeClaimVolumeSource, out *v1.PersistentVolumeClaimVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ClaimName = in.ClaimName
 out.ReadOnly = in.ReadOnly
 return nil
}
func Convert_core_PersistentVolumeClaimVolumeSource_To_v1_PersistentVolumeClaimVolumeSource(in *core.PersistentVolumeClaimVolumeSource, out *v1.PersistentVolumeClaimVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_PersistentVolumeClaimVolumeSource_To_v1_PersistentVolumeClaimVolumeSource(in, out, s)
}
func autoConvert_v1_PersistentVolumeList_To_core_PersistentVolumeList(in *v1.PersistentVolumeList, out *core.PersistentVolumeList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]core.PersistentVolume, len(*in))
  for i := range *in {
   if err := Convert_v1_PersistentVolume_To_core_PersistentVolume(&(*in)[i], &(*out)[i], s); err != nil {
    return err
   }
  }
 } else {
  out.Items = nil
 }
 return nil
}
func Convert_v1_PersistentVolumeList_To_core_PersistentVolumeList(in *v1.PersistentVolumeList, out *core.PersistentVolumeList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_PersistentVolumeList_To_core_PersistentVolumeList(in, out, s)
}
func autoConvert_core_PersistentVolumeList_To_v1_PersistentVolumeList(in *core.PersistentVolumeList, out *v1.PersistentVolumeList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]v1.PersistentVolume, len(*in))
  for i := range *in {
   if err := Convert_core_PersistentVolume_To_v1_PersistentVolume(&(*in)[i], &(*out)[i], s); err != nil {
    return err
   }
  }
 } else {
  out.Items = nil
 }
 return nil
}
func Convert_core_PersistentVolumeList_To_v1_PersistentVolumeList(in *core.PersistentVolumeList, out *v1.PersistentVolumeList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_PersistentVolumeList_To_v1_PersistentVolumeList(in, out, s)
}
func autoConvert_v1_PersistentVolumeSource_To_core_PersistentVolumeSource(in *v1.PersistentVolumeSource, out *core.PersistentVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.GCEPersistentDisk = (*core.GCEPersistentDiskVolumeSource)(unsafe.Pointer(in.GCEPersistentDisk))
 out.AWSElasticBlockStore = (*core.AWSElasticBlockStoreVolumeSource)(unsafe.Pointer(in.AWSElasticBlockStore))
 out.HostPath = (*core.HostPathVolumeSource)(unsafe.Pointer(in.HostPath))
 out.Glusterfs = (*core.GlusterfsPersistentVolumeSource)(unsafe.Pointer(in.Glusterfs))
 out.NFS = (*core.NFSVolumeSource)(unsafe.Pointer(in.NFS))
 out.RBD = (*core.RBDPersistentVolumeSource)(unsafe.Pointer(in.RBD))
 out.ISCSI = (*core.ISCSIPersistentVolumeSource)(unsafe.Pointer(in.ISCSI))
 out.Cinder = (*core.CinderPersistentVolumeSource)(unsafe.Pointer(in.Cinder))
 out.CephFS = (*core.CephFSPersistentVolumeSource)(unsafe.Pointer(in.CephFS))
 out.FC = (*core.FCVolumeSource)(unsafe.Pointer(in.FC))
 out.Flocker = (*core.FlockerVolumeSource)(unsafe.Pointer(in.Flocker))
 out.FlexVolume = (*core.FlexPersistentVolumeSource)(unsafe.Pointer(in.FlexVolume))
 out.AzureFile = (*core.AzureFilePersistentVolumeSource)(unsafe.Pointer(in.AzureFile))
 out.VsphereVolume = (*core.VsphereVirtualDiskVolumeSource)(unsafe.Pointer(in.VsphereVolume))
 out.Quobyte = (*core.QuobyteVolumeSource)(unsafe.Pointer(in.Quobyte))
 out.AzureDisk = (*core.AzureDiskVolumeSource)(unsafe.Pointer(in.AzureDisk))
 out.PhotonPersistentDisk = (*core.PhotonPersistentDiskVolumeSource)(unsafe.Pointer(in.PhotonPersistentDisk))
 out.PortworxVolume = (*core.PortworxVolumeSource)(unsafe.Pointer(in.PortworxVolume))
 out.ScaleIO = (*core.ScaleIOPersistentVolumeSource)(unsafe.Pointer(in.ScaleIO))
 out.Local = (*core.LocalVolumeSource)(unsafe.Pointer(in.Local))
 out.StorageOS = (*core.StorageOSPersistentVolumeSource)(unsafe.Pointer(in.StorageOS))
 out.CSI = (*core.CSIPersistentVolumeSource)(unsafe.Pointer(in.CSI))
 return nil
}
func Convert_v1_PersistentVolumeSource_To_core_PersistentVolumeSource(in *v1.PersistentVolumeSource, out *core.PersistentVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_PersistentVolumeSource_To_core_PersistentVolumeSource(in, out, s)
}
func autoConvert_core_PersistentVolumeSource_To_v1_PersistentVolumeSource(in *core.PersistentVolumeSource, out *v1.PersistentVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.GCEPersistentDisk = (*v1.GCEPersistentDiskVolumeSource)(unsafe.Pointer(in.GCEPersistentDisk))
 out.AWSElasticBlockStore = (*v1.AWSElasticBlockStoreVolumeSource)(unsafe.Pointer(in.AWSElasticBlockStore))
 out.HostPath = (*v1.HostPathVolumeSource)(unsafe.Pointer(in.HostPath))
 out.Glusterfs = (*v1.GlusterfsPersistentVolumeSource)(unsafe.Pointer(in.Glusterfs))
 out.NFS = (*v1.NFSVolumeSource)(unsafe.Pointer(in.NFS))
 out.RBD = (*v1.RBDPersistentVolumeSource)(unsafe.Pointer(in.RBD))
 out.Quobyte = (*v1.QuobyteVolumeSource)(unsafe.Pointer(in.Quobyte))
 out.ISCSI = (*v1.ISCSIPersistentVolumeSource)(unsafe.Pointer(in.ISCSI))
 out.FlexVolume = (*v1.FlexPersistentVolumeSource)(unsafe.Pointer(in.FlexVolume))
 out.Cinder = (*v1.CinderPersistentVolumeSource)(unsafe.Pointer(in.Cinder))
 out.CephFS = (*v1.CephFSPersistentVolumeSource)(unsafe.Pointer(in.CephFS))
 out.FC = (*v1.FCVolumeSource)(unsafe.Pointer(in.FC))
 out.Flocker = (*v1.FlockerVolumeSource)(unsafe.Pointer(in.Flocker))
 out.AzureFile = (*v1.AzureFilePersistentVolumeSource)(unsafe.Pointer(in.AzureFile))
 out.VsphereVolume = (*v1.VsphereVirtualDiskVolumeSource)(unsafe.Pointer(in.VsphereVolume))
 out.AzureDisk = (*v1.AzureDiskVolumeSource)(unsafe.Pointer(in.AzureDisk))
 out.PhotonPersistentDisk = (*v1.PhotonPersistentDiskVolumeSource)(unsafe.Pointer(in.PhotonPersistentDisk))
 out.PortworxVolume = (*v1.PortworxVolumeSource)(unsafe.Pointer(in.PortworxVolume))
 out.ScaleIO = (*v1.ScaleIOPersistentVolumeSource)(unsafe.Pointer(in.ScaleIO))
 out.Local = (*v1.LocalVolumeSource)(unsafe.Pointer(in.Local))
 out.StorageOS = (*v1.StorageOSPersistentVolumeSource)(unsafe.Pointer(in.StorageOS))
 out.CSI = (*v1.CSIPersistentVolumeSource)(unsafe.Pointer(in.CSI))
 return nil
}
func Convert_core_PersistentVolumeSource_To_v1_PersistentVolumeSource(in *core.PersistentVolumeSource, out *v1.PersistentVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_PersistentVolumeSource_To_v1_PersistentVolumeSource(in, out, s)
}
func autoConvert_v1_PersistentVolumeSpec_To_core_PersistentVolumeSpec(in *v1.PersistentVolumeSpec, out *core.PersistentVolumeSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Capacity = *(*core.ResourceList)(unsafe.Pointer(&in.Capacity))
 if err := Convert_v1_PersistentVolumeSource_To_core_PersistentVolumeSource(&in.PersistentVolumeSource, &out.PersistentVolumeSource, s); err != nil {
  return err
 }
 out.AccessModes = *(*[]core.PersistentVolumeAccessMode)(unsafe.Pointer(&in.AccessModes))
 out.ClaimRef = (*core.ObjectReference)(unsafe.Pointer(in.ClaimRef))
 out.PersistentVolumeReclaimPolicy = core.PersistentVolumeReclaimPolicy(in.PersistentVolumeReclaimPolicy)
 out.StorageClassName = in.StorageClassName
 out.MountOptions = *(*[]string)(unsafe.Pointer(&in.MountOptions))
 out.VolumeMode = (*core.PersistentVolumeMode)(unsafe.Pointer(in.VolumeMode))
 out.NodeAffinity = (*core.VolumeNodeAffinity)(unsafe.Pointer(in.NodeAffinity))
 return nil
}
func Convert_v1_PersistentVolumeSpec_To_core_PersistentVolumeSpec(in *v1.PersistentVolumeSpec, out *core.PersistentVolumeSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_PersistentVolumeSpec_To_core_PersistentVolumeSpec(in, out, s)
}
func autoConvert_core_PersistentVolumeSpec_To_v1_PersistentVolumeSpec(in *core.PersistentVolumeSpec, out *v1.PersistentVolumeSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Capacity = *(*v1.ResourceList)(unsafe.Pointer(&in.Capacity))
 if err := Convert_core_PersistentVolumeSource_To_v1_PersistentVolumeSource(&in.PersistentVolumeSource, &out.PersistentVolumeSource, s); err != nil {
  return err
 }
 out.AccessModes = *(*[]v1.PersistentVolumeAccessMode)(unsafe.Pointer(&in.AccessModes))
 out.ClaimRef = (*v1.ObjectReference)(unsafe.Pointer(in.ClaimRef))
 out.PersistentVolumeReclaimPolicy = v1.PersistentVolumeReclaimPolicy(in.PersistentVolumeReclaimPolicy)
 out.StorageClassName = in.StorageClassName
 out.MountOptions = *(*[]string)(unsafe.Pointer(&in.MountOptions))
 out.VolumeMode = (*v1.PersistentVolumeMode)(unsafe.Pointer(in.VolumeMode))
 out.NodeAffinity = (*v1.VolumeNodeAffinity)(unsafe.Pointer(in.NodeAffinity))
 return nil
}
func Convert_core_PersistentVolumeSpec_To_v1_PersistentVolumeSpec(in *core.PersistentVolumeSpec, out *v1.PersistentVolumeSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_PersistentVolumeSpec_To_v1_PersistentVolumeSpec(in, out, s)
}
func autoConvert_v1_PersistentVolumeStatus_To_core_PersistentVolumeStatus(in *v1.PersistentVolumeStatus, out *core.PersistentVolumeStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Phase = core.PersistentVolumePhase(in.Phase)
 out.Message = in.Message
 out.Reason = in.Reason
 return nil
}
func Convert_v1_PersistentVolumeStatus_To_core_PersistentVolumeStatus(in *v1.PersistentVolumeStatus, out *core.PersistentVolumeStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_PersistentVolumeStatus_To_core_PersistentVolumeStatus(in, out, s)
}
func autoConvert_core_PersistentVolumeStatus_To_v1_PersistentVolumeStatus(in *core.PersistentVolumeStatus, out *v1.PersistentVolumeStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Phase = v1.PersistentVolumePhase(in.Phase)
 out.Message = in.Message
 out.Reason = in.Reason
 return nil
}
func Convert_core_PersistentVolumeStatus_To_v1_PersistentVolumeStatus(in *core.PersistentVolumeStatus, out *v1.PersistentVolumeStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_PersistentVolumeStatus_To_v1_PersistentVolumeStatus(in, out, s)
}
func autoConvert_v1_PhotonPersistentDiskVolumeSource_To_core_PhotonPersistentDiskVolumeSource(in *v1.PhotonPersistentDiskVolumeSource, out *core.PhotonPersistentDiskVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.PdID = in.PdID
 out.FSType = in.FSType
 return nil
}
func Convert_v1_PhotonPersistentDiskVolumeSource_To_core_PhotonPersistentDiskVolumeSource(in *v1.PhotonPersistentDiskVolumeSource, out *core.PhotonPersistentDiskVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_PhotonPersistentDiskVolumeSource_To_core_PhotonPersistentDiskVolumeSource(in, out, s)
}
func autoConvert_core_PhotonPersistentDiskVolumeSource_To_v1_PhotonPersistentDiskVolumeSource(in *core.PhotonPersistentDiskVolumeSource, out *v1.PhotonPersistentDiskVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.PdID = in.PdID
 out.FSType = in.FSType
 return nil
}
func Convert_core_PhotonPersistentDiskVolumeSource_To_v1_PhotonPersistentDiskVolumeSource(in *core.PhotonPersistentDiskVolumeSource, out *v1.PhotonPersistentDiskVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_PhotonPersistentDiskVolumeSource_To_v1_PhotonPersistentDiskVolumeSource(in, out, s)
}
func autoConvert_v1_Pod_To_core_Pod(in *v1.Pod, out *core.Pod, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_v1_PodSpec_To_core_PodSpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 if err := Convert_v1_PodStatus_To_core_PodStatus(&in.Status, &out.Status, s); err != nil {
  return err
 }
 return nil
}
func autoConvert_core_Pod_To_v1_Pod(in *core.Pod, out *v1.Pod, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_core_PodSpec_To_v1_PodSpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 if err := Convert_core_PodStatus_To_v1_PodStatus(&in.Status, &out.Status, s); err != nil {
  return err
 }
 return nil
}
func autoConvert_v1_PodAffinity_To_core_PodAffinity(in *v1.PodAffinity, out *core.PodAffinity, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.RequiredDuringSchedulingIgnoredDuringExecution = *(*[]core.PodAffinityTerm)(unsafe.Pointer(&in.RequiredDuringSchedulingIgnoredDuringExecution))
 out.PreferredDuringSchedulingIgnoredDuringExecution = *(*[]core.WeightedPodAffinityTerm)(unsafe.Pointer(&in.PreferredDuringSchedulingIgnoredDuringExecution))
 return nil
}
func Convert_v1_PodAffinity_To_core_PodAffinity(in *v1.PodAffinity, out *core.PodAffinity, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_PodAffinity_To_core_PodAffinity(in, out, s)
}
func autoConvert_core_PodAffinity_To_v1_PodAffinity(in *core.PodAffinity, out *v1.PodAffinity, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.RequiredDuringSchedulingIgnoredDuringExecution = *(*[]v1.PodAffinityTerm)(unsafe.Pointer(&in.RequiredDuringSchedulingIgnoredDuringExecution))
 out.PreferredDuringSchedulingIgnoredDuringExecution = *(*[]v1.WeightedPodAffinityTerm)(unsafe.Pointer(&in.PreferredDuringSchedulingIgnoredDuringExecution))
 return nil
}
func Convert_core_PodAffinity_To_v1_PodAffinity(in *core.PodAffinity, out *v1.PodAffinity, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_PodAffinity_To_v1_PodAffinity(in, out, s)
}
func autoConvert_v1_PodAffinityTerm_To_core_PodAffinityTerm(in *v1.PodAffinityTerm, out *core.PodAffinityTerm, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.LabelSelector = (*metav1.LabelSelector)(unsafe.Pointer(in.LabelSelector))
 out.Namespaces = *(*[]string)(unsafe.Pointer(&in.Namespaces))
 out.TopologyKey = in.TopologyKey
 return nil
}
func Convert_v1_PodAffinityTerm_To_core_PodAffinityTerm(in *v1.PodAffinityTerm, out *core.PodAffinityTerm, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_PodAffinityTerm_To_core_PodAffinityTerm(in, out, s)
}
func autoConvert_core_PodAffinityTerm_To_v1_PodAffinityTerm(in *core.PodAffinityTerm, out *v1.PodAffinityTerm, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.LabelSelector = (*metav1.LabelSelector)(unsafe.Pointer(in.LabelSelector))
 out.Namespaces = *(*[]string)(unsafe.Pointer(&in.Namespaces))
 out.TopologyKey = in.TopologyKey
 return nil
}
func Convert_core_PodAffinityTerm_To_v1_PodAffinityTerm(in *core.PodAffinityTerm, out *v1.PodAffinityTerm, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_PodAffinityTerm_To_v1_PodAffinityTerm(in, out, s)
}
func autoConvert_v1_PodAntiAffinity_To_core_PodAntiAffinity(in *v1.PodAntiAffinity, out *core.PodAntiAffinity, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.RequiredDuringSchedulingIgnoredDuringExecution = *(*[]core.PodAffinityTerm)(unsafe.Pointer(&in.RequiredDuringSchedulingIgnoredDuringExecution))
 out.PreferredDuringSchedulingIgnoredDuringExecution = *(*[]core.WeightedPodAffinityTerm)(unsafe.Pointer(&in.PreferredDuringSchedulingIgnoredDuringExecution))
 return nil
}
func Convert_v1_PodAntiAffinity_To_core_PodAntiAffinity(in *v1.PodAntiAffinity, out *core.PodAntiAffinity, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_PodAntiAffinity_To_core_PodAntiAffinity(in, out, s)
}
func autoConvert_core_PodAntiAffinity_To_v1_PodAntiAffinity(in *core.PodAntiAffinity, out *v1.PodAntiAffinity, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.RequiredDuringSchedulingIgnoredDuringExecution = *(*[]v1.PodAffinityTerm)(unsafe.Pointer(&in.RequiredDuringSchedulingIgnoredDuringExecution))
 out.PreferredDuringSchedulingIgnoredDuringExecution = *(*[]v1.WeightedPodAffinityTerm)(unsafe.Pointer(&in.PreferredDuringSchedulingIgnoredDuringExecution))
 return nil
}
func Convert_core_PodAntiAffinity_To_v1_PodAntiAffinity(in *core.PodAntiAffinity, out *v1.PodAntiAffinity, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_PodAntiAffinity_To_v1_PodAntiAffinity(in, out, s)
}
func autoConvert_v1_PodAttachOptions_To_core_PodAttachOptions(in *v1.PodAttachOptions, out *core.PodAttachOptions, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Stdin = in.Stdin
 out.Stdout = in.Stdout
 out.Stderr = in.Stderr
 out.TTY = in.TTY
 out.Container = in.Container
 return nil
}
func Convert_v1_PodAttachOptions_To_core_PodAttachOptions(in *v1.PodAttachOptions, out *core.PodAttachOptions, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_PodAttachOptions_To_core_PodAttachOptions(in, out, s)
}
func autoConvert_core_PodAttachOptions_To_v1_PodAttachOptions(in *core.PodAttachOptions, out *v1.PodAttachOptions, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Stdin = in.Stdin
 out.Stdout = in.Stdout
 out.Stderr = in.Stderr
 out.TTY = in.TTY
 out.Container = in.Container
 return nil
}
func Convert_core_PodAttachOptions_To_v1_PodAttachOptions(in *core.PodAttachOptions, out *v1.PodAttachOptions, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_PodAttachOptions_To_v1_PodAttachOptions(in, out, s)
}
func autoConvert_v1_PodCondition_To_core_PodCondition(in *v1.PodCondition, out *core.PodCondition, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Type = core.PodConditionType(in.Type)
 out.Status = core.ConditionStatus(in.Status)
 out.LastProbeTime = in.LastProbeTime
 out.LastTransitionTime = in.LastTransitionTime
 out.Reason = in.Reason
 out.Message = in.Message
 return nil
}
func Convert_v1_PodCondition_To_core_PodCondition(in *v1.PodCondition, out *core.PodCondition, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_PodCondition_To_core_PodCondition(in, out, s)
}
func autoConvert_core_PodCondition_To_v1_PodCondition(in *core.PodCondition, out *v1.PodCondition, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Type = v1.PodConditionType(in.Type)
 out.Status = v1.ConditionStatus(in.Status)
 out.LastProbeTime = in.LastProbeTime
 out.LastTransitionTime = in.LastTransitionTime
 out.Reason = in.Reason
 out.Message = in.Message
 return nil
}
func Convert_core_PodCondition_To_v1_PodCondition(in *core.PodCondition, out *v1.PodCondition, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_PodCondition_To_v1_PodCondition(in, out, s)
}
func autoConvert_v1_PodDNSConfig_To_core_PodDNSConfig(in *v1.PodDNSConfig, out *core.PodDNSConfig, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Nameservers = *(*[]string)(unsafe.Pointer(&in.Nameservers))
 out.Searches = *(*[]string)(unsafe.Pointer(&in.Searches))
 out.Options = *(*[]core.PodDNSConfigOption)(unsafe.Pointer(&in.Options))
 return nil
}
func Convert_v1_PodDNSConfig_To_core_PodDNSConfig(in *v1.PodDNSConfig, out *core.PodDNSConfig, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_PodDNSConfig_To_core_PodDNSConfig(in, out, s)
}
func autoConvert_core_PodDNSConfig_To_v1_PodDNSConfig(in *core.PodDNSConfig, out *v1.PodDNSConfig, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Nameservers = *(*[]string)(unsafe.Pointer(&in.Nameservers))
 out.Searches = *(*[]string)(unsafe.Pointer(&in.Searches))
 out.Options = *(*[]v1.PodDNSConfigOption)(unsafe.Pointer(&in.Options))
 return nil
}
func Convert_core_PodDNSConfig_To_v1_PodDNSConfig(in *core.PodDNSConfig, out *v1.PodDNSConfig, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_PodDNSConfig_To_v1_PodDNSConfig(in, out, s)
}
func autoConvert_v1_PodDNSConfigOption_To_core_PodDNSConfigOption(in *v1.PodDNSConfigOption, out *core.PodDNSConfigOption, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Name = in.Name
 out.Value = (*string)(unsafe.Pointer(in.Value))
 return nil
}
func Convert_v1_PodDNSConfigOption_To_core_PodDNSConfigOption(in *v1.PodDNSConfigOption, out *core.PodDNSConfigOption, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_PodDNSConfigOption_To_core_PodDNSConfigOption(in, out, s)
}
func autoConvert_core_PodDNSConfigOption_To_v1_PodDNSConfigOption(in *core.PodDNSConfigOption, out *v1.PodDNSConfigOption, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Name = in.Name
 out.Value = (*string)(unsafe.Pointer(in.Value))
 return nil
}
func Convert_core_PodDNSConfigOption_To_v1_PodDNSConfigOption(in *core.PodDNSConfigOption, out *v1.PodDNSConfigOption, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_PodDNSConfigOption_To_v1_PodDNSConfigOption(in, out, s)
}
func autoConvert_v1_PodExecOptions_To_core_PodExecOptions(in *v1.PodExecOptions, out *core.PodExecOptions, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Stdin = in.Stdin
 out.Stdout = in.Stdout
 out.Stderr = in.Stderr
 out.TTY = in.TTY
 out.Container = in.Container
 out.Command = *(*[]string)(unsafe.Pointer(&in.Command))
 return nil
}
func Convert_v1_PodExecOptions_To_core_PodExecOptions(in *v1.PodExecOptions, out *core.PodExecOptions, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_PodExecOptions_To_core_PodExecOptions(in, out, s)
}
func autoConvert_core_PodExecOptions_To_v1_PodExecOptions(in *core.PodExecOptions, out *v1.PodExecOptions, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Stdin = in.Stdin
 out.Stdout = in.Stdout
 out.Stderr = in.Stderr
 out.TTY = in.TTY
 out.Container = in.Container
 out.Command = *(*[]string)(unsafe.Pointer(&in.Command))
 return nil
}
func Convert_core_PodExecOptions_To_v1_PodExecOptions(in *core.PodExecOptions, out *v1.PodExecOptions, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_PodExecOptions_To_v1_PodExecOptions(in, out, s)
}
func autoConvert_v1_PodList_To_core_PodList(in *v1.PodList, out *core.PodList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]core.Pod, len(*in))
  for i := range *in {
   if err := Convert_v1_Pod_To_core_Pod(&(*in)[i], &(*out)[i], s); err != nil {
    return err
   }
  }
 } else {
  out.Items = nil
 }
 return nil
}
func Convert_v1_PodList_To_core_PodList(in *v1.PodList, out *core.PodList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_PodList_To_core_PodList(in, out, s)
}
func autoConvert_core_PodList_To_v1_PodList(in *core.PodList, out *v1.PodList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]v1.Pod, len(*in))
  for i := range *in {
   if err := Convert_core_Pod_To_v1_Pod(&(*in)[i], &(*out)[i], s); err != nil {
    return err
   }
  }
 } else {
  out.Items = nil
 }
 return nil
}
func Convert_core_PodList_To_v1_PodList(in *core.PodList, out *v1.PodList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_PodList_To_v1_PodList(in, out, s)
}
func autoConvert_v1_PodLogOptions_To_core_PodLogOptions(in *v1.PodLogOptions, out *core.PodLogOptions, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Container = in.Container
 out.Follow = in.Follow
 out.Previous = in.Previous
 out.SinceSeconds = (*int64)(unsafe.Pointer(in.SinceSeconds))
 out.SinceTime = (*metav1.Time)(unsafe.Pointer(in.SinceTime))
 out.Timestamps = in.Timestamps
 out.TailLines = (*int64)(unsafe.Pointer(in.TailLines))
 out.LimitBytes = (*int64)(unsafe.Pointer(in.LimitBytes))
 return nil
}
func Convert_v1_PodLogOptions_To_core_PodLogOptions(in *v1.PodLogOptions, out *core.PodLogOptions, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_PodLogOptions_To_core_PodLogOptions(in, out, s)
}
func autoConvert_core_PodLogOptions_To_v1_PodLogOptions(in *core.PodLogOptions, out *v1.PodLogOptions, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Container = in.Container
 out.Follow = in.Follow
 out.Previous = in.Previous
 out.SinceSeconds = (*int64)(unsafe.Pointer(in.SinceSeconds))
 out.SinceTime = (*metav1.Time)(unsafe.Pointer(in.SinceTime))
 out.Timestamps = in.Timestamps
 out.TailLines = (*int64)(unsafe.Pointer(in.TailLines))
 out.LimitBytes = (*int64)(unsafe.Pointer(in.LimitBytes))
 return nil
}
func Convert_core_PodLogOptions_To_v1_PodLogOptions(in *core.PodLogOptions, out *v1.PodLogOptions, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_PodLogOptions_To_v1_PodLogOptions(in, out, s)
}
func autoConvert_v1_PodPortForwardOptions_To_core_PodPortForwardOptions(in *v1.PodPortForwardOptions, out *core.PodPortForwardOptions, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Ports = *(*[]int32)(unsafe.Pointer(&in.Ports))
 return nil
}
func Convert_v1_PodPortForwardOptions_To_core_PodPortForwardOptions(in *v1.PodPortForwardOptions, out *core.PodPortForwardOptions, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_PodPortForwardOptions_To_core_PodPortForwardOptions(in, out, s)
}
func autoConvert_core_PodPortForwardOptions_To_v1_PodPortForwardOptions(in *core.PodPortForwardOptions, out *v1.PodPortForwardOptions, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Ports = *(*[]int32)(unsafe.Pointer(&in.Ports))
 return nil
}
func Convert_core_PodPortForwardOptions_To_v1_PodPortForwardOptions(in *core.PodPortForwardOptions, out *v1.PodPortForwardOptions, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_PodPortForwardOptions_To_v1_PodPortForwardOptions(in, out, s)
}
func autoConvert_v1_PodProxyOptions_To_core_PodProxyOptions(in *v1.PodProxyOptions, out *core.PodProxyOptions, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Path = in.Path
 return nil
}
func Convert_v1_PodProxyOptions_To_core_PodProxyOptions(in *v1.PodProxyOptions, out *core.PodProxyOptions, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_PodProxyOptions_To_core_PodProxyOptions(in, out, s)
}
func autoConvert_core_PodProxyOptions_To_v1_PodProxyOptions(in *core.PodProxyOptions, out *v1.PodProxyOptions, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Path = in.Path
 return nil
}
func Convert_core_PodProxyOptions_To_v1_PodProxyOptions(in *core.PodProxyOptions, out *v1.PodProxyOptions, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_PodProxyOptions_To_v1_PodProxyOptions(in, out, s)
}
func autoConvert_v1_PodReadinessGate_To_core_PodReadinessGate(in *v1.PodReadinessGate, out *core.PodReadinessGate, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ConditionType = core.PodConditionType(in.ConditionType)
 return nil
}
func Convert_v1_PodReadinessGate_To_core_PodReadinessGate(in *v1.PodReadinessGate, out *core.PodReadinessGate, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_PodReadinessGate_To_core_PodReadinessGate(in, out, s)
}
func autoConvert_core_PodReadinessGate_To_v1_PodReadinessGate(in *core.PodReadinessGate, out *v1.PodReadinessGate, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ConditionType = v1.PodConditionType(in.ConditionType)
 return nil
}
func Convert_core_PodReadinessGate_To_v1_PodReadinessGate(in *core.PodReadinessGate, out *v1.PodReadinessGate, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_PodReadinessGate_To_v1_PodReadinessGate(in, out, s)
}
func autoConvert_v1_PodSecurityContext_To_core_PodSecurityContext(in *v1.PodSecurityContext, out *core.PodSecurityContext, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.SELinuxOptions = (*core.SELinuxOptions)(unsafe.Pointer(in.SELinuxOptions))
 out.RunAsUser = (*int64)(unsafe.Pointer(in.RunAsUser))
 out.RunAsGroup = (*int64)(unsafe.Pointer(in.RunAsGroup))
 out.RunAsNonRoot = (*bool)(unsafe.Pointer(in.RunAsNonRoot))
 out.SupplementalGroups = *(*[]int64)(unsafe.Pointer(&in.SupplementalGroups))
 out.FSGroup = (*int64)(unsafe.Pointer(in.FSGroup))
 out.Sysctls = *(*[]core.Sysctl)(unsafe.Pointer(&in.Sysctls))
 return nil
}
func autoConvert_core_PodSecurityContext_To_v1_PodSecurityContext(in *core.PodSecurityContext, out *v1.PodSecurityContext, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.SELinuxOptions = (*v1.SELinuxOptions)(unsafe.Pointer(in.SELinuxOptions))
 out.RunAsUser = (*int64)(unsafe.Pointer(in.RunAsUser))
 out.RunAsGroup = (*int64)(unsafe.Pointer(in.RunAsGroup))
 out.RunAsNonRoot = (*bool)(unsafe.Pointer(in.RunAsNonRoot))
 out.SupplementalGroups = *(*[]int64)(unsafe.Pointer(&in.SupplementalGroups))
 out.FSGroup = (*int64)(unsafe.Pointer(in.FSGroup))
 out.Sysctls = *(*[]v1.Sysctl)(unsafe.Pointer(&in.Sysctls))
 return nil
}
func autoConvert_v1_PodSignature_To_core_PodSignature(in *v1.PodSignature, out *core.PodSignature, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.PodController = (*metav1.OwnerReference)(unsafe.Pointer(in.PodController))
 return nil
}
func Convert_v1_PodSignature_To_core_PodSignature(in *v1.PodSignature, out *core.PodSignature, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_PodSignature_To_core_PodSignature(in, out, s)
}
func autoConvert_core_PodSignature_To_v1_PodSignature(in *core.PodSignature, out *v1.PodSignature, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.PodController = (*metav1.OwnerReference)(unsafe.Pointer(in.PodController))
 return nil
}
func Convert_core_PodSignature_To_v1_PodSignature(in *core.PodSignature, out *v1.PodSignature, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_PodSignature_To_v1_PodSignature(in, out, s)
}
func autoConvert_v1_PodSpec_To_core_PodSpec(in *v1.PodSpec, out *core.PodSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in.Volumes != nil {
  in, out := &in.Volumes, &out.Volumes
  *out = make([]core.Volume, len(*in))
  for i := range *in {
   if err := Convert_v1_Volume_To_core_Volume(&(*in)[i], &(*out)[i], s); err != nil {
    return err
   }
  }
 } else {
  out.Volumes = nil
 }
 if in.InitContainers != nil {
  in, out := &in.InitContainers, &out.InitContainers
  *out = make([]core.Container, len(*in))
  for i := range *in {
   if err := Convert_v1_Container_To_core_Container(&(*in)[i], &(*out)[i], s); err != nil {
    return err
   }
  }
 } else {
  out.InitContainers = nil
 }
 if in.Containers != nil {
  in, out := &in.Containers, &out.Containers
  *out = make([]core.Container, len(*in))
  for i := range *in {
   if err := Convert_v1_Container_To_core_Container(&(*in)[i], &(*out)[i], s); err != nil {
    return err
   }
  }
 } else {
  out.Containers = nil
 }
 out.RestartPolicy = core.RestartPolicy(in.RestartPolicy)
 out.TerminationGracePeriodSeconds = (*int64)(unsafe.Pointer(in.TerminationGracePeriodSeconds))
 out.ActiveDeadlineSeconds = (*int64)(unsafe.Pointer(in.ActiveDeadlineSeconds))
 out.DNSPolicy = core.DNSPolicy(in.DNSPolicy)
 out.NodeSelector = *(*map[string]string)(unsafe.Pointer(&in.NodeSelector))
 out.ServiceAccountName = in.ServiceAccountName
 out.AutomountServiceAccountToken = (*bool)(unsafe.Pointer(in.AutomountServiceAccountToken))
 out.NodeName = in.NodeName
 if in.SecurityContext != nil {
  in, out := &in.SecurityContext, &out.SecurityContext
  *out = new(core.PodSecurityContext)
  if err := Convert_v1_PodSecurityContext_To_core_PodSecurityContext(*in, *out, s); err != nil {
   return err
  }
 } else {
  out.SecurityContext = nil
 }
 out.ImagePullSecrets = *(*[]core.LocalObjectReference)(unsafe.Pointer(&in.ImagePullSecrets))
 out.Hostname = in.Hostname
 out.Subdomain = in.Subdomain
 out.Affinity = (*core.Affinity)(unsafe.Pointer(in.Affinity))
 out.SchedulerName = in.SchedulerName
 out.Tolerations = *(*[]core.Toleration)(unsafe.Pointer(&in.Tolerations))
 out.HostAliases = *(*[]core.HostAlias)(unsafe.Pointer(&in.HostAliases))
 out.PriorityClassName = in.PriorityClassName
 out.Priority = (*int32)(unsafe.Pointer(in.Priority))
 out.DNSConfig = (*core.PodDNSConfig)(unsafe.Pointer(in.DNSConfig))
 out.ReadinessGates = *(*[]core.PodReadinessGate)(unsafe.Pointer(&in.ReadinessGates))
 out.RuntimeClassName = (*string)(unsafe.Pointer(in.RuntimeClassName))
 out.EnableServiceLinks = (*bool)(unsafe.Pointer(in.EnableServiceLinks))
 return nil
}
func autoConvert_core_PodSpec_To_v1_PodSpec(in *core.PodSpec, out *v1.PodSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in.Volumes != nil {
  in, out := &in.Volumes, &out.Volumes
  *out = make([]v1.Volume, len(*in))
  for i := range *in {
   if err := Convert_core_Volume_To_v1_Volume(&(*in)[i], &(*out)[i], s); err != nil {
    return err
   }
  }
 } else {
  out.Volumes = nil
 }
 if in.InitContainers != nil {
  in, out := &in.InitContainers, &out.InitContainers
  *out = make([]v1.Container, len(*in))
  for i := range *in {
   if err := Convert_core_Container_To_v1_Container(&(*in)[i], &(*out)[i], s); err != nil {
    return err
   }
  }
 } else {
  out.InitContainers = nil
 }
 if in.Containers != nil {
  in, out := &in.Containers, &out.Containers
  *out = make([]v1.Container, len(*in))
  for i := range *in {
   if err := Convert_core_Container_To_v1_Container(&(*in)[i], &(*out)[i], s); err != nil {
    return err
   }
  }
 } else {
  out.Containers = nil
 }
 out.RestartPolicy = v1.RestartPolicy(in.RestartPolicy)
 out.TerminationGracePeriodSeconds = (*int64)(unsafe.Pointer(in.TerminationGracePeriodSeconds))
 out.ActiveDeadlineSeconds = (*int64)(unsafe.Pointer(in.ActiveDeadlineSeconds))
 out.DNSPolicy = v1.DNSPolicy(in.DNSPolicy)
 out.NodeSelector = *(*map[string]string)(unsafe.Pointer(&in.NodeSelector))
 out.ServiceAccountName = in.ServiceAccountName
 out.AutomountServiceAccountToken = (*bool)(unsafe.Pointer(in.AutomountServiceAccountToken))
 out.NodeName = in.NodeName
 if in.SecurityContext != nil {
  in, out := &in.SecurityContext, &out.SecurityContext
  *out = new(v1.PodSecurityContext)
  if err := Convert_core_PodSecurityContext_To_v1_PodSecurityContext(*in, *out, s); err != nil {
   return err
  }
 } else {
  out.SecurityContext = nil
 }
 out.ImagePullSecrets = *(*[]v1.LocalObjectReference)(unsafe.Pointer(&in.ImagePullSecrets))
 out.Hostname = in.Hostname
 out.Subdomain = in.Subdomain
 out.Affinity = (*v1.Affinity)(unsafe.Pointer(in.Affinity))
 out.SchedulerName = in.SchedulerName
 out.Tolerations = *(*[]v1.Toleration)(unsafe.Pointer(&in.Tolerations))
 out.HostAliases = *(*[]v1.HostAlias)(unsafe.Pointer(&in.HostAliases))
 out.PriorityClassName = in.PriorityClassName
 out.Priority = (*int32)(unsafe.Pointer(in.Priority))
 out.DNSConfig = (*v1.PodDNSConfig)(unsafe.Pointer(in.DNSConfig))
 out.ReadinessGates = *(*[]v1.PodReadinessGate)(unsafe.Pointer(&in.ReadinessGates))
 out.RuntimeClassName = (*string)(unsafe.Pointer(in.RuntimeClassName))
 out.EnableServiceLinks = (*bool)(unsafe.Pointer(in.EnableServiceLinks))
 return nil
}
func autoConvert_v1_PodStatus_To_core_PodStatus(in *v1.PodStatus, out *core.PodStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Phase = core.PodPhase(in.Phase)
 out.Conditions = *(*[]core.PodCondition)(unsafe.Pointer(&in.Conditions))
 out.Message = in.Message
 out.Reason = in.Reason
 out.NominatedNodeName = in.NominatedNodeName
 out.HostIP = in.HostIP
 out.PodIP = in.PodIP
 out.StartTime = (*metav1.Time)(unsafe.Pointer(in.StartTime))
 out.InitContainerStatuses = *(*[]core.ContainerStatus)(unsafe.Pointer(&in.InitContainerStatuses))
 out.ContainerStatuses = *(*[]core.ContainerStatus)(unsafe.Pointer(&in.ContainerStatuses))
 out.QOSClass = core.PodQOSClass(in.QOSClass)
 return nil
}
func Convert_v1_PodStatus_To_core_PodStatus(in *v1.PodStatus, out *core.PodStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_PodStatus_To_core_PodStatus(in, out, s)
}
func autoConvert_core_PodStatus_To_v1_PodStatus(in *core.PodStatus, out *v1.PodStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Phase = v1.PodPhase(in.Phase)
 out.Conditions = *(*[]v1.PodCondition)(unsafe.Pointer(&in.Conditions))
 out.Message = in.Message
 out.Reason = in.Reason
 out.NominatedNodeName = in.NominatedNodeName
 out.HostIP = in.HostIP
 out.PodIP = in.PodIP
 out.StartTime = (*metav1.Time)(unsafe.Pointer(in.StartTime))
 out.QOSClass = v1.PodQOSClass(in.QOSClass)
 out.InitContainerStatuses = *(*[]v1.ContainerStatus)(unsafe.Pointer(&in.InitContainerStatuses))
 out.ContainerStatuses = *(*[]v1.ContainerStatus)(unsafe.Pointer(&in.ContainerStatuses))
 return nil
}
func Convert_core_PodStatus_To_v1_PodStatus(in *core.PodStatus, out *v1.PodStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_PodStatus_To_v1_PodStatus(in, out, s)
}
func autoConvert_v1_PodStatusResult_To_core_PodStatusResult(in *v1.PodStatusResult, out *core.PodStatusResult, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_v1_PodStatus_To_core_PodStatus(&in.Status, &out.Status, s); err != nil {
  return err
 }
 return nil
}
func Convert_v1_PodStatusResult_To_core_PodStatusResult(in *v1.PodStatusResult, out *core.PodStatusResult, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_PodStatusResult_To_core_PodStatusResult(in, out, s)
}
func autoConvert_core_PodStatusResult_To_v1_PodStatusResult(in *core.PodStatusResult, out *v1.PodStatusResult, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_core_PodStatus_To_v1_PodStatus(&in.Status, &out.Status, s); err != nil {
  return err
 }
 return nil
}
func Convert_core_PodStatusResult_To_v1_PodStatusResult(in *core.PodStatusResult, out *v1.PodStatusResult, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_PodStatusResult_To_v1_PodStatusResult(in, out, s)
}
func autoConvert_v1_PodTemplate_To_core_PodTemplate(in *v1.PodTemplate, out *core.PodTemplate, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_v1_PodTemplateSpec_To_core_PodTemplateSpec(&in.Template, &out.Template, s); err != nil {
  return err
 }
 return nil
}
func Convert_v1_PodTemplate_To_core_PodTemplate(in *v1.PodTemplate, out *core.PodTemplate, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_PodTemplate_To_core_PodTemplate(in, out, s)
}
func autoConvert_core_PodTemplate_To_v1_PodTemplate(in *core.PodTemplate, out *v1.PodTemplate, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_core_PodTemplateSpec_To_v1_PodTemplateSpec(&in.Template, &out.Template, s); err != nil {
  return err
 }
 return nil
}
func Convert_core_PodTemplate_To_v1_PodTemplate(in *core.PodTemplate, out *v1.PodTemplate, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_PodTemplate_To_v1_PodTemplate(in, out, s)
}
func autoConvert_v1_PodTemplateList_To_core_PodTemplateList(in *v1.PodTemplateList, out *core.PodTemplateList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]core.PodTemplate, len(*in))
  for i := range *in {
   if err := Convert_v1_PodTemplate_To_core_PodTemplate(&(*in)[i], &(*out)[i], s); err != nil {
    return err
   }
  }
 } else {
  out.Items = nil
 }
 return nil
}
func Convert_v1_PodTemplateList_To_core_PodTemplateList(in *v1.PodTemplateList, out *core.PodTemplateList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_PodTemplateList_To_core_PodTemplateList(in, out, s)
}
func autoConvert_core_PodTemplateList_To_v1_PodTemplateList(in *core.PodTemplateList, out *v1.PodTemplateList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]v1.PodTemplate, len(*in))
  for i := range *in {
   if err := Convert_core_PodTemplate_To_v1_PodTemplate(&(*in)[i], &(*out)[i], s); err != nil {
    return err
   }
  }
 } else {
  out.Items = nil
 }
 return nil
}
func Convert_core_PodTemplateList_To_v1_PodTemplateList(in *core.PodTemplateList, out *v1.PodTemplateList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_PodTemplateList_To_v1_PodTemplateList(in, out, s)
}
func autoConvert_v1_PodTemplateSpec_To_core_PodTemplateSpec(in *v1.PodTemplateSpec, out *core.PodTemplateSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_v1_PodSpec_To_core_PodSpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 return nil
}
func autoConvert_core_PodTemplateSpec_To_v1_PodTemplateSpec(in *core.PodTemplateSpec, out *v1.PodTemplateSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_core_PodSpec_To_v1_PodSpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 return nil
}
func autoConvert_v1_PortworxVolumeSource_To_core_PortworxVolumeSource(in *v1.PortworxVolumeSource, out *core.PortworxVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.VolumeID = in.VolumeID
 out.FSType = in.FSType
 out.ReadOnly = in.ReadOnly
 return nil
}
func Convert_v1_PortworxVolumeSource_To_core_PortworxVolumeSource(in *v1.PortworxVolumeSource, out *core.PortworxVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_PortworxVolumeSource_To_core_PortworxVolumeSource(in, out, s)
}
func autoConvert_core_PortworxVolumeSource_To_v1_PortworxVolumeSource(in *core.PortworxVolumeSource, out *v1.PortworxVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.VolumeID = in.VolumeID
 out.FSType = in.FSType
 out.ReadOnly = in.ReadOnly
 return nil
}
func Convert_core_PortworxVolumeSource_To_v1_PortworxVolumeSource(in *core.PortworxVolumeSource, out *v1.PortworxVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_PortworxVolumeSource_To_v1_PortworxVolumeSource(in, out, s)
}
func autoConvert_v1_Preconditions_To_core_Preconditions(in *v1.Preconditions, out *core.Preconditions, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.UID = (*types.UID)(unsafe.Pointer(in.UID))
 return nil
}
func Convert_v1_Preconditions_To_core_Preconditions(in *v1.Preconditions, out *core.Preconditions, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_Preconditions_To_core_Preconditions(in, out, s)
}
func autoConvert_core_Preconditions_To_v1_Preconditions(in *core.Preconditions, out *v1.Preconditions, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.UID = (*types.UID)(unsafe.Pointer(in.UID))
 return nil
}
func Convert_core_Preconditions_To_v1_Preconditions(in *core.Preconditions, out *v1.Preconditions, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_Preconditions_To_v1_Preconditions(in, out, s)
}
func autoConvert_v1_PreferAvoidPodsEntry_To_core_PreferAvoidPodsEntry(in *v1.PreferAvoidPodsEntry, out *core.PreferAvoidPodsEntry, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := Convert_v1_PodSignature_To_core_PodSignature(&in.PodSignature, &out.PodSignature, s); err != nil {
  return err
 }
 out.EvictionTime = in.EvictionTime
 out.Reason = in.Reason
 out.Message = in.Message
 return nil
}
func Convert_v1_PreferAvoidPodsEntry_To_core_PreferAvoidPodsEntry(in *v1.PreferAvoidPodsEntry, out *core.PreferAvoidPodsEntry, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_PreferAvoidPodsEntry_To_core_PreferAvoidPodsEntry(in, out, s)
}
func autoConvert_core_PreferAvoidPodsEntry_To_v1_PreferAvoidPodsEntry(in *core.PreferAvoidPodsEntry, out *v1.PreferAvoidPodsEntry, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := Convert_core_PodSignature_To_v1_PodSignature(&in.PodSignature, &out.PodSignature, s); err != nil {
  return err
 }
 out.EvictionTime = in.EvictionTime
 out.Reason = in.Reason
 out.Message = in.Message
 return nil
}
func Convert_core_PreferAvoidPodsEntry_To_v1_PreferAvoidPodsEntry(in *core.PreferAvoidPodsEntry, out *v1.PreferAvoidPodsEntry, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_PreferAvoidPodsEntry_To_v1_PreferAvoidPodsEntry(in, out, s)
}
func autoConvert_v1_PreferredSchedulingTerm_To_core_PreferredSchedulingTerm(in *v1.PreferredSchedulingTerm, out *core.PreferredSchedulingTerm, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Weight = in.Weight
 if err := Convert_v1_NodeSelectorTerm_To_core_NodeSelectorTerm(&in.Preference, &out.Preference, s); err != nil {
  return err
 }
 return nil
}
func Convert_v1_PreferredSchedulingTerm_To_core_PreferredSchedulingTerm(in *v1.PreferredSchedulingTerm, out *core.PreferredSchedulingTerm, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_PreferredSchedulingTerm_To_core_PreferredSchedulingTerm(in, out, s)
}
func autoConvert_core_PreferredSchedulingTerm_To_v1_PreferredSchedulingTerm(in *core.PreferredSchedulingTerm, out *v1.PreferredSchedulingTerm, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Weight = in.Weight
 if err := Convert_core_NodeSelectorTerm_To_v1_NodeSelectorTerm(&in.Preference, &out.Preference, s); err != nil {
  return err
 }
 return nil
}
func Convert_core_PreferredSchedulingTerm_To_v1_PreferredSchedulingTerm(in *core.PreferredSchedulingTerm, out *v1.PreferredSchedulingTerm, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_PreferredSchedulingTerm_To_v1_PreferredSchedulingTerm(in, out, s)
}
func autoConvert_v1_Probe_To_core_Probe(in *v1.Probe, out *core.Probe, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := Convert_v1_Handler_To_core_Handler(&in.Handler, &out.Handler, s); err != nil {
  return err
 }
 out.InitialDelaySeconds = in.InitialDelaySeconds
 out.TimeoutSeconds = in.TimeoutSeconds
 out.PeriodSeconds = in.PeriodSeconds
 out.SuccessThreshold = in.SuccessThreshold
 out.FailureThreshold = in.FailureThreshold
 return nil
}
func Convert_v1_Probe_To_core_Probe(in *v1.Probe, out *core.Probe, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_Probe_To_core_Probe(in, out, s)
}
func autoConvert_core_Probe_To_v1_Probe(in *core.Probe, out *v1.Probe, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := Convert_core_Handler_To_v1_Handler(&in.Handler, &out.Handler, s); err != nil {
  return err
 }
 out.InitialDelaySeconds = in.InitialDelaySeconds
 out.TimeoutSeconds = in.TimeoutSeconds
 out.PeriodSeconds = in.PeriodSeconds
 out.SuccessThreshold = in.SuccessThreshold
 out.FailureThreshold = in.FailureThreshold
 return nil
}
func Convert_core_Probe_To_v1_Probe(in *core.Probe, out *v1.Probe, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_Probe_To_v1_Probe(in, out, s)
}
func autoConvert_v1_ProjectedVolumeSource_To_core_ProjectedVolumeSource(in *v1.ProjectedVolumeSource, out *core.ProjectedVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in.Sources != nil {
  in, out := &in.Sources, &out.Sources
  *out = make([]core.VolumeProjection, len(*in))
  for i := range *in {
   if err := Convert_v1_VolumeProjection_To_core_VolumeProjection(&(*in)[i], &(*out)[i], s); err != nil {
    return err
   }
  }
 } else {
  out.Sources = nil
 }
 out.DefaultMode = (*int32)(unsafe.Pointer(in.DefaultMode))
 return nil
}
func Convert_v1_ProjectedVolumeSource_To_core_ProjectedVolumeSource(in *v1.ProjectedVolumeSource, out *core.ProjectedVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_ProjectedVolumeSource_To_core_ProjectedVolumeSource(in, out, s)
}
func autoConvert_core_ProjectedVolumeSource_To_v1_ProjectedVolumeSource(in *core.ProjectedVolumeSource, out *v1.ProjectedVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in.Sources != nil {
  in, out := &in.Sources, &out.Sources
  *out = make([]v1.VolumeProjection, len(*in))
  for i := range *in {
   if err := Convert_core_VolumeProjection_To_v1_VolumeProjection(&(*in)[i], &(*out)[i], s); err != nil {
    return err
   }
  }
 } else {
  out.Sources = nil
 }
 out.DefaultMode = (*int32)(unsafe.Pointer(in.DefaultMode))
 return nil
}
func Convert_core_ProjectedVolumeSource_To_v1_ProjectedVolumeSource(in *core.ProjectedVolumeSource, out *v1.ProjectedVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_ProjectedVolumeSource_To_v1_ProjectedVolumeSource(in, out, s)
}
func autoConvert_v1_QuobyteVolumeSource_To_core_QuobyteVolumeSource(in *v1.QuobyteVolumeSource, out *core.QuobyteVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Registry = in.Registry
 out.Volume = in.Volume
 out.ReadOnly = in.ReadOnly
 out.User = in.User
 out.Group = in.Group
 return nil
}
func Convert_v1_QuobyteVolumeSource_To_core_QuobyteVolumeSource(in *v1.QuobyteVolumeSource, out *core.QuobyteVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_QuobyteVolumeSource_To_core_QuobyteVolumeSource(in, out, s)
}
func autoConvert_core_QuobyteVolumeSource_To_v1_QuobyteVolumeSource(in *core.QuobyteVolumeSource, out *v1.QuobyteVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Registry = in.Registry
 out.Volume = in.Volume
 out.ReadOnly = in.ReadOnly
 out.User = in.User
 out.Group = in.Group
 return nil
}
func Convert_core_QuobyteVolumeSource_To_v1_QuobyteVolumeSource(in *core.QuobyteVolumeSource, out *v1.QuobyteVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_QuobyteVolumeSource_To_v1_QuobyteVolumeSource(in, out, s)
}
func autoConvert_v1_RBDPersistentVolumeSource_To_core_RBDPersistentVolumeSource(in *v1.RBDPersistentVolumeSource, out *core.RBDPersistentVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.CephMonitors = *(*[]string)(unsafe.Pointer(&in.CephMonitors))
 out.RBDImage = in.RBDImage
 out.FSType = in.FSType
 out.RBDPool = in.RBDPool
 out.RadosUser = in.RadosUser
 out.Keyring = in.Keyring
 out.SecretRef = (*core.SecretReference)(unsafe.Pointer(in.SecretRef))
 out.ReadOnly = in.ReadOnly
 return nil
}
func Convert_v1_RBDPersistentVolumeSource_To_core_RBDPersistentVolumeSource(in *v1.RBDPersistentVolumeSource, out *core.RBDPersistentVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_RBDPersistentVolumeSource_To_core_RBDPersistentVolumeSource(in, out, s)
}
func autoConvert_core_RBDPersistentVolumeSource_To_v1_RBDPersistentVolumeSource(in *core.RBDPersistentVolumeSource, out *v1.RBDPersistentVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.CephMonitors = *(*[]string)(unsafe.Pointer(&in.CephMonitors))
 out.RBDImage = in.RBDImage
 out.FSType = in.FSType
 out.RBDPool = in.RBDPool
 out.RadosUser = in.RadosUser
 out.Keyring = in.Keyring
 out.SecretRef = (*v1.SecretReference)(unsafe.Pointer(in.SecretRef))
 out.ReadOnly = in.ReadOnly
 return nil
}
func Convert_core_RBDPersistentVolumeSource_To_v1_RBDPersistentVolumeSource(in *core.RBDPersistentVolumeSource, out *v1.RBDPersistentVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_RBDPersistentVolumeSource_To_v1_RBDPersistentVolumeSource(in, out, s)
}
func autoConvert_v1_RBDVolumeSource_To_core_RBDVolumeSource(in *v1.RBDVolumeSource, out *core.RBDVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.CephMonitors = *(*[]string)(unsafe.Pointer(&in.CephMonitors))
 out.RBDImage = in.RBDImage
 out.FSType = in.FSType
 out.RBDPool = in.RBDPool
 out.RadosUser = in.RadosUser
 out.Keyring = in.Keyring
 out.SecretRef = (*core.LocalObjectReference)(unsafe.Pointer(in.SecretRef))
 out.ReadOnly = in.ReadOnly
 return nil
}
func Convert_v1_RBDVolumeSource_To_core_RBDVolumeSource(in *v1.RBDVolumeSource, out *core.RBDVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_RBDVolumeSource_To_core_RBDVolumeSource(in, out, s)
}
func autoConvert_core_RBDVolumeSource_To_v1_RBDVolumeSource(in *core.RBDVolumeSource, out *v1.RBDVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.CephMonitors = *(*[]string)(unsafe.Pointer(&in.CephMonitors))
 out.RBDImage = in.RBDImage
 out.FSType = in.FSType
 out.RBDPool = in.RBDPool
 out.RadosUser = in.RadosUser
 out.Keyring = in.Keyring
 out.SecretRef = (*v1.LocalObjectReference)(unsafe.Pointer(in.SecretRef))
 out.ReadOnly = in.ReadOnly
 return nil
}
func Convert_core_RBDVolumeSource_To_v1_RBDVolumeSource(in *core.RBDVolumeSource, out *v1.RBDVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_RBDVolumeSource_To_v1_RBDVolumeSource(in, out, s)
}
func autoConvert_v1_RangeAllocation_To_core_RangeAllocation(in *v1.RangeAllocation, out *core.RangeAllocation, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 out.Range = in.Range
 out.Data = *(*[]byte)(unsafe.Pointer(&in.Data))
 return nil
}
func Convert_v1_RangeAllocation_To_core_RangeAllocation(in *v1.RangeAllocation, out *core.RangeAllocation, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_RangeAllocation_To_core_RangeAllocation(in, out, s)
}
func autoConvert_core_RangeAllocation_To_v1_RangeAllocation(in *core.RangeAllocation, out *v1.RangeAllocation, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 out.Range = in.Range
 out.Data = *(*[]byte)(unsafe.Pointer(&in.Data))
 return nil
}
func Convert_core_RangeAllocation_To_v1_RangeAllocation(in *core.RangeAllocation, out *v1.RangeAllocation, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_RangeAllocation_To_v1_RangeAllocation(in, out, s)
}
func autoConvert_v1_ReplicationController_To_core_ReplicationController(in *v1.ReplicationController, out *core.ReplicationController, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_v1_ReplicationControllerSpec_To_core_ReplicationControllerSpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 if err := Convert_v1_ReplicationControllerStatus_To_core_ReplicationControllerStatus(&in.Status, &out.Status, s); err != nil {
  return err
 }
 return nil
}
func Convert_v1_ReplicationController_To_core_ReplicationController(in *v1.ReplicationController, out *core.ReplicationController, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_ReplicationController_To_core_ReplicationController(in, out, s)
}
func autoConvert_core_ReplicationController_To_v1_ReplicationController(in *core.ReplicationController, out *v1.ReplicationController, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_core_ReplicationControllerSpec_To_v1_ReplicationControllerSpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 if err := Convert_core_ReplicationControllerStatus_To_v1_ReplicationControllerStatus(&in.Status, &out.Status, s); err != nil {
  return err
 }
 return nil
}
func Convert_core_ReplicationController_To_v1_ReplicationController(in *core.ReplicationController, out *v1.ReplicationController, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_ReplicationController_To_v1_ReplicationController(in, out, s)
}
func autoConvert_v1_ReplicationControllerCondition_To_core_ReplicationControllerCondition(in *v1.ReplicationControllerCondition, out *core.ReplicationControllerCondition, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Type = core.ReplicationControllerConditionType(in.Type)
 out.Status = core.ConditionStatus(in.Status)
 out.LastTransitionTime = in.LastTransitionTime
 out.Reason = in.Reason
 out.Message = in.Message
 return nil
}
func Convert_v1_ReplicationControllerCondition_To_core_ReplicationControllerCondition(in *v1.ReplicationControllerCondition, out *core.ReplicationControllerCondition, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_ReplicationControllerCondition_To_core_ReplicationControllerCondition(in, out, s)
}
func autoConvert_core_ReplicationControllerCondition_To_v1_ReplicationControllerCondition(in *core.ReplicationControllerCondition, out *v1.ReplicationControllerCondition, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Type = v1.ReplicationControllerConditionType(in.Type)
 out.Status = v1.ConditionStatus(in.Status)
 out.LastTransitionTime = in.LastTransitionTime
 out.Reason = in.Reason
 out.Message = in.Message
 return nil
}
func Convert_core_ReplicationControllerCondition_To_v1_ReplicationControllerCondition(in *core.ReplicationControllerCondition, out *v1.ReplicationControllerCondition, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_ReplicationControllerCondition_To_v1_ReplicationControllerCondition(in, out, s)
}
func autoConvert_v1_ReplicationControllerList_To_core_ReplicationControllerList(in *v1.ReplicationControllerList, out *core.ReplicationControllerList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]core.ReplicationController, len(*in))
  for i := range *in {
   if err := Convert_v1_ReplicationController_To_core_ReplicationController(&(*in)[i], &(*out)[i], s); err != nil {
    return err
   }
  }
 } else {
  out.Items = nil
 }
 return nil
}
func Convert_v1_ReplicationControllerList_To_core_ReplicationControllerList(in *v1.ReplicationControllerList, out *core.ReplicationControllerList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_ReplicationControllerList_To_core_ReplicationControllerList(in, out, s)
}
func autoConvert_core_ReplicationControllerList_To_v1_ReplicationControllerList(in *core.ReplicationControllerList, out *v1.ReplicationControllerList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]v1.ReplicationController, len(*in))
  for i := range *in {
   if err := Convert_core_ReplicationController_To_v1_ReplicationController(&(*in)[i], &(*out)[i], s); err != nil {
    return err
   }
  }
 } else {
  out.Items = nil
 }
 return nil
}
func Convert_core_ReplicationControllerList_To_v1_ReplicationControllerList(in *core.ReplicationControllerList, out *v1.ReplicationControllerList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_ReplicationControllerList_To_v1_ReplicationControllerList(in, out, s)
}
func autoConvert_v1_ReplicationControllerSpec_To_core_ReplicationControllerSpec(in *v1.ReplicationControllerSpec, out *core.ReplicationControllerSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := metav1.Convert_Pointer_int32_To_int32(&in.Replicas, &out.Replicas, s); err != nil {
  return err
 }
 out.MinReadySeconds = in.MinReadySeconds
 out.Selector = *(*map[string]string)(unsafe.Pointer(&in.Selector))
 if in.Template != nil {
  in, out := &in.Template, &out.Template
  *out = new(core.PodTemplateSpec)
  if err := Convert_v1_PodTemplateSpec_To_core_PodTemplateSpec(*in, *out, s); err != nil {
   return err
  }
 } else {
  out.Template = nil
 }
 return nil
}
func autoConvert_core_ReplicationControllerSpec_To_v1_ReplicationControllerSpec(in *core.ReplicationControllerSpec, out *v1.ReplicationControllerSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := metav1.Convert_int32_To_Pointer_int32(&in.Replicas, &out.Replicas, s); err != nil {
  return err
 }
 out.MinReadySeconds = in.MinReadySeconds
 out.Selector = *(*map[string]string)(unsafe.Pointer(&in.Selector))
 if in.Template != nil {
  in, out := &in.Template, &out.Template
  *out = new(v1.PodTemplateSpec)
  if err := Convert_core_PodTemplateSpec_To_v1_PodTemplateSpec(*in, *out, s); err != nil {
   return err
  }
 } else {
  out.Template = nil
 }
 return nil
}
func autoConvert_v1_ReplicationControllerStatus_To_core_ReplicationControllerStatus(in *v1.ReplicationControllerStatus, out *core.ReplicationControllerStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Replicas = in.Replicas
 out.FullyLabeledReplicas = in.FullyLabeledReplicas
 out.ReadyReplicas = in.ReadyReplicas
 out.AvailableReplicas = in.AvailableReplicas
 out.ObservedGeneration = in.ObservedGeneration
 out.Conditions = *(*[]core.ReplicationControllerCondition)(unsafe.Pointer(&in.Conditions))
 return nil
}
func Convert_v1_ReplicationControllerStatus_To_core_ReplicationControllerStatus(in *v1.ReplicationControllerStatus, out *core.ReplicationControllerStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_ReplicationControllerStatus_To_core_ReplicationControllerStatus(in, out, s)
}
func autoConvert_core_ReplicationControllerStatus_To_v1_ReplicationControllerStatus(in *core.ReplicationControllerStatus, out *v1.ReplicationControllerStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Replicas = in.Replicas
 out.FullyLabeledReplicas = in.FullyLabeledReplicas
 out.ReadyReplicas = in.ReadyReplicas
 out.AvailableReplicas = in.AvailableReplicas
 out.ObservedGeneration = in.ObservedGeneration
 out.Conditions = *(*[]v1.ReplicationControllerCondition)(unsafe.Pointer(&in.Conditions))
 return nil
}
func Convert_core_ReplicationControllerStatus_To_v1_ReplicationControllerStatus(in *core.ReplicationControllerStatus, out *v1.ReplicationControllerStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_ReplicationControllerStatus_To_v1_ReplicationControllerStatus(in, out, s)
}
func autoConvert_v1_ResourceFieldSelector_To_core_ResourceFieldSelector(in *v1.ResourceFieldSelector, out *core.ResourceFieldSelector, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ContainerName = in.ContainerName
 out.Resource = in.Resource
 out.Divisor = in.Divisor
 return nil
}
func Convert_v1_ResourceFieldSelector_To_core_ResourceFieldSelector(in *v1.ResourceFieldSelector, out *core.ResourceFieldSelector, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_ResourceFieldSelector_To_core_ResourceFieldSelector(in, out, s)
}
func autoConvert_core_ResourceFieldSelector_To_v1_ResourceFieldSelector(in *core.ResourceFieldSelector, out *v1.ResourceFieldSelector, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ContainerName = in.ContainerName
 out.Resource = in.Resource
 out.Divisor = in.Divisor
 return nil
}
func Convert_core_ResourceFieldSelector_To_v1_ResourceFieldSelector(in *core.ResourceFieldSelector, out *v1.ResourceFieldSelector, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_ResourceFieldSelector_To_v1_ResourceFieldSelector(in, out, s)
}
func autoConvert_v1_ResourceQuota_To_core_ResourceQuota(in *v1.ResourceQuota, out *core.ResourceQuota, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_v1_ResourceQuotaSpec_To_core_ResourceQuotaSpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 if err := Convert_v1_ResourceQuotaStatus_To_core_ResourceQuotaStatus(&in.Status, &out.Status, s); err != nil {
  return err
 }
 return nil
}
func Convert_v1_ResourceQuota_To_core_ResourceQuota(in *v1.ResourceQuota, out *core.ResourceQuota, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_ResourceQuota_To_core_ResourceQuota(in, out, s)
}
func autoConvert_core_ResourceQuota_To_v1_ResourceQuota(in *core.ResourceQuota, out *v1.ResourceQuota, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_core_ResourceQuotaSpec_To_v1_ResourceQuotaSpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 if err := Convert_core_ResourceQuotaStatus_To_v1_ResourceQuotaStatus(&in.Status, &out.Status, s); err != nil {
  return err
 }
 return nil
}
func Convert_core_ResourceQuota_To_v1_ResourceQuota(in *core.ResourceQuota, out *v1.ResourceQuota, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_ResourceQuota_To_v1_ResourceQuota(in, out, s)
}
func autoConvert_v1_ResourceQuotaList_To_core_ResourceQuotaList(in *v1.ResourceQuotaList, out *core.ResourceQuotaList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 out.Items = *(*[]core.ResourceQuota)(unsafe.Pointer(&in.Items))
 return nil
}
func Convert_v1_ResourceQuotaList_To_core_ResourceQuotaList(in *v1.ResourceQuotaList, out *core.ResourceQuotaList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_ResourceQuotaList_To_core_ResourceQuotaList(in, out, s)
}
func autoConvert_core_ResourceQuotaList_To_v1_ResourceQuotaList(in *core.ResourceQuotaList, out *v1.ResourceQuotaList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 out.Items = *(*[]v1.ResourceQuota)(unsafe.Pointer(&in.Items))
 return nil
}
func Convert_core_ResourceQuotaList_To_v1_ResourceQuotaList(in *core.ResourceQuotaList, out *v1.ResourceQuotaList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_ResourceQuotaList_To_v1_ResourceQuotaList(in, out, s)
}
func autoConvert_v1_ResourceQuotaSpec_To_core_ResourceQuotaSpec(in *v1.ResourceQuotaSpec, out *core.ResourceQuotaSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Hard = *(*core.ResourceList)(unsafe.Pointer(&in.Hard))
 out.Scopes = *(*[]core.ResourceQuotaScope)(unsafe.Pointer(&in.Scopes))
 out.ScopeSelector = (*core.ScopeSelector)(unsafe.Pointer(in.ScopeSelector))
 return nil
}
func Convert_v1_ResourceQuotaSpec_To_core_ResourceQuotaSpec(in *v1.ResourceQuotaSpec, out *core.ResourceQuotaSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_ResourceQuotaSpec_To_core_ResourceQuotaSpec(in, out, s)
}
func autoConvert_core_ResourceQuotaSpec_To_v1_ResourceQuotaSpec(in *core.ResourceQuotaSpec, out *v1.ResourceQuotaSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Hard = *(*v1.ResourceList)(unsafe.Pointer(&in.Hard))
 out.Scopes = *(*[]v1.ResourceQuotaScope)(unsafe.Pointer(&in.Scopes))
 out.ScopeSelector = (*v1.ScopeSelector)(unsafe.Pointer(in.ScopeSelector))
 return nil
}
func Convert_core_ResourceQuotaSpec_To_v1_ResourceQuotaSpec(in *core.ResourceQuotaSpec, out *v1.ResourceQuotaSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_ResourceQuotaSpec_To_v1_ResourceQuotaSpec(in, out, s)
}
func autoConvert_v1_ResourceQuotaStatus_To_core_ResourceQuotaStatus(in *v1.ResourceQuotaStatus, out *core.ResourceQuotaStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Hard = *(*core.ResourceList)(unsafe.Pointer(&in.Hard))
 out.Used = *(*core.ResourceList)(unsafe.Pointer(&in.Used))
 return nil
}
func Convert_v1_ResourceQuotaStatus_To_core_ResourceQuotaStatus(in *v1.ResourceQuotaStatus, out *core.ResourceQuotaStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_ResourceQuotaStatus_To_core_ResourceQuotaStatus(in, out, s)
}
func autoConvert_core_ResourceQuotaStatus_To_v1_ResourceQuotaStatus(in *core.ResourceQuotaStatus, out *v1.ResourceQuotaStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Hard = *(*v1.ResourceList)(unsafe.Pointer(&in.Hard))
 out.Used = *(*v1.ResourceList)(unsafe.Pointer(&in.Used))
 return nil
}
func Convert_core_ResourceQuotaStatus_To_v1_ResourceQuotaStatus(in *core.ResourceQuotaStatus, out *v1.ResourceQuotaStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_ResourceQuotaStatus_To_v1_ResourceQuotaStatus(in, out, s)
}
func autoConvert_v1_ResourceRequirements_To_core_ResourceRequirements(in *v1.ResourceRequirements, out *core.ResourceRequirements, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Limits = *(*core.ResourceList)(unsafe.Pointer(&in.Limits))
 out.Requests = *(*core.ResourceList)(unsafe.Pointer(&in.Requests))
 return nil
}
func Convert_v1_ResourceRequirements_To_core_ResourceRequirements(in *v1.ResourceRequirements, out *core.ResourceRequirements, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_ResourceRequirements_To_core_ResourceRequirements(in, out, s)
}
func autoConvert_core_ResourceRequirements_To_v1_ResourceRequirements(in *core.ResourceRequirements, out *v1.ResourceRequirements, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Limits = *(*v1.ResourceList)(unsafe.Pointer(&in.Limits))
 out.Requests = *(*v1.ResourceList)(unsafe.Pointer(&in.Requests))
 return nil
}
func Convert_core_ResourceRequirements_To_v1_ResourceRequirements(in *core.ResourceRequirements, out *v1.ResourceRequirements, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_ResourceRequirements_To_v1_ResourceRequirements(in, out, s)
}
func autoConvert_v1_SELinuxOptions_To_core_SELinuxOptions(in *v1.SELinuxOptions, out *core.SELinuxOptions, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.User = in.User
 out.Role = in.Role
 out.Type = in.Type
 out.Level = in.Level
 return nil
}
func Convert_v1_SELinuxOptions_To_core_SELinuxOptions(in *v1.SELinuxOptions, out *core.SELinuxOptions, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_SELinuxOptions_To_core_SELinuxOptions(in, out, s)
}
func autoConvert_core_SELinuxOptions_To_v1_SELinuxOptions(in *core.SELinuxOptions, out *v1.SELinuxOptions, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.User = in.User
 out.Role = in.Role
 out.Type = in.Type
 out.Level = in.Level
 return nil
}
func Convert_core_SELinuxOptions_To_v1_SELinuxOptions(in *core.SELinuxOptions, out *v1.SELinuxOptions, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_SELinuxOptions_To_v1_SELinuxOptions(in, out, s)
}
func autoConvert_v1_ScaleIOPersistentVolumeSource_To_core_ScaleIOPersistentVolumeSource(in *v1.ScaleIOPersistentVolumeSource, out *core.ScaleIOPersistentVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Gateway = in.Gateway
 out.System = in.System
 out.SecretRef = (*core.SecretReference)(unsafe.Pointer(in.SecretRef))
 out.SSLEnabled = in.SSLEnabled
 out.ProtectionDomain = in.ProtectionDomain
 out.StoragePool = in.StoragePool
 out.StorageMode = in.StorageMode
 out.VolumeName = in.VolumeName
 out.FSType = in.FSType
 out.ReadOnly = in.ReadOnly
 return nil
}
func Convert_v1_ScaleIOPersistentVolumeSource_To_core_ScaleIOPersistentVolumeSource(in *v1.ScaleIOPersistentVolumeSource, out *core.ScaleIOPersistentVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_ScaleIOPersistentVolumeSource_To_core_ScaleIOPersistentVolumeSource(in, out, s)
}
func autoConvert_core_ScaleIOPersistentVolumeSource_To_v1_ScaleIOPersistentVolumeSource(in *core.ScaleIOPersistentVolumeSource, out *v1.ScaleIOPersistentVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Gateway = in.Gateway
 out.System = in.System
 out.SecretRef = (*v1.SecretReference)(unsafe.Pointer(in.SecretRef))
 out.SSLEnabled = in.SSLEnabled
 out.ProtectionDomain = in.ProtectionDomain
 out.StoragePool = in.StoragePool
 out.StorageMode = in.StorageMode
 out.VolumeName = in.VolumeName
 out.FSType = in.FSType
 out.ReadOnly = in.ReadOnly
 return nil
}
func Convert_core_ScaleIOPersistentVolumeSource_To_v1_ScaleIOPersistentVolumeSource(in *core.ScaleIOPersistentVolumeSource, out *v1.ScaleIOPersistentVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_ScaleIOPersistentVolumeSource_To_v1_ScaleIOPersistentVolumeSource(in, out, s)
}
func autoConvert_v1_ScaleIOVolumeSource_To_core_ScaleIOVolumeSource(in *v1.ScaleIOVolumeSource, out *core.ScaleIOVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Gateway = in.Gateway
 out.System = in.System
 out.SecretRef = (*core.LocalObjectReference)(unsafe.Pointer(in.SecretRef))
 out.SSLEnabled = in.SSLEnabled
 out.ProtectionDomain = in.ProtectionDomain
 out.StoragePool = in.StoragePool
 out.StorageMode = in.StorageMode
 out.VolumeName = in.VolumeName
 out.FSType = in.FSType
 out.ReadOnly = in.ReadOnly
 return nil
}
func Convert_v1_ScaleIOVolumeSource_To_core_ScaleIOVolumeSource(in *v1.ScaleIOVolumeSource, out *core.ScaleIOVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_ScaleIOVolumeSource_To_core_ScaleIOVolumeSource(in, out, s)
}
func autoConvert_core_ScaleIOVolumeSource_To_v1_ScaleIOVolumeSource(in *core.ScaleIOVolumeSource, out *v1.ScaleIOVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Gateway = in.Gateway
 out.System = in.System
 out.SecretRef = (*v1.LocalObjectReference)(unsafe.Pointer(in.SecretRef))
 out.SSLEnabled = in.SSLEnabled
 out.ProtectionDomain = in.ProtectionDomain
 out.StoragePool = in.StoragePool
 out.StorageMode = in.StorageMode
 out.VolumeName = in.VolumeName
 out.FSType = in.FSType
 out.ReadOnly = in.ReadOnly
 return nil
}
func Convert_core_ScaleIOVolumeSource_To_v1_ScaleIOVolumeSource(in *core.ScaleIOVolumeSource, out *v1.ScaleIOVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_ScaleIOVolumeSource_To_v1_ScaleIOVolumeSource(in, out, s)
}
func autoConvert_v1_ScopeSelector_To_core_ScopeSelector(in *v1.ScopeSelector, out *core.ScopeSelector, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.MatchExpressions = *(*[]core.ScopedResourceSelectorRequirement)(unsafe.Pointer(&in.MatchExpressions))
 return nil
}
func Convert_v1_ScopeSelector_To_core_ScopeSelector(in *v1.ScopeSelector, out *core.ScopeSelector, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_ScopeSelector_To_core_ScopeSelector(in, out, s)
}
func autoConvert_core_ScopeSelector_To_v1_ScopeSelector(in *core.ScopeSelector, out *v1.ScopeSelector, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.MatchExpressions = *(*[]v1.ScopedResourceSelectorRequirement)(unsafe.Pointer(&in.MatchExpressions))
 return nil
}
func Convert_core_ScopeSelector_To_v1_ScopeSelector(in *core.ScopeSelector, out *v1.ScopeSelector, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_ScopeSelector_To_v1_ScopeSelector(in, out, s)
}
func autoConvert_v1_ScopedResourceSelectorRequirement_To_core_ScopedResourceSelectorRequirement(in *v1.ScopedResourceSelectorRequirement, out *core.ScopedResourceSelectorRequirement, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ScopeName = core.ResourceQuotaScope(in.ScopeName)
 out.Operator = core.ScopeSelectorOperator(in.Operator)
 out.Values = *(*[]string)(unsafe.Pointer(&in.Values))
 return nil
}
func Convert_v1_ScopedResourceSelectorRequirement_To_core_ScopedResourceSelectorRequirement(in *v1.ScopedResourceSelectorRequirement, out *core.ScopedResourceSelectorRequirement, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_ScopedResourceSelectorRequirement_To_core_ScopedResourceSelectorRequirement(in, out, s)
}
func autoConvert_core_ScopedResourceSelectorRequirement_To_v1_ScopedResourceSelectorRequirement(in *core.ScopedResourceSelectorRequirement, out *v1.ScopedResourceSelectorRequirement, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ScopeName = v1.ResourceQuotaScope(in.ScopeName)
 out.Operator = v1.ScopeSelectorOperator(in.Operator)
 out.Values = *(*[]string)(unsafe.Pointer(&in.Values))
 return nil
}
func Convert_core_ScopedResourceSelectorRequirement_To_v1_ScopedResourceSelectorRequirement(in *core.ScopedResourceSelectorRequirement, out *v1.ScopedResourceSelectorRequirement, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_ScopedResourceSelectorRequirement_To_v1_ScopedResourceSelectorRequirement(in, out, s)
}
func autoConvert_v1_Secret_To_core_Secret(in *v1.Secret, out *core.Secret, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 out.Data = *(*map[string][]byte)(unsafe.Pointer(&in.Data))
 out.Type = core.SecretType(in.Type)
 return nil
}
func autoConvert_core_Secret_To_v1_Secret(in *core.Secret, out *v1.Secret, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 out.Data = *(*map[string][]byte)(unsafe.Pointer(&in.Data))
 out.Type = v1.SecretType(in.Type)
 return nil
}
func Convert_core_Secret_To_v1_Secret(in *core.Secret, out *v1.Secret, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_Secret_To_v1_Secret(in, out, s)
}
func autoConvert_v1_SecretEnvSource_To_core_SecretEnvSource(in *v1.SecretEnvSource, out *core.SecretEnvSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := Convert_v1_LocalObjectReference_To_core_LocalObjectReference(&in.LocalObjectReference, &out.LocalObjectReference, s); err != nil {
  return err
 }
 out.Optional = (*bool)(unsafe.Pointer(in.Optional))
 return nil
}
func Convert_v1_SecretEnvSource_To_core_SecretEnvSource(in *v1.SecretEnvSource, out *core.SecretEnvSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_SecretEnvSource_To_core_SecretEnvSource(in, out, s)
}
func autoConvert_core_SecretEnvSource_To_v1_SecretEnvSource(in *core.SecretEnvSource, out *v1.SecretEnvSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := Convert_core_LocalObjectReference_To_v1_LocalObjectReference(&in.LocalObjectReference, &out.LocalObjectReference, s); err != nil {
  return err
 }
 out.Optional = (*bool)(unsafe.Pointer(in.Optional))
 return nil
}
func Convert_core_SecretEnvSource_To_v1_SecretEnvSource(in *core.SecretEnvSource, out *v1.SecretEnvSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_SecretEnvSource_To_v1_SecretEnvSource(in, out, s)
}
func autoConvert_v1_SecretKeySelector_To_core_SecretKeySelector(in *v1.SecretKeySelector, out *core.SecretKeySelector, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := Convert_v1_LocalObjectReference_To_core_LocalObjectReference(&in.LocalObjectReference, &out.LocalObjectReference, s); err != nil {
  return err
 }
 out.Key = in.Key
 out.Optional = (*bool)(unsafe.Pointer(in.Optional))
 return nil
}
func Convert_v1_SecretKeySelector_To_core_SecretKeySelector(in *v1.SecretKeySelector, out *core.SecretKeySelector, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_SecretKeySelector_To_core_SecretKeySelector(in, out, s)
}
func autoConvert_core_SecretKeySelector_To_v1_SecretKeySelector(in *core.SecretKeySelector, out *v1.SecretKeySelector, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := Convert_core_LocalObjectReference_To_v1_LocalObjectReference(&in.LocalObjectReference, &out.LocalObjectReference, s); err != nil {
  return err
 }
 out.Key = in.Key
 out.Optional = (*bool)(unsafe.Pointer(in.Optional))
 return nil
}
func Convert_core_SecretKeySelector_To_v1_SecretKeySelector(in *core.SecretKeySelector, out *v1.SecretKeySelector, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_SecretKeySelector_To_v1_SecretKeySelector(in, out, s)
}
func autoConvert_v1_SecretList_To_core_SecretList(in *v1.SecretList, out *core.SecretList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]core.Secret, len(*in))
  for i := range *in {
   if err := Convert_v1_Secret_To_core_Secret(&(*in)[i], &(*out)[i], s); err != nil {
    return err
   }
  }
 } else {
  out.Items = nil
 }
 return nil
}
func Convert_v1_SecretList_To_core_SecretList(in *v1.SecretList, out *core.SecretList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_SecretList_To_core_SecretList(in, out, s)
}
func autoConvert_core_SecretList_To_v1_SecretList(in *core.SecretList, out *v1.SecretList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]v1.Secret, len(*in))
  for i := range *in {
   if err := Convert_core_Secret_To_v1_Secret(&(*in)[i], &(*out)[i], s); err != nil {
    return err
   }
  }
 } else {
  out.Items = nil
 }
 return nil
}
func Convert_core_SecretList_To_v1_SecretList(in *core.SecretList, out *v1.SecretList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_SecretList_To_v1_SecretList(in, out, s)
}
func autoConvert_v1_SecretProjection_To_core_SecretProjection(in *v1.SecretProjection, out *core.SecretProjection, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := Convert_v1_LocalObjectReference_To_core_LocalObjectReference(&in.LocalObjectReference, &out.LocalObjectReference, s); err != nil {
  return err
 }
 out.Items = *(*[]core.KeyToPath)(unsafe.Pointer(&in.Items))
 out.Optional = (*bool)(unsafe.Pointer(in.Optional))
 return nil
}
func Convert_v1_SecretProjection_To_core_SecretProjection(in *v1.SecretProjection, out *core.SecretProjection, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_SecretProjection_To_core_SecretProjection(in, out, s)
}
func autoConvert_core_SecretProjection_To_v1_SecretProjection(in *core.SecretProjection, out *v1.SecretProjection, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := Convert_core_LocalObjectReference_To_v1_LocalObjectReference(&in.LocalObjectReference, &out.LocalObjectReference, s); err != nil {
  return err
 }
 out.Items = *(*[]v1.KeyToPath)(unsafe.Pointer(&in.Items))
 out.Optional = (*bool)(unsafe.Pointer(in.Optional))
 return nil
}
func Convert_core_SecretProjection_To_v1_SecretProjection(in *core.SecretProjection, out *v1.SecretProjection, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_SecretProjection_To_v1_SecretProjection(in, out, s)
}
func autoConvert_v1_SecretReference_To_core_SecretReference(in *v1.SecretReference, out *core.SecretReference, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Name = in.Name
 out.Namespace = in.Namespace
 return nil
}
func Convert_v1_SecretReference_To_core_SecretReference(in *v1.SecretReference, out *core.SecretReference, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_SecretReference_To_core_SecretReference(in, out, s)
}
func autoConvert_core_SecretReference_To_v1_SecretReference(in *core.SecretReference, out *v1.SecretReference, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Name = in.Name
 out.Namespace = in.Namespace
 return nil
}
func Convert_core_SecretReference_To_v1_SecretReference(in *core.SecretReference, out *v1.SecretReference, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_SecretReference_To_v1_SecretReference(in, out, s)
}
func autoConvert_v1_SecretVolumeSource_To_core_SecretVolumeSource(in *v1.SecretVolumeSource, out *core.SecretVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.SecretName = in.SecretName
 out.Items = *(*[]core.KeyToPath)(unsafe.Pointer(&in.Items))
 out.DefaultMode = (*int32)(unsafe.Pointer(in.DefaultMode))
 out.Optional = (*bool)(unsafe.Pointer(in.Optional))
 return nil
}
func Convert_v1_SecretVolumeSource_To_core_SecretVolumeSource(in *v1.SecretVolumeSource, out *core.SecretVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_SecretVolumeSource_To_core_SecretVolumeSource(in, out, s)
}
func autoConvert_core_SecretVolumeSource_To_v1_SecretVolumeSource(in *core.SecretVolumeSource, out *v1.SecretVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.SecretName = in.SecretName
 out.Items = *(*[]v1.KeyToPath)(unsafe.Pointer(&in.Items))
 out.DefaultMode = (*int32)(unsafe.Pointer(in.DefaultMode))
 out.Optional = (*bool)(unsafe.Pointer(in.Optional))
 return nil
}
func Convert_core_SecretVolumeSource_To_v1_SecretVolumeSource(in *core.SecretVolumeSource, out *v1.SecretVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_SecretVolumeSource_To_v1_SecretVolumeSource(in, out, s)
}
func autoConvert_v1_SecurityContext_To_core_SecurityContext(in *v1.SecurityContext, out *core.SecurityContext, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Capabilities = (*core.Capabilities)(unsafe.Pointer(in.Capabilities))
 out.Privileged = (*bool)(unsafe.Pointer(in.Privileged))
 out.SELinuxOptions = (*core.SELinuxOptions)(unsafe.Pointer(in.SELinuxOptions))
 out.RunAsUser = (*int64)(unsafe.Pointer(in.RunAsUser))
 out.RunAsGroup = (*int64)(unsafe.Pointer(in.RunAsGroup))
 out.RunAsNonRoot = (*bool)(unsafe.Pointer(in.RunAsNonRoot))
 out.ReadOnlyRootFilesystem = (*bool)(unsafe.Pointer(in.ReadOnlyRootFilesystem))
 out.AllowPrivilegeEscalation = (*bool)(unsafe.Pointer(in.AllowPrivilegeEscalation))
 out.ProcMount = (*core.ProcMountType)(unsafe.Pointer(in.ProcMount))
 return nil
}
func Convert_v1_SecurityContext_To_core_SecurityContext(in *v1.SecurityContext, out *core.SecurityContext, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_SecurityContext_To_core_SecurityContext(in, out, s)
}
func autoConvert_core_SecurityContext_To_v1_SecurityContext(in *core.SecurityContext, out *v1.SecurityContext, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Capabilities = (*v1.Capabilities)(unsafe.Pointer(in.Capabilities))
 out.Privileged = (*bool)(unsafe.Pointer(in.Privileged))
 out.SELinuxOptions = (*v1.SELinuxOptions)(unsafe.Pointer(in.SELinuxOptions))
 out.RunAsUser = (*int64)(unsafe.Pointer(in.RunAsUser))
 out.RunAsGroup = (*int64)(unsafe.Pointer(in.RunAsGroup))
 out.RunAsNonRoot = (*bool)(unsafe.Pointer(in.RunAsNonRoot))
 out.ReadOnlyRootFilesystem = (*bool)(unsafe.Pointer(in.ReadOnlyRootFilesystem))
 out.AllowPrivilegeEscalation = (*bool)(unsafe.Pointer(in.AllowPrivilegeEscalation))
 out.ProcMount = (*v1.ProcMountType)(unsafe.Pointer(in.ProcMount))
 return nil
}
func autoConvert_v1_SerializedReference_To_core_SerializedReference(in *v1.SerializedReference, out *core.SerializedReference, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := Convert_v1_ObjectReference_To_core_ObjectReference(&in.Reference, &out.Reference, s); err != nil {
  return err
 }
 return nil
}
func Convert_v1_SerializedReference_To_core_SerializedReference(in *v1.SerializedReference, out *core.SerializedReference, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_SerializedReference_To_core_SerializedReference(in, out, s)
}
func autoConvert_core_SerializedReference_To_v1_SerializedReference(in *core.SerializedReference, out *v1.SerializedReference, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := Convert_core_ObjectReference_To_v1_ObjectReference(&in.Reference, &out.Reference, s); err != nil {
  return err
 }
 return nil
}
func Convert_core_SerializedReference_To_v1_SerializedReference(in *core.SerializedReference, out *v1.SerializedReference, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_SerializedReference_To_v1_SerializedReference(in, out, s)
}
func autoConvert_v1_Service_To_core_Service(in *v1.Service, out *core.Service, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_v1_ServiceSpec_To_core_ServiceSpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 if err := Convert_v1_ServiceStatus_To_core_ServiceStatus(&in.Status, &out.Status, s); err != nil {
  return err
 }
 return nil
}
func Convert_v1_Service_To_core_Service(in *v1.Service, out *core.Service, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_Service_To_core_Service(in, out, s)
}
func autoConvert_core_Service_To_v1_Service(in *core.Service, out *v1.Service, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_core_ServiceSpec_To_v1_ServiceSpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 if err := Convert_core_ServiceStatus_To_v1_ServiceStatus(&in.Status, &out.Status, s); err != nil {
  return err
 }
 return nil
}
func Convert_core_Service_To_v1_Service(in *core.Service, out *v1.Service, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_Service_To_v1_Service(in, out, s)
}
func autoConvert_v1_ServiceAccount_To_core_ServiceAccount(in *v1.ServiceAccount, out *core.ServiceAccount, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 out.Secrets = *(*[]core.ObjectReference)(unsafe.Pointer(&in.Secrets))
 out.ImagePullSecrets = *(*[]core.LocalObjectReference)(unsafe.Pointer(&in.ImagePullSecrets))
 out.AutomountServiceAccountToken = (*bool)(unsafe.Pointer(in.AutomountServiceAccountToken))
 return nil
}
func Convert_v1_ServiceAccount_To_core_ServiceAccount(in *v1.ServiceAccount, out *core.ServiceAccount, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_ServiceAccount_To_core_ServiceAccount(in, out, s)
}
func autoConvert_core_ServiceAccount_To_v1_ServiceAccount(in *core.ServiceAccount, out *v1.ServiceAccount, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 out.Secrets = *(*[]v1.ObjectReference)(unsafe.Pointer(&in.Secrets))
 out.ImagePullSecrets = *(*[]v1.LocalObjectReference)(unsafe.Pointer(&in.ImagePullSecrets))
 out.AutomountServiceAccountToken = (*bool)(unsafe.Pointer(in.AutomountServiceAccountToken))
 return nil
}
func Convert_core_ServiceAccount_To_v1_ServiceAccount(in *core.ServiceAccount, out *v1.ServiceAccount, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_ServiceAccount_To_v1_ServiceAccount(in, out, s)
}
func autoConvert_v1_ServiceAccountList_To_core_ServiceAccountList(in *v1.ServiceAccountList, out *core.ServiceAccountList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 out.Items = *(*[]core.ServiceAccount)(unsafe.Pointer(&in.Items))
 return nil
}
func Convert_v1_ServiceAccountList_To_core_ServiceAccountList(in *v1.ServiceAccountList, out *core.ServiceAccountList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_ServiceAccountList_To_core_ServiceAccountList(in, out, s)
}
func autoConvert_core_ServiceAccountList_To_v1_ServiceAccountList(in *core.ServiceAccountList, out *v1.ServiceAccountList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 out.Items = *(*[]v1.ServiceAccount)(unsafe.Pointer(&in.Items))
 return nil
}
func Convert_core_ServiceAccountList_To_v1_ServiceAccountList(in *core.ServiceAccountList, out *v1.ServiceAccountList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_ServiceAccountList_To_v1_ServiceAccountList(in, out, s)
}
func autoConvert_v1_ServiceAccountTokenProjection_To_core_ServiceAccountTokenProjection(in *v1.ServiceAccountTokenProjection, out *core.ServiceAccountTokenProjection, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Audience = in.Audience
 if err := metav1.Convert_Pointer_int64_To_int64(&in.ExpirationSeconds, &out.ExpirationSeconds, s); err != nil {
  return err
 }
 out.Path = in.Path
 return nil
}
func Convert_v1_ServiceAccountTokenProjection_To_core_ServiceAccountTokenProjection(in *v1.ServiceAccountTokenProjection, out *core.ServiceAccountTokenProjection, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_ServiceAccountTokenProjection_To_core_ServiceAccountTokenProjection(in, out, s)
}
func autoConvert_core_ServiceAccountTokenProjection_To_v1_ServiceAccountTokenProjection(in *core.ServiceAccountTokenProjection, out *v1.ServiceAccountTokenProjection, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Audience = in.Audience
 if err := metav1.Convert_int64_To_Pointer_int64(&in.ExpirationSeconds, &out.ExpirationSeconds, s); err != nil {
  return err
 }
 out.Path = in.Path
 return nil
}
func Convert_core_ServiceAccountTokenProjection_To_v1_ServiceAccountTokenProjection(in *core.ServiceAccountTokenProjection, out *v1.ServiceAccountTokenProjection, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_ServiceAccountTokenProjection_To_v1_ServiceAccountTokenProjection(in, out, s)
}
func autoConvert_v1_ServiceList_To_core_ServiceList(in *v1.ServiceList, out *core.ServiceList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]core.Service, len(*in))
  for i := range *in {
   if err := Convert_v1_Service_To_core_Service(&(*in)[i], &(*out)[i], s); err != nil {
    return err
   }
  }
 } else {
  out.Items = nil
 }
 return nil
}
func Convert_v1_ServiceList_To_core_ServiceList(in *v1.ServiceList, out *core.ServiceList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_ServiceList_To_core_ServiceList(in, out, s)
}
func autoConvert_core_ServiceList_To_v1_ServiceList(in *core.ServiceList, out *v1.ServiceList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]v1.Service, len(*in))
  for i := range *in {
   if err := Convert_core_Service_To_v1_Service(&(*in)[i], &(*out)[i], s); err != nil {
    return err
   }
  }
 } else {
  out.Items = nil
 }
 return nil
}
func Convert_core_ServiceList_To_v1_ServiceList(in *core.ServiceList, out *v1.ServiceList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_ServiceList_To_v1_ServiceList(in, out, s)
}
func autoConvert_v1_ServicePort_To_core_ServicePort(in *v1.ServicePort, out *core.ServicePort, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Name = in.Name
 out.Protocol = core.Protocol(in.Protocol)
 out.Port = in.Port
 out.TargetPort = in.TargetPort
 out.NodePort = in.NodePort
 return nil
}
func Convert_v1_ServicePort_To_core_ServicePort(in *v1.ServicePort, out *core.ServicePort, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_ServicePort_To_core_ServicePort(in, out, s)
}
func autoConvert_core_ServicePort_To_v1_ServicePort(in *core.ServicePort, out *v1.ServicePort, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Name = in.Name
 out.Protocol = v1.Protocol(in.Protocol)
 out.Port = in.Port
 out.TargetPort = in.TargetPort
 out.NodePort = in.NodePort
 return nil
}
func Convert_core_ServicePort_To_v1_ServicePort(in *core.ServicePort, out *v1.ServicePort, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_ServicePort_To_v1_ServicePort(in, out, s)
}
func autoConvert_v1_ServiceProxyOptions_To_core_ServiceProxyOptions(in *v1.ServiceProxyOptions, out *core.ServiceProxyOptions, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Path = in.Path
 return nil
}
func Convert_v1_ServiceProxyOptions_To_core_ServiceProxyOptions(in *v1.ServiceProxyOptions, out *core.ServiceProxyOptions, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_ServiceProxyOptions_To_core_ServiceProxyOptions(in, out, s)
}
func autoConvert_core_ServiceProxyOptions_To_v1_ServiceProxyOptions(in *core.ServiceProxyOptions, out *v1.ServiceProxyOptions, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Path = in.Path
 return nil
}
func Convert_core_ServiceProxyOptions_To_v1_ServiceProxyOptions(in *core.ServiceProxyOptions, out *v1.ServiceProxyOptions, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_ServiceProxyOptions_To_v1_ServiceProxyOptions(in, out, s)
}
func autoConvert_v1_ServiceSpec_To_core_ServiceSpec(in *v1.ServiceSpec, out *core.ServiceSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Ports = *(*[]core.ServicePort)(unsafe.Pointer(&in.Ports))
 out.Selector = *(*map[string]string)(unsafe.Pointer(&in.Selector))
 out.ClusterIP = in.ClusterIP
 out.Type = core.ServiceType(in.Type)
 out.ExternalIPs = *(*[]string)(unsafe.Pointer(&in.ExternalIPs))
 out.SessionAffinity = core.ServiceAffinity(in.SessionAffinity)
 out.LoadBalancerIP = in.LoadBalancerIP
 out.LoadBalancerSourceRanges = *(*[]string)(unsafe.Pointer(&in.LoadBalancerSourceRanges))
 out.ExternalName = in.ExternalName
 out.ExternalTrafficPolicy = core.ServiceExternalTrafficPolicyType(in.ExternalTrafficPolicy)
 out.HealthCheckNodePort = in.HealthCheckNodePort
 out.PublishNotReadyAddresses = in.PublishNotReadyAddresses
 out.SessionAffinityConfig = (*core.SessionAffinityConfig)(unsafe.Pointer(in.SessionAffinityConfig))
 return nil
}
func Convert_v1_ServiceSpec_To_core_ServiceSpec(in *v1.ServiceSpec, out *core.ServiceSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_ServiceSpec_To_core_ServiceSpec(in, out, s)
}
func autoConvert_core_ServiceSpec_To_v1_ServiceSpec(in *core.ServiceSpec, out *v1.ServiceSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Type = v1.ServiceType(in.Type)
 out.Ports = *(*[]v1.ServicePort)(unsafe.Pointer(&in.Ports))
 out.Selector = *(*map[string]string)(unsafe.Pointer(&in.Selector))
 out.ClusterIP = in.ClusterIP
 out.ExternalName = in.ExternalName
 out.ExternalIPs = *(*[]string)(unsafe.Pointer(&in.ExternalIPs))
 out.LoadBalancerIP = in.LoadBalancerIP
 out.SessionAffinity = v1.ServiceAffinity(in.SessionAffinity)
 out.SessionAffinityConfig = (*v1.SessionAffinityConfig)(unsafe.Pointer(in.SessionAffinityConfig))
 out.LoadBalancerSourceRanges = *(*[]string)(unsafe.Pointer(&in.LoadBalancerSourceRanges))
 out.ExternalTrafficPolicy = v1.ServiceExternalTrafficPolicyType(in.ExternalTrafficPolicy)
 out.HealthCheckNodePort = in.HealthCheckNodePort
 out.PublishNotReadyAddresses = in.PublishNotReadyAddresses
 return nil
}
func Convert_core_ServiceSpec_To_v1_ServiceSpec(in *core.ServiceSpec, out *v1.ServiceSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_ServiceSpec_To_v1_ServiceSpec(in, out, s)
}
func autoConvert_v1_ServiceStatus_To_core_ServiceStatus(in *v1.ServiceStatus, out *core.ServiceStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := Convert_v1_LoadBalancerStatus_To_core_LoadBalancerStatus(&in.LoadBalancer, &out.LoadBalancer, s); err != nil {
  return err
 }
 return nil
}
func Convert_v1_ServiceStatus_To_core_ServiceStatus(in *v1.ServiceStatus, out *core.ServiceStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_ServiceStatus_To_core_ServiceStatus(in, out, s)
}
func autoConvert_core_ServiceStatus_To_v1_ServiceStatus(in *core.ServiceStatus, out *v1.ServiceStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := Convert_core_LoadBalancerStatus_To_v1_LoadBalancerStatus(&in.LoadBalancer, &out.LoadBalancer, s); err != nil {
  return err
 }
 return nil
}
func Convert_core_ServiceStatus_To_v1_ServiceStatus(in *core.ServiceStatus, out *v1.ServiceStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_ServiceStatus_To_v1_ServiceStatus(in, out, s)
}
func autoConvert_v1_SessionAffinityConfig_To_core_SessionAffinityConfig(in *v1.SessionAffinityConfig, out *core.SessionAffinityConfig, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ClientIP = (*core.ClientIPConfig)(unsafe.Pointer(in.ClientIP))
 return nil
}
func Convert_v1_SessionAffinityConfig_To_core_SessionAffinityConfig(in *v1.SessionAffinityConfig, out *core.SessionAffinityConfig, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_SessionAffinityConfig_To_core_SessionAffinityConfig(in, out, s)
}
func autoConvert_core_SessionAffinityConfig_To_v1_SessionAffinityConfig(in *core.SessionAffinityConfig, out *v1.SessionAffinityConfig, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ClientIP = (*v1.ClientIPConfig)(unsafe.Pointer(in.ClientIP))
 return nil
}
func Convert_core_SessionAffinityConfig_To_v1_SessionAffinityConfig(in *core.SessionAffinityConfig, out *v1.SessionAffinityConfig, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_SessionAffinityConfig_To_v1_SessionAffinityConfig(in, out, s)
}
func autoConvert_v1_StorageOSPersistentVolumeSource_To_core_StorageOSPersistentVolumeSource(in *v1.StorageOSPersistentVolumeSource, out *core.StorageOSPersistentVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.VolumeName = in.VolumeName
 out.VolumeNamespace = in.VolumeNamespace
 out.FSType = in.FSType
 out.ReadOnly = in.ReadOnly
 out.SecretRef = (*core.ObjectReference)(unsafe.Pointer(in.SecretRef))
 return nil
}
func Convert_v1_StorageOSPersistentVolumeSource_To_core_StorageOSPersistentVolumeSource(in *v1.StorageOSPersistentVolumeSource, out *core.StorageOSPersistentVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_StorageOSPersistentVolumeSource_To_core_StorageOSPersistentVolumeSource(in, out, s)
}
func autoConvert_core_StorageOSPersistentVolumeSource_To_v1_StorageOSPersistentVolumeSource(in *core.StorageOSPersistentVolumeSource, out *v1.StorageOSPersistentVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.VolumeName = in.VolumeName
 out.VolumeNamespace = in.VolumeNamespace
 out.FSType = in.FSType
 out.ReadOnly = in.ReadOnly
 out.SecretRef = (*v1.ObjectReference)(unsafe.Pointer(in.SecretRef))
 return nil
}
func Convert_core_StorageOSPersistentVolumeSource_To_v1_StorageOSPersistentVolumeSource(in *core.StorageOSPersistentVolumeSource, out *v1.StorageOSPersistentVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_StorageOSPersistentVolumeSource_To_v1_StorageOSPersistentVolumeSource(in, out, s)
}
func autoConvert_v1_StorageOSVolumeSource_To_core_StorageOSVolumeSource(in *v1.StorageOSVolumeSource, out *core.StorageOSVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.VolumeName = in.VolumeName
 out.VolumeNamespace = in.VolumeNamespace
 out.FSType = in.FSType
 out.ReadOnly = in.ReadOnly
 out.SecretRef = (*core.LocalObjectReference)(unsafe.Pointer(in.SecretRef))
 return nil
}
func Convert_v1_StorageOSVolumeSource_To_core_StorageOSVolumeSource(in *v1.StorageOSVolumeSource, out *core.StorageOSVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_StorageOSVolumeSource_To_core_StorageOSVolumeSource(in, out, s)
}
func autoConvert_core_StorageOSVolumeSource_To_v1_StorageOSVolumeSource(in *core.StorageOSVolumeSource, out *v1.StorageOSVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.VolumeName = in.VolumeName
 out.VolumeNamespace = in.VolumeNamespace
 out.FSType = in.FSType
 out.ReadOnly = in.ReadOnly
 out.SecretRef = (*v1.LocalObjectReference)(unsafe.Pointer(in.SecretRef))
 return nil
}
func Convert_core_StorageOSVolumeSource_To_v1_StorageOSVolumeSource(in *core.StorageOSVolumeSource, out *v1.StorageOSVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_StorageOSVolumeSource_To_v1_StorageOSVolumeSource(in, out, s)
}
func autoConvert_v1_Sysctl_To_core_Sysctl(in *v1.Sysctl, out *core.Sysctl, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Name = in.Name
 out.Value = in.Value
 return nil
}
func Convert_v1_Sysctl_To_core_Sysctl(in *v1.Sysctl, out *core.Sysctl, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_Sysctl_To_core_Sysctl(in, out, s)
}
func autoConvert_core_Sysctl_To_v1_Sysctl(in *core.Sysctl, out *v1.Sysctl, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Name = in.Name
 out.Value = in.Value
 return nil
}
func Convert_core_Sysctl_To_v1_Sysctl(in *core.Sysctl, out *v1.Sysctl, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_Sysctl_To_v1_Sysctl(in, out, s)
}
func autoConvert_v1_TCPSocketAction_To_core_TCPSocketAction(in *v1.TCPSocketAction, out *core.TCPSocketAction, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Port = in.Port
 out.Host = in.Host
 return nil
}
func Convert_v1_TCPSocketAction_To_core_TCPSocketAction(in *v1.TCPSocketAction, out *core.TCPSocketAction, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_TCPSocketAction_To_core_TCPSocketAction(in, out, s)
}
func autoConvert_core_TCPSocketAction_To_v1_TCPSocketAction(in *core.TCPSocketAction, out *v1.TCPSocketAction, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Port = in.Port
 out.Host = in.Host
 return nil
}
func Convert_core_TCPSocketAction_To_v1_TCPSocketAction(in *core.TCPSocketAction, out *v1.TCPSocketAction, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_TCPSocketAction_To_v1_TCPSocketAction(in, out, s)
}
func autoConvert_v1_Taint_To_core_Taint(in *v1.Taint, out *core.Taint, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Key = in.Key
 out.Value = in.Value
 out.Effect = core.TaintEffect(in.Effect)
 out.TimeAdded = (*metav1.Time)(unsafe.Pointer(in.TimeAdded))
 return nil
}
func Convert_v1_Taint_To_core_Taint(in *v1.Taint, out *core.Taint, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_Taint_To_core_Taint(in, out, s)
}
func autoConvert_core_Taint_To_v1_Taint(in *core.Taint, out *v1.Taint, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Key = in.Key
 out.Value = in.Value
 out.Effect = v1.TaintEffect(in.Effect)
 out.TimeAdded = (*metav1.Time)(unsafe.Pointer(in.TimeAdded))
 return nil
}
func Convert_core_Taint_To_v1_Taint(in *core.Taint, out *v1.Taint, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_Taint_To_v1_Taint(in, out, s)
}
func autoConvert_v1_Toleration_To_core_Toleration(in *v1.Toleration, out *core.Toleration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Key = in.Key
 out.Operator = core.TolerationOperator(in.Operator)
 out.Value = in.Value
 out.Effect = core.TaintEffect(in.Effect)
 out.TolerationSeconds = (*int64)(unsafe.Pointer(in.TolerationSeconds))
 return nil
}
func Convert_v1_Toleration_To_core_Toleration(in *v1.Toleration, out *core.Toleration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_Toleration_To_core_Toleration(in, out, s)
}
func autoConvert_core_Toleration_To_v1_Toleration(in *core.Toleration, out *v1.Toleration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Key = in.Key
 out.Operator = v1.TolerationOperator(in.Operator)
 out.Value = in.Value
 out.Effect = v1.TaintEffect(in.Effect)
 out.TolerationSeconds = (*int64)(unsafe.Pointer(in.TolerationSeconds))
 return nil
}
func Convert_core_Toleration_To_v1_Toleration(in *core.Toleration, out *v1.Toleration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_Toleration_To_v1_Toleration(in, out, s)
}
func autoConvert_v1_TopologySelectorLabelRequirement_To_core_TopologySelectorLabelRequirement(in *v1.TopologySelectorLabelRequirement, out *core.TopologySelectorLabelRequirement, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Key = in.Key
 out.Values = *(*[]string)(unsafe.Pointer(&in.Values))
 return nil
}
func Convert_v1_TopologySelectorLabelRequirement_To_core_TopologySelectorLabelRequirement(in *v1.TopologySelectorLabelRequirement, out *core.TopologySelectorLabelRequirement, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_TopologySelectorLabelRequirement_To_core_TopologySelectorLabelRequirement(in, out, s)
}
func autoConvert_core_TopologySelectorLabelRequirement_To_v1_TopologySelectorLabelRequirement(in *core.TopologySelectorLabelRequirement, out *v1.TopologySelectorLabelRequirement, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Key = in.Key
 out.Values = *(*[]string)(unsafe.Pointer(&in.Values))
 return nil
}
func Convert_core_TopologySelectorLabelRequirement_To_v1_TopologySelectorLabelRequirement(in *core.TopologySelectorLabelRequirement, out *v1.TopologySelectorLabelRequirement, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_TopologySelectorLabelRequirement_To_v1_TopologySelectorLabelRequirement(in, out, s)
}
func autoConvert_v1_TopologySelectorTerm_To_core_TopologySelectorTerm(in *v1.TopologySelectorTerm, out *core.TopologySelectorTerm, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.MatchLabelExpressions = *(*[]core.TopologySelectorLabelRequirement)(unsafe.Pointer(&in.MatchLabelExpressions))
 return nil
}
func Convert_v1_TopologySelectorTerm_To_core_TopologySelectorTerm(in *v1.TopologySelectorTerm, out *core.TopologySelectorTerm, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_TopologySelectorTerm_To_core_TopologySelectorTerm(in, out, s)
}
func autoConvert_core_TopologySelectorTerm_To_v1_TopologySelectorTerm(in *core.TopologySelectorTerm, out *v1.TopologySelectorTerm, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.MatchLabelExpressions = *(*[]v1.TopologySelectorLabelRequirement)(unsafe.Pointer(&in.MatchLabelExpressions))
 return nil
}
func Convert_core_TopologySelectorTerm_To_v1_TopologySelectorTerm(in *core.TopologySelectorTerm, out *v1.TopologySelectorTerm, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_TopologySelectorTerm_To_v1_TopologySelectorTerm(in, out, s)
}
func autoConvert_v1_TypedLocalObjectReference_To_core_TypedLocalObjectReference(in *v1.TypedLocalObjectReference, out *core.TypedLocalObjectReference, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.APIGroup = (*string)(unsafe.Pointer(in.APIGroup))
 out.Kind = in.Kind
 out.Name = in.Name
 return nil
}
func Convert_v1_TypedLocalObjectReference_To_core_TypedLocalObjectReference(in *v1.TypedLocalObjectReference, out *core.TypedLocalObjectReference, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_TypedLocalObjectReference_To_core_TypedLocalObjectReference(in, out, s)
}
func autoConvert_core_TypedLocalObjectReference_To_v1_TypedLocalObjectReference(in *core.TypedLocalObjectReference, out *v1.TypedLocalObjectReference, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.APIGroup = (*string)(unsafe.Pointer(in.APIGroup))
 out.Kind = in.Kind
 out.Name = in.Name
 return nil
}
func Convert_core_TypedLocalObjectReference_To_v1_TypedLocalObjectReference(in *core.TypedLocalObjectReference, out *v1.TypedLocalObjectReference, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_TypedLocalObjectReference_To_v1_TypedLocalObjectReference(in, out, s)
}
func autoConvert_v1_Volume_To_core_Volume(in *v1.Volume, out *core.Volume, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Name = in.Name
 if err := Convert_v1_VolumeSource_To_core_VolumeSource(&in.VolumeSource, &out.VolumeSource, s); err != nil {
  return err
 }
 return nil
}
func Convert_v1_Volume_To_core_Volume(in *v1.Volume, out *core.Volume, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_Volume_To_core_Volume(in, out, s)
}
func autoConvert_core_Volume_To_v1_Volume(in *core.Volume, out *v1.Volume, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Name = in.Name
 if err := Convert_core_VolumeSource_To_v1_VolumeSource(&in.VolumeSource, &out.VolumeSource, s); err != nil {
  return err
 }
 return nil
}
func Convert_core_Volume_To_v1_Volume(in *core.Volume, out *v1.Volume, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_Volume_To_v1_Volume(in, out, s)
}
func autoConvert_v1_VolumeDevice_To_core_VolumeDevice(in *v1.VolumeDevice, out *core.VolumeDevice, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Name = in.Name
 out.DevicePath = in.DevicePath
 return nil
}
func Convert_v1_VolumeDevice_To_core_VolumeDevice(in *v1.VolumeDevice, out *core.VolumeDevice, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_VolumeDevice_To_core_VolumeDevice(in, out, s)
}
func autoConvert_core_VolumeDevice_To_v1_VolumeDevice(in *core.VolumeDevice, out *v1.VolumeDevice, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Name = in.Name
 out.DevicePath = in.DevicePath
 return nil
}
func Convert_core_VolumeDevice_To_v1_VolumeDevice(in *core.VolumeDevice, out *v1.VolumeDevice, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_VolumeDevice_To_v1_VolumeDevice(in, out, s)
}
func autoConvert_v1_VolumeMount_To_core_VolumeMount(in *v1.VolumeMount, out *core.VolumeMount, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Name = in.Name
 out.ReadOnly = in.ReadOnly
 out.MountPath = in.MountPath
 out.SubPath = in.SubPath
 out.MountPropagation = (*core.MountPropagationMode)(unsafe.Pointer(in.MountPropagation))
 return nil
}
func Convert_v1_VolumeMount_To_core_VolumeMount(in *v1.VolumeMount, out *core.VolumeMount, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_VolumeMount_To_core_VolumeMount(in, out, s)
}
func autoConvert_core_VolumeMount_To_v1_VolumeMount(in *core.VolumeMount, out *v1.VolumeMount, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Name = in.Name
 out.ReadOnly = in.ReadOnly
 out.MountPath = in.MountPath
 out.SubPath = in.SubPath
 out.MountPropagation = (*v1.MountPropagationMode)(unsafe.Pointer(in.MountPropagation))
 return nil
}
func Convert_core_VolumeMount_To_v1_VolumeMount(in *core.VolumeMount, out *v1.VolumeMount, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_VolumeMount_To_v1_VolumeMount(in, out, s)
}
func autoConvert_v1_VolumeNodeAffinity_To_core_VolumeNodeAffinity(in *v1.VolumeNodeAffinity, out *core.VolumeNodeAffinity, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Required = (*core.NodeSelector)(unsafe.Pointer(in.Required))
 return nil
}
func Convert_v1_VolumeNodeAffinity_To_core_VolumeNodeAffinity(in *v1.VolumeNodeAffinity, out *core.VolumeNodeAffinity, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_VolumeNodeAffinity_To_core_VolumeNodeAffinity(in, out, s)
}
func autoConvert_core_VolumeNodeAffinity_To_v1_VolumeNodeAffinity(in *core.VolumeNodeAffinity, out *v1.VolumeNodeAffinity, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Required = (*v1.NodeSelector)(unsafe.Pointer(in.Required))
 return nil
}
func Convert_core_VolumeNodeAffinity_To_v1_VolumeNodeAffinity(in *core.VolumeNodeAffinity, out *v1.VolumeNodeAffinity, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_VolumeNodeAffinity_To_v1_VolumeNodeAffinity(in, out, s)
}
func autoConvert_v1_VolumeProjection_To_core_VolumeProjection(in *v1.VolumeProjection, out *core.VolumeProjection, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Secret = (*core.SecretProjection)(unsafe.Pointer(in.Secret))
 out.DownwardAPI = (*core.DownwardAPIProjection)(unsafe.Pointer(in.DownwardAPI))
 out.ConfigMap = (*core.ConfigMapProjection)(unsafe.Pointer(in.ConfigMap))
 if in.ServiceAccountToken != nil {
  in, out := &in.ServiceAccountToken, &out.ServiceAccountToken
  *out = new(core.ServiceAccountTokenProjection)
  if err := Convert_v1_ServiceAccountTokenProjection_To_core_ServiceAccountTokenProjection(*in, *out, s); err != nil {
   return err
  }
 } else {
  out.ServiceAccountToken = nil
 }
 return nil
}
func Convert_v1_VolumeProjection_To_core_VolumeProjection(in *v1.VolumeProjection, out *core.VolumeProjection, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_VolumeProjection_To_core_VolumeProjection(in, out, s)
}
func autoConvert_core_VolumeProjection_To_v1_VolumeProjection(in *core.VolumeProjection, out *v1.VolumeProjection, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Secret = (*v1.SecretProjection)(unsafe.Pointer(in.Secret))
 out.DownwardAPI = (*v1.DownwardAPIProjection)(unsafe.Pointer(in.DownwardAPI))
 out.ConfigMap = (*v1.ConfigMapProjection)(unsafe.Pointer(in.ConfigMap))
 if in.ServiceAccountToken != nil {
  in, out := &in.ServiceAccountToken, &out.ServiceAccountToken
  *out = new(v1.ServiceAccountTokenProjection)
  if err := Convert_core_ServiceAccountTokenProjection_To_v1_ServiceAccountTokenProjection(*in, *out, s); err != nil {
   return err
  }
 } else {
  out.ServiceAccountToken = nil
 }
 return nil
}
func Convert_core_VolumeProjection_To_v1_VolumeProjection(in *core.VolumeProjection, out *v1.VolumeProjection, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_VolumeProjection_To_v1_VolumeProjection(in, out, s)
}
func autoConvert_v1_VolumeSource_To_core_VolumeSource(in *v1.VolumeSource, out *core.VolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.HostPath = (*core.HostPathVolumeSource)(unsafe.Pointer(in.HostPath))
 out.EmptyDir = (*core.EmptyDirVolumeSource)(unsafe.Pointer(in.EmptyDir))
 out.GCEPersistentDisk = (*core.GCEPersistentDiskVolumeSource)(unsafe.Pointer(in.GCEPersistentDisk))
 out.AWSElasticBlockStore = (*core.AWSElasticBlockStoreVolumeSource)(unsafe.Pointer(in.AWSElasticBlockStore))
 out.GitRepo = (*core.GitRepoVolumeSource)(unsafe.Pointer(in.GitRepo))
 out.Secret = (*core.SecretVolumeSource)(unsafe.Pointer(in.Secret))
 out.NFS = (*core.NFSVolumeSource)(unsafe.Pointer(in.NFS))
 out.ISCSI = (*core.ISCSIVolumeSource)(unsafe.Pointer(in.ISCSI))
 out.Glusterfs = (*core.GlusterfsVolumeSource)(unsafe.Pointer(in.Glusterfs))
 out.PersistentVolumeClaim = (*core.PersistentVolumeClaimVolumeSource)(unsafe.Pointer(in.PersistentVolumeClaim))
 out.RBD = (*core.RBDVolumeSource)(unsafe.Pointer(in.RBD))
 out.FlexVolume = (*core.FlexVolumeSource)(unsafe.Pointer(in.FlexVolume))
 out.Cinder = (*core.CinderVolumeSource)(unsafe.Pointer(in.Cinder))
 out.CephFS = (*core.CephFSVolumeSource)(unsafe.Pointer(in.CephFS))
 out.Flocker = (*core.FlockerVolumeSource)(unsafe.Pointer(in.Flocker))
 out.DownwardAPI = (*core.DownwardAPIVolumeSource)(unsafe.Pointer(in.DownwardAPI))
 out.FC = (*core.FCVolumeSource)(unsafe.Pointer(in.FC))
 out.AzureFile = (*core.AzureFileVolumeSource)(unsafe.Pointer(in.AzureFile))
 out.ConfigMap = (*core.ConfigMapVolumeSource)(unsafe.Pointer(in.ConfigMap))
 out.VsphereVolume = (*core.VsphereVirtualDiskVolumeSource)(unsafe.Pointer(in.VsphereVolume))
 out.Quobyte = (*core.QuobyteVolumeSource)(unsafe.Pointer(in.Quobyte))
 out.AzureDisk = (*core.AzureDiskVolumeSource)(unsafe.Pointer(in.AzureDisk))
 out.PhotonPersistentDisk = (*core.PhotonPersistentDiskVolumeSource)(unsafe.Pointer(in.PhotonPersistentDisk))
 if in.Projected != nil {
  in, out := &in.Projected, &out.Projected
  *out = new(core.ProjectedVolumeSource)
  if err := Convert_v1_ProjectedVolumeSource_To_core_ProjectedVolumeSource(*in, *out, s); err != nil {
   return err
  }
 } else {
  out.Projected = nil
 }
 out.PortworxVolume = (*core.PortworxVolumeSource)(unsafe.Pointer(in.PortworxVolume))
 out.ScaleIO = (*core.ScaleIOVolumeSource)(unsafe.Pointer(in.ScaleIO))
 out.StorageOS = (*core.StorageOSVolumeSource)(unsafe.Pointer(in.StorageOS))
 return nil
}
func Convert_v1_VolumeSource_To_core_VolumeSource(in *v1.VolumeSource, out *core.VolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_VolumeSource_To_core_VolumeSource(in, out, s)
}
func autoConvert_core_VolumeSource_To_v1_VolumeSource(in *core.VolumeSource, out *v1.VolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.HostPath = (*v1.HostPathVolumeSource)(unsafe.Pointer(in.HostPath))
 out.EmptyDir = (*v1.EmptyDirVolumeSource)(unsafe.Pointer(in.EmptyDir))
 out.GCEPersistentDisk = (*v1.GCEPersistentDiskVolumeSource)(unsafe.Pointer(in.GCEPersistentDisk))
 out.AWSElasticBlockStore = (*v1.AWSElasticBlockStoreVolumeSource)(unsafe.Pointer(in.AWSElasticBlockStore))
 out.GitRepo = (*v1.GitRepoVolumeSource)(unsafe.Pointer(in.GitRepo))
 out.Secret = (*v1.SecretVolumeSource)(unsafe.Pointer(in.Secret))
 out.NFS = (*v1.NFSVolumeSource)(unsafe.Pointer(in.NFS))
 out.ISCSI = (*v1.ISCSIVolumeSource)(unsafe.Pointer(in.ISCSI))
 out.Glusterfs = (*v1.GlusterfsVolumeSource)(unsafe.Pointer(in.Glusterfs))
 out.PersistentVolumeClaim = (*v1.PersistentVolumeClaimVolumeSource)(unsafe.Pointer(in.PersistentVolumeClaim))
 out.RBD = (*v1.RBDVolumeSource)(unsafe.Pointer(in.RBD))
 out.Quobyte = (*v1.QuobyteVolumeSource)(unsafe.Pointer(in.Quobyte))
 out.FlexVolume = (*v1.FlexVolumeSource)(unsafe.Pointer(in.FlexVolume))
 out.Cinder = (*v1.CinderVolumeSource)(unsafe.Pointer(in.Cinder))
 out.CephFS = (*v1.CephFSVolumeSource)(unsafe.Pointer(in.CephFS))
 out.Flocker = (*v1.FlockerVolumeSource)(unsafe.Pointer(in.Flocker))
 out.DownwardAPI = (*v1.DownwardAPIVolumeSource)(unsafe.Pointer(in.DownwardAPI))
 out.FC = (*v1.FCVolumeSource)(unsafe.Pointer(in.FC))
 out.AzureFile = (*v1.AzureFileVolumeSource)(unsafe.Pointer(in.AzureFile))
 out.ConfigMap = (*v1.ConfigMapVolumeSource)(unsafe.Pointer(in.ConfigMap))
 out.VsphereVolume = (*v1.VsphereVirtualDiskVolumeSource)(unsafe.Pointer(in.VsphereVolume))
 out.AzureDisk = (*v1.AzureDiskVolumeSource)(unsafe.Pointer(in.AzureDisk))
 out.PhotonPersistentDisk = (*v1.PhotonPersistentDiskVolumeSource)(unsafe.Pointer(in.PhotonPersistentDisk))
 if in.Projected != nil {
  in, out := &in.Projected, &out.Projected
  *out = new(v1.ProjectedVolumeSource)
  if err := Convert_core_ProjectedVolumeSource_To_v1_ProjectedVolumeSource(*in, *out, s); err != nil {
   return err
  }
 } else {
  out.Projected = nil
 }
 out.PortworxVolume = (*v1.PortworxVolumeSource)(unsafe.Pointer(in.PortworxVolume))
 out.ScaleIO = (*v1.ScaleIOVolumeSource)(unsafe.Pointer(in.ScaleIO))
 out.StorageOS = (*v1.StorageOSVolumeSource)(unsafe.Pointer(in.StorageOS))
 return nil
}
func Convert_core_VolumeSource_To_v1_VolumeSource(in *core.VolumeSource, out *v1.VolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_VolumeSource_To_v1_VolumeSource(in, out, s)
}
func autoConvert_v1_VsphereVirtualDiskVolumeSource_To_core_VsphereVirtualDiskVolumeSource(in *v1.VsphereVirtualDiskVolumeSource, out *core.VsphereVirtualDiskVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.VolumePath = in.VolumePath
 out.FSType = in.FSType
 out.StoragePolicyName = in.StoragePolicyName
 out.StoragePolicyID = in.StoragePolicyID
 return nil
}
func Convert_v1_VsphereVirtualDiskVolumeSource_To_core_VsphereVirtualDiskVolumeSource(in *v1.VsphereVirtualDiskVolumeSource, out *core.VsphereVirtualDiskVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_VsphereVirtualDiskVolumeSource_To_core_VsphereVirtualDiskVolumeSource(in, out, s)
}
func autoConvert_core_VsphereVirtualDiskVolumeSource_To_v1_VsphereVirtualDiskVolumeSource(in *core.VsphereVirtualDiskVolumeSource, out *v1.VsphereVirtualDiskVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.VolumePath = in.VolumePath
 out.FSType = in.FSType
 out.StoragePolicyName = in.StoragePolicyName
 out.StoragePolicyID = in.StoragePolicyID
 return nil
}
func Convert_core_VsphereVirtualDiskVolumeSource_To_v1_VsphereVirtualDiskVolumeSource(in *core.VsphereVirtualDiskVolumeSource, out *v1.VsphereVirtualDiskVolumeSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_VsphereVirtualDiskVolumeSource_To_v1_VsphereVirtualDiskVolumeSource(in, out, s)
}
func autoConvert_v1_WeightedPodAffinityTerm_To_core_WeightedPodAffinityTerm(in *v1.WeightedPodAffinityTerm, out *core.WeightedPodAffinityTerm, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Weight = in.Weight
 if err := Convert_v1_PodAffinityTerm_To_core_PodAffinityTerm(&in.PodAffinityTerm, &out.PodAffinityTerm, s); err != nil {
  return err
 }
 return nil
}
func Convert_v1_WeightedPodAffinityTerm_To_core_WeightedPodAffinityTerm(in *v1.WeightedPodAffinityTerm, out *core.WeightedPodAffinityTerm, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1_WeightedPodAffinityTerm_To_core_WeightedPodAffinityTerm(in, out, s)
}
func autoConvert_core_WeightedPodAffinityTerm_To_v1_WeightedPodAffinityTerm(in *core.WeightedPodAffinityTerm, out *v1.WeightedPodAffinityTerm, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Weight = in.Weight
 if err := Convert_core_PodAffinityTerm_To_v1_PodAffinityTerm(&in.PodAffinityTerm, &out.PodAffinityTerm, s); err != nil {
  return err
 }
 return nil
}
func Convert_core_WeightedPodAffinityTerm_To_v1_WeightedPodAffinityTerm(in *core.WeightedPodAffinityTerm, out *v1.WeightedPodAffinityTerm, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_core_WeightedPodAffinityTerm_To_v1_WeightedPodAffinityTerm(in, out, s)
}
