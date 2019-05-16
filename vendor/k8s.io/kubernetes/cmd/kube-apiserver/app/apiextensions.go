package app

import (
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	apiextensionsapiserver "k8s.io/apiextensions-apiserver/pkg/apiserver"
	apiextensionsoptions "k8s.io/apiextensions-apiserver/pkg/cmd/server/options"
	"k8s.io/apiserver/pkg/admission"
	"k8s.io/apiserver/pkg/features"
	genericapiserver "k8s.io/apiserver/pkg/server"
	genericoptions "k8s.io/apiserver/pkg/server/options"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	"k8s.io/apiserver/pkg/util/webhook"
	kubeexternalinformers "k8s.io/client-go/informers"
	"k8s.io/kubernetes/cmd/kube-apiserver/app/options"
)

func createAPIExtensionsConfig(kubeAPIServerConfig genericapiserver.Config, externalInformers kubeexternalinformers.SharedInformerFactory, pluginInitializers []admission.PluginInitializer, commandOptions *options.ServerRunOptions, masterCount int, serviceResolver webhook.ServiceResolver, authResolverWrapper webhook.AuthenticationInfoResolverWrapper) (*apiextensionsapiserver.Config, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	genericConfig := kubeAPIServerConfig
	commandOptions.Admission.ApplyTo(&genericConfig, externalInformers, genericConfig.LoopbackClientConfig, apiextensionsapiserver.Scheme, pluginInitializers...)
	etcdOptions := *commandOptions.Etcd
	etcdOptions.StorageConfig.Paging = utilfeature.DefaultFeatureGate.Enabled(features.APIListChunking)
	etcdOptions.StorageConfig.Codec = apiextensionsapiserver.Codecs.LegacyCodec(v1beta1.SchemeGroupVersion)
	genericConfig.RESTOptionsGetter = &genericoptions.SimpleRestOptionsFactory{Options: etcdOptions}
	if err := commandOptions.APIEnablement.ApplyTo(&genericConfig, apiextensionsapiserver.DefaultAPIResourceConfigSource(), apiextensionsapiserver.Scheme); err != nil {
		return nil, err
	}
	apiextensionsConfig := &apiextensionsapiserver.Config{GenericConfig: &genericapiserver.RecommendedConfig{Config: genericConfig, SharedInformerFactory: externalInformers}, ExtraConfig: apiextensionsapiserver.ExtraConfig{CRDRESTOptionsGetter: apiextensionsoptions.NewCRDRESTOptionsGetter(etcdOptions), MasterCount: masterCount, AuthResolverWrapper: authResolverWrapper, ServiceResolver: serviceResolver}}
	return apiextensionsConfig, nil
}
func createAPIExtensionsServer(apiextensionsConfig *apiextensionsapiserver.Config, delegateAPIServer genericapiserver.DelegationTarget) (*apiextensionsapiserver.CustomResourceDefinitions, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return apiextensionsConfig.Complete().New(delegateAPIServer)
}
