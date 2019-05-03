package deployment

import (
 "fmt"
 "strconv"
 "k8s.io/klog"
 apps "k8s.io/api/apps/v1"
 "k8s.io/api/core/v1"
 extensions "k8s.io/api/extensions/v1beta1"
 deploymentutil "k8s.io/kubernetes/pkg/controller/deployment/util"
)

func (dc *DeploymentController) rollback(d *apps.Deployment, rsList []*apps.ReplicaSet) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 newRS, allOldRSs, err := dc.getAllReplicaSetsAndSyncRevision(d, rsList, true)
 if err != nil {
  return err
 }
 allRSs := append(allOldRSs, newRS)
 rollbackTo := getRollbackTo(d)
 if rollbackTo.Revision == 0 {
  if rollbackTo.Revision = deploymentutil.LastRevision(allRSs); rollbackTo.Revision == 0 {
   dc.emitRollbackWarningEvent(d, deploymentutil.RollbackRevisionNotFound, "Unable to find last revision.")
   return dc.updateDeploymentAndClearRollbackTo(d)
  }
 }
 for _, rs := range allRSs {
  v, err := deploymentutil.Revision(rs)
  if err != nil {
   klog.V(4).Infof("Unable to extract revision from deployment's replica set %q: %v", rs.Name, err)
   continue
  }
  if v == rollbackTo.Revision {
   klog.V(4).Infof("Found replica set %q with desired revision %d", rs.Name, v)
   performedRollback, err := dc.rollbackToTemplate(d, rs)
   if performedRollback && err == nil {
    dc.emitRollbackNormalEvent(d, fmt.Sprintf("Rolled back deployment %q to revision %d", d.Name, rollbackTo.Revision))
   }
   return err
  }
 }
 dc.emitRollbackWarningEvent(d, deploymentutil.RollbackRevisionNotFound, "Unable to find the revision to rollback to.")
 return dc.updateDeploymentAndClearRollbackTo(d)
}
func (dc *DeploymentController) rollbackToTemplate(d *apps.Deployment, rs *apps.ReplicaSet) (bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 performedRollback := false
 if !deploymentutil.EqualIgnoreHash(&d.Spec.Template, &rs.Spec.Template) {
  klog.V(4).Infof("Rolling back deployment %q to template spec %+v", d.Name, rs.Spec.Template.Spec)
  deploymentutil.SetFromReplicaSetTemplate(d, rs.Spec.Template)
  deploymentutil.SetDeploymentAnnotationsTo(d, rs)
  performedRollback = true
 } else {
  klog.V(4).Infof("Rolling back to a revision that contains the same template as current deployment %q, skipping rollback...", d.Name)
  eventMsg := fmt.Sprintf("The rollback revision contains the same template as current deployment %q", d.Name)
  dc.emitRollbackWarningEvent(d, deploymentutil.RollbackTemplateUnchanged, eventMsg)
 }
 return performedRollback, dc.updateDeploymentAndClearRollbackTo(d)
}
func (dc *DeploymentController) emitRollbackWarningEvent(d *apps.Deployment, reason, message string) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 dc.eventRecorder.Eventf(d, v1.EventTypeWarning, reason, message)
}
func (dc *DeploymentController) emitRollbackNormalEvent(d *apps.Deployment, message string) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 dc.eventRecorder.Eventf(d, v1.EventTypeNormal, deploymentutil.RollbackDone, message)
}
func (dc *DeploymentController) updateDeploymentAndClearRollbackTo(d *apps.Deployment) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(4).Infof("Cleans up rollbackTo of deployment %q", d.Name)
 setRollbackTo(d, nil)
 _, err := dc.client.AppsV1().Deployments(d.Namespace).Update(d)
 return err
}
func getRollbackTo(d *apps.Deployment) *extensions.RollbackConfig {
 _logClusterCodePath()
 defer _logClusterCodePath()
 revision := d.Annotations[apps.DeprecatedRollbackTo]
 if revision == "" {
  return nil
 }
 revision64, err := strconv.ParseInt(revision, 10, 64)
 if err != nil {
  return nil
 }
 return &extensions.RollbackConfig{Revision: revision64}
}
func setRollbackTo(d *apps.Deployment, rollbackTo *extensions.RollbackConfig) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if rollbackTo == nil {
  delete(d.Annotations, apps.DeprecatedRollbackTo)
  return
 }
 if d.Annotations == nil {
  d.Annotations = make(map[string]string)
 }
 d.Annotations[apps.DeprecatedRollbackTo] = strconv.FormatInt(rollbackTo.Revision, 10)
}
