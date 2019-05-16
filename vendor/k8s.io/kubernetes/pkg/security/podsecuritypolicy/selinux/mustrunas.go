package selinux

import (
	"fmt"
	goformat "fmt"
	policy "k8s.io/api/policy/v1beta1"
	"k8s.io/apimachinery/pkg/util/validation/field"
	api "k8s.io/kubernetes/pkg/apis/core"
	"k8s.io/kubernetes/pkg/apis/core/v1"
	"k8s.io/kubernetes/pkg/security/podsecuritypolicy/util"
	goos "os"
	godefaultruntime "runtime"
	"sort"
	"strings"
	gotime "time"
)

type mustRunAs struct{ opts *api.SELinuxOptions }

var _ SELinuxStrategy = &mustRunAs{}

func NewMustRunAs(options *policy.SELinuxStrategyOptions) (SELinuxStrategy, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if options == nil {
		return nil, fmt.Errorf("MustRunAs requires SELinuxContextStrategyOptions")
	}
	if options.SELinuxOptions == nil {
		return nil, fmt.Errorf("MustRunAs requires SELinuxOptions")
	}
	internalSELinuxOptions := &api.SELinuxOptions{}
	if err := v1.Convert_v1_SELinuxOptions_To_core_SELinuxOptions(options.SELinuxOptions, internalSELinuxOptions, nil); err != nil {
		return nil, err
	}
	return &mustRunAs{opts: internalSELinuxOptions}, nil
}
func (s *mustRunAs) Generate(_ *api.Pod, _ *api.Container) (*api.SELinuxOptions, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return s.opts, nil
}
func (s *mustRunAs) Validate(fldPath *field.Path, _ *api.Pod, _ *api.Container, seLinux *api.SELinuxOptions) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	allErrs := field.ErrorList{}
	if seLinux == nil {
		allErrs = append(allErrs, field.Required(fldPath, ""))
		return allErrs
	}
	if !equalLevels(s.opts.Level, seLinux.Level) {
		detail := fmt.Sprintf("must be %s", s.opts.Level)
		allErrs = append(allErrs, field.Invalid(fldPath.Child("level"), seLinux.Level, detail))
	}
	if seLinux.Role != s.opts.Role {
		detail := fmt.Sprintf("must be %s", s.opts.Role)
		allErrs = append(allErrs, field.Invalid(fldPath.Child("role"), seLinux.Role, detail))
	}
	if seLinux.Type != s.opts.Type {
		detail := fmt.Sprintf("must be %s", s.opts.Type)
		allErrs = append(allErrs, field.Invalid(fldPath.Child("type"), seLinux.Type, detail))
	}
	if seLinux.User != s.opts.User {
		detail := fmt.Sprintf("must be %s", s.opts.User)
		allErrs = append(allErrs, field.Invalid(fldPath.Child("user"), seLinux.User, detail))
	}
	return allErrs
}
func equalLevels(expected, actual string) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
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
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return expected == actual
}
func equalCategories(expected, actual string) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	expectedCategories := strings.Split(expected, ",")
	actualCategories := strings.Split(actual, ",")
	sort.Strings(expectedCategories)
	sort.Strings(actualCategories)
	return util.EqualStringSlices(expectedCategories, actualCategories)
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
