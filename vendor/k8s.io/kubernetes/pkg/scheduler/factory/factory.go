package factory

import (
	"fmt"
	goformat "fmt"
	"k8s.io/api/core/v1"
	storagev1 "k8s.io/api/storage/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apimachinery/pkg/util/wait"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	appsinformers "k8s.io/client-go/informers/apps/v1"
	coreinformers "k8s.io/client-go/informers/core/v1"
	policyinformers "k8s.io/client-go/informers/policy/v1beta1"
	storageinformers "k8s.io/client-go/informers/storage/v1"
	clientset "k8s.io/client-go/kubernetes"
	appslisters "k8s.io/client-go/listers/apps/v1"
	corelisters "k8s.io/client-go/listers/core/v1"
	policylisters "k8s.io/client-go/listers/policy/v1beta1"
	storagelisters "k8s.io/client-go/listers/storage/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	"k8s.io/klog"
	podutil "k8s.io/kubernetes/pkg/api/v1/pod"
	"k8s.io/kubernetes/pkg/apis/core/helper"
	"k8s.io/kubernetes/pkg/features"
	kubeletapis "k8s.io/kubernetes/pkg/kubelet/apis"
	"k8s.io/kubernetes/pkg/scheduler/algorithm"
	"k8s.io/kubernetes/pkg/scheduler/algorithm/predicates"
	schedulerapi "k8s.io/kubernetes/pkg/scheduler/api"
	"k8s.io/kubernetes/pkg/scheduler/api/validation"
	"k8s.io/kubernetes/pkg/scheduler/core"
	"k8s.io/kubernetes/pkg/scheduler/core/equivalence"
	schedulerinternalcache "k8s.io/kubernetes/pkg/scheduler/internal/cache"
	cachedebugger "k8s.io/kubernetes/pkg/scheduler/internal/cache/debugger"
	internalqueue "k8s.io/kubernetes/pkg/scheduler/internal/queue"
	"k8s.io/kubernetes/pkg/scheduler/util"
	"k8s.io/kubernetes/pkg/scheduler/volumebinder"
	"os"
	goos "os"
	"os/signal"
	"reflect"
	godefaultruntime "runtime"
	"time"
	gotime "time"
)

const (
	initialGetBackoff = 100 * time.Millisecond
	maximalGetBackoff = time.Minute
)

var (
	serviceAffinitySet            = sets.NewString(predicates.CheckServiceAffinityPred)
	matchInterPodAffinitySet      = sets.NewString(predicates.MatchInterPodAffinityPred)
	generalPredicatesSets         = sets.NewString(predicates.GeneralPred)
	noDiskConflictSet             = sets.NewString(predicates.NoDiskConflictPred)
	maxPDVolumeCountPredicateKeys = []string{predicates.MaxGCEPDVolumeCountPred, predicates.MaxAzureDiskVolumeCountPred, predicates.MaxEBSVolumeCountPred, predicates.MaxCinderVolumeCountPred}
)

type Binder interface {
	Bind(binding *v1.Binding) error
}
type PodConditionUpdater interface {
	Update(pod *v1.Pod, podCondition *v1.PodCondition) error
}
type Config struct {
	SchedulerCache      schedulerinternalcache.Cache
	Ecache              *equivalence.Cache
	NodeLister          algorithm.NodeLister
	Algorithm           algorithm.ScheduleAlgorithm
	GetBinder           func(pod *v1.Pod) Binder
	PodConditionUpdater PodConditionUpdater
	PodPreemptor        PodPreemptor
	NextPod             func() *v1.Pod
	WaitForCacheSync    func() bool
	Error               func(*v1.Pod, error)
	Recorder            record.EventRecorder
	StopEverything      <-chan struct{}
	VolumeBinder        *volumebinder.VolumeBinder
	DisablePreemption   bool
	SchedulingQueue     internalqueue.SchedulingQueue
}
type PodPreemptor interface {
	GetUpdatedPod(pod *v1.Pod) (*v1.Pod, error)
	DeletePod(pod *v1.Pod) error
	SetNominatedNodeName(pod *v1.Pod, nominatedNode string) error
	RemoveNominatedNodeName(pod *v1.Pod) error
}
type Configurator interface {
	GetHardPodAffinitySymmetricWeight() int32
	MakeDefaultErrorFunc(backoff *util.PodBackoff, podQueue internalqueue.SchedulingQueue) func(pod *v1.Pod, err error)
	GetPredicateMetadataProducer() (algorithm.PredicateMetadataProducer, error)
	GetPredicates(predicateKeys sets.String) (map[string]algorithm.FitPredicate, error)
	GetNodeLister() corelisters.NodeLister
	GetClient() clientset.Interface
	GetScheduledPodLister() corelisters.PodLister
	Create() (*Config, error)
	CreateFromProvider(providerName string) (*Config, error)
	CreateFromConfig(policy schedulerapi.Policy) (*Config, error)
	CreateFromKeys(predicateKeys, priorityKeys sets.String, extenders []algorithm.SchedulerExtender) (*Config, error)
}
type configFactory struct {
	client                         clientset.Interface
	podQueue                       internalqueue.SchedulingQueue
	scheduledPodLister             corelisters.PodLister
	podLister                      algorithm.PodLister
	nodeLister                     corelisters.NodeLister
	pVLister                       corelisters.PersistentVolumeLister
	pVCLister                      corelisters.PersistentVolumeClaimLister
	serviceLister                  corelisters.ServiceLister
	controllerLister               corelisters.ReplicationControllerLister
	replicaSetLister               appslisters.ReplicaSetLister
	statefulSetLister              appslisters.StatefulSetLister
	pdbLister                      policylisters.PodDisruptionBudgetLister
	storageClassLister             storagelisters.StorageClassLister
	StopEverything                 <-chan struct{}
	scheduledPodsHasSynced         cache.InformerSynced
	schedulerCache                 schedulerinternalcache.Cache
	schedulerName                  string
	hardPodAffinitySymmetricWeight int32
	equivalencePodCache            *equivalence.Cache
	enableEquivalenceClassCache    bool
	volumeBinder                   *volumebinder.VolumeBinder
	alwaysCheckAllPredicates       bool
	disablePreemption              bool
	percentageOfNodesToScore       int32
}
type ConfigFactoryArgs struct {
	SchedulerName                  string
	Client                         clientset.Interface
	NodeInformer                   coreinformers.NodeInformer
	PodInformer                    coreinformers.PodInformer
	PvInformer                     coreinformers.PersistentVolumeInformer
	PvcInformer                    coreinformers.PersistentVolumeClaimInformer
	ReplicationControllerInformer  coreinformers.ReplicationControllerInformer
	ReplicaSetInformer             appsinformers.ReplicaSetInformer
	StatefulSetInformer            appsinformers.StatefulSetInformer
	ServiceInformer                coreinformers.ServiceInformer
	PdbInformer                    policyinformers.PodDisruptionBudgetInformer
	StorageClassInformer           storageinformers.StorageClassInformer
	HardPodAffinitySymmetricWeight int32
	EnableEquivalenceClassCache    bool
	DisablePreemption              bool
	PercentageOfNodesToScore       int32
	BindTimeoutSeconds             int64
	StopCh                         <-chan struct{}
}

func NewConfigFactory(args *ConfigFactoryArgs) Configurator {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	stopEverything := args.StopCh
	if stopEverything == nil {
		stopEverything = wait.NeverStop
	}
	schedulerCache := schedulerinternalcache.New(30*time.Second, stopEverything)
	var storageClassLister storagelisters.StorageClassLister
	if args.StorageClassInformer != nil {
		storageClassLister = args.StorageClassInformer.Lister()
	}
	c := &configFactory{client: args.Client, podLister: schedulerCache, podQueue: internalqueue.NewSchedulingQueue(stopEverything), nodeLister: args.NodeInformer.Lister(), pVLister: args.PvInformer.Lister(), pVCLister: args.PvcInformer.Lister(), serviceLister: args.ServiceInformer.Lister(), controllerLister: args.ReplicationControllerInformer.Lister(), replicaSetLister: args.ReplicaSetInformer.Lister(), statefulSetLister: args.StatefulSetInformer.Lister(), pdbLister: args.PdbInformer.Lister(), storageClassLister: storageClassLister, schedulerCache: schedulerCache, StopEverything: stopEverything, schedulerName: args.SchedulerName, hardPodAffinitySymmetricWeight: args.HardPodAffinitySymmetricWeight, enableEquivalenceClassCache: args.EnableEquivalenceClassCache, disablePreemption: args.DisablePreemption, percentageOfNodesToScore: args.PercentageOfNodesToScore}
	c.scheduledPodsHasSynced = args.PodInformer.Informer().HasSynced
	args.PodInformer.Informer().AddEventHandler(cache.FilteringResourceEventHandler{FilterFunc: func(obj interface{}) bool {
		switch t := obj.(type) {
		case *v1.Pod:
			return assignedPod(t)
		case cache.DeletedFinalStateUnknown:
			if pod, ok := t.Obj.(*v1.Pod); ok {
				return assignedPod(pod)
			}
			runtime.HandleError(fmt.Errorf("unable to convert object %T to *v1.Pod in %T", obj, c))
			return false
		default:
			runtime.HandleError(fmt.Errorf("unable to handle object in %T: %T", c, obj))
			return false
		}
	}, Handler: cache.ResourceEventHandlerFuncs{AddFunc: c.addPodToCache, UpdateFunc: c.updatePodInCache, DeleteFunc: c.deletePodFromCache}})
	args.PodInformer.Informer().AddEventHandler(cache.FilteringResourceEventHandler{FilterFunc: func(obj interface{}) bool {
		switch t := obj.(type) {
		case *v1.Pod:
			return !assignedPod(t) && responsibleForPod(t, args.SchedulerName)
		case cache.DeletedFinalStateUnknown:
			if pod, ok := t.Obj.(*v1.Pod); ok {
				return !assignedPod(pod) && responsibleForPod(pod, args.SchedulerName)
			}
			runtime.HandleError(fmt.Errorf("unable to convert object %T to *v1.Pod in %T", obj, c))
			return false
		default:
			runtime.HandleError(fmt.Errorf("unable to handle object in %T: %T", c, obj))
			return false
		}
	}, Handler: cache.ResourceEventHandlerFuncs{AddFunc: c.addPodToSchedulingQueue, UpdateFunc: c.updatePodInSchedulingQueue, DeleteFunc: c.deletePodFromSchedulingQueue}})
	c.scheduledPodLister = assignedPodLister{args.PodInformer.Lister()}
	args.NodeInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{AddFunc: c.addNodeToCache, UpdateFunc: c.updateNodeInCache, DeleteFunc: c.deleteNodeFromCache})
	args.PvInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{AddFunc: c.onPvAdd, UpdateFunc: c.onPvUpdate, DeleteFunc: c.onPvDelete})
	args.PvcInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{AddFunc: c.onPvcAdd, UpdateFunc: c.onPvcUpdate, DeleteFunc: c.onPvcDelete})
	args.ServiceInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{AddFunc: c.onServiceAdd, UpdateFunc: c.onServiceUpdate, DeleteFunc: c.onServiceDelete})
	if utilfeature.DefaultFeatureGate.Enabled(features.VolumeScheduling) {
		c.volumeBinder = volumebinder.NewVolumeBinder(args.Client, args.PvcInformer, args.PvInformer, args.StorageClassInformer, time.Duration(args.BindTimeoutSeconds)*time.Second)
		args.StorageClassInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{AddFunc: c.onStorageClassAdd, DeleteFunc: c.onStorageClassDelete})
	}
	debugger := cachedebugger.New(args.NodeInformer.Lister(), args.PodInformer.Lister(), c.schedulerCache, c.podQueue)
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, compareSignal)
	go func() {
		for {
			select {
			case <-c.StopEverything:
				c.podQueue.Close()
				return
			case <-ch:
				debugger.Comparer.Compare()
				debugger.Dumper.DumpAll()
			}
		}
	}()
	return c
}
func (c *configFactory) skipPodUpdate(pod *v1.Pod) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	isAssumed, err := c.schedulerCache.IsAssumedPod(pod)
	if err != nil {
		runtime.HandleError(fmt.Errorf("failed to check whether pod %s/%s is assumed: %v", pod.Namespace, pod.Name, err))
		return false
	}
	if !isAssumed {
		return false
	}
	assumedPod, err := c.schedulerCache.GetPod(pod)
	if err != nil {
		runtime.HandleError(fmt.Errorf("failed to get assumed pod %s/%s from cache: %v", pod.Namespace, pod.Name, err))
		return false
	}
	f := func(pod *v1.Pod) *v1.Pod {
		p := pod.DeepCopy()
		p.ResourceVersion = ""
		p.Spec.NodeName = ""
		p.Annotations = nil
		return p
	}
	assumedPodCopy, podCopy := f(assumedPod), f(pod)
	if !reflect.DeepEqual(assumedPodCopy, podCopy) {
		return false
	}
	klog.V(3).Infof("Skipping pod %s/%s update", pod.Namespace, pod.Name)
	return true
}
func (c *configFactory) onPvAdd(obj interface{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if c.enableEquivalenceClassCache {
		pv, ok := obj.(*v1.PersistentVolume)
		if !ok {
			klog.Errorf("cannot convert to *v1.PersistentVolume: %v", obj)
			return
		}
		c.invalidatePredicatesForPv(pv)
	}
	c.podQueue.MoveAllToActiveQueue()
}
func (c *configFactory) onPvUpdate(old, new interface{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if c.enableEquivalenceClassCache {
		newPV, ok := new.(*v1.PersistentVolume)
		if !ok {
			klog.Errorf("cannot convert to *v1.PersistentVolume: %v", new)
			return
		}
		oldPV, ok := old.(*v1.PersistentVolume)
		if !ok {
			klog.Errorf("cannot convert to *v1.PersistentVolume: %v", old)
			return
		}
		c.invalidatePredicatesForPvUpdate(oldPV, newPV)
	}
	c.podQueue.MoveAllToActiveQueue()
}
func (c *configFactory) invalidatePredicatesForPvUpdate(oldPV, newPV *v1.PersistentVolume) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	invalidPredicates := sets.NewString()
	if utilfeature.DefaultFeatureGate.Enabled(features.VolumeScheduling) {
		invalidPredicates.Insert(predicates.CheckVolumeBindingPred)
	}
	for k, v := range newPV.Labels {
		if isZoneRegionLabel(k) && !reflect.DeepEqual(v, oldPV.Labels[k]) {
			invalidPredicates.Insert(predicates.NoVolumeZoneConflictPred)
			break
		}
	}
	c.equivalencePodCache.InvalidatePredicates(invalidPredicates)
}
func isZoneRegionLabel(k string) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return k == kubeletapis.LabelZoneFailureDomain || k == kubeletapis.LabelZoneRegion
}
func (c *configFactory) onPvDelete(obj interface{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if c.enableEquivalenceClassCache {
		var pv *v1.PersistentVolume
		switch t := obj.(type) {
		case *v1.PersistentVolume:
			pv = t
		case cache.DeletedFinalStateUnknown:
			var ok bool
			pv, ok = t.Obj.(*v1.PersistentVolume)
			if !ok {
				klog.Errorf("cannot convert to *v1.PersistentVolume: %v", t.Obj)
				return
			}
		default:
			klog.Errorf("cannot convert to *v1.PersistentVolume: %v", t)
			return
		}
		c.invalidatePredicatesForPv(pv)
	}
}
func (c *configFactory) invalidatePredicatesForPv(pv *v1.PersistentVolume) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	invalidPredicates := sets.NewString()
	if pv.Spec.AWSElasticBlockStore != nil {
		invalidPredicates.Insert(predicates.MaxEBSVolumeCountPred)
	}
	if pv.Spec.GCEPersistentDisk != nil {
		invalidPredicates.Insert(predicates.MaxGCEPDVolumeCountPred)
	}
	if pv.Spec.AzureDisk != nil {
		invalidPredicates.Insert(predicates.MaxAzureDiskVolumeCountPred)
	}
	if pv.Spec.CSI != nil && utilfeature.DefaultFeatureGate.Enabled(features.AttachVolumeLimit) {
		invalidPredicates.Insert(predicates.MaxCSIVolumeCountPred)
	}
	for k := range pv.Labels {
		if isZoneRegionLabel(k) {
			invalidPredicates.Insert(predicates.NoVolumeZoneConflictPred)
			break
		}
	}
	if utilfeature.DefaultFeatureGate.Enabled(features.VolumeScheduling) {
		invalidPredicates.Insert(predicates.CheckVolumeBindingPred)
	}
	c.equivalencePodCache.InvalidatePredicates(invalidPredicates)
}
func (c *configFactory) onPvcAdd(obj interface{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if c.enableEquivalenceClassCache {
		pvc, ok := obj.(*v1.PersistentVolumeClaim)
		if !ok {
			klog.Errorf("cannot convert to *v1.PersistentVolumeClaim: %v", obj)
			return
		}
		c.invalidatePredicatesForPvc(pvc)
	}
	c.podQueue.MoveAllToActiveQueue()
}
func (c *configFactory) onPvcUpdate(old, new interface{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if !utilfeature.DefaultFeatureGate.Enabled(features.VolumeScheduling) {
		return
	}
	if c.enableEquivalenceClassCache {
		newPVC, ok := new.(*v1.PersistentVolumeClaim)
		if !ok {
			klog.Errorf("cannot convert to *v1.PersistentVolumeClaim: %v", new)
			return
		}
		oldPVC, ok := old.(*v1.PersistentVolumeClaim)
		if !ok {
			klog.Errorf("cannot convert to *v1.PersistentVolumeClaim: %v", old)
			return
		}
		c.invalidatePredicatesForPvcUpdate(oldPVC, newPVC)
	}
	c.podQueue.MoveAllToActiveQueue()
}
func (c *configFactory) onPvcDelete(obj interface{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if c.enableEquivalenceClassCache {
		var pvc *v1.PersistentVolumeClaim
		switch t := obj.(type) {
		case *v1.PersistentVolumeClaim:
			pvc = t
		case cache.DeletedFinalStateUnknown:
			var ok bool
			pvc, ok = t.Obj.(*v1.PersistentVolumeClaim)
			if !ok {
				klog.Errorf("cannot convert to *v1.PersistentVolumeClaim: %v", t.Obj)
				return
			}
		default:
			klog.Errorf("cannot convert to *v1.PersistentVolumeClaim: %v", t)
			return
		}
		c.invalidatePredicatesForPvc(pvc)
	}
}
func (c *configFactory) invalidatePredicatesForPvc(pvc *v1.PersistentVolumeClaim) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	invalidPredicates := sets.NewString(maxPDVolumeCountPredicateKeys...)
	if utilfeature.DefaultFeatureGate.Enabled(features.AttachVolumeLimit) {
		invalidPredicates.Insert(predicates.MaxCSIVolumeCountPred)
	}
	invalidPredicates.Insert(predicates.NoVolumeZoneConflictPred)
	if utilfeature.DefaultFeatureGate.Enabled(features.VolumeScheduling) {
		invalidPredicates.Insert(predicates.CheckVolumeBindingPred)
	}
	c.equivalencePodCache.InvalidatePredicates(invalidPredicates)
}
func (c *configFactory) invalidatePredicatesForPvcUpdate(old, new *v1.PersistentVolumeClaim) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	invalidPredicates := sets.NewString()
	if old.Spec.VolumeName != new.Spec.VolumeName {
		if utilfeature.DefaultFeatureGate.Enabled(features.VolumeScheduling) {
			invalidPredicates.Insert(predicates.CheckVolumeBindingPred)
		}
		invalidPredicates.Insert(maxPDVolumeCountPredicateKeys...)
		if utilfeature.DefaultFeatureGate.Enabled(features.AttachVolumeLimit) {
			invalidPredicates.Insert(predicates.MaxCSIVolumeCountPred)
		}
	}
	c.equivalencePodCache.InvalidatePredicates(invalidPredicates)
}
func (c *configFactory) onStorageClassAdd(obj interface{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	sc, ok := obj.(*storagev1.StorageClass)
	if !ok {
		klog.Errorf("cannot convert to *storagev1.StorageClass: %v", obj)
		return
	}
	if sc.VolumeBindingMode != nil && *sc.VolumeBindingMode == storagev1.VolumeBindingWaitForFirstConsumer {
		c.podQueue.MoveAllToActiveQueue()
	}
}
func (c *configFactory) onStorageClassDelete(obj interface{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if c.enableEquivalenceClassCache {
		var sc *storagev1.StorageClass
		switch t := obj.(type) {
		case *storagev1.StorageClass:
			sc = t
		case cache.DeletedFinalStateUnknown:
			var ok bool
			sc, ok = t.Obj.(*storagev1.StorageClass)
			if !ok {
				klog.Errorf("cannot convert to *storagev1.StorageClass: %v", t.Obj)
				return
			}
		default:
			klog.Errorf("cannot convert to *storagev1.StorageClass: %v", t)
			return
		}
		c.invalidatePredicatesForStorageClass(sc)
	}
}
func (c *configFactory) invalidatePredicatesForStorageClass(sc *storagev1.StorageClass) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	invalidPredicates := sets.NewString()
	if utilfeature.DefaultFeatureGate.Enabled(features.VolumeScheduling) {
		if sc.VolumeBindingMode != nil && *sc.VolumeBindingMode == storagev1.VolumeBindingWaitForFirstConsumer {
			invalidPredicates.Insert(predicates.CheckVolumeBindingPred)
			invalidPredicates.Insert(predicates.NoVolumeZoneConflictPred)
		}
	}
	c.equivalencePodCache.InvalidatePredicates(invalidPredicates)
}
func (c *configFactory) onServiceAdd(obj interface{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if c.enableEquivalenceClassCache {
		c.equivalencePodCache.InvalidatePredicates(serviceAffinitySet)
	}
	c.podQueue.MoveAllToActiveQueue()
}
func (c *configFactory) onServiceUpdate(oldObj interface{}, newObj interface{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if c.enableEquivalenceClassCache {
		oldService := oldObj.(*v1.Service)
		newService := newObj.(*v1.Service)
		if !reflect.DeepEqual(oldService.Spec.Selector, newService.Spec.Selector) {
			c.equivalencePodCache.InvalidatePredicates(serviceAffinitySet)
		}
	}
	c.podQueue.MoveAllToActiveQueue()
}
func (c *configFactory) onServiceDelete(obj interface{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if c.enableEquivalenceClassCache {
		c.equivalencePodCache.InvalidatePredicates(serviceAffinitySet)
	}
	c.podQueue.MoveAllToActiveQueue()
}
func (c *configFactory) GetNodeLister() corelisters.NodeLister {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return c.nodeLister
}
func (c *configFactory) GetHardPodAffinitySymmetricWeight() int32 {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return c.hardPodAffinitySymmetricWeight
}
func (c *configFactory) GetSchedulerName() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return c.schedulerName
}
func (c *configFactory) GetClient() clientset.Interface {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return c.client
}
func (c *configFactory) GetScheduledPodLister() corelisters.PodLister {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return c.scheduledPodLister
}
func (c *configFactory) addPodToCache(obj interface{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	pod, ok := obj.(*v1.Pod)
	if !ok {
		klog.Errorf("cannot convert to *v1.Pod: %v", obj)
		return
	}
	if err := c.schedulerCache.AddPod(pod); err != nil {
		klog.Errorf("scheduler cache AddPod failed: %v", err)
	}
	c.podQueue.AssignedPodAdded(pod)
}
func (c *configFactory) updatePodInCache(oldObj, newObj interface{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	oldPod, ok := oldObj.(*v1.Pod)
	if !ok {
		klog.Errorf("cannot convert oldObj to *v1.Pod: %v", oldObj)
		return
	}
	newPod, ok := newObj.(*v1.Pod)
	if !ok {
		klog.Errorf("cannot convert newObj to *v1.Pod: %v", newObj)
		return
	}
	if err := c.schedulerCache.UpdatePod(oldPod, newPod); err != nil {
		klog.Errorf("scheduler cache UpdatePod failed: %v", err)
	}
	c.invalidateCachedPredicatesOnUpdatePod(newPod, oldPod)
	c.podQueue.AssignedPodUpdated(newPod)
}
func (c *configFactory) addPodToSchedulingQueue(obj interface{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if err := c.podQueue.Add(obj.(*v1.Pod)); err != nil {
		runtime.HandleError(fmt.Errorf("unable to queue %T: %v", obj, err))
	}
}
func (c *configFactory) updatePodInSchedulingQueue(oldObj, newObj interface{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	pod := newObj.(*v1.Pod)
	if c.skipPodUpdate(pod) {
		return
	}
	if err := c.podQueue.Update(oldObj.(*v1.Pod), pod); err != nil {
		runtime.HandleError(fmt.Errorf("unable to update %T: %v", newObj, err))
	}
}
func (c *configFactory) deletePodFromSchedulingQueue(obj interface{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var pod *v1.Pod
	switch t := obj.(type) {
	case *v1.Pod:
		pod = obj.(*v1.Pod)
	case cache.DeletedFinalStateUnknown:
		var ok bool
		pod, ok = t.Obj.(*v1.Pod)
		if !ok {
			runtime.HandleError(fmt.Errorf("unable to convert object %T to *v1.Pod in %T", obj, c))
			return
		}
	default:
		runtime.HandleError(fmt.Errorf("unable to handle object in %T: %T", c, obj))
		return
	}
	if err := c.podQueue.Delete(pod); err != nil {
		runtime.HandleError(fmt.Errorf("unable to dequeue %T: %v", obj, err))
	}
	if c.volumeBinder != nil {
		c.volumeBinder.DeletePodBindings(pod)
	}
}
func (c *configFactory) invalidateCachedPredicatesOnUpdatePod(newPod *v1.Pod, oldPod *v1.Pod) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if c.enableEquivalenceClassCache {
		if len(newPod.Spec.NodeName) != 0 && newPod.Spec.NodeName == oldPod.Spec.NodeName {
			if !reflect.DeepEqual(oldPod.GetLabels(), newPod.GetLabels()) {
				c.equivalencePodCache.InvalidatePredicates(matchInterPodAffinitySet)
			}
			if !reflect.DeepEqual(predicates.GetResourceRequest(newPod), predicates.GetResourceRequest(oldPod)) {
				c.equivalencePodCache.InvalidatePredicatesOnNode(newPod.Spec.NodeName, generalPredicatesSets)
			}
		}
	}
}
func (c *configFactory) deletePodFromCache(obj interface{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var pod *v1.Pod
	switch t := obj.(type) {
	case *v1.Pod:
		pod = t
	case cache.DeletedFinalStateUnknown:
		var ok bool
		pod, ok = t.Obj.(*v1.Pod)
		if !ok {
			klog.Errorf("cannot convert to *v1.Pod: %v", t.Obj)
			return
		}
	default:
		klog.Errorf("cannot convert to *v1.Pod: %v", t)
		return
	}
	if err := c.schedulerCache.RemovePod(pod); err != nil {
		klog.Errorf("scheduler cache RemovePod failed: %v", err)
	}
	c.invalidateCachedPredicatesOnDeletePod(pod)
	c.podQueue.MoveAllToActiveQueue()
}
func (c *configFactory) invalidateCachedPredicatesOnDeletePod(pod *v1.Pod) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if c.enableEquivalenceClassCache {
		c.equivalencePodCache.InvalidateCachedPredicateItemForPodAdd(pod, pod.Spec.NodeName)
		c.equivalencePodCache.InvalidatePredicates(matchInterPodAffinitySet)
		for _, volume := range pod.Spec.Volumes {
			if volume.GCEPersistentDisk != nil || volume.AWSElasticBlockStore != nil || volume.RBD != nil || volume.ISCSI != nil {
				c.equivalencePodCache.InvalidatePredicatesOnNode(pod.Spec.NodeName, noDiskConflictSet)
			}
		}
	}
}
func (c *configFactory) addNodeToCache(obj interface{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	node, ok := obj.(*v1.Node)
	if !ok {
		klog.Errorf("cannot convert to *v1.Node: %v", obj)
		return
	}
	if c.enableEquivalenceClassCache {
		c.equivalencePodCache.GetNodeCache(node.GetName())
	}
	if err := c.schedulerCache.AddNode(node); err != nil {
		klog.Errorf("scheduler cache AddNode failed: %v", err)
	}
	c.podQueue.MoveAllToActiveQueue()
}
func (c *configFactory) updateNodeInCache(oldObj, newObj interface{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	oldNode, ok := oldObj.(*v1.Node)
	if !ok {
		klog.Errorf("cannot convert oldObj to *v1.Node: %v", oldObj)
		return
	}
	newNode, ok := newObj.(*v1.Node)
	if !ok {
		klog.Errorf("cannot convert newObj to *v1.Node: %v", newObj)
		return
	}
	if err := c.schedulerCache.UpdateNode(oldNode, newNode); err != nil {
		klog.Errorf("scheduler cache UpdateNode failed: %v", err)
	}
	c.invalidateCachedPredicatesOnNodeUpdate(newNode, oldNode)
	if c.podQueue.NumUnschedulablePods() == 0 || nodeSchedulingPropertiesChanged(newNode, oldNode) {
		c.podQueue.MoveAllToActiveQueue()
	}
}
func (c *configFactory) invalidateCachedPredicatesOnNodeUpdate(newNode *v1.Node, oldNode *v1.Node) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if c.enableEquivalenceClassCache {
		invalidPredicates := sets.NewString()
		if !reflect.DeepEqual(oldNode.Status.Allocatable, newNode.Status.Allocatable) {
			invalidPredicates.Insert(predicates.GeneralPred)
		}
		if !reflect.DeepEqual(oldNode.GetLabels(), newNode.GetLabels()) {
			invalidPredicates.Insert(predicates.GeneralPred, predicates.CheckServiceAffinityPred)
			for k, v := range oldNode.GetLabels() {
				if v != newNode.GetLabels()[k] {
					invalidPredicates.Insert(predicates.MatchInterPodAffinityPred)
				}
				if isZoneRegionLabel(k) {
					if v != newNode.GetLabels()[k] {
						invalidPredicates.Insert(predicates.NoVolumeZoneConflictPred)
					}
				}
			}
		}
		oldTaints, oldErr := helper.GetTaintsFromNodeAnnotations(oldNode.GetAnnotations())
		if oldErr != nil {
			klog.Errorf("Failed to get taints from old node annotation for equivalence cache")
		}
		newTaints, newErr := helper.GetTaintsFromNodeAnnotations(newNode.GetAnnotations())
		if newErr != nil {
			klog.Errorf("Failed to get taints from new node annotation for equivalence cache")
		}
		if !reflect.DeepEqual(oldTaints, newTaints) || !reflect.DeepEqual(oldNode.Spec.Taints, newNode.Spec.Taints) {
			invalidPredicates.Insert(predicates.PodToleratesNodeTaintsPred)
		}
		if !reflect.DeepEqual(oldNode.Status.Conditions, newNode.Status.Conditions) {
			oldConditions := make(map[v1.NodeConditionType]v1.ConditionStatus)
			newConditions := make(map[v1.NodeConditionType]v1.ConditionStatus)
			for _, cond := range oldNode.Status.Conditions {
				oldConditions[cond.Type] = cond.Status
			}
			for _, cond := range newNode.Status.Conditions {
				newConditions[cond.Type] = cond.Status
			}
			if oldConditions[v1.NodeMemoryPressure] != newConditions[v1.NodeMemoryPressure] {
				invalidPredicates.Insert(predicates.CheckNodeMemoryPressurePred)
			}
			if oldConditions[v1.NodeDiskPressure] != newConditions[v1.NodeDiskPressure] {
				invalidPredicates.Insert(predicates.CheckNodeDiskPressurePred)
			}
			if oldConditions[v1.NodePIDPressure] != newConditions[v1.NodePIDPressure] {
				invalidPredicates.Insert(predicates.CheckNodePIDPressurePred)
			}
			if oldConditions[v1.NodeReady] != newConditions[v1.NodeReady] || oldConditions[v1.NodeOutOfDisk] != newConditions[v1.NodeOutOfDisk] || oldConditions[v1.NodeNetworkUnavailable] != newConditions[v1.NodeNetworkUnavailable] {
				invalidPredicates.Insert(predicates.CheckNodeConditionPred)
			}
		}
		if newNode.Spec.Unschedulable != oldNode.Spec.Unschedulable {
			invalidPredicates.Insert(predicates.CheckNodeConditionPred)
		}
		c.equivalencePodCache.InvalidatePredicatesOnNode(newNode.GetName(), invalidPredicates)
	}
}
func nodeSchedulingPropertiesChanged(newNode *v1.Node, oldNode *v1.Node) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if nodeSpecUnschedulableChanged(newNode, oldNode) {
		return true
	}
	if nodeAllocatableChanged(newNode, oldNode) {
		return true
	}
	if nodeLabelsChanged(newNode, oldNode) {
		return true
	}
	if nodeTaintsChanged(newNode, oldNode) {
		return true
	}
	if nodeConditionsChanged(newNode, oldNode) {
		return true
	}
	return false
}
func nodeAllocatableChanged(newNode *v1.Node, oldNode *v1.Node) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return !reflect.DeepEqual(oldNode.Status.Allocatable, newNode.Status.Allocatable)
}
func nodeLabelsChanged(newNode *v1.Node, oldNode *v1.Node) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return !reflect.DeepEqual(oldNode.GetLabels(), newNode.GetLabels())
}
func nodeTaintsChanged(newNode *v1.Node, oldNode *v1.Node) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return !reflect.DeepEqual(newNode.Spec.Taints, oldNode.Spec.Taints)
}
func nodeConditionsChanged(newNode *v1.Node, oldNode *v1.Node) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	strip := func(conditions []v1.NodeCondition) map[v1.NodeConditionType]v1.ConditionStatus {
		conditionStatuses := make(map[v1.NodeConditionType]v1.ConditionStatus, len(conditions))
		for i := range conditions {
			conditionStatuses[conditions[i].Type] = conditions[i].Status
		}
		return conditionStatuses
	}
	return !reflect.DeepEqual(strip(oldNode.Status.Conditions), strip(newNode.Status.Conditions))
}
func nodeSpecUnschedulableChanged(newNode *v1.Node, oldNode *v1.Node) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return newNode.Spec.Unschedulable != oldNode.Spec.Unschedulable && newNode.Spec.Unschedulable == false
}
func (c *configFactory) deleteNodeFromCache(obj interface{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var node *v1.Node
	switch t := obj.(type) {
	case *v1.Node:
		node = t
	case cache.DeletedFinalStateUnknown:
		var ok bool
		node, ok = t.Obj.(*v1.Node)
		if !ok {
			klog.Errorf("cannot convert to *v1.Node: %v", t.Obj)
			return
		}
	default:
		klog.Errorf("cannot convert to *v1.Node: %v", t)
		return
	}
	if err := c.schedulerCache.RemoveNode(node); err != nil {
		klog.Errorf("scheduler cache RemoveNode failed: %v", err)
	}
	if c.enableEquivalenceClassCache {
		c.equivalencePodCache.InvalidateAllPredicatesOnNode(node.GetName())
	}
}
func (c *configFactory) Create() (*Config, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return c.CreateFromProvider(DefaultProvider)
}
func (c *configFactory) CreateFromProvider(providerName string) (*Config, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(2).Infof("Creating scheduler from algorithm provider '%v'", providerName)
	provider, err := GetAlgorithmProvider(providerName)
	if err != nil {
		return nil, err
	}
	return c.CreateFromKeys(provider.FitPredicateKeys, provider.PriorityFunctionKeys, []algorithm.SchedulerExtender{})
}
func (c *configFactory) CreateFromConfig(policy schedulerapi.Policy) (*Config, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(2).Infof("Creating scheduler from configuration: %v", policy)
	if err := validation.ValidatePolicy(policy); err != nil {
		return nil, err
	}
	predicateKeys := sets.NewString()
	if policy.Predicates == nil {
		klog.V(2).Infof("Using predicates from algorithm provider '%v'", DefaultProvider)
		provider, err := GetAlgorithmProvider(DefaultProvider)
		if err != nil {
			return nil, err
		}
		predicateKeys = provider.FitPredicateKeys
	} else {
		for _, predicate := range policy.Predicates {
			klog.V(2).Infof("Registering predicate: %s", predicate.Name)
			predicateKeys.Insert(RegisterCustomFitPredicate(predicate))
		}
	}
	priorityKeys := sets.NewString()
	if policy.Priorities == nil {
		klog.V(2).Infof("Using priorities from algorithm provider '%v'", DefaultProvider)
		provider, err := GetAlgorithmProvider(DefaultProvider)
		if err != nil {
			return nil, err
		}
		priorityKeys = provider.PriorityFunctionKeys
	} else {
		for _, priority := range policy.Priorities {
			klog.V(2).Infof("Registering priority: %s", priority.Name)
			priorityKeys.Insert(RegisterCustomPriorityFunction(priority))
		}
	}
	var extenders []algorithm.SchedulerExtender
	if len(policy.ExtenderConfigs) != 0 {
		ignoredExtendedResources := sets.NewString()
		for ii := range policy.ExtenderConfigs {
			klog.V(2).Infof("Creating extender with config %+v", policy.ExtenderConfigs[ii])
			extender, err := core.NewHTTPExtender(&policy.ExtenderConfigs[ii])
			if err != nil {
				return nil, err
			}
			extenders = append(extenders, extender)
			for _, r := range policy.ExtenderConfigs[ii].ManagedResources {
				if r.IgnoredByScheduler {
					ignoredExtendedResources.Insert(string(r.Name))
				}
			}
		}
		predicates.RegisterPredicateMetadataProducerWithExtendedResourceOptions(ignoredExtendedResources)
	}
	if policy.HardPodAffinitySymmetricWeight != 0 {
		c.hardPodAffinitySymmetricWeight = policy.HardPodAffinitySymmetricWeight
	}
	if policy.AlwaysCheckAllPredicates {
		c.alwaysCheckAllPredicates = policy.AlwaysCheckAllPredicates
	}
	return c.CreateFromKeys(predicateKeys, priorityKeys, extenders)
}
func (c *configFactory) getBinderFunc(extenders []algorithm.SchedulerExtender) func(pod *v1.Pod) Binder {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var extenderBinder algorithm.SchedulerExtender
	for i := range extenders {
		if extenders[i].IsBinder() {
			extenderBinder = extenders[i]
			break
		}
	}
	defaultBinder := &binder{c.client}
	return func(pod *v1.Pod) Binder {
		if extenderBinder != nil && extenderBinder.IsInterested(pod) {
			return extenderBinder
		}
		return defaultBinder
	}
}
func (c *configFactory) CreateFromKeys(predicateKeys, priorityKeys sets.String, extenders []algorithm.SchedulerExtender) (*Config, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(2).Infof("Creating scheduler with fit predicates '%v' and priority functions '%v'", predicateKeys, priorityKeys)
	if c.GetHardPodAffinitySymmetricWeight() < 1 || c.GetHardPodAffinitySymmetricWeight() > 100 {
		return nil, fmt.Errorf("invalid hardPodAffinitySymmetricWeight: %d, must be in the range 1-100", c.GetHardPodAffinitySymmetricWeight())
	}
	predicateFuncs, err := c.GetPredicates(predicateKeys)
	if err != nil {
		return nil, err
	}
	priorityConfigs, err := c.GetPriorityFunctionConfigs(priorityKeys)
	if err != nil {
		return nil, err
	}
	priorityMetaProducer, err := c.GetPriorityMetadataProducer()
	if err != nil {
		return nil, err
	}
	predicateMetaProducer, err := c.GetPredicateMetadataProducer()
	if err != nil {
		return nil, err
	}
	if c.enableEquivalenceClassCache {
		c.equivalencePodCache = equivalence.NewCache(predicates.Ordering())
		klog.Info("Created equivalence class cache")
	}
	algo := core.NewGenericScheduler(c.schedulerCache, c.equivalencePodCache, c.podQueue, predicateFuncs, predicateMetaProducer, priorityConfigs, priorityMetaProducer, extenders, c.volumeBinder, c.pVCLister, c.pdbLister, c.alwaysCheckAllPredicates, c.disablePreemption, c.percentageOfNodesToScore)
	podBackoff := util.CreateDefaultPodBackoff()
	return &Config{SchedulerCache: c.schedulerCache, Ecache: c.equivalencePodCache, NodeLister: &nodeLister{c.nodeLister}, Algorithm: algo, GetBinder: c.getBinderFunc(extenders), PodConditionUpdater: &podConditionUpdater{c.client}, PodPreemptor: &podPreemptor{c.client}, WaitForCacheSync: func() bool {
		return cache.WaitForCacheSync(c.StopEverything, c.scheduledPodsHasSynced)
	}, NextPod: func() *v1.Pod {
		return c.getNextPod()
	}, Error: c.MakeDefaultErrorFunc(podBackoff, c.podQueue), StopEverything: c.StopEverything, VolumeBinder: c.volumeBinder, SchedulingQueue: c.podQueue}, nil
}

type nodeLister struct{ corelisters.NodeLister }

func (n *nodeLister) List() ([]*v1.Node, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return n.NodeLister.List(labels.Everything())
}
func (c *configFactory) GetPriorityFunctionConfigs(priorityKeys sets.String) ([]algorithm.PriorityConfig, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	pluginArgs, err := c.getPluginArgs()
	if err != nil {
		return nil, err
	}
	return getPriorityFunctionConfigs(priorityKeys, *pluginArgs)
}
func (c *configFactory) GetPriorityMetadataProducer() (algorithm.PriorityMetadataProducer, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	pluginArgs, err := c.getPluginArgs()
	if err != nil {
		return nil, err
	}
	return getPriorityMetadataProducer(*pluginArgs)
}
func (c *configFactory) GetPredicateMetadataProducer() (algorithm.PredicateMetadataProducer, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	pluginArgs, err := c.getPluginArgs()
	if err != nil {
		return nil, err
	}
	return getPredicateMetadataProducer(*pluginArgs)
}
func (c *configFactory) GetPredicates(predicateKeys sets.String) (map[string]algorithm.FitPredicate, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	pluginArgs, err := c.getPluginArgs()
	if err != nil {
		return nil, err
	}
	return getFitPredicateFunctions(predicateKeys, *pluginArgs)
}
func (c *configFactory) getPluginArgs() (*PluginFactoryArgs, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &PluginFactoryArgs{PodLister: c.podLister, ServiceLister: c.serviceLister, ControllerLister: c.controllerLister, ReplicaSetLister: c.replicaSetLister, StatefulSetLister: c.statefulSetLister, NodeLister: &nodeLister{c.nodeLister}, PDBLister: c.pdbLister, NodeInfo: &predicates.CachedNodeInfo{NodeLister: c.nodeLister}, PVInfo: &predicates.CachedPersistentVolumeInfo{PersistentVolumeLister: c.pVLister}, PVCInfo: &predicates.CachedPersistentVolumeClaimInfo{PersistentVolumeClaimLister: c.pVCLister}, StorageClassInfo: &predicates.CachedStorageClassInfo{StorageClassLister: c.storageClassLister}, VolumeBinder: c.volumeBinder, HardPodAffinitySymmetricWeight: c.hardPodAffinitySymmetricWeight}, nil
}
func (c *configFactory) getNextPod() *v1.Pod {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	pod, err := c.podQueue.Pop()
	if err == nil {
		klog.V(4).Infof("About to try and schedule pod %v/%v", pod.Namespace, pod.Name)
		return pod
	}
	klog.Errorf("Error while retrieving next pod from scheduling queue: %v", err)
	return nil
}
func assignedPod(pod *v1.Pod) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return len(pod.Spec.NodeName) != 0
}
func responsibleForPod(pod *v1.Pod, schedulerName string) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return schedulerName == pod.Spec.SchedulerName
}

type assignedPodLister struct{ corelisters.PodLister }

func (l assignedPodLister) List(selector labels.Selector) ([]*v1.Pod, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	list, err := l.PodLister.List(selector)
	if err != nil {
		return nil, err
	}
	filtered := make([]*v1.Pod, 0, len(list))
	for _, pod := range list {
		if len(pod.Spec.NodeName) > 0 {
			filtered = append(filtered, pod)
		}
	}
	return filtered, nil
}
func (l assignedPodLister) Pods(namespace string) corelisters.PodNamespaceLister {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return assignedPodNamespaceLister{l.PodLister.Pods(namespace)}
}

type assignedPodNamespaceLister struct{ corelisters.PodNamespaceLister }

func (l assignedPodNamespaceLister) List(selector labels.Selector) (ret []*v1.Pod, err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	list, err := l.PodNamespaceLister.List(selector)
	if err != nil {
		return nil, err
	}
	filtered := make([]*v1.Pod, 0, len(list))
	for _, pod := range list {
		if len(pod.Spec.NodeName) > 0 {
			filtered = append(filtered, pod)
		}
	}
	return filtered, nil
}
func (l assignedPodNamespaceLister) Get(name string) (*v1.Pod, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	pod, err := l.PodNamespaceLister.Get(name)
	if err != nil {
		return nil, err
	}
	if len(pod.Spec.NodeName) > 0 {
		return pod, nil
	}
	return nil, errors.NewNotFound(schema.GroupResource{Resource: string(v1.ResourcePods)}, name)
}

type podInformer struct{ informer cache.SharedIndexInformer }

func (i *podInformer) Informer() cache.SharedIndexInformer {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return i.informer
}
func (i *podInformer) Lister() corelisters.PodLister {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return corelisters.NewPodLister(i.informer.GetIndexer())
}
func NewPodInformer(client clientset.Interface, resyncPeriod time.Duration) coreinformers.PodInformer {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	selector := fields.ParseSelectorOrDie("status.phase!=" + string(v1.PodSucceeded) + ",status.phase!=" + string(v1.PodFailed))
	lw := cache.NewListWatchFromClient(client.CoreV1().RESTClient(), string(v1.ResourcePods), metav1.NamespaceAll, selector)
	return &podInformer{informer: cache.NewSharedIndexInformer(lw, &v1.Pod{}, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc})}
}
func (c *configFactory) MakeDefaultErrorFunc(backoff *util.PodBackoff, podQueue internalqueue.SchedulingQueue) func(pod *v1.Pod, err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return func(pod *v1.Pod, err error) {
		if err == core.ErrNoNodesAvailable {
			klog.V(4).Infof("Unable to schedule %v/%v: no nodes are registered to the cluster; waiting", pod.Namespace, pod.Name)
		} else {
			if _, ok := err.(*core.FitError); ok {
				klog.V(4).Infof("Unable to schedule %v/%v: no fit: %v; waiting", pod.Namespace, pod.Name, err)
			} else if errors.IsNotFound(err) {
				if errStatus, ok := err.(errors.APIStatus); ok && errStatus.Status().Details.Kind == "node" {
					nodeName := errStatus.Status().Details.Name
					_, err := c.client.CoreV1().Nodes().Get(nodeName, metav1.GetOptions{})
					if err != nil && errors.IsNotFound(err) {
						node := v1.Node{ObjectMeta: metav1.ObjectMeta{Name: nodeName}}
						c.schedulerCache.RemoveNode(&node)
						if c.enableEquivalenceClassCache {
							c.equivalencePodCache.InvalidateAllPredicatesOnNode(nodeName)
						}
					}
				}
			} else {
				klog.Errorf("Error scheduling %v/%v: %v; retrying", pod.Namespace, pod.Name, err)
			}
		}
		backoff.Gc()
		podSchedulingCycle := podQueue.SchedulingCycle()
		go func() {
			defer runtime.HandleCrash()
			podID := types.NamespacedName{Namespace: pod.Namespace, Name: pod.Name}
			origPod := pod
			if !util.PodPriorityEnabled() {
				entry := backoff.GetEntry(podID)
				if !entry.TryWait(backoff.MaxDuration()) {
					klog.Warningf("Request for pod %v already in flight, abandoning", podID)
					return
				}
			}
			getBackoff := initialGetBackoff
			for {
				pod, err := c.client.CoreV1().Pods(podID.Namespace).Get(podID.Name, metav1.GetOptions{})
				if err == nil {
					if len(pod.Spec.NodeName) == 0 {
						podQueue.AddUnschedulableIfNotPresent(pod, podSchedulingCycle)
					} else {
						if c.volumeBinder != nil {
							c.volumeBinder.DeletePodBindings(pod)
						}
					}
					break
				}
				if errors.IsNotFound(err) {
					klog.Warningf("A pod %v no longer exists", podID)
					if c.volumeBinder != nil {
						c.volumeBinder.DeletePodBindings(origPod)
					}
					return
				}
				klog.Errorf("Error getting pod %v for retry: %v; retrying...", podID, err)
				if getBackoff = getBackoff * 2; getBackoff > maximalGetBackoff {
					getBackoff = maximalGetBackoff
				}
				time.Sleep(getBackoff)
			}
		}()
	}
}

type nodeEnumerator struct{ *v1.NodeList }

func (ne *nodeEnumerator) Len() int {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if ne.NodeList == nil {
		return 0
	}
	return len(ne.Items)
}
func (ne *nodeEnumerator) Get(index int) interface{} {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &ne.Items[index]
}

type binder struct{ Client clientset.Interface }

func (b *binder) Bind(binding *v1.Binding) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(3).Infof("Attempting to bind %v to %v", binding.Name, binding.Target.Name)
	return b.Client.CoreV1().Pods(binding.Namespace).Bind(binding)
}

type podConditionUpdater struct{ Client clientset.Interface }

func (p *podConditionUpdater) Update(pod *v1.Pod, condition *v1.PodCondition) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(3).Infof("Updating pod condition for %s/%s to (%s==%s)", pod.Namespace, pod.Name, condition.Type, condition.Status)
	if podutil.UpdatePodCondition(&pod.Status, condition) {
		_, err := p.Client.CoreV1().Pods(pod.Namespace).UpdateStatus(pod)
		return err
	}
	return nil
}

type podPreemptor struct{ Client clientset.Interface }

func (p *podPreemptor) GetUpdatedPod(pod *v1.Pod) (*v1.Pod, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return p.Client.CoreV1().Pods(pod.Namespace).Get(pod.Name, metav1.GetOptions{})
}
func (p *podPreemptor) DeletePod(pod *v1.Pod) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return p.Client.CoreV1().Pods(pod.Namespace).Delete(pod.Name, &metav1.DeleteOptions{})
}
func (p *podPreemptor) SetNominatedNodeName(pod *v1.Pod, nominatedNodeName string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	podCopy := pod.DeepCopy()
	podCopy.Status.NominatedNodeName = nominatedNodeName
	_, err := p.Client.CoreV1().Pods(pod.Namespace).UpdateStatus(podCopy)
	return err
}
func (p *podPreemptor) RemoveNominatedNodeName(pod *v1.Pod) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(pod.Status.NominatedNodeName) == 0 {
		return nil
	}
	return p.SetNominatedNodeName(pod, "")
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
