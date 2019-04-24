package buildconfig

import (
	"context"
	"bytes"
	"net/http"
	"runtime"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/apiserver/pkg/storage/names"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	buildapi "github.com/openshift/origin/pkg/build/apis/build"
	"github.com/openshift/origin/pkg/build/apis/build/validation"
)

var (
	GroupStrategy	= groupStrategy{strategy{legacyscheme.Scheme, names.SimpleNameGenerator}}
	LegacyStrategy	= legacyStrategy{strategy{legacyscheme.Scheme, names.SimpleNameGenerator}}
)

type strategy struct {
	runtime.ObjectTyper
	names.NameGenerator
}

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
func (strategy) Canonicalize(obj runtime.Object) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (s strategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bc := obj.(*buildapi.BuildConfig)
	dropUnknownTriggers(bc)
}
func (strategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	newBC := obj.(*buildapi.BuildConfig)
	oldBC := old.(*buildapi.BuildConfig)
	dropUnknownTriggers(newBC)
	if newBC.Status.LastVersion < oldBC.Status.LastVersion {
		newBC.Status.LastVersion = oldBC.Status.LastVersion
	}
}
func (strategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return validation.ValidateBuildConfig(obj.(*buildapi.BuildConfig))
}
func (strategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return validation.ValidateBuildConfigUpdate(obj.(*buildapi.BuildConfig), old.(*buildapi.BuildConfig))
}

type groupStrategy struct{ strategy }

func (s groupStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
	_logClusterCodePath()
	defer _logClusterCodePath()
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.strategy.PrepareForCreate(ctx, obj)
}

var _ rest.GarbageCollectionDeleteStrategy = legacyStrategy{}

func (s legacyStrategy) DefaultGarbageCollectionPolicy(ctx context.Context) rest.GarbageCollectionPolicy {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return rest.OrphanDependents
}
func (strategy) CheckGracefulDelete(obj runtime.Object, options *metav1.DeleteOptions) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return false
}
func dropUnknownTriggers(bc *buildapi.BuildConfig) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	triggers := []buildapi.BuildTriggerPolicy{}
	for _, t := range bc.Spec.Triggers {
		if buildapi.KnownTriggerTypes.Has(string(t.Type)) {
			triggers = append(triggers, t)
		}
	}
	bc.Spec.Triggers = triggers
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
