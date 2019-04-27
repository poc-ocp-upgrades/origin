package util

import (
	"reflect"
	"testing"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/diff"
	"github.com/google/gofuzz"
	projectapi "github.com/openshift/origin/pkg/project/apis/project"
)

func TestProjectFidelity(t *testing.T) {
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
	f := fuzz.New().NilChance(0)
	p := &projectapi.Project{}
	for i := 0; i < 100; i++ {
		f.Fuzz(p)
		p.TypeMeta = metav1.TypeMeta{}
		namespace := ConvertProjectToExternal(p)
		p2 := ConvertNamespaceFromExternal(namespace)
		if !reflect.DeepEqual(p, p2) {
			t.Errorf("project data not preserved; the diff is %s", diff.ObjectDiff(p, p2))
		}
	}
}
func TestNamespaceFidelity(t *testing.T) {
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
	f := fuzz.New().NilChance(0)
	n := &corev1.Namespace{}
	for i := 0; i < 100; i++ {
		f.Fuzz(n)
		n.TypeMeta = metav1.TypeMeta{}
		project := ConvertNamespaceFromExternal(n)
		n2 := ConvertProjectToExternal(project)
		if !reflect.DeepEqual(n, n2) {
			t.Errorf("namespace data not preserved; the diff is %s", diff.ObjectDiff(n, n2))
		}
	}
}
