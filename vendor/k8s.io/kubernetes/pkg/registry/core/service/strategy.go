package service

import (
 "context"
 "fmt"
 "k8s.io/apimachinery/pkg/runtime"
 "k8s.io/apimachinery/pkg/util/validation/field"
 "k8s.io/apiserver/pkg/storage/names"
 "k8s.io/kubernetes/pkg/api/legacyscheme"
 api "k8s.io/kubernetes/pkg/apis/core"
 "k8s.io/kubernetes/pkg/apis/core/validation"
)

type svcStrategy struct {
 runtime.ObjectTyper
 names.NameGenerator
}

var Strategy = svcStrategy{legacyscheme.Scheme, names.SimpleNameGenerator}

func (svcStrategy) NamespaceScoped() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return true
}
func (svcStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 service := obj.(*api.Service)
 service.Status = api.ServiceStatus{}
}
func (svcStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 newService := obj.(*api.Service)
 oldService := old.(*api.Service)
 newService.Status = oldService.Status
}
func (svcStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 service := obj.(*api.Service)
 return validation.ValidateService(service)
}
func (svcStrategy) Canonicalize(obj runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
}
func (svcStrategy) AllowCreateOnUpdate() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return true
}
func (svcStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return validation.ValidateServiceUpdate(obj.(*api.Service), old.(*api.Service))
}
func (svcStrategy) AllowUnconditionalUpdate() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return true
}
func (svcStrategy) Export(ctx context.Context, obj runtime.Object, exact bool) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 t, ok := obj.(*api.Service)
 if !ok {
  return fmt.Errorf("unexpected object: %v", obj)
 }
 t.Status = api.ServiceStatus{}
 if exact {
  return nil
 }
 if t.Spec.ClusterIP != api.ClusterIPNone {
  t.Spec.ClusterIP = ""
 }
 if t.Spec.Type == api.ServiceTypeNodePort {
  for i := range t.Spec.Ports {
   t.Spec.Ports[i].NodePort = 0
  }
 }
 return nil
}

type serviceStatusStrategy struct{ svcStrategy }

var StatusStrategy = serviceStatusStrategy{Strategy}

func (serviceStatusStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 newService := obj.(*api.Service)
 oldService := old.(*api.Service)
 newService.Spec = oldService.Spec
}
func (serviceStatusStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return validation.ValidateServiceStatusUpdate(obj.(*api.Service), old.(*api.Service))
}
