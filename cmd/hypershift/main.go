package main

import (
	goflag "flag"
	"fmt"
	goformat "fmt"
	"github.com/openshift/library-go/pkg/serviceability"
	"github.com/openshift/origin/pkg/cmd/openshift-apiserver"
	"github.com/openshift/origin/pkg/cmd/openshift-controller-manager"
	"github.com/openshift/origin/pkg/cmd/openshift-integrated-oauth-server"
	"github.com/openshift/origin/pkg/cmd/openshift-kube-apiserver"
	"github.com/openshift/origin/pkg/cmd/openshift-network-controller"
	"github.com/openshift/origin/pkg/version"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	genericapiserver "k8s.io/apiserver/pkg/server"
	utilflag "k8s.io/apiserver/pkg/util/flag"
	"k8s.io/apiserver/pkg/util/logs"
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
	pflag.CommandLine.SetNormalizeFunc(utilflag.WordSepNormalizeFunc)
	pflag.CommandLine.AddGoFlagSet(goflag.CommandLine)
	logs.InitLogs()
	defer logs.FlushLogs()
	defer serviceability.BehaviorOnPanic(os.Getenv("OPENSHIFT_ON_PANIC"), version.Get())()
	defer serviceability.Profile(os.Getenv("OPENSHIFT_PROFILE")).Stop()
	if len(os.Getenv("GOMAXPROCS")) == 0 {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}
	command := NewHyperShiftCommand(stopCh)
	if err := command.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
func NewHyperShiftCommand(stopCh <-chan struct{}) *cobra.Command {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	cmd := &cobra.Command{Use: "hypershift", Short: "Combined server command for OpenShift", Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(1)
	}}
	startOpenShiftAPIServer := openshift_apiserver.NewOpenShiftAPIServerCommand(openshift_apiserver.RecommendedStartAPIServerName, "hypershift", os.Stdout, os.Stderr, stopCh)
	cmd.AddCommand(startOpenShiftAPIServer)
	startOpenShiftKubeAPIServer := openshift_kube_apiserver.NewOpenShiftKubeAPIServerServerCommand(openshift_kube_apiserver.RecommendedStartAPIServerName, "hypershift", os.Stdout, os.Stderr, stopCh)
	cmd.AddCommand(startOpenShiftKubeAPIServer)
	startOpenShiftControllerManager := openshift_controller_manager.NewOpenShiftControllerManagerCommand(openshift_controller_manager.RecommendedStartControllerManagerName, "hypershift", os.Stdout, os.Stderr)
	cmd.AddCommand(startOpenShiftControllerManager)
	startOpenShiftNetworkController := openshift_network_controller.NewOpenShiftNetworkControllerCommand(openshift_network_controller.RecommendedStartNetworkControllerName, "hypershift", os.Stdout, os.Stderr)
	cmd.AddCommand(startOpenShiftNetworkController)
	startOsin := openshift_integrated_oauth_server.NewOsinServer(os.Stdout, os.Stderr, stopCh)
	startOsin.Use = "openshift-osinserver"
	startOsin.Deprecated = "will be removed in 4.0"
	startOsin.Hidden = true
	cmd.AddCommand(startOsin)
	return cmd
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
