package app

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
	"k8s.io/klog"
	kappsv1 "k8s.io/api/apps/v1"
	kappsv1beta2 "k8s.io/api/apps/v1beta2"
	corev1 "k8s.io/api/core/v1"
	extensionsv1beta1 "k8s.io/api/extensions/v1beta1"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/intstr"
	kuval "k8s.io/apimachinery/pkg/util/validation"
	kapi "k8s.io/kubernetes/pkg/apis/core"
	"k8s.io/kubernetes/pkg/apis/core/validation"
	appsv1 "github.com/openshift/api/apps/v1"
	"github.com/openshift/api/build"
	buildv1 "github.com/openshift/api/build/v1"
	"github.com/openshift/api/image"
	imagev1 "github.com/openshift/api/image/v1"
	imagev1typedclient "github.com/openshift/client-go/image/clientset/versioned/typed/image/v1"
	"github.com/openshift/origin/pkg/api/legacy"
	imageapi "github.com/openshift/origin/pkg/image/apis/image"
	"github.com/openshift/origin/pkg/oc/lib/newapp"
	routeapi "github.com/openshift/origin/pkg/route/apis/route"
	"github.com/openshift/origin/pkg/util/docker/dockerfile"
)

type PipelineBuilder interface {
	To(string) PipelineBuilder
	NewBuildPipeline(string, *ImageRef, *SourceRepository, bool) (*Pipeline, error)
	NewImagePipeline(string, *ImageRef) (*Pipeline, error)
}

func NewPipelineBuilder(name string, environment Environment, dockerStrategyOptions *buildv1.DockerStrategyOptions, outputDocker bool) PipelineBuilder {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &pipelineBuilder{nameGenerator: NewUniqueNameGenerator(name), environment: environment, outputDocker: outputDocker, dockerStrategyOptions: dockerStrategyOptions}
}

type pipelineBuilder struct {
	nameGenerator		UniqueNameGenerator
	environment		Environment
	outputDocker		bool
	to			string
	dockerStrategyOptions	*buildv1.DockerStrategyOptions
}

func (pb *pipelineBuilder) To(name string) PipelineBuilder {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pb.to = name
	return pb
}
func (pb *pipelineBuilder) NewBuildPipeline(from string, input *ImageRef, sourceRepository *SourceRepository, binary bool) (*Pipeline, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	strategy, source, err := StrategyAndSourceForRepository(sourceRepository, input)
	if err != nil {
		return nil, fmt.Errorf("can't build %q: %v", from, err)
	}
	var name string
	output := &ImageRef{OutputImage: true, AsImageStream: !pb.outputDocker}
	if len(pb.to) > 0 {
		outputImageRef, err := imageapi.ParseDockerImageReference(pb.to)
		if err != nil {
			return nil, err
		}
		output.Reference = outputImageRef
		name, err = pb.nameGenerator.Generate(NameSuggestions{source, output, input})
		if err != nil {
			return nil, err
		}
	} else {
		name, err = pb.nameGenerator.Generate(NameSuggestions{source, input})
		if err != nil {
			return nil, err
		}
		output.Reference = imageapi.DockerImageReference{Name: name, Tag: imageapi.DefaultImageTag}
	}
	source.Name = name
	if sourceRepository.GetStrategy() == newapp.StrategyDocker && sourceRepository.Info() != nil {
		node := sourceRepository.Info().Dockerfile.AST()
		ports := dockerfile.LastExposedPorts(node)
		if len(ports) > 0 {
			if input.Info == nil {
				input.Info = &imageapi.DockerImage{Config: &imageapi.DockerConfig{}}
			}
			input.Info.Config.ExposedPorts = map[string]struct{}{}
			for _, p := range ports {
				input.Info.Config.ExposedPorts[p] = struct{}{}
			}
		}
	}
	if input != nil {
		output.Info = input.Info
	}
	build := &BuildRef{Source: source, Input: input, Strategy: strategy, Output: output, Env: pb.environment, DockerStrategyOptions: pb.dockerStrategyOptions, Binary: binary}
	return &Pipeline{Name: name, From: from, InputImage: input, Image: output, Build: build}, nil
}
func (pb *pipelineBuilder) NewImagePipeline(from string, input *ImageRef) (*Pipeline, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	name, err := pb.nameGenerator.Generate(input)
	if err != nil {
		return nil, err
	}
	input.ObjectName = name
	return &Pipeline{Name: name, From: from, Image: input}, nil
}

type Pipeline struct {
	Name		string
	From		string
	InputImage	*ImageRef
	Build		*BuildRef
	Image		*ImageRef
	Deployment	*DeploymentConfigRef
	Labels		map[string]string
}

func (p *Pipeline) NeedsDeployment(env Environment, labels map[string]string, asTest bool) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if p.Deployment != nil {
		return nil
	}
	p.Deployment = &DeploymentConfigRef{Name: p.Name, Images: []*ImageRef{p.Image}, Env: env, Labels: labels, AsTest: asTest}
	return nil
}
func (p *Pipeline) Objects(accept, objectAccept Acceptor) (Objects, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	objects := Objects{}
	if p.InputImage != nil && p.InputImage.AsImageStream && accept.Accept(p.InputImage) {
		repo, err := p.InputImage.ImageStream()
		if err != nil {
			return nil, err
		}
		if objectAccept.Accept(repo) {
			objects = append(objects, repo)
		} else {
			tag, err := p.InputImage.ImageStreamTag()
			if err != nil {
				return nil, err
			}
			if objectAccept.Accept(tag) && accept.Accept(tag) {
				objects = append(objects, tag)
			}
		}
	}
	if p.Image != nil && p.Image.AsImageStream && accept.Accept(p.Image) {
		repo, err := p.Image.ImageStream()
		if err != nil {
			return nil, err
		}
		if objectAccept.Accept(repo) {
			objects = append(objects, repo)
		} else {
			tag, err := p.Image.ImageStreamTag()
			if err != nil {
				return nil, err
			}
			if objectAccept.Accept(tag) {
				objects = append(objects, tag)
			}
		}
	}
	if p.Build != nil && accept.Accept(p.Build) {
		build, err := p.Build.BuildConfig()
		if err != nil {
			return nil, err
		}
		if objectAccept.Accept(build) {
			objects = append(objects, build)
		}
		if p.Build.Source != nil && p.Build.Source.SourceImage != nil && p.Build.Source.SourceImage.AsImageStream && accept.Accept(p.Build.Source.SourceImage) {
			srcImage, err := p.Build.Source.SourceImage.ImageStream()
			if err != nil {
				return nil, err
			}
			if objectAccept.Accept(srcImage) {
				objects = append(objects, srcImage)
			}
		}
	}
	if p.Deployment != nil && accept.Accept(p.Deployment) {
		dc, err := p.Deployment.DeploymentConfig()
		if err != nil {
			return nil, err
		}
		if objectAccept.Accept(dc) {
			objects = append(objects, dc)
		}
	}
	return objects, nil
}

type PipelineGroup []*Pipeline

func (g PipelineGroup) Reduce() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var deployment *DeploymentConfigRef
	for _, p := range g {
		if p.Deployment == nil || p.Deployment == deployment {
			continue
		}
		if deployment == nil {
			deployment = p.Deployment
		} else {
			deployment.Images = append(deployment.Images, p.Deployment.Images...)
			deployment.Env = NewEnvironment(deployment.Env, p.Deployment.Env)
			p.Deployment = deployment
		}
	}
	return nil
}
func (g PipelineGroup) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := []string{}
	for _, p := range g {
		s = append(s, p.From)
	}
	return strings.Join(s, "+")
}
func MakeSimpleName(name string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	name = strings.ToLower(name)
	name = invalidServiceChars.ReplaceAllString(name, "")
	name = strings.TrimFunc(name, func(r rune) bool {
		return r == '-'
	})
	if len(name) > kuval.DNS1035LabelMaxLength {
		name = name[:kuval.DNS1035LabelMaxLength]
	}
	return name
}

var invalidServiceChars = regexp.MustCompile("[^-a-z0-9]")

func makeValidServiceName(name string) (string, string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(validation.ValidateServiceName(name, false)) == 0 {
		return name, ""
	}
	name = MakeSimpleName(name)
	if len(name) == 0 {
		return "", "service-"
	}
	return name, ""
}

type sortablePorts []corev1.ContainerPort

func (s sortablePorts) Len() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return len(s)
}
func (s sortablePorts) Swap(i, j int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s[i], s[j] = s[j], s[i]
}
func (s sortablePorts) Less(i, j int) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s[i].ContainerPort < s[j].ContainerPort
}
func portName(port int, protocol corev1.Protocol) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if protocol == "" {
		protocol = corev1.ProtocolTCP
	}
	return strings.ToLower(fmt.Sprintf("%d-%s", port, protocol))
}
func GenerateService(meta metav1.ObjectMeta, selector map[string]string) *corev1.Service {
	_logClusterCodePath()
	defer _logClusterCodePath()
	name, generateName := makeValidServiceName(meta.Name)
	svc := &corev1.Service{TypeMeta: metav1.TypeMeta{APIVersion: corev1.SchemeGroupVersion.String(), Kind: "Service"}, ObjectMeta: metav1.ObjectMeta{Name: name, GenerateName: generateName, Labels: meta.Labels}, Spec: corev1.ServiceSpec{Selector: selector}}
	return svc
}
func AllContainerPorts(containers ...corev1.Container) []corev1.ContainerPort {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var ports []corev1.ContainerPort
	for _, container := range containers {
		ports = append(ports, container.Ports...)
	}
	sort.Sort(sortablePorts(ports))
	return ports
}
func UniqueContainerToServicePorts(ports []corev1.ContainerPort) []corev1.ServicePort {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var result []corev1.ServicePort
	svcPorts := map[string]struct{}{}
	for _, p := range ports {
		name := portName(int(p.ContainerPort), p.Protocol)
		_, exists := svcPorts[name]
		if exists {
			continue
		}
		svcPorts[name] = struct{}{}
		result = append(result, corev1.ServicePort{Name: name, Port: p.ContainerPort, Protocol: p.Protocol, TargetPort: intstr.FromInt(int(p.ContainerPort))})
	}
	return result
}
func AddServices(objects Objects, firstPortOnly bool) Objects {
	_logClusterCodePath()
	defer _logClusterCodePath()
	svcs := []runtime.Object{}
	for _, o := range objects {
		switch t := o.(type) {
		case *appsv1.DeploymentConfig:
			svc := addService(t.Spec.Template.Spec.Containers, t.ObjectMeta, t.Spec.Selector, firstPortOnly)
			if svc != nil {
				svcs = append(svcs, svc)
			}
		case *kappsv1.DaemonSet:
			svc := addService(t.Spec.Template.Spec.Containers, t.ObjectMeta, t.Spec.Template.Labels, firstPortOnly)
			if svc != nil {
				svcs = append(svcs, svc)
			}
		case *extensionsv1beta1.DaemonSet:
			svc := addService(t.Spec.Template.Spec.Containers, t.ObjectMeta, t.Spec.Template.Labels, firstPortOnly)
			if svc != nil {
				svcs = append(svcs, svc)
			}
		case *kappsv1beta2.DaemonSet:
			svc := addService(t.Spec.Template.Spec.Containers, t.ObjectMeta, t.Spec.Template.Labels, firstPortOnly)
			if svc != nil {
				svcs = append(svcs, svc)
			}
		}
	}
	return append(objects, svcs...)
}
func addService(containers []corev1.Container, objectMeta metav1.ObjectMeta, selector map[string]string, firstPortOnly bool) *corev1.Service {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ports := UniqueContainerToServicePorts(AllContainerPorts(containers...))
	if len(ports) == 0 {
		return nil
	}
	if firstPortOnly {
		ports = ports[:1]
	}
	svc := GenerateService(objectMeta, selector)
	svc.Spec.Ports = ports
	return svc
}
func AddRoutes(objects Objects) Objects {
	_logClusterCodePath()
	defer _logClusterCodePath()
	routes := []runtime.Object{}
	for _, o := range objects {
		switch t := o.(type) {
		case *kapi.Service:
			routes = append(routes, &routeapi.Route{ObjectMeta: metav1.ObjectMeta{Name: t.Name, Labels: t.Labels}, Spec: routeapi.RouteSpec{To: routeapi.RouteTargetReference{Name: t.Name}}})
		}
	}
	return append(objects, routes...)
}

type acceptNew struct{}

var AcceptNew Acceptor = acceptNew{}

func (acceptNew) Accept(from interface{}) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, meta, err := objectMetaData(from)
	if err != nil {
		return false
	}
	if len(meta.GetResourceVersion()) > 0 {
		return false
	}
	return true
}

type acceptUnique struct {
	typer	runtime.ObjectTyper
	objects	map[string]struct{}
}

func (a *acceptUnique) Accept(from interface{}) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, meta, err := objectMetaData(from)
	if err != nil {
		return false
	}
	gvk, _, err := a.typer.ObjectKinds(obj)
	if err != nil {
		return false
	}
	key := fmt.Sprintf("%s/%s/%s", gvk[0].Kind, meta.GetNamespace(), meta.GetName())
	_, exists := a.objects[key]
	if exists {
		return false
	}
	a.objects[key] = struct{}{}
	return true
}
func NewAcceptUnique(typer runtime.ObjectTyper) Acceptor {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &acceptUnique{typer: typer, objects: map[string]struct{}{}}
}

type acceptNonExistentImageStream struct {
	typer		runtime.ObjectTyper
	getter		imagev1typedclient.ImageV1Interface
	namespace	string
}

func (a *acceptNonExistentImageStream) Accept(from interface{}) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, _, err := objectMetaData(from)
	if err != nil {
		return false
	}
	gvk, _, err := a.typer.ObjectKinds(obj)
	if err != nil {
		return false
	}
	gk := gvk[0].GroupKind()
	if !(image.Kind("ImageStream") == gk || legacy.Kind("ImageStream") == gk) {
		return true
	}
	is, ok := from.(*imagev1.ImageStream)
	if !ok {
		klog.V(4).Infof("type cast to image stream %#v not right for an unanticipated reason", from)
		return true
	}
	namespace := a.namespace
	if len(is.Namespace) > 0 {
		namespace = is.Namespace
	}
	imgstrm, err := a.getter.ImageStreams(namespace).Get(is.Name, metav1.GetOptions{})
	if err == nil && imgstrm != nil {
		klog.V(4).Infof("acceptor determined that imagestream %s in namespace %s exists so don't accept: %#v", is.Name, namespace, imgstrm)
		return false
	}
	return true
}
func NewAcceptNonExistentImageStream(typer runtime.ObjectTyper, getter imagev1typedclient.ImageV1Interface, namespace string) Acceptor {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &acceptNonExistentImageStream{typer: typer, getter: getter, namespace: namespace}
}

type acceptNonExistentImageStreamTag struct {
	typer		runtime.ObjectTyper
	getter		imagev1typedclient.ImageV1Interface
	namespace	string
}

func (a *acceptNonExistentImageStreamTag) Accept(from interface{}) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, _, err := objectMetaData(from)
	if err != nil {
		return false
	}
	gvk, _, err := a.typer.ObjectKinds(obj)
	if err != nil {
		return false
	}
	gk := gvk[0].GroupKind()
	if !(image.Kind("ImageStreamTag") == gk || legacy.Kind("ImageStreamTag") == gk) {
		return true
	}
	ist, ok := from.(*imagev1.ImageStreamTag)
	if !ok {
		klog.V(4).Infof("type cast to imagestreamtag %#v not right for an unanticipated reason", from)
		return true
	}
	namespace := a.namespace
	if len(ist.Namespace) > 0 {
		namespace = ist.Namespace
	}
	tag, err := a.getter.ImageStreamTags(namespace).Get(ist.Name, metav1.GetOptions{})
	if err == nil && tag != nil {
		klog.V(4).Infof("acceptor determined that imagestreamtag %s in namespace %s exists so don't accept: %#v", ist.Name, namespace, tag)
		return false
	}
	return true
}
func NewAcceptNonExistentImageStreamTag(typer runtime.ObjectTyper, getter imagev1typedclient.ImageV1Interface, namespace string) Acceptor {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &acceptNonExistentImageStreamTag{typer: typer, getter: getter, namespace: namespace}
}
func objectMetaData(raw interface{}) (runtime.Object, metav1.Object, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, ok := raw.(runtime.Object)
	if !ok {
		return nil, nil, fmt.Errorf("%#v is not a runtime.Object", raw)
	}
	meta, err := meta.Accessor(obj)
	if err != nil {
		return nil, nil, err
	}
	return obj, meta, nil
}

type acceptBuildConfigs struct{ typer runtime.ObjectTyper }

func (a *acceptBuildConfigs) Accept(from interface{}) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, _, err := objectMetaData(from)
	if err != nil {
		return false
	}
	gvk, _, err := a.typer.ObjectKinds(obj)
	if err != nil {
		return false
	}
	gk := gvk[0].GroupKind()
	return build.Kind("BuildConfig") == gk || image.Kind("ImageStream") == gk
}
func NewAcceptBuildConfigs(typer runtime.ObjectTyper) Acceptor {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &acceptBuildConfigs{typer: typer}
}

type Acceptors []Acceptor

func (aa Acceptors) Accept(from interface{}) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, a := range aa {
		if !a.Accept(from) {
			return false
		}
	}
	return true
}

type acceptAll struct{}

var AcceptAll Acceptor = acceptAll{}

func (acceptAll) Accept(_ interface{}) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return true
}

type Objects []runtime.Object
type Acceptor interface{ Accept(from interface{}) bool }
type acceptFirst struct{ handled map[interface{}]struct{} }

func NewAcceptFirst() Acceptor {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &acceptFirst{make(map[interface{}]struct{})}
}
func (s *acceptFirst) Accept(from interface{}) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if _, ok := s.handled[from]; ok {
		return false
	}
	s.handled[from] = struct{}{}
	return true
}
