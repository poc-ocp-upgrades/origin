package v1alpha1

import (
 unsafe "unsafe"
 v1alpha1 "k8s.io/api/rbac/v1alpha1"
 v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 conversion "k8s.io/apimachinery/pkg/conversion"
 runtime "k8s.io/apimachinery/pkg/runtime"
 rbac "k8s.io/kubernetes/pkg/apis/rbac"
)

func init() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 localSchemeBuilder.Register(RegisterConversions)
}
func RegisterConversions(s *runtime.Scheme) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := s.AddGeneratedConversionFunc((*v1alpha1.AggregationRule)(nil), (*rbac.AggregationRule)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1alpha1_AggregationRule_To_rbac_AggregationRule(a.(*v1alpha1.AggregationRule), b.(*rbac.AggregationRule), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*rbac.AggregationRule)(nil), (*v1alpha1.AggregationRule)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_rbac_AggregationRule_To_v1alpha1_AggregationRule(a.(*rbac.AggregationRule), b.(*v1alpha1.AggregationRule), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1alpha1.ClusterRole)(nil), (*rbac.ClusterRole)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1alpha1_ClusterRole_To_rbac_ClusterRole(a.(*v1alpha1.ClusterRole), b.(*rbac.ClusterRole), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*rbac.ClusterRole)(nil), (*v1alpha1.ClusterRole)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_rbac_ClusterRole_To_v1alpha1_ClusterRole(a.(*rbac.ClusterRole), b.(*v1alpha1.ClusterRole), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1alpha1.ClusterRoleBinding)(nil), (*rbac.ClusterRoleBinding)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1alpha1_ClusterRoleBinding_To_rbac_ClusterRoleBinding(a.(*v1alpha1.ClusterRoleBinding), b.(*rbac.ClusterRoleBinding), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*rbac.ClusterRoleBinding)(nil), (*v1alpha1.ClusterRoleBinding)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_rbac_ClusterRoleBinding_To_v1alpha1_ClusterRoleBinding(a.(*rbac.ClusterRoleBinding), b.(*v1alpha1.ClusterRoleBinding), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1alpha1.ClusterRoleBindingList)(nil), (*rbac.ClusterRoleBindingList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1alpha1_ClusterRoleBindingList_To_rbac_ClusterRoleBindingList(a.(*v1alpha1.ClusterRoleBindingList), b.(*rbac.ClusterRoleBindingList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*rbac.ClusterRoleBindingList)(nil), (*v1alpha1.ClusterRoleBindingList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_rbac_ClusterRoleBindingList_To_v1alpha1_ClusterRoleBindingList(a.(*rbac.ClusterRoleBindingList), b.(*v1alpha1.ClusterRoleBindingList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1alpha1.ClusterRoleList)(nil), (*rbac.ClusterRoleList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1alpha1_ClusterRoleList_To_rbac_ClusterRoleList(a.(*v1alpha1.ClusterRoleList), b.(*rbac.ClusterRoleList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*rbac.ClusterRoleList)(nil), (*v1alpha1.ClusterRoleList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_rbac_ClusterRoleList_To_v1alpha1_ClusterRoleList(a.(*rbac.ClusterRoleList), b.(*v1alpha1.ClusterRoleList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1alpha1.PolicyRule)(nil), (*rbac.PolicyRule)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1alpha1_PolicyRule_To_rbac_PolicyRule(a.(*v1alpha1.PolicyRule), b.(*rbac.PolicyRule), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*rbac.PolicyRule)(nil), (*v1alpha1.PolicyRule)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_rbac_PolicyRule_To_v1alpha1_PolicyRule(a.(*rbac.PolicyRule), b.(*v1alpha1.PolicyRule), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1alpha1.Role)(nil), (*rbac.Role)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1alpha1_Role_To_rbac_Role(a.(*v1alpha1.Role), b.(*rbac.Role), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*rbac.Role)(nil), (*v1alpha1.Role)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_rbac_Role_To_v1alpha1_Role(a.(*rbac.Role), b.(*v1alpha1.Role), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1alpha1.RoleBinding)(nil), (*rbac.RoleBinding)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1alpha1_RoleBinding_To_rbac_RoleBinding(a.(*v1alpha1.RoleBinding), b.(*rbac.RoleBinding), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*rbac.RoleBinding)(nil), (*v1alpha1.RoleBinding)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_rbac_RoleBinding_To_v1alpha1_RoleBinding(a.(*rbac.RoleBinding), b.(*v1alpha1.RoleBinding), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1alpha1.RoleBindingList)(nil), (*rbac.RoleBindingList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1alpha1_RoleBindingList_To_rbac_RoleBindingList(a.(*v1alpha1.RoleBindingList), b.(*rbac.RoleBindingList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*rbac.RoleBindingList)(nil), (*v1alpha1.RoleBindingList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_rbac_RoleBindingList_To_v1alpha1_RoleBindingList(a.(*rbac.RoleBindingList), b.(*v1alpha1.RoleBindingList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1alpha1.RoleList)(nil), (*rbac.RoleList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1alpha1_RoleList_To_rbac_RoleList(a.(*v1alpha1.RoleList), b.(*rbac.RoleList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*rbac.RoleList)(nil), (*v1alpha1.RoleList)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_rbac_RoleList_To_v1alpha1_RoleList(a.(*rbac.RoleList), b.(*v1alpha1.RoleList), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1alpha1.RoleRef)(nil), (*rbac.RoleRef)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1alpha1_RoleRef_To_rbac_RoleRef(a.(*v1alpha1.RoleRef), b.(*rbac.RoleRef), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*rbac.RoleRef)(nil), (*v1alpha1.RoleRef)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_rbac_RoleRef_To_v1alpha1_RoleRef(a.(*rbac.RoleRef), b.(*v1alpha1.RoleRef), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1alpha1.Subject)(nil), (*rbac.Subject)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1alpha1_Subject_To_rbac_Subject(a.(*v1alpha1.Subject), b.(*rbac.Subject), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*rbac.Subject)(nil), (*v1alpha1.Subject)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_rbac_Subject_To_v1alpha1_Subject(a.(*rbac.Subject), b.(*v1alpha1.Subject), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*rbac.Subject)(nil), (*v1alpha1.Subject)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_rbac_Subject_To_v1alpha1_Subject(a.(*rbac.Subject), b.(*v1alpha1.Subject), scope)
 }); err != nil {
  return err
 }
 if err := s.AddConversionFunc((*v1alpha1.Subject)(nil), (*rbac.Subject)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1alpha1_Subject_To_rbac_Subject(a.(*v1alpha1.Subject), b.(*rbac.Subject), scope)
 }); err != nil {
  return err
 }
 return nil
}
func autoConvert_v1alpha1_AggregationRule_To_rbac_AggregationRule(in *v1alpha1.AggregationRule, out *rbac.AggregationRule, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ClusterRoleSelectors = *(*[]v1.LabelSelector)(unsafe.Pointer(&in.ClusterRoleSelectors))
 return nil
}
func Convert_v1alpha1_AggregationRule_To_rbac_AggregationRule(in *v1alpha1.AggregationRule, out *rbac.AggregationRule, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1alpha1_AggregationRule_To_rbac_AggregationRule(in, out, s)
}
func autoConvert_rbac_AggregationRule_To_v1alpha1_AggregationRule(in *rbac.AggregationRule, out *v1alpha1.AggregationRule, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ClusterRoleSelectors = *(*[]v1.LabelSelector)(unsafe.Pointer(&in.ClusterRoleSelectors))
 return nil
}
func Convert_rbac_AggregationRule_To_v1alpha1_AggregationRule(in *rbac.AggregationRule, out *v1alpha1.AggregationRule, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_rbac_AggregationRule_To_v1alpha1_AggregationRule(in, out, s)
}
func autoConvert_v1alpha1_ClusterRole_To_rbac_ClusterRole(in *v1alpha1.ClusterRole, out *rbac.ClusterRole, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 out.Rules = *(*[]rbac.PolicyRule)(unsafe.Pointer(&in.Rules))
 out.AggregationRule = (*rbac.AggregationRule)(unsafe.Pointer(in.AggregationRule))
 return nil
}
func Convert_v1alpha1_ClusterRole_To_rbac_ClusterRole(in *v1alpha1.ClusterRole, out *rbac.ClusterRole, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1alpha1_ClusterRole_To_rbac_ClusterRole(in, out, s)
}
func autoConvert_rbac_ClusterRole_To_v1alpha1_ClusterRole(in *rbac.ClusterRole, out *v1alpha1.ClusterRole, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 out.Rules = *(*[]v1alpha1.PolicyRule)(unsafe.Pointer(&in.Rules))
 out.AggregationRule = (*v1alpha1.AggregationRule)(unsafe.Pointer(in.AggregationRule))
 return nil
}
func Convert_rbac_ClusterRole_To_v1alpha1_ClusterRole(in *rbac.ClusterRole, out *v1alpha1.ClusterRole, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_rbac_ClusterRole_To_v1alpha1_ClusterRole(in, out, s)
}
func autoConvert_v1alpha1_ClusterRoleBinding_To_rbac_ClusterRoleBinding(in *v1alpha1.ClusterRoleBinding, out *rbac.ClusterRoleBinding, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if in.Subjects != nil {
  in, out := &in.Subjects, &out.Subjects
  *out = make([]rbac.Subject, len(*in))
  for i := range *in {
   if err := Convert_v1alpha1_Subject_To_rbac_Subject(&(*in)[i], &(*out)[i], s); err != nil {
    return err
   }
  }
 } else {
  out.Subjects = nil
 }
 if err := Convert_v1alpha1_RoleRef_To_rbac_RoleRef(&in.RoleRef, &out.RoleRef, s); err != nil {
  return err
 }
 return nil
}
func Convert_v1alpha1_ClusterRoleBinding_To_rbac_ClusterRoleBinding(in *v1alpha1.ClusterRoleBinding, out *rbac.ClusterRoleBinding, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1alpha1_ClusterRoleBinding_To_rbac_ClusterRoleBinding(in, out, s)
}
func autoConvert_rbac_ClusterRoleBinding_To_v1alpha1_ClusterRoleBinding(in *rbac.ClusterRoleBinding, out *v1alpha1.ClusterRoleBinding, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if in.Subjects != nil {
  in, out := &in.Subjects, &out.Subjects
  *out = make([]v1alpha1.Subject, len(*in))
  for i := range *in {
   if err := Convert_rbac_Subject_To_v1alpha1_Subject(&(*in)[i], &(*out)[i], s); err != nil {
    return err
   }
  }
 } else {
  out.Subjects = nil
 }
 if err := Convert_rbac_RoleRef_To_v1alpha1_RoleRef(&in.RoleRef, &out.RoleRef, s); err != nil {
  return err
 }
 return nil
}
func Convert_rbac_ClusterRoleBinding_To_v1alpha1_ClusterRoleBinding(in *rbac.ClusterRoleBinding, out *v1alpha1.ClusterRoleBinding, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_rbac_ClusterRoleBinding_To_v1alpha1_ClusterRoleBinding(in, out, s)
}
func autoConvert_v1alpha1_ClusterRoleBindingList_To_rbac_ClusterRoleBindingList(in *v1alpha1.ClusterRoleBindingList, out *rbac.ClusterRoleBindingList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]rbac.ClusterRoleBinding, len(*in))
  for i := range *in {
   if err := Convert_v1alpha1_ClusterRoleBinding_To_rbac_ClusterRoleBinding(&(*in)[i], &(*out)[i], s); err != nil {
    return err
   }
  }
 } else {
  out.Items = nil
 }
 return nil
}
func Convert_v1alpha1_ClusterRoleBindingList_To_rbac_ClusterRoleBindingList(in *v1alpha1.ClusterRoleBindingList, out *rbac.ClusterRoleBindingList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1alpha1_ClusterRoleBindingList_To_rbac_ClusterRoleBindingList(in, out, s)
}
func autoConvert_rbac_ClusterRoleBindingList_To_v1alpha1_ClusterRoleBindingList(in *rbac.ClusterRoleBindingList, out *v1alpha1.ClusterRoleBindingList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]v1alpha1.ClusterRoleBinding, len(*in))
  for i := range *in {
   if err := Convert_rbac_ClusterRoleBinding_To_v1alpha1_ClusterRoleBinding(&(*in)[i], &(*out)[i], s); err != nil {
    return err
   }
  }
 } else {
  out.Items = nil
 }
 return nil
}
func Convert_rbac_ClusterRoleBindingList_To_v1alpha1_ClusterRoleBindingList(in *rbac.ClusterRoleBindingList, out *v1alpha1.ClusterRoleBindingList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_rbac_ClusterRoleBindingList_To_v1alpha1_ClusterRoleBindingList(in, out, s)
}
func autoConvert_v1alpha1_ClusterRoleList_To_rbac_ClusterRoleList(in *v1alpha1.ClusterRoleList, out *rbac.ClusterRoleList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 out.Items = *(*[]rbac.ClusterRole)(unsafe.Pointer(&in.Items))
 return nil
}
func Convert_v1alpha1_ClusterRoleList_To_rbac_ClusterRoleList(in *v1alpha1.ClusterRoleList, out *rbac.ClusterRoleList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1alpha1_ClusterRoleList_To_rbac_ClusterRoleList(in, out, s)
}
func autoConvert_rbac_ClusterRoleList_To_v1alpha1_ClusterRoleList(in *rbac.ClusterRoleList, out *v1alpha1.ClusterRoleList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 out.Items = *(*[]v1alpha1.ClusterRole)(unsafe.Pointer(&in.Items))
 return nil
}
func Convert_rbac_ClusterRoleList_To_v1alpha1_ClusterRoleList(in *rbac.ClusterRoleList, out *v1alpha1.ClusterRoleList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_rbac_ClusterRoleList_To_v1alpha1_ClusterRoleList(in, out, s)
}
func autoConvert_v1alpha1_PolicyRule_To_rbac_PolicyRule(in *v1alpha1.PolicyRule, out *rbac.PolicyRule, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Verbs = *(*[]string)(unsafe.Pointer(&in.Verbs))
 out.APIGroups = *(*[]string)(unsafe.Pointer(&in.APIGroups))
 out.Resources = *(*[]string)(unsafe.Pointer(&in.Resources))
 out.ResourceNames = *(*[]string)(unsafe.Pointer(&in.ResourceNames))
 out.NonResourceURLs = *(*[]string)(unsafe.Pointer(&in.NonResourceURLs))
 return nil
}
func Convert_v1alpha1_PolicyRule_To_rbac_PolicyRule(in *v1alpha1.PolicyRule, out *rbac.PolicyRule, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1alpha1_PolicyRule_To_rbac_PolicyRule(in, out, s)
}
func autoConvert_rbac_PolicyRule_To_v1alpha1_PolicyRule(in *rbac.PolicyRule, out *v1alpha1.PolicyRule, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Verbs = *(*[]string)(unsafe.Pointer(&in.Verbs))
 out.APIGroups = *(*[]string)(unsafe.Pointer(&in.APIGroups))
 out.Resources = *(*[]string)(unsafe.Pointer(&in.Resources))
 out.ResourceNames = *(*[]string)(unsafe.Pointer(&in.ResourceNames))
 out.NonResourceURLs = *(*[]string)(unsafe.Pointer(&in.NonResourceURLs))
 return nil
}
func Convert_rbac_PolicyRule_To_v1alpha1_PolicyRule(in *rbac.PolicyRule, out *v1alpha1.PolicyRule, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_rbac_PolicyRule_To_v1alpha1_PolicyRule(in, out, s)
}
func autoConvert_v1alpha1_Role_To_rbac_Role(in *v1alpha1.Role, out *rbac.Role, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 out.Rules = *(*[]rbac.PolicyRule)(unsafe.Pointer(&in.Rules))
 return nil
}
func Convert_v1alpha1_Role_To_rbac_Role(in *v1alpha1.Role, out *rbac.Role, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1alpha1_Role_To_rbac_Role(in, out, s)
}
func autoConvert_rbac_Role_To_v1alpha1_Role(in *rbac.Role, out *v1alpha1.Role, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 out.Rules = *(*[]v1alpha1.PolicyRule)(unsafe.Pointer(&in.Rules))
 return nil
}
func Convert_rbac_Role_To_v1alpha1_Role(in *rbac.Role, out *v1alpha1.Role, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_rbac_Role_To_v1alpha1_Role(in, out, s)
}
func autoConvert_v1alpha1_RoleBinding_To_rbac_RoleBinding(in *v1alpha1.RoleBinding, out *rbac.RoleBinding, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if in.Subjects != nil {
  in, out := &in.Subjects, &out.Subjects
  *out = make([]rbac.Subject, len(*in))
  for i := range *in {
   if err := Convert_v1alpha1_Subject_To_rbac_Subject(&(*in)[i], &(*out)[i], s); err != nil {
    return err
   }
  }
 } else {
  out.Subjects = nil
 }
 if err := Convert_v1alpha1_RoleRef_To_rbac_RoleRef(&in.RoleRef, &out.RoleRef, s); err != nil {
  return err
 }
 return nil
}
func Convert_v1alpha1_RoleBinding_To_rbac_RoleBinding(in *v1alpha1.RoleBinding, out *rbac.RoleBinding, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1alpha1_RoleBinding_To_rbac_RoleBinding(in, out, s)
}
func autoConvert_rbac_RoleBinding_To_v1alpha1_RoleBinding(in *rbac.RoleBinding, out *v1alpha1.RoleBinding, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if in.Subjects != nil {
  in, out := &in.Subjects, &out.Subjects
  *out = make([]v1alpha1.Subject, len(*in))
  for i := range *in {
   if err := Convert_rbac_Subject_To_v1alpha1_Subject(&(*in)[i], &(*out)[i], s); err != nil {
    return err
   }
  }
 } else {
  out.Subjects = nil
 }
 if err := Convert_rbac_RoleRef_To_v1alpha1_RoleRef(&in.RoleRef, &out.RoleRef, s); err != nil {
  return err
 }
 return nil
}
func Convert_rbac_RoleBinding_To_v1alpha1_RoleBinding(in *rbac.RoleBinding, out *v1alpha1.RoleBinding, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_rbac_RoleBinding_To_v1alpha1_RoleBinding(in, out, s)
}
func autoConvert_v1alpha1_RoleBindingList_To_rbac_RoleBindingList(in *v1alpha1.RoleBindingList, out *rbac.RoleBindingList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]rbac.RoleBinding, len(*in))
  for i := range *in {
   if err := Convert_v1alpha1_RoleBinding_To_rbac_RoleBinding(&(*in)[i], &(*out)[i], s); err != nil {
    return err
   }
  }
 } else {
  out.Items = nil
 }
 return nil
}
func Convert_v1alpha1_RoleBindingList_To_rbac_RoleBindingList(in *v1alpha1.RoleBindingList, out *rbac.RoleBindingList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1alpha1_RoleBindingList_To_rbac_RoleBindingList(in, out, s)
}
func autoConvert_rbac_RoleBindingList_To_v1alpha1_RoleBindingList(in *rbac.RoleBindingList, out *v1alpha1.RoleBindingList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]v1alpha1.RoleBinding, len(*in))
  for i := range *in {
   if err := Convert_rbac_RoleBinding_To_v1alpha1_RoleBinding(&(*in)[i], &(*out)[i], s); err != nil {
    return err
   }
  }
 } else {
  out.Items = nil
 }
 return nil
}
func Convert_rbac_RoleBindingList_To_v1alpha1_RoleBindingList(in *rbac.RoleBindingList, out *v1alpha1.RoleBindingList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_rbac_RoleBindingList_To_v1alpha1_RoleBindingList(in, out, s)
}
func autoConvert_v1alpha1_RoleList_To_rbac_RoleList(in *v1alpha1.RoleList, out *rbac.RoleList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 out.Items = *(*[]rbac.Role)(unsafe.Pointer(&in.Items))
 return nil
}
func Convert_v1alpha1_RoleList_To_rbac_RoleList(in *v1alpha1.RoleList, out *rbac.RoleList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1alpha1_RoleList_To_rbac_RoleList(in, out, s)
}
func autoConvert_rbac_RoleList_To_v1alpha1_RoleList(in *rbac.RoleList, out *v1alpha1.RoleList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ListMeta = in.ListMeta
 out.Items = *(*[]v1alpha1.Role)(unsafe.Pointer(&in.Items))
 return nil
}
func Convert_rbac_RoleList_To_v1alpha1_RoleList(in *rbac.RoleList, out *v1alpha1.RoleList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_rbac_RoleList_To_v1alpha1_RoleList(in, out, s)
}
func autoConvert_v1alpha1_RoleRef_To_rbac_RoleRef(in *v1alpha1.RoleRef, out *rbac.RoleRef, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.APIGroup = in.APIGroup
 out.Kind = in.Kind
 out.Name = in.Name
 return nil
}
func Convert_v1alpha1_RoleRef_To_rbac_RoleRef(in *v1alpha1.RoleRef, out *rbac.RoleRef, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1alpha1_RoleRef_To_rbac_RoleRef(in, out, s)
}
func autoConvert_rbac_RoleRef_To_v1alpha1_RoleRef(in *rbac.RoleRef, out *v1alpha1.RoleRef, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.APIGroup = in.APIGroup
 out.Kind = in.Kind
 out.Name = in.Name
 return nil
}
func Convert_rbac_RoleRef_To_v1alpha1_RoleRef(in *rbac.RoleRef, out *v1alpha1.RoleRef, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_rbac_RoleRef_To_v1alpha1_RoleRef(in, out, s)
}
func autoConvert_v1alpha1_Subject_To_rbac_Subject(in *v1alpha1.Subject, out *rbac.Subject, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Kind = in.Kind
 out.Name = in.Name
 out.Namespace = in.Namespace
 return nil
}
func autoConvert_rbac_Subject_To_v1alpha1_Subject(in *rbac.Subject, out *v1alpha1.Subject, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Kind = in.Kind
 out.Name = in.Name
 out.Namespace = in.Namespace
 return nil
}
