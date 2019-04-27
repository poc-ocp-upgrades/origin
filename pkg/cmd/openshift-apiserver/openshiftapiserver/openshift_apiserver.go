package openshiftapiserver

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	restful "github.com/emicklei/go-restful"
	"k8s.io/klog"
	kapierror "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	genericapiserver "k8s.io/apiserver/pkg/server"
	genericmux "k8s.io/apiserver/pkg/server/mux"
	kubeinformers "k8s.io/client-go/informers"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/restmapper"
	openapicontroller "k8s.io/kube-aggregator/pkg/controllers/openapi"
	openapiaggregator "k8s.io/kube-aggregator/pkg/controllers/openapi/aggregator"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	kapi "k8s.io/kubernetes/pkg/apis/core"
	coreclient "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/core/internalversion"
	rbacrest "k8s.io/kubernetes/pkg/registry/rbac/rest"
	rbacregistryvalidation "k8s.io/kubernetes/pkg/registry/rbac/validation"
	rbacauthorizer "k8s.io/kubernetes/plugin/pkg/auth/authorizer/rbac"
	openshiftcontrolplanev1 "github.com/openshift/api/openshiftcontrolplane/v1"
	quotainformer "github.com/openshift/client-go/quota/informers/externalversions"
	securityv1informer "github.com/openshift/client-go/security/informers/externalversions"
	oappsapiserver "github.com/openshift/origin/pkg/apps/apiserver"
	authorizationapiserver "github.com/openshift/origin/pkg/authorization/apiserver"
	buildapiserver "github.com/openshift/origin/pkg/build/apiserver"
	"github.com/openshift/origin/pkg/cmd/openshift-apiserver/openshiftapiserver/configprocessing"
	"github.com/openshift/origin/pkg/cmd/server/bootstrappolicy"
	imageapiserver "github.com/openshift/origin/pkg/image/apiserver"
	"github.com/openshift/origin/pkg/image/apiserver/registryhostname"
	networkapiserver "github.com/openshift/origin/pkg/network/apiserver"
	oauthapiserver "github.com/openshift/origin/pkg/oauth/apiserver"
	projectapiserver "github.com/openshift/origin/pkg/project/apiserver"
	projectauth "github.com/openshift/origin/pkg/project/auth"
	projectcache "github.com/openshift/origin/pkg/project/cache"
	quotaapiserver "github.com/openshift/origin/pkg/quota/apiserver"
	"github.com/openshift/origin/pkg/quota/controller/clusterquotamapping"
	routeapiserver "github.com/openshift/origin/pkg/route/apiserver"
	routeallocationcontroller "github.com/openshift/origin/pkg/route/controller/allocation"
	securityapiserver "github.com/openshift/origin/pkg/security/apiserver"
	securityclient "github.com/openshift/origin/pkg/security/generated/internalclientset"
	templateapiserver "github.com/openshift/origin/pkg/template/apiserver"
	userapiserver "github.com/openshift/origin/pkg/user/apiserver"
	"github.com/openshift/origin/pkg/version"
	_ "github.com/openshift/origin/pkg/api/install"
)

type OpenshiftAPIExtraConfig struct {
	InformerStart				func(stopCh <-chan struct{})
	KubeAPIServerClientConfig		*restclient.Config
	KubeInformers				kubeinformers.SharedInformerFactory
	QuotaInformers				quotainformer.SharedInformerFactory
	SecurityInformers			securityv1informer.SharedInformerFactory
	RuleResolver				rbacregistryvalidation.AuthorizationRuleResolver
	SubjectLocator				rbacauthorizer.SubjectLocator
	RegistryHostnameRetriever		registryhostname.RegistryHostnameRetriever
	AllowedRegistriesForImport		openshiftcontrolplanev1.AllowedRegistries
	MaxImagesBulkImportedPerRepository	int
	AdditionalTrustedCA			[]byte
	RouteAllocator				*routeallocationcontroller.RouteAllocationController
	ProjectAuthorizationCache		*projectauth.AuthorizationCache
	ProjectCache				*projectcache.ProjectCache
	ProjectRequestTemplate			string
	ProjectRequestMessage			string
	RESTMapper				*restmapper.DeferredDiscoveryRESTMapper
	ServiceAccountMethod			string
	ClusterQuotaMappingController		*clusterquotamapping.ClusterQuotaMappingController
}

func (c *OpenshiftAPIExtraConfig) Validate() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	ret := []error{}
	if c.KubeInformers == nil {
		ret = append(ret, fmt.Errorf("KubeInformers is required"))
	}
	if c.QuotaInformers == nil {
		ret = append(ret, fmt.Errorf("QuotaInformers is required"))
	}
	if c.SecurityInformers == nil {
		ret = append(ret, fmt.Errorf("SecurityInformers is required"))
	}
	if c.RuleResolver == nil {
		ret = append(ret, fmt.Errorf("RuleResolver is required"))
	}
	if c.SubjectLocator == nil {
		ret = append(ret, fmt.Errorf("SubjectLocator is required"))
	}
	if c.RegistryHostnameRetriever == nil {
		ret = append(ret, fmt.Errorf("RegistryHostnameRetriever is required"))
	}
	if c.RouteAllocator == nil {
		ret = append(ret, fmt.Errorf("RouteAllocator is required"))
	}
	if c.ProjectAuthorizationCache == nil {
		ret = append(ret, fmt.Errorf("ProjectAuthorizationCache is required"))
	}
	if c.ProjectCache == nil {
		ret = append(ret, fmt.Errorf("ProjectCache is required"))
	}
	if c.ClusterQuotaMappingController == nil {
		ret = append(ret, fmt.Errorf("ClusterQuotaMappingController is required"))
	}
	if c.RESTMapper == nil {
		ret = append(ret, fmt.Errorf("RESTMapper is required"))
	}
	return utilerrors.NewAggregate(ret)
}

type OpenshiftAPIConfig struct {
	GenericConfig	*genericapiserver.RecommendedConfig
	ExtraConfig	OpenshiftAPIExtraConfig
}
type OpenshiftAPIServer struct {
	GenericAPIServer *genericapiserver.GenericAPIServer
}
type completedConfig struct {
	GenericConfig	genericapiserver.CompletedConfig
	ExtraConfig	*OpenshiftAPIExtraConfig
}
type CompletedConfig struct{ *completedConfig }

func (c *OpenshiftAPIConfig) Complete() completedConfig {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cfg := completedConfig{c.GenericConfig.Complete(), &c.ExtraConfig}
	return cfg
}
func (c *completedConfig) withAppsAPIServer(delegateAPIServer genericapiserver.DelegationTarget) (genericapiserver.DelegationTarget, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cfg := &oappsapiserver.AppsServerConfig{GenericConfig: &genericapiserver.RecommendedConfig{Config: *c.GenericConfig.Config, SharedInformerFactory: c.GenericConfig.SharedInformerFactory}, ExtraConfig: oappsapiserver.ExtraConfig{KubeAPIServerClientConfig: c.ExtraConfig.KubeAPIServerClientConfig, Codecs: legacyscheme.Codecs, Scheme: legacyscheme.Scheme}}
	config := cfg.Complete()
	server, err := config.New(delegateAPIServer)
	if err != nil {
		return nil, err
	}
	server.GenericAPIServer.PrepareRun()
	return server.GenericAPIServer, nil
}
func (c *completedConfig) withAuthorizationAPIServer(delegateAPIServer genericapiserver.DelegationTarget) (genericapiserver.DelegationTarget, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cfg := &authorizationapiserver.AuthorizationAPIServerConfig{GenericConfig: &genericapiserver.RecommendedConfig{Config: *c.GenericConfig.Config, SharedInformerFactory: c.GenericConfig.SharedInformerFactory}, ExtraConfig: authorizationapiserver.ExtraConfig{KubeAPIServerClientConfig: c.ExtraConfig.KubeAPIServerClientConfig, KubeInformers: c.ExtraConfig.KubeInformers, RuleResolver: c.ExtraConfig.RuleResolver, SubjectLocator: c.ExtraConfig.SubjectLocator, Codecs: legacyscheme.Codecs, Scheme: legacyscheme.Scheme}}
	config := cfg.Complete()
	server, err := config.New(delegateAPIServer)
	if err != nil {
		return nil, err
	}
	server.GenericAPIServer.PrepareRun()
	return server.GenericAPIServer, nil
}
func (c *completedConfig) withBuildAPIServer(delegateAPIServer genericapiserver.DelegationTarget) (genericapiserver.DelegationTarget, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cfg := &buildapiserver.BuildServerConfig{GenericConfig: &genericapiserver.RecommendedConfig{Config: *c.GenericConfig.Config, SharedInformerFactory: c.GenericConfig.SharedInformerFactory}, ExtraConfig: buildapiserver.ExtraConfig{KubeAPIServerClientConfig: c.ExtraConfig.KubeAPIServerClientConfig, Codecs: legacyscheme.Codecs, Scheme: legacyscheme.Scheme}}
	config := cfg.Complete()
	server, err := config.New(delegateAPIServer)
	if err != nil {
		return nil, err
	}
	server.GenericAPIServer.PrepareRun()
	return server.GenericAPIServer, nil
}
func (c *completedConfig) withImageAPIServer(delegateAPIServer genericapiserver.DelegationTarget) (genericapiserver.DelegationTarget, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cfg := &imageapiserver.ImageAPIServerConfig{GenericConfig: &genericapiserver.RecommendedConfig{Config: *c.GenericConfig.Config, SharedInformerFactory: c.GenericConfig.SharedInformerFactory}, ExtraConfig: imageapiserver.ExtraConfig{KubeAPIServerClientConfig: c.ExtraConfig.KubeAPIServerClientConfig, RegistryHostnameRetriever: c.ExtraConfig.RegistryHostnameRetriever, AllowedRegistriesForImport: c.ExtraConfig.AllowedRegistriesForImport, MaxImagesBulkImportedPerRepository: c.ExtraConfig.MaxImagesBulkImportedPerRepository, Codecs: legacyscheme.Codecs, Scheme: legacyscheme.Scheme, AdditionalTrustedCA: c.ExtraConfig.AdditionalTrustedCA}}
	config := cfg.Complete()
	server, err := config.New(delegateAPIServer)
	if err != nil {
		return nil, err
	}
	server.GenericAPIServer.PrepareRun()
	return server.GenericAPIServer, nil
}
func (c *completedConfig) withNetworkAPIServer(delegateAPIServer genericapiserver.DelegationTarget) (genericapiserver.DelegationTarget, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cfg := &networkapiserver.NetworkAPIServerConfig{GenericConfig: &genericapiserver.RecommendedConfig{Config: *c.GenericConfig.Config, SharedInformerFactory: c.GenericConfig.SharedInformerFactory}, ExtraConfig: networkapiserver.ExtraConfig{Codecs: legacyscheme.Codecs, Scheme: legacyscheme.Scheme}}
	config := cfg.Complete()
	server, err := config.New(delegateAPIServer)
	if err != nil {
		return nil, err
	}
	server.GenericAPIServer.PrepareRun()
	return server.GenericAPIServer, nil
}
func (c *completedConfig) withOAuthAPIServer(delegateAPIServer genericapiserver.DelegationTarget) (genericapiserver.DelegationTarget, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cfg := &oauthapiserver.OAuthAPIServerConfig{GenericConfig: &genericapiserver.RecommendedConfig{Config: *c.GenericConfig.Config, SharedInformerFactory: c.GenericConfig.SharedInformerFactory}, ExtraConfig: oauthapiserver.ExtraConfig{KubeAPIServerClientConfig: c.ExtraConfig.KubeAPIServerClientConfig, ServiceAccountMethod: c.ExtraConfig.ServiceAccountMethod, Codecs: legacyscheme.Codecs, Scheme: legacyscheme.Scheme}}
	config := cfg.Complete()
	server, err := config.New(delegateAPIServer)
	if err != nil {
		return nil, err
	}
	server.GenericAPIServer.PrepareRun()
	return server.GenericAPIServer, nil
}
func (c *completedConfig) withProjectAPIServer(delegateAPIServer genericapiserver.DelegationTarget) (genericapiserver.DelegationTarget, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cfg := &projectapiserver.ProjectAPIServerConfig{GenericConfig: &genericapiserver.RecommendedConfig{Config: *c.GenericConfig.Config, SharedInformerFactory: c.GenericConfig.SharedInformerFactory}, ExtraConfig: projectapiserver.ExtraConfig{KubeAPIServerClientConfig: c.ExtraConfig.KubeAPIServerClientConfig, ProjectAuthorizationCache: c.ExtraConfig.ProjectAuthorizationCache, ProjectCache: c.ExtraConfig.ProjectCache, ProjectRequestTemplate: c.ExtraConfig.ProjectRequestTemplate, ProjectRequestMessage: c.ExtraConfig.ProjectRequestMessage, RESTMapper: c.ExtraConfig.RESTMapper, Codecs: legacyscheme.Codecs, Scheme: legacyscheme.Scheme}}
	config := cfg.Complete()
	server, err := config.New(delegateAPIServer)
	if err != nil {
		return nil, err
	}
	server.GenericAPIServer.PrepareRun()
	return server.GenericAPIServer, nil
}
func (c *completedConfig) withQuotaAPIServer(delegateAPIServer genericapiserver.DelegationTarget) (genericapiserver.DelegationTarget, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cfg := &quotaapiserver.QuotaAPIServerConfig{GenericConfig: &genericapiserver.RecommendedConfig{Config: *c.GenericConfig.Config, SharedInformerFactory: c.GenericConfig.SharedInformerFactory}, ExtraConfig: quotaapiserver.ExtraConfig{ClusterQuotaMappingController: c.ExtraConfig.ClusterQuotaMappingController, QuotaInformers: c.ExtraConfig.QuotaInformers, Codecs: legacyscheme.Codecs, Scheme: legacyscheme.Scheme}}
	config := cfg.Complete()
	server, err := config.New(delegateAPIServer)
	if err != nil {
		return nil, err
	}
	server.GenericAPIServer.PrepareRun()
	return server.GenericAPIServer, nil
}
func (c *completedConfig) withRouteAPIServer(delegateAPIServer genericapiserver.DelegationTarget) (genericapiserver.DelegationTarget, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cfg := &routeapiserver.RouteAPIServerConfig{GenericConfig: &genericapiserver.RecommendedConfig{Config: *c.GenericConfig.Config, SharedInformerFactory: c.GenericConfig.SharedInformerFactory}, ExtraConfig: routeapiserver.ExtraConfig{KubeAPIServerClientConfig: c.ExtraConfig.KubeAPIServerClientConfig, RouteAllocator: c.ExtraConfig.RouteAllocator, Codecs: legacyscheme.Codecs, Scheme: legacyscheme.Scheme}}
	config := cfg.Complete()
	server, err := config.New(delegateAPIServer)
	if err != nil {
		return nil, err
	}
	server.GenericAPIServer.PrepareRun()
	return server.GenericAPIServer, nil
}
func (c *completedConfig) withSecurityAPIServer(delegateAPIServer genericapiserver.DelegationTarget) (genericapiserver.DelegationTarget, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cfg := &securityapiserver.SecurityAPIServerConfig{GenericConfig: &genericapiserver.RecommendedConfig{Config: *c.GenericConfig.Config, SharedInformerFactory: c.GenericConfig.SharedInformerFactory}, ExtraConfig: securityapiserver.ExtraConfig{KubeAPIServerClientConfig: c.ExtraConfig.KubeAPIServerClientConfig, SecurityInformers: c.ExtraConfig.SecurityInformers, KubeInformers: c.ExtraConfig.KubeInformers, Authorizer: c.GenericConfig.Authorization.Authorizer, Codecs: legacyscheme.Codecs, Scheme: legacyscheme.Scheme}}
	config := cfg.Complete()
	server, err := config.New(delegateAPIServer)
	if err != nil {
		return nil, err
	}
	server.GenericAPIServer.PrepareRun()
	return server.GenericAPIServer, nil
}
func (c *completedConfig) withTemplateAPIServer(delegateAPIServer genericapiserver.DelegationTarget) (genericapiserver.DelegationTarget, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cfg := &templateapiserver.TemplateConfig{GenericConfig: &genericapiserver.RecommendedConfig{Config: *c.GenericConfig.Config, SharedInformerFactory: c.GenericConfig.SharedInformerFactory}, ExtraConfig: templateapiserver.ExtraConfig{KubeAPIServerClientConfig: c.ExtraConfig.KubeAPIServerClientConfig, Codecs: legacyscheme.Codecs, Scheme: legacyscheme.Scheme}}
	config := cfg.Complete()
	server, err := config.New(delegateAPIServer)
	if err != nil {
		return nil, err
	}
	server.GenericAPIServer.PrepareRun()
	return server.GenericAPIServer, nil
}
func (c *completedConfig) withUserAPIServer(delegateAPIServer genericapiserver.DelegationTarget) (genericapiserver.DelegationTarget, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cfg := &userapiserver.UserConfig{GenericConfig: &genericapiserver.RecommendedConfig{Config: *c.GenericConfig.Config, SharedInformerFactory: c.GenericConfig.SharedInformerFactory}, ExtraConfig: userapiserver.ExtraConfig{Codecs: legacyscheme.Codecs, Scheme: legacyscheme.Scheme}}
	config := cfg.Complete()
	server, err := config.New(delegateAPIServer)
	if err != nil {
		return nil, err
	}
	server.GenericAPIServer.PrepareRun()
	return server.GenericAPIServer, nil
}
func (c *completedConfig) withOpenAPIAggregationController(delegatedAPIServer *genericapiserver.GenericAPIServer) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	delegatedAPIServer.RemoveOpenAPIData()
	specDownloader := openapiaggregator.NewDownloader()
	openAPIAggregator, err := openapiaggregator.BuildAndRegisterAggregator(&specDownloader, delegatedAPIServer, delegatedAPIServer.Handler.GoRestfulContainer.RegisteredWebServices(), configprocessing.DefaultOpenAPIConfig(nil), delegatedAPIServer.Handler.NonGoRestfulMux)
	if err != nil {
		return err
	}
	openAPIAggregationController := openapicontroller.NewAggregationController(&specDownloader, openAPIAggregator)
	delegatedAPIServer.AddPostStartHook("apiservice-openapi-controller", func(context genericapiserver.PostStartHookContext) error {
		go openAPIAggregationController.Run(context.StopCh)
		return nil
	})
	return nil
}

type apiServerAppenderFunc func(delegateAPIServer genericapiserver.DelegationTarget) (genericapiserver.DelegationTarget, error)

func addAPIServerOrDie(delegateAPIServer genericapiserver.DelegationTarget, apiServerAppenderFn apiServerAppenderFunc) genericapiserver.DelegationTarget {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	delegateAPIServer, err := apiServerAppenderFn(delegateAPIServer)
	if err != nil {
		klog.Fatal(err)
	}
	return delegateAPIServer
}
func (c completedConfig) New(delegationTarget genericapiserver.DelegationTarget, keepRemovedNetworkingAPIs bool) (*OpenshiftAPIServer, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	delegateAPIServer := delegationTarget
	delegateAPIServer = addAPIServerOrDie(delegateAPIServer, c.withAppsAPIServer)
	delegateAPIServer = addAPIServerOrDie(delegateAPIServer, c.withAuthorizationAPIServer)
	delegateAPIServer = addAPIServerOrDie(delegateAPIServer, c.withBuildAPIServer)
	delegateAPIServer = addAPIServerOrDie(delegateAPIServer, c.withImageAPIServer)
	if keepRemovedNetworkingAPIs {
		delegateAPIServer = addAPIServerOrDie(delegateAPIServer, c.withNetworkAPIServer)
	}
	delegateAPIServer = addAPIServerOrDie(delegateAPIServer, c.withOAuthAPIServer)
	delegateAPIServer = addAPIServerOrDie(delegateAPIServer, c.withProjectAPIServer)
	delegateAPIServer = addAPIServerOrDie(delegateAPIServer, c.withQuotaAPIServer)
	delegateAPIServer = addAPIServerOrDie(delegateAPIServer, c.withRouteAPIServer)
	delegateAPIServer = addAPIServerOrDie(delegateAPIServer, c.withSecurityAPIServer)
	delegateAPIServer = addAPIServerOrDie(delegateAPIServer, c.withTemplateAPIServer)
	delegateAPIServer = addAPIServerOrDie(delegateAPIServer, c.withUserAPIServer)
	genericServer, err := c.GenericConfig.New("openshift-apiserver", delegateAPIServer)
	if err != nil {
		return nil, err
	}
	if err := c.withOpenAPIAggregationController(genericServer); err != nil {
		return nil, err
	}
	s := &OpenshiftAPIServer{GenericAPIServer: genericServer}
	addReadinessCheckRoute(s.GenericAPIServer.Handler.NonGoRestfulMux, "/healthz/ready", c.ExtraConfig.ProjectAuthorizationCache.ReadyForAccess)
	AddOpenshiftVersionRoute(s.GenericAPIServer.Handler.GoRestfulContainer, "/version/openshift")
	s.GenericAPIServer.AddPostStartHookOrDie("authorization.openshift.io-bootstrapclusterroles", func(context genericapiserver.PostStartHookContext) error {
		newContext := genericapiserver.PostStartHookContext{LoopbackClientConfig: c.ExtraConfig.KubeAPIServerClientConfig, StopCh: context.StopCh}
		return bootstrapData(bootstrappolicy.Policy()).EnsureRBACPolicy()(newContext)
	})
	s.GenericAPIServer.AddPostStartHookOrDie("authorization.openshift.io-ensureopenshift-infra", c.EnsureOpenShiftInfraNamespace)
	s.GenericAPIServer.AddPostStartHookOrDie("project.openshift.io-projectcache", c.startProjectCache)
	s.GenericAPIServer.AddPostStartHookOrDie("project.openshift.io-projectauthorizationcache", c.startProjectAuthorizationCache)
	s.GenericAPIServer.AddPostStartHookOrDie("security.openshift.io-bootstrapscc", c.bootstrapSCC)
	s.GenericAPIServer.AddPostStartHookOrDie("openshift.io-startinformers", func(context genericapiserver.PostStartHookContext) error {
		c.ExtraConfig.InformerStart(context.StopCh)
		return nil
	})
	s.GenericAPIServer.AddPostStartHookOrDie("openshift.io-restmapperupdater", func(context genericapiserver.PostStartHookContext) error {
		go func() {
			wait.Until(func() {
				c.ExtraConfig.RESTMapper.Reset()
			}, 10*time.Second, context.StopCh)
		}()
		return nil
	})
	s.GenericAPIServer.AddPostStartHookOrDie("quota.openshift.io-clusterquotamapping", func(context genericapiserver.PostStartHookContext) error {
		go c.ExtraConfig.ClusterQuotaMappingController.Run(5, context.StopCh)
		return nil
	})
	return s, nil
}
func addReadinessCheckRoute(mux *genericmux.PathRecorderMux, path string, readyFunc func() bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	mux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
		if readyFunc() {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("ok"))
		} else {
			w.WriteHeader(http.StatusServiceUnavailable)
		}
	})
}
func AddOpenshiftVersionRoute(container *restful.Container, path string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	versionInfo, err := json.MarshalIndent(version.Get(), "", "  ")
	if err != nil {
		klog.Errorf("Unable to initialize version route: %v", err)
		return
	}
	ws := new(restful.WebService)
	ws.Path(path)
	ws.Doc("git code version from which this is built")
	ws.Route(ws.GET("/").To(func(_ *restful.Request, resp *restful.Response) {
		writeJSON(resp, versionInfo)
	}).Doc("get the code version").Operation("getCodeVersion").Produces(restful.MIME_JSON))
	container.Add(ws)
}
func writeJSON(resp *restful.Response, json []byte) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	resp.ResponseWriter.Header().Set("Content-Type", "application/json")
	resp.ResponseWriter.WriteHeader(http.StatusOK)
	resp.ResponseWriter.Write(json)
}
func (c *completedConfig) startProjectCache(context genericapiserver.PostStartHookContext) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	klog.Infof("Using default project node label selector: %s", c.ExtraConfig.ProjectCache.DefaultNodeSelector)
	go c.ExtraConfig.ProjectCache.Run(context.StopCh)
	return nil
}
func (c *completedConfig) startProjectAuthorizationCache(context genericapiserver.PostStartHookContext) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	period := 1 * time.Second
	c.ExtraConfig.ProjectAuthorizationCache.Run(period)
	return nil
}
func (c *completedConfig) bootstrapSCC(context genericapiserver.PostStartHookContext) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	ns := bootstrappolicy.DefaultOpenShiftInfraNamespace
	bootstrapSCCGroups, bootstrapSCCUsers := bootstrappolicy.GetBoostrapSCCAccess(ns)
	var securityClient securityclient.Interface
	err := wait.Poll(1*time.Second, 30*time.Second, func() (bool, error) {
		var err error
		securityClient, err = securityclient.NewForConfig(context.LoopbackClientConfig)
		if err != nil {
			utilruntime.HandleError(fmt.Errorf("unable to initialize client: %v", err))
			return false, nil
		}
		return true, nil
	})
	if err != nil {
		utilruntime.HandleError(fmt.Errorf("error getting client: %v", err))
		return err
	}
	for _, scc := range bootstrappolicy.GetBootstrapSecurityContextConstraints(bootstrapSCCGroups, bootstrapSCCUsers) {
		_, err := securityClient.Security().SecurityContextConstraints().Create(scc)
		if kapierror.IsAlreadyExists(err) {
			continue
		}
		if err != nil {
			utilruntime.HandleError(fmt.Errorf("unable to create default security context constraint %s.  Got error: %v", scc.Name, err))
			continue
		}
		klog.Infof("Created default security context constraint %s", scc.Name)
	}
	return nil
}
func (c *completedConfig) EnsureOpenShiftInfraNamespace(context genericapiserver.PostStartHookContext) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	namespaceName := bootstrappolicy.DefaultOpenShiftInfraNamespace
	var coreClient coreclient.CoreInterface
	err := wait.Poll(1*time.Second, 30*time.Second, func() (bool, error) {
		var err error
		coreClient, err = coreclient.NewForConfig(c.ExtraConfig.KubeAPIServerClientConfig)
		if err != nil {
			utilruntime.HandleError(fmt.Errorf("unable to initialize client: %v", err))
			return false, nil
		}
		return true, nil
	})
	if err != nil {
		utilruntime.HandleError(fmt.Errorf("error getting client: %v", err))
		return err
	}
	_, err = coreClient.Namespaces().Create(&kapi.Namespace{ObjectMeta: metav1.ObjectMeta{Name: namespaceName}})
	if err != nil && !kapierror.IsAlreadyExists(err) {
		utilruntime.HandleError(fmt.Errorf("error creating namespace %q: %v", namespaceName, err))
		return err
	}
	_, err = coreClient.ServiceAccounts(namespaceName).Create(&kapi.ServiceAccount{ObjectMeta: metav1.ObjectMeta{Name: bootstrappolicy.InfraNodeBootstrapServiceAccountName}})
	if err != nil && !kapierror.IsAlreadyExists(err) {
		klog.Errorf("Error creating service account %s/%s: %v", namespaceName, bootstrappolicy.InfraNodeBootstrapServiceAccountName, err)
	}
	return nil
}
func bootstrapData(data *bootstrappolicy.PolicyData) *rbacrest.PolicyData {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &rbacrest.PolicyData{ClusterRoles: data.ClusterRoles, ClusterRoleBindings: data.ClusterRoleBindings, Roles: data.Roles, RoleBindings: data.RoleBindings, ClusterRolesToAggregate: data.ClusterRolesToAggregate}
}
