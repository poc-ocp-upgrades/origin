package securitycontext

import (
 "fmt"
 "strings"
 "k8s.io/api/core/v1"
)

func HasPrivilegedRequest(container *v1.Container) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if container.SecurityContext == nil {
  return false
 }
 if container.SecurityContext.Privileged == nil {
  return false
 }
 return *container.SecurityContext.Privileged
}
func HasCapabilitiesRequest(container *v1.Container) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if container.SecurityContext == nil {
  return false
 }
 if container.SecurityContext.Capabilities == nil {
  return false
 }
 return len(container.SecurityContext.Capabilities.Add) > 0 || len(container.SecurityContext.Capabilities.Drop) > 0
}

const expectedSELinuxFields = 4

func ParseSELinuxOptions(context string) (*v1.SELinuxOptions, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 fields := strings.SplitN(context, ":", expectedSELinuxFields)
 if len(fields) != expectedSELinuxFields {
  return nil, fmt.Errorf("expected %v fields in selinux; got %v (context: %v)", expectedSELinuxFields, len(fields), context)
 }
 return &v1.SELinuxOptions{User: fields[0], Role: fields[1], Type: fields[2], Level: fields[3]}, nil
}
func DetermineEffectiveSecurityContext(pod *v1.Pod, container *v1.Container) *v1.SecurityContext {
 _logClusterCodePath()
 defer _logClusterCodePath()
 effectiveSc := securityContextFromPodSecurityContext(pod)
 containerSc := container.SecurityContext
 if effectiveSc == nil && containerSc == nil {
  return &v1.SecurityContext{}
 }
 if effectiveSc != nil && containerSc == nil {
  return effectiveSc
 }
 if effectiveSc == nil && containerSc != nil {
  return containerSc
 }
 if containerSc.SELinuxOptions != nil {
  effectiveSc.SELinuxOptions = new(v1.SELinuxOptions)
  *effectiveSc.SELinuxOptions = *containerSc.SELinuxOptions
 }
 if containerSc.Capabilities != nil {
  effectiveSc.Capabilities = new(v1.Capabilities)
  *effectiveSc.Capabilities = *containerSc.Capabilities
 }
 if containerSc.Privileged != nil {
  effectiveSc.Privileged = new(bool)
  *effectiveSc.Privileged = *containerSc.Privileged
 }
 if containerSc.RunAsUser != nil {
  effectiveSc.RunAsUser = new(int64)
  *effectiveSc.RunAsUser = *containerSc.RunAsUser
 }
 if containerSc.RunAsGroup != nil {
  effectiveSc.RunAsGroup = new(int64)
  *effectiveSc.RunAsGroup = *containerSc.RunAsGroup
 }
 if containerSc.RunAsNonRoot != nil {
  effectiveSc.RunAsNonRoot = new(bool)
  *effectiveSc.RunAsNonRoot = *containerSc.RunAsNonRoot
 }
 if containerSc.ReadOnlyRootFilesystem != nil {
  effectiveSc.ReadOnlyRootFilesystem = new(bool)
  *effectiveSc.ReadOnlyRootFilesystem = *containerSc.ReadOnlyRootFilesystem
 }
 if containerSc.AllowPrivilegeEscalation != nil {
  effectiveSc.AllowPrivilegeEscalation = new(bool)
  *effectiveSc.AllowPrivilegeEscalation = *containerSc.AllowPrivilegeEscalation
 }
 if containerSc.ProcMount != nil {
  effectiveSc.ProcMount = new(v1.ProcMountType)
  *effectiveSc.ProcMount = *containerSc.ProcMount
 }
 return effectiveSc
}
func securityContextFromPodSecurityContext(pod *v1.Pod) *v1.SecurityContext {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if pod.Spec.SecurityContext == nil {
  return nil
 }
 synthesized := &v1.SecurityContext{}
 if pod.Spec.SecurityContext.SELinuxOptions != nil {
  synthesized.SELinuxOptions = &v1.SELinuxOptions{}
  *synthesized.SELinuxOptions = *pod.Spec.SecurityContext.SELinuxOptions
 }
 if pod.Spec.SecurityContext.RunAsUser != nil {
  synthesized.RunAsUser = new(int64)
  *synthesized.RunAsUser = *pod.Spec.SecurityContext.RunAsUser
 }
 if pod.Spec.SecurityContext.RunAsGroup != nil {
  synthesized.RunAsGroup = new(int64)
  *synthesized.RunAsGroup = *pod.Spec.SecurityContext.RunAsGroup
 }
 if pod.Spec.SecurityContext.RunAsNonRoot != nil {
  synthesized.RunAsNonRoot = new(bool)
  *synthesized.RunAsNonRoot = *pod.Spec.SecurityContext.RunAsNonRoot
 }
 return synthesized
}
func AddNoNewPrivileges(sc *v1.SecurityContext) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if sc == nil {
  return false
 }
 if sc.AllowPrivilegeEscalation == nil {
  return false
 }
 return !*sc.AllowPrivilegeEscalation
}

var (
 defaultMaskedPaths   = []string{"/proc/acpi", "/proc/kcore", "/proc/keys", "/proc/latency_stats", "/proc/timer_list", "/proc/timer_stats", "/proc/sched_debug", "/proc/scsi", "/sys/firmware"}
 defaultReadonlyPaths = []string{"/proc/asound", "/proc/bus", "/proc/fs", "/proc/irq", "/proc/sys", "/proc/sysrq-trigger"}
)

func ConvertToRuntimeMaskedPaths(opt *v1.ProcMountType) []string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if opt != nil && *opt == v1.UnmaskedProcMount {
  return []string{}
 }
 return defaultMaskedPaths
}
func ConvertToRuntimeReadonlyPaths(opt *v1.ProcMountType) []string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if opt != nil && *opt == v1.UnmaskedProcMount {
  return []string{}
 }
 return defaultReadonlyPaths
}
