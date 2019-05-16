package service

import (
	"context"
	"fmt"
	goformat "fmt"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apimachinery/pkg/util/wait"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	coreinformers "k8s.io/client-go/informers/core/v1"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	v1core "k8s.io/client-go/kubernetes/typed/core/v1"
	corelisters "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/workqueue"
	cloudprovider "k8s.io/cloud-provider"
	"k8s.io/klog"
	v1helper "k8s.io/kubernetes/pkg/apis/core/v1/helper"
	"k8s.io/kubernetes/pkg/controller"
	kubefeatures "k8s.io/kubernetes/pkg/features"
	"k8s.io/kubernetes/pkg/util/metrics"
	goos "os"
	"reflect"
	godefaultruntime "runtime"
	"sync"
	"time"
	gotime "time"
)

const (
	serviceSyncPeriod            = 30 * time.Second
	nodeSyncPeriod               = 100 * time.Second
	minRetryDelay                = 5 * time.Second
	maxRetryDelay                = 300 * time.Second
	clientRetryCount             = 5
	clientRetryInterval          = 5 * time.Second
	LabelNodeRoleMaster          = "node-role.kubernetes.io/master"
	LabelNodeRoleExcludeBalancer = "alpha.service-controller.kubernetes.io/exclude-balancer"
)

type cachedService struct{ state *v1.Service }
type serviceCache struct {
	mu         sync.Mutex
	serviceMap map[string]*cachedService
}
type ServiceController struct {
	cloud               cloudprovider.Interface
	knownHosts          []*v1.Node
	servicesToUpdate    []*v1.Service
	kubeClient          clientset.Interface
	clusterName         string
	balancer            cloudprovider.LoadBalancer
	cache               *serviceCache
	serviceLister       corelisters.ServiceLister
	serviceListerSynced cache.InformerSynced
	eventBroadcaster    record.EventBroadcaster
	eventRecorder       record.EventRecorder
	nodeLister          corelisters.NodeLister
	nodeListerSynced    cache.InformerSynced
	queue               workqueue.RateLimitingInterface
}

func New(cloud cloudprovider.Interface, kubeClient clientset.Interface, serviceInformer coreinformers.ServiceInformer, nodeInformer coreinformers.NodeInformer, clusterName string) (*ServiceController, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	broadcaster := record.NewBroadcaster()
	broadcaster.StartLogging(klog.Infof)
	broadcaster.StartRecordingToSink(&v1core.EventSinkImpl{Interface: kubeClient.CoreV1().Events("")})
	recorder := broadcaster.NewRecorder(scheme.Scheme, v1.EventSource{Component: "service-controller"})
	if kubeClient != nil && kubeClient.CoreV1().RESTClient().GetRateLimiter() != nil {
		if err := metrics.RegisterMetricAndTrackRateLimiterUsage("service_controller", kubeClient.CoreV1().RESTClient().GetRateLimiter()); err != nil {
			return nil, err
		}
	}
	s := &ServiceController{cloud: cloud, knownHosts: []*v1.Node{}, kubeClient: kubeClient, clusterName: clusterName, cache: &serviceCache{serviceMap: make(map[string]*cachedService)}, eventBroadcaster: broadcaster, eventRecorder: recorder, nodeLister: nodeInformer.Lister(), nodeListerSynced: nodeInformer.Informer().HasSynced, queue: workqueue.NewNamedRateLimitingQueue(workqueue.NewItemExponentialFailureRateLimiter(minRetryDelay, maxRetryDelay), "service")}
	serviceInformer.Informer().AddEventHandlerWithResyncPeriod(cache.ResourceEventHandlerFuncs{AddFunc: s.enqueueService, UpdateFunc: func(old, cur interface{}) {
		oldSvc, ok1 := old.(*v1.Service)
		curSvc, ok2 := cur.(*v1.Service)
		if ok1 && ok2 && s.needsUpdate(oldSvc, curSvc) {
			s.enqueueService(cur)
		}
	}, DeleteFunc: s.enqueueService}, serviceSyncPeriod)
	s.serviceLister = serviceInformer.Lister()
	s.serviceListerSynced = serviceInformer.Informer().HasSynced
	if err := s.init(); err != nil {
		return nil, err
	}
	return s, nil
}
func (s *ServiceController) enqueueService(obj interface{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	key, err := controller.KeyFunc(obj)
	if err != nil {
		klog.Errorf("Couldn't get key for object %#v: %v", obj, err)
		return
	}
	s.queue.Add(key)
}
func (s *ServiceController) Run(stopCh <-chan struct{}, workers int) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	defer runtime.HandleCrash()
	defer s.queue.ShutDown()
	klog.Info("Starting service controller")
	defer klog.Info("Shutting down service controller")
	if !controller.WaitForCacheSync("service", stopCh, s.serviceListerSynced, s.nodeListerSynced) {
		return
	}
	for i := 0; i < workers; i++ {
		go wait.Until(s.worker, time.Second, stopCh)
	}
	go wait.Until(s.nodeSyncLoop, nodeSyncPeriod, stopCh)
	<-stopCh
}
func (s *ServiceController) worker() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for s.processNextWorkItem() {
	}
}
func (s *ServiceController) processNextWorkItem() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	key, quit := s.queue.Get()
	if quit {
		return false
	}
	defer s.queue.Done(key)
	err := s.syncService(key.(string))
	if err == nil {
		s.queue.Forget(key)
		return true
	}
	runtime.HandleError(fmt.Errorf("error processing service %v (will retry): %v", key, err))
	s.queue.AddRateLimited(key)
	return true
}
func (s *ServiceController) init() error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if s.cloud == nil {
		return fmt.Errorf("WARNING: no cloud provider provided, services of type LoadBalancer will fail")
	}
	balancer, ok := s.cloud.LoadBalancer()
	if !ok {
		return fmt.Errorf("the cloud provider does not support external load balancers")
	}
	s.balancer = balancer
	return nil
}
func (s *ServiceController) processServiceUpdate(cachedService *cachedService, service *v1.Service, key string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if cachedService.state != nil {
		if cachedService.state.UID != service.UID {
			err := s.processLoadBalancerDelete(cachedService, key)
			if err != nil {
				return err
			}
		}
	}
	cachedService.state = service
	err := s.createLoadBalancerIfNeeded(key, service)
	if err != nil {
		eventType := "CreatingLoadBalancerFailed"
		message := "Error creating load balancer (will retry): "
		if !wantsLoadBalancer(service) {
			eventType = "CleanupLoadBalancerFailed"
			message = "Error cleaning up load balancer (will retry): "
		}
		message += err.Error()
		s.eventRecorder.Event(service, v1.EventTypeWarning, eventType, message)
		return err
	}
	s.cache.set(key, cachedService)
	return nil
}
func (s *ServiceController) createLoadBalancerIfNeeded(key string, service *v1.Service) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	previousState := v1helper.LoadBalancerStatusDeepCopy(&service.Status.LoadBalancer)
	var newState *v1.LoadBalancerStatus
	var err error
	if !wantsLoadBalancer(service) {
		if v1helper.LoadBalancerStatusEqual(previousState, &v1.LoadBalancerStatus{}) {
			return nil
		}
		klog.V(3).Infof("Getting load balancer for service %s", key)
		_, exists, err := s.balancer.GetLoadBalancer(context.TODO(), s.clusterName, service)
		if err != nil {
			return fmt.Errorf("error getting LB for service %s: %v", key, err)
		}
		if exists {
			klog.Infof("Deleting existing load balancer for service %s that no longer needs a load balancer.", key)
			s.eventRecorder.Event(service, v1.EventTypeNormal, "DeletingLoadBalancer", "Deleting load balancer")
			if err := s.balancer.EnsureLoadBalancerDeleted(context.TODO(), s.clusterName, service); err != nil {
				return err
			}
			s.eventRecorder.Event(service, v1.EventTypeNormal, "DeletedLoadBalancer", "Deleted load balancer")
		}
		newState = &v1.LoadBalancerStatus{}
	} else {
		klog.V(2).Infof("Ensuring LB for service %s", key)
		s.eventRecorder.Event(service, v1.EventTypeNormal, "EnsuringLoadBalancer", "Ensuring load balancer")
		newState, err = s.ensureLoadBalancer(service)
		if err != nil {
			return fmt.Errorf("failed to ensure load balancer for service %s: %v", key, err)
		}
		s.eventRecorder.Event(service, v1.EventTypeNormal, "EnsuredLoadBalancer", "Ensured load balancer")
	}
	if !v1helper.LoadBalancerStatusEqual(previousState, newState) {
		service = service.DeepCopy()
		service.Status.LoadBalancer = *newState
		if err := s.persistUpdate(service); err != nil {
			if errors.IsConflict(err) {
				return fmt.Errorf("not persisting update to service '%s/%s' that has been changed since we received it: %v", service.Namespace, service.Name, err)
			}
			runtime.HandleError(fmt.Errorf("failed to persist service %q updated status to apiserver, even after retries. Giving up: %v", key, err))
			return nil
		}
	} else {
		klog.V(2).Infof("Not persisting unchanged LoadBalancerStatus for service %s to registry.", key)
	}
	return nil
}
func (s *ServiceController) persistUpdate(service *v1.Service) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var err error
	for i := 0; i < clientRetryCount; i++ {
		_, err = s.kubeClient.CoreV1().Services(service.Namespace).UpdateStatus(service)
		if err == nil {
			return nil
		}
		if errors.IsNotFound(err) {
			klog.Infof("Not persisting update to service '%s/%s' that no longer exists: %v", service.Namespace, service.Name, err)
			return nil
		}
		if errors.IsConflict(err) {
			return err
		}
		klog.Warningf("Failed to persist updated LoadBalancerStatus to service '%s/%s' after creating its load balancer: %v", service.Namespace, service.Name, err)
		time.Sleep(clientRetryInterval)
	}
	return err
}
func (s *ServiceController) ensureLoadBalancer(service *v1.Service) (*v1.LoadBalancerStatus, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	nodes, err := s.nodeLister.ListWithPredicate(getNodeConditionPredicate())
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		s.eventRecorder.Eventf(service, v1.EventTypeWarning, "UnAvailableLoadBalancer", "There are no available nodes for LoadBalancer service %s/%s", service.Namespace, service.Name)
	}
	return s.balancer.EnsureLoadBalancer(context.TODO(), s.clusterName, service, nodes)
}
func (s *serviceCache) ListKeys() []string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	s.mu.Lock()
	defer s.mu.Unlock()
	keys := make([]string, 0, len(s.serviceMap))
	for k := range s.serviceMap {
		keys = append(keys, k)
	}
	return keys
}
func (s *serviceCache) GetByKey(key string) (interface{}, bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	s.mu.Lock()
	defer s.mu.Unlock()
	if v, ok := s.serviceMap[key]; ok {
		return v, true, nil
	}
	return nil, false, nil
}
func (s *serviceCache) allServices() []*v1.Service {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	s.mu.Lock()
	defer s.mu.Unlock()
	services := make([]*v1.Service, 0, len(s.serviceMap))
	for _, v := range s.serviceMap {
		services = append(services, v.state)
	}
	return services
}
func (s *serviceCache) get(serviceName string) (*cachedService, bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	s.mu.Lock()
	defer s.mu.Unlock()
	service, ok := s.serviceMap[serviceName]
	return service, ok
}
func (s *serviceCache) getOrCreate(serviceName string) *cachedService {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	s.mu.Lock()
	defer s.mu.Unlock()
	service, ok := s.serviceMap[serviceName]
	if !ok {
		service = &cachedService{}
		s.serviceMap[serviceName] = service
	}
	return service
}
func (s *serviceCache) set(serviceName string, service *cachedService) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	s.mu.Lock()
	defer s.mu.Unlock()
	s.serviceMap[serviceName] = service
}
func (s *serviceCache) delete(serviceName string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.serviceMap, serviceName)
}
func (s *ServiceController) needsUpdate(oldService *v1.Service, newService *v1.Service) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if !wantsLoadBalancer(oldService) && !wantsLoadBalancer(newService) {
		return false
	}
	if wantsLoadBalancer(oldService) != wantsLoadBalancer(newService) {
		s.eventRecorder.Eventf(newService, v1.EventTypeNormal, "Type", "%v -> %v", oldService.Spec.Type, newService.Spec.Type)
		return true
	}
	if wantsLoadBalancer(newService) && !reflect.DeepEqual(oldService.Spec.LoadBalancerSourceRanges, newService.Spec.LoadBalancerSourceRanges) {
		s.eventRecorder.Eventf(newService, v1.EventTypeNormal, "LoadBalancerSourceRanges", "%v -> %v", oldService.Spec.LoadBalancerSourceRanges, newService.Spec.LoadBalancerSourceRanges)
		return true
	}
	if !portsEqualForLB(oldService, newService) || oldService.Spec.SessionAffinity != newService.Spec.SessionAffinity {
		return true
	}
	if !loadBalancerIPsAreEqual(oldService, newService) {
		s.eventRecorder.Eventf(newService, v1.EventTypeNormal, "LoadbalancerIP", "%v -> %v", oldService.Spec.LoadBalancerIP, newService.Spec.LoadBalancerIP)
		return true
	}
	if len(oldService.Spec.ExternalIPs) != len(newService.Spec.ExternalIPs) {
		s.eventRecorder.Eventf(newService, v1.EventTypeNormal, "ExternalIP", "Count: %v -> %v", len(oldService.Spec.ExternalIPs), len(newService.Spec.ExternalIPs))
		return true
	}
	for i := range oldService.Spec.ExternalIPs {
		if oldService.Spec.ExternalIPs[i] != newService.Spec.ExternalIPs[i] {
			s.eventRecorder.Eventf(newService, v1.EventTypeNormal, "ExternalIP", "Added: %v", newService.Spec.ExternalIPs[i])
			return true
		}
	}
	if !reflect.DeepEqual(oldService.Annotations, newService.Annotations) {
		return true
	}
	if oldService.UID != newService.UID {
		s.eventRecorder.Eventf(newService, v1.EventTypeNormal, "UID", "%v -> %v", oldService.UID, newService.UID)
		return true
	}
	if oldService.Spec.ExternalTrafficPolicy != newService.Spec.ExternalTrafficPolicy {
		s.eventRecorder.Eventf(newService, v1.EventTypeNormal, "ExternalTrafficPolicy", "%v -> %v", oldService.Spec.ExternalTrafficPolicy, newService.Spec.ExternalTrafficPolicy)
		return true
	}
	if oldService.Spec.HealthCheckNodePort != newService.Spec.HealthCheckNodePort {
		s.eventRecorder.Eventf(newService, v1.EventTypeNormal, "HealthCheckNodePort", "%v -> %v", oldService.Spec.HealthCheckNodePort, newService.Spec.HealthCheckNodePort)
		return true
	}
	return false
}
func (s *ServiceController) loadBalancerName(service *v1.Service) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return s.balancer.GetLoadBalancerName(context.TODO(), "", service)
}
func getPortsForLB(service *v1.Service) ([]*v1.ServicePort, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var protocol v1.Protocol
	ports := []*v1.ServicePort{}
	for i := range service.Spec.Ports {
		sp := &service.Spec.Ports[i]
		ports = append(ports, sp)
		if protocol == "" {
			protocol = sp.Protocol
		} else if protocol != sp.Protocol && wantsLoadBalancer(service) {
			return nil, fmt.Errorf("mixed protocol external load balancers are not supported")
		}
	}
	return ports, nil
}
func portsEqualForLB(x, y *v1.Service) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	xPorts, err := getPortsForLB(x)
	if err != nil {
		return false
	}
	yPorts, err := getPortsForLB(y)
	if err != nil {
		return false
	}
	return portSlicesEqualForLB(xPorts, yPorts)
}
func portSlicesEqualForLB(x, y []*v1.ServicePort) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if !portEqualForLB(x[i], y[i]) {
			return false
		}
	}
	return true
}
func portEqualForLB(x, y *v1.ServicePort) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if x.Name != y.Name {
		return false
	}
	if x.Protocol != y.Protocol {
		return false
	}
	if x.Port != y.Port {
		return false
	}
	if x.NodePort != y.NodePort {
		return false
	}
	return true
}
func nodeNames(nodes []*v1.Node) sets.String {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ret := sets.NewString()
	for _, node := range nodes {
		ret.Insert(node.Name)
	}
	return ret
}
func nodeSlicesEqualForLB(x, y []*v1.Node) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(x) != len(y) {
		return false
	}
	return nodeNames(x).Equal(nodeNames(y))
}
func getNodeConditionPredicate() corelisters.NodeConditionPredicate {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return func(node *v1.Node) bool {
		if node.Spec.Unschedulable {
			return false
		}
		if _, hasMasterRoleLabel := node.Labels[LabelNodeRoleMaster]; hasMasterRoleLabel {
			return false
		}
		if utilfeature.DefaultFeatureGate.Enabled(kubefeatures.ServiceNodeExclusion) {
			if _, hasExcludeBalancerLabel := node.Labels[LabelNodeRoleExcludeBalancer]; hasExcludeBalancerLabel {
				return false
			}
		}
		if len(node.Status.Conditions) == 0 {
			return false
		}
		for _, cond := range node.Status.Conditions {
			if cond.Type == v1.NodeReady && cond.Status != v1.ConditionTrue {
				klog.V(4).Infof("Ignoring node %v with %v condition status %v", node.Name, cond.Type, cond.Status)
				return false
			}
		}
		return true
	}
}
func (s *ServiceController) nodeSyncLoop() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	newHosts, err := s.nodeLister.ListWithPredicate(getNodeConditionPredicate())
	if err != nil {
		klog.Errorf("Failed to retrieve current set of nodes from node lister: %v", err)
		return
	}
	if nodeSlicesEqualForLB(newHosts, s.knownHosts) {
		s.servicesToUpdate = s.updateLoadBalancerHosts(s.servicesToUpdate, newHosts)
		return
	}
	klog.Infof("Detected change in list of current cluster nodes. New node set: %v", nodeNames(newHosts))
	s.servicesToUpdate = s.cache.allServices()
	numServices := len(s.servicesToUpdate)
	s.servicesToUpdate = s.updateLoadBalancerHosts(s.servicesToUpdate, newHosts)
	klog.Infof("Successfully updated %d out of %d load balancers to direct traffic to the updated set of nodes", numServices-len(s.servicesToUpdate), numServices)
	s.knownHosts = newHosts
}
func (s *ServiceController) updateLoadBalancerHosts(services []*v1.Service, hosts []*v1.Node) (servicesToRetry []*v1.Service) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for _, service := range services {
		func() {
			if service == nil {
				return
			}
			if err := s.lockedUpdateLoadBalancerHosts(service, hosts); err != nil {
				klog.Errorf("External error while updating load balancer: %v.", err)
				servicesToRetry = append(servicesToRetry, service)
			}
		}()
	}
	return servicesToRetry
}
func (s *ServiceController) lockedUpdateLoadBalancerHosts(service *v1.Service, hosts []*v1.Node) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if !wantsLoadBalancer(service) {
		return nil
	}
	err := s.balancer.UpdateLoadBalancer(context.TODO(), s.clusterName, service, hosts)
	if err == nil {
		if len(hosts) == 0 {
			s.eventRecorder.Eventf(service, v1.EventTypeWarning, "UnAvailableLoadBalancer", "There are no available nodes for LoadBalancer service %s/%s", service.Namespace, service.Name)
		} else {
			s.eventRecorder.Event(service, v1.EventTypeNormal, "UpdatedLoadBalancer", "Updated load balancer with new hosts")
		}
		return nil
	}
	if _, exists, err := s.balancer.GetLoadBalancer(context.TODO(), s.clusterName, service); err != nil {
		klog.Errorf("External error while checking if load balancer %q exists: name, %v", s.balancer.GetLoadBalancerName(context.TODO(), s.clusterName, service), err)
	} else if !exists {
		return nil
	}
	s.eventRecorder.Eventf(service, v1.EventTypeWarning, "LoadBalancerUpdateFailed", "Error updating load balancer with new hosts %v: %v", nodeNames(hosts), err)
	return err
}
func wantsLoadBalancer(service *v1.Service) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return service.Spec.Type == v1.ServiceTypeLoadBalancer
}
func loadBalancerIPsAreEqual(oldService, newService *v1.Service) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return oldService.Spec.LoadBalancerIP == newService.Spec.LoadBalancerIP
}
func (s *ServiceController) syncService(key string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	startTime := time.Now()
	var cachedService *cachedService
	defer func() {
		klog.V(4).Infof("Finished syncing service %q (%v)", key, time.Since(startTime))
	}()
	namespace, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		return err
	}
	service, err := s.serviceLister.Services(namespace).Get(name)
	switch {
	case errors.IsNotFound(err):
		klog.Infof("Service has been deleted %v. Attempting to cleanup load balancer resources", key)
		err = s.processServiceDeletion(key)
	case err != nil:
		klog.Infof("Unable to retrieve service %v from store: %v", key, err)
	default:
		cachedService = s.cache.getOrCreate(key)
		err = s.processServiceUpdate(cachedService, service, key)
	}
	return err
}
func (s *ServiceController) processServiceDeletion(key string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	cachedService, ok := s.cache.get(key)
	if !ok {
		klog.Errorf("service %s not in cache even though the watcher thought it was. Ignoring the deletion", key)
		return nil
	}
	return s.processLoadBalancerDelete(cachedService, key)
}
func (s *ServiceController) processLoadBalancerDelete(cachedService *cachedService, key string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	service := cachedService.state
	if !wantsLoadBalancer(service) {
		return nil
	}
	s.eventRecorder.Event(service, v1.EventTypeNormal, "DeletingLoadBalancer", "Deleting load balancer")
	err := s.balancer.EnsureLoadBalancerDeleted(context.TODO(), s.clusterName, service)
	if err != nil {
		s.eventRecorder.Eventf(service, v1.EventTypeWarning, "DeletingLoadBalancerFailed", "Error deleting load balancer (will retry): %v", err)
		return err
	}
	s.eventRecorder.Event(service, v1.EventTypeNormal, "DeletedLoadBalancer", "Deleted load balancer")
	s.cache.delete(key)
	return nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
