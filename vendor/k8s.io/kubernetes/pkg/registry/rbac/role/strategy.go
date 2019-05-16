package role

import (
	"context"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/apiserver/pkg/storage/names"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	"k8s.io/kubernetes/pkg/apis/rbac"
	"k8s.io/kubernetes/pkg/apis/rbac/validation"
)

type strategy struct {
	runtime.ObjectTyper
	names.NameGenerator
}

var Strategy = strategy{legacyscheme.Scheme, names.SimpleNameGenerator}
var _ rest.RESTCreateStrategy = Strategy
var _ rest.RESTUpdateStrategy = Strategy

func (strategy) NamespaceScoped() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return true
}
func (strategy) AllowCreateOnUpdate() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return true
}
func (strategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	_ = obj.(*rbac.Role)
}
func (strategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	newRole := obj.(*rbac.Role)
	oldRole := old.(*rbac.Role)
	_, _ = newRole, oldRole
}
func (strategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	role := obj.(*rbac.Role)
	return validation.ValidateRole(role)
}
func (strategy) Canonicalize(obj runtime.Object) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	_ = obj.(*rbac.Role)
}
func (strategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	newObj := obj.(*rbac.Role)
	errorList := validation.ValidateRole(newObj)
	return append(errorList, validation.ValidateRoleUpdate(newObj, old.(*rbac.Role))...)
}
func (strategy) AllowUnconditionalUpdate() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return true
}
