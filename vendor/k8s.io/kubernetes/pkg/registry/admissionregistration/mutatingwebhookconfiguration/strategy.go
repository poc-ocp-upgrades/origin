package mutatingwebhookconfiguration

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

type mutatingWebhookConfigurationStrategy struct {
	runtime.ObjectTyper
	names.NameGenerator
}

var Strategy = mutatingWebhookConfigurationStrategy{legacyscheme.Scheme, names.SimpleNameGenerator}

func (mutatingWebhookConfigurationStrategy) NamespaceScoped() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return false
}
func (mutatingWebhookConfigurationStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ic := obj.(*admissionregistration.MutatingWebhookConfiguration)
	ic.Generation = 1
}
func (mutatingWebhookConfigurationStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	newIC := obj.(*admissionregistration.MutatingWebhookConfiguration)
	oldIC := old.(*admissionregistration.MutatingWebhookConfiguration)
	if !reflect.DeepEqual(oldIC.Webhooks, newIC.Webhooks) {
		newIC.Generation = oldIC.Generation + 1
	}
}
func (mutatingWebhookConfigurationStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ic := obj.(*admissionregistration.MutatingWebhookConfiguration)
	return validation.ValidateMutatingWebhookConfiguration(ic)
}
func (mutatingWebhookConfigurationStrategy) Canonicalize(obj runtime.Object) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
}
func (mutatingWebhookConfigurationStrategy) AllowCreateOnUpdate() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return false
}
func (mutatingWebhookConfigurationStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	validationErrorList := validation.ValidateMutatingWebhookConfiguration(obj.(*admissionregistration.MutatingWebhookConfiguration))
	updateErrorList := validation.ValidateMutatingWebhookConfigurationUpdate(obj.(*admissionregistration.MutatingWebhookConfiguration), old.(*admissionregistration.MutatingWebhookConfiguration))
	return append(validationErrorList, updateErrorList...)
}
func (mutatingWebhookConfigurationStrategy) AllowUnconditionalUpdate() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return false
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
