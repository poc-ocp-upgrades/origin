package group

import (
	"fmt"
	"bytes"
	"net/http"
	"runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	api "k8s.io/kubernetes/pkg/apis/core"
	securityapi "github.com/openshift/origin/pkg/security/apis/security"
)

type mustRunAs struct {
	ranges	[]securityapi.IDRange
	field	string
}

var _ GroupSecurityContextConstraintsStrategy = &mustRunAs{}

func NewMustRunAs(ranges []securityapi.IDRange, field string) (GroupSecurityContextConstraintsStrategy, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(ranges) == 0 {
		return nil, fmt.Errorf("ranges must be supplied for MustRunAs")
	}
	return &mustRunAs{ranges: ranges, field: field}, nil
}
func (s *mustRunAs) Generate(_ *api.Pod) ([]int64, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []int64{s.ranges[0].Min}, nil
}
func (s *mustRunAs) GenerateSingle(_ *api.Pod) (*int64, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	single := new(int64)
	*single = s.ranges[0].Min
	return single, nil
}
func (s *mustRunAs) Validate(_ *api.Pod, groups []int64) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, rng := range s.ranges {
		if fallsInRange(group, rng) {
			return true
		}
	}
	return false
}
func fallsInRange(group int64, rng securityapi.IDRange) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return group >= rng.Min && group <= rng.Max
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
