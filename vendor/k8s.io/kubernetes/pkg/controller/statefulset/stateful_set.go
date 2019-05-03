package statefulset

import (
 "fmt"
 "reflect"
 "time"
 apps "k8s.io/api/apps/v1"
 "k8s.io/api/core/v1"
 "k8s.io/apimachinery/pkg/api/errors"
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 "k8s.io/apimachinery/pkg/labels"
 utilruntime "k8s.io/apimachinery/pkg/util/runtime"
 "k8s.io/apimachinery/pkg/util/wait"
 appsinformers "k8s.io/client-go/informers/apps/v1"
 coreinformers "k8s.io/client-go/informers/core/v1"
 clientset "k8s.io/client-go/kubernetes"
 "k8s.io/client-go/kubernetes/scheme"
 v1core "k8s.io/client-go/kubernetes/typed/core/v1"
 appslisters "k8s.io/client-go/listers/apps/v1"
 corelisters "k8s.io/client-go/listers/core/v1"
 "k8s.io/client-go/tools/cache"
 "k8s.io/client-go/tools/record"
 "k8s.io/client-go/util/workqueue"
 "k8s.io/kubernetes/pkg/controller"
 "k8s.io/kubernetes/pkg/controller/history"
 "k8s.io/klog"
)

const (
 statefulSetResyncPeriod = 30 * time.Second
)

var controllerKind = apps.SchemeGroupVersion.WithKind("StatefulSet")

type StatefulSetController struct {
 kubeClient      clientset.Interface
 control         StatefulSetControlInterface
 podControl      controller.PodControlInterface
 podLister       corelisters.PodLister
 podListerSynced cache.InformerSynced
 setLister       appslisters.StatefulSetLister
 setListerSynced cache.InformerSynced
 pvcListerSynced cache.InformerSynced
 revListerSynced cache.InformerSynced
 queue           workqueue.RateLimitingInterface
}

func NewStatefulSetController(podInformer coreinformers.PodInformer, setInformer appsinformers.StatefulSetInformer, pvcInformer coreinformers.PersistentVolumeClaimInformer, revInformer appsinformers.ControllerRevisionInformer, kubeClient clientset.Interface) *StatefulSetController {
 _logClusterCodePath()
 defer _logClusterCodePath()
 eventBroadcaster := record.NewBroadcaster()
 eventBroadcaster.StartLogging(klog.Infof)
 eventBroadcaster.StartRecordingToSink(&v1core.EventSinkImpl{Interface: kubeClient.CoreV1().Events("")})
 recorder := eventBroadcaster.NewRecorder(scheme.Scheme, v1.EventSource{Component: "statefulset-controller"})
 ssc := &StatefulSetController{kubeClient: kubeClient, control: NewDefaultStatefulSetControl(NewRealStatefulPodControl(kubeClient, setInformer.Lister(), podInformer.Lister(), pvcInformer.Lister(), recorder), NewRealStatefulSetStatusUpdater(kubeClient, setInformer.Lister()), history.NewHistory(kubeClient, revInformer.Lister()), recorder), pvcListerSynced: pvcInformer.Informer().HasSynced, queue: workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "statefulset"), podControl: controller.RealPodControl{KubeClient: kubeClient, Recorder: recorder}, revListerSynced: revInformer.Informer().HasSynced}
 podInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{AddFunc: ssc.addPod, UpdateFunc: ssc.updatePod, DeleteFunc: ssc.deletePod})
 ssc.podLister = podInformer.Lister()
 ssc.podListerSynced = podInformer.Informer().HasSynced
 setInformer.Informer().AddEventHandlerWithResyncPeriod(cache.ResourceEventHandlerFuncs{AddFunc: ssc.enqueueStatefulSet, UpdateFunc: func(old, cur interface{}) {
  oldPS := old.(*apps.StatefulSet)
  curPS := cur.(*apps.StatefulSet)
  if oldPS.Status.Replicas != curPS.Status.Replicas {
   klog.V(4).Infof("Observed updated replica count for StatefulSet: %v, %d->%d", curPS.Name, oldPS.Status.Replicas, curPS.Status.Replicas)
  }
  ssc.enqueueStatefulSet(cur)
 }, DeleteFunc: ssc.enqueueStatefulSet}, statefulSetResyncPeriod)
 ssc.setLister = setInformer.Lister()
 ssc.setListerSynced = setInformer.Informer().HasSynced
 return ssc
}
func (ssc *StatefulSetController) Run(workers int, stopCh <-chan struct{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 defer utilruntime.HandleCrash()
 defer ssc.queue.ShutDown()
 klog.Infof("Starting stateful set controller")
 defer klog.Infof("Shutting down statefulset controller")
 if !controller.WaitForCacheSync("stateful set", stopCh, ssc.podListerSynced, ssc.setListerSynced, ssc.pvcListerSynced, ssc.revListerSynced) {
  return
 }
 for i := 0; i < workers; i++ {
  go wait.Until(ssc.worker, time.Second, stopCh)
 }
 <-stopCh
}
func (ssc *StatefulSetController) addPod(obj interface{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 pod := obj.(*v1.Pod)
 if pod.DeletionTimestamp != nil {
  ssc.deletePod(pod)
  return
 }
 if controllerRef := metav1.GetControllerOf(pod); controllerRef != nil {
  set := ssc.resolveControllerRef(pod.Namespace, controllerRef)
  if set == nil {
   return
  }
  klog.V(4).Infof("Pod %s created, labels: %+v", pod.Name, pod.Labels)
  ssc.enqueueStatefulSet(set)
  return
 }
 sets := ssc.getStatefulSetsForPod(pod)
 if len(sets) == 0 {
  return
 }
 klog.V(4).Infof("Orphan Pod %s created, labels: %+v", pod.Name, pod.Labels)
 for _, set := range sets {
  ssc.enqueueStatefulSet(set)
 }
}
func (ssc *StatefulSetController) updatePod(old, cur interface{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 curPod := cur.(*v1.Pod)
 oldPod := old.(*v1.Pod)
 if curPod.ResourceVersion == oldPod.ResourceVersion {
  return
 }
 labelChanged := !reflect.DeepEqual(curPod.Labels, oldPod.Labels)
 curControllerRef := metav1.GetControllerOf(curPod)
 oldControllerRef := metav1.GetControllerOf(oldPod)
 controllerRefChanged := !reflect.DeepEqual(curControllerRef, oldControllerRef)
 if controllerRefChanged && oldControllerRef != nil {
  if set := ssc.resolveControllerRef(oldPod.Namespace, oldControllerRef); set != nil {
   ssc.enqueueStatefulSet(set)
  }
 }
 if curControllerRef != nil {
  set := ssc.resolveControllerRef(curPod.Namespace, curControllerRef)
  if set == nil {
   return
  }
  klog.V(4).Infof("Pod %s updated, objectMeta %+v -> %+v.", curPod.Name, oldPod.ObjectMeta, curPod.ObjectMeta)
  ssc.enqueueStatefulSet(set)
  return
 }
 if labelChanged || controllerRefChanged {
  sets := ssc.getStatefulSetsForPod(curPod)
  if len(sets) == 0 {
   return
  }
  klog.V(4).Infof("Orphan Pod %s updated, objectMeta %+v -> %+v.", curPod.Name, oldPod.ObjectMeta, curPod.ObjectMeta)
  for _, set := range sets {
   ssc.enqueueStatefulSet(set)
  }
 }
}
func (ssc *StatefulSetController) deletePod(obj interface{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 pod, ok := obj.(*v1.Pod)
 if !ok {
  tombstone, ok := obj.(cache.DeletedFinalStateUnknown)
  if !ok {
   utilruntime.HandleError(fmt.Errorf("couldn't get object from tombstone %+v", obj))
   return
  }
  pod, ok = tombstone.Obj.(*v1.Pod)
  if !ok {
   utilruntime.HandleError(fmt.Errorf("tombstone contained object that is not a pod %+v", obj))
   return
  }
 }
 controllerRef := metav1.GetControllerOf(pod)
 if controllerRef == nil {
  return
 }
 set := ssc.resolveControllerRef(pod.Namespace, controllerRef)
 if set == nil {
  return
 }
 klog.V(4).Infof("Pod %s/%s deleted through %v.", pod.Namespace, pod.Name, utilruntime.GetCaller())
 ssc.enqueueStatefulSet(set)
}
func (ssc *StatefulSetController) getPodsForStatefulSet(set *apps.StatefulSet, selector labels.Selector) ([]*v1.Pod, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 pods, err := ssc.podLister.Pods(set.Namespace).List(labels.Everything())
 if err != nil {
  return nil, err
 }
 filter := func(pod *v1.Pod) bool {
  return isMemberOf(set, pod)
 }
 canAdoptFunc := controller.RecheckDeletionTimestamp(func() (metav1.Object, error) {
  fresh, err := ssc.kubeClient.AppsV1().StatefulSets(set.Namespace).Get(set.Name, metav1.GetOptions{})
  if err != nil {
   return nil, err
  }
  if fresh.UID != set.UID {
   return nil, fmt.Errorf("original StatefulSet %v/%v is gone: got uid %v, wanted %v", set.Namespace, set.Name, fresh.UID, set.UID)
  }
  return fresh, nil
 })
 cm := controller.NewPodControllerRefManager(ssc.podControl, set, selector, controllerKind, canAdoptFunc)
 return cm.ClaimPods(pods, filter)
}
func (ssc *StatefulSetController) adoptOrphanRevisions(set *apps.StatefulSet) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 revisions, err := ssc.control.ListRevisions(set)
 if err != nil {
  return err
 }
 hasOrphans := false
 for i := range revisions {
  if metav1.GetControllerOf(revisions[i]) == nil {
   hasOrphans = true
   break
  }
 }
 if hasOrphans {
  fresh, err := ssc.kubeClient.AppsV1().StatefulSets(set.Namespace).Get(set.Name, metav1.GetOptions{})
  if err != nil {
   return err
  }
  if fresh.UID != set.UID {
   return fmt.Errorf("original StatefulSet %v/%v is gone: got uid %v, wanted %v", set.Namespace, set.Name, fresh.UID, set.UID)
  }
  return ssc.control.AdoptOrphanRevisions(set, revisions)
 }
 return nil
}
func (ssc *StatefulSetController) getStatefulSetsForPod(pod *v1.Pod) []*apps.StatefulSet {
 _logClusterCodePath()
 defer _logClusterCodePath()
 sets, err := ssc.setLister.GetPodStatefulSets(pod)
 if err != nil {
  return nil
 }
 if len(sets) > 1 {
  utilruntime.HandleError(fmt.Errorf("user error: more than one StatefulSet is selecting pods with labels: %+v", pod.Labels))
 }
 return sets
}
func (ssc *StatefulSetController) resolveControllerRef(namespace string, controllerRef *metav1.OwnerReference) *apps.StatefulSet {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if controllerRef.Kind != controllerKind.Kind {
  return nil
 }
 set, err := ssc.setLister.StatefulSets(namespace).Get(controllerRef.Name)
 if err != nil {
  return nil
 }
 if set.UID != controllerRef.UID {
  return nil
 }
 return set
}
func (ssc *StatefulSetController) enqueueStatefulSet(obj interface{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 key, err := controller.KeyFunc(obj)
 if err != nil {
  utilruntime.HandleError(fmt.Errorf("Couldn't get key for object %+v: %v", obj, err))
  return
 }
 ssc.queue.Add(key)
}
func (ssc *StatefulSetController) processNextWorkItem() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 key, quit := ssc.queue.Get()
 if quit {
  return false
 }
 defer ssc.queue.Done(key)
 if err := ssc.sync(key.(string)); err != nil {
  utilruntime.HandleError(fmt.Errorf("Error syncing StatefulSet %v, requeuing: %v", key.(string), err))
  ssc.queue.AddRateLimited(key)
 } else {
  ssc.queue.Forget(key)
 }
 return true
}
func (ssc *StatefulSetController) worker() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for ssc.processNextWorkItem() {
 }
}
func (ssc *StatefulSetController) sync(key string) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 startTime := time.Now()
 defer func() {
  klog.V(4).Infof("Finished syncing statefulset %q (%v)", key, time.Since(startTime))
 }()
 namespace, name, err := cache.SplitMetaNamespaceKey(key)
 if err != nil {
  return err
 }
 set, err := ssc.setLister.StatefulSets(namespace).Get(name)
 if errors.IsNotFound(err) {
  klog.Infof("StatefulSet has been deleted %v", key)
  return nil
 }
 if err != nil {
  utilruntime.HandleError(fmt.Errorf("unable to retrieve StatefulSet %v from store: %v", key, err))
  return err
 }
 selector, err := metav1.LabelSelectorAsSelector(set.Spec.Selector)
 if err != nil {
  utilruntime.HandleError(fmt.Errorf("error converting StatefulSet %v selector: %v", key, err))
  return nil
 }
 if err := ssc.adoptOrphanRevisions(set); err != nil {
  return err
 }
 pods, err := ssc.getPodsForStatefulSet(set, selector)
 if err != nil {
  return err
 }
 return ssc.syncStatefulSet(set, pods)
}
func (ssc *StatefulSetController) syncStatefulSet(set *apps.StatefulSet, pods []*v1.Pod) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(4).Infof("Syncing StatefulSet %v/%v with %d pods", set.Namespace, set.Name, len(pods))
 if err := ssc.control.UpdateStatefulSet(set.DeepCopy(), pods); err != nil {
  return err
 }
 klog.V(4).Infof("Successfully synced StatefulSet %s/%s successful", set.Namespace, set.Name)
 return nil
}
