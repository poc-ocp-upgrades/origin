package util

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/diff"
	intstrutil "k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	scaleclient "k8s.io/client-go/scale"
	autoscalingv1 "k8s.io/kubernetes/pkg/apis/autoscaling/v1"
	kapiv1 "k8s.io/kubernetes/pkg/apis/core/v1"
	"k8s.io/kubernetes/pkg/kubectl"
	appsv1 "github.com/openshift/api/apps/v1"
	"github.com/openshift/origin/pkg/api/apihelpers"
)

type rcMapper struct{}

func (rcMapper) ResourceFor(gvr schema.GroupVersionResource) (schema.GroupVersionResource, error) {
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
	if gvr.Group == "" && gvr.Resource == "replicationcontrollers" {
		return kapiv1.SchemeGroupVersion.WithResource("replicationcontrollers"), nil
	}
	return schema.GroupVersionResource{}, fmt.Errorf("unknown replication controller resource: %#v", gvr)
}
func (rcMapper) ScaleForResource(gvr schema.GroupVersionResource) (schema.GroupVersionKind, error) {
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
	if gvr == kapiv1.SchemeGroupVersion.WithResource("replicationcontrollers") {
		return autoscalingv1.SchemeGroupVersion.WithKind("Scale"), nil
	}
	return schema.GroupVersionKind{}, fmt.Errorf("unknown replication controller resource: %#v", gvr)
}
func DecodeDeploymentConfig(controller metav1.ObjectMetaAccessor) (*appsv1.DeploymentConfig, error) {
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
	encodedConfig, exists := controller.GetObjectMeta().GetAnnotations()[appsv1.DeploymentEncodedConfigAnnotation]
	if !exists {
		return nil, fmt.Errorf("object %s does not have encoded deployment config annotation", controller.GetObjectMeta().GetName())
	}
	config, err := runtime.Decode(annotationDecoder, []byte(encodedConfig))
	if err != nil {
		return nil, err
	}
	externalConfig, ok := config.(*appsv1.DeploymentConfig)
	if !ok {
		return nil, fmt.Errorf("object %+v is not v1.DeploymentConfig", config)
	}
	return externalConfig, nil
}
func RolloutExceededTimeoutSeconds(config *appsv1.DeploymentConfig, latestRC *v1.ReplicationController) bool {
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
	timeoutSeconds := GetTimeoutSecondsForStrategy(config)
	if timeoutSeconds <= 0 {
		return false
	}
	return int64(time.Since(latestRC.CreationTimestamp.Time).Seconds()) > timeoutSeconds
}
func GetTimeoutSecondsForStrategy(config *appsv1.DeploymentConfig) int64 {
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
	var timeoutSeconds int64
	switch config.Spec.Strategy.Type {
	case appsv1.DeploymentStrategyTypeRolling:
		timeoutSeconds = DefaultRollingTimeoutSeconds
		if t := config.Spec.Strategy.RollingParams.TimeoutSeconds; t != nil {
			timeoutSeconds = *t
		}
	case appsv1.DeploymentStrategyTypeRecreate:
		timeoutSeconds = DefaultRecreateTimeoutSeconds
		if t := config.Spec.Strategy.RecreateParams.TimeoutSeconds; t != nil {
			timeoutSeconds = *t
		}
	case appsv1.DeploymentStrategyTypeCustom:
		timeoutSeconds = DefaultRecreateTimeoutSeconds
	}
	return timeoutSeconds
}
func NewReplicationControllerScaler(client kubernetes.Interface) kubectl.Scaler {
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
	return kubectl.NewScaler(NewReplicationControllerScaleClient(client))
}
func NewReplicationControllerScaleClient(client kubernetes.Interface) scaleclient.ScalesGetter {
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
	return scaleclient.New(client.CoreV1().RESTClient(), rcMapper{}, dynamic.LegacyAPIPathResolverFunc, rcMapper{})
}
func DeployerPodNameForDeployment(deployment string) string {
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
	return apihelpers.GetPodName(deployment, "deploy")
}
func NewDeploymentCondition(condType appsv1.DeploymentConditionType, status v1.ConditionStatus, reason string, message string) *appsv1.DeploymentCondition {
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
	return &appsv1.DeploymentCondition{Type: condType, Status: status, LastUpdateTime: metav1.Now(), LastTransitionTime: metav1.Now(), Reason: reason, Message: message}
}
func SetDeploymentCondition(status *appsv1.DeploymentConfigStatus, condition appsv1.DeploymentCondition) {
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
	currentCond := GetDeploymentCondition(*status, condition.Type)
	if currentCond != nil && currentCond.Status == condition.Status && currentCond.Reason == condition.Reason {
		return
	}
	if currentCond != nil && currentCond.Status == condition.Status {
		condition.LastTransitionTime = currentCond.LastTransitionTime
	}
	newConditions := filterOutCondition(status.Conditions, condition.Type)
	status.Conditions = append(newConditions, condition)
}
func RemoveDeploymentCondition(status *appsv1.DeploymentConfigStatus, condType appsv1.DeploymentConditionType) {
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
	status.Conditions = filterOutCondition(status.Conditions, condType)
}
func filterOutCondition(conditions []appsv1.DeploymentCondition, condType appsv1.DeploymentConditionType) []appsv1.DeploymentCondition {
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
	var newConditions []appsv1.DeploymentCondition
	for _, c := range conditions {
		if c.Type == condType {
			continue
		}
		newConditions = append(newConditions, c)
	}
	return newConditions
}
func IsOwnedByConfig(obj metav1.Object) bool {
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
	_, ok := obj.GetAnnotations()[appsv1.DeploymentConfigAnnotation]
	return ok
}
func DeploymentsForCleanup(configuration *appsv1.DeploymentConfig, deployments []*v1.ReplicationController) []v1.ReplicationController {
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
	sort.Sort(sort.Reverse(ByLatestVersionDesc(deployments)))
	relevantDeployments := []v1.ReplicationController{}
	activeDeployment := ActiveDeployment(deployments)
	if activeDeployment == nil {
		for i := range deployments {
			deployment := deployments[i]
			if deploymentVersionFor(deployment) != configuration.Status.LatestVersion {
				relevantDeployments = append(relevantDeployments, *deployment)
			}
		}
	} else {
		for i := range deployments {
			deployment := deployments[i]
			if deployment != activeDeployment && deploymentVersionFor(deployment) < deploymentVersionFor(activeDeployment) {
				relevantDeployments = append(relevantDeployments, *deployment)
			}
		}
	}
	return relevantDeployments
}
func RecordConfigChangeCause(config *appsv1.DeploymentConfig) {
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
	config.Status.Details = &appsv1.DeploymentDetails{Causes: []appsv1.DeploymentCause{{Type: appsv1.DeploymentTriggerOnConfigChange}}, Message: "config change"}
}
func RecordImageChangeCauses(config *appsv1.DeploymentConfig, imageNames []string) {
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
	config.Status.Details = &appsv1.DeploymentDetails{Message: "image change"}
	for _, imageName := range imageNames {
		config.Status.Details.Causes = append(config.Status.Details.Causes, appsv1.DeploymentCause{Type: appsv1.DeploymentTriggerOnImageChange, ImageTrigger: &appsv1.DeploymentCauseImageTrigger{From: v1.ObjectReference{Kind: "DockerImage", Name: imageName}}})
	}
}
func HasUpdatedImages(dc *appsv1.DeploymentConfig, rc *v1.ReplicationController) (bool, []string) {
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
	updatedImages := []string{}
	rcImages := sets.NewString()
	for _, c := range rc.Spec.Template.Spec.Containers {
		rcImages.Insert(c.Image)
	}
	for _, c := range dc.Spec.Template.Spec.Containers {
		if !rcImages.Has(c.Image) {
			updatedImages = append(updatedImages, c.Image)
		}
	}
	if len(updatedImages) == 0 {
		return false, nil
	}
	return true, updatedImages
}
func HasLatestPodTemplate(currentConfig *appsv1.DeploymentConfig, rc *v1.ReplicationController) (bool, string, error) {
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
	latestConfig, err := DecodeDeploymentConfig(rc)
	if err != nil {
		return true, "", err
	}
	if reflect.DeepEqual(currentConfig.Spec.Template, latestConfig.Spec.Template) {
		return true, "", nil
	}
	return false, diff.ObjectReflectDiff(currentConfig.Spec.Template, latestConfig.Spec.Template), nil
}
func IsProgressing(config *appsv1.DeploymentConfig, newStatus *appsv1.DeploymentConfigStatus) bool {
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
	oldStatusOldReplicas := config.Status.Replicas - config.Status.UpdatedReplicas
	newStatusOldReplicas := newStatus.Replicas - newStatus.UpdatedReplicas
	return (newStatus.UpdatedReplicas > config.Status.UpdatedReplicas) || (newStatusOldReplicas < oldStatusOldReplicas)
}
func LabelForDeployment(deployment *v1.ReplicationController) string {
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
	return fmt.Sprintf("%s/%s", deployment.Namespace, deployment.Name)
}
func LabelForDeploymentConfig(config runtime.Object) string {
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
	accessor, _ := meta.Accessor(config)
	return fmt.Sprintf("%s/%s", accessor.GetNamespace(), accessor.GetName())
}
func DeploymentDesiredReplicas(obj runtime.Object) (int32, bool) {
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
	return int32AnnotationFor(obj, appsv1.DesiredReplicasAnnotation)
}
func LatestDeploymentNameForConfig(config *appsv1.DeploymentConfig) string {
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
	return LatestDeploymentNameForConfigAndVersion(config.Name, config.Status.LatestVersion)
}
func DeploymentNameForConfigVersion(name string, version int64) string {
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
	return fmt.Sprintf("%s-%d", name, version)
}
func LatestDeploymentNameForConfigAndVersion(name string, version int64) string {
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
	return fmt.Sprintf("%s-%d", name, version)
}
func DeployerPodNameFor(obj runtime.Object) string {
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
	return AnnotationFor(obj, appsv1.DeploymentPodAnnotation)
}
func DeploymentConfigNameFor(obj runtime.Object) string {
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
	return AnnotationFor(obj, appsv1.DeploymentConfigAnnotation)
}
func DeploymentStatusReasonFor(obj runtime.Object) string {
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
	return AnnotationFor(obj, appsv1.DeploymentStatusReasonAnnotation)
}
func DeleteStatusReasons(rc *v1.ReplicationController) {
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
	delete(rc.Annotations, appsv1.DeploymentStatusReasonAnnotation)
	delete(rc.Annotations, appsv1.DeploymentCancelledAnnotation)
}
func SetCancelledByUserReason(rc *v1.ReplicationController) {
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
	rc.Annotations[appsv1.DeploymentCancelledAnnotation] = "true"
	rc.Annotations[appsv1.DeploymentStatusReasonAnnotation] = deploymentCancelledByUser
}
func SetCancelledByNewerDeployment(rc *v1.ReplicationController) {
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
	rc.Annotations[appsv1.DeploymentCancelledAnnotation] = "true"
	rc.Annotations[appsv1.DeploymentStatusReasonAnnotation] = deploymentCancelledNewerDeploymentExists
}
func HasSynced(dc *appsv1.DeploymentConfig, generation int64) bool {
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
	return dc.Status.ObservedGeneration >= generation
}
func HasChangeTrigger(config *appsv1.DeploymentConfig) bool {
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
	for _, trigger := range config.Spec.Triggers {
		if trigger.Type == appsv1.DeploymentTriggerOnConfigChange {
			return true
		}
	}
	return false
}
func HasTrigger(config *appsv1.DeploymentConfig) bool {
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
	return HasChangeTrigger(config) || HasImageChangeTrigger(config)
}
func HasLastTriggeredImage(config *appsv1.DeploymentConfig) bool {
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
	hasImageTrigger := false
	for _, trigger := range config.Spec.Triggers {
		if trigger.Type == appsv1.DeploymentTriggerOnImageChange {
			hasImageTrigger = true
			if len(trigger.ImageChangeParams.LastTriggeredImage) == 0 {
				return false
			}
		}
	}
	return hasImageTrigger
}
func IsInitialDeployment(config *appsv1.DeploymentConfig) bool {
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
	return config.Status.LatestVersion == 0
}
func IsRollingConfig(config *appsv1.DeploymentConfig) bool {
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
	return config.Spec.Strategy.Type == appsv1.DeploymentStrategyTypeRolling
}
func ResolveFenceposts(maxSurge, maxUnavailable *intstrutil.IntOrString, desired int32) (int32, int32, error) {
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
	surge, err := intstrutil.GetValueFromIntOrPercent(maxSurge, int(desired), true)
	if err != nil {
		return 0, 0, err
	}
	unavailable, err := intstrutil.GetValueFromIntOrPercent(maxUnavailable, int(desired), false)
	if err != nil {
		return 0, 0, err
	}
	if surge == 0 && unavailable == 0 {
		unavailable = 1
	}
	return int32(surge), int32(unavailable), nil
}
func MaxUnavailable(config *appsv1.DeploymentConfig) int32 {
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
	if !IsRollingConfig(config) {
		return int32(0)
	}
	_, maxUnavailable, _ := ResolveFenceposts(config.Spec.Strategy.RollingParams.MaxSurge, config.Spec.Strategy.RollingParams.MaxUnavailable, config.Spec.Replicas)
	return maxUnavailable
}
func MaxSurge(config appsv1.DeploymentConfig) int32 {
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
	if !IsRollingConfig(&config) {
		return int32(0)
	}
	maxSurge, _, _ := ResolveFenceposts(config.Spec.Strategy.RollingParams.MaxSurge, config.Spec.Strategy.RollingParams.MaxUnavailable, config.Spec.Replicas)
	return maxSurge
}
func AnnotationFor(obj runtime.Object, key string) string {
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
	objectMeta, err := meta.Accessor(obj)
	if err != nil {
		return ""
	}
	if objectMeta == nil || reflect.ValueOf(objectMeta).IsNil() {
		return ""
	}
	return objectMeta.GetAnnotations()[key]
}
func ActiveDeployment(input []*v1.ReplicationController) *v1.ReplicationController {
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
	var activeDeployment *v1.ReplicationController
	var lastCompleteDeploymentVersion int64 = 0
	for i := range input {
		deployment := input[i]
		deploymentVersion := DeploymentVersionFor(deployment)
		if IsCompleteDeployment(deployment) && deploymentVersion > lastCompleteDeploymentVersion {
			activeDeployment = deployment
			lastCompleteDeploymentVersion = deploymentVersion
		}
	}
	return activeDeployment
}
func ConfigSelector(name string) labels.Selector {
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
	return labels.SelectorFromValidatedSet(labels.Set{appsv1.DeploymentConfigAnnotation: name})
}
func IsCompleteDeployment(deployment runtime.Object) bool {
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
	return DeploymentStatusFor(deployment) == appsv1.DeploymentStatusComplete
}
func IsFailedDeployment(deployment runtime.Object) bool {
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
	return DeploymentStatusFor(deployment) == appsv1.DeploymentStatusFailed
}
func IsTerminatedDeployment(deployment runtime.Object) bool {
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
	return IsCompleteDeployment(deployment) || IsFailedDeployment(deployment)
}
func IsDeploymentCancelled(deployment runtime.Object) bool {
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
	value := AnnotationFor(deployment, appsv1.DeploymentCancelledAnnotation)
	return strings.EqualFold(value, "true")
}
func DeployerPodSelector(name string) labels.Selector {
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
	return labels.SelectorFromValidatedSet(labels.Set{appsv1.DeployerPodForDeploymentLabel: name})
}
func DeploymentStatusFor(deployment runtime.Object) appsv1.DeploymentStatus {
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
	return appsv1.DeploymentStatus(AnnotationFor(deployment, appsv1.DeploymentStatusAnnotation))
}
func SetDeploymentLatestVersionAnnotation(rc *v1.ReplicationController, version string) {
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
	if rc.Annotations == nil {
		rc.Annotations = map[string]string{}
	}
	rc.Annotations[appsv1.DeploymentVersionAnnotation] = version
}
func DeploymentVersionFor(obj runtime.Object) int64 {
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
	v, err := strconv.ParseInt(AnnotationFor(obj, appsv1.DeploymentVersionAnnotation), 10, 64)
	if err != nil {
		return -1
	}
	return v
}
func DeploymentNameFor(obj runtime.Object) string {
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
	return AnnotationFor(obj, appsv1.DeploymentAnnotation)
}
func deploymentVersionFor(obj runtime.Object) int64 {
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
	v, err := strconv.ParseInt(AnnotationFor(obj, appsv1.DeploymentVersionAnnotation), 10, 64)
	if err != nil {
		return -1
	}
	return v
}
func LatestDeploymentInfo(config *appsv1.DeploymentConfig, deployments []*v1.ReplicationController) (bool, *v1.ReplicationController) {
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
	if config.Status.LatestVersion == 0 || len(deployments) == 0 {
		return false, nil
	}
	sort.Sort(ByLatestVersionDesc(deployments))
	candidate := deployments[0]
	return deploymentVersionFor(candidate) == config.Status.LatestVersion, candidate
}
func GetDeploymentCondition(status appsv1.DeploymentConfigStatus, condType appsv1.DeploymentConditionType) *appsv1.DeploymentCondition {
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
	for i := range status.Conditions {
		c := status.Conditions[i]
		if c.Type == condType {
			return &c
		}
	}
	return nil
}
func GetReplicaCountForDeployments(deployments []*v1.ReplicationController) int32 {
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
	totalReplicaCount := int32(0)
	for _, deployment := range deployments {
		count := deployment.Spec.Replicas
		if count == nil {
			continue
		}
		totalReplicaCount += *count
	}
	return totalReplicaCount
}
func GetStatusReplicaCountForDeployments(deployments []*v1.ReplicationController) int32 {
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
	totalReplicaCount := int32(0)
	for _, deployment := range deployments {
		totalReplicaCount += deployment.Status.Replicas
	}
	return totalReplicaCount
}
func GetReadyReplicaCountForReplicationControllers(replicationControllers []*v1.ReplicationController) int32 {
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
	totalReadyReplicas := int32(0)
	for _, rc := range replicationControllers {
		if rc != nil {
			totalReadyReplicas += rc.Status.ReadyReplicas
		}
	}
	return totalReadyReplicas
}
func GetAvailableReplicaCountForReplicationControllers(replicationControllers []*v1.ReplicationController) int32 {
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
	totalAvailableReplicas := int32(0)
	for _, rc := range replicationControllers {
		if rc != nil {
			totalAvailableReplicas += rc.Status.AvailableReplicas
		}
	}
	return totalAvailableReplicas
}
func HasImageChangeTrigger(config *appsv1.DeploymentConfig) bool {
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
	for _, trigger := range config.Spec.Triggers {
		if trigger.Type == appsv1.DeploymentTriggerOnImageChange {
			return true
		}
	}
	return false
}
func CanTransitionPhase(current, next appsv1.DeploymentStatus) bool {
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
	switch current {
	case appsv1.DeploymentStatusNew:
		switch next {
		case appsv1.DeploymentStatusPending, appsv1.DeploymentStatusRunning, appsv1.DeploymentStatusFailed, appsv1.DeploymentStatusComplete:
			return true
		}
	case appsv1.DeploymentStatusPending:
		switch next {
		case appsv1.DeploymentStatusRunning, appsv1.DeploymentStatusFailed, appsv1.DeploymentStatusComplete:
			return true
		}
	case appsv1.DeploymentStatusRunning:
		switch next {
		case appsv1.DeploymentStatusFailed, appsv1.DeploymentStatusComplete:
			return true
		}
	}
	return false
}
func int32AnnotationFor(obj runtime.Object, key string) (int32, bool) {
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
	s := AnnotationFor(obj, key)
	if len(s) == 0 {
		return 0, false
	}
	i, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return 0, false
	}
	return int32(i), true
}

type ByLatestVersionAsc []*v1.ReplicationController

func (d ByLatestVersionAsc) Len() int {
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
	return len(d)
}
func (d ByLatestVersionAsc) Swap(i, j int) {
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
	d[i], d[j] = d[j], d[i]
}
func (d ByLatestVersionAsc) Less(i, j int) bool {
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
	return DeploymentVersionFor(d[i]) < DeploymentVersionFor(d[j])
}

type ByLatestVersionDesc []*v1.ReplicationController

func (d ByLatestVersionDesc) Len() int {
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
	return len(d)
}
func (d ByLatestVersionDesc) Swap(i, j int) {
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
	d[i], d[j] = d[j], d[i]
}
func (d ByLatestVersionDesc) Less(i, j int) bool {
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
	return DeploymentVersionFor(d[j]) < DeploymentVersionFor(d[i])
}
