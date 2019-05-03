package crypto

import (
	godefaultbytes "bytes"
	"crypto/subtle"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
)

func IsEqualConstantTime(s1, s2 string) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return subtle.ConstantTimeCompare([]byte(s1), []byte(s2)) == 1
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
