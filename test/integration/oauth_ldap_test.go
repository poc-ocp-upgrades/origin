package integration

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apimachinery/pkg/util/uuid"
	restclient "k8s.io/client-go/rest"
	"github.com/openshift/origin/pkg/cmd/server/admin"
	configapi "github.com/openshift/origin/pkg/cmd/server/apis/config"
	configapilatest "github.com/openshift/origin/pkg/cmd/server/apis/config/latest"
	authapi "github.com/openshift/origin/pkg/oauthserver/api"
	"github.com/openshift/origin/pkg/oc/lib/tokencmd"
	userclient "github.com/openshift/origin/pkg/user/generated/internalclientset/typed/user/internalversion"
	testutil "github.com/openshift/origin/test/util"
	testserver "github.com/openshift/origin/test/util/server"
	"github.com/vjeantet/ldapserver"
)

func TestOAuthLDAP(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var (
		randomSuffix	= string(uuid.NewUUID())
		providerName	= "myldapprovider"
		bindDN		= "uid=admin,ou=company,ou=" + randomSuffix
		bindPassword	= "admin-password-" + randomSuffix
		searchDN	= "ou=company,ou=" + randomSuffix
		searchAttr	= "myuid" + randomSuffix
		searchScope	= "one"
		searchFilter	= "(myAttr=myValue)"
		nameAttr1	= "missing-name-attr"
		nameAttr2	= "a-display-name" + randomSuffix
		idAttr1		= "missing-id-attr"
		idAttr2		= "dn"
		emailAttr1	= "missing-attr"
		emailAttr2	= "c-mail" + randomSuffix
		loginAttr1	= "missing-attr"
		loginAttr2	= "d-mylogin" + randomSuffix
		myUserUID	= "myuser"
		myUserName	= "My User, Jr."
		myUserEmail	= "myuser@example.com"
		myUserDN	= searchAttr + "=" + myUserUID + "," + searchDN
		myUserPassword	= "myuser-password-" + randomSuffix
	)
	expectedAttributes := [][]byte{}
	for _, attr := range sets.NewString(searchAttr, nameAttr1, nameAttr2, idAttr1, idAttr2, emailAttr1, emailAttr2, loginAttr1, loginAttr2).List() {
		expectedAttributes = append(expectedAttributes, []byte(attr))
	}
	expectedSearchRequest := ldapserver.SearchRequest{BaseObject: []byte(searchDN), Scope: ldapserver.SearchRequestSingleLevel, DerefAliases: 0, SizeLimit: 2, TimeLimit: 0, TypesOnly: false, Attributes: expectedAttributes, Filter: fmt.Sprintf("(&%s(%s=%s))", searchFilter, searchAttr, myUserUID)}
	ldapAddress, err := testserver.FindAvailableBindAddress(8389, 8400)
	if err != nil {
		t.Fatalf("could not allocate LDAP bind address: %v", err)
	}
	ldapServer := testutil.NewTestLDAPServer()
	ldapServer.SetPassword(bindDN, bindPassword)
	ldapServer.Start(ldapAddress)
	defer ldapServer.Stop()
	masterOptions, err := testserver.DefaultMasterOptions()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer testserver.CleanupMasterEtcd(t, masterOptions)
	bindPasswordFile, err := ioutil.TempFile("", "bindPassword")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer os.Remove(bindPasswordFile.Name())
	bindPasswordKeyFile, err := ioutil.TempFile("", "bindPasswordKey")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer os.Remove(bindPasswordKeyFile.Name())
	encryptOpts := &admin.EncryptOptions{CleartextData: []byte(bindPassword), EncryptedFile: bindPasswordFile.Name(), GenKeyFile: bindPasswordKeyFile.Name()}
	if err := encryptOpts.Encrypt(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	masterOptions.OAuthConfig.IdentityProviders[0] = configapi.IdentityProvider{Name: providerName, UseAsChallenger: true, UseAsLogin: true, MappingMethod: "claim", Provider: &configapi.LDAPPasswordIdentityProvider{URL: fmt.Sprintf("ldap://%s/%s?%s?%s?%s", ldapAddress, searchDN, searchAttr, searchScope, searchFilter), BindDN: bindDN, BindPassword: configapi.StringSource{StringSourceSpec: configapi.StringSourceSpec{File: bindPasswordFile.Name(), KeyFile: bindPasswordKeyFile.Name()}}, Insecure: true, CA: "", Attributes: configapi.LDAPAttributeMapping{ID: []string{idAttr1, idAttr2}, PreferredUsername: []string{loginAttr1, loginAttr2}, Name: []string{nameAttr1, nameAttr2}, Email: []string{emailAttr1, emailAttr2}}}}
	serializedOptions, err := configapilatest.WriteYAML(masterOptions)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	deserializedObject, err := configapilatest.ReadYAML(bytes.NewBuffer(serializedOptions))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if deserializedOptions, ok := deserializedObject.(*configapi.MasterConfig); !ok {
		t.Fatalf("unexpected object: %v", deserializedObject)
	} else {
		masterOptions = deserializedOptions
	}
	clusterAdminKubeConfig, err := testserver.StartConfiguredMaster(masterOptions)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	clusterAdminClientConfig, err := testutil.GetClusterAdminClientConfig(clusterAdminKubeConfig)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	anonConfig := restclient.Config{}
	anonConfig.Host = clusterAdminClientConfig.Host
	anonConfig.CAFile = clusterAdminClientConfig.CAFile
	anonConfig.CAData = clusterAdminClientConfig.CAData
	ldapServer.ResetRequests()
	if _, err := tokencmd.RequestToken(&anonConfig, nil, myUserUID, myUserPassword); err == nil {
		t.Error("Expected error, got none")
	}
	if len(ldapServer.BindRequests) != 1 {
		t.Errorf("Expected a single bind request for the search phase, got %d:\n%#v", len(ldapServer.BindRequests), ldapServer.BindRequests)
	}
	if len(ldapServer.SearchRequests) != 1 {
		t.Errorf("Expected a single search request, got %d:\n%#v", len(ldapServer.BindRequests), ldapServer.BindRequests)
	}
	ldapServer.SetPassword(myUserDN, myUserPassword)
	ldapServer.AddSearchResult(myUserDN, map[string]string{emailAttr2: myUserEmail, nameAttr2: myUserName, loginAttr2: myUserUID})
	ldapServer.ResetRequests()
	if _, err := tokencmd.RequestToken(&anonConfig, nil, myUserUID, "badpassword"); err == nil {
		t.Error("Expected error, got none")
	}
	if len(ldapServer.BindRequests) != 2 {
		t.Errorf("Expected a bind request for the search phase and a failed bind request for the auth phase, got %d:\n%#v", len(ldapServer.BindRequests), ldapServer.BindRequests)
	}
	if len(ldapServer.SearchRequests) != 1 {
		t.Errorf("Expected a single search request, got %d:\n%#v", len(ldapServer.BindRequests), ldapServer.BindRequests)
	}
	ldapServer.ResetRequests()
	accessToken, err := tokencmd.RequestToken(&anonConfig, nil, myUserUID, myUserPassword)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if len(accessToken) == 0 {
		t.Errorf("Expected access token, got none")
	}
	if len(ldapServer.BindRequests) != 2 {
		t.Errorf("Expected a bind request for the search phase and a failed bind request for the auth phase, got %d:\n%#v", len(ldapServer.BindRequests), ldapServer.BindRequests)
	}
	if len(ldapServer.SearchRequests) != 1 {
		t.Errorf("Expected a single search request, got %d:\n%#v", len(ldapServer.BindRequests), ldapServer.BindRequests)
	}
	if !reflect.DeepEqual(expectedSearchRequest.BaseObject, ldapServer.SearchRequests[0].BaseObject) {
		t.Errorf("Expected search base DN\n\t%#v\ngot\n\t%#v", string(expectedSearchRequest.BaseObject), string(ldapServer.SearchRequests[0].BaseObject))
	}
	if !reflect.DeepEqual(expectedSearchRequest.Filter, ldapServer.SearchRequests[0].Filter) {
		t.Errorf("Expected search filter\n\t%#v\ngot\n\t%#v", string(expectedSearchRequest.Filter), string(ldapServer.SearchRequests[0].Filter))
	}
	{
		expectedAttrs := []string{}
		for _, a := range expectedSearchRequest.Attributes {
			expectedAttrs = append(expectedAttrs, string(a))
		}
		actualAttrs := []string{}
		for _, a := range ldapServer.SearchRequests[0].Attributes {
			actualAttrs = append(actualAttrs, string(a))
		}
		if !reflect.DeepEqual(expectedAttrs, actualAttrs) {
			t.Errorf("Expected search attributes\n\t%#v\ngot\n\t%#v", expectedAttrs, actualAttrs)
		}
	}
	userConfig := anonConfig
	userConfig.BearerToken = accessToken
	user, err := userclient.NewForConfigOrDie(&userConfig).Users().Get("~", metav1.GetOptions{})
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if user.Name != myUserUID {
		t.Fatalf("Expected %s as the user, got %v", myUserUID, user)
	}
	identity, err := userclient.NewForConfigOrDie(clusterAdminClientConfig).Identities().Get(fmt.Sprintf("%s:%s", providerName, myUserDN), metav1.GetOptions{})
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if identity.ProviderUserName != myUserDN {
		t.Errorf("Expected %q, got %q", myUserDN, identity.ProviderUserName)
	}
	if v := identity.Extra[authapi.IdentityDisplayNameKey]; v != myUserName {
		t.Errorf("Expected %q, got %q", myUserName, v)
	}
	if v := identity.Extra[authapi.IdentityPreferredUsernameKey]; v != myUserUID {
		t.Errorf("Expected %q, got %q", myUserUID, v)
	}
	if v := identity.Extra[authapi.IdentityEmailKey]; v != myUserEmail {
		t.Errorf("Expected %q, got %q", myUserEmail, v)
	}
}
