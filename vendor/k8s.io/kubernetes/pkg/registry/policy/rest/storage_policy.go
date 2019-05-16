package rest

import (
	goformat "fmt"
	policyapiv1beta1 "k8s.io/api/policy/v1beta1"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/registry/rest"
	genericapiserver "k8s.io/apiserver/pkg/server"
	serverstorage "k8s.io/apiserver/pkg/server/storage"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	"k8s.io/kubernetes/pkg/apis/policy"
	poddisruptionbudgetstore "k8s.io/kubernetes/pkg/registry/policy/poddisruptionbudget/storage"
	pspstore "k8s.io/kubernetes/pkg/registry/policy/podsecuritypolicy/storage"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

type RESTStorageProvider struct{}

func (p RESTStorageProvider) NewRESTStorage(apiResourceConfigSource serverstorage.APIResourceConfigSource, restOptionsGetter generic.RESTOptionsGetter) (genericapiserver.APIGroupInfo, bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	apiGroupInfo := genericapiserver.NewDefaultAPIGroupInfo(policy.GroupName, legacyscheme.Scheme, legacyscheme.ParameterCodec, legacyscheme.Codecs)
	if apiResourceConfigSource.VersionEnabled(policyapiv1beta1.SchemeGroupVersion) {
		apiGroupInfo.VersionedResourcesStorageMap[policyapiv1beta1.SchemeGroupVersion.Version] = p.v1beta1Storage(apiResourceConfigSource, restOptionsGetter)
	}
	return apiGroupInfo, true
}
func (p RESTStorageProvider) v1beta1Storage(apiResourceConfigSource serverstorage.APIResourceConfigSource, restOptionsGetter generic.RESTOptionsGetter) map[string]rest.Storage {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	storage := map[string]rest.Storage{}
	poddisruptionbudgetStorage, poddisruptionbudgetStatusStorage := poddisruptionbudgetstore.NewREST(restOptionsGetter)
	storage["poddisruptionbudgets"] = poddisruptionbudgetStorage
	storage["poddisruptionbudgets/status"] = poddisruptionbudgetStatusStorage
	storage["podsecuritypolicies"] = pspstore.NewREST(restOptionsGetter)
	return storage
}
func (p RESTStorageProvider) GroupName() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return policy.GroupName
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
