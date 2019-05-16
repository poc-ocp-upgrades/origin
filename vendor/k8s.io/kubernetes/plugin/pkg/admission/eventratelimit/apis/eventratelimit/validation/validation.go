package validation

import (
	goformat "fmt"
	"k8s.io/apimachinery/pkg/util/validation/field"
	eventratelimitapi "k8s.io/kubernetes/plugin/pkg/admission/eventratelimit/apis/eventratelimit"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

var limitTypes = map[eventratelimitapi.LimitType]bool{eventratelimitapi.ServerLimitType: true, eventratelimitapi.NamespaceLimitType: true, eventratelimitapi.UserLimitType: true, eventratelimitapi.SourceAndObjectLimitType: true}

func ValidateConfiguration(config *eventratelimitapi.Configuration) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	allErrs := field.ErrorList{}
	limitsPath := field.NewPath("limits")
	if len(config.Limits) == 0 {
		allErrs = append(allErrs, field.Invalid(limitsPath, config.Limits, "must not be empty"))
	}
	for i, limit := range config.Limits {
		idxPath := limitsPath.Index(i)
		if !limitTypes[limit.Type] {
			allowedValues := make([]string, len(limitTypes))
			i := 0
			for limitType := range limitTypes {
				allowedValues[i] = string(limitType)
				i++
			}
			allErrs = append(allErrs, field.NotSupported(idxPath.Child("type"), limit.Type, allowedValues))
		}
		if limit.Burst <= 0 {
			allErrs = append(allErrs, field.Invalid(idxPath.Child("burst"), limit.Burst, "must be positive"))
		}
		if limit.QPS <= 0 {
			allErrs = append(allErrs, field.Invalid(idxPath.Child("qps"), limit.QPS, "must be positive"))
		}
		if limit.Type != eventratelimitapi.ServerLimitType {
			if limit.CacheSize < 0 {
				allErrs = append(allErrs, field.Invalid(idxPath.Child("cacheSize"), limit.CacheSize, "must not be negative"))
			}
		}
	}
	return allErrs
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
