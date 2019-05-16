package rollback

import (
	"fmt"
	appsapi "github.com/openshift/origin/pkg/apps/apis/apps"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
)

type RollbackGenerator interface {
	GenerateRollback(from, to *appsapi.DeploymentConfig, spec *appsapi.DeploymentConfigRollbackSpec) (*appsapi.DeploymentConfig, error)
}

func NewRollbackGenerator() RollbackGenerator {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &rollbackGenerator{}
}

type rollbackGenerator struct{}

func (g *rollbackGenerator) GenerateRollback(from, to *appsapi.DeploymentConfig, spec *appsapi.DeploymentConfigRollbackSpec) (*appsapi.DeploymentConfig, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	rollback := &appsapi.DeploymentConfig{}
	if err := legacyscheme.Scheme.Convert(&from, &rollback, nil); err != nil {
		return nil, fmt.Errorf("couldn't clone 'from' DeploymentConfig: %v", err)
	}
	if spec.IncludeTemplate {
		if err := legacyscheme.Scheme.Convert(&to.Spec.Template, &rollback.Spec.Template, nil); err != nil {
			return nil, fmt.Errorf("couldn't copy template to rollback:: %v", err)
		}
	}
	if spec.IncludeReplicationMeta {
		rollback.Spec.Replicas = to.Spec.Replicas
		rollback.Spec.Selector = map[string]string{}
		for k, v := range to.Spec.Selector {
			rollback.Spec.Selector[k] = v
		}
	}
	if spec.IncludeTriggers {
		if err := legacyscheme.Scheme.Convert(&to.Spec.Triggers, &rollback.Spec.Triggers, nil); err != nil {
			return nil, fmt.Errorf("couldn't copy triggers to rollback:: %v", err)
		}
	}
	if spec.IncludeStrategy {
		if err := legacyscheme.Scheme.Convert(&to.Spec.Strategy, &rollback.Spec.Strategy, nil); err != nil {
			return nil, fmt.Errorf("couldn't copy strategy to rollback:: %v", err)
		}
	}
	for _, trigger := range rollback.Spec.Triggers {
		if trigger.Type == appsapi.DeploymentTriggerOnImageChange {
			trigger.ImageChangeParams.Automatic = false
		}
	}
	rollback.Status.LatestVersion++
	return rollback, nil
}
