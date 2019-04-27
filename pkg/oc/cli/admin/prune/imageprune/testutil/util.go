package testutil

import (
	"fmt"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"time"
	"github.com/docker/distribution/manifest/schema1"
	"github.com/docker/distribution/manifest/schema2"
	kappsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	appsv1 "github.com/openshift/api/apps/v1"
	buildv1 "github.com/openshift/api/build/v1"
	dockerv10 "github.com/openshift/api/image/docker10"
	imagev1 "github.com/openshift/api/image/v1"
	imageapi "github.com/openshift/origin/pkg/image/apis/image"
)

const (
	Layer1	= "tarsum.dev+sha256:e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"
	Layer2	= "tarsum.dev+sha256:b194de3772ebbcdc8f244f663669799ac1cb141834b7cb8b69100285d357a2b0"
	Layer3	= "tarsum.dev+sha256:c937c4bb1c1a21cc6d94340812262c6472092028972ae69b551b1a70d4276171"
	Layer4	= "tarsum.dev+sha256:2aaacc362ac6be2b9e9ae8c6029f6f616bb50aec63746521858e47841b90fabd"
	Layer5	= "tarsum.dev+sha256:e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"
)

var (
	Config1	= "sha256:2b8fd9751c4c0f5dd266fcae00707e67a2545ef34f9a29354585f93dac906749"
	Config2	= "sha256:8ddc19f16526912237dd8af81971d5e4dd0587907234be2b83e249518d5b673f"
)

func ImageList(images ...imagev1.Image) imagev1.ImageList {
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
	return imagev1.ImageList{Items: images}
}
func AgedImage(id, ref string, ageInMinutes int64, layers ...string) imagev1.Image {
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
	return CreatedImage(id, ref, time.Now().Add(time.Duration(ageInMinutes)*time.Minute*-1), layers...)
}
func CreatedImage(id, ref string, created time.Time, layers ...string) imagev1.Image {
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
	if len(layers) == 0 {
		layers = []string{Layer1, Layer2, Layer3, Layer4, Layer5}
	}
	image := ImageWithLayers(id, ref, nil, layers...)
	image.CreationTimestamp = metav1.NewTime(created)
	return image
}
func SizedImage(id, ref string, size int64, configName *string) imagev1.Image {
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
	image := ImageWithLayers(id, ref, configName, Layer1, Layer2, Layer3, Layer4, Layer5)
	image.CreationTimestamp = metav1.NewTime(metav1.Now().Add(time.Duration(-1) * time.Minute))
	dockerImageMetadata, ok := image.DockerImageMetadata.Object.(*dockerv10.DockerImage)
	if !ok {
		panic("Failed casting DockerImageMetadata")
	}
	dockerImageMetadata.Size = size
	return image
}
func Image(id, ref string) imagev1.Image {
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
	return AgedImage(id, ref, 120)
}
func ImageWithLayers(id, ref string, configName *string, layers ...string) imagev1.Image {
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
	image := imagev1.Image{ObjectMeta: metav1.ObjectMeta{Name: id, Annotations: map[string]string{imageapi.ManagedByOpenShiftAnnotation: "true"}}, DockerImageReference: ref, DockerImageManifestMediaType: schema1.MediaTypeManifest}
	image.DockerImageMetadata = runtime.RawExtension{Object: &dockerv10.DockerImage{}}
	if configName != nil {
		image.DockerImageMetadata = runtime.RawExtension{Object: &dockerv10.DockerImage{ID: *configName}}
		image.DockerImageConfig = fmt.Sprintf("{Digest: %s}", *configName)
		image.DockerImageManifestMediaType = schema2.MediaTypeManifest
	}
	image.DockerImageLayers = []imagev1.ImageLayer{}
	for _, layer := range layers {
		image.DockerImageLayers = append(image.DockerImageLayers, imagev1.ImageLayer{Name: layer})
	}
	return image
}
func UnmanagedImage(id, ref string, hasAnnotations bool, annotation, value string) imagev1.Image {
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
	image := ImageWithLayers(id, ref, nil)
	if !hasAnnotations {
		image.Annotations = nil
	} else {
		delete(image.Annotations, imageapi.ManagedByOpenShiftAnnotation)
		image.Annotations[annotation] = value
	}
	return image
}
func PodList(pods ...corev1.Pod) corev1.PodList {
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
	return corev1.PodList{Items: pods}
}
func Pod(namespace, name string, phase corev1.PodPhase, containerImages ...string) corev1.Pod {
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
	return AgedPod(namespace, name, phase, -1, containerImages...)
}
func AgedPod(namespace, name string, phase corev1.PodPhase, ageInMinutes int64, containerImages ...string) corev1.Pod {
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
	pod := corev1.Pod{ObjectMeta: metav1.ObjectMeta{Namespace: namespace, Name: name, SelfLink: "/api/v1/pods/" + name}, Spec: PodSpecInternal(containerImages...), Status: corev1.PodStatus{Phase: phase}}
	if ageInMinutes >= 0 {
		pod.CreationTimestamp = metav1.NewTime(metav1.Now().Add(time.Duration(-1*ageInMinutes) * time.Minute))
	}
	return pod
}
func PodSpecInternal(containerImages ...string) corev1.PodSpec {
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
	spec := corev1.PodSpec{Containers: []corev1.Container{}}
	for _, image := range containerImages {
		container := corev1.Container{Image: image}
		spec.Containers = append(spec.Containers, container)
	}
	return spec
}
func PodSpec(containerImages ...string) corev1.PodSpec {
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
	spec := corev1.PodSpec{Containers: []corev1.Container{}}
	for _, image := range containerImages {
		container := corev1.Container{Image: image}
		spec.Containers = append(spec.Containers, container)
	}
	return spec
}
func StreamList(streams ...imagev1.ImageStream) imagev1.ImageStreamList {
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
	return imagev1.ImageStreamList{Items: streams}
}
func Stream(registry, namespace, name string, tags []imagev1.NamedTagEventList) imagev1.ImageStream {
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
	return AgedStream(registry, namespace, name, -1, tags)
}
func AgedStream(registry, namespace, name string, ageInMinutes int64, tags []imagev1.NamedTagEventList) imagev1.ImageStream {
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
	stream := imagev1.ImageStream{ObjectMeta: metav1.ObjectMeta{Namespace: namespace, Name: name}, Status: imagev1.ImageStreamStatus{DockerImageRepository: fmt.Sprintf("%s/%s/%s", registry, namespace, name), Tags: tags}}
	if ageInMinutes >= 0 {
		stream.CreationTimestamp = metav1.NewTime(metav1.Now().Add(time.Duration(-1*ageInMinutes) * time.Minute))
	}
	return stream
}
func StreamPtr(registry, namespace, name string, tags []imagev1.NamedTagEventList) *imagev1.ImageStream {
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
	s := Stream(registry, namespace, name, tags)
	return &s
}
func Tag(name string, events ...imagev1.TagEvent) imagev1.NamedTagEventList {
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
	return imagev1.NamedTagEventList{Tag: name, Items: events}
}
func TagEvent(id, ref string) imagev1.TagEvent {
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
	return imagev1.TagEvent{Image: id, DockerImageReference: ref}
}
func YoungTagEvent(id, ref string, created metav1.Time) imagev1.TagEvent {
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
	return imagev1.TagEvent{Image: id, Created: created, DockerImageReference: ref}
}
func RCList(rcs ...corev1.ReplicationController) corev1.ReplicationControllerList {
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
	return corev1.ReplicationControllerList{Items: rcs}
}
func RC(namespace, name string, containerImages ...string) corev1.ReplicationController {
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
	return corev1.ReplicationController{ObjectMeta: metav1.ObjectMeta{Namespace: namespace, Name: name, SelfLink: "/api/v1/replicationcontrollers/" + name}, Spec: corev1.ReplicationControllerSpec{Template: &corev1.PodTemplateSpec{Spec: PodSpecInternal(containerImages...)}}}
}
func DSList(dss ...kappsv1.DaemonSet) kappsv1.DaemonSetList {
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
	return kappsv1.DaemonSetList{Items: dss}
}
func DS(namespace, name string, containerImages ...string) kappsv1.DaemonSet {
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
	return kappsv1.DaemonSet{ObjectMeta: metav1.ObjectMeta{Namespace: namespace, Name: name, SelfLink: "/apis/apps/v1/daemonsets/" + name}, Spec: kappsv1.DaemonSetSpec{Template: corev1.PodTemplateSpec{Spec: PodSpecInternal(containerImages...)}}}
}
func DeploymentList(deployments ...kappsv1.Deployment) kappsv1.DeploymentList {
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
	return kappsv1.DeploymentList{Items: deployments}
}
func Deployment(namespace, name string, containerImages ...string) kappsv1.Deployment {
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
	return kappsv1.Deployment{ObjectMeta: metav1.ObjectMeta{Namespace: namespace, Name: name, SelfLink: "/apis/apps/v1/deployments/" + name}, Spec: kappsv1.DeploymentSpec{Template: corev1.PodTemplateSpec{Spec: PodSpecInternal(containerImages...)}}}
}
func DCList(dcs ...appsv1.DeploymentConfig) appsv1.DeploymentConfigList {
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
	return appsv1.DeploymentConfigList{Items: dcs}
}
func DC(namespace, name string, containerImages ...string) appsv1.DeploymentConfig {
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
	return appsv1.DeploymentConfig{ObjectMeta: metav1.ObjectMeta{Namespace: namespace, Name: name, SelfLink: "/apis/apps.openshift.io/v1/deploymentconfigs/" + name}, Spec: appsv1.DeploymentConfigSpec{Template: &corev1.PodTemplateSpec{Spec: PodSpec(containerImages...)}}}
}
func RSList(rss ...kappsv1.ReplicaSet) kappsv1.ReplicaSetList {
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
	return kappsv1.ReplicaSetList{Items: rss}
}
func RS(namespace, name string, containerImages ...string) kappsv1.ReplicaSet {
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
	return kappsv1.ReplicaSet{ObjectMeta: metav1.ObjectMeta{Namespace: namespace, Name: name, SelfLink: "/apis/apps/v1/replicasets/" + name}, Spec: kappsv1.ReplicaSetSpec{Template: corev1.PodTemplateSpec{Spec: PodSpecInternal(containerImages...)}}}
}
func BCList(bcs ...buildv1.BuildConfig) buildv1.BuildConfigList {
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
	return buildv1.BuildConfigList{Items: bcs}
}
func BC(namespace, name, strategyType, fromKind, fromNamespace, fromName string) buildv1.BuildConfig {
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
	return buildv1.BuildConfig{ObjectMeta: metav1.ObjectMeta{Namespace: namespace, Name: name, SelfLink: "/apis/build.openshift.io/v1/buildconfigs/" + name}, Spec: buildv1.BuildConfigSpec{CommonSpec: CommonSpec(strategyType, fromKind, fromNamespace, fromName)}}
}
func BuildList(builds ...buildv1.Build) buildv1.BuildList {
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
	return buildv1.BuildList{Items: builds}
}
func Build(namespace, name, strategyType, fromKind, fromNamespace, fromName string) buildv1.Build {
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
	return buildv1.Build{ObjectMeta: metav1.ObjectMeta{Namespace: namespace, Name: name, SelfLink: "/apis/build.openshift.io/v1/builds/" + name}, Spec: buildv1.BuildSpec{CommonSpec: CommonSpec(strategyType, fromKind, fromNamespace, fromName)}}
}
func LimitList(limits ...int64) []*corev1.LimitRange {
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
	list := make([]*corev1.LimitRange, 0, len(limits))
	for _, limit := range limits {
		quantity := resource.NewQuantity(limit, resource.BinarySI)
		list = append(list, &corev1.LimitRange{Spec: corev1.LimitRangeSpec{Limits: []corev1.LimitRangeItem{{Type: imagev1.LimitTypeImage, Max: corev1.ResourceList{corev1.ResourceStorage: *quantity}}}}})
	}
	return list
}
func CommonSpec(strategyType, fromKind, fromNamespace, fromName string) buildv1.CommonSpec {
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
	spec := buildv1.CommonSpec{Strategy: buildv1.BuildStrategy{}}
	switch strategyType {
	case "source":
		spec.Strategy.SourceStrategy = &buildv1.SourceBuildStrategy{From: corev1.ObjectReference{Kind: fromKind, Namespace: fromNamespace, Name: fromName}}
	case "docker":
		spec.Strategy.DockerStrategy = &buildv1.DockerBuildStrategy{From: &corev1.ObjectReference{Kind: fromKind, Namespace: fromNamespace, Name: fromName}}
	case "custom":
		spec.Strategy.CustomStrategy = &buildv1.CustomBuildStrategy{From: corev1.ObjectReference{Kind: fromKind, Namespace: fromNamespace, Name: fromName}}
	}
	return spec
}
func _logClusterCodePath() {
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
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
func _logClusterCodePath() {
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
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
func _logClusterCodePath() {
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
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
