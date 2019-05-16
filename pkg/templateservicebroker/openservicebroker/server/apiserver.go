package server

import (
	"fmt"
	goformat "fmt"
	templateapiv1 "github.com/openshift/api/template/v1"
	templateclientset "github.com/openshift/client-go/template/clientset/versioned"
	templateinformer "github.com/openshift/client-go/template/informers/externalversions"
	templateapi "github.com/openshift/origin/pkg/template/apis/template"
	templateservicebroker "github.com/openshift/origin/pkg/templateservicebroker/servicebroker"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	genericapiserver "k8s.io/apiserver/pkg/server"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/kubernetes/pkg/apis/core/install"
	"k8s.io/kubernetes/pkg/controller"
	goos "os"
	godefaultruntime "runtime"
	"time"
	gotime "time"
)

var (
	Scheme             = runtime.NewScheme()
	Codecs             = serializer.NewCodecFactory(Scheme)
	unversionedVersion = schema.GroupVersion{Group: "", Version: "v1"}
	unversionedTypes   = []runtime.Object{&metav1.Status{}, &metav1.WatchEvent{}, &metav1.APIVersions{}, &metav1.APIGroupList{}, &metav1.APIGroup{}, &metav1.APIResourceList{}}
)

func init() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	install.Install(Scheme)
	metav1.AddToGroupVersion(Scheme, schema.GroupVersion{Group: "", Version: "v1"})
	Scheme.AddUnversionedTypes(unversionedVersion, unversionedTypes...)
}

type ExtraConfig struct{ TemplateNamespaces []string }
type TemplateServiceBrokerConfig struct {
	GenericConfig *genericapiserver.RecommendedConfig
	ExtraConfig   ExtraConfig
}
type TemplateServiceBrokerServer struct {
	GenericAPIServer *genericapiserver.GenericAPIServer
}
type completedTemplateServiceBrokerConfig struct {
	GenericConfig genericapiserver.CompletedConfig
	ExtraConfig   *ExtraConfig
}

func (c *TemplateServiceBrokerConfig) Complete() completedTemplateServiceBrokerConfig {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	cfg := completedTemplateServiceBrokerConfig{c.GenericConfig.Complete(), &c.ExtraConfig}
	return cfg
}
func (c completedTemplateServiceBrokerConfig) New(delegationTarget genericapiserver.DelegationTarget) (*TemplateServiceBrokerServer, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	genericServer, err := c.GenericConfig.New("template-service-broker", delegationTarget)
	if err != nil {
		return nil, err
	}
	s := &TemplateServiceBrokerServer{GenericAPIServer: genericServer}
	clientConfig, err := restclient.InClusterConfig()
	if err != nil {
		return nil, err
	}
	templateClient, err := templateclientset.NewForConfig(clientConfig)
	if err != nil {
		return nil, err
	}
	templateInformers := templateinformer.NewSharedInformerFactory(templateClient, 5*time.Minute)
	templateInformers.Template().V1().Templates().Informer().AddIndexers(cache.Indexers{templateapi.TemplateUIDIndex: func(obj interface{}) ([]string, error) {
		return []string{string(obj.(*templateapiv1.Template).UID)}, nil
	}})
	broker, err := templateservicebroker.NewBroker(clientConfig, templateInformers.Template().V1().Templates(), c.ExtraConfig.TemplateNamespaces)
	if err != nil {
		return nil, err
	}
	if err := s.GenericAPIServer.AddPostStartHook("template-service-broker-synctemplates", func(context genericapiserver.PostStartHookContext) error {
		templateInformers.Start(context.StopCh)
		if !controller.WaitForCacheSync("tsb", context.StopCh, templateInformers.Template().V1().Templates().Informer().HasSynced) {
			return fmt.Errorf("unable to sync caches")
		}
		return nil
	}); err != nil {
		return nil, err
	}
	Route(s.GenericAPIServer.Handler.GoRestfulContainer, templateapi.ServiceBrokerRoot, broker)
	return s, nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
