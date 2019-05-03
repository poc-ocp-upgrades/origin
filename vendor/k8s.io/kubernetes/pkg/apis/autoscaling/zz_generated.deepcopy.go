package autoscaling

import (
 v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 runtime "k8s.io/apimachinery/pkg/runtime"
)

func (in *CrossVersionObjectReference) DeepCopyInto(out *CrossVersionObjectReference) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 return
}
func (in *CrossVersionObjectReference) DeepCopy() *CrossVersionObjectReference {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(CrossVersionObjectReference)
 in.DeepCopyInto(out)
 return out
}
func (in *ExternalMetricSource) DeepCopyInto(out *ExternalMetricSource) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 in.Metric.DeepCopyInto(&out.Metric)
 in.Target.DeepCopyInto(&out.Target)
 return
}
func (in *ExternalMetricSource) DeepCopy() *ExternalMetricSource {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ExternalMetricSource)
 in.DeepCopyInto(out)
 return out
}
func (in *ExternalMetricStatus) DeepCopyInto(out *ExternalMetricStatus) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 in.Metric.DeepCopyInto(&out.Metric)
 in.Current.DeepCopyInto(&out.Current)
 return
}
func (in *ExternalMetricStatus) DeepCopy() *ExternalMetricStatus {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ExternalMetricStatus)
 in.DeepCopyInto(out)
 return out
}
func (in *HorizontalPodAutoscaler) DeepCopyInto(out *HorizontalPodAutoscaler) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
 in.Spec.DeepCopyInto(&out.Spec)
 in.Status.DeepCopyInto(&out.Status)
 return
}
func (in *HorizontalPodAutoscaler) DeepCopy() *HorizontalPodAutoscaler {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(HorizontalPodAutoscaler)
 in.DeepCopyInto(out)
 return out
}
func (in *HorizontalPodAutoscaler) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *HorizontalPodAutoscalerCondition) DeepCopyInto(out *HorizontalPodAutoscalerCondition) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
 return
}
func (in *HorizontalPodAutoscalerCondition) DeepCopy() *HorizontalPodAutoscalerCondition {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(HorizontalPodAutoscalerCondition)
 in.DeepCopyInto(out)
 return out
}
func (in *HorizontalPodAutoscalerList) DeepCopyInto(out *HorizontalPodAutoscalerList) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]HorizontalPodAutoscaler, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 return
}
func (in *HorizontalPodAutoscalerList) DeepCopy() *HorizontalPodAutoscalerList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(HorizontalPodAutoscalerList)
 in.DeepCopyInto(out)
 return out
}
func (in *HorizontalPodAutoscalerList) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *HorizontalPodAutoscalerSpec) DeepCopyInto(out *HorizontalPodAutoscalerSpec) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.ScaleTargetRef = in.ScaleTargetRef
 if in.MinReplicas != nil {
  in, out := &in.MinReplicas, &out.MinReplicas
  *out = new(int32)
  **out = **in
 }
 if in.Metrics != nil {
  in, out := &in.Metrics, &out.Metrics
  *out = make([]MetricSpec, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 return
}
func (in *HorizontalPodAutoscalerSpec) DeepCopy() *HorizontalPodAutoscalerSpec {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(HorizontalPodAutoscalerSpec)
 in.DeepCopyInto(out)
 return out
}
func (in *HorizontalPodAutoscalerStatus) DeepCopyInto(out *HorizontalPodAutoscalerStatus) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.ObservedGeneration != nil {
  in, out := &in.ObservedGeneration, &out.ObservedGeneration
  *out = new(int64)
  **out = **in
 }
 if in.LastScaleTime != nil {
  in, out := &in.LastScaleTime, &out.LastScaleTime
  *out = (*in).DeepCopy()
 }
 if in.CurrentMetrics != nil {
  in, out := &in.CurrentMetrics, &out.CurrentMetrics
  *out = make([]MetricStatus, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 if in.Conditions != nil {
  in, out := &in.Conditions, &out.Conditions
  *out = make([]HorizontalPodAutoscalerCondition, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 return
}
func (in *HorizontalPodAutoscalerStatus) DeepCopy() *HorizontalPodAutoscalerStatus {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(HorizontalPodAutoscalerStatus)
 in.DeepCopyInto(out)
 return out
}
func (in *MetricIdentifier) DeepCopyInto(out *MetricIdentifier) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Selector != nil {
  in, out := &in.Selector, &out.Selector
  *out = new(v1.LabelSelector)
  (*in).DeepCopyInto(*out)
 }
 return
}
func (in *MetricIdentifier) DeepCopy() *MetricIdentifier {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(MetricIdentifier)
 in.DeepCopyInto(out)
 return out
}
func (in *MetricSpec) DeepCopyInto(out *MetricSpec) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Object != nil {
  in, out := &in.Object, &out.Object
  *out = new(ObjectMetricSource)
  (*in).DeepCopyInto(*out)
 }
 if in.Pods != nil {
  in, out := &in.Pods, &out.Pods
  *out = new(PodsMetricSource)
  (*in).DeepCopyInto(*out)
 }
 if in.Resource != nil {
  in, out := &in.Resource, &out.Resource
  *out = new(ResourceMetricSource)
  (*in).DeepCopyInto(*out)
 }
 if in.External != nil {
  in, out := &in.External, &out.External
  *out = new(ExternalMetricSource)
  (*in).DeepCopyInto(*out)
 }
 return
}
func (in *MetricSpec) DeepCopy() *MetricSpec {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(MetricSpec)
 in.DeepCopyInto(out)
 return out
}
func (in *MetricStatus) DeepCopyInto(out *MetricStatus) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Object != nil {
  in, out := &in.Object, &out.Object
  *out = new(ObjectMetricStatus)
  (*in).DeepCopyInto(*out)
 }
 if in.Pods != nil {
  in, out := &in.Pods, &out.Pods
  *out = new(PodsMetricStatus)
  (*in).DeepCopyInto(*out)
 }
 if in.Resource != nil {
  in, out := &in.Resource, &out.Resource
  *out = new(ResourceMetricStatus)
  (*in).DeepCopyInto(*out)
 }
 if in.External != nil {
  in, out := &in.External, &out.External
  *out = new(ExternalMetricStatus)
  (*in).DeepCopyInto(*out)
 }
 return
}
func (in *MetricStatus) DeepCopy() *MetricStatus {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(MetricStatus)
 in.DeepCopyInto(out)
 return out
}
func (in *MetricTarget) DeepCopyInto(out *MetricTarget) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Value != nil {
  in, out := &in.Value, &out.Value
  x := (*in).DeepCopy()
  *out = &x
 }
 if in.AverageValue != nil {
  in, out := &in.AverageValue, &out.AverageValue
  x := (*in).DeepCopy()
  *out = &x
 }
 if in.AverageUtilization != nil {
  in, out := &in.AverageUtilization, &out.AverageUtilization
  *out = new(int32)
  **out = **in
 }
 return
}
func (in *MetricTarget) DeepCopy() *MetricTarget {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(MetricTarget)
 in.DeepCopyInto(out)
 return out
}
func (in *MetricValueStatus) DeepCopyInto(out *MetricValueStatus) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Value != nil {
  in, out := &in.Value, &out.Value
  x := (*in).DeepCopy()
  *out = &x
 }
 if in.AverageValue != nil {
  in, out := &in.AverageValue, &out.AverageValue
  x := (*in).DeepCopy()
  *out = &x
 }
 if in.AverageUtilization != nil {
  in, out := &in.AverageUtilization, &out.AverageUtilization
  *out = new(int32)
  **out = **in
 }
 return
}
func (in *MetricValueStatus) DeepCopy() *MetricValueStatus {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(MetricValueStatus)
 in.DeepCopyInto(out)
 return out
}
func (in *ObjectMetricSource) DeepCopyInto(out *ObjectMetricSource) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.DescribedObject = in.DescribedObject
 in.Target.DeepCopyInto(&out.Target)
 in.Metric.DeepCopyInto(&out.Metric)
 return
}
func (in *ObjectMetricSource) DeepCopy() *ObjectMetricSource {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ObjectMetricSource)
 in.DeepCopyInto(out)
 return out
}
func (in *ObjectMetricStatus) DeepCopyInto(out *ObjectMetricStatus) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 in.Metric.DeepCopyInto(&out.Metric)
 in.Current.DeepCopyInto(&out.Current)
 out.DescribedObject = in.DescribedObject
 return
}
func (in *ObjectMetricStatus) DeepCopy() *ObjectMetricStatus {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ObjectMetricStatus)
 in.DeepCopyInto(out)
 return out
}
func (in *PodsMetricSource) DeepCopyInto(out *PodsMetricSource) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 in.Metric.DeepCopyInto(&out.Metric)
 in.Target.DeepCopyInto(&out.Target)
 return
}
func (in *PodsMetricSource) DeepCopy() *PodsMetricSource {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(PodsMetricSource)
 in.DeepCopyInto(out)
 return out
}
func (in *PodsMetricStatus) DeepCopyInto(out *PodsMetricStatus) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 in.Metric.DeepCopyInto(&out.Metric)
 in.Current.DeepCopyInto(&out.Current)
 return
}
func (in *PodsMetricStatus) DeepCopy() *PodsMetricStatus {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(PodsMetricStatus)
 in.DeepCopyInto(out)
 return out
}
func (in *ResourceMetricSource) DeepCopyInto(out *ResourceMetricSource) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 in.Target.DeepCopyInto(&out.Target)
 return
}
func (in *ResourceMetricSource) DeepCopy() *ResourceMetricSource {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ResourceMetricSource)
 in.DeepCopyInto(out)
 return out
}
func (in *ResourceMetricStatus) DeepCopyInto(out *ResourceMetricStatus) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 in.Current.DeepCopyInto(&out.Current)
 return
}
func (in *ResourceMetricStatus) DeepCopy() *ResourceMetricStatus {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ResourceMetricStatus)
 in.DeepCopyInto(out)
 return out
}
func (in *Scale) DeepCopyInto(out *Scale) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
 out.Spec = in.Spec
 out.Status = in.Status
 return
}
func (in *Scale) DeepCopy() *Scale {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(Scale)
 in.DeepCopyInto(out)
 return out
}
func (in *Scale) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *ScaleSpec) DeepCopyInto(out *ScaleSpec) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 return
}
func (in *ScaleSpec) DeepCopy() *ScaleSpec {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ScaleSpec)
 in.DeepCopyInto(out)
 return out
}
func (in *ScaleStatus) DeepCopyInto(out *ScaleStatus) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 return
}
func (in *ScaleStatus) DeepCopy() *ScaleStatus {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ScaleStatus)
 in.DeepCopyInto(out)
 return out
}
