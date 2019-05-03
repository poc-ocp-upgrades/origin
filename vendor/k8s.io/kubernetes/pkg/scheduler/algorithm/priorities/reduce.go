package priorities

import (
 "k8s.io/api/core/v1"
 "k8s.io/kubernetes/pkg/scheduler/algorithm"
 schedulerapi "k8s.io/kubernetes/pkg/scheduler/api"
 schedulercache "k8s.io/kubernetes/pkg/scheduler/cache"
)

func NormalizeReduce(maxPriority int, reverse bool) algorithm.PriorityReduceFunction {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return func(_ *v1.Pod, _ interface{}, _ map[string]*schedulercache.NodeInfo, result schedulerapi.HostPriorityList) error {
  var maxCount int
  for i := range result {
   if result[i].Score > maxCount {
    maxCount = result[i].Score
   }
  }
  if maxCount == 0 {
   if reverse {
    for i := range result {
     result[i].Score = maxPriority
    }
   }
   return nil
  }
  for i := range result {
   score := result[i].Score
   score = maxPriority * score / maxCount
   if reverse {
    score = maxPriority - score
   }
   result[i].Score = score
  }
  return nil
 }
}
