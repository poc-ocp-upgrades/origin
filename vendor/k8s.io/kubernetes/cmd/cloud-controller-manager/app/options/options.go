package options

import (
	"fmt"
	goformat "fmt"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	apiserveroptions "k8s.io/apiserver/pkg/server/options"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	apiserverflag "k8s.io/apiserver/pkg/util/flag"
	"k8s.io/client-go/informers"
	clientset "k8s.io/client-go/kubernetes"
	v1core "k8s.io/client-go/kubernetes/typed/core/v1"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/record"
	"k8s.io/klog"
	ccmconfig "k8s.io/kubernetes/cmd/cloud-controller-manager/app/apis/config"
	ccmconfigscheme "k8s.io/kubernetes/cmd/cloud-controller-manager/app/apis/config/scheme"
	ccmconfigv1alpha1 "k8s.io/kubernetes/cmd/cloud-controller-manager/app/apis/config/v1alpha1"
	cloudcontrollerconfig "k8s.io/kubernetes/cmd/cloud-controller-manager/app/config"
	cmoptions "k8s.io/kubernetes/cmd/controller-manager/app/options"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	"k8s.io/kubernetes/pkg/controller"
	_ "k8s.io/kubernetes/pkg/features"
	"k8s.io/kubernetes/pkg/master/ports"
	"math/rand"
	"net"
	goos "os"
	godefaultruntime "runtime"
	"time"
	gotime "time"
)

const (
	CloudControllerManagerUserAgent           = "cloud-controller-manager"
	DefaultInsecureCloudControllerManagerPort = 0
)

type CloudControllerManagerOptions struct {
	Generic                   *cmoptions.GenericControllerManagerConfigurationOptions
	KubeCloudShared           *cmoptions.KubeCloudSharedOptions
	ServiceController         *cmoptions.ServiceControllerOptions
	SecureServing             *apiserveroptions.SecureServingOptionsWithLoopback
	InsecureServing           *apiserveroptions.DeprecatedInsecureServingOptionsWithLoopback
	Authentication            *apiserveroptions.DelegatingAuthenticationOptions
	Authorization             *apiserveroptions.DelegatingAuthorizationOptions
	Master                    string
	Kubeconfig                string
	NodeStatusUpdateFrequency metav1.Duration
}

func NewCloudControllerManagerOptions() (*CloudControllerManagerOptions, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	componentConfig, err := NewDefaultComponentConfig(DefaultInsecureCloudControllerManagerPort)
	if err != nil {
		return nil, err
	}
	s := CloudControllerManagerOptions{Generic: cmoptions.NewGenericControllerManagerConfigurationOptions(componentConfig.Generic), KubeCloudShared: cmoptions.NewKubeCloudSharedOptions(componentConfig.KubeCloudShared), ServiceController: &cmoptions.ServiceControllerOptions{ConcurrentServiceSyncs: componentConfig.ServiceController.ConcurrentServiceSyncs}, SecureServing: apiserveroptions.NewSecureServingOptions().WithLoopback(), InsecureServing: (&apiserveroptions.DeprecatedInsecureServingOptions{BindAddress: net.ParseIP(componentConfig.Generic.Address), BindPort: int(componentConfig.Generic.Port), BindNetwork: "tcp"}).WithLoopback(), Authentication: apiserveroptions.NewDelegatingAuthenticationOptions(), Authorization: apiserveroptions.NewDelegatingAuthorizationOptions(), NodeStatusUpdateFrequency: componentConfig.NodeStatusUpdateFrequency}
	s.Authentication.RemoteKubeConfigFileOptional = true
	s.Authorization.RemoteKubeConfigFileOptional = true
	s.Authorization.AlwaysAllowPaths = []string{"/healthz"}
	s.SecureServing.ServerCert.CertDirectory = ""
	s.SecureServing.ServerCert.PairName = "cloud-controller-manager"
	s.SecureServing.BindPort = ports.CloudControllerManagerPort
	return &s, nil
}
func NewDefaultComponentConfig(insecurePort int32) (*ccmconfig.CloudControllerManagerConfiguration, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	versioned := &ccmconfigv1alpha1.CloudControllerManagerConfiguration{}
	ccmconfigscheme.Scheme.Default(versioned)
	internal := &ccmconfig.CloudControllerManagerConfiguration{}
	if err := ccmconfigscheme.Scheme.Convert(versioned, internal, nil); err != nil {
		return nil, err
	}
	internal.Generic.Port = insecurePort
	return internal, nil
}
func (o *CloudControllerManagerOptions) Flags() apiserverflag.NamedFlagSets {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	fss := apiserverflag.NamedFlagSets{}
	o.Generic.AddFlags(&fss, []string{}, []string{})
	fss.FlagSet("generic").MarkHidden("controllers")
	o.KubeCloudShared.AddFlags(fss.FlagSet("generic"))
	o.ServiceController.AddFlags(fss.FlagSet("service controller"))
	o.SecureServing.AddFlags(fss.FlagSet("secure serving"))
	o.InsecureServing.AddUnqualifiedFlags(fss.FlagSet("insecure serving"))
	o.Authentication.AddFlags(fss.FlagSet("authentication"))
	o.Authorization.AddFlags(fss.FlagSet("authorization"))
	fs := fss.FlagSet("misc")
	fs.StringVar(&o.Master, "master", o.Master, "The address of the Kubernetes API server (overrides any value in kubeconfig).")
	fs.StringVar(&o.Kubeconfig, "kubeconfig", o.Kubeconfig, "Path to kubeconfig file with authorization and master location information.")
	fs.DurationVar(&o.NodeStatusUpdateFrequency.Duration, "node-status-update-frequency", o.NodeStatusUpdateFrequency.Duration, "Specifies how often the controller updates nodes' status.")
	utilfeature.DefaultFeatureGate.AddFlag(fss.FlagSet("generic"))
	return fss
}
func (o *CloudControllerManagerOptions) ApplyTo(c *cloudcontrollerconfig.Config, userAgent string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var err error
	if err = o.Generic.ApplyTo(&c.ComponentConfig.Generic); err != nil {
		return err
	}
	if err = o.KubeCloudShared.ApplyTo(&c.ComponentConfig.KubeCloudShared); err != nil {
		return err
	}
	if err = o.ServiceController.ApplyTo(&c.ComponentConfig.ServiceController); err != nil {
		return err
	}
	if err = o.InsecureServing.ApplyTo(&c.InsecureServing, &c.LoopbackClientConfig); err != nil {
		return err
	}
	if err = o.SecureServing.ApplyTo(&c.SecureServing, &c.LoopbackClientConfig); err != nil {
		return err
	}
	if o.SecureServing.BindPort != 0 || o.SecureServing.Listener != nil {
		if err = o.Authentication.ApplyTo(&c.Authentication, c.SecureServing, nil); err != nil {
			return err
		}
		if err = o.Authorization.ApplyTo(&c.Authorization); err != nil {
			return err
		}
	}
	c.Kubeconfig, err = clientcmd.BuildConfigFromFlags(o.Master, o.Kubeconfig)
	if err != nil {
		return err
	}
	c.Kubeconfig.ContentConfig.ContentType = o.Generic.ClientConnection.ContentType
	c.Kubeconfig.QPS = o.Generic.ClientConnection.QPS
	c.Kubeconfig.Burst = int(o.Generic.ClientConnection.Burst)
	c.Client, err = clientset.NewForConfig(restclient.AddUserAgent(c.Kubeconfig, userAgent))
	if err != nil {
		return err
	}
	c.LeaderElectionClient = clientset.NewForConfigOrDie(restclient.AddUserAgent(c.Kubeconfig, "leader-election"))
	c.EventRecorder = createRecorder(c.Client, userAgent)
	rootClientBuilder := controller.SimpleControllerClientBuilder{ClientConfig: c.Kubeconfig}
	if c.ComponentConfig.KubeCloudShared.UseServiceAccountCredentials {
		c.ClientBuilder = controller.SAControllerClientBuilder{ClientConfig: restclient.AnonymousClientConfig(c.Kubeconfig), CoreClient: c.Client.CoreV1(), AuthenticationClient: c.Client.AuthenticationV1(), Namespace: metav1.NamespaceSystem}
	} else {
		c.ClientBuilder = rootClientBuilder
	}
	c.VersionedClient = rootClientBuilder.ClientOrDie("shared-informers")
	c.SharedInformers = informers.NewSharedInformerFactory(c.VersionedClient, resyncPeriod(c)())
	c.ComponentConfig.Generic.Port = int32(o.InsecureServing.BindPort)
	c.ComponentConfig.Generic.Address = o.InsecureServing.BindAddress.String()
	c.ComponentConfig.NodeStatusUpdateFrequency = o.NodeStatusUpdateFrequency
	return nil
}
func (o *CloudControllerManagerOptions) Validate() error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	errors := []error{}
	errors = append(errors, o.Generic.Validate(nil, nil)...)
	errors = append(errors, o.KubeCloudShared.Validate()...)
	errors = append(errors, o.ServiceController.Validate()...)
	errors = append(errors, o.SecureServing.Validate()...)
	errors = append(errors, o.InsecureServing.Validate()...)
	errors = append(errors, o.Authentication.Validate()...)
	errors = append(errors, o.Authorization.Validate()...)
	if len(o.KubeCloudShared.CloudProvider.Name) == 0 {
		errors = append(errors, fmt.Errorf("--cloud-provider cannot be empty"))
	}
	return utilerrors.NewAggregate(errors)
}
func resyncPeriod(c *cloudcontrollerconfig.Config) func() time.Duration {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return func() time.Duration {
		factor := rand.Float64() + 1
		return time.Duration(float64(c.ComponentConfig.Generic.MinResyncPeriod.Nanoseconds()) * factor)
	}
}
func (o *CloudControllerManagerOptions) Config() (*cloudcontrollerconfig.Config, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if err := o.Validate(); err != nil {
		return nil, err
	}
	if err := o.SecureServing.MaybeDefaultWithSelfSignedCerts("localhost", nil, []net.IP{net.ParseIP("127.0.0.1")}); err != nil {
		return nil, fmt.Errorf("error creating self-signed certificates: %v", err)
	}
	c := &cloudcontrollerconfig.Config{}
	if err := o.ApplyTo(c, CloudControllerManagerUserAgent); err != nil {
		return nil, err
	}
	return c, nil
}
func createRecorder(kubeClient clientset.Interface, userAgent string) record.EventRecorder {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	eventBroadcaster := record.NewBroadcaster()
	eventBroadcaster.StartLogging(klog.Infof)
	eventBroadcaster.StartRecordingToSink(&v1core.EventSinkImpl{Interface: kubeClient.CoreV1().Events("")})
	return eventBroadcaster.NewRecorder(legacyscheme.Scheme, v1.EventSource{Component: userAgent})
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
