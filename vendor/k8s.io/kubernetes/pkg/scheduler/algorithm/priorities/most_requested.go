package priorities

import (
	schedulerapi "k8s.io/kubernetes/pkg/scheduler/api"
	schedulercache "k8s.io/kubernetes/pkg/scheduler/cache"
)

var (
	mostResourcePriority     = &ResourceAllocationPriority{"MostResourceAllocation", mostResourceScorer}
	MostRequestedPriorityMap = mostResourcePriority.PriorityMap
)

func mostResourceScorer(requested, allocable *schedulercache.Resource, includeVolumes bool, requestedVolumes int, allocatableVolumes int) int64 {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return (mostRequestedScore(requested.MilliCPU, allocable.MilliCPU) + mostRequestedScore(requested.Memory, allocable.Memory)) / 2
}
func mostRequestedScore(requested, capacity int64) int64 {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if capacity == 0 {
		return 0
	}
	if requested > capacity {
		return 0
	}
	return (requested * schedulerapi.MaxPriority) / capacity
}
