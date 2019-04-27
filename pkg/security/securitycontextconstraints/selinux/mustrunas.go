package selinux

import (
	"fmt"
	"sort"
	"strings"
	"k8s.io/apimachinery/pkg/util/validation/field"
	api "k8s.io/kubernetes/pkg/apis/core"
	securityapi "github.com/openshift/origin/pkg/security/apis/security"
	"github.com/openshift/origin/pkg/security/securitycontextconstraints/util"
)

type mustRunAs struct {
	opts *securityapi.SELinuxContextStrategyOptions
}

var _ SELinuxSecurityContextConstraintsStrategy = &mustRunAs{}

func NewMustRunAs(options *securityapi.SELinuxContextStrategyOptions) (SELinuxSecurityContextConstraintsStrategy, error) {
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
		return nil, fmt.Errorf("MustRunAs requires SELinuxContextStrategyOptions")
	}
	if options.SELinuxOptions == nil {
		return nil, fmt.Errorf("MustRunAs requires SELinuxOptions")
	}
	return &mustRunAs{opts: options}, nil
}
func (s *mustRunAs) Generate(_ *api.Pod, _ *api.Container) (*api.SELinuxOptions, error) {
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
	return s.opts.SELinuxOptions, nil
}
func (s *mustRunAs) Validate(fldPath *field.Path, _ *api.Pod, _ *api.Container, seLinux *api.SELinuxOptions) field.ErrorList {
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
	if seLinux == nil {
		allErrs = append(allErrs, field.Required(fldPath, ""))
		return allErrs
	}
	if !equalLevels(s.opts.SELinuxOptions.Level, seLinux.Level) {
		detail := fmt.Sprintf("must be %s", s.opts.SELinuxOptions.Level)
		allErrs = append(allErrs, field.Invalid(fldPath.Child("level"), seLinux.Level, detail))
	}
	if seLinux.Role != s.opts.SELinuxOptions.Role {
		detail := fmt.Sprintf("must be %s", s.opts.SELinuxOptions.Role)
		allErrs = append(allErrs, field.Invalid(fldPath.Child("role"), seLinux.Role, detail))
	}
	if seLinux.Type != s.opts.SELinuxOptions.Type {
		detail := fmt.Sprintf("must be %s", s.opts.SELinuxOptions.Type)
		allErrs = append(allErrs, field.Invalid(fldPath.Child("type"), seLinux.Type, detail))
	}
	if seLinux.User != s.opts.SELinuxOptions.User {
		detail := fmt.Sprintf("must be %s", s.opts.SELinuxOptions.User)
		allErrs = append(allErrs, field.Invalid(fldPath.Child("user"), seLinux.User, detail))
	}
	return allErrs
}
func equalLevels(expected, actual string) bool {
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
	if expected == actual {
		return true
	}
	expectedParts := strings.SplitN(expected, ":", 2)
	actualParts := strings.SplitN(actual, ":", 2)
	if len(expectedParts) != 2 || len(actualParts) != 2 {
		return false
	}
	if !equalSensitivity(expectedParts[0], actualParts[0]) {
		return false
	}
	if !equalCategories(expectedParts[1], actualParts[1]) {
		return false
	}
	return true
}
func equalSensitivity(expected, actual string) bool {
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
	return expected == actual
}
func equalCategories(expected, actual string) bool {
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
	expectedCategories := strings.Split(expected, ",")
	actualCategories := strings.Split(actual, ",")
	sort.Strings(expectedCategories)
	sort.Strings(actualCategories)
	return util.EqualStringSlices(expectedCategories, actualCategories)
}
