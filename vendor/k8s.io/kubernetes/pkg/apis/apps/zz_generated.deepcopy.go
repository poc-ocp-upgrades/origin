package apps

import (
 v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 runtime "k8s.io/apimachinery/pkg/runtime"
 core "k8s.io/kubernetes/pkg/apis/core"
)

func (in *ControllerRevision) DeepCopyInto(out *ControllerRevision) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
 if in.Data != nil {
  out.Data = in.Data.DeepCopyObject()
 }
 return
}
func (in *ControllerRevision) DeepCopy() *ControllerRevision {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ControllerRevision)
 in.DeepCopyInto(out)
 return out
}
func (in *ControllerRevision) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *ControllerRevisionList) DeepCopyInto(out *ControllerRevisionList) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]ControllerRevision, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 return
}
func (in *ControllerRevisionList) DeepCopy() *ControllerRevisionList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ControllerRevisionList)
 in.DeepCopyInto(out)
 return out
}
func (in *ControllerRevisionList) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *DaemonSet) DeepCopyInto(out *DaemonSet) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
 in.Spec.DeepCopyInto(&out.Spec)
 in.Status.DeepCopyInto(&out.Status)
 return
}
func (in *DaemonSet) DeepCopy() *DaemonSet {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(DaemonSet)
 in.DeepCopyInto(out)
 return out
}
func (in *DaemonSet) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *DaemonSetCondition) DeepCopyInto(out *DaemonSetCondition) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
 return
}
func (in *DaemonSetCondition) DeepCopy() *DaemonSetCondition {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(DaemonSetCondition)
 in.DeepCopyInto(out)
 return out
}
func (in *DaemonSetList) DeepCopyInto(out *DaemonSetList) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]DaemonSet, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 return
}
func (in *DaemonSetList) DeepCopy() *DaemonSetList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(DaemonSetList)
 in.DeepCopyInto(out)
 return out
}
func (in *DaemonSetList) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *DaemonSetSpec) DeepCopyInto(out *DaemonSetSpec) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Selector != nil {
  in, out := &in.Selector, &out.Selector
  *out = new(v1.LabelSelector)
  (*in).DeepCopyInto(*out)
 }
 in.Template.DeepCopyInto(&out.Template)
 in.UpdateStrategy.DeepCopyInto(&out.UpdateStrategy)
 if in.RevisionHistoryLimit != nil {
  in, out := &in.RevisionHistoryLimit, &out.RevisionHistoryLimit
  *out = new(int32)
  **out = **in
 }
 return
}
func (in *DaemonSetSpec) DeepCopy() *DaemonSetSpec {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(DaemonSetSpec)
 in.DeepCopyInto(out)
 return out
}
func (in *DaemonSetStatus) DeepCopyInto(out *DaemonSetStatus) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.CollisionCount != nil {
  in, out := &in.CollisionCount, &out.CollisionCount
  *out = new(int32)
  **out = **in
 }
 if in.Conditions != nil {
  in, out := &in.Conditions, &out.Conditions
  *out = make([]DaemonSetCondition, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 return
}
func (in *DaemonSetStatus) DeepCopy() *DaemonSetStatus {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(DaemonSetStatus)
 in.DeepCopyInto(out)
 return out
}
func (in *DaemonSetUpdateStrategy) DeepCopyInto(out *DaemonSetUpdateStrategy) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.RollingUpdate != nil {
  in, out := &in.RollingUpdate, &out.RollingUpdate
  *out = new(RollingUpdateDaemonSet)
  **out = **in
 }
 return
}
func (in *DaemonSetUpdateStrategy) DeepCopy() *DaemonSetUpdateStrategy {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(DaemonSetUpdateStrategy)
 in.DeepCopyInto(out)
 return out
}
func (in *Deployment) DeepCopyInto(out *Deployment) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
 in.Spec.DeepCopyInto(&out.Spec)
 in.Status.DeepCopyInto(&out.Status)
 return
}
func (in *Deployment) DeepCopy() *Deployment {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(Deployment)
 in.DeepCopyInto(out)
 return out
}
func (in *Deployment) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *DeploymentCondition) DeepCopyInto(out *DeploymentCondition) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
 in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
 return
}
func (in *DeploymentCondition) DeepCopy() *DeploymentCondition {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(DeploymentCondition)
 in.DeepCopyInto(out)
 return out
}
func (in *DeploymentList) DeepCopyInto(out *DeploymentList) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]Deployment, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 return
}
func (in *DeploymentList) DeepCopy() *DeploymentList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(DeploymentList)
 in.DeepCopyInto(out)
 return out
}
func (in *DeploymentList) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *DeploymentRollback) DeepCopyInto(out *DeploymentRollback) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 if in.UpdatedAnnotations != nil {
  in, out := &in.UpdatedAnnotations, &out.UpdatedAnnotations
  *out = make(map[string]string, len(*in))
  for key, val := range *in {
   (*out)[key] = val
  }
 }
 out.RollbackTo = in.RollbackTo
 return
}
func (in *DeploymentRollback) DeepCopy() *DeploymentRollback {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(DeploymentRollback)
 in.DeepCopyInto(out)
 return out
}
func (in *DeploymentRollback) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *DeploymentSpec) DeepCopyInto(out *DeploymentSpec) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Selector != nil {
  in, out := &in.Selector, &out.Selector
  *out = new(v1.LabelSelector)
  (*in).DeepCopyInto(*out)
 }
 in.Template.DeepCopyInto(&out.Template)
 in.Strategy.DeepCopyInto(&out.Strategy)
 if in.RevisionHistoryLimit != nil {
  in, out := &in.RevisionHistoryLimit, &out.RevisionHistoryLimit
  *out = new(int32)
  **out = **in
 }
 if in.RollbackTo != nil {
  in, out := &in.RollbackTo, &out.RollbackTo
  *out = new(RollbackConfig)
  **out = **in
 }
 if in.ProgressDeadlineSeconds != nil {
  in, out := &in.ProgressDeadlineSeconds, &out.ProgressDeadlineSeconds
  *out = new(int32)
  **out = **in
 }
 return
}
func (in *DeploymentSpec) DeepCopy() *DeploymentSpec {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(DeploymentSpec)
 in.DeepCopyInto(out)
 return out
}
func (in *DeploymentStatus) DeepCopyInto(out *DeploymentStatus) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Conditions != nil {
  in, out := &in.Conditions, &out.Conditions
  *out = make([]DeploymentCondition, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 if in.CollisionCount != nil {
  in, out := &in.CollisionCount, &out.CollisionCount
  *out = new(int32)
  **out = **in
 }
 return
}
func (in *DeploymentStatus) DeepCopy() *DeploymentStatus {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(DeploymentStatus)
 in.DeepCopyInto(out)
 return out
}
func (in *DeploymentStrategy) DeepCopyInto(out *DeploymentStrategy) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.RollingUpdate != nil {
  in, out := &in.RollingUpdate, &out.RollingUpdate
  *out = new(RollingUpdateDeployment)
  **out = **in
 }
 return
}
func (in *DeploymentStrategy) DeepCopy() *DeploymentStrategy {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(DeploymentStrategy)
 in.DeepCopyInto(out)
 return out
}
func (in *ReplicaSet) DeepCopyInto(out *ReplicaSet) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
 in.Spec.DeepCopyInto(&out.Spec)
 in.Status.DeepCopyInto(&out.Status)
 return
}
func (in *ReplicaSet) DeepCopy() *ReplicaSet {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ReplicaSet)
 in.DeepCopyInto(out)
 return out
}
func (in *ReplicaSet) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *ReplicaSetCondition) DeepCopyInto(out *ReplicaSetCondition) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
 return
}
func (in *ReplicaSetCondition) DeepCopy() *ReplicaSetCondition {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ReplicaSetCondition)
 in.DeepCopyInto(out)
 return out
}
func (in *ReplicaSetList) DeepCopyInto(out *ReplicaSetList) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]ReplicaSet, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 return
}
func (in *ReplicaSetList) DeepCopy() *ReplicaSetList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ReplicaSetList)
 in.DeepCopyInto(out)
 return out
}
func (in *ReplicaSetList) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *ReplicaSetSpec) DeepCopyInto(out *ReplicaSetSpec) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Selector != nil {
  in, out := &in.Selector, &out.Selector
  *out = new(v1.LabelSelector)
  (*in).DeepCopyInto(*out)
 }
 in.Template.DeepCopyInto(&out.Template)
 return
}
func (in *ReplicaSetSpec) DeepCopy() *ReplicaSetSpec {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ReplicaSetSpec)
 in.DeepCopyInto(out)
 return out
}
func (in *ReplicaSetStatus) DeepCopyInto(out *ReplicaSetStatus) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Conditions != nil {
  in, out := &in.Conditions, &out.Conditions
  *out = make([]ReplicaSetCondition, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 return
}
func (in *ReplicaSetStatus) DeepCopy() *ReplicaSetStatus {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ReplicaSetStatus)
 in.DeepCopyInto(out)
 return out
}
func (in *RollbackConfig) DeepCopyInto(out *RollbackConfig) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 return
}
func (in *RollbackConfig) DeepCopy() *RollbackConfig {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(RollbackConfig)
 in.DeepCopyInto(out)
 return out
}
func (in *RollingUpdateDaemonSet) DeepCopyInto(out *RollingUpdateDaemonSet) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.MaxUnavailable = in.MaxUnavailable
 return
}
func (in *RollingUpdateDaemonSet) DeepCopy() *RollingUpdateDaemonSet {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(RollingUpdateDaemonSet)
 in.DeepCopyInto(out)
 return out
}
func (in *RollingUpdateDeployment) DeepCopyInto(out *RollingUpdateDeployment) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.MaxUnavailable = in.MaxUnavailable
 out.MaxSurge = in.MaxSurge
 return
}
func (in *RollingUpdateDeployment) DeepCopy() *RollingUpdateDeployment {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(RollingUpdateDeployment)
 in.DeepCopyInto(out)
 return out
}
func (in *RollingUpdateStatefulSetStrategy) DeepCopyInto(out *RollingUpdateStatefulSetStrategy) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 return
}
func (in *RollingUpdateStatefulSetStrategy) DeepCopy() *RollingUpdateStatefulSetStrategy {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(RollingUpdateStatefulSetStrategy)
 in.DeepCopyInto(out)
 return out
}
func (in *StatefulSet) DeepCopyInto(out *StatefulSet) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
 in.Spec.DeepCopyInto(&out.Spec)
 in.Status.DeepCopyInto(&out.Status)
 return
}
func (in *StatefulSet) DeepCopy() *StatefulSet {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(StatefulSet)
 in.DeepCopyInto(out)
 return out
}
func (in *StatefulSet) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *StatefulSetCondition) DeepCopyInto(out *StatefulSetCondition) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
 return
}
func (in *StatefulSetCondition) DeepCopy() *StatefulSetCondition {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(StatefulSetCondition)
 in.DeepCopyInto(out)
 return out
}
func (in *StatefulSetList) DeepCopyInto(out *StatefulSetList) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]StatefulSet, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 return
}
func (in *StatefulSetList) DeepCopy() *StatefulSetList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(StatefulSetList)
 in.DeepCopyInto(out)
 return out
}
func (in *StatefulSetList) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *StatefulSetSpec) DeepCopyInto(out *StatefulSetSpec) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Selector != nil {
  in, out := &in.Selector, &out.Selector
  *out = new(v1.LabelSelector)
  (*in).DeepCopyInto(*out)
 }
 in.Template.DeepCopyInto(&out.Template)
 if in.VolumeClaimTemplates != nil {
  in, out := &in.VolumeClaimTemplates, &out.VolumeClaimTemplates
  *out = make([]core.PersistentVolumeClaim, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 in.UpdateStrategy.DeepCopyInto(&out.UpdateStrategy)
 if in.RevisionHistoryLimit != nil {
  in, out := &in.RevisionHistoryLimit, &out.RevisionHistoryLimit
  *out = new(int32)
  **out = **in
 }
 return
}
func (in *StatefulSetSpec) DeepCopy() *StatefulSetSpec {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(StatefulSetSpec)
 in.DeepCopyInto(out)
 return out
}
func (in *StatefulSetStatus) DeepCopyInto(out *StatefulSetStatus) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.ObservedGeneration != nil {
  in, out := &in.ObservedGeneration, &out.ObservedGeneration
  *out = new(int64)
  **out = **in
 }
 if in.CollisionCount != nil {
  in, out := &in.CollisionCount, &out.CollisionCount
  *out = new(int32)
  **out = **in
 }
 if in.Conditions != nil {
  in, out := &in.Conditions, &out.Conditions
  *out = make([]StatefulSetCondition, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 return
}
func (in *StatefulSetStatus) DeepCopy() *StatefulSetStatus {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(StatefulSetStatus)
 in.DeepCopyInto(out)
 return out
}
func (in *StatefulSetUpdateStrategy) DeepCopyInto(out *StatefulSetUpdateStrategy) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.RollingUpdate != nil {
  in, out := &in.RollingUpdate, &out.RollingUpdate
  *out = new(RollingUpdateStatefulSetStrategy)
  **out = **in
 }
 return
}
func (in *StatefulSetUpdateStrategy) DeepCopy() *StatefulSetUpdateStrategy {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(StatefulSetUpdateStrategy)
 in.DeepCopyInto(out)
 return out
}
