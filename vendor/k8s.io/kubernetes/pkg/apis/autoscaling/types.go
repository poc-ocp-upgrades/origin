package autoscaling

import (
 "k8s.io/apimachinery/pkg/api/resource"
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 api "k8s.io/kubernetes/pkg/apis/core"
)

type Scale struct {
 metav1.TypeMeta
 metav1.ObjectMeta
 Spec   ScaleSpec
 Status ScaleStatus
}
type ScaleSpec struct{ Replicas int32 }
type ScaleStatus struct {
 Replicas int32
 Selector string
}
type CrossVersionObjectReference struct {
 Kind       string
 Name       string
 APIVersion string
}
type HorizontalPodAutoscalerSpec struct {
 ScaleTargetRef CrossVersionObjectReference
 MinReplicas    *int32
 MaxReplicas    int32
 Metrics        []MetricSpec
}
type MetricSourceType string

var (
 ObjectMetricSourceType   MetricSourceType = "Object"
 PodsMetricSourceType     MetricSourceType = "Pods"
 ResourceMetricSourceType MetricSourceType = "Resource"
 ExternalMetricSourceType MetricSourceType = "External"
)

type MetricSpec struct {
 Type     MetricSourceType
 Object   *ObjectMetricSource
 Pods     *PodsMetricSource
 Resource *ResourceMetricSource
 External *ExternalMetricSource
}
type ObjectMetricSource struct {
 DescribedObject CrossVersionObjectReference
 Target          MetricTarget
 Metric          MetricIdentifier
}
type PodsMetricSource struct {
 Metric MetricIdentifier
 Target MetricTarget
}
type ResourceMetricSource struct {
 Name   api.ResourceName
 Target MetricTarget
}
type ExternalMetricSource struct {
 Metric MetricIdentifier
 Target MetricTarget
}
type MetricIdentifier struct {
 Name     string
 Selector *metav1.LabelSelector
}
type MetricTarget struct {
 Type               MetricTargetType
 Value              *resource.Quantity
 AverageValue       *resource.Quantity
 AverageUtilization *int32
}
type MetricTargetType string

var (
 UtilizationMetricType  MetricTargetType = "Utilization"
 ValueMetricType        MetricTargetType = "Value"
 AverageValueMetricType MetricTargetType = "AverageValue"
)

type HorizontalPodAutoscalerStatus struct {
 ObservedGeneration *int64
 LastScaleTime      *metav1.Time
 CurrentReplicas    int32
 DesiredReplicas    int32
 CurrentMetrics     []MetricStatus
 Conditions         []HorizontalPodAutoscalerCondition
}
type ConditionStatus string

const (
 ConditionTrue    ConditionStatus = "True"
 ConditionFalse   ConditionStatus = "False"
 ConditionUnknown ConditionStatus = "Unknown"
)

type HorizontalPodAutoscalerConditionType string

var (
 ScalingActive  HorizontalPodAutoscalerConditionType = "ScalingActive"
 AbleToScale    HorizontalPodAutoscalerConditionType = "AbleToScale"
 ScalingLimited HorizontalPodAutoscalerConditionType = "ScalingLimited"
)

type HorizontalPodAutoscalerCondition struct {
 Type               HorizontalPodAutoscalerConditionType
 Status             ConditionStatus
 LastTransitionTime metav1.Time
 Reason             string
 Message            string
}
type MetricStatus struct {
 Type     MetricSourceType
 Object   *ObjectMetricStatus
 Pods     *PodsMetricStatus
 Resource *ResourceMetricStatus
 External *ExternalMetricStatus
}
type ObjectMetricStatus struct {
 Metric          MetricIdentifier
 Current         MetricValueStatus
 DescribedObject CrossVersionObjectReference
}
type PodsMetricStatus struct {
 Metric  MetricIdentifier
 Current MetricValueStatus
}
type ResourceMetricStatus struct {
 Name    api.ResourceName
 Current MetricValueStatus
}
type ExternalMetricStatus struct {
 Metric  MetricIdentifier
 Current MetricValueStatus
}
type MetricValueStatus struct {
 Value              *resource.Quantity
 AverageValue       *resource.Quantity
 AverageUtilization *int32
}
type HorizontalPodAutoscaler struct {
 metav1.TypeMeta
 metav1.ObjectMeta
 Spec   HorizontalPodAutoscalerSpec
 Status HorizontalPodAutoscalerStatus
}
type HorizontalPodAutoscalerList struct {
 metav1.TypeMeta
 metav1.ListMeta
 Items []HorizontalPodAutoscaler
}
