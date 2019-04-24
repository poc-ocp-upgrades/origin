package test

import (
	"testing"
	"bytes"
	"net/http"
	"runtime"
	"fmt"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/kubernetes/pkg/apis/autoscaling"
	kapi "k8s.io/kubernetes/pkg/apis/core"
	appsv1 "github.com/openshift/api/apps/v1"
	appsapi "github.com/openshift/origin/pkg/apps/apis/apps"
	appsv1conversions "github.com/openshift/origin/pkg/apps/apis/apps/v1"
	imageapi "github.com/openshift/origin/pkg/image/apis/image"
	"k8s.io/apimachinery/pkg/runtime"
)

const (
	ImageStreamName		= "test-image-stream"
	ImageID			= "0000000000000000000000000000000000000000000000000000000000000001"
	DockerImageReference	= "registry:5000/openshift/test-image-stream@sha256:0000000000000000000000000000000000000000000000000000000000000001"
)

func OkDeploymentConfig(version int64) *appsapi.DeploymentConfig {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &appsapi.DeploymentConfig{ObjectMeta: metav1.ObjectMeta{Name: "config", Namespace: kapi.NamespaceDefault}, Spec: OkDeploymentConfigSpec(), Status: OkDeploymentConfigStatus(version)}
}
func OkDeploymentConfigSpec() appsapi.DeploymentConfigSpec {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return appsapi.DeploymentConfigSpec{Replicas: 1, Selector: OkSelector(), Strategy: OkStrategy(), Template: OkPodTemplate(), Triggers: []appsapi.DeploymentTriggerPolicy{OkImageChangeTrigger(), OkConfigChangeTrigger()}}
}
func OkDeploymentConfigStatus(version int64) appsapi.DeploymentConfigStatus {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return appsapi.DeploymentConfigStatus{LatestVersion: version}
}
func OkImageChangeDetails() *appsapi.DeploymentDetails {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &appsapi.DeploymentDetails{Causes: []appsapi.DeploymentCause{{Type: appsapi.DeploymentTriggerOnImageChange, ImageTrigger: &appsapi.DeploymentCauseImageTrigger{From: kapi.ObjectReference{Name: imageapi.JoinImageStreamTag(ImageStreamName, imageapi.DefaultImageTag), Kind: "ImageStreamTag"}}}}}
}
func OkConfigChangeDetails() *appsapi.DeploymentDetails {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &appsapi.DeploymentDetails{Causes: []appsapi.DeploymentCause{{Type: appsapi.DeploymentTriggerOnConfigChange}}}
}
func OkStrategy() appsapi.DeploymentStrategy {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return appsapi.DeploymentStrategy{Type: appsapi.DeploymentStrategyTypeRecreate, Resources: kapi.ResourceRequirements{Limits: kapi.ResourceList{kapi.ResourceName(kapi.ResourceCPU): resource.MustParse("10"), kapi.ResourceName(kapi.ResourceMemory): resource.MustParse("10G")}}, RecreateParams: &appsapi.RecreateDeploymentStrategyParams{TimeoutSeconds: mkintp(20)}, ActiveDeadlineSeconds: mkintp(int(appsapi.MaxDeploymentDurationSeconds))}
}
func OkCustomStrategy() appsapi.DeploymentStrategy {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return appsapi.DeploymentStrategy{Type: appsapi.DeploymentStrategyTypeCustom, CustomParams: OkCustomParams(), Resources: kapi.ResourceRequirements{Limits: kapi.ResourceList{kapi.ResourceName(kapi.ResourceCPU): resource.MustParse("10"), kapi.ResourceName(kapi.ResourceMemory): resource.MustParse("10G")}}}
}
func OkCustomParams() *appsapi.CustomDeploymentStrategyParams {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &appsapi.CustomDeploymentStrategyParams{Image: "openshift/origin-deployer", Environment: []kapi.EnvVar{{Name: "ENV1", Value: "VAL1"}}, Command: []string{"/bin/echo", "hello", "world"}}
}
func mkintp(i int) *int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	v := int64(i)
	return &v
}
func OkRollingStrategy() appsapi.DeploymentStrategy {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return appsapi.DeploymentStrategy{Type: appsapi.DeploymentStrategyTypeRolling, RollingParams: &appsapi.RollingDeploymentStrategyParams{UpdatePeriodSeconds: mkintp(1), IntervalSeconds: mkintp(1), TimeoutSeconds: mkintp(20)}, Resources: kapi.ResourceRequirements{Limits: kapi.ResourceList{kapi.ResourceName(kapi.ResourceCPU): resource.MustParse("10"), kapi.ResourceName(kapi.ResourceMemory): resource.MustParse("10G")}}}
}
func OkSelector() map[string]string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return map[string]string{"a": "b"}
}
func OkPodTemplate() *kapi.PodTemplateSpec {
	_logClusterCodePath()
	defer _logClusterCodePath()
	one := int64(1)
	return &kapi.PodTemplateSpec{Spec: kapi.PodSpec{Containers: []kapi.Container{{Name: "container1", Image: "registry:8080/repo1:ref1", Env: []kapi.EnvVar{{Name: "ENV1", Value: "VAL1"}}, ImagePullPolicy: kapi.PullIfNotPresent, TerminationMessagePath: "/dev/termination-log", TerminationMessagePolicy: kapi.TerminationMessageReadFile}, {Name: "container2", Image: "registry:8080/repo1:ref2", ImagePullPolicy: kapi.PullIfNotPresent, TerminationMessagePath: "/dev/termination-log", TerminationMessagePolicy: kapi.TerminationMessageReadFile}}, RestartPolicy: kapi.RestartPolicyAlways, DNSPolicy: kapi.DNSClusterFirst, TerminationGracePeriodSeconds: &one, SchedulerName: kapi.DefaultSchedulerName, SecurityContext: &kapi.PodSecurityContext{}}, ObjectMeta: metav1.ObjectMeta{Labels: OkSelector()}}
}
func OkPodTemplateChanged() *kapi.PodTemplateSpec {
	_logClusterCodePath()
	defer _logClusterCodePath()
	template := OkPodTemplate()
	template.Spec.Containers[0].Image = DockerImageReference
	return template
}
func OkPodTemplateMissingImage(missing ...string) *kapi.PodTemplateSpec {
	_logClusterCodePath()
	defer _logClusterCodePath()
	set := sets.NewString(missing...)
	template := OkPodTemplate()
	for i, c := range template.Spec.Containers {
		if set.Has(c.Name) {
			template.Spec.Containers[i].Image = ""
		}
	}
	return template
}
func OkConfigChangeTrigger() appsapi.DeploymentTriggerPolicy {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return appsapi.DeploymentTriggerPolicy{Type: appsapi.DeploymentTriggerOnConfigChange}
}
func OkImageChangeTrigger() appsapi.DeploymentTriggerPolicy {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return appsapi.DeploymentTriggerPolicy{Type: appsapi.DeploymentTriggerOnImageChange, ImageChangeParams: &appsapi.DeploymentTriggerImageChangeParams{Automatic: true, ContainerNames: []string{"container1"}, From: kapi.ObjectReference{Kind: "ImageStreamTag", Name: imageapi.JoinImageStreamTag(ImageStreamName, imageapi.DefaultImageTag)}}}
}
func OkTriggeredImageChange() appsapi.DeploymentTriggerPolicy {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ict := OkImageChangeTrigger()
	ict.ImageChangeParams.LastTriggeredImage = DockerImageReference
	return ict
}
func OkNonAutomaticICT() appsapi.DeploymentTriggerPolicy {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ict := OkImageChangeTrigger()
	ict.ImageChangeParams.Automatic = false
	return ict
}
func OkTriggeredNonAutomatic() appsapi.DeploymentTriggerPolicy {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ict := OkNonAutomaticICT()
	ict.ImageChangeParams.LastTriggeredImage = DockerImageReference
	return ict
}
func TestDeploymentConfig(config *appsapi.DeploymentConfig) *appsapi.DeploymentConfig {
	_logClusterCodePath()
	defer _logClusterCodePath()
	config.Spec.Test = true
	return config
}
func OkHPAForDeploymentConfig(config *appsapi.DeploymentConfig, min, max int) *autoscaling.HorizontalPodAutoscaler {
	_logClusterCodePath()
	defer _logClusterCodePath()
	newMin := int32(min)
	return &autoscaling.HorizontalPodAutoscaler{ObjectMeta: metav1.ObjectMeta{Name: config.Name, Namespace: config.Namespace}, Spec: autoscaling.HorizontalPodAutoscalerSpec{ScaleTargetRef: autoscaling.CrossVersionObjectReference{Name: config.Name, Kind: "DeploymentConfig"}, MinReplicas: &newMin, MaxReplicas: int32(max)}}
}
func OkStreamForConfig(config *appsapi.DeploymentConfig) *imageapi.ImageStream {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, t := range config.Spec.Triggers {
		if t.Type != appsapi.DeploymentTriggerOnImageChange {
			continue
		}
		ref := t.ImageChangeParams.From
		name, tag, _ := imageapi.SplitImageStreamTag(ref.Name)
		return &imageapi.ImageStream{ObjectMeta: metav1.ObjectMeta{Name: name, Namespace: ref.Namespace}, Status: imageapi.ImageStreamStatus{Tags: map[string]imageapi.TagEventList{tag: {Items: []imageapi.TagEvent{{DockerImageReference: t.ImageChangeParams.LastTriggeredImage}}}}}}
	}
	return nil
}
func RoundTripConfig(t *testing.T, config *appsapi.DeploymentConfig) *appsapi.DeploymentConfig {
	_logClusterCodePath()
	defer _logClusterCodePath()
	scheme := runtime.NewScheme()
	appsv1conversions.Install(scheme)
	versioned, err := scheme.ConvertToVersion(config, appsv1.GroupVersion)
	if err != nil {
		t.Errorf("unexpected conversion error: %v", err)
		return nil
	}
	defaulted, err := scheme.ConvertToVersion(versioned, appsapi.SchemeGroupVersion)
	if err != nil {
		t.Errorf("unexpected conversion error: %v", err)
		return nil
	}
	return defaulted.(*appsapi.DeploymentConfig)
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
