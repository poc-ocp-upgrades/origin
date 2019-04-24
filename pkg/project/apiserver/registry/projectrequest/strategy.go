package projectrequest

import (
	"context"
	"bytes"
	"net/http"
	"runtime"
	"fmt"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	projectapi "github.com/openshift/origin/pkg/project/apis/project"
	projectvalidation "github.com/openshift/origin/pkg/project/apis/project/validation"
)

type strategy struct{ runtime.ObjectTyper }

var Strategy = strategy{legacyscheme.Scheme}

func (strategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (strategy) NamespaceScoped() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return false
}
func (strategy) GenerateName(base string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return base
}
func (strategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (strategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	projectrequest := obj.(*projectapi.ProjectRequest)
	return projectvalidation.ValidateProjectRequest(projectrequest)
}
func (strategy) ValidateUpdate(ctx context.Context, obj runtime.Object, old runtime.Object) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
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
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
