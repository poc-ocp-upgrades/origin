package main

import (
	"flag"
	goformat "fmt"
	"github.com/openshift/library-go/pkg/serviceability"
	_ "github.com/openshift/origin/pkg/api/install"
	tsbcmd "github.com/openshift/origin/pkg/templateservicebroker/cmd/server"
	"github.com/openshift/origin/pkg/version"
	genericapiserver "k8s.io/apiserver/pkg/server"
	"k8s.io/apiserver/pkg/util/logs"
	"k8s.io/klog"
	_ "k8s.io/kubernetes/pkg/apis/autoscaling/install"
	_ "k8s.io/kubernetes/pkg/apis/batch/install"
	_ "k8s.io/kubernetes/pkg/apis/core/install"
	_ "k8s.io/kubernetes/pkg/apis/extensions/install"
	"math/rand"
	"os"
	goos "os"
	"runtime"
	godefaultruntime "runtime"
	"time"
	gotime "time"
)

func main() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	stopCh := genericapiserver.SetupSignalHandler()
	rand.Seed(time.Now().UTC().UnixNano())
	logs.InitLogs()
	defer logs.FlushLogs()
	if len(os.Getenv("GOMAXPROCS")) == 0 {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}
	defer serviceability.BehaviorOnPanic(os.Getenv("OPENSHIFT_ON_PANIC"), version.Get())()
	defer serviceability.Profile(os.Getenv("OPENSHIFT_PROFILE")).Stop()
	cmd := tsbcmd.NewCommandStartTemplateServiceBrokerServer(os.Stdout, os.Stderr, stopCh)
	cmd.Flags().AddGoFlagSet(flag.CommandLine)
	if err := cmd.Execute(); err != nil {
		klog.Fatal(err)
	}
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
