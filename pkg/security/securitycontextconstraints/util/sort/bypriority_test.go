package sort

import (
	"sort"
	"testing"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	securityv1 "github.com/openshift/api/security/v1"
)

func TestByPriority(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	sccs := []*securityv1.SecurityContextConstraints{testSCC("one", 1), testSCC("two", 2), testSCC("three", 3), testSCC("negative", -1), testSCC("super", 100)}
	expected := []string{"super", "three", "two", "one", "negative"}
	sort.Sort(ByPriority(sccs))
	for i, scc := range sccs {
		if scc.Name != expected[i] {
			t.Errorf("sort by priority found %s at element %d but expected %s", scc.Name, i, expected[i])
		}
	}
}
func TestByPrioritiesScore(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	privilegedSCC := testSCC("privileged", 1)
	privilegedSCC.AllowPrivilegedContainer = true
	nonPriviledSCC := testSCC("nonprivileged", 1)
	hostDirSCC := testSCC("hostdir", 1)
	hostDirSCC.Volumes = []securityv1.FSType{securityv1.FSTypeHostPath}
	sccs := []*securityv1.SecurityContextConstraints{nonPriviledSCC, privilegedSCC, hostDirSCC}
	expected := []string{"nonprivileged", "hostdir", "privileged"}
	sort.Sort(ByPriority(sccs))
	for i, scc := range sccs {
		if scc.Name != expected[i] {
			t.Errorf("sort by score found %s at element %d but expected %s", scc.Name, i, expected[i])
		}
	}
}
func TestByPrioritiesName(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	sccs := []*securityv1.SecurityContextConstraints{testSCC("e", 1), testSCC("d", 1), testSCC("a", 1), testSCC("c", 1), testSCC("b", 1)}
	expected := []string{"a", "b", "c", "d", "e"}
	sort.Sort(ByPriority(sccs))
	for i, scc := range sccs {
		if scc.Name != expected[i] {
			t.Errorf("sort by priority found %s at element %d but expected %s", scc.Name, i, expected[i])
		}
	}
}
func TestByPrioritiesMixedSCCs(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	privilegedSCC := testSCC("privileged", 1)
	privilegedSCC.AllowPrivilegedContainer = true
	nonPriviledSCC := testSCC("nonprivileged", 1)
	sccs := []*securityv1.SecurityContextConstraints{testSCC("priorityB", 5), testSCC("priorityA", 5), testSCC("super", 100), privilegedSCC, nonPriviledSCC}
	expected := []string{"super", "priorityA", "priorityB", "nonprivileged", "privileged"}
	sort.Sort(ByPriority(sccs))
	for i, scc := range sccs {
		if scc.Name != expected[i] {
			t.Errorf("sort by priority found %s at element %d but expected %s", scc.Name, i, expected[i])
		}
	}
}
func testSCC(name string, priority int) *securityv1.SecurityContextConstraints {
	_logClusterCodePath()
	defer _logClusterCodePath()
	newPriority := int32(priority)
	return &securityv1.SecurityContextConstraints{ObjectMeta: metav1.ObjectMeta{Name: name}, Priority: &newPriority}
}
