package v1alpha1

import (
 runtime "k8s.io/apimachinery/pkg/runtime"
 v1alpha1 "k8s.io/kube-controller-manager/config/v1alpha1"
)

func RegisterDefaults(scheme *runtime.Scheme) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 scheme.AddTypeDefaultingFunc(&v1alpha1.KubeControllerManagerConfiguration{}, func(obj interface{}) {
  SetObjectDefaults_KubeControllerManagerConfiguration(obj.(*v1alpha1.KubeControllerManagerConfiguration))
 })
 return nil
}
func SetObjectDefaults_KubeControllerManagerConfiguration(in *v1alpha1.KubeControllerManagerConfiguration) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 SetDefaults_KubeControllerManagerConfiguration(in)
 SetDefaults_KubeCloudSharedConfiguration(&in.KubeCloudShared)
 SetDefaults_CSRSigningControllerConfiguration(&in.CSRSigningController)
 SetDefaults_DaemonSetControllerConfiguration(&in.DaemonSetController)
 SetDefaults_DeploymentControllerConfiguration(&in.DeploymentController)
 SetDefaults_EndpointControllerConfiguration(&in.EndpointController)
 SetDefaults_GarbageCollectorControllerConfiguration(&in.GarbageCollectorController)
 SetDefaults_HPAControllerConfiguration(&in.HPAController)
 SetDefaults_JobControllerConfiguration(&in.JobController)
 SetDefaults_NamespaceControllerConfiguration(&in.NamespaceController)
 SetDefaults_NodeLifecycleControllerConfiguration(&in.NodeLifecycleController)
 SetDefaults_VolumeConfiguration(&in.PersistentVolumeBinderController.VolumeConfiguration)
 SetDefaults_PersistentVolumeRecyclerConfiguration(&in.PersistentVolumeBinderController.VolumeConfiguration.PersistentVolumeRecyclerConfiguration)
 SetDefaults_PodGCControllerConfiguration(&in.PodGCController)
 SetDefaults_ReplicaSetControllerConfiguration(&in.ReplicaSetController)
 SetDefaults_ReplicationControllerConfiguration(&in.ReplicationController)
 SetDefaults_ResourceQuotaControllerConfiguration(&in.ResourceQuotaController)
 SetDefaults_ServiceControllerConfiguration(&in.ServiceController)
}
