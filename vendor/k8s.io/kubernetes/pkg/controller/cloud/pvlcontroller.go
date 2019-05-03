package cloud

import (
 "context"
 "encoding/json"
 "fmt"
 "time"
 "k8s.io/klog"
 "k8s.io/api/core/v1"
 "k8s.io/apimachinery/pkg/api/errors"
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 "k8s.io/apimachinery/pkg/runtime"
 "k8s.io/apimachinery/pkg/types"
 utilruntime "k8s.io/apimachinery/pkg/util/runtime"
 "k8s.io/apimachinery/pkg/util/strategicpatch"
 "k8s.io/apimachinery/pkg/util/wait"
 "k8s.io/apimachinery/pkg/watch"
 utilfeature "k8s.io/apiserver/pkg/util/feature"
 v1helper "k8s.io/kubernetes/pkg/apis/core/v1/helper"
 "k8s.io/kubernetes/pkg/features"
 kubeletapis "k8s.io/kubernetes/pkg/kubelet/apis"
 volumeutil "k8s.io/kubernetes/pkg/volume/util"
 "k8s.io/client-go/kubernetes"
 corelisters "k8s.io/client-go/listers/core/v1"
 "k8s.io/client-go/tools/cache"
 "k8s.io/client-go/util/workqueue"
 cloudprovider "k8s.io/cloud-provider"
 "k8s.io/kubernetes/pkg/controller"
)

const initializerName = "pvlabel.kubernetes.io"

type PersistentVolumeLabelController struct {
 cloud         cloudprovider.Interface
 kubeClient    kubernetes.Interface
 pvlController cache.Controller
 pvlIndexer    cache.Indexer
 volumeLister  corelisters.PersistentVolumeLister
 syncHandler   func(key string) error
 queue         workqueue.RateLimitingInterface
}

func NewPersistentVolumeLabelController(kubeClient kubernetes.Interface, cloud cloudprovider.Interface) *PersistentVolumeLabelController {
 _logClusterCodePath()
 defer _logClusterCodePath()
 pvlc := &PersistentVolumeLabelController{cloud: cloud, kubeClient: kubeClient, queue: workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "pvLabels")}
 pvlc.syncHandler = pvlc.addLabelsAndAffinity
 pvlc.pvlIndexer, pvlc.pvlController = cache.NewIndexerInformer(&cache.ListWatch{ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
  options.IncludeUninitialized = true
  return kubeClient.CoreV1().PersistentVolumes().List(options)
 }, WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
  options.IncludeUninitialized = true
  return kubeClient.CoreV1().PersistentVolumes().Watch(options)
 }}, &v1.PersistentVolume{}, 0, cache.ResourceEventHandlerFuncs{AddFunc: func(obj interface{}) {
  key, err := cache.MetaNamespaceKeyFunc(obj)
  if err == nil {
   pvlc.queue.Add(key)
  }
 }}, cache.Indexers{})
 pvlc.volumeLister = corelisters.NewPersistentVolumeLister(pvlc.pvlIndexer)
 return pvlc
}
func (pvlc *PersistentVolumeLabelController) Run(threadiness int, stopCh <-chan struct{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 defer utilruntime.HandleCrash()
 defer pvlc.queue.ShutDown()
 klog.Infof("Starting PersistentVolumeLabelController")
 defer klog.Infof("Shutting down PersistentVolumeLabelController")
 go pvlc.pvlController.Run(stopCh)
 if !controller.WaitForCacheSync("persistent volume label", stopCh, pvlc.pvlController.HasSynced) {
  return
 }
 for i := 0; i < threadiness; i++ {
  go wait.Until(pvlc.runWorker, time.Second, stopCh)
 }
 <-stopCh
}
func (pvlc *PersistentVolumeLabelController) runWorker() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for pvlc.processNextWorkItem() {
 }
}
func (pvlc *PersistentVolumeLabelController) processNextWorkItem() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 keyObj, quit := pvlc.queue.Get()
 if quit {
  return false
 }
 defer pvlc.queue.Done(keyObj)
 key := keyObj.(string)
 err := pvlc.syncHandler(key)
 if err == nil {
  pvlc.queue.Forget(key)
  return true
 }
 utilruntime.HandleError(fmt.Errorf("%v failed with : %v", key, err))
 pvlc.queue.AddRateLimited(key)
 return true
}
func (pvlc *PersistentVolumeLabelController) addLabelsAndAffinity(key string) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 _, name, err := cache.SplitMetaNamespaceKey(key)
 if err != nil {
  return fmt.Errorf("error getting name of volume %q to get volume from informer: %v", key, err)
 }
 volume, err := pvlc.volumeLister.Get(name)
 if errors.IsNotFound(err) {
  return nil
 } else if err != nil {
  return fmt.Errorf("error getting volume %s from informer: %v", name, err)
 }
 return pvlc.addLabelsAndAffinityToVolume(volume)
}
func (pvlc *PersistentVolumeLabelController) addLabelsAndAffinityToVolume(vol *v1.PersistentVolume) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var volumeLabels map[string]string
 if needsInitialization(vol.Initializers, initializerName) {
  if labeler, ok := (pvlc.cloud).(cloudprovider.PVLabeler); ok {
   labels, err := labeler.GetLabelsForVolume(context.TODO(), vol)
   if err != nil {
    return fmt.Errorf("error querying volume %v: %v", vol.Spec, err)
   }
   volumeLabels = labels
  } else {
   klog.V(4).Info("cloud provider does not support PVLabeler")
  }
  return pvlc.updateVolume(vol, volumeLabels)
 }
 return nil
}
func (pvlc *PersistentVolumeLabelController) createPatch(vol *v1.PersistentVolume, volLabels map[string]string) ([]byte, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 volName := vol.Name
 newVolume := vol.DeepCopyObject().(*v1.PersistentVolume)
 populateAffinity := utilfeature.DefaultFeatureGate.Enabled(features.VolumeScheduling) && len(volLabels) != 0
 if newVolume.Labels == nil {
  newVolume.Labels = make(map[string]string)
 }
 requirements := make([]v1.NodeSelectorRequirement, 0)
 for k, v := range volLabels {
  newVolume.Labels[k] = v
  if populateAffinity {
   var values []string
   if k == kubeletapis.LabelZoneFailureDomain {
    zones, err := volumeutil.LabelZonesToSet(v)
    if err != nil {
     return nil, fmt.Errorf("failed to convert label string for Zone: %s to a Set", v)
    }
    values = zones.List()
   } else {
    values = []string{v}
   }
   requirements = append(requirements, v1.NodeSelectorRequirement{Key: k, Operator: v1.NodeSelectorOpIn, Values: values})
  }
 }
 if populateAffinity {
  if newVolume.Spec.NodeAffinity == nil {
   newVolume.Spec.NodeAffinity = new(v1.VolumeNodeAffinity)
  }
  if newVolume.Spec.NodeAffinity.Required == nil {
   newVolume.Spec.NodeAffinity.Required = new(v1.NodeSelector)
  }
  if len(newVolume.Spec.NodeAffinity.Required.NodeSelectorTerms) == 0 {
   newVolume.Spec.NodeAffinity.Required.NodeSelectorTerms = make([]v1.NodeSelectorTerm, 1)
  }
  if v1helper.NodeSelectorRequirementKeysExistInNodeSelectorTerms(requirements, newVolume.Spec.NodeAffinity.Required.NodeSelectorTerms) {
   klog.V(4).Infof("NodeSelectorRequirements for cloud labels %v conflict with existing NodeAffinity %v. Skipping addition of NodeSelectorRequirements for cloud labels.", requirements, newVolume.Spec.NodeAffinity)
  } else {
   for _, req := range requirements {
    for i := range newVolume.Spec.NodeAffinity.Required.NodeSelectorTerms {
     newVolume.Spec.NodeAffinity.Required.NodeSelectorTerms[i].MatchExpressions = append(newVolume.Spec.NodeAffinity.Required.NodeSelectorTerms[i].MatchExpressions, req)
    }
   }
  }
 }
 newVolume.Initializers = removeInitializer(newVolume.Initializers, initializerName)
 klog.V(4).Infof("removed initializer on PersistentVolume %s", newVolume.Name)
 oldData, err := json.Marshal(vol)
 if err != nil {
  return nil, fmt.Errorf("failed to marshal old persistentvolume %#v for persistentvolume %q: %v", vol, volName, err)
 }
 newData, err := json.Marshal(newVolume)
 if err != nil {
  return nil, fmt.Errorf("failed to marshal new persistentvolume %#v for persistentvolume %q: %v", newVolume, volName, err)
 }
 patch, err := strategicpatch.CreateTwoWayMergePatch(oldData, newData, v1.PersistentVolume{})
 if err != nil {
  return nil, fmt.Errorf("failed to create patch for persistentvolume %q: %v", volName, err)
 }
 return patch, nil
}
func (pvlc *PersistentVolumeLabelController) updateVolume(vol *v1.PersistentVolume, volLabels map[string]string) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 volName := vol.Name
 klog.V(4).Infof("updating PersistentVolume %s", volName)
 patchBytes, err := pvlc.createPatch(vol, volLabels)
 if err != nil {
  return err
 }
 _, err = pvlc.kubeClient.CoreV1().PersistentVolumes().Patch(string(volName), types.StrategicMergePatchType, patchBytes)
 if err != nil {
  return fmt.Errorf("failed to update PersistentVolume %s: %v", volName, err)
 }
 klog.V(4).Infof("updated PersistentVolume %s", volName)
 return nil
}
func removeInitializer(initializers *metav1.Initializers, name string) *metav1.Initializers {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if initializers == nil {
  return nil
 }
 var updated []metav1.Initializer
 for _, pending := range initializers.Pending {
  if pending.Name != name {
   updated = append(updated, pending)
  }
 }
 if len(updated) == len(initializers.Pending) {
  return initializers
 }
 if len(updated) == 0 {
  return nil
 }
 return &metav1.Initializers{Pending: updated}
}
func needsInitialization(initializers *metav1.Initializers, name string) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if initializers == nil {
  return false
 }
 if len(initializers.Pending) == 0 {
  return false
 }
 return initializers.Pending[0].Name == name
}
