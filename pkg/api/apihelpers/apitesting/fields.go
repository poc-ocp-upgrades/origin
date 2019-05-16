package apitesting

import (
	goformat "fmt"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	goos "os"
	godefaultruntime "runtime"
	"testing"
	gotime "time"
)

type FieldKeyCheck struct {
	SchemeBuilder            runtime.SchemeBuilder
	Kind                     schema.GroupVersionKind
	AllowedExternalFieldKeys []string
	FieldKeyEvaluatorFn      FieldKeyEvaluator
}

func (f FieldKeyCheck) Check(t *testing.T) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	scheme := runtime.NewScheme()
	f.SchemeBuilder.AddToScheme(scheme)
	internalObj, err := scheme.New(f.Kind.GroupKind().WithVersion(runtime.APIVersionInternal))
	if err != nil {
		t.Errorf("unable to new up %v", f.Kind)
	}
	for _, externalFieldKey := range f.AllowedExternalFieldKeys {
		internalFieldKey, _, err := scheme.ConvertFieldLabel(f.Kind, externalFieldKey, "")
		if err != nil {
			t.Errorf("illegal field conversion %q for %v", externalFieldKey, f.Kind)
			continue
		}
		if internalFieldKey == "metadata.name" {
			continue
		}
		fieldSet := fields.Set{}
		if err := f.FieldKeyEvaluatorFn(internalObj, fieldSet); err != nil {
			t.Errorf("unable to valuate field keys for %v: %v", f.Kind, err)
			continue
		}
		found := false
		for actualInternalFieldKey := range fieldSet {
			if internalFieldKey == actualInternalFieldKey {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("%q converted to %q which has no internal field key match for %v", externalFieldKey, internalFieldKey, f.Kind)
			continue
		}
	}
}

type FieldKeyEvaluator func(obj runtime.Object, fieldSet fields.Set) error

func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
