package tokens

import (
	"fmt"
	"bytes"
	"net/http"
	"runtime"
)

const (
	URLToken		= "${url}"
	ServerRelativeURLToken	= "${server-relative-url}"
	QueryToken		= "${query}"
)

func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
