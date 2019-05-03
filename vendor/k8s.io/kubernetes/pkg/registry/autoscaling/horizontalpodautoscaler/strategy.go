package horizontalpodautoscaler

import (
 "context"
 "k8s.io/apimachinery/pkg/runtime"
 "k8s.io/apimachinery/pkg/util/validation/field"
 "k8s.io/apiserver/pkg/storage/names"
 "k8s.io/kubernetes/pkg/api/legacyscheme"
 "k8s.io/kubernetes/pkg/apis/autoscaling"
 "k8s.io/kubernetes/pkg/apis/autoscaling/validation"
)

type autoscalerStrategy struct {
 runtime.ObjectTyper
 names.NameGenerator
}

var Strategy = autoscalerStrategy{legacyscheme.Scheme, names.SimpleNameGenerator}

func (autoscalerStrategy) NamespaceScoped() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return true
}
func (autoscalerStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 newHPA := obj.(*autoscaling.HorizontalPodAutoscaler)
 newHPA.Status = autoscaling.HorizontalPodAutoscalerStatus{}
}
func (autoscalerStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 autoscaler := obj.(*autoscaling.HorizontalPodAutoscaler)
 return validation.ValidateHorizontalPodAutoscaler(autoscaler)
}
func (autoscalerStrategy) Canonicalize(obj runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
}
func (autoscalerStrategy) AllowCreateOnUpdate() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return false
}
func (autoscalerStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 newHPA := obj.(*autoscaling.HorizontalPodAutoscaler)
 oldHPA := old.(*autoscaling.HorizontalPodAutoscaler)
 newHPA.Status = oldHPA.Status
}
func (autoscalerStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return validation.ValidateHorizontalPodAutoscalerUpdate(obj.(*autoscaling.HorizontalPodAutoscaler), old.(*autoscaling.HorizontalPodAutoscaler))
}
func (autoscalerStrategy) AllowUnconditionalUpdate() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return true
}

type autoscalerStatusStrategy struct{ autoscalerStrategy }

var StatusStrategy = autoscalerStatusStrategy{Strategy}

func (autoscalerStatusStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 newAutoscaler := obj.(*autoscaling.HorizontalPodAutoscaler)
 oldAutoscaler := old.(*autoscaling.HorizontalPodAutoscaler)
 newAutoscaler.Spec = oldAutoscaler.Spec
}
func (autoscalerStatusStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return validation.ValidateHorizontalPodAutoscalerStatusUpdate(obj.(*autoscaling.HorizontalPodAutoscaler), old.(*autoscaling.HorizontalPodAutoscaler))
}
