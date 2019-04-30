package imagereferencemutators

import (
	"fmt"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	buildv1 "github.com/openshift/api/build/v1"
)

type ImageReferenceMutateFunc func(ref *corev1.ObjectReference) error
type ImageReferenceMutator interface {
	Mutate(fn ImageReferenceMutateFunc) field.ErrorList
}

var errNoImageMutator = fmt.Errorf("no list of images available for this object")

func GetImageReferenceMutator(obj, old runtime.Object) (ImageReferenceMutator, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch t := obj.(type) {
	case *buildv1.Build:
		if oldT, ok := old.(*buildv1.Build); ok && oldT != nil {
			return &buildSpecMutator{spec: &t.Spec.CommonSpec, oldSpec: &oldT.Spec.CommonSpec, path: field.NewPath("spec")}, nil
		}
		return &buildSpecMutator{spec: &t.Spec.CommonSpec, path: field.NewPath("spec")}, nil
	case *buildv1.BuildConfig:
		if oldT, ok := old.(*buildv1.BuildConfig); ok && oldT != nil {
			return &buildSpecMutator{spec: &t.Spec.CommonSpec, oldSpec: &oldT.Spec.CommonSpec, path: field.NewPath("spec")}, nil
		}
		return &buildSpecMutator{spec: &t.Spec.CommonSpec, path: field.NewPath("spec")}, nil
	default:
		if spec, path, err := GetPodSpec(obj); err == nil {
			if old == nil {
				return &podSpecMutator{spec: spec, path: path}, nil
			}
			oldSpec, _, err := GetPodSpec(old)
			if err != nil {
				return nil, fmt.Errorf("old and new pod spec objects were not of the same type %T != %T: %v", obj, old, err)
			}
			return &podSpecMutator{spec: spec, oldSpec: oldSpec, path: path}, nil
		}
		if spec, path, err := GetPodSpecV1(obj); err == nil {
			if old == nil {
				return &podSpecV1Mutator{spec: spec, path: path}, nil
			}
			oldSpec, _, err := GetPodSpecV1(old)
			if err != nil {
				return nil, fmt.Errorf("old and new pod spec objects were not of the same type %T != %T: %v", obj, old, err)
			}
			return &podSpecV1Mutator{spec: spec, oldSpec: oldSpec, path: path}, nil
		}
		return nil, errNoImageMutator
	}
}

type AnnotationAccessor interface {
	Annotations() map[string]string
	SetAnnotations(map[string]string)
	TemplateAnnotations() (map[string]string, bool)
	SetTemplateAnnotations(map[string]string) bool
}
type annotationsAccessor struct {
	object		metav1.Object
	template	metav1.Object
}

func (a annotationsAccessor) Annotations() map[string]string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return a.object.GetAnnotations()
}
func (a annotationsAccessor) TemplateAnnotations() (map[string]string, bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if a.template == nil {
		return nil, false
	}
	return a.template.GetAnnotations(), true
}
func (a annotationsAccessor) SetAnnotations(annotations map[string]string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	a.object.SetAnnotations(annotations)
}
func (a annotationsAccessor) SetTemplateAnnotations(annotations map[string]string) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if a.template == nil {
		return false
	}
	a.template.SetAnnotations(annotations)
	return true
}
func GetAnnotationAccessor(obj runtime.Object) (AnnotationAccessor, bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch t := obj.(type) {
	case metav1.Object:
		templateObject, _ := GetTemplateMetaObject(obj)
		return annotationsAccessor{object: t, template: templateObject}, true
	default:
		return nil, false
	}
}
