package syncgroups

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
	"testing"
	"gopkg.in/ldap.v2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	clienttesting "k8s.io/client-go/testing"
	userv1 "github.com/openshift/api/user/v1"
	fakeuserclient "github.com/openshift/client-go/user/clientset/versioned/fake"
	fakeuserv1client "github.com/openshift/client-go/user/clientset/versioned/typed/user/v1/fake"
	"github.com/openshift/origin/pkg/oauthserver/ldaputil"
	"github.com/openshift/origin/pkg/oc/lib/groupsync/interfaces"
)

func TestMakeOpenShiftGroup(t *testing.T) {
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
	syncer := &LDAPGroupSyncer{Out: ioutil.Discard, Err: ioutil.Discard, Host: "test-host:port", GroupNameMapper: &TestGroupNameMapper{NameMapping: map[string]string{"alfa": "zulu"}}}
	tcs := map[string]struct {
		ldapGroupUID	string
		usernames	[]string
		startingGroups	[]runtime.Object
		expectedGroup	*userv1.Group
		expectedErr	string
	}{"bad ldapGroupUID": {ldapGroupUID: "bravo", expectedErr: "no name found for group: bravo"}, "good": {ldapGroupUID: "alfa", usernames: []string{"valerie"}, expectedGroup: &userv1.Group{ObjectMeta: metav1.ObjectMeta{Name: "zulu", Annotations: map[string]string{ldaputil.LDAPURLAnnotation: "test-host:port", ldaputil.LDAPUIDAnnotation: "alfa"}, Labels: map[string]string{ldaputil.LDAPHostLabel: "test-host"}}, Users: []string{"valerie"}}}, "replaced good": {ldapGroupUID: "alfa", usernames: []string{"valerie"}, expectedGroup: &userv1.Group{ObjectMeta: metav1.ObjectMeta{Name: "zulu", Annotations: map[string]string{ldaputil.LDAPURLAnnotation: "test-host:port", ldaputil.LDAPUIDAnnotation: "alfa"}, Labels: map[string]string{ldaputil.LDAPHostLabel: "test-host"}}, Users: []string{"valerie"}}, startingGroups: []runtime.Object{&userv1.Group{ObjectMeta: metav1.ObjectMeta{Name: "zulu", Annotations: map[string]string{ldaputil.LDAPURLAnnotation: "test-host:port", ldaputil.LDAPUIDAnnotation: "alfa"}, Labels: map[string]string{ldaputil.LDAPHostLabel: "test-host"}}, Users: []string{"other-user"}}}}, "conflicting uid": {ldapGroupUID: "alfa", usernames: []string{"valerie"}, startingGroups: []runtime.Object{&userv1.Group{ObjectMeta: metav1.ObjectMeta{Name: "zulu", Annotations: map[string]string{ldaputil.LDAPURLAnnotation: "test-host:port", ldaputil.LDAPUIDAnnotation: "bravo"}, Labels: map[string]string{ldaputil.LDAPHostLabel: "test-host"}}, Users: []string{"other-user"}}}, expectedErr: `group "zulu": openshift.io/ldap.uid annotation did not match LDAP UID: wanted alfa, got bravo`}, "conflicting host": {ldapGroupUID: "alfa", usernames: []string{"valerie"}, startingGroups: []runtime.Object{&userv1.Group{ObjectMeta: metav1.ObjectMeta{Name: "zulu", Annotations: map[string]string{ldaputil.LDAPURLAnnotation: "bad-host:port", ldaputil.LDAPUIDAnnotation: "alfa"}, Labels: map[string]string{ldaputil.LDAPHostLabel: "bad-host"}}, Users: []string{"other-user"}}}, expectedErr: `group "zulu": openshift.io/ldap.host label did not match sync host: wanted test-host, got bad-host`}, "conflicting port": {ldapGroupUID: "alfa", usernames: []string{"valerie"}, startingGroups: []runtime.Object{&userv1.Group{ObjectMeta: metav1.ObjectMeta{Name: "zulu", Annotations: map[string]string{ldaputil.LDAPURLAnnotation: "test-host:port2", ldaputil.LDAPUIDAnnotation: "alfa"}, Labels: map[string]string{ldaputil.LDAPHostLabel: "test-host"}}, Users: []string{"other-user"}}}, expectedErr: `group "zulu": openshift.io/ldap.url annotation did not match sync host: wanted test-host:port, got test-host:port2`}}
	for name, tc := range tcs {
		fakeClient := &fakeuserv1client.FakeUserV1{Fake: &(fakeuserclient.NewSimpleClientset(tc.startingGroups...).Fake)}
		syncer.GroupClient = fakeClient.Groups()
		actualGroup, err := syncer.makeOpenShiftGroup(tc.ldapGroupUID, tc.usernames)
		if err != nil && len(tc.expectedErr) == 0 {
			t.Errorf("%s: unexpected error %v", name, err)
		} else if err == nil && len(tc.expectedErr) != 0 {
			t.Errorf("%s: expected %v, got nil", name, tc.expectedErr)
		} else if err != nil {
			if e, a := tc.expectedErr, err.Error(); e != a {
				t.Errorf("%s: expected %v, got %v", name, e, a)
			}
		}
		if actualGroup != nil {
			delete(actualGroup.Annotations, ldaputil.LDAPSyncTimeAnnotation)
		}
		if !reflect.DeepEqual(tc.expectedGroup, actualGroup) {
			t.Errorf("%s: expected %v, got %v", name, tc.expectedGroup, actualGroup)
		}
	}
}

const (
	Group1UID		string	= "group1"
	Group2UID		string	= "group2"
	Group3UID		string	= "group3"
	UserNameAttribute	string	= "cn"
	Member1UID		string	= "member1"
	Member2UID		string	= "member2"
	Member3UID		string	= "member3"
	Member4UID		string	= "member4"
	BaseDN			string	= "dc=example,dc=com"
)

var Member1 *ldap.Entry = &ldap.Entry{DN: UserNameAttribute + "=" + Member1UID + "," + BaseDN, Attributes: []*ldap.EntryAttribute{{Name: UserNameAttribute, Values: []string{Member1UID}, ByteValues: [][]byte{[]byte(Member1UID)}}}}
var Member2 *ldap.Entry = &ldap.Entry{DN: UserNameAttribute + "=" + Member2UID + "," + BaseDN, Attributes: []*ldap.EntryAttribute{{Name: UserNameAttribute, Values: []string{Member2UID}, ByteValues: [][]byte{[]byte(Member2UID)}}}}
var Member3 *ldap.Entry = &ldap.Entry{DN: UserNameAttribute + "=" + Member3UID + "," + BaseDN, Attributes: []*ldap.EntryAttribute{{Name: UserNameAttribute, Values: []string{Member3UID}, ByteValues: [][]byte{[]byte(Member3UID)}}}}
var Member4 *ldap.Entry = &ldap.Entry{DN: UserNameAttribute + "=" + Member4UID + "," + BaseDN, Attributes: []*ldap.EntryAttribute{{Name: UserNameAttribute, Values: []string{Member4UID}, ByteValues: [][]byte{[]byte(Member4UID)}}}}
var Group1Members []*ldap.Entry = []*ldap.Entry{Member1, Member2}
var Group2Members []*ldap.Entry = []*ldap.Entry{Member2, Member3}
var Group3Members []*ldap.Entry = []*ldap.Entry{Member3, Member4}

func TestGoodSync(t *testing.T) {
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
	testGroupSyncer, tc := newTestSyncer()
	_, errs := testGroupSyncer.Sync()
	for _, err := range errs {
		t.Errorf("unexpected sync error: %v", err)
	}
	checkClientForGroups(tc, newDefaultOpenShiftGroups(testGroupSyncer.Host), t)
}
func TestListFails(t *testing.T) {
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
	testGroupSyncer, _ := newTestSyncer()
	testGroupSyncer.GroupLister.(*TestGroupLister).err = errors.New("error during listing")
	groups, errs := testGroupSyncer.Sync()
	if len(errs) != 1 {
		t.Errorf("unexpected sync error: %v", errs)
	} else if errs[0] != testGroupSyncer.GroupLister.(*TestGroupLister).err {
		t.Errorf("unexpected sync error: %v", errs)
	}
	if groups != nil {
		t.Errorf("unexpected groups %v", groups)
	}
}
func TestMissingLDAPGroupUIDMapping(t *testing.T) {
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
	testGroupSyncer, tc := newTestSyncer()
	testGroupSyncer.GroupLister.(*TestGroupLister).GroupUIDs = append(testGroupSyncer.GroupLister.(*TestGroupLister).GroupUIDs, "ldapgroupwithnouid")
	_, errs := testGroupSyncer.Sync()
	if len(errs) != 1 {
		t.Errorf("unexpected sync error: %v", errs)
	} else if e, a := "no members found for group: ldapgroupwithnouid", errs[0].Error(); e != a {
		t.Errorf("expected %v, got %v", e, a)
	}
	checkClientForGroups(tc, newDefaultOpenShiftGroups(testGroupSyncer.Host), t)
}
func checkClientForGroups(tc *fakeuserv1client.FakeUserV1, expectedGroups []*userv1.Group, t *testing.T) {
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
	actualGroups := extractActualGroups(tc)
	for _, expectedGroup := range expectedGroups {
		if !groupExists(actualGroups, expectedGroup) {
			t.Errorf("did not find %v, got %v", expectedGroup, actualGroups)
		}
	}
}
func groupExists(haystack []*userv1.Group, needle *userv1.Group) bool {
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
	for _, actual := range haystack {
		actualGroup := actual.DeepCopy()
		delete(actualGroup.Annotations, ldaputil.LDAPSyncTimeAnnotation)
		if reflect.DeepEqual(needle, actualGroup) {
			return true
		}
	}
	return false
}
func extractActualGroups(tc *fakeuserv1client.FakeUserV1) []*userv1.Group {
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
	ret := []*userv1.Group{}
	for _, genericAction := range tc.Actions() {
		switch action := genericAction.(type) {
		case clienttesting.CreateAction:
			ret = append(ret, action.GetObject().(*userv1.Group))
		case clienttesting.UpdateAction:
			ret = append(ret, action.GetObject().(*userv1.Group))
		}
	}
	return ret
}
func newDefaultOpenShiftGroups(host string) []*userv1.Group {
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
	return []*userv1.Group{{ObjectMeta: metav1.ObjectMeta{Name: "os" + Group1UID, Annotations: map[string]string{ldaputil.LDAPURLAnnotation: host, ldaputil.LDAPUIDAnnotation: Group1UID}, Labels: map[string]string{ldaputil.LDAPHostLabel: strings.Split(host, ":")[0]}}, Users: []string{Member1UID, Member2UID}}, {ObjectMeta: metav1.ObjectMeta{Name: "os" + Group2UID, Annotations: map[string]string{ldaputil.LDAPURLAnnotation: host, ldaputil.LDAPUIDAnnotation: Group2UID}, Labels: map[string]string{ldaputil.LDAPHostLabel: strings.Split(host, ":")[0]}}, Users: []string{Member2UID, Member3UID}}, {ObjectMeta: metav1.ObjectMeta{Name: "os" + Group3UID, Annotations: map[string]string{ldaputil.LDAPURLAnnotation: host, ldaputil.LDAPUIDAnnotation: Group3UID}, Labels: map[string]string{ldaputil.LDAPHostLabel: strings.Split(host, ":")[0]}}, Users: []string{Member3UID, Member4UID}}}
}
func newTestSyncer() (*LDAPGroupSyncer, *fakeuserv1client.FakeUserV1) {
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
	tc := &fakeuserv1client.FakeUserV1{Fake: &(fakeuserclient.NewSimpleClientset().Fake)}
	tc.PrependReactor("create", "groups", func(action clienttesting.Action) (handled bool, ret runtime.Object, err error) {
		createAction := action.(clienttesting.CreateAction)
		return true, createAction.GetObject(), nil
	})
	tc.PrependReactor("update", "groups", func(action clienttesting.Action) (handled bool, ret runtime.Object, err error) {
		updateAction := action.(clienttesting.UpdateAction)
		return true, updateAction.GetObject(), nil
	})
	return &LDAPGroupSyncer{GroupLister: newTestLister(), GroupMemberExtractor: newTestMemberExtractor(), UserNameMapper: newTestUserNameMapper(), GroupNameMapper: newTestGroupNameMapper(), GroupClient: tc.Groups(), Host: newTestHost(), Out: ioutil.Discard, Err: ioutil.Discard}, tc
}
func newTestHost() string {
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
	return "test.host:port"
}
func newTestLister() interfaces.LDAPGroupLister {
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
	return &TestGroupLister{GroupUIDs: []string{Group1UID, Group2UID, Group3UID}}
}
func newTestMemberExtractor() interfaces.LDAPMemberExtractor {
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
	return &TestGroupMemberExtractor{MemberMapping: map[string][]*ldap.Entry{Group1UID: Group1Members, Group2UID: Group2Members, Group3UID: Group3Members}}
}
func newTestUserNameMapper() interfaces.LDAPUserNameMapper {
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
	return &TestUserNameMapper{NameAttributes: []string{UserNameAttribute}}
}
func newTestGroupNameMapper() interfaces.LDAPGroupNameMapper {
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
	return &TestGroupNameMapper{NameMapping: map[string]string{Group1UID: "os" + Group1UID, Group2UID: "os" + Group2UID, Group3UID: "os" + Group3UID}}
}

var _ interfaces.LDAPGroupLister = &TestGroupLister{}
var _ interfaces.LDAPMemberExtractor = &TestGroupMemberExtractor{}
var _ interfaces.LDAPUserNameMapper = &TestUserNameMapper{}
var _ interfaces.LDAPGroupNameMapper = &TestGroupNameMapper{}

type TestGroupLister struct {
	GroupUIDs	[]string
	err		error
}

func (l *TestGroupLister) ListGroups() ([]string, error) {
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
	if l.err != nil {
		return nil, l.err
	}
	return l.GroupUIDs, nil
}

type TestGroupMemberExtractor struct{ MemberMapping map[string][]*ldap.Entry }

func (e *TestGroupMemberExtractor) ExtractMembers(ldapGroupUID string) ([]*ldap.Entry, error) {
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
	members, exist := e.MemberMapping[ldapGroupUID]
	if !exist {
		return nil, fmt.Errorf("no members found for group: %s", ldapGroupUID)
	}
	return members, nil
}

type TestUserNameMapper struct{ NameAttributes []string }

func (m *TestUserNameMapper) UserNameFor(user *ldap.Entry) (string, error) {
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
	openShiftUserName := ldaputil.GetAttributeValue(user, m.NameAttributes)
	if len(openShiftUserName) == 0 {
		return "", fmt.Errorf("the user entry (%v) does not map to a OpenShift User name with the given mapping", user)
	}
	return openShiftUserName, nil
}

type TestGroupNameMapper struct{ NameMapping map[string]string }

func (m *TestGroupNameMapper) GroupNameFor(ldapGroupUID string) (string, error) {
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
	name, exists := m.NameMapping[ldapGroupUID]
	if !exists {
		return "", fmt.Errorf("no name found for group: %s", ldapGroupUID)
	}
	return name, nil
}
