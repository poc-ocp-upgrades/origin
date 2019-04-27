package builds

import (
	"sort"
	"k8s.io/apimachinery/pkg/util/sets"
	buildv1 "github.com/openshift/api/build/v1"
	"github.com/openshift/origin/pkg/oc/lib/buildapihelpers"
)

type Resolver interface {
	Resolve() ([]*buildv1.Build, error)
}
type mergeResolver struct{ resolvers []Resolver }

func (m *mergeResolver) Resolve() ([]*buildv1.Build, error) {
	_logClusterCodePath()
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
	results := []*buildv1.Build{}
	for _, resolver := range m.resolvers {
		builds, err := resolver.Resolve()
		if err != nil {
			return nil, err
		}
		results = append(results, builds...)
	}
	return results, nil
}
func NewOrphanBuildResolver(dataSet DataSet, BuildPhaseFilter []buildv1.BuildPhase) Resolver {
	_logClusterCodePath()
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
	filter := sets.NewString()
	for _, BuildPhase := range BuildPhaseFilter {
		filter.Insert(string(BuildPhase))
	}
	return &orphanBuildResolver{dataSet: dataSet, BuildPhaseFilter: filter}
}

type orphanBuildResolver struct {
	dataSet			DataSet
	BuildPhaseFilter	sets.String
}

func (o *orphanBuildResolver) Resolve() ([]*buildv1.Build, error) {
	_logClusterCodePath()
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
	builds, err := o.dataSet.ListBuilds()
	if err != nil {
		return nil, err
	}
	results := []*buildv1.Build{}
	for _, build := range builds {
		if !o.BuildPhaseFilter.Has(string(build.Status.Phase)) {
			continue
		}
		isOrphan := false
		if build.Status.Config == nil {
			isOrphan = true
		} else {
			_, exists, _ := o.dataSet.GetBuildConfig(build)
			isOrphan = !exists
		}
		if isOrphan {
			results = append(results, build)
		}
	}
	return results, nil
}

type perBuildConfigResolver struct {
	dataSet		DataSet
	keepComplete	int
	keepFailed	int
}

func NewPerBuildConfigResolver(dataSet DataSet, keepComplete int, keepFailed int) Resolver {
	_logClusterCodePath()
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
	return &perBuildConfigResolver{dataSet: dataSet, keepComplete: keepComplete, keepFailed: keepFailed}
}
func (o *perBuildConfigResolver) Resolve() ([]*buildv1.Build, error) {
	_logClusterCodePath()
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
	buildConfigs, err := o.dataSet.ListBuildConfigs()
	if err != nil {
		return nil, err
	}
	completeStates := sets.NewString(string(buildv1.BuildPhaseComplete))
	failedStates := sets.NewString(string(buildv1.BuildPhaseFailed), string(buildv1.BuildPhaseError), string(buildv1.BuildPhaseCancelled))
	prunableBuilds := []*buildv1.Build{}
	for _, buildConfig := range buildConfigs {
		builds, err := o.dataSet.ListBuildsByBuildConfig(buildConfig)
		if err != nil {
			return nil, err
		}
		var completeBuilds, failedBuilds []*buildv1.Build
		for _, build := range builds {
			if completeStates.Has(string(build.Status.Phase)) {
				completeBuilds = append(completeBuilds, build)
			} else if failedStates.Has(string(build.Status.Phase)) {
				failedBuilds = append(failedBuilds, build)
			}
		}
		sort.Sort(sort.Reverse(buildapihelpers.BuildPtrSliceByCreationTimestamp(completeBuilds)))
		sort.Sort(sort.Reverse(buildapihelpers.BuildPtrSliceByCreationTimestamp(failedBuilds)))
		if o.keepComplete >= 0 && o.keepComplete < len(completeBuilds) {
			prunableBuilds = append(prunableBuilds, completeBuilds[o.keepComplete:]...)
		}
		if o.keepFailed >= 0 && o.keepFailed < len(failedBuilds) {
			prunableBuilds = append(prunableBuilds, failedBuilds[o.keepFailed:]...)
		}
	}
	return prunableBuilds, nil
}
