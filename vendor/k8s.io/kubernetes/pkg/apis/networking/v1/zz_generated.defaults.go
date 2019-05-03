package v1

import (
 v1 "k8s.io/api/networking/v1"
 runtime "k8s.io/apimachinery/pkg/runtime"
)

func RegisterDefaults(scheme *runtime.Scheme) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 scheme.AddTypeDefaultingFunc(&v1.NetworkPolicy{}, func(obj interface{}) {
  SetObjectDefaults_NetworkPolicy(obj.(*v1.NetworkPolicy))
 })
 scheme.AddTypeDefaultingFunc(&v1.NetworkPolicyList{}, func(obj interface{}) {
  SetObjectDefaults_NetworkPolicyList(obj.(*v1.NetworkPolicyList))
 })
 return nil
}
func SetObjectDefaults_NetworkPolicy(in *v1.NetworkPolicy) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 SetDefaults_NetworkPolicy(in)
 for i := range in.Spec.Ingress {
  a := &in.Spec.Ingress[i]
  for j := range a.Ports {
   b := &a.Ports[j]
   SetDefaults_NetworkPolicyPort(b)
  }
 }
 for i := range in.Spec.Egress {
  a := &in.Spec.Egress[i]
  for j := range a.Ports {
   b := &a.Ports[j]
   SetDefaults_NetworkPolicyPort(b)
  }
 }
}
func SetObjectDefaults_NetworkPolicyList(in *v1.NetworkPolicyList) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for i := range in.Items {
  a := &in.Items[i]
  SetObjectDefaults_NetworkPolicy(a)
 }
}
