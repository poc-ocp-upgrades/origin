package group

import (
	"fmt"
	goformat "fmt"
	policy "k8s.io/api/policy/v1beta1"
	"k8s.io/apimachinery/pkg/util/validation/field"
	psputil "k8s.io/kubernetes/pkg/security/podsecuritypolicy/util"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

func ValidateGroupsInRanges(fldPath *field.Path, ranges []policy.IDRange, groups []int64) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	allErrs := field.ErrorList{}
	for _, group := range groups {
		if !isGroupInRanges(group, ranges) {
			detail := fmt.Sprintf("group %d must be in the ranges: %v", group, ranges)
			allErrs = append(allErrs, field.Invalid(fldPath, groups, detail))
		}
	}
	return allErrs
}
func isGroupInRanges(group int64, ranges []policy.IDRange) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for _, rng := range ranges {
		if psputil.GroupFallsInRange(group, rng) {
			return true
		}
	}
	return false
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
