package rolebinding

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
 _logClusterCodePath()
 defer _logClusterCodePath()
 return true
}
func (strategy) AllowCreateOnUpdate() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return true
}
func (strategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 _ = obj.(*rbac.RoleBinding)
}
func (strategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 newRoleBinding := obj.(*rbac.RoleBinding)
 oldRoleBinding := old.(*rbac.RoleBinding)
 _, _ = newRoleBinding, oldRoleBinding
}
func (strategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 roleBinding := obj.(*rbac.RoleBinding)
 return validation.ValidateRoleBinding(roleBinding)
}
func (strategy) Canonicalize(obj runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 _ = obj.(*rbac.RoleBinding)
}
func (strategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 newObj := obj.(*rbac.RoleBinding)
 errorList := validation.ValidateRoleBinding(newObj)
 return append(errorList, validation.ValidateRoleBindingUpdate(newObj, old.(*rbac.RoleBinding))...)
}
func (strategy) AllowUnconditionalUpdate() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return true
}
