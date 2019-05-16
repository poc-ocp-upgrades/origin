package crypto

import (
	"crypto/subtle"
	goformat "fmt"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

func IsEqualConstantTime(s1, s2 string) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return subtle.ConstantTimeCompare([]byte(s1), []byte(s2)) == 1
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
