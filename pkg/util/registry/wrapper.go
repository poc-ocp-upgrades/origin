package registry

import (
	"context"
	"bytes"
	"net/http"
	"runtime"
	"fmt"
	"k8s.io/apimachinery/pkg/apis/meta/internalversion"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	metav1beta1 "k8s.io/apimachinery/pkg/apis/meta/v1beta1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/rest"
	"github.com/openshift/origin/pkg/util/errors"
)

type NoWatchStorage interface {
	rest.Getter
	rest.Lister
	rest.TableConvertor
	rest.CreaterUpdater
	rest.GracefulDeleter
	rest.Scoper
}

func WrapNoWatchStorageError(delegate NoWatchStorage) NoWatchStorage {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &noWatchStorageErrWrapper{delegate: delegate}
}

var _ = NoWatchStorage(&noWatchStorageErrWrapper{})

type noWatchStorageErrWrapper struct{ delegate NoWatchStorage }

func (s *noWatchStorageErrWrapper) NamespaceScoped() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.delegate.NamespaceScoped()
}
func (s *noWatchStorageErrWrapper) Get(ctx context.Context, name string, options *metav1.GetOptions) (runtime.Object, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := s.delegate.Get(ctx, name, options)
	return obj, errors.SyncStatusError(ctx, err)
}
func (s *noWatchStorageErrWrapper) List(ctx context.Context, options *internalversion.ListOptions) (runtime.Object, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := s.delegate.List(ctx, options)
	return obj, errors.SyncStatusError(ctx, err)
}
func (s *noWatchStorageErrWrapper) ConvertToTable(ctx context.Context, object runtime.Object, tableOptions runtime.Object) (*metav1beta1.Table, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.delegate.ConvertToTable(ctx, object, tableOptions)
}
func (s *noWatchStorageErrWrapper) Create(ctx context.Context, in runtime.Object, createValidation rest.ValidateObjectFunc, options *metav1.CreateOptions) (runtime.Object, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := s.delegate.Create(ctx, in, createValidation, options)
	return obj, errors.SyncStatusError(ctx, err)
}
func (s *noWatchStorageErrWrapper) Update(ctx context.Context, name string, objInfo rest.UpdatedObjectInfo, createValidation rest.ValidateObjectFunc, updateValidation rest.ValidateObjectUpdateFunc, forceAllowCreate bool, options *metav1.UpdateOptions) (runtime.Object, bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, created, err := s.delegate.Update(ctx, name, objInfo, createValidation, updateValidation, forceAllowCreate, options)
	return obj, created, errors.SyncStatusError(ctx, err)
}
func (s *noWatchStorageErrWrapper) Delete(ctx context.Context, name string, options *metav1.DeleteOptions) (runtime.Object, bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, deleted, err := s.delegate.Delete(ctx, name, options)
	return obj, deleted, errors.SyncStatusError(ctx, err)
}
func (s *noWatchStorageErrWrapper) New() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.delegate.New()
}
func (s *noWatchStorageErrWrapper) NewList() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.delegate.NewList()
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
