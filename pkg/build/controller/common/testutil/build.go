package test

import (
	goformat "fmt"
	buildv1 "github.com/openshift/api/build/v1"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

type TestBuild buildv1.Build

func Build() *TestBuild {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
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
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	b.Spec.Strategy.DockerStrategy = nil
	b.Spec.Strategy.SourceStrategy = nil
	b.Spec.Strategy.CustomStrategy = nil
	b.Spec.Strategy.JenkinsPipelineStrategy = nil
}
func (b *TestBuild) WithDockerStrategy() *TestBuild {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	b.clearStrategy()
	b.Spec.Strategy.DockerStrategy = &buildv1.DockerBuildStrategy{}
	return b
}
func (b *TestBuild) WithSourceStrategy() *TestBuild {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	b.clearStrategy()
	strategy := &buildv1.SourceBuildStrategy{}
	strategy.From.Name = "builder/image"
	strategy.From.Kind = "DockerImage"
	b.Spec.Strategy.SourceStrategy = strategy
	return b
}
func (b *TestBuild) WithCustomStrategy() *TestBuild {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	b.clearStrategy()
	strategy := &buildv1.CustomBuildStrategy{}
	strategy.From.Name = "builder/image"
	strategy.From.Kind = "DockerImage"
	b.Spec.Strategy.CustomStrategy = strategy
	return b
}
func (b *TestBuild) WithImageLabels(labels []buildv1.ImageLabel) *TestBuild {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	b.Spec.Output.ImageLabels = labels
	return b
}
func (b *TestBuild) WithNodeSelector(ns map[string]string) *TestBuild {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	b.Spec.NodeSelector = ns
	return b
}
func (b *TestBuild) AsBuild() *buildv1.Build {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return (*buildv1.Build)(b)
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
