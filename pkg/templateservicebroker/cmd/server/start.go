package server

import (
	"fmt"
	goformat "fmt"
	"github.com/openshift/origin/pkg/template/servicebroker/apis/config"
	configinstall "github.com/openshift/origin/pkg/template/servicebroker/apis/config/install"
	"github.com/openshift/origin/pkg/templateservicebroker/openservicebroker/server"
	"github.com/spf13/cobra"
	"io"
	"io/ioutil"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	genericapiserver "k8s.io/apiserver/pkg/server"
	genericapiserveroptions "k8s.io/apiserver/pkg/server/options"
	"k8s.io/apiserver/pkg/util/webhook"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	aggregatorapiserver "k8s.io/kube-aggregator/pkg/apiserver"
	"k8s.io/kubernetes/pkg/kubectl/cmd/util"
	"net"
	goos "os"
	godefaultruntime "runtime"
	"time"
	gotime "time"
)

type TemplateServiceBrokerServerOptions struct {
	SecureServing  *genericapiserveroptions.SecureServingOptionsWithLoopback
	Authentication *genericapiserveroptions.DelegatingAuthenticationOptions
	Authorization  *genericapiserveroptions.DelegatingAuthorizationOptions
	Audit          *genericapiserveroptions.AuditOptions
	Features       *genericapiserveroptions.FeatureOptions
	StdOut         io.Writer
	StdErr         io.Writer
	TSBConfig      *config.TemplateServiceBrokerConfig
}

func NewTemplateServiceBrokerServerOptions(out, errOut io.Writer) *TemplateServiceBrokerServerOptions {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	o := &TemplateServiceBrokerServerOptions{SecureServing: genericapiserveroptions.NewSecureServingOptions().WithLoopback(), Authentication: genericapiserveroptions.NewDelegatingAuthenticationOptions(), Authorization: genericapiserveroptions.NewDelegatingAuthorizationOptions(), Audit: genericapiserveroptions.NewAuditOptions(), Features: genericapiserveroptions.NewFeatureOptions(), StdOut: out, StdErr: errOut}
	return o
}
func NewCommandStartTemplateServiceBrokerServer(out, errOut io.Writer, stopCh <-chan struct{}) *cobra.Command {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	o := NewTemplateServiceBrokerServerOptions(out, errOut)
	cmd := &cobra.Command{Use: "template-service-broker", Short: "Launch a template service broker server", Long: "Launch a template service broker server", RunE: func(c *cobra.Command, args []string) error {
		if err := o.Complete(c); err != nil {
			return err
		}
		if err := o.Validate(args); err != nil {
			return err
		}
		if err := o.RunTemplateServiceBrokerServer(stopCh); err != nil {
			return err
		}
		return nil
	}}
	flags := cmd.Flags()
	o.SecureServing.AddFlags(flags)
	o.Authentication.AddFlags(flags)
	o.Authorization.AddFlags(flags)
	o.Audit.AddFlags(flags)
	o.Features.AddFlags(flags)
	flags.String("config", "", "filename containing the TemplateServiceBrokerConfig")
	return cmd
}
func (o TemplateServiceBrokerServerOptions) Validate(args []string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if o.TSBConfig == nil {
		return fmt.Errorf("missing config: specify --config")
	}
	if len(o.TSBConfig.TemplateNamespaces) == 0 {
		return fmt.Errorf("templateNamespaces are required")
	}
	return nil
}
func (o *TemplateServiceBrokerServerOptions) Complete(cmd *cobra.Command) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	configFile := util.GetFlagString(cmd, "config")
	if len(configFile) > 0 {
		content, err := ioutil.ReadFile(configFile)
		if err != nil {
			return err
		}
		configObj, err := runtime.Decode(configCodecs.UniversalDecoder(), content)
		if err != nil {
			return err
		}
		config, ok := configObj.(*config.TemplateServiceBrokerConfig)
		if !ok {
			return fmt.Errorf("unexpected type: %T", configObj)
		}
		o.TSBConfig = config
	}
	return nil
}
func (o TemplateServiceBrokerServerOptions) Config() (*server.TemplateServiceBrokerConfig, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if err := o.SecureServing.MaybeDefaultWithSelfSignedCerts("localhost", nil, []net.IP{net.ParseIP("127.0.0.1")}); err != nil {
		return nil, fmt.Errorf("error creating self-signed certificates: %v", err)
	}
	kubeClientConfig, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}
	kubeClient, err := kubernetes.NewForConfig(kubeClientConfig)
	if err != nil {
		return nil, err
	}
	serverConfig := genericapiserver.NewRecommendedConfig(server.Codecs)
	serverConfig.ClientConfig = kubeClientConfig
	serverConfig.SharedInformerFactory = informers.NewSharedInformerFactory(kubeClient, 10*time.Hour)
	if err := o.SecureServing.ApplyTo(&serverConfig.SecureServing, &serverConfig.LoopbackClientConfig); err != nil {
		return nil, err
	}
	if err := o.Authentication.ApplyTo(&serverConfig.Authentication, serverConfig.SecureServing, serverConfig.OpenAPIConfig); err != nil {
		return nil, err
	}
	if err := o.Authorization.ApplyTo(&serverConfig.Authorization); err != nil {
		return nil, err
	}
	authInfoResolverWrapper := webhook.NewDefaultAuthenticationInfoResolverWrapper(nil, serverConfig.Config.LoopbackClientConfig)
	if err := o.Audit.ApplyTo(&serverConfig.Config, serverConfig.Config.LoopbackClientConfig, serverConfig.SharedInformerFactory, genericapiserveroptions.NewProcessInfo("template-service-broker", "openshift-template-service-broker"), &genericapiserveroptions.WebhookOptions{AuthInfoResolverWrapper: authInfoResolverWrapper, ServiceResolver: aggregatorapiserver.NewClusterIPServiceResolver(serverConfig.SharedInformerFactory.Core().V1().Services().Lister())}); err != nil {
		return nil, err
	}
	if err := o.Features.ApplyTo(&serverConfig.Config); err != nil {
		return nil, err
	}
	serverConfig.EnableMetrics = true
	config := &server.TemplateServiceBrokerConfig{GenericConfig: serverConfig, ExtraConfig: server.ExtraConfig{TemplateNamespaces: o.TSBConfig.TemplateNamespaces}}
	return config, nil
}
func (o TemplateServiceBrokerServerOptions) RunTemplateServiceBrokerServer(stopCh <-chan struct{}) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	config, err := o.Config()
	if err != nil {
		return err
	}
	server, err := config.Complete().New(genericapiserver.NewEmptyDelegate())
	if err != nil {
		return err
	}
	return server.GenericAPIServer.PrepareRun().Run(stopCh)
}

var (
	configScheme = runtime.NewScheme()
	configCodecs = serializer.NewCodecFactory(configScheme)
)

func init() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	configinstall.Install(configScheme)
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
