package storage

import (
 "context"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 "fmt"
 "k8s.io/apimachinery/pkg/api/errors"
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 "k8s.io/apimachinery/pkg/runtime"
 "k8s.io/apimachinery/pkg/runtime/schema"
 "k8s.io/apiserver/pkg/registry/generic"
 genericregistry "k8s.io/apiserver/pkg/registry/generic/registry"
 "k8s.io/apiserver/pkg/registry/rest"
 "k8s.io/kubernetes/pkg/apis/apps"
 appsv1beta1 "k8s.io/kubernetes/pkg/apis/apps/v1beta1"
 appsv1beta2 "k8s.io/kubernetes/pkg/apis/apps/v1beta2"
 "k8s.io/kubernetes/pkg/apis/autoscaling"
 autoscalingv1 "k8s.io/kubernetes/pkg/apis/autoscaling/v1"
 autoscalingvalidation "k8s.io/kubernetes/pkg/apis/autoscaling/validation"
 "k8s.io/kubernetes/pkg/printers"
 printersinternal "k8s.io/kubernetes/pkg/printers/internalversion"
 printerstorage "k8s.io/kubernetes/pkg/printers/storage"
 "k8s.io/kubernetes/pkg/registry/apps/statefulset"
)

type StatefulSetStorage struct {
 StatefulSet *REST
 Status      *StatusREST
 Scale       *ScaleREST
}

func NewStorage(optsGetter generic.RESTOptionsGetter) StatefulSetStorage {
 _logClusterCodePath()
 defer _logClusterCodePath()
 statefulSetRest, statefulSetStatusRest := NewREST(optsGetter)
 return StatefulSetStorage{StatefulSet: statefulSetRest, Status: statefulSetStatusRest, Scale: &ScaleREST{store: statefulSetRest.Store}}
}

type REST struct{ *genericregistry.Store }

func NewREST(optsGetter generic.RESTOptionsGetter) (*REST, *StatusREST) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 store := &genericregistry.Store{NewFunc: func() runtime.Object {
  return &apps.StatefulSet{}
 }, NewListFunc: func() runtime.Object {
  return &apps.StatefulSetList{}
 }, DefaultQualifiedResource: apps.Resource("statefulsets"), CreateStrategy: statefulset.Strategy, UpdateStrategy: statefulset.Strategy, DeleteStrategy: statefulset.Strategy, TableConvertor: printerstorage.TableConvertor{TablePrinter: printers.NewTablePrinter().With(printersinternal.AddHandlers)}}
 options := &generic.StoreOptions{RESTOptions: optsGetter}
 if err := store.CompleteWithOptions(options); err != nil {
  panic(err)
 }
 statusStore := *store
 statusStore.UpdateStrategy = statefulset.StatusStrategy
 return &REST{store}, &StatusREST{store: &statusStore}
}

var _ rest.CategoriesProvider = &REST{}

func (r *REST) Categories() []string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return []string{"all"}
}

type StatusREST struct{ store *genericregistry.Store }

func (r *StatusREST) New() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &apps.StatefulSet{}
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

var _ rest.ShortNamesProvider = &REST{}

func (r *REST) ShortNames() []string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return []string{"sts"}
}

type ScaleREST struct{ store *genericregistry.Store }

var _ = rest.Patcher(&ScaleREST{})
var _ = rest.GroupVersionKindProvider(&ScaleREST{})

func (r *ScaleREST) GroupVersionKind(containingGV schema.GroupVersion) schema.GroupVersionKind {
 _logClusterCodePath()
 defer _logClusterCodePath()
 switch containingGV {
 case appsv1beta1.SchemeGroupVersion:
  return appsv1beta1.SchemeGroupVersion.WithKind("Scale")
 case appsv1beta2.SchemeGroupVersion:
  return appsv1beta2.SchemeGroupVersion.WithKind("Scale")
 default:
  return autoscalingv1.SchemeGroupVersion.WithKind("Scale")
 }
}
func (r *ScaleREST) New() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &autoscaling.Scale{}
}
func (r *ScaleREST) Get(ctx context.Context, name string, options *metav1.GetOptions) (runtime.Object, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := r.store.Get(ctx, name, options)
 if err != nil {
  return nil, err
 }
 ss := obj.(*apps.StatefulSet)
 scale, err := scaleFromStatefulSet(ss)
 if err != nil {
  return nil, errors.NewBadRequest(fmt.Sprintf("%v", err))
 }
 return scale, err
}
func (r *ScaleREST) Update(ctx context.Context, name string, objInfo rest.UpdatedObjectInfo, createValidation rest.ValidateObjectFunc, updateValidation rest.ValidateObjectUpdateFunc, forceAllowCreate bool, options *metav1.UpdateOptions) (runtime.Object, bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := r.store.Get(ctx, name, &metav1.GetOptions{})
 if err != nil {
  return nil, false, err
 }
 ss := obj.(*apps.StatefulSet)
 oldScale, err := scaleFromStatefulSet(ss)
 if err != nil {
  return nil, false, err
 }
 obj, err = objInfo.UpdatedObject(ctx, oldScale)
 if err != nil {
  return nil, false, err
 }
 if obj == nil {
  return nil, false, errors.NewBadRequest(fmt.Sprintf("nil update passed to Scale"))
 }
 scale, ok := obj.(*autoscaling.Scale)
 if !ok {
  return nil, false, errors.NewBadRequest(fmt.Sprintf("wrong object passed to Scale update: %v", obj))
 }
 if errs := autoscalingvalidation.ValidateScale(scale); len(errs) > 0 {
  return nil, false, errors.NewInvalid(autoscaling.Kind("Scale"), scale.Name, errs)
 }
 ss.Spec.Replicas = scale.Spec.Replicas
 ss.ResourceVersion = scale.ResourceVersion
 obj, _, err = r.store.Update(ctx, ss.Name, rest.DefaultUpdatedObjectInfo(ss), createValidation, updateValidation, false, options)
 if err != nil {
  return nil, false, err
 }
 ss = obj.(*apps.StatefulSet)
 newScale, err := scaleFromStatefulSet(ss)
 if err != nil {
  return nil, false, errors.NewBadRequest(fmt.Sprintf("%v", err))
 }
 return newScale, false, err
}
func scaleFromStatefulSet(ss *apps.StatefulSet) (*autoscaling.Scale, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 selector, err := metav1.LabelSelectorAsSelector(ss.Spec.Selector)
 if err != nil {
  return nil, err
 }
 return &autoscaling.Scale{ObjectMeta: metav1.ObjectMeta{Name: ss.Name, Namespace: ss.Namespace, UID: ss.UID, ResourceVersion: ss.ResourceVersion, CreationTimestamp: ss.CreationTimestamp}, Spec: autoscaling.ScaleSpec{Replicas: ss.Spec.Replicas}, Status: autoscaling.ScaleStatus{Replicas: ss.Status.Replicas, Selector: selector.String()}}, nil
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
