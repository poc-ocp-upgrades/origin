package rest

import (
	goformat "fmt"
	auditv1alpha1 "k8s.io/api/auditregistration/v1alpha1"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/registry/rest"
	genericapiserver "k8s.io/apiserver/pkg/server"
	serverstorage "k8s.io/apiserver/pkg/server/storage"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	auditstorage "k8s.io/kubernetes/pkg/registry/auditregistration/auditsink/storage"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

type RESTStorageProvider struct{}

func (p RESTStorageProvider) NewRESTStorage(apiResourceConfigSource serverstorage.APIResourceConfigSource, restOptionsGetter generic.RESTOptionsGetter) (genericapiserver.APIGroupInfo, bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	apiGroupInfo := genericapiserver.NewDefaultAPIGroupInfo(auditv1alpha1.GroupName, legacyscheme.Scheme, legacyscheme.ParameterCodec, legacyscheme.Codecs)
	if apiResourceConfigSource.VersionEnabled(auditv1alpha1.SchemeGroupVersion) {
		apiGroupInfo.VersionedResourcesStorageMap[auditv1alpha1.SchemeGroupVersion.Version] = p.v1alpha1Storage(apiResourceConfigSource, restOptionsGetter)
	}
	return apiGroupInfo, true
}
func (p RESTStorageProvider) v1alpha1Storage(apiResourceConfigSource serverstorage.APIResourceConfigSource, restOptionsGetter generic.RESTOptionsGetter) map[string]rest.Storage {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	storage := map[string]rest.Storage{}
	s := auditstorage.NewREST(restOptionsGetter)
	storage["auditsinks"] = s
	return storage
}
func (p RESTStorageProvider) GroupName() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return auditv1alpha1.GroupName
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
