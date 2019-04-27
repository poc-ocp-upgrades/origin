package scope

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	authorizationapi "github.com/openshift/origin/pkg/authorization/apis/authorization"
)

func TestUserEvaluator(t *testing.T) {
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
	testCases := []struct {
		name		string
		scopes		[]string
		err		string
		numRules	int
	}{{name: "missing-part", scopes: []string{UserIndicator}, err: "no scope evaluator found", numRules: 1}, {name: "bad-part", scopes: []string{UserIndicator + "foo"}, err: "no scope evaluator found", numRules: 1}, {name: "info", scopes: []string{UserInfo}, numRules: 2}, {name: "one-error", scopes: []string{UserIndicator, UserInfo}, err: "no scope evaluator found", numRules: 2}, {name: "access", scopes: []string{UserAccessCheck}, numRules: 3}, {name: "both", scopes: []string{UserInfo, UserAccessCheck}, numRules: 4}, {name: "list--scoped-projects", scopes: []string{UserListScopedProjects}, numRules: 2}}
	for _, tc := range testCases {
		actualRules, actualErr := ScopesToRules(tc.scopes, "namespace", nil)
		switch {
		case len(tc.err) == 0 && actualErr == nil:
		case len(tc.err) == 0 && actualErr != nil:
			t.Errorf("%s: unexpected error: %v", tc.name, actualErr)
		case len(tc.err) != 0 && actualErr == nil:
			t.Errorf("%s: missing error: %v", tc.name, tc.err)
		case len(tc.err) != 0 && actualErr != nil:
			if !strings.Contains(actualErr.Error(), tc.err) {
				t.Errorf("%s: expected %v, got %v", tc.name, tc.err, actualErr)
			}
		}
		if len(actualRules) != tc.numRules {
			t.Errorf("%s: expected %v, got %v", tc.name, tc.numRules, len(actualRules))
		}
	}
}
func TestClusterRoleEvaluator(t *testing.T) {
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
	testCases := []struct {
		name		string
		scopes		[]string
		namespace	string
		clusterRoles	[]rbacv1.ClusterRole
		policyGetterErr	error
		numRules	int
		err		string
	}{{name: "bad-format-1", scopes: []string{ClusterRoleIndicator}, err: "bad format for", numRules: 1}, {name: "bad-format-2", scopes: []string{ClusterRoleIndicator + "foo"}, err: "bad format for", numRules: 1}, {name: "bad-format-3", scopes: []string{ClusterRoleIndicator + ":ns"}, err: "bad format for", numRules: 1}, {name: "bad-format-4", scopes: []string{ClusterRoleIndicator + "foo:"}, err: "bad format for", numRules: 1}, {name: "missing-role", policyGetterErr: fmt.Errorf(`clusterrole "missing" not found`), clusterRoles: []rbacv1.ClusterRole{{ObjectMeta: metav1.ObjectMeta{Name: "admin"}, Rules: []rbacv1.PolicyRule{{}}}}, scopes: []string{ClusterRoleIndicator + "missing:*"}, err: `clusterrole "missing" not found`, numRules: 1}, {name: "mismatched-namespace", clusterRoles: []rbacv1.ClusterRole{{ObjectMeta: metav1.ObjectMeta{Name: "admin"}, Rules: []rbacv1.PolicyRule{{}}}}, namespace: "current-ns", scopes: []string{ClusterRoleIndicator + "admin:mismatch"}, numRules: 1}, {name: "all-namespaces", clusterRoles: []rbacv1.ClusterRole{{ObjectMeta: metav1.ObjectMeta{Name: "admin"}, Rules: []rbacv1.PolicyRule{{}}}}, namespace: "current-ns", scopes: []string{ClusterRoleIndicator + "admin:*"}, numRules: 2}, {name: "matching-namespaces", clusterRoles: []rbacv1.ClusterRole{{ObjectMeta: metav1.ObjectMeta{Name: "admin"}, Rules: []rbacv1.PolicyRule{{}}}}, namespace: "current-ns", scopes: []string{ClusterRoleIndicator + "admin:current-ns"}, numRules: 2}, {name: "colon-role", clusterRoles: []rbacv1.ClusterRole{{ObjectMeta: metav1.ObjectMeta{Name: "admin:two"}, Rules: []rbacv1.PolicyRule{{}}}}, namespace: "current-ns", scopes: []string{ClusterRoleIndicator + "admin:two:current-ns"}, numRules: 2}, {name: "getter-error", policyGetterErr: fmt.Errorf("some bad thing happened"), namespace: "current-ns", scopes: []string{ClusterRoleIndicator + "admin:two:current-ns"}, err: `some bad thing happened`, numRules: 1}}
	for _, tc := range testCases {
		actualRules, actualErr := ScopesToRules(tc.scopes, tc.namespace, &fakePolicyGetter{clusterRoles: tc.clusterRoles, err: tc.policyGetterErr})
		switch {
		case len(tc.err) == 0 && actualErr == nil:
		case len(tc.err) == 0 && actualErr != nil:
			t.Errorf("%s: unexpected error: %v", tc.name, actualErr)
		case len(tc.err) != 0 && actualErr == nil:
			t.Errorf("%s: missing error: %v", tc.name, tc.err)
		case len(tc.err) != 0 && actualErr != nil:
			if !strings.Contains(actualErr.Error(), tc.err) {
				t.Errorf("%s: expected %v, got %v", tc.name, tc.err, actualErr)
			}
		}
		if len(actualRules) != tc.numRules {
			t.Errorf("%s: expected %v, got %v", tc.name, tc.numRules, len(actualRules))
		}
	}
}
func TestEscalationProtection(t *testing.T) {
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
	testCases := []struct {
		name		string
		scopes		[]string
		namespace	string
		clusterRoles	[]rbacv1.ClusterRole
		expectedRules	[]rbacv1.PolicyRule
	}{{name: "simple match secrets", clusterRoles: []rbacv1.ClusterRole{{ObjectMeta: metav1.ObjectMeta{Name: "admin"}, Rules: []rbacv1.PolicyRule{{APIGroups: []string{""}, Resources: []string{"pods", "secrets"}}}}}, expectedRules: []rbacv1.PolicyRule{authorizationapi.DiscoveryRule, {APIGroups: []string{""}, Resources: []string{"pods"}}}, scopes: []string{ClusterRoleIndicator + "admin:*"}}, {name: "no longer match old group secrets", clusterRoles: []rbacv1.ClusterRole{{ObjectMeta: metav1.ObjectMeta{Name: "admin"}, Rules: []rbacv1.PolicyRule{{APIGroups: []string{}, Resources: []string{"pods", "secrets"}}}}}, expectedRules: []rbacv1.PolicyRule{authorizationapi.DiscoveryRule, {APIGroups: []string{}, Resources: []string{"pods", "secrets"}}}, scopes: []string{ClusterRoleIndicator + "admin:*"}}, {name: "skip non-matching group secrets", clusterRoles: []rbacv1.ClusterRole{{ObjectMeta: metav1.ObjectMeta{Name: "admin"}, Rules: []rbacv1.PolicyRule{{APIGroups: []string{"foo"}, Resources: []string{"pods", "secrets"}}}}}, expectedRules: []rbacv1.PolicyRule{authorizationapi.DiscoveryRule, {APIGroups: []string{"foo"}, Resources: []string{"pods", "secrets"}}}, scopes: []string{ClusterRoleIndicator + "admin:*"}}, {name: "access tokens", clusterRoles: []rbacv1.ClusterRole{{ObjectMeta: metav1.ObjectMeta{Name: "admin"}, Rules: []rbacv1.PolicyRule{{APIGroups: []string{"", "and-foo"}, Resources: []string{"pods", "oauthaccesstokens"}}}}}, expectedRules: []rbacv1.PolicyRule{authorizationapi.DiscoveryRule, {APIGroups: []string{"", "and-foo"}, Resources: []string{"pods"}}}, scopes: []string{ClusterRoleIndicator + "admin:*"}}, {name: "allow the escalation", clusterRoles: []rbacv1.ClusterRole{{ObjectMeta: metav1.ObjectMeta{Name: "admin"}, Rules: []rbacv1.PolicyRule{{APIGroups: []string{""}, Resources: []string{"pods", "secrets"}}}}}, expectedRules: []rbacv1.PolicyRule{authorizationapi.DiscoveryRule, {APIGroups: []string{""}, Resources: []string{"pods", "secrets"}}}, scopes: []string{ClusterRoleIndicator + "admin:*:!"}}}
	for _, tc := range testCases {
		actualRules, actualErr := ScopesToRules(tc.scopes, "ns-01", &fakePolicyGetter{clusterRoles: tc.clusterRoles})
		if actualErr != nil {
			t.Errorf("%s: unexpected error: %v", tc.name, actualErr)
		}
		if !reflect.DeepEqual(actualRules, tc.expectedRules) {
			t.Errorf("%s: expected %v, got %v", tc.name, tc.expectedRules, actualRules)
		}
	}
}

type fakePolicyGetter struct {
	clusterRoles	[]rbacv1.ClusterRole
	err		error
}

func (f *fakePolicyGetter) List(label labels.Selector) ([]*rbacv1.ClusterRole, error) {
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
	ret := []*rbacv1.ClusterRole{}
	for _, v := range f.clusterRoles {
		ret = append(ret, &v)
	}
	return ret, f.err
}
func (f *fakePolicyGetter) Get(id string) (*rbacv1.ClusterRole, error) {
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
	for _, v := range f.clusterRoles {
		if v.ObjectMeta.Name == id {
			return &v, nil
		}
	}
	return &rbacv1.ClusterRole{}, f.err
}
