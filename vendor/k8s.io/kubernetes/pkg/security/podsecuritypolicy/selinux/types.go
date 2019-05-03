package selinux

import (
 "k8s.io/apimachinery/pkg/util/validation/field"
 api "k8s.io/kubernetes/pkg/apis/core"
)

type SELinuxStrategy interface {
 Generate(pod *api.Pod, container *api.Container) (*api.SELinuxOptions, error)
 Validate(fldPath *field.Path, pod *api.Pod, container *api.Container, options *api.SELinuxOptions) field.ErrorList
}
