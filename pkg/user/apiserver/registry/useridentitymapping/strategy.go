package useridentitymapping

import (
	"context"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	userapi "github.com/openshift/origin/pkg/user/apis/user"
	"github.com/openshift/origin/pkg/user/apis/user/validation"
)

type userIdentityMappingStrategy struct{ runtime.ObjectTyper }

var Strategy = userIdentityMappingStrategy{legacyscheme.Scheme}

func (s userIdentityMappingStrategy) NamespaceScoped() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return false
}
func (userIdentityMappingStrategy) GenerateName(base string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return base
}
func (userIdentityMappingStrategy) AllowCreateOnUpdate() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return true
}
func (userIdentityMappingStrategy) AllowUnconditionalUpdate() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return false
}
func (s userIdentityMappingStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	mapping := obj.(*userapi.UserIdentityMapping)
	if len(mapping.Name) == 0 {
		mapping.Name = mapping.Identity.Name
	}
	mapping.Namespace = ""
	mapping.ResourceVersion = ""
	mapping.Identity.Namespace = ""
	mapping.Identity.Kind = ""
	mapping.Identity.UID = ""
	mapping.User.Namespace = ""
	mapping.User.Kind = ""
	mapping.User.UID = ""
}
func (s userIdentityMappingStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	mapping := obj.(*userapi.UserIdentityMapping)
	if len(mapping.Name) == 0 {
		mapping.Name = mapping.Identity.Name
	}
	mapping.Namespace = ""
	mapping.Identity.Namespace = ""
	mapping.Identity.Kind = ""
	mapping.Identity.UID = ""
	mapping.User.Namespace = ""
	mapping.User.Kind = ""
	mapping.User.UID = ""
}
func (s userIdentityMappingStrategy) Canonicalize(obj runtime.Object) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (s userIdentityMappingStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return validation.ValidateUserIdentityMapping(obj.(*userapi.UserIdentityMapping))
}
func (s userIdentityMappingStrategy) ValidateUpdate(ctx context.Context, obj runtime.Object, old runtime.Object) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return validation.ValidateUserIdentityMappingUpdate(obj.(*userapi.UserIdentityMapping), old.(*userapi.UserIdentityMapping))
}
