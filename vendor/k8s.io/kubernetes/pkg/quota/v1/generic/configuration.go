package generic

import (
	goformat "fmt"
	"k8s.io/apimachinery/pkg/runtime/schema"
	quota "k8s.io/kubernetes/pkg/quota/v1"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

type simpleConfiguration struct {
	evaluators       []quota.Evaluator
	ignoredResources map[schema.GroupResource]struct{}
}

func NewConfiguration(evaluators []quota.Evaluator, ignoredResources map[schema.GroupResource]struct{}) quota.Configuration {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &simpleConfiguration{evaluators: evaluators, ignoredResources: ignoredResources}
}
func (c *simpleConfiguration) IgnoredResources() map[schema.GroupResource]struct{} {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return c.ignoredResources
}
func (c *simpleConfiguration) Evaluators() []quota.Evaluator {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return c.evaluators
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
