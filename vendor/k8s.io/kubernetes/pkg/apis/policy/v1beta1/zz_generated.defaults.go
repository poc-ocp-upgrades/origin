package v1beta1

import (
 v1beta1 "k8s.io/api/policy/v1beta1"
 runtime "k8s.io/apimachinery/pkg/runtime"
)

func RegisterDefaults(scheme *runtime.Scheme) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 scheme.AddTypeDefaultingFunc(&v1beta1.PodSecurityPolicy{}, func(obj interface{}) {
  SetObjectDefaults_PodSecurityPolicy(obj.(*v1beta1.PodSecurityPolicy))
 })
 scheme.AddTypeDefaultingFunc(&v1beta1.PodSecurityPolicyList{}, func(obj interface{}) {
  SetObjectDefaults_PodSecurityPolicyList(obj.(*v1beta1.PodSecurityPolicyList))
 })
 return nil
}
func SetObjectDefaults_PodSecurityPolicy(in *v1beta1.PodSecurityPolicy) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 SetDefaults_PodSecurityPolicySpec(&in.Spec)
}
func SetObjectDefaults_PodSecurityPolicyList(in *v1beta1.PodSecurityPolicyList) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for i := range in.Items {
  a := &in.Items[i]
  SetObjectDefaults_PodSecurityPolicy(a)
 }
}
