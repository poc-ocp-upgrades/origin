package redirect

import (
	goformat "fmt"
	"net/url"
	goos "os"
	godefaultruntime "runtime"
	"strings"
	gotime "time"
)

func IsServerRelativeURL(then string) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(then) == 0 {
		return false
	}
	u, err := url.Parse(then)
	if err != nil {
		return false
	}
	return len(u.Scheme) == 0 && len(u.Host) == 0 && strings.HasPrefix(u.Path, "/")
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
