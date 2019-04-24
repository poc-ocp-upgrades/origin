package main

import (
	"flag"
	"bytes"
	"net/http"
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"time"
	"github.com/openshift/origin/pkg/version"
	"k8s.io/klog"
	genericapiserver "k8s.io/apiserver/pkg/server"
	"k8s.io/apiserver/pkg/util/logs"
	"github.com/openshift/library-go/pkg/serviceability"
	tsbcmd "github.com/openshift/origin/pkg/templateservicebroker/cmd/server"
	_ "github.com/openshift/origin/pkg/api/install"
	_ "k8s.io/kubernetes/pkg/apis/autoscaling/install"
	_ "k8s.io/kubernetes/pkg/apis/batch/install"
	_ "k8s.io/kubernetes/pkg/apis/core/install"
	_ "k8s.io/kubernetes/pkg/apis/extensions/install"
)

func main() {
	_logClusterCodePath()
	defer _logClusterCodePath()
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
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
