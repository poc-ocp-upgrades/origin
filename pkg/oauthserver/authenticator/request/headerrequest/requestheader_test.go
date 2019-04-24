package headerrequest

import (
	"net/http"
	"reflect"
	"testing"
	"github.com/openshift/origin/pkg/oauthserver/api"
	"k8s.io/apiserver/pkg/authentication/user"
)

type TestUserIdentityMapper struct{ Identity api.UserIdentityInfo }

func (m *TestUserIdentityMapper) UserFor(identityInfo api.UserIdentityInfo) (user.Info, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	m.Identity = identityInfo
	username := identityInfo.GetProviderUserName()
	if preferredUsername := identityInfo.GetExtra()[api.IdentityPreferredUsernameKey]; len(preferredUsername) > 0 {
		username = preferredUsername
	}
	return &user.DefaultInfo{Name: username}, nil
}
func TestRequestHeader(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	testcases := map[string]struct {
		Config			Config
		RequestHeaders		http.Header
		ExpectedUsername	string
		ExpectedIdentity	api.UserIdentityInfo
	}{"empty": {ExpectedUsername: ""}, "no match": {Config: Config{IDHeaders: []string{"X-Remote-User"}}, ExpectedUsername: ""}, "match": {Config: Config{IDHeaders: []string{"X-Remote-User"}}, RequestHeaders: http.Header{"X-Remote-User": {"Bob"}}, ExpectedUsername: "Bob"}, "exact match": {Config: Config{IDHeaders: []string{"X-Remote-User"}}, RequestHeaders: http.Header{"Prefixed-X-Remote-User-With-Suffix": {"Bob"}, "X-Remote-User-With-Suffix": {"Bob"}}, ExpectedUsername: ""}, "first match": {Config: Config{IDHeaders: []string{"X-Remote-User", "A-Second-X-Remote-User", "Another-X-Remote-User"}}, RequestHeaders: http.Header{"X-Remote-User": {"", "First header, second value"}, "A-Second-X-Remote-User": {"Second header, first value", "Second header, second value"}, "Another-X-Remote-User": {"Third header, first value"}}, ExpectedUsername: "Second header, first value"}, "case-insensitive": {Config: Config{IDHeaders: []string{"x-REMOTE-user"}}, RequestHeaders: http.Header{"X-Remote-User": {"Bob"}}, ExpectedUsername: "Bob"}, "extended attributes": {Config: Config{IDHeaders: []string{"x-id", "x-id2"}, PreferredUsernameHeaders: []string{"x-preferred-username", "x-preferred-username2"}, EmailHeaders: []string{"x-email", "x-email2"}, NameHeaders: []string{"x-name", "x-name2"}}, RequestHeaders: http.Header{"X-Id2": {"12345"}, "X-Preferred-Username2": {"bob"}, "X-Email2": {"bob@example.com"}, "X-Name2": {"Bob"}}, ExpectedUsername: "bob", ExpectedIdentity: &api.DefaultUserIdentityInfo{ProviderName: "testprovider", ProviderUserName: "12345", Extra: map[string]string{api.IdentityDisplayNameKey: "Bob", api.IdentityEmailKey: "bob@example.com", api.IdentityPreferredUsernameKey: "bob"}}}}
	for k, testcase := range testcases {
		mapper := &TestUserIdentityMapper{}
		auth := NewAuthenticator("testprovider", &testcase.Config, mapper)
		req := &http.Request{Header: testcase.RequestHeaders}
		authResponse, ok, err := auth.AuthenticateRequest(req)
		if testcase.ExpectedUsername == "" {
			if ok {
				t.Errorf("%s: Didn't expect user, authentication succeeded", k)
				continue
			}
		}
		if testcase.ExpectedUsername != "" {
			if err != nil {
				t.Errorf("%s: Expected user, got error: %v", k, err)
				continue
			}
			if !ok {
				t.Errorf("%s: Expected user, auth failed", k)
				continue
			}
			if testcase.ExpectedUsername != authResponse.User.GetName() {
				t.Errorf("%s: Expected username %s, got %s", k, testcase.ExpectedUsername, authResponse.User.GetName())
				continue
			}
		}
		if testcase.ExpectedIdentity != nil {
			if !reflect.DeepEqual(testcase.ExpectedIdentity.GetExtra(), mapper.Identity.GetExtra()) {
				t.Errorf("%s: Expected %#v, got %#v", k, testcase.ExpectedIdentity.GetExtra(), mapper.Identity.GetExtra())
			}
			if !reflect.DeepEqual(testcase.ExpectedIdentity.GetProviderUserName(), mapper.Identity.GetProviderUserName()) {
				t.Errorf("%s: Expected %#v, got %#v", k, testcase.ExpectedIdentity.GetProviderUserName(), mapper.Identity.GetProviderUserName())
			}
		}
	}
}
