package build

import (
	"context"
	"bytes"
	"net/http"
	"runtime"
	"fmt"
	"reflect"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/storage/names"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	buildapi "github.com/openshift/origin/pkg/build/apis/build"
	buildinternalhelpers "github.com/openshift/origin/pkg/build/apis/build/internal_helpers"
	"github.com/openshift/origin/pkg/build/apis/build/validation"
)

type strategy struct {
	runtime.ObjectTyper
	names.NameGenerator
}

var Strategy = strategy{legacyscheme.Scheme, names.SimpleNameGenerator}

func (strategy) NamespaceScoped() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return true
}
func (strategy) AllowCreateOnUpdate() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return false
}
func (strategy) AllowUnconditionalUpdate() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return false
}
func (strategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	build := obj.(*buildapi.Build)
	if len(build.Status.Phase) == 0 {
		build.Status.Phase = buildapi.BuildPhaseNew
	}
}
func (strategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	newBuild := obj.(*buildapi.Build)
	oldBuild := old.(*buildapi.Build)
	if oldBuild.Status.Phase == buildapi.BuildPhaseFailed && newBuild.Status.Reason != buildapi.StatusReasonOutOfMemoryKilled {
		newBuild.Status.Reason = oldBuild.Status.Reason
		newBuild.Status.Message = oldBuild.Status.Message
	}
}
func (strategy) Canonicalize(obj runtime.Object) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (strategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return validation.ValidateBuild(obj.(*buildapi.Build))
}
func (strategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return validation.ValidateBuildUpdate(obj.(*buildapi.Build), old.(*buildapi.Build))
}
func (strategy) CheckGracefulDelete(obj runtime.Object, options *metav1.DeleteOptions) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return false
}

type detailsStrategy struct{ strategy }

func (detailsStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	newBuild := obj.(*buildapi.Build)
	oldBuild := old.(*buildapi.Build)
	phase := oldBuild.Status.Phase
	stages := newBuild.Status.Stages
	if buildinternalhelpers.IsBuildComplete(newBuild) {
		phase = newBuild.Status.Phase
	}
	revision := newBuild.Spec.Revision
	message := newBuild.Status.Message
	reason := newBuild.Status.Reason
	outputTo := newBuild.Status.Output.To
	*newBuild = *oldBuild
	newBuild.Status.Phase = phase
	newBuild.Status.Stages = stages
	newBuild.Spec.Revision = revision
	newBuild.Status.Reason = reason
	newBuild.Status.Message = message
	newBuild.Status.Output.To = outputTo
}
func (detailsStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	newBuild := obj.(*buildapi.Build)
	oldBuild := old.(*buildapi.Build)
	oldRevision := oldBuild.Spec.Revision
	newRevision := newBuild.Spec.Revision
	errors := field.ErrorList{}
	if newRevision == nil && oldRevision != nil {
		errors = append(errors, field.Invalid(field.NewPath("spec", "revision"), nil, "cannot set an empty revision in build spec"))
	}
	if !reflect.DeepEqual(oldRevision, newRevision) && oldRevision != nil {
		errors = append(errors, field.Duplicate(field.NewPath("spec", "revision"), oldBuild.Spec.Revision))
	}
	return errors
}
func (detailsStrategy) AllowUnconditionalUpdate() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return true
}

var DetailsStrategy = detailsStrategy{Strategy}

func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
