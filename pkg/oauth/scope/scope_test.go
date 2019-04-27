package scope

import (
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
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
	checkAdd(t, []string{}, []string{}, []string{})
	checkAdd(t, []string{"A"}, []string{}, []string{"A"})
	checkAdd(t, []string{"A"}, []string{"A"}, []string{"A"})
	checkAdd(t, []string{"B", "A"}, []string{"A", "B"}, []string{"A", "B"})
	checkAdd(t, []string{"B", "A"}, []string{"C", "A", "B"}, []string{"A", "B", "C"})
	checkAdd(t, []string{}, []string{"C", "A", "B"}, []string{"A", "B", "C"})
}
func checkAdd(t *testing.T, s1, s2, expected []string) {
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
	actual := Add(s1, s2)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v + %v to be %v, but got %v", s1, s2, expected, actual)
	}
}
func TestCovers(t *testing.T) {
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
	checkCovers(t, []string{}, []string{}, true)
	checkCovers(t, []string{"A"}, []string{}, false)
	checkCovers(t, []string{"B", "A"}, []string{}, false)
	checkCovers(t, []string{}, []string{"B", "A"}, false)
	checkCovers(t, []string{"A"}, []string{"A"}, true)
	checkCovers(t, []string{"B", "A"}, []string{"A"}, true)
	checkCovers(t, []string{"B", "A"}, []string{"A", "B"}, true)
	checkCovers(t, []string{"B", "A", "C"}, []string{"A", "B"}, true)
	checkCovers(t, []string{}, []string{"A"}, false)
	checkCovers(t, []string{"B"}, []string{"A"}, false)
	checkCovers(t, []string{"A", "B"}, []string{"A", "C"}, false)
}
func checkCovers(t *testing.T, has, requested []string, expected bool) {
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
	actual := Covers(has, requested)
	if actual != expected {
		if expected {
			t.Errorf("Expected %v to cover %v, but it did not", has, requested)
		} else {
			t.Errorf("Expected %v to not cover %v, but it did", has, requested)
		}
	}
}
