package v1beta1

import (
 "math"
 "k8s.io/api/core/v1"
 extensionsv1beta1 "k8s.io/api/extensions/v1beta1"
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 "k8s.io/apimachinery/pkg/runtime"
 "k8s.io/apimachinery/pkg/util/intstr"
)

func addDefaultingFuncs(scheme *runtime.Scheme) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return RegisterDefaults(scheme)
}
func SetDefaults_DaemonSet(obj *extensionsv1beta1.DaemonSet) {
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
 updateStrategy := &obj.Spec.UpdateStrategy
 if updateStrategy.Type == "" {
  updateStrategy.Type = extensionsv1beta1.OnDeleteDaemonSetStrategyType
 }
 if updateStrategy.Type == extensionsv1beta1.RollingUpdateDaemonSetStrategyType {
  if updateStrategy.RollingUpdate == nil {
   rollingUpdate := extensionsv1beta1.RollingUpdateDaemonSet{}
   updateStrategy.RollingUpdate = &rollingUpdate
  }
  if updateStrategy.RollingUpdate.MaxUnavailable == nil {
   maxUnavailable := intstr.FromInt(1)
   updateStrategy.RollingUpdate.MaxUnavailable = &maxUnavailable
  }
 }
 if obj.Spec.RevisionHistoryLimit == nil {
  obj.Spec.RevisionHistoryLimit = new(int32)
  *obj.Spec.RevisionHistoryLimit = 10
 }
}
func SetDefaults_PodSecurityPolicySpec(obj *extensionsv1beta1.PodSecurityPolicySpec) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if obj.AllowPrivilegeEscalation == nil {
  t := true
  obj.AllowPrivilegeEscalation = &t
 }
}
func SetDefaults_Deployment(obj *extensionsv1beta1.Deployment) {
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
  strategy.Type = extensionsv1beta1.RollingUpdateDeploymentStrategyType
 }
 if strategy.Type == extensionsv1beta1.RollingUpdateDeploymentStrategyType || strategy.RollingUpdate != nil {
  if strategy.RollingUpdate == nil {
   rollingUpdate := extensionsv1beta1.RollingUpdateDeployment{}
   strategy.RollingUpdate = &rollingUpdate
  }
  if strategy.RollingUpdate.MaxUnavailable == nil {
   maxUnavailable := intstr.FromInt(1)
   strategy.RollingUpdate.MaxUnavailable = &maxUnavailable
  }
  if strategy.RollingUpdate.MaxSurge == nil {
   maxSurge := intstr.FromInt(1)
   strategy.RollingUpdate.MaxSurge = &maxSurge
  }
 }
 if obj.Spec.ProgressDeadlineSeconds == nil {
  obj.Spec.ProgressDeadlineSeconds = new(int32)
  *obj.Spec.ProgressDeadlineSeconds = math.MaxInt32
 }
 if obj.Spec.RevisionHistoryLimit == nil {
  obj.Spec.RevisionHistoryLimit = new(int32)
  *obj.Spec.RevisionHistoryLimit = math.MaxInt32
 }
}
func SetDefaults_ReplicaSet(obj *extensionsv1beta1.ReplicaSet) {
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
}
func SetDefaults_NetworkPolicy(obj *extensionsv1beta1.NetworkPolicy) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for _, i := range obj.Spec.Ingress {
  for _, p := range i.Ports {
   if p.Protocol == nil {
    proto := v1.ProtocolTCP
    p.Protocol = &proto
   }
  }
 }
 if len(obj.Spec.PolicyTypes) == 0 {
  obj.Spec.PolicyTypes = []extensionsv1beta1.PolicyType{extensionsv1beta1.PolicyTypeIngress}
  if len(obj.Spec.Egress) != 0 {
   obj.Spec.PolicyTypes = append(obj.Spec.PolicyTypes, extensionsv1beta1.PolicyTypeEgress)
  }
 }
}
