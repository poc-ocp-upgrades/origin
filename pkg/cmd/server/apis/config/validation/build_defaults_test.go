package validation

import (
	"testing"
	"k8s.io/apimachinery/pkg/util/validation/field"
	kapi "k8s.io/kubernetes/pkg/apis/core"
	buildapi "github.com/openshift/origin/pkg/build/apis/build"
	configapi "github.com/openshift/origin/pkg/cmd/server/apis/config"
)

func TestValidateBuildDefaultsConfig(t *testing.T) {
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
	tests := []struct {
		config		*configapi.BuildDefaultsConfig
		errExpected	bool
		errField	string
		errType		field.ErrorType
	}{{config: &configapi.BuildDefaultsConfig{GitHTTPProxy: "http://valid.url", GitHTTPSProxy: "https://valid.url", Env: []kapi.EnvVar{{Name: "VAR1", Value: "VALUE1"}, {Name: "VAR2", Value: "VALUE2"}}}, errExpected: false}, {config: &configapi.BuildDefaultsConfig{GitHTTPProxy: "some!@#$%^&*()url", GitHTTPSProxy: "https://valid.url"}, errExpected: true, errField: "gitHTTPProxy", errType: field.ErrorTypeInvalid}, {config: &configapi.BuildDefaultsConfig{GitHTTPProxy: "https://valid.url", GitHTTPSProxy: "some!@#$%^&*()url"}, errExpected: true, errField: "gitHTTPSProxy", errType: field.ErrorTypeInvalid}, {config: &configapi.BuildDefaultsConfig{Env: []kapi.EnvVar{{Name: "", Value: "test"}}}, errExpected: true, errField: "env[0].name", errType: field.ErrorTypeRequired}, {config: &configapi.BuildDefaultsConfig{Env: []kapi.EnvVar{{Name: " invalid,name", Value: "test"}}}, errExpected: true, errField: "env[0].name", errType: field.ErrorTypeInvalid}, {config: &configapi.BuildDefaultsConfig{Env: []kapi.EnvVar{{Name: "name", ValueFrom: &kapi.EnvVarSource{ResourceFieldRef: &kapi.ResourceFieldSelector{ContainerName: "name", Resource: "resource"}}}}}, errExpected: true, errField: "env[0].valueFrom.ResourceFieldRef", errType: field.ErrorTypeInvalid}, {config: &configapi.BuildDefaultsConfig{ImageLabels: []buildapi.ImageLabel{{Name: "", Value: "empty"}}}, errExpected: true, errField: "imageLabels[0].name", errType: field.ErrorTypeRequired}, {config: &configapi.BuildDefaultsConfig{ImageLabels: []buildapi.ImageLabel{{Name: "\tč;", Value: "????"}}}, errExpected: true, errField: "imageLabels[0].name", errType: field.ErrorTypeInvalid}, {config: &configapi.BuildDefaultsConfig{ImageLabels: []buildapi.ImageLabel{{Name: "name", Value: "Jan"}, {Name: "name", Value: "Elvis"}}}, errExpected: true, errField: "imageLabels[1].name", errType: field.ErrorTypeInvalid}, {config: &configapi.BuildDefaultsConfig{NodeSelector: map[string]string{"A": "B"}}, errExpected: false}, {config: &configapi.BuildDefaultsConfig{NodeSelector: map[string]string{"A@B!": "C"}}, errExpected: true, errField: "nodeSelector[A@B!]", errType: field.ErrorTypeInvalid}, {config: &configapi.BuildDefaultsConfig{Annotations: map[string]string{"A": "B"}}, errExpected: false}, {config: &configapi.BuildDefaultsConfig{Annotations: map[string]string{"A B": "C"}}, errExpected: true, errField: "annotations", errType: field.ErrorTypeInvalid}}
	for i, tc := range tests {
		result := ValidateBuildDefaultsConfig(tc.config)
		if !tc.errExpected {
			if len(result) > 0 {
				t.Errorf("%d: unexpected error: %v", i, result.ToAggregate())
			}
			continue
		}
		if tc.errExpected && len(result) == 0 {
			t.Errorf("%d: did not get expected error", i)
			continue
		}
		err := result[0]
		if err.Type != tc.errType {
			t.Errorf("%d: unexpected error type: %v", i, err.Type)
		}
		if err.Field != tc.errField {
			t.Errorf("%d: unexpected error field: %v", i, err.Field)
		}
	}
}
