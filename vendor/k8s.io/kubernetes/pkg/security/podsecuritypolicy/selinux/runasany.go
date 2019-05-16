package selinux

import (
	policy "k8s.io/api/policy/v1beta1"
	"k8s.io/apimachinery/pkg/util/validation/field"
	api "k8s.io/kubernetes/pkg/apis/core"
)

type runAsAny struct{}

var _ SELinuxStrategy = &runAsAny{}

func NewRunAsAny(options *policy.SELinuxStrategyOptions) (SELinuxStrategy, error) {
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
