package auditsink

import (
	"context"
	goformat "fmt"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/storage/names"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	audit "k8s.io/kubernetes/pkg/apis/auditregistration"
	"k8s.io/kubernetes/pkg/apis/auditregistration/validation"
	goos "os"
	"reflect"
	godefaultruntime "runtime"
	gotime "time"
)

type auditSinkStrategy struct {
	runtime.ObjectTyper
	names.NameGenerator
}

var Strategy = auditSinkStrategy{legacyscheme.Scheme, names.SimpleNameGenerator}

func (auditSinkStrategy) NamespaceScoped() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return false
}
func (auditSinkStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ic := obj.(*audit.AuditSink)
	ic.Generation = 1
}
func (auditSinkStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	newIC := obj.(*audit.AuditSink)
	oldIC := old.(*audit.AuditSink)
	if !reflect.DeepEqual(oldIC.Spec, newIC.Spec) {
		newIC.Generation = oldIC.Generation + 1
	}
}
func (auditSinkStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ic := obj.(*audit.AuditSink)
	return validation.ValidateAuditSink(ic)
}
func (auditSinkStrategy) Canonicalize(obj runtime.Object) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
}
func (auditSinkStrategy) AllowCreateOnUpdate() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return false
}
func (auditSinkStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	validationErrorList := validation.ValidateAuditSink(obj.(*audit.AuditSink))
	updateErrorList := validation.ValidateAuditSinkUpdate(obj.(*audit.AuditSink), old.(*audit.AuditSink))
	return append(validationErrorList, updateErrorList...)
}
func (auditSinkStrategy) AllowUnconditionalUpdate() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return false
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
