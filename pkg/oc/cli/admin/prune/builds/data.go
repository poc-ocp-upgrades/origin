package builds

import (
	"fmt"
	"time"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/cache"
	buildv1 "github.com/openshift/api/build/v1"
)

func BuildByBuildConfigIndexFunc(obj interface{}) ([]string, error) {
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
	build, ok := obj.(*buildv1.Build)
	if !ok {
		return nil, fmt.Errorf("not a build: %v", build)
	}
	config := build.Status.Config
	if config == nil {
		return []string{"orphan"}, nil
	}
	return []string{config.Namespace + "/" + config.Name}, nil
}

type Filter interface {
	Filter(builds []*buildv1.Build) []*buildv1.Build
}
type andFilter struct{ filterPredicates []FilterPredicate }

func (a *andFilter) Filter(builds []*buildv1.Build) []*buildv1.Build {
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
	results := []*buildv1.Build{}
	for _, build := range builds {
		include := true
		for _, filterPredicate := range a.filterPredicates {
			include = include && filterPredicate(build)
		}
		if include {
			results = append(results, build)
		}
	}
	return results
}

type FilterPredicate func(build *buildv1.Build) bool

func NewFilterBeforePredicate(d time.Duration) FilterPredicate {
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
	now := metav1.Now()
	before := metav1.NewTime(now.Time.Add(-1 * d))
	return func(build *buildv1.Build) bool {
		return build.CreationTimestamp.Before(&before)
	}
}

type DataSet interface {
	GetBuildConfig(build *buildv1.Build) (*buildv1.BuildConfig, bool, error)
	ListBuildConfigs() ([]*buildv1.BuildConfig, error)
	ListBuilds() ([]*buildv1.Build, error)
	ListBuildsByBuildConfig(buildConfig *buildv1.BuildConfig) ([]*buildv1.Build, error)
}
type dataSet struct {
	buildConfigStore	cache.Store
	buildIndexer		cache.Indexer
}

func NewDataSet(buildConfigs []*buildv1.BuildConfig, builds []*buildv1.Build) DataSet {
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
	buildConfigStore := cache.NewStore(cache.MetaNamespaceKeyFunc)
	for _, buildConfig := range buildConfigs {
		buildConfigStore.Add(buildConfig)
	}
	buildIndexer := cache.NewIndexer(cache.MetaNamespaceKeyFunc, cache.Indexers{"buildConfig": BuildByBuildConfigIndexFunc})
	for _, build := range builds {
		buildIndexer.Add(build)
	}
	return &dataSet{buildConfigStore: buildConfigStore, buildIndexer: buildIndexer}
}
func (d *dataSet) GetBuildConfig(build *buildv1.Build) (*buildv1.BuildConfig, bool, error) {
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
	config := build.Status.Config
	if config == nil {
		return nil, false, nil
	}
	var buildConfig *buildv1.BuildConfig
	key := &buildv1.BuildConfig{ObjectMeta: metav1.ObjectMeta{Name: config.Name, Namespace: config.Namespace}}
	item, exists, err := d.buildConfigStore.Get(key)
	if exists {
		buildConfig = item.(*buildv1.BuildConfig)
	}
	return buildConfig, exists, err
}
func (d *dataSet) ListBuildConfigs() ([]*buildv1.BuildConfig, error) {
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
	results := []*buildv1.BuildConfig{}
	for _, item := range d.buildConfigStore.List() {
		results = append(results, item.(*buildv1.BuildConfig))
	}
	return results, nil
}
func (d *dataSet) ListBuilds() ([]*buildv1.Build, error) {
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
	results := []*buildv1.Build{}
	for _, item := range d.buildIndexer.List() {
		results = append(results, item.(*buildv1.Build))
	}
	return results, nil
}
func (d *dataSet) ListBuildsByBuildConfig(buildConfig *buildv1.BuildConfig) ([]*buildv1.Build, error) {
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
	results := []*buildv1.Build{}
	key := &buildv1.Build{}
	key.Status.Config = &corev1.ObjectReference{Name: buildConfig.Name, Namespace: buildConfig.Namespace}
	items, err := d.buildIndexer.Index("buildConfig", key)
	if err != nil {
		return nil, err
	}
	for _, item := range items {
		results = append(results, item.(*buildv1.Build))
	}
	return results, nil
}
