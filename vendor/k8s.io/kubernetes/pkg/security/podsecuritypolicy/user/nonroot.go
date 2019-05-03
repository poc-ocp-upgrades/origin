package user

import (
 policy "k8s.io/api/policy/v1beta1"
 "k8s.io/apimachinery/pkg/util/validation/field"
 api "k8s.io/kubernetes/pkg/apis/core"
)

type nonRoot struct{}

var _ RunAsUserStrategy = &nonRoot{}

func NewRunAsNonRoot(options *policy.RunAsUserStrategyOptions) (RunAsUserStrategy, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &nonRoot{}, nil
}
func (s *nonRoot) Generate(pod *api.Pod, container *api.Container) (*int64, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return nil, nil
}
func (s *nonRoot) Validate(scPath *field.Path, _ *api.Pod, _ *api.Container, runAsNonRoot *bool, runAsUser *int64) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if runAsNonRoot == nil && runAsUser == nil {
  allErrs = append(allErrs, field.Required(scPath.Child("runAsNonRoot"), "must be true"))
  return allErrs
 }
 if runAsNonRoot != nil && *runAsNonRoot == false {
  allErrs = append(allErrs, field.Invalid(scPath.Child("runAsNonRoot"), *runAsNonRoot, "must be true"))
  return allErrs
 }
 if runAsUser != nil && *runAsUser == 0 {
  allErrs = append(allErrs, field.Invalid(scPath.Child("runAsUser"), *runAsUser, "running with the root UID is forbidden"))
  return allErrs
 }
 return allErrs
}
