package user

import (
	"fmt"
	"k8s.io/apimachinery/pkg/util/validation/field"
	api "k8s.io/kubernetes/pkg/apis/core"
	securityapi "github.com/openshift/origin/pkg/security/apis/security"
)

type mustRunAs struct {
	opts *securityapi.RunAsUserStrategyOptions
}

var _ RunAsUserSecurityContextConstraintsStrategy = &mustRunAs{}

func NewMustRunAs(options *securityapi.RunAsUserStrategyOptions) (RunAsUserSecurityContextConstraintsStrategy, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if options == nil {
		return nil, fmt.Errorf("MustRunAs requires run as user options")
	}
	if options.UID == nil {
		return nil, fmt.Errorf("MustRunAs requires a UID")
	}
	return &mustRunAs{opts: options}, nil
}
func (s *mustRunAs) Generate(pod *api.Pod, container *api.Container) (*int64, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.opts.UID, nil
}
func (s *mustRunAs) Validate(fldPath *field.Path, _ *api.Pod, _ *api.Container, runAsNonRoot *bool, runAsUser *int64) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	allErrs := field.ErrorList{}
	if runAsUser == nil {
		allErrs = append(allErrs, field.Required(fldPath.Child("runAsUser"), ""))
		return allErrs
	}
	if *s.opts.UID != *runAsUser {
		detail := fmt.Sprintf("must be: %v", *s.opts.UID)
		allErrs = append(allErrs, field.Invalid(fldPath.Child("runAsUser"), *runAsUser, detail))
		return allErrs
	}
	return allErrs
}
