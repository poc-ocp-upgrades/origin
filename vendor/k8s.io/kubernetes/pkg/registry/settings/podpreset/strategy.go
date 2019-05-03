package podpreset

import (
 "context"
 "k8s.io/apimachinery/pkg/runtime"
 "k8s.io/apimachinery/pkg/util/validation/field"
 "k8s.io/apiserver/pkg/storage/names"
 "k8s.io/kubernetes/pkg/api/legacyscheme"
 "k8s.io/kubernetes/pkg/apis/settings"
 "k8s.io/kubernetes/pkg/apis/settings/validation"
)

type podPresetStrategy struct {
 runtime.ObjectTyper
 names.NameGenerator
}

var Strategy = podPresetStrategy{legacyscheme.Scheme, names.SimpleNameGenerator}

func (podPresetStrategy) NamespaceScoped() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return true
}
func (podPresetStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 pip := obj.(*settings.PodPreset)
 pip.Generation = 1
}
func (podPresetStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 newPodPreset := obj.(*settings.PodPreset)
 oldPodPreset := old.(*settings.PodPreset)
 newPodPreset.Spec = oldPodPreset.Spec
}
func (podPresetStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 pip := obj.(*settings.PodPreset)
 return validation.ValidatePodPreset(pip)
}
func (podPresetStrategy) Canonicalize(obj runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
}
func (podPresetStrategy) AllowCreateOnUpdate() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return false
}
func (podPresetStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 validationErrorList := validation.ValidatePodPreset(obj.(*settings.PodPreset))
 updateErrorList := validation.ValidatePodPresetUpdate(obj.(*settings.PodPreset), old.(*settings.PodPreset))
 return append(validationErrorList, updateErrorList...)
}
func (podPresetStrategy) AllowUnconditionalUpdate() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return true
}
