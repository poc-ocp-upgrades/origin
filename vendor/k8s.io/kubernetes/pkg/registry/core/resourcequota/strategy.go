package resourcequota

import (
 "context"
 "k8s.io/apimachinery/pkg/runtime"
 "k8s.io/apimachinery/pkg/util/validation/field"
 "k8s.io/apiserver/pkg/storage/names"
 "k8s.io/kubernetes/pkg/api/legacyscheme"
 api "k8s.io/kubernetes/pkg/apis/core"
 "k8s.io/kubernetes/pkg/apis/core/validation"
)

type resourcequotaStrategy struct {
 runtime.ObjectTyper
 names.NameGenerator
}

var Strategy = resourcequotaStrategy{legacyscheme.Scheme, names.SimpleNameGenerator}

func (resourcequotaStrategy) NamespaceScoped() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return true
}
func (resourcequotaStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 resourcequota := obj.(*api.ResourceQuota)
 resourcequota.Status = api.ResourceQuotaStatus{}
}
func (resourcequotaStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 newResourcequota := obj.(*api.ResourceQuota)
 oldResourcequota := old.(*api.ResourceQuota)
 newResourcequota.Status = oldResourcequota.Status
}
func (resourcequotaStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 resourcequota := obj.(*api.ResourceQuota)
 return validation.ValidateResourceQuota(resourcequota)
}
func (resourcequotaStrategy) Canonicalize(obj runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
}
func (resourcequotaStrategy) AllowCreateOnUpdate() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return false
}
func (resourcequotaStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 errorList := validation.ValidateResourceQuota(obj.(*api.ResourceQuota))
 return append(errorList, validation.ValidateResourceQuotaUpdate(obj.(*api.ResourceQuota), old.(*api.ResourceQuota))...)
}
func (resourcequotaStrategy) AllowUnconditionalUpdate() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return true
}

type resourcequotaStatusStrategy struct{ resourcequotaStrategy }

var StatusStrategy = resourcequotaStatusStrategy{Strategy}

func (resourcequotaStatusStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 newResourcequota := obj.(*api.ResourceQuota)
 oldResourcequota := old.(*api.ResourceQuota)
 newResourcequota.Spec = oldResourcequota.Spec
}
func (resourcequotaStatusStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return validation.ValidateResourceQuotaStatusUpdate(obj.(*api.ResourceQuota), old.(*api.ResourceQuota))
}
