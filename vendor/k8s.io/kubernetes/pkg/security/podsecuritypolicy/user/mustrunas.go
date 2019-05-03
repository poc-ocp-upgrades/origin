package user

import (
 "fmt"
 policy "k8s.io/api/policy/v1beta1"
 "k8s.io/apimachinery/pkg/util/validation/field"
 api "k8s.io/kubernetes/pkg/apis/core"
 psputil "k8s.io/kubernetes/pkg/security/podsecuritypolicy/util"
)

type mustRunAs struct {
 opts *policy.RunAsUserStrategyOptions
}

func NewMustRunAs(options *policy.RunAsUserStrategyOptions) (RunAsUserStrategy, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if options == nil {
  return nil, fmt.Errorf("MustRunAs requires run as user options")
 }
 if len(options.Ranges) == 0 {
  return nil, fmt.Errorf("MustRunAs requires at least one range")
 }
 return &mustRunAs{opts: options}, nil
}
func (s *mustRunAs) Generate(pod *api.Pod, container *api.Container) (*int64, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &s.opts.Ranges[0].Min, nil
}
func (s *mustRunAs) Validate(scPath *field.Path, _ *api.Pod, _ *api.Container, runAsNonRoot *bool, runAsUser *int64) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if runAsUser == nil {
  allErrs = append(allErrs, field.Required(scPath.Child("runAsUser"), ""))
  return allErrs
 }
 if !s.isValidUID(*runAsUser) {
  detail := fmt.Sprintf("must be in the ranges: %v", s.opts.Ranges)
  allErrs = append(allErrs, field.Invalid(scPath.Child("runAsUser"), *runAsUser, detail))
 }
 return allErrs
}
func (s *mustRunAs) isValidUID(id int64) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for _, rng := range s.opts.Ranges {
  if psputil.UserFallsInRange(id, rng) {
   return true
  }
 }
 return false
}
