package build

import (
	"reflect"
	"testing"
	kapi "k8s.io/kubernetes/pkg/apis/core"
)

func TestLogOptionsDrift(t *testing.T) {
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
	popts := reflect.TypeOf(kapi.PodLogOptions{})
	bopts := reflect.TypeOf(BuildLogOptions{})
	for i := 0; i < popts.NumField(); i++ {
		name := popts.Field(i).Name
		boptsField, found := bopts.FieldByName(name)
		if !found {
			t.Errorf("buildLogOptions drifting from podLogOptions! Field %q wasn't found!", name)
		}
		if should, is := popts.Field(i).Type, boptsField.Type; is != should {
			t.Errorf("buildLogOptions drifting from podLogOptions! Field %q should be a %s but is %s!", name, should.String(), is.String())
		}
	}
}
