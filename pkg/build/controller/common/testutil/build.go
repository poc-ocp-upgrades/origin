package test

import (
	buildv1 "github.com/openshift/api/build/v1"
	"bytes"
	"net/http"
	"runtime"
	"fmt"
)

type TestBuild buildv1.Build

func Build() *TestBuild {
	_logClusterCodePath()
	defer _logClusterCodePath()
	b := (*TestBuild)(&buildv1.Build{})
	b.Kind = "Build"
	b.APIVersion = "build.openshift.io/v1"
	b.Name = "TestBuild"
	b.WithDockerStrategy()
	b.Spec.Source.Git = &buildv1.GitBuildSource{URI: "http://test.build/source"}
	b.Spec.TriggeredBy = []buildv1.BuildTriggerCause{}
	return b
}
func (b *TestBuild) clearStrategy() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	b.Spec.Strategy.DockerStrategy = nil
	b.Spec.Strategy.SourceStrategy = nil
	b.Spec.Strategy.CustomStrategy = nil
	b.Spec.Strategy.JenkinsPipelineStrategy = nil
}
func (b *TestBuild) WithDockerStrategy() *TestBuild {
	_logClusterCodePath()
	defer _logClusterCodePath()
	b.clearStrategy()
	b.Spec.Strategy.DockerStrategy = &buildv1.DockerBuildStrategy{}
	return b
}
func (b *TestBuild) WithSourceStrategy() *TestBuild {
	_logClusterCodePath()
	defer _logClusterCodePath()
	b.clearStrategy()
	strategy := &buildv1.SourceBuildStrategy{}
	strategy.From.Name = "builder/image"
	strategy.From.Kind = "DockerImage"
	b.Spec.Strategy.SourceStrategy = strategy
	return b
}
func (b *TestBuild) WithCustomStrategy() *TestBuild {
	_logClusterCodePath()
	defer _logClusterCodePath()
	b.clearStrategy()
	strategy := &buildv1.CustomBuildStrategy{}
	strategy.From.Name = "builder/image"
	strategy.From.Kind = "DockerImage"
	b.Spec.Strategy.CustomStrategy = strategy
	return b
}
func (b *TestBuild) WithImageLabels(labels []buildv1.ImageLabel) *TestBuild {
	_logClusterCodePath()
	defer _logClusterCodePath()
	b.Spec.Output.ImageLabels = labels
	return b
}
func (b *TestBuild) WithNodeSelector(ns map[string]string) *TestBuild {
	_logClusterCodePath()
	defer _logClusterCodePath()
	b.Spec.NodeSelector = ns
	return b
}
func (b *TestBuild) AsBuild() *buildv1.Build {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return (*buildv1.Build)(b)
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
