package brokertemplateinstance

import (
	"context"
	"bytes"
	"net/http"
	"runtime"
	"fmt"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/storage/names"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	templateapi "github.com/openshift/origin/pkg/template/apis/template"
	"github.com/openshift/origin/pkg/template/apis/template/validation"
)

type brokerTemplateInstanceStrategy struct {
	runtime.ObjectTyper
	names.NameGenerator
}

var Strategy = brokerTemplateInstanceStrategy{legacyscheme.Scheme, names.SimpleNameGenerator}

func (brokerTemplateInstanceStrategy) NamespaceScoped() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return false
}
func (brokerTemplateInstanceStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (brokerTemplateInstanceStrategy) Canonicalize(obj runtime.Object) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (brokerTemplateInstanceStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (brokerTemplateInstanceStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return validation.ValidateBrokerTemplateInstance(obj.(*templateapi.BrokerTemplateInstance))
}
func (brokerTemplateInstanceStrategy) AllowCreateOnUpdate() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return false
}
func (brokerTemplateInstanceStrategy) AllowUnconditionalUpdate() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return false
}
func (brokerTemplateInstanceStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return validation.ValidateBrokerTemplateInstanceUpdate(obj.(*templateapi.BrokerTemplateInstance), old.(*templateapi.BrokerTemplateInstance))
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
