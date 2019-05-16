package apiserver

import (
	"fmt"
	goformat "fmt"
	routeapiv1 "github.com/openshift/api/route/v1"
	routeetcd "github.com/openshift/origin/pkg/route/apiserver/registry/route/etcd"
	routeallocationcontroller "github.com/openshift/origin/pkg/route/controller/allocation"
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
	RouteAllocator            *routeallocationcontroller.RouteAllocationController
	Scheme                    *runtime.Scheme
	Codecs                    serializer.CodecFactory
	makeV1Storage             sync.Once
	v1Storage                 map[string]rest.Storage
	v1StorageErr              error
}
type RouteAPIServerConfig struct {
	GenericConfig *genericapiserver.RecommendedConfig
	ExtraConfig   ExtraConfig
}
type RouteAPIServer struct {
	GenericAPIServer *genericapiserver.GenericAPIServer
}
type completedConfig struct {
	GenericConfig genericapiserver.CompletedConfig
	ExtraConfig   *ExtraConfig
}
type CompletedConfig struct{ *completedConfig }

func (c *RouteAPIServerConfig) Complete() completedConfig {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	cfg := completedConfig{c.GenericConfig.Complete(), &c.ExtraConfig}
	return cfg
}
func (c completedConfig) New(delegationTarget genericapiserver.DelegationTarget) (*RouteAPIServer, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	genericServer, err := c.GenericConfig.New("route.openshift.io-apiserver", delegationTarget)
	if err != nil {
		return nil, err
	}
	s := &RouteAPIServer{GenericAPIServer: genericServer}
	v1Storage, err := c.V1RESTStorage()
	if err != nil {
		return nil, err
	}
	apiGroupInfo := genericapiserver.NewDefaultAPIGroupInfo(routeapiv1.GroupName, c.ExtraConfig.Scheme, metav1.ParameterCodec, c.ExtraConfig.Codecs)
	apiGroupInfo.VersionedResourcesStorageMap[routeapiv1.SchemeGroupVersion.Version] = v1Storage
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
	routeStorage, routeStatusStorage, err := routeetcd.NewREST(c.GenericConfig.RESTOptionsGetter, c.ExtraConfig.RouteAllocator, authorizationClient.SubjectAccessReviews())
	if err != nil {
		return nil, fmt.Errorf("error building REST storage: %v", err)
	}
	v1Storage := map[string]rest.Storage{}
	v1Storage["routes"] = routeStorage
	v1Storage["routes/status"] = routeStatusStorage
	return v1Storage, nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
