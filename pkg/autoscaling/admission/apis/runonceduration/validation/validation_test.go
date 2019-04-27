package validation

import (
	"testing"
	"github.com/openshift/origin/pkg/autoscaling/admission/apis/runonceduration"
)

func TestRunOnceDurationConfigValidation(t *testing.T) {
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
	var invalidSecs int64 = -1
	invalidConfig := &runonceduration.RunOnceDurationConfig{ActiveDeadlineSecondsLimit: &invalidSecs}
	errs := ValidateRunOnceDurationConfig(invalidConfig)
	if len(errs) == 0 {
		t.Errorf("Did not get expected error on invalid config")
	}
	var validSecs int64 = 5
	validConfig := &runonceduration.RunOnceDurationConfig{ActiveDeadlineSecondsLimit: &validSecs}
	errs = ValidateRunOnceDurationConfig(validConfig)
	if len(errs) > 0 {
		t.Errorf("Unexpected error on valid config")
	}
}
