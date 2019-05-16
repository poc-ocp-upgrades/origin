package horizontalpodautoscaler

import (
	"context"
	goformat "fmt"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/storage/names"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	"k8s.io/kubernetes/pkg/apis/autoscaling"
	"k8s.io/kubernetes/pkg/apis/autoscaling/validation"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

type autoscalerStrategy struct {
	runtime.ObjectTyper
	names.NameGenerator
}

var Strategy = autoscalerStrategy{legacyscheme.Scheme, names.SimpleNameGenerator}

func (autoscalerStrategy) NamespaceScoped() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return true
}
func (autoscalerStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	newHPA := obj.(*autoscaling.HorizontalPodAutoscaler)
	newHPA.Status = autoscaling.HorizontalPodAutoscalerStatus{}
}
func (autoscalerStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	autoscaler := obj.(*autoscaling.HorizontalPodAutoscaler)
	return validation.ValidateHorizontalPodAutoscaler(autoscaler)
}
func (autoscalerStrategy) Canonicalize(obj runtime.Object) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
}
func (autoscalerStrategy) AllowCreateOnUpdate() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return false
}
func (autoscalerStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	newHPA := obj.(*autoscaling.HorizontalPodAutoscaler)
	oldHPA := old.(*autoscaling.HorizontalPodAutoscaler)
	newHPA.Status = oldHPA.Status
}
func (autoscalerStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return validation.ValidateHorizontalPodAutoscalerUpdate(obj.(*autoscaling.HorizontalPodAutoscaler), old.(*autoscaling.HorizontalPodAutoscaler))
}
func (autoscalerStrategy) AllowUnconditionalUpdate() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return true
}

type autoscalerStatusStrategy struct{ autoscalerStrategy }

var StatusStrategy = autoscalerStatusStrategy{Strategy}

func (autoscalerStatusStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	newAutoscaler := obj.(*autoscaling.HorizontalPodAutoscaler)
	oldAutoscaler := old.(*autoscaling.HorizontalPodAutoscaler)
	newAutoscaler.Spec = oldAutoscaler.Spec
}
func (autoscalerStatusStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return validation.ValidateHorizontalPodAutoscalerStatusUpdate(obj.(*autoscaling.HorizontalPodAutoscaler), old.(*autoscaling.HorizontalPodAutoscaler))
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
