package mutatingwebhookconfiguration

import (
 "context"
 "reflect"
 "k8s.io/apimachinery/pkg/runtime"
 "k8s.io/apimachinery/pkg/util/validation/field"
 "k8s.io/apiserver/pkg/storage/names"
 "k8s.io/kubernetes/pkg/api/legacyscheme"
 "k8s.io/kubernetes/pkg/apis/admissionregistration"
 "k8s.io/kubernetes/pkg/apis/admissionregistration/validation"
)

type mutatingWebhookConfigurationStrategy struct {
 runtime.ObjectTyper
 names.NameGenerator
}

var Strategy = mutatingWebhookConfigurationStrategy{legacyscheme.Scheme, names.SimpleNameGenerator}

func (mutatingWebhookConfigurationStrategy) NamespaceScoped() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return false
}
func (mutatingWebhookConfigurationStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ic := obj.(*admissionregistration.MutatingWebhookConfiguration)
 ic.Generation = 1
}
func (mutatingWebhookConfigurationStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 newIC := obj.(*admissionregistration.MutatingWebhookConfiguration)
 oldIC := old.(*admissionregistration.MutatingWebhookConfiguration)
 if !reflect.DeepEqual(oldIC.Webhooks, newIC.Webhooks) {
  newIC.Generation = oldIC.Generation + 1
 }
}
func (mutatingWebhookConfigurationStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ic := obj.(*admissionregistration.MutatingWebhookConfiguration)
 return validation.ValidateMutatingWebhookConfiguration(ic)
}
func (mutatingWebhookConfigurationStrategy) Canonicalize(obj runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
}
func (mutatingWebhookConfigurationStrategy) AllowCreateOnUpdate() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return false
}
func (mutatingWebhookConfigurationStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 validationErrorList := validation.ValidateMutatingWebhookConfiguration(obj.(*admissionregistration.MutatingWebhookConfiguration))
 updateErrorList := validation.ValidateMutatingWebhookConfigurationUpdate(obj.(*admissionregistration.MutatingWebhookConfiguration), old.(*admissionregistration.MutatingWebhookConfiguration))
 return append(validationErrorList, updateErrorList...)
}
func (mutatingWebhookConfigurationStrategy) AllowUnconditionalUpdate() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return false
}
