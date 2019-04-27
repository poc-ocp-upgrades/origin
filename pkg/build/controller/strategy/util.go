package strategy

import (
	"fmt"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/apis/policy"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kvalidation "k8s.io/apimachinery/pkg/util/validation"
	buildv1 "github.com/openshift/api/build/v1"
	"github.com/openshift/origin/pkg/api/apihelpers"
	"github.com/openshift/origin/pkg/build/buildapihelpers"
	buildutil "github.com/openshift/origin/pkg/build/util"
	"github.com/openshift/origin/pkg/image/apis/image/reference"
)

const (
	dockerSocketPath			= "/var/run/docker.sock"
	sourceSecretMountPath			= "/var/run/secrets/openshift.io/source"
	DockerPushSecretMountPath		= "/var/run/secrets/openshift.io/push"
	DockerPullSecretMountPath		= "/var/run/secrets/openshift.io/pull"
	ConfigMapBuildSourceBaseMountPath	= "/var/run/configs/openshift.io/build"
	ConfigMapBuildSystemConfigsMountPath	= "/var/run/configs/openshift.io/build-system"
	ConfigMapCertsMountPath			= "/var/run/configs/openshift.io/certs"
	SecretBuildSourceBaseMountPath		= "/var/run/secrets/openshift.io/build"
	SourceImagePullSecretMountPath		= "/var/run/secrets/openshift.io/source-image"
	ExtractImageContentContainer		= "extract-image-content"
	GitCloneContainer			= "git-clone"
)
const (
	CustomBuild	= "custom-build"
	DockerBuild	= "docker-build"
	StiBuild	= "sti-build"
)

var BuildContainerNames = []string{CustomBuild, StiBuild, DockerBuild}
var (
	BuildControllerRefKind = buildv1.GroupVersion.WithKind("Build")
)
var hostPortRegex = regexp.MustCompile("\\.\\.(\\d+)$")

type FatalError struct{ Reason string }

func (e *FatalError) Error() string {
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
	return fmt.Sprintf("fatal error: %s", e.Reason)
}
func IsFatal(err error) bool {
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
	_, isFatal := err.(*FatalError)
	return isFatal
}
func setupDockerSocket(pod *corev1.Pod) {
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
	dockerSocketVolume := corev1.Volume{Name: "docker-socket", VolumeSource: corev1.VolumeSource{HostPath: &corev1.HostPathVolumeSource{Path: dockerSocketPath}}}
	dockerSocketVolumeMount := corev1.VolumeMount{Name: "docker-socket", MountPath: dockerSocketPath}
	pod.Spec.Volumes = append(pod.Spec.Volumes, dockerSocketVolume)
	pod.Spec.Containers[0].VolumeMounts = append(pod.Spec.Containers[0].VolumeMounts, dockerSocketVolumeMount)
	for i, initContainer := range pod.Spec.InitContainers {
		if initContainer.Name == ExtractImageContentContainer {
			pod.Spec.InitContainers[i].VolumeMounts = append(pod.Spec.InitContainers[i].VolumeMounts, dockerSocketVolumeMount)
			break
		}
	}
}
func mountConfigMapVolume(pod *corev1.Pod, container *corev1.Container, configMapName, mountPath, volumeSuffix string) {
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
	mountVolume(pod, container, configMapName, mountPath, volumeSuffix, policy.ConfigMap)
}
func mountSecretVolume(pod *corev1.Pod, container *corev1.Container, secretName, mountPath, volumeSuffix string) {
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
	mountVolume(pod, container, secretName, mountPath, volumeSuffix, policy.Secret)
}
func mountVolume(pod *corev1.Pod, container *corev1.Container, objName, mountPath, volumeSuffix string, fsType policy.FSType) {
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
	volumeName := apihelpers.GetName(objName, volumeSuffix, kvalidation.DNS1123LabelMaxLength)
	volumeName = strings.Replace(volumeName, ".", "-", -1)
	volumeExists := false
	for _, v := range pod.Spec.Volumes {
		if v.Name == volumeName {
			volumeExists = true
			break
		}
	}
	mode := int32(0600)
	if !volumeExists {
		volume := makeVolume(volumeName, objName, mode, fsType)
		pod.Spec.Volumes = append(pod.Spec.Volumes, volume)
	}
	volumeMount := corev1.VolumeMount{Name: volumeName, MountPath: mountPath, ReadOnly: true}
	container.VolumeMounts = append(container.VolumeMounts, volumeMount)
}
func makeVolume(volumeName, refName string, mode int32, fsType policy.FSType) corev1.Volume {
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
	vol := corev1.Volume{Name: volumeName, VolumeSource: corev1.VolumeSource{}}
	switch fsType {
	case policy.ConfigMap:
		vol.VolumeSource.ConfigMap = &corev1.ConfigMapVolumeSource{LocalObjectReference: corev1.LocalObjectReference{Name: refName}, DefaultMode: &mode}
	case policy.EmptyDir:
		vol.VolumeSource.EmptyDir = &corev1.EmptyDirVolumeSource{}
	case policy.Secret:
		vol.VolumeSource.Secret = &corev1.SecretVolumeSource{SecretName: refName, DefaultMode: &mode}
	default:
		klog.V(3).Infof("File system %s is not supported for volumes. Using empty directory instead.", fsType)
		vol.VolumeSource.EmptyDir = &corev1.EmptyDirVolumeSource{}
	}
	return vol
}
func setupDockerSecrets(pod *corev1.Pod, container *corev1.Container, pushSecret, pullSecret *corev1.LocalObjectReference, imageSources []buildv1.ImageSource) {
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
	if pushSecret != nil {
		mountSecretVolume(pod, container, pushSecret.Name, DockerPushSecretMountPath, "push")
		container.Env = append(container.Env, []corev1.EnvVar{{Name: "PUSH_DOCKERCFG_PATH", Value: DockerPushSecretMountPath}}...)
		klog.V(3).Infof("%s will be used for docker push in %s", DockerPushSecretMountPath, pod.Name)
	}
	if pullSecret != nil {
		mountSecretVolume(pod, container, pullSecret.Name, DockerPullSecretMountPath, "pull")
		container.Env = append(container.Env, []corev1.EnvVar{{Name: "PULL_DOCKERCFG_PATH", Value: DockerPullSecretMountPath}}...)
		klog.V(3).Infof("%s will be used for docker pull in %s", DockerPullSecretMountPath, pod.Name)
	}
	for i, imageSource := range imageSources {
		if imageSource.PullSecret == nil {
			continue
		}
		mountPath := filepath.Join(SourceImagePullSecretMountPath, strconv.Itoa(i))
		mountSecretVolume(pod, container, imageSource.PullSecret.Name, mountPath, fmt.Sprintf("%s%d", "source-image", i))
		container.Env = append(container.Env, []corev1.EnvVar{{Name: fmt.Sprintf("%s%d", "PULL_SOURCE_DOCKERCFG_PATH_", i), Value: mountPath}}...)
		klog.V(3).Infof("%s will be used for docker pull in %s", mountPath, pod.Name)
	}
}
func setupSourceSecrets(pod *corev1.Pod, container *corev1.Container, sourceSecret *corev1.LocalObjectReference) {
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
	if sourceSecret == nil {
		return
	}
	mountSecretVolume(pod, container, sourceSecret.Name, sourceSecretMountPath, "source")
	klog.V(3).Infof("Installed source secrets in %s, in Pod %s/%s", sourceSecretMountPath, pod.Namespace, pod.Name)
	container.Env = append(container.Env, []corev1.EnvVar{{Name: "SOURCE_SECRET_PATH", Value: sourceSecretMountPath}}...)
}
func setupInputConfigMaps(pod *corev1.Pod, container *corev1.Container, configs []buildv1.ConfigMapBuildSource) {
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
	for _, c := range configs {
		mountConfigMapVolume(pod, container, c.ConfigMap.Name, filepath.Join(ConfigMapBuildSourceBaseMountPath, c.ConfigMap.Name), "build")
		klog.V(3).Infof("%s will be used as a build config in %s", c.ConfigMap.Name, ConfigMapBuildSourceBaseMountPath)
	}
}
func setupInputSecrets(pod *corev1.Pod, container *corev1.Container, secrets []buildv1.SecretBuildSource) {
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
	for _, s := range secrets {
		mountSecretVolume(pod, container, s.Secret.Name, filepath.Join(SecretBuildSourceBaseMountPath, s.Secret.Name), "build")
		klog.V(3).Infof("%s will be used as a build secret in %s", s.Secret.Name, SecretBuildSourceBaseMountPath)
	}
}
func addSourceEnvVars(source buildv1.BuildSource, output *[]corev1.EnvVar) {
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
	sourceVars := []corev1.EnvVar{}
	if source.Git != nil {
		sourceVars = append(sourceVars, corev1.EnvVar{Name: "SOURCE_REPOSITORY", Value: source.Git.URI})
		sourceVars = append(sourceVars, corev1.EnvVar{Name: "SOURCE_URI", Value: source.Git.URI})
	}
	if len(source.ContextDir) > 0 {
		sourceVars = append(sourceVars, corev1.EnvVar{Name: "SOURCE_CONTEXT_DIR", Value: source.ContextDir})
	}
	if source.Git != nil && len(source.Git.Ref) > 0 {
		sourceVars = append(sourceVars, corev1.EnvVar{Name: "SOURCE_REF", Value: source.Git.Ref})
	}
	*output = append(*output, sourceVars...)
}
func addOutputEnvVars(buildOutput *corev1.ObjectReference, output *[]corev1.EnvVar) error {
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
	if buildOutput == nil {
		return nil
	}
	if buildOutput.Kind != "DockerImage" {
		return fmt.Errorf("invalid build output kind %s, must be DockerImage", buildOutput.Kind)
	}
	ref, err := reference.Parse(buildOutput.Name)
	if err != nil {
		return err
	}
	registry := ref.Registry
	ref.Registry = ""
	image := ref.String()
	outputVars := []corev1.EnvVar{{Name: "OUTPUT_REGISTRY", Value: registry}, {Name: "OUTPUT_IMAGE", Value: image}}
	*output = append(*output, outputVars...)
	return nil
}
func setupAdditionalSecrets(pod *corev1.Pod, container *corev1.Container, secrets []buildv1.SecretSpec) {
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
	for _, secretSpec := range secrets {
		mountSecretVolume(pod, container, secretSpec.SecretSource.Name, secretSpec.MountPath, "secret")
		klog.V(3).Infof("Installed additional secret in %s, in Pod %s/%s", secretSpec.MountPath, pod.Namespace, pod.Name)
	}
}
func getPodLabels(build *buildv1.Build) map[string]string {
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
	return map[string]string{buildutil.BuildLabel: buildapihelpers.LabelValue(build.Name)}
}
func makeOwnerReference(build *buildv1.Build) metav1.OwnerReference {
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
	t := true
	return metav1.OwnerReference{APIVersion: BuildControllerRefKind.GroupVersion().String(), Kind: BuildControllerRefKind.Kind, Name: build.Name, UID: build.UID, Controller: &t}
}
func setOwnerReference(pod *corev1.Pod, build *buildv1.Build) {
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
	pod.OwnerReferences = []metav1.OwnerReference{makeOwnerReference(build)}
}
func HasOwnerReference(pod *corev1.Pod, build *buildv1.Build) bool {
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
	ref := makeOwnerReference(build)
	for _, r := range pod.OwnerReferences {
		if reflect.DeepEqual(r, ref) {
			return true
		}
	}
	return false
}
func copyEnvVarSlice(in []corev1.EnvVar) []corev1.EnvVar {
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
	out := make([]corev1.EnvVar, len(in))
	copy(out, in)
	return out
}
func setupContainersConfigs(build *buildv1.Build, pod *corev1.Pod) {
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
	const volumeName = "build-system-configs"
	const configDir = ConfigMapBuildSystemConfigsMountPath
	exists := false
	for _, v := range pod.Spec.Volumes {
		if v.Name == volumeName {
			exists = true
			break
		}
	}
	if !exists {
		cmSource := &corev1.ConfigMapVolumeSource{LocalObjectReference: corev1.LocalObjectReference{Name: buildapihelpers.GetBuildSystemConfigMapName(build)}}
		pod.Spec.Volumes = append(pod.Spec.Volumes, corev1.Volume{Name: volumeName, VolumeSource: corev1.VolumeSource{ConfigMap: cmSource}})
		containers := make([]corev1.Container, len(pod.Spec.Containers))
		for i, c := range pod.Spec.Containers {
			containers[i] = updateConfigsForContainer(c, volumeName, configDir)
		}
		pod.Spec.Containers = containers
		if len(pod.Spec.InitContainers) > 0 {
			initContainers := make([]corev1.Container, len(pod.Spec.InitContainers))
			for i, c := range pod.Spec.InitContainers {
				initContainers[i] = updateConfigsForContainer(c, volumeName, configDir)
			}
			pod.Spec.InitContainers = initContainers
		}
	}
}
func updateConfigsForContainer(c corev1.Container, volumeName string, configDir string) corev1.Container {
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
	c.VolumeMounts = append(c.VolumeMounts, corev1.VolumeMount{Name: volumeName, MountPath: configDir, ReadOnly: true})
	registriesConfPath := filepath.Join(configDir, buildutil.RegistryConfKey)
	signaturePolicyPath := filepath.Join(configDir, buildutil.SignaturePolicyKey)
	registriesDirPath := filepath.Join(configDir, "registries.d")
	storageConfPath := filepath.Join(configDir, "storage.conf")
	c.Env = append(c.Env, corev1.EnvVar{Name: "BUILD_REGISTRIES_CONF_PATH", Value: registriesConfPath})
	c.Env = append(c.Env, corev1.EnvVar{Name: "BUILD_REGISTRIES_DIR_PATH", Value: registriesDirPath})
	c.Env = append(c.Env, corev1.EnvVar{Name: "BUILD_SIGNATURE_POLICY_PATH", Value: signaturePolicyPath})
	c.Env = append(c.Env, corev1.EnvVar{Name: "BUILD_STORAGE_CONF_PATH", Value: storageConfPath})
	return c
}
func setupContainersStorage(pod *corev1.Pod, container *corev1.Container) {
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
	exists := false
	for _, v := range pod.Spec.Volumes {
		if v.Name == "container-storage-root" {
			exists = true
			break
		}
	}
	if !exists {
		pod.Spec.Volumes = append(pod.Spec.Volumes, corev1.Volume{Name: "container-storage-root", VolumeSource: corev1.VolumeSource{EmptyDir: &corev1.EmptyDirVolumeSource{}}})
	}
	container.VolumeMounts = append(container.VolumeMounts, corev1.VolumeMount{Name: "container-storage-root", MountPath: "/var/lib/containers/storage"})
	container.Env = append(container.Env, corev1.EnvVar{Name: "BUILD_STORAGE_DRIVER", Value: "overlay"})
	container.Env = append(container.Env, corev1.EnvVar{Name: "BUILD_ISOLATION", Value: "chroot"})
}
func setupContainersNodeStorage(pod *corev1.Pod, container *corev1.Container) {
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
	exists := false
	for _, v := range pod.Spec.Volumes {
		if v.Name == "node-storage-root" {
			exists = true
			break
		}
	}
	if !exists {
		pod.Spec.Volumes = append(pod.Spec.Volumes, corev1.Volume{Name: "node-storage-root", VolumeSource: corev1.VolumeSource{HostPath: &corev1.HostPathVolumeSource{Path: "/var/lib/containers/storage"}}})
	}
	container.VolumeMounts = append(container.VolumeMounts, corev1.VolumeMount{Name: "node-storage-root", MountPath: "/var/lib/containers/storage"})
	container.Env = append(container.Env, corev1.EnvVar{Name: "BUILD_STORAGE_DRIVER", Value: "overlay"})
	container.Env = append(container.Env, corev1.EnvVar{Name: "BUILD_ISOLATION", Value: "chroot"})
}
func setupBuildCAs(build *buildv1.Build, pod *corev1.Pod, additionalCAs map[string]string, internalRegistryHost string) {
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
	casExist := false
	for _, v := range pod.Spec.Volumes {
		if v.Name == "build-ca-bundles" {
			casExist = true
			break
		}
	}
	if !casExist {
		cmSource := &corev1.ConfigMapVolumeSource{LocalObjectReference: corev1.LocalObjectReference{Name: buildapihelpers.GetBuildCAConfigMapName(build)}, Items: []corev1.KeyToPath{{Key: buildutil.ServiceCAKey, Path: fmt.Sprintf("certs.d/%s/ca.crt", internalRegistryHost)}}}
		for key := range additionalCAs {
			mountDir := hostPortRegex.ReplaceAllString(key, ":$1")
			cmSource.Items = append(cmSource.Items, corev1.KeyToPath{Key: key, Path: fmt.Sprintf("certs.d/%s/ca.crt", mountDir)})
		}
		pod.Spec.Volumes = append(pod.Spec.Volumes, corev1.Volume{Name: "build-ca-bundles", VolumeSource: corev1.VolumeSource{ConfigMap: cmSource}})
		containers := make([]corev1.Container, len(pod.Spec.Containers))
		for i, c := range pod.Spec.Containers {
			c.VolumeMounts = append(c.VolumeMounts, corev1.VolumeMount{Name: "build-ca-bundles", MountPath: ConfigMapCertsMountPath})
			containers[i] = c
		}
		pod.Spec.Containers = containers
	}
}
func setupBlobCache(pod *corev1.Pod) {
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
	const volume = "build-blob-cache"
	const mountPath = buildutil.BuildBlobsContentCache
	exists := false
	for _, v := range pod.Spec.Volumes {
		if v.Name == volume {
			exists = true
			break
		}
	}
	if !exists {
		pod.Spec.Volumes = append(pod.Spec.Volumes, corev1.Volume{Name: volume, VolumeSource: corev1.VolumeSource{EmptyDir: &corev1.EmptyDirVolumeSource{}}})
		containers := make([]corev1.Container, len(pod.Spec.Containers))
		for i, c := range pod.Spec.Containers {
			c.VolumeMounts = append(c.VolumeMounts, corev1.VolumeMount{Name: volume, MountPath: mountPath})
			c.Env = append(c.Env, corev1.EnvVar{Name: "BUILD_BLOBCACHE_DIR", Value: mountPath})
			containers[i] = c
		}
		pod.Spec.Containers = containers
		initContainers := make([]corev1.Container, len(pod.Spec.InitContainers))
		for i, ic := range pod.Spec.InitContainers {
			ic.VolumeMounts = append(ic.VolumeMounts, corev1.VolumeMount{Name: volume, MountPath: mountPath})
			ic.Env = append(ic.Env, corev1.EnvVar{Name: "BUILD_BLOBCACHE_DIR", Value: mountPath})
			initContainers[i] = ic
		}
		pod.Spec.InitContainers = initContainers
	}
}
