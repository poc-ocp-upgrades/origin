package builds

import (
	"time"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog"
	buildv1 "github.com/openshift/api/build/v1"
	buildv1client "github.com/openshift/client-go/build/clientset/versioned/typed/build/v1"
)

type Pruner interface {
	Prune(deleter BuildDeleter) error
}
type BuildDeleter interface {
	DeleteBuild(build *buildv1.Build) error
}
type pruner struct{ resolver Resolver }

var _ Pruner = &pruner{}

type PrunerOptions struct {
	KeepYoungerThan	time.Duration
	Orphans		bool
	KeepComplete	int
	KeepFailed	int
	BuildConfigs	[]*buildv1.BuildConfig
	Builds		[]*buildv1.Build
}

func NewPruner(options PrunerOptions) Pruner {
	_logClusterCodePath()
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
	klog.V(1).Infof("Creating build pruner with keepYoungerThan=%v, orphans=%v, keepComplete=%v, keepFailed=%v", options.KeepYoungerThan, options.Orphans, options.KeepComplete, options.KeepFailed)
	filter := &andFilter{filterPredicates: []FilterPredicate{NewFilterBeforePredicate(options.KeepYoungerThan)}}
	builds := filter.Filter(options.Builds)
	dataSet := NewDataSet(options.BuildConfigs, builds)
	resolvers := []Resolver{}
	if options.Orphans {
		inactiveBuildStatus := []buildv1.BuildPhase{buildv1.BuildPhaseCancelled, buildv1.BuildPhaseComplete, buildv1.BuildPhaseError, buildv1.BuildPhaseFailed}
		resolvers = append(resolvers, NewOrphanBuildResolver(dataSet, inactiveBuildStatus))
	}
	resolvers = append(resolvers, NewPerBuildConfigResolver(dataSet, options.KeepComplete, options.KeepFailed))
	return &pruner{resolver: &mergeResolver{resolvers: resolvers}}
}
func (p *pruner) Prune(deleter BuildDeleter) error {
	_logClusterCodePath()
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
	builds, err := p.resolver.Resolve()
	if err != nil {
		return err
	}
	for _, build := range builds {
		if err := deleter.DeleteBuild(build); err != nil {
			return err
		}
	}
	return nil
}
func NewBuildDeleter(client buildv1client.BuildsGetter) BuildDeleter {
	_logClusterCodePath()
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
	return &buildDeleter{client: client}
}

type buildDeleter struct{ client buildv1client.BuildsGetter }

var _ BuildDeleter = &buildDeleter{}

func (c *buildDeleter) DeleteBuild(build *buildv1.Build) error {
	_logClusterCodePath()
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
	return c.client.Builds(build.Namespace).Delete(build.Name, &metav1.DeleteOptions{})
}
