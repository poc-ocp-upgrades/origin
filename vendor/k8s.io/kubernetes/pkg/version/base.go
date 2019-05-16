package version

import (
	goformat "fmt"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

var (
	gitMajor     string
	gitMinor     string
	gitVersion   = "v0.0.0-master+$Format:%h$"
	gitCommit    = "$Format:%H$"
	gitTreeState = ""
	buildDate    = "1970-01-01T00:00:00Z"
)

func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
