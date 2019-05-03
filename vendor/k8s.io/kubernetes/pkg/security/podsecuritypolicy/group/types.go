package group

import (
 "k8s.io/apimachinery/pkg/util/validation/field"
 api "k8s.io/kubernetes/pkg/apis/core"
)

type GroupStrategy interface {
 Generate(pod *api.Pod) ([]int64, error)
 GenerateSingle(pod *api.Pod) (*int64, error)
 Validate(fldPath *field.Path, pod *api.Pod, groups []int64) field.ErrorList
}
