package integration

import (
	"fmt"
	"testing"
	"time"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	watchapi "k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/util/retry"
	appsv1 "github.com/openshift/api/apps/v1"
	appsclient "github.com/openshift/client-go/apps/clientset/versioned"
	appsutil "github.com/openshift/origin/pkg/apps/util"
	appstest "github.com/openshift/origin/pkg/apps/util/test"
	imageapi "github.com/openshift/origin/pkg/image/apis/image"
	imageclient "github.com/openshift/origin/pkg/image/generated/internalclientset"
	testutil "github.com/openshift/origin/test/util"
	testserver "github.com/openshift/origin/test/util/server"
)

const maxUpdateRetries = 10

func TestTriggers_manual(t *testing.T) {
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
	const namespace = "test-triggers-manual"
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
	kc, adminConfig, err := testutil.GetClientForUser(clusterAdminClientConfig, "my-test-user")
	if err != nil {
		t.Fatal(err)
	}
	adminAppsClient := appsclient.NewForConfigOrDie(adminConfig).AppsV1()
	config := appstest.OkDeploymentConfig(0)
	config.Namespace = namespace
	config.Spec.Triggers = []appsv1.DeploymentTriggerPolicy{{Type: "Manual"}}
	dc, err := adminAppsClient.DeploymentConfigs(namespace).Create(config)
	if err != nil {
		t.Fatalf("Couldn't create DeploymentConfig: %v %#v", err, config)
	}
	rcWatch, err := kc.CoreV1().ReplicationControllers(namespace).Watch(metav1.ListOptions{ResourceVersion: dc.ResourceVersion})
	if err != nil {
		t.Fatalf("Couldn't subscribe to Deployments: %v", err)
	}
	defer rcWatch.Stop()
	request := &appsv1.DeploymentRequest{Name: config.Name, Latest: false, Force: true}
	retryErr := retry.RetryOnConflict(wait.Backoff{Steps: maxUpdateRetries}, func() error {
		var err error
		config, err = adminAppsClient.DeploymentConfigs(namespace).Instantiate(config.Name, request)
		return err
	})
	if retryErr != nil {
		t.Fatalf("Couldn't instantiate deployment config %q: %v", request.Name, err)
	}
	if config.Status.LatestVersion != 1 {
		t.Fatal("Instantiated deployment config should have version 1")
	}
	if config.Status.Details == nil || len(config.Status.Details.Causes) == 0 {
		t.Fatal("Instantiated deployment config should have a cause of deployment")
	}
	gotType := config.Status.Details.Causes[0].Type
	if gotType != "Manual" {
		t.Fatalf("Instantiated deployment config should have a %q cause of deployment instead of %q", "Manual", gotType)
	}
	event := <-rcWatch.ResultChan()
	if e, a := watchapi.Added, event.Type; e != a {
		t.Fatalf("expected watch event type %s, got %s", e, a)
	}
	deployment := event.Object.(*corev1.ReplicationController)
	if e, a := config.Name, appsutil.DeploymentConfigNameFor(deployment); e != a {
		t.Fatalf("Expected deployment annotated with deploymentConfig '%s', got '%s'", e, a)
	}
	if e, a := int64(1), appsutil.DeploymentVersionFor(deployment); e != a {
		t.Fatalf("Deployment annotation version does not match: %#v", deployment)
	}
}
func TestTriggers_imageChange(t *testing.T) {
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
	const registryHostname = "registry:8080"
	testutil.SetAdditionalAllowedRegistries(registryHostname)
	masterConfig, clusterAdminKubeConfig, err := testserver.StartTestMaster()
	if err != nil {
		t.Fatalf("error starting master: %v", err)
	}
	defer testserver.CleanupMasterEtcd(t, masterConfig)
	openshiftClusterAdminClientConfig, err := testutil.GetClusterAdminClientConfig(clusterAdminKubeConfig)
	if err != nil {
		t.Fatalf("error getting cluster admin client config: %v", err)
	}
	_, projectAdminClientConfig, err := testserver.CreateNewProject(openshiftClusterAdminClientConfig, testutil.Namespace(), "bob")
	if err != nil {
		t.Fatalf("error creating project: %v", err)
	}
	projectAdminAppsClient := appsclient.NewForConfigOrDie(projectAdminClientConfig).AppsV1()
	projectAdminImageClient := imageclient.NewForConfigOrDie(projectAdminClientConfig).Image()
	imageStream := &imageapi.ImageStream{ObjectMeta: metav1.ObjectMeta{Name: appstest.ImageStreamName}}
	config := appstest.OkDeploymentConfig(0)
	config.Namespace = testutil.Namespace()
	config.Spec.Triggers = []appsv1.DeploymentTriggerPolicy{appstest.OkImageChangeTrigger()}
	configWatch, err := projectAdminAppsClient.DeploymentConfigs(testutil.Namespace()).Watch(metav1.ListOptions{})
	if err != nil {
		t.Fatalf("Couldn't subscribe to deploymentconfigs %v", err)
	}
	defer configWatch.Stop()
	if imageStream, err = projectAdminImageClient.ImageStreams(testutil.Namespace()).Create(imageStream); err != nil {
		t.Fatalf("Couldn't create imagestream: %v", err)
	}
	imageWatch, err := projectAdminImageClient.ImageStreams(testutil.Namespace()).Watch(metav1.ListOptions{})
	if err != nil {
		t.Fatalf("Couldn't subscribe to imagestreams: %v", err)
	}
	defer imageWatch.Stop()
	updatedImage := fmt.Sprintf("sha256:%s", appstest.ImageID)
	updatedPullSpec := fmt.Sprintf("%s/%s/%s@%s", registryHostname, testutil.Namespace(), appstest.ImageStreamName, updatedImage)
	createTagEvent := func() {
		mapping := &imageapi.ImageStreamMapping{ObjectMeta: metav1.ObjectMeta{Name: imageStream.Name}, Tag: imageapi.DefaultImageTag, Image: imageapi.Image{ObjectMeta: metav1.ObjectMeta{Name: updatedImage}, DockerImageReference: updatedPullSpec}}
		if _, err := projectAdminImageClient.ImageStreamMappings(testutil.Namespace()).Create(mapping); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		t.Log("Waiting for image stream mapping to be reflected in the image stream status...")
	statusLoop:
		for {
			select {
			case event := <-imageWatch.ResultChan():
				stream := event.Object.(*imageapi.ImageStream)
				if _, ok := stream.Status.Tags[imageapi.DefaultImageTag]; ok {
					t.Logf("imagestream %q now has status with tags: %#v", stream.Name, stream.Status.Tags)
					break statusLoop
				}
				t.Logf("Still waiting for latest tag status on imagestream %q", stream.Name)
			}
		}
	}
	if config, err = projectAdminAppsClient.DeploymentConfigs(testutil.Namespace()).Create(config); err != nil {
		t.Fatalf("Couldn't create deploymentconfig: %v", err)
	}
	createTagEvent()
	var newConfig *appsv1.DeploymentConfig
	t.Log("Waiting for a new deployment config in response to imagestream update")
waitForNewConfig:
	for {
		select {
		case event := <-configWatch.ResultChan():
			if event.Type == watchapi.Modified {
				newConfig = event.Object.(*appsv1.DeploymentConfig)
				if newConfig.Status.LatestVersion > 0 {
					if e, a := updatedPullSpec, newConfig.Spec.Template.Spec.Containers[0].Image; e != a {
						t.Fatalf("unexpected image for pod template container 0; expected %q, got %q", e, a)
					}
					break waitForNewConfig
				}
				t.Log("Still waiting for a new deployment config in response to imagestream update")
			}
		}
	}
}
func TestTriggers_imageChange_nonAutomatic(t *testing.T) {
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
	const registryHostname = "registry:8080"
	testutil.SetAdditionalAllowedRegistries(registryHostname, "registry:5000")
	masterConfig, clusterAdminKubeConfig, err := testserver.StartTestMaster()
	if err != nil {
		t.Fatalf("error starting master: %v", err)
	}
	defer testserver.CleanupMasterEtcd(t, masterConfig)
	openshiftClusterAdminClientConfig, err := testutil.GetClusterAdminClientConfig(clusterAdminKubeConfig)
	if err != nil {
		t.Fatalf("error getting cluster admin client config: %v", err)
	}
	_, adminConfig, err := testserver.CreateNewProject(openshiftClusterAdminClientConfig, testutil.Namespace(), "bob")
	if err != nil {
		t.Fatalf("error creating project: %v", err)
	}
	adminAppsClient := appsclient.NewForConfigOrDie(adminConfig).AppsV1()
	adminImageClient := imageclient.NewForConfigOrDie(adminConfig).Image()
	imageStream := &imageapi.ImageStream{ObjectMeta: metav1.ObjectMeta{Name: appstest.ImageStreamName}}
	if imageStream, err = adminImageClient.ImageStreams(testutil.Namespace()).Create(imageStream); err != nil {
		t.Fatalf("Couldn't create imagestream: %v", err)
	}
	imageWatch, err := adminImageClient.ImageStreams(testutil.Namespace()).Watch(metav1.ListOptions{})
	if err != nil {
		t.Fatalf("Couldn't subscribe to imagestreams: %v", err)
	}
	defer imageWatch.Stop()
	image := fmt.Sprintf("sha256:%s", appstest.ImageID)
	pullSpec := fmt.Sprintf("registry:5000/%s/%s@%s", testutil.Namespace(), appstest.ImageStreamName, image)
	mapping := &imageapi.ImageStreamMapping{ObjectMeta: metav1.ObjectMeta{Name: imageStream.Name}, Tag: imageapi.DefaultImageTag, Image: imageapi.Image{ObjectMeta: metav1.ObjectMeta{Name: image}, DockerImageReference: pullSpec}}
	createTagEvent := func(mapping *imageapi.ImageStreamMapping) {
		if _, err := adminImageClient.ImageStreamMappings(testutil.Namespace()).Create(mapping); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		t.Log("Waiting for image stream mapping to be reflected in the image stream status...")
		timeout := time.After(time.Minute)
		for {
			select {
			case event := <-imageWatch.ResultChan():
				stream := event.Object.(*imageapi.ImageStream)
				tagEventList, ok := stream.Status.Tags[imageapi.DefaultImageTag]
				if ok && len(tagEventList.Items) > 0 && tagEventList.Items[0].DockerImageReference == mapping.Image.DockerImageReference {
					t.Logf("imagestream %q now has status with tags: %#v", stream.Name, stream.Status.Tags)
					return
				}
				if len(tagEventList.Items) > 0 {
					t.Logf("want: %s, got: %s", mapping.Image.DockerImageReference, tagEventList.Items[0].DockerImageReference)
				}
				t.Logf("Still waiting for latest tag status update on imagestream %q with tags: %#v", stream.Name, tagEventList)
			case <-timeout:
				t.Fatalf("timed out waiting for image stream %q to be updated", imageStream.Name)
			}
		}
	}
	configWatch, err := adminAppsClient.DeploymentConfigs(testutil.Namespace()).Watch(metav1.ListOptions{})
	if err != nil {
		t.Fatalf("Couldn't subscribe to deploymentconfigs: %v", err)
	}
	defer configWatch.Stop()
	config := appstest.OkDeploymentConfig(0)
	config.Namespace = testutil.Namespace()
	config.Spec.Triggers = []appsv1.DeploymentTriggerPolicy{appstest.OkImageChangeTrigger()}
	config.Spec.Triggers[0].ImageChangeParams.Automatic = false
	if config, err = adminAppsClient.DeploymentConfigs(testutil.Namespace()).Create(config); err != nil {
		t.Fatalf("Couldn't create deploymentconfig: %v", err)
	}
	createTagEvent(mapping)
	var newConfig *appsv1.DeploymentConfig
	t.Log("Waiting for the first imagestream update - no deployment should run")
	timeout := time.After(20 * time.Second)
out:
	for {
		select {
		case event := <-configWatch.ResultChan():
			if event.Type != watchapi.Modified {
				continue
			}
			newConfig = event.Object.(*appsv1.DeploymentConfig)
			if newConfig.Status.LatestVersion > 0 {
				t.Fatalf("unexpected latestVersion update - the config has no config change trigger")
			}
		case <-timeout:
			break out
		}
	}
	t.Log("Waiting for the second imagestream update - no deployment should run")
	mapping.Image.Name = "sha256:0000000000000000000000000000000000000000000000000000000000000321"
	mapping.Image.DockerImageReference = fmt.Sprintf("%s/%s/%s@%s", registryHostname, testutil.Namespace(), appstest.ImageStreamName, mapping.Image.Name)
	createTagEvent(mapping)
	timeout = time.After(20 * time.Second)
loop:
	for {
		select {
		case event := <-configWatch.ResultChan():
			if event.Type != watchapi.Modified {
				continue
			}
			newConfig = event.Object.(*appsv1.DeploymentConfig)
			if newConfig.Status.LatestVersion > 0 {
				t.Fatalf("unexpected latestVersion update - the config has no config change trigger")
			}
		case <-timeout:
			break loop
		}
	}
	t.Log("Instantiate the deployment config - the latest image should be picked up and a new deployment should run")
	request := &appsv1.DeploymentRequest{Name: config.Name, Latest: true, Force: true}
	retryErr := retry.RetryOnConflict(wait.Backoff{Steps: maxUpdateRetries}, func() error {
		var err error
		config, err = adminAppsClient.DeploymentConfigs(testutil.Namespace()).Instantiate(config.Name, request)
		return err
	})
	if retryErr != nil {
		t.Fatalf("Couldn't instantiate deployment config %q: %v", request.Name, err)
	}
	config, err = adminAppsClient.DeploymentConfigs(config.Namespace).Get(config.Name, metav1.GetOptions{})
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if exp, got := mapping.Image.DockerImageReference, config.Spec.Template.Spec.Containers[0].Image; exp != got {
		t.Fatalf("Expected image %q instead of %q to be updated in deployment config %q", exp, got, config.Name)
	}
	if exp, got := int64(1), config.Status.LatestVersion; exp != got {
		t.Fatalf("Expected latestVersion for deployment config %q to be %d, got %d", config.Name, exp, got)
	}
	if config.Status.Details == nil || len(config.Status.Details.Causes) == 0 {
		t.Fatalf("Expected a cause of deployment for deployment config %q", config.Name)
	}
	if gotType, expectedType := config.Status.Details.Causes[0].Type, "Manual"; string(gotType) != expectedType {
		t.Fatalf("Instantiated deployment config should have a %q cause of deployment instead of %q", expectedType, gotType)
	}
}
func TestTriggers_MultipleICTs(t *testing.T) {
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
	const registryHostname = "registry:8080"
	testutil.SetAdditionalAllowedRegistries(registryHostname)
	masterConfig, clusterAdminKubeConfig, err := testserver.StartTestMaster()
	if err != nil {
		t.Fatalf("error starting master: %v", err)
	}
	defer testserver.CleanupMasterEtcd(t, masterConfig)
	openshiftClusterAdminClientConfig, err := testutil.GetClusterAdminClientConfig(clusterAdminKubeConfig)
	if err != nil {
		t.Fatalf("error getting cluster admin client config: %v", err)
	}
	_, adminConfig, err := testserver.CreateNewProject(openshiftClusterAdminClientConfig, testutil.Namespace(), "bob")
	if err != nil {
		t.Fatalf("error creating project: %v", err)
	}
	adminAppsClient := appsclient.NewForConfigOrDie(adminConfig).AppsV1()
	adminImageClient := imageclient.NewForConfigOrDie(adminConfig).Image()
	imageStream := &imageapi.ImageStream{ObjectMeta: metav1.ObjectMeta{Name: appstest.ImageStreamName}}
	secondImageStream := &imageapi.ImageStream{ObjectMeta: metav1.ObjectMeta{Name: "sample"}}
	config := appstest.OkDeploymentConfig(0)
	config.Namespace = testutil.Namespace()
	firstTrigger := appstest.OkImageChangeTrigger()
	secondTrigger := appstest.OkImageChangeTrigger()
	secondTrigger.ImageChangeParams.ContainerNames = []string{"container2"}
	secondTrigger.ImageChangeParams.From.Name = imageapi.JoinImageStreamTag("sample", imageapi.DefaultImageTag)
	config.Spec.Triggers = []appsv1.DeploymentTriggerPolicy{firstTrigger, secondTrigger}
	configWatch, err := adminAppsClient.DeploymentConfigs(testutil.Namespace()).Watch(metav1.ListOptions{})
	if err != nil {
		t.Fatalf("Couldn't subscribe to deploymentconfigs %v", err)
	}
	defer configWatch.Stop()
	if imageStream, err = adminImageClient.ImageStreams(testutil.Namespace()).Create(imageStream); err != nil {
		t.Fatalf("Couldn't create imagestream %q: %v", imageStream.Name, err)
	}
	if secondImageStream, err = adminImageClient.ImageStreams(testutil.Namespace()).Create(secondImageStream); err != nil {
		t.Fatalf("Couldn't create imagestream %q: %v", secondImageStream.Name, err)
	}
	imageWatch, err := adminImageClient.ImageStreams(testutil.Namespace()).Watch(metav1.ListOptions{})
	if err != nil {
		t.Fatalf("Couldn't subscribe to imagestreams: %v", err)
	}
	defer imageWatch.Stop()
	updatedImage := fmt.Sprintf("sha256:%s", appstest.ImageID)
	updatedPullSpec := fmt.Sprintf("%s/%s/%s@%s", registryHostname, testutil.Namespace(), appstest.ImageStreamName, updatedImage)
	createTagEvent := func(name, tag, image, pullSpec string) {
		mapping := &imageapi.ImageStreamMapping{ObjectMeta: metav1.ObjectMeta{Name: name}, Tag: tag, Image: imageapi.Image{ObjectMeta: metav1.ObjectMeta{Name: image}, DockerImageReference: pullSpec}}
		if _, err := adminImageClient.ImageStreamMappings(testutil.Namespace()).Create(mapping); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		t.Log("Waiting for image stream mapping to be reflected in the image stream status...")
	statusLoop:
		for {
			select {
			case event := <-imageWatch.ResultChan():
				stream := event.Object.(*imageapi.ImageStream)
				if stream.Name != name {
					continue
				}
				if _, ok := stream.Status.Tags[tag]; ok {
					t.Logf("imagestream %q now has status with tags: %#v", stream.Name, stream.Status.Tags)
					break statusLoop
				}
				t.Logf("Still waiting for latest tag status on imagestream %q", stream.Name)
			}
		}
	}
	if config, err = adminAppsClient.DeploymentConfigs(testutil.Namespace()).Create(config); err != nil {
		t.Fatalf("Couldn't create deploymentconfig: %v", err)
	}
	timeout := time.After(30 * time.Second)
	t.Log("Should not trigger a new deployment in response to the first imagestream update")
	createTagEvent(imageStream.Name, imageapi.DefaultImageTag, updatedImage, updatedPullSpec)
out:
	for {
		select {
		case event := <-configWatch.ResultChan():
			if event.Type != watchapi.Modified {
				continue
			}
			newConfig := event.Object.(*appsv1.DeploymentConfig)
			if newConfig.Status.LatestVersion > 0 {
				t.Fatalf("unexpected latestVersion update: %#v", newConfig)
			}
			container := newConfig.Spec.Template.Spec.Containers[0]
			if e, a := updatedPullSpec, container.Image; e == a {
				t.Fatalf("unexpected image update: %#v", newConfig)
			}
		case <-timeout:
			break out
		}
	}
	timeout = time.After(30 * time.Second)
	t.Log("Should trigger a new deployment in response to the second imagestream update")
	secondImage := "sampleimage"
	secondPullSpec := "samplepullspec"
	createTagEvent(secondImageStream.Name, imageapi.DefaultImageTag, secondImage, secondPullSpec)
	for {
	inner:
		select {
		case event := <-configWatch.ResultChan():
			if event.Type != watchapi.Modified {
				continue
			}
			newConfig := event.Object.(*appsv1.DeploymentConfig)
			switch {
			case newConfig.Status.LatestVersion == 0:
				t.Logf("Wating for latestVersion to update to 1")
				break inner
			case newConfig.Status.LatestVersion > 1:
				t.Fatalf("unexpected latestVersion %d for %#v", newConfig.Status.LatestVersion, newConfig)
			default:
			}
			container := newConfig.Spec.Template.Spec.Containers[0]
			if e, a := updatedPullSpec, container.Image; e != a {
				t.Fatalf("unexpected image for pod template container %q; expected %q, got %q", container.Name, e, a)
			}
			container = newConfig.Spec.Template.Spec.Containers[1]
			if e, a := secondPullSpec, container.Image; e != a {
				t.Fatalf("unexpected image for pod template container %q; expected %q, got %q", container.Name, e, a)
			}
			return
		case <-timeout:
			t.Fatalf("timed out waiting for the second image update to happen")
		}
	}
}
func TestTriggers_configChange(t *testing.T) {
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
	const namespace = "test-triggers-configchange"
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
	kc, adminConfig, err := testutil.GetClientForUser(clusterAdminClientConfig, "my-test-user")
	if err != nil {
		t.Fatal(err)
	}
	adminAppsClient := appsclient.NewForConfigOrDie(adminConfig).AppsV1()
	config := appstest.OkDeploymentConfig(0)
	config.Namespace = namespace
	config.Spec.Triggers = []appsv1.DeploymentTriggerPolicy{appstest.OkConfigChangeTrigger()}
	rcWatch, err := kc.CoreV1().ReplicationControllers(namespace).Watch(metav1.ListOptions{})
	if err != nil {
		t.Fatalf("Couldn't subscribe to Deployments %v", err)
	}
	defer rcWatch.Stop()
	config, err = adminAppsClient.DeploymentConfigs(namespace).Create(config)
	if err != nil {
		t.Fatalf("Couldn't create DeploymentConfig: %v", err)
	}
	event := <-rcWatch.ResultChan()
	if e, a := watchapi.Added, event.Type; e != a {
		t.Fatalf("expected watch event type %s, got %s", e, a)
	}
	deployment := event.Object.(*corev1.ReplicationController)
	if e, a := config.Name, appsutil.DeploymentConfigNameFor(deployment); e != a {
		t.Fatalf("Expected deployment annotated with deploymentConfig '%s', got '%s'", e, a)
	}
	retryErr := retry.RetryOnConflict(wait.Backoff{Steps: maxUpdateRetries}, func() error {
		liveDeployment, err := kc.CoreV1().ReplicationControllers(deployment.Namespace).Get(deployment.Name, metav1.GetOptions{})
		if err != nil {
			return err
		}
		liveDeployment.Annotations[appsv1.DeploymentStatusAnnotation] = string(appsv1.DeploymentStatusComplete)
		_, err = kc.CoreV1().ReplicationControllers(namespace).Update(liveDeployment)
		return err
	})
	if retryErr != nil {
		t.Fatal(retryErr)
	}
	event = <-rcWatch.ResultChan()
	if e, a := watchapi.Modified, event.Type; e != a {
		t.Fatalf("expected watch event type %s, got %s", e, a)
	}
	assertEnvVarEquals("ENV1", "VAL1", deployment, t)
	retryErr = retry.RetryOnConflict(wait.Backoff{Steps: maxUpdateRetries}, func() error {
		latest, err := adminAppsClient.DeploymentConfigs(namespace).Get(config.Name, metav1.GetOptions{})
		if err != nil {
			return err
		}
		for i, e := range latest.Spec.Template.Spec.Containers[0].Env {
			if e.Name == "ENV1" {
				latest.Spec.Template.Spec.Containers[0].Env[i].Value = "UPDATED"
				break
			}
		}
		_, err = adminAppsClient.DeploymentConfigs(namespace).Update(latest)
		return err
	})
	if retryErr != nil {
		t.Fatal(retryErr)
	}
	if retryErr := retry.RetryOnConflict(wait.Backoff{Steps: maxUpdateRetries}, func() error {
		newConfig, err := adminAppsClient.DeploymentConfigs(namespace).Get(config.Name, metav1.GetOptions{})
		if err != nil {
			return err
		}
		newConfig.Spec.Template.Spec.Containers[0].Env[0].Value = "UPDATED"
		_, err = adminAppsClient.DeploymentConfigs(namespace).Update(newConfig)
		return err
	}); retryErr != nil {
		t.Fatal(retryErr)
	}
	var newDeployment *corev1.ReplicationController
	for {
		event = <-rcWatch.ResultChan()
		if event.Type != watchapi.Added {
			continue
		}
		newDeployment = event.Object.(*corev1.ReplicationController)
		break
	}
	assertEnvVarEquals("ENV1", "UPDATED", newDeployment, t)
	if newDeployment.Name == deployment.Name {
		t.Fatalf("expected new deployment; old=%s, new=%s", deployment.Name, newDeployment.Name)
	}
}
func assertEnvVarEquals(name string, value string, deployment *corev1.ReplicationController, t *testing.T) {
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
	env := deployment.Spec.Template.Spec.Containers[0].Env
	for _, e := range env {
		if e.Name == name && e.Value == value {
			return
		}
	}
	t.Fatalf("Expected env var with name %s and value %s", name, value)
}
