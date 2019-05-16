package poddisruptionbudget

import (
	"context"
	goformat "fmt"
	apiequality "k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/storage/names"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	"k8s.io/kubernetes/pkg/apis/policy"
	"k8s.io/kubernetes/pkg/apis/policy/validation"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

type podDisruptionBudgetStrategy struct {
	runtime.ObjectTyper
	names.NameGenerator
}

var Strategy = podDisruptionBudgetStrategy{legacyscheme.Scheme, names.SimpleNameGenerator}

func (podDisruptionBudgetStrategy) NamespaceScoped() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return true
}
func (podDisruptionBudgetStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	podDisruptionBudget := obj.(*policy.PodDisruptionBudget)
	podDisruptionBudget.Status = policy.PodDisruptionBudgetStatus{}
	podDisruptionBudget.Generation = 1
}
func (podDisruptionBudgetStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	newPodDisruptionBudget := obj.(*policy.PodDisruptionBudget)
	oldPodDisruptionBudget := old.(*policy.PodDisruptionBudget)
	newPodDisruptionBudget.Status = oldPodDisruptionBudget.Status
	if !apiequality.Semantic.DeepEqual(oldPodDisruptionBudget.Spec, newPodDisruptionBudget.Spec) {
		newPodDisruptionBudget.Generation = oldPodDisruptionBudget.Generation + 1
	}
}
func (podDisruptionBudgetStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	podDisruptionBudget := obj.(*policy.PodDisruptionBudget)
	return validation.ValidatePodDisruptionBudget(podDisruptionBudget)
}
func (podDisruptionBudgetStrategy) Canonicalize(obj runtime.Object) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
}
func (podDisruptionBudgetStrategy) AllowCreateOnUpdate() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return false
}
func (podDisruptionBudgetStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	validationErrorList := validation.ValidatePodDisruptionBudget(obj.(*policy.PodDisruptionBudget))
	updateErrorList := validation.ValidatePodDisruptionBudgetUpdate(obj.(*policy.PodDisruptionBudget), old.(*policy.PodDisruptionBudget))
	return append(validationErrorList, updateErrorList...)
}
func (podDisruptionBudgetStrategy) AllowUnconditionalUpdate() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return false
}

type podDisruptionBudgetStatusStrategy struct{ podDisruptionBudgetStrategy }

var StatusStrategy = podDisruptionBudgetStatusStrategy{Strategy}

func (podDisruptionBudgetStatusStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	newPodDisruptionBudget := obj.(*policy.PodDisruptionBudget)
	oldPodDisruptionBudget := old.(*policy.PodDisruptionBudget)
	newPodDisruptionBudget.Spec = oldPodDisruptionBudget.Spec
}
func (podDisruptionBudgetStatusStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return field.ErrorList{}
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
