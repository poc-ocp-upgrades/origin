package config

import (
 v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 runtime "k8s.io/apimachinery/pkg/runtime"
)

func (in ConfigurationMap) DeepCopyInto(out *ConfigurationMap) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 {
  in := &in
  *out = make(ConfigurationMap, len(*in))
  for key, val := range *in {
   (*out)[key] = val
  }
  return
 }
}
func (in ConfigurationMap) DeepCopy() ConfigurationMap {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(ConfigurationMap)
 in.DeepCopyInto(out)
 return *out
}
func (in *KubeProxyConfiguration) DeepCopyInto(out *KubeProxyConfiguration) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.TypeMeta = in.TypeMeta
 if in.FeatureGates != nil {
  in, out := &in.FeatureGates, &out.FeatureGates
  *out = make(map[string]bool, len(*in))
  for key, val := range *in {
   (*out)[key] = val
  }
 }
 out.ClientConnection = in.ClientConnection
 in.IPTables.DeepCopyInto(&out.IPTables)
 in.IPVS.DeepCopyInto(&out.IPVS)
 if in.OOMScoreAdj != nil {
  in, out := &in.OOMScoreAdj, &out.OOMScoreAdj
  *out = new(int32)
  **out = **in
 }
 out.UDPIdleTimeout = in.UDPIdleTimeout
 in.Conntrack.DeepCopyInto(&out.Conntrack)
 out.ConfigSyncPeriod = in.ConfigSyncPeriod
 if in.NodePortAddresses != nil {
  in, out := &in.NodePortAddresses, &out.NodePortAddresses
  *out = make([]string, len(*in))
  copy(*out, *in)
 }
 return
}
func (in *KubeProxyConfiguration) DeepCopy() *KubeProxyConfiguration {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(KubeProxyConfiguration)
 in.DeepCopyInto(out)
 return out
}
func (in *KubeProxyConfiguration) DeepCopyObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c := in.DeepCopy(); c != nil {
  return c
 }
 return nil
}
func (in *KubeProxyConntrackConfiguration) DeepCopyInto(out *KubeProxyConntrackConfiguration) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.Max != nil {
  in, out := &in.Max, &out.Max
  *out = new(int32)
  **out = **in
 }
 if in.MaxPerCore != nil {
  in, out := &in.MaxPerCore, &out.MaxPerCore
  *out = new(int32)
  **out = **in
 }
 if in.Min != nil {
  in, out := &in.Min, &out.Min
  *out = new(int32)
  **out = **in
 }
 if in.TCPEstablishedTimeout != nil {
  in, out := &in.TCPEstablishedTimeout, &out.TCPEstablishedTimeout
  *out = new(v1.Duration)
  **out = **in
 }
 if in.TCPCloseWaitTimeout != nil {
  in, out := &in.TCPCloseWaitTimeout, &out.TCPCloseWaitTimeout
  *out = new(v1.Duration)
  **out = **in
 }
 return
}
func (in *KubeProxyConntrackConfiguration) DeepCopy() *KubeProxyConntrackConfiguration {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(KubeProxyConntrackConfiguration)
 in.DeepCopyInto(out)
 return out
}
func (in *KubeProxyIPTablesConfiguration) DeepCopyInto(out *KubeProxyIPTablesConfiguration) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 if in.MasqueradeBit != nil {
  in, out := &in.MasqueradeBit, &out.MasqueradeBit
  *out = new(int32)
  **out = **in
 }
 out.SyncPeriod = in.SyncPeriod
 out.MinSyncPeriod = in.MinSyncPeriod
 return
}
func (in *KubeProxyIPTablesConfiguration) DeepCopy() *KubeProxyIPTablesConfiguration {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(KubeProxyIPTablesConfiguration)
 in.DeepCopyInto(out)
 return out
}
func (in *KubeProxyIPVSConfiguration) DeepCopyInto(out *KubeProxyIPVSConfiguration) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 *out = *in
 out.SyncPeriod = in.SyncPeriod
 out.MinSyncPeriod = in.MinSyncPeriod
 if in.ExcludeCIDRs != nil {
  in, out := &in.ExcludeCIDRs, &out.ExcludeCIDRs
  *out = make([]string, len(*in))
  copy(*out, *in)
 }
 return
}
func (in *KubeProxyIPVSConfiguration) DeepCopy() *KubeProxyIPVSConfiguration {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in == nil {
  return nil
 }
 out := new(KubeProxyIPVSConfiguration)
 in.DeepCopyInto(out)
 return out
}
