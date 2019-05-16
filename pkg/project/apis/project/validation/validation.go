package validation

import (
	goformat "fmt"
	oapi "github.com/openshift/origin/pkg/api"
	projectapi "github.com/openshift/origin/pkg/project/apis/project"
	"github.com/openshift/origin/pkg/util/labelselector"
	"k8s.io/apimachinery/pkg/api/validation/path"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/kubernetes/pkg/apis/core/validation"
	goos "os"
	"reflect"
	godefaultruntime "runtime"
	"strings"
	gotime "time"
)

func ValidateProjectName(name string, prefix bool) []string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if reasons := path.ValidatePathSegmentName(name, prefix); len(reasons) != 0 {
		return reasons
	}
	if len(name) < 2 {
		return []string{"must be at least 2 characters long"}
	}
	if reasons := validation.ValidateNamespaceName(name, false); len(reasons) != 0 {
		return reasons
	}
	return nil
}
func ValidateProject(project *projectapi.Project) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	result := validation.ValidateObjectMeta(&project.ObjectMeta, false, ValidateProjectName, field.NewPath("metadata"))
	if !validateNoNewLineOrTab(project.Annotations[oapi.OpenShiftDisplayName]) {
		result = append(result, field.Invalid(field.NewPath("metadata", "annotations").Key(oapi.OpenShiftDisplayName), project.Annotations[oapi.OpenShiftDisplayName], "may not contain a new line or tab"))
	}
	result = append(result, validateNodeSelector(project)...)
	return result
}
func validateNoNewLineOrTab(s string) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return !(strings.Contains(s, "\n") || strings.Contains(s, "\t"))
}
func ValidateProjectUpdate(newProject *projectapi.Project, oldProject *projectapi.Project) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	allErrs := validation.ValidateObjectMetaUpdate(&newProject.ObjectMeta, &oldProject.ObjectMeta, field.NewPath("metadata"))
	allErrs = append(allErrs, ValidateProject(newProject)...)
	if !reflect.DeepEqual(newProject.Spec.Finalizers, oldProject.Spec.Finalizers) {
		allErrs = append(allErrs, field.Invalid(field.NewPath("spec", "finalizers"), oldProject.Spec.Finalizers, "field is immutable"))
	}
	if !reflect.DeepEqual(newProject.Status, oldProject.Status) {
		allErrs = append(allErrs, field.Invalid(field.NewPath("status"), oldProject.Spec.Finalizers, "field is immutable"))
	}
	for name, value := range newProject.Annotations {
		if name == oapi.OpenShiftDisplayName || name == oapi.OpenShiftDescription {
			continue
		}
		if value != oldProject.Annotations[name] {
			allErrs = append(allErrs, field.Invalid(field.NewPath("metadata", "annotations").Key(name), value, "field is immutable, try updating the namespace"))
		}
	}
	for name, value := range oldProject.Annotations {
		if name == oapi.OpenShiftDisplayName || name == oapi.OpenShiftDescription {
			continue
		}
		if _, inNew := newProject.Annotations[name]; !inNew {
			allErrs = append(allErrs, field.Invalid(field.NewPath("metadata", "annotations").Key(name), value, "field is immutable, try updating the namespace"))
		}
	}
	for name, value := range newProject.Labels {
		if value != oldProject.Labels[name] {
			allErrs = append(allErrs, field.Invalid(field.NewPath("metadata", "labels").Key(name), value, "field is immutable, , try updating the namespace"))
		}
	}
	for name, value := range oldProject.Labels {
		if _, inNew := newProject.Labels[name]; !inNew {
			allErrs = append(allErrs, field.Invalid(field.NewPath("metadata", "labels").Key(name), value, "field is immutable, try updating the namespace"))
		}
	}
	return allErrs
}
func ValidateProjectRequest(request *projectapi.ProjectRequest) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	project := &projectapi.Project{}
	project.ObjectMeta = request.ObjectMeta
	return ValidateProject(project)
}
func validateNodeSelector(p *projectapi.Project) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	allErrs := field.ErrorList{}
	if len(p.Annotations) > 0 {
		if selector, ok := p.Annotations[projectapi.ProjectNodeSelector]; ok {
			if _, err := labelselector.Parse(selector); err != nil {
				allErrs = append(allErrs, field.Invalid(field.NewPath("nodeSelector"), p.Annotations[projectapi.ProjectNodeSelector], "must be a valid label selector"))
			}
		}
	}
	return allErrs
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
