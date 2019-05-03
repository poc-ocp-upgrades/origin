package lease

import (
 "context"
 "k8s.io/apimachinery/pkg/runtime"
 "k8s.io/apimachinery/pkg/util/validation/field"
 "k8s.io/apiserver/pkg/storage/names"
 "k8s.io/kubernetes/pkg/api/legacyscheme"
 "k8s.io/kubernetes/pkg/apis/coordination"
 "k8s.io/kubernetes/pkg/apis/coordination/validation"
)

type leaseStrategy struct {
 runtime.ObjectTyper
 names.NameGenerator
}

var Strategy = leaseStrategy{legacyscheme.Scheme, names.SimpleNameGenerator}

func (leaseStrategy) NamespaceScoped() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return true
}
func (leaseStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
}
func (leaseStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
}
func (leaseStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 lease := obj.(*coordination.Lease)
 return validation.ValidateLease(lease)
}
func (leaseStrategy) Canonicalize(obj runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
}
func (leaseStrategy) AllowCreateOnUpdate() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return true
}
func (leaseStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return validation.ValidateLeaseUpdate(obj.(*coordination.Lease), old.(*coordination.Lease))
}
func (leaseStrategy) AllowUnconditionalUpdate() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return false
}
