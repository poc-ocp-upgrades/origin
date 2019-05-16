package priorities

import (
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kubernetes/pkg/scheduler/algorithm"
	schedulerapi "k8s.io/kubernetes/pkg/scheduler/api"
	schedulercache "k8s.io/kubernetes/pkg/scheduler/cache"
)

func makeNode(node string, milliCPU, memory int64) *v1.Node {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &v1.Node{ObjectMeta: metav1.ObjectMeta{Name: node}, Status: v1.NodeStatus{Capacity: v1.ResourceList{v1.ResourceCPU: *resource.NewMilliQuantity(milliCPU, resource.DecimalSI), v1.ResourceMemory: *resource.NewQuantity(memory, resource.BinarySI)}, Allocatable: v1.ResourceList{v1.ResourceCPU: *resource.NewMilliQuantity(milliCPU, resource.DecimalSI), v1.ResourceMemory: *resource.NewQuantity(memory, resource.BinarySI)}}}
}
func priorityFunction(mapFn algorithm.PriorityMapFunction, reduceFn algorithm.PriorityReduceFunction, metaData interface{}) algorithm.PriorityFunction {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return func(pod *v1.Pod, nodeNameToInfo map[string]*schedulercache.NodeInfo, nodes []*v1.Node) (schedulerapi.HostPriorityList, error) {
		result := make(schedulerapi.HostPriorityList, 0, len(nodes))
		for i := range nodes {
			hostResult, err := mapFn(pod, metaData, nodeNameToInfo[nodes[i].Name])
			if err != nil {
				return nil, err
			}
			result = append(result, hostResult)
		}
		if reduceFn != nil {
			if err := reduceFn(pod, metaData, nodeNameToInfo, result); err != nil {
				return nil, err
			}
		}
		return result, nil
	}
}
