package registrytest

import (
	"k8s.io/apiserver/pkg/registry/rest"
	"reflect"
	"testing"
)

func AssertShortNames(t *testing.T, storage rest.ShortNamesProvider, expected []string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	actual := storage.ShortNames()
	ok := reflect.DeepEqual(actual, expected)
	if !ok {
		t.Errorf("short names not equal. expected = %v actual = %v", expected, actual)
	}
}
