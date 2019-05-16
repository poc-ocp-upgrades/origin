package buildconfig

import (
	"context"
	goformat "fmt"
	buildapi "github.com/openshift/origin/pkg/build/apis/build"
	"github.com/openshift/origin/pkg/build/apis/build/validation"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/apiserver/pkg/storage/names"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

var (
	GroupStrategy  = groupStrategy{strategy{legacyscheme.Scheme, names.SimpleNameGenerator}}
	LegacyStrategy = legacyStrategy{strategy{legacyscheme.Scheme, names.SimpleNameGenerator}}
)

type strategy struct {
	runtime.ObjectTyper
	names.NameGenerator
}

func (strategy) NamespaceScoped() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return true
}
func (strategy) AllowCreateOnUpdate() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return false
}
func (strategy) AllowUnconditionalUpdate() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return false
}
func (strategy) Canonicalize(obj runtime.Object) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
}
func (s strategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	bc := obj.(*buildapi.BuildConfig)
	dropUnknownTriggers(bc)
}
func (strategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	newBC := obj.(*buildapi.BuildConfig)
	oldBC := old.(*buildapi.BuildConfig)
	dropUnknownTriggers(newBC)
	if newBC.Status.LastVersion < oldBC.Status.LastVersion {
		newBC.Status.LastVersion = oldBC.Status.LastVersion
	}
}
func (strategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return validation.ValidateBuildConfig(obj.(*buildapi.BuildConfig))
}
func (strategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return validation.ValidateBuildConfigUpdate(obj.(*buildapi.BuildConfig), old.(*buildapi.BuildConfig))
}

type groupStrategy struct{ strategy }

func (s groupStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	s.strategy.PrepareForCreate(ctx, obj)
	bc := obj.(*buildapi.BuildConfig)
	if bc.Spec.SuccessfulBuildsHistoryLimit == nil {
		v := buildapi.DefaultSuccessfulBuildsHistoryLimit
		bc.Spec.SuccessfulBuildsHistoryLimit = &v
	}
	if bc.Spec.FailedBuildsHistoryLimit == nil {
		v := buildapi.DefaultFailedBuildsHistoryLimit
		bc.Spec.FailedBuildsHistoryLimit = &v
	}
}

type legacyStrategy struct{ strategy }

func (s legacyStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	s.strategy.PrepareForCreate(ctx, obj)
}

var _ rest.GarbageCollectionDeleteStrategy = legacyStrategy{}

func (s legacyStrategy) DefaultGarbageCollectionPolicy(ctx context.Context) rest.GarbageCollectionPolicy {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return rest.OrphanDependents
}
func (strategy) CheckGracefulDelete(obj runtime.Object, options *metav1.DeleteOptions) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return false
}
func dropUnknownTriggers(bc *buildapi.BuildConfig) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	triggers := []buildapi.BuildTriggerPolicy{}
	for _, t := range bc.Spec.Triggers {
		if buildapi.KnownTriggerTypes.Has(string(t.Type)) {
			triggers = append(triggers, t)
		}
	}
	bc.Spec.Triggers = triggers
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
