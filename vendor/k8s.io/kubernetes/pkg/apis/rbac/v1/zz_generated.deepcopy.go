package v1

func (in *ClusterRoleBindingBuilder) DeepCopyInto(out *ClusterRoleBindingBuilder) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 in.ClusterRoleBinding.DeepCopyInto(&out.ClusterRoleBinding)
 return
}
func (in *ClusterRoleBindingBuilder) DeepCopy() *ClusterRoleBindingBuilder {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ClusterRoleBindingBuilder)
 in.DeepCopyInto(out)
 return out
}
func (in *PolicyRuleBuilder) DeepCopyInto(out *PolicyRuleBuilder) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 in.PolicyRule.DeepCopyInto(&out.PolicyRule)
 return
}
func (in *PolicyRuleBuilder) DeepCopy() *PolicyRuleBuilder {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(PolicyRuleBuilder)
 in.DeepCopyInto(out)
 return out
}
func (in *RoleBindingBuilder) DeepCopyInto(out *RoleBindingBuilder) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 in.RoleBinding.DeepCopyInto(&out.RoleBinding)
 return
}
func (in *RoleBindingBuilder) DeepCopy() *RoleBindingBuilder {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(RoleBindingBuilder)
 in.DeepCopyInto(out)
 return out
}
func (in SortableRuleSlice) DeepCopyInto(out *SortableRuleSlice) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 {
  in := &in
  *out = make(SortableRuleSlice, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
  return
 }
}
func (in SortableRuleSlice) DeepCopy() SortableRuleSlice {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(SortableRuleSlice)
 in.DeepCopyInto(out)
 return *out
}
