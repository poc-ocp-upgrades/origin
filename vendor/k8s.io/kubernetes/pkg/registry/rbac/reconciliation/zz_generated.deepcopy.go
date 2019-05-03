package reconciliation

import (
 v1 "k8s.io/api/rbac/v1"
)

func (in *ClusterRoleBindingAdapter) DeepCopyInto(out *ClusterRoleBindingAdapter) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.ClusterRoleBinding != nil {
  in, out := &in.ClusterRoleBinding, &out.ClusterRoleBinding
  *out = new(v1.ClusterRoleBinding)
  (*in).DeepCopyInto(*out)
 }
 return
}
func (in *ClusterRoleBindingAdapter) DeepCopy() *ClusterRoleBindingAdapter {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ClusterRoleBindingAdapter)
 in.DeepCopyInto(out)
 return out
}
func (in ClusterRoleBindingAdapter) DeepCopyRoleBinding() RoleBinding {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return *in.DeepCopy()
}
func (in *ClusterRoleRuleOwner) DeepCopyInto(out *ClusterRoleRuleOwner) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.ClusterRole != nil {
  in, out := &in.ClusterRole, &out.ClusterRole
  *out = new(v1.ClusterRole)
  (*in).DeepCopyInto(*out)
 }
 return
}
func (in *ClusterRoleRuleOwner) DeepCopy() *ClusterRoleRuleOwner {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ClusterRoleRuleOwner)
 in.DeepCopyInto(out)
 return out
}
func (in ClusterRoleRuleOwner) DeepCopyRuleOwner() RuleOwner {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return *in.DeepCopy()
}
func (in *RoleBindingAdapter) DeepCopyInto(out *RoleBindingAdapter) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.RoleBinding != nil {
  in, out := &in.RoleBinding, &out.RoleBinding
  *out = new(v1.RoleBinding)
  (*in).DeepCopyInto(*out)
 }
 return
}
func (in *RoleBindingAdapter) DeepCopy() *RoleBindingAdapter {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(RoleBindingAdapter)
 in.DeepCopyInto(out)
 return out
}
func (in RoleBindingAdapter) DeepCopyRoleBinding() RoleBinding {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return *in.DeepCopy()
}
func (in *RoleRuleOwner) DeepCopyInto(out *RoleRuleOwner) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Role != nil {
  in, out := &in.Role, &out.Role
  *out = new(v1.Role)
  (*in).DeepCopyInto(*out)
 }
 return
}
func (in *RoleRuleOwner) DeepCopy() *RoleRuleOwner {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(RoleRuleOwner)
 in.DeepCopyInto(out)
 return out
}
func (in RoleRuleOwner) DeepCopyRuleOwner() RuleOwner {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return *in.DeepCopy()
}
