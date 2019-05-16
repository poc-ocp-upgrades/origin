package group

import (
	"fmt"
	goformat "fmt"
	securityapi "github.com/openshift/origin/pkg/security/apis/security"
	"k8s.io/apimachinery/pkg/util/validation/field"
	api "k8s.io/kubernetes/pkg/apis/core"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

type mustRunAs struct {
	ranges []securityapi.IDRange
	field  string
}

var _ GroupSecurityContextConstraintsStrategy = &mustRunAs{}

func NewMustRunAs(ranges []securityapi.IDRange, field string) (GroupSecurityContextConstraintsStrategy, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(ranges) == 0 {
		return nil, fmt.Errorf("ranges must be supplied for MustRunAs")
	}
	return &mustRunAs{ranges: ranges, field: field}, nil
}
func (s *mustRunAs) Generate(_ *api.Pod) ([]int64, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return []int64{s.ranges[0].Min}, nil
}
func (s *mustRunAs) GenerateSingle(_ *api.Pod) (*int64, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	single := new(int64)
	*single = s.ranges[0].Min
	return single, nil
}
func (s *mustRunAs) Validate(_ *api.Pod, groups []int64) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	allErrs := field.ErrorList{}
	if len(groups) == 0 && len(s.ranges) > 0 {
		allErrs = append(allErrs, field.Invalid(field.NewPath(s.field), groups, "unable to validate empty groups against required ranges"))
	}
	for _, group := range groups {
		if !s.isGroupValid(group) {
			detail := fmt.Sprintf("%d is not an allowed group", group)
			allErrs = append(allErrs, field.Invalid(field.NewPath(s.field), groups, detail))
		}
	}
	return allErrs
}
func (s *mustRunAs) isGroupValid(group int64) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for _, rng := range s.ranges {
		if fallsInRange(group, rng) {
			return true
		}
	}
	return false
}
func fallsInRange(group int64, rng securityapi.IDRange) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return group >= rng.Min && group <= rng.Max
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
