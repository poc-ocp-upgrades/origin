package statefulset

import (
	"context"
	goformat "fmt"
	appsv1beta1 "k8s.io/api/apps/v1beta1"
	appsv1beta2 "k8s.io/api/apps/v1beta2"
	apiequality "k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/validation/field"
	genericapirequest "k8s.io/apiserver/pkg/endpoints/request"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/apiserver/pkg/storage/names"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	"k8s.io/kubernetes/pkg/api/pod"
	"k8s.io/kubernetes/pkg/apis/apps"
	"k8s.io/kubernetes/pkg/apis/apps/validation"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

type statefulSetStrategy struct {
	runtime.ObjectTyper
	names.NameGenerator
}

var Strategy = statefulSetStrategy{legacyscheme.Scheme, names.SimpleNameGenerator}

func (statefulSetStrategy) DefaultGarbageCollectionPolicy(ctx context.Context) rest.GarbageCollectionPolicy {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if requestInfo, found := genericapirequest.RequestInfoFrom(ctx); found {
		groupVersion := schema.GroupVersion{Group: requestInfo.APIGroup, Version: requestInfo.APIVersion}
		switch groupVersion {
		case appsv1beta1.SchemeGroupVersion, appsv1beta2.SchemeGroupVersion:
			return rest.OrphanDependents
		default:
			return rest.DeleteDependents
		}
	}
	return rest.OrphanDependents
}
func (statefulSetStrategy) NamespaceScoped() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return true
}
func (statefulSetStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	statefulSet := obj.(*apps.StatefulSet)
	statefulSet.Status = apps.StatefulSetStatus{}
	statefulSet.Generation = 1
	pod.DropDisabledAlphaFields(&statefulSet.Spec.Template.Spec)
}
func (statefulSetStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	newStatefulSet := obj.(*apps.StatefulSet)
	oldStatefulSet := old.(*apps.StatefulSet)
	newStatefulSet.Status = oldStatefulSet.Status
	pod.DropDisabledAlphaFields(&newStatefulSet.Spec.Template.Spec)
	pod.DropDisabledAlphaFields(&oldStatefulSet.Spec.Template.Spec)
	if !apiequality.Semantic.DeepEqual(oldStatefulSet.Spec, newStatefulSet.Spec) {
		newStatefulSet.Generation = oldStatefulSet.Generation + 1
	}
}
func (statefulSetStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	statefulSet := obj.(*apps.StatefulSet)
	return validation.ValidateStatefulSet(statefulSet)
}
func (statefulSetStrategy) Canonicalize(obj runtime.Object) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
}
func (statefulSetStrategy) AllowCreateOnUpdate() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return false
}
func (statefulSetStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	validationErrorList := validation.ValidateStatefulSet(obj.(*apps.StatefulSet))
	updateErrorList := validation.ValidateStatefulSetUpdate(obj.(*apps.StatefulSet), old.(*apps.StatefulSet))
	return append(validationErrorList, updateErrorList...)
}
func (statefulSetStrategy) AllowUnconditionalUpdate() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return true
}

type statefulSetStatusStrategy struct{ statefulSetStrategy }

var StatusStrategy = statefulSetStatusStrategy{Strategy}

func (statefulSetStatusStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	newStatefulSet := obj.(*apps.StatefulSet)
	oldStatefulSet := old.(*apps.StatefulSet)
	newStatefulSet.Spec = oldStatefulSet.Spec
}
func (statefulSetStatusStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return validation.ValidateStatefulSetStatusUpdate(obj.(*apps.StatefulSet), old.(*apps.StatefulSet))
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
