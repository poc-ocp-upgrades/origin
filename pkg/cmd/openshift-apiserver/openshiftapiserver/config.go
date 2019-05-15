package openshiftapiserver

import (
	"fmt"
	openshiftcontrolplanev1 "github.com/openshift/api/openshiftcontrolplane/v1"
	"github.com/openshift/library-go/pkg/config/helpers"
	"github.com/openshift/origin/pkg/cmd/configflags"
	"github.com/openshift/origin/pkg/cmd/openshift-apiserver/openshiftadmission"
	"github.com/openshift/origin/pkg/cmd/openshift-apiserver/openshiftapiserver/configprocessing"
	"github.com/openshift/origin/pkg/image/apiserver/registryhostname"
	usercache "github.com/openshift/origin/pkg/user/cache"
	"github.com/openshift/origin/pkg/version"
	"github.com/spf13/pflag"
	"io/ioutil"
	"k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/apimachinery/pkg/util/sets"
	genericapiserver "k8s.io/apiserver/pkg/server"
	genericapiserveroptions "k8s.io/apiserver/pkg/server/options"
	"k8s.io/apiserver/pkg/util/webhook"
	cacheddiscovery "k8s.io/client-go/discovery/cached"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/restmapper"
	"k8s.io/client-go/tools/cache"
	"k8s.io/klog"
	aggregatorapiserver "k8s.io/kube-aggregator/pkg/apiserver"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	"net/http"
	"time"
)

func NewOpenshiftAPIConfig(config *openshiftcontrolplanev1.OpenShiftAPIServerConfig) (*OpenshiftAPIConfig, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	kubeClientConfig, err := helpers.GetKubeClientConfig(config.KubeClientConfig)
	if err != nil {
		return nil, err
	}
	kubeClient, err := kubernetes.NewForConfig(kubeClientConfig)
	if err != nil {
		return nil, err
	}
	kubeInformers := informers.NewSharedInformerFactory(kubeClient, 10*time.Minute)
	openshiftVersion := version.Get()
	restOptsGetter, err := NewRESTOptionsGetter(config.APIServerArguments, config.StorageConfig)
	if err != nil {
		return nil, err
	}
	genericConfig := genericapiserver.NewRecommendedConfig(legacyscheme.Codecs)
	genericConfig.SharedInformerFactory = kubeInformers
	genericConfig.ClientConfig = kubeClientConfig
	genericConfig.CorsAllowedOriginList = config.CORSAllowedOrigins
	genericConfig.Version = &openshiftVersion
	genericConfig.ExternalAddress = "apiserver.openshift-apiserver.svc"
	genericConfig.BuildHandlerChainFunc = OpenshiftHandlerChain
	genericConfig.RequestInfoResolver = configprocessing.OpenshiftRequestInfoResolver()
	genericConfig.OpenAPIConfig = configprocessing.DefaultOpenAPIConfig(nil)
	genericConfig.RESTOptionsGetter = restOptsGetter
	genericConfig.RequestTimeout = time.Duration(60) * time.Second
	genericConfig.MinRequestTimeout = int(config.ServingInfo.RequestTimeoutSeconds)
	genericConfig.MaxRequestsInFlight = int(config.ServingInfo.MaxRequestsInFlight)
	genericConfig.MaxMutatingRequestsInFlight = int(config.ServingInfo.MaxRequestsInFlight / 2)
	genericConfig.LongRunningFunc = configprocessing.IsLongRunningRequest
	servingOptions, err := configprocessing.ToServingOptions(config.ServingInfo)
	if err != nil {
		return nil, err
	}
	if err := servingOptions.ApplyTo(&genericConfig.Config.SecureServing, &genericConfig.Config.LoopbackClientConfig); err != nil {
		return nil, err
	}
	authenticationOptions := genericapiserveroptions.NewDelegatingAuthenticationOptions()
	if len(config.AggregatorConfig.ClientCA) > 0 {
		authenticationOptions.ClientCert.ClientCA = config.ServingInfo.ClientCA
		authenticationOptions.RequestHeader.ClientCAFile = config.AggregatorConfig.ClientCA
		authenticationOptions.RequestHeader.AllowedNames = config.AggregatorConfig.AllowedNames
		authenticationOptions.RequestHeader.UsernameHeaders = config.AggregatorConfig.UsernameHeaders
		authenticationOptions.RequestHeader.GroupHeaders = config.AggregatorConfig.GroupHeaders
		authenticationOptions.RequestHeader.ExtraHeaderPrefixes = config.AggregatorConfig.ExtraHeaderPrefixes
	}
	authenticationOptions.RemoteKubeConfigFile = config.KubeClientConfig.KubeConfig
	if err := authenticationOptions.ApplyTo(&genericConfig.Authentication, genericConfig.SecureServing, genericConfig.OpenAPIConfig); err != nil {
		return nil, err
	}
	authorizationOptions := genericapiserveroptions.NewDelegatingAuthorizationOptions().WithAlwaysAllowPaths("/healthz", "/healthz/").WithAlwaysAllowGroups("system:masters")
	authorizationOptions.RemoteKubeConfigFile = config.KubeClientConfig.KubeConfig
	if err := authorizationOptions.ApplyTo(&genericConfig.Authorization); err != nil {
		return nil, err
	}
	informers, err := NewInformers(kubeInformers, kubeClientConfig, genericConfig.LoopbackClientConfig)
	if err != nil {
		return nil, err
	}
	if err := informers.GetOpenshiftUserInformers().User().V1().Groups().Informer().AddIndexers(cache.Indexers{usercache.ByUserIndexName: usercache.ByUserIndexKeys}); err != nil {
		return nil, err
	}
	authInfoResolverWrapper := webhook.NewDefaultAuthenticationInfoResolverWrapper(nil, genericConfig.LoopbackClientConfig)
	auditFlags := configflags.AuditFlags(&config.AuditConfig, configflags.ArgsWithPrefix(config.APIServerArguments, "audit-"))
	auditOpt := genericapiserveroptions.NewAuditOptions()
	fs := pflag.NewFlagSet("audit", pflag.ContinueOnError)
	auditOpt.AddFlags(fs)
	if err := fs.Parse(configflags.ToFlagSlice(auditFlags)); err != nil {
		return nil, err
	}
	if errs := auditOpt.Validate(); len(errs) > 0 {
		return nil, errors.NewAggregate(errs)
	}
	if err := auditOpt.ApplyTo(&genericConfig.Config, genericConfig.Config.LoopbackClientConfig, informers.kubernetesInformers, genericapiserveroptions.NewProcessInfo("openshift-apiserver", "openshift-apiserver"), &genericapiserveroptions.WebhookOptions{AuthInfoResolverWrapper: authInfoResolverWrapper, ServiceResolver: aggregatorapiserver.NewClusterIPServiceResolver(informers.kubernetesInformers.Core().V1().Services().Lister())}); err != nil {
		return nil, err
	}
	projectCache, err := NewProjectCache(informers.kubernetesInformers.Core().V1().Namespaces(), kubeClientConfig, config.ProjectConfig.DefaultNodeSelector)
	if err != nil {
		return nil, err
	}
	clusterQuotaMappingController := NewClusterQuotaMappingController(informers.kubernetesInformers.Core().V1().Namespaces(), informers.quotaInformers.Quota().V1().ClusterResourceQuotas())
	discoveryClient := cacheddiscovery.NewMemCacheClient(kubeClient.Discovery())
	restMapper := restmapper.NewDeferredDiscoveryRESTMapper(discoveryClient)
	admissionInitializer, err := openshiftadmission.NewPluginInitializer(config.ImagePolicyConfig.ExternalRegistryHostnames, config.ImagePolicyConfig.InternalRegistryHostname, config.CloudProviderFile, kubeClientConfig, informers, genericConfig.Authorization.Authorizer, projectCache, restMapper, clusterQuotaMappingController)
	if err != nil {
		return nil, err
	}
	admissionConfigFile, cleanup, err := openshiftadmission.ToAdmissionConfigFile(config.AdmissionConfig.PluginConfig)
	defer cleanup()
	if err != nil {
		return nil, err
	}
	admissionOptions := genericapiserveroptions.NewAdmissionOptions()
	admissionOptions.DefaultOffPlugins = sets.String{}
	admissionOptions.RecommendedPluginOrder = openshiftadmission.OpenShiftAdmissionPlugins
	admissionOptions.Plugins = openshiftadmission.OriginAdmissionPlugins
	admissionOptions.EnablePlugins = config.AdmissionConfig.EnabledAdmissionPlugins
	admissionOptions.DisablePlugins = config.AdmissionConfig.DisabledAdmissionPlugins
	admissionOptions.ConfigFile = admissionConfigFile
	admissionOptions.ApplyTo(&genericConfig.Config, kubeInformers, kubeClientConfig, legacyscheme.Scheme, admissionInitializer)
	var externalRegistryHostname string
	if len(config.ImagePolicyConfig.ExternalRegistryHostnames) > 0 {
		externalRegistryHostname = config.ImagePolicyConfig.ExternalRegistryHostnames[0]
	}
	registryHostnameRetriever, err := registryhostname.DefaultRegistryHostnameRetriever(kubeClientConfig, externalRegistryHostname, config.ImagePolicyConfig.InternalRegistryHostname)
	if err != nil {
		return nil, err
	}
	var caData []byte
	if len(config.ImagePolicyConfig.AdditionalTrustedCA) != 0 {
		klog.V(2).Infof("Image import using additional CA path: %s", config.ImagePolicyConfig.AdditionalTrustedCA)
		var err error
		caData, err = ioutil.ReadFile(config.ImagePolicyConfig.AdditionalTrustedCA)
		if err != nil {
			return nil, fmt.Errorf("failed to read CA bundle %s for image importing: %v", config.ImagePolicyConfig.AdditionalTrustedCA, err)
		}
	}
	subjectLocator := NewSubjectLocator(informers.GetKubernetesInformers().Rbac().V1())
	projectAuthorizationCache := NewProjectAuthorizationCache(subjectLocator, informers.GetKubernetesInformers().Core().V1().Namespaces(), informers.GetKubernetesInformers().Rbac().V1())
	routeAllocator, err := configprocessing.RouteAllocator(config.RoutingConfig.Subdomain)
	if err != nil {
		return nil, err
	}
	ruleResolver := NewRuleResolver(informers.kubernetesInformers.Rbac().V1())
	ret := &OpenshiftAPIConfig{GenericConfig: genericConfig, ExtraConfig: OpenshiftAPIExtraConfig{InformerStart: informers.Start, KubeAPIServerClientConfig: kubeClientConfig, KubeInformers: kubeInformers, QuotaInformers: informers.quotaInformers, SecurityInformers: informers.securityInformers, RuleResolver: ruleResolver, SubjectLocator: subjectLocator, RegistryHostnameRetriever: registryHostnameRetriever, AllowedRegistriesForImport: config.ImagePolicyConfig.AllowedRegistriesForImport, MaxImagesBulkImportedPerRepository: config.ImagePolicyConfig.MaxImagesBulkImportedPerRepository, AdditionalTrustedCA: caData, RouteAllocator: routeAllocator, ProjectAuthorizationCache: projectAuthorizationCache, ProjectCache: projectCache, ProjectRequestTemplate: config.ProjectConfig.ProjectRequestTemplate, ProjectRequestMessage: config.ProjectConfig.ProjectRequestMessage, ClusterQuotaMappingController: clusterQuotaMappingController, RESTMapper: restMapper, ServiceAccountMethod: string(config.ServiceAccountOAuthGrantMethod)}}
	return ret, ret.ExtraConfig.Validate()
}
func OpenshiftHandlerChain(apiHandler http.Handler, genericConfig *genericapiserver.Config) http.Handler {
	_logClusterCodePath()
	defer _logClusterCodePath()
	handler := genericapiserver.DefaultBuildHandlerChain(apiHandler, genericConfig)
	handler = configprocessing.WithCacheControl(handler, "no-store")
	return handler
}
