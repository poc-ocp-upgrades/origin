package validation

import (
	goformat "fmt"
	"k8s.io/apimachinery/pkg/util/validation/field"
	resourcequotaapi "k8s.io/kubernetes/plugin/pkg/admission/resourcequota/apis/resourcequota"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

func ValidateConfiguration(config *resourcequotaapi.Configuration) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	allErrs := field.ErrorList{}
	fldPath := field.NewPath("limitedResources")
	for i, limitedResource := range config.LimitedResources {
		idxPath := fldPath.Index(i)
		if len(limitedResource.Resource) == 0 {
			allErrs = append(allErrs, field.Required(idxPath.Child("resource"), ""))
		}
	}
	return allErrs
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
