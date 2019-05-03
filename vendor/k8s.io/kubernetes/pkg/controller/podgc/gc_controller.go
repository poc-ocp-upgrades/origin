package podgc

import (
 "sort"
 "sync"
 "time"
 "k8s.io/api/core/v1"
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 "k8s.io/apimachinery/pkg/labels"
 utilruntime "k8s.io/apimachinery/pkg/util/runtime"
 "k8s.io/apimachinery/pkg/util/sets"
 "k8s.io/apimachinery/pkg/util/wait"
 coreinformers "k8s.io/client-go/informers/core/v1"
 clientset "k8s.io/client-go/kubernetes"
 corelisters "k8s.io/client-go/listers/core/v1"
 "k8s.io/client-go/tools/cache"
 "k8s.io/kubernetes/pkg/controller"
 "k8s.io/kubernetes/pkg/util/metrics"
 "k8s.io/klog"
)

const (
 gcCheckPeriod = 20 * time.Second
)

type PodGCController struct {
 kubeClient             clientset.Interface
 podLister              corelisters.PodLister
 podListerSynced        cache.InformerSynced
 deletePod              func(namespace, name string) error
 terminatedPodThreshold int
}

func NewPodGC(kubeClient clientset.Interface, podInformer coreinformers.PodInformer, terminatedPodThreshold int) *PodGCController {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if kubeClient != nil && kubeClient.CoreV1().RESTClient().GetRateLimiter() != nil {
  metrics.RegisterMetricAndTrackRateLimiterUsage("gc_controller", kubeClient.CoreV1().RESTClient().GetRateLimiter())
 }
 gcc := &PodGCController{kubeClient: kubeClient, terminatedPodThreshold: terminatedPodThreshold, deletePod: func(namespace, name string) error {
  klog.Infof("PodGC is force deleting Pod: %v/%v", namespace, name)
  return kubeClient.CoreV1().Pods(namespace).Delete(name, metav1.NewDeleteOptions(0))
 }}
 gcc.podLister = podInformer.Lister()
 gcc.podListerSynced = podInformer.Informer().HasSynced
 return gcc
}
func (gcc *PodGCController) Run(stop <-chan struct{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 defer utilruntime.HandleCrash()
 klog.Infof("Starting GC controller")
 defer klog.Infof("Shutting down GC controller")
 if !controller.WaitForCacheSync("GC", stop, gcc.podListerSynced) {
  return
 }
 go wait.Until(gcc.gc, gcCheckPeriod, stop)
 <-stop
}
func (gcc *PodGCController) gc() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 pods, err := gcc.podLister.List(labels.Everything())
 if err != nil {
  klog.Errorf("Error while listing all Pods: %v", err)
  return
 }
 if gcc.terminatedPodThreshold > 0 {
  gcc.gcTerminated(pods)
 }
 gcc.gcOrphaned(pods)
 gcc.gcUnscheduledTerminating(pods)
}
func isPodTerminated(pod *v1.Pod) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if phase := pod.Status.Phase; phase != v1.PodPending && phase != v1.PodRunning && phase != v1.PodUnknown {
  return true
 }
 return false
}
func (gcc *PodGCController) gcTerminated(pods []*v1.Pod) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 terminatedPods := []*v1.Pod{}
 for _, pod := range pods {
  if isPodTerminated(pod) {
   terminatedPods = append(terminatedPods, pod)
  }
 }
 terminatedPodCount := len(terminatedPods)
 sort.Sort(byCreationTimestamp(terminatedPods))
 deleteCount := terminatedPodCount - gcc.terminatedPodThreshold
 if deleteCount > terminatedPodCount {
  deleteCount = terminatedPodCount
 }
 if deleteCount > 0 {
  klog.Infof("garbage collecting %v pods", deleteCount)
 }
 var wait sync.WaitGroup
 for i := 0; i < deleteCount; i++ {
  wait.Add(1)
  go func(namespace string, name string) {
   defer wait.Done()
   if err := gcc.deletePod(namespace, name); err != nil {
    defer utilruntime.HandleError(err)
   }
  }(terminatedPods[i].Namespace, terminatedPods[i].Name)
 }
 wait.Wait()
}
func (gcc *PodGCController) gcOrphaned(pods []*v1.Pod) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(4).Infof("GC'ing orphaned")
 nodes, err := gcc.kubeClient.CoreV1().Nodes().List(metav1.ListOptions{})
 if err != nil {
  return
 }
 nodeNames := sets.NewString()
 for i := range nodes.Items {
  nodeNames.Insert(nodes.Items[i].Name)
 }
 for _, pod := range pods {
  if pod.Spec.NodeName == "" {
   continue
  }
  if nodeNames.Has(pod.Spec.NodeName) {
   continue
  }
  klog.V(2).Infof("Found orphaned Pod %v/%v assigned to the Node %v. Deleting.", pod.Namespace, pod.Name, pod.Spec.NodeName)
  if err := gcc.deletePod(pod.Namespace, pod.Name); err != nil {
   utilruntime.HandleError(err)
  } else {
   klog.V(0).Infof("Forced deletion of orphaned Pod %v/%v succeeded", pod.Namespace, pod.Name)
  }
 }
}
func (gcc *PodGCController) gcUnscheduledTerminating(pods []*v1.Pod) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(4).Infof("GC'ing unscheduled pods which are terminating.")
 for _, pod := range pods {
  if pod.DeletionTimestamp == nil || len(pod.Spec.NodeName) > 0 {
   continue
  }
  klog.V(2).Infof("Found unscheduled terminating Pod %v/%v not assigned to any Node. Deleting.", pod.Namespace, pod.Name)
  if err := gcc.deletePod(pod.Namespace, pod.Name); err != nil {
   utilruntime.HandleError(err)
  } else {
   klog.V(0).Infof("Forced deletion of unscheduled terminating Pod %v/%v succeeded", pod.Namespace, pod.Name)
  }
 }
}

type byCreationTimestamp []*v1.Pod

func (o byCreationTimestamp) Len() int {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return len(o)
}
func (o byCreationTimestamp) Swap(i, j int) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 o[i], o[j] = o[j], o[i]
}
func (o byCreationTimestamp) Less(i, j int) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if o[i].CreationTimestamp.Equal(&o[j].CreationTimestamp) {
  return o[i].Name < o[j].Name
 }
 return o[i].CreationTimestamp.Before(&o[j].CreationTimestamp)
}
