package util

import (
	"strconv"
	"strings"
	appsv1 "github.com/openshift/api/apps/v1"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
)

var (
	deploymentConfigControllerRefKind = appsv1.GroupVersion.WithKind("DeploymentConfig")
)

func newControllerRef(config *appsv1.DeploymentConfig) *metav1.OwnerReference {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	blockOwnerDeletion := true
	isController := true
	return &metav1.OwnerReference{APIVersion: deploymentConfigControllerRefKind.GroupVersion().String(), Kind: deploymentConfigControllerRefKind.Kind, Name: config.Name, UID: config.UID, BlockOwnerDeletion: &blockOwnerDeletion, Controller: &isController}
}
func MakeDeployment(config *appsv1.DeploymentConfig) (*v1.ReplicationController, error) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	encodedConfig, err := runtime.Encode(annotationEncoder, config)
	if err != nil {
		return nil, err
	}
	deploymentName := LatestDeploymentNameForConfig(config)
	podSpec := config.Spec.Template.Spec.DeepCopy()
	for i := range podSpec.Containers {
		podSpec.Containers[i].Image = strings.TrimSpace(podSpec.Containers[i].Image)
	}
	controllerLabels := make(labels.Set)
	for k, v := range config.Labels {
		controllerLabels[k] = v
	}
	controllerLabels[appsv1.DeploymentConfigAnnotation] = config.Name
	selector := map[string]string{}
	for k, v := range config.Spec.Selector {
		selector[k] = v
	}
	selector[DeploymentConfigLabel] = config.Name
	selector[DeploymentLabel] = deploymentName
	podLabels := make(labels.Set)
	for k, v := range config.Spec.Template.Labels {
		podLabels[k] = v
	}
	podLabels[DeploymentConfigLabel] = config.Name
	podLabels[DeploymentLabel] = deploymentName
	podAnnotations := make(labels.Set)
	for k, v := range config.Spec.Template.Annotations {
		podAnnotations[k] = v
	}
	podAnnotations[appsv1.DeploymentAnnotation] = deploymentName
	podAnnotations[appsv1.DeploymentConfigAnnotation] = config.Name
	podAnnotations[appsv1.DeploymentVersionAnnotation] = strconv.FormatInt(config.Status.LatestVersion, 10)
	controllerRef := newControllerRef(config)
	zero := int32(0)
	deployment := &v1.ReplicationController{ObjectMeta: metav1.ObjectMeta{Name: deploymentName, Namespace: config.Namespace, Annotations: map[string]string{appsv1.DeploymentConfigAnnotation: config.Name, appsv1.DeploymentEncodedConfigAnnotation: string(encodedConfig), appsv1.DeploymentStatusAnnotation: string(appsv1.DeploymentStatusNew), appsv1.DeploymentVersionAnnotation: strconv.FormatInt(config.Status.LatestVersion, 10), appsv1.DesiredReplicasAnnotation: strconv.Itoa(int(config.Spec.Replicas)), DeploymentReplicasAnnotation: strconv.Itoa(0)}, Labels: controllerLabels, OwnerReferences: []metav1.OwnerReference{*controllerRef}}, Spec: v1.ReplicationControllerSpec{Replicas: &zero, Selector: selector, MinReadySeconds: config.Spec.MinReadySeconds, Template: &v1.PodTemplateSpec{ObjectMeta: metav1.ObjectMeta{Labels: podLabels, Annotations: podAnnotations}, Spec: *podSpec}}}
	if config.Status.Details != nil && len(config.Status.Details.Message) > 0 {
		deployment.Annotations[appsv1.DeploymentStatusReasonAnnotation] = config.Status.Details.Message
	}
	if value, ok := config.Annotations[DeploymentIgnorePodAnnotation]; ok {
		deployment.Annotations[DeploymentIgnorePodAnnotation] = value
	}
	return deployment, nil
}
