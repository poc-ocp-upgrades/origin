package securitycontext

import (
 "k8s.io/api/core/v1"
 api "k8s.io/kubernetes/pkg/apis/core"
)

func ValidSecurityContextWithContainerDefaults() *v1.SecurityContext {
 _logClusterCodePath()
 defer _logClusterCodePath()
 priv := false
 return &v1.SecurityContext{Capabilities: &v1.Capabilities{}, Privileged: &priv}
}
func ValidInternalSecurityContextWithContainerDefaults() *api.SecurityContext {
 _logClusterCodePath()
 defer _logClusterCodePath()
 priv := false
 dpm := api.DefaultProcMount
 return &api.SecurityContext{Capabilities: &api.Capabilities{}, Privileged: &priv, ProcMount: &dpm}
}
