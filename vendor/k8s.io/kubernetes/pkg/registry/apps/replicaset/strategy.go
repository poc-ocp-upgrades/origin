package replicaset

import (
 "context"
 "fmt"
 "strconv"
 appsv1beta2 "k8s.io/api/apps/v1beta2"
 extensionsv1beta1 "k8s.io/api/extensions/v1beta1"
 apiequality "k8s.io/apimachinery/pkg/api/equality"
 apivalidation "k8s.io/apimachinery/pkg/api/validation"
 "k8s.io/apimachinery/pkg/fields"
 "k8s.io/apimachinery/pkg/labels"
 "k8s.io/apimachinery/pkg/runtime"
 "k8s.io/apimachinery/pkg/runtime/schema"
 "k8s.io/apimachinery/pkg/util/validation/field"
 genericapirequest "k8s.io/apiserver/pkg/endpoints/request"
 "k8s.io/apiserver/pkg/registry/generic"
 "k8s.io/apiserver/pkg/registry/rest"
 apistorage "k8s.io/apiserver/pkg/storage"
 "k8s.io/apiserver/pkg/storage/names"
 "k8s.io/kubernetes/pkg/api/legacyscheme"
 "k8s.io/kubernetes/pkg/api/pod"
 "k8s.io/kubernetes/pkg/apis/apps"
 "k8s.io/kubernetes/pkg/apis/apps/validation"
)

type rsStrategy struct {
 runtime.ObjectTyper
 names.NameGenerator
}

var Strategy = rsStrategy{legacyscheme.Scheme, names.SimpleNameGenerator}

func (rsStrategy) DefaultGarbageCollectionPolicy(ctx context.Context) rest.GarbageCollectionPolicy {
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
func (rsStrategy) NamespaceScoped() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return true
}
func (rsStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 rs := obj.(*apps.ReplicaSet)
 rs.Status = apps.ReplicaSetStatus{}
 rs.Generation = 1
 pod.DropDisabledAlphaFields(&rs.Spec.Template.Spec)
}
func (rsStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 newRS := obj.(*apps.ReplicaSet)
 oldRS := old.(*apps.ReplicaSet)
 newRS.Status = oldRS.Status
 pod.DropDisabledAlphaFields(&newRS.Spec.Template.Spec)
 pod.DropDisabledAlphaFields(&oldRS.Spec.Template.Spec)
 if !apiequality.Semantic.DeepEqual(oldRS.Spec, newRS.Spec) {
  newRS.Generation = oldRS.Generation + 1
 }
}
func (rsStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 rs := obj.(*apps.ReplicaSet)
 return validation.ValidateReplicaSet(rs)
}
func (rsStrategy) Canonicalize(obj runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
}
func (rsStrategy) AllowCreateOnUpdate() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return false
}
func (rsStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 newReplicaSet := obj.(*apps.ReplicaSet)
 oldReplicaSet := old.(*apps.ReplicaSet)
 allErrs := validation.ValidateReplicaSet(obj.(*apps.ReplicaSet))
 allErrs = append(allErrs, validation.ValidateReplicaSetUpdate(newReplicaSet, oldReplicaSet)...)
 if requestInfo, found := genericapirequest.RequestInfoFrom(ctx); found {
  groupVersion := schema.GroupVersion{Group: requestInfo.APIGroup, Version: requestInfo.APIVersion}
  switch groupVersion {
  case extensionsv1beta1.SchemeGroupVersion:
  default:
   allErrs = append(allErrs, apivalidation.ValidateImmutableField(newReplicaSet.Spec.Selector, oldReplicaSet.Spec.Selector, field.NewPath("spec").Child("selector"))...)
  }
 }
 return allErrs
}
func (rsStrategy) AllowUnconditionalUpdate() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return true
}
func ReplicaSetToSelectableFields(rs *apps.ReplicaSet) fields.Set {
 _logClusterCodePath()
 defer _logClusterCodePath()
 objectMetaFieldsSet := generic.ObjectMetaFieldsSet(&rs.ObjectMeta, true)
 rsSpecificFieldsSet := fields.Set{"status.replicas": strconv.Itoa(int(rs.Status.Replicas))}
 return generic.MergeFieldsSets(objectMetaFieldsSet, rsSpecificFieldsSet)
}
func GetAttrs(obj runtime.Object) (labels.Set, fields.Set, bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 rs, ok := obj.(*apps.ReplicaSet)
 if !ok {
  return nil, nil, false, fmt.Errorf("given object is not a ReplicaSet.")
 }
 return labels.Set(rs.ObjectMeta.Labels), ReplicaSetToSelectableFields(rs), rs.Initializers != nil, nil
}
func MatchReplicaSet(label labels.Selector, field fields.Selector) apistorage.SelectionPredicate {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return apistorage.SelectionPredicate{Label: label, Field: field, GetAttrs: GetAttrs}
}

type rsStatusStrategy struct{ rsStrategy }

var StatusStrategy = rsStatusStrategy{Strategy}

func (rsStatusStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 newRS := obj.(*apps.ReplicaSet)
 oldRS := old.(*apps.ReplicaSet)
 newRS.Spec = oldRS.Spec
}
func (rsStatusStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return validation.ValidateReplicaSetStatusUpdate(obj.(*apps.ReplicaSet), old.(*apps.ReplicaSet))
}
