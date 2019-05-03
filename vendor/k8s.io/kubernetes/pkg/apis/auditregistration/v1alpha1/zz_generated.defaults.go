package v1alpha1

import (
 v1alpha1 "k8s.io/api/auditregistration/v1alpha1"
 runtime "k8s.io/apimachinery/pkg/runtime"
)

func RegisterDefaults(scheme *runtime.Scheme) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 scheme.AddTypeDefaultingFunc(&v1alpha1.AuditSink{}, func(obj interface{}) {
  SetObjectDefaults_AuditSink(obj.(*v1alpha1.AuditSink))
 })
 scheme.AddTypeDefaultingFunc(&v1alpha1.AuditSinkList{}, func(obj interface{}) {
  SetObjectDefaults_AuditSinkList(obj.(*v1alpha1.AuditSinkList))
 })
 return nil
}
func SetObjectDefaults_AuditSink(in *v1alpha1.AuditSink) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 SetDefaults_AuditSink(in)
}
func SetObjectDefaults_AuditSinkList(in *v1alpha1.AuditSinkList) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for i := range in.Items {
  a := &in.Items[i]
  SetObjectDefaults_AuditSink(a)
 }
}
