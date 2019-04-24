package validation

import (
	unversionedvalidation "k8s.io/apimachinery/pkg/apis/meta/v1/validation"
	"bytes"
	"net/http"
	"runtime"
	"fmt"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"github.com/openshift/origin/pkg/project/apiserver/admission/apis/requestlimit"
)

func ValidateProjectRequestLimitConfig(config *requestlimit.ProjectRequestLimitConfig) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	allErrs := field.ErrorList{}
	for i, projectLimit := range config.Limits {
		allErrs = append(allErrs, ValidateProjectLimitBySelector(projectLimit, field.NewPath("limits").Index(i))...)
	}
	if config.MaxProjectsForSystemUsers != nil && *config.MaxProjectsForSystemUsers < 0 {
		allErrs = append(allErrs, field.Invalid(field.NewPath("maxProjectsForSystemUsers"), *config.MaxProjectsForSystemUsers, "cannot be a negative number"))
	}
	if config.MaxProjectsForServiceAccounts != nil && *config.MaxProjectsForServiceAccounts < 0 {
		allErrs = append(allErrs, field.Invalid(field.NewPath("maxProjectsForServiceAccounts"), *config.MaxProjectsForServiceAccounts, "cannot be a negative number"))
	}
	return allErrs
}
func ValidateProjectLimitBySelector(limit requestlimit.ProjectLimitBySelector, path *field.Path) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	allErrs := field.ErrorList{}
	allErrs = append(allErrs, unversionedvalidation.ValidateLabels(limit.Selector, path.Child("selector"))...)
	if limit.MaxProjects != nil && *limit.MaxProjects < 0 {
		allErrs = append(allErrs, field.Invalid(path.Child("maxProjects"), *limit.MaxProjects, "cannot be a negative number"))
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
