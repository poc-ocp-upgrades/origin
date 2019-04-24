package selinux

import (
	"k8s.io/apimachinery/pkg/util/validation/field"
	api "k8s.io/kubernetes/pkg/apis/core"
	securityapi "github.com/openshift/origin/pkg/security/apis/security"
)

type runAsAny struct{}

var _ SELinuxSecurityContextConstraintsStrategy = &runAsAny{}

func NewRunAsAny(options *securityapi.SELinuxContextStrategyOptions) (SELinuxSecurityContextConstraintsStrategy, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &runAsAny{}, nil
}
func (s *runAsAny) Generate(pod *api.Pod, container *api.Container) (*api.SELinuxOptions, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil, nil
}
func (s *runAsAny) Validate(fldPath *field.Path, _ *api.Pod, _ *api.Container, options *api.SELinuxOptions) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return field.ErrorList{}
}
