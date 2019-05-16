package group

import (
	"fmt"
	policy "k8s.io/api/policy/v1beta1"
	"k8s.io/apimachinery/pkg/util/validation/field"
	api "k8s.io/kubernetes/pkg/apis/core"
)

type mustRunAs struct{ ranges []policy.IDRange }

var _ GroupStrategy = &mustRunAs{}

func NewMustRunAs(ranges []policy.IDRange) (GroupStrategy, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(ranges) == 0 {
		return nil, fmt.Errorf("ranges must be supplied for MustRunAs")
	}
	return &mustRunAs{ranges: ranges}, nil
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
func (s *mustRunAs) Validate(fldPath *field.Path, _ *api.Pod, groups []int64) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	allErrs := field.ErrorList{}
	if len(groups) == 0 && len(s.ranges) > 0 {
		allErrs = append(allErrs, field.Invalid(fldPath, groups, "unable to validate empty groups against required ranges"))
	}
	allErrs = append(allErrs, ValidateGroupsInRanges(fldPath, s.ranges, groups)...)
	return allErrs
}
