package persistentvolume

import (
	"k8s.io/api/core/v1"
	"sync"
)

type PodBindingCache interface {
	UpdateBindings(pod *v1.Pod, node string, bindings []*bindingInfo)
	GetBindings(pod *v1.Pod, node string) []*bindingInfo
	UpdateProvisionedPVCs(pod *v1.Pod, node string, provisionings []*v1.PersistentVolumeClaim)
	GetProvisionedPVCs(pod *v1.Pod, node string) []*v1.PersistentVolumeClaim
	DeleteBindings(pod *v1.Pod)
}
type podBindingCache struct {
	rwMutex          sync.RWMutex
	bindingDecisions map[string]nodeDecisions
}
type nodeDecisions map[string]nodeDecision
type nodeDecision struct {
	bindings      []*bindingInfo
	provisionings []*v1.PersistentVolumeClaim
}

func NewPodBindingCache() PodBindingCache {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &podBindingCache{bindingDecisions: map[string]nodeDecisions{}}
}
func (c *podBindingCache) DeleteBindings(pod *v1.Pod) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	c.rwMutex.Lock()
	defer c.rwMutex.Unlock()
	podName := getPodName(pod)
	if _, ok := c.bindingDecisions[podName]; ok {
		delete(c.bindingDecisions, podName)
		VolumeBindingRequestSchedulerBinderCache.WithLabelValues("delete").Inc()
	}
}
func (c *podBindingCache) UpdateBindings(pod *v1.Pod, node string, bindings []*bindingInfo) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	c.rwMutex.Lock()
	defer c.rwMutex.Unlock()
	podName := getPodName(pod)
	decisions, ok := c.bindingDecisions[podName]
	if !ok {
		decisions = nodeDecisions{}
		c.bindingDecisions[podName] = decisions
	}
	decision, ok := decisions[node]
	if !ok {
		decision = nodeDecision{bindings: bindings}
		VolumeBindingRequestSchedulerBinderCache.WithLabelValues("add").Inc()
	} else {
		decision.bindings = bindings
	}
	decisions[node] = decision
}
func (c *podBindingCache) GetBindings(pod *v1.Pod, node string) []*bindingInfo {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	c.rwMutex.RLock()
	defer c.rwMutex.RUnlock()
	podName := getPodName(pod)
	decisions, ok := c.bindingDecisions[podName]
	if !ok {
		return nil
	}
	decision, ok := decisions[node]
	if !ok {
		return nil
	}
	return decision.bindings
}
func (c *podBindingCache) UpdateProvisionedPVCs(pod *v1.Pod, node string, pvcs []*v1.PersistentVolumeClaim) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	c.rwMutex.Lock()
	defer c.rwMutex.Unlock()
	podName := getPodName(pod)
	decisions, ok := c.bindingDecisions[podName]
	if !ok {
		decisions = nodeDecisions{}
		c.bindingDecisions[podName] = decisions
	}
	decision, ok := decisions[node]
	if !ok {
		decision = nodeDecision{provisionings: pvcs}
	} else {
		decision.provisionings = pvcs
	}
	decisions[node] = decision
}
func (c *podBindingCache) GetProvisionedPVCs(pod *v1.Pod, node string) []*v1.PersistentVolumeClaim {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	c.rwMutex.RLock()
	defer c.rwMutex.RUnlock()
	podName := getPodName(pod)
	decisions, ok := c.bindingDecisions[podName]
	if !ok {
		return nil
	}
	decision, ok := decisions[node]
	if !ok {
		return nil
	}
	return decision.provisionings
}
