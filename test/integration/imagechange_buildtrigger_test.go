package integration

import (
	"testing"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	watchapi "k8s.io/apimachinery/pkg/watch"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/cli-runtime/pkg/genericclioptions/printers"
	"k8s.io/client-go/kubernetes"
	rbacv1client "k8s.io/client-go/kubernetes/typed/rbac/v1"
	"k8s.io/client-go/rest"
	kapi "k8s.io/kubernetes/pkg/apis/core"
	"github.com/openshift/api/build"
	buildv1 "github.com/openshift/api/build/v1"
	buildv1client "github.com/openshift/client-go/build/clientset/versioned"
	"github.com/openshift/origin/pkg/cmd/server/bootstrappolicy"
	imageapi "github.com/openshift/origin/pkg/image/apis/image"
	imageclient "github.com/openshift/origin/pkg/image/generated/internalclientset"
	"github.com/openshift/origin/pkg/oc/cli/admin/policy"
	testutil "github.com/openshift/origin/test/util"
	testserver "github.com/openshift/origin/test/util/server"
)

const (
	streamName		= "test-image-trigger-repo"
	tag			= "latest"
	registryHostname	= "registry:8000"
)

func TestSimpleImageChangeBuildTriggerFromImageStreamTagSTI(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	testutil.SetAdditionalAllowedRegistries(registryHostname)
	_, _, projectAdminConfig, fn := setup(t)
	defer fn()
	imageStream := mockImageStream2(tag)
	imageStreamMapping := mockImageStreamMapping(imageStream.Name, "someimage", tag, registryHostname+"/openshift/test-image-trigger:"+tag)
	strategy := stiStrategy("ImageStreamTag", streamName+":"+tag)
	config := imageChangeBuildConfig("sti-imagestreamtag", strategy)
	runTest(t, "SimpleImageChangeBuildTriggerFromImageStreamTagSTI", projectAdminConfig, imageStream, imageStreamMapping, config, tag)
}
func TestSimpleImageChangeBuildTriggerFromImageStreamTagSTIWithConfigChange(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	testutil.SetAdditionalAllowedRegistries(registryHostname)
	_, _, projectAdminConfig, fn := setup(t)
	defer fn()
	imageStream := mockImageStream2(tag)
	imageStreamMapping := mockImageStreamMapping(imageStream.Name, "someimage", tag, registryHostname+"/openshift/test-image-trigger:"+tag)
	strategy := stiStrategy("ImageStreamTag", streamName+":"+tag)
	config := imageChangeBuildConfigWithConfigChange("sti-imagestreamtag", strategy)
	runTest(t, "SimpleImageChangeBuildTriggerFromImageStreamTagSTI", projectAdminConfig, imageStream, imageStreamMapping, config, tag)
}
func TestSimpleImageChangeBuildTriggerFromImageStreamTagDocker(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	testutil.SetAdditionalAllowedRegistries(registryHostname)
	_, _, projectAdminConfig, fn := setup(t)
	defer fn()
	imageStream := mockImageStream2(tag)
	imageStreamMapping := mockImageStreamMapping(imageStream.Name, "someimage", tag, registryHostname+"/openshift/test-image-trigger:"+tag)
	strategy := dockerStrategy("ImageStreamTag", streamName+":"+tag)
	config := imageChangeBuildConfig("docker-imagestreamtag", strategy)
	runTest(t, "SimpleImageChangeBuildTriggerFromImageStreamTagDocker", projectAdminConfig, imageStream, imageStreamMapping, config, tag)
}
func TestSimpleImageChangeBuildTriggerFromImageStreamTagDockerWithConfigChange(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	testutil.SetAdditionalAllowedRegistries(registryHostname)
	_, _, projectAdminConfig, fn := setup(t)
	defer fn()
	imageStream := mockImageStream2(tag)
	imageStreamMapping := mockImageStreamMapping(imageStream.Name, "someimage", tag, registryHostname+"/openshift/test-image-trigger:"+tag)
	strategy := dockerStrategy("ImageStreamTag", streamName+":"+tag)
	config := imageChangeBuildConfigWithConfigChange("docker-imagestreamtag", strategy)
	runTest(t, "SimpleImageChangeBuildTriggerFromImageStreamTagDocker", projectAdminConfig, imageStream, imageStreamMapping, config, tag)
}
func TestSimpleImageChangeBuildTriggerFromImageStreamTagCustom(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	testutil.SetAdditionalAllowedRegistries(registryHostname)
	clusterAdminClientConfig, projectAdminKubeClient, projectAdminConfig, fn := setup(t)
	defer fn()
	subjects := []rbacv1.Subject{{APIGroup: rbacv1.GroupName, Kind: rbacv1.GroupKind, Name: bootstrappolicy.AuthenticatedGroup}}
	options := policy.RoleModificationOptions{RoleName: bootstrappolicy.BuildStrategyCustomRoleName, RoleKind: "ClusterRole", RbacClient: rbacv1client.NewForConfigOrDie(clusterAdminClientConfig), Subjects: subjects, PrintFlags: genericclioptions.NewPrintFlags(""), ToPrinter: func(string) (printers.ResourcePrinter, error) {
		return printers.NewDiscardingPrinter(), nil
	}}
	options.AddRole()
	if err := testutil.WaitForPolicyUpdate(projectAdminKubeClient.AuthorizationV1(), testutil.Namespace(), "create", build.Resource(bootstrappolicy.CustomBuildResource), true); err != nil {
		t.Fatal(err)
	}
	imageStream := mockImageStream2(tag)
	imageStreamMapping := mockImageStreamMapping(imageStream.Name, "someimage", tag, registryHostname+"/openshift/test-image-trigger:"+tag)
	strategy := customStrategy("ImageStreamTag", streamName+":"+tag)
	config := imageChangeBuildConfig("custom-imagestreamtag", strategy)
	runTest(t, "SimpleImageChangeBuildTriggerFromImageStreamTagCustom", projectAdminConfig, imageStream, imageStreamMapping, config, tag)
}
func TestSimpleImageChangeBuildTriggerFromImageStreamTagCustomWithConfigChange(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	testutil.SetAdditionalAllowedRegistries(registryHostname)
	clusterAdminClientConfig, projectAdminKubeClient, projectAdminConfig, fn := setup(t)
	defer fn()
	subjects := []rbacv1.Subject{{APIGroup: rbacv1.GroupName, Kind: rbacv1.GroupKind, Name: bootstrappolicy.AuthenticatedGroup}}
	options := policy.RoleModificationOptions{RoleName: bootstrappolicy.BuildStrategyCustomRoleName, RoleKind: "ClusterRole", RbacClient: rbacv1client.NewForConfigOrDie(clusterAdminClientConfig), Subjects: subjects, PrintFlags: genericclioptions.NewPrintFlags(""), ToPrinter: func(string) (printers.ResourcePrinter, error) {
		return printers.NewDiscardingPrinter(), nil
	}}
	options.AddRole()
	if err := testutil.WaitForPolicyUpdate(projectAdminKubeClient.AuthorizationV1(), testutil.Namespace(), "create", build.Resource(bootstrappolicy.CustomBuildResource), true); err != nil {
		t.Fatal(err)
	}
	imageStream := mockImageStream2(tag)
	imageStreamMapping := mockImageStreamMapping(imageStream.Name, "someimage", tag, registryHostname+"/openshift/test-image-trigger:"+tag)
	strategy := customStrategy("ImageStreamTag", streamName+":"+tag)
	config := imageChangeBuildConfigWithConfigChange("custom-imagestreamtag", strategy)
	runTest(t, "SimpleImageChangeBuildTriggerFromImageStreamTagCustom", projectAdminConfig, imageStream, imageStreamMapping, config, tag)
}
func dockerStrategy(kind, name string) buildv1.BuildStrategy {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return buildv1.BuildStrategy{DockerStrategy: &buildv1.DockerBuildStrategy{From: &corev1.ObjectReference{Kind: kind, Name: name}}}
}
func stiStrategy(kind, name string) buildv1.BuildStrategy {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return buildv1.BuildStrategy{SourceStrategy: &buildv1.SourceBuildStrategy{From: corev1.ObjectReference{Kind: kind, Name: name}}}
}
func customStrategy(kind, name string) buildv1.BuildStrategy {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return buildv1.BuildStrategy{CustomStrategy: &buildv1.CustomBuildStrategy{From: corev1.ObjectReference{Kind: kind, Name: name}}}
}
func imageChangeBuildConfig(name string, strategy buildv1.BuildStrategy) *buildv1.BuildConfig {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &buildv1.BuildConfig{ObjectMeta: metav1.ObjectMeta{Name: name, Namespace: testutil.Namespace(), Labels: map[string]string{"testlabel": "testvalue"}}, Spec: buildv1.BuildConfigSpec{RunPolicy: buildv1.BuildRunPolicyParallel, CommonSpec: buildv1.CommonSpec{Source: buildv1.BuildSource{Git: &buildv1.GitBuildSource{URI: "git://github.com/openshift/ruby-hello-world.git"}, ContextDir: "contextimage"}, Strategy: strategy, Output: buildv1.BuildOutput{To: &corev1.ObjectReference{Kind: "ImageStreamTag", Name: "test-image-trigger-repo:outputtag"}}}, Triggers: []buildv1.BuildTriggerPolicy{{Type: buildv1.ImageChangeBuildTriggerType, ImageChange: &buildv1.ImageChangeTrigger{}}}}}
}
func imageChangeBuildConfigWithConfigChange(name string, strategy buildv1.BuildStrategy) *buildv1.BuildConfig {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bc := imageChangeBuildConfig(name, strategy)
	bc.Spec.Triggers = append(bc.Spec.Triggers, buildv1.BuildTriggerPolicy{Type: buildv1.ConfigChangeBuildTriggerType})
	return bc
}
func mockImageStream2(tag string) *imageapi.ImageStream {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &imageapi.ImageStream{ObjectMeta: metav1.ObjectMeta{Name: "test-image-trigger-repo"}, Spec: imageapi.ImageStreamSpec{DockerImageRepository: registryHostname + "/openshift/test-image-trigger", Tags: map[string]imageapi.TagReference{tag: {From: &kapi.ObjectReference{Kind: "DockerImage", Name: registryHostname + "/openshift/test-image-trigger:" + tag}}}}}
}
func mockImageStreamMapping(stream, image, tag, reference string) *imageapi.ImageStreamMapping {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &imageapi.ImageStreamMapping{ObjectMeta: metav1.ObjectMeta{Name: stream}, Tag: tag, Image: imageapi.Image{ObjectMeta: metav1.ObjectMeta{Name: image}, DockerImageReference: reference}}
}
func setup(t *testing.T) (*rest.Config, kubernetes.Interface, *rest.Config, func()) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	masterConfig, clusterAdminKubeConfigFile, err := testserver.StartTestMaster()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	clusterAdminKubeConfig, err := testutil.GetClusterAdminClientConfig(clusterAdminKubeConfigFile)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	projectKubeAdminClient, projectAdminConfig, err := testserver.CreateNewProject(clusterAdminKubeConfig, testutil.Namespace(), testutil.Namespace())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	return clusterAdminKubeConfig, projectKubeAdminClient, projectAdminConfig, func() {
		testserver.CleanupMasterEtcd(t, masterConfig)
	}
}
func runTest(t *testing.T, testname string, projectAdminClientConfig *rest.Config, imageStream *imageapi.ImageStream, imageStreamMapping *imageapi.ImageStreamMapping, config *buildv1.BuildConfig, tag string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	projectAdminBuildClient := buildv1client.NewForConfigOrDie(projectAdminClientConfig).BuildV1()
	projectAdminImageClient := imageclient.NewForConfigOrDie(projectAdminClientConfig).Image()
	created, err := projectAdminBuildClient.BuildConfigs(testutil.Namespace()).Create(config)
	if err != nil {
		t.Fatalf("Couldn't create BuildConfig: %v", err)
	}
	buildWatch, err := projectAdminBuildClient.Builds(testutil.Namespace()).Watch(metav1.ListOptions{ResourceVersion: created.ResourceVersion})
	if err != nil {
		t.Fatalf("Couldn't subscribe to Builds %v", err)
	}
	defer buildWatch.Stop()
	buildConfigWatch, err := projectAdminBuildClient.BuildConfigs(testutil.Namespace()).Watch(metav1.ListOptions{ResourceVersion: created.ResourceVersion})
	if err != nil {
		t.Fatalf("Couldn't subscribe to BuildConfigs %v", err)
	}
	defer buildConfigWatch.Stop()
	imageStream, err = projectAdminImageClient.ImageStreams(testutil.Namespace()).Create(imageStream)
	if err != nil {
		t.Fatalf("Couldn't create ImageStream: %v", err)
	}
	_, err = projectAdminImageClient.ImageStreamMappings(testutil.Namespace()).Create(imageStreamMapping)
	if err != nil {
		t.Fatalf("Couldn't create Image: %v", err)
	}
	event := <-buildWatch.ResultChan()
	if e, a := watchapi.Added, event.Type; e != a {
		t.Fatalf("expected watch event type %s, got %s", e, a)
	}
	newBuild := event.Object.(*buildv1.Build)
	build1Name := newBuild.Name
	strategy := newBuild.Spec.Strategy
	switch {
	case strategy.SourceStrategy != nil:
		if strategy.SourceStrategy.From.Name != registryHostname+"/openshift/test-image-trigger:"+tag {
			i, _ := projectAdminImageClient.ImageStreams(testutil.Namespace()).Get(imageStream.Name, metav1.GetOptions{})
			bc, _ := projectAdminBuildClient.BuildConfigs(testutil.Namespace()).Get(config.Name, metav1.GetOptions{})
			t.Fatalf("Expected build with base image %s, got %s\n, imagerepo is %v\ntrigger is %#v\n", registryHostname+"/openshift/test-image-trigger:"+tag, strategy.SourceStrategy.From.Name, i, bc.Spec.Triggers[0].ImageChange)
		}
	case strategy.DockerStrategy != nil:
		if strategy.DockerStrategy.From.Name != registryHostname+"/openshift/test-image-trigger:"+tag {
			i, _ := projectAdminImageClient.ImageStreams(testutil.Namespace()).Get(imageStream.Name, metav1.GetOptions{})
			bc, _ := projectAdminBuildClient.BuildConfigs(testutil.Namespace()).Get(config.Name, metav1.GetOptions{})
			t.Fatalf("Expected build with base image %s, got %s\n, imagerepo is %v\ntrigger is %#v\n", registryHostname+"/openshift/test-image-trigger:"+tag, strategy.DockerStrategy.From.Name, i, bc.Spec.Triggers[0].ImageChange)
		}
	case strategy.CustomStrategy != nil:
		if strategy.CustomStrategy.From.Name != registryHostname+"/openshift/test-image-trigger:"+tag {
			i, _ := projectAdminImageClient.ImageStreams(testutil.Namespace()).Get(imageStream.Name, metav1.GetOptions{})
			bc, _ := projectAdminBuildClient.BuildConfigs(testutil.Namespace()).Get(config.Name, metav1.GetOptions{})
			t.Fatalf("Expected build with base image %s, got %s\n, imagerepo is %v\ntrigger is %#v\n", registryHostname+"/openshift/test-image-trigger:"+tag, strategy.CustomStrategy.From.Name, i, bc.Spec.Triggers[0].ImageChange)
		}
	}
	event = <-buildWatch.ResultChan()
	if e, a := watchapi.Modified, event.Type; e != a {
		t.Fatalf("expected watch event type %s, got %s: %#v", e, a, event.Object)
	}
	newBuild = event.Object.(*buildv1.Build)
	if newBuild.Spec.Output.To.Name != "test-image-trigger-repo:outputtag" {
		t.Fatalf("unexpected build output: %#v %#v", newBuild.Spec.Output.To, newBuild.Spec.Output)
	}
	if newBuild.Labels["testlabel"] != "testvalue" {
		t.Fatalf("Expected build with label %s=%s from build config got %s=%s", "testlabel", "testvalue", "testlabel", newBuild.Labels["testlabel"])
	}
	<-buildConfigWatch.ResultChan()
	updatedConfig, err := projectAdminBuildClient.BuildConfigs(testutil.Namespace()).Get(config.Name, metav1.GetOptions{})
	if err != nil {
		t.Fatalf("Couldn't get BuildConfig: %v", err)
	}
	if updatedConfig.Spec.Triggers[0].ImageChange.LastTriggeredImageID != registryHostname+"/openshift/test-image-trigger:"+tag {
		t.Errorf("Expected imageID equal to pull spec, got %#v", updatedConfig.Spec.Triggers[0].ImageChange)
	}
	if _, err := projectAdminImageClient.ImageStreamMappings(testutil.Namespace()).Create(&imageapi.ImageStreamMapping{ObjectMeta: metav1.ObjectMeta{Namespace: testutil.Namespace(), Name: imageStream.Name}, Tag: tag, Image: imageapi.Image{ObjectMeta: metav1.ObjectMeta{Name: "ref-2-random"}, DockerImageReference: registryHostname + "/openshift/test-image-trigger:ref-2-random"}}); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	for {
		event = <-buildWatch.ResultChan()
		newBuild = event.Object.(*buildv1.Build)
		if newBuild.Name != build1Name {
			break
		}
	}
	if e, a := watchapi.Added, event.Type; e != a {
		t.Fatalf("expected watch event type %s, got %s", e, a)
	}
	strategy = newBuild.Spec.Strategy
	switch {
	case strategy.SourceStrategy != nil:
		if strategy.SourceStrategy.From.Name != registryHostname+"/openshift/test-image-trigger:ref-2-random" {
			i, _ := projectAdminImageClient.ImageStreams(testutil.Namespace()).Get(imageStream.Name, metav1.GetOptions{})
			bc, _ := projectAdminBuildClient.BuildConfigs(testutil.Namespace()).Get(config.Name, metav1.GetOptions{})
			t.Fatalf("Expected build with base image %s, got %s\n, imagerepo is %v\trigger is %#v\n", registryHostname+"/openshift/test-image-trigger:ref-2-random", strategy.SourceStrategy.From.Name, i, bc.Spec.Triggers[3].ImageChange)
		}
	case strategy.DockerStrategy != nil:
		if strategy.DockerStrategy.From.Name != registryHostname+"/openshift/test-image-trigger:ref-2-random" {
			i, _ := projectAdminImageClient.ImageStreams(testutil.Namespace()).Get(imageStream.Name, metav1.GetOptions{})
			bc, _ := projectAdminBuildClient.BuildConfigs(testutil.Namespace()).Get(config.Name, metav1.GetOptions{})
			t.Fatalf("Expected build with base image %s, got %s\n, imagerepo is %v\trigger is %#v\n", registryHostname+"/openshift/test-image-trigger:ref-2-random", strategy.DockerStrategy.From.Name, i, bc.Spec.Triggers[3].ImageChange)
		}
	case strategy.CustomStrategy != nil:
		if strategy.CustomStrategy.From.Name != registryHostname+"/openshift/test-image-trigger:ref-2-random" {
			i, _ := projectAdminImageClient.ImageStreams(testutil.Namespace()).Get(imageStream.Name, metav1.GetOptions{})
			bc, _ := projectAdminBuildClient.BuildConfigs(testutil.Namespace()).Get(config.Name, metav1.GetOptions{})
			t.Fatalf("Expected build with base image %s, got %s\n, imagerepo is %v\trigger is %#v\n", registryHostname+"/openshift/test-image-trigger:ref-2-random", strategy.CustomStrategy.From.Name, i, bc.Spec.Triggers[3].ImageChange)
		}
	}
	for {
		event = <-buildWatch.ResultChan()
		newBuild = event.Object.(*buildv1.Build)
		if newBuild.Name != build1Name {
			break
		}
	}
	if e, a := watchapi.Modified, event.Type; e != a {
		t.Fatalf("expected watch event type %s, got %s", e, a)
	}
	if newBuild.Spec.Output.To.Name != "test-image-trigger-repo:outputtag" {
		t.Fatalf("unexpected build output: %#v %#v", newBuild.Spec.Output.To, newBuild.Spec.Output)
	}
	if newBuild.Labels["testlabel"] != "testvalue" {
		t.Fatalf("Expected build with label %s=%s from build config got %s=%s", "testlabel", "testvalue", "testlabel", newBuild.Labels["testlabel"])
	}
	<-buildConfigWatch.ResultChan()
	updatedConfig, err = projectAdminBuildClient.BuildConfigs(testutil.Namespace()).Get(config.Name, metav1.GetOptions{})
	if err != nil {
		t.Fatalf("Couldn't get BuildConfig: %v", err)
	}
	if e, a := registryHostname+"/openshift/test-image-trigger:ref-2-random", updatedConfig.Spec.Triggers[0].ImageChange.LastTriggeredImageID; e != a {
		t.Errorf("unexpected trigger id: expected %v, got %v", e, a)
	}
}
func TestMultipleImageChangeBuildTriggers(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	testutil.SetAdditionalAllowedRegistries("registry:5000")
	mockImageStream := func(name, tag string) *imageapi.ImageStream {
		return &imageapi.ImageStream{ObjectMeta: metav1.ObjectMeta{Name: name}, Spec: imageapi.ImageStreamSpec{DockerImageRepository: "registry:5000/openshift/" + name, Tags: map[string]imageapi.TagReference{tag: {From: &kapi.ObjectReference{Kind: "DockerImage", Name: "registry:5000/openshift/" + name + ":" + tag}}}}}
	}
	mockStreamMapping := func(name, tag string) *imageapi.ImageStreamMapping {
		return &imageapi.ImageStreamMapping{ObjectMeta: metav1.ObjectMeta{Name: name}, Tag: tag, Image: imageapi.Image{ObjectMeta: metav1.ObjectMeta{Name: name}, DockerImageReference: "registry:5000/openshift/" + name + ":" + tag}}
	}
	multipleImageChangeBuildConfig := func() *buildv1.BuildConfig {
		strategy := stiStrategy("ImageStreamTag", "image1:tag1")
		bc := imageChangeBuildConfig("multi-image-trigger", strategy)
		bc.Spec.CommonSpec.Output.To.Name = "image1:outputtag"
		bc.Spec.Triggers = []buildv1.BuildTriggerPolicy{{Type: buildv1.ImageChangeBuildTriggerType, ImageChange: &buildv1.ImageChangeTrigger{}}, {Type: buildv1.ImageChangeBuildTriggerType, ImageChange: &buildv1.ImageChangeTrigger{From: &corev1.ObjectReference{Name: "image2:tag2", Kind: "ImageStreamTag"}}}, {Type: buildv1.ImageChangeBuildTriggerType, ImageChange: &buildv1.ImageChangeTrigger{From: &corev1.ObjectReference{Name: "image3:tag3", Kind: "ImageStreamTag"}}}}
		return bc
	}
	_, _, projectAdminConfig, fn := setup(t)
	defer fn()
	config := multipleImageChangeBuildConfig()
	triggersToTest := []struct {
		triggerIndex	int
		name		string
		tag		string
	}{{triggerIndex: 0, name: "image1", tag: "tag1"}, {triggerIndex: 1, name: "image2", tag: "tag2"}, {triggerIndex: 2, name: "image3", tag: "tag3"}}
	projectAdminBuildClient := buildv1client.NewForConfigOrDie(projectAdminConfig).BuildV1()
	projectAdminImageClient := imageclient.NewForConfigOrDie(projectAdminConfig).Image()
	created, err := projectAdminBuildClient.BuildConfigs(testutil.Namespace()).Create(config)
	if err != nil {
		t.Fatalf("Couldn't create BuildConfig: %v", err)
	}
	buildWatch, err := projectAdminBuildClient.Builds(testutil.Namespace()).Watch(metav1.ListOptions{ResourceVersion: created.ResourceVersion})
	if err != nil {
		t.Fatalf("Couldn't subscribe to Builds %v", err)
	}
	defer buildWatch.Stop()
	buildConfigWatch, err := projectAdminBuildClient.BuildConfigs(testutil.Namespace()).Watch(metav1.ListOptions{ResourceVersion: created.ResourceVersion})
	if err != nil {
		t.Fatalf("Couldn't subscribe to BuildConfigs %v", err)
	}
	defer buildConfigWatch.Stop()
	ignoreBuilds := make(map[string]struct{})
	for _, tc := range triggersToTest {
		imageStream := mockImageStream(tc.name, tc.tag)
		imageStreamMapping := mockStreamMapping(tc.name, tc.tag)
		imageStream, err = projectAdminImageClient.ImageStreams(testutil.Namespace()).Create(imageStream)
		if err != nil {
			t.Fatalf("Couldn't create ImageStream: %v", err)
		}
		_, err = projectAdminImageClient.ImageStreamMappings(testutil.Namespace()).Create(imageStreamMapping)
		if err != nil {
			t.Fatalf("Couldn't create Image: %v", err)
		}
		var newBuild *buildv1.Build
		var event watchapi.Event
		newBuild, event = filterEvents(t, ignoreBuilds, buildWatch)
		if e, a := watchapi.Added, event.Type; e != a {
			t.Fatalf("expected watch event type %s, got %s", e, a)
		}
		trigger := config.Spec.Triggers[tc.triggerIndex]
		if trigger.ImageChange.From == nil {
			strategy := newBuild.Spec.Strategy
			switch {
			case strategy.SourceStrategy != nil:
				if strategy.SourceStrategy.From.Name != "registry:5000/openshift/"+tc.name+":"+tc.tag {
					i, _ := projectAdminImageClient.ImageStreams(testutil.Namespace()).Get(imageStream.Name, metav1.GetOptions{})
					bc, _ := projectAdminBuildClient.BuildConfigs(testutil.Namespace()).Get(config.Name, metav1.GetOptions{})
					t.Fatalf("Expected build with base image %s, got %s\n, imagerepo is %v\ntrigger is %#v", "registry:5000/openshift/"+tc.name+":"+tc.tag, strategy.SourceStrategy.From.Name, i, bc.Spec.Triggers[tc.triggerIndex].ImageChange)
				}
			case strategy.DockerStrategy != nil:
				if strategy.DockerStrategy.From.Name != registryHostname+"/openshift/"+tc.name+":"+tc.tag {
					i, _ := projectAdminImageClient.ImageStreams(testutil.Namespace()).Get(imageStream.Name, metav1.GetOptions{})
					bc, _ := projectAdminBuildClient.BuildConfigs(testutil.Namespace()).Get(config.Name, metav1.GetOptions{})
					t.Fatalf("Expected build with base image %s, got %s\n, imagerepo is %v\ntrigger is %#v", "registry:5000/openshift/"+tc.name+":"+tag, strategy.DockerStrategy.From.Name, i, bc.Spec.Triggers[tc.triggerIndex].ImageChange)
				}
			case strategy.CustomStrategy != nil:
				if strategy.CustomStrategy.From.Name != registryHostname+"/openshift/"+tc.name+":"+tag {
					i, _ := projectAdminImageClient.ImageStreams(testutil.Namespace()).Get(imageStream.Name, metav1.GetOptions{})
					bc, _ := projectAdminBuildClient.BuildConfigs(testutil.Namespace()).Get(config.Name, metav1.GetOptions{})
					t.Fatalf("Expected build with base image %s, got %s\n, imagerepo is %v\ntrigger is %#v", "registry:5000/openshift/"+tc.name+":"+tag, strategy.CustomStrategy.From.Name, i, bc.Spec.Triggers[tc.triggerIndex].ImageChange)
				}
			}
		}
		newBuild, event = filterEvents(t, ignoreBuilds, buildWatch)
		if e, a := watchapi.Modified, event.Type; e != a {
			t.Fatalf("expected watch event type %s, got %s", e, a)
		}
		if newBuild.Spec.Output.To.Name != "image1:outputtag" {
			t.Fatalf("unexpected build output: %#v %#v", newBuild.Spec.Output.To, newBuild.Spec.Output)
		}
		<-buildConfigWatch.ResultChan()
		updatedConfig, err := projectAdminBuildClient.BuildConfigs(testutil.Namespace()).Get(config.Name, metav1.GetOptions{})
		if err != nil {
			t.Fatalf("Couldn't get BuildConfig: %v", err)
		}
		if updatedConfig.Spec.Triggers[tc.triggerIndex].ImageChange.LastTriggeredImageID != "registry:5000/openshift/"+tc.name+":"+tc.tag {
			t.Fatalf("Expected imageID equal to pull spec, got %#v", updatedConfig.Spec.Triggers[0].ImageChange)
		}
		ignoreBuilds[newBuild.Name] = struct{}{}
	}
}
func filterEvents(t *testing.T, ignoreBuilds map[string]struct{}, buildWatch watchapi.Interface) (newBuild *buildv1.Build, event watchapi.Event) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for {
		event = <-buildWatch.ResultChan()
		var ok bool
		newBuild, ok = event.Object.(*buildv1.Build)
		if !ok {
			t.Errorf("unexpected event type (not a Build): %v", event.Object)
		}
		if _, exists := ignoreBuilds[newBuild.Name]; !exists {
			break
		}
	}
	return
}
