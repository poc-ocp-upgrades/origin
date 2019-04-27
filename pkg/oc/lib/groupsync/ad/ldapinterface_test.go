package ad

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
	"gopkg.in/ldap.v2"
	"github.com/openshift/origin/pkg/oauthserver/ldaputil"
	"github.com/openshift/origin/pkg/oauthserver/ldaputil/testclient"
)

func newTestADLDAPInterface(client ldap.Client) *ADLDAPInterface {
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
	userQuery := ldaputil.LDAPQuery{BaseDN: "ou=users,dc=example,dc=com", Scope: ldaputil.ScopeWholeSubtree, DerefAliases: ldaputil.DerefAliasesAlways, TimeLimit: 0, Filter: "objectClass=inetOrgPerson"}
	groupMembershipAttributes := []string{"memberOf"}
	userNameAttributes := []string{"cn"}
	return NewADLDAPInterface(testclient.NewConfig(client), userQuery, groupMembershipAttributes, userNameAttributes)
}
func newTestUser(CN, groupUID string) *ldap.Entry {
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
	return ldap.NewEntry(fmt.Sprintf("cn=%s,ou=users,dc=example,dc=com", CN), map[string][]string{"cn": {CN}, "memberOf": {groupUID}})
}
func TestExtractMembers(t *testing.T) {
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
	var testCases = []struct {
		name		string
		cacheSeed	map[string][]*ldap.Entry
		client		ldap.Client
		expectedError	error
		expectedMembers	[]*ldap.Entry
	}{{name: "members cached", cacheSeed: map[string][]*ldap.Entry{"testGroup": {newTestUser("testUser", "testGroup")}}, expectedError: nil, expectedMembers: []*ldap.Entry{newTestUser("testUser", "testGroup")}}, {name: "user query error", client: testclient.NewMatchingSearchErrorClient(testclient.New(), "ou=users,dc=example,dc=com", errors.New("generic search error")), expectedError: errors.New("generic search error"), expectedMembers: nil}, {name: "no errors", client: testclient.NewDNMappingClient(testclient.New(), map[string][]*ldap.Entry{"ou=users,dc=example,dc=com": {newTestUser("testUser", "testGroup")}}), expectedError: nil, expectedMembers: []*ldap.Entry{newTestUser("testUser", "testGroup")}}}
	for _, testCase := range testCases {
		ldapInterface := newTestADLDAPInterface(testCase.client)
		if len(testCase.cacheSeed) > 0 {
			ldapInterface.ldapGroupToLDAPMembers = testCase.cacheSeed
		}
		members, err := ldapInterface.ExtractMembers("testGroup")
		if !reflect.DeepEqual(err, testCase.expectedError) {
			t.Errorf("%s: incorrect error returned:\n\texpected:\n\t%v\n\tgot:\n\t%v\n", testCase.name, testCase.expectedError, err)
		}
		if !reflect.DeepEqual(members, testCase.expectedMembers) {
			t.Errorf("%s: incorrect members returned:\n\texpected:\n\t%v\n\tgot:\n\t%v\n", testCase.name, testCase.expectedMembers, members)
		}
	}
}
func TestListGroups(t *testing.T) {
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
	client := testclient.NewDNMappingClient(testclient.New(), map[string][]*ldap.Entry{"ou=users,dc=example,dc=com": {newTestUser("testUser", "testGroup")}})
	ldapInterface := newTestADLDAPInterface(client)
	groups, err := ldapInterface.ListGroups()
	if !reflect.DeepEqual(err, nil) {
		t.Errorf("listing groups: incorrect error returned:\n\texpected:\n\t%v\n\tgot:\n\t%v\n", nil, err)
	}
	if !reflect.DeepEqual(groups, []string{"testGroup"}) {
		t.Errorf("listing groups: incorrect group list:\n\texpected:\n\t%v\n\tgot:\n\t%v\n", []string{"testGroup"}, groups)
	}
}
func TestPopulateCache(t *testing.T) {
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
	var testCases = []struct {
		name			string
		cacheSeed		map[string][]*ldap.Entry
		searchDNOverride	string
		client			ldap.Client
		expectedError		error
		expectedCache		map[string][]*ldap.Entry
	}{{name: "cache already populated", cacheSeed: map[string][]*ldap.Entry{"testGroup": {newTestUser("testUser", "testGroup")}}, expectedError: nil, expectedCache: map[string][]*ldap.Entry{"testGroup": {newTestUser("testUser", "testGroup")}}}, {name: "user query error", client: testclient.NewMatchingSearchErrorClient(testclient.New(), "ou=users,dc=example,dc=com", errors.New("generic search error")), expectedError: errors.New("generic search error"), expectedCache: make(map[string][]*ldap.Entry)}, {name: "cache populated correctly", client: testclient.NewDNMappingClient(testclient.New(), map[string][]*ldap.Entry{"ou=users,dc=example,dc=com": {newTestUser("testUser", "testGroup")}}), expectedError: nil, expectedCache: map[string][]*ldap.Entry{"testGroup": {newTestUser("testUser", "testGroup")}}}}
	for _, testCase := range testCases {
		ldapInterface := newTestADLDAPInterface(testCase.client)
		if len(testCase.cacheSeed) > 0 {
			ldapInterface.ldapGroupToLDAPMembers = testCase.cacheSeed
			ldapInterface.cacheFullyPopulated = true
		}
		err := ldapInterface.populateCache()
		if !reflect.DeepEqual(err, testCase.expectedError) {
			t.Errorf("%s: incorrect error returned:\n\texpected:\n\t%v\n\tgot:\n\t%v\n", testCase.name, testCase.expectedError, err)
		}
		if !reflect.DeepEqual(testCase.expectedCache, ldapInterface.ldapGroupToLDAPMembers) {
			t.Errorf("%s: incorrect cache state:\n\texpected:\n\t%v\n\tgot:\n\t%v\n", testCase.name, testCase.expectedCache, ldapInterface.ldapGroupToLDAPMembers)
		}
	}
}
func TestPopulateCacheAfterExtractMembers(t *testing.T) {
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
	client := testclient.NewDNMappingClient(testclient.New(), map[string][]*ldap.Entry{"ou=users,dc=example,dc=com": {newTestUser("testUser", "testGroup")}})
	ldapInterface := newTestADLDAPInterface(client)
	_, err := ldapInterface.ExtractMembers("testGroup")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	client.(*testclient.DNMappingClient).DNMapping["ou=users,dc=example,dc=com"] = []*ldap.Entry{newTestUser("testUser", "testGroup"), newTestUser("testUser2", "testGroup2")}
	expectedCache := map[string][]*ldap.Entry{"testGroup": {newTestUser("testUser", "testGroup")}, "testGroup2": {newTestUser("testUser2", "testGroup2")}}
	err = ldapInterface.populateCache()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if !reflect.DeepEqual(expectedCache, ldapInterface.ldapGroupToLDAPMembers) {
		t.Errorf("incorrect cache state:\n\texpected:\n\t%v\n\tgot:\n\t%v\n", expectedCache, ldapInterface.ldapGroupToLDAPMembers)
	}
}
