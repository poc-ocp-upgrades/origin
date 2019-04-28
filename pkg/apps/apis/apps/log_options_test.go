package apps

import (
	"reflect"
	"testing"
	kapi "k8s.io/kubernetes/pkg/apis/core"
)

func TestLogOptionsDrift(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	popts := reflect.TypeOf(kapi.PodLogOptions{})
	dopts := reflect.TypeOf(DeploymentLogOptions{})
	for i := 0; i < popts.NumField(); i++ {
		name := popts.Field(i).Name
		doptsField, found := dopts.FieldByName(name)
		if !found {
			t.Errorf("deploymentLogOptions drifting from podLogOptions! Field %q wasn't found!", name)
		}
		if should, is := popts.Field(i).Type, doptsField.Type; is != should {
			t.Errorf("deploymentLogOptions drifting from podLogOptions! Field %q should be a %s but is %s!", name, should.String(), is.String())
		}
	}
}
