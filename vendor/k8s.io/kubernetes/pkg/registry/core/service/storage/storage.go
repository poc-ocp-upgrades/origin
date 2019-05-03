package storage

import (
 "context"
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 "k8s.io/apimachinery/pkg/runtime"
 "k8s.io/apiserver/pkg/registry/generic"
 genericregistry "k8s.io/apiserver/pkg/registry/generic/registry"
 "k8s.io/apiserver/pkg/registry/rest"
 api "k8s.io/kubernetes/pkg/apis/core"
 "k8s.io/kubernetes/pkg/printers"
 printersinternal "k8s.io/kubernetes/pkg/printers/internalversion"
 printerstorage "k8s.io/kubernetes/pkg/printers/storage"
 "k8s.io/kubernetes/pkg/registry/core/service"
)

type GenericREST struct{ *genericregistry.Store }

func NewGenericREST(optsGetter generic.RESTOptionsGetter) (*GenericREST, *StatusREST) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 store := &genericregistry.Store{NewFunc: func() runtime.Object {
  return &api.Service{}
 }, NewListFunc: func() runtime.Object {
  return &api.ServiceList{}
 }, DefaultQualifiedResource: api.Resource("services"), ReturnDeletedObject: true, CreateStrategy: service.Strategy, UpdateStrategy: service.Strategy, DeleteStrategy: service.Strategy, ExportStrategy: service.Strategy, TableConvertor: printerstorage.TableConvertor{TablePrinter: printers.NewTablePrinter().With(printersinternal.AddHandlers)}}
 options := &generic.StoreOptions{RESTOptions: optsGetter}
 if err := store.CompleteWithOptions(options); err != nil {
  panic(err)
 }
 statusStore := *store
 statusStore.UpdateStrategy = service.StatusStrategy
 return &GenericREST{store}, &StatusREST{store: &statusStore}
}

var (
 _ rest.ShortNamesProvider = &GenericREST{}
 _ rest.CategoriesProvider = &GenericREST{}
)

func (r *GenericREST) ShortNames() []string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return []string{"svc"}
}
func (r *GenericREST) Categories() []string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return []string{"all"}
}

type StatusREST struct{ store *genericregistry.Store }

func (r *StatusREST) New() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &api.Service{}
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
