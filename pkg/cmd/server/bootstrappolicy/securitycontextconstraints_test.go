package bootstrappolicy

import (
	"reflect"
	"sort"
	"testing"
	"k8s.io/apiserver/pkg/authentication/serviceaccount"
	securityv1 "github.com/openshift/api/security/v1"
	sccutil "github.com/openshift/origin/pkg/security/securitycontextconstraints/util"
	sccsort "github.com/openshift/origin/pkg/security/securitycontextconstraints/util/sort"
)

func TestBootstrappedConstraints(t *testing.T) {
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
	expectedConstraintNames := []string{SecurityContextConstraintsAnyUID, SecurityContextConstraintRestricted, SecurityContextConstraintNonRoot, SecurityContextConstraintHostMountAndAnyUID, SecurityContextConstraintsHostNetwork, SecurityContextConstraintHostNS, SecurityContextConstraintPrivileged}
	expectedGroups, expectedUsers := getExpectedAccess()
	expectedVolumes := []securityv1.FSType{securityv1.FSTypeEmptyDir, securityv1.FSTypeSecret, securityv1.FSTypeDownwardAPI, securityv1.FSTypeConfigMap, securityv1.FSTypePersistentVolumeClaim}
	groups, users := GetBoostrapSCCAccess(DefaultOpenShiftInfraNamespace)
	bootstrappedConstraints := GetBootstrapSecurityContextConstraints(groups, users)
	if len(expectedConstraintNames) != len(bootstrappedConstraints) {
		t.Errorf("unexpected number of constraints: found %d, wanted %d", len(bootstrappedConstraints), len(expectedConstraintNames))
	}
	bootstrappedConstraintsExternal, err := sccsort.ByPriorityConvert(bootstrappedConstraints)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	sort.Sort(bootstrappedConstraintsExternal)
	for i, constraint := range bootstrappedConstraintsExternal {
		if constraint.Name != expectedConstraintNames[i] {
			t.Errorf("unexpected contraint no. %d (by priority).  Found %v, wanted %v", i, constraint.Name, expectedConstraintNames[i])
		}
		g := expectedGroups[constraint.Name]
		if g == nil {
			g = []string{}
		}
		if !reflect.DeepEqual(g, constraint.Groups) {
			t.Errorf("unexpected group access for %s.  Found %v, wanted %v", constraint.Name, constraint.Groups, g)
		}
		u := expectedUsers[constraint.Name]
		if u == nil {
			u = []string{}
		}
		if !reflect.DeepEqual(u, constraint.Users) {
			t.Errorf("unexpected user access for %s.  Found %v, wanted %v", constraint.Name, constraint.Users, u)
		}
		for _, expectedVolume := range expectedVolumes {
			if !sccutil.SCCAllowsFSType(constraint, expectedVolume) {
				t.Errorf("%s does not support %v which is required for all default SCCs", constraint.Name, expectedVolume)
			}
		}
	}
}
func TestBootstrappedConstraintsWithAddedUser(t *testing.T) {
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
	expectedGroups, expectedUsers := getExpectedAccess()
	groups, users := GetBoostrapSCCAccess(DefaultOpenShiftInfraNamespace)
	users[SecurityContextConstraintPrivileged] = append(users[SecurityContextConstraintPrivileged], "foo")
	bootstrappedConstraints := GetBootstrapSecurityContextConstraints(groups, users)
	expectedUsers[SecurityContextConstraintPrivileged] = append(expectedUsers[SecurityContextConstraintPrivileged], "foo")
	for _, constraint := range bootstrappedConstraints {
		g := expectedGroups[constraint.Name]
		if g == nil {
			g = []string{}
		}
		if !reflect.DeepEqual(g, constraint.Groups) {
			t.Errorf("unexpected group access for %s.  Found %v, wanted %v", constraint.Name, constraint.Groups, g)
		}
		u := expectedUsers[constraint.Name]
		if u == nil {
			u = []string{}
		}
		if !reflect.DeepEqual(u, constraint.Users) {
			t.Errorf("unexpected user access for %s.  Found %v, wanted %v", constraint.Name, constraint.Users, u)
		}
	}
}
func getExpectedAccess() (map[string][]string, map[string][]string) {
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
	groups := map[string][]string{SecurityContextConstraintPrivileged: {ClusterAdminGroup, NodesGroup, MastersGroup}, SecurityContextConstraintsAnyUID: {ClusterAdminGroup}, SecurityContextConstraintRestricted: {AuthenticatedGroup}}
	buildControllerUsername := serviceaccount.MakeUsername(DefaultOpenShiftInfraNamespace, InfraBuildControllerServiceAccountName)
	pvRecyclerControllerUsername := serviceaccount.MakeUsername(DefaultOpenShiftInfraNamespace, InfraPersistentVolumeRecyclerControllerServiceAccountName)
	users := map[string][]string{SecurityContextConstraintPrivileged: {SystemAdminUsername, buildControllerUsername}, SecurityContextConstraintHostMountAndAnyUID: {pvRecyclerControllerUsername}}
	return groups, users
}
