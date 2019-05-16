package deployconfig

import (
	"context"
	goformat "fmt"
	appsapi "github.com/openshift/origin/pkg/apps/apis/apps"
	"github.com/openshift/origin/pkg/apps/apis/apps/validation"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/apiserver/pkg/storage/names"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	goos "os"
	"reflect"
	godefaultruntime "runtime"
	gotime "time"
)

type strategy struct {
	runtime.ObjectTyper
	names.NameGenerator
}

var CommonStrategy = strategy{legacyscheme.Scheme, names.SimpleNameGenerator}
var LegacyStrategy = legacyStrategy{CommonStrategy}
var GroupStrategy = groupStrategy{CommonStrategy}

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
func (s strategy) Export(ctx context.Context, obj runtime.Object, exact bool) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	s.PrepareForCreate(ctx, obj)
	return nil
}
func (strategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	dc := obj.(*appsapi.DeploymentConfig)
	dc.Generation = 1
	dc.Status = appsapi.DeploymentConfigStatus{}
	for i := range dc.Spec.Triggers {
		if params := dc.Spec.Triggers[i].ImageChangeParams; params != nil {
			params.LastTriggeredImage = ""
		}
	}
}
func (strategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	newDc := obj.(*appsapi.DeploymentConfig)
	oldDc := old.(*appsapi.DeploymentConfig)
	newVersion := newDc.Status.LatestVersion
	oldVersion := oldDc.Status.LatestVersion
	newDc.Status = oldDc.Status
	if newVersion == oldVersion+1 {
		newDc.Status.LatestVersion = newVersion
	}
	if !reflect.DeepEqual(oldDc.Spec, newDc.Spec) || newDc.Status.LatestVersion != oldDc.Status.LatestVersion {
		newDc.Generation = oldDc.Generation + 1
	}
}
func (strategy) Canonicalize(obj runtime.Object) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
}
func (strategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return validation.ValidateDeploymentConfig(obj.(*appsapi.DeploymentConfig))
}
func (strategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return validation.ValidateDeploymentConfigUpdate(obj.(*appsapi.DeploymentConfig), old.(*appsapi.DeploymentConfig))
}
func (strategy) CheckGracefulDelete(obj runtime.Object, options *metav1.DeleteOptions) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return false
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

type groupStrategy struct{ strategy }

func (s groupStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	s.strategy.PrepareForCreate(ctx, obj)
	dc := obj.(*appsapi.DeploymentConfig)
	appsV1DeploymentConfigLayeredDefaults(dc)
}

type statusStrategy struct{ strategy }

var StatusStrategy = statusStrategy{CommonStrategy}

func (statusStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	newDc := obj.(*appsapi.DeploymentConfig)
	oldDc := old.(*appsapi.DeploymentConfig)
	newDc.Spec = oldDc.Spec
	newDc.Labels = oldDc.Labels
}
func (statusStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return validation.ValidateDeploymentConfigStatusUpdate(obj.(*appsapi.DeploymentConfig), old.(*appsapi.DeploymentConfig))
}
func appsV1DeploymentConfigLayeredDefaults(dc *appsapi.DeploymentConfig) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if dc.Spec.RevisionHistoryLimit == nil {
		v := appsapi.DefaultRevisionHistoryLimit
		dc.Spec.RevisionHistoryLimit = &v
	}
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
