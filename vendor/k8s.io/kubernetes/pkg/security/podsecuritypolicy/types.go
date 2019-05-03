package podsecuritypolicy

import (
 policy "k8s.io/api/policy/v1beta1"
 "k8s.io/apimachinery/pkg/util/validation/field"
 api "k8s.io/kubernetes/pkg/apis/core"
 "k8s.io/kubernetes/pkg/security/podsecuritypolicy/apparmor"
 "k8s.io/kubernetes/pkg/security/podsecuritypolicy/capabilities"
 "k8s.io/kubernetes/pkg/security/podsecuritypolicy/group"
 "k8s.io/kubernetes/pkg/security/podsecuritypolicy/seccomp"
 "k8s.io/kubernetes/pkg/security/podsecuritypolicy/selinux"
 "k8s.io/kubernetes/pkg/security/podsecuritypolicy/sysctl"
 "k8s.io/kubernetes/pkg/security/podsecuritypolicy/user"
)

type Provider interface {
 DefaultPodSecurityContext(pod *api.Pod) error
 DefaultContainerSecurityContext(pod *api.Pod, container *api.Container) error
 ValidatePod(pod *api.Pod) field.ErrorList
 ValidateContainer(pod *api.Pod, container *api.Container, containerPath *field.Path) field.ErrorList
 GetPSPName() string
}
type StrategyFactory interface {
 CreateStrategies(psp *policy.PodSecurityPolicy, namespace string) (*ProviderStrategies, error)
}
type ProviderStrategies struct {
 RunAsUserStrategy         user.RunAsUserStrategy
 RunAsGroupStrategy        group.GroupStrategy
 SELinuxStrategy           selinux.SELinuxStrategy
 AppArmorStrategy          apparmor.Strategy
 FSGroupStrategy           group.GroupStrategy
 SupplementalGroupStrategy group.GroupStrategy
 CapabilitiesStrategy      capabilities.Strategy
 SysctlsStrategy           sysctl.SysctlsStrategy
 SeccompStrategy           seccomp.Strategy
}
