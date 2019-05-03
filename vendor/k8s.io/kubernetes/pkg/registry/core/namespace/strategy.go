package namespace

import (
 "context"
 "fmt"
 "k8s.io/apimachinery/pkg/fields"
 "k8s.io/apimachinery/pkg/labels"
 "k8s.io/apimachinery/pkg/runtime"
 "k8s.io/apimachinery/pkg/util/validation/field"
 "k8s.io/apiserver/pkg/registry/generic"
 apistorage "k8s.io/apiserver/pkg/storage"
 "k8s.io/apiserver/pkg/storage/names"
 "k8s.io/kubernetes/pkg/api/legacyscheme"
 api "k8s.io/kubernetes/pkg/apis/core"
 "k8s.io/kubernetes/pkg/apis/core/validation"
)

type namespaceStrategy struct {
 runtime.ObjectTyper
 names.NameGenerator
}

var Strategy = namespaceStrategy{legacyscheme.Scheme, names.SimpleNameGenerator}

func (namespaceStrategy) NamespaceScoped() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return false
}
func (namespaceStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 namespace := obj.(*api.Namespace)
 namespace.Status = api.NamespaceStatus{Phase: api.NamespaceActive}
 hasKubeFinalizer := false
 for i := range namespace.Spec.Finalizers {
  if namespace.Spec.Finalizers[i] == api.FinalizerKubernetes {
   hasKubeFinalizer = true
   break
  }
 }
 if !hasKubeFinalizer {
  if len(namespace.Spec.Finalizers) == 0 {
   namespace.Spec.Finalizers = []api.FinalizerName{api.FinalizerKubernetes}
  } else {
   namespace.Spec.Finalizers = append(namespace.Spec.Finalizers, api.FinalizerKubernetes)
  }
 }
}
func (namespaceStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 newNamespace := obj.(*api.Namespace)
 oldNamespace := old.(*api.Namespace)
 newNamespace.Spec.Finalizers = oldNamespace.Spec.Finalizers
 newNamespace.Status = oldNamespace.Status
}
func (namespaceStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 namespace := obj.(*api.Namespace)
 return validation.ValidateNamespace(namespace)
}
func (namespaceStrategy) Canonicalize(obj runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
}
func (namespaceStrategy) AllowCreateOnUpdate() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return false
}
func (namespaceStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 errorList := validation.ValidateNamespace(obj.(*api.Namespace))
 return append(errorList, validation.ValidateNamespaceUpdate(obj.(*api.Namespace), old.(*api.Namespace))...)
}
func (namespaceStrategy) AllowUnconditionalUpdate() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return true
}

type namespaceStatusStrategy struct{ namespaceStrategy }

var StatusStrategy = namespaceStatusStrategy{Strategy}

func (namespaceStatusStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 newNamespace := obj.(*api.Namespace)
 oldNamespace := old.(*api.Namespace)
 newNamespace.Spec = oldNamespace.Spec
}
func (namespaceStatusStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return validation.ValidateNamespaceStatusUpdate(obj.(*api.Namespace), old.(*api.Namespace))
}

type namespaceFinalizeStrategy struct{ namespaceStrategy }

var FinalizeStrategy = namespaceFinalizeStrategy{Strategy}

func (namespaceFinalizeStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return validation.ValidateNamespaceFinalizeUpdate(obj.(*api.Namespace), old.(*api.Namespace))
}
func (namespaceFinalizeStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 newNamespace := obj.(*api.Namespace)
 oldNamespace := old.(*api.Namespace)
 newNamespace.Status = oldNamespace.Status
}
func GetAttrs(obj runtime.Object) (labels.Set, fields.Set, bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 namespaceObj, ok := obj.(*api.Namespace)
 if !ok {
  return nil, nil, false, fmt.Errorf("not a namespace")
 }
 return labels.Set(namespaceObj.Labels), NamespaceToSelectableFields(namespaceObj), namespaceObj.Initializers != nil, nil
}
func MatchNamespace(label labels.Selector, field fields.Selector) apistorage.SelectionPredicate {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return apistorage.SelectionPredicate{Label: label, Field: field, GetAttrs: GetAttrs}
}
func NamespaceToSelectableFields(namespace *api.Namespace) fields.Set {
 _logClusterCodePath()
 defer _logClusterCodePath()
 objectMetaFieldsSet := generic.ObjectMetaFieldsSet(&namespace.ObjectMeta, false)
 specificFieldsSet := fields.Set{"status.phase": string(namespace.Status.Phase), "name": namespace.Name}
 return generic.MergeFieldsSets(objectMetaFieldsSet, specificFieldsSet)
}
