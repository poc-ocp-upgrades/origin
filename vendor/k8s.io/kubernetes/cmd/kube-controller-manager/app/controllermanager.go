package app

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"k8s.io/apimachinery/pkg/runtime/schema"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apimachinery/pkg/util/uuid"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/apiserver/pkg/server"
	"k8s.io/apiserver/pkg/server/healthz"
	"k8s.io/apiserver/pkg/server/mux"
	apiserverflag "k8s.io/apiserver/pkg/util/flag"
	"k8s.io/apiserver/pkg/util/globalflag"
	cacheddiscovery "k8s.io/client-go/discovery/cached"
	"k8s.io/client-go/informers"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/restmapper"
	"k8s.io/client-go/tools/leaderelection"
	"k8s.io/client-go/tools/leaderelection/resourcelock"
	certutil "k8s.io/client-go/util/cert"
	cloudprovider "k8s.io/cloud-provider"
	"k8s.io/klog"
	genericcontrollermanager "k8s.io/kubernetes/cmd/controller-manager/app"
	cmoptions "k8s.io/kubernetes/cmd/controller-manager/app/options"
	"k8s.io/kubernetes/cmd/kube-controller-manager/app/config"
	"k8s.io/kubernetes/cmd/kube-controller-manager/app/options"
	"k8s.io/kubernetes/pkg/controller"
	kubectrlmgrconfig "k8s.io/kubernetes/pkg/controller/apis/config"
	serviceaccountcontroller "k8s.io/kubernetes/pkg/controller/serviceaccount"
	"k8s.io/kubernetes/pkg/serviceaccount"
	"k8s.io/kubernetes/pkg/util/configz"
	"k8s.io/kubernetes/pkg/version"
	"k8s.io/kubernetes/pkg/version/verflag"
	"math/rand"
	"net/http"
	"os"
	"time"
)

const (
	ControllerStartJitter = 1.0
	ConfigzName           = "kubecontrollermanager.config.k8s.io"
)

type ControllerLoopMode int

const (
	IncludeCloudLoops ControllerLoopMode = iota
	ExternalLoops
)

func NewControllerManagerCommand(stopCh <-chan struct{}) *cobra.Command {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	s, err := options.NewKubeControllerManagerOptions()
	if err != nil {
		klog.Fatalf("unable to initialize command options: %v", err)
	}
	cmd := &cobra.Command{Use: "kube-controller-manager", Long: `The Kubernetes controller manager is a daemon that embeds
the core control loops shipped with Kubernetes. In applications of robotics and
automation, a control loop is a non-terminating loop that regulates the state of
the system. In Kubernetes, a controller is a control loop that watches the shared
state of the cluster through the apiserver and makes changes attempting to move the
current state towards the desired state. Examples of controllers that ship with
Kubernetes today are the replication controller, endpoints controller, namespace
controller, and serviceaccounts controller.`, Run: func(cmd *cobra.Command, args []string) {
		verflag.PrintAndExitIfRequested()
		if err := ShimFlagsForOpenShift(s); err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		c, err := s.Config(KnownControllers(), ControllersDisabledByDefault.List())
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		if err := ShimForOpenShift(s, c); err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		if err := Run(c.Complete(), stopCh); err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
	}}
	fs := cmd.Flags()
	namedFlagSets := s.Flags(KnownControllers(), ControllersDisabledByDefault.List())
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
func ResyncPeriod(c *config.CompletedConfig) func() time.Duration {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return func() time.Duration {
		factor := rand.Float64() + 1
		return time.Duration(float64(c.ComponentConfig.Generic.MinResyncPeriod.Nanoseconds()) * factor)
	}
}
func Run(c *config.CompletedConfig, stopCh <-chan struct{}) error {
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
	if cfgz, err := configz.New(ConfigzName); err == nil {
		cfgz.Set(c.ComponentConfig)
	} else {
		klog.Errorf("unable to register configz: %c", err)
	}
	var checks []healthz.HealthzChecker
	var electionChecker *leaderelection.HealthzAdaptor
	if c.ComponentConfig.Generic.LeaderElection.LeaderElect {
		electionChecker = leaderelection.NewLeaderHealthzAdaptor(time.Second * 20)
		checks = append(checks, electionChecker)
	}
	var unsecuredMux *mux.PathRecorderMux
	if c.SecureServing != nil {
		unsecuredMux = genericcontrollermanager.NewBaseHandler(&c.ComponentConfig.Generic.Debugging, checks...)
		handler := genericcontrollermanager.BuildHandlerChain(unsecuredMux, &c.Authorization, &c.Authentication)
		if serverStoppedCh, err := c.SecureServing.Serve(handler, 0, ctx.Done()); err != nil {
			return err
		} else {
			defer func() {
				cancel()
				<-serverStoppedCh
			}()
		}
	}
	if c.InsecureServing != nil {
		unsecuredMux = genericcontrollermanager.NewBaseHandler(&c.ComponentConfig.Generic.Debugging, checks...)
		insecureSuperuserAuthn := server.AuthenticationInfo{Authenticator: &server.InsecureSuperuser{}}
		handler := genericcontrollermanager.BuildHandlerChain(unsecuredMux, nil, &insecureSuperuserAuthn)
		if err := c.InsecureServing.Serve(handler, 0, ctx.Done()); err != nil {
			return err
		}
	}
	run := func(ctx context.Context) {
		rootClientBuilder := controller.SimpleControllerClientBuilder{ClientConfig: c.Kubeconfig}
		var clientBuilder controller.ControllerClientBuilder
		if c.ComponentConfig.KubeCloudShared.UseServiceAccountCredentials {
			if len(c.ComponentConfig.SAController.ServiceAccountKeyFile) == 0 {
				klog.Warningf("--use-service-account-credentials was specified without providing a --service-account-private-key-file")
			}
			clientBuilder = controller.SAControllerClientBuilder{ClientConfig: restclient.AnonymousClientConfig(c.Kubeconfig), CoreClient: c.Client.CoreV1(), AuthenticationClient: c.Client.AuthenticationV1(), Namespace: "kube-system"}
		} else {
			clientBuilder = rootClientBuilder
		}
		controllerContext, err := CreateControllerContext(c, rootClientBuilder, clientBuilder, ctx.Done())
		if err != nil {
			klog.Fatalf("error building controller context: %v", err)
		}
		saTokenControllerInitFunc := serviceAccountTokenControllerStarter{rootClientBuilder: rootClientBuilder}.startServiceAccountTokenController
		if err := createPVRecyclerSA(c.OpenShiftContext.OpenShiftConfig, rootClientBuilder); err != nil {
			klog.Fatalf("error creating recycler serviceaccount: %v", err)
		}
		if err := StartControllers(controllerContext, saTokenControllerInitFunc, NewControllerInitializers(controllerContext.LoopMode), unsecuredMux); err != nil {
			klog.Fatalf("error starting controllers: %v", err)
		}
		controllerContext.InformerFactory.Start(controllerContext.Stop)
		close(controllerContext.InformersStarted)
		select {}
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
	rl, err := resourcelock.New(c.ComponentConfig.Generic.LeaderElection.ResourceLock, "kube-system", "kube-controller-manager", c.LeaderElectionClient.CoreV1(), resourcelock.ResourceLockConfig{Identity: id, EventRecorder: c.EventRecorder})
	if err != nil {
		klog.Fatalf("error creating lock: %v", err)
	}
	leaderelection.RunOrDie(ctx, leaderelection.LeaderElectionConfig{Lock: rl, LeaseDuration: c.ComponentConfig.Generic.LeaderElection.LeaseDuration.Duration, RenewDeadline: c.ComponentConfig.Generic.LeaderElection.RenewDeadline.Duration, RetryPeriod: c.ComponentConfig.Generic.LeaderElection.RetryPeriod.Duration, Callbacks: leaderelection.LeaderCallbacks{OnStartedLeading: run, OnStoppedLeading: func() {
		cancel()
		utilruntime.HandleError(fmt.Errorf("leaderelection lost"))
	}}, WatchDog: electionChecker, Name: "kube-controller-manager"})
	return nil
}

type ControllerContext struct {
	OpenShiftContext   config.OpenShiftContext
	ClientBuilder      controller.ControllerClientBuilder
	InformerFactory    informers.SharedInformerFactory
	ComponentConfig    kubectrlmgrconfig.KubeControllerManagerConfiguration
	RESTMapper         *restmapper.DeferredDiscoveryRESTMapper
	AvailableResources map[schema.GroupVersionResource]bool
	Cloud              cloudprovider.Interface
	LoopMode           ControllerLoopMode
	Stop               <-chan struct{}
	InformersStarted   chan struct{}
	ResyncPeriod       func() time.Duration
}

func (c ControllerContext) IsControllerEnabled(name string) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return IsControllerEnabled(name, ControllersDisabledByDefault, c.ComponentConfig.Generic.Controllers...)
}
func IsControllerEnabled(name string, disabledByDefaultControllers sets.String, controllers ...string) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	hasStar := false
	for _, ctrl := range controllers {
		if ctrl == name {
			return true
		}
		if ctrl == "-"+name {
			return false
		}
		if ctrl == "*" {
			hasStar = true
		}
	}
	if !hasStar {
		return false
	}
	if disabledByDefaultControllers.Has(name) {
		return false
	}
	return true
}

type InitFunc func(ctx ControllerContext) (debuggingHandler http.Handler, enabled bool, err error)

func KnownControllers() []string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ret := sets.StringKeySet(NewControllerInitializers(IncludeCloudLoops))
	ret.Insert(saTokenControllerName)
	return ret.List()
}

var ControllersDisabledByDefault = sets.NewString("bootstrapsigner", "tokencleaner")

const (
	saTokenControllerName = "serviceaccount-token"
)

func NewControllerInitializers(loopMode ControllerLoopMode) map[string]InitFunc {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	controllers := map[string]InitFunc{}
	controllers["endpoint"] = startEndpointController
	controllers["replicationcontroller"] = startReplicationController
	controllers["podgc"] = startPodGCController
	controllers["resourcequota"] = startResourceQuotaController
	controllers["namespace"] = startNamespaceController
	controllers["serviceaccount"] = startServiceAccountController
	controllers["garbagecollector"] = startGarbageCollectorController
	controllers["daemonset"] = startDaemonSetController
	controllers["job"] = startJobController
	controllers["deployment"] = startDeploymentController
	controllers["replicaset"] = startReplicaSetController
	controllers["horizontalpodautoscaling"] = startHPAController
	controllers["disruption"] = startDisruptionController
	controllers["statefulset"] = startStatefulSetController
	controllers["cronjob"] = startCronJobController
	controllers["csrsigning"] = startCSRSigningController
	controllers["csrapproving"] = startCSRApprovingController
	controllers["csrcleaner"] = startCSRCleanerController
	controllers["ttl"] = startTTLController
	controllers["bootstrapsigner"] = startBootstrapSignerController
	controllers["tokencleaner"] = startTokenCleanerController
	controllers["nodeipam"] = startNodeIpamController
	if loopMode == IncludeCloudLoops {
		controllers["service"] = startServiceController
		controllers["route"] = startRouteController
	}
	controllers["nodelifecycle"] = startNodeLifecycleController
	controllers["persistentvolume-binder"] = startPersistentVolumeBinderController
	controllers["attachdetach"] = startAttachDetachController
	controllers["persistentvolume-expander"] = startVolumeExpandController
	controllers["clusterrole-aggregation"] = startClusterRoleAggregrationController
	controllers["pvc-protection"] = startPVCProtectionController
	controllers["pv-protection"] = startPVProtectionController
	controllers["ttl-after-finished"] = startTTLAfterFinishedController
	controllers["root-ca-cert-publisher"] = startRootCACertPublisher
	return controllers
}
func GetAvailableResources(clientBuilder controller.ControllerClientBuilder) (map[schema.GroupVersionResource]bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	client := clientBuilder.ClientOrDie("controller-discovery")
	discoveryClient := client.Discovery()
	resourceMap, err := discoveryClient.ServerResources()
	if err != nil {
		utilruntime.HandleError(fmt.Errorf("unable to get all supported resources from server: %v", err))
	}
	if len(resourceMap) == 0 {
		return nil, fmt.Errorf("unable to get any supported resources from server")
	}
	allResources := map[schema.GroupVersionResource]bool{}
	for _, apiResourceList := range resourceMap {
		version, err := schema.ParseGroupVersion(apiResourceList.GroupVersion)
		if err != nil {
			return nil, err
		}
		for _, apiResource := range apiResourceList.APIResources {
			allResources[version.WithResource(apiResource.Name)] = true
		}
	}
	return allResources, nil
}
func CreateControllerContext(s *config.CompletedConfig, rootClientBuilder, clientBuilder controller.ControllerClientBuilder, stop <-chan struct{}) (ControllerContext, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	versionedClient := rootClientBuilder.ClientOrDie("shared-informers")
	var sharedInformers informers.SharedInformerFactory
	if InformerFactoryOverride == nil {
		sharedInformers = informers.NewSharedInformerFactory(versionedClient, ResyncPeriod(s)())
	} else {
		sharedInformers = InformerFactoryOverride
	}
	if err := genericcontrollermanager.WaitForAPIServer(versionedClient.Discovery().RESTClient(), 10*time.Second); err != nil {
		return ControllerContext{}, fmt.Errorf("failed to wait for apiserver being healthy: %v", err)
	}
	discoveryClient := rootClientBuilder.ClientOrDie("controller-discovery")
	cachedClient := cacheddiscovery.NewMemCacheClient(discoveryClient.Discovery())
	restMapper := restmapper.NewDeferredDiscoveryRESTMapper(cachedClient)
	go wait.Until(func() {
		restMapper.Reset()
	}, 30*time.Second, stop)
	availableResources, err := GetAvailableResources(rootClientBuilder)
	if err != nil {
		return ControllerContext{}, err
	}
	cloud, loopMode, err := createCloudProvider(s.ComponentConfig.KubeCloudShared.CloudProvider.Name, s.ComponentConfig.KubeCloudShared.ExternalCloudVolumePlugin, s.ComponentConfig.KubeCloudShared.CloudProvider.CloudConfigFile, s.ComponentConfig.KubeCloudShared.AllowUntaggedCloud, sharedInformers)
	if err != nil {
		return ControllerContext{}, err
	}
	ctx := ControllerContext{OpenShiftContext: s.OpenShiftContext, ClientBuilder: clientBuilder, InformerFactory: sharedInformers, ComponentConfig: s.ComponentConfig, RESTMapper: restMapper, AvailableResources: availableResources, Cloud: cloud, LoopMode: loopMode, Stop: stop, InformersStarted: make(chan struct{}), ResyncPeriod: ResyncPeriod(s)}
	return ctx, nil
}
func StartControllers(ctx ControllerContext, startSATokenController InitFunc, controllers map[string]InitFunc, unsecuredMux *mux.PathRecorderMux) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if _, _, err := startSATokenController(ctx); err != nil {
		return err
	}
	if ctx.Cloud != nil {
		ctx.Cloud.Initialize(ctx.ClientBuilder, ctx.Stop)
	}
	for controllerName, initFn := range controllers {
		if !ctx.IsControllerEnabled(controllerName) {
			klog.Warningf("%q is disabled", controllerName)
			continue
		}
		time.Sleep(wait.Jitter(ctx.ComponentConfig.Generic.ControllerStartInterval.Duration, ControllerStartJitter))
		klog.V(1).Infof("Starting %q", controllerName)
		debugHandler, started, err := initFn(ctx)
		if err != nil {
			klog.Errorf("Error starting %q", controllerName)
			return err
		}
		if !started {
			klog.Warningf("Skipping %q", controllerName)
			continue
		}
		if debugHandler != nil && unsecuredMux != nil {
			basePath := "/debug/controllers/" + controllerName
			unsecuredMux.UnlistedHandle(basePath, http.StripPrefix(basePath, debugHandler))
			unsecuredMux.UnlistedHandlePrefix(basePath+"/", http.StripPrefix(basePath, debugHandler))
		}
		klog.Infof("Started %q", controllerName)
	}
	return nil
}

type serviceAccountTokenControllerStarter struct {
	rootClientBuilder controller.ControllerClientBuilder
}

func (c serviceAccountTokenControllerStarter) startServiceAccountTokenController(ctx ControllerContext) (http.Handler, bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if !ctx.IsControllerEnabled(saTokenControllerName) {
		klog.Warningf("%q is disabled", saTokenControllerName)
		return nil, false, nil
	}
	if len(ctx.ComponentConfig.SAController.ServiceAccountKeyFile) == 0 {
		klog.Warningf("%q is disabled because there is no private key", saTokenControllerName)
		return nil, false, nil
	}
	privateKey, err := certutil.PrivateKeyFromFile(ctx.ComponentConfig.SAController.ServiceAccountKeyFile)
	if err != nil {
		return nil, true, fmt.Errorf("error reading key for service account token controller: %v", err)
	}
	var rootCA []byte
	if ctx.ComponentConfig.SAController.RootCAFile != "" {
		if rootCA, err = readCA(ctx.ComponentConfig.SAController.RootCAFile); err != nil {
			return nil, true, fmt.Errorf("error parsing root-ca-file at %s: %v", ctx.ComponentConfig.SAController.RootCAFile, err)
		}
	} else {
		rootCA = c.rootClientBuilder.ConfigOrDie("tokens-controller").CAData
	}
	tokenGenerator, err := serviceaccount.JWTTokenGenerator(serviceaccount.LegacyIssuer, privateKey)
	if err != nil {
		return nil, false, fmt.Errorf("failed to build token generator: %v", err)
	}
	controller, err := serviceaccountcontroller.NewTokensController(ctx.InformerFactory.Core().V1().ServiceAccounts(), ctx.InformerFactory.Core().V1().Secrets(), c.rootClientBuilder.ClientOrDie("tokens-controller"), applyOpenShiftServiceServingCertCA(serviceaccountcontroller.TokensControllerOptions{TokenGenerator: tokenGenerator, RootCA: rootCA}))
	if err != nil {
		return nil, true, fmt.Errorf("error creating Tokens controller: %v", err)
	}
	go controller.Run(int(ctx.ComponentConfig.SAController.ConcurrentSATokenSyncs), ctx.Stop)
	ctx.InformerFactory.Start(ctx.Stop)
	return nil, true, nil
}
func readCA(file string) ([]byte, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	rootCA, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	if _, err := certutil.ParseCertsPEM(rootCA); err != nil {
		return nil, err
	}
	return rootCA, err
}
