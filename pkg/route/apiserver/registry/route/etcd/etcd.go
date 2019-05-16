package etcd

import (
	"context"
	goformat "fmt"
	routegroup "github.com/openshift/api/route"
	printersinternal "github.com/openshift/origin/pkg/printers/internalversion"
	"github.com/openshift/origin/pkg/route"
	routeapi "github.com/openshift/origin/pkg/route/apis/route"
	routeregistry "github.com/openshift/origin/pkg/route/apiserver/registry/route"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/registry/generic/registry"
	"k8s.io/apiserver/pkg/registry/rest"
	kapirest "k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/apiserver/pkg/storage"
	"k8s.io/kubernetes/pkg/printers"
	printerstorage "k8s.io/kubernetes/pkg/printers/storage"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

type REST struct{ *registry.Store }

var _ rest.StandardStorage = &REST{}
var _ rest.CategoriesProvider = &REST{}

func (r *REST) Categories() []string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return []string{"all"}
}
func NewREST(optsGetter generic.RESTOptionsGetter, allocator route.RouteAllocator, sarClient routeregistry.SubjectAccessReviewInterface) (*REST, *StatusREST, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	strategy := routeregistry.NewStrategy(allocator, sarClient)
	store := &registry.Store{NewFunc: func() runtime.Object {
		return &routeapi.Route{}
	}, NewListFunc: func() runtime.Object {
		return &routeapi.RouteList{}
	}, DefaultQualifiedResource: routegroup.Resource("routes"), TableConvertor: printerstorage.TableConvertor{TablePrinter: printers.NewTablePrinter().With(printersinternal.AddHandlers)}, CreateStrategy: strategy, UpdateStrategy: strategy, DeleteStrategy: strategy}
	options := &generic.StoreOptions{RESTOptions: optsGetter, AttrFunc: storage.AttrFunc(storage.DefaultNamespaceScopedAttr).WithFieldMutation(routeapi.RouteFieldSelector)}
	if err := store.CompleteWithOptions(options); err != nil {
		return nil, nil, err
	}
	statusStore := *store
	statusStore.UpdateStrategy = routeregistry.StatusStrategy
	return &REST{store}, &StatusREST{&statusStore}, nil
}

type StatusREST struct{ store *registry.Store }

var _ = kapirest.Patcher(&StatusREST{})

func (r *StatusREST) New() runtime.Object {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &routeapi.Route{}
}
func (r *StatusREST) Get(ctx context.Context, name string, options *metav1.GetOptions) (runtime.Object, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return r.store.Get(ctx, name, options)
}
func (r *StatusREST) Update(ctx context.Context, name string, objInfo kapirest.UpdatedObjectInfo, createValidation rest.ValidateObjectFunc, updateValidation rest.ValidateObjectUpdateFunc, forceAllowCreate bool, options *metav1.UpdateOptions) (runtime.Object, bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return r.store.Update(ctx, name, objInfo, createValidation, updateValidation, forceAllowCreate, options)
}

type LegacyREST struct{ *REST }

func (r *LegacyREST) Categories() []string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return []string{}
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
