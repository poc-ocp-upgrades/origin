package endpoint

import (
	"fmt"
	goformat "fmt"
	"k8s.io/api/core/v1"
	apiequality "k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apimachinery/pkg/util/wait"
	coreinformers "k8s.io/client-go/informers/core/v1"
	clientset "k8s.io/client-go/kubernetes"
	corelisters "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/leaderelection/resourcelock"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/api/v1/endpoints"
	podutil "k8s.io/kubernetes/pkg/api/v1/pod"
	api "k8s.io/kubernetes/pkg/apis/core"
	"k8s.io/kubernetes/pkg/controller"
	"k8s.io/kubernetes/pkg/util/metrics"
	goos "os"
	"reflect"
	godefaultruntime "runtime"
	"strconv"
	"time"
	gotime "time"
)

const (
	maxRetries                         = 15
	TolerateUnreadyEndpointsAnnotation = "service.alpha.kubernetes.io/tolerate-unready-endpoints"
)

func NewEndpointController(podInformer coreinformers.PodInformer, serviceInformer coreinformers.ServiceInformer, endpointsInformer coreinformers.EndpointsInformer, client clientset.Interface) *EndpointController {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if client != nil && client.CoreV1().RESTClient().GetRateLimiter() != nil {
		metrics.RegisterMetricAndTrackRateLimiterUsage("endpoint_controller", client.CoreV1().RESTClient().GetRateLimiter())
	}
	e := &EndpointController{client: client, queue: workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "endpoint"), workerLoopPeriod: time.Second}
	serviceInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{AddFunc: e.enqueueService, UpdateFunc: func(old, cur interface{}) {
		e.enqueueService(cur)
	}, DeleteFunc: e.enqueueService})
	e.serviceLister = serviceInformer.Lister()
	e.servicesSynced = serviceInformer.Informer().HasSynced
	podInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{AddFunc: e.addPod, UpdateFunc: e.updatePod, DeleteFunc: e.deletePod})
	e.podLister = podInformer.Lister()
	e.podsSynced = podInformer.Informer().HasSynced
	e.endpointsLister = endpointsInformer.Lister()
	e.endpointsSynced = endpointsInformer.Informer().HasSynced
	return e
}

type EndpointController struct {
	client           clientset.Interface
	serviceLister    corelisters.ServiceLister
	servicesSynced   cache.InformerSynced
	podLister        corelisters.PodLister
	podsSynced       cache.InformerSynced
	endpointsLister  corelisters.EndpointsLister
	endpointsSynced  cache.InformerSynced
	queue            workqueue.RateLimitingInterface
	workerLoopPeriod time.Duration
}

func (e *EndpointController) Run(workers int, stopCh <-chan struct{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	defer utilruntime.HandleCrash()
	defer e.queue.ShutDown()
	klog.Infof("Starting endpoint controller")
	defer klog.Infof("Shutting down endpoint controller")
	if !controller.WaitForCacheSync("endpoint", stopCh, e.podsSynced, e.servicesSynced, e.endpointsSynced) {
		return
	}
	for i := 0; i < workers; i++ {
		go wait.Until(e.worker, e.workerLoopPeriod, stopCh)
	}
	go func() {
		defer utilruntime.HandleCrash()
		e.checkLeftoverEndpoints()
	}()
	<-stopCh
}
func (e *EndpointController) getPodServiceMemberships(pod *v1.Pod) (sets.String, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	set := sets.String{}
	services, err := e.serviceLister.GetPodServices(pod)
	if err != nil {
		return set, nil
	}
	for i := range services {
		key, err := controller.KeyFunc(services[i])
		if err != nil {
			return nil, err
		}
		set.Insert(key)
	}
	return set, nil
}
func (e *EndpointController) addPod(obj interface{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	pod := obj.(*v1.Pod)
	services, err := e.getPodServiceMemberships(pod)
	if err != nil {
		utilruntime.HandleError(fmt.Errorf("Unable to get pod %s/%s's service memberships: %v", pod.Namespace, pod.Name, err))
		return
	}
	for key := range services {
		e.queue.Add(key)
	}
}
func podToEndpointAddress(pod *v1.Pod) *v1.EndpointAddress {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &v1.EndpointAddress{IP: pod.Status.PodIP, NodeName: &pod.Spec.NodeName, TargetRef: &v1.ObjectReference{Kind: "Pod", Namespace: pod.ObjectMeta.Namespace, Name: pod.ObjectMeta.Name, UID: pod.ObjectMeta.UID, ResourceVersion: pod.ObjectMeta.ResourceVersion}}
}
func podChanged(oldPod, newPod *v1.Pod) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if newPod.DeletionTimestamp != oldPod.DeletionTimestamp {
		return true
	}
	if podutil.IsPodReady(oldPod) != podutil.IsPodReady(newPod) {
		return true
	}
	newEndpointAddress := podToEndpointAddress(newPod)
	oldEndpointAddress := podToEndpointAddress(oldPod)
	newEndpointAddress.TargetRef.ResourceVersion = ""
	oldEndpointAddress.TargetRef.ResourceVersion = ""
	if reflect.DeepEqual(newEndpointAddress, oldEndpointAddress) {
		return false
	}
	return true
}
func determineNeededServiceUpdates(oldServices, services sets.String, podChanged bool) sets.String {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if podChanged {
		services = services.Union(oldServices)
	} else {
		services = services.Difference(oldServices).Union(oldServices.Difference(services))
	}
	return services
}
func (e *EndpointController) updatePod(old, cur interface{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	newPod := cur.(*v1.Pod)
	oldPod := old.(*v1.Pod)
	if newPod.ResourceVersion == oldPod.ResourceVersion {
		return
	}
	podChangedFlag := podChanged(oldPod, newPod)
	labelsChanged := false
	if !reflect.DeepEqual(newPod.Labels, oldPod.Labels) || !hostNameAndDomainAreEqual(newPod, oldPod) {
		labelsChanged = true
	}
	if !podChangedFlag && !labelsChanged {
		return
	}
	services, err := e.getPodServiceMemberships(newPod)
	if err != nil {
		utilruntime.HandleError(fmt.Errorf("Unable to get pod %v/%v's service memberships: %v", newPod.Namespace, newPod.Name, err))
		return
	}
	if labelsChanged {
		oldServices, err := e.getPodServiceMemberships(oldPod)
		if err != nil {
			utilruntime.HandleError(fmt.Errorf("Unable to get pod %v/%v's service memberships: %v", oldPod.Namespace, oldPod.Name, err))
			return
		}
		services = determineNeededServiceUpdates(oldServices, services, podChangedFlag)
	}
	for key := range services {
		e.queue.Add(key)
	}
}
func hostNameAndDomainAreEqual(pod1, pod2 *v1.Pod) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return pod1.Spec.Hostname == pod2.Spec.Hostname && pod1.Spec.Subdomain == pod2.Spec.Subdomain
}
func (e *EndpointController) deletePod(obj interface{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if _, ok := obj.(*v1.Pod); ok {
		e.addPod(obj)
		return
	}
	tombstone, ok := obj.(cache.DeletedFinalStateUnknown)
	if !ok {
		utilruntime.HandleError(fmt.Errorf("Couldn't get object from tombstone %#v", obj))
		return
	}
	pod, ok := tombstone.Obj.(*v1.Pod)
	if !ok {
		utilruntime.HandleError(fmt.Errorf("Tombstone contained object that is not a Pod: %#v", obj))
		return
	}
	klog.V(4).Infof("Enqueuing services of deleted pod %s/%s having final state unrecorded", pod.Namespace, pod.Name)
	e.addPod(pod)
}
func (e *EndpointController) enqueueService(obj interface{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	key, err := controller.KeyFunc(obj)
	if err != nil {
		utilruntime.HandleError(fmt.Errorf("Couldn't get key for object %+v: %v", obj, err))
		return
	}
	e.queue.Add(key)
}
func (e *EndpointController) worker() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for e.processNextWorkItem() {
	}
}
func (e *EndpointController) processNextWorkItem() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	eKey, quit := e.queue.Get()
	if quit {
		return false
	}
	defer e.queue.Done(eKey)
	err := e.syncService(eKey.(string))
	e.handleErr(err, eKey)
	return true
}
func (e *EndpointController) handleErr(err error, key interface{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if err == nil {
		e.queue.Forget(key)
		return
	}
	if e.queue.NumRequeues(key) < maxRetries {
		klog.V(2).Infof("Error syncing endpoints for service %q, retrying. Error: %v", key, err)
		e.queue.AddRateLimited(key)
		return
	}
	klog.Warningf("Dropping service %q out of the queue: %v", key, err)
	e.queue.Forget(key)
	utilruntime.HandleError(err)
}
func (e *EndpointController) syncService(key string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	startTime := time.Now()
	defer func() {
		klog.V(4).Infof("Finished syncing service %q endpoints. (%v)", key, time.Since(startTime))
	}()
	namespace, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		return err
	}
	service, err := e.serviceLister.Services(namespace).Get(name)
	if err != nil {
		err = e.client.CoreV1().Endpoints(namespace).Delete(name, nil)
		if err != nil && !errors.IsNotFound(err) {
			return err
		}
		return nil
	}
	if service.Spec.Selector == nil {
		return nil
	}
	klog.V(5).Infof("About to update endpoints for service %q", key)
	pods, err := e.podLister.Pods(service.Namespace).List(labels.Set(service.Spec.Selector).AsSelectorPreValidated())
	if err != nil {
		return err
	}
	tolerateUnreadyEndpoints := service.Spec.PublishNotReadyAddresses
	if v, ok := service.Annotations[TolerateUnreadyEndpointsAnnotation]; ok {
		b, err := strconv.ParseBool(v)
		if err == nil {
			tolerateUnreadyEndpoints = b
		} else {
			utilruntime.HandleError(fmt.Errorf("Failed to parse annotation %v: %v", TolerateUnreadyEndpointsAnnotation, err))
		}
	}
	subsets := []v1.EndpointSubset{}
	var totalReadyEps int = 0
	var totalNotReadyEps int = 0
	for _, pod := range pods {
		if len(pod.Status.PodIP) == 0 {
			klog.V(5).Infof("Failed to find an IP for pod %s/%s", pod.Namespace, pod.Name)
			continue
		}
		if !tolerateUnreadyEndpoints && pod.DeletionTimestamp != nil {
			klog.V(5).Infof("Pod is being deleted %s/%s", pod.Namespace, pod.Name)
			continue
		}
		epa := *podToEndpointAddress(pod)
		hostname := pod.Spec.Hostname
		if len(hostname) > 0 && pod.Spec.Subdomain == service.Name && service.Namespace == pod.Namespace {
			epa.Hostname = hostname
		}
		if len(service.Spec.Ports) == 0 {
			if service.Spec.ClusterIP == api.ClusterIPNone {
				subsets, totalReadyEps, totalNotReadyEps = addEndpointSubset(subsets, pod, epa, nil, tolerateUnreadyEndpoints)
			}
		} else {
			for i := range service.Spec.Ports {
				servicePort := &service.Spec.Ports[i]
				portName := servicePort.Name
				portProto := servicePort.Protocol
				portNum, err := podutil.FindPort(pod, servicePort)
				if err != nil {
					klog.V(4).Infof("Failed to find port for service %s/%s: %v", service.Namespace, service.Name, err)
					continue
				}
				var readyEps, notReadyEps int
				epp := &v1.EndpointPort{Name: portName, Port: int32(portNum), Protocol: portProto}
				subsets, readyEps, notReadyEps = addEndpointSubset(subsets, pod, epa, epp, tolerateUnreadyEndpoints)
				totalReadyEps = totalReadyEps + readyEps
				totalNotReadyEps = totalNotReadyEps + notReadyEps
			}
		}
	}
	subsets = endpoints.RepackSubsets(subsets)
	currentEndpoints, err := e.endpointsLister.Endpoints(service.Namespace).Get(service.Name)
	if err != nil {
		if errors.IsNotFound(err) {
			currentEndpoints = &v1.Endpoints{ObjectMeta: metav1.ObjectMeta{Name: service.Name, Labels: service.Labels}}
		} else {
			return err
		}
	}
	createEndpoints := len(currentEndpoints.ResourceVersion) == 0
	if !createEndpoints && apiequality.Semantic.DeepEqual(currentEndpoints.Subsets, subsets) && apiequality.Semantic.DeepEqual(currentEndpoints.Labels, service.Labels) {
		klog.V(5).Infof("endpoints are equal for %s/%s, skipping update", service.Namespace, service.Name)
		return nil
	}
	newEndpoints := currentEndpoints.DeepCopy()
	newEndpoints.Subsets = subsets
	newEndpoints.Labels = service.Labels
	if newEndpoints.Annotations == nil {
		newEndpoints.Annotations = make(map[string]string)
	}
	klog.V(4).Infof("Update endpoints for %v/%v, ready: %d not ready: %d", service.Namespace, service.Name, totalReadyEps, totalNotReadyEps)
	if createEndpoints {
		_, err = e.client.CoreV1().Endpoints(service.Namespace).Create(newEndpoints)
	} else {
		_, err = e.client.CoreV1().Endpoints(service.Namespace).Update(newEndpoints)
	}
	if err != nil {
		if createEndpoints && errors.IsForbidden(err) {
			klog.V(5).Infof("Forbidden from creating endpoints: %v", err)
		}
		return err
	}
	return nil
}
func (e *EndpointController) checkLeftoverEndpoints() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	list, err := e.endpointsLister.List(labels.Everything())
	if err != nil {
		utilruntime.HandleError(fmt.Errorf("Unable to list endpoints (%v); orphaned endpoints will not be cleaned up. (They're pretty harmless, but you can restart this component if you want another attempt made.)", err))
		return
	}
	for _, ep := range list {
		if _, ok := ep.Annotations[resourcelock.LeaderElectionRecordAnnotationKey]; ok {
			continue
		}
		key, err := controller.KeyFunc(ep)
		if err != nil {
			utilruntime.HandleError(fmt.Errorf("Unable to get key for endpoint %#v", ep))
			continue
		}
		e.queue.Add(key)
	}
}
func addEndpointSubset(subsets []v1.EndpointSubset, pod *v1.Pod, epa v1.EndpointAddress, epp *v1.EndpointPort, tolerateUnreadyEndpoints bool) ([]v1.EndpointSubset, int, int) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var readyEps int = 0
	var notReadyEps int = 0
	ports := []v1.EndpointPort{}
	if epp != nil {
		ports = append(ports, *epp)
	}
	if tolerateUnreadyEndpoints || podutil.IsPodReady(pod) {
		subsets = append(subsets, v1.EndpointSubset{Addresses: []v1.EndpointAddress{epa}, Ports: ports})
		readyEps++
	} else if shouldPodBeInEndpoints(pod) {
		klog.V(5).Infof("Pod is out of service: %s/%s", pod.Namespace, pod.Name)
		subsets = append(subsets, v1.EndpointSubset{NotReadyAddresses: []v1.EndpointAddress{epa}, Ports: ports})
		notReadyEps++
	}
	return subsets, readyEps, notReadyEps
}
func shouldPodBeInEndpoints(pod *v1.Pod) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	switch pod.Spec.RestartPolicy {
	case v1.RestartPolicyNever:
		return pod.Status.Phase != v1.PodFailed && pod.Status.Phase != v1.PodSucceeded
	case v1.RestartPolicyOnFailure:
		return pod.Status.Phase != v1.PodSucceeded
	default:
		return true
	}
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
