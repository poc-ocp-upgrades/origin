package endpoint

import (
 "context"
 "k8s.io/apimachinery/pkg/runtime"
 "k8s.io/apimachinery/pkg/util/validation/field"
 "k8s.io/apiserver/pkg/storage/names"
 endptspkg "k8s.io/kubernetes/pkg/api/endpoints"
 "k8s.io/kubernetes/pkg/api/legacyscheme"
 api "k8s.io/kubernetes/pkg/apis/core"
 "k8s.io/kubernetes/pkg/apis/core/validation"
)

type endpointsStrategy struct {
 runtime.ObjectTyper
 names.NameGenerator
}

var Strategy = endpointsStrategy{legacyscheme.Scheme, names.SimpleNameGenerator}

func (endpointsStrategy) NamespaceScoped() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return true
}
func (endpointsStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
}
func (endpointsStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
}
func (endpointsStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return validation.ValidateEndpoints(obj.(*api.Endpoints))
}
func (endpointsStrategy) Canonicalize(obj runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 endpoints := obj.(*api.Endpoints)
 endpoints.Subsets = endptspkg.RepackSubsets(endpoints.Subsets)
}
func (endpointsStrategy) AllowCreateOnUpdate() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return true
}
func (endpointsStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 errorList := validation.ValidateEndpoints(obj.(*api.Endpoints))
 return append(errorList, validation.ValidateEndpointsUpdate(obj.(*api.Endpoints), old.(*api.Endpoints))...)
}
func (endpointsStrategy) AllowUnconditionalUpdate() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return true
}
