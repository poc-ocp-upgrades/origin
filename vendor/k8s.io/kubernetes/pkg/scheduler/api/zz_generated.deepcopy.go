package api

import (
 v1 "k8s.io/api/core/v1"
 runtime "k8s.io/apimachinery/pkg/runtime"
 rest "k8s.io/client-go/rest"
)

func (in *ExtenderArgs) DeepCopyInto(out *ExtenderArgs) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Pod != nil {
  in, out := &in.Pod, &out.Pod
  *out = new(v1.Pod)
  (*in).DeepCopyInto(*out)
 }
 if in.Nodes != nil {
  in, out := &in.Nodes, &out.Nodes
  *out = new(v1.NodeList)
  (*in).DeepCopyInto(*out)
 }
 if in.NodeNames != nil {
  in, out := &in.NodeNames, &out.NodeNames
  *out = new([]string)
  if **in != nil {
   in, out := *in, *out
   *out = make([]string, len(*in))
   copy(*out, *in)
  }
 }
 return
}
func (in *ExtenderArgs) DeepCopy() *ExtenderArgs {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ExtenderArgs)
 in.DeepCopyInto(out)
 return out
}
func (in *ExtenderBindingArgs) DeepCopyInto(out *ExtenderBindingArgs) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 return
}
func (in *ExtenderBindingArgs) DeepCopy() *ExtenderBindingArgs {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ExtenderBindingArgs)
 in.DeepCopyInto(out)
 return out
}
func (in *ExtenderBindingResult) DeepCopyInto(out *ExtenderBindingResult) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 return
}
func (in *ExtenderBindingResult) DeepCopy() *ExtenderBindingResult {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ExtenderBindingResult)
 in.DeepCopyInto(out)
 return out
}
func (in *ExtenderConfig) DeepCopyInto(out *ExtenderConfig) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.TLSConfig != nil {
  in, out := &in.TLSConfig, &out.TLSConfig
  *out = new(rest.TLSClientConfig)
  (*in).DeepCopyInto(*out)
 }
 if in.ManagedResources != nil {
  in, out := &in.ManagedResources, &out.ManagedResources
  *out = make([]ExtenderManagedResource, len(*in))
  copy(*out, *in)
 }
 return
}
func (in *ExtenderConfig) DeepCopy() *ExtenderConfig {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ExtenderConfig)
 in.DeepCopyInto(out)
 return out
}
func (in *ExtenderFilterResult) DeepCopyInto(out *ExtenderFilterResult) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Nodes != nil {
  in, out := &in.Nodes, &out.Nodes
  *out = new(v1.NodeList)
  (*in).DeepCopyInto(*out)
 }
 if in.NodeNames != nil {
  in, out := &in.NodeNames, &out.NodeNames
  *out = new([]string)
  if **in != nil {
   in, out := *in, *out
   *out = make([]string, len(*in))
   copy(*out, *in)
  }
 }
 if in.FailedNodes != nil {
  in, out := &in.FailedNodes, &out.FailedNodes
  *out = make(FailedNodesMap, len(*in))
  for key, val := range *in {
   (*out)[key] = val
  }
 }
 return
}
func (in *ExtenderFilterResult) DeepCopy() *ExtenderFilterResult {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ExtenderFilterResult)
 in.DeepCopyInto(out)
 return out
}
func (in *ExtenderManagedResource) DeepCopyInto(out *ExtenderManagedResource) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 return
}
func (in *ExtenderManagedResource) DeepCopy() *ExtenderManagedResource {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ExtenderManagedResource)
 in.DeepCopyInto(out)
 return out
}
func (in *ExtenderPreemptionArgs) DeepCopyInto(out *ExtenderPreemptionArgs) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Pod != nil {
  in, out := &in.Pod, &out.Pod
  *out = new(v1.Pod)
  (*in).DeepCopyInto(*out)
 }
 if in.NodeNameToVictims != nil {
  in, out := &in.NodeNameToVictims, &out.NodeNameToVictims
  *out = make(map[string]*Victims, len(*in))
  for key, val := range *in {
   var outVal *Victims
   if val == nil {
    (*out)[key] = nil
   } else {
    in, out := &val, &outVal
    *out = new(Victims)
    (*in).DeepCopyInto(*out)
   }
   (*out)[key] = outVal
  }
 }
 if in.NodeNameToMetaVictims != nil {
  in, out := &in.NodeNameToMetaVictims, &out.NodeNameToMetaVictims
  *out = make(map[string]*MetaVictims, len(*in))
  for key, val := range *in {
   var outVal *MetaVictims
   if val == nil {
    (*out)[key] = nil
   } else {
    in, out := &val, &outVal
    *out = new(MetaVictims)
    (*in).DeepCopyInto(*out)
   }
   (*out)[key] = outVal
  }
 }
 return
}
func (in *ExtenderPreemptionArgs) DeepCopy() *ExtenderPreemptionArgs {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ExtenderPreemptionArgs)
 in.DeepCopyInto(out)
 return out
}
func (in *ExtenderPreemptionResult) DeepCopyInto(out *ExtenderPreemptionResult) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.NodeNameToMetaVictims != nil {
  in, out := &in.NodeNameToMetaVictims, &out.NodeNameToMetaVictims
  *out = make(map[string]*MetaVictims, len(*in))
  for key, val := range *in {
   var outVal *MetaVictims
   if val == nil {
    (*out)[key] = nil
   } else {
    in, out := &val, &outVal
    *out = new(MetaVictims)
    (*in).DeepCopyInto(*out)
   }
   (*out)[key] = outVal
  }
 }
 return
}
func (in *ExtenderPreemptionResult) DeepCopy() *ExtenderPreemptionResult {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ExtenderPreemptionResult)
 in.DeepCopyInto(out)
 return out
}
func (in FailedNodesMap) DeepCopyInto(out *FailedNodesMap) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 {
  in := &in
  *out = make(FailedNodesMap, len(*in))
  for key, val := range *in {
   (*out)[key] = val
  }
  return
 }
}
func (in FailedNodesMap) DeepCopy() FailedNodesMap {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(FailedNodesMap)
 in.DeepCopyInto(out)
 return *out
}
func (in *HostPriority) DeepCopyInto(out *HostPriority) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 return
}
func (in *HostPriority) DeepCopy() *HostPriority {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(HostPriority)
 in.DeepCopyInto(out)
 return out
}
func (in HostPriorityList) DeepCopyInto(out *HostPriorityList) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 {
  in := &in
  *out = make(HostPriorityList, len(*in))
  copy(*out, *in)
  return
 }
}
func (in HostPriorityList) DeepCopy() HostPriorityList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(HostPriorityList)
 in.DeepCopyInto(out)
 return *out
}
func (in *LabelPreference) DeepCopyInto(out *LabelPreference) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 return
}
func (in *LabelPreference) DeepCopy() *LabelPreference {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(LabelPreference)
 in.DeepCopyInto(out)
 return out
}
func (in *LabelsPresence) DeepCopyInto(out *LabelsPresence) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Labels != nil {
  in, out := &in.Labels, &out.Labels
  *out = make([]string, len(*in))
  copy(*out, *in)
 }
 return
}
func (in *LabelsPresence) DeepCopy() *LabelsPresence {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(LabelsPresence)
 in.DeepCopyInto(out)
 return out
}
func (in *MetaPod) DeepCopyInto(out *MetaPod) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 return
}
func (in *MetaPod) DeepCopy() *MetaPod {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(MetaPod)
 in.DeepCopyInto(out)
 return out
}
func (in *MetaVictims) DeepCopyInto(out *MetaVictims) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Pods != nil {
  in, out := &in.Pods, &out.Pods
  *out = make([]*MetaPod, len(*in))
  for i := range *in {
   if (*in)[i] != nil {
    in, out := &(*in)[i], &(*out)[i]
    *out = new(MetaPod)
    **out = **in
   }
  }
 }
 return
}
func (in *MetaVictims) DeepCopy() *MetaVictims {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(MetaVictims)
 in.DeepCopyInto(out)
 return out
}
func (in *Policy) DeepCopyInto(out *Policy) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 if in.Predicates != nil {
  in, out := &in.Predicates, &out.Predicates
  *out = make([]PredicatePolicy, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 if in.Priorities != nil {
  in, out := &in.Priorities, &out.Priorities
  *out = make([]PriorityPolicy, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 if in.ExtenderConfigs != nil {
  in, out := &in.ExtenderConfigs, &out.ExtenderConfigs
  *out = make([]ExtenderConfig, len(*in))
  for i := range *in {
   (*in)[i].DeepCopyInto(&(*out)[i])
  }
 }
 return
}
func (in *Policy) DeepCopy() *Policy {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(Policy)
 in.DeepCopyInto(out)
 return out
}
func (in *Policy) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *PredicateArgument) DeepCopyInto(out *PredicateArgument) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.ServiceAffinity != nil {
  in, out := &in.ServiceAffinity, &out.ServiceAffinity
  *out = new(ServiceAffinity)
  (*in).DeepCopyInto(*out)
 }
 if in.LabelsPresence != nil {
  in, out := &in.LabelsPresence, &out.LabelsPresence
  *out = new(LabelsPresence)
  (*in).DeepCopyInto(*out)
 }
 return
}
func (in *PredicateArgument) DeepCopy() *PredicateArgument {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(PredicateArgument)
 in.DeepCopyInto(out)
 return out
}
func (in *PredicatePolicy) DeepCopyInto(out *PredicatePolicy) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Argument != nil {
  in, out := &in.Argument, &out.Argument
  *out = new(PredicateArgument)
  (*in).DeepCopyInto(*out)
 }
 return
}
func (in *PredicatePolicy) DeepCopy() *PredicatePolicy {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(PredicatePolicy)
 in.DeepCopyInto(out)
 return out
}
func (in *PriorityArgument) DeepCopyInto(out *PriorityArgument) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.ServiceAntiAffinity != nil {
  in, out := &in.ServiceAntiAffinity, &out.ServiceAntiAffinity
  *out = new(ServiceAntiAffinity)
  **out = **in
 }
 if in.LabelPreference != nil {
  in, out := &in.LabelPreference, &out.LabelPreference
  *out = new(LabelPreference)
  **out = **in
 }
 if in.RequestedToCapacityRatioArguments != nil {
  in, out := &in.RequestedToCapacityRatioArguments, &out.RequestedToCapacityRatioArguments
  *out = new(RequestedToCapacityRatioArguments)
  (*in).DeepCopyInto(*out)
 }
 return
}
func (in *PriorityArgument) DeepCopy() *PriorityArgument {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(PriorityArgument)
 in.DeepCopyInto(out)
 return out
}
func (in *PriorityPolicy) DeepCopyInto(out *PriorityPolicy) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Argument != nil {
  in, out := &in.Argument, &out.Argument
  *out = new(PriorityArgument)
  (*in).DeepCopyInto(*out)
 }
 return
}
func (in *PriorityPolicy) DeepCopy() *PriorityPolicy {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(PriorityPolicy)
 in.DeepCopyInto(out)
 return out
}
func (in *RequestedToCapacityRatioArguments) DeepCopyInto(out *RequestedToCapacityRatioArguments) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.UtilizationShape != nil {
  in, out := &in.UtilizationShape, &out.UtilizationShape
  *out = make([]UtilizationShapePoint, len(*in))
  copy(*out, *in)
 }
 return
}
func (in *RequestedToCapacityRatioArguments) DeepCopy() *RequestedToCapacityRatioArguments {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(RequestedToCapacityRatioArguments)
 in.DeepCopyInto(out)
 return out
}
func (in *ServiceAffinity) DeepCopyInto(out *ServiceAffinity) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Labels != nil {
  in, out := &in.Labels, &out.Labels
  *out = make([]string, len(*in))
  copy(*out, *in)
 }
 return
}
func (in *ServiceAffinity) DeepCopy() *ServiceAffinity {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ServiceAffinity)
 in.DeepCopyInto(out)
 return out
}
func (in *ServiceAntiAffinity) DeepCopyInto(out *ServiceAntiAffinity) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 return
}
func (in *ServiceAntiAffinity) DeepCopy() *ServiceAntiAffinity {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ServiceAntiAffinity)
 in.DeepCopyInto(out)
 return out
}
func (in *UtilizationShapePoint) DeepCopyInto(out *UtilizationShapePoint) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 return
}
func (in *UtilizationShapePoint) DeepCopy() *UtilizationShapePoint {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(UtilizationShapePoint)
 in.DeepCopyInto(out)
 return out
}
func (in *Victims) DeepCopyInto(out *Victims) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Pods != nil {
  in, out := &in.Pods, &out.Pods
  *out = make([]*v1.Pod, len(*in))
  for i := range *in {
   if (*in)[i] != nil {
    in, out := &(*in)[i], &(*out)[i]
    *out = new(v1.Pod)
    (*in).DeepCopyInto(*out)
   }
  }
 }
 return
}
func (in *Victims) DeepCopy() *Victims {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(Victims)
 in.DeepCopyInto(out)
 return out
}
