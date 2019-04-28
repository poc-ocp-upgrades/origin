package test

import (
	"testing"
	"github.com/openshift/origin/pkg/build/buildscheme"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/admission"
	"k8s.io/client-go/scale/scheme"
	buildv1 "github.com/openshift/api/build/v1"
	buildutil "github.com/openshift/origin/pkg/build/util"
)

type TestPod corev1.Pod

func Pod() *TestPod {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return (*TestPod)(&corev1.Pod{})
}
func (p *TestPod) WithAnnotation(name, value string) *TestPod {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if p.Annotations == nil {
		p.Annotations = map[string]string{}
	}
	p.Annotations[name] = value
	return p
}
func (p *TestPod) WithEnvVar(name, value string) *TestPod {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(p.Spec.InitContainers) == 0 {
		p.Spec.InitContainers = append(p.Spec.InitContainers, corev1.Container{})
	}
	if len(p.Spec.Containers) == 0 {
		p.Spec.Containers = append(p.Spec.Containers, corev1.Container{})
	}
	p.Spec.InitContainers[0].Env = append(p.Spec.InitContainers[0].Env, corev1.EnvVar{Name: name, Value: value})
	p.Spec.Containers[0].Env = append(p.Spec.Containers[0].Env, corev1.EnvVar{Name: name, Value: value})
	return p
}
func (p *TestPod) WithBuild(t *testing.T, build *buildv1.Build) *TestPod {
	_logClusterCodePath()
	defer _logClusterCodePath()
	encodedBuild, err := runtime.Encode(buildscheme.Encoder, build)
	if err != nil {
		t.Fatalf("%v", err)
	}
	return p.WithAnnotation(buildutil.BuildAnnotation, build.Name).WithEnvVar("BUILD", string(encodedBuild))
}
func (p *TestPod) InitEnvValue(name string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(p.Spec.InitContainers) == 0 {
		return ""
	}
	for _, ev := range p.Spec.InitContainers[0].Env {
		if ev.Name == name {
			return ev.Value
		}
	}
	return ""
}
func (p *TestPod) EnvValue(name string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(p.Spec.Containers) == 0 {
		return ""
	}
	for _, ev := range p.Spec.Containers[0].Env {
		if ev.Name == name {
			return ev.Value
		}
	}
	return ""
}
func (p *TestPod) GetBuild(t *testing.T) *buildv1.Build {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := runtime.Decode(buildscheme.Decoder, []byte(p.EnvValue("BUILD")))
	if err != nil {
		t.Fatalf("Could not decode build: %v", err)
	}
	build, ok := obj.(*buildv1.Build)
	if !ok {
		t.Fatalf("Not a build object: %#v", obj)
	}
	return build
}
func (p *TestPod) ToAttributes() admission.Attributes {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return admission.NewAttributesRecord((*corev1.Pod)(p), nil, scheme.Kind("Pod").WithVersion("version"), "default", "TestPod", corev1.Resource("pods").WithVersion("version"), "", admission.Create, false, nil)
}
func (p *TestPod) AsPod() *corev1.Pod {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return (*corev1.Pod)(p)
}
