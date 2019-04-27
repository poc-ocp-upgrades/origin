package securitycontextconstraints

import (
	"k8s.io/apimachinery/pkg/util/validation/field"
	api "k8s.io/kubernetes/pkg/apis/core"
)

type SecurityContextConstraintsProvider interface {
	CreatePodSecurityContext(pod *api.Pod) (*api.PodSecurityContext, map[string]string, error)
	CreateContainerSecurityContext(pod *api.Pod, container *api.Container) (*api.SecurityContext, error)
	ValidatePodSecurityContext(pod *api.Pod, fldPath *field.Path) field.ErrorList
	ValidateContainerSecurityContext(pod *api.Pod, container *api.Container, fldPath *field.Path) field.ErrorList
	GetSCCName() string
	GetSCCUsers() []string
	GetSCCGroups() []string
}
