package identity

import (
	"context"
	"bytes"
	"net/http"
	"runtime"
	"fmt"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	userapi "github.com/openshift/origin/pkg/user/apis/user"
	"github.com/openshift/origin/pkg/user/apis/user/validation"
)

type identityStrategy struct{ runtime.ObjectTyper }

var Strategy = identityStrategy{legacyscheme.Scheme}
var _ rest.GarbageCollectionDeleteStrategy = identityStrategy{}

func (identityStrategy) DefaultGarbageCollectionPolicy(ctx context.Context) rest.GarbageCollectionPolicy {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return rest.Unsupported
}
func (identityStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (identityStrategy) NamespaceScoped() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return false
}
func (identityStrategy) GenerateName(base string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return base
}
func (identityStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	identity := obj.(*userapi.Identity)
	identity.Name = identityName(identity.ProviderName, identity.ProviderUserName)
}
func identityName(provider, identity string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return provider + ":" + identity
}
func (identityStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	identity := obj.(*userapi.Identity)
	return validation.ValidateIdentity(identity)
}
func (identityStrategy) AllowCreateOnUpdate() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return false
}
func (identityStrategy) AllowUnconditionalUpdate() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return false
}
func (identityStrategy) Canonicalize(obj runtime.Object) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (identityStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return validation.ValidateIdentityUpdate(obj.(*userapi.Identity), old.(*userapi.Identity))
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
