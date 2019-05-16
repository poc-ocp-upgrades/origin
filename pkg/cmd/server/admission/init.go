package admission

import (
	goformat "fmt"
	quotainformer "github.com/openshift/client-go/quota/informers/externalversions/quota/v1"
	securityv1informer "github.com/openshift/client-go/security/informers/externalversions/security/v1"
	userinformer "github.com/openshift/client-go/user/informers/externalversions"
	"github.com/openshift/origin/pkg/image/apiserver/registryhostname"
	"github.com/openshift/origin/pkg/project/cache"
	"github.com/openshift/origin/pkg/quota/controller/clusterquotamapping"
	"k8s.io/apiserver/pkg/admission"
	restclient "k8s.io/client-go/rest"
	quota "k8s.io/kubernetes/pkg/quota/v1"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

type PluginInitializer struct {
	ProjectCache                 *cache.ProjectCache
	DefaultNodeSelector          string
	OriginQuotaRegistry          quota.Registry
	RESTClientConfig             restclient.Config
	ClusterResourceQuotaInformer quotainformer.ClusterResourceQuotaInformer
	ClusterQuotaMapper           clusterquotamapping.ClusterQuotaMapper
	RegistryHostnameRetriever    registryhostname.RegistryHostnameRetriever
	SecurityInformers            securityv1informer.SecurityContextConstraintsInformer
	UserInformers                userinformer.SharedInformerFactory
}

func (i *PluginInitializer) Initialize(plugin admission.Interface) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if wantsProjectCache, ok := plugin.(WantsProjectCache); ok {
		wantsProjectCache.SetProjectCache(i.ProjectCache)
	}
	if castObj, ok := plugin.(WantsDefaultNodeSelector); ok {
		castObj.SetDefaultNodeSelector(i.DefaultNodeSelector)
	}
	if wantsOriginQuotaRegistry, ok := plugin.(WantsOriginQuotaRegistry); ok {
		wantsOriginQuotaRegistry.SetOriginQuotaRegistry(i.OriginQuotaRegistry)
	}
	if wantsRESTClientConfig, ok := plugin.(WantsRESTClientConfig); ok {
		wantsRESTClientConfig.SetRESTClientConfig(i.RESTClientConfig)
	}
	if wantsClusterQuota, ok := plugin.(WantsClusterQuota); ok {
		wantsClusterQuota.SetClusterQuota(i.ClusterQuotaMapper, i.ClusterResourceQuotaInformer)
	}
	if wantsSecurityInformer, ok := plugin.(WantsSecurityInformer); ok {
		wantsSecurityInformer.SetSecurityInformers(i.SecurityInformers)
	}
	if wantsDefaultRegistryFunc, ok := plugin.(WantsDefaultRegistryFunc); ok {
		wantsDefaultRegistryFunc.SetDefaultRegistryFunc(i.RegistryHostnameRetriever.InternalRegistryHostname)
	}
	if wantsUserInformer, ok := plugin.(WantsUserInformer); ok {
		wantsUserInformer.SetUserInformer(i.UserInformers)
	}
}
func Validate(plugins []admission.Interface) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for _, plugin := range plugins {
		if validater, ok := plugin.(admission.InitializationValidator); ok {
			err := validater.ValidateInitialization()
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
