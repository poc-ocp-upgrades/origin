package podautoscaler

import (
	"fmt"
	goformat "fmt"
	autoscalingv1 "k8s.io/api/autoscaling/v1"
	autoscalingv2 "k8s.io/api/autoscaling/v2beta2"
	"k8s.io/api/core/v1"
	apiequality "k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/api/errors"
	apimeta "k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	autoscalinginformers "k8s.io/client-go/informers/autoscaling/v1"
	coreinformers "k8s.io/client-go/informers/core/v1"
	"k8s.io/client-go/kubernetes/scheme"
	autoscalingclient "k8s.io/client-go/kubernetes/typed/autoscaling/v1"
	v1core "k8s.io/client-go/kubernetes/typed/core/v1"
	autoscalinglisters "k8s.io/client-go/listers/autoscaling/v1"
	corelisters "k8s.io/client-go/listers/core/v1"
	scaleclient "k8s.io/client-go/scale"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	"k8s.io/kubernetes/pkg/controller"
	metricsclient "k8s.io/kubernetes/pkg/controller/podautoscaler/metrics"
	"math"
	goos "os"
	godefaultruntime "runtime"
	"time"
	gotime "time"
)

var (
	scaleUpLimitFactor  = 2.0
	scaleUpLimitMinimum = 4.0
)

type timestampedRecommendation struct {
	recommendation int32
	timestamp      time.Time
}
type HorizontalController struct {
	scaleNamespacer              scaleclient.ScalesGetter
	hpaNamespacer                autoscalingclient.HorizontalPodAutoscalersGetter
	mapper                       apimeta.RESTMapper
	replicaCalc                  *ReplicaCalculator
	eventRecorder                record.EventRecorder
	downscaleStabilisationWindow time.Duration
	hpaLister                    autoscalinglisters.HorizontalPodAutoscalerLister
	hpaListerSynced              cache.InformerSynced
	podLister                    corelisters.PodLister
	podListerSynced              cache.InformerSynced
	queue                        workqueue.RateLimitingInterface
	recommendations              map[string][]timestampedRecommendation
}

func NewHorizontalController(evtNamespacer v1core.EventsGetter, scaleNamespacer scaleclient.ScalesGetter, hpaNamespacer autoscalingclient.HorizontalPodAutoscalersGetter, mapper apimeta.RESTMapper, metricsClient metricsclient.MetricsClient, hpaInformer autoscalinginformers.HorizontalPodAutoscalerInformer, podInformer coreinformers.PodInformer, resyncPeriod time.Duration, downscaleStabilisationWindow time.Duration, tolerance float64, cpuInitializationPeriod, delayOfInitialReadinessStatus time.Duration) *HorizontalController {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	broadcaster := record.NewBroadcaster()
	broadcaster.StartLogging(klog.Infof)
	broadcaster.StartRecordingToSink(&v1core.EventSinkImpl{Interface: evtNamespacer.Events("")})
	recorder := broadcaster.NewRecorder(scheme.Scheme, v1.EventSource{Component: "horizontal-pod-autoscaler"})
	hpaController := &HorizontalController{eventRecorder: recorder, scaleNamespacer: scaleNamespacer, hpaNamespacer: hpaNamespacer, downscaleStabilisationWindow: downscaleStabilisationWindow, queue: workqueue.NewNamedRateLimitingQueue(NewDefaultHPARateLimiter(resyncPeriod), "horizontalpodautoscaler"), mapper: mapper, recommendations: map[string][]timestampedRecommendation{}}
	hpaInformer.Informer().AddEventHandlerWithResyncPeriod(cache.ResourceEventHandlerFuncs{AddFunc: hpaController.enqueueHPA, UpdateFunc: hpaController.updateHPA, DeleteFunc: hpaController.deleteHPA}, resyncPeriod)
	hpaController.hpaLister = hpaInformer.Lister()
	hpaController.hpaListerSynced = hpaInformer.Informer().HasSynced
	hpaController.podLister = podInformer.Lister()
	hpaController.podListerSynced = podInformer.Informer().HasSynced
	replicaCalc := NewReplicaCalculator(metricsClient, hpaController.podLister, tolerance, cpuInitializationPeriod, delayOfInitialReadinessStatus)
	hpaController.replicaCalc = replicaCalc
	return hpaController
}
func (a *HorizontalController) Run(stopCh <-chan struct{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	defer utilruntime.HandleCrash()
	defer a.queue.ShutDown()
	klog.Infof("Starting HPA controller")
	defer klog.Infof("Shutting down HPA controller")
	if !controller.WaitForCacheSync("HPA", stopCh, a.hpaListerSynced, a.podListerSynced) {
		return
	}
	go wait.Until(a.worker, time.Second, stopCh)
	<-stopCh
}
func (a *HorizontalController) updateHPA(old, cur interface{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	a.enqueueHPA(cur)
}
func (a *HorizontalController) enqueueHPA(obj interface{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	key, err := controller.KeyFunc(obj)
	if err != nil {
		utilruntime.HandleError(fmt.Errorf("couldn't get key for object %+v: %v", obj, err))
		return
	}
	a.queue.AddRateLimited(key)
}
func (a *HorizontalController) deleteHPA(obj interface{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	key, err := controller.KeyFunc(obj)
	if err != nil {
		utilruntime.HandleError(fmt.Errorf("couldn't get key for object %+v: %v", obj, err))
		return
	}
	a.queue.Forget(key)
}
func (a *HorizontalController) worker() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for a.processNextWorkItem() {
	}
	klog.Infof("horizontal pod autoscaler controller worker shutting down")
}
func (a *HorizontalController) processNextWorkItem() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	key, quit := a.queue.Get()
	if quit {
		return false
	}
	defer a.queue.Done(key)
	deleted, err := a.reconcileKey(key.(string))
	if err != nil {
		utilruntime.HandleError(err)
	}
	if !deleted {
		a.queue.AddRateLimited(key)
	}
	return true
}
func (a *HorizontalController) computeReplicasForMetrics(hpa *autoscalingv2.HorizontalPodAutoscaler, scale *autoscalingv1.Scale, metricSpecs []autoscalingv2.MetricSpec) (replicas int32, metric string, statuses []autoscalingv2.MetricStatus, timestamp time.Time, err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	currentReplicas := scale.Status.Replicas
	statuses = make([]autoscalingv2.MetricStatus, len(metricSpecs))
	for i, metricSpec := range metricSpecs {
		if scale.Status.Selector == "" {
			errMsg := "selector is required"
			a.eventRecorder.Event(hpa, v1.EventTypeWarning, "SelectorRequired", errMsg)
			setCondition(hpa, autoscalingv2.ScalingActive, v1.ConditionFalse, "InvalidSelector", "the HPA target's scale is missing a selector")
			return 0, "", nil, time.Time{}, fmt.Errorf(errMsg)
		}
		selector, err := labels.Parse(scale.Status.Selector)
		if err != nil {
			errMsg := fmt.Sprintf("couldn't convert selector into a corresponding internal selector object: %v", err)
			a.eventRecorder.Event(hpa, v1.EventTypeWarning, "InvalidSelector", errMsg)
			setCondition(hpa, autoscalingv2.ScalingActive, v1.ConditionFalse, "InvalidSelector", errMsg)
			return 0, "", nil, time.Time{}, fmt.Errorf(errMsg)
		}
		var replicaCountProposal int32
		var timestampProposal time.Time
		var metricNameProposal string
		switch metricSpec.Type {
		case autoscalingv2.ObjectMetricSourceType:
			metricSelector, err := metav1.LabelSelectorAsSelector(metricSpec.Object.Metric.Selector)
			if err != nil {
				a.eventRecorder.Event(hpa, v1.EventTypeWarning, "FailedGetObjectMetric", err.Error())
				setCondition(hpa, autoscalingv2.ScalingActive, v1.ConditionFalse, "FailedGetObjectMetric", "the HPA was unable to compute the replica count: %v", err)
				return 0, "", nil, time.Time{}, fmt.Errorf("failed to get object metric value: %v", err)
			}
			replicaCountProposal, timestampProposal, metricNameProposal, err = a.computeStatusForObjectMetric(currentReplicas, metricSpec, hpa, selector, &statuses[i], metricSelector)
			if err != nil {
				return 0, "", nil, time.Time{}, fmt.Errorf("failed to get object metric value: %v", err)
			}
		case autoscalingv2.PodsMetricSourceType:
			metricSelector, err := metav1.LabelSelectorAsSelector(metricSpec.Pods.Metric.Selector)
			if err != nil {
				a.eventRecorder.Event(hpa, v1.EventTypeWarning, "FailedGetPodsMetric", err.Error())
				setCondition(hpa, autoscalingv2.ScalingActive, v1.ConditionFalse, "FailedGetPodsMetric", "the HPA was unable to compute the replica count: %v", err)
				return 0, "", nil, time.Time{}, fmt.Errorf("failed to get pods metric value: %v", err)
			}
			replicaCountProposal, timestampProposal, metricNameProposal, err = a.computeStatusForPodsMetric(currentReplicas, metricSpec, hpa, selector, &statuses[i], metricSelector)
			if err != nil {
				return 0, "", nil, time.Time{}, fmt.Errorf("failed to get object metric value: %v", err)
			}
		case autoscalingv2.ResourceMetricSourceType:
			replicaCountProposal, timestampProposal, metricNameProposal, err = a.computeStatusForResourceMetric(currentReplicas, metricSpec, hpa, selector, &statuses[i])
			if err != nil {
				return 0, "", nil, time.Time{}, err
			}
		case autoscalingv2.ExternalMetricSourceType:
			replicaCountProposal, timestampProposal, metricNameProposal, err = a.computeStatusForExternalMetric(currentReplicas, metricSpec, hpa, selector, &statuses[i])
			if err != nil {
				return 0, "", nil, time.Time{}, err
			}
		default:
			errMsg := fmt.Sprintf("unknown metric source type %q", string(metricSpec.Type))
			a.eventRecorder.Event(hpa, v1.EventTypeWarning, "InvalidMetricSourceType", errMsg)
			setCondition(hpa, autoscalingv2.ScalingActive, v1.ConditionFalse, "InvalidMetricSourceType", "the HPA was unable to compute the replica count: %s", errMsg)
			return 0, "", nil, time.Time{}, fmt.Errorf(errMsg)
		}
		if replicas == 0 || replicaCountProposal > replicas {
			timestamp = timestampProposal
			replicas = replicaCountProposal
			metric = metricNameProposal
		}
	}
	setCondition(hpa, autoscalingv2.ScalingActive, v1.ConditionTrue, "ValidMetricFound", "the HPA was able to successfully calculate a replica count from %s", metric)
	return replicas, metric, statuses, timestamp, nil
}
func (a *HorizontalController) reconcileKey(key string) (deleted bool, err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	namespace, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		return true, err
	}
	hpa, err := a.hpaLister.HorizontalPodAutoscalers(namespace).Get(name)
	if errors.IsNotFound(err) {
		klog.Infof("Horizontal Pod Autoscaler %s has been deleted in %s", name, namespace)
		delete(a.recommendations, key)
		return true, nil
	}
	return false, a.reconcileAutoscaler(hpa, key)
}
func (a *HorizontalController) computeStatusForObjectMetric(currentReplicas int32, metricSpec autoscalingv2.MetricSpec, hpa *autoscalingv2.HorizontalPodAutoscaler, selector labels.Selector, status *autoscalingv2.MetricStatus, metricSelector labels.Selector) (int32, time.Time, string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	replicaCountProposal, utilizationProposal, timestampProposal, err := a.replicaCalc.GetObjectMetricReplicas(currentReplicas, metricSpec.Object.Target.Value.MilliValue(), metricSpec.Object.Metric.Name, hpa.Namespace, &metricSpec.Object.DescribedObject, selector, metricSelector)
	if err != nil {
		a.eventRecorder.Event(hpa, v1.EventTypeWarning, "FailedGetObjectMetric", err.Error())
		setCondition(hpa, autoscalingv2.ScalingActive, v1.ConditionFalse, "FailedGetObjectMetric", "the HPA was unable to compute the replica count: %v", err)
		return 0, timestampProposal, "", err
	}
	*status = autoscalingv2.MetricStatus{Type: autoscalingv2.ObjectMetricSourceType, Object: &autoscalingv2.ObjectMetricStatus{DescribedObject: metricSpec.Object.DescribedObject, Metric: autoscalingv2.MetricIdentifier{Name: metricSpec.Object.Metric.Name, Selector: metricSpec.Object.Metric.Selector}, Current: autoscalingv2.MetricValueStatus{Value: resource.NewMilliQuantity(utilizationProposal, resource.DecimalSI)}}}
	return replicaCountProposal, timestampProposal, fmt.Sprintf("%s metric %s", metricSpec.Object.DescribedObject.Kind, metricSpec.Object.Metric.Name), nil
}
func (a *HorizontalController) computeStatusForPodsMetric(currentReplicas int32, metricSpec autoscalingv2.MetricSpec, hpa *autoscalingv2.HorizontalPodAutoscaler, selector labels.Selector, status *autoscalingv2.MetricStatus, metricSelector labels.Selector) (int32, time.Time, string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	replicaCountProposal, utilizationProposal, timestampProposal, err := a.replicaCalc.GetMetricReplicas(currentReplicas, metricSpec.Pods.Target.AverageValue.MilliValue(), metricSpec.Pods.Metric.Name, hpa.Namespace, selector, metricSelector)
	if err != nil {
		a.eventRecorder.Event(hpa, v1.EventTypeWarning, "FailedGetPodsMetric", err.Error())
		setCondition(hpa, autoscalingv2.ScalingActive, v1.ConditionFalse, "FailedGetPodsMetric", "the HPA was unable to compute the replica count: %v", err)
		return 0, timestampProposal, "", err
	}
	*status = autoscalingv2.MetricStatus{Type: autoscalingv2.PodsMetricSourceType, Pods: &autoscalingv2.PodsMetricStatus{Metric: autoscalingv2.MetricIdentifier{Name: metricSpec.Pods.Metric.Name, Selector: metricSpec.Pods.Metric.Selector}, Current: autoscalingv2.MetricValueStatus{AverageValue: resource.NewMilliQuantity(utilizationProposal, resource.DecimalSI)}}}
	return replicaCountProposal, timestampProposal, fmt.Sprintf("pods metric %s", metricSpec.Pods.Metric.Name), nil
}
func (a *HorizontalController) computeStatusForResourceMetric(currentReplicas int32, metricSpec autoscalingv2.MetricSpec, hpa *autoscalingv2.HorizontalPodAutoscaler, selector labels.Selector, status *autoscalingv2.MetricStatus) (int32, time.Time, string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if metricSpec.Resource.Target.AverageValue != nil {
		var rawProposal int64
		replicaCountProposal, rawProposal, timestampProposal, err := a.replicaCalc.GetRawResourceReplicas(currentReplicas, metricSpec.Resource.Target.AverageValue.MilliValue(), metricSpec.Resource.Name, hpa.Namespace, selector)
		if err != nil {
			a.eventRecorder.Event(hpa, v1.EventTypeWarning, "FailedGetResourceMetric", err.Error())
			setCondition(hpa, autoscalingv2.ScalingActive, v1.ConditionFalse, "FailedGetResourceMetric", "the HPA was unable to compute the replica count: %v", err)
			return 0, time.Time{}, "", fmt.Errorf("failed to get %s utilization: %v", metricSpec.Resource.Name, err)
		}
		metricNameProposal := fmt.Sprintf("%s resource", metricSpec.Resource.Name)
		status = &autoscalingv2.MetricStatus{Type: autoscalingv2.ResourceMetricSourceType, Resource: &autoscalingv2.ResourceMetricStatus{Name: metricSpec.Resource.Name, Current: autoscalingv2.MetricValueStatus{AverageValue: resource.NewMilliQuantity(rawProposal, resource.DecimalSI)}}}
		return replicaCountProposal, timestampProposal, metricNameProposal, nil
	} else {
		if metricSpec.Resource.Target.AverageUtilization == nil {
			errMsg := "invalid resource metric source: neither a utilization target nor a value target was set"
			a.eventRecorder.Event(hpa, v1.EventTypeWarning, "FailedGetResourceMetric", errMsg)
			setCondition(hpa, autoscalingv2.ScalingActive, v1.ConditionFalse, "FailedGetResourceMetric", "the HPA was unable to compute the replica count: %s", errMsg)
			return 0, time.Time{}, "", fmt.Errorf(errMsg)
		}
		targetUtilization := *metricSpec.Resource.Target.AverageUtilization
		var percentageProposal int32
		var rawProposal int64
		replicaCountProposal, percentageProposal, rawProposal, timestampProposal, err := a.replicaCalc.GetResourceReplicas(currentReplicas, targetUtilization, metricSpec.Resource.Name, hpa.Namespace, selector)
		if err != nil {
			a.eventRecorder.Event(hpa, v1.EventTypeWarning, "FailedGetResourceMetric", err.Error())
			setCondition(hpa, autoscalingv2.ScalingActive, v1.ConditionFalse, "FailedGetResourceMetric", "the HPA was unable to compute the replica count: %v", err)
			return 0, time.Time{}, "", fmt.Errorf("failed to get %s utilization: %v", metricSpec.Resource.Name, err)
		}
		metricNameProposal := fmt.Sprintf("%s resource utilization (percentage of request)", metricSpec.Resource.Name)
		*status = autoscalingv2.MetricStatus{Type: autoscalingv2.ResourceMetricSourceType, Resource: &autoscalingv2.ResourceMetricStatus{Name: metricSpec.Resource.Name, Current: autoscalingv2.MetricValueStatus{AverageUtilization: &percentageProposal, AverageValue: resource.NewMilliQuantity(rawProposal, resource.DecimalSI)}}}
		return replicaCountProposal, timestampProposal, metricNameProposal, nil
	}
}
func (a *HorizontalController) computeStatusForExternalMetric(currentReplicas int32, metricSpec autoscalingv2.MetricSpec, hpa *autoscalingv2.HorizontalPodAutoscaler, selector labels.Selector, status *autoscalingv2.MetricStatus) (int32, time.Time, string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if metricSpec.External.Target.AverageValue != nil {
		replicaCountProposal, utilizationProposal, timestampProposal, err := a.replicaCalc.GetExternalPerPodMetricReplicas(currentReplicas, metricSpec.External.Target.AverageValue.MilliValue(), metricSpec.External.Metric.Name, hpa.Namespace, metricSpec.External.Metric.Selector)
		if err != nil {
			a.eventRecorder.Event(hpa, v1.EventTypeWarning, "FailedGetExternalMetric", err.Error())
			setCondition(hpa, autoscalingv2.ScalingActive, v1.ConditionFalse, "FailedGetExternalMetric", "the HPA was unable to compute the replica count: %v", err)
			return 0, time.Time{}, "", fmt.Errorf("failed to get %s external metric: %v", metricSpec.External.Metric.Name, err)
		}
		*status = autoscalingv2.MetricStatus{Type: autoscalingv2.ExternalMetricSourceType, External: &autoscalingv2.ExternalMetricStatus{Metric: autoscalingv2.MetricIdentifier{Name: metricSpec.External.Metric.Name, Selector: metricSpec.External.Metric.Selector}, Current: autoscalingv2.MetricValueStatus{AverageValue: resource.NewMilliQuantity(utilizationProposal, resource.DecimalSI)}}}
		return replicaCountProposal, timestampProposal, fmt.Sprintf("external metric %s(%+v)", metricSpec.External.Metric.Name, metricSpec.External.Metric.Selector), nil
	}
	if metricSpec.External.Target.Value != nil {
		replicaCountProposal, utilizationProposal, timestampProposal, err := a.replicaCalc.GetExternalMetricReplicas(currentReplicas, metricSpec.External.Target.Value.MilliValue(), metricSpec.External.Metric.Name, hpa.Namespace, metricSpec.External.Metric.Selector, selector)
		if err != nil {
			a.eventRecorder.Event(hpa, v1.EventTypeWarning, "FailedGetExternalMetric", err.Error())
			setCondition(hpa, autoscalingv2.ScalingActive, v1.ConditionFalse, "FailedGetExternalMetric", "the HPA was unable to compute the replica count: %v", err)
			return 0, time.Time{}, "", fmt.Errorf("failed to get external metric %s: %v", metricSpec.External.Metric.Name, err)
		}
		*status = autoscalingv2.MetricStatus{Type: autoscalingv2.ExternalMetricSourceType, External: &autoscalingv2.ExternalMetricStatus{Metric: autoscalingv2.MetricIdentifier{Name: metricSpec.External.Metric.Name, Selector: metricSpec.External.Metric.Selector}, Current: autoscalingv2.MetricValueStatus{Value: resource.NewMilliQuantity(utilizationProposal, resource.DecimalSI)}}}
		return replicaCountProposal, timestampProposal, fmt.Sprintf("external metric %s(%+v)", metricSpec.External.Metric.Name, metricSpec.External.Metric.Selector), nil
	}
	errMsg := "invalid external metric source: neither a value target nor an average value target was set"
	a.eventRecorder.Event(hpa, v1.EventTypeWarning, "FailedGetExternalMetric", errMsg)
	setCondition(hpa, autoscalingv2.ScalingActive, v1.ConditionFalse, "FailedGetExternalMetric", "the HPA was unable to compute the replica count: %s", errMsg)
	return 0, time.Time{}, "", fmt.Errorf(errMsg)
}
func (a *HorizontalController) recordInitialRecommendation(currentReplicas int32, key string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if a.recommendations[key] == nil {
		a.recommendations[key] = []timestampedRecommendation{{currentReplicas, time.Now()}}
	}
}
func (a *HorizontalController) reconcileAutoscaler(hpav1Shared *autoscalingv1.HorizontalPodAutoscaler, key string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	hpav1 := hpav1Shared.DeepCopy()
	hpaRaw, err := unsafeConvertToVersionVia(hpav1, autoscalingv2.SchemeGroupVersion)
	if err != nil {
		a.eventRecorder.Event(hpav1, v1.EventTypeWarning, "FailedConvertHPA", err.Error())
		return fmt.Errorf("failed to convert the given HPA to %s: %v", autoscalingv2.SchemeGroupVersion.String(), err)
	}
	hpa := hpaRaw.(*autoscalingv2.HorizontalPodAutoscaler)
	hpaStatusOriginal := hpa.Status.DeepCopy()
	reference := fmt.Sprintf("%s/%s/%s", hpa.Spec.ScaleTargetRef.Kind, hpa.Namespace, hpa.Spec.ScaleTargetRef.Name)
	targetGV, err := schema.ParseGroupVersion(hpa.Spec.ScaleTargetRef.APIVersion)
	if err != nil {
		a.eventRecorder.Event(hpa, v1.EventTypeWarning, "FailedGetScale", err.Error())
		setCondition(hpa, autoscalingv2.AbleToScale, v1.ConditionFalse, "FailedGetScale", "the HPA controller was unable to get the target's current scale: %v", err)
		a.updateStatusIfNeeded(hpaStatusOriginal, hpa)
		return fmt.Errorf("invalid API version in scale target reference: %v", err)
	}
	targetGK := schema.GroupKind{Group: targetGV.Group, Kind: hpa.Spec.ScaleTargetRef.Kind}
	mappings, err := a.mapper.RESTMappings(targetGK)
	mappings, err = overrideMappingsForOapiDeploymentConfig(mappings, err, targetGK)
	if err != nil {
		a.eventRecorder.Event(hpa, v1.EventTypeWarning, "FailedGetScale", err.Error())
		setCondition(hpa, autoscalingv2.AbleToScale, v1.ConditionFalse, "FailedGetScale", "the HPA controller was unable to get the target's current scale: %v", err)
		a.updateStatusIfNeeded(hpaStatusOriginal, hpa)
		return fmt.Errorf("unable to determine resource for scale target reference: %v", err)
	}
	scale, targetGR, err := a.scaleForResourceMappings(hpa.Namespace, hpa.Spec.ScaleTargetRef.Name, mappings)
	if err != nil {
		a.eventRecorder.Event(hpa, v1.EventTypeWarning, "FailedGetScale", err.Error())
		setCondition(hpa, autoscalingv2.AbleToScale, v1.ConditionFalse, "FailedGetScale", "the HPA controller was unable to get the target's current scale: %v", err)
		a.updateStatusIfNeeded(hpaStatusOriginal, hpa)
		return fmt.Errorf("failed to query scale subresource for %s: %v", reference, err)
	}
	setCondition(hpa, autoscalingv2.AbleToScale, v1.ConditionTrue, "SucceededGetScale", "the HPA controller was able to get the target's current scale")
	currentReplicas := scale.Status.Replicas
	a.recordInitialRecommendation(currentReplicas, key)
	var metricStatuses []autoscalingv2.MetricStatus
	metricDesiredReplicas := int32(0)
	metricName := ""
	metricTimestamp := time.Time{}
	desiredReplicas := int32(0)
	rescaleReason := ""
	timestamp := time.Now()
	rescale := true
	if scale.Spec.Replicas == 0 {
		desiredReplicas = 0
		rescale = false
		setCondition(hpa, autoscalingv2.ScalingActive, v1.ConditionFalse, "ScalingDisabled", "scaling is disabled since the replica count of the target is zero")
	} else if currentReplicas > hpa.Spec.MaxReplicas {
		rescaleReason = "Current number of replicas above Spec.MaxReplicas"
		desiredReplicas = hpa.Spec.MaxReplicas
	} else if hpa.Spec.MinReplicas != nil && currentReplicas < *hpa.Spec.MinReplicas {
		rescaleReason = "Current number of replicas below Spec.MinReplicas"
		desiredReplicas = *hpa.Spec.MinReplicas
	} else if currentReplicas == 0 {
		rescaleReason = "Current number of replicas must be greater than 0"
		desiredReplicas = 1
	} else {
		metricDesiredReplicas, metricName, metricStatuses, metricTimestamp, err = a.computeReplicasForMetrics(hpa, scale, hpa.Spec.Metrics)
		if err != nil {
			a.setCurrentReplicasInStatus(hpa, currentReplicas)
			if err := a.updateStatusIfNeeded(hpaStatusOriginal, hpa); err != nil {
				utilruntime.HandleError(err)
			}
			a.eventRecorder.Event(hpa, v1.EventTypeWarning, "FailedComputeMetricsReplicas", err.Error())
			return fmt.Errorf("failed to compute desired number of replicas based on listed metrics for %s: %v", reference, err)
		}
		klog.V(4).Infof("proposing %v desired replicas (based on %s from %s) for %s", metricDesiredReplicas, metricName, timestamp, reference)
		rescaleMetric := ""
		if metricDesiredReplicas > desiredReplicas {
			desiredReplicas = metricDesiredReplicas
			timestamp = metricTimestamp
			rescaleMetric = metricName
		}
		if desiredReplicas > currentReplicas {
			rescaleReason = fmt.Sprintf("%s above target", rescaleMetric)
		}
		if desiredReplicas < currentReplicas {
			rescaleReason = "All metrics below target"
		}
		desiredReplicas = a.normalizeDesiredReplicas(hpa, key, currentReplicas, desiredReplicas)
		rescale = desiredReplicas != currentReplicas
	}
	if rescale {
		scale.Spec.Replicas = desiredReplicas
		_, err = a.scaleNamespacer.Scales(hpa.Namespace).Update(targetGR, scale)
		if err != nil {
			a.eventRecorder.Eventf(hpa, v1.EventTypeWarning, "FailedRescale", "New size: %d; reason: %s; error: %v", desiredReplicas, rescaleReason, err.Error())
			setCondition(hpa, autoscalingv2.AbleToScale, v1.ConditionFalse, "FailedUpdateScale", "the HPA controller was unable to update the target scale: %v", err)
			a.setCurrentReplicasInStatus(hpa, currentReplicas)
			if err := a.updateStatusIfNeeded(hpaStatusOriginal, hpa); err != nil {
				utilruntime.HandleError(err)
			}
			return fmt.Errorf("failed to rescale %s: %v", reference, err)
		}
		setCondition(hpa, autoscalingv2.AbleToScale, v1.ConditionTrue, "SucceededRescale", "the HPA controller was able to update the target scale to %d", desiredReplicas)
		a.eventRecorder.Eventf(hpa, v1.EventTypeNormal, "SuccessfulRescale", "New size: %d; reason: %s", desiredReplicas, rescaleReason)
		klog.Infof("Successful rescale of %s, old size: %d, new size: %d, reason: %s", hpa.Name, currentReplicas, desiredReplicas, rescaleReason)
	} else {
		klog.V(4).Infof("decided not to scale %s to %v (last scale time was %s)", reference, desiredReplicas, hpa.Status.LastScaleTime)
		desiredReplicas = currentReplicas
	}
	a.setStatus(hpa, currentReplicas, desiredReplicas, metricStatuses, rescale)
	return a.updateStatusIfNeeded(hpaStatusOriginal, hpa)
}
func (a *HorizontalController) stabilizeRecommendation(key string, prenormalizedDesiredReplicas int32) int32 {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	maxRecommendation := prenormalizedDesiredReplicas
	foundOldSample := false
	oldSampleIndex := 0
	cutoff := time.Now().Add(-a.downscaleStabilisationWindow)
	for i, rec := range a.recommendations[key] {
		if rec.timestamp.Before(cutoff) {
			foundOldSample = true
			oldSampleIndex = i
		} else if rec.recommendation > maxRecommendation {
			maxRecommendation = rec.recommendation
		}
	}
	if foundOldSample {
		a.recommendations[key][oldSampleIndex] = timestampedRecommendation{prenormalizedDesiredReplicas, time.Now()}
	} else {
		a.recommendations[key] = append(a.recommendations[key], timestampedRecommendation{prenormalizedDesiredReplicas, time.Now()})
	}
	return maxRecommendation
}
func (a *HorizontalController) normalizeDesiredReplicas(hpa *autoscalingv2.HorizontalPodAutoscaler, key string, currentReplicas int32, prenormalizedDesiredReplicas int32) int32 {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	stabilizedRecommendation := a.stabilizeRecommendation(key, prenormalizedDesiredReplicas)
	if stabilizedRecommendation != prenormalizedDesiredReplicas {
		setCondition(hpa, autoscalingv2.AbleToScale, v1.ConditionTrue, "ScaleDownStabilized", "recent recommendations were higher than current one, applying the highest recent recommendation")
	} else {
		setCondition(hpa, autoscalingv2.AbleToScale, v1.ConditionTrue, "ReadyForNewScale", "recommended size matches current size")
	}
	var minReplicas int32
	if hpa.Spec.MinReplicas != nil {
		minReplicas = *hpa.Spec.MinReplicas
	} else {
		minReplicas = 0
	}
	desiredReplicas, condition, reason := convertDesiredReplicasWithRules(currentReplicas, stabilizedRecommendation, minReplicas, hpa.Spec.MaxReplicas)
	if desiredReplicas == stabilizedRecommendation {
		setCondition(hpa, autoscalingv2.ScalingLimited, v1.ConditionFalse, condition, reason)
	} else {
		setCondition(hpa, autoscalingv2.ScalingLimited, v1.ConditionTrue, condition, reason)
	}
	return desiredReplicas
}
func convertDesiredReplicasWithRules(currentReplicas, desiredReplicas, hpaMinReplicas, hpaMaxReplicas int32) (int32, string, string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var minimumAllowedReplicas int32
	var maximumAllowedReplicas int32
	var possibleLimitingCondition string
	var possibleLimitingReason string
	if hpaMinReplicas == 0 {
		minimumAllowedReplicas = 1
		possibleLimitingReason = "the desired replica count is zero"
	} else {
		minimumAllowedReplicas = hpaMinReplicas
		possibleLimitingReason = "the desired replica count is less than the minimum replica count"
	}
	scaleUpLimit := calculateScaleUpLimit(currentReplicas)
	if hpaMaxReplicas > scaleUpLimit {
		maximumAllowedReplicas = scaleUpLimit
		possibleLimitingCondition = "ScaleUpLimit"
		possibleLimitingReason = "the desired replica count is increasing faster than the maximum scale rate"
	} else {
		maximumAllowedReplicas = hpaMaxReplicas
		possibleLimitingCondition = "TooManyReplicas"
		possibleLimitingReason = "the desired replica count is more than the maximum replica count"
	}
	if desiredReplicas < minimumAllowedReplicas {
		possibleLimitingCondition = "TooFewReplicas"
		return minimumAllowedReplicas, possibleLimitingCondition, possibleLimitingReason
	} else if desiredReplicas > maximumAllowedReplicas {
		return maximumAllowedReplicas, possibleLimitingCondition, possibleLimitingReason
	}
	return desiredReplicas, "DesiredWithinRange", "the desired count is within the acceptable range"
}
func calculateScaleUpLimit(currentReplicas int32) int32 {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return int32(math.Max(scaleUpLimitFactor*float64(currentReplicas), scaleUpLimitMinimum))
}
func (a *HorizontalController) scaleForResourceMappings(namespace, name string, mappings []*apimeta.RESTMapping) (*autoscalingv1.Scale, schema.GroupResource, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var firstErr error
	for i, mapping := range mappings {
		targetGR := mapping.Resource.GroupResource()
		scale, err := a.scaleNamespacer.Scales(namespace).Get(targetGR, name)
		if err == nil {
			return scale, targetGR, nil
		}
		if i == 0 {
			firstErr = err
		}
	}
	if firstErr == nil {
		firstErr = fmt.Errorf("unrecognized resource")
	}
	return nil, schema.GroupResource{}, firstErr
}
func (a *HorizontalController) setCurrentReplicasInStatus(hpa *autoscalingv2.HorizontalPodAutoscaler, currentReplicas int32) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	a.setStatus(hpa, currentReplicas, hpa.Status.DesiredReplicas, hpa.Status.CurrentMetrics, false)
}
func (a *HorizontalController) setStatus(hpa *autoscalingv2.HorizontalPodAutoscaler, currentReplicas, desiredReplicas int32, metricStatuses []autoscalingv2.MetricStatus, rescale bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	hpa.Status = autoscalingv2.HorizontalPodAutoscalerStatus{CurrentReplicas: currentReplicas, DesiredReplicas: desiredReplicas, LastScaleTime: hpa.Status.LastScaleTime, CurrentMetrics: metricStatuses, Conditions: hpa.Status.Conditions}
	if rescale {
		now := metav1.NewTime(time.Now())
		hpa.Status.LastScaleTime = &now
	}
}
func (a *HorizontalController) updateStatusIfNeeded(oldStatus *autoscalingv2.HorizontalPodAutoscalerStatus, newHPA *autoscalingv2.HorizontalPodAutoscaler) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if apiequality.Semantic.DeepEqual(oldStatus, &newHPA.Status) {
		return nil
	}
	return a.updateStatus(newHPA)
}
func (a *HorizontalController) updateStatus(hpa *autoscalingv2.HorizontalPodAutoscaler) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	hpaRaw, err := unsafeConvertToVersionVia(hpa, autoscalingv1.SchemeGroupVersion)
	if err != nil {
		a.eventRecorder.Event(hpa, v1.EventTypeWarning, "FailedConvertHPA", err.Error())
		return fmt.Errorf("failed to convert the given HPA to %s: %v", autoscalingv2.SchemeGroupVersion.String(), err)
	}
	hpav1 := hpaRaw.(*autoscalingv1.HorizontalPodAutoscaler)
	_, err = a.hpaNamespacer.HorizontalPodAutoscalers(hpav1.Namespace).UpdateStatus(hpav1)
	if err != nil {
		a.eventRecorder.Event(hpa, v1.EventTypeWarning, "FailedUpdateStatus", err.Error())
		return fmt.Errorf("failed to update status for %s: %v", hpa.Name, err)
	}
	klog.V(2).Infof("Successfully updated status for %s", hpa.Name)
	return nil
}
func unsafeConvertToVersionVia(obj runtime.Object, externalVersion schema.GroupVersion) (runtime.Object, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	objInt, err := legacyscheme.Scheme.UnsafeConvertToVersion(obj, schema.GroupVersion{Group: externalVersion.Group, Version: runtime.APIVersionInternal})
	if err != nil {
		return nil, fmt.Errorf("failed to convert the given object to the internal version: %v", err)
	}
	objExt, err := legacyscheme.Scheme.UnsafeConvertToVersion(objInt, externalVersion)
	if err != nil {
		return nil, fmt.Errorf("failed to convert the given object back to the external version: %v", err)
	}
	return objExt, err
}
func setCondition(hpa *autoscalingv2.HorizontalPodAutoscaler, conditionType autoscalingv2.HorizontalPodAutoscalerConditionType, status v1.ConditionStatus, reason, message string, args ...interface{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	hpa.Status.Conditions = setConditionInList(hpa.Status.Conditions, conditionType, status, reason, message, args...)
}
func setConditionInList(inputList []autoscalingv2.HorizontalPodAutoscalerCondition, conditionType autoscalingv2.HorizontalPodAutoscalerConditionType, status v1.ConditionStatus, reason, message string, args ...interface{}) []autoscalingv2.HorizontalPodAutoscalerCondition {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	resList := inputList
	var existingCond *autoscalingv2.HorizontalPodAutoscalerCondition
	for i, condition := range resList {
		if condition.Type == conditionType {
			existingCond = &resList[i]
			break
		}
	}
	if existingCond == nil {
		resList = append(resList, autoscalingv2.HorizontalPodAutoscalerCondition{Type: conditionType})
		existingCond = &resList[len(resList)-1]
	}
	if existingCond.Status != status {
		existingCond.LastTransitionTime = metav1.Now()
	}
	existingCond.Status = status
	existingCond.Reason = reason
	existingCond.Message = fmt.Sprintf(message, args...)
	return resList
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
