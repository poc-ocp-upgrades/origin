package v1alpha1

import (
 runtime "k8s.io/apimachinery/pkg/runtime"
 v1alpha1 "k8s.io/kube-scheduler/config/v1alpha1"
)

func RegisterDefaults(scheme *runtime.Scheme) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 scheme.AddTypeDefaultingFunc(&v1alpha1.KubeSchedulerConfiguration{}, func(obj interface{}) {
  SetObjectDefaults_KubeSchedulerConfiguration(obj.(*v1alpha1.KubeSchedulerConfiguration))
 })
 return nil
}
func SetObjectDefaults_KubeSchedulerConfiguration(in *v1alpha1.KubeSchedulerConfiguration) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 SetDefaults_KubeSchedulerConfiguration(in)
}
