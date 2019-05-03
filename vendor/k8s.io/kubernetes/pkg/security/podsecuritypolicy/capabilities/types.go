package capabilities

import (
 "k8s.io/apimachinery/pkg/util/validation/field"
 api "k8s.io/kubernetes/pkg/apis/core"
)

type Strategy interface {
 Generate(pod *api.Pod, container *api.Container) (*api.Capabilities, error)
 Validate(fldPath *field.Path, pod *api.Pod, container *api.Container, capabilities *api.Capabilities) field.ErrorList
}
