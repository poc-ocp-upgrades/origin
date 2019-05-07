package priorities

import (
	"fmt"
	"k8s.io/api/core/v1"
	v1helper "k8s.io/kubernetes/pkg/apis/core/v1/helper"
	schedulerapi "k8s.io/kubernetes/pkg/scheduler/api"
	schedulercache "k8s.io/kubernetes/pkg/scheduler/cache"
)

func countIntolerableTaintsPreferNoSchedule(taints []v1.Taint, tolerations []v1.Toleration) (intolerableTaints int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, taint := range taints {
		if taint.Effect != v1.TaintEffectPreferNoSchedule {
			continue
		}
		if !v1helper.TolerationsTolerateTaint(tolerations, &taint) {
			intolerableTaints++
		}
	}
	return
}
func getAllTolerationPreferNoSchedule(tolerations []v1.Toleration) (tolerationList []v1.Toleration) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, toleration := range tolerations {
		if len(toleration.Effect) == 0 || toleration.Effect == v1.TaintEffectPreferNoSchedule {
			tolerationList = append(tolerationList, toleration)
		}
	}
	return
}
func ComputeTaintTolerationPriorityMap(pod *v1.Pod, meta interface{}, nodeInfo *schedulercache.NodeInfo) (schedulerapi.HostPriority, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	node := nodeInfo.Node()
	if node == nil {
		return schedulerapi.HostPriority{}, fmt.Errorf("node not found")
	}
	var tolerationsPreferNoSchedule []v1.Toleration
	if priorityMeta, ok := meta.(*priorityMetadata); ok {
		tolerationsPreferNoSchedule = priorityMeta.podTolerations
	} else {
		tolerationsPreferNoSchedule = getAllTolerationPreferNoSchedule(pod.Spec.Tolerations)
	}
	return schedulerapi.HostPriority{Host: node.Name, Score: countIntolerableTaintsPreferNoSchedule(node.Spec.Taints, tolerationsPreferNoSchedule)}, nil
}

var ComputeTaintTolerationPriorityReduce = NormalizeReduce(schedulerapi.MaxPriority, true)
