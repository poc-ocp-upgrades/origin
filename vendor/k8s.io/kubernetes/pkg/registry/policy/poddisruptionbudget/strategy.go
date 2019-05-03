package poddisruptionbudget

import (
 "context"
 apiequality "k8s.io/apimachinery/pkg/api/equality"
 "k8s.io/apimachinery/pkg/runtime"
 "k8s.io/apimachinery/pkg/util/validation/field"
 "k8s.io/apiserver/pkg/storage/names"
 "k8s.io/kubernetes/pkg/api/legacyscheme"
 "k8s.io/kubernetes/pkg/apis/policy"
 "k8s.io/kubernetes/pkg/apis/policy/validation"
)

type podDisruptionBudgetStrategy struct {
 runtime.ObjectTyper
 names.NameGenerator
}

var Strategy = podDisruptionBudgetStrategy{legacyscheme.Scheme, names.SimpleNameGenerator}

func (podDisruptionBudgetStrategy) NamespaceScoped() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return true
}
func (podDisruptionBudgetStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 podDisruptionBudget := obj.(*policy.PodDisruptionBudget)
 podDisruptionBudget.Status = policy.PodDisruptionBudgetStatus{}
 podDisruptionBudget.Generation = 1
}
func (podDisruptionBudgetStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 newPodDisruptionBudget := obj.(*policy.PodDisruptionBudget)
 oldPodDisruptionBudget := old.(*policy.PodDisruptionBudget)
 newPodDisruptionBudget.Status = oldPodDisruptionBudget.Status
 if !apiequality.Semantic.DeepEqual(oldPodDisruptionBudget.Spec, newPodDisruptionBudget.Spec) {
  newPodDisruptionBudget.Generation = oldPodDisruptionBudget.Generation + 1
 }
}
func (podDisruptionBudgetStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 podDisruptionBudget := obj.(*policy.PodDisruptionBudget)
 return validation.ValidatePodDisruptionBudget(podDisruptionBudget)
}
func (podDisruptionBudgetStrategy) Canonicalize(obj runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
}
func (podDisruptionBudgetStrategy) AllowCreateOnUpdate() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return false
}
func (podDisruptionBudgetStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 validationErrorList := validation.ValidatePodDisruptionBudget(obj.(*policy.PodDisruptionBudget))
 updateErrorList := validation.ValidatePodDisruptionBudgetUpdate(obj.(*policy.PodDisruptionBudget), old.(*policy.PodDisruptionBudget))
 return append(validationErrorList, updateErrorList...)
}
func (podDisruptionBudgetStrategy) AllowUnconditionalUpdate() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return false
}

type podDisruptionBudgetStatusStrategy struct{ podDisruptionBudgetStrategy }

var StatusStrategy = podDisruptionBudgetStatusStrategy{Strategy}

func (podDisruptionBudgetStatusStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 newPodDisruptionBudget := obj.(*policy.PodDisruptionBudget)
 oldPodDisruptionBudget := old.(*policy.PodDisruptionBudget)
 newPodDisruptionBudget.Spec = oldPodDisruptionBudget.Spec
}
func (podDisruptionBudgetStatusStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return field.ErrorList{}
}
