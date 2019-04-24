package rolebindingrestriction

import (
	"context"
	"bytes"
	"net/http"
	"runtime"
	"fmt"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/apiserver/pkg/storage/names"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	authorizationapi "github.com/openshift/origin/pkg/authorization/apis/authorization"
	"github.com/openshift/origin/pkg/authorization/apis/authorization/validation"
)

type strategy struct {
	runtime.ObjectTyper
	names.NameGenerator
}

var Strategy = strategy{legacyscheme.Scheme, names.SimpleNameGenerator}
var _ rest.GarbageCollectionDeleteStrategy = strategy{}

func (strategy) DefaultGarbageCollectionPolicy(ctx context.Context) rest.GarbageCollectionPolicy {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return rest.Unsupported
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
func (strategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_ = obj.(*authorizationapi.RoleBindingRestriction)
}
func (strategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_ = obj.(*authorizationapi.RoleBindingRestriction)
	_ = old.(*authorizationapi.RoleBindingRestriction)
}
func (strategy) Canonicalize(obj runtime.Object) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (strategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return validation.ValidateRoleBindingRestriction(obj.(*authorizationapi.RoleBindingRestriction))
}
func (strategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return validation.ValidateRoleBindingRestrictionUpdate(obj.(*authorizationapi.RoleBindingRestriction), old.(*authorizationapi.RoleBindingRestriction))
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
