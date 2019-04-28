package securitycontextconstraints

import (
	"context"
	"fmt"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/registry/rest"
	apistorage "k8s.io/apiserver/pkg/storage"
	"k8s.io/apiserver/pkg/storage/names"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	securityapi "github.com/openshift/origin/pkg/security/apis/security"
	"github.com/openshift/origin/pkg/security/apis/security/validation"
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
	return false
}
func (strategy) AllowCreateOnUpdate() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return false
}
func (strategy) AllowUnconditionalUpdate() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return true
}
func (strategy) PrepareForCreate(_ context.Context, obj runtime.Object) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (strategy) PrepareForUpdate(_ context.Context, obj, old runtime.Object) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (strategy) Canonicalize(obj runtime.Object) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	scc := obj.(*securityapi.SecurityContextConstraints)
	scc.Users = uniqueStrings(scc.Users)
	scc.Groups = uniqueStrings(scc.Groups)
}
func uniqueStrings(values []string) []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(values) < 2 {
		return values
	}
	updated := make([]string, 0, len(values))
	existing := make(map[string]struct{})
	for _, value := range values {
		if _, ok := existing[value]; ok {
			continue
		}
		existing[value] = struct{}{}
		updated = append(updated, value)
	}
	return updated
}
func (strategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return validation.ValidateSecurityContextConstraints(obj.(*securityapi.SecurityContextConstraints))
}
func (strategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return validation.ValidateSecurityContextConstraintsUpdate(obj.(*securityapi.SecurityContextConstraints), old.(*securityapi.SecurityContextConstraints))
}
func GetAttrs(obj runtime.Object) (labels.Set, fields.Set, bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	scc, ok := obj.(*securityapi.SecurityContextConstraints)
	if !ok {
		return nil, nil, false, fmt.Errorf("not SecurityContextConstraints")
	}
	return labels.Set(scc.Labels), SelectableFields(scc), scc.Initializers != nil, nil
}
func Matcher(label labels.Selector, field fields.Selector) apistorage.SelectionPredicate {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return apistorage.SelectionPredicate{Label: label, Field: field, GetAttrs: func(obj runtime.Object) (labels.Set, fields.Set, bool, error) {
		scc, ok := obj.(*securityapi.SecurityContextConstraints)
		if !ok {
			return nil, nil, false, fmt.Errorf("not a securitycontextconstraint")
		}
		return labels.Set(scc.Labels), SelectableFields(scc), scc.Initializers != nil, nil
	}}
}
func SelectableFields(obj *securityapi.SecurityContextConstraints) fields.Set {
	_logClusterCodePath()
	defer _logClusterCodePath()
	objectMetaFieldsSet := generic.ObjectMetaFieldsSet(&obj.ObjectMeta, true)
	return objectMetaFieldsSet
}
