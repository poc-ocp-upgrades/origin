package servicebroker

import (
	templateclientset "github.com/openshift/client-go/template/clientset/versioned"
	v1template "github.com/openshift/client-go/template/clientset/versioned/typed/template/v1"
	templateinformer "github.com/openshift/client-go/template/informers/externalversions/template/v1"
	templatelister "github.com/openshift/client-go/template/listers/template/v1"
	"github.com/openshift/origin/pkg/templateservicebroker/openservicebroker/api"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/discovery"
	cacheddiscovery "k8s.io/client-go/discovery/cached"
	"k8s.io/client-go/dynamic"
	kclientset "k8s.io/client-go/kubernetes"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/restmapper"
	"os"
	"time"
)

type Broker struct {
	kc                 kclientset.Interface
	templateclient     v1template.TemplateV1Interface
	lister             templatelister.TemplateLister
	hasSynced          func() bool
	templateNamespaces map[string]struct{}
	restmapper         meta.RESTMapper
	dynamicClient      dynamic.Interface
	gcCreateDelay      time.Duration
}

var _ api.Broker = &Broker{}

func NewBroker(saKubeClientConfig *restclient.Config, informer templateinformer.TemplateInformer, namespaces []string) (*Broker, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	templateNamespaces := map[string]struct{}{}
	for _, namespace := range namespaces {
		templateNamespaces[namespace] = struct{}{}
	}
	kubeClient, err := kclientset.NewForConfig(saKubeClientConfig)
	if err != nil {
		return nil, err
	}
	templateClient, err := templateclientset.NewForConfig(saKubeClientConfig)
	if err != nil {
		return nil, err
	}
	dynamicClient, err := dynamic.NewForConfig(saKubeClientConfig)
	if err != nil {
		return nil, err
	}
	discoveryConfig := restclient.CopyConfig(saKubeClientConfig)
	discoveryConfig.Burst = 100
	discoveryClient, err := discovery.NewDiscoveryClientForConfig(discoveryConfig)
	if err != nil {
		return nil, err
	}
	cachedDiscovery := cacheddiscovery.NewMemCacheClient(discoveryClient)
	restMapper := restmapper.NewDeferredDiscoveryRESTMapper(cachedDiscovery)
	restMapper.Reset()
	go wait.Until(restMapper.Reset, 30*time.Second, wait.NeverStop)
	delay := 5 * time.Second
	value := os.Getenv("TEMPLATE_SERVICE_BROKER_GC_DELAY")
	if len(value) != 0 {
		if v, err := time.ParseDuration(value); err == nil {
			delay = v
		}
	}
	b := &Broker{kc: kubeClient, templateclient: templateClient.TemplateV1(), lister: informer.Lister(), hasSynced: informer.Informer().HasSynced, templateNamespaces: templateNamespaces, restmapper: restMapper, dynamicClient: dynamicClient, gcCreateDelay: delay}
	return b, nil
}
