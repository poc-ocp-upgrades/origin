package integration

import (
	"fmt"
	"testing"
	"time"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/util/wait"
	discocache "k8s.io/client-go/discovery/cached"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/restmapper"
	"k8s.io/client-go/scale"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	"github.com/openshift/api/apps"
	appsv1 "github.com/openshift/api/apps/v1"
	appsclient "github.com/openshift/client-go/apps/clientset/versioned"
	appsutil "github.com/openshift/origin/pkg/apps/util"
	appstest "github.com/openshift/origin/pkg/apps/util/test"
	testutil "github.com/openshift/origin/test/util"
	testserver "github.com/openshift/origin/test/util/server"
)

func TestDeployScale(t *testing.T) {
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
	const namespace = "test-deploy-scale"
	masterConfig, clusterAdminKubeConfig, err := testserver.StartTestMaster()
	if err != nil {
		t.Fatal(err)
	}
	defer testserver.CleanupMasterEtcd(t, masterConfig)
	clusterAdminClientConfig, err := testutil.GetClusterAdminClientConfig(clusterAdminKubeConfig)
	if err != nil {
		t.Fatal(err)
	}
	_, _, err = testserver.CreateNewProject(clusterAdminClientConfig, namespace, "my-test-user")
	if err != nil {
		t.Fatal(err)
	}
	_, adminConfig, err := testutil.GetClientForUser(clusterAdminClientConfig, "my-test-user")
	if err != nil {
		t.Fatal(err)
	}
	adminAppsClient := appsclient.NewForConfigOrDie(adminConfig)
	config := appstest.OkDeploymentConfig(0)
	config.Namespace = namespace
	config.Spec.Triggers = []appsv1.DeploymentTriggerPolicy{}
	config.Spec.Replicas = 1
	dc, err := adminAppsClient.AppsV1().DeploymentConfigs(namespace).Create(config)
	if err != nil {
		t.Fatalf("Couldn't create DeploymentConfig: %v %#v", err, config)
	}
	generation := dc.Generation
	{
		scalePath := fmt.Sprintf("/apis/apps.openshift.io/v1/namespaces/%s/deploymentconfigs/%s/scale", dc.Namespace, dc.Name)
		scale := &unstructured.Unstructured{}
		if err := adminAppsClient.RESTClient().Get().AbsPath(scalePath).Do().Into(scale); err != nil {
			t.Fatal(err)
		}
		if scale.GetAPIVersion() != "extensions/v1beta1" {
			t.Fatalf("Expected extensions/v1beta1, got %v", scale.GetAPIVersion())
		}
		scaleBytes, err := scale.MarshalJSON()
		if err != nil {
			t.Fatal(err)
		}
		if err := adminAppsClient.RESTClient().Put().AbsPath(scalePath).Body(scaleBytes).Do().Error(); err != nil {
			t.Fatal(err)
		}
	}
	condition := func() (bool, error) {
		config, err := adminAppsClient.AppsV1().DeploymentConfigs(namespace).Get(dc.Name, metav1.GetOptions{})
		if err != nil {
			return false, nil
		}
		externalConfig := &appsv1.DeploymentConfig{}
		if err := legacyscheme.Scheme.Convert(config, externalConfig, nil); err != nil {
			panic(err)
		}
		return appsutil.HasSynced(externalConfig, generation), nil
	}
	if err := wait.PollImmediate(500*time.Millisecond, 10*time.Second, condition); err != nil {
		t.Fatalf("Deployment config never synced: %v", err)
	}
	restMapper := restmapper.NewDeferredDiscoveryRESTMapper(discocache.NewMemCacheClient(adminAppsClient.Discovery()))
	restMapper.Reset()
	scaleKindResolver := scale.NewDiscoveryScaleKindResolver(adminAppsClient.Discovery())
	scaleClient, err := scale.NewForConfig(adminConfig, restMapper, dynamic.LegacyAPIPathResolverFunc, scaleKindResolver)
	if err != nil {
		t.Fatal(err)
	}
	scale, err := scaleClient.Scales(namespace).Get(apps.Resource("deploymentconfigs"), config.Name)
	if err != nil {
		t.Fatalf("Couldn't get DeploymentConfig scale: %v", err)
	}
	if scale.Spec.Replicas != 1 {
		t.Fatalf("Expected scale.spec.replicas=1, got %#v", scale)
	}
	scaleUpdate := scale.DeepCopy()
	scaleUpdate.Spec.Replicas = 3
	updatedScale, err := scaleClient.Scales(namespace).Update(apps.Resource("deploymentconfigs"), scaleUpdate)
	if err != nil {
		t.Fatalf("Couldn't update DeploymentConfig scale to %#v: %v", scaleUpdate, err)
	}
	if updatedScale.Spec.Replicas != 3 {
		t.Fatalf("Expected scale.spec.replicas=3, got %#v", scale)
	}
	persistedScale, err := scaleClient.Scales(namespace).Get(apps.Resource("deploymentconfigs"), config.Name)
	if err != nil {
		t.Fatalf("Couldn't get DeploymentConfig scale: %v", err)
	}
	if persistedScale.Spec.Replicas != 3 {
		t.Fatalf("Expected scale.spec.replicas=3, got %#v", scale)
	}
}
