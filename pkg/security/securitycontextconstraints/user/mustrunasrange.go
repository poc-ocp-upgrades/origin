package user

import (
	"fmt"
	"k8s.io/apimachinery/pkg/util/validation/field"
	api "k8s.io/kubernetes/pkg/apis/core"
	securityapi "github.com/openshift/origin/pkg/security/apis/security"
)

type mustRunAsRange struct {
	opts *securityapi.RunAsUserStrategyOptions
}

var _ RunAsUserSecurityContextConstraintsStrategy = &mustRunAsRange{}

func NewMustRunAsRange(options *securityapi.RunAsUserStrategyOptions) (RunAsUserSecurityContextConstraintsStrategy, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if options == nil {
		return nil, fmt.Errorf("MustRunAsRange requires run as user options")
	}
	if options.UIDRangeMin == nil {
		return nil, fmt.Errorf("MustRunAsRange requires a UIDRangeMin")
	}
	if options.UIDRangeMax == nil {
		return nil, fmt.Errorf("MustRunAsRange requires a UIDRangeMax")
	}
	return &mustRunAsRange{opts: options}, nil
}
func (s *mustRunAsRange) Generate(pod *api.Pod, container *api.Container) (*int64, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.opts.UIDRangeMin, nil
}
func (s *mustRunAsRange) Validate(fldPath *field.Path, _ *api.Pod, _ *api.Container, runAsNonRoot *bool, runAsUser *int64) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	allErrs := field.ErrorList{}
	if runAsUser == nil {
		allErrs = append(allErrs, field.Required(fldPath.Child("runAsUser"), ""))
		return allErrs
	}
	if *runAsUser < *s.opts.UIDRangeMin || *runAsUser > *s.opts.UIDRangeMax {
		detail := fmt.Sprintf("must be in the ranges: [%v, %v]", *s.opts.UIDRangeMin, *s.opts.UIDRangeMax)
		allErrs = append(allErrs, field.Invalid(fldPath.Child("runAsUser"), *runAsUser, detail))
		return allErrs
	}
	return allErrs
}
