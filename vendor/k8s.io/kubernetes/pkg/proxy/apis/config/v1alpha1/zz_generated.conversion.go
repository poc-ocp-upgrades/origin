package v1alpha1

import (
 unsafe "unsafe"
 configv1alpha1 "k8s.io/apimachinery/pkg/apis/config/v1alpha1"
 v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 conversion "k8s.io/apimachinery/pkg/conversion"
 runtime "k8s.io/apimachinery/pkg/runtime"
 v1alpha1 "k8s.io/kube-proxy/config/v1alpha1"
 config "k8s.io/kubernetes/pkg/proxy/apis/config"
)

func init() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 localSchemeBuilder.Register(RegisterConversions)
}
func RegisterConversions(s *runtime.Scheme) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := s.AddGeneratedConversionFunc((*v1alpha1.KubeProxyConfiguration)(nil), (*config.KubeProxyConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1alpha1_KubeProxyConfiguration_To_config_KubeProxyConfiguration(a.(*v1alpha1.KubeProxyConfiguration), b.(*config.KubeProxyConfiguration), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*config.KubeProxyConfiguration)(nil), (*v1alpha1.KubeProxyConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_config_KubeProxyConfiguration_To_v1alpha1_KubeProxyConfiguration(a.(*config.KubeProxyConfiguration), b.(*v1alpha1.KubeProxyConfiguration), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1alpha1.KubeProxyConntrackConfiguration)(nil), (*config.KubeProxyConntrackConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1alpha1_KubeProxyConntrackConfiguration_To_config_KubeProxyConntrackConfiguration(a.(*v1alpha1.KubeProxyConntrackConfiguration), b.(*config.KubeProxyConntrackConfiguration), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*config.KubeProxyConntrackConfiguration)(nil), (*v1alpha1.KubeProxyConntrackConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_config_KubeProxyConntrackConfiguration_To_v1alpha1_KubeProxyConntrackConfiguration(a.(*config.KubeProxyConntrackConfiguration), b.(*v1alpha1.KubeProxyConntrackConfiguration), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1alpha1.KubeProxyIPTablesConfiguration)(nil), (*config.KubeProxyIPTablesConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1alpha1_KubeProxyIPTablesConfiguration_To_config_KubeProxyIPTablesConfiguration(a.(*v1alpha1.KubeProxyIPTablesConfiguration), b.(*config.KubeProxyIPTablesConfiguration), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*config.KubeProxyIPTablesConfiguration)(nil), (*v1alpha1.KubeProxyIPTablesConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_config_KubeProxyIPTablesConfiguration_To_v1alpha1_KubeProxyIPTablesConfiguration(a.(*config.KubeProxyIPTablesConfiguration), b.(*v1alpha1.KubeProxyIPTablesConfiguration), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1alpha1.KubeProxyIPVSConfiguration)(nil), (*config.KubeProxyIPVSConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1alpha1_KubeProxyIPVSConfiguration_To_config_KubeProxyIPVSConfiguration(a.(*v1alpha1.KubeProxyIPVSConfiguration), b.(*config.KubeProxyIPVSConfiguration), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*config.KubeProxyIPVSConfiguration)(nil), (*v1alpha1.KubeProxyIPVSConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_config_KubeProxyIPVSConfiguration_To_v1alpha1_KubeProxyIPVSConfiguration(a.(*config.KubeProxyIPVSConfiguration), b.(*v1alpha1.KubeProxyIPVSConfiguration), scope)
 }); err != nil {
  return err
 }
 return nil
}
func autoConvert_v1alpha1_KubeProxyConfiguration_To_config_KubeProxyConfiguration(in *v1alpha1.KubeProxyConfiguration, out *config.KubeProxyConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.FeatureGates = *(*map[string]bool)(unsafe.Pointer(&in.FeatureGates))
 out.BindAddress = in.BindAddress
 out.HealthzBindAddress = in.HealthzBindAddress
 out.MetricsBindAddress = in.MetricsBindAddress
 out.EnableProfiling = in.EnableProfiling
 out.ClusterCIDR = in.ClusterCIDR
 out.HostnameOverride = in.HostnameOverride
 if err := configv1alpha1.Convert_v1alpha1_ClientConnectionConfiguration_To_config_ClientConnectionConfiguration(&in.ClientConnection, &out.ClientConnection, s); err != nil {
  return err
 }
 if err := Convert_v1alpha1_KubeProxyIPTablesConfiguration_To_config_KubeProxyIPTablesConfiguration(&in.IPTables, &out.IPTables, s); err != nil {
  return err
 }
 if err := Convert_v1alpha1_KubeProxyIPVSConfiguration_To_config_KubeProxyIPVSConfiguration(&in.IPVS, &out.IPVS, s); err != nil {
  return err
 }
 out.OOMScoreAdj = (*int32)(unsafe.Pointer(in.OOMScoreAdj))
 out.Mode = config.ProxyMode(in.Mode)
 out.PortRange = in.PortRange
 out.ResourceContainer = in.ResourceContainer
 out.UDPIdleTimeout = in.UDPIdleTimeout
 if err := Convert_v1alpha1_KubeProxyConntrackConfiguration_To_config_KubeProxyConntrackConfiguration(&in.Conntrack, &out.Conntrack, s); err != nil {
  return err
 }
 out.ConfigSyncPeriod = in.ConfigSyncPeriod
 out.NodePortAddresses = *(*[]string)(unsafe.Pointer(&in.NodePortAddresses))
 return nil
}
func Convert_v1alpha1_KubeProxyConfiguration_To_config_KubeProxyConfiguration(in *v1alpha1.KubeProxyConfiguration, out *config.KubeProxyConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1alpha1_KubeProxyConfiguration_To_config_KubeProxyConfiguration(in, out, s)
}
func autoConvert_config_KubeProxyConfiguration_To_v1alpha1_KubeProxyConfiguration(in *config.KubeProxyConfiguration, out *v1alpha1.KubeProxyConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.FeatureGates = *(*map[string]bool)(unsafe.Pointer(&in.FeatureGates))
 out.BindAddress = in.BindAddress
 out.HealthzBindAddress = in.HealthzBindAddress
 out.MetricsBindAddress = in.MetricsBindAddress
 out.EnableProfiling = in.EnableProfiling
 out.ClusterCIDR = in.ClusterCIDR
 out.HostnameOverride = in.HostnameOverride
 if err := configv1alpha1.Convert_config_ClientConnectionConfiguration_To_v1alpha1_ClientConnectionConfiguration(&in.ClientConnection, &out.ClientConnection, s); err != nil {
  return err
 }
 if err := Convert_config_KubeProxyIPTablesConfiguration_To_v1alpha1_KubeProxyIPTablesConfiguration(&in.IPTables, &out.IPTables, s); err != nil {
  return err
 }
 if err := Convert_config_KubeProxyIPVSConfiguration_To_v1alpha1_KubeProxyIPVSConfiguration(&in.IPVS, &out.IPVS, s); err != nil {
  return err
 }
 out.OOMScoreAdj = (*int32)(unsafe.Pointer(in.OOMScoreAdj))
 out.Mode = v1alpha1.ProxyMode(in.Mode)
 out.PortRange = in.PortRange
 out.ResourceContainer = in.ResourceContainer
 out.UDPIdleTimeout = in.UDPIdleTimeout
 if err := Convert_config_KubeProxyConntrackConfiguration_To_v1alpha1_KubeProxyConntrackConfiguration(&in.Conntrack, &out.Conntrack, s); err != nil {
  return err
 }
 out.ConfigSyncPeriod = in.ConfigSyncPeriod
 out.NodePortAddresses = *(*[]string)(unsafe.Pointer(&in.NodePortAddresses))
 return nil
}
func Convert_config_KubeProxyConfiguration_To_v1alpha1_KubeProxyConfiguration(in *config.KubeProxyConfiguration, out *v1alpha1.KubeProxyConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_config_KubeProxyConfiguration_To_v1alpha1_KubeProxyConfiguration(in, out, s)
}
func autoConvert_v1alpha1_KubeProxyConntrackConfiguration_To_config_KubeProxyConntrackConfiguration(in *v1alpha1.KubeProxyConntrackConfiguration, out *config.KubeProxyConntrackConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Max = (*int32)(unsafe.Pointer(in.Max))
 out.MaxPerCore = (*int32)(unsafe.Pointer(in.MaxPerCore))
 out.Min = (*int32)(unsafe.Pointer(in.Min))
 out.TCPEstablishedTimeout = (*v1.Duration)(unsafe.Pointer(in.TCPEstablishedTimeout))
 out.TCPCloseWaitTimeout = (*v1.Duration)(unsafe.Pointer(in.TCPCloseWaitTimeout))
 return nil
}
func Convert_v1alpha1_KubeProxyConntrackConfiguration_To_config_KubeProxyConntrackConfiguration(in *v1alpha1.KubeProxyConntrackConfiguration, out *config.KubeProxyConntrackConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1alpha1_KubeProxyConntrackConfiguration_To_config_KubeProxyConntrackConfiguration(in, out, s)
}
func autoConvert_config_KubeProxyConntrackConfiguration_To_v1alpha1_KubeProxyConntrackConfiguration(in *config.KubeProxyConntrackConfiguration, out *v1alpha1.KubeProxyConntrackConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Max = (*int32)(unsafe.Pointer(in.Max))
 out.MaxPerCore = (*int32)(unsafe.Pointer(in.MaxPerCore))
 out.Min = (*int32)(unsafe.Pointer(in.Min))
 out.TCPEstablishedTimeout = (*v1.Duration)(unsafe.Pointer(in.TCPEstablishedTimeout))
 out.TCPCloseWaitTimeout = (*v1.Duration)(unsafe.Pointer(in.TCPCloseWaitTimeout))
 return nil
}
func Convert_config_KubeProxyConntrackConfiguration_To_v1alpha1_KubeProxyConntrackConfiguration(in *config.KubeProxyConntrackConfiguration, out *v1alpha1.KubeProxyConntrackConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_config_KubeProxyConntrackConfiguration_To_v1alpha1_KubeProxyConntrackConfiguration(in, out, s)
}
func autoConvert_v1alpha1_KubeProxyIPTablesConfiguration_To_config_KubeProxyIPTablesConfiguration(in *v1alpha1.KubeProxyIPTablesConfiguration, out *config.KubeProxyIPTablesConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.MasqueradeBit = (*int32)(unsafe.Pointer(in.MasqueradeBit))
 out.MasqueradeAll = in.MasqueradeAll
 out.SyncPeriod = in.SyncPeriod
 out.MinSyncPeriod = in.MinSyncPeriod
 return nil
}
func Convert_v1alpha1_KubeProxyIPTablesConfiguration_To_config_KubeProxyIPTablesConfiguration(in *v1alpha1.KubeProxyIPTablesConfiguration, out *config.KubeProxyIPTablesConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1alpha1_KubeProxyIPTablesConfiguration_To_config_KubeProxyIPTablesConfiguration(in, out, s)
}
func autoConvert_config_KubeProxyIPTablesConfiguration_To_v1alpha1_KubeProxyIPTablesConfiguration(in *config.KubeProxyIPTablesConfiguration, out *v1alpha1.KubeProxyIPTablesConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.MasqueradeBit = (*int32)(unsafe.Pointer(in.MasqueradeBit))
 out.MasqueradeAll = in.MasqueradeAll
 out.SyncPeriod = in.SyncPeriod
 out.MinSyncPeriod = in.MinSyncPeriod
 return nil
}
func Convert_config_KubeProxyIPTablesConfiguration_To_v1alpha1_KubeProxyIPTablesConfiguration(in *config.KubeProxyIPTablesConfiguration, out *v1alpha1.KubeProxyIPTablesConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_config_KubeProxyIPTablesConfiguration_To_v1alpha1_KubeProxyIPTablesConfiguration(in, out, s)
}
func autoConvert_v1alpha1_KubeProxyIPVSConfiguration_To_config_KubeProxyIPVSConfiguration(in *v1alpha1.KubeProxyIPVSConfiguration, out *config.KubeProxyIPVSConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.SyncPeriod = in.SyncPeriod
 out.MinSyncPeriod = in.MinSyncPeriod
 out.Scheduler = in.Scheduler
 out.ExcludeCIDRs = *(*[]string)(unsafe.Pointer(&in.ExcludeCIDRs))
 return nil
}
func Convert_v1alpha1_KubeProxyIPVSConfiguration_To_config_KubeProxyIPVSConfiguration(in *v1alpha1.KubeProxyIPVSConfiguration, out *config.KubeProxyIPVSConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1alpha1_KubeProxyIPVSConfiguration_To_config_KubeProxyIPVSConfiguration(in, out, s)
}
func autoConvert_config_KubeProxyIPVSConfiguration_To_v1alpha1_KubeProxyIPVSConfiguration(in *config.KubeProxyIPVSConfiguration, out *v1alpha1.KubeProxyIPVSConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.SyncPeriod = in.SyncPeriod
 out.MinSyncPeriod = in.MinSyncPeriod
 out.Scheduler = in.Scheduler
 out.ExcludeCIDRs = *(*[]string)(unsafe.Pointer(&in.ExcludeCIDRs))
 return nil
}
func Convert_config_KubeProxyIPVSConfiguration_To_v1alpha1_KubeProxyIPVSConfiguration(in *config.KubeProxyIPVSConfiguration, out *v1alpha1.KubeProxyIPVSConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_config_KubeProxyIPVSConfiguration_To_v1alpha1_KubeProxyIPVSConfiguration(in, out, s)
}
