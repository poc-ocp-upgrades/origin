package event

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

type eventStrategy struct {
 runtime.ObjectTyper
 names.NameGenerator
}

var Strategy = eventStrategy{legacyscheme.Scheme, names.SimpleNameGenerator}

func (eventStrategy) NamespaceScoped() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return true
}
func (eventStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
}
func (eventStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
}
func (eventStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 event := obj.(*api.Event)
 return validation.ValidateEvent(event)
}
func (eventStrategy) AllowCreateOnUpdate() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return false
}
func (eventStrategy) Canonicalize(obj runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
}
func (eventStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 event := obj.(*api.Event)
 return validation.ValidateEvent(event)
}
func (eventStrategy) AllowUnconditionalUpdate() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return true
}
func SelectableFields(pip *api.Event) fields.Set {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return generic.ObjectMetaFieldsSet(&pip.ObjectMeta, true)
}
func GetAttrs(obj runtime.Object) (labels.Set, fields.Set, bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 pip, ok := obj.(*api.Event)
 if !ok {
  return nil, nil, false, fmt.Errorf("given object is not a Event")
 }
 return labels.Set(pip.ObjectMeta.Labels), SelectableFields(pip), pip.Initializers != nil, nil
}
func Matcher(label labels.Selector, field fields.Selector) apistorage.SelectionPredicate {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return apistorage.SelectionPredicate{Label: label, Field: field, GetAttrs: GetAttrs}
}
