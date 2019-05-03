package persistentvolume

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
 pvutil "k8s.io/kubernetes/pkg/api/persistentvolume"
 api "k8s.io/kubernetes/pkg/apis/core"
 "k8s.io/kubernetes/pkg/apis/core/validation"
 volumevalidation "k8s.io/kubernetes/pkg/volume/validation"
)

type persistentvolumeStrategy struct {
 runtime.ObjectTyper
 names.NameGenerator
}

var Strategy = persistentvolumeStrategy{legacyscheme.Scheme, names.SimpleNameGenerator}

func (persistentvolumeStrategy) NamespaceScoped() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return false
}
func (persistentvolumeStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 pv := obj.(*api.PersistentVolume)
 pv.Status = api.PersistentVolumeStatus{}
 pvutil.DropDisabledFields(&pv.Spec, nil)
}
func (persistentvolumeStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 persistentvolume := obj.(*api.PersistentVolume)
 errorList := validation.ValidatePersistentVolume(persistentvolume)
 return append(errorList, volumevalidation.ValidatePersistentVolume(persistentvolume)...)
}
func (persistentvolumeStrategy) Canonicalize(obj runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
}
func (persistentvolumeStrategy) AllowCreateOnUpdate() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return false
}
func (persistentvolumeStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 newPv := obj.(*api.PersistentVolume)
 oldPv := old.(*api.PersistentVolume)
 newPv.Status = oldPv.Status
 pvutil.DropDisabledFields(&newPv.Spec, &oldPv.Spec)
}
func (persistentvolumeStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 newPv := obj.(*api.PersistentVolume)
 errorList := validation.ValidatePersistentVolume(newPv)
 errorList = append(errorList, volumevalidation.ValidatePersistentVolume(newPv)...)
 return append(errorList, validation.ValidatePersistentVolumeUpdate(newPv, old.(*api.PersistentVolume))...)
}
func (persistentvolumeStrategy) AllowUnconditionalUpdate() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return true
}

type persistentvolumeStatusStrategy struct{ persistentvolumeStrategy }

var StatusStrategy = persistentvolumeStatusStrategy{Strategy}

func (persistentvolumeStatusStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 newPv := obj.(*api.PersistentVolume)
 oldPv := old.(*api.PersistentVolume)
 newPv.Spec = oldPv.Spec
}
func (persistentvolumeStatusStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return validation.ValidatePersistentVolumeStatusUpdate(obj.(*api.PersistentVolume), old.(*api.PersistentVolume))
}
func GetAttrs(obj runtime.Object) (labels.Set, fields.Set, bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 persistentvolumeObj, ok := obj.(*api.PersistentVolume)
 if !ok {
  return nil, nil, false, fmt.Errorf("not a persistentvolume")
 }
 return labels.Set(persistentvolumeObj.Labels), PersistentVolumeToSelectableFields(persistentvolumeObj), persistentvolumeObj.Initializers != nil, nil
}
func MatchPersistentVolumes(label labels.Selector, field fields.Selector) storage.SelectionPredicate {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return storage.SelectionPredicate{Label: label, Field: field, GetAttrs: GetAttrs}
}
func PersistentVolumeToSelectableFields(persistentvolume *api.PersistentVolume) fields.Set {
 _logClusterCodePath()
 defer _logClusterCodePath()
 objectMetaFieldsSet := generic.ObjectMetaFieldsSet(&persistentvolume.ObjectMeta, false)
 specificFieldsSet := fields.Set{"name": persistentvolume.Name}
 return generic.MergeFieldsSets(objectMetaFieldsSet, specificFieldsSet)
}
