package admission

import (
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apiserver/pkg/admission"
	"k8s.io/apiserver/pkg/authorization/authorizer"
	"k8s.io/apiserver/pkg/util/webhook"
	quota "k8s.io/kubernetes/pkg/quota/v1"
)

type WantsCloudConfig interface{ SetCloudConfig([]byte) }
type WantsRESTMapper interface{ SetRESTMapper(meta.RESTMapper) }
type WantsQuotaConfiguration interface {
	SetQuotaConfiguration(quota.Configuration)
	admission.InitializationValidator
}
type PluginInitializer struct {
	authorizer                        authorizer.Authorizer
	cloudConfig                       []byte
	restMapper                        meta.RESTMapper
	quotaConfiguration                quota.Configuration
	serviceResolver                   webhook.ServiceResolver
	authenticationInfoResolverWrapper webhook.AuthenticationInfoResolverWrapper
}

var _ admission.PluginInitializer = &PluginInitializer{}

func NewPluginInitializer(cloudConfig []byte, restMapper meta.RESTMapper, quotaConfiguration quota.Configuration) *PluginInitializer {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &PluginInitializer{cloudConfig: cloudConfig, restMapper: restMapper, quotaConfiguration: quotaConfiguration}
}
func (i *PluginInitializer) Initialize(plugin admission.Interface) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if wants, ok := plugin.(WantsCloudConfig); ok {
		wants.SetCloudConfig(i.cloudConfig)
	}
	if wants, ok := plugin.(WantsRESTMapper); ok {
		wants.SetRESTMapper(i.restMapper)
	}
	if wants, ok := plugin.(WantsQuotaConfiguration); ok {
		wants.SetQuotaConfiguration(i.quotaConfiguration)
	}
}
