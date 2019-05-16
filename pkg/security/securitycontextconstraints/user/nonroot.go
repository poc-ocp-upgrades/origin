package user

import (
	securityapi "github.com/openshift/origin/pkg/security/apis/security"
	"k8s.io/apimachinery/pkg/util/validation/field"
	api "k8s.io/kubernetes/pkg/apis/core"
)

type nonRoot struct{}

var _ RunAsUserSecurityContextConstraintsStrategy = &nonRoot{}

func NewRunAsNonRoot(options *securityapi.RunAsUserStrategyOptions) (RunAsUserSecurityContextConstraintsStrategy, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &nonRoot{}, nil
}
func (s *nonRoot) Generate(pod *api.Pod, container *api.Container) (*int64, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return nil, nil
}
func (s *nonRoot) Validate(fldPath *field.Path, _ *api.Pod, _ *api.Container, runAsNonRoot *bool, runAsUser *int64) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	allErrs := field.ErrorList{}
	if runAsNonRoot == nil && runAsUser == nil {
		allErrs = append(allErrs, field.Required(fldPath.Child("runAsNonRoot"), "must be true"))
		return allErrs
	}
	if runAsNonRoot != nil && *runAsNonRoot == false {
		allErrs = append(allErrs, field.Invalid(fldPath.Child("runAsNonRoot"), *runAsNonRoot, "must be true"))
		return allErrs
	}
	if runAsUser != nil && *runAsUser == 0 {
		allErrs = append(allErrs, field.Invalid(fldPath.Child("runAsUser"), *runAsUser, "running with the root UID is forbidden"))
		return allErrs
	}
	return allErrs
}
