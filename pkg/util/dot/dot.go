package dot

import (
	"fmt"
	goformat "fmt"
	goos "os"
	godefaultruntime "runtime"
	"strings"
	gotime "time"
)

func Quote(id string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return fmt.Sprintf(`"%s"`, strings.Replace(id, `"`, `\"`, -1))
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
