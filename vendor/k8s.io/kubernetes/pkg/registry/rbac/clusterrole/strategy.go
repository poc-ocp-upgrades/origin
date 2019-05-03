package clusterrole

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
 return false
}
func (strategy) AllowCreateOnUpdate() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return true
}
func (strategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 _ = obj.(*rbac.ClusterRole)
}
func (strategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 newClusterRole := obj.(*rbac.ClusterRole)
 oldClusterRole := old.(*rbac.ClusterRole)
 _, _ = newClusterRole, oldClusterRole
}
func (strategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 clusterRole := obj.(*rbac.ClusterRole)
 return validation.ValidateClusterRole(clusterRole)
}
func (strategy) Canonicalize(obj runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 _ = obj.(*rbac.ClusterRole)
}
func (strategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 newObj := obj.(*rbac.ClusterRole)
 errorList := validation.ValidateClusterRole(newObj)
 return append(errorList, validation.ValidateClusterRoleUpdate(newObj, old.(*rbac.ClusterRole))...)
}
func (strategy) AllowUnconditionalUpdate() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return true
}
