package v1

import (
 autoscalingv1 "k8s.io/api/autoscaling/v1"
 "k8s.io/apimachinery/pkg/runtime"
)

func addDefaultingFuncs(scheme *runtime.Scheme) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return RegisterDefaults(scheme)
}
func SetDefaults_HorizontalPodAutoscaler(obj *autoscalingv1.HorizontalPodAutoscaler) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if obj.Spec.MinReplicas == nil {
  minReplicas := int32(1)
  obj.Spec.MinReplicas = &minReplicas
 }
}
