package priorities

import (
	"fmt"
	"k8s.io/api/core/v1"
	schedulerapi "k8s.io/kubernetes/pkg/scheduler/api"
	schedulercache "k8s.io/kubernetes/pkg/scheduler/cache"
	"k8s.io/kubernetes/pkg/util/parsers"
	"strings"
)

const (
	mb           int64 = 1024 * 1024
	minThreshold int64 = 23 * mb
	maxThreshold int64 = 1000 * mb
)

func ImageLocalityPriorityMap(pod *v1.Pod, meta interface{}, nodeInfo *schedulercache.NodeInfo) (schedulerapi.HostPriority, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	node := nodeInfo.Node()
	if node == nil {
		return schedulerapi.HostPriority{}, fmt.Errorf("node not found")
	}
	var score int
	if priorityMeta, ok := meta.(*priorityMetadata); ok {
		score = calculatePriority(sumImageScores(nodeInfo, pod.Spec.Containers, priorityMeta.totalNumNodes))
	} else {
		score = 0
	}
	return schedulerapi.HostPriority{Host: node.Name, Score: score}, nil
}
func calculatePriority(sumScores int64) int {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if sumScores < minThreshold {
		sumScores = minThreshold
	} else if sumScores > maxThreshold {
		sumScores = maxThreshold
	}
	return int(int64(schedulerapi.MaxPriority) * (sumScores - minThreshold) / (maxThreshold - minThreshold))
}
func sumImageScores(nodeInfo *schedulercache.NodeInfo, containers []v1.Container, totalNumNodes int) int64 {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var sum int64
	imageStates := nodeInfo.ImageStates()
	for _, container := range containers {
		if state, ok := imageStates[normalizedImageName(container.Image)]; ok {
			sum += scaledImageScore(state, totalNumNodes)
		}
	}
	return sum
}
func scaledImageScore(imageState *schedulercache.ImageStateSummary, totalNumNodes int) int64 {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	spread := float64(imageState.NumNodes) / float64(totalNumNodes)
	return int64(float64(imageState.Size) * spread)
}
func normalizedImageName(name string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if strings.LastIndex(name, ":") <= strings.LastIndex(name, "/") {
		name = name + ":" + parsers.DefaultImageTag
	}
	return name
}
