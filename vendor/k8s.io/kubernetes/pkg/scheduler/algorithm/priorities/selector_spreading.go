package priorities

import (
	"fmt"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/scheduler/algorithm"
	schedulerapi "k8s.io/kubernetes/pkg/scheduler/api"
	schedulercache "k8s.io/kubernetes/pkg/scheduler/cache"
	utilnode "k8s.io/kubernetes/pkg/util/node"
)

const zoneWeighting float64 = 2.0 / 3.0

type SelectorSpread struct {
	serviceLister     algorithm.ServiceLister
	controllerLister  algorithm.ControllerLister
	replicaSetLister  algorithm.ReplicaSetLister
	statefulSetLister algorithm.StatefulSetLister
}

func NewSelectorSpreadPriority(serviceLister algorithm.ServiceLister, controllerLister algorithm.ControllerLister, replicaSetLister algorithm.ReplicaSetLister, statefulSetLister algorithm.StatefulSetLister) (algorithm.PriorityMapFunction, algorithm.PriorityReduceFunction) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	selectorSpread := &SelectorSpread{serviceLister: serviceLister, controllerLister: controllerLister, replicaSetLister: replicaSetLister, statefulSetLister: statefulSetLister}
	return selectorSpread.CalculateSpreadPriorityMap, selectorSpread.CalculateSpreadPriorityReduce
}
func (s *SelectorSpread) CalculateSpreadPriorityMap(pod *v1.Pod, meta interface{}, nodeInfo *schedulercache.NodeInfo) (schedulerapi.HostPriority, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var selectors []labels.Selector
	node := nodeInfo.Node()
	if node == nil {
		return schedulerapi.HostPriority{}, fmt.Errorf("node not found")
	}
	priorityMeta, ok := meta.(*priorityMetadata)
	if ok {
		selectors = priorityMeta.podSelectors
	} else {
		selectors = getSelectors(pod, s.serviceLister, s.controllerLister, s.replicaSetLister, s.statefulSetLister)
	}
	if len(selectors) == 0 {
		return schedulerapi.HostPriority{Host: node.Name, Score: int(0)}, nil
	}
	count := int(0)
	for _, nodePod := range nodeInfo.Pods() {
		if pod.Namespace != nodePod.Namespace {
			continue
		}
		if nodePod.DeletionTimestamp != nil {
			klog.V(4).Infof("skipping pending-deleted pod: %s/%s", nodePod.Namespace, nodePod.Name)
			continue
		}
		for _, selector := range selectors {
			if selector.Matches(labels.Set(nodePod.ObjectMeta.Labels)) {
				count++
				break
			}
		}
	}
	return schedulerapi.HostPriority{Host: node.Name, Score: int(count)}, nil
}
func (s *SelectorSpread) CalculateSpreadPriorityReduce(pod *v1.Pod, meta interface{}, nodeNameToInfo map[string]*schedulercache.NodeInfo, result schedulerapi.HostPriorityList) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	countsByZone := make(map[string]int, 10)
	maxCountByZone := int(0)
	maxCountByNodeName := int(0)
	for i := range result {
		if result[i].Score > maxCountByNodeName {
			maxCountByNodeName = result[i].Score
		}
		zoneID := utilnode.GetZoneKey(nodeNameToInfo[result[i].Host].Node())
		if zoneID == "" {
			continue
		}
		countsByZone[zoneID] += result[i].Score
	}
	for zoneID := range countsByZone {
		if countsByZone[zoneID] > maxCountByZone {
			maxCountByZone = countsByZone[zoneID]
		}
	}
	haveZones := len(countsByZone) != 0
	maxCountByNodeNameFloat64 := float64(maxCountByNodeName)
	maxCountByZoneFloat64 := float64(maxCountByZone)
	MaxPriorityFloat64 := float64(schedulerapi.MaxPriority)
	for i := range result {
		fScore := MaxPriorityFloat64
		if maxCountByNodeName > 0 {
			fScore = MaxPriorityFloat64 * (float64(maxCountByNodeName-result[i].Score) / maxCountByNodeNameFloat64)
		}
		if haveZones {
			zoneID := utilnode.GetZoneKey(nodeNameToInfo[result[i].Host].Node())
			if zoneID != "" {
				zoneScore := MaxPriorityFloat64
				if maxCountByZone > 0 {
					zoneScore = MaxPriorityFloat64 * (float64(maxCountByZone-countsByZone[zoneID]) / maxCountByZoneFloat64)
				}
				fScore = (fScore * (1.0 - zoneWeighting)) + (zoneWeighting * zoneScore)
			}
		}
		result[i].Score = int(fScore)
		if klog.V(10) {
			klog.Infof("%v -> %v: SelectorSpreadPriority, Score: (%d)", pod.Name, result[i].Host, int(fScore))
		}
	}
	return nil
}

type ServiceAntiAffinity struct {
	podLister     algorithm.PodLister
	serviceLister algorithm.ServiceLister
	label         string
}

func NewServiceAntiAffinityPriority(podLister algorithm.PodLister, serviceLister algorithm.ServiceLister, label string) (algorithm.PriorityMapFunction, algorithm.PriorityReduceFunction) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	antiAffinity := &ServiceAntiAffinity{podLister: podLister, serviceLister: serviceLister, label: label}
	return antiAffinity.CalculateAntiAffinityPriorityMap, antiAffinity.CalculateAntiAffinityPriorityReduce
}
func (s *ServiceAntiAffinity) getNodeClassificationByLabels(nodes []*v1.Node) (map[string]string, []string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	labeledNodes := map[string]string{}
	nonLabeledNodes := []string{}
	for _, node := range nodes {
		if labels.Set(node.Labels).Has(s.label) {
			label := labels.Set(node.Labels).Get(s.label)
			labeledNodes[node.Name] = label
		} else {
			nonLabeledNodes = append(nonLabeledNodes, node.Name)
		}
	}
	return labeledNodes, nonLabeledNodes
}
func filteredPod(namespace string, selector labels.Selector, nodeInfo *schedulercache.NodeInfo) (pods []*v1.Pod) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if nodeInfo.Pods() == nil || len(nodeInfo.Pods()) == 0 || selector == nil {
		return []*v1.Pod{}
	}
	for _, pod := range nodeInfo.Pods() {
		if namespace == pod.Namespace && pod.DeletionTimestamp == nil && selector.Matches(labels.Set(pod.Labels)) {
			pods = append(pods, pod)
		}
	}
	return
}
func (s *ServiceAntiAffinity) CalculateAntiAffinityPriorityMap(pod *v1.Pod, meta interface{}, nodeInfo *schedulercache.NodeInfo) (schedulerapi.HostPriority, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var firstServiceSelector labels.Selector
	node := nodeInfo.Node()
	if node == nil {
		return schedulerapi.HostPriority{}, fmt.Errorf("node not found")
	}
	priorityMeta, ok := meta.(*priorityMetadata)
	if ok {
		firstServiceSelector = priorityMeta.podFirstServiceSelector
	} else {
		firstServiceSelector = getFirstServiceSelector(pod, s.serviceLister)
	}
	matchedPodsOfNode := filteredPod(pod.Namespace, firstServiceSelector, nodeInfo)
	return schedulerapi.HostPriority{Host: node.Name, Score: int(len(matchedPodsOfNode))}, nil
}
func (s *ServiceAntiAffinity) CalculateAntiAffinityPriorityReduce(pod *v1.Pod, meta interface{}, nodeNameToInfo map[string]*schedulercache.NodeInfo, result schedulerapi.HostPriorityList) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var numServicePods int
	var label string
	podCounts := map[string]int{}
	labelNodesStatus := map[string]string{}
	maxPriorityFloat64 := float64(schedulerapi.MaxPriority)
	for _, hostPriority := range result {
		numServicePods += hostPriority.Score
		if !labels.Set(nodeNameToInfo[hostPriority.Host].Node().Labels).Has(s.label) {
			continue
		}
		label = labels.Set(nodeNameToInfo[hostPriority.Host].Node().Labels).Get(s.label)
		labelNodesStatus[hostPriority.Host] = label
		podCounts[label] += hostPriority.Score
	}
	for i, hostPriority := range result {
		label, ok := labelNodesStatus[hostPriority.Host]
		if !ok {
			result[i].Host = hostPriority.Host
			result[i].Score = int(0)
			continue
		}
		fScore := maxPriorityFloat64
		if numServicePods > 0 {
			fScore = maxPriorityFloat64 * (float64(numServicePods-podCounts[label]) / float64(numServicePods))
		}
		result[i].Host = hostPriority.Host
		result[i].Score = int(fScore)
	}
	return nil
}
