package v1alpha1

import (
 rbacv1alpha1 "k8s.io/api/rbac/v1alpha1"
 "k8s.io/apimachinery/pkg/runtime"
)

func addDefaultingFuncs(scheme *runtime.Scheme) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return RegisterDefaults(scheme)
}
func SetDefaults_ClusterRoleBinding(obj *rbacv1alpha1.ClusterRoleBinding) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if len(obj.RoleRef.APIGroup) == 0 {
  obj.RoleRef.APIGroup = GroupName
 }
}
func SetDefaults_RoleBinding(obj *rbacv1alpha1.RoleBinding) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if len(obj.RoleRef.APIGroup) == 0 {
  obj.RoleRef.APIGroup = GroupName
 }
}
func SetDefaults_Subject(obj *rbacv1alpha1.Subject) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if len(obj.APIVersion) == 0 {
  switch obj.Kind {
  case rbacv1alpha1.ServiceAccountKind:
   obj.APIVersion = "v1"
  case rbacv1alpha1.UserKind:
   obj.APIVersion = SchemeGroupVersion.String()
  case rbacv1alpha1.GroupKind:
   obj.APIVersion = SchemeGroupVersion.String()
  }
 }
}
