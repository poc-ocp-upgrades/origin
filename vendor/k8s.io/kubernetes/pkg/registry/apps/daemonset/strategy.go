package daemonset

import (
 "context"
 appsv1beta2 "k8s.io/api/apps/v1beta2"
 extensionsv1beta1 "k8s.io/api/extensions/v1beta1"
 apiequality "k8s.io/apimachinery/pkg/api/equality"
 apivalidation "k8s.io/apimachinery/pkg/api/validation"
 "k8s.io/apimachinery/pkg/runtime"
 "k8s.io/apimachinery/pkg/runtime/schema"
 "k8s.io/apimachinery/pkg/util/validation/field"
 genericapirequest "k8s.io/apiserver/pkg/endpoints/request"
 "k8s.io/apiserver/pkg/registry/rest"
 "k8s.io/apiserver/pkg/storage/names"
 "k8s.io/kubernetes/pkg/api/legacyscheme"
 "k8s.io/kubernetes/pkg/api/pod"
 "k8s.io/kubernetes/pkg/apis/apps"
 "k8s.io/kubernetes/pkg/apis/apps/validation"
)

type daemonSetStrategy struct {
 runtime.ObjectTyper
 names.NameGenerator
}

var Strategy = daemonSetStrategy{legacyscheme.Scheme, names.SimpleNameGenerator}

func (daemonSetStrategy) DefaultGarbageCollectionPolicy(ctx context.Context) rest.GarbageCollectionPolicy {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if requestInfo, found := genericapirequest.RequestInfoFrom(ctx); found {
  groupVersion := schema.GroupVersion{Group: requestInfo.APIGroup, Version: requestInfo.APIVersion}
  switch groupVersion {
  case extensionsv1beta1.SchemeGroupVersion, appsv1beta2.SchemeGroupVersion:
   return rest.OrphanDependents
  default:
   return rest.DeleteDependents
  }
 }
 return rest.OrphanDependents
}
func (daemonSetStrategy) NamespaceScoped() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return true
}
func (daemonSetStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 daemonSet := obj.(*apps.DaemonSet)
 daemonSet.Status = apps.DaemonSetStatus{}
 daemonSet.Generation = 1
 if daemonSet.Spec.TemplateGeneration < 1 {
  daemonSet.Spec.TemplateGeneration = 1
 }
 pod.DropDisabledAlphaFields(&daemonSet.Spec.Template.Spec)
}
func (daemonSetStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 newDaemonSet := obj.(*apps.DaemonSet)
 oldDaemonSet := old.(*apps.DaemonSet)
 pod.DropDisabledAlphaFields(&newDaemonSet.Spec.Template.Spec)
 pod.DropDisabledAlphaFields(&oldDaemonSet.Spec.Template.Spec)
 newDaemonSet.Status = oldDaemonSet.Status
 newDaemonSet.Spec.TemplateGeneration = oldDaemonSet.Spec.TemplateGeneration
 if !apiequality.Semantic.DeepEqual(oldDaemonSet.Spec.Template, newDaemonSet.Spec.Template) {
  newDaemonSet.Spec.TemplateGeneration = oldDaemonSet.Spec.TemplateGeneration + 1
  newDaemonSet.Generation = oldDaemonSet.Generation + 1
  return
 }
 if !apiequality.Semantic.DeepEqual(oldDaemonSet.Spec, newDaemonSet.Spec) {
  newDaemonSet.Generation = oldDaemonSet.Generation + 1
 }
}
func (daemonSetStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 daemonSet := obj.(*apps.DaemonSet)
 return validation.ValidateDaemonSet(daemonSet)
}
func (daemonSetStrategy) Canonicalize(obj runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
}
func (daemonSetStrategy) AllowCreateOnUpdate() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return false
}
func (daemonSetStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 newDaemonSet := obj.(*apps.DaemonSet)
 oldDaemonSet := old.(*apps.DaemonSet)
 allErrs := validation.ValidateDaemonSet(obj.(*apps.DaemonSet))
 allErrs = append(allErrs, validation.ValidateDaemonSetUpdate(newDaemonSet, oldDaemonSet)...)
 if requestInfo, found := genericapirequest.RequestInfoFrom(ctx); found {
  groupVersion := schema.GroupVersion{Group: requestInfo.APIGroup, Version: requestInfo.APIVersion}
  switch groupVersion {
  case extensionsv1beta1.SchemeGroupVersion:
  default:
   allErrs = append(allErrs, apivalidation.ValidateImmutableField(newDaemonSet.Spec.Selector, oldDaemonSet.Spec.Selector, field.NewPath("spec").Child("selector"))...)
  }
 }
 return allErrs
}
func (daemonSetStrategy) AllowUnconditionalUpdate() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return true
}

type daemonSetStatusStrategy struct{ daemonSetStrategy }

var StatusStrategy = daemonSetStatusStrategy{Strategy}

func (daemonSetStatusStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 newDaemonSet := obj.(*apps.DaemonSet)
 oldDaemonSet := old.(*apps.DaemonSet)
 newDaemonSet.Spec = oldDaemonSet.Spec
}
func (daemonSetStatusStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return validation.ValidateDaemonSetStatusUpdate(obj.(*apps.DaemonSet), old.(*apps.DaemonSet))
}
