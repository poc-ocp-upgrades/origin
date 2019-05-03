package persistentvolumeclaim

import (
 "context"
 "fmt"
 "k8s.io/apimachinery/pkg/fields"
 "k8s.io/apimachinery/pkg/labels"
 "k8s.io/apimachinery/pkg/runtime"
 "k8s.io/apimachinery/pkg/util/validation/field"
 "k8s.io/apiserver/pkg/registry/generic"
 "k8s.io/apiserver/pkg/storage"
 "k8s.io/apiserver/pkg/storage/names"
 "k8s.io/kubernetes/pkg/api/legacyscheme"
 pvcutil "k8s.io/kubernetes/pkg/api/persistentvolumeclaim"
 api "k8s.io/kubernetes/pkg/apis/core"
 "k8s.io/kubernetes/pkg/apis/core/validation"
)

type persistentvolumeclaimStrategy struct {
 runtime.ObjectTyper
 names.NameGenerator
}

var Strategy = persistentvolumeclaimStrategy{legacyscheme.Scheme, names.SimpleNameGenerator}

func (persistentvolumeclaimStrategy) NamespaceScoped() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return true
}
func (persistentvolumeclaimStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 pvc := obj.(*api.PersistentVolumeClaim)
 pvc.Status = api.PersistentVolumeClaimStatus{}
 pvcutil.DropDisabledAlphaFields(&pvc.Spec)
}
func (persistentvolumeclaimStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 pvc := obj.(*api.PersistentVolumeClaim)
 return validation.ValidatePersistentVolumeClaim(pvc)
}
func (persistentvolumeclaimStrategy) Canonicalize(obj runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
}
func (persistentvolumeclaimStrategy) AllowCreateOnUpdate() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return false
}
func (persistentvolumeclaimStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 newPvc := obj.(*api.PersistentVolumeClaim)
 oldPvc := old.(*api.PersistentVolumeClaim)
 newPvc.Status = oldPvc.Status
 pvcutil.DropDisabledAlphaFields(&newPvc.Spec)
 pvcutil.DropDisabledAlphaFields(&oldPvc.Spec)
}
func (persistentvolumeclaimStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 errorList := validation.ValidatePersistentVolumeClaim(obj.(*api.PersistentVolumeClaim))
 return append(errorList, validation.ValidatePersistentVolumeClaimUpdate(obj.(*api.PersistentVolumeClaim), old.(*api.PersistentVolumeClaim))...)
}
func (persistentvolumeclaimStrategy) AllowUnconditionalUpdate() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return true
}

type persistentvolumeclaimStatusStrategy struct{ persistentvolumeclaimStrategy }

var StatusStrategy = persistentvolumeclaimStatusStrategy{Strategy}

func (persistentvolumeclaimStatusStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 newPv := obj.(*api.PersistentVolumeClaim)
 oldPv := old.(*api.PersistentVolumeClaim)
 newPv.Spec = oldPv.Spec
}
func (persistentvolumeclaimStatusStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return validation.ValidatePersistentVolumeClaimStatusUpdate(obj.(*api.PersistentVolumeClaim), old.(*api.PersistentVolumeClaim))
}
func GetAttrs(obj runtime.Object) (labels.Set, fields.Set, bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 persistentvolumeclaimObj, ok := obj.(*api.PersistentVolumeClaim)
 if !ok {
  return nil, nil, false, fmt.Errorf("not a persistentvolumeclaim")
 }
 return labels.Set(persistentvolumeclaimObj.Labels), PersistentVolumeClaimToSelectableFields(persistentvolumeclaimObj), persistentvolumeclaimObj.Initializers != nil, nil
}
func MatchPersistentVolumeClaim(label labels.Selector, field fields.Selector) storage.SelectionPredicate {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return storage.SelectionPredicate{Label: label, Field: field, GetAttrs: GetAttrs}
}
func PersistentVolumeClaimToSelectableFields(persistentvolumeclaim *api.PersistentVolumeClaim) fields.Set {
 _logClusterCodePath()
 defer _logClusterCodePath()
 objectMetaFieldsSet := generic.ObjectMetaFieldsSet(&persistentvolumeclaim.ObjectMeta, true)
 specificFieldsSet := fields.Set{"name": persistentvolumeclaim.Name}
 return generic.MergeFieldsSets(objectMetaFieldsSet, specificFieldsSet)
}
