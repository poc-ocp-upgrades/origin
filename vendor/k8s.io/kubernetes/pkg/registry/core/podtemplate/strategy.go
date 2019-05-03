package podtemplate

import (
 "context"
 "k8s.io/apimachinery/pkg/runtime"
 "k8s.io/apimachinery/pkg/util/validation/field"
 "k8s.io/apiserver/pkg/storage/names"
 "k8s.io/kubernetes/pkg/api/legacyscheme"
 "k8s.io/kubernetes/pkg/api/pod"
 api "k8s.io/kubernetes/pkg/apis/core"
 "k8s.io/kubernetes/pkg/apis/core/validation"
)

type podTemplateStrategy struct {
 runtime.ObjectTyper
 names.NameGenerator
}

var Strategy = podTemplateStrategy{legacyscheme.Scheme, names.SimpleNameGenerator}

func (podTemplateStrategy) NamespaceScoped() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return true
}
func (podTemplateStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 template := obj.(*api.PodTemplate)
 pod.DropDisabledAlphaFields(&template.Template.Spec)
}
func (podTemplateStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 pod := obj.(*api.PodTemplate)
 return validation.ValidatePodTemplate(pod)
}
func (podTemplateStrategy) Canonicalize(obj runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
}
func (podTemplateStrategy) AllowCreateOnUpdate() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return false
}
func (podTemplateStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 newTemplate := obj.(*api.PodTemplate)
 oldTemplate := old.(*api.PodTemplate)
 pod.DropDisabledAlphaFields(&newTemplate.Template.Spec)
 pod.DropDisabledAlphaFields(&oldTemplate.Template.Spec)
}
func (podTemplateStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return validation.ValidatePodTemplateUpdate(obj.(*api.PodTemplate), old.(*api.PodTemplate))
}
func (podTemplateStrategy) AllowUnconditionalUpdate() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return true
}
func (podTemplateStrategy) Export(ctx context.Context, obj runtime.Object, exact bool) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return nil
}
