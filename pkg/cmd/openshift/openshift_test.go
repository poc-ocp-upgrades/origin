package openshift

import (
	"testing"
)

func TestCommandFor(t *testing.T) {
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
	cmd := CommandFor("unknown")
	if cmd.Use != "openshift" {
		t.Errorf("expected command to be openshift: %#v", cmd)
	}
}
