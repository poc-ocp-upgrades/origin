package main

import (
	goformat "fmt"
	"github.com/openshift/origin/pkg/cmd/flagtypes"
	"github.com/openshift/origin/pkg/cmd/openshift-sdn"
	"k8s.io/apiserver/pkg/util/logs"
	"math/rand"
	"os"
	goos "os"
	godefaultruntime "runtime"
	"time"
	gotime "time"
)

func main() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	logs.InitLogs()
	defer logs.FlushLogs()
	rand.Seed(time.Now().UTC().UnixNano())
	cmd := openshift_sdn.NewOpenShiftSDNCommand("openshift-sdn", os.Stderr)
	flagtypes.GLog(cmd.PersistentFlags())
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
