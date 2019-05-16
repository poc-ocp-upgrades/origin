package priorities

import (
	schedulerapi "k8s.io/kubernetes/pkg/scheduler/api"
	schedulercache "k8s.io/kubernetes/pkg/scheduler/cache"
)

var (
	leastResourcePriority     = &ResourceAllocationPriority{"LeastResourceAllocation", leastResourceScorer}
	LeastRequestedPriorityMap = leastResourcePriority.PriorityMap
)

func leastResourceScorer(requested, allocable *schedulercache.Resource, includeVolumes bool, requestedVolumes int, allocatableVolumes int) int64 {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return (leastRequestedScore(requested.MilliCPU, allocable.MilliCPU) + leastRequestedScore(requested.Memory, allocable.Memory)) / 2
}
func leastRequestedScore(requested, capacity int64) int64 {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if capacity == 0 {
		return 0
	}
	if requested > capacity {
		return 0
	}
	return ((capacity - requested) * int64(schedulerapi.MaxPriority)) / capacity
}
