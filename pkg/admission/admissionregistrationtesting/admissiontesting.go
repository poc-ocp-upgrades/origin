package admissionregistrationtesting

import (
	"fmt"
	goformat "fmt"
	"k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apiserver/pkg/admission"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

func AdmissionRegistrationTest(registeredAdmission *admission.Plugins, orderedAdmissionPlugins []string, defaultOffPlugins sets.String) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	errs := []error{}
	registeredPlugins := sets.NewString(registeredAdmission.Registered()...)
	orderedAdmissionPluginsSet := sets.NewString(orderedAdmissionPlugins...)
	if diff := orderedAdmissionPluginsSet.Difference(registeredPlugins); len(diff) > 0 {
		errs = append(errs, fmt.Errorf("registered plugins missing admission plugins:  %v", diff.List()))
	}
	if diff := defaultOffPlugins.Difference(orderedAdmissionPluginsSet); len(diff) > 0 {
		errs = append(errs, fmt.Errorf("ordered admission plugins missing defaultOff plugins: %v", diff.List()))
	}
	return errors.NewAggregate(errs)
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
