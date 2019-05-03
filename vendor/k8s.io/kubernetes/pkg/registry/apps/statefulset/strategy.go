package statefulset

import (
 "context"
 appsv1beta1 "k8s.io/api/apps/v1beta1"
 appsv1beta2 "k8s.io/api/apps/v1beta2"
 apiequality "k8s.io/apimachinery/pkg/api/equality"
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

type statefulSetStrategy struct {
 runtime.ObjectTyper
 names.NameGenerator
}

var Strategy = statefulSetStrategy{legacyscheme.Scheme, names.SimpleNameGenerator}

func (statefulSetStrategy) DefaultGarbageCollectionPolicy(ctx context.Context) rest.GarbageCollectionPolicy {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if requestInfo, found := genericapirequest.RequestInfoFrom(ctx); found {
  groupVersion := schema.GroupVersion{Group: requestInfo.APIGroup, Version: requestInfo.APIVersion}
  switch groupVersion {
  case appsv1beta1.SchemeGroupVersion, appsv1beta2.SchemeGroupVersion:
   return rest.OrphanDependents
  default:
   return rest.DeleteDependents
  }
 }
 return rest.OrphanDependents
}
func (statefulSetStrategy) NamespaceScoped() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return true
}
func (statefulSetStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 statefulSet := obj.(*apps.StatefulSet)
 statefulSet.Status = apps.StatefulSetStatus{}
 statefulSet.Generation = 1
 pod.DropDisabledAlphaFields(&statefulSet.Spec.Template.Spec)
}
func (statefulSetStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 newStatefulSet := obj.(*apps.StatefulSet)
 oldStatefulSet := old.(*apps.StatefulSet)
 newStatefulSet.Status = oldStatefulSet.Status
 pod.DropDisabledAlphaFields(&newStatefulSet.Spec.Template.Spec)
 pod.DropDisabledAlphaFields(&oldStatefulSet.Spec.Template.Spec)
 if !apiequality.Semantic.DeepEqual(oldStatefulSet.Spec, newStatefulSet.Spec) {
  newStatefulSet.Generation = oldStatefulSet.Generation + 1
 }
}
func (statefulSetStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 statefulSet := obj.(*apps.StatefulSet)
 return validation.ValidateStatefulSet(statefulSet)
}
func (statefulSetStrategy) Canonicalize(obj runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
}
func (statefulSetStrategy) AllowCreateOnUpdate() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return false
}
func (statefulSetStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 validationErrorList := validation.ValidateStatefulSet(obj.(*apps.StatefulSet))
 updateErrorList := validation.ValidateStatefulSetUpdate(obj.(*apps.StatefulSet), old.(*apps.StatefulSet))
 return append(validationErrorList, updateErrorList...)
}
func (statefulSetStrategy) AllowUnconditionalUpdate() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return true
}

type statefulSetStatusStrategy struct{ statefulSetStrategy }

var StatusStrategy = statefulSetStatusStrategy{Strategy}

func (statefulSetStatusStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 newStatefulSet := obj.(*apps.StatefulSet)
 oldStatefulSet := old.(*apps.StatefulSet)
 newStatefulSet.Spec = oldStatefulSet.Spec
}
func (statefulSetStatusStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return validation.ValidateStatefulSetStatusUpdate(obj.(*apps.StatefulSet), old.(*apps.StatefulSet))
}
