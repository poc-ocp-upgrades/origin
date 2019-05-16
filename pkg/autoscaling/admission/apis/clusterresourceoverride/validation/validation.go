package validation

import (
	goformat "fmt"
	"github.com/openshift/origin/pkg/autoscaling/admission/apis/clusterresourceoverride"
	"k8s.io/apimachinery/pkg/util/validation/field"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

func Validate(config *clusterresourceoverride.ClusterResourceOverrideConfig) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	allErrs := field.ErrorList{}
	if config == nil {
		return allErrs
	}
	if config.LimitCPUToMemoryPercent == 0 && config.CPURequestToLimitPercent == 0 && config.MemoryRequestToLimitPercent == 0 {
		allErrs = append(allErrs, field.Forbidden(field.NewPath(clusterresourceoverride.PluginName), "plugin enabled but no percentages were specified"))
	}
	if config.LimitCPUToMemoryPercent < 0 {
		allErrs = append(allErrs, field.Invalid(field.NewPath(clusterresourceoverride.PluginName, "LimitCPUToMemoryPercent"), config.LimitCPUToMemoryPercent, "must be positive"))
	}
	if config.CPURequestToLimitPercent < 0 || config.CPURequestToLimitPercent > 100 {
		allErrs = append(allErrs, field.Invalid(field.NewPath(clusterresourceoverride.PluginName, "CPURequestToLimitPercent"), config.CPURequestToLimitPercent, "must be between 0 and 100"))
	}
	if config.MemoryRequestToLimitPercent < 0 || config.MemoryRequestToLimitPercent > 100 {
		allErrs = append(allErrs, field.Invalid(field.NewPath(clusterresourceoverride.PluginName, "MemoryRequestToLimitPercent"), config.MemoryRequestToLimitPercent, "must be between 0 and 100"))
	}
	return allErrs
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
