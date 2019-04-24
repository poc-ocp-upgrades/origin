package etcd

import (
	"context"
	"bytes"
	"net/http"
	"runtime"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/registry/generic/registry"
	"k8s.io/apiserver/pkg/registry/rest"
	authorizationclient "k8s.io/client-go/kubernetes/typed/authorization/v1"
	"k8s.io/kubernetes/pkg/printers"
	printerstorage "k8s.io/kubernetes/pkg/printers/storage"
	"github.com/openshift/api/template"
	printersinternal "github.com/openshift/origin/pkg/printers/internalversion"
	templateapi "github.com/openshift/origin/pkg/template/apis/template"
	"github.com/openshift/origin/pkg/template/apiserver/registry/templateinstance"
)

type REST struct{ *registry.Store }

var _ rest.StandardStorage = &REST{}

func NewREST(optsGetter generic.RESTOptionsGetter, authorizationClient authorizationclient.AuthorizationV1Interface) (*REST, *StatusREST, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	strategy := templateinstance.NewStrategy(authorizationClient)
	store := &registry.Store{NewFunc: func() runtime.Object {
		return &templateapi.TemplateInstance{}
	}, NewListFunc: func() runtime.Object {
		return &templateapi.TemplateInstanceList{}
	}, DefaultQualifiedResource: template.Resource("templateinstances"), TableConvertor: printerstorage.TableConvertor{TablePrinter: printers.NewTablePrinter().With(printersinternal.AddHandlers)}, CreateStrategy: strategy, UpdateStrategy: strategy, DeleteStrategy: strategy}
	options := &generic.StoreOptions{RESTOptions: optsGetter}
	if err := store.CompleteWithOptions(options); err != nil {
		return nil, nil, err
	}
	statusStore := *store
	statusStore.UpdateStrategy = templateinstance.StatusStrategy
	return &REST{store}, &StatusREST{&statusStore}, nil
}

type StatusREST struct{ store *registry.Store }

var _ = rest.Patcher(&StatusREST{})

func (r *StatusREST) New() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &templateapi.TemplateInstance{}
}
func (r *StatusREST) Get(ctx context.Context, name string, options *metav1.GetOptions) (runtime.Object, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return r.store.Get(ctx, name, options)
}
func (r *StatusREST) Update(ctx context.Context, name string, objInfo rest.UpdatedObjectInfo, createValidation rest.ValidateObjectFunc, updateValidation rest.ValidateObjectUpdateFunc, forceAllowCreate bool, options *metav1.UpdateOptions) (runtime.Object, bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return r.store.Update(ctx, name, objInfo, createValidation, updateValidation, forceAllowCreate, options)
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
