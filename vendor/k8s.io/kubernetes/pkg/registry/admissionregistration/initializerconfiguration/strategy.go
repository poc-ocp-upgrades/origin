package initializerconfiguration

import (
	"context"
	goformat "fmt"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/storage/names"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	"k8s.io/kubernetes/pkg/apis/admissionregistration"
	"k8s.io/kubernetes/pkg/apis/admissionregistration/validation"
	goos "os"
	"reflect"
	godefaultruntime "runtime"
	gotime "time"
)

type initializerConfigurationStrategy struct {
	runtime.ObjectTyper
	names.NameGenerator
}

var Strategy = initializerConfigurationStrategy{legacyscheme.Scheme, names.SimpleNameGenerator}

func (initializerConfigurationStrategy) NamespaceScoped() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return false
}
func (initializerConfigurationStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ic := obj.(*admissionregistration.InitializerConfiguration)
	ic.Generation = 1
}
func (initializerConfigurationStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	newIC := obj.(*admissionregistration.InitializerConfiguration)
	oldIC := old.(*admissionregistration.InitializerConfiguration)
	if !reflect.DeepEqual(oldIC.Initializers, newIC.Initializers) {
		newIC.Generation = oldIC.Generation + 1
	}
}
func (initializerConfigurationStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ic := obj.(*admissionregistration.InitializerConfiguration)
	return validation.ValidateInitializerConfiguration(ic)
}
func (initializerConfigurationStrategy) Canonicalize(obj runtime.Object) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
}
func (initializerConfigurationStrategy) AllowCreateOnUpdate() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return false
}
func (initializerConfigurationStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	validationErrorList := validation.ValidateInitializerConfiguration(obj.(*admissionregistration.InitializerConfiguration))
	updateErrorList := validation.ValidateInitializerConfigurationUpdate(obj.(*admissionregistration.InitializerConfiguration), old.(*admissionregistration.InitializerConfiguration))
	return append(validationErrorList, updateErrorList...)
}
func (initializerConfigurationStrategy) AllowUnconditionalUpdate() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return false
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
