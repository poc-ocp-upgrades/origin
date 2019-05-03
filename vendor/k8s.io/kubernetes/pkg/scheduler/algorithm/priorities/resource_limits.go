package priorities

import (
 "fmt"
 "k8s.io/api/core/v1"
 schedulerapi "k8s.io/kubernetes/pkg/scheduler/api"
 schedulercache "k8s.io/kubernetes/pkg/scheduler/cache"
 "k8s.io/klog"
)

func ResourceLimitsPriorityMap(pod *v1.Pod, meta interface{}, nodeInfo *schedulercache.NodeInfo) (schedulerapi.HostPriority, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 node := nodeInfo.Node()
 if node == nil {
  return schedulerapi.HostPriority{}, fmt.Errorf("node not found")
 }
 allocatableResources := nodeInfo.AllocatableResource()
 podLimits := getResourceLimits(pod)
 cpuScore := computeScore(podLimits.MilliCPU, allocatableResources.MilliCPU)
 memScore := computeScore(podLimits.Memory, allocatableResources.Memory)
 score := int(0)
 if cpuScore == 1 || memScore == 1 {
  score = 1
 }
 if klog.V(10) {
  klog.Infof("%v -> %v: Resource Limits Priority, allocatable %d millicores %d memory bytes, pod limits %d millicores %d memory bytes, score %d", pod.Name, node.Name, allocatableResources.MilliCPU, allocatableResources.Memory, podLimits.MilliCPU, podLimits.Memory, score)
 }
 return schedulerapi.HostPriority{Host: node.Name, Score: score}, nil
}
func computeScore(limit, allocatable int64) int64 {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if limit != 0 && allocatable != 0 && limit <= allocatable {
  return 1
 }
 return 0
}
func getResourceLimits(pod *v1.Pod) *schedulercache.Resource {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result := &schedulercache.Resource{}
 for _, container := range pod.Spec.Containers {
  result.Add(container.Resources.Limits)
 }
 for _, container := range pod.Spec.InitContainers {
  result.SetMaxResource(container.Resources.Limits)
 }
 return result
}
