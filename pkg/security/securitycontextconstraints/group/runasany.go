package group

import (
	"k8s.io/apimachinery/pkg/util/validation/field"
	api "k8s.io/kubernetes/pkg/apis/core"
)

type runAsAny struct{}

var _ GroupSecurityContextConstraintsStrategy = &runAsAny{}

func NewRunAsAny() (GroupSecurityContextConstraintsStrategy, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &runAsAny{}, nil
}
func (s *runAsAny) Generate(_ *api.Pod) ([]int64, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return nil, nil
}
func (s *runAsAny) GenerateSingle(_ *api.Pod) (*int64, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return nil, nil
}
func (s *runAsAny) Validate(_ *api.Pod, groups []int64) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return field.ErrorList{}
}
