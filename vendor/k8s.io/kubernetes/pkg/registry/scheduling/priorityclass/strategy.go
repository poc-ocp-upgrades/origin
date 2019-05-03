package priorityclass

import (
 "context"
 "k8s.io/apimachinery/pkg/runtime"
 "k8s.io/apimachinery/pkg/util/validation/field"
 "k8s.io/apiserver/pkg/storage/names"
 "k8s.io/kubernetes/pkg/api/legacyscheme"
 "k8s.io/kubernetes/pkg/apis/scheduling"
 "k8s.io/kubernetes/pkg/apis/scheduling/validation"
)

type priorityClassStrategy struct {
 runtime.ObjectTyper
 names.NameGenerator
}

var Strategy = priorityClassStrategy{legacyscheme.Scheme, names.SimpleNameGenerator}

func (priorityClassStrategy) NamespaceScoped() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return false
}
func (priorityClassStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 pc := obj.(*scheduling.PriorityClass)
 pc.Generation = 1
}
func (priorityClassStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 _ = obj.(*scheduling.PriorityClass)
 _ = old.(*scheduling.PriorityClass)
}
func (priorityClassStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 pc := obj.(*scheduling.PriorityClass)
 return validation.ValidatePriorityClass(pc)
}
func (priorityClassStrategy) Canonicalize(obj runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
}
func (priorityClassStrategy) AllowCreateOnUpdate() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return false
}
func (priorityClassStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return validation.ValidatePriorityClassUpdate(obj.(*scheduling.PriorityClass), old.(*scheduling.PriorityClass))
}
func (priorityClassStrategy) AllowUnconditionalUpdate() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return true
}
