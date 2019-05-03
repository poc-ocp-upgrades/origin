package networkpolicy

import (
 "context"
 "reflect"
 "k8s.io/apimachinery/pkg/runtime"
 "k8s.io/apimachinery/pkg/util/validation/field"
 "k8s.io/apiserver/pkg/storage/names"
 "k8s.io/kubernetes/pkg/api/legacyscheme"
 "k8s.io/kubernetes/pkg/apis/networking"
 "k8s.io/kubernetes/pkg/apis/networking/validation"
)

type networkPolicyStrategy struct {
 runtime.ObjectTyper
 names.NameGenerator
}

var Strategy = networkPolicyStrategy{legacyscheme.Scheme, names.SimpleNameGenerator}

func (networkPolicyStrategy) NamespaceScoped() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return true
}
func (networkPolicyStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 networkPolicy := obj.(*networking.NetworkPolicy)
 networkPolicy.Generation = 1
}
func (networkPolicyStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 newNetworkPolicy := obj.(*networking.NetworkPolicy)
 oldNetworkPolicy := old.(*networking.NetworkPolicy)
 if !reflect.DeepEqual(oldNetworkPolicy.Spec, newNetworkPolicy.Spec) {
  newNetworkPolicy.Generation = oldNetworkPolicy.Generation + 1
 }
}
func (networkPolicyStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 networkPolicy := obj.(*networking.NetworkPolicy)
 return validation.ValidateNetworkPolicy(networkPolicy)
}
func (networkPolicyStrategy) Canonicalize(obj runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
}
func (networkPolicyStrategy) AllowCreateOnUpdate() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return false
}
func (networkPolicyStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 validationErrorList := validation.ValidateNetworkPolicy(obj.(*networking.NetworkPolicy))
 updateErrorList := validation.ValidateNetworkPolicyUpdate(obj.(*networking.NetworkPolicy), old.(*networking.NetworkPolicy))
 return append(validationErrorList, updateErrorList...)
}
func (networkPolicyStrategy) AllowUnconditionalUpdate() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return true
}
