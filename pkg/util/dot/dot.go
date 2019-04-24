package dot

import (
	"fmt"
	"bytes"
	"net/http"
	"runtime"
	"strings"
)

func Quote(id string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fmt.Sprintf(`"%s"`, strings.Replace(id, `"`, `\"`, -1))
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
