package strategy

import (
	"fmt"
	"strings"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	buildv1 "github.com/openshift/api/build/v1"
	securityv1 "github.com/openshift/api/security/v1"
	securityclient "github.com/openshift/client-go/security/clientset/versioned/typed/security/v1"
	"github.com/openshift/origin/pkg/build/buildapihelpers"
	buildutil "github.com/openshift/origin/pkg/build/util"
)

type SourceBuildStrategy struct {
	Image		string
	SecurityClient	securityclient.SecurityV1Interface
}

var DefaultDropCaps = []string{"KILL", "MKNOD", "SETGID", "SETUID"}

func (bs *SourceBuildStrategy) CreateBuildPod(build *buildv1.Build, additionalCAs map[string]string, internalRegistryHost string) (*corev1.Pod, error) {
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
	data, err := runtime.Encode(buildJSONCodec, build)
	if err != nil {
		return nil, fmt.Errorf("failed to encode the Build %s/%s: %v", build.Namespace, build.Name, err)
	}
	containerEnv := []corev1.EnvVar{{Name: "BUILD", Value: string(data)}, {Name: "LANG", Value: "en_US.utf8"}}
	addSourceEnvVars(build.Spec.Source, &containerEnv)
	strategy := build.Spec.Strategy.SourceStrategy
	if len(strategy.Env) > 0 {
		buildutil.MergeTrustedEnvWithoutDuplicates(strategy.Env, &containerEnv, true)
	}
	if !bs.canRunAsRoot(build) {
		containerEnv = append(containerEnv, corev1.EnvVar{Name: buildutil.AllowedUIDs, Value: "1-"})
		containerEnv = append(containerEnv, corev1.EnvVar{Name: buildutil.DropCapabilities, Value: strings.Join(DefaultDropCaps, ",")})
	}
	serviceAccount := build.Spec.ServiceAccount
	if len(serviceAccount) == 0 {
		serviceAccount = buildutil.BuilderServiceAccountName
	}
	privileged := true
	pod := &corev1.Pod{ObjectMeta: metav1.ObjectMeta{Name: buildapihelpers.GetBuildPodName(build), Namespace: build.Namespace, Labels: getPodLabels(build)}, Spec: corev1.PodSpec{ServiceAccountName: serviceAccount, Containers: []corev1.Container{{Name: StiBuild, Image: bs.Image, Command: []string{"openshift-sti-build"}, Env: copyEnvVarSlice(containerEnv), SecurityContext: &corev1.SecurityContext{Privileged: &privileged}, TerminationMessagePolicy: corev1.TerminationMessageFallbackToLogsOnError, VolumeMounts: []corev1.VolumeMount{{Name: "buildworkdir", MountPath: buildutil.BuildWorkDirMount}, {Name: "buildcachedir", MountPath: buildutil.BuildBlobsMetaCache}}, ImagePullPolicy: corev1.PullIfNotPresent, Resources: build.Spec.Resources}}, Volumes: []corev1.Volume{{Name: "buildcachedir", VolumeSource: corev1.VolumeSource{HostPath: &corev1.HostPathVolumeSource{Path: buildutil.BuildBlobsMetaCache}}}, {Name: "buildworkdir", VolumeSource: corev1.VolumeSource{EmptyDir: &corev1.EmptyDirVolumeSource{}}}}, RestartPolicy: corev1.RestartPolicyNever, NodeSelector: build.Spec.NodeSelector}}
	if build.Spec.Source.Git != nil || build.Spec.Source.Binary != nil {
		gitCloneContainer := corev1.Container{Name: GitCloneContainer, Image: bs.Image, Command: []string{"openshift-git-clone"}, Env: copyEnvVarSlice(containerEnv), TerminationMessagePolicy: corev1.TerminationMessageFallbackToLogsOnError, VolumeMounts: []corev1.VolumeMount{{Name: "buildworkdir", MountPath: buildutil.BuildWorkDirMount}}, ImagePullPolicy: corev1.PullIfNotPresent, Resources: build.Spec.Resources}
		if build.Spec.Source.Binary != nil {
			gitCloneContainer.Stdin = true
			gitCloneContainer.StdinOnce = true
		}
		setupSourceSecrets(pod, &gitCloneContainer, build.Spec.Source.SourceSecret)
		pod.Spec.InitContainers = append(pod.Spec.InitContainers, gitCloneContainer)
	}
	if len(build.Spec.Source.Images) > 0 {
		extractImageContentContainer := corev1.Container{Name: ExtractImageContentContainer, Image: bs.Image, Command: []string{"openshift-extract-image-content"}, Env: copyEnvVarSlice(containerEnv), SecurityContext: &corev1.SecurityContext{Privileged: &privileged}, TerminationMessagePolicy: corev1.TerminationMessageFallbackToLogsOnError, VolumeMounts: []corev1.VolumeMount{{Name: "buildworkdir", MountPath: buildutil.BuildWorkDirMount}, {Name: "buildcachedir", MountPath: buildutil.BuildBlobsMetaCache}}, ImagePullPolicy: corev1.PullIfNotPresent, Resources: build.Spec.Resources}
		setupDockerSecrets(pod, &extractImageContentContainer, build.Spec.Output.PushSecret, strategy.PullSecret, build.Spec.Source.Images)
		setupContainersStorage(pod, &extractImageContentContainer)
		pod.Spec.InitContainers = append(pod.Spec.InitContainers, extractImageContentContainer)
	}
	pod.Spec.InitContainers = append(pod.Spec.InitContainers, corev1.Container{Name: "manage-dockerfile", Image: bs.Image, Command: []string{"openshift-manage-dockerfile"}, Env: copyEnvVarSlice(containerEnv), TerminationMessagePolicy: corev1.TerminationMessageFallbackToLogsOnError, VolumeMounts: []corev1.VolumeMount{{Name: "buildworkdir", MountPath: buildutil.BuildWorkDirMount}}, ImagePullPolicy: corev1.PullIfNotPresent, Resources: build.Spec.Resources})
	if build.Spec.CompletionDeadlineSeconds != nil {
		pod.Spec.ActiveDeadlineSeconds = build.Spec.CompletionDeadlineSeconds
	}
	setOwnerReference(pod, build)
	setupDockerSecrets(pod, &pod.Spec.Containers[0], build.Spec.Output.PushSecret, strategy.PullSecret, build.Spec.Source.Images)
	setupInputSecrets(pod, &pod.Spec.Containers[0], build.Spec.Source.Secrets)
	setupInputConfigMaps(pod, &pod.Spec.Containers[0], build.Spec.Source.ConfigMaps)
	setupContainersConfigs(build, pod)
	setupBuildCAs(build, pod, additionalCAs, internalRegistryHost)
	setupContainersStorage(pod, &pod.Spec.Containers[0])
	setupBlobCache(pod)
	return pod, nil
}
func (bs *SourceBuildStrategy) canRunAsRoot(build *buildv1.Build) bool {
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
	rootUser := int64(0)
	review, err := bs.SecurityClient.PodSecurityPolicySubjectReviews(build.Namespace).Create(&securityv1.PodSecurityPolicySubjectReview{Spec: securityv1.PodSecurityPolicySubjectReviewSpec{Template: corev1.PodTemplateSpec{Spec: corev1.PodSpec{ServiceAccountName: build.Spec.ServiceAccount, Containers: []corev1.Container{{Name: "fake", Image: "fake", SecurityContext: &corev1.SecurityContext{RunAsUser: &rootUser}}}}}}})
	if err != nil {
		utilruntime.HandleError(err)
		return false
	}
	return review.Status.AllowedBy != nil
}
