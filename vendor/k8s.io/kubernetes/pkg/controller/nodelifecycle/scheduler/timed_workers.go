package scheduler

import (
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/klog"
	"sync"
	"time"
)

type WorkArgs struct{ NamespacedName types.NamespacedName }

func (w *WorkArgs) KeyFromWorkArgs() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return w.NamespacedName.String()
}
func NewWorkArgs(name, namespace string) *WorkArgs {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &WorkArgs{types.NamespacedName{Namespace: namespace, Name: name}}
}

type TimedWorker struct {
	WorkItem  *WorkArgs
	CreatedAt time.Time
	FireAt    time.Time
	Timer     *time.Timer
}

func CreateWorker(args *WorkArgs, createdAt time.Time, fireAt time.Time, f func(args *WorkArgs) error) *TimedWorker {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	delay := fireAt.Sub(createdAt)
	if delay <= 0 {
		go f(args)
		return nil
	}
	timer := time.AfterFunc(delay, func() {
		f(args)
	})
	return &TimedWorker{WorkItem: args, CreatedAt: createdAt, FireAt: fireAt, Timer: timer}
}
func (w *TimedWorker) Cancel() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if w != nil {
		w.Timer.Stop()
	}
}

type TimedWorkerQueue struct {
	sync.Mutex
	workers  map[string]*TimedWorker
	workFunc func(args *WorkArgs) error
}

func CreateWorkerQueue(f func(args *WorkArgs) error) *TimedWorkerQueue {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &TimedWorkerQueue{workers: make(map[string]*TimedWorker), workFunc: f}
}
func (q *TimedWorkerQueue) getWrappedWorkerFunc(key string) func(args *WorkArgs) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return func(args *WorkArgs) error {
		err := q.workFunc(args)
		q.Lock()
		defer q.Unlock()
		if err == nil {
			q.workers[key] = nil
		} else {
			delete(q.workers, key)
		}
		return err
	}
}
func (q *TimedWorkerQueue) AddWork(args *WorkArgs, createdAt time.Time, fireAt time.Time) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	key := args.KeyFromWorkArgs()
	klog.V(4).Infof("Adding TimedWorkerQueue item %v at %v to be fired at %v", key, createdAt, fireAt)
	q.Lock()
	defer q.Unlock()
	if _, exists := q.workers[key]; exists {
		klog.Warningf("Trying to add already existing work for %+v. Skipping.", args)
		return
	}
	worker := CreateWorker(args, createdAt, fireAt, q.getWrappedWorkerFunc(key))
	q.workers[key] = worker
}
func (q *TimedWorkerQueue) CancelWork(key string) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	q.Lock()
	defer q.Unlock()
	worker, found := q.workers[key]
	result := false
	if found {
		klog.V(4).Infof("Cancelling TimedWorkerQueue item %v at %v", key, time.Now())
		if worker != nil {
			result = true
			worker.Cancel()
		}
		delete(q.workers, key)
	}
	return result
}
func (q *TimedWorkerQueue) GetWorkerUnsafe(key string) *TimedWorker {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	q.Lock()
	defer q.Unlock()
	return q.workers[key]
}
