package rangeallocations

import (
	"context"
	securityapi "github.com/openshift/origin/pkg/security/apis/security"
	apimachineryvalidation "k8s.io/apimachinery/pkg/api/validation"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/apiserver/pkg/storage/names"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	"k8s.io/kubernetes/pkg/apis/core/validation"
)

type strategy struct {
	runtime.ObjectTyper
	names.NameGenerator
}

var strategyInstance = strategy{legacyscheme.Scheme, names.SimpleNameGenerator}
var _ rest.RESTCreateStrategy = strategyInstance
var _ rest.RESTUpdateStrategy = strategyInstance

func (strategy) NamespaceScoped() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return false
}
func (strategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_ = obj.(*securityapi.RangeAllocation)
}
func (strategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cfg := obj.(*securityapi.RangeAllocation)
	return validation.ValidateObjectMeta(&cfg.ObjectMeta, false, apimachineryvalidation.NameIsDNSSubdomain, field.NewPath("metadata"))
}
func (strategy) Canonicalize(obj runtime.Object) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (strategy) AllowCreateOnUpdate() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return false
}
func (strategy) PrepareForUpdate(ctx context.Context, newObj, oldObj runtime.Object) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_ = oldObj.(*securityapi.RangeAllocation)
	_ = newObj.(*securityapi.RangeAllocation)
}
func (strategy) AllowUnconditionalUpdate() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return false
}
func (strategy) ValidateUpdate(ctx context.Context, newObj, oldObj runtime.Object) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	oldCfg, newCfg := oldObj.(*securityapi.RangeAllocation), newObj.(*securityapi.RangeAllocation)
	return validation.ValidateObjectMetaUpdate(&newCfg.ObjectMeta, &oldCfg.ObjectMeta, field.NewPath("metadata"))
}
