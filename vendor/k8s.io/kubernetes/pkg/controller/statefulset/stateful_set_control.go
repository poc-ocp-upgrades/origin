package statefulset

import (
 "math"
 "sort"
 "k8s.io/klog"
 apps "k8s.io/api/apps/v1"
 "k8s.io/api/core/v1"
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 "k8s.io/client-go/tools/record"
 "k8s.io/kubernetes/pkg/controller/history"
)

type StatefulSetControlInterface interface {
 UpdateStatefulSet(set *apps.StatefulSet, pods []*v1.Pod) error
 ListRevisions(set *apps.StatefulSet) ([]*apps.ControllerRevision, error)
 AdoptOrphanRevisions(set *apps.StatefulSet, revisions []*apps.ControllerRevision) error
}

func NewDefaultStatefulSetControl(podControl StatefulPodControlInterface, statusUpdater StatefulSetStatusUpdaterInterface, controllerHistory history.Interface, recorder record.EventRecorder) StatefulSetControlInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &defaultStatefulSetControl{podControl, statusUpdater, controllerHistory, recorder}
}

type defaultStatefulSetControl struct {
 podControl        StatefulPodControlInterface
 statusUpdater     StatefulSetStatusUpdaterInterface
 controllerHistory history.Interface
 recorder          record.EventRecorder
}

func (ssc *defaultStatefulSetControl) UpdateStatefulSet(set *apps.StatefulSet, pods []*v1.Pod) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 revisions, err := ssc.ListRevisions(set)
 if err != nil {
  return err
 }
 history.SortControllerRevisions(revisions)
 currentRevision, updateRevision, collisionCount, err := ssc.getStatefulSetRevisions(set, revisions)
 if err != nil {
  return err
 }
 status, err := ssc.updateStatefulSet(set, currentRevision, updateRevision, collisionCount, pods)
 if err != nil {
  return err
 }
 err = ssc.updateStatefulSetStatus(set, status)
 if err != nil {
  return err
 }
 klog.V(4).Infof("StatefulSet %s/%s pod status replicas=%d ready=%d current=%d updated=%d", set.Namespace, set.Name, status.Replicas, status.ReadyReplicas, status.CurrentReplicas, status.UpdatedReplicas)
 klog.V(4).Infof("StatefulSet %s/%s revisions current=%s update=%s", set.Namespace, set.Name, status.CurrentRevision, status.UpdateRevision)
 return ssc.truncateHistory(set, pods, revisions, currentRevision, updateRevision)
}
func (ssc *defaultStatefulSetControl) ListRevisions(set *apps.StatefulSet) ([]*apps.ControllerRevision, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 selector, err := metav1.LabelSelectorAsSelector(set.Spec.Selector)
 if err != nil {
  return nil, err
 }
 return ssc.controllerHistory.ListControllerRevisions(set, selector)
}
func (ssc *defaultStatefulSetControl) AdoptOrphanRevisions(set *apps.StatefulSet, revisions []*apps.ControllerRevision) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for i := range revisions {
  adopted, err := ssc.controllerHistory.AdoptControllerRevision(set, controllerKind, revisions[i])
  if err != nil {
   return err
  }
  revisions[i] = adopted
 }
 return nil
}
func (ssc *defaultStatefulSetControl) truncateHistory(set *apps.StatefulSet, pods []*v1.Pod, revisions []*apps.ControllerRevision, current *apps.ControllerRevision, update *apps.ControllerRevision) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 history := make([]*apps.ControllerRevision, 0, len(revisions))
 live := map[string]bool{current.Name: true, update.Name: true}
 for i := range pods {
  live[getPodRevision(pods[i])] = true
 }
 for i := range revisions {
  if !live[revisions[i].Name] {
   history = append(history, revisions[i])
  }
 }
 historyLen := len(history)
 historyLimit := int(*set.Spec.RevisionHistoryLimit)
 if historyLen <= historyLimit {
  return nil
 }
 history = history[:(historyLen - historyLimit)]
 for i := 0; i < len(history); i++ {
  if err := ssc.controllerHistory.DeleteControllerRevision(history[i]); err != nil {
   return err
  }
 }
 return nil
}
func (ssc *defaultStatefulSetControl) getStatefulSetRevisions(set *apps.StatefulSet, revisions []*apps.ControllerRevision) (*apps.ControllerRevision, *apps.ControllerRevision, int32, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var currentRevision, updateRevision *apps.ControllerRevision
 revisionCount := len(revisions)
 history.SortControllerRevisions(revisions)
 var collisionCount int32
 if set.Status.CollisionCount != nil {
  collisionCount = *set.Status.CollisionCount
 }
 updateRevision, err := newRevision(set, nextRevision(revisions), &collisionCount)
 if err != nil {
  return nil, nil, collisionCount, err
 }
 equalRevisions := history.FindEqualRevisions(revisions, updateRevision)
 equalCount := len(equalRevisions)
 if equalCount > 0 && history.EqualRevision(revisions[revisionCount-1], equalRevisions[equalCount-1]) {
  updateRevision = revisions[revisionCount-1]
 } else if equalCount > 0 {
  updateRevision, err = ssc.controllerHistory.UpdateControllerRevision(equalRevisions[equalCount-1], updateRevision.Revision)
  if err != nil {
   return nil, nil, collisionCount, err
  }
 } else {
  updateRevision, err = ssc.controllerHistory.CreateControllerRevision(set, updateRevision, &collisionCount)
  if err != nil {
   return nil, nil, collisionCount, err
  }
 }
 for i := range revisions {
  if revisions[i].Name == set.Status.CurrentRevision {
   currentRevision = revisions[i]
  }
 }
 if currentRevision == nil {
  currentRevision = updateRevision
 }
 return currentRevision, updateRevision, collisionCount, nil
}
func (ssc *defaultStatefulSetControl) updateStatefulSet(set *apps.StatefulSet, currentRevision *apps.ControllerRevision, updateRevision *apps.ControllerRevision, collisionCount int32, pods []*v1.Pod) (*apps.StatefulSetStatus, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 currentSet, err := ApplyRevision(set, currentRevision)
 if err != nil {
  return nil, err
 }
 updateSet, err := ApplyRevision(set, updateRevision)
 if err != nil {
  return nil, err
 }
 status := apps.StatefulSetStatus{}
 status.ObservedGeneration = set.Generation
 status.CurrentRevision = currentRevision.Name
 status.UpdateRevision = updateRevision.Name
 status.CollisionCount = new(int32)
 *status.CollisionCount = collisionCount
 replicaCount := int(*set.Spec.Replicas)
 replicas := make([]*v1.Pod, replicaCount)
 condemned := make([]*v1.Pod, 0, len(pods))
 unhealthy := 0
 firstUnhealthyOrdinal := math.MaxInt32
 var firstUnhealthyPod *v1.Pod
 for i := range pods {
  status.Replicas++
  if isRunningAndReady(pods[i]) {
   status.ReadyReplicas++
  }
  if isCreated(pods[i]) && !isTerminating(pods[i]) {
   if getPodRevision(pods[i]) == currentRevision.Name {
    status.CurrentReplicas++
   }
   if getPodRevision(pods[i]) == updateRevision.Name {
    status.UpdatedReplicas++
   }
  }
  if ord := getOrdinal(pods[i]); 0 <= ord && ord < replicaCount {
   replicas[ord] = pods[i]
  } else if ord >= replicaCount {
   condemned = append(condemned, pods[i])
  }
 }
 for ord := 0; ord < replicaCount; ord++ {
  if replicas[ord] == nil {
   replicas[ord] = newVersionedStatefulSetPod(currentSet, updateSet, currentRevision.Name, updateRevision.Name, ord)
  }
 }
 sort.Sort(ascendingOrdinal(condemned))
 for i := range replicas {
  if !isHealthy(replicas[i]) {
   unhealthy++
   if ord := getOrdinal(replicas[i]); ord < firstUnhealthyOrdinal {
    firstUnhealthyOrdinal = ord
    firstUnhealthyPod = replicas[i]
   }
  }
 }
 for i := range condemned {
  if !isHealthy(condemned[i]) {
   unhealthy++
   if ord := getOrdinal(condemned[i]); ord < firstUnhealthyOrdinal {
    firstUnhealthyOrdinal = ord
    firstUnhealthyPod = condemned[i]
   }
  }
 }
 if unhealthy > 0 {
  klog.V(4).Infof("StatefulSet %s/%s has %d unhealthy Pods starting with %s", set.Namespace, set.Name, unhealthy, firstUnhealthyPod.Name)
 }
 if set.DeletionTimestamp != nil {
  return &status, nil
 }
 monotonic := !allowsBurst(set)
 for i := range replicas {
  if isFailed(replicas[i]) {
   ssc.recorder.Eventf(set, v1.EventTypeWarning, "RecreatingFailedPod", "StatefulSet %s/%s is recreating failed Pod %s", set.Namespace, set.Name, replicas[i].Name)
   if err := ssc.podControl.DeleteStatefulPod(set, replicas[i]); err != nil {
    return &status, err
   }
   if getPodRevision(replicas[i]) == currentRevision.Name {
    status.CurrentReplicas--
   }
   if getPodRevision(replicas[i]) == updateRevision.Name {
    status.UpdatedReplicas--
   }
   status.Replicas--
   replicas[i] = newVersionedStatefulSetPod(currentSet, updateSet, currentRevision.Name, updateRevision.Name, i)
  }
  if !isCreated(replicas[i]) {
   if err := ssc.podControl.CreateStatefulPod(set, replicas[i]); err != nil {
    return &status, err
   }
   status.Replicas++
   if getPodRevision(replicas[i]) == currentRevision.Name {
    status.CurrentReplicas++
   }
   if getPodRevision(replicas[i]) == updateRevision.Name {
    status.UpdatedReplicas++
   }
   if monotonic {
    return &status, nil
   }
   continue
  }
  if isTerminating(replicas[i]) && monotonic {
   klog.V(4).Infof("StatefulSet %s/%s is waiting for Pod %s to Terminate", set.Namespace, set.Name, replicas[i].Name)
   return &status, nil
  }
  if !isRunningAndReady(replicas[i]) && monotonic {
   klog.V(4).Infof("StatefulSet %s/%s is waiting for Pod %s to be Running and Ready", set.Namespace, set.Name, replicas[i].Name)
   return &status, nil
  }
  if identityMatches(set, replicas[i]) && storageMatches(set, replicas[i]) {
   continue
  }
  replica := replicas[i].DeepCopy()
  if err := ssc.podControl.UpdateStatefulPod(updateSet, replica); err != nil {
   return &status, err
  }
 }
 for target := len(condemned) - 1; target >= 0; target-- {
  if isTerminating(condemned[target]) {
   klog.V(4).Infof("StatefulSet %s/%s is waiting for Pod %s to Terminate prior to scale down", set.Namespace, set.Name, condemned[target].Name)
   if monotonic {
    return &status, nil
   }
   continue
  }
  if !isRunningAndReady(condemned[target]) && monotonic && condemned[target] != firstUnhealthyPod {
   klog.V(4).Infof("StatefulSet %s/%s is waiting for Pod %s to be Running and Ready prior to scale down", set.Namespace, set.Name, firstUnhealthyPod.Name)
   return &status, nil
  }
  klog.V(2).Infof("StatefulSet %s/%s terminating Pod %s for scale down", set.Namespace, set.Name, condemned[target].Name)
  if err := ssc.podControl.DeleteStatefulPod(set, condemned[target]); err != nil {
   return &status, err
  }
  if getPodRevision(condemned[target]) == currentRevision.Name {
   status.CurrentReplicas--
  }
  if getPodRevision(condemned[target]) == updateRevision.Name {
   status.UpdatedReplicas--
  }
  if monotonic {
   return &status, nil
  }
 }
 if set.Spec.UpdateStrategy.Type == apps.OnDeleteStatefulSetStrategyType {
  return &status, nil
 }
 updateMin := 0
 if set.Spec.UpdateStrategy.RollingUpdate != nil {
  updateMin = int(*set.Spec.UpdateStrategy.RollingUpdate.Partition)
 }
 for target := len(replicas) - 1; target >= updateMin; target-- {
  if getPodRevision(replicas[target]) != updateRevision.Name && !isTerminating(replicas[target]) {
   klog.V(2).Infof("StatefulSet %s/%s terminating Pod %s for update", set.Namespace, set.Name, replicas[target].Name)
   err := ssc.podControl.DeleteStatefulPod(set, replicas[target])
   status.CurrentReplicas--
   return &status, err
  }
  if !isHealthy(replicas[target]) {
   klog.V(4).Infof("StatefulSet %s/%s is waiting for Pod %s to update", set.Namespace, set.Name, replicas[target].Name)
   return &status, nil
  }
 }
 return &status, nil
}
func (ssc *defaultStatefulSetControl) updateStatefulSetStatus(set *apps.StatefulSet, status *apps.StatefulSetStatus) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 completeRollingUpdate(set, status)
 if !inconsistentStatus(set, status) {
  return nil
 }
 set = set.DeepCopy()
 if err := ssc.statusUpdater.UpdateStatefulSetStatus(set, status); err != nil {
  return err
 }
 return nil
}

var _ StatefulSetControlInterface = &defaultStatefulSetControl{}
