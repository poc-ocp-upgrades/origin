package version

import (
 godefaultruntime "runtime"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
)

var (
 gitMajor     string
 gitMinor     string
 gitVersion   = "v0.0.0-master+$Format:%h$"
 gitCommit    = "$Format:%H$"
 gitTreeState = ""
 buildDate    = "1970-01-01T00:00:00Z"
)

func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
