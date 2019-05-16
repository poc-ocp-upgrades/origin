package priorities

import (
	goformat "fmt"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	"k8s.io/kubernetes/pkg/features"
	schedulerapi "k8s.io/kubernetes/pkg/scheduler/api"
	schedulercache "k8s.io/kubernetes/pkg/scheduler/cache"
	"math"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

var (
	balancedResourcePriority      = &ResourceAllocationPriority{"BalancedResourceAllocation", balancedResourceScorer}
	BalancedResourceAllocationMap = balancedResourcePriority.PriorityMap
)

func balancedResourceScorer(requested, allocable *schedulercache.Resource, includeVolumes bool, requestedVolumes int, allocatableVolumes int) int64 {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	cpuFraction := fractionOfCapacity(requested.MilliCPU, allocable.MilliCPU)
	memoryFraction := fractionOfCapacity(requested.Memory, allocable.Memory)
	if includeVolumes && utilfeature.DefaultFeatureGate.Enabled(features.BalanceAttachedNodeVolumes) && allocatableVolumes > 0 {
		volumeFraction := float64(requestedVolumes) / float64(allocatableVolumes)
		if cpuFraction >= 1 || memoryFraction >= 1 || volumeFraction >= 1 {
			return 0
		}
		mean := (cpuFraction + memoryFraction + volumeFraction) / float64(3)
		variance := float64((((cpuFraction - mean) * (cpuFraction - mean)) + ((memoryFraction - mean) * (memoryFraction - mean)) + ((volumeFraction - mean) * (volumeFraction - mean))) / float64(3))
		return int64((1 - variance) * float64(schedulerapi.MaxPriority))
	}
	if cpuFraction >= 1 || memoryFraction >= 1 {
		return 0
	}
	diff := math.Abs(cpuFraction - memoryFraction)
	return int64((1 - diff) * float64(schedulerapi.MaxPriority))
}
func fractionOfCapacity(requested, capacity int64) float64 {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if capacity == 0 {
		return 1
	}
	return float64(requested) / float64(capacity)
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
