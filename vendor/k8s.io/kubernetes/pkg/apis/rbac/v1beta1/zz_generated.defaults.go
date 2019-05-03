package v1beta1

import (
 v1beta1 "k8s.io/api/rbac/v1beta1"
 runtime "k8s.io/apimachinery/pkg/runtime"
)

func RegisterDefaults(scheme *runtime.Scheme) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 scheme.AddTypeDefaultingFunc(&v1beta1.ClusterRoleBinding{}, func(obj interface{}) {
  SetObjectDefaults_ClusterRoleBinding(obj.(*v1beta1.ClusterRoleBinding))
 })
 scheme.AddTypeDefaultingFunc(&v1beta1.ClusterRoleBindingList{}, func(obj interface{}) {
  SetObjectDefaults_ClusterRoleBindingList(obj.(*v1beta1.ClusterRoleBindingList))
 })
 scheme.AddTypeDefaultingFunc(&v1beta1.RoleBinding{}, func(obj interface{}) {
  SetObjectDefaults_RoleBinding(obj.(*v1beta1.RoleBinding))
 })
 scheme.AddTypeDefaultingFunc(&v1beta1.RoleBindingList{}, func(obj interface{}) {
  SetObjectDefaults_RoleBindingList(obj.(*v1beta1.RoleBindingList))
 })
 return nil
}
func SetObjectDefaults_ClusterRoleBinding(in *v1beta1.ClusterRoleBinding) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 SetDefaults_ClusterRoleBinding(in)
 for i := range in.Subjects {
  a := &in.Subjects[i]
  SetDefaults_Subject(a)
 }
}
func SetObjectDefaults_ClusterRoleBindingList(in *v1beta1.ClusterRoleBindingList) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for i := range in.Items {
  a := &in.Items[i]
  SetObjectDefaults_ClusterRoleBinding(a)
 }
}
func SetObjectDefaults_RoleBinding(in *v1beta1.RoleBinding) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 SetDefaults_RoleBinding(in)
 for i := range in.Subjects {
  a := &in.Subjects[i]
  SetDefaults_Subject(a)
 }
}
func SetObjectDefaults_RoleBindingList(in *v1beta1.RoleBindingList) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for i := range in.Items {
  a := &in.Items[i]
  SetObjectDefaults_RoleBinding(a)
 }
}
