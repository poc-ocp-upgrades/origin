package priorities

import (
 "fmt"
 schedulerapi "k8s.io/kubernetes/pkg/scheduler/api"
 schedulercache "k8s.io/kubernetes/pkg/scheduler/cache"
)

type FunctionShape []FunctionShapePoint
type FunctionShapePoint struct {
 Utilization int64
 Score       int64
}

var (
 defaultFunctionShape, _ = NewFunctionShape([]FunctionShapePoint{{0, 10}, {100, 0}})
)

const (
 minUtilization = 0
 maxUtilization = 100
 minScore       = 0
 maxScore       = schedulerapi.MaxPriority
)

func NewFunctionShape(points []FunctionShapePoint) (FunctionShape, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 n := len(points)
 if n == 0 {
  return nil, fmt.Errorf("at least one point must be specified")
 }
 for i := 1; i < n; i++ {
  if points[i-1].Utilization >= points[i].Utilization {
   return nil, fmt.Errorf("utilization values must be sorted. Utilization[%d]==%d >= Utilization[%d]==%d", i-1, points[i-1].Utilization, i, points[i].Utilization)
  }
 }
 for i, point := range points {
  if point.Utilization < minUtilization {
   return nil, fmt.Errorf("utilization values must not be less than %d. Utilization[%d]==%d", minUtilization, i, point.Utilization)
  }
  if point.Utilization > maxUtilization {
   return nil, fmt.Errorf("utilization values must not be greater than %d. Utilization[%d]==%d", maxUtilization, i, point.Utilization)
  }
  if point.Score < minScore {
   return nil, fmt.Errorf("score values must not be less than %d. Score[%d]==%d", minScore, i, point.Score)
  }
  if point.Score > maxScore {
   return nil, fmt.Errorf("score valuses not be greater than %d. Score[%d]==%d", maxScore, i, point.Score)
  }
 }
 pointsCopy := make(FunctionShape, n)
 copy(pointsCopy, points)
 return pointsCopy, nil
}
func RequestedToCapacityRatioResourceAllocationPriorityDefault() *ResourceAllocationPriority {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return RequestedToCapacityRatioResourceAllocationPriority(defaultFunctionShape)
}
func RequestedToCapacityRatioResourceAllocationPriority(scoringFunctionShape FunctionShape) *ResourceAllocationPriority {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &ResourceAllocationPriority{"RequestedToCapacityRatioResourceAllocationPriority", buildRequestedToCapacityRatioScorerFunction(scoringFunctionShape)}
}
func buildRequestedToCapacityRatioScorerFunction(scoringFunctionShape FunctionShape) func(*schedulercache.Resource, *schedulercache.Resource, bool, int, int) int64 {
 _logClusterCodePath()
 defer _logClusterCodePath()
 rawScoringFunction := buildBrokenLinearFunction(scoringFunctionShape)
 resourceScoringFunction := func(requested, capacity int64) int64 {
  if capacity == 0 || requested > capacity {
   return rawScoringFunction(maxUtilization)
  }
  return rawScoringFunction(maxUtilization - (capacity-requested)*maxUtilization/capacity)
 }
 return func(requested, allocable *schedulercache.Resource, includeVolumes bool, requestedVolumes int, allocatableVolumes int) int64 {
  cpuScore := resourceScoringFunction(requested.MilliCPU, allocable.MilliCPU)
  memoryScore := resourceScoringFunction(requested.Memory, allocable.Memory)
  return (cpuScore + memoryScore) / 2
 }
}
func buildBrokenLinearFunction(shape FunctionShape) func(int64) int64 {
 _logClusterCodePath()
 defer _logClusterCodePath()
 n := len(shape)
 return func(p int64) int64 {
  for i := 0; i < n; i++ {
   if p <= shape[i].Utilization {
    if i == 0 {
     return shape[0].Score
    }
    return shape[i-1].Score + (shape[i].Score-shape[i-1].Score)*(p-shape[i-1].Utilization)/(shape[i].Utilization-shape[i-1].Utilization)
   }
  }
  return shape[n-1].Score
 }
}
