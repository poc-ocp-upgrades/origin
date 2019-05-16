package resourcequota

import (
	"fmt"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/clock"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/controller"
	quota "k8s.io/kubernetes/pkg/quota/v1"
	"k8s.io/kubernetes/pkg/quota/v1/evaluator/core"
	"k8s.io/kubernetes/pkg/quota/v1/generic"
	"sync"
	"time"
)

type eventType int

func (e eventType) String() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	switch e {
	case addEvent:
		return "add"
	case updateEvent:
		return "update"
	case deleteEvent:
		return "delete"
	default:
		return fmt.Sprintf("unknown(%d)", int(e))
	}
}

const (
	addEvent eventType = iota
	updateEvent
	deleteEvent
)

type event struct {
	eventType eventType
	obj       interface{}
	oldObj    interface{}
	gvr       schema.GroupVersionResource
}
type QuotaMonitor struct {
	monitors          monitors
	monitorLock       sync.Mutex
	informersStarted  <-chan struct{}
	stopCh            <-chan struct{}
	running           bool
	resourceChanges   workqueue.RateLimitingInterface
	informerFactory   InformerFactory
	ignoredResources  map[schema.GroupResource]struct{}
	resyncPeriod      controller.ResyncPeriodFunc
	replenishmentFunc ReplenishmentFunc
	registry          quota.Registry
}

func NewQuotaMonitor(informersStarted <-chan struct{}, informerFactory InformerFactory, ignoredResources map[schema.GroupResource]struct{}, resyncPeriod controller.ResyncPeriodFunc, replenishmentFunc ReplenishmentFunc, registry quota.Registry) *QuotaMonitor {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &QuotaMonitor{informersStarted: informersStarted, informerFactory: informerFactory, ignoredResources: ignoredResources, resourceChanges: workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "resource_quota_controller_resource_changes"), resyncPeriod: resyncPeriod, replenishmentFunc: replenishmentFunc, registry: registry}
}

type monitor struct {
	controller cache.Controller
	stopCh     chan struct{}
}

func (m *monitor) Run() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	m.controller.Run(m.stopCh)
}

type monitors map[schema.GroupVersionResource]*monitor

func (qm *QuotaMonitor) controllerFor(resource schema.GroupVersionResource) (cache.Controller, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	clock := clock.RealClock{}
	handlers := cache.ResourceEventHandlerFuncs{UpdateFunc: func(oldObj, newObj interface{}) {
		notifyUpdate := false
		switch resource.GroupResource() {
		case schema.GroupResource{Resource: "pods"}:
			oldPod := oldObj.(*v1.Pod)
			newPod := newObj.(*v1.Pod)
			notifyUpdate = core.QuotaV1Pod(oldPod, clock) && !core.QuotaV1Pod(newPod, clock)
		case schema.GroupResource{Resource: "services"}:
			oldService := oldObj.(*v1.Service)
			newService := newObj.(*v1.Service)
			notifyUpdate = core.GetQuotaServiceType(oldService) != core.GetQuotaServiceType(newService)
		}
		if notifyUpdate {
			event := &event{eventType: updateEvent, obj: newObj, oldObj: oldObj, gvr: resource}
			qm.resourceChanges.Add(event)
		}
	}, DeleteFunc: func(obj interface{}) {
		if deletedFinalStateUnknown, ok := obj.(cache.DeletedFinalStateUnknown); ok {
			obj = deletedFinalStateUnknown.Obj
		}
		event := &event{eventType: deleteEvent, obj: obj, gvr: resource}
		qm.resourceChanges.Add(event)
	}}
	shared, err := qm.informerFactory.ForResource(resource)
	if err == nil {
		klog.V(4).Infof("QuotaMonitor using a shared informer for resource %q", resource.String())
		shared.Informer().AddEventHandlerWithResyncPeriod(handlers, qm.resyncPeriod())
		return shared.Informer().GetController(), nil
	}
	klog.V(4).Infof("QuotaMonitor unable to use a shared informer for resource %q: %v", resource.String(), err)
	return nil, fmt.Errorf("unable to monitor quota for resource %q", resource.String())
}
func (qm *QuotaMonitor) SyncMonitors(resources map[schema.GroupVersionResource]struct{}) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	qm.monitorLock.Lock()
	defer qm.monitorLock.Unlock()
	toRemove := qm.monitors
	if toRemove == nil {
		toRemove = monitors{}
	}
	current := monitors{}
	errs := []error{}
	kept := 0
	added := 0
	for resource := range resources {
		if _, ok := qm.ignoredResources[resource.GroupResource()]; ok {
			continue
		}
		if m, ok := toRemove[resource]; ok {
			current[resource] = m
			delete(toRemove, resource)
			kept++
			continue
		}
		c, err := qm.controllerFor(resource)
		if err != nil {
			errs = append(errs, fmt.Errorf("couldn't start monitor for resource %q: %v", resource, err))
			continue
		}
		evaluator := qm.registry.Get(resource.GroupResource())
		if evaluator == nil {
			listerFunc := generic.ListerFuncForResourceFunc(qm.informerFactory.ForResource)
			listResourceFunc := generic.ListResourceUsingListerFunc(listerFunc, resource)
			evaluator = generic.NewObjectCountEvaluator(resource.GroupResource(), listResourceFunc, "")
			qm.registry.Add(evaluator)
			klog.Infof("QuotaMonitor created object count evaluator for %s", resource.GroupResource())
		}
		current[resource] = &monitor{controller: c}
		added++
	}
	qm.monitors = current
	for _, monitor := range toRemove {
		if monitor.stopCh != nil {
			close(monitor.stopCh)
		}
	}
	klog.V(4).Infof("quota synced monitors; added %d, kept %d, removed %d", added, kept, len(toRemove))
	return utilerrors.NewAggregate(errs)
}
func (qm *QuotaMonitor) StartMonitors() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	qm.monitorLock.Lock()
	defer qm.monitorLock.Unlock()
	if !qm.running {
		return
	}
	<-qm.informersStarted
	monitors := qm.monitors
	started := 0
	for _, monitor := range monitors {
		if monitor.stopCh == nil {
			monitor.stopCh = make(chan struct{})
			qm.informerFactory.Start(qm.stopCh)
			go monitor.Run()
			started++
		}
	}
	klog.V(4).Infof("QuotaMonitor started %d new monitors, %d currently running", started, len(monitors))
}
func (qm *QuotaMonitor) IsSynced() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	qm.monitorLock.Lock()
	defer qm.monitorLock.Unlock()
	if len(qm.monitors) == 0 {
		klog.V(4).Info("quota monitor not synced: no monitors")
		return false
	}
	for resource, monitor := range qm.monitors {
		if !monitor.controller.HasSynced() {
			klog.V(4).Infof("quota monitor not synced: %v", resource)
			return false
		}
	}
	return true
}
func (qm *QuotaMonitor) Run(stopCh <-chan struct{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.Infof("QuotaMonitor running")
	defer klog.Infof("QuotaMonitor stopping")
	qm.monitorLock.Lock()
	qm.stopCh = stopCh
	qm.running = true
	qm.monitorLock.Unlock()
	qm.StartMonitors()
	wait.Until(qm.runProcessResourceChanges, 1*time.Second, stopCh)
	qm.monitorLock.Lock()
	defer qm.monitorLock.Unlock()
	monitors := qm.monitors
	stopped := 0
	for _, monitor := range monitors {
		if monitor.stopCh != nil {
			stopped++
			close(monitor.stopCh)
		}
	}
	klog.Infof("QuotaMonitor stopped %d of %d monitors", stopped, len(monitors))
}
func (qm *QuotaMonitor) runProcessResourceChanges() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for qm.processResourceChanges() {
	}
}
func (qm *QuotaMonitor) processResourceChanges() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	item, quit := qm.resourceChanges.Get()
	if quit {
		return false
	}
	defer qm.resourceChanges.Done(item)
	event, ok := item.(*event)
	if !ok {
		utilruntime.HandleError(fmt.Errorf("expect a *event, got %v", item))
		return true
	}
	obj := event.obj
	accessor, err := meta.Accessor(obj)
	if err != nil {
		utilruntime.HandleError(fmt.Errorf("cannot access obj: %v", err))
		return true
	}
	klog.V(4).Infof("QuotaMonitor process object: %s, namespace %s, name %s, uid %s, event type %v", event.gvr.String(), accessor.GetNamespace(), accessor.GetName(), string(accessor.GetUID()), event.eventType)
	qm.replenishmentFunc(event.gvr.GroupResource(), accessor.GetNamespace())
	return true
}
