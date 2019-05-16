package scheduler

import (
	"fmt"
	goformat "fmt"
	configv1 "github.com/openshift/api/config/v1"
	"github.com/openshift/origin/pkg/admission/customresourcevalidation"
	"io"
	"k8s.io/apimachinery/pkg/api/validation"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/admission"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

const PluginName = "config.openshift.io/ValidateScheduler"

func Register(plugins *admission.Plugins) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	plugins.Register(PluginName, func(config io.Reader) (admission.Interface, error) {
		return customresourcevalidation.NewValidator(map[schema.GroupResource]bool{configv1.Resource("schedulers"): true}, map[schema.GroupVersionKind]customresourcevalidation.ObjectValidator{configv1.GroupVersion.WithKind("Scheduler"): schedulerV1{}})
	})
}
func toSchedulerV1(uncastObj runtime.Object) (*configv1.Scheduler, field.ErrorList) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if uncastObj == nil {
		return nil, nil
	}
	allErrs := field.ErrorList{}
	obj, ok := uncastObj.(*configv1.Scheduler)
	if !ok {
		return nil, append(allErrs, field.NotSupported(field.NewPath("kind"), fmt.Sprintf("%T", uncastObj), []string{"Scheduler"}), field.NotSupported(field.NewPath("apiVersion"), fmt.Sprintf("%T", uncastObj), []string{"config.openshift.io/v1"}))
	}
	return obj, nil
}

type schedulerV1 struct{}

func validateSchedulerSpec(spec configv1.SchedulerSpec) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	allErrs := field.ErrorList{}
	if name := spec.Policy.Name; len(name) > 0 {
		for _, msg := range validation.NameIsDNSSubdomain(spec.Policy.Name, false) {
			allErrs = append(allErrs, field.Invalid(field.NewPath("spec.Policy.name"), name, msg))
		}
	}
	return allErrs
}
func (schedulerV1) ValidateCreate(uncastObj runtime.Object) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	obj, allErrs := toSchedulerV1(uncastObj)
	if len(allErrs) > 0 {
		return allErrs
	}
	allErrs = append(allErrs, validation.ValidateObjectMeta(&obj.ObjectMeta, false, customresourcevalidation.RequireNameCluster, field.NewPath("metadata"))...)
	allErrs = append(allErrs, validateSchedulerSpec(obj.Spec)...)
	return allErrs
}
func (schedulerV1) ValidateUpdate(uncastObj runtime.Object, uncastOldObj runtime.Object) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	obj, allErrs := toSchedulerV1(uncastObj)
	if len(allErrs) > 0 {
		return allErrs
	}
	oldObj, allErrs := toSchedulerV1(uncastOldObj)
	if len(allErrs) > 0 {
		return allErrs
	}
	allErrs = append(allErrs, validation.ValidateObjectMetaUpdate(&obj.ObjectMeta, &oldObj.ObjectMeta, field.NewPath("metadata"))...)
	allErrs = append(allErrs, validateSchedulerSpec(obj.Spec)...)
	return allErrs
}
func (schedulerV1) ValidateStatusUpdate(uncastObj runtime.Object, uncastOldObj runtime.Object) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	obj, errs := toSchedulerV1(uncastObj)
	if len(errs) > 0 {
		return errs
	}
	oldObj, errs := toSchedulerV1(uncastOldObj)
	if len(errs) > 0 {
		return errs
	}
	errs = append(errs, validation.ValidateObjectMetaUpdate(&obj.ObjectMeta, &oldObj.ObjectMeta, field.NewPath("metadata"))...)
	return errs
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
