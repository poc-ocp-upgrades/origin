package integration

import (
	"testing"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	restclient "k8s.io/client-go/rest"
	"github.com/openshift/origin/pkg/oc/lib/tokencmd"
	testutil "github.com/openshift/origin/test/util"
	testserver "github.com/openshift/origin/test/util/server"
)

func TestOAuthDisabled(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	masterOptions, err := testserver.DefaultMasterOptions()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer testserver.CleanupMasterEtcd(t, masterOptions)
	masterOptions.OAuthConfig = nil
	clusterAdminKubeConfig, err := testserver.StartConfiguredMaster(masterOptions)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	client, err := testutil.GetClusterAdminKubeClient(clusterAdminKubeConfig)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	clientConfig, err := testutil.GetClusterAdminClientConfig(clusterAdminKubeConfig)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	namespaces, err := client.CoreV1().Namespaces().List(metav1.ListOptions{})
	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}
	if len(namespaces.Items) == 0 {
		t.Errorf("Expected namespaces, got none")
	}
	anonConfig := restclient.Config{}
	anonConfig.Host = clientConfig.Host
	anonConfig.CAFile = clientConfig.CAFile
	anonConfig.CAData = clientConfig.CAData
	if _, err := tokencmd.RequestToken(&anonConfig, nil, "username", "password"); err == nil {
		t.Error("Expected error, got none")
	}
}
