package garbagecollector

import (
	"fmt"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/controller/garbagecollector/metaonly"
	"reflect"
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
	gvk       schema.GroupVersionKind
}
type GraphBuilder struct {
	restMapper       meta.RESTMapper
	monitors         monitors
	monitorLock      sync.RWMutex
	informersStarted <-chan struct{}
	stopCh           <-chan struct{}
	running          bool
	dynamicClient    dynamic.Interface
	graphChanges     workqueue.RateLimitingInterface
	uidToNode        *concurrentUIDToNode
	attemptToDelete  workqueue.RateLimitingInterface
	attemptToOrphan  workqueue.RateLimitingInterface
	absentOwnerCache *UIDCache
	sharedInformers  informers.SharedInformerFactory
	ignoredResources map[schema.GroupResource]struct{}
}
type monitor struct {
	controller cache.Controller
	store      cache.Store
	stopCh     chan struct{}
}

func (m *monitor) Run() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	m.controller.Run(m.stopCh)
}

type monitors map[schema.GroupVersionResource]*monitor

func listWatcher(client dynamic.Interface, resource schema.GroupVersionResource) *cache.ListWatch {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &cache.ListWatch{ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
		return client.Resource(resource).List(options)
	}, WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
		return client.Resource(resource).Watch(options)
	}}
}
func (gb *GraphBuilder) controllerFor(resource schema.GroupVersionResource, kind schema.GroupVersionKind) (cache.Controller, cache.Store, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	handlers := cache.ResourceEventHandlerFuncs{AddFunc: func(obj interface{}) {
		event := &event{eventType: addEvent, obj: obj, gvk: kind}
		gb.graphChanges.Add(event)
	}, UpdateFunc: func(oldObj, newObj interface{}) {
		event := &event{eventType: updateEvent, obj: newObj, oldObj: oldObj, gvk: kind}
		gb.graphChanges.Add(event)
	}, DeleteFunc: func(obj interface{}) {
		if deletedFinalStateUnknown, ok := obj.(cache.DeletedFinalStateUnknown); ok {
			obj = deletedFinalStateUnknown.Obj
		}
		event := &event{eventType: deleteEvent, obj: obj, gvk: kind}
		gb.graphChanges.Add(event)
	}}
	shared, err := gb.sharedInformers.ForResource(resource)
	if err == nil {
		klog.V(4).Infof("using a shared informer for resource %q, kind %q", resource.String(), kind.String())
		shared.Informer().AddEventHandlerWithResyncPeriod(handlers, ResourceResyncTime)
		return shared.Informer().GetController(), shared.Informer().GetStore(), nil
	} else {
		klog.V(4).Infof("unable to use a shared informer for resource %q, kind %q: %v", resource.String(), kind.String(), err)
	}
	klog.V(5).Infof("create storage for resource %s", resource)
	store, monitor := cache.NewInformer(listWatcher(gb.dynamicClient, resource), nil, ResourceResyncTime, handlers)
	return monitor, store, nil
}
func (gb *GraphBuilder) syncMonitors(resources map[schema.GroupVersionResource]struct{}) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	gb.monitorLock.Lock()
	defer gb.monitorLock.Unlock()
	toRemove := gb.monitors
	if toRemove == nil {
		toRemove = monitors{}
	}
	current := monitors{}
	errs := []error{}
	kept := 0
	added := 0
	for resource := range resources {
		if _, ok := gb.ignoredResources[resource.GroupResource()]; ok {
			continue
		}
		if m, ok := toRemove[resource]; ok {
			current[resource] = m
			delete(toRemove, resource)
			kept++
			continue
		}
		kind, err := gb.restMapper.KindFor(resource)
		if err != nil {
			errs = append(errs, fmt.Errorf("couldn't look up resource %q: %v", resource, err))
			continue
		}
		c, s, err := gb.controllerFor(resource, kind)
		if err != nil {
			errs = append(errs, fmt.Errorf("couldn't start monitor for resource %q: %v", resource, err))
			continue
		}
		current[resource] = &monitor{store: s, controller: c}
		added++
	}
	gb.monitors = current
	for _, monitor := range toRemove {
		if monitor.stopCh != nil {
			close(monitor.stopCh)
		}
	}
	klog.V(4).Infof("synced monitors; added %d, kept %d, removed %d", added, kept, len(toRemove))
	return utilerrors.NewAggregate(errs)
}
func (gb *GraphBuilder) startMonitors() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	gb.monitorLock.Lock()
	defer gb.monitorLock.Unlock()
	if !gb.running {
		return
	}
	<-gb.informersStarted
	monitors := gb.monitors
	started := 0
	for _, monitor := range monitors {
		if monitor.stopCh == nil {
			monitor.stopCh = make(chan struct{})
			gb.sharedInformers.Start(gb.stopCh)
			go monitor.Run()
			started++
		}
	}
	klog.V(4).Infof("started %d new monitors, %d currently running", started, len(monitors))
}
func (gb *GraphBuilder) IsSynced() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	gb.monitorLock.Lock()
	defer gb.monitorLock.Unlock()
	if len(gb.monitors) == 0 {
		klog.V(4).Info("garbage controller monitor not synced: no monitors")
		return false
	}
	for resource, monitor := range gb.monitors {
		if !monitor.controller.HasSynced() {
			klog.V(4).Infof("garbage controller monitor not yet synced: %+v", resource)
			return false
		}
	}
	return true
}
func (gb *GraphBuilder) Run(stopCh <-chan struct{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.Infof("GraphBuilder running")
	defer klog.Infof("GraphBuilder stopping")
	gb.monitorLock.Lock()
	gb.stopCh = stopCh
	gb.running = true
	gb.monitorLock.Unlock()
	gb.startMonitors()
	wait.Until(gb.runProcessGraphChanges, 1*time.Second, stopCh)
	gb.monitorLock.Lock()
	defer gb.monitorLock.Unlock()
	monitors := gb.monitors
	stopped := 0
	for _, monitor := range monitors {
		if monitor.stopCh != nil {
			stopped++
			close(monitor.stopCh)
		}
	}
	gb.monitors = nil
	klog.Infof("stopped %d of %d monitors", stopped, len(monitors))
}

var ignoredResources = map[schema.GroupResource]struct{}{{Group: "", Resource: "events"}: {}}

func DefaultIgnoredResources() map[schema.GroupResource]struct{} {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return ignoredResources
}
func (gb *GraphBuilder) enqueueVirtualDeleteEvent(ref objectReference) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	gb.graphChanges.Add(&event{eventType: deleteEvent, obj: &metaonly.MetadataOnlyObject{TypeMeta: metav1.TypeMeta{APIVersion: ref.APIVersion, Kind: ref.Kind}, ObjectMeta: metav1.ObjectMeta{Namespace: ref.Namespace, UID: ref.UID, Name: ref.Name}}})
}
func (gb *GraphBuilder) addDependentToOwners(n *node, owners []metav1.OwnerReference) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for _, owner := range owners {
		ownerNode, ok := gb.uidToNode.Read(owner.UID)
		if !ok {
			ownerNode = &node{identity: objectReference{OwnerReference: owner, Namespace: n.identity.Namespace}, dependents: make(map[*node]struct{}), virtual: true}
			klog.V(5).Infof("add virtual node.identity: %s\n\n", ownerNode.identity)
			gb.uidToNode.Write(ownerNode)
		}
		ownerNode.addDependent(n)
		if !ok {
			gb.attemptToDelete.Add(ownerNode)
		}
	}
}
func (gb *GraphBuilder) insertNode(n *node) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	gb.uidToNode.Write(n)
	gb.addDependentToOwners(n, n.owners)
}
func (gb *GraphBuilder) removeDependentFromOwners(n *node, owners []metav1.OwnerReference) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for _, owner := range owners {
		ownerNode, ok := gb.uidToNode.Read(owner.UID)
		if !ok {
			continue
		}
		ownerNode.deleteDependent(n)
	}
}
func (gb *GraphBuilder) removeNode(n *node) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	gb.uidToNode.Delete(n.identity.UID)
	gb.removeDependentFromOwners(n, n.owners)
}

type ownerRefPair struct {
	oldRef metav1.OwnerReference
	newRef metav1.OwnerReference
}

func referencesDiffs(old []metav1.OwnerReference, new []metav1.OwnerReference) (added []metav1.OwnerReference, removed []metav1.OwnerReference, changed []ownerRefPair) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	oldUIDToRef := make(map[string]metav1.OwnerReference)
	for _, value := range old {
		oldUIDToRef[string(value.UID)] = value
	}
	oldUIDSet := sets.StringKeySet(oldUIDToRef)
	newUIDToRef := make(map[string]metav1.OwnerReference)
	for _, value := range new {
		newUIDToRef[string(value.UID)] = value
	}
	newUIDSet := sets.StringKeySet(newUIDToRef)
	addedUID := newUIDSet.Difference(oldUIDSet)
	removedUID := oldUIDSet.Difference(newUIDSet)
	intersection := oldUIDSet.Intersection(newUIDSet)
	for uid := range addedUID {
		added = append(added, newUIDToRef[uid])
	}
	for uid := range removedUID {
		removed = append(removed, oldUIDToRef[uid])
	}
	for uid := range intersection {
		if !reflect.DeepEqual(oldUIDToRef[uid], newUIDToRef[uid]) {
			changed = append(changed, ownerRefPair{oldRef: oldUIDToRef[uid], newRef: newUIDToRef[uid]})
		}
	}
	return added, removed, changed
}
func deletionStarts(oldObj interface{}, newAccessor metav1.Object) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if oldObj == nil {
		if newAccessor.GetDeletionTimestamp() == nil {
			return false
		}
		return true
	}
	oldAccessor, err := meta.Accessor(oldObj)
	if err != nil {
		utilruntime.HandleError(fmt.Errorf("cannot access oldObj: %v", err))
		return false
	}
	return beingDeleted(newAccessor) && !beingDeleted(oldAccessor)
}
func beingDeleted(accessor metav1.Object) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return accessor.GetDeletionTimestamp() != nil
}
func hasDeleteDependentsFinalizer(accessor metav1.Object) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	finalizers := accessor.GetFinalizers()
	for _, finalizer := range finalizers {
		if finalizer == metav1.FinalizerDeleteDependents {
			return true
		}
	}
	return false
}
func hasOrphanFinalizer(accessor metav1.Object) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	finalizers := accessor.GetFinalizers()
	for _, finalizer := range finalizers {
		if finalizer == metav1.FinalizerOrphanDependents {
			return true
		}
	}
	return false
}
func startsWaitingForDependentsDeleted(oldObj interface{}, newAccessor metav1.Object) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return deletionStarts(oldObj, newAccessor) && hasDeleteDependentsFinalizer(newAccessor)
}
func startsWaitingForDependentsOrphaned(oldObj interface{}, newAccessor metav1.Object) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return deletionStarts(oldObj, newAccessor) && hasOrphanFinalizer(newAccessor)
}
func (gb *GraphBuilder) addUnblockedOwnersToDeleteQueue(removed []metav1.OwnerReference, changed []ownerRefPair) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for _, ref := range removed {
		if ref.BlockOwnerDeletion != nil && *ref.BlockOwnerDeletion {
			node, found := gb.uidToNode.Read(ref.UID)
			if !found {
				klog.V(5).Infof("cannot find %s in uidToNode", ref.UID)
				continue
			}
			gb.attemptToDelete.Add(node)
		}
	}
	for _, c := range changed {
		wasBlocked := c.oldRef.BlockOwnerDeletion != nil && *c.oldRef.BlockOwnerDeletion
		isUnblocked := c.newRef.BlockOwnerDeletion == nil || (c.newRef.BlockOwnerDeletion != nil && !*c.newRef.BlockOwnerDeletion)
		if wasBlocked && isUnblocked {
			node, found := gb.uidToNode.Read(c.newRef.UID)
			if !found {
				klog.V(5).Infof("cannot find %s in uidToNode", c.newRef.UID)
				continue
			}
			gb.attemptToDelete.Add(node)
		}
	}
}
func (gb *GraphBuilder) processTransitions(oldObj interface{}, newAccessor metav1.Object, n *node) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if startsWaitingForDependentsOrphaned(oldObj, newAccessor) {
		klog.V(5).Infof("add %s to the attemptToOrphan", n.identity)
		gb.attemptToOrphan.Add(n)
		return
	}
	if startsWaitingForDependentsDeleted(oldObj, newAccessor) {
		klog.V(2).Infof("add %s to the attemptToDelete, because it's waiting for its dependents to be deleted", n.identity)
		n.markDeletingDependents()
		for dep := range n.dependents {
			gb.attemptToDelete.Add(dep)
		}
		gb.attemptToDelete.Add(n)
	}
}
func (gb *GraphBuilder) runProcessGraphChanges() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for gb.processGraphChanges() {
	}
}
func (gb *GraphBuilder) processGraphChanges() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	item, quit := gb.graphChanges.Get()
	if quit {
		return false
	}
	defer gb.graphChanges.Done(item)
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
	klog.V(5).Infof("GraphBuilder process object: %s/%s, namespace %s, name %s, uid %s, event type %v", event.gvk.GroupVersion().String(), event.gvk.Kind, accessor.GetNamespace(), accessor.GetName(), string(accessor.GetUID()), event.eventType)
	existingNode, found := gb.uidToNode.Read(accessor.GetUID())
	if found {
		existingNode.markObserved()
	}
	switch {
	case (event.eventType == addEvent || event.eventType == updateEvent) && !found:
		newNode := &node{identity: objectReference{OwnerReference: metav1.OwnerReference{APIVersion: event.gvk.GroupVersion().String(), Kind: event.gvk.Kind, UID: accessor.GetUID(), Name: accessor.GetName()}, Namespace: accessor.GetNamespace()}, dependents: make(map[*node]struct{}), owners: accessor.GetOwnerReferences(), deletingDependents: beingDeleted(accessor) && hasDeleteDependentsFinalizer(accessor), beingDeleted: beingDeleted(accessor)}
		gb.insertNode(newNode)
		gb.processTransitions(event.oldObj, accessor, newNode)
	case (event.eventType == addEvent || event.eventType == updateEvent) && found:
		added, removed, changed := referencesDiffs(existingNode.owners, accessor.GetOwnerReferences())
		if len(added) != 0 || len(removed) != 0 || len(changed) != 0 {
			gb.addUnblockedOwnersToDeleteQueue(removed, changed)
			existingNode.owners = accessor.GetOwnerReferences()
			gb.addDependentToOwners(existingNode, added)
			gb.removeDependentFromOwners(existingNode, removed)
		}
		if beingDeleted(accessor) {
			existingNode.markBeingDeleted()
		}
		gb.processTransitions(event.oldObj, accessor, existingNode)
	case event.eventType == deleteEvent:
		if !found {
			klog.V(5).Infof("%v doesn't exist in the graph, this shouldn't happen", accessor.GetUID())
			return true
		}
		gb.removeNode(existingNode)
		existingNode.dependentsLock.RLock()
		defer existingNode.dependentsLock.RUnlock()
		if len(existingNode.dependents) > 0 {
			gb.absentOwnerCache.Add(accessor.GetUID())
		}
		for dep := range existingNode.dependents {
			gb.attemptToDelete.Add(dep)
		}
		for _, owner := range existingNode.owners {
			ownerNode, found := gb.uidToNode.Read(owner.UID)
			if !found || !ownerNode.isDeletingDependents() {
				continue
			}
			gb.attemptToDelete.Add(ownerNode)
		}
	}
	return true
}
