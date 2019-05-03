package storageclass

import (
 "context"
 "k8s.io/apimachinery/pkg/runtime"
 "k8s.io/apimachinery/pkg/util/validation/field"
 "k8s.io/apiserver/pkg/storage/names"
 utilfeature "k8s.io/apiserver/pkg/util/feature"
 "k8s.io/kubernetes/pkg/api/legacyscheme"
 "k8s.io/kubernetes/pkg/apis/storage"
 storageutil "k8s.io/kubernetes/pkg/apis/storage/util"
 "k8s.io/kubernetes/pkg/apis/storage/validation"
 "k8s.io/kubernetes/pkg/features"
)

type storageClassStrategy struct {
 runtime.ObjectTyper
 names.NameGenerator
}

var Strategy = storageClassStrategy{legacyscheme.Scheme, names.SimpleNameGenerator}

func (storageClassStrategy) NamespaceScoped() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return false
}
func (storageClassStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 class := obj.(*storage.StorageClass)
 if !utilfeature.DefaultFeatureGate.Enabled(features.ExpandPersistentVolumes) {
  class.AllowVolumeExpansion = nil
 }
 storageutil.DropDisabledAlphaFields(class)
}
func (storageClassStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 storageClass := obj.(*storage.StorageClass)
 return validation.ValidateStorageClass(storageClass)
}
func (storageClassStrategy) Canonicalize(obj runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
}
func (storageClassStrategy) AllowCreateOnUpdate() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return false
}
func (storageClassStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 newClass := obj.(*storage.StorageClass)
 oldClass := old.(*storage.StorageClass)
 if !utilfeature.DefaultFeatureGate.Enabled(features.ExpandPersistentVolumes) {
  newClass.AllowVolumeExpansion = nil
  oldClass.AllowVolumeExpansion = nil
 }
 storageutil.DropDisabledAlphaFields(oldClass)
 storageutil.DropDisabledAlphaFields(newClass)
}
func (storageClassStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 errorList := validation.ValidateStorageClass(obj.(*storage.StorageClass))
 return append(errorList, validation.ValidateStorageClassUpdate(obj.(*storage.StorageClass), old.(*storage.StorageClass))...)
}
func (storageClassStrategy) AllowUnconditionalUpdate() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return true
}
