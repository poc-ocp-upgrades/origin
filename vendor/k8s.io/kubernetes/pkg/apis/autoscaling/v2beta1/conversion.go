package v2beta1

import (
	goformat "fmt"
	autoscalingv2beta1 "k8s.io/api/autoscaling/v2beta1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/conversion"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/kubernetes/pkg/apis/autoscaling"
	core "k8s.io/kubernetes/pkg/apis/core"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

func addConversionFuncs(scheme *runtime.Scheme) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	err := scheme.AddConversionFuncs(Convert_autoscaling_ExternalMetricSource_To_v2beta1_ExternalMetricSource, Convert_v2beta1_ExternalMetricSource_To_autoscaling_ExternalMetricSource, Convert_autoscaling_ObjectMetricSource_To_v2beta1_ObjectMetricSource, Convert_v2beta1_ObjectMetricSource_To_autoscaling_ObjectMetricSource, Convert_autoscaling_PodsMetricSource_To_v2beta1_PodsMetricSource, Convert_v2beta1_PodsMetricSource_To_autoscaling_PodsMetricSource, Convert_autoscaling_ExternalMetricStatus_To_v2beta1_ExternalMetricStatus, Convert_v2beta1_ExternalMetricStatus_To_autoscaling_ExternalMetricStatus, Convert_autoscaling_ObjectMetricStatus_To_v2beta1_ObjectMetricStatus, Convert_v2beta1_ObjectMetricStatus_To_autoscaling_ObjectMetricStatus, Convert_autoscaling_PodsMetricStatus_To_v2beta1_PodsMetricStatus, Convert_v2beta1_PodsMetricStatus_To_autoscaling_PodsMetricStatus, Convert_autoscaling_ResourceMetricSource_To_v2beta1_ResourceMetricSource, Convert_v2beta1_ResourceMetricSource_To_autoscaling_ResourceMetricSource, Convert_autoscaling_MetricTarget_To_v2beta1_CrossVersionObjectReference, Convert_v2beta1_CrossVersionObjectReference_To_autoscaling_MetricTarget, Convert_autoscaling_ResourceMetricStatus_To_v2beta1_ResourceMetricStatus, Convert_v2beta1_ResourceMetricStatus_To_autoscaling_ResourceMetricStatus, Convert_autoscaling_HorizontalPodAutoscaler_To_v2beta1_HorizontalPodAutoscaler, Convert_v2beta1_HorizontalPodAutoscaler_To_autoscaling_HorizontalPodAutoscaler)
	if err != nil {
		return err
	}
	return nil
}
func Convert_autoscaling_MetricTarget_To_v2beta1_CrossVersionObjectReference(in *autoscaling.MetricTarget, out *autoscalingv2beta1.CrossVersionObjectReference, s conversion.Scope) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return nil
}
func Convert_v2beta1_CrossVersionObjectReference_To_autoscaling_MetricTarget(in *autoscalingv2beta1.CrossVersionObjectReference, out *autoscaling.MetricTarget, s conversion.Scope) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return nil
}
func Convert_v2beta1_ResourceMetricStatus_To_autoscaling_ResourceMetricStatus(in *autoscalingv2beta1.ResourceMetricStatus, out *autoscaling.ResourceMetricStatus, s conversion.Scope) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	out.Name = core.ResourceName(in.Name)
	utilization := in.CurrentAverageUtilization
	averageValue := in.CurrentAverageValue
	out.Current = autoscaling.MetricValueStatus{AverageValue: &averageValue, AverageUtilization: utilization}
	return nil
}
func Convert_autoscaling_ResourceMetricStatus_To_v2beta1_ResourceMetricStatus(in *autoscaling.ResourceMetricStatus, out *autoscalingv2beta1.ResourceMetricStatus, s conversion.Scope) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	out.Name = v1.ResourceName(in.Name)
	out.CurrentAverageUtilization = in.Current.AverageUtilization
	if in.Current.AverageValue != nil {
		out.CurrentAverageValue = *in.Current.AverageValue
	}
	return nil
}
func Convert_v2beta1_ResourceMetricSource_To_autoscaling_ResourceMetricSource(in *autoscalingv2beta1.ResourceMetricSource, out *autoscaling.ResourceMetricSource, s conversion.Scope) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	out.Name = core.ResourceName(in.Name)
	utilization := in.TargetAverageUtilization
	averageValue := in.TargetAverageValue
	var metricType autoscaling.MetricTargetType
	if utilization == nil {
		metricType = autoscaling.AverageValueMetricType
	} else {
		metricType = autoscaling.UtilizationMetricType
	}
	out.Target = autoscaling.MetricTarget{Type: metricType, AverageValue: averageValue, AverageUtilization: utilization}
	return nil
}
func Convert_autoscaling_ResourceMetricSource_To_v2beta1_ResourceMetricSource(in *autoscaling.ResourceMetricSource, out *autoscalingv2beta1.ResourceMetricSource, s conversion.Scope) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	out.Name = v1.ResourceName(in.Name)
	out.TargetAverageUtilization = in.Target.AverageUtilization
	out.TargetAverageValue = in.Target.AverageValue
	return nil
}
func Convert_autoscaling_ExternalMetricSource_To_v2beta1_ExternalMetricSource(in *autoscaling.ExternalMetricSource, out *autoscalingv2beta1.ExternalMetricSource, s conversion.Scope) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	out.MetricName = in.Metric.Name
	out.TargetValue = in.Target.Value
	out.TargetAverageValue = in.Target.AverageValue
	out.MetricSelector = in.Metric.Selector
	return nil
}
func Convert_v2beta1_ExternalMetricSource_To_autoscaling_ExternalMetricSource(in *autoscalingv2beta1.ExternalMetricSource, out *autoscaling.ExternalMetricSource, s conversion.Scope) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	value := in.TargetValue
	averageValue := in.TargetAverageValue
	var metricType autoscaling.MetricTargetType
	if value == nil {
		metricType = autoscaling.AverageValueMetricType
	} else {
		metricType = autoscaling.ValueMetricType
	}
	out.Target = autoscaling.MetricTarget{Type: metricType, Value: value, AverageValue: averageValue}
	out.Metric = autoscaling.MetricIdentifier{Name: in.MetricName, Selector: in.MetricSelector}
	return nil
}
func Convert_autoscaling_ObjectMetricSource_To_v2beta1_ObjectMetricSource(in *autoscaling.ObjectMetricSource, out *autoscalingv2beta1.ObjectMetricSource, s conversion.Scope) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if in.Target.Value != nil {
		out.TargetValue = *in.Target.Value
	}
	out.AverageValue = in.Target.AverageValue
	out.Target = autoscalingv2beta1.CrossVersionObjectReference{Kind: in.DescribedObject.Kind, Name: in.DescribedObject.Name, APIVersion: in.DescribedObject.APIVersion}
	out.MetricName = in.Metric.Name
	out.Selector = in.Metric.Selector
	return nil
}
func Convert_v2beta1_ObjectMetricSource_To_autoscaling_ObjectMetricSource(in *autoscalingv2beta1.ObjectMetricSource, out *autoscaling.ObjectMetricSource, s conversion.Scope) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var metricType autoscaling.MetricTargetType
	if in.AverageValue == nil {
		metricType = autoscaling.ValueMetricType
	} else {
		metricType = autoscaling.AverageValueMetricType
	}
	out.Target = autoscaling.MetricTarget{Type: metricType, Value: &in.TargetValue, AverageValue: in.AverageValue}
	out.DescribedObject = autoscaling.CrossVersionObjectReference{Kind: in.Target.Kind, Name: in.Target.Name, APIVersion: in.Target.APIVersion}
	out.Metric = autoscaling.MetricIdentifier{Name: in.MetricName, Selector: in.Selector}
	return nil
}
func Convert_autoscaling_PodsMetricSource_To_v2beta1_PodsMetricSource(in *autoscaling.PodsMetricSource, out *autoscalingv2beta1.PodsMetricSource, s conversion.Scope) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if in.Target.AverageValue != nil {
		targetAverageValue := *in.Target.AverageValue
		out.TargetAverageValue = targetAverageValue
	}
	out.MetricName = in.Metric.Name
	out.Selector = in.Metric.Selector
	return nil
}
func Convert_v2beta1_PodsMetricSource_To_autoscaling_PodsMetricSource(in *autoscalingv2beta1.PodsMetricSource, out *autoscaling.PodsMetricSource, s conversion.Scope) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	targetAverageValue := &in.TargetAverageValue
	var metricType autoscaling.MetricTargetType
	metricType = autoscaling.AverageValueMetricType
	out.Target = autoscaling.MetricTarget{Type: metricType, AverageValue: targetAverageValue}
	out.Metric = autoscaling.MetricIdentifier{Name: in.MetricName, Selector: in.Selector}
	return nil
}
func Convert_autoscaling_ExternalMetricStatus_To_v2beta1_ExternalMetricStatus(in *autoscaling.ExternalMetricStatus, out *autoscalingv2beta1.ExternalMetricStatus, s conversion.Scope) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if &in.Current.AverageValue != nil {
		out.CurrentAverageValue = in.Current.AverageValue
	}
	out.MetricName = in.Metric.Name
	if in.Current.Value != nil {
		out.CurrentValue = *in.Current.Value
	}
	out.MetricSelector = in.Metric.Selector
	return nil
}
func Convert_v2beta1_ExternalMetricStatus_To_autoscaling_ExternalMetricStatus(in *autoscalingv2beta1.ExternalMetricStatus, out *autoscaling.ExternalMetricStatus, s conversion.Scope) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	value := in.CurrentValue
	averageValue := in.CurrentAverageValue
	out.Current = autoscaling.MetricValueStatus{Value: &value, AverageValue: averageValue}
	out.Metric = autoscaling.MetricIdentifier{Name: in.MetricName, Selector: in.MetricSelector}
	return nil
}
func Convert_autoscaling_ObjectMetricStatus_To_v2beta1_ObjectMetricStatus(in *autoscaling.ObjectMetricStatus, out *autoscalingv2beta1.ObjectMetricStatus, s conversion.Scope) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if in.Current.Value != nil {
		out.CurrentValue = *in.Current.Value
	}
	out.Target = autoscalingv2beta1.CrossVersionObjectReference{Kind: in.DescribedObject.Kind, Name: in.DescribedObject.Name, APIVersion: in.DescribedObject.APIVersion}
	out.MetricName = in.Metric.Name
	out.Selector = in.Metric.Selector
	if in.Current.AverageValue != nil {
		currentAverageValue := *in.Current.AverageValue
		out.AverageValue = &currentAverageValue
	}
	return nil
}
func Convert_v2beta1_ObjectMetricStatus_To_autoscaling_ObjectMetricStatus(in *autoscalingv2beta1.ObjectMetricStatus, out *autoscaling.ObjectMetricStatus, s conversion.Scope) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	out.Current = autoscaling.MetricValueStatus{Value: &in.CurrentValue, AverageValue: in.AverageValue}
	out.DescribedObject = autoscaling.CrossVersionObjectReference{Kind: in.Target.Kind, Name: in.Target.Name, APIVersion: in.Target.APIVersion}
	out.Metric = autoscaling.MetricIdentifier{Name: in.MetricName, Selector: in.Selector}
	return nil
}
func Convert_autoscaling_PodsMetricStatus_To_v2beta1_PodsMetricStatus(in *autoscaling.PodsMetricStatus, out *autoscalingv2beta1.PodsMetricStatus, s conversion.Scope) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if in.Current.AverageValue != nil {
		out.CurrentAverageValue = *in.Current.AverageValue
	}
	out.MetricName = in.Metric.Name
	out.Selector = in.Metric.Selector
	return nil
}
func Convert_v2beta1_PodsMetricStatus_To_autoscaling_PodsMetricStatus(in *autoscalingv2beta1.PodsMetricStatus, out *autoscaling.PodsMetricStatus, s conversion.Scope) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	out.Current = autoscaling.MetricValueStatus{AverageValue: &in.CurrentAverageValue}
	out.Metric = autoscaling.MetricIdentifier{Name: in.MetricName, Selector: in.Selector}
	return nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
