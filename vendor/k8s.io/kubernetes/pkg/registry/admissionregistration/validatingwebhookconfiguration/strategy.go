package validatingwebhookconfiguration

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

type validatingWebhookConfigurationStrategy struct {
	runtime.ObjectTyper
	names.NameGenerator
}

var Strategy = validatingWebhookConfigurationStrategy{legacyscheme.Scheme, names.SimpleNameGenerator}

func (validatingWebhookConfigurationStrategy) NamespaceScoped() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return false
}
func (validatingWebhookConfigurationStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ic := obj.(*admissionregistration.ValidatingWebhookConfiguration)
	ic.Generation = 1
}
func (validatingWebhookConfigurationStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	newIC := obj.(*admissionregistration.ValidatingWebhookConfiguration)
	oldIC := old.(*admissionregistration.ValidatingWebhookConfiguration)
	if !reflect.DeepEqual(oldIC.Webhooks, newIC.Webhooks) {
		newIC.Generation = oldIC.Generation + 1
	}
}
func (validatingWebhookConfigurationStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ic := obj.(*admissionregistration.ValidatingWebhookConfiguration)
	return validation.ValidateValidatingWebhookConfiguration(ic)
}
func (validatingWebhookConfigurationStrategy) Canonicalize(obj runtime.Object) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
}
func (validatingWebhookConfigurationStrategy) AllowCreateOnUpdate() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return false
}
func (validatingWebhookConfigurationStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	validationErrorList := validation.ValidateValidatingWebhookConfiguration(obj.(*admissionregistration.ValidatingWebhookConfiguration))
	updateErrorList := validation.ValidateValidatingWebhookConfigurationUpdate(obj.(*admissionregistration.ValidatingWebhookConfiguration), old.(*admissionregistration.ValidatingWebhookConfiguration))
	return append(validationErrorList, updateErrorList...)
}
func (validatingWebhookConfigurationStrategy) AllowUnconditionalUpdate() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return false
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
