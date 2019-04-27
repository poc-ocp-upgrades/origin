package openshiftadmission

import (
	"testing"
	"k8s.io/apimachinery/pkg/util/sets"
	"github.com/openshift/origin/pkg/admission/admissionregistrationtesting"
)

func TestAdmissionRegistration(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	err := admissionregistrationtesting.AdmissionRegistrationTest(OriginAdmissionPlugins, OpenShiftAdmissionPlugins, sets.String{})
	if err != nil {
		t.Fatal(err)
	}
}
