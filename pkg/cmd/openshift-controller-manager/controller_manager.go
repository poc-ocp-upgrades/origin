package openshift_controller_manager

import (
	"context"
	"fmt"
	openshiftcontrolplanev1 "github.com/openshift/api/openshiftcontrolplane/v1"
	origincontrollers "github.com/openshift/origin/pkg/cmd/openshift-controller-manager/controller"
	"github.com/openshift/origin/pkg/cmd/util"
	"github.com/openshift/origin/pkg/cmd/util/variable"
	"github.com/openshift/origin/pkg/version"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	v1core "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/leaderelection"
	"k8s.io/client-go/tools/leaderelection/resourcelock"
	"k8s.io/client-go/tools/record"
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	_ "k8s.io/kubernetes/pkg/client/metrics/prometheus"
	"net/http"
	"os"
	"time"
)

func RunOpenShiftControllerManager(config *openshiftcontrolplanev1.OpenShiftControllerManagerConfig, clientConfig *rest.Config) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	util.InitLogrus()
	kubeClient, err := kubernetes.NewForConfig(clientConfig)
	if err != nil {
		return err
	}
	if config.ServingInfo != nil {
		klog.Infof("Starting controllers on %s (%s)", config.ServingInfo.BindAddress, version.Get().String())
		if err := origincontrollers.RunControllerServer(*config.ServingInfo, kubeClient); err != nil {
			return err
		}
	}
	{
		imageTemplate := variable.NewDefaultImageTemplate()
		imageTemplate.Format = config.Deployer.ImageTemplateFormat.Format
		imageTemplate.Latest = config.Deployer.ImageTemplateFormat.Latest
		klog.Infof("DeploymentConfig controller using images from %q", imageTemplate.ExpandOrDie("<component>"))
	}
	{
		imageTemplate := variable.NewDefaultImageTemplate()
		imageTemplate.Format = config.Build.ImageTemplateFormat.Format
		imageTemplate.Latest = config.Build.ImageTemplateFormat.Latest
		klog.Infof("Build controller using images from %q", imageTemplate.ExpandOrDie("<component>"))
	}
	originControllerManager := func(ctx context.Context) {
		if err := WaitForHealthyAPIServer(kubeClient.Discovery().RESTClient()); err != nil {
			klog.Fatal(err)
		}
		controllerContext, err := origincontrollers.NewControllerContext(*config, clientConfig, ctx.Done())
		if err != nil {
			klog.Fatal(err)
		}
		if err := startControllers(controllerContext); err != nil {
			klog.Fatal(err)
		}
		controllerContext.StartInformers(ctx.Done())
	}
	eventBroadcaster := record.NewBroadcaster()
	eventBroadcaster.StartLogging(klog.Infof)
	eventBroadcaster.StartRecordingToSink(&v1core.EventSinkImpl{Interface: kubeClient.CoreV1().Events("")})
	eventRecorder := eventBroadcaster.NewRecorder(legacyscheme.Scheme, v1.EventSource{Component: "openshift-controller-manager"})
	id, err := os.Hostname()
	if err != nil {
		return err
	}
	rl, err := resourcelock.New("configmaps", "openshift-controller-manager", "openshift-master-controllers", kubeClient.CoreV1(), resourcelock.ResourceLockConfig{Identity: id, EventRecorder: eventRecorder})
	if err != nil {
		return err
	}
	go leaderelection.RunOrDie(context.Background(), leaderelection.LeaderElectionConfig{Lock: rl, LeaseDuration: config.LeaderElection.LeaseDuration.Duration, RenewDeadline: config.LeaderElection.RenewDeadline.Duration, RetryPeriod: config.LeaderElection.RetryPeriod.Duration, Callbacks: leaderelection.LeaderCallbacks{OnStartedLeading: originControllerManager, OnStoppedLeading: func() {
		klog.Fatalf("leaderelection lost")
	}}})
	return nil
}
func WaitForHealthyAPIServer(client rest.Interface) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var healthzContent string
	err := wait.PollImmediate(time.Second, 5*time.Minute, func() (bool, error) {
		healthStatus := 0
		resp := client.Get().AbsPath("/healthz").Do().StatusCode(&healthStatus)
		if healthStatus != http.StatusOK {
			klog.Errorf("Server isn't healthy yet. Waiting a little while.")
			return false, nil
		}
		content, _ := resp.Raw()
		healthzContent = string(content)
		return true, nil
	})
	if err != nil {
		return fmt.Errorf("server unhealthy: %v: %v", healthzContent, err)
	}
	return nil
}
func startControllers(controllerContext *origincontrollers.ControllerContext) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for controllerName, initFn := range origincontrollers.ControllerInitializers {
		if !controllerContext.IsControllerEnabled(controllerName) {
			klog.Warningf("%q is disabled", controllerName)
			continue
		}
		klog.V(1).Infof("Starting %q", controllerName)
		started, err := initFn(controllerContext)
		if err != nil {
			klog.Fatalf("Error starting %q (%v)", controllerName, err)
			return err
		}
		if !started {
			klog.Warningf("Skipping %q", controllerName)
			continue
		}
		klog.Infof("Started %q", controllerName)
	}
	klog.Infof("Started Origin Controllers")
	return nil
}
