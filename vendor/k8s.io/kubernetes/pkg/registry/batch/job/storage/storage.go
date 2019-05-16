package storage

import (
	"context"
	goformat "fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/generic"
	genericregistry "k8s.io/apiserver/pkg/registry/generic/registry"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/kubernetes/pkg/apis/batch"
	"k8s.io/kubernetes/pkg/printers"
	printersinternal "k8s.io/kubernetes/pkg/printers/internalversion"
	printerstorage "k8s.io/kubernetes/pkg/printers/storage"
	"k8s.io/kubernetes/pkg/registry/batch/job"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

type JobStorage struct {
	Job    *REST
	Status *StatusREST
}

func NewStorage(optsGetter generic.RESTOptionsGetter) JobStorage {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	jobRest, jobStatusRest := NewREST(optsGetter)
	return JobStorage{Job: jobRest, Status: jobStatusRest}
}

type REST struct{ *genericregistry.Store }

func NewREST(optsGetter generic.RESTOptionsGetter) (*REST, *StatusREST) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	store := &genericregistry.Store{NewFunc: func() runtime.Object {
		return &batch.Job{}
	}, NewListFunc: func() runtime.Object {
		return &batch.JobList{}
	}, PredicateFunc: job.MatchJob, DefaultQualifiedResource: batch.Resource("jobs"), CreateStrategy: job.Strategy, UpdateStrategy: job.Strategy, DeleteStrategy: job.Strategy, TableConvertor: printerstorage.TableConvertor{TablePrinter: printers.NewTablePrinter().With(printersinternal.AddHandlers)}}
	options := &generic.StoreOptions{RESTOptions: optsGetter, AttrFunc: job.GetAttrs}
	if err := store.CompleteWithOptions(options); err != nil {
		panic(err)
	}
	statusStore := *store
	statusStore.UpdateStrategy = job.StatusStrategy
	return &REST{store}, &StatusREST{store: &statusStore}
}

var _ rest.CategoriesProvider = &REST{}

func (r *REST) Categories() []string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return []string{"all"}
}

type StatusREST struct{ store *genericregistry.Store }

func (r *StatusREST) New() runtime.Object {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &batch.Job{}
}
func (r *StatusREST) Get(ctx context.Context, name string, options *metav1.GetOptions) (runtime.Object, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return r.store.Get(ctx, name, options)
}
func (r *StatusREST) Update(ctx context.Context, name string, objInfo rest.UpdatedObjectInfo, createValidation rest.ValidateObjectFunc, updateValidation rest.ValidateObjectUpdateFunc, forceAllowCreate bool, options *metav1.UpdateOptions) (runtime.Object, bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return r.store.Update(ctx, name, objInfo, createValidation, updateValidation, false, options)
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
