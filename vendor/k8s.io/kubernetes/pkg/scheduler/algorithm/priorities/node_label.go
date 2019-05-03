package priorities

import (
 "fmt"
 "k8s.io/api/core/v1"
 "k8s.io/apimachinery/pkg/labels"
 "k8s.io/kubernetes/pkg/scheduler/algorithm"
 schedulerapi "k8s.io/kubernetes/pkg/scheduler/api"
 schedulercache "k8s.io/kubernetes/pkg/scheduler/cache"
)

type NodeLabelPrioritizer struct {
 label    string
 presence bool
}

func NewNodeLabelPriority(label string, presence bool) (algorithm.PriorityMapFunction, algorithm.PriorityReduceFunction) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 labelPrioritizer := &NodeLabelPrioritizer{label: label, presence: presence}
 return labelPrioritizer.CalculateNodeLabelPriorityMap, nil
}
func (n *NodeLabelPrioritizer) CalculateNodeLabelPriorityMap(pod *v1.Pod, meta interface{}, nodeInfo *schedulercache.NodeInfo) (schedulerapi.HostPriority, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 node := nodeInfo.Node()
 if node == nil {
  return schedulerapi.HostPriority{}, fmt.Errorf("node not found")
 }
 exists := labels.Set(node.Labels).Has(n.label)
 score := 0
 if (exists && n.presence) || (!exists && !n.presence) {
  score = schedulerapi.MaxPriority
 }
 return schedulerapi.HostPriority{Host: node.Name, Score: score}, nil
}
