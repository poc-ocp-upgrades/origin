package user

import (
 "k8s.io/apimachinery/pkg/util/validation/field"
 api "k8s.io/kubernetes/pkg/apis/core"
)

type RunAsUserStrategy interface {
 Generate(pod *api.Pod, container *api.Container) (*int64, error)
 Validate(scPath *field.Path, pod *api.Pod, container *api.Container, runAsNonRoot *bool, runAsUser *int64) field.ErrorList
}
