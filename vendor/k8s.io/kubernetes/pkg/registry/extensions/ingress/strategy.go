package ingress

import (
	"context"
	goformat "fmt"
	apiequality "k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/storage/names"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	"k8s.io/kubernetes/pkg/apis/extensions"
	"k8s.io/kubernetes/pkg/apis/extensions/validation"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

type ingressStrategy struct {
	runtime.ObjectTyper
	names.NameGenerator
}

var Strategy = ingressStrategy{legacyscheme.Scheme, names.SimpleNameGenerator}

func (ingressStrategy) NamespaceScoped() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return true
}
func (ingressStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ingress := obj.(*extensions.Ingress)
	ingress.Status = extensions.IngressStatus{}
	ingress.Generation = 1
}
func (ingressStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	newIngress := obj.(*extensions.Ingress)
	oldIngress := old.(*extensions.Ingress)
	newIngress.Status = oldIngress.Status
	if !apiequality.Semantic.DeepEqual(oldIngress.Spec, newIngress.Spec) {
		newIngress.Generation = oldIngress.Generation + 1
	}
}
func (ingressStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ingress := obj.(*extensions.Ingress)
	err := validation.ValidateIngress(ingress)
	return err
}
func (ingressStrategy) Canonicalize(obj runtime.Object) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
}
func (ingressStrategy) AllowCreateOnUpdate() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return false
}
func (ingressStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	validationErrorList := validation.ValidateIngress(obj.(*extensions.Ingress))
	updateErrorList := validation.ValidateIngressUpdate(obj.(*extensions.Ingress), old.(*extensions.Ingress))
	return append(validationErrorList, updateErrorList...)
}
func (ingressStrategy) AllowUnconditionalUpdate() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return true
}

type ingressStatusStrategy struct{ ingressStrategy }

var StatusStrategy = ingressStatusStrategy{Strategy}

func (ingressStatusStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	newIngress := obj.(*extensions.Ingress)
	oldIngress := old.(*extensions.Ingress)
	newIngress.Spec = oldIngress.Spec
}
func (ingressStatusStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return validation.ValidateIngressStatusUpdate(obj.(*extensions.Ingress), old.(*extensions.Ingress))
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
