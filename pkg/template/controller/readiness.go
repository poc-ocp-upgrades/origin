package controller

import (
	"fmt"
	"strconv"
	kappsv1 "k8s.io/api/apps/v1"
	kappsv1beta1 "k8s.io/api/apps/v1beta1"
	kappsv1beta2 "k8s.io/api/apps/v1beta2"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	kextensionsv1beta1 "k8s.io/api/extensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	deploymentutil "k8s.io/kubernetes/pkg/controller/deployment/util"
	appsv1 "github.com/openshift/api/apps/v1"
	buildv1 "github.com/openshift/api/build/v1"
	routev1 "github.com/openshift/api/route/v1"
	buildv1client "github.com/openshift/client-go/build/clientset/versioned"
	"github.com/openshift/origin/pkg/api/legacy"
	appsutil "github.com/openshift/origin/pkg/apps/util"
	buildutil "github.com/openshift/origin/pkg/build/util"
)

var readinessScheme = runtime.NewScheme()

func init() {
	_logClusterCodePath()
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
	kappsv1.AddToScheme(readinessScheme)
	kappsv1beta1.AddToScheme(readinessScheme)
	kappsv1beta2.AddToScheme(readinessScheme)
	kextensionsv1beta1.AddToScheme(readinessScheme)
	batchv1.AddToScheme(readinessScheme)
	corev1.AddToScheme(readinessScheme)
	appsv1.Install(readinessScheme)
	buildv1.Install(readinessScheme)
	routev1.Install(readinessScheme)
}
func checkBuildReadiness(obj runtime.Object) (bool, bool, error) {
	_logClusterCodePath()
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
	b, ok := obj.(*buildv1.Build)
	if !ok {
		return false, false, fmt.Errorf("object %T is not v1.Build", obj)
	}
	ready := buildutil.IsTerminalPhase(b.Status.Phase) && b.Status.Phase == buildv1.BuildPhaseComplete
	failed := buildutil.IsTerminalPhase(b.Status.Phase) && b.Status.Phase != buildv1.BuildPhaseComplete
	return ready, failed, nil
}
func checkBuildConfigReadiness(oc buildv1client.Interface, obj runtime.Object) (bool, bool, error) {
	_logClusterCodePath()
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
	bc, ok := obj.(*buildv1.BuildConfig)
	if !ok {
		return false, false, fmt.Errorf("object %T is not v1.BuildConfig", obj)
	}
	builds, err := oc.BuildV1().Builds(bc.Namespace).List(metav1.ListOptions{LabelSelector: buildutil.BuildConfigSelector(bc.Name).String()})
	if err != nil {
		return false, false, err
	}
	for _, item := range builds.Items {
		if item.Annotations[buildutil.BuildNumberAnnotation] == strconv.FormatInt(bc.Status.LastVersion, 10) {
			return checkBuildReadiness(&item)
		}
	}
	return false, false, nil
}

type deploymentCondition struct {
	status	corev1.ConditionStatus
	reason	string
}

func newDeploymentCondition(status corev1.ConditionStatus, reason string) *deploymentCondition {
	_logClusterCodePath()
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
	return &deploymentCondition{status: status, reason: reason}
}
func checkDeploymentReadiness(obj runtime.Object) (bool, bool, error) {
	_logClusterCodePath()
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
	var (
		isSynced		bool
		progressing, available	*deploymentCondition
	)
	switch d := obj.(type) {
	case *kappsv1.Deployment:
		isSynced = d.Status.ObservedGeneration == d.Generation
		for _, condition := range d.Status.Conditions {
			switch condition.Type {
			case kappsv1.DeploymentProgressing:
				progressing = newDeploymentCondition(condition.Status, condition.Reason)
			case kappsv1.DeploymentAvailable:
				available = newDeploymentCondition(condition.Status, condition.Reason)
			}
		}
	case *kappsv1beta1.Deployment:
		isSynced = d.Status.ObservedGeneration == d.Generation
		for _, condition := range d.Status.Conditions {
			switch condition.Type {
			case kappsv1beta1.DeploymentProgressing:
				progressing = newDeploymentCondition(condition.Status, condition.Reason)
			case kappsv1beta1.DeploymentAvailable:
				available = newDeploymentCondition(condition.Status, condition.Reason)
			}
		}
	case *kappsv1beta2.Deployment:
		isSynced = d.Status.ObservedGeneration == d.Generation
		for _, condition := range d.Status.Conditions {
			switch condition.Type {
			case kappsv1beta2.DeploymentProgressing:
				progressing = newDeploymentCondition(condition.Status, condition.Reason)
			case kappsv1beta2.DeploymentAvailable:
				available = newDeploymentCondition(condition.Status, condition.Reason)
			}
		}
	case *kextensionsv1beta1.Deployment:
		isSynced = d.Status.ObservedGeneration == d.Generation
		for _, condition := range d.Status.Conditions {
			switch condition.Type {
			case kextensionsv1beta1.DeploymentProgressing:
				progressing = newDeploymentCondition(condition.Status, condition.Reason)
			case kextensionsv1beta1.DeploymentAvailable:
				available = newDeploymentCondition(condition.Status, condition.Reason)
			}
		}
	default:
		return false, false, fmt.Errorf("unsupported deployment version: %T", d)
	}
	if !isSynced || progressing == nil {
		return false, false, nil
	}
	ready := progressing.status == corev1.ConditionTrue && progressing.reason == deploymentutil.NewRSAvailableReason && available != nil && available.status == corev1.ConditionTrue
	failed := progressing.status == corev1.ConditionFalse
	return ready, failed, nil
}
func checkDeploymentConfigReadiness(obj runtime.Object) (bool, bool, error) {
	_logClusterCodePath()
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
	dc, ok := obj.(*appsv1.DeploymentConfig)
	if !ok {
		return false, false, fmt.Errorf("object %T is not v1.DeploymentConfig", obj)
	}
	var progressing, available *appsv1.DeploymentCondition
	for i, condition := range dc.Status.Conditions {
		switch condition.Type {
		case appsv1.DeploymentProgressing:
			progressing = &dc.Status.Conditions[i]
		case appsv1.DeploymentAvailable:
			available = &dc.Status.Conditions[i]
		}
	}
	ready := dc.Status.ObservedGeneration == dc.Generation && progressing != nil && progressing.Status == corev1.ConditionTrue && progressing.Reason == appsutil.NewRcAvailableReason && available != nil && available.Status == corev1.ConditionTrue
	failed := dc.Status.ObservedGeneration == dc.Generation && progressing != nil && progressing.Status == corev1.ConditionFalse
	return ready, failed, nil
}
func checkJobReadiness(obj runtime.Object) (bool, bool, error) {
	_logClusterCodePath()
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
	var (
		hasCompletionTime	bool
		isJobFailed		bool
	)
	switch j := obj.(type) {
	case *batchv1.Job:
		hasCompletionTime = j.Status.CompletionTime != nil
		isJobFailed = j.Status.Failed > 0
	default:
		return false, false, fmt.Errorf("unsupported job version: %T", j)
	}
	return hasCompletionTime, isJobFailed, nil
}
func checkStatefulSetReadiness(obj runtime.Object) (bool, bool, error) {
	_logClusterCodePath()
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
	var (
		isSynced		bool
		hasReplicasReady	bool
	)
	switch s := obj.(type) {
	case *kappsv1.StatefulSet:
		isSynced = s.Status.ObservedGeneration == s.Generation
		hasReplicasReady = s.Spec.Replicas != nil && s.Status.ReadyReplicas == *s.Spec.Replicas
	case *kappsv1beta1.StatefulSet:
		isSynced = s.Status.ObservedGeneration != nil && *s.Status.ObservedGeneration == s.Generation
		hasReplicasReady = s.Spec.Replicas != nil && s.Status.ReadyReplicas == *s.Spec.Replicas
	case *kappsv1beta2.StatefulSet:
		isSynced = s.Status.ObservedGeneration == s.Generation
		hasReplicasReady = s.Spec.Replicas != nil && s.Status.ReadyReplicas == *s.Spec.Replicas
	default:
		return false, false, fmt.Errorf("unsupported statefulset version: %T", s)
	}
	return isSynced && hasReplicasReady, false, nil
}
func checkRouteReadiness(obj runtime.Object) (bool, bool, error) {
	_logClusterCodePath()
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
	route, ok := obj.(*routev1.Route)
	if !ok {
		return false, false, fmt.Errorf("object %T is not v1.Route", obj)
	}
	return len(route.Spec.Host) > 0, false, nil
}
func groupVersionKind(gv schema.GroupVersion, kind string) schema.GroupVersionKind {
	_logClusterCodePath()
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
	return schema.GroupVersionKind{Group: gv.Group, Version: gv.Version, Kind: kind}
}

var readinessCheckers = map[schema.GroupVersionKind]func(runtime.Object) (bool, bool, error){groupVersionKind(buildv1.GroupVersion, "Build"): checkBuildReadiness, groupVersionKind(appsv1.GroupVersion, "DeploymentConfig"): checkDeploymentConfigReadiness, groupVersionKind(routev1.GroupVersion, "Route"): checkRouteReadiness, legacy.GroupVersionKind("Build"): checkBuildReadiness, legacy.GroupVersionKind("DeploymentConfig"): checkDeploymentConfigReadiness, legacy.GroupVersionKind("Route"): checkRouteReadiness, groupVersionKind(kappsv1.SchemeGroupVersion, "Deployment"): checkDeploymentReadiness, groupVersionKind(kappsv1beta1.SchemeGroupVersion, "Deployment"): checkDeploymentReadiness, groupVersionKind(kappsv1beta2.SchemeGroupVersion, "Deployment"): checkDeploymentReadiness, groupVersionKind(kextensionsv1beta1.SchemeGroupVersion, "Deployment"): checkDeploymentReadiness, groupVersionKind(kappsv1.SchemeGroupVersion, "StatefulSet"): checkStatefulSetReadiness, groupVersionKind(kappsv1beta1.SchemeGroupVersion, "StatefulSet"): checkStatefulSetReadiness, groupVersionKind(kappsv1beta2.SchemeGroupVersion, "StatefulSet"): checkStatefulSetReadiness, groupVersionKind(batchv1.SchemeGroupVersion, "Job"): checkJobReadiness}

func CanCheckReadiness(ref corev1.ObjectReference) bool {
	_logClusterCodePath()
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
	switch ref.GroupVersionKind() {
	case groupVersionKind(buildv1.GroupVersion, "BuildConfig"), groupVersionKind(legacy.GroupVersion, "BuildConfig"):
		return true
	}
	_, found := readinessCheckers[ref.GroupVersionKind()]
	return found
}
func CheckReadiness(oc buildv1client.Interface, ref corev1.ObjectReference, obj *unstructured.Unstructured) (bool, bool, error) {
	_logClusterCodePath()
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
	castObj, err := readinessScheme.New(ref.GroupVersionKind())
	if err != nil {
		return false, false, err
	}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(obj.Object, castObj); err != nil {
		return false, false, err
	}
	switch ref.GroupVersionKind() {
	case groupVersionKind(buildv1.GroupVersion, "BuildConfig"), groupVersionKind(legacy.GroupVersion, "BuildConfig"):
		return checkBuildConfigReadiness(oc, castObj)
	}
	readinessCheckFunc, ok := readinessCheckers[ref.GroupVersionKind()]
	if !ok {
		return false, false, fmt.Errorf("readiness check for %+v is not defined", ref.GroupVersionKind())
	}
	return readinessCheckFunc(castObj)
}
