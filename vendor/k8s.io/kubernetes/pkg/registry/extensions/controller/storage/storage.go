package storage

import (
	"context"
	"fmt"
	goformat "fmt"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/generic"
	genericregistry "k8s.io/apiserver/pkg/registry/generic/registry"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/kubernetes/pkg/apis/autoscaling"
	autoscalingvalidation "k8s.io/kubernetes/pkg/apis/autoscaling/validation"
	api "k8s.io/kubernetes/pkg/apis/core"
	"k8s.io/kubernetes/pkg/apis/extensions"
	controllerstore "k8s.io/kubernetes/pkg/registry/core/replicationcontroller/storage"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

type ContainerStorage struct {
	ReplicationController *RcREST
	Scale                 *ScaleREST
}

func NewStorage(optsGetter generic.RESTOptionsGetter) ContainerStorage {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	controllerREST, _ := controllerstore.NewREST(optsGetter)
	return ContainerStorage{ReplicationController: &RcREST{}, Scale: &ScaleREST{store: controllerREST.Store}}
}

type ScaleREST struct{ store *genericregistry.Store }

var _ = rest.Patcher(&ScaleREST{})

func (r *ScaleREST) New() runtime.Object {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &autoscaling.Scale{}
}
func (r *ScaleREST) Get(ctx context.Context, name string, options *metav1.GetOptions) (runtime.Object, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	obj, err := r.store.Get(ctx, name, options)
	if err != nil {
		return nil, errors.NewNotFound(extensions.Resource("replicationcontrollers/scale"), name)
	}
	rc := obj.(*api.ReplicationController)
	return scaleFromRC(rc), nil
}
func (r *ScaleREST) Update(ctx context.Context, name string, objInfo rest.UpdatedObjectInfo, createValidation rest.ValidateObjectFunc, updateValidation rest.ValidateObjectUpdateFunc, forceAllowCreate bool, options *metav1.UpdateOptions) (runtime.Object, bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	obj, err := r.store.Get(ctx, name, &metav1.GetOptions{})
	if err != nil {
		return nil, false, errors.NewNotFound(extensions.Resource("replicationcontrollers/scale"), name)
	}
	rc := obj.(*api.ReplicationController)
	oldScale := scaleFromRC(rc)
	obj, err = objInfo.UpdatedObject(ctx, oldScale)
	if obj == nil {
		return nil, false, errors.NewBadRequest(fmt.Sprintf("nil update passed to Scale"))
	}
	scale, ok := obj.(*autoscaling.Scale)
	if !ok {
		return nil, false, errors.NewBadRequest(fmt.Sprintf("wrong object passed to Scale update: %v", obj))
	}
	if errs := autoscalingvalidation.ValidateScale(scale); len(errs) > 0 {
		return nil, false, errors.NewInvalid(extensions.Kind("Scale"), scale.Name, errs)
	}
	rc.Spec.Replicas = scale.Spec.Replicas
	rc.ResourceVersion = scale.ResourceVersion
	obj, _, err = r.store.Update(ctx, rc.Name, rest.DefaultUpdatedObjectInfo(rc), createValidation, updateValidation, false, options)
	if err != nil {
		return nil, false, errors.NewConflict(extensions.Resource("replicationcontrollers/scale"), scale.Name, err)
	}
	rc = obj.(*api.ReplicationController)
	return scaleFromRC(rc), false, nil
}
func scaleFromRC(rc *api.ReplicationController) *autoscaling.Scale {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &autoscaling.Scale{ObjectMeta: metav1.ObjectMeta{Name: rc.Name, Namespace: rc.Namespace, UID: rc.UID, ResourceVersion: rc.ResourceVersion, CreationTimestamp: rc.CreationTimestamp}, Spec: autoscaling.ScaleSpec{Replicas: rc.Spec.Replicas}, Status: autoscaling.ScaleStatus{Replicas: rc.Status.Replicas, Selector: labels.SelectorFromSet(labels.Set(rc.Spec.Selector)).String()}}
}

type RcREST struct{}

func (r *RcREST) NamespaceScoped() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return true
}
func (r *RcREST) New() runtime.Object {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &extensions.ReplicationControllerDummy{}
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
