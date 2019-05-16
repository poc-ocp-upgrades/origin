package replicationcontroller

import (
	"context"
	"fmt"
	goformat "fmt"
	apiequality "k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/registry/rest"
	apistorage "k8s.io/apiserver/pkg/storage"
	"k8s.io/apiserver/pkg/storage/names"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	"k8s.io/kubernetes/pkg/api/pod"
	api "k8s.io/kubernetes/pkg/apis/core"
	"k8s.io/kubernetes/pkg/apis/core/helper"
	"k8s.io/kubernetes/pkg/apis/core/validation"
	goos "os"
	godefaultruntime "runtime"
	"strconv"
	"strings"
	gotime "time"
)

type rcStrategy struct {
	runtime.ObjectTyper
	names.NameGenerator
}

var Strategy = rcStrategy{legacyscheme.Scheme, names.SimpleNameGenerator}

func (rcStrategy) DefaultGarbageCollectionPolicy(ctx context.Context) rest.GarbageCollectionPolicy {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return rest.OrphanDependents
}
func (rcStrategy) NamespaceScoped() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return true
}
func (rcStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	controller := obj.(*api.ReplicationController)
	controller.Status = api.ReplicationControllerStatus{}
	controller.Generation = 1
	if controller.Spec.Template != nil {
		pod.DropDisabledAlphaFields(&controller.Spec.Template.Spec)
	}
}
func (rcStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	newController := obj.(*api.ReplicationController)
	oldController := old.(*api.ReplicationController)
	newController.Status = oldController.Status
	if oldController.Spec.Template != nil {
		pod.DropDisabledAlphaFields(&oldController.Spec.Template.Spec)
	}
	if newController.Spec.Template != nil {
		pod.DropDisabledAlphaFields(&newController.Spec.Template.Spec)
	}
	if !apiequality.Semantic.DeepEqual(oldController.Spec, newController.Spec) {
		newController.Generation = oldController.Generation + 1
	}
}
func (rcStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	controller := obj.(*api.ReplicationController)
	return validation.ValidateReplicationController(controller)
}
func (rcStrategy) Canonicalize(obj runtime.Object) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
}
func (rcStrategy) AllowCreateOnUpdate() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return false
}
func (rcStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	oldRc := old.(*api.ReplicationController)
	newRc := obj.(*api.ReplicationController)
	validationErrorList := validation.ValidateReplicationController(newRc)
	updateErrorList := validation.ValidateReplicationControllerUpdate(newRc, oldRc)
	errs := append(validationErrorList, updateErrorList...)
	for key, value := range helper.NonConvertibleFields(oldRc.Annotations) {
		parts := strings.Split(key, "/")
		if len(parts) != 2 {
			continue
		}
		brokenField := parts[1]
		switch {
		case strings.Contains(brokenField, "selector"):
			if !apiequality.Semantic.DeepEqual(oldRc.Spec.Selector, newRc.Spec.Selector) {
				errs = append(errs, field.Invalid(field.NewPath("spec").Child("selector"), newRc.Spec.Selector, "cannot update non-convertible selector"))
			}
		default:
			errs = append(errs, &field.Error{Type: field.ErrorTypeNotFound, BadValue: value, Field: brokenField, Detail: "unknown non-convertible field"})
		}
	}
	return errs
}
func (rcStrategy) AllowUnconditionalUpdate() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return true
}
func ControllerToSelectableFields(controller *api.ReplicationController) fields.Set {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	objectMetaFieldsSet := generic.ObjectMetaFieldsSet(&controller.ObjectMeta, true)
	controllerSpecificFieldsSet := fields.Set{"status.replicas": strconv.Itoa(int(controller.Status.Replicas))}
	return generic.MergeFieldsSets(objectMetaFieldsSet, controllerSpecificFieldsSet)
}
func GetAttrs(obj runtime.Object) (labels.Set, fields.Set, bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	rc, ok := obj.(*api.ReplicationController)
	if !ok {
		return nil, nil, false, fmt.Errorf("given object is not a replication controller.")
	}
	return labels.Set(rc.ObjectMeta.Labels), ControllerToSelectableFields(rc), rc.Initializers != nil, nil
}
func MatchController(label labels.Selector, field fields.Selector) apistorage.SelectionPredicate {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return apistorage.SelectionPredicate{Label: label, Field: field, GetAttrs: GetAttrs}
}

type rcStatusStrategy struct{ rcStrategy }

var StatusStrategy = rcStatusStrategy{Strategy}

func (rcStatusStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	newRc := obj.(*api.ReplicationController)
	oldRc := old.(*api.ReplicationController)
	newRc.Spec = oldRc.Spec
}
func (rcStatusStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return validation.ValidateReplicationControllerStatusUpdate(obj.(*api.ReplicationController), old.(*api.ReplicationController))
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
