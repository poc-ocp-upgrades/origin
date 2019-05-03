package admission

import (
	quotainformer "github.com/openshift/client-go/quota/informers/externalversions/quota/v1"
	securityv1informer "github.com/openshift/client-go/security/informers/externalversions"
	userinformer "github.com/openshift/client-go/user/informers/externalversions"
	"github.com/openshift/origin/pkg/project/cache"
	"github.com/openshift/origin/pkg/quota/controller/clusterquotamapping"
	"k8s.io/apiserver/pkg/admission"
	restclient "k8s.io/client-go/rest"
	quota "k8s.io/kubernetes/pkg/quota/v1"
)

type WantsProjectCache interface {
	SetProjectCache(*cache.ProjectCache)
	admission.InitializationValidator
}
type WantsDefaultNodeSelector interface {
	SetDefaultNodeSelector(string)
	admission.InitializationValidator
}
type WantsOriginQuotaRegistry interface {
	SetOriginQuotaRegistry(quota.Registry)
	admission.InitializationValidator
}
type WantsRESTClientConfig interface {
	SetRESTClientConfig(restclient.Config)
	admission.InitializationValidator
}
type WantsClusterQuota interface {
	SetClusterQuota(clusterquotamapping.ClusterQuotaMapper, quotainformer.ClusterResourceQuotaInformer)
	admission.InitializationValidator
}
type WantsSecurityInformer interface {
	SetSecurityInformers(securityv1informer.SharedInformerFactory)
	admission.InitializationValidator
}
type WantsDefaultRegistryFunc interface {
	SetDefaultRegistryFunc(func() (string, bool))
	admission.InitializationValidator
}
type WantsUserInformer interface {
	SetUserInformer(userinformer.SharedInformerFactory)
	admission.InitializationValidator
}
