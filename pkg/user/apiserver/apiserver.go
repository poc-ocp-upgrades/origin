package apiserver

import (
	goformat "fmt"
	userapiv1 "github.com/openshift/api/user/v1"
	userclient "github.com/openshift/client-go/user/clientset/versioned"
	groupetcd "github.com/openshift/origin/pkg/user/apiserver/registry/group/etcd"
	identityetcd "github.com/openshift/origin/pkg/user/apiserver/registry/identity/etcd"
	useretcd "github.com/openshift/origin/pkg/user/apiserver/registry/user/etcd"
	"github.com/openshift/origin/pkg/user/apiserver/registry/useridentitymapping"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/apiserver/pkg/registry/rest"
	genericapiserver "k8s.io/apiserver/pkg/server"
	goos "os"
	godefaultruntime "runtime"
	"sync"
	gotime "time"
)

type ExtraConfig struct {
	Scheme        *runtime.Scheme
	Codecs        serializer.CodecFactory
	makeV1Storage sync.Once
	v1Storage     map[string]rest.Storage
	v1StorageErr  error
}
type UserConfig struct {
	GenericConfig *genericapiserver.RecommendedConfig
	ExtraConfig   ExtraConfig
}
type UserServer struct {
	GenericAPIServer *genericapiserver.GenericAPIServer
}
type completedConfig struct {
	GenericConfig genericapiserver.CompletedConfig
	ExtraConfig   *ExtraConfig
}
type CompletedConfig struct{ *completedConfig }

func (c *UserConfig) Complete() completedConfig {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	cfg := completedConfig{c.GenericConfig.Complete(), &c.ExtraConfig}
	return cfg
}
func (c completedConfig) New(delegationTarget genericapiserver.DelegationTarget) (*UserServer, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	genericServer, err := c.GenericConfig.New("user.openshift.io-apiserver", delegationTarget)
	if err != nil {
		return nil, err
	}
	s := &UserServer{GenericAPIServer: genericServer}
	v1Storage, err := c.V1RESTStorage()
	if err != nil {
		return nil, err
	}
	apiGroupInfo := genericapiserver.NewDefaultAPIGroupInfo(userapiv1.GroupName, c.ExtraConfig.Scheme, metav1.ParameterCodec, c.ExtraConfig.Codecs)
	apiGroupInfo.VersionedResourcesStorageMap[userapiv1.SchemeGroupVersion.Version] = v1Storage
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
	userClient, err := userclient.NewForConfig(c.GenericConfig.LoopbackClientConfig)
	if err != nil {
		return nil, err
	}
	userStorage, err := useretcd.NewREST(c.GenericConfig.RESTOptionsGetter)
	if err != nil {
		return nil, err
	}
	identityStorage, err := identityetcd.NewREST(c.GenericConfig.RESTOptionsGetter)
	if err != nil {
		return nil, err
	}
	userIdentityMappingStorage := useridentitymapping.NewREST(userClient.UserV1().Users(), userClient.UserV1().Identities())
	groupStorage, err := groupetcd.NewREST(c.GenericConfig.RESTOptionsGetter)
	if err != nil {
		return nil, err
	}
	v1Storage := map[string]rest.Storage{}
	v1Storage["users"] = userStorage
	v1Storage["groups"] = groupStorage
	v1Storage["identities"] = identityStorage
	v1Storage["userIdentityMappings"] = userIdentityMappingStorage
	return v1Storage, nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
