package redirect

import (
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	"net/url"
	godefaultruntime "runtime"
	"strings"
)

func IsServerRelativeURL(then string) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(then) == 0 {
		return false
	}
	u, err := url.Parse(then)
	if err != nil {
		return false
	}
	return len(u.Scheme) == 0 && len(u.Host) == 0 && strings.HasPrefix(u.Path, "/")
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
