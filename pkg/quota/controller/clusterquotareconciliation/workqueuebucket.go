package clusterquotareconciliation

import (
	"sync"
	"k8s.io/client-go/util/workqueue"
)

type BucketingWorkQueue interface {
	AddWithData(key interface{}, data ...interface{})
	AddWithDataRateLimited(key interface{}, data ...interface{})
	GetWithData() (key interface{}, data []interface{}, quit bool)
	Done(key interface{})
	Forget(key interface{})
	ShutDown()
}

func NewBucketingWorkQueue(name string) BucketingWorkQueue {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &workQueueBucket{queue: workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), name), work: map[interface{}][]interface{}{}, dirtyWork: map[interface{}][]interface{}{}, inProgress: map[interface{}]bool{}}
}

type workQueueBucket struct {
	queue		workqueue.RateLimitingInterface
	workLock	sync.Mutex
	work		map[interface{}][]interface{}
	dirtyWork	map[interface{}][]interface{}
	inProgress	map[interface{}]bool
}

func (e *workQueueBucket) AddWithData(key interface{}, data ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	e.workLock.Lock()
	defer e.workLock.Unlock()
	e.queue.Add(key)
	if e.inProgress[key] {
		e.dirtyWork[key] = append(e.dirtyWork[key], data...)
		return
	}
	e.work[key] = append(e.work[key], data...)
}
func (e *workQueueBucket) AddWithDataRateLimited(key interface{}, data ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	e.workLock.Lock()
	defer e.workLock.Unlock()
	e.queue.AddRateLimited(key)
	if e.inProgress[key] {
		e.dirtyWork[key] = append(e.dirtyWork[key], data...)
		return
	}
	e.work[key] = append(e.work[key], data...)
}
func (e *workQueueBucket) Done(key interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	e.workLock.Lock()
	defer e.workLock.Unlock()
	e.queue.Done(key)
	e.work[key] = e.dirtyWork[key]
	delete(e.dirtyWork, key)
	delete(e.inProgress, key)
}
func (e *workQueueBucket) Forget(key interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	e.queue.Forget(key)
}
func (e *workQueueBucket) GetWithData() (interface{}, []interface{}, bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	key, shutdown := e.queue.Get()
	if shutdown {
		return nil, []interface{}{}, shutdown
	}
	e.workLock.Lock()
	defer e.workLock.Unlock()
	work := e.work[key]
	delete(e.work, key)
	delete(e.dirtyWork, key)
	e.inProgress[key] = true
	if len(work) != 0 {
		return key, work, false
	}
	return key, []interface{}{}, false
}
func (e *workQueueBucket) ShutDown() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	e.queue.ShutDown()
}
