package deployment

import (
 "context"
 appsv1beta1 "k8s.io/api/apps/v1beta1"
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

type deploymentStrategy struct {
 runtime.ObjectTyper
 names.NameGenerator
}

var Strategy = deploymentStrategy{legacyscheme.Scheme, names.SimpleNameGenerator}

func (deploymentStrategy) DefaultGarbageCollectionPolicy(ctx context.Context) rest.GarbageCollectionPolicy {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if requestInfo, found := genericapirequest.RequestInfoFrom(ctx); found {
  groupVersion := schema.GroupVersion{Group: requestInfo.APIGroup, Version: requestInfo.APIVersion}
  switch groupVersion {
  case extensionsv1beta1.SchemeGroupVersion, appsv1beta1.SchemeGroupVersion, appsv1beta2.SchemeGroupVersion:
   return rest.OrphanDependents
  default:
   return rest.DeleteDependents
  }
 }
 return rest.OrphanDependents
}
func (deploymentStrategy) NamespaceScoped() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return true
}
func (deploymentStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 deployment := obj.(*apps.Deployment)
 deployment.Status = apps.DeploymentStatus{}
 deployment.Generation = 1
 pod.DropDisabledAlphaFields(&deployment.Spec.Template.Spec)
}
func (deploymentStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 deployment := obj.(*apps.Deployment)
 return validation.ValidateDeployment(deployment)
}
func (deploymentStrategy) Canonicalize(obj runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
}
func (deploymentStrategy) AllowCreateOnUpdate() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return false
}
func (deploymentStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 newDeployment := obj.(*apps.Deployment)
 oldDeployment := old.(*apps.Deployment)
 newDeployment.Status = oldDeployment.Status
 pod.DropDisabledAlphaFields(&newDeployment.Spec.Template.Spec)
 pod.DropDisabledAlphaFields(&oldDeployment.Spec.Template.Spec)
 if !apiequality.Semantic.DeepEqual(newDeployment.Spec, oldDeployment.Spec) || !apiequality.Semantic.DeepEqual(newDeployment.Annotations, oldDeployment.Annotations) {
  newDeployment.Generation = oldDeployment.Generation + 1
 }
}
func (deploymentStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 newDeployment := obj.(*apps.Deployment)
 oldDeployment := old.(*apps.Deployment)
 allErrs := validation.ValidateDeploymentUpdate(newDeployment, oldDeployment)
 if requestInfo, found := genericapirequest.RequestInfoFrom(ctx); found {
  groupVersion := schema.GroupVersion{Group: requestInfo.APIGroup, Version: requestInfo.APIVersion}
  switch groupVersion {
  case appsv1beta1.SchemeGroupVersion, extensionsv1beta1.SchemeGroupVersion:
  default:
   allErrs = append(allErrs, apivalidation.ValidateImmutableField(newDeployment.Spec.Selector, oldDeployment.Spec.Selector, field.NewPath("spec").Child("selector"))...)
  }
 }
 return allErrs
}
func (deploymentStrategy) AllowUnconditionalUpdate() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return true
}

type deploymentStatusStrategy struct{ deploymentStrategy }

var StatusStrategy = deploymentStatusStrategy{Strategy}

func (deploymentStatusStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 newDeployment := obj.(*apps.Deployment)
 oldDeployment := old.(*apps.Deployment)
 newDeployment.Spec = oldDeployment.Spec
 newDeployment.Labels = oldDeployment.Labels
}
func (deploymentStatusStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return validation.ValidateDeploymentStatusUpdate(obj.(*apps.Deployment), old.(*apps.Deployment))
}
