package imagereferencemutators

import (
	"fmt"
	kappsv1 "k8s.io/api/apps/v1"
	kappsv1beta1 "k8s.io/api/apps/v1beta1"
	kappsv1beta2 "k8s.io/api/apps/v1beta2"
	batchv1 "k8s.io/api/batch/v1"
	batchv1beta1 "k8s.io/api/batch/v1beta1"
	batchv2alpha1 "k8s.io/api/batch/v2alpha1"
	corev1 "k8s.io/api/core/v1"
	extensionsv1beta1 "k8s.io/api/extensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/kubernetes/pkg/apis/apps"
	"k8s.io/kubernetes/pkg/apis/batch"
	kapi "k8s.io/kubernetes/pkg/apis/core"
	appsv1 "github.com/openshift/api/apps/v1"
	securityv1 "github.com/openshift/api/security/v1"
	appsapi "github.com/openshift/origin/pkg/apps/apis/apps"
	securityapi "github.com/openshift/origin/pkg/security/apis/security"
)

type ContainerMutator interface {
	GetName() string
	GetImage() string
	SetImage(image string)
}
type PodSpecReferenceMutator interface {
	GetContainerByIndex(init bool, i int) (ContainerMutator, bool)
	GetContainerByName(name string) (ContainerMutator, bool)
	Path() *field.Path
}

func GetPodSpecReferenceMutator(obj runtime.Object) (PodSpecReferenceMutator, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if spec, path, err := GetPodSpec(obj); err == nil {
		return &podSpecMutator{spec: spec, path: path}, nil
	}
	if spec, path, err := GetPodSpecV1(obj); err == nil {
		return &podSpecV1Mutator{spec: spec, path: path}, nil
	}
	return nil, errNoImageMutator
}

var errNoPodSpec = fmt.Errorf("No PodSpec available for this object")

func GetPodSpec(obj runtime.Object) (*kapi.PodSpec, *field.Path, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch r := obj.(type) {
	case *kapi.Pod:
		return &r.Spec, field.NewPath("spec"), nil
	case *kapi.PodTemplate:
		return &r.Template.Spec, field.NewPath("template", "spec"), nil
	case *kapi.ReplicationController:
		if r.Spec.Template != nil {
			return &r.Spec.Template.Spec, field.NewPath("spec", "template", "spec"), nil
		}
	case *apps.DaemonSet:
		return &r.Spec.Template.Spec, field.NewPath("spec", "template", "spec"), nil
	case *apps.Deployment:
		return &r.Spec.Template.Spec, field.NewPath("spec", "template", "spec"), nil
	case *apps.ReplicaSet:
		return &r.Spec.Template.Spec, field.NewPath("spec", "template", "spec"), nil
	case *batch.Job:
		return &r.Spec.Template.Spec, field.NewPath("spec", "template", "spec"), nil
	case *batch.CronJob:
		return &r.Spec.JobTemplate.Spec.Template.Spec, field.NewPath("spec", "jobTemplate", "spec", "template", "spec"), nil
	case *batch.JobTemplate:
		return &r.Template.Spec.Template.Spec, field.NewPath("template", "spec", "template", "spec"), nil
	case *apps.StatefulSet:
		return &r.Spec.Template.Spec, field.NewPath("spec", "template", "spec"), nil
	case *securityapi.PodSecurityPolicySubjectReview:
		return &r.Spec.Template.Spec, field.NewPath("spec", "template", "spec"), nil
	case *securityapi.PodSecurityPolicySelfSubjectReview:
		return &r.Spec.Template.Spec, field.NewPath("spec", "template", "spec"), nil
	case *securityapi.PodSecurityPolicyReview:
		return &r.Spec.Template.Spec, field.NewPath("spec", "template", "spec"), nil
	case *appsapi.DeploymentConfig:
		if r.Spec.Template != nil {
			return &r.Spec.Template.Spec, field.NewPath("spec", "template", "spec"), nil
		}
	}
	return nil, nil, errNoPodSpec
}
func GetPodSpecV1(obj runtime.Object) (*corev1.PodSpec, *field.Path, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch r := obj.(type) {
	case *corev1.Pod:
		return &r.Spec, field.NewPath("spec"), nil
	case *corev1.PodTemplate:
		return &r.Template.Spec, field.NewPath("template", "spec"), nil
	case *corev1.ReplicationController:
		if r.Spec.Template != nil {
			return &r.Spec.Template.Spec, field.NewPath("spec", "template", "spec"), nil
		}
	case *extensionsv1beta1.DaemonSet:
		return &r.Spec.Template.Spec, field.NewPath("spec", "template", "spec"), nil
	case *kappsv1.DaemonSet:
		return &r.Spec.Template.Spec, field.NewPath("spec", "template", "spec"), nil
	case *kappsv1beta2.DaemonSet:
		return &r.Spec.Template.Spec, field.NewPath("spec", "template", "spec"), nil
	case *extensionsv1beta1.Deployment:
		return &r.Spec.Template.Spec, field.NewPath("spec", "template", "spec"), nil
	case *kappsv1.Deployment:
		return &r.Spec.Template.Spec, field.NewPath("spec", "template", "spec"), nil
	case *kappsv1beta1.Deployment:
		return &r.Spec.Template.Spec, field.NewPath("spec", "template", "spec"), nil
	case *kappsv1beta2.Deployment:
		return &r.Spec.Template.Spec, field.NewPath("spec", "template", "spec"), nil
	case *extensionsv1beta1.ReplicaSet:
		return &r.Spec.Template.Spec, field.NewPath("spec", "template", "spec"), nil
	case *kappsv1.ReplicaSet:
		return &r.Spec.Template.Spec, field.NewPath("spec", "template", "spec"), nil
	case *kappsv1beta2.ReplicaSet:
		return &r.Spec.Template.Spec, field.NewPath("spec", "template", "spec"), nil
	case *batchv1.Job:
		return &r.Spec.Template.Spec, field.NewPath("spec", "template", "spec"), nil
	case *batchv2alpha1.CronJob:
		return &r.Spec.JobTemplate.Spec.Template.Spec, field.NewPath("spec", "jobTemplate", "spec", "template", "spec"), nil
	case *batchv1beta1.CronJob:
		return &r.Spec.JobTemplate.Spec.Template.Spec, field.NewPath("spec", "jobTemplate", "spec", "template", "spec"), nil
	case *batchv2alpha1.JobTemplate:
		return &r.Template.Spec.Template.Spec, field.NewPath("template", "spec", "template", "spec"), nil
	case *batchv1beta1.JobTemplate:
		return &r.Template.Spec.Template.Spec, field.NewPath("template", "spec", "template", "spec"), nil
	case *kappsv1.StatefulSet:
		return &r.Spec.Template.Spec, field.NewPath("spec", "template", "spec"), nil
	case *kappsv1beta1.StatefulSet:
		return &r.Spec.Template.Spec, field.NewPath("spec", "template", "spec"), nil
	case *kappsv1beta2.StatefulSet:
		return &r.Spec.Template.Spec, field.NewPath("spec", "template", "spec"), nil
	case *securityv1.PodSecurityPolicySubjectReview:
		return &r.Spec.Template.Spec, field.NewPath("spec", "template", "spec"), nil
	case *securityv1.PodSecurityPolicySelfSubjectReview:
		return &r.Spec.Template.Spec, field.NewPath("spec", "template", "spec"), nil
	case *securityv1.PodSecurityPolicyReview:
		return &r.Spec.Template.Spec, field.NewPath("spec", "template", "spec"), nil
	case *appsv1.DeploymentConfig:
		if r.Spec.Template != nil {
			return &r.Spec.Template.Spec, field.NewPath("spec", "template", "spec"), nil
		}
	}
	return nil, nil, errNoPodSpec
}
func GetTemplateMetaObject(obj runtime.Object) (metav1.Object, bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch r := obj.(type) {
	case *kapi.PodTemplate:
		return &r.Template.ObjectMeta, true
	case *corev1.PodTemplate:
		return &r.Template.ObjectMeta, true
	case *kapi.ReplicationController:
		if r.Spec.Template != nil {
			return &r.Spec.Template.ObjectMeta, true
		}
	case *corev1.ReplicationController:
		if r.Spec.Template != nil {
			return &r.Spec.Template.ObjectMeta, true
		}
	case *apps.DaemonSet:
		return &r.Spec.Template.ObjectMeta, true
	case *extensionsv1beta1.DaemonSet:
		return &r.Spec.Template.ObjectMeta, true
	case *kappsv1.DaemonSet:
		return &r.Spec.Template.ObjectMeta, true
	case *kappsv1beta2.DaemonSet:
		return &r.Spec.Template.ObjectMeta, true
	case *apps.Deployment:
		return &r.Spec.Template.ObjectMeta, true
	case *extensionsv1beta1.Deployment:
		return &r.Spec.Template.ObjectMeta, true
	case *kappsv1.Deployment:
		return &r.Spec.Template.ObjectMeta, true
	case *kappsv1beta1.Deployment:
		return &r.Spec.Template.ObjectMeta, true
	case *kappsv1beta2.Deployment:
		return &r.Spec.Template.ObjectMeta, true
	case *apps.ReplicaSet:
		return &r.Spec.Template.ObjectMeta, true
	case *extensionsv1beta1.ReplicaSet:
		return &r.Spec.Template.ObjectMeta, true
	case *kappsv1.ReplicaSet:
		return &r.Spec.Template.ObjectMeta, true
	case *kappsv1beta2.ReplicaSet:
		return &r.Spec.Template.ObjectMeta, true
	case *batch.Job:
		return &r.Spec.Template.ObjectMeta, true
	case *batchv1.Job:
		return &r.Spec.Template.ObjectMeta, true
	case *batch.CronJob:
		return &r.Spec.JobTemplate.Spec.Template.ObjectMeta, true
	case *batchv2alpha1.CronJob:
		return &r.Spec.JobTemplate.Spec.Template.ObjectMeta, true
	case *batchv1beta1.CronJob:
		return &r.Spec.JobTemplate.Spec.Template.ObjectMeta, true
	case *batch.JobTemplate:
		return &r.Template.Spec.Template.ObjectMeta, true
	case *batchv2alpha1.JobTemplate:
		return &r.Template.Spec.Template.ObjectMeta, true
	case *batchv1beta1.JobTemplate:
		return &r.Template.Spec.Template.ObjectMeta, true
	case *apps.StatefulSet:
		return &r.Spec.Template.ObjectMeta, true
	case *kappsv1.StatefulSet:
		return &r.Spec.Template.ObjectMeta, true
	case *kappsv1beta1.StatefulSet:
		return &r.Spec.Template.ObjectMeta, true
	case *kappsv1beta2.StatefulSet:
		return &r.Spec.Template.ObjectMeta, true
	case *securityapi.PodSecurityPolicySubjectReview:
		return &r.Spec.Template.ObjectMeta, true
	case *securityv1.PodSecurityPolicySubjectReview:
		return &r.Spec.Template.ObjectMeta, true
	case *securityapi.PodSecurityPolicySelfSubjectReview:
		return &r.Spec.Template.ObjectMeta, true
	case *securityv1.PodSecurityPolicySelfSubjectReview:
		return &r.Spec.Template.ObjectMeta, true
	case *securityapi.PodSecurityPolicyReview:
		return &r.Spec.Template.ObjectMeta, true
	case *securityv1.PodSecurityPolicyReview:
		return &r.Spec.Template.ObjectMeta, true
	case *appsapi.DeploymentConfig:
		if r.Spec.Template != nil {
			return &r.Spec.Template.ObjectMeta, true
		}
	case *appsv1.DeploymentConfig:
		if r.Spec.Template != nil {
			return &r.Spec.Template.ObjectMeta, true
		}
	}
	return nil, false
}

type containerMutator struct{ *kapi.Container }

func (m containerMutator) GetName() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return m.Name
}
func (m containerMutator) GetImage() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return m.Image
}
func (m containerMutator) SetImage(image string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	m.Image = image
}

type containerV1Mutator struct{ *corev1.Container }

func (m containerV1Mutator) GetName() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return m.Name
}
func (m containerV1Mutator) GetImage() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return m.Image
}
func (m containerV1Mutator) SetImage(image string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	m.Image = image
}

type podSpecMutator struct {
	spec	*kapi.PodSpec
	oldSpec	*kapi.PodSpec
	path	*field.Path
}

func (m *podSpecMutator) Path() *field.Path {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return m.path
}
func hasIdenticalPodSpecImage(spec *kapi.PodSpec, containerName, image string) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if spec == nil {
		return false
	}
	for i := range spec.InitContainers {
		if spec.InitContainers[i].Name == containerName {
			return spec.InitContainers[i].Image == image
		}
	}
	for i := range spec.Containers {
		if spec.Containers[i].Name == containerName {
			return spec.Containers[i].Image == image
		}
	}
	return false
}
func (m *podSpecMutator) Mutate(fn ImageReferenceMutateFunc) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var errs field.ErrorList
	for i := range m.spec.InitContainers {
		container := &m.spec.InitContainers[i]
		if hasIdenticalPodSpecImage(m.oldSpec, container.Name, container.Image) {
			continue
		}
		ref := corev1.ObjectReference{Kind: "DockerImage", Name: container.Image}
		if err := fn(&ref); err != nil {
			errs = append(errs, fieldErrorOrInternal(err, m.path.Child("initContainers").Index(i).Child("image")))
			continue
		}
		if ref.Kind != "DockerImage" {
			errs = append(errs, fieldErrorOrInternal(fmt.Errorf("pod specs may only contain references to docker images, not %q", ref.Kind), m.path.Child("initContainers").Index(i).Child("image")))
			continue
		}
		container.Image = ref.Name
	}
	for i := range m.spec.Containers {
		container := &m.spec.Containers[i]
		if hasIdenticalPodSpecImage(m.oldSpec, container.Name, container.Image) {
			continue
		}
		ref := corev1.ObjectReference{Kind: "DockerImage", Name: container.Image}
		if err := fn(&ref); err != nil {
			errs = append(errs, fieldErrorOrInternal(err, m.path.Child("containers").Index(i).Child("image")))
			continue
		}
		if ref.Kind != "DockerImage" {
			errs = append(errs, fieldErrorOrInternal(fmt.Errorf("pod specs may only contain references to docker images, not %q", ref.Kind), m.path.Child("containers").Index(i).Child("image")))
			continue
		}
		container.Image = ref.Name
	}
	return errs
}
func (m *podSpecMutator) GetContainerByName(name string) (ContainerMutator, bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	spec := m.spec
	for i := range spec.InitContainers {
		if name != spec.InitContainers[i].Name {
			continue
		}
		return containerMutator{&spec.InitContainers[i]}, true
	}
	for i := range spec.Containers {
		if name != spec.Containers[i].Name {
			continue
		}
		return containerMutator{&spec.Containers[i]}, true
	}
	return nil, false
}
func (m *podSpecMutator) GetContainerByIndex(init bool, i int) (ContainerMutator, bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var container *kapi.Container
	spec := m.spec
	if init {
		if i < 0 || i >= len(spec.InitContainers) {
			return nil, false
		}
		container = &spec.InitContainers[i]
	} else {
		if i < 0 || i >= len(spec.Containers) {
			return nil, false
		}
		container = &spec.Containers[i]
	}
	return containerMutator{container}, true
}

type podSpecV1Mutator struct {
	spec	*corev1.PodSpec
	oldSpec	*corev1.PodSpec
	path	*field.Path
}

func (m *podSpecV1Mutator) Path() *field.Path {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return m.path
}
func hasIdenticalPodSpecV1Image(spec *corev1.PodSpec, containerName, image string) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if spec == nil {
		return false
	}
	for i := range spec.InitContainers {
		if spec.InitContainers[i].Name == containerName {
			return spec.InitContainers[i].Image == image
		}
	}
	for i := range spec.Containers {
		if spec.Containers[i].Name == containerName {
			return spec.Containers[i].Image == image
		}
	}
	return false
}
func (m *podSpecV1Mutator) Mutate(fn ImageReferenceMutateFunc) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var errs field.ErrorList
	for i := range m.spec.InitContainers {
		container := &m.spec.InitContainers[i]
		if hasIdenticalPodSpecV1Image(m.oldSpec, container.Name, container.Image) {
			continue
		}
		ref := corev1.ObjectReference{Kind: "DockerImage", Name: container.Image}
		if err := fn(&ref); err != nil {
			errs = append(errs, fieldErrorOrInternal(err, m.path.Child("initContainers").Index(i).Child("image")))
			continue
		}
		if ref.Kind != "DockerImage" {
			errs = append(errs, fieldErrorOrInternal(fmt.Errorf("pod specs may only contain references to docker images, not %q", ref.Kind), m.path.Child("initContainers").Index(i).Child("image")))
			continue
		}
		container.Image = ref.Name
	}
	for i := range m.spec.Containers {
		container := &m.spec.Containers[i]
		if hasIdenticalPodSpecV1Image(m.oldSpec, container.Name, container.Image) {
			continue
		}
		ref := corev1.ObjectReference{Kind: "DockerImage", Name: container.Image}
		if err := fn(&ref); err != nil {
			errs = append(errs, fieldErrorOrInternal(err, m.path.Child("containers").Index(i).Child("image")))
			continue
		}
		if ref.Kind != "DockerImage" {
			errs = append(errs, fieldErrorOrInternal(fmt.Errorf("pod specs may only contain references to docker images, not %q", ref.Kind), m.path.Child("containers").Index(i).Child("image")))
			continue
		}
		container.Image = ref.Name
	}
	return errs
}
func (m *podSpecV1Mutator) GetContainerByName(name string) (ContainerMutator, bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	spec := m.spec
	for i := range spec.InitContainers {
		if name != spec.InitContainers[i].Name {
			continue
		}
		return containerV1Mutator{&spec.InitContainers[i]}, true
	}
	for i := range spec.Containers {
		if name != spec.Containers[i].Name {
			continue
		}
		return containerV1Mutator{&spec.Containers[i]}, true
	}
	return nil, false
}
func (m *podSpecV1Mutator) GetContainerByIndex(init bool, i int) (ContainerMutator, bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var container *corev1.Container
	spec := m.spec
	if init {
		if i < 0 || i >= len(spec.InitContainers) {
			return nil, false
		}
		container = &spec.InitContainers[i]
	} else {
		if i < 0 || i >= len(spec.Containers) {
			return nil, false
		}
		container = &spec.Containers[i]
	}
	return containerV1Mutator{container}, true
}
