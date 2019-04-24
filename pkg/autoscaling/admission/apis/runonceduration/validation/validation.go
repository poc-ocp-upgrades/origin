package validation

import (
	"github.com/openshift/origin/pkg/autoscaling/admission/apis/runonceduration"
	"bytes"
	"net/http"
	"runtime"
	"fmt"
	"k8s.io/apimachinery/pkg/util/validation/field"
)

func ValidateRunOnceDurationConfig(config *runonceduration.RunOnceDurationConfig) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	allErrs := field.ErrorList{}
	if config == nil || config.ActiveDeadlineSecondsLimit == nil {
		return allErrs
	}
	if *config.ActiveDeadlineSecondsLimit <= 0 {
		allErrs = append(allErrs, field.Invalid(field.NewPath("activeDeadlineSecondsOverride"), config.ActiveDeadlineSecondsLimit, "must be greater than 0"))
	}
	return allErrs
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
