package v2beta1

import (
 v2beta1 "k8s.io/api/autoscaling/v2beta1"
 runtime "k8s.io/apimachinery/pkg/runtime"
)

func RegisterDefaults(scheme *runtime.Scheme) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 scheme.AddTypeDefaultingFunc(&v2beta1.HorizontalPodAutoscaler{}, func(obj interface{}) {
  SetObjectDefaults_HorizontalPodAutoscaler(obj.(*v2beta1.HorizontalPodAutoscaler))
 })
 scheme.AddTypeDefaultingFunc(&v2beta1.HorizontalPodAutoscalerList{}, func(obj interface{}) {
  SetObjectDefaults_HorizontalPodAutoscalerList(obj.(*v2beta1.HorizontalPodAutoscalerList))
 })
 return nil
}
func SetObjectDefaults_HorizontalPodAutoscaler(in *v2beta1.HorizontalPodAutoscaler) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 SetDefaults_HorizontalPodAutoscaler(in)
}
func SetObjectDefaults_HorizontalPodAutoscalerList(in *v2beta1.HorizontalPodAutoscalerList) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for i := range in.Items {
  a := &in.Items[i]
  SetObjectDefaults_HorizontalPodAutoscaler(a)
 }
}
