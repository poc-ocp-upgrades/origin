package v2beta2

import (
	goformat "fmt"
	autoscalingv2beta2 "k8s.io/api/autoscaling/v2beta2"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/kubernetes/pkg/apis/autoscaling"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

func addDefaultingFuncs(scheme *runtime.Scheme) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return RegisterDefaults(scheme)
}
func SetDefaults_HorizontalPodAutoscaler(obj *autoscalingv2beta2.HorizontalPodAutoscaler) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if obj.Spec.MinReplicas == nil {
		minReplicas := int32(1)
		obj.Spec.MinReplicas = &minReplicas
	}
	if len(obj.Spec.Metrics) == 0 {
		utilizationDefaultVal := int32(autoscaling.DefaultCPUUtilization)
		obj.Spec.Metrics = []autoscalingv2beta2.MetricSpec{{Type: autoscalingv2beta2.ResourceMetricSourceType, Resource: &autoscalingv2beta2.ResourceMetricSource{Name: v1.ResourceCPU, Target: autoscalingv2beta2.MetricTarget{Type: autoscalingv2beta2.UtilizationMetricType, AverageUtilization: &utilizationDefaultVal}}}}
	}
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
