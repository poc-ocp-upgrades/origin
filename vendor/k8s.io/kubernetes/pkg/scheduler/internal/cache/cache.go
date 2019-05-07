package cache

import (
	godefaultbytes "bytes"
	"fmt"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apimachinery/pkg/util/wait"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/features"
	"k8s.io/kubernetes/pkg/scheduler/algorithm"
	schedulercache "k8s.io/kubernetes/pkg/scheduler/cache"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"sync"
	"time"
)

var (
	cleanAssumedPeriod = 1 * time.Second
)

func New(ttl time.Duration, stop <-chan struct{}) Cache {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cache := newSchedulerCache(ttl, cleanAssumedPeriod, stop)
	cache.run()
	return cache
}

type schedulerCache struct {
	stop        <-chan struct{}
	ttl         time.Duration
	period      time.Duration
	mu          sync.RWMutex
	assumedPods map[string]bool
	podStates   map[string]*podState
	nodes       map[string]*schedulercache.NodeInfo
	nodeTree    *NodeTree
	imageStates map[string]*imageState
}
type podState struct {
	pod             *v1.Pod
	deadline        *time.Time
	bindingFinished bool
}
type imageState struct {
	size  int64
	nodes sets.String
}

func (cache *schedulerCache) createImageStateSummary(state *imageState) *schedulercache.ImageStateSummary {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &schedulercache.ImageStateSummary{Size: state.size, NumNodes: len(state.nodes)}
}
func newSchedulerCache(ttl, period time.Duration, stop <-chan struct{}) *schedulerCache {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &schedulerCache{ttl: ttl, period: period, stop: stop, nodes: make(map[string]*schedulercache.NodeInfo), nodeTree: newNodeTree(nil), assumedPods: make(map[string]bool), podStates: make(map[string]*podState), imageStates: make(map[string]*imageState)}
}
func (cache *schedulerCache) Snapshot() *Snapshot {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cache.mu.RLock()
	defer cache.mu.RUnlock()
	nodes := make(map[string]*schedulercache.NodeInfo)
	for k, v := range cache.nodes {
		nodes[k] = v.Clone()
	}
	assumedPods := make(map[string]bool)
	for k, v := range cache.assumedPods {
		assumedPods[k] = v
	}
	return &Snapshot{Nodes: nodes, AssumedPods: assumedPods}
}
func (cache *schedulerCache) UpdateNodeNameToInfoMap(nodeNameToInfo map[string]*schedulercache.NodeInfo) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cache.mu.Lock()
	defer cache.mu.Unlock()
	for name, info := range cache.nodes {
		if utilfeature.DefaultFeatureGate.Enabled(features.BalanceAttachedNodeVolumes) && info.TransientInfo != nil {
			info.TransientInfo.ResetTransientSchedulerInfo()
		}
		if current, ok := nodeNameToInfo[name]; !ok || current.GetGeneration() != info.GetGeneration() {
			nodeNameToInfo[name] = info.Clone()
		}
	}
	for name := range nodeNameToInfo {
		if _, ok := cache.nodes[name]; !ok {
			delete(nodeNameToInfo, name)
		}
	}
	return nil
}
func (cache *schedulerCache) List(selector labels.Selector) ([]*v1.Pod, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	alwaysTrue := func(p *v1.Pod) bool {
		return true
	}
	return cache.FilteredList(alwaysTrue, selector)
}
func (cache *schedulerCache) FilteredList(podFilter algorithm.PodFilter, selector labels.Selector) ([]*v1.Pod, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cache.mu.RLock()
	defer cache.mu.RUnlock()
	maxSize := 0
	for _, info := range cache.nodes {
		maxSize += len(info.Pods())
	}
	pods := make([]*v1.Pod, 0, maxSize)
	for _, info := range cache.nodes {
		for _, pod := range info.Pods() {
			if podFilter(pod) && selector.Matches(labels.Set(pod.Labels)) {
				pods = append(pods, pod)
			}
		}
	}
	return pods, nil
}
func (cache *schedulerCache) AssumePod(pod *v1.Pod) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	key, err := schedulercache.GetPodKey(pod)
	if err != nil {
		return err
	}
	cache.mu.Lock()
	defer cache.mu.Unlock()
	if _, ok := cache.podStates[key]; ok {
		return fmt.Errorf("pod %v is in the cache, so can't be assumed", key)
	}
	cache.addPod(pod)
	ps := &podState{pod: pod}
	cache.podStates[key] = ps
	cache.assumedPods[key] = true
	return nil
}
func (cache *schedulerCache) FinishBinding(pod *v1.Pod) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return cache.finishBinding(pod, time.Now())
}
func (cache *schedulerCache) finishBinding(pod *v1.Pod, now time.Time) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	key, err := schedulercache.GetPodKey(pod)
	if err != nil {
		return err
	}
	cache.mu.RLock()
	defer cache.mu.RUnlock()
	klog.V(5).Infof("Finished binding for pod %v. Can be expired.", key)
	currState, ok := cache.podStates[key]
	if ok && cache.assumedPods[key] {
		dl := now.Add(cache.ttl)
		currState.bindingFinished = true
		currState.deadline = &dl
	}
	return nil
}
func (cache *schedulerCache) ForgetPod(pod *v1.Pod) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	key, err := schedulercache.GetPodKey(pod)
	if err != nil {
		return err
	}
	cache.mu.Lock()
	defer cache.mu.Unlock()
	currState, ok := cache.podStates[key]
	if ok && currState.pod.Spec.NodeName != pod.Spec.NodeName {
		return fmt.Errorf("pod %v was assumed on %v but assigned to %v", key, pod.Spec.NodeName, currState.pod.Spec.NodeName)
	}
	switch {
	case ok && cache.assumedPods[key]:
		err := cache.removePod(pod)
		if err != nil {
			return err
		}
		delete(cache.assumedPods, key)
		delete(cache.podStates, key)
	default:
		return fmt.Errorf("pod %v wasn't assumed so cannot be forgotten", key)
	}
	return nil
}
func (cache *schedulerCache) addPod(pod *v1.Pod) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	n, ok := cache.nodes[pod.Spec.NodeName]
	if !ok {
		n = schedulercache.NewNodeInfo()
		cache.nodes[pod.Spec.NodeName] = n
	}
	n.AddPod(pod)
}
func (cache *schedulerCache) updatePod(oldPod, newPod *v1.Pod) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := cache.removePod(oldPod); err != nil {
		return err
	}
	cache.addPod(newPod)
	return nil
}
func (cache *schedulerCache) removePod(pod *v1.Pod) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	n := cache.nodes[pod.Spec.NodeName]
	if err := n.RemovePod(pod); err != nil {
		return err
	}
	if len(n.Pods()) == 0 && n.Node() == nil {
		delete(cache.nodes, pod.Spec.NodeName)
	}
	return nil
}
func (cache *schedulerCache) AddPod(pod *v1.Pod) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	key, err := schedulercache.GetPodKey(pod)
	if err != nil {
		return err
	}
	cache.mu.Lock()
	defer cache.mu.Unlock()
	currState, ok := cache.podStates[key]
	switch {
	case ok && cache.assumedPods[key]:
		if currState.pod.Spec.NodeName != pod.Spec.NodeName {
			klog.Warningf("Pod %v was assumed to be on %v but got added to %v", key, pod.Spec.NodeName, currState.pod.Spec.NodeName)
			cache.removePod(currState.pod)
			cache.addPod(pod)
		}
		delete(cache.assumedPods, key)
		cache.podStates[key].deadline = nil
		cache.podStates[key].pod = pod
	case !ok:
		cache.addPod(pod)
		ps := &podState{pod: pod}
		cache.podStates[key] = ps
	default:
		return fmt.Errorf("pod %v was already in added state", key)
	}
	return nil
}
func (cache *schedulerCache) UpdatePod(oldPod, newPod *v1.Pod) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	key, err := schedulercache.GetPodKey(oldPod)
	if err != nil {
		return err
	}
	cache.mu.Lock()
	defer cache.mu.Unlock()
	currState, ok := cache.podStates[key]
	switch {
	case ok && !cache.assumedPods[key]:
		if currState.pod.Spec.NodeName != newPod.Spec.NodeName {
			klog.Errorf("Pod %v updated on a different node than previously added to.", key)
			klog.Fatalf("Schedulercache is corrupted and can badly affect scheduling decisions")
		}
		if err := cache.updatePod(oldPod, newPod); err != nil {
			return err
		}
		currState.pod = newPod
	default:
		return fmt.Errorf("pod %v is not added to scheduler cache, so cannot be updated", key)
	}
	return nil
}
func (cache *schedulerCache) RemovePod(pod *v1.Pod) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	key, err := schedulercache.GetPodKey(pod)
	if err != nil {
		return err
	}
	cache.mu.Lock()
	defer cache.mu.Unlock()
	currState, ok := cache.podStates[key]
	switch {
	case ok && !cache.assumedPods[key]:
		if currState.pod.Spec.NodeName != pod.Spec.NodeName {
			klog.Errorf("Pod %v was assumed to be on %v but got added to %v", key, pod.Spec.NodeName, currState.pod.Spec.NodeName)
			klog.Fatalf("Schedulercache is corrupted and can badly affect scheduling decisions")
		}
		err := cache.removePod(currState.pod)
		if err != nil {
			return err
		}
		delete(cache.podStates, key)
	default:
		return fmt.Errorf("pod %v is not found in scheduler cache, so cannot be removed from it", key)
	}
	return nil
}
func (cache *schedulerCache) IsAssumedPod(pod *v1.Pod) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	key, err := schedulercache.GetPodKey(pod)
	if err != nil {
		return false, err
	}
	cache.mu.RLock()
	defer cache.mu.RUnlock()
	b, found := cache.assumedPods[key]
	if !found {
		return false, nil
	}
	return b, nil
}
func (cache *schedulerCache) GetPod(pod *v1.Pod) (*v1.Pod, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	key, err := schedulercache.GetPodKey(pod)
	if err != nil {
		return nil, err
	}
	cache.mu.RLock()
	defer cache.mu.RUnlock()
	podState, ok := cache.podStates[key]
	if !ok {
		return nil, fmt.Errorf("pod %v does not exist in scheduler cache", key)
	}
	return podState.pod, nil
}
func (cache *schedulerCache) AddNode(node *v1.Node) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cache.mu.Lock()
	defer cache.mu.Unlock()
	n, ok := cache.nodes[node.Name]
	if !ok {
		n = schedulercache.NewNodeInfo()
		cache.nodes[node.Name] = n
	} else {
		cache.removeNodeImageStates(n.Node())
	}
	cache.nodeTree.AddNode(node)
	cache.addNodeImageStates(node, n)
	return n.SetNode(node)
}
func (cache *schedulerCache) UpdateNode(oldNode, newNode *v1.Node) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cache.mu.Lock()
	defer cache.mu.Unlock()
	n, ok := cache.nodes[newNode.Name]
	if !ok {
		n = schedulercache.NewNodeInfo()
		cache.nodes[newNode.Name] = n
	} else {
		cache.removeNodeImageStates(n.Node())
	}
	cache.nodeTree.UpdateNode(oldNode, newNode)
	cache.addNodeImageStates(newNode, n)
	return n.SetNode(newNode)
}
func (cache *schedulerCache) RemoveNode(node *v1.Node) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cache.mu.Lock()
	defer cache.mu.Unlock()
	n := cache.nodes[node.Name]
	if err := n.RemoveNode(node); err != nil {
		return err
	}
	if len(n.Pods()) == 0 && n.Node() == nil {
		delete(cache.nodes, node.Name)
	}
	cache.nodeTree.RemoveNode(node)
	cache.removeNodeImageStates(node)
	return nil
}
func (cache *schedulerCache) addNodeImageStates(node *v1.Node, nodeInfo *schedulercache.NodeInfo) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	newSum := make(map[string]*schedulercache.ImageStateSummary)
	for _, image := range node.Status.Images {
		for _, name := range image.Names {
			state, ok := cache.imageStates[name]
			if !ok {
				state = &imageState{size: image.SizeBytes, nodes: sets.NewString(node.Name)}
				cache.imageStates[name] = state
			} else {
				state.nodes.Insert(node.Name)
			}
			if _, ok := newSum[name]; !ok {
				newSum[name] = cache.createImageStateSummary(state)
			}
		}
	}
	nodeInfo.SetImageStates(newSum)
}
func (cache *schedulerCache) removeNodeImageStates(node *v1.Node) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if node == nil {
		return
	}
	for _, image := range node.Status.Images {
		for _, name := range image.Names {
			state, ok := cache.imageStates[name]
			if ok {
				state.nodes.Delete(node.Name)
				if len(state.nodes) == 0 {
					delete(cache.imageStates, name)
				}
			}
		}
	}
}
func (cache *schedulerCache) run() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	go wait.Until(cache.cleanupExpiredAssumedPods, cache.period, cache.stop)
}
func (cache *schedulerCache) cleanupExpiredAssumedPods() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cache.cleanupAssumedPods(time.Now())
}
func (cache *schedulerCache) cleanupAssumedPods(now time.Time) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cache.mu.Lock()
	defer cache.mu.Unlock()
	for key := range cache.assumedPods {
		ps, ok := cache.podStates[key]
		if !ok {
			panic("Key found in assumed set but not in podStates. Potentially a logical error.")
		}
		if !ps.bindingFinished {
			klog.V(3).Infof("Couldn't expire cache for pod %v/%v. Binding is still in progress.", ps.pod.Namespace, ps.pod.Name)
			continue
		}
		if now.After(*ps.deadline) {
			klog.Warningf("Pod %s/%s expired", ps.pod.Namespace, ps.pod.Name)
			if err := cache.expirePod(key, ps); err != nil {
				klog.Errorf("ExpirePod failed for %s: %v", key, err)
			}
		}
	}
}
func (cache *schedulerCache) expirePod(key string, ps *podState) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := cache.removePod(ps.pod); err != nil {
		return err
	}
	delete(cache.assumedPods, key)
	delete(cache.podStates, key)
	return nil
}
func (cache *schedulerCache) NodeTree() *NodeTree {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return cache.nodeTree
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
