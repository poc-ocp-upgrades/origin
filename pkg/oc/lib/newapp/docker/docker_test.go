package docker

import (
	"testing"
	"github.com/spf13/pflag"
)

func TestNewHelper(t *testing.T) {
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
	flags := pflag.NewFlagSet("test", pflag.ContinueOnError)
	helper := NewHelper()
	helper.InstallFlags(flags)
}
