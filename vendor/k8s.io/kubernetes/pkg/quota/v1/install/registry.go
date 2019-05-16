package install

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	quota "k8s.io/kubernetes/pkg/quota/v1"
	core "k8s.io/kubernetes/pkg/quota/v1/evaluator/core"
	generic "k8s.io/kubernetes/pkg/quota/v1/generic"
)

func NewQuotaConfigurationForAdmission() quota.Configuration {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	evaluators := core.NewEvaluators(nil)
	return generic.NewConfiguration(evaluators, DefaultIgnoredResources())
}
func NewQuotaConfigurationForControllers(f quota.ListerForResourceFunc) quota.Configuration {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	evaluators := core.NewEvaluators(f)
	return generic.NewConfiguration(evaluators, DefaultIgnoredResources())
}

var ignoredResources = map[schema.GroupResource]struct{}{{Group: "", Resource: "events"}: {}}

func DefaultIgnoredResources() map[schema.GroupResource]struct{} {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return ignoredResources
}
