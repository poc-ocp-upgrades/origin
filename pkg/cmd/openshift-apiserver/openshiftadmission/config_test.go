package openshiftadmission

import (
	"strings"
	"testing"
	"k8s.io/apiserver/pkg/admission"
)

func TestAdmissionPluginNames(t *testing.T) {
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
	originAdmissionPlugins := admission.NewPlugins()
	RegisterOpenshiftAdmissionPlugins(originAdmissionPlugins)
	for _, plugin := range originAdmissionPlugins.Registered() {
		if !strings.Contains(plugin, "openshift.io/") {
			t.Errorf("openshift admission plugins must be prefixed with openshift.io/ %v", plugin)
		}
	}
}
