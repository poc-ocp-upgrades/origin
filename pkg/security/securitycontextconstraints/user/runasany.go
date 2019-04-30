package user

import (
	"k8s.io/apimachinery/pkg/util/validation/field"
	api "k8s.io/kubernetes/pkg/apis/core"
	securityapi "github.com/openshift/origin/pkg/security/apis/security"
)

type runAsAny struct{}

var _ RunAsUserSecurityContextConstraintsStrategy = &runAsAny{}

func NewRunAsAny(options *securityapi.RunAsUserStrategyOptions) (RunAsUserSecurityContextConstraintsStrategy, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &runAsAny{}, nil
}
func (s *runAsAny) Generate(pod *api.Pod, container *api.Container) (*int64, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil, nil
}
func (s *runAsAny) Validate(fldPath *field.Path, _ *api.Pod, _ *api.Container, runAsNonRoot *bool, runAsUser *int64) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return field.ErrorList{}
}
