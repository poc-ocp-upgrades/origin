package secret

import (
 "context"
 "fmt"
 "k8s.io/apimachinery/pkg/api/errors"
 "k8s.io/apimachinery/pkg/fields"
 "k8s.io/apimachinery/pkg/labels"
 "k8s.io/apimachinery/pkg/runtime"
 "k8s.io/apimachinery/pkg/util/validation/field"
 "k8s.io/apiserver/pkg/registry/generic"
 "k8s.io/apiserver/pkg/registry/rest"
 pkgstorage "k8s.io/apiserver/pkg/storage"
 "k8s.io/apiserver/pkg/storage/names"
 "k8s.io/kubernetes/pkg/api/legacyscheme"
 api "k8s.io/kubernetes/pkg/apis/core"
 "k8s.io/kubernetes/pkg/apis/core/validation"
)

type strategy struct {
 runtime.ObjectTyper
 names.NameGenerator
}

var Strategy = strategy{legacyscheme.Scheme, names.SimpleNameGenerator}
var _ = rest.RESTCreateStrategy(Strategy)
var _ = rest.RESTUpdateStrategy(Strategy)

func (strategy) NamespaceScoped() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return true
}
func (strategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
}
func (strategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return validation.ValidateSecret(obj.(*api.Secret))
}
func (strategy) Canonicalize(obj runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
}
func (strategy) AllowCreateOnUpdate() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return false
}
func (strategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
}
func (strategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return validation.ValidateSecretUpdate(obj.(*api.Secret), old.(*api.Secret))
}
func (strategy) AllowUnconditionalUpdate() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return true
}
func (s strategy) Export(ctx context.Context, obj runtime.Object, exact bool) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 t, ok := obj.(*api.Secret)
 if !ok {
  return fmt.Errorf("unexpected object: %v", obj)
 }
 s.PrepareForCreate(ctx, obj)
 if exact {
  return nil
 }
 if t.Type == api.SecretTypeServiceAccountToken || len(t.Annotations[api.ServiceAccountUIDKey]) > 0 {
  errs := []*field.Error{field.Invalid(field.NewPath("type"), t, "can not export service account secrets")}
  return errors.NewInvalid(api.Kind("Secret"), t.Name, errs)
 }
 return nil
}
func GetAttrs(obj runtime.Object) (labels.Set, fields.Set, bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 secret, ok := obj.(*api.Secret)
 if !ok {
  return nil, nil, false, fmt.Errorf("not a secret")
 }
 return labels.Set(secret.Labels), SelectableFields(secret), secret.Initializers != nil, nil
}
func Matcher(label labels.Selector, field fields.Selector) pkgstorage.SelectionPredicate {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return pkgstorage.SelectionPredicate{Label: label, Field: field, GetAttrs: GetAttrs, IndexFields: []string{"metadata.name"}}
}
func SecretNameTriggerFunc(obj runtime.Object) []pkgstorage.MatchValue {
 _logClusterCodePath()
 defer _logClusterCodePath()
 secret := obj.(*api.Secret)
 result := pkgstorage.MatchValue{IndexName: "metadata.name", Value: secret.ObjectMeta.Name}
 return []pkgstorage.MatchValue{result}
}
func SelectableFields(obj *api.Secret) fields.Set {
 _logClusterCodePath()
 defer _logClusterCodePath()
 objectMetaFieldsSet := generic.ObjectMetaFieldsSet(&obj.ObjectMeta, true)
 secretSpecificFieldsSet := fields.Set{"type": string(obj.Type)}
 return generic.MergeFieldsSets(objectMetaFieldsSet, secretSpecificFieldsSet)
}
