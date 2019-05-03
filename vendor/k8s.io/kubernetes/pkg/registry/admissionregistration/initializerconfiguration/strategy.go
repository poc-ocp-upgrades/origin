package initializerconfiguration

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

type initializerConfigurationStrategy struct {
 runtime.ObjectTyper
 names.NameGenerator
}

var Strategy = initializerConfigurationStrategy{legacyscheme.Scheme, names.SimpleNameGenerator}

func (initializerConfigurationStrategy) NamespaceScoped() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return false
}
func (initializerConfigurationStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ic := obj.(*admissionregistration.InitializerConfiguration)
 ic.Generation = 1
}
func (initializerConfigurationStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 newIC := obj.(*admissionregistration.InitializerConfiguration)
 oldIC := old.(*admissionregistration.InitializerConfiguration)
 if !reflect.DeepEqual(oldIC.Initializers, newIC.Initializers) {
  newIC.Generation = oldIC.Generation + 1
 }
}
func (initializerConfigurationStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ic := obj.(*admissionregistration.InitializerConfiguration)
 return validation.ValidateInitializerConfiguration(ic)
}
func (initializerConfigurationStrategy) Canonicalize(obj runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
}
func (initializerConfigurationStrategy) AllowCreateOnUpdate() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return false
}
func (initializerConfigurationStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 validationErrorList := validation.ValidateInitializerConfiguration(obj.(*admissionregistration.InitializerConfiguration))
 updateErrorList := validation.ValidateInitializerConfigurationUpdate(obj.(*admissionregistration.InitializerConfiguration), old.(*admissionregistration.InitializerConfiguration))
 return append(validationErrorList, updateErrorList...)
}
func (initializerConfigurationStrategy) AllowUnconditionalUpdate() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return false
}
