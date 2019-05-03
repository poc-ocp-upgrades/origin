package podsecuritypolicy

import (
 "fmt"
 "strings"
 corev1 "k8s.io/api/core/v1"
 policy "k8s.io/api/policy/v1beta1"
 "k8s.io/apimachinery/pkg/util/validation/field"
 utilfeature "k8s.io/apiserver/pkg/util/feature"
 api "k8s.io/kubernetes/pkg/apis/core"
 "k8s.io/kubernetes/pkg/features"
 psputil "k8s.io/kubernetes/pkg/security/podsecuritypolicy/util"
 "k8s.io/kubernetes/pkg/securitycontext"
)

type simpleProvider struct {
 psp        *policy.PodSecurityPolicy
 strategies *ProviderStrategies
}

var _ Provider = &simpleProvider{}

func NewSimpleProvider(psp *policy.PodSecurityPolicy, namespace string, strategyFactory StrategyFactory) (Provider, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if psp == nil {
  return nil, fmt.Errorf("NewSimpleProvider requires a PodSecurityPolicy")
 }
 if strategyFactory == nil {
  return nil, fmt.Errorf("NewSimpleProvider requires a StrategyFactory")
 }
 strategies, err := strategyFactory.CreateStrategies(psp, namespace)
 if err != nil {
  return nil, err
 }
 return &simpleProvider{psp: psp, strategies: strategies}, nil
}
func (s *simpleProvider) DefaultPodSecurityContext(pod *api.Pod) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 sc := securitycontext.NewPodSecurityContextMutator(pod.Spec.SecurityContext)
 if sc.SupplementalGroups() == nil {
  supGroups, err := s.strategies.SupplementalGroupStrategy.Generate(pod)
  if err != nil {
   return err
  }
  sc.SetSupplementalGroups(supGroups)
 }
 if sc.FSGroup() == nil {
  fsGroup, err := s.strategies.FSGroupStrategy.GenerateSingle(pod)
  if err != nil {
   return err
  }
  sc.SetFSGroup(fsGroup)
 }
 if sc.SELinuxOptions() == nil {
  seLinux, err := s.strategies.SELinuxStrategy.Generate(pod, nil)
  if err != nil {
   return err
  }
  sc.SetSELinuxOptions(seLinux)
 }
 seccompProfile, err := s.strategies.SeccompStrategy.Generate(pod.Annotations, pod)
 if err != nil {
  return err
 }
 if seccompProfile != "" {
  if pod.Annotations == nil {
   pod.Annotations = map[string]string{}
  }
  pod.Annotations[api.SeccompPodAnnotationKey] = seccompProfile
 }
 pod.Spec.SecurityContext = sc.PodSecurityContext()
 return nil
}
func (s *simpleProvider) DefaultContainerSecurityContext(pod *api.Pod, container *api.Container) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 sc := securitycontext.NewEffectiveContainerSecurityContextMutator(securitycontext.NewPodSecurityContextAccessor(pod.Spec.SecurityContext), securitycontext.NewContainerSecurityContextMutator(container.SecurityContext))
 if sc.RunAsUser() == nil {
  uid, err := s.strategies.RunAsUserStrategy.Generate(pod, container)
  if err != nil {
   return err
  }
  sc.SetRunAsUser(uid)
 }
 if utilfeature.DefaultFeatureGate.Enabled(features.RunAsGroup) {
  if sc.RunAsGroup() == nil {
   gid, err := s.strategies.RunAsGroupStrategy.GenerateSingle(pod)
   if err != nil {
    return err
   }
   sc.SetRunAsGroup(gid)
  }
 }
 if sc.SELinuxOptions() == nil {
  seLinux, err := s.strategies.SELinuxStrategy.Generate(pod, container)
  if err != nil {
   return err
  }
  sc.SetSELinuxOptions(seLinux)
 }
 annotations, err := s.strategies.AppArmorStrategy.Generate(pod.Annotations, container)
 if err != nil {
  return err
 }
 if sc.RunAsNonRoot() == nil && sc.RunAsUser() == nil && s.psp.Spec.RunAsUser.Rule == policy.RunAsUserStrategyMustRunAsNonRoot {
  nonRoot := true
  sc.SetRunAsNonRoot(&nonRoot)
 }
 caps, err := s.strategies.CapabilitiesStrategy.Generate(pod, container)
 if err != nil {
  return err
 }
 sc.SetCapabilities(caps)
 if s.psp.Spec.ReadOnlyRootFilesystem && sc.ReadOnlyRootFilesystem() == nil {
  readOnlyRootFS := true
  sc.SetReadOnlyRootFilesystem(&readOnlyRootFS)
 }
 if s.psp.Spec.DefaultAllowPrivilegeEscalation != nil && sc.AllowPrivilegeEscalation() == nil {
  sc.SetAllowPrivilegeEscalation(s.psp.Spec.DefaultAllowPrivilegeEscalation)
 }
 if !*s.psp.Spec.AllowPrivilegeEscalation && sc.AllowPrivilegeEscalation() == nil {
  sc.SetAllowPrivilegeEscalation(s.psp.Spec.AllowPrivilegeEscalation)
 }
 pod.Annotations = annotations
 container.SecurityContext = sc.ContainerSecurityContext()
 return nil
}
func (s *simpleProvider) ValidatePod(pod *api.Pod) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 sc := securitycontext.NewPodSecurityContextAccessor(pod.Spec.SecurityContext)
 scPath := field.NewPath("spec", "securityContext")
 var fsGroups []int64
 if fsGroup := sc.FSGroup(); fsGroup != nil {
  fsGroups = []int64{*fsGroup}
 }
 allErrs = append(allErrs, s.strategies.FSGroupStrategy.Validate(scPath.Child("fsGroup"), pod, fsGroups)...)
 allErrs = append(allErrs, s.strategies.SupplementalGroupStrategy.Validate(scPath.Child("supplementalGroups"), pod, sc.SupplementalGroups())...)
 allErrs = append(allErrs, s.strategies.SeccompStrategy.ValidatePod(pod)...)
 allErrs = append(allErrs, s.strategies.SELinuxStrategy.Validate(scPath.Child("seLinuxOptions"), pod, nil, sc.SELinuxOptions())...)
 if !s.psp.Spec.HostNetwork && sc.HostNetwork() {
  allErrs = append(allErrs, field.Invalid(scPath.Child("hostNetwork"), sc.HostNetwork(), "Host network is not allowed to be used"))
 }
 if !s.psp.Spec.HostPID && sc.HostPID() {
  allErrs = append(allErrs, field.Invalid(scPath.Child("hostPID"), sc.HostPID(), "Host PID is not allowed to be used"))
 }
 if !s.psp.Spec.HostIPC && sc.HostIPC() {
  allErrs = append(allErrs, field.Invalid(scPath.Child("hostIPC"), sc.HostIPC(), "Host IPC is not allowed to be used"))
 }
 allErrs = append(allErrs, s.strategies.SysctlsStrategy.Validate(pod)...)
 if len(pod.Spec.Volumes) > 0 {
  allowsAllVolumeTypes := psputil.PSPAllowsAllVolumes(s.psp)
  allowedVolumes := psputil.FSTypeToStringSet(s.psp.Spec.Volumes)
  for i, v := range pod.Spec.Volumes {
   fsType, err := psputil.GetVolumeFSType(v)
   if err != nil {
    allErrs = append(allErrs, field.Invalid(field.NewPath("spec", "volumes").Index(i), string(fsType), err.Error()))
    continue
   }
   if !allowsAllVolumeTypes && !allowedVolumes.Has(string(fsType)) {
    allErrs = append(allErrs, field.Invalid(field.NewPath("spec", "volumes").Index(i), string(fsType), fmt.Sprintf("%s volumes are not allowed to be used", string(fsType))))
    continue
   }
   if fsType == policy.HostPath {
    allows, mustBeReadOnly := psputil.AllowsHostVolumePath(s.psp, v.HostPath.Path)
    if !allows {
     allErrs = append(allErrs, field.Invalid(field.NewPath("spec", "volumes").Index(i).Child("hostPath", "pathPrefix"), v.HostPath.Path, fmt.Sprintf("is not allowed to be used")))
    } else if mustBeReadOnly {
     for i, c := range pod.Spec.InitContainers {
      for j, cv := range c.VolumeMounts {
       if cv.Name == v.Name && !cv.ReadOnly {
        allErrs = append(allErrs, field.Invalid(field.NewPath("spec", "initContainers").Index(i).Child("volumeMounts").Index(j).Child("readOnly"), cv.ReadOnly, "must be read-only"))
       }
      }
     }
     for i, c := range pod.Spec.Containers {
      for j, cv := range c.VolumeMounts {
       if cv.Name == v.Name && !cv.ReadOnly {
        allErrs = append(allErrs, field.Invalid(field.NewPath("spec", "containers").Index(i).Child("volumeMounts").Index(j).Child("readOnly"), cv.ReadOnly, "must be read-only"))
       }
      }
     }
    }
   }
   if fsType == policy.FlexVolume && len(s.psp.Spec.AllowedFlexVolumes) > 0 {
    found := false
    driver := v.FlexVolume.Driver
    for _, allowedFlexVolume := range s.psp.Spec.AllowedFlexVolumes {
     if driver == allowedFlexVolume.Driver {
      found = true
      break
     }
    }
    if !found {
     allErrs = append(allErrs, field.Invalid(field.NewPath("spec", "volumes").Index(i).Child("driver"), driver, "Flexvolume driver is not allowed to be used"))
    }
   }
  }
 }
 return allErrs
}
func (s *simpleProvider) ValidateContainer(pod *api.Pod, container *api.Container, containerPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 podSC := securitycontext.NewPodSecurityContextAccessor(pod.Spec.SecurityContext)
 sc := securitycontext.NewEffectiveContainerSecurityContextAccessor(podSC, securitycontext.NewContainerSecurityContextMutator(container.SecurityContext))
 scPath := containerPath.Child("securityContext")
 allErrs = append(allErrs, s.strategies.RunAsUserStrategy.Validate(scPath, pod, container, sc.RunAsNonRoot(), sc.RunAsUser())...)
 if utilfeature.DefaultFeatureGate.Enabled(features.RunAsGroup) {
  var runAsGroups []int64
  if sc.RunAsGroup() != nil {
   runAsGroups = []int64{*sc.RunAsGroup()}
  }
  allErrs = append(allErrs, s.strategies.RunAsGroupStrategy.Validate(scPath, pod, runAsGroups)...)
 }
 allErrs = append(allErrs, s.strategies.SELinuxStrategy.Validate(scPath.Child("seLinuxOptions"), pod, container, sc.SELinuxOptions())...)
 allErrs = append(allErrs, s.strategies.AppArmorStrategy.Validate(pod, container)...)
 allErrs = append(allErrs, s.strategies.SeccompStrategy.ValidateContainer(pod, container)...)
 privileged := sc.Privileged()
 if !s.psp.Spec.Privileged && privileged != nil && *privileged {
  allErrs = append(allErrs, field.Invalid(scPath.Child("privileged"), *privileged, "Privileged containers are not allowed"))
 }
 procMount := sc.ProcMount()
 allowedProcMounts := s.psp.Spec.AllowedProcMountTypes
 if len(allowedProcMounts) == 0 {
  allowedProcMounts = []corev1.ProcMountType{corev1.DefaultProcMount}
 }
 foundProcMountType := false
 for _, pm := range allowedProcMounts {
  if string(pm) == string(procMount) {
   foundProcMountType = true
  }
 }
 if !foundProcMountType {
  allErrs = append(allErrs, field.Invalid(scPath.Child("procMount"), procMount, "ProcMountType is not allowed"))
 }
 allErrs = append(allErrs, s.strategies.CapabilitiesStrategy.Validate(scPath.Child("capabilities"), pod, container, sc.Capabilities())...)
 allErrs = append(allErrs, s.hasInvalidHostPort(container, containerPath)...)
 if s.psp.Spec.ReadOnlyRootFilesystem {
  readOnly := sc.ReadOnlyRootFilesystem()
  if readOnly == nil {
   allErrs = append(allErrs, field.Invalid(scPath.Child("readOnlyRootFilesystem"), readOnly, "ReadOnlyRootFilesystem may not be nil and must be set to true"))
  } else if !*readOnly {
   allErrs = append(allErrs, field.Invalid(scPath.Child("readOnlyRootFilesystem"), *readOnly, "ReadOnlyRootFilesystem must be set to true"))
  }
 }
 allowEscalation := sc.AllowPrivilegeEscalation()
 if !*s.psp.Spec.AllowPrivilegeEscalation && (allowEscalation == nil || *allowEscalation) {
  allErrs = append(allErrs, field.Invalid(scPath.Child("allowPrivilegeEscalation"), allowEscalation, "Allowing privilege escalation for containers is not allowed"))
 }
 return allErrs
}
func (s *simpleProvider) hasInvalidHostPort(container *api.Container, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 for _, cp := range container.Ports {
  if cp.HostPort > 0 && !s.isValidHostPort(cp.HostPort) {
   detail := fmt.Sprintf("Host port %d is not allowed to be used. Allowed ports: [%s]", cp.HostPort, hostPortRangesToString(s.psp.Spec.HostPorts))
   allErrs = append(allErrs, field.Invalid(fldPath.Child("hostPort"), cp.HostPort, detail))
  }
 }
 return allErrs
}
func (s *simpleProvider) isValidHostPort(port int32) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for _, hostPortRange := range s.psp.Spec.HostPorts {
  if port >= hostPortRange.Min && port <= hostPortRange.Max {
   return true
  }
 }
 return false
}
func (s *simpleProvider) GetPSPName() string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return s.psp.Name
}
func hostPortRangesToString(ranges []policy.HostPortRange) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 formattedString := ""
 if ranges != nil {
  strRanges := []string{}
  for _, r := range ranges {
   if r.Min == r.Max {
    strRanges = append(strRanges, fmt.Sprintf("%d", r.Min))
   } else {
    strRanges = append(strRanges, fmt.Sprintf("%d-%d", r.Min, r.Max))
   }
  }
  formattedString = strings.Join(strRanges, ",")
 }
 return formattedString
}
