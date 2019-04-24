package config

import (
	"reflect"
	"testing"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func TestGetKubeAPIServerFlagAPIEnablement(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCases := []struct {
		name		string
		flagValue	[]string
		expected	map[schema.GroupVersion]bool
	}{{name: "skip bad", flagValue: []string{"api/legacy=true", "foobar/is/bad=true", "foo/v1=true=true", "good/v1=false"}, expected: map[schema.GroupVersion]bool{{Group: "good", Version: "v1"}: false}}, {name: "good", flagValue: []string{"good/v2=true", "good/v1=false"}, expected: map[schema.GroupVersion]bool{{Group: "good", Version: "v1"}: false, {Group: "good", Version: "v2"}: true}}}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := GetKubeAPIServerFlagAPIEnablement(tc.flagValue)
			if !reflect.DeepEqual(tc.expected, actual) {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}
func TestGetEnabledAPIVersionsForGroup(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCases := []struct {
		name				string
		flagValue			[]string
		disabledAPIGroupVersions	map[string][]string
		apiGroup			string
		expected			[]string
	}{{name: "enable unknown from flag", apiGroup: "good", flagValue: []string{"good/v2=true", "good/v1=false"}, expected: []string{"v2"}}, {name: "enable from flag, disabled by disable", apiGroup: "good", disabledAPIGroupVersions: map[string][]string{"good": {"v2"}}, flagValue: []string{"good/v2=true", "good/v1=false"}, expected: []string{}}, {name: "enable by default, disable by flag", apiGroup: "batch", flagValue: []string{"batch/v1=false"}, expected: []string{"v1beta1"}}, {name: "enable by default, no disable", apiGroup: "batch", expected: []string{"v1", "v1beta1"}}, {name: "enable settings", apiGroup: "settings.k8s.io", flagValue: []string{"settings.k8s.io/v1alpha1=true"}, expected: []string{"v1alpha1"}}}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			config := KubernetesMasterConfig{DisabledAPIGroupVersions: tc.disabledAPIGroupVersions, APIServerArguments: map[string][]string{kubeAPIEnablementFlag: tc.flagValue}}
			actual := GetEnabledAPIVersionsForGroup(config, tc.apiGroup)
			if !reflect.DeepEqual(tc.expected, actual) {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}
func TestGetDisabledAPIVersionsForGroup(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCases := []struct {
		name				string
		flagValue			[]string
		disabledAPIGroupVersions	map[string][]string
		apiGroup			string
		expected			[]string
	}{{name: "enable unknown from flag", apiGroup: "good", flagValue: []string{"good/v2=true", "good/v1=false"}, expected: []string{"v1"}}, {name: "enable from flag, disabled by disable", apiGroup: "good", disabledAPIGroupVersions: map[string][]string{"good": {"v2"}}, flagValue: []string{"good/v2=true", "good/v1=false"}, expected: []string{"v1", "v2"}}, {name: "enable by default, disable by flag", apiGroup: "batch", flagValue: []string{"batch/v1=false"}, expected: []string{"v1"}}, {name: "enable by default, no disable", apiGroup: "batch", expected: []string{}}}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			config := KubernetesMasterConfig{DisabledAPIGroupVersions: tc.disabledAPIGroupVersions, APIServerArguments: map[string][]string{kubeAPIEnablementFlag: tc.flagValue}}
			actual := GetDisabledAPIVersionsForGroup(config, tc.apiGroup)
			if !reflect.DeepEqual(tc.expected, actual) {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}
