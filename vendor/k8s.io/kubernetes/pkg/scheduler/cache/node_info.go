package cache

import (
	"errors"
	"fmt"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/klog"
	v1helper "k8s.io/kubernetes/pkg/apis/core/v1/helper"
	priorityutil "k8s.io/kubernetes/pkg/scheduler/algorithm/priorities/util"
	"sync"
	"sync/atomic"
)

var (
	emptyResource = Resource{}
	generation    int64
)

type ImageStateSummary struct {
	Size     int64
	NumNodes int
}
type NodeInfo struct {
	node                    *v1.Node
	pods                    []*v1.Pod
	podsWithAffinity        []*v1.Pod
	usedPorts               HostPortInfo
	requestedResource       *Resource
	nonzeroRequest          *Resource
	allocatableResource     *Resource
	taints                  []v1.Taint
	taintsErr               error
	imageStates             map[string]*ImageStateSummary
	TransientInfo           *TransientSchedulerInfo
	memoryPressureCondition v1.ConditionStatus
	diskPressureCondition   v1.ConditionStatus
	pidPressureCondition    v1.ConditionStatus
	generation              int64
}

func initializeNodeTransientInfo() nodeTransientInfo {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nodeTransientInfo{AllocatableVolumesCount: 0, RequestedVolumes: 0}
}
func nextGeneration() int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return atomic.AddInt64(&generation, 1)
}

type nodeTransientInfo struct {
	AllocatableVolumesCount int
	RequestedVolumes        int
}
type TransientSchedulerInfo struct {
	TransientLock sync.Mutex
	TransNodeInfo nodeTransientInfo
}

func NewTransientSchedulerInfo() *TransientSchedulerInfo {
	_logClusterCodePath()
	defer _logClusterCodePath()
	tsi := &TransientSchedulerInfo{TransNodeInfo: initializeNodeTransientInfo()}
	return tsi
}
func (transientSchedInfo *TransientSchedulerInfo) ResetTransientSchedulerInfo() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	transientSchedInfo.TransientLock.Lock()
	defer transientSchedInfo.TransientLock.Unlock()
	transientSchedInfo.TransNodeInfo.AllocatableVolumesCount = 0
	transientSchedInfo.TransNodeInfo.RequestedVolumes = 0
}

type Resource struct {
	MilliCPU         int64
	Memory           int64
	EphemeralStorage int64
	AllowedPodNumber int
	ScalarResources  map[v1.ResourceName]int64
}

func NewResource(rl v1.ResourceList) *Resource {
	_logClusterCodePath()
	defer _logClusterCodePath()
	r := &Resource{}
	r.Add(rl)
	return r
}
func (r *Resource) Add(rl v1.ResourceList) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if r == nil {
		return
	}
	for rName, rQuant := range rl {
		switch rName {
		case v1.ResourceCPU:
			r.MilliCPU += rQuant.MilliValue()
		case v1.ResourceMemory:
			r.Memory += rQuant.Value()
		case v1.ResourcePods:
			r.AllowedPodNumber += int(rQuant.Value())
		case v1.ResourceEphemeralStorage:
			r.EphemeralStorage += rQuant.Value()
		default:
			if v1helper.IsScalarResourceName(rName) {
				r.AddScalar(rName, rQuant.Value())
			}
		}
	}
}
func (r *Resource) ResourceList() v1.ResourceList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result := v1.ResourceList{v1.ResourceCPU: *resource.NewMilliQuantity(r.MilliCPU, resource.DecimalSI), v1.ResourceMemory: *resource.NewQuantity(r.Memory, resource.BinarySI), v1.ResourcePods: *resource.NewQuantity(int64(r.AllowedPodNumber), resource.BinarySI), v1.ResourceEphemeralStorage: *resource.NewQuantity(r.EphemeralStorage, resource.BinarySI)}
	for rName, rQuant := range r.ScalarResources {
		if v1helper.IsHugePageResourceName(rName) {
			result[rName] = *resource.NewQuantity(rQuant, resource.BinarySI)
		} else {
			result[rName] = *resource.NewQuantity(rQuant, resource.DecimalSI)
		}
	}
	return result
}
func (r *Resource) Clone() *Resource {
	_logClusterCodePath()
	defer _logClusterCodePath()
	res := &Resource{MilliCPU: r.MilliCPU, Memory: r.Memory, AllowedPodNumber: r.AllowedPodNumber, EphemeralStorage: r.EphemeralStorage}
	if r.ScalarResources != nil {
		res.ScalarResources = make(map[v1.ResourceName]int64)
		for k, v := range r.ScalarResources {
			res.ScalarResources[k] = v
		}
	}
	return res
}
func (r *Resource) AddScalar(name v1.ResourceName, quantity int64) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	r.SetScalar(name, r.ScalarResources[name]+quantity)
}
func (r *Resource) SetScalar(name v1.ResourceName, quantity int64) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if r.ScalarResources == nil {
		r.ScalarResources = map[v1.ResourceName]int64{}
	}
	r.ScalarResources[name] = quantity
}
func (r *Resource) SetMaxResource(rl v1.ResourceList) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if r == nil {
		return
	}
	for rName, rQuantity := range rl {
		switch rName {
		case v1.ResourceMemory:
			if mem := rQuantity.Value(); mem > r.Memory {
				r.Memory = mem
			}
		case v1.ResourceCPU:
			if cpu := rQuantity.MilliValue(); cpu > r.MilliCPU {
				r.MilliCPU = cpu
			}
		case v1.ResourceEphemeralStorage:
			if ephemeralStorage := rQuantity.Value(); ephemeralStorage > r.EphemeralStorage {
				r.EphemeralStorage = ephemeralStorage
			}
		default:
			if v1helper.IsScalarResourceName(rName) {
				value := rQuantity.Value()
				if value > r.ScalarResources[rName] {
					r.SetScalar(rName, value)
				}
			}
		}
	}
}
func NewNodeInfo(pods ...*v1.Pod) *NodeInfo {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ni := &NodeInfo{requestedResource: &Resource{}, nonzeroRequest: &Resource{}, allocatableResource: &Resource{}, TransientInfo: NewTransientSchedulerInfo(), generation: nextGeneration(), usedPorts: make(HostPortInfo), imageStates: make(map[string]*ImageStateSummary)}
	for _, pod := range pods {
		ni.AddPod(pod)
	}
	return ni
}
func (n *NodeInfo) Node() *v1.Node {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if n == nil {
		return nil
	}
	return n.node
}
func (n *NodeInfo) Pods() []*v1.Pod {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if n == nil {
		return nil
	}
	return n.pods
}
func (n *NodeInfo) SetPods(pods []*v1.Pod) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	n.pods = pods
}
func (n *NodeInfo) UsedPorts() HostPortInfo {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if n == nil {
		return nil
	}
	return n.usedPorts
}
func (n *NodeInfo) SetUsedPorts(newUsedPorts HostPortInfo) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	n.usedPorts = newUsedPorts
}
func (n *NodeInfo) ImageStates() map[string]*ImageStateSummary {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if n == nil {
		return nil
	}
	return n.imageStates
}
func (n *NodeInfo) SetImageStates(newImageStates map[string]*ImageStateSummary) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	n.imageStates = newImageStates
}
func (n *NodeInfo) PodsWithAffinity() []*v1.Pod {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if n == nil {
		return nil
	}
	return n.podsWithAffinity
}
func (n *NodeInfo) AllowedPodNumber() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if n == nil || n.allocatableResource == nil {
		return 0
	}
	return n.allocatableResource.AllowedPodNumber
}
func (n *NodeInfo) Taints() ([]v1.Taint, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if n == nil {
		return nil, nil
	}
	return n.taints, n.taintsErr
}
func (n *NodeInfo) SetTaints(newTaints []v1.Taint) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	n.taints = newTaints
}
func (n *NodeInfo) MemoryPressureCondition() v1.ConditionStatus {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if n == nil {
		return v1.ConditionUnknown
	}
	return n.memoryPressureCondition
}
func (n *NodeInfo) DiskPressureCondition() v1.ConditionStatus {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if n == nil {
		return v1.ConditionUnknown
	}
	return n.diskPressureCondition
}
func (n *NodeInfo) PIDPressureCondition() v1.ConditionStatus {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if n == nil {
		return v1.ConditionUnknown
	}
	return n.pidPressureCondition
}
func (n *NodeInfo) RequestedResource() Resource {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if n == nil {
		return emptyResource
	}
	return *n.requestedResource
}
func (n *NodeInfo) SetRequestedResource(newResource *Resource) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	n.requestedResource = newResource
}
func (n *NodeInfo) NonZeroRequest() Resource {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if n == nil {
		return emptyResource
	}
	return *n.nonzeroRequest
}
func (n *NodeInfo) SetNonZeroRequest(newResource *Resource) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	n.nonzeroRequest = newResource
}
func (n *NodeInfo) AllocatableResource() Resource {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if n == nil {
		return emptyResource
	}
	return *n.allocatableResource
}
func (n *NodeInfo) SetAllocatableResource(allocatableResource *Resource) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	n.allocatableResource = allocatableResource
	n.generation = nextGeneration()
}
func (n *NodeInfo) GetGeneration() int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if n == nil {
		return 0
	}
	return n.generation
}
func (n *NodeInfo) SetGeneration(newGeneration int64) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	n.generation = newGeneration
}
func (n *NodeInfo) Clone() *NodeInfo {
	_logClusterCodePath()
	defer _logClusterCodePath()
	clone := &NodeInfo{node: n.node, requestedResource: n.requestedResource.Clone(), nonzeroRequest: n.nonzeroRequest.Clone(), allocatableResource: n.allocatableResource.Clone(), taintsErr: n.taintsErr, TransientInfo: n.TransientInfo, memoryPressureCondition: n.memoryPressureCondition, diskPressureCondition: n.diskPressureCondition, pidPressureCondition: n.pidPressureCondition, usedPorts: make(HostPortInfo), imageStates: n.imageStates, generation: n.generation}
	if len(n.pods) > 0 {
		clone.pods = append([]*v1.Pod(nil), n.pods...)
	}
	if len(n.usedPorts) > 0 {
		for ip, portMap := range n.usedPorts {
			clone.usedPorts[ip] = make(map[ProtocolPort]struct{})
			for protocolPort, v := range portMap {
				clone.usedPorts[ip][protocolPort] = v
			}
		}
	}
	if len(n.podsWithAffinity) > 0 {
		clone.podsWithAffinity = append([]*v1.Pod(nil), n.podsWithAffinity...)
	}
	if len(n.taints) > 0 {
		clone.taints = append([]v1.Taint(nil), n.taints...)
	}
	return clone
}
func (n *NodeInfo) VolumeLimits() map[v1.ResourceName]int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	volumeLimits := map[v1.ResourceName]int64{}
	for k, v := range n.AllocatableResource().ScalarResources {
		if v1helper.IsAttachableVolumeResourceName(k) {
			volumeLimits[k] = v
		}
	}
	return volumeLimits
}
func (n *NodeInfo) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	podKeys := make([]string, len(n.pods))
	for i, pod := range n.pods {
		podKeys[i] = pod.Name
	}
	return fmt.Sprintf("&NodeInfo{Pods:%v, RequestedResource:%#v, NonZeroRequest: %#v, UsedPort: %#v, AllocatableResource:%#v}", podKeys, n.requestedResource, n.nonzeroRequest, n.usedPorts, n.allocatableResource)
}
func hasPodAffinityConstraints(pod *v1.Pod) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	affinity := pod.Spec.Affinity
	return affinity != nil && (affinity.PodAffinity != nil || affinity.PodAntiAffinity != nil)
}
func (n *NodeInfo) AddPod(pod *v1.Pod) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	res, non0CPU, non0Mem := calculateResource(pod)
	n.requestedResource.MilliCPU += res.MilliCPU
	n.requestedResource.Memory += res.Memory
	n.requestedResource.EphemeralStorage += res.EphemeralStorage
	if n.requestedResource.ScalarResources == nil && len(res.ScalarResources) > 0 {
		n.requestedResource.ScalarResources = map[v1.ResourceName]int64{}
	}
	for rName, rQuant := range res.ScalarResources {
		n.requestedResource.ScalarResources[rName] += rQuant
	}
	n.nonzeroRequest.MilliCPU += non0CPU
	n.nonzeroRequest.Memory += non0Mem
	n.pods = append(n.pods, pod)
	if hasPodAffinityConstraints(pod) {
		n.podsWithAffinity = append(n.podsWithAffinity, pod)
	}
	n.UpdateUsedPorts(pod, true)
	n.generation = nextGeneration()
}
func (n *NodeInfo) RemovePod(pod *v1.Pod) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	k1, err := GetPodKey(pod)
	if err != nil {
		return err
	}
	for i := range n.podsWithAffinity {
		k2, err := GetPodKey(n.podsWithAffinity[i])
		if err != nil {
			klog.Errorf("Cannot get pod key, err: %v", err)
			continue
		}
		if k1 == k2 {
			n.podsWithAffinity[i] = n.podsWithAffinity[len(n.podsWithAffinity)-1]
			n.podsWithAffinity = n.podsWithAffinity[:len(n.podsWithAffinity)-1]
			break
		}
	}
	for i := range n.pods {
		k2, err := GetPodKey(n.pods[i])
		if err != nil {
			klog.Errorf("Cannot get pod key, err: %v", err)
			continue
		}
		if k1 == k2 {
			n.pods[i] = n.pods[len(n.pods)-1]
			n.pods = n.pods[:len(n.pods)-1]
			res, non0CPU, non0Mem := calculateResource(pod)
			n.requestedResource.MilliCPU -= res.MilliCPU
			n.requestedResource.Memory -= res.Memory
			n.requestedResource.EphemeralStorage -= res.EphemeralStorage
			if len(res.ScalarResources) > 0 && n.requestedResource.ScalarResources == nil {
				n.requestedResource.ScalarResources = map[v1.ResourceName]int64{}
			}
			for rName, rQuant := range res.ScalarResources {
				n.requestedResource.ScalarResources[rName] -= rQuant
			}
			n.nonzeroRequest.MilliCPU -= non0CPU
			n.nonzeroRequest.Memory -= non0Mem
			n.UpdateUsedPorts(pod, false)
			n.generation = nextGeneration()
			return nil
		}
	}
	return fmt.Errorf("no corresponding pod %s in pods of node %s", pod.Name, n.node.Name)
}
func calculateResource(pod *v1.Pod) (res Resource, non0CPU int64, non0Mem int64) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	resPtr := &res
	for _, c := range pod.Spec.Containers {
		resPtr.Add(c.Resources.Requests)
		non0CPUReq, non0MemReq := priorityutil.GetNonzeroRequests(&c.Resources.Requests)
		non0CPU += non0CPUReq
		non0Mem += non0MemReq
	}
	return
}
func (n *NodeInfo) UpdateUsedPorts(pod *v1.Pod, add bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for j := range pod.Spec.Containers {
		container := &pod.Spec.Containers[j]
		for k := range container.Ports {
			podPort := &container.Ports[k]
			if add {
				n.usedPorts.Add(podPort.HostIP, string(podPort.Protocol), podPort.HostPort)
			} else {
				n.usedPorts.Remove(podPort.HostIP, string(podPort.Protocol), podPort.HostPort)
			}
		}
	}
}
func (n *NodeInfo) SetNode(node *v1.Node) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	n.node = node
	n.allocatableResource = NewResource(node.Status.Allocatable)
	n.taints = node.Spec.Taints
	for i := range node.Status.Conditions {
		cond := &node.Status.Conditions[i]
		switch cond.Type {
		case v1.NodeMemoryPressure:
			n.memoryPressureCondition = cond.Status
		case v1.NodeDiskPressure:
			n.diskPressureCondition = cond.Status
		case v1.NodePIDPressure:
			n.pidPressureCondition = cond.Status
		default:
		}
	}
	n.TransientInfo = NewTransientSchedulerInfo()
	n.generation = nextGeneration()
	return nil
}
func (n *NodeInfo) RemoveNode(node *v1.Node) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	n.node = nil
	n.allocatableResource = &Resource{}
	n.taints, n.taintsErr = nil, nil
	n.memoryPressureCondition = v1.ConditionUnknown
	n.diskPressureCondition = v1.ConditionUnknown
	n.pidPressureCondition = v1.ConditionUnknown
	n.imageStates = make(map[string]*ImageStateSummary)
	n.generation = nextGeneration()
	return nil
}
func (n *NodeInfo) FilterOutPods(pods []*v1.Pod) []*v1.Pod {
	_logClusterCodePath()
	defer _logClusterCodePath()
	node := n.Node()
	if node == nil {
		return pods
	}
	filtered := make([]*v1.Pod, 0, len(pods))
	for _, p := range pods {
		if p.Spec.NodeName != node.Name {
			filtered = append(filtered, p)
			continue
		}
		podKey, _ := GetPodKey(p)
		for _, np := range n.Pods() {
			npodkey, _ := GetPodKey(np)
			if npodkey == podKey {
				filtered = append(filtered, p)
				break
			}
		}
	}
	return filtered
}
func GetPodKey(pod *v1.Pod) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	uid := string(pod.UID)
	if len(uid) == 0 {
		return "", errors.New("Cannot get cache key for pod with empty UID")
	}
	return uid, nil
}
func (n *NodeInfo) Filter(pod *v1.Pod) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if pod.Spec.NodeName != n.node.Name {
		return true
	}
	for _, p := range n.pods {
		if p.Name == pod.Name && p.Namespace == pod.Namespace {
			return true
		}
	}
	return false
}
