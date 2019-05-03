package v1beta1

import (
 appsv1beta1 "k8s.io/api/apps/v1beta1"
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 "k8s.io/apimachinery/pkg/runtime"
 "k8s.io/apimachinery/pkg/util/intstr"
)

func addDefaultingFuncs(scheme *runtime.Scheme) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return RegisterDefaults(scheme)
}
func SetDefaults_StatefulSet(obj *appsv1beta1.StatefulSet) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if len(obj.Spec.PodManagementPolicy) == 0 {
  obj.Spec.PodManagementPolicy = appsv1beta1.OrderedReadyPodManagement
 }
 if obj.Spec.UpdateStrategy.Type == "" {
  obj.Spec.UpdateStrategy.Type = appsv1beta1.OnDeleteStatefulSetStrategyType
 }
 labels := obj.Spec.Template.Labels
 if labels != nil {
  if obj.Spec.Selector == nil {
   obj.Spec.Selector = &metav1.LabelSelector{MatchLabels: labels}
  }
  if len(obj.Labels) == 0 {
   obj.Labels = labels
  }
 }
 if obj.Spec.Replicas == nil {
  obj.Spec.Replicas = new(int32)
  *obj.Spec.Replicas = 1
 }
 if obj.Spec.RevisionHistoryLimit == nil {
  obj.Spec.RevisionHistoryLimit = new(int32)
  *obj.Spec.RevisionHistoryLimit = 10
 }
 if obj.Spec.UpdateStrategy.Type == appsv1beta1.RollingUpdateStatefulSetStrategyType && obj.Spec.UpdateStrategy.RollingUpdate != nil && obj.Spec.UpdateStrategy.RollingUpdate.Partition == nil {
  obj.Spec.UpdateStrategy.RollingUpdate.Partition = new(int32)
  *obj.Spec.UpdateStrategy.RollingUpdate.Partition = 0
 }
}
func SetDefaults_Deployment(obj *appsv1beta1.Deployment) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 labels := obj.Spec.Template.Labels
 if labels != nil {
  if obj.Spec.Selector == nil {
   obj.Spec.Selector = &metav1.LabelSelector{MatchLabels: labels}
  }
  if len(obj.Labels) == 0 {
   obj.Labels = labels
  }
 }
 if obj.Spec.Replicas == nil {
  obj.Spec.Replicas = new(int32)
  *obj.Spec.Replicas = 1
 }
 strategy := &obj.Spec.Strategy
 if strategy.Type == "" {
  strategy.Type = appsv1beta1.RollingUpdateDeploymentStrategyType
 }
 if strategy.Type == appsv1beta1.RollingUpdateDeploymentStrategyType {
  if strategy.RollingUpdate == nil {
   rollingUpdate := appsv1beta1.RollingUpdateDeployment{}
   strategy.RollingUpdate = &rollingUpdate
  }
  if strategy.RollingUpdate.MaxUnavailable == nil {
   maxUnavailable := intstr.FromString("25%")
   strategy.RollingUpdate.MaxUnavailable = &maxUnavailable
  }
  if strategy.RollingUpdate.MaxSurge == nil {
   maxSurge := intstr.FromString("25%")
   strategy.RollingUpdate.MaxSurge = &maxSurge
  }
 }
 if obj.Spec.RevisionHistoryLimit == nil {
  obj.Spec.RevisionHistoryLimit = new(int32)
  *obj.Spec.RevisionHistoryLimit = 2
 }
 if obj.Spec.ProgressDeadlineSeconds == nil {
  obj.Spec.ProgressDeadlineSeconds = new(int32)
  *obj.Spec.ProgressDeadlineSeconds = 600
 }
}
