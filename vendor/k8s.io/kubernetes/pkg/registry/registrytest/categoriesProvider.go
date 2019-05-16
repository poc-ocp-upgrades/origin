package registrytest

import (
	goformat "fmt"
	"k8s.io/apiserver/pkg/registry/rest"
	goos "os"
	"reflect"
	godefaultruntime "runtime"
	"testing"
	gotime "time"
)

func AssertCategories(t *testing.T, storage rest.CategoriesProvider, expected []string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	actual := storage.Categories()
	ok := reflect.DeepEqual(actual, expected)
	if !ok {
		t.Errorf("categories are not equal. expected = %v actual = %v", expected, actual)
	}
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
