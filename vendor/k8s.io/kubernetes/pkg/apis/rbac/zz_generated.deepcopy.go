package rbac

import (
 v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 runtime "k8s.io/apimachinery/pkg/runtime"
)

func (in *AggregationRule) DeepCopyInto(out *AggregationRule) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.ClusterRoleSelectors != nil {
  in, out := &in.ClusterRoleSelectors, &out.ClusterRoleSelectors
  *out = make([]v1.LabelSelector, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 return
}
func (in *AggregationRule) DeepCopy() *AggregationRule {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(AggregationRule)
 in.DeepCopyInto(out)
 return out
}
func (in *ClusterRole) DeepCopyInto(out *ClusterRole) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
 if in.Rules != nil {
  in, out := &in.Rules, &out.Rules
  *out = make([]PolicyRule, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 if in.AggregationRule != nil {
  in, out := &in.AggregationRule, &out.AggregationRule
  *out = new(AggregationRule)
  (*in).DeepCopyInto(*out)
 }
 return
}
func (in *ClusterRole) DeepCopy() *ClusterRole {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ClusterRole)
 in.DeepCopyInto(out)
 return out
}
func (in *ClusterRole) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *ClusterRoleBinding) DeepCopyInto(out *ClusterRoleBinding) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
 if in.Subjects != nil {
  in, out := &in.Subjects, &out.Subjects
  *out = make([]Subject, len(*in))
  copy(*out, *in)
 }
 out.RoleRef = in.RoleRef
 return
}
func (in *ClusterRoleBinding) DeepCopy() *ClusterRoleBinding {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ClusterRoleBinding)
 in.DeepCopyInto(out)
 return out
}
func (in *ClusterRoleBinding) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *ClusterRoleBindingList) DeepCopyInto(out *ClusterRoleBindingList) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]ClusterRoleBinding, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 return
}
func (in *ClusterRoleBindingList) DeepCopy() *ClusterRoleBindingList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ClusterRoleBindingList)
 in.DeepCopyInto(out)
 return out
}
func (in *ClusterRoleBindingList) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *ClusterRoleList) DeepCopyInto(out *ClusterRoleList) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]ClusterRole, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 return
}
func (in *ClusterRoleList) DeepCopy() *ClusterRoleList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ClusterRoleList)
 in.DeepCopyInto(out)
 return out
}
func (in *ClusterRoleList) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *PolicyRule) DeepCopyInto(out *PolicyRule) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Verbs != nil {
  in, out := &in.Verbs, &out.Verbs
  *out = make([]string, len(*in))
  copy(*out, *in)
 }
 if in.APIGroups != nil {
  in, out := &in.APIGroups, &out.APIGroups
  *out = make([]string, len(*in))
  copy(*out, *in)
 }
 if in.Resources != nil {
  in, out := &in.Resources, &out.Resources
  *out = make([]string, len(*in))
  copy(*out, *in)
 }
 if in.ResourceNames != nil {
  in, out := &in.ResourceNames, &out.ResourceNames
  *out = make([]string, len(*in))
  copy(*out, *in)
 }
 if in.NonResourceURLs != nil {
  in, out := &in.NonResourceURLs, &out.NonResourceURLs
  *out = make([]string, len(*in))
  copy(*out, *in)
 }
 return
}
func (in *PolicyRule) DeepCopy() *PolicyRule {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(PolicyRule)
 in.DeepCopyInto(out)
 return out
}
func (in *Role) DeepCopyInto(out *Role) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
 if in.Rules != nil {
  in, out := &in.Rules, &out.Rules
  *out = make([]PolicyRule, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 return
}
func (in *Role) DeepCopy() *Role {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(Role)
 in.DeepCopyInto(out)
 return out
}
func (in *Role) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *RoleBinding) DeepCopyInto(out *RoleBinding) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
 if in.Subjects != nil {
  in, out := &in.Subjects, &out.Subjects
  *out = make([]Subject, len(*in))
  copy(*out, *in)
 }
 out.RoleRef = in.RoleRef
 return
}
func (in *RoleBinding) DeepCopy() *RoleBinding {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(RoleBinding)
 in.DeepCopyInto(out)
 return out
}
func (in *RoleBinding) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *RoleBindingList) DeepCopyInto(out *RoleBindingList) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]RoleBinding, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 return
}
func (in *RoleBindingList) DeepCopy() *RoleBindingList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(RoleBindingList)
 in.DeepCopyInto(out)
 return out
}
func (in *RoleBindingList) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *RoleList) DeepCopyInto(out *RoleList) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]Role, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 return
}
func (in *RoleList) DeepCopy() *RoleList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(RoleList)
 in.DeepCopyInto(out)
 return out
}
func (in *RoleList) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *RoleRef) DeepCopyInto(out *RoleRef) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 return
}
func (in *RoleRef) DeepCopy() *RoleRef {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(RoleRef)
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
func (in *Subject) DeepCopyInto(out *Subject) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 return
}
func (in *Subject) DeepCopy() *Subject {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(Subject)
 in.DeepCopyInto(out)
 return out
}
