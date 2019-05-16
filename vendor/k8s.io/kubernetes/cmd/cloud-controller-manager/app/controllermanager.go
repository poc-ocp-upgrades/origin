package app

import (
	"context"
	"fmt"
	goformat "fmt"
	"github.com/spf13/cobra"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/uuid"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/apiserver/pkg/server"
	"k8s.io/apiserver/pkg/server/healthz"
	apiserverflag "k8s.io/apiserver/pkg/util/flag"
	"k8s.io/apiserver/pkg/util/globalflag"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/leaderelection"
	"k8s.io/client-go/tools/leaderelection/resourcelock"
	cloudprovider "k8s.io/cloud-provider"
	"k8s.io/klog"
	cloudcontrollerconfig "k8s.io/kubernetes/cmd/cloud-controller-manager/app/config"
	"k8s.io/kubernetes/cmd/cloud-controller-manager/app/options"
	genericcontrollermanager "k8s.io/kubernetes/cmd/controller-manager/app"
	cmoptions "k8s.io/kubernetes/cmd/controller-manager/app/options"
	cloudcontrollers "k8s.io/kubernetes/pkg/controller/cloud"
	routecontroller "k8s.io/kubernetes/pkg/controller/route"
	servicecontroller "k8s.io/kubernetes/pkg/controller/service"
	"k8s.io/kubernetes/pkg/util/configz"
	utilflag "k8s.io/kubernetes/pkg/util/flag"
	"k8s.io/kubernetes/pkg/version"
	"k8s.io/kubernetes/pkg/version/verflag"
	"net"
	"os"
	goos "os"
	godefaultruntime "runtime"
	"strings"
	"time"
	gotime "time"
)

const (
	ControllerStartJitter = 1.0
	ConfigzName           = "cloudcontrollermanager.config.k8s.io"
)

func NewCloudControllerManagerCommand() *cobra.Command {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	s, err := options.NewCloudControllerManagerOptions()
	if err != nil {
		klog.Fatalf("unable to initialize command options: %v", err)
	}
	cmd := &cobra.Command{Use: "cloud-controller-manager", Long: `The Cloud controller manager is a daemon that embeds
the cloud specific control loops shipped with Kubernetes.`, Run: func(cmd *cobra.Command, args []string) {
		verflag.PrintAndExitIfRequested()
		utilflag.PrintFlags(cmd.Flags())
		c, err := s.Config()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		if err := Run(c.Complete(), wait.NeverStop); err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
	}}
	fs := cmd.Flags()
	namedFlagSets := s.Flags()
	verflag.AddFlags(namedFlagSets.FlagSet("global"))
	globalflag.AddGlobalFlags(namedFlagSets.FlagSet("global"), cmd.Name())
	cmoptions.AddCustomGlobalFlags(namedFlagSets.FlagSet("generic"))
	for _, f := range namedFlagSets.FlagSets {
		fs.AddFlagSet(f)
	}
	usageFmt := "Usage:\n  %s\n"
	cols, _, _ := apiserverflag.TerminalSize(cmd.OutOrStdout())
	cmd.SetUsageFunc(func(cmd *cobra.Command) error {
		fmt.Fprintf(cmd.OutOrStderr(), usageFmt, cmd.UseLine())
		apiserverflag.PrintSections(cmd.OutOrStderr(), namedFlagSets, cols)
		return nil
	})
	cmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		fmt.Fprintf(cmd.OutOrStdout(), "%s\n\n"+usageFmt, cmd.Long, cmd.UseLine())
		apiserverflag.PrintSections(cmd.OutOrStdout(), namedFlagSets, cols)
	})
	return cmd
}
func Run(c *cloudcontrollerconfig.CompletedConfig, stopCh <-chan struct{}) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()
	go func() {
		select {
		case <-stopCh:
			cancel()
		case <-ctx.Done():
		}
	}()
	klog.Infof("Version: %+v", version.Get())
	cloud, err := cloudprovider.InitCloudProvider(c.ComponentConfig.KubeCloudShared.CloudProvider.Name, c.ComponentConfig.KubeCloudShared.CloudProvider.CloudConfigFile)
	if err != nil {
		klog.Fatalf("Cloud provider could not be initialized: %v", err)
	}
	if cloud == nil {
		klog.Fatalf("cloud provider is nil")
	}
	if cloud.HasClusterID() == false {
		if c.ComponentConfig.KubeCloudShared.AllowUntaggedCloud == true {
			klog.Warning("detected a cluster without a ClusterID.  A ClusterID will be required in the future.  Please tag your cluster to avoid any future issues")
		} else {
			klog.Fatalf("no ClusterID found.  A ClusterID is required for the cloud provider to function properly.  This check can be bypassed by setting the allow-untagged-cloud option")
		}
	}
	if cz, err := configz.New(ConfigzName); err == nil {
		cz.Set(c.ComponentConfig)
	} else {
		klog.Errorf("unable to register configz: %c", err)
	}
	var checks []healthz.HealthzChecker
	var electionChecker *leaderelection.HealthzAdaptor
	if c.ComponentConfig.Generic.LeaderElection.LeaderElect {
		electionChecker = leaderelection.NewLeaderHealthzAdaptor(time.Second * 20)
		checks = append(checks, electionChecker)
	}
	if c.SecureServing != nil {
		unsecuredMux := genericcontrollermanager.NewBaseHandler(&c.ComponentConfig.Generic.Debugging, checks...)
		handler := genericcontrollermanager.BuildHandlerChain(unsecuredMux, &c.Authorization, &c.Authentication)
		if serverStoppedCh, err := c.SecureServing.Serve(handler, 0, stopCh); err != nil {
			return err
		} else {
			defer func() {
				cancel()
				<-serverStoppedCh
			}()
		}
	}
	if c.InsecureServing != nil {
		unsecuredMux := genericcontrollermanager.NewBaseHandler(&c.ComponentConfig.Generic.Debugging, checks...)
		insecureSuperuserAuthn := server.AuthenticationInfo{Authenticator: &server.InsecureSuperuser{}}
		handler := genericcontrollermanager.BuildHandlerChain(unsecuredMux, nil, &insecureSuperuserAuthn)
		if err := c.InsecureServing.Serve(handler, 0, stopCh); err != nil {
			return err
		}
	}
	run := func(ctx context.Context) {
		if err := startControllers(c, ctx.Done(), cloud); err != nil {
			klog.Fatalf("error running controllers: %v", err)
		}
	}
	if !c.ComponentConfig.Generic.LeaderElection.LeaderElect {
		run(context.TODO())
		panic("unreachable")
	}
	id, err := os.Hostname()
	if err != nil {
		return err
	}
	id = id + "_" + string(uuid.NewUUID())
	rl, err := resourcelock.New(c.ComponentConfig.Generic.LeaderElection.ResourceLock, "kube-system", "cloud-controller-manager", c.LeaderElectionClient.CoreV1(), resourcelock.ResourceLockConfig{Identity: id, EventRecorder: c.EventRecorder})
	if err != nil {
		klog.Fatalf("error creating lock: %v", err)
	}
	leaderelection.RunOrDie(context.TODO(), leaderelection.LeaderElectionConfig{Lock: rl, LeaseDuration: c.ComponentConfig.Generic.LeaderElection.LeaseDuration.Duration, RenewDeadline: c.ComponentConfig.Generic.LeaderElection.RenewDeadline.Duration, RetryPeriod: c.ComponentConfig.Generic.LeaderElection.RetryPeriod.Duration, Callbacks: leaderelection.LeaderCallbacks{OnStartedLeading: run, OnStoppedLeading: func() {
		cancel()
		utilruntime.HandleError(fmt.Errorf("leaderelection lost"))
	}}, WatchDog: electionChecker, Name: "cloud-controller-manager"})
	return nil
}
func startControllers(c *cloudcontrollerconfig.CompletedConfig, stop <-chan struct{}, cloud cloudprovider.Interface) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	client := func(serviceAccountName string) kubernetes.Interface {
		return c.ClientBuilder.ClientOrDie(serviceAccountName)
	}
	if cloud != nil {
		cloud.Initialize(c.ClientBuilder, stop)
	}
	nodeController := cloudcontrollers.NewCloudNodeController(c.SharedInformers.Core().V1().Nodes(), client("cloud-node-controller"), cloud, c.ComponentConfig.KubeCloudShared.NodeMonitorPeriod.Duration, c.ComponentConfig.NodeStatusUpdateFrequency.Duration)
	nodeController.Run(stop)
	time.Sleep(wait.Jitter(c.ComponentConfig.Generic.ControllerStartInterval.Duration, ControllerStartJitter))
	pvlController := cloudcontrollers.NewPersistentVolumeLabelController(client("pvl-controller"), cloud)
	go pvlController.Run(5, stop)
	time.Sleep(wait.Jitter(c.ComponentConfig.Generic.ControllerStartInterval.Duration, ControllerStartJitter))
	serviceController, err := servicecontroller.New(cloud, client("service-controller"), c.SharedInformers.Core().V1().Services(), c.SharedInformers.Core().V1().Nodes(), c.ComponentConfig.KubeCloudShared.ClusterName)
	if err != nil {
		klog.Errorf("Failed to start service controller: %v", err)
	} else {
		go serviceController.Run(stop, int(c.ComponentConfig.ServiceController.ConcurrentServiceSyncs))
		time.Sleep(wait.Jitter(c.ComponentConfig.Generic.ControllerStartInterval.Duration, ControllerStartJitter))
	}
	if c.ComponentConfig.KubeCloudShared.AllocateNodeCIDRs && c.ComponentConfig.KubeCloudShared.ConfigureCloudRoutes {
		if routes, ok := cloud.Routes(); !ok {
			klog.Warning("configure-cloud-routes is set, but cloud provider does not support routes. Will not configure cloud provider routes.")
		} else {
			var clusterCIDR *net.IPNet
			if len(strings.TrimSpace(c.ComponentConfig.KubeCloudShared.ClusterCIDR)) != 0 {
				_, clusterCIDR, err = net.ParseCIDR(c.ComponentConfig.KubeCloudShared.ClusterCIDR)
				if err != nil {
					klog.Warningf("Unsuccessful parsing of cluster CIDR %v: %v", c.ComponentConfig.KubeCloudShared.ClusterCIDR, err)
				}
			}
			routeController := routecontroller.New(routes, client("route-controller"), c.SharedInformers.Core().V1().Nodes(), c.ComponentConfig.KubeCloudShared.ClusterName, clusterCIDR)
			go routeController.Run(stop, c.ComponentConfig.KubeCloudShared.RouteReconciliationPeriod.Duration)
			time.Sleep(wait.Jitter(c.ComponentConfig.Generic.ControllerStartInterval.Duration, ControllerStartJitter))
		}
	} else {
		klog.Infof("Will not configure cloud provider routes for allocate-node-cidrs: %v, configure-cloud-routes: %v.", c.ComponentConfig.KubeCloudShared.AllocateNodeCIDRs, c.ComponentConfig.KubeCloudShared.ConfigureCloudRoutes)
	}
	err = genericcontrollermanager.WaitForAPIServer(c.VersionedClient.Discovery().RESTClient(), 10*time.Second)
	if err != nil {
		klog.Fatalf("Failed to wait for apiserver being healthy: %v", err)
	}
	c.SharedInformers.Start(stop)
	select {}
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
