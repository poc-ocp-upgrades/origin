package validation

import (
	"testing"
	"k8s.io/apimachinery/pkg/util/validation/field"
	buildapi "github.com/openshift/origin/pkg/build/apis/build"
	configapi "github.com/openshift/origin/pkg/cmd/server/apis/config"
)

func TestValidateBuildOverridesConfig(t *testing.T) {
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
		config		*configapi.BuildOverridesConfig
		errExpected	bool
		errField	string
		errType		field.ErrorType
	}{{config: &configapi.BuildOverridesConfig{ImageLabels: []buildapi.ImageLabel{{Name: "A", Value: "B"}}}, errExpected: false}, {config: &configapi.BuildOverridesConfig{ImageLabels: []buildapi.ImageLabel{{Name: "", Value: "empty"}}}, errExpected: true, errField: "imageLabels[0].name", errType: field.ErrorTypeRequired}, {config: &configapi.BuildOverridesConfig{ImageLabels: []buildapi.ImageLabel{{Name: "\tč;", Value: "????"}}}, errExpected: true, errField: "imageLabels[0].name", errType: field.ErrorTypeInvalid}, {config: &configapi.BuildOverridesConfig{ImageLabels: []buildapi.ImageLabel{{Name: "name", Value: "Jan"}, {Name: "name", Value: "Elvis"}}}, errExpected: true, errField: "imageLabels[1].name", errType: field.ErrorTypeInvalid}, {config: &configapi.BuildOverridesConfig{NodeSelector: map[string]string{"A": "B"}}, errExpected: false}, {config: &configapi.BuildOverridesConfig{NodeSelector: map[string]string{"A@B!": "C"}}, errExpected: true, errField: "nodeSelector[A@B!]", errType: field.ErrorTypeInvalid}, {config: &configapi.BuildOverridesConfig{Annotations: map[string]string{"A": "B"}}, errExpected: false}, {config: &configapi.BuildOverridesConfig{Annotations: map[string]string{"A B": "C"}}, errExpected: true, errField: "annotations", errType: field.ErrorTypeInvalid}}
	for i, tc := range tests {
		result := ValidateBuildOverridesConfig(tc.config)
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
