package cmd

import (
	"errors"
	"fmt"
	"strings"
	kutilerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/klog"
	"github.com/openshift/library-go/pkg/git"
	"github.com/openshift/origin/pkg/oc/lib/newapp"
	"github.com/openshift/origin/pkg/oc/lib/newapp/app"
	dockerfileutil "github.com/openshift/origin/pkg/util/docker/dockerfile"
)

type Resolvers struct {
	DockerSearcher			app.Searcher
	ImageStreamSearcher		app.Searcher
	ImageStreamByAnnotationSearcher	app.Searcher
	TemplateSearcher		app.Searcher
	TemplateFileSearcher		app.Searcher
	AllowMissingImages		bool
	Detector			app.Detector
}

func (r *Resolvers) ImageSourceResolver() app.Resolver {
	_logClusterCodePath()
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
	resolver := app.PerfectMatchWeightedResolver{}
	if r.ImageStreamByAnnotationSearcher != nil {
		resolver = append(resolver, app.WeightedResolver{Searcher: r.ImageStreamByAnnotationSearcher, Weight: 0.0})
	}
	if r.ImageStreamSearcher != nil {
		resolver = append(resolver, app.WeightedResolver{Searcher: r.ImageStreamSearcher, Weight: 1.0})
	}
	if r.DockerSearcher != nil {
		resolver = append(resolver, app.WeightedResolver{Searcher: r.DockerSearcher, Weight: 2.0})
	}
	return resolver
}
func (r *Resolvers) DockerfileResolver() app.Resolver {
	_logClusterCodePath()
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
	resolver := app.PerfectMatchWeightedResolver{}
	if r.ImageStreamSearcher != nil {
		resolver = append(resolver, app.WeightedResolver{Searcher: r.ImageStreamSearcher, Weight: 0.0})
	}
	if r.DockerSearcher != nil {
		resolver = append(resolver, app.WeightedResolver{Searcher: r.DockerSearcher, Weight: 1.0})
	}
	if r.AllowMissingImages {
		resolver = append(resolver, app.WeightedResolver{Searcher: &app.MissingImageSearcher{}, Weight: 100.0})
	}
	return resolver
}
func (r *Resolvers) PipelineResolver() app.Resolver {
	_logClusterCodePath()
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
	return app.PipelineResolver{}
}
func (r *Resolvers) SourceResolver() app.Resolver {
	_logClusterCodePath()
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
	resolver := app.PerfectMatchWeightedResolver{}
	if r.ImageStreamSearcher != nil {
		resolver = append(resolver, app.WeightedResolver{Searcher: r.ImageStreamSearcher, Weight: 0.0})
	}
	if r.ImageStreamByAnnotationSearcher != nil {
		resolver = append(resolver, app.WeightedResolver{Searcher: r.ImageStreamByAnnotationSearcher, Weight: 1.0})
	}
	if r.DockerSearcher != nil {
		resolver = append(resolver, app.WeightedResolver{Searcher: r.DockerSearcher, Weight: 2.0})
	}
	return resolver
}

type ComponentInputs struct {
	SourceRepositories	[]string
	Components		[]string
	ImageStreams		[]string
	DockerImages		[]string
	Templates		[]string
	TemplateFiles		[]string
	Groups			[]string
}
type ResolvedComponents struct {
	Components	app.ComponentReferences
	Repositories	app.SourceRepositories
}

func Resolve(appConfig *AppConfig) (*ResolvedComponents, error) {
	_logClusterCodePath()
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
	r := &appConfig.Resolvers
	c := &appConfig.ComponentInputs
	g := &appConfig.GenerationInputs
	s := &appConfig.SourceRepositories
	i := &appConfig.ImageStreams
	b := &app.ReferenceBuilder{}
	if err := AddComponentInputsToRefBuilder(b, r, c, g, s, i); err != nil {
		return nil, err
	}
	components, repositories, errs := b.Result()
	if len(errs) > 0 {
		return nil, kutilerrors.NewAggregate(errs)
	}
	imageComp, repositories, err := AddImageSourceRepository(repositories, r.ImageSourceResolver(), g)
	if err != nil {
		return nil, err
	}
	c = nil
	if len(g.ContextDir) > 0 && len(repositories) > 0 {
		klog.V(5).Infof("Setting contextDir on all repositories to %v", g.ContextDir)
		for _, repo := range repositories {
			repo.SetContextDir(g.ContextDir)
		}
	}
	if g.Strategy != newapp.StrategyUnspecified && len(repositories) > 0 {
		klog.V(5).Infof("Setting build strategy on all repositories to %v", g.Strategy)
		for _, repo := range repositories {
			repo.SetStrategy(g.Strategy)
		}
	}
	if g.Strategy != newapp.StrategyUnspecified && len(repositories) == 0 && !g.BinaryBuild {
		return nil, errors.New("--strategy is specified and none of the arguments provided could be classified as a source code location")
	}
	if g.BinaryBuild && (len(repositories) > 0 || components.HasSource()) {
		return nil, errors.New("specifying binary builds and source repositories at the same time is not allowed")
	}
	componentsIncludingImageComps := components
	if imageComp != nil {
		componentsIncludingImageComps = append(components, imageComp)
	}
	if err := componentsIncludingImageComps.Resolve(); err != nil {
		return nil, err
	}
	if err := detectPartialMatches(componentsIncludingImageComps); err != nil {
		return nil, err
	}
	components, err = InferBuildTypes(components, g)
	if err != nil {
		return nil, err
	}
	if err := EnsureHasSource(components.NeedsSource(), repositories.NotUsed(), g); err != nil {
		return nil, err
	}
	sourceComponents, err := AddMissingComponentsToRefBuilder(b, repositories.NotUsed(), r.DockerfileResolver(), r.SourceResolver(), r.PipelineResolver(), g)
	if err != nil {
		return nil, err
	}
	if err := sourceComponents.Resolve(); err != nil {
		return nil, err
	}
	components = append(components, sourceComponents...)
	klog.V(4).Infof("Code [%v]", repositories)
	klog.V(4).Infof("Components [%v]", components)
	return &ResolvedComponents{Components: components, Repositories: repositories}, nil
}
func AddSourceRepositoriesToRefBuilder(b *app.ReferenceBuilder, c *ComponentInputs, g *GenerationInputs, s, i *[]string) (app.SourceRepositories, error) {
	_logClusterCodePath()
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
	strategy := g.Strategy
	if strategy == newapp.StrategyUnspecified {
		strategy = newapp.StrategySource
	}
	if git.IsGitInstalled() || len(c.SourceRepositories) > 0 {
		for _, s := range c.SourceRepositories {
			if repo, ok := b.AddSourceRepository(s, strategy); ok {
				repo.SetContextDir(g.ContextDir)
			}
		}
	} else if len(c.Components) > 0 && len(*i) > 0 && len(*s) == 0 || len(c.Components) > 0 && len(*i) == 0 && len(*s) > 0 {
		for _, s := range c.Components {
			if repo, ok := b.AddSourceRepository(s, strategy); ok {
				repo.SetContextDir(g.ContextDir)
				c.Components = []string{}
			}
		}
	}
	if len(g.Dockerfile) > 0 {
		if g.Strategy != newapp.StrategyUnspecified && g.Strategy != newapp.StrategyDocker {
			return nil, errors.New("when directly referencing a Dockerfile, the strategy must must be 'docker'")
		}
		if err := AddDockerfileToSourceRepositories(b, g.Dockerfile); err != nil {
			return nil, err
		}
	}
	_, result, errs := b.Result()
	return result, kutilerrors.NewAggregate(errs)
}
func AddDockerfileToSourceRepositories(b *app.ReferenceBuilder, dockerfile string) error {
	_logClusterCodePath()
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
	_, repos, errs := b.Result()
	if err := kutilerrors.NewAggregate(errs); err != nil {
		return err
	}
	switch len(repos) {
	case 0:
		repo, err := app.NewSourceRepositoryForDockerfile(dockerfile)
		if err != nil {
			return fmt.Errorf("provided Dockerfile is not valid: %v", err)
		}
		b.AddExistingSourceRepository(repo)
	case 1:
		if err := repos[0].AddDockerfile(dockerfile); err != nil {
			return fmt.Errorf("provided Dockerfile is not valid: %v", err)
		}
	default:
		return errors.New("--dockerfile cannot be used with multiple source repositories")
	}
	return nil
}
func DetectSource(repositories []*app.SourceRepository, d app.Detector, g *GenerationInputs) error {
	_logClusterCodePath()
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
	errs := []error{}
	for _, repo := range repositories {
		err := repo.Detect(d, g.Strategy == newapp.StrategyDocker || g.Strategy == newapp.StrategyPipeline)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		switch g.Strategy {
		case newapp.StrategyDocker:
			if repo.Info().Dockerfile == nil {
				errs = append(errs, errors.New("No Dockerfile was found in the repository and the requested build strategy is 'docker'"))
			}
		case newapp.StrategyPipeline:
			if !repo.Info().Jenkinsfile {
				errs = append(errs, errors.New("No Jenkinsfile was found in the repository and the requested build strategy is 'pipeline'"))
			}
		default:
			if repo.Info().Dockerfile == nil && !repo.Info().Jenkinsfile && len(repo.Info().Types) == 0 {
				errs = append(errs, errors.New("No language matched the source repository"))
			}
		}
	}
	return kutilerrors.NewAggregate(errs)
}
func AddComponentInputsToRefBuilder(b *app.ReferenceBuilder, r *Resolvers, c *ComponentInputs, g *GenerationInputs, s, i *[]string) error {
	_logClusterCodePath()
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
	repositories, err := AddSourceRepositoriesToRefBuilder(b, c, g, s, i)
	if err != nil {
		return err
	}
	if err := DetectSource(repositories, r.Detector, g); err != nil {
		return err
	}
	b.AddComponents(c.DockerImages, func(input *app.ComponentInput) app.ComponentReference {
		input.Argument = fmt.Sprintf("--docker-image=%q", input.From)
		input.Searcher = r.DockerSearcher
		if r.DockerSearcher != nil {
			resolver := app.PerfectMatchWeightedResolver{}
			resolver = append(resolver, app.WeightedResolver{Searcher: r.DockerSearcher, Weight: 0.0})
			if r.AllowMissingImages {
				resolver = append(resolver, app.WeightedResolver{Searcher: app.MissingImageSearcher{}, Weight: 100.0})
			}
			input.Resolver = resolver
		}
		return input
	})
	b.AddComponents(c.ImageStreams, func(input *app.ComponentInput) app.ComponentReference {
		input.Argument = fmt.Sprintf("--image-stream=%q", input.From)
		input.Searcher = r.ImageStreamSearcher
		if r.ImageStreamSearcher != nil {
			resolver := app.PerfectMatchWeightedResolver{app.WeightedResolver{Searcher: r.ImageStreamSearcher}}
			input.Resolver = resolver
		}
		return input
	})
	b.AddComponents(c.Templates, func(input *app.ComponentInput) app.ComponentReference {
		input.Argument = fmt.Sprintf("--template=%q", input.From)
		input.Searcher = r.TemplateSearcher
		if r.TemplateSearcher != nil {
			input.Resolver = app.HighestUniqueScoreResolver{Searcher: r.TemplateSearcher}
		}
		return input
	})
	b.AddComponents(c.TemplateFiles, func(input *app.ComponentInput) app.ComponentReference {
		input.Argument = fmt.Sprintf("--file=%q", input.From)
		if r.TemplateFileSearcher != nil {
			input.Resolver = app.FirstMatchResolver{Searcher: r.TemplateFileSearcher}
		}
		input.Searcher = r.TemplateFileSearcher
		return input
	})
	b.AddComponents(c.Components, func(input *app.ComponentInput) app.ComponentReference {
		resolver := app.PerfectMatchWeightedResolver{}
		searcher := app.MultiWeightedSearcher{}
		if r.ImageStreamSearcher != nil {
			resolver = append(resolver, app.WeightedResolver{Searcher: r.ImageStreamSearcher, Weight: 0.0})
			searcher = append(searcher, app.WeightedSearcher{Searcher: r.ImageStreamSearcher, Weight: 0.0})
		}
		if r.TemplateSearcher != nil && !input.ExpectToBuild {
			resolver = append(resolver, app.WeightedResolver{Searcher: r.TemplateSearcher, Weight: 0.0})
			searcher = append(searcher, app.WeightedSearcher{Searcher: r.TemplateSearcher, Weight: 0.0})
		}
		if r.TemplateFileSearcher != nil && !input.ExpectToBuild {
			resolver = append(resolver, app.WeightedResolver{Searcher: r.TemplateFileSearcher, Weight: 0.0})
		}
		if r.DockerSearcher != nil {
			resolver = append(resolver, app.WeightedResolver{Searcher: r.DockerSearcher, Weight: 2.0})
			searcher = append(searcher, app.WeightedSearcher{Searcher: r.DockerSearcher, Weight: 1.0})
		}
		if r.AllowMissingImages {
			resolver = append(resolver, app.WeightedResolver{Searcher: app.MissingImageSearcher{}, Weight: 100.0})
		}
		input.Resolver = resolver
		input.Searcher = searcher
		return input
	})
	b.AddGroups(c.Groups)
	return nil
}
func AddImageSourceRepository(sourceRepos app.SourceRepositories, r app.Resolver, g *GenerationInputs) (app.ComponentReference, app.SourceRepositories, error) {
	_logClusterCodePath()
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
	if len(g.SourceImage) == 0 {
		return nil, sourceRepos, nil
	}
	paths := strings.SplitN(g.SourceImagePath, ":", 2)
	var sourcePath, destPath string
	switch len(paths) {
	case 1:
		sourcePath = paths[0]
	case 2:
		sourcePath = paths[0]
		destPath = paths[1]
	}
	compRef, _, err := app.NewComponentInput(g.SourceImage)
	if err != nil {
		return nil, nil, err
	}
	compRef.Resolver = r
	switch len(sourceRepos) {
	case 0:
		sourceRepos = append(sourceRepos, app.NewImageSourceRepository(compRef, sourcePath, destPath))
	case 1:
		sourceRepos[0].SetSourceImage(compRef)
		sourceRepos[0].SetSourceImagePath(sourcePath, destPath)
	default:
		return nil, nil, errors.New("--source-image cannot be used with multiple source repositories")
	}
	return compRef, sourceRepos, nil
}
func detectPartialMatches(components app.ComponentReferences) error {
	_logClusterCodePath()
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
	errs := []error{}
	for _, ref := range components {
		input := ref.Input()
		if input.ResolvedMatch.Score != 0.0 {
			errs = append(errs, fmt.Errorf("component %q had only a partial match of %q - if this is the value you want to use, specify it explicitly", input.From, input.ResolvedMatch.Name))
		}
	}
	return kutilerrors.NewAggregate(errs)
}
func InferBuildTypes(components app.ComponentReferences, g *GenerationInputs) (app.ComponentReferences, error) {
	_logClusterCodePath()
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
	errs := []error{}
	for _, ref := range components {
		input := ref.Input()
		input.ResolvedMatch.Builder = app.IsBuilderMatch(input.ResolvedMatch)
		generatorInput, err := app.GeneratorInputFromMatch(input.ResolvedMatch)
		if err != nil && !g.AllowGenerationErrors {
			errs = append(errs, err)
			continue
		}
		input.ResolvedMatch.GeneratorInput = generatorInput
		if g.Strategy != newapp.StrategyUnspecified && input.Uses != nil {
			input.Uses.SetStrategy(g.Strategy)
		}
		if g.ExpectToBuild || (input.ResolvedMatch.Builder && g.Strategy != newapp.StrategyDocker) {
			input.ExpectToBuild = true
		}
		switch {
		case input.ExpectToBuild && input.ResolvedMatch.IsTemplate():
			errs = append(errs, errors.New("template with source code explicitly attached is not supported - you must either specify the template and source code separately or attach an image to the source code using the '[image]~[code]' form"))
			continue
		}
	}
	if len(components) == 0 && g.BinaryBuild && g.Strategy == newapp.StrategySource {
		return nil, errors.New("you must provide a builder image when using the source strategy with a binary build")
	}
	if len(components) == 0 && g.BinaryBuild {
		if len(g.Name) == 0 {
			return nil, errors.New("you must provide a --name when you don't specify a source repository or base image")
		}
		ref := &app.ComponentInput{From: "--binary", Argument: "--binary", Value: g.Name, ScratchImage: true, ExpectToBuild: true}
		components = append(components, ref)
	}
	return components, kutilerrors.NewAggregate(errs)
}
func EnsureHasSource(components app.ComponentReferences, repositories app.SourceRepositories, g *GenerationInputs) error {
	_logClusterCodePath()
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
	if len(components) == 0 {
		return nil
	}
	switch {
	case len(repositories) > 1:
		if len(components) == 1 {
			component := components[0]
			suggestions := ""
			for _, repo := range repositories {
				suggestions += fmt.Sprintf("%s~%s\n", component, repo)
			}
			return fmt.Errorf("there are multiple code locations provided - use one of the following suggestions to declare which code goes with the image:\n%s", suggestions)
		}
		return fmt.Errorf("the following images require source code: %s\n"+" and the following repositories are not used: %s\nUse '[image]~[repo]' to declare which code goes with which image", components, repositories)
	case len(repositories) == 1:
		klog.V(2).Infof("Using %q as the source for build", repositories[0])
		for _, component := range components {
			klog.V(2).Infof("Pairing with component %v", component)
			component.Input().Use(repositories[0])
			repositories[0].UsedBy(component)
		}
	default:
		switch {
		case g.BinaryBuild:
			for _, component := range components {
				input := component.Input()
				if input.Uses != nil {
					continue
				}
				strategy := newapp.StrategySource
				isBuilder := input.ResolvedMatch != nil && input.ResolvedMatch.Builder
				if g.Strategy == newapp.StrategyDocker || (g.Strategy == newapp.StrategyUnspecified && !isBuilder) {
					strategy = newapp.StrategyDocker
				}
				repo := app.NewBinarySourceRepository(strategy)
				input.Use(repo)
				repo.UsedBy(input)
				input.ExpectToBuild = true
			}
		case g.ExpectToBuild:
			return errors.New("you must specify at least one source repository URL, provide a Dockerfile, or indicate you wish to use binary builds")
		default:
			for _, component := range components {
				component.Input().ExpectToBuild = false
			}
		}
	}
	return nil
}
func AddMissingComponentsToRefBuilder(b *app.ReferenceBuilder, repositories app.SourceRepositories, dockerfileResolver, sourceResolver, pipelineResolver app.Resolver, g *GenerationInputs) (app.ComponentReferences, error) {
	_logClusterCodePath()
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
	errs := []error{}
	result := app.ComponentReferences{}
	for _, repo := range repositories {
		info := repo.Info()
		switch {
		case info == nil:
			errs = append(errs, fmt.Errorf("source not detected for repository %q", repo))
			continue
		case info.Jenkinsfile && (g.Strategy == newapp.StrategyUnspecified || g.Strategy == newapp.StrategyPipeline):
			refs := b.AddComponents([]string{"pipeline"}, func(input *app.ComponentInput) app.ComponentReference {
				input.Resolver = pipelineResolver
				input.Use(repo)
				input.ExpectToBuild = true
				repo.UsedBy(input)
				repo.SetStrategy(newapp.StrategyPipeline)
				return input
			})
			result = append(result, refs...)
		case info.Dockerfile != nil && (g.Strategy == newapp.StrategyUnspecified || g.Strategy == newapp.StrategyDocker):
			node := info.Dockerfile.AST()
			baseImage := dockerfileutil.LastBaseImage(node)
			if baseImage == "" {
				errs = append(errs, fmt.Errorf("the Dockerfile in the repository %q has no FROM instruction", info.Path))
				continue
			}
			refs := b.AddComponents([]string{baseImage}, func(input *app.ComponentInput) app.ComponentReference {
				input.Resolver = dockerfileResolver
				input.Use(repo)
				input.ExpectToBuild = true
				repo.UsedBy(input)
				repo.SetStrategy(newapp.StrategyDocker)
				return input
			})
			result = append(result, refs...)
		default:
			if len(info.Types) == 0 {
				errs = append(errs, fmt.Errorf("no language was detected for repository at %q; please specify a builder image to use with your repository: [builder-image]~%s", repo, repo))
				continue
			}
			refs := b.AddComponents([]string{info.Types[0].Term()}, func(input *app.ComponentInput) app.ComponentReference {
				input.Resolver = sourceResolver
				input.ExpectToBuild = true
				input.Use(repo)
				repo.UsedBy(input)
				return input
			})
			result = append(result, refs...)
		}
	}
	return result, kutilerrors.NewAggregate(errs)
}
