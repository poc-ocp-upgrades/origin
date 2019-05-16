package app

import (
	"fmt"
	goformat "fmt"
	apiextensionsinformers "k8s.io/apiextensions-apiserver/pkg/client/informers/internalversion"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apiserver/pkg/admission"
	"k8s.io/apiserver/pkg/features"
	genericapiserver "k8s.io/apiserver/pkg/server"
	"k8s.io/apiserver/pkg/server/healthz"
	genericoptions "k8s.io/apiserver/pkg/server/options"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	kubeexternalinformers "k8s.io/client-go/informers"
	"k8s.io/client-go/tools/cache"
	"k8s.io/klog"
	"k8s.io/kube-aggregator/pkg/apis/apiregistration"
	"k8s.io/kube-aggregator/pkg/apis/apiregistration/v1"
	"k8s.io/kube-aggregator/pkg/apis/apiregistration/v1beta1"
	aggregatorapiserver "k8s.io/kube-aggregator/pkg/apiserver"
	aggregatorscheme "k8s.io/kube-aggregator/pkg/apiserver/scheme"
	apiregistrationclient "k8s.io/kube-aggregator/pkg/client/clientset_generated/internalclientset/typed/apiregistration/internalversion"
	informers "k8s.io/kube-aggregator/pkg/client/informers/internalversion/apiregistration/internalversion"
	"k8s.io/kube-aggregator/pkg/controllers/autoregister"
	"k8s.io/kubernetes/cmd/kube-apiserver/app/options"
	"k8s.io/kubernetes/pkg/master/controller/crdregistration"
	"net/http"
	goos "os"
	godefaultruntime "runtime"
	"strings"
	"sync"
	gotime "time"
)

func createAggregatorConfig(kubeAPIServerConfig genericapiserver.Config, commandOptions *options.ServerRunOptions, externalInformers kubeexternalinformers.SharedInformerFactory, serviceResolver aggregatorapiserver.ServiceResolver, proxyTransport *http.Transport, pluginInitializers []admission.PluginInitializer) (*aggregatorapiserver.Config, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	genericConfig := kubeAPIServerConfig
	commandOptions.Admission.ApplyTo(&genericConfig, externalInformers, genericConfig.LoopbackClientConfig, aggregatorscheme.Scheme, pluginInitializers...)
	genericConfig.EnableSwaggerUI = false
	genericConfig.SwaggerConfig = nil
	etcdOptions := *commandOptions.Etcd
	etcdOptions.StorageConfig.Paging = utilfeature.DefaultFeatureGate.Enabled(features.APIListChunking)
	etcdOptions.StorageConfig.Codec = aggregatorscheme.Codecs.LegacyCodec(v1beta1.SchemeGroupVersion, v1.SchemeGroupVersion)
	genericConfig.RESTOptionsGetter = &genericoptions.SimpleRestOptionsFactory{Options: etcdOptions}
	if err := commandOptions.APIEnablement.ApplyTo(&genericConfig, aggregatorapiserver.DefaultAPIResourceConfigSource(), aggregatorscheme.Scheme); err != nil {
		return nil, err
	}
	aggregatorConfig := &aggregatorapiserver.Config{GenericConfig: &genericapiserver.RecommendedConfig{Config: genericConfig, SharedInformerFactory: externalInformers}, ExtraConfig: aggregatorapiserver.ExtraConfig{ProxyClientCert: commandOptions.ProxyClientCertFile, ProxyClientKey: commandOptions.ProxyClientKeyFile, ServiceResolver: serviceResolver, ProxyTransport: proxyTransport}}
	return aggregatorConfig, nil
}
func createAggregatorServer(aggregatorConfig *aggregatorapiserver.Config, delegateAPIServer genericapiserver.DelegationTarget, apiExtensionInformers apiextensionsinformers.SharedInformerFactory) (*aggregatorapiserver.APIAggregator, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	aggregatorServer, err := aggregatorConfig.Complete().NewWithDelegate(delegateAPIServer)
	if err != nil {
		return nil, err
	}
	apiRegistrationClient, err := apiregistrationclient.NewForConfig(aggregatorConfig.GenericConfig.LoopbackClientConfig)
	if err != nil {
		return nil, err
	}
	autoRegistrationController := autoregister.NewAutoRegisterController(aggregatorServer.APIRegistrationInformers.Apiregistration().InternalVersion().APIServices(), apiRegistrationClient)
	apiServices := apiServicesToRegister(delegateAPIServer, autoRegistrationController)
	crdRegistrationController := crdregistration.NewAutoRegistrationController(apiExtensionInformers.Apiextensions().InternalVersion().CustomResourceDefinitions(), autoRegistrationController)
	aggregatorServer.GenericAPIServer.AddPostStartHook("kube-apiserver-autoregistration", func(context genericapiserver.PostStartHookContext) error {
		go crdRegistrationController.Run(5, context.StopCh)
		go func() {
			if aggregatorConfig.GenericConfig.MergedResourceConfig.AnyVersionForGroupEnabled("apiextensions.k8s.io") {
				crdRegistrationController.WaitForInitialSync()
			}
			autoRegistrationController.Run(5, context.StopCh)
		}()
		return nil
	})
	aggregatorServer.GenericAPIServer.AddHealthzChecks(makeAPIServiceAvailableHealthzCheck("autoregister-completion", apiServices, aggregatorServer.APIRegistrationInformers.Apiregistration().InternalVersion().APIServices()))
	return aggregatorServer, nil
}
func makeAPIService(gv schema.GroupVersion) *apiregistration.APIService {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	apiServicePriority, ok := apiVersionPriorities[gv]
	if !ok {
		klog.Infof("Skipping APIService creation for %v", gv)
		return nil
	}
	return &apiregistration.APIService{ObjectMeta: metav1.ObjectMeta{Name: gv.Version + "." + gv.Group}, Spec: apiregistration.APIServiceSpec{Group: gv.Group, Version: gv.Version, GroupPriorityMinimum: apiServicePriority.group, VersionPriority: apiServicePriority.version}}
}
func makeAPIServiceAvailableHealthzCheck(name string, apiServices []*apiregistration.APIService, apiServiceInformer informers.APIServiceInformer) healthz.HealthzChecker {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	pendingServiceNamesLock := &sync.RWMutex{}
	pendingServiceNames := sets.NewString()
	for _, service := range apiServices {
		pendingServiceNames.Insert(service.Name)
	}
	handleAPIServiceChange := func(service *apiregistration.APIService) {
		pendingServiceNamesLock.Lock()
		defer pendingServiceNamesLock.Unlock()
		if !pendingServiceNames.Has(service.Name) {
			return
		}
		if apiregistration.IsAPIServiceConditionTrue(service, apiregistration.Available) {
			pendingServiceNames.Delete(service.Name)
		}
	}
	apiServiceInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{AddFunc: func(obj interface{}) {
		handleAPIServiceChange(obj.(*apiregistration.APIService))
	}, UpdateFunc: func(old, new interface{}) {
		handleAPIServiceChange(new.(*apiregistration.APIService))
	}})
	return healthz.NamedCheck(name, func(r *http.Request) error {
		pendingServiceNamesLock.RLock()
		defer pendingServiceNamesLock.RUnlock()
		if pendingServiceNames.Len() > 0 {
			return fmt.Errorf("missing APIService: %v", pendingServiceNames.List())
		}
		return nil
	})
}

type priority struct {
	group   int32
	version int32
}

var apiVersionPriorities = map[schema.GroupVersion]priority{{Group: "", Version: "v1"}: {group: 18000, version: 1}, {Group: "extensions", Version: "v1beta1"}: {group: 17900, version: 1}, {Group: "apps", Version: "v1beta1"}: {group: 17800, version: 1}, {Group: "apps", Version: "v1beta2"}: {group: 17800, version: 9}, {Group: "apps", Version: "v1"}: {group: 17800, version: 15}, {Group: "events.k8s.io", Version: "v1beta1"}: {group: 17750, version: 5}, {Group: "authentication.k8s.io", Version: "v1"}: {group: 17700, version: 15}, {Group: "authentication.k8s.io", Version: "v1beta1"}: {group: 17700, version: 9}, {Group: "authorization.k8s.io", Version: "v1"}: {group: 17600, version: 15}, {Group: "authorization.k8s.io", Version: "v1beta1"}: {group: 17600, version: 9}, {Group: "autoscaling", Version: "v1"}: {group: 17500, version: 15}, {Group: "autoscaling", Version: "v2beta1"}: {group: 17500, version: 9}, {Group: "autoscaling", Version: "v2beta2"}: {group: 17500, version: 1}, {Group: "batch", Version: "v1"}: {group: 17400, version: 15}, {Group: "batch", Version: "v1beta1"}: {group: 17400, version: 9}, {Group: "batch", Version: "v2alpha1"}: {group: 17400, version: 9}, {Group: "certificates.k8s.io", Version: "v1beta1"}: {group: 17300, version: 9}, {Group: "networking.k8s.io", Version: "v1"}: {group: 17200, version: 15}, {Group: "policy", Version: "v1beta1"}: {group: 17100, version: 9}, {Group: "rbac.authorization.k8s.io", Version: "v1"}: {group: 17000, version: 15}, {Group: "rbac.authorization.k8s.io", Version: "v1beta1"}: {group: 17000, version: 12}, {Group: "rbac.authorization.k8s.io", Version: "v1alpha1"}: {group: 17000, version: 9}, {Group: "settings.k8s.io", Version: "v1alpha1"}: {group: 16900, version: 9}, {Group: "storage.k8s.io", Version: "v1"}: {group: 16800, version: 15}, {Group: "storage.k8s.io", Version: "v1beta1"}: {group: 16800, version: 9}, {Group: "storage.k8s.io", Version: "v1alpha1"}: {group: 16800, version: 1}, {Group: "apiextensions.k8s.io", Version: "v1beta1"}: {group: 16700, version: 9}, {Group: "admissionregistration.k8s.io", Version: "v1"}: {group: 16700, version: 15}, {Group: "admissionregistration.k8s.io", Version: "v1beta1"}: {group: 16700, version: 12}, {Group: "admissionregistration.k8s.io", Version: "v1alpha1"}: {group: 16700, version: 9}, {Group: "scheduling.k8s.io", Version: "v1beta1"}: {group: 16600, version: 12}, {Group: "scheduling.k8s.io", Version: "v1alpha1"}: {group: 16600, version: 9}, {Group: "coordination.k8s.io", Version: "v1beta1"}: {group: 16500, version: 9}, {Group: "auditregistration.k8s.io", Version: "v1alpha1"}: {group: 16400, version: 1}}

func apiServicesToRegister(delegateAPIServer genericapiserver.DelegationTarget, registration autoregister.AutoAPIServiceRegistration) []*apiregistration.APIService {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	apiServices := []*apiregistration.APIService{}
	for _, curr := range delegateAPIServer.ListedPaths() {
		if curr == "/api/v1" {
			apiService := makeAPIService(schema.GroupVersion{Group: "", Version: "v1"})
			registration.AddAPIServiceToSyncOnStart(apiService)
			apiServices = append(apiServices, apiService)
			continue
		}
		if !strings.HasPrefix(curr, "/apis/") {
			continue
		}
		tokens := strings.Split(curr, "/")
		if len(tokens) != 4 {
			continue
		}
		apiService := makeAPIService(schema.GroupVersion{Group: tokens[2], Version: tokens[3]})
		if apiService == nil {
			continue
		}
		registration.AddAPIServiceToSyncOnStart(apiService)
		apiServices = append(apiServices, apiService)
	}
	return apiServices
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
