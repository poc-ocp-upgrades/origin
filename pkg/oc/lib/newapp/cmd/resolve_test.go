package cmd

import (
	"testing"
	"github.com/openshift/origin/pkg/oc/lib/newapp"
	"github.com/openshift/origin/pkg/oc/lib/newapp/app"
)

func TestResolveJenkinsfileAndDockerfile(t *testing.T) {
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
	dockerfile, _ := app.NewDockerfile("FROM centos\n")
	i := app.SourceRepositoryInfo{Dockerfile: dockerfile, Jenkinsfile: true}
	repo := app.SourceRepository{}
	repo.SetInfo(&i)
	repositories := app.SourceRepositories{&repo}
	resolvers := Resolvers{}
	componentrefs, err := AddMissingComponentsToRefBuilder(&app.ReferenceBuilder{}, repositories, resolvers.DockerfileResolver(), resolvers.SourceResolver(), resolvers.PipelineResolver(), &GenerationInputs{})
	checkResolveResult(t, componentrefs, err, newapp.StrategyPipeline)
}
func TestResolveJenkinsfileAndSource(t *testing.T) {
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
	i := app.SourceRepositoryInfo{Jenkinsfile: true, Types: []app.SourceLanguageType{{Platform: "foo"}}}
	repo := app.SourceRepository{}
	repo.SetInfo(&i)
	repositories := app.SourceRepositories{&repo}
	resolvers := Resolvers{}
	componentrefs, err := AddMissingComponentsToRefBuilder(&app.ReferenceBuilder{}, repositories, resolvers.DockerfileResolver(), resolvers.SourceResolver(), resolvers.PipelineResolver(), &GenerationInputs{})
	checkResolveResult(t, componentrefs, err, newapp.StrategyPipeline)
}
func TestResolveDockerfileAndSource(t *testing.T) {
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
	dockerfile, _ := app.NewDockerfile("FROM centos\n")
	i := app.SourceRepositoryInfo{Dockerfile: dockerfile, Types: []app.SourceLanguageType{{Platform: "foo"}}}
	repo := app.SourceRepository{}
	repo.SetInfo(&i)
	repositories := app.SourceRepositories{&repo}
	resolvers := Resolvers{}
	componentrefs, err := AddMissingComponentsToRefBuilder(&app.ReferenceBuilder{}, repositories, resolvers.DockerfileResolver(), resolvers.SourceResolver(), resolvers.PipelineResolver(), &GenerationInputs{})
	checkResolveResult(t, componentrefs, err, newapp.StrategyDocker)
}
func TestBinaryContentFlagGeneratesDummySource(t *testing.T) {
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
	component := app.ComponentInput{Value: "foo", From: "--binary", Argument: "--binary"}
	refs := app.ComponentReferences{&component}
	input := GenerationInputs{BinaryBuild: true, ExpectToBuild: true}
	err := EnsureHasSource(refs, nil, &input)
	if err != nil {
		t.Fatal(err)
	}
	if component.Uses == nil {
		t.Fatal("expected source repository to be created")
	}
	if !component.Uses.InUse() {
		t.Fatal("expected source repository to be in use")
	}
}
func checkResolveResult(t *testing.T, componentrefs app.ComponentReferences, err error, strategy newapp.Strategy) {
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
	if err != nil {
		t.Fatal(err)
	}
	if len(componentrefs) != 1 {
		t.Fatal("expected len(componentrefs) == 1")
	}
	if componentrefs[0].Input().Uses == nil {
		t.Fatal("expected non-nil componentrefs[0].Input().Uses")
	}
	if componentrefs[0].Input().Uses.GetStrategy() != strategy {
		t.Fatalf("expected componentrefs[0].Input().Uses.GetStrategy() == %s", strategy)
	}
}
