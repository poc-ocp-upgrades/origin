package deployment

import (
 "fmt"
 "reflect"
 "sort"
 "strconv"
 apps "k8s.io/api/apps/v1"
 "k8s.io/api/core/v1"
 "k8s.io/apimachinery/pkg/api/errors"
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 "k8s.io/klog"
 "k8s.io/kubernetes/pkg/controller"
 deploymentutil "k8s.io/kubernetes/pkg/controller/deployment/util"
 labelsutil "k8s.io/kubernetes/pkg/util/labels"
)

func (dc *DeploymentController) syncStatusOnly(d *apps.Deployment, rsList []*apps.ReplicaSet) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 newRS, oldRSs, err := dc.getAllReplicaSetsAndSyncRevision(d, rsList, false)
 if err != nil {
  return err
 }
 allRSs := append(oldRSs, newRS)
 return dc.syncDeploymentStatus(allRSs, newRS, d)
}
func (dc *DeploymentController) sync(d *apps.Deployment, rsList []*apps.ReplicaSet) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 newRS, oldRSs, err := dc.getAllReplicaSetsAndSyncRevision(d, rsList, false)
 if err != nil {
  return err
 }
 if err := dc.scale(d, newRS, oldRSs); err != nil {
  return err
 }
 if d.Spec.Paused && getRollbackTo(d) == nil {
  if err := dc.cleanupDeployment(oldRSs, d); err != nil {
   return err
  }
 }
 allRSs := append(oldRSs, newRS)
 return dc.syncDeploymentStatus(allRSs, newRS, d)
}
func (dc *DeploymentController) checkPausedConditions(d *apps.Deployment) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if !deploymentutil.HasProgressDeadline(d) {
  return nil
 }
 cond := deploymentutil.GetDeploymentCondition(d.Status, apps.DeploymentProgressing)
 if cond != nil && cond.Reason == deploymentutil.TimedOutReason {
  return nil
 }
 pausedCondExists := cond != nil && cond.Reason == deploymentutil.PausedDeployReason
 needsUpdate := false
 if d.Spec.Paused && !pausedCondExists {
  condition := deploymentutil.NewDeploymentCondition(apps.DeploymentProgressing, v1.ConditionUnknown, deploymentutil.PausedDeployReason, "Deployment is paused")
  deploymentutil.SetDeploymentCondition(&d.Status, *condition)
  needsUpdate = true
 } else if !d.Spec.Paused && pausedCondExists {
  condition := deploymentutil.NewDeploymentCondition(apps.DeploymentProgressing, v1.ConditionUnknown, deploymentutil.ResumedDeployReason, "Deployment is resumed")
  deploymentutil.SetDeploymentCondition(&d.Status, *condition)
  needsUpdate = true
 }
 if !needsUpdate {
  return nil
 }
 var err error
 d, err = dc.client.AppsV1().Deployments(d.Namespace).UpdateStatus(d)
 return err
}
func (dc *DeploymentController) getAllReplicaSetsAndSyncRevision(d *apps.Deployment, rsList []*apps.ReplicaSet, createIfNotExisted bool) (*apps.ReplicaSet, []*apps.ReplicaSet, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 _, allOldRSs := deploymentutil.FindOldReplicaSets(d, rsList)
 newRS, err := dc.getNewReplicaSet(d, rsList, allOldRSs, createIfNotExisted)
 if err != nil {
  return nil, nil, err
 }
 return newRS, allOldRSs, nil
}
func (dc *DeploymentController) getNewReplicaSet(d *apps.Deployment, rsList, oldRSs []*apps.ReplicaSet, createIfNotExisted bool) (*apps.ReplicaSet, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 existingNewRS := deploymentutil.FindNewReplicaSet(d, rsList)
 maxOldRevision := deploymentutil.MaxRevision(oldRSs)
 newRevision := strconv.FormatInt(maxOldRevision+1, 10)
 if existingNewRS != nil {
  rsCopy := existingNewRS.DeepCopy()
  annotationsUpdated := deploymentutil.SetNewReplicaSetAnnotations(d, rsCopy, newRevision, true)
  minReadySecondsNeedsUpdate := rsCopy.Spec.MinReadySeconds != d.Spec.MinReadySeconds
  if annotationsUpdated || minReadySecondsNeedsUpdate {
   rsCopy.Spec.MinReadySeconds = d.Spec.MinReadySeconds
   return dc.client.AppsV1().ReplicaSets(rsCopy.ObjectMeta.Namespace).Update(rsCopy)
  }
  needsUpdate := deploymentutil.SetDeploymentRevision(d, rsCopy.Annotations[deploymentutil.RevisionAnnotation])
  cond := deploymentutil.GetDeploymentCondition(d.Status, apps.DeploymentProgressing)
  if deploymentutil.HasProgressDeadline(d) && cond == nil {
   msg := fmt.Sprintf("Found new replica set %q", rsCopy.Name)
   condition := deploymentutil.NewDeploymentCondition(apps.DeploymentProgressing, v1.ConditionTrue, deploymentutil.FoundNewRSReason, msg)
   deploymentutil.SetDeploymentCondition(&d.Status, *condition)
   needsUpdate = true
  }
  if needsUpdate {
   var err error
   if d, err = dc.client.AppsV1().Deployments(d.Namespace).UpdateStatus(d); err != nil {
    return nil, err
   }
  }
  return rsCopy, nil
 }
 if !createIfNotExisted {
  return nil, nil
 }
 newRSTemplate := *d.Spec.Template.DeepCopy()
 podTemplateSpecHash := controller.ComputeHash(&newRSTemplate, d.Status.CollisionCount)
 newRSTemplate.Labels = labelsutil.CloneAndAddLabel(d.Spec.Template.Labels, apps.DefaultDeploymentUniqueLabelKey, podTemplateSpecHash)
 newRSSelector := labelsutil.CloneSelectorAndAddLabel(d.Spec.Selector, apps.DefaultDeploymentUniqueLabelKey, podTemplateSpecHash)
 newRS := apps.ReplicaSet{ObjectMeta: metav1.ObjectMeta{Name: d.Name + "-" + podTemplateSpecHash, Namespace: d.Namespace, OwnerReferences: []metav1.OwnerReference{*metav1.NewControllerRef(d, controllerKind)}, Labels: newRSTemplate.Labels}, Spec: apps.ReplicaSetSpec{Replicas: new(int32), MinReadySeconds: d.Spec.MinReadySeconds, Selector: newRSSelector, Template: newRSTemplate}}
 allRSs := append(oldRSs, &newRS)
 newReplicasCount, err := deploymentutil.NewRSNewReplicas(d, allRSs, &newRS)
 if err != nil {
  return nil, err
 }
 *(newRS.Spec.Replicas) = newReplicasCount
 deploymentutil.SetNewReplicaSetAnnotations(d, &newRS, newRevision, false)
 alreadyExists := false
 createdRS, err := dc.client.AppsV1().ReplicaSets(d.Namespace).Create(&newRS)
 switch {
 case errors.IsAlreadyExists(err):
  alreadyExists = true
  rs, rsErr := dc.rsLister.ReplicaSets(newRS.Namespace).Get(newRS.Name)
  if rsErr != nil {
   return nil, rsErr
  }
  controllerRef := metav1.GetControllerOf(rs)
  if controllerRef != nil && controllerRef.UID == d.UID && deploymentutil.EqualIgnoreHash(&d.Spec.Template, &rs.Spec.Template) {
   createdRS = rs
   err = nil
   break
  }
  if d.Status.CollisionCount == nil {
   d.Status.CollisionCount = new(int32)
  }
  preCollisionCount := *d.Status.CollisionCount
  *d.Status.CollisionCount++
  _, dErr := dc.client.AppsV1().Deployments(d.Namespace).UpdateStatus(d)
  if dErr == nil {
   klog.V(2).Infof("Found a hash collision for deployment %q - bumping collisionCount (%d->%d) to resolve it", d.Name, preCollisionCount, *d.Status.CollisionCount)
  }
  return nil, err
 case err != nil:
  msg := fmt.Sprintf("Failed to create new replica set %q: %v", newRS.Name, err)
  if deploymentutil.HasProgressDeadline(d) {
   cond := deploymentutil.NewDeploymentCondition(apps.DeploymentProgressing, v1.ConditionFalse, deploymentutil.FailedRSCreateReason, msg)
   deploymentutil.SetDeploymentCondition(&d.Status, *cond)
   _, _ = dc.client.AppsV1().Deployments(d.Namespace).UpdateStatus(d)
  }
  dc.eventRecorder.Eventf(d, v1.EventTypeWarning, deploymentutil.FailedRSCreateReason, msg)
  return nil, err
 }
 if !alreadyExists && newReplicasCount > 0 {
  dc.eventRecorder.Eventf(d, v1.EventTypeNormal, "ScalingReplicaSet", "Scaled up replica set %s to %d", createdRS.Name, newReplicasCount)
 }
 needsUpdate := deploymentutil.SetDeploymentRevision(d, newRevision)
 if !alreadyExists && deploymentutil.HasProgressDeadline(d) {
  msg := fmt.Sprintf("Created new replica set %q", createdRS.Name)
  condition := deploymentutil.NewDeploymentCondition(apps.DeploymentProgressing, v1.ConditionTrue, deploymentutil.NewReplicaSetReason, msg)
  deploymentutil.SetDeploymentCondition(&d.Status, *condition)
  needsUpdate = true
 }
 if needsUpdate {
  _, err = dc.client.AppsV1().Deployments(d.Namespace).UpdateStatus(d)
 }
 return createdRS, err
}
func (dc *DeploymentController) scale(deployment *apps.Deployment, newRS *apps.ReplicaSet, oldRSs []*apps.ReplicaSet) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if activeOrLatest := deploymentutil.FindActiveOrLatest(newRS, oldRSs); activeOrLatest != nil {
  if *(activeOrLatest.Spec.Replicas) == *(deployment.Spec.Replicas) {
   return nil
  }
  _, _, err := dc.scaleReplicaSetAndRecordEvent(activeOrLatest, *(deployment.Spec.Replicas), deployment)
  return err
 }
 if deploymentutil.IsSaturated(deployment, newRS) {
  for _, old := range controller.FilterActiveReplicaSets(oldRSs) {
   if _, _, err := dc.scaleReplicaSetAndRecordEvent(old, 0, deployment); err != nil {
    return err
   }
  }
  return nil
 }
 if deploymentutil.IsRollingUpdate(deployment) {
  allRSs := controller.FilterActiveReplicaSets(append(oldRSs, newRS))
  allRSsReplicas := deploymentutil.GetReplicaCountForReplicaSets(allRSs)
  allowedSize := int32(0)
  if *(deployment.Spec.Replicas) > 0 {
   allowedSize = *(deployment.Spec.Replicas) + deploymentutil.MaxSurge(*deployment)
  }
  deploymentReplicasToAdd := allowedSize - allRSsReplicas
  var scalingOperation string
  switch {
  case deploymentReplicasToAdd > 0:
   sort.Sort(controller.ReplicaSetsBySizeNewer(allRSs))
   scalingOperation = "up"
  case deploymentReplicasToAdd < 0:
   sort.Sort(controller.ReplicaSetsBySizeOlder(allRSs))
   scalingOperation = "down"
  }
  deploymentReplicasAdded := int32(0)
  nameToSize := make(map[string]int32)
  for i := range allRSs {
   rs := allRSs[i]
   if deploymentReplicasToAdd != 0 {
    proportion := deploymentutil.GetProportion(rs, *deployment, deploymentReplicasToAdd, deploymentReplicasAdded)
    nameToSize[rs.Name] = *(rs.Spec.Replicas) + proportion
    deploymentReplicasAdded += proportion
   } else {
    nameToSize[rs.Name] = *(rs.Spec.Replicas)
   }
  }
  for i := range allRSs {
   rs := allRSs[i]
   if i == 0 && deploymentReplicasToAdd != 0 {
    leftover := deploymentReplicasToAdd - deploymentReplicasAdded
    nameToSize[rs.Name] = nameToSize[rs.Name] + leftover
    if nameToSize[rs.Name] < 0 {
     nameToSize[rs.Name] = 0
    }
   }
   if _, _, err := dc.scaleReplicaSet(rs, nameToSize[rs.Name], deployment, scalingOperation); err != nil {
    return err
   }
  }
 }
 return nil
}
func (dc *DeploymentController) scaleReplicaSetAndRecordEvent(rs *apps.ReplicaSet, newScale int32, deployment *apps.Deployment) (bool, *apps.ReplicaSet, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if *(rs.Spec.Replicas) == newScale {
  return false, rs, nil
 }
 var scalingOperation string
 if *(rs.Spec.Replicas) < newScale {
  scalingOperation = "up"
 } else {
  scalingOperation = "down"
 }
 scaled, newRS, err := dc.scaleReplicaSet(rs, newScale, deployment, scalingOperation)
 return scaled, newRS, err
}
func (dc *DeploymentController) scaleReplicaSet(rs *apps.ReplicaSet, newScale int32, deployment *apps.Deployment, scalingOperation string) (bool, *apps.ReplicaSet, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 sizeNeedsUpdate := *(rs.Spec.Replicas) != newScale
 annotationsNeedUpdate := deploymentutil.ReplicasAnnotationsNeedUpdate(rs, *(deployment.Spec.Replicas), *(deployment.Spec.Replicas)+deploymentutil.MaxSurge(*deployment))
 scaled := false
 var err error
 if sizeNeedsUpdate || annotationsNeedUpdate {
  rsCopy := rs.DeepCopy()
  *(rsCopy.Spec.Replicas) = newScale
  deploymentutil.SetReplicasAnnotations(rsCopy, *(deployment.Spec.Replicas), *(deployment.Spec.Replicas)+deploymentutil.MaxSurge(*deployment))
  rs, err = dc.client.AppsV1().ReplicaSets(rsCopy.Namespace).Update(rsCopy)
  if err == nil && sizeNeedsUpdate {
   scaled = true
   dc.eventRecorder.Eventf(deployment, v1.EventTypeNormal, "ScalingReplicaSet", "Scaled %s replica set %s to %d", scalingOperation, rs.Name, newScale)
  }
 }
 return scaled, rs, err
}
func (dc *DeploymentController) cleanupDeployment(oldRSs []*apps.ReplicaSet, deployment *apps.Deployment) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if !deploymentutil.HasRevisionHistoryLimit(deployment) {
  return nil
 }
 aliveFilter := func(rs *apps.ReplicaSet) bool {
  return rs != nil && rs.ObjectMeta.DeletionTimestamp == nil
 }
 cleanableRSes := controller.FilterReplicaSets(oldRSs, aliveFilter)
 diff := int32(len(cleanableRSes)) - *deployment.Spec.RevisionHistoryLimit
 if diff <= 0 {
  return nil
 }
 sort.Sort(controller.ReplicaSetsByCreationTimestamp(cleanableRSes))
 klog.V(4).Infof("Looking to cleanup old replica sets for deployment %q", deployment.Name)
 for i := int32(0); i < diff; i++ {
  rs := cleanableRSes[i]
  if rs.Status.Replicas != 0 || *(rs.Spec.Replicas) != 0 || rs.Generation > rs.Status.ObservedGeneration || rs.DeletionTimestamp != nil {
   continue
  }
  klog.V(4).Infof("Trying to cleanup replica set %q for deployment %q", rs.Name, deployment.Name)
  if err := dc.client.AppsV1().ReplicaSets(rs.Namespace).Delete(rs.Name, nil); err != nil && !errors.IsNotFound(err) {
   return err
  }
 }
 return nil
}
func (dc *DeploymentController) syncDeploymentStatus(allRSs []*apps.ReplicaSet, newRS *apps.ReplicaSet, d *apps.Deployment) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 newStatus := calculateStatus(allRSs, newRS, d)
 if reflect.DeepEqual(d.Status, newStatus) {
  return nil
 }
 newDeployment := d
 newDeployment.Status = newStatus
 _, err := dc.client.AppsV1().Deployments(newDeployment.Namespace).UpdateStatus(newDeployment)
 return err
}
func calculateStatus(allRSs []*apps.ReplicaSet, newRS *apps.ReplicaSet, deployment *apps.Deployment) apps.DeploymentStatus {
 _logClusterCodePath()
 defer _logClusterCodePath()
 availableReplicas := deploymentutil.GetAvailableReplicaCountForReplicaSets(allRSs)
 totalReplicas := deploymentutil.GetReplicaCountForReplicaSets(allRSs)
 unavailableReplicas := totalReplicas - availableReplicas
 if unavailableReplicas < 0 {
  unavailableReplicas = 0
 }
 status := apps.DeploymentStatus{ObservedGeneration: deployment.Generation, Replicas: deploymentutil.GetActualReplicaCountForReplicaSets(allRSs), UpdatedReplicas: deploymentutil.GetActualReplicaCountForReplicaSets([]*apps.ReplicaSet{newRS}), ReadyReplicas: deploymentutil.GetReadyReplicaCountForReplicaSets(allRSs), AvailableReplicas: availableReplicas, UnavailableReplicas: unavailableReplicas, CollisionCount: deployment.Status.CollisionCount}
 conditions := deployment.Status.Conditions
 for i := range conditions {
  status.Conditions = append(status.Conditions, conditions[i])
 }
 if availableReplicas >= *(deployment.Spec.Replicas)-deploymentutil.MaxUnavailable(*deployment) {
  minAvailability := deploymentutil.NewDeploymentCondition(apps.DeploymentAvailable, v1.ConditionTrue, deploymentutil.MinimumReplicasAvailable, "Deployment has minimum availability.")
  deploymentutil.SetDeploymentCondition(&status, *minAvailability)
 } else {
  noMinAvailability := deploymentutil.NewDeploymentCondition(apps.DeploymentAvailable, v1.ConditionFalse, deploymentutil.MinimumReplicasUnavailable, "Deployment does not have minimum availability.")
  deploymentutil.SetDeploymentCondition(&status, *noMinAvailability)
 }
 return status
}
func (dc *DeploymentController) isScalingEvent(d *apps.Deployment, rsList []*apps.ReplicaSet) (bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 newRS, oldRSs, err := dc.getAllReplicaSetsAndSyncRevision(d, rsList, false)
 if err != nil {
  return false, err
 }
 allRSs := append(oldRSs, newRS)
 for _, rs := range controller.FilterActiveReplicaSets(allRSs) {
  desired, ok := deploymentutil.GetDesiredReplicasAnnotation(rs)
  if !ok {
   continue
  }
  if desired != *(d.Spec.Replicas) {
   return true, nil
  }
 }
 return false, nil
}
