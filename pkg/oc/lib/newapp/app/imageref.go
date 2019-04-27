package app

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"github.com/moby/buildkit/frontend/dockerfile/parser"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kvalidation "k8s.io/apimachinery/pkg/util/validation"
	appsv1 "github.com/openshift/api/apps/v1"
	buildv1 "github.com/openshift/api/build/v1"
	dockerv10 "github.com/openshift/api/image/docker10"
	imagev1 "github.com/openshift/api/image/v1"
	"github.com/openshift/origin/pkg/api/apihelpers"
	imageapi "github.com/openshift/origin/pkg/image/apis/image"
	imageutil "github.com/openshift/origin/pkg/image/util"
	"github.com/openshift/origin/pkg/util/docker/dockerfile"
	"github.com/openshift/origin/pkg/util/portutils"
)

type ImageRefGenerator interface {
	FromName(name string) (*ImageRef, error)
	FromNameAndPorts(name string, ports []string) (*ImageRef, error)
	FromStream(repo *imagev1.ImageStream, tag string) (*ImageRef, error)
	FromDockerfile(name string, dir string, context string) (*ImageRef, error)
}
type SecretAccessor interface {
	Token() (string, error)
	CACert() (string, error)
}
type imageRefGenerator struct{}

func NewImageRefGenerator() ImageRefGenerator {
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
	return &imageRefGenerator{}
}
func (g *imageRefGenerator) FromName(name string) (*ImageRef, error) {
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
	ref, err := imageapi.ParseDockerImageReference(name)
	if err != nil {
		return nil, err
	}
	return &ImageRef{Reference: ref, Info: &dockerv10.DockerImage{Config: &imageapi.DockerConfig{}}}, nil
}
func (g *imageRefGenerator) FromNameAndPorts(name string, ports []string) (*ImageRef, error) {
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
	present := struct{}{}
	imageRef, err := g.FromName(name)
	if err != nil {
		return nil, err
	}
	exposedPorts := map[string]struct{}{}
	for _, p := range ports {
		exposedPorts[p] = present
	}
	imageRef.Info = &dockerv10.DockerImage{Config: &imageapi.DockerConfig{ExposedPorts: exposedPorts}}
	return imageRef, nil
}
func (g *imageRefGenerator) FromDockerfile(name string, dir string, context string) (*ImageRef, error) {
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
	file, err := os.Open(filepath.Join(dir, context, "Dockerfile"))
	if err != nil {
		return nil, err
	}
	defer file.Close()
	node, err := parser.Parse(file)
	if err != nil {
		return nil, err
	}
	ports := dockerfile.LastExposedPorts(node.AST)
	return g.FromNameAndPorts(name, ports)
}
func (g *imageRefGenerator) FromStream(stream *imagev1.ImageStream, tag string) (*ImageRef, error) {
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
	imageRef := &ImageRef{Stream: stream}
	if tagged := imageutil.LatestTaggedImage(stream, tag); tagged != nil {
		if ref, err := imageapi.ParseDockerImageReference(tagged.DockerImageReference); err == nil {
			imageRef.ResolvedReference = &ref
			imageRef.Reference = ref
		}
	}
	if pullSpec := stream.Status.DockerImageRepository; len(pullSpec) != 0 {
		ref, err := imageapi.ParseDockerImageReference(pullSpec)
		if err != nil {
			return nil, err
		}
		imageRef.Reference = ref
	}
	switch {
	case len(tag) > 0:
		imageRef.Reference.Tag = tag
	case len(tag) == 0 && len(imageRef.Reference.Tag) == 0:
		imageRef.Reference.Tag = imageapi.DefaultImageTag
	}
	return imageRef, nil
}

type ImageRef struct {
	Reference		imageapi.DockerImageReference
	ResolvedReference	*imageapi.DockerImageReference
	AsResolvedImage		bool
	AsImageStream		bool
	OutputImage		bool
	Insecure		bool
	HasEmptyDir		bool
	TagDirectly		bool
	Tag			string
	InternalDefaultTag	string
	Env			Environment
	ObjectName		string
	ContainerFn		func(*corev1.Container)
	Stream			*imagev1.ImageStream
	Info			*dockerv10.DockerImage
}

func (r *ImageRef) Exists() bool {
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
	return r.Stream != nil
}
func (r *ImageRef) ObjectReference() corev1.ObjectReference {
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
	switch {
	case r.Stream != nil:
		return corev1.ObjectReference{Kind: "ImageStreamTag", Name: imageapi.JoinImageStreamTag(r.Stream.Name, r.Reference.Tag), Namespace: r.Stream.Namespace}
	case r.AsImageStream:
		name, _ := r.SuggestName()
		return corev1.ObjectReference{Kind: "ImageStreamTag", Name: imageapi.JoinImageStreamTag(name, r.InternalTag())}
	default:
		return corev1.ObjectReference{Kind: "DockerImage", Name: r.PullSpec()}
	}
}
func (r *ImageRef) InternalTag() string {
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
	tag := r.Tag
	if len(tag) == 0 {
		tag = r.Reference.Tag
	}
	if len(tag) == 0 {
		tag = r.InternalDefaultTag
	}
	if len(tag) == 0 {
		tag = imageapi.DefaultImageTag
	}
	return tag
}
func (r *ImageRef) PullSpec() string {
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
	if r.AsResolvedImage && r.ResolvedReference != nil {
		return r.ResolvedReference.Exact()
	}
	return r.Reference.Exact()
}
func (r *ImageRef) RepoName() string {
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
	name := r.Reference.Namespace
	if len(name) > 0 {
		name += "/"
	}
	name += r.Reference.Name
	return name
}
func (r *ImageRef) SuggestName() (string, bool) {
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
	if r == nil {
		return "", false
	}
	if len(r.ObjectName) > 0 {
		return r.ObjectName, true
	}
	if r.Stream != nil {
		return r.Stream.Name, true
	}
	if len(r.Reference.Name) > 0 {
		return r.Reference.Name, true
	}
	return "", false
}
func (r *ImageRef) SuggestNamespace() string {
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
	if r == nil {
		return ""
	}
	if len(r.ObjectName) > 0 {
		return ""
	}
	if r.Stream != nil {
		return r.Stream.Namespace
	}
	return ""
}
func (r *ImageRef) BuildOutput() (*buildv1.BuildOutput, error) {
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
	if r == nil {
		return &buildv1.BuildOutput{}, nil
	}
	if !r.AsImageStream {
		return &buildv1.BuildOutput{To: &corev1.ObjectReference{Kind: "DockerImage", Name: r.Reference.String()}}, nil
	}
	imageRepo, err := r.ImageStream()
	if err != nil {
		return nil, err
	}
	return &buildv1.BuildOutput{To: &corev1.ObjectReference{Kind: "ImageStreamTag", Name: imageapi.JoinImageStreamTag(imageRepo.Name, r.Reference.Tag)}}, nil
}
func (r *ImageRef) BuildTriggers() []buildv1.BuildTriggerPolicy {
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
	if r.Stream == nil && !r.AsImageStream {
		return nil
	}
	return []buildv1.BuildTriggerPolicy{{Type: buildv1.ImageChangeBuildTriggerType, ImageChange: &buildv1.ImageChangeTrigger{}}}
}
func (r *ImageRef) ImageStream() (*imagev1.ImageStream, error) {
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
	if r.Stream != nil {
		return r.Stream, nil
	}
	name, ok := r.SuggestName()
	if !ok {
		return nil, fmt.Errorf("unable to suggest an ImageStream name for %q", r.Reference.String())
	}
	stream := &imagev1.ImageStream{TypeMeta: metav1.TypeMeta{APIVersion: imagev1.SchemeGroupVersion.String(), Kind: "ImageStream"}, ObjectMeta: metav1.ObjectMeta{Name: name}}
	if r.OutputImage {
		return stream, nil
	}
	if !r.TagDirectly {
		stream.Spec.DockerImageRepository = r.Reference.AsRepository().String()
		if r.Insecure {
			stream.ObjectMeta.Annotations = map[string]string{imageapi.InsecureRepositoryAnnotation: "true"}
		}
		return stream, nil
	}
	if stream.Spec.Tags == nil {
		stream.Spec.Tags = []imagev1.TagReference{}
	}
	stream.Spec.Tags = append(stream.Spec.Tags, imagev1.TagReference{Name: r.InternalTag(), Annotations: map[string]string{"openshift.io/imported-from": r.Reference.Exact()}, From: &corev1.ObjectReference{Kind: "DockerImage", Name: r.PullSpec()}, ImportPolicy: imagev1.TagImportPolicy{Insecure: r.Insecure}})
	return stream, nil
}
func (r *ImageRef) ImageStreamTag() (*imagev1.ImageStreamTag, error) {
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
	name, ok := r.SuggestName()
	if !ok {
		return nil, fmt.Errorf("unable to suggest an ImageStream name for %q", r.Reference.String())
	}
	istname := imageapi.JoinImageStreamTag(name, r.Reference.Tag)
	ist := &imagev1.ImageStreamTag{TypeMeta: metav1.TypeMeta{APIVersion: imagev1.SchemeGroupVersion.String(), Kind: "ImageStreamTag"}, ObjectMeta: metav1.ObjectMeta{Name: istname, Namespace: r.SuggestNamespace(), Annotations: map[string]string{"openshift.io/imported-from": r.Reference.Exact()}}, Tag: &imagev1.TagReference{Name: r.InternalTag(), From: &corev1.ObjectReference{Kind: "DockerImage", Name: r.PullSpec()}, ImportPolicy: imagev1.TagImportPolicy{Insecure: r.Insecure}}}
	return ist, nil
}
func (r *ImageRef) DeployableContainer() (container *corev1.Container, triggers []appsv1.DeploymentTriggerPolicy, err error) {
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
	name, ok := r.SuggestName()
	if !ok {
		return nil, nil, fmt.Errorf("unable to suggest a container name for the image %q", r.Reference.String())
	}
	if r.AsImageStream {
		triggers = []appsv1.DeploymentTriggerPolicy{{Type: appsv1.DeploymentTriggerOnImageChange, ImageChangeParams: &appsv1.DeploymentTriggerImageChangeParams{Automatic: true, ContainerNames: []string{name}, From: r.ObjectReference()}}}
	}
	container = &corev1.Container{Name: name, Image: r.PullSpec(), Env: r.Env.List()}
	if r.ContainerFn != nil {
		r.ContainerFn(container)
		return container, triggers, nil
	}
	if r.Info != nil && r.Info.Config != nil {
		ports := []string{}
		for exposed := range r.Info.Config.ExposedPorts {
			ports = append(ports, strings.Split(exposed, " ")...)
		}
		dockerPorts, _ := portutils.FilterPortAndProtocolArray(ports)
		for _, dp := range dockerPorts {
			intPort, _ := strconv.Atoi(dp.Port())
			container.Ports = append(container.Ports, corev1.ContainerPort{ContainerPort: int32(intPort), Protocol: corev1.Protocol(strings.ToUpper(dp.Proto()))})
		}
		maxDigits := len(fmt.Sprintf("%d", len(r.Info.Config.Volumes)))
		baseName := apihelpers.GetName(container.Name, volumeNameInfix, kvalidation.LabelValueMaxLength-maxDigits-1)
		i := 1
		for volume := range r.Info.Config.Volumes {
			r.HasEmptyDir = true
			container.VolumeMounts = append(container.VolumeMounts, corev1.VolumeMount{Name: fmt.Sprintf("%s-%d", baseName, i), ReadOnly: false, MountPath: volume})
			i++
		}
	}
	return container, triggers, nil
}
func (r *ImageRef) InstallablePod(generatorInput GeneratorInput, secretAccessor SecretAccessor, serviceAccountName string) (*corev1.Pod, *corev1.Secret, error) {
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
	name, ok := r.SuggestName()
	if !ok {
		return nil, nil, fmt.Errorf("can't suggest a name for the provided image %q", r.Reference.Exact())
	}
	meta := metav1.ObjectMeta{Name: fmt.Sprintf("%s-install", name)}
	container, _, err := r.DeployableContainer()
	if err != nil {
		return nil, nil, fmt.Errorf("can't generate an installable container: %v", err)
	}
	container.Name = "install"
	namespaceEnv := corev1.EnvVar{Name: "POD_NAMESPACE", ValueFrom: &corev1.EnvVarSource{FieldRef: &corev1.ObjectFieldSelector{APIVersion: "v1", FieldPath: "metadata.namespace"}}}
	container.Env = append([]corev1.EnvVar{namespaceEnv}, container.Env...)
	deadline := int64(60 * 60 * 4)
	pod := &corev1.Pod{TypeMeta: metav1.TypeMeta{APIVersion: metav1.SchemeGroupVersion.String(), Kind: "Pod"}, ObjectMeta: meta, Spec: corev1.PodSpec{RestartPolicy: corev1.RestartPolicyNever, ActiveDeadlineSeconds: &deadline}}
	var secret *corev1.Secret
	if token := generatorInput.Token; token != nil {
		if token.ServiceAccount {
			pod.Spec.ServiceAccountName = serviceAccountName
		}
		if token.Env != nil {
			containerToken, err := secretAccessor.Token()
			if err != nil {
				return nil, nil, err
			}
			container.Env = append(container.Env, corev1.EnvVar{Name: *token.Env, Value: containerToken})
		}
		if token.File != nil {
			containerToken, err := secretAccessor.Token()
			if err != nil {
				return nil, nil, err
			}
			crt, err := secretAccessor.CACert()
			if err != nil {
				return nil, nil, err
			}
			secret = &corev1.Secret{TypeMeta: metav1.TypeMeta{APIVersion: metav1.SchemeGroupVersion.String(), Kind: "Secret"}, ObjectMeta: meta, Type: "kubernetes.io/token", Data: map[string][]byte{corev1.ServiceAccountTokenKey: []byte(containerToken)}}
			if len(crt) > 0 {
				secret.Data[corev1.ServiceAccountRootCAKey] = []byte(crt)
			}
			pod.Spec.Volumes = append(pod.Spec.Volumes, corev1.Volume{Name: "generate-token", VolumeSource: corev1.VolumeSource{Secret: &corev1.SecretVolumeSource{SecretName: meta.Name}}})
			container.VolumeMounts = append(container.VolumeMounts, corev1.VolumeMount{Name: "generate-token", MountPath: *token.File})
		}
	}
	pod.Spec.Containers = []corev1.Container{*container}
	return pod, secret, nil
}
