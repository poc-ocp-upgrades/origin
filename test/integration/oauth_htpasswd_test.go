package integration

import (
	"io/ioutil"
	"os"
	"testing"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	restclient "k8s.io/client-go/rest"
	configapi "github.com/openshift/origin/pkg/cmd/server/apis/config"
	"github.com/openshift/origin/pkg/oc/lib/tokencmd"
	userclient "github.com/openshift/origin/pkg/user/generated/internalclientset/typed/user/internalversion"
	testutil "github.com/openshift/origin/test/util"
	testserver "github.com/openshift/origin/test/util/server"
)

func TestOAuthHTPasswd(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	htpasswdFile, err := ioutil.TempFile("", "test.htpasswd")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer os.Remove(htpasswdFile.Name())
	masterOptions, err := testserver.DefaultMasterOptions()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer testserver.CleanupMasterEtcd(t, masterOptions)
	masterOptions.OAuthConfig.IdentityProviders[0] = configapi.IdentityProvider{Name: "htpasswd", UseAsChallenger: true, UseAsLogin: true, MappingMethod: "claim", Provider: &configapi.HTPasswdPasswordIdentityProvider{File: htpasswdFile.Name()}}
	clusterAdminKubeConfig, err := testserver.StartConfiguredMaster(masterOptions)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	clientConfig, err := testutil.GetClusterAdminClientConfig(clusterAdminKubeConfig)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	anonConfig := restclient.Config{}
	anonConfig.Host = clientConfig.Host
	anonConfig.CAFile = clientConfig.CAFile
	anonConfig.CAData = clientConfig.CAData
	if _, err := tokencmd.RequestToken(&anonConfig, nil, "username", "password"); err == nil {
		t.Error("Expected error, got none")
	}
	userpass := "username:$apr1$4Ci5I8yc$85R9vc4fOgzAULsldiUuv."
	ioutil.WriteFile(htpasswdFile.Name(), []byte(userpass), os.FileMode(0600))
	accessToken, err := tokencmd.RequestToken(&anonConfig, nil, "username", "password")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if len(accessToken) == 0 {
		t.Errorf("Expected access token, got none")
	}
	userConfig := anonConfig
	userConfig.BearerToken = accessToken
	userClient, err := userclient.NewForConfig(&userConfig)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	user, err := userClient.Users().Get("~", metav1.GetOptions{})
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if user.Name != "username" {
		t.Fatalf("Expected username as the user, got %v", user)
	}
}
