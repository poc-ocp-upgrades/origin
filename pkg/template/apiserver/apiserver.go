package apiserver

import (
	goformat "fmt"
	templateapiv1 "github.com/openshift/api/template/v1"
	brokertemplateinstanceetcd "github.com/openshift/origin/pkg/template/apiserver/registry/brokertemplateinstance/etcd"
	templateregistry "github.com/openshift/origin/pkg/template/apiserver/registry/template"
	templateetcd "github.com/openshift/origin/pkg/template/apiserver/registry/template/etcd"
	templateinstanceetcd "github.com/openshift/origin/pkg/template/apiserver/registry/templateinstance/etcd"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/apiserver/pkg/registry/rest"
	genericapiserver "k8s.io/apiserver/pkg/server"
	authorizationclient "k8s.io/client-go/kubernetes/typed/authorization/v1"
	restclient "k8s.io/client-go/rest"
	goos "os"
	godefaultruntime "runtime"
	"sync"
	gotime "time"
)

type ExtraConfig struct {
	KubeAPIServerClientConfig *restclient.Config
	Scheme                    *runtime.Scheme
	Codecs                    serializer.CodecFactory
	makeV1Storage             sync.Once
	v1Storage                 map[string]rest.Storage
	v1StorageErr              error
}
type TemplateConfig struct {
	GenericConfig *genericapiserver.RecommendedConfig
	ExtraConfig   ExtraConfig
}
type TemplateServer struct {
	GenericAPIServer *genericapiserver.GenericAPIServer
}
type completedConfig struct {
	GenericConfig genericapiserver.CompletedConfig
	ExtraConfig   *ExtraConfig
}
type CompletedConfig struct{ *completedConfig }

func (c *TemplateConfig) Complete() completedConfig {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	cfg := completedConfig{c.GenericConfig.Complete(), &c.ExtraConfig}
	return cfg
}
func (c completedConfig) New(delegationTarget genericapiserver.DelegationTarget) (*TemplateServer, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	genericServer, err := c.GenericConfig.New("template.openshift.io-apiserver", delegationTarget)
	if err != nil {
		return nil, err
	}
	s := &TemplateServer{GenericAPIServer: genericServer}
	v1Storage, err := c.V1RESTStorage()
	if err != nil {
		return nil, err
	}
	apiGroupInfo := genericapiserver.NewDefaultAPIGroupInfo(templateapiv1.GroupName, c.ExtraConfig.Scheme, metav1.ParameterCodec, c.ExtraConfig.Codecs)
	apiGroupInfo.VersionedResourcesStorageMap[templateapiv1.SchemeGroupVersion.Version] = v1Storage
	if err := s.GenericAPIServer.InstallAPIGroup(&apiGroupInfo); err != nil {
		return nil, err
	}
	return s, nil
}
func (c *completedConfig) V1RESTStorage() (map[string]rest.Storage, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	c.ExtraConfig.makeV1Storage.Do(func() {
		c.ExtraConfig.v1Storage, c.ExtraConfig.v1StorageErr = c.newV1RESTStorage()
	})
	return c.ExtraConfig.v1Storage, c.ExtraConfig.v1StorageErr
}
func (c *completedConfig) newV1RESTStorage() (map[string]rest.Storage, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	authorizationClient, err := authorizationclient.NewForConfig(c.ExtraConfig.KubeAPIServerClientConfig)
	if err != nil {
		return nil, err
	}
	templateStorage, err := templateetcd.NewREST(c.GenericConfig.RESTOptionsGetter)
	if err != nil {
		return nil, err
	}
	templateInstanceStorage, templateInstanceStatusStorage, err := templateinstanceetcd.NewREST(c.GenericConfig.RESTOptionsGetter, authorizationClient)
	if err != nil {
		return nil, err
	}
	brokerTemplateInstanceStorage, err := brokertemplateinstanceetcd.NewREST(c.GenericConfig.RESTOptionsGetter)
	if err != nil {
		return nil, err
	}
	v1Storage := map[string]rest.Storage{}
	v1Storage["processedTemplates"] = templateregistry.NewREST()
	v1Storage["templates"] = templateStorage
	v1Storage["templateinstances"] = templateInstanceStorage
	v1Storage["templateinstances/status"] = templateInstanceStatusStorage
	v1Storage["brokertemplateinstances"] = brokerTemplateInstanceStorage
	return v1Storage, nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
