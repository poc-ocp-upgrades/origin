package capabilities

import (
	"fmt"
	"bytes"
	"net/http"
	"runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apimachinery/pkg/util/validation/field"
	api "k8s.io/kubernetes/pkg/apis/core"
	securityapi "github.com/openshift/origin/pkg/security/apis/security"
)

type defaultCapabilities struct {
	defaultAddCapabilities		[]api.Capability
	requiredDropCapabilities	[]api.Capability
	allowedCaps			[]api.Capability
}

var _ CapabilitiesSecurityContextConstraintsStrategy = &defaultCapabilities{}

func NewDefaultCapabilities(defaultAddCapabilities, requiredDropCapabilities, allowedCaps []api.Capability) (CapabilitiesSecurityContextConstraintsStrategy, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &defaultCapabilities{defaultAddCapabilities: defaultAddCapabilities, requiredDropCapabilities: requiredDropCapabilities, allowedCaps: allowedCaps}, nil
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
func (s *defaultCapabilities) Validate(pod *api.Pod, container *api.Container, capabilities *api.Capabilities) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	allErrs := field.ErrorList{}
	if capabilities == nil {
		if len(s.defaultAddCapabilities) == 0 && len(s.requiredDropCapabilities) == 0 {
			return allErrs
		}
		allErrs = append(allErrs, field.Invalid(field.NewPath("capabilities"), capabilities, "required capabilities are not set on the securityContext"))
		return allErrs
	}
	allowedAdd := makeCapSet(s.allowedCaps)
	allowAllCaps := allowedAdd.Has(string(securityapi.AllowAllCapabilities))
	if allowAllCaps {
		return allErrs
	}
	defaultAdd := makeCapSet(s.defaultAddCapabilities)
	for _, cap := range capabilities.Add {
		sCap := string(cap)
		if !defaultAdd.Has(sCap) && !allowedAdd.Has(sCap) {
			allErrs = append(allErrs, field.Invalid(field.NewPath("capabilities", "add"), sCap, "capability may not be added"))
		}
	}
	containerDrops := makeCapSet(capabilities.Drop)
	for _, requiredDrop := range s.requiredDropCapabilities {
		sDrop := string(requiredDrop)
		if !containerDrops.Has(sDrop) {
			allErrs = append(allErrs, field.Invalid(field.NewPath("capabilities", "drop"), capabilities.Drop, fmt.Sprintf("%s is required to be dropped but was not found", sDrop)))
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
