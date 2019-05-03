package limitrange

import (
 "context"
 "k8s.io/apimachinery/pkg/runtime"
 "k8s.io/apimachinery/pkg/util/uuid"
 "k8s.io/apimachinery/pkg/util/validation/field"
 "k8s.io/apiserver/pkg/storage/names"
 "k8s.io/kubernetes/pkg/api/legacyscheme"
 api "k8s.io/kubernetes/pkg/apis/core"
 "k8s.io/kubernetes/pkg/apis/core/validation"
)

type limitrangeStrategy struct {
 runtime.ObjectTyper
 names.NameGenerator
}

var Strategy = limitrangeStrategy{legacyscheme.Scheme, names.SimpleNameGenerator}

func (limitrangeStrategy) NamespaceScoped() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return true
}
func (limitrangeStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 limitRange := obj.(*api.LimitRange)
 if len(limitRange.Name) == 0 {
  limitRange.Name = string(uuid.NewUUID())
 }
}
func (limitrangeStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
}
func (limitrangeStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 limitRange := obj.(*api.LimitRange)
 return validation.ValidateLimitRange(limitRange)
}
func (limitrangeStrategy) Canonicalize(obj runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
}
func (limitrangeStrategy) AllowCreateOnUpdate() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return true
}
func (limitrangeStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 limitRange := obj.(*api.LimitRange)
 return validation.ValidateLimitRange(limitRange)
}
func (limitrangeStrategy) AllowUnconditionalUpdate() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return true
}
func (limitrangeStrategy) Export(context.Context, runtime.Object, bool) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return nil
}
