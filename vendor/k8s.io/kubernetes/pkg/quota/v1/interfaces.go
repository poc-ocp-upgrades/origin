package quota

import (
	goformat "fmt"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apiserver/pkg/admission"
	"k8s.io/client-go/tools/cache"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

type UsageStatsOptions struct {
	Namespace     string
	Scopes        []corev1.ResourceQuotaScope
	Resources     []corev1.ResourceName
	ScopeSelector *corev1.ScopeSelector
}
type UsageStats struct{ Used corev1.ResourceList }
type Evaluator interface {
	Constraints(required []corev1.ResourceName, item runtime.Object) error
	GroupResource() schema.GroupResource
	Handles(operation admission.Attributes) bool
	Matches(resourceQuota *corev1.ResourceQuota, item runtime.Object) (bool, error)
	MatchingScopes(item runtime.Object, scopes []corev1.ScopedResourceSelectorRequirement) ([]corev1.ScopedResourceSelectorRequirement, error)
	UncoveredQuotaScopes(limitedScopes []corev1.ScopedResourceSelectorRequirement, matchedQuotaScopes []corev1.ScopedResourceSelectorRequirement) ([]corev1.ScopedResourceSelectorRequirement, error)
	MatchingResources(input []corev1.ResourceName) []corev1.ResourceName
	Usage(item runtime.Object) (corev1.ResourceList, error)
	UsageStats(options UsageStatsOptions) (UsageStats, error)
}
type Configuration interface {
	IgnoredResources() map[schema.GroupResource]struct{}
	Evaluators() []Evaluator
}
type Registry interface {
	Add(e Evaluator)
	Remove(e Evaluator)
	Get(gr schema.GroupResource) Evaluator
	List() []Evaluator
}
type ListerForResourceFunc func(schema.GroupVersionResource) (cache.GenericLister, error)

func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
