package v1alpha1

import (
 unsafe "unsafe"
 configv1alpha1 "k8s.io/apimachinery/pkg/apis/config/v1alpha1"
 conversion "k8s.io/apimachinery/pkg/conversion"
 runtime "k8s.io/apimachinery/pkg/runtime"
 apisconfigv1alpha1 "k8s.io/apiserver/pkg/apis/config/v1alpha1"
 v1alpha1 "k8s.io/kube-scheduler/config/v1alpha1"
 config "k8s.io/kubernetes/pkg/scheduler/apis/config"
)

func init() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 localSchemeBuilder.Register(RegisterConversions)
}
func RegisterConversions(s *runtime.Scheme) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := s.AddGeneratedConversionFunc((*v1alpha1.KubeSchedulerConfiguration)(nil), (*config.KubeSchedulerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1alpha1_KubeSchedulerConfiguration_To_config_KubeSchedulerConfiguration(a.(*v1alpha1.KubeSchedulerConfiguration), b.(*config.KubeSchedulerConfiguration), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*config.KubeSchedulerConfiguration)(nil), (*v1alpha1.KubeSchedulerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_config_KubeSchedulerConfiguration_To_v1alpha1_KubeSchedulerConfiguration(a.(*config.KubeSchedulerConfiguration), b.(*v1alpha1.KubeSchedulerConfiguration), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1alpha1.KubeSchedulerLeaderElectionConfiguration)(nil), (*config.KubeSchedulerLeaderElectionConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1alpha1_KubeSchedulerLeaderElectionConfiguration_To_config_KubeSchedulerLeaderElectionConfiguration(a.(*v1alpha1.KubeSchedulerLeaderElectionConfiguration), b.(*config.KubeSchedulerLeaderElectionConfiguration), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*config.KubeSchedulerLeaderElectionConfiguration)(nil), (*v1alpha1.KubeSchedulerLeaderElectionConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_config_KubeSchedulerLeaderElectionConfiguration_To_v1alpha1_KubeSchedulerLeaderElectionConfiguration(a.(*config.KubeSchedulerLeaderElectionConfiguration), b.(*v1alpha1.KubeSchedulerLeaderElectionConfiguration), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1alpha1.SchedulerAlgorithmSource)(nil), (*config.SchedulerAlgorithmSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1alpha1_SchedulerAlgorithmSource_To_config_SchedulerAlgorithmSource(a.(*v1alpha1.SchedulerAlgorithmSource), b.(*config.SchedulerAlgorithmSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*config.SchedulerAlgorithmSource)(nil), (*v1alpha1.SchedulerAlgorithmSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_config_SchedulerAlgorithmSource_To_v1alpha1_SchedulerAlgorithmSource(a.(*config.SchedulerAlgorithmSource), b.(*v1alpha1.SchedulerAlgorithmSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1alpha1.SchedulerPolicyConfigMapSource)(nil), (*config.SchedulerPolicyConfigMapSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1alpha1_SchedulerPolicyConfigMapSource_To_config_SchedulerPolicyConfigMapSource(a.(*v1alpha1.SchedulerPolicyConfigMapSource), b.(*config.SchedulerPolicyConfigMapSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*config.SchedulerPolicyConfigMapSource)(nil), (*v1alpha1.SchedulerPolicyConfigMapSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_config_SchedulerPolicyConfigMapSource_To_v1alpha1_SchedulerPolicyConfigMapSource(a.(*config.SchedulerPolicyConfigMapSource), b.(*v1alpha1.SchedulerPolicyConfigMapSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1alpha1.SchedulerPolicyFileSource)(nil), (*config.SchedulerPolicyFileSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1alpha1_SchedulerPolicyFileSource_To_config_SchedulerPolicyFileSource(a.(*v1alpha1.SchedulerPolicyFileSource), b.(*config.SchedulerPolicyFileSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*config.SchedulerPolicyFileSource)(nil), (*v1alpha1.SchedulerPolicyFileSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_config_SchedulerPolicyFileSource_To_v1alpha1_SchedulerPolicyFileSource(a.(*config.SchedulerPolicyFileSource), b.(*v1alpha1.SchedulerPolicyFileSource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*v1alpha1.SchedulerPolicySource)(nil), (*config.SchedulerPolicySource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_v1alpha1_SchedulerPolicySource_To_config_SchedulerPolicySource(a.(*v1alpha1.SchedulerPolicySource), b.(*config.SchedulerPolicySource), scope)
 }); err != nil {
  return err
 }
 if err := s.AddGeneratedConversionFunc((*config.SchedulerPolicySource)(nil), (*v1alpha1.SchedulerPolicySource)(nil), func(a, b interface{}, scope conversion.Scope) error {
  return Convert_config_SchedulerPolicySource_To_v1alpha1_SchedulerPolicySource(a.(*config.SchedulerPolicySource), b.(*v1alpha1.SchedulerPolicySource), scope)
 }); err != nil {
  return err
 }
 return nil
}
func autoConvert_v1alpha1_KubeSchedulerConfiguration_To_config_KubeSchedulerConfiguration(in *v1alpha1.KubeSchedulerConfiguration, out *config.KubeSchedulerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.SchedulerName = in.SchedulerName
 if err := Convert_v1alpha1_SchedulerAlgorithmSource_To_config_SchedulerAlgorithmSource(&in.AlgorithmSource, &out.AlgorithmSource, s); err != nil {
  return err
 }
 out.HardPodAffinitySymmetricWeight = in.HardPodAffinitySymmetricWeight
 if err := Convert_v1alpha1_KubeSchedulerLeaderElectionConfiguration_To_config_KubeSchedulerLeaderElectionConfiguration(&in.LeaderElection, &out.LeaderElection, s); err != nil {
  return err
 }
 if err := configv1alpha1.Convert_v1alpha1_ClientConnectionConfiguration_To_config_ClientConnectionConfiguration(&in.ClientConnection, &out.ClientConnection, s); err != nil {
  return err
 }
 out.HealthzBindAddress = in.HealthzBindAddress
 out.MetricsBindAddress = in.MetricsBindAddress
 if err := apisconfigv1alpha1.Convert_v1alpha1_DebuggingConfiguration_To_config_DebuggingConfiguration(&in.DebuggingConfiguration, &out.DebuggingConfiguration, s); err != nil {
  return err
 }
 out.DisablePreemption = in.DisablePreemption
 out.PercentageOfNodesToScore = in.PercentageOfNodesToScore
 out.FailureDomains = in.FailureDomains
 out.BindTimeoutSeconds = (*int64)(unsafe.Pointer(in.BindTimeoutSeconds))
 return nil
}
func Convert_v1alpha1_KubeSchedulerConfiguration_To_config_KubeSchedulerConfiguration(in *v1alpha1.KubeSchedulerConfiguration, out *config.KubeSchedulerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1alpha1_KubeSchedulerConfiguration_To_config_KubeSchedulerConfiguration(in, out, s)
}
func autoConvert_config_KubeSchedulerConfiguration_To_v1alpha1_KubeSchedulerConfiguration(in *config.KubeSchedulerConfiguration, out *v1alpha1.KubeSchedulerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.SchedulerName = in.SchedulerName
 if err := Convert_config_SchedulerAlgorithmSource_To_v1alpha1_SchedulerAlgorithmSource(&in.AlgorithmSource, &out.AlgorithmSource, s); err != nil {
  return err
 }
 out.HardPodAffinitySymmetricWeight = in.HardPodAffinitySymmetricWeight
 if err := Convert_config_KubeSchedulerLeaderElectionConfiguration_To_v1alpha1_KubeSchedulerLeaderElectionConfiguration(&in.LeaderElection, &out.LeaderElection, s); err != nil {
  return err
 }
 if err := configv1alpha1.Convert_config_ClientConnectionConfiguration_To_v1alpha1_ClientConnectionConfiguration(&in.ClientConnection, &out.ClientConnection, s); err != nil {
  return err
 }
 out.HealthzBindAddress = in.HealthzBindAddress
 out.MetricsBindAddress = in.MetricsBindAddress
 if err := apisconfigv1alpha1.Convert_config_DebuggingConfiguration_To_v1alpha1_DebuggingConfiguration(&in.DebuggingConfiguration, &out.DebuggingConfiguration, s); err != nil {
  return err
 }
 out.DisablePreemption = in.DisablePreemption
 out.PercentageOfNodesToScore = in.PercentageOfNodesToScore
 out.FailureDomains = in.FailureDomains
 out.BindTimeoutSeconds = (*int64)(unsafe.Pointer(in.BindTimeoutSeconds))
 return nil
}
func Convert_config_KubeSchedulerConfiguration_To_v1alpha1_KubeSchedulerConfiguration(in *config.KubeSchedulerConfiguration, out *v1alpha1.KubeSchedulerConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_config_KubeSchedulerConfiguration_To_v1alpha1_KubeSchedulerConfiguration(in, out, s)
}
func autoConvert_v1alpha1_KubeSchedulerLeaderElectionConfiguration_To_config_KubeSchedulerLeaderElectionConfiguration(in *v1alpha1.KubeSchedulerLeaderElectionConfiguration, out *config.KubeSchedulerLeaderElectionConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := apisconfigv1alpha1.Convert_v1alpha1_LeaderElectionConfiguration_To_config_LeaderElectionConfiguration(&in.LeaderElectionConfiguration, &out.LeaderElectionConfiguration, s); err != nil {
  return err
 }
 out.LockObjectNamespace = in.LockObjectNamespace
 out.LockObjectName = in.LockObjectName
 return nil
}
func Convert_v1alpha1_KubeSchedulerLeaderElectionConfiguration_To_config_KubeSchedulerLeaderElectionConfiguration(in *v1alpha1.KubeSchedulerLeaderElectionConfiguration, out *config.KubeSchedulerLeaderElectionConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1alpha1_KubeSchedulerLeaderElectionConfiguration_To_config_KubeSchedulerLeaderElectionConfiguration(in, out, s)
}
func autoConvert_config_KubeSchedulerLeaderElectionConfiguration_To_v1alpha1_KubeSchedulerLeaderElectionConfiguration(in *config.KubeSchedulerLeaderElectionConfiguration, out *v1alpha1.KubeSchedulerLeaderElectionConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := apisconfigv1alpha1.Convert_config_LeaderElectionConfiguration_To_v1alpha1_LeaderElectionConfiguration(&in.LeaderElectionConfiguration, &out.LeaderElectionConfiguration, s); err != nil {
  return err
 }
 out.LockObjectNamespace = in.LockObjectNamespace
 out.LockObjectName = in.LockObjectName
 return nil
}
func Convert_config_KubeSchedulerLeaderElectionConfiguration_To_v1alpha1_KubeSchedulerLeaderElectionConfiguration(in *config.KubeSchedulerLeaderElectionConfiguration, out *v1alpha1.KubeSchedulerLeaderElectionConfiguration, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_config_KubeSchedulerLeaderElectionConfiguration_To_v1alpha1_KubeSchedulerLeaderElectionConfiguration(in, out, s)
}
func autoConvert_v1alpha1_SchedulerAlgorithmSource_To_config_SchedulerAlgorithmSource(in *v1alpha1.SchedulerAlgorithmSource, out *config.SchedulerAlgorithmSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Policy = (*config.SchedulerPolicySource)(unsafe.Pointer(in.Policy))
 out.Provider = (*string)(unsafe.Pointer(in.Provider))
 return nil
}
func Convert_v1alpha1_SchedulerAlgorithmSource_To_config_SchedulerAlgorithmSource(in *v1alpha1.SchedulerAlgorithmSource, out *config.SchedulerAlgorithmSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1alpha1_SchedulerAlgorithmSource_To_config_SchedulerAlgorithmSource(in, out, s)
}
func autoConvert_config_SchedulerAlgorithmSource_To_v1alpha1_SchedulerAlgorithmSource(in *config.SchedulerAlgorithmSource, out *v1alpha1.SchedulerAlgorithmSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Policy = (*v1alpha1.SchedulerPolicySource)(unsafe.Pointer(in.Policy))
 out.Provider = (*string)(unsafe.Pointer(in.Provider))
 return nil
}
func Convert_config_SchedulerAlgorithmSource_To_v1alpha1_SchedulerAlgorithmSource(in *config.SchedulerAlgorithmSource, out *v1alpha1.SchedulerAlgorithmSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_config_SchedulerAlgorithmSource_To_v1alpha1_SchedulerAlgorithmSource(in, out, s)
}
func autoConvert_v1alpha1_SchedulerPolicyConfigMapSource_To_config_SchedulerPolicyConfigMapSource(in *v1alpha1.SchedulerPolicyConfigMapSource, out *config.SchedulerPolicyConfigMapSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Namespace = in.Namespace
 out.Name = in.Name
 return nil
}
func Convert_v1alpha1_SchedulerPolicyConfigMapSource_To_config_SchedulerPolicyConfigMapSource(in *v1alpha1.SchedulerPolicyConfigMapSource, out *config.SchedulerPolicyConfigMapSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1alpha1_SchedulerPolicyConfigMapSource_To_config_SchedulerPolicyConfigMapSource(in, out, s)
}
func autoConvert_config_SchedulerPolicyConfigMapSource_To_v1alpha1_SchedulerPolicyConfigMapSource(in *config.SchedulerPolicyConfigMapSource, out *v1alpha1.SchedulerPolicyConfigMapSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Namespace = in.Namespace
 out.Name = in.Name
 return nil
}
func Convert_config_SchedulerPolicyConfigMapSource_To_v1alpha1_SchedulerPolicyConfigMapSource(in *config.SchedulerPolicyConfigMapSource, out *v1alpha1.SchedulerPolicyConfigMapSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_config_SchedulerPolicyConfigMapSource_To_v1alpha1_SchedulerPolicyConfigMapSource(in, out, s)
}
func autoConvert_v1alpha1_SchedulerPolicyFileSource_To_config_SchedulerPolicyFileSource(in *v1alpha1.SchedulerPolicyFileSource, out *config.SchedulerPolicyFileSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Path = in.Path
 return nil
}
func Convert_v1alpha1_SchedulerPolicyFileSource_To_config_SchedulerPolicyFileSource(in *v1alpha1.SchedulerPolicyFileSource, out *config.SchedulerPolicyFileSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1alpha1_SchedulerPolicyFileSource_To_config_SchedulerPolicyFileSource(in, out, s)
}
func autoConvert_config_SchedulerPolicyFileSource_To_v1alpha1_SchedulerPolicyFileSource(in *config.SchedulerPolicyFileSource, out *v1alpha1.SchedulerPolicyFileSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Path = in.Path
 return nil
}
func Convert_config_SchedulerPolicyFileSource_To_v1alpha1_SchedulerPolicyFileSource(in *config.SchedulerPolicyFileSource, out *v1alpha1.SchedulerPolicyFileSource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_config_SchedulerPolicyFileSource_To_v1alpha1_SchedulerPolicyFileSource(in, out, s)
}
func autoConvert_v1alpha1_SchedulerPolicySource_To_config_SchedulerPolicySource(in *v1alpha1.SchedulerPolicySource, out *config.SchedulerPolicySource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.File = (*config.SchedulerPolicyFileSource)(unsafe.Pointer(in.File))
 out.ConfigMap = (*config.SchedulerPolicyConfigMapSource)(unsafe.Pointer(in.ConfigMap))
 return nil
}
func Convert_v1alpha1_SchedulerPolicySource_To_config_SchedulerPolicySource(in *v1alpha1.SchedulerPolicySource, out *config.SchedulerPolicySource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_v1alpha1_SchedulerPolicySource_To_config_SchedulerPolicySource(in, out, s)
}
func autoConvert_config_SchedulerPolicySource_To_v1alpha1_SchedulerPolicySource(in *config.SchedulerPolicySource, out *v1alpha1.SchedulerPolicySource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.File = (*v1alpha1.SchedulerPolicyFileSource)(unsafe.Pointer(in.File))
 out.ConfigMap = (*v1alpha1.SchedulerPolicyConfigMapSource)(unsafe.Pointer(in.ConfigMap))
 return nil
}
func Convert_config_SchedulerPolicySource_To_v1alpha1_SchedulerPolicySource(in *config.SchedulerPolicySource, out *v1alpha1.SchedulerPolicySource, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return autoConvert_config_SchedulerPolicySource_To_v1alpha1_SchedulerPolicySource(in, out, s)
}
