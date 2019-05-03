package config

import (
 runtime "k8s.io/apimachinery/pkg/runtime"
)

func (in *KubeSchedulerConfiguration) DeepCopyInto(out *KubeSchedulerConfiguration) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 in.AlgorithmSource.DeepCopyInto(&out.AlgorithmSource)
 out.LeaderElection = in.LeaderElection
 out.ClientConnection = in.ClientConnection
 out.DebuggingConfiguration = in.DebuggingConfiguration
 if in.BindTimeoutSeconds != nil {
  in, out := &in.BindTimeoutSeconds, &out.BindTimeoutSeconds
  *out = new(int64)
  **out = **in
 }
 return
}
func (in *KubeSchedulerConfiguration) DeepCopy() *KubeSchedulerConfiguration {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(KubeSchedulerConfiguration)
 in.DeepCopyInto(out)
 return out
}
func (in *KubeSchedulerConfiguration) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *KubeSchedulerLeaderElectionConfiguration) DeepCopyInto(out *KubeSchedulerLeaderElectionConfiguration) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.LeaderElectionConfiguration = in.LeaderElectionConfiguration
 return
}
func (in *KubeSchedulerLeaderElectionConfiguration) DeepCopy() *KubeSchedulerLeaderElectionConfiguration {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(KubeSchedulerLeaderElectionConfiguration)
 in.DeepCopyInto(out)
 return out
}
func (in *SchedulerAlgorithmSource) DeepCopyInto(out *SchedulerAlgorithmSource) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Policy != nil {
  in, out := &in.Policy, &out.Policy
  *out = new(SchedulerPolicySource)
  (*in).DeepCopyInto(*out)
 }
 if in.Provider != nil {
  in, out := &in.Provider, &out.Provider
  *out = new(string)
  **out = **in
 }
 return
}
func (in *SchedulerAlgorithmSource) DeepCopy() *SchedulerAlgorithmSource {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(SchedulerAlgorithmSource)
 in.DeepCopyInto(out)
 return out
}
func (in *SchedulerPolicyConfigMapSource) DeepCopyInto(out *SchedulerPolicyConfigMapSource) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 return
}
func (in *SchedulerPolicyConfigMapSource) DeepCopy() *SchedulerPolicyConfigMapSource {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(SchedulerPolicyConfigMapSource)
 in.DeepCopyInto(out)
 return out
}
func (in *SchedulerPolicyFileSource) DeepCopyInto(out *SchedulerPolicyFileSource) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 return
}
func (in *SchedulerPolicyFileSource) DeepCopy() *SchedulerPolicyFileSource {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(SchedulerPolicyFileSource)
 in.DeepCopyInto(out)
 return out
}
func (in *SchedulerPolicySource) DeepCopyInto(out *SchedulerPolicySource) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.File != nil {
  in, out := &in.File, &out.File
  *out = new(SchedulerPolicyFileSource)
  **out = **in
 }
 if in.ConfigMap != nil {
  in, out := &in.ConfigMap, &out.ConfigMap
  *out = new(SchedulerPolicyConfigMapSource)
  **out = **in
 }
 return
}
func (in *SchedulerPolicySource) DeepCopy() *SchedulerPolicySource {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(SchedulerPolicySource)
 in.DeepCopyInto(out)
 return out
}
