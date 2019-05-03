package podsecuritypolicy

import (
 "context"
 "k8s.io/apimachinery/pkg/runtime"
 "k8s.io/apimachinery/pkg/util/validation/field"
 "k8s.io/apiserver/pkg/registry/rest"
 "k8s.io/apiserver/pkg/storage/names"
 "k8s.io/kubernetes/pkg/api/legacyscheme"
 psputil "k8s.io/kubernetes/pkg/api/podsecuritypolicy"
 "k8s.io/kubernetes/pkg/apis/policy"
 "k8s.io/kubernetes/pkg/apis/policy/validation"
)

type strategy struct {
 runtime.ObjectTyper
 names.NameGenerator
}

var Strategy = strategy{legacyscheme.Scheme, names.SimpleNameGenerator}
var _ = rest.RESTCreateStrategy(Strategy)
var _ = rest.RESTUpdateStrategy(Strategy)

func (strategy) NamespaceScoped() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return false
}
func (strategy) AllowCreateOnUpdate() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return false
}
func (strategy) AllowUnconditionalUpdate() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return true
}
func (strategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 psp := obj.(*policy.PodSecurityPolicy)
 psputil.DropDisabledAlphaFields(&psp.Spec)
}
func (strategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 newPsp := obj.(*policy.PodSecurityPolicy)
 oldPsp := old.(*policy.PodSecurityPolicy)
 psputil.DropDisabledAlphaFields(&newPsp.Spec)
 psputil.DropDisabledAlphaFields(&oldPsp.Spec)
}
func (strategy) Canonicalize(obj runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
}
func (strategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return validation.ValidatePodSecurityPolicy(obj.(*policy.PodSecurityPolicy))
}
func (strategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return validation.ValidatePodSecurityPolicyUpdate(old.(*policy.PodSecurityPolicy), obj.(*policy.PodSecurityPolicy))
}
