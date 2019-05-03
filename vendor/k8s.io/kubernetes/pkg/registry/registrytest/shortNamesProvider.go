package registrytest

import (
 "reflect"
 "testing"
 "k8s.io/apiserver/pkg/registry/rest"
)

func AssertShortNames(t *testing.T, storage rest.ShortNamesProvider, expected []string) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 actual := storage.ShortNames()
 ok := reflect.DeepEqual(actual, expected)
 if !ok {
  t.Errorf("short names not equal. expected = %v actual = %v", expected, actual)
 }
}
