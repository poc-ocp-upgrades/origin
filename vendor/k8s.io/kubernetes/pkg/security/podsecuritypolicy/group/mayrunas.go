package group

import (
 "fmt"
 policy "k8s.io/api/policy/v1beta1"
 "k8s.io/apimachinery/pkg/util/validation/field"
 api "k8s.io/kubernetes/pkg/apis/core"
)

type mayRunAs struct{ ranges []policy.IDRange }

var _ GroupStrategy = &mayRunAs{}

func NewMayRunAs(ranges []policy.IDRange) (GroupStrategy, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if len(ranges) == 0 {
  return nil, fmt.Errorf("ranges must be supplied for MayRunAs")
 }
 return &mayRunAs{ranges: ranges}, nil
}
func (s *mayRunAs) Generate(_ *api.Pod) ([]int64, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return nil, nil
}
func (s *mayRunAs) GenerateSingle(_ *api.Pod) (*int64, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return nil, nil
}
func (s *mayRunAs) Validate(fldPath *field.Path, _ *api.Pod, groups []int64) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return ValidateGroupsInRanges(fldPath, s.ranges, groups)
}
