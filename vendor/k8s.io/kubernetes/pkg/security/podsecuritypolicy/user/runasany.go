package user

import (
 policy "k8s.io/api/policy/v1beta1"
 "k8s.io/apimachinery/pkg/util/validation/field"
 api "k8s.io/kubernetes/pkg/apis/core"
)

type runAsAny struct{}

var _ RunAsUserStrategy = &runAsAny{}

func NewRunAsAny(options *policy.RunAsUserStrategyOptions) (RunAsUserStrategy, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &runAsAny{}, nil
}
func (s *runAsAny) Generate(pod *api.Pod, container *api.Container) (*int64, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return nil, nil
}
func (s *runAsAny) Validate(_ *field.Path, _ *api.Pod, _ *api.Container, runAsNonRoot *bool, runAsUser *int64) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return field.ErrorList{}
}
