package identitymapper

import (
	"testing"
	"github.com/davecgh/go-spew/spew"
	"k8s.io/apimachinery/pkg/runtime"
	clienttesting "k8s.io/client-go/testing"
	userv1fakeclient "github.com/openshift/client-go/user/clientset/versioned/fake"
	authapi "github.com/openshift/origin/pkg/oauthserver/api"
)

func TestLookup(t *testing.T) {
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
	testcases := map[string]struct {
		ProviderName		string
		ProviderUserName	string
		ExistingResources	[]runtime.Object
		ValidateActions		func(t *testing.T, actions []clienttesting.Action)
		ExpectedError		bool
		ExpectedUserName	string
	}{"no valid mapping": {ProviderName: "idp", ProviderUserName: "bob", ExistingResources: []runtime.Object{makeUser("bobUserUID", "bob", "idp:bob")}, ValidateActions: func(t *testing.T, actions []clienttesting.Action) {
		if len(actions) != 1 {
			t.Fatal(spew.Sdump(actions))
		}
		if !actions[0].Matches("get", "useridentitymappings") {
			t.Error(spew.Sdump(actions))
		}
	}, ExpectedError: true}, "existing identity, user reference": {ProviderName: "idp", ProviderUserName: "bob", ExistingResources: []runtime.Object{makeUserIdentityMapping("bobIdentityUID", "idp", "bob", "bobUserUID", "bob"), makeUser("bobUserUID", "bob", "idp:bob")}, ValidateActions: func(t *testing.T, actions []clienttesting.Action) {
		if len(actions) != 2 {
			t.Fatal(spew.Sdump(actions))
		}
		if !actions[0].Matches("get", "useridentitymappings") {
			t.Error(spew.Sdump(actions))
		}
		if !actions[1].Matches("get", "users") || actions[1].(clienttesting.GetAction).GetName() != "bob" {
			t.Error(spew.Sdump(actions))
		}
	}, ExpectedUserName: "bob"}}
	for k, tc := range testcases {
		t.Run(k, func(t *testing.T) {
			fakeClientSet := userv1fakeclient.NewSimpleClientset(tc.ExistingResources...)
			lookupMapper := &lookupIdentityMapper{mappings: fakeClientSet.UserV1().UserIdentityMappings(), users: fakeClientSet.UserV1().Users()}
			identity := authapi.NewDefaultUserIdentityInfo(tc.ProviderName, tc.ProviderUserName)
			user, err := lookupMapper.UserFor(identity)
			if tc.ExpectedError != (err != nil) {
				t.Fatalf("Expected error=%v, got %v", tc.ExpectedError, err)
			}
			if !tc.ExpectedError && user.GetName() != tc.ExpectedUserName {
				t.Fatalf("Expected username %v, got %v", tc.ExpectedUserName, user.GetName())
			}
			tc.ValidateActions(t, fakeClientSet.Actions())
		})
	}
}
