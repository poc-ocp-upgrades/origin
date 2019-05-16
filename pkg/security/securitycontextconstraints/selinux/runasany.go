package selinux

import (
	securityapi "github.com/openshift/origin/pkg/security/apis/security"
	"k8s.io/apimachinery/pkg/util/validation/field"
	api "k8s.io/kubernetes/pkg/apis/core"
)

type runAsAny struct{}

var _ SELinuxSecurityContextConstraintsStrategy = &runAsAny{}

func NewRunAsAny(options *securityapi.SELinuxContextStrategyOptions) (SELinuxSecurityContextConstraintsStrategy, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &runAsAny{}, nil
}
func (s *runAsAny) Generate(pod *api.Pod, container *api.Container) (*api.SELinuxOptions, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return nil, nil
}
func (s *runAsAny) Validate(fldPath *field.Path, _ *api.Pod, _ *api.Container, options *api.SELinuxOptions) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return field.ErrorList{}
}
