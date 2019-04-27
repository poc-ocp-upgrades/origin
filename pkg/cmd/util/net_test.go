package util

import (
	"reflect"
	"testing"
)

func TestHostnameMatchSpecCandidates(t *testing.T) {
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
	testcases := []struct {
		Hostname	string
		ExpectedSpecs	[]string
	}{{Hostname: "", ExpectedSpecs: nil}, {Hostname: "a", ExpectedSpecs: []string{"a", "*"}}, {Hostname: "foo.bar", ExpectedSpecs: []string{"foo.bar", "*.bar", "*.*"}}}
	for _, tc := range testcases {
		specs := HostnameMatchSpecCandidates(tc.Hostname)
		if !reflect.DeepEqual(specs, tc.ExpectedSpecs) {
			t.Errorf("%s: Expected %#v, got %#v", tc.Hostname, tc.ExpectedSpecs, specs)
		}
	}
}
func TestHostnameMatches(t *testing.T) {
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
	testcases := []struct {
		Hostname	string
		Spec		string
		ExpectedMatch	bool
	}{{Hostname: "", Spec: "", ExpectedMatch: false}, {Hostname: "a", Spec: "", ExpectedMatch: false}, {Hostname: "a", Spec: "a", ExpectedMatch: true}, {Hostname: "a", Spec: "*", ExpectedMatch: true}, {Hostname: "a", Spec: "*.a", ExpectedMatch: false}, {Hostname: "a", Spec: "*.*", ExpectedMatch: false}, {Hostname: "a.b", Spec: "a.b", ExpectedMatch: true}, {Hostname: "a.b", Spec: "*.b", ExpectedMatch: true}, {Hostname: "a.b", Spec: "*.*", ExpectedMatch: true}, {Hostname: "a.b", Spec: "a.*", ExpectedMatch: false}, {Hostname: "a.b", Spec: "*.a.b", ExpectedMatch: false}}
	for i, tc := range testcases {
		matches := HostnameMatches(tc.Hostname, tc.Spec)
		if matches != tc.ExpectedMatch {
			t.Errorf("%d: Expected match=%v, got %v (hostname=%s, specs=%v)", i, tc.ExpectedMatch, matches, tc.Hostname, tc.Spec)
		}
	}
}
