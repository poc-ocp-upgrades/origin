package fuzzer

import (
	goformat "fmt"
	fuzz "github.com/google/gofuzz"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtimeserializer "k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/kubernetes/pkg/apis/autoscaling"
	api "k8s.io/kubernetes/pkg/apis/core"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

var Funcs = func(codecs runtimeserializer.CodecFactory) []interface{} {
	return []interface{}{func(s *autoscaling.ScaleStatus, c fuzz.Continue) {
		c.FuzzNoCustom(s)
		metaSelector := &metav1.LabelSelector{}
		c.Fuzz(metaSelector)
		labelSelector, _ := metav1.LabelSelectorAsSelector(metaSelector)
		s.Selector = labelSelector.String()
	}, func(s *autoscaling.HorizontalPodAutoscalerSpec, c fuzz.Continue) {
		c.FuzzNoCustom(s)
		minReplicas := int32(c.Rand.Int31())
		s.MinReplicas = &minReplicas
		randomQuantity := func() resource.Quantity {
			var q resource.Quantity
			c.Fuzz(&q)
			_ = q.String()
			return q
		}
		var podMetricID autoscaling.MetricIdentifier
		var objMetricID autoscaling.MetricIdentifier
		c.Fuzz(&podMetricID)
		c.Fuzz(&objMetricID)
		targetUtilization := int32(c.RandUint64())
		averageValue := randomQuantity()
		s.Metrics = []autoscaling.MetricSpec{{Type: autoscaling.PodsMetricSourceType, Pods: &autoscaling.PodsMetricSource{Metric: podMetricID, Target: autoscaling.MetricTarget{Type: autoscaling.AverageValueMetricType, AverageValue: &averageValue}}}, {Type: autoscaling.ObjectMetricSourceType, Object: &autoscaling.ObjectMetricSource{Metric: objMetricID, Target: autoscaling.MetricTarget{Type: autoscaling.ValueMetricType, Value: &averageValue}}}, {Type: autoscaling.ResourceMetricSourceType, Resource: &autoscaling.ResourceMetricSource{Name: api.ResourceCPU, Target: autoscaling.MetricTarget{Type: autoscaling.UtilizationMetricType, AverageUtilization: &targetUtilization}}}}
	}, func(s *autoscaling.HorizontalPodAutoscalerStatus, c fuzz.Continue) {
		c.FuzzNoCustom(s)
		randomQuantity := func() resource.Quantity {
			var q resource.Quantity
			c.Fuzz(&q)
			_ = q.String()
			return q
		}
		averageValue := randomQuantity()
		currentUtilization := int32(c.RandUint64())
		s.CurrentMetrics = []autoscaling.MetricStatus{{Type: autoscaling.PodsMetricSourceType, Pods: &autoscaling.PodsMetricStatus{Metric: autoscaling.MetricIdentifier{Name: c.RandString()}, Current: autoscaling.MetricValueStatus{AverageValue: &averageValue}}}, {Type: autoscaling.ResourceMetricSourceType, Resource: &autoscaling.ResourceMetricStatus{Name: api.ResourceCPU, Current: autoscaling.MetricValueStatus{AverageUtilization: &currentUtilization, AverageValue: &averageValue}}}}
	}}
}

func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
