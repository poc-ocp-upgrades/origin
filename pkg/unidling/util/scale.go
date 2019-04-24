package util

import (
	"fmt"
	"bytes"
	"net/http"
	"runtime"
	"k8s.io/klog"
	"github.com/openshift/origin/pkg/api/legacy"
	autoscalingv1 "k8s.io/api/autoscaling/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/strategicpatch"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/scale"
	appsv1 "github.com/openshift/api/apps/v1"
	appsclient "github.com/openshift/client-go/apps/clientset/versioned/typed/apps/v1"
	unidlingapi "github.com/openshift/origin/pkg/unidling/api"
)

type AnnotationFunc func(currentReplicas int32, annotations map[string]string)

func NewScaleAnnotater(scales scale.ScalesGetter, mapper meta.RESTMapper, dcs appsclient.DeploymentConfigsGetter, rcs corev1client.ReplicationControllersGetter, changeAnnots AnnotationFunc) *ScaleAnnotater {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &ScaleAnnotater{mapper: mapper, scales: scales, dcs: dcs, rcs: rcs, ChangeAnnotations: changeAnnots}
}

type ScaleAnnotater struct {
	mapper			meta.RESTMapper
	scales			scale.ScalesGetter
	dcs			appsclient.DeploymentConfigsGetter
	rcs			corev1client.ReplicationControllersGetter
	ChangeAnnotations	AnnotationFunc
}
type ScaleUpdater interface {
	Update(*ScaleAnnotater, runtime.Object, *autoscalingv1.Scale) error
}
type scaleUpdater struct {
	encoder		runtime.Encoder
	namespace	string
	dcGetter	appsclient.DeploymentConfigsGetter
	rcGetter	corev1client.ReplicationControllersGetter
}

func NewScaleUpdater(encoder runtime.Encoder, namespace string, dcGetter appsclient.DeploymentConfigsGetter, rcGetter corev1client.ReplicationControllersGetter) ScaleUpdater {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return scaleUpdater{encoder: encoder, namespace: namespace, dcGetter: dcGetter, rcGetter: rcGetter}
}
func (s scaleUpdater) Update(annotator *ScaleAnnotater, obj runtime.Object, scale *autoscalingv1.Scale) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var (
		err				error
		patchBytes, originalObj, newObj	[]byte
	)
	originalObj, err = runtime.Encode(s.encoder, obj)
	if err != nil {
		return err
	}
	switch typedObj := obj.(type) {
	case *appsv1.DeploymentConfig:
		if typedObj.Annotations == nil {
			typedObj.Annotations = make(map[string]string)
		}
		annotator.ChangeAnnotations(typedObj.Spec.Replicas, typedObj.Annotations)
		typedObj.Spec.Replicas = scale.Spec.Replicas
		newObj, err = runtime.Encode(s.encoder, typedObj)
		if err != nil {
			return err
		}
		patchBytes, err = strategicpatch.CreateTwoWayMergePatch(originalObj, newObj, &appsv1.DeploymentConfig{})
		if err != nil {
			return err
		}
		_, err = s.dcGetter.DeploymentConfigs(s.namespace).Patch(typedObj.Name, types.StrategicMergePatchType, patchBytes)
	case *corev1.ReplicationController:
		if typedObj.Annotations == nil {
			typedObj.Annotations = make(map[string]string)
		}
		annotator.ChangeAnnotations(*typedObj.Spec.Replicas, typedObj.Annotations)
		typedObj.Spec.Replicas = &scale.Spec.Replicas
		newObj, err = runtime.Encode(s.encoder, typedObj)
		if err != nil {
			return err
		}
		patchBytes, err = strategicpatch.CreateTwoWayMergePatch(originalObj, newObj, &corev1.ReplicationController{})
		if err != nil {
			return err
		}
		_, err = s.rcGetter.ReplicationControllers(s.namespace).Patch(typedObj.Name, types.StrategicMergePatchType, patchBytes)
	}
	return err
}
func (c *ScaleAnnotater) GetObjectWithScale(namespace string, ref unidlingapi.CrossGroupObjectReference) (runtime.Object, *autoscalingv1.Scale, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var obj runtime.Object
	var err error
	var scale *autoscalingv1.Scale
	switch {
	case ref.Kind == "DeploymentConfig" && (ref.Group == appsv1.GroupName || ref.Group == legacy.GroupName):
		var dc *appsv1.DeploymentConfig
		dc, err = c.dcs.DeploymentConfigs(namespace).Get(ref.Name, metav1.GetOptions{})
		if err != nil {
			return nil, nil, err
		}
		obj = dc
	case ref.Kind == "ReplicationController" && ref.Group == corev1.GroupName:
		var rc *corev1.ReplicationController
		rc, err = c.rcs.ReplicationControllers(namespace).Get(ref.Name, metav1.GetOptions{})
		if err != nil {
			return nil, nil, err
		}
		obj = rc
	}
	mappings, err := c.mapper.RESTMappings(schema.GroupKind{Group: ref.Group, Kind: ref.Kind})
	if err != nil {
		return nil, nil, err
	}
	for _, mapping := range mappings {
		scale, err = c.scales.Scales(namespace).Get(mapping.Resource.GroupResource(), ref.Name)
		if err != nil {
			return nil, nil, err
		}
	}
	return obj, scale, err
}
func (c *ScaleAnnotater) UpdateObjectScale(updater ScaleUpdater, namespace string, ref unidlingapi.CrossGroupObjectReference, obj runtime.Object, scale *autoscalingv1.Scale) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var err error
	mappings, err := c.mapper.RESTMappings(schema.GroupKind{Group: ref.Group, Kind: ref.Kind})
	if err != nil {
		return err
	}
	if len(mappings) == 0 {
		return fmt.Errorf("cannot locate resource for %s.%s/%s", ref.Kind, ref.Group, ref.Name)
	}
	for _, mapping := range mappings {
		if obj == nil {
			_, err = c.scales.Scales(namespace).Update(mapping.Resource.GroupResource(), scale)
			return err
		}
		switch obj.(type) {
		case *appsv1.DeploymentConfig, *corev1.ReplicationController:
			return updater.Update(c, obj, scale)
		default:
			klog.V(2).Infof("Unidling unknown type %t: using scale interface and not removing annotations", obj)
			_, err = c.scales.Scales(namespace).Update(mapping.Resource.GroupResource(), scale)
		}
	}
	return err
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
