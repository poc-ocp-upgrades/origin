package util

import (
	"fmt"
	"reflect"
	kmeta "k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	appsv1 "github.com/openshift/api/apps/v1"
)

const (
	OverwriteExistingDstKey	= 1 << iota
	ErrorOnExistingDstKey
	ErrorOnDifferentDstKeyValue
)

func AddObjectLabelsWithFlags(obj runtime.Object, labels labels.Set, flags int) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if labels == nil {
		return nil
	}
	accessor, err := kmeta.Accessor(obj)
	if err != nil {
		if _, ok := obj.(*unstructured.Unstructured); !ok {
			return err
		}
	} else {
		metaLabels := accessor.GetLabels()
		if metaLabels == nil {
			metaLabels = make(map[string]string)
		}
		switch objType := obj.(type) {
		case *appsv1.DeploymentConfig:
			if err := addDeploymentConfigNestedLabels(objType, labels, flags); err != nil {
				return fmt.Errorf("unable to add nested labels to %s/%s: %v", obj.GetObjectKind().GroupVersionKind(), accessor.GetName(), err)
			}
		}
		if err := MergeInto(metaLabels, labels, flags); err != nil {
			return fmt.Errorf("unable to add labels to %s/%s: %v", obj.GetObjectKind().GroupVersionKind(), accessor.GetName(), err)
		}
		accessor.SetLabels(metaLabels)
		return nil
	}
	if unstruct, ok := obj.(*unstructured.Unstructured); ok && unstruct.Object != nil {
		if obj, ok := unstruct.Object["metadata"]; ok {
			if m, ok := obj.(map[string]interface{}); ok {
				existing := make(map[string]string)
				if l, ok := m["labels"]; ok {
					if found, ok := interfaceToStringMap(l); ok {
						existing = found
					}
				}
				if err := MergeInto(existing, labels, flags); err != nil {
					return err
				}
				m["labels"] = mapToGeneric(existing)
			}
			return nil
		}
		if obj, ok := unstruct.Object["labels"]; ok {
			existing := make(map[string]string)
			if found, ok := interfaceToStringMap(obj); ok {
				existing = found
			}
			if err := MergeInto(existing, labels, flags); err != nil {
				return err
			}
			unstruct.Object["labels"] = mapToGeneric(existing)
			return nil
		}
	}
	return nil
}
func AddObjectLabels(obj runtime.Object, labels labels.Set) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return AddObjectLabelsWithFlags(obj, labels, OverwriteExistingDstKey)
}
func AddObjectAnnotations(obj runtime.Object, annotations map[string]string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(annotations) == 0 {
		return nil
	}
	accessor, err := kmeta.Accessor(obj)
	if err != nil {
		if _, ok := obj.(*unstructured.Unstructured); !ok {
			return err
		}
	} else {
		metaAnnotations := accessor.GetAnnotations()
		if metaAnnotations == nil {
			metaAnnotations = make(map[string]string)
		}
		switch objType := obj.(type) {
		case *appsv1.DeploymentConfig:
			if err := addDeploymentConfigNestedAnnotations(objType, annotations); err != nil {
				return fmt.Errorf("unable to add nested annotations to %s/%s: %v", obj.GetObjectKind().GroupVersionKind(), accessor.GetName(), err)
			}
		}
		MergeInto(metaAnnotations, annotations, OverwriteExistingDstKey)
		accessor.SetAnnotations(metaAnnotations)
		return nil
	}
	if unstruct, ok := obj.(*unstructured.Unstructured); ok && unstruct.Object != nil {
		if obj, ok := unstruct.Object["metadata"]; ok {
			if m, ok := obj.(map[string]interface{}); ok {
				existing := make(map[string]string)
				if l, ok := m["annotations"]; ok {
					if found, ok := interfaceToStringMap(l); ok {
						existing = found
					}
				}
				if err := MergeInto(existing, annotations, OverwriteExistingDstKey); err != nil {
					return err
				}
				m["annotations"] = mapToGeneric(existing)
			}
			return nil
		}
		if obj, ok := unstruct.Object["annotations"]; ok {
			existing := make(map[string]string)
			if found, ok := interfaceToStringMap(obj); ok {
				existing = found
			}
			if err := MergeInto(existing, annotations, OverwriteExistingDstKey); err != nil {
				return err
			}
			unstruct.Object["annotations"] = mapToGeneric(existing)
			return nil
		}
	}
	return nil
}
func addDeploymentConfigNestedLabels(obj *appsv1.DeploymentConfig, labels labels.Set, flags int) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if obj.Spec.Template == nil {
		return nil
	}
	if obj.Spec.Template.Labels == nil {
		obj.Spec.Template.Labels = make(map[string]string)
	}
	if err := MergeInto(obj.Spec.Template.Labels, labels, flags); err != nil {
		return fmt.Errorf("unable to add labels to Template.DeploymentConfig.Template.ControllerTemplate.Template: %v", err)
	}
	return nil
}
func addDeploymentConfigNestedAnnotations(obj *appsv1.DeploymentConfig, annotations map[string]string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if obj.Spec.Template == nil {
		return nil
	}
	if obj.Spec.Template.Annotations == nil {
		obj.Spec.Template.Annotations = make(map[string]string)
	}
	if err := MergeInto(obj.Spec.Template.Annotations, annotations, OverwriteExistingDstKey); err != nil {
		return fmt.Errorf("unable to add annotations to Template.DeploymentConfig.Template.ControllerTemplate.Template: %v", err)
	}
	return nil
}
func interfaceToStringMap(obj interface{}) (map[string]string, bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if obj == nil {
		return nil, false
	}
	lm, ok := obj.(map[string]interface{})
	if !ok {
		return nil, false
	}
	existing := make(map[string]string)
	for k, v := range lm {
		switch t := v.(type) {
		case string:
			existing[k] = t
		}
	}
	return existing, true
}
func mapToGeneric(obj map[string]string) map[string]interface{} {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if obj == nil {
		return nil
	}
	res := make(map[string]interface{})
	for k, v := range obj {
		res[k] = v
	}
	return res
}
func MergeInto(dst, src interface{}, flags int) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	dstVal := reflect.ValueOf(dst)
	srcVal := reflect.ValueOf(src)
	if dstVal.Kind() != reflect.Map {
		return fmt.Errorf("dst is not a valid map: %v", dstVal.Kind())
	}
	if srcVal.Kind() != reflect.Map {
		return fmt.Errorf("src is not a valid map: %v", srcVal.Kind())
	}
	if dstTyp, srcTyp := dstVal.Type(), srcVal.Type(); !dstTyp.AssignableTo(srcTyp) {
		return fmt.Errorf("type mismatch, can't assign '%v' to '%v'", srcTyp, dstTyp)
	}
	if dstVal.IsNil() {
		return fmt.Errorf("dst value is nil")
	}
	if srcVal.IsNil() {
		return nil
	}
	for _, k := range srcVal.MapKeys() {
		if dstVal.MapIndex(k).IsValid() {
			if flags&ErrorOnExistingDstKey != 0 {
				return fmt.Errorf("dst key already set (ErrorOnExistingDstKey=1), '%v'='%v'", k, dstVal.MapIndex(k))
			}
			if dstVal.MapIndex(k).String() != srcVal.MapIndex(k).String() {
				if flags&ErrorOnDifferentDstKeyValue != 0 {
					return fmt.Errorf("dst key already set to a different value (ErrorOnDifferentDstKeyValue=1), '%v'='%v'", k, dstVal.MapIndex(k))
				}
				if flags&OverwriteExistingDstKey != 0 {
					dstVal.SetMapIndex(k, srcVal.MapIndex(k))
				}
			}
		} else {
			dstVal.SetMapIndex(k, srcVal.MapIndex(k))
		}
	}
	return nil
}
