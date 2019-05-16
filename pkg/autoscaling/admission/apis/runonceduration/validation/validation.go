package validation

import (
	goformat "fmt"
	"github.com/openshift/origin/pkg/autoscaling/admission/apis/runonceduration"
	"k8s.io/apimachinery/pkg/util/validation/field"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

func ValidateRunOnceDurationConfig(config *runonceduration.RunOnceDurationConfig) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	allErrs := field.ErrorList{}
	if config == nil || config.ActiveDeadlineSecondsLimit == nil {
		return allErrs
	}
	if *config.ActiveDeadlineSecondsLimit <= 0 {
		allErrs = append(allErrs, field.Invalid(field.NewPath("activeDeadlineSecondsOverride"), config.ActiveDeadlineSecondsLimit, "must be greater than 0"))
	}
	return allErrs
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
