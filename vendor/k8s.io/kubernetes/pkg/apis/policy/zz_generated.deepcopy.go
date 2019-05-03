package policy

import (
 v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 runtime "k8s.io/apimachinery/pkg/runtime"
 intstr "k8s.io/apimachinery/pkg/util/intstr"
 core "k8s.io/kubernetes/pkg/apis/core"
)

func (in *AllowedFlexVolume) DeepCopyInto(out *AllowedFlexVolume) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 return
}
func (in *AllowedFlexVolume) DeepCopy() *AllowedFlexVolume {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(AllowedFlexVolume)
 in.DeepCopyInto(out)
 return out
}
func (in *AllowedHostPath) DeepCopyInto(out *AllowedHostPath) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 return
}
func (in *AllowedHostPath) DeepCopy() *AllowedHostPath {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(AllowedHostPath)
 in.DeepCopyInto(out)
 return out
}
func (in *Eviction) DeepCopyInto(out *Eviction) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
 if in.DeleteOptions != nil {
  in, out := &in.DeleteOptions, &out.DeleteOptions
  *out = new(v1.DeleteOptions)
  (*in).DeepCopyInto(*out)
 }
 return
}
func (in *Eviction) DeepCopy() *Eviction {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(Eviction)
 in.DeepCopyInto(out)
 return out
}
func (in *Eviction) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *FSGroupStrategyOptions) DeepCopyInto(out *FSGroupStrategyOptions) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Ranges != nil {
  in, out := &in.Ranges, &out.Ranges
  *out = make([]IDRange, len(*in))
  copy(*out, *in)
 }
 return
}
func (in *FSGroupStrategyOptions) DeepCopy() *FSGroupStrategyOptions {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(FSGroupStrategyOptions)
 in.DeepCopyInto(out)
 return out
}
func (in *HostPortRange) DeepCopyInto(out *HostPortRange) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 return
}
func (in *HostPortRange) DeepCopy() *HostPortRange {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(HostPortRange)
 in.DeepCopyInto(out)
 return out
}
func (in *IDRange) DeepCopyInto(out *IDRange) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 return
}
func (in *IDRange) DeepCopy() *IDRange {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(IDRange)
 in.DeepCopyInto(out)
 return out
}
func (in *PodDisruptionBudget) DeepCopyInto(out *PodDisruptionBudget) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
 in.Spec.DeepCopyInto(&out.Spec)
 in.Status.DeepCopyInto(&out.Status)
 return
}
func (in *PodDisruptionBudget) DeepCopy() *PodDisruptionBudget {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(PodDisruptionBudget)
 in.DeepCopyInto(out)
 return out
}
func (in *PodDisruptionBudget) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *PodDisruptionBudgetList) DeepCopyInto(out *PodDisruptionBudgetList) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]PodDisruptionBudget, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 return
}
func (in *PodDisruptionBudgetList) DeepCopy() *PodDisruptionBudgetList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(PodDisruptionBudgetList)
 in.DeepCopyInto(out)
 return out
}
func (in *PodDisruptionBudgetList) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *PodDisruptionBudgetSpec) DeepCopyInto(out *PodDisruptionBudgetSpec) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.MinAvailable != nil {
  in, out := &in.MinAvailable, &out.MinAvailable
  *out = new(intstr.IntOrString)
  **out = **in
 }
 if in.Selector != nil {
  in, out := &in.Selector, &out.Selector
  *out = new(v1.LabelSelector)
  (*in).DeepCopyInto(*out)
 }
 if in.MaxUnavailable != nil {
  in, out := &in.MaxUnavailable, &out.MaxUnavailable
  *out = new(intstr.IntOrString)
  **out = **in
 }
 return
}
func (in *PodDisruptionBudgetSpec) DeepCopy() *PodDisruptionBudgetSpec {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(PodDisruptionBudgetSpec)
 in.DeepCopyInto(out)
 return out
}
func (in *PodDisruptionBudgetStatus) DeepCopyInto(out *PodDisruptionBudgetStatus) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.DisruptedPods != nil {
  in, out := &in.DisruptedPods, &out.DisruptedPods
  *out = make(map[string]v1.Time, len(*in))
  for key, val := range *in {
   (*out)[key] = *val.DeepCopy()
  }
 }
 return
}
func (in *PodDisruptionBudgetStatus) DeepCopy() *PodDisruptionBudgetStatus {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(PodDisruptionBudgetStatus)
 in.DeepCopyInto(out)
 return out
}
func (in *PodSecurityPolicy) DeepCopyInto(out *PodSecurityPolicy) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
 in.Spec.DeepCopyInto(&out.Spec)
 return
}
func (in *PodSecurityPolicy) DeepCopy() *PodSecurityPolicy {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(PodSecurityPolicy)
 in.DeepCopyInto(out)
 return out
}
func (in *PodSecurityPolicy) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *PodSecurityPolicyList) DeepCopyInto(out *PodSecurityPolicyList) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 out.ListMeta = in.ListMeta
 if in.Items != nil {
  in, out := &in.Items, &out.Items
  *out = make([]PodSecurityPolicy, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 return
}
func (in *PodSecurityPolicyList) DeepCopy() *PodSecurityPolicyList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(PodSecurityPolicyList)
 in.DeepCopyInto(out)
 return out
}
func (in *PodSecurityPolicyList) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *PodSecurityPolicySpec) DeepCopyInto(out *PodSecurityPolicySpec) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.DefaultAddCapabilities != nil {
  in, out := &in.DefaultAddCapabilities, &out.DefaultAddCapabilities
  *out = make([]core.Capability, len(*in))
  copy(*out, *in)
 }
 if in.RequiredDropCapabilities != nil {
  in, out := &in.RequiredDropCapabilities, &out.RequiredDropCapabilities
  *out = make([]core.Capability, len(*in))
  copy(*out, *in)
 }
 if in.AllowedCapabilities != nil {
  in, out := &in.AllowedCapabilities, &out.AllowedCapabilities
  *out = make([]core.Capability, len(*in))
  copy(*out, *in)
 }
 if in.Volumes != nil {
  in, out := &in.Volumes, &out.Volumes
  *out = make([]FSType, len(*in))
  copy(*out, *in)
 }
 if in.HostPorts != nil {
  in, out := &in.HostPorts, &out.HostPorts
  *out = make([]HostPortRange, len(*in))
  copy(*out, *in)
 }
 in.SELinux.DeepCopyInto(&out.SELinux)
 in.RunAsUser.DeepCopyInto(&out.RunAsUser)
 if in.RunAsGroup != nil {
  in, out := &in.RunAsGroup, &out.RunAsGroup
  *out = new(RunAsGroupStrategyOptions)
  (*in).DeepCopyInto(*out)
 }
 in.SupplementalGroups.DeepCopyInto(&out.SupplementalGroups)
 in.FSGroup.DeepCopyInto(&out.FSGroup)
 if in.DefaultAllowPrivilegeEscalation != nil {
  in, out := &in.DefaultAllowPrivilegeEscalation, &out.DefaultAllowPrivilegeEscalation
  *out = new(bool)
  **out = **in
 }
 if in.AllowedHostPaths != nil {
  in, out := &in.AllowedHostPaths, &out.AllowedHostPaths
  *out = make([]AllowedHostPath, len(*in))
  copy(*out, *in)
 }
 if in.AllowedFlexVolumes != nil {
  in, out := &in.AllowedFlexVolumes, &out.AllowedFlexVolumes
  *out = make([]AllowedFlexVolume, len(*in))
  copy(*out, *in)
 }
 if in.AllowedUnsafeSysctls != nil {
  in, out := &in.AllowedUnsafeSysctls, &out.AllowedUnsafeSysctls
  *out = make([]string, len(*in))
  copy(*out, *in)
 }
 if in.ForbiddenSysctls != nil {
  in, out := &in.ForbiddenSysctls, &out.ForbiddenSysctls
  *out = make([]string, len(*in))
  copy(*out, *in)
 }
 if in.AllowedProcMountTypes != nil {
  in, out := &in.AllowedProcMountTypes, &out.AllowedProcMountTypes
  *out = make([]core.ProcMountType, len(*in))
  copy(*out, *in)
 }
 return
}
func (in *PodSecurityPolicySpec) DeepCopy() *PodSecurityPolicySpec {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(PodSecurityPolicySpec)
 in.DeepCopyInto(out)
 return out
}
func (in *RunAsGroupStrategyOptions) DeepCopyInto(out *RunAsGroupStrategyOptions) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Ranges != nil {
  in, out := &in.Ranges, &out.Ranges
  *out = make([]IDRange, len(*in))
  copy(*out, *in)
 }
 return
}
func (in *RunAsGroupStrategyOptions) DeepCopy() *RunAsGroupStrategyOptions {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(RunAsGroupStrategyOptions)
 in.DeepCopyInto(out)
 return out
}
func (in *RunAsUserStrategyOptions) DeepCopyInto(out *RunAsUserStrategyOptions) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Ranges != nil {
  in, out := &in.Ranges, &out.Ranges
  *out = make([]IDRange, len(*in))
  copy(*out, *in)
 }
 return
}
func (in *RunAsUserStrategyOptions) DeepCopy() *RunAsUserStrategyOptions {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(RunAsUserStrategyOptions)
 in.DeepCopyInto(out)
 return out
}
func (in *SELinuxStrategyOptions) DeepCopyInto(out *SELinuxStrategyOptions) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.SELinuxOptions != nil {
  in, out := &in.SELinuxOptions, &out.SELinuxOptions
  *out = new(core.SELinuxOptions)
  **out = **in
 }
 return
}
func (in *SELinuxStrategyOptions) DeepCopy() *SELinuxStrategyOptions {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(SELinuxStrategyOptions)
 in.DeepCopyInto(out)
 return out
}
func (in *SupplementalGroupsStrategyOptions) DeepCopyInto(out *SupplementalGroupsStrategyOptions) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Ranges != nil {
  in, out := &in.Ranges, &out.Ranges
  *out = make([]IDRange, len(*in))
  copy(*out, *in)
 }
 return
}
func (in *SupplementalGroupsStrategyOptions) DeepCopy() *SupplementalGroupsStrategyOptions {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(SupplementalGroupsStrategyOptions)
 in.DeepCopyInto(out)
 return out
}
