package redirect

import (
	"net/url"
	"bytes"
	"net/http"
	"runtime"
	"fmt"
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
