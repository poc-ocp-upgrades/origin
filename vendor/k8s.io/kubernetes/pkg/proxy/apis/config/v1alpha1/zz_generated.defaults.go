package v1alpha1

import (
 runtime "k8s.io/apimachinery/pkg/runtime"
 v1alpha1 "k8s.io/kube-proxy/config/v1alpha1"
)

func RegisterDefaults(scheme *runtime.Scheme) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 scheme.AddTypeDefaultingFunc(&v1alpha1.KubeProxyConfiguration{}, func(obj interface{}) {
  SetObjectDefaults_KubeProxyConfiguration(obj.(*v1alpha1.KubeProxyConfiguration))
 })
 return nil
}
func SetObjectDefaults_KubeProxyConfiguration(in *v1alpha1.KubeProxyConfiguration) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 SetDefaults_KubeProxyConfiguration(in)
}
