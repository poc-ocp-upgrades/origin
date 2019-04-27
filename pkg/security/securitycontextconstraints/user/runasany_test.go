package user

import (
	"testing"
	securityapi "github.com/openshift/origin/pkg/security/apis/security"
)

func TestRunAsAnyOptions(t *testing.T) {
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
	_, err := NewRunAsAny(nil)
	if err != nil {
		t.Fatalf("unexpected error initializing NewRunAsAny %v", err)
	}
	_, err = NewRunAsAny(&securityapi.RunAsUserStrategyOptions{})
	if err != nil {
		t.Errorf("unexpected error initializing NewRunAsAny %v", err)
	}
}
func TestRunAsAnyGenerate(t *testing.T) {
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
	s, err := NewRunAsAny(&securityapi.RunAsUserStrategyOptions{})
	if err != nil {
		t.Fatalf("unexpected error initializing NewRunAsAny %v", err)
	}
	uid, err := s.Generate(nil, nil)
	if uid != nil {
		t.Errorf("expected nil uid but got %d", *uid)
	}
	if err != nil {
		t.Errorf("unexpected error generating uid %v", err)
	}
}
func TestRunAsAnyValidate(t *testing.T) {
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
	s, err := NewRunAsAny(&securityapi.RunAsUserStrategyOptions{})
	if err != nil {
		t.Fatalf("unexpected error initializing NewRunAsAny %v", err)
	}
	errs := s.Validate(nil, nil, nil, nil, nil)
	if len(errs) != 0 {
		t.Errorf("unexpected errors validating with ")
	}
}
