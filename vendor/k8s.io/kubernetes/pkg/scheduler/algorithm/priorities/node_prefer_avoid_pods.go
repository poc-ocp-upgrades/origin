package priorities

import (
	"fmt"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1helper "k8s.io/kubernetes/pkg/apis/core/v1/helper"
	schedulerapi "k8s.io/kubernetes/pkg/scheduler/api"
	schedulercache "k8s.io/kubernetes/pkg/scheduler/cache"
)

func CalculateNodePreferAvoidPodsPriorityMap(pod *v1.Pod, meta interface{}, nodeInfo *schedulercache.NodeInfo) (schedulerapi.HostPriority, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	node := nodeInfo.Node()
	if node == nil {
		return schedulerapi.HostPriority{}, fmt.Errorf("node not found")
	}
	var controllerRef *metav1.OwnerReference
	if priorityMeta, ok := meta.(*priorityMetadata); ok {
		controllerRef = priorityMeta.controllerRef
	} else {
		controllerRef = metav1.GetControllerOf(pod)
	}
	if controllerRef != nil {
		if controllerRef.Kind != "ReplicationController" && controllerRef.Kind != "ReplicaSet" {
			controllerRef = nil
		}
	}
	if controllerRef == nil {
		return schedulerapi.HostPriority{Host: node.Name, Score: schedulerapi.MaxPriority}, nil
	}
	avoids, err := v1helper.GetAvoidPodsFromNodeAnnotations(node.Annotations)
	if err != nil {
		return schedulerapi.HostPriority{Host: node.Name, Score: schedulerapi.MaxPriority}, nil
	}
	for i := range avoids.PreferAvoidPods {
		avoid := &avoids.PreferAvoidPods[i]
		if avoid.PodSignature.PodController.Kind == controllerRef.Kind && avoid.PodSignature.PodController.UID == controllerRef.UID {
			return schedulerapi.HostPriority{Host: node.Name, Score: 0}, nil
		}
	}
	return schedulerapi.HostPriority{Host: node.Name, Score: schedulerapi.MaxPriority}, nil
}
