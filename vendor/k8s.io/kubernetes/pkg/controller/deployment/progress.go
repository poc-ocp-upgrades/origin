package deployment

import (
 "fmt"
 "reflect"
 "time"
 "k8s.io/klog"
 apps "k8s.io/api/apps/v1"
 "k8s.io/api/core/v1"
 "k8s.io/kubernetes/pkg/controller/deployment/util"
)

func (dc *DeploymentController) syncRolloutStatus(allRSs []*apps.ReplicaSet, newRS *apps.ReplicaSet, d *apps.Deployment) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 newStatus := calculateStatus(allRSs, newRS, d)
 if !util.HasProgressDeadline(d) {
  util.RemoveDeploymentCondition(&newStatus, apps.DeploymentProgressing)
 }
 currentCond := util.GetDeploymentCondition(d.Status, apps.DeploymentProgressing)
 isCompleteDeployment := newStatus.Replicas == newStatus.UpdatedReplicas && currentCond != nil && currentCond.Reason == util.NewRSAvailableReason
 if util.HasProgressDeadline(d) && !isCompleteDeployment {
  switch {
  case util.DeploymentComplete(d, &newStatus):
   msg := fmt.Sprintf("Deployment %q has successfully progressed.", d.Name)
   if newRS != nil {
    msg = fmt.Sprintf("ReplicaSet %q has successfully progressed.", newRS.Name)
   }
   condition := util.NewDeploymentCondition(apps.DeploymentProgressing, v1.ConditionTrue, util.NewRSAvailableReason, msg)
   util.SetDeploymentCondition(&newStatus, *condition)
  case util.DeploymentProgressing(d, &newStatus):
   msg := fmt.Sprintf("Deployment %q is progressing.", d.Name)
   if newRS != nil {
    msg = fmt.Sprintf("ReplicaSet %q is progressing.", newRS.Name)
   }
   condition := util.NewDeploymentCondition(apps.DeploymentProgressing, v1.ConditionTrue, util.ReplicaSetUpdatedReason, msg)
   if currentCond != nil {
    if currentCond.Status == v1.ConditionTrue {
     condition.LastTransitionTime = currentCond.LastTransitionTime
    }
    util.RemoveDeploymentCondition(&newStatus, apps.DeploymentProgressing)
   }
   util.SetDeploymentCondition(&newStatus, *condition)
  case util.DeploymentTimedOut(d, &newStatus):
   msg := fmt.Sprintf("Deployment %q has timed out progressing.", d.Name)
   if newRS != nil {
    msg = fmt.Sprintf("ReplicaSet %q has timed out progressing.", newRS.Name)
   }
   condition := util.NewDeploymentCondition(apps.DeploymentProgressing, v1.ConditionFalse, util.TimedOutReason, msg)
   util.SetDeploymentCondition(&newStatus, *condition)
  }
 }
 if replicaFailureCond := dc.getReplicaFailures(allRSs, newRS); len(replicaFailureCond) > 0 {
  util.SetDeploymentCondition(&newStatus, replicaFailureCond[0])
 } else {
  util.RemoveDeploymentCondition(&newStatus, apps.DeploymentReplicaFailure)
 }
 if reflect.DeepEqual(d.Status, newStatus) {
  dc.requeueStuckDeployment(d, newStatus)
  return nil
 }
 newDeployment := d
 newDeployment.Status = newStatus
 _, err := dc.client.AppsV1().Deployments(newDeployment.Namespace).UpdateStatus(newDeployment)
 return err
}
func (dc *DeploymentController) getReplicaFailures(allRSs []*apps.ReplicaSet, newRS *apps.ReplicaSet) []apps.DeploymentCondition {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var conditions []apps.DeploymentCondition
 if newRS != nil {
  for _, c := range newRS.Status.Conditions {
   if c.Type != apps.ReplicaSetReplicaFailure {
    continue
   }
   conditions = append(conditions, util.ReplicaSetToDeploymentCondition(c))
  }
 }
 if len(conditions) > 0 {
  return conditions
 }
 for i := range allRSs {
  rs := allRSs[i]
  if rs == nil {
   continue
  }
  for _, c := range rs.Status.Conditions {
   if c.Type != apps.ReplicaSetReplicaFailure {
    continue
   }
   conditions = append(conditions, util.ReplicaSetToDeploymentCondition(c))
  }
 }
 return conditions
}

var nowFn = func() time.Time {
 return time.Now()
}

func (dc *DeploymentController) requeueStuckDeployment(d *apps.Deployment, newStatus apps.DeploymentStatus) time.Duration {
 _logClusterCodePath()
 defer _logClusterCodePath()
 currentCond := util.GetDeploymentCondition(d.Status, apps.DeploymentProgressing)
 if !util.HasProgressDeadline(d) || currentCond == nil {
  return time.Duration(-1)
 }
 if util.DeploymentComplete(d, &newStatus) || currentCond.Reason == util.TimedOutReason {
  return time.Duration(-1)
 }
 after := currentCond.LastUpdateTime.Time.Add(time.Duration(*d.Spec.ProgressDeadlineSeconds) * time.Second).Sub(nowFn())
 if after < time.Second {
  klog.V(4).Infof("Queueing up deployment %q for a progress check now", d.Name)
  dc.enqueueRateLimited(d)
  return time.Duration(0)
 }
 klog.V(4).Infof("Queueing up deployment %q for a progress check after %ds", d.Name, int(after.Seconds()))
 dc.enqueueAfter(d, after+time.Second)
 return after
}
