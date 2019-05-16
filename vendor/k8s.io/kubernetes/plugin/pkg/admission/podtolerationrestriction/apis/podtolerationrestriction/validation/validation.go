package validation

import (
	"fmt"
	goformat "fmt"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/kubernetes/pkg/apis/core/validation"
	internalapi "k8s.io/kubernetes/plugin/pkg/admission/podtolerationrestriction/apis/podtolerationrestriction"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

func ValidateConfiguration(config *internalapi.Configuration) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	allErrs := field.ErrorList{}
	fldpath := field.NewPath("podtolerationrestriction")
	allErrs = append(allErrs, validation.ValidateTolerations(config.Default, fldpath.Child("default"))...)
	allErrs = append(allErrs, validation.ValidateTolerations(config.Whitelist, fldpath.Child("whitelist"))...)
	if len(allErrs) > 0 {
		return fmt.Errorf("invalid config: %v", allErrs)
	}
	return nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
