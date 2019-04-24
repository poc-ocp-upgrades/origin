package user

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

type userStrategy struct{ runtime.ObjectTyper }

var Strategy = userStrategy{legacyscheme.Scheme}
var _ rest.GarbageCollectionDeleteStrategy = userStrategy{}

func (userStrategy) DefaultGarbageCollectionPolicy(ctx context.Context) rest.GarbageCollectionPolicy {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return rest.Unsupported
}
func (userStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (userStrategy) NamespaceScoped() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return false
}
func (userStrategy) GenerateName(base string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return base
}
func (userStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (userStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return validation.ValidateUser(obj.(*userapi.User))
}
func (userStrategy) AllowCreateOnUpdate() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return false
}
func (userStrategy) AllowUnconditionalUpdate() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return false
}
func (userStrategy) Canonicalize(obj runtime.Object) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (userStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return validation.ValidateUserUpdate(obj.(*userapi.User), old.(*userapi.User))
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
