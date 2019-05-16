package storage

import (
	"context"
	"fmt"
	goformat "fmt"
	"k8s.io/apimachinery/pkg/api/errors"
	metainternalversion "k8s.io/apimachinery/pkg/apis/meta/internalversion"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	metav1beta1 "k8s.io/apimachinery/pkg/apis/meta/v1beta1"
	"k8s.io/apimachinery/pkg/runtime"
	utilnet "k8s.io/apimachinery/pkg/util/net"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apimachinery/pkg/watch"
	genericapirequest "k8s.io/apiserver/pkg/endpoints/request"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/apiserver/pkg/util/dryrun"
	"k8s.io/klog"
	apiservice "k8s.io/kubernetes/pkg/api/service"
	api "k8s.io/kubernetes/pkg/apis/core"
	"k8s.io/kubernetes/pkg/apis/core/helper"
	"k8s.io/kubernetes/pkg/apis/core/validation"
	registry "k8s.io/kubernetes/pkg/registry/core/service"
	"k8s.io/kubernetes/pkg/registry/core/service/ipallocator"
	"k8s.io/kubernetes/pkg/registry/core/service/portallocator"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	goos "os"
	godefaultruntime "runtime"
	"strconv"
	gotime "time"
)

type REST struct {
	services         ServiceStorage
	endpoints        EndpointsStorage
	serviceIPs       ipallocator.Interface
	serviceNodePorts portallocator.Interface
	proxyTransport   http.RoundTripper
	pods             rest.Getter
}
type ServiceNodePort struct {
	Protocol api.Protocol
	NodePort int32
}
type ServiceStorage interface {
	rest.Scoper
	rest.Getter
	rest.Lister
	rest.CreaterUpdater
	rest.GracefulDeleter
	rest.Watcher
	rest.TableConvertor
	rest.Exporter
}
type EndpointsStorage interface {
	rest.Getter
	rest.GracefulDeleter
}

func NewREST(services ServiceStorage, endpoints EndpointsStorage, pods rest.Getter, serviceIPs ipallocator.Interface, serviceNodePorts portallocator.Interface, proxyTransport http.RoundTripper) (*REST, *registry.ProxyREST) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	rest := &REST{services: services, endpoints: endpoints, serviceIPs: serviceIPs, serviceNodePorts: serviceNodePorts, proxyTransport: proxyTransport, pods: pods}
	return rest, &registry.ProxyREST{Redirector: rest, ProxyTransport: proxyTransport}
}

var (
	_ ServiceStorage          = &REST{}
	_ rest.CategoriesProvider = &REST{}
	_ rest.ShortNamesProvider = &REST{}
)

func (rs *REST) ShortNames() []string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return []string{"svc"}
}
func (rs *REST) Categories() []string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return []string{"all"}
}
func (rs *REST) NamespaceScoped() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return rs.services.NamespaceScoped()
}
func (rs *REST) New() runtime.Object {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return rs.services.New()
}
func (rs *REST) NewList() runtime.Object {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return rs.services.NewList()
}
func (rs *REST) Get(ctx context.Context, name string, options *metav1.GetOptions) (runtime.Object, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return rs.services.Get(ctx, name, options)
}
func (rs *REST) List(ctx context.Context, options *metainternalversion.ListOptions) (runtime.Object, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return rs.services.List(ctx, options)
}
func (rs *REST) Watch(ctx context.Context, options *metainternalversion.ListOptions) (watch.Interface, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return rs.services.Watch(ctx, options)
}
func (rs *REST) Export(ctx context.Context, name string, opts metav1.ExportOptions) (runtime.Object, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return rs.services.Export(ctx, name, opts)
}
func (rs *REST) Create(ctx context.Context, obj runtime.Object, createValidation rest.ValidateObjectFunc, options *metav1.CreateOptions) (runtime.Object, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	service := obj.(*api.Service)
	if err := rest.BeforeCreate(registry.Strategy, ctx, obj); err != nil {
		return nil, err
	}
	releaseServiceIP := false
	defer func() {
		if releaseServiceIP {
			if helper.IsServiceIPSet(service) {
				rs.serviceIPs.Release(net.ParseIP(service.Spec.ClusterIP))
			}
		}
	}()
	var err error
	if !dryrun.IsDryRun(options.DryRun) {
		if service.Spec.Type != api.ServiceTypeExternalName {
			if releaseServiceIP, err = initClusterIP(service, rs.serviceIPs); err != nil {
				return nil, err
			}
		}
	}
	nodePortOp := portallocator.StartOperation(rs.serviceNodePorts, dryrun.IsDryRun(options.DryRun))
	defer nodePortOp.Finish()
	if service.Spec.Type == api.ServiceTypeNodePort || service.Spec.Type == api.ServiceTypeLoadBalancer {
		if err := initNodePorts(service, nodePortOp); err != nil {
			return nil, err
		}
	}
	if apiservice.NeedsHealthCheck(service) {
		if err := allocateHealthCheckNodePort(service, nodePortOp); err != nil {
			return nil, errors.NewInternalError(err)
		}
	}
	if errs := validation.ValidateServiceExternalTrafficFieldsCombination(service); len(errs) > 0 {
		return nil, errors.NewInvalid(api.Kind("Service"), service.Name, errs)
	}
	out, err := rs.services.Create(ctx, service, createValidation, options)
	if err != nil {
		err = rest.CheckGeneratedNameError(registry.Strategy, err, service)
	}
	if err == nil {
		el := nodePortOp.Commit()
		if el != nil {
			utilruntime.HandleError(fmt.Errorf("error(s) committing service node-ports changes: %v", el))
		}
		releaseServiceIP = false
	}
	return out, err
}
func (rs *REST) Delete(ctx context.Context, id string, options *metav1.DeleteOptions) (runtime.Object, bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	obj, _, err := rs.services.Delete(ctx, id, options)
	if err != nil {
		return nil, false, err
	}
	svc := obj.(*api.Service)
	if !dryrun.IsDryRun(options.DryRun) {
		_, _, err = rs.endpoints.Delete(ctx, id, &metav1.DeleteOptions{})
		if err != nil && !errors.IsNotFound(err) {
			return nil, false, err
		}
		rs.releaseAllocatedResources(svc)
	}
	details := &metav1.StatusDetails{Name: svc.Name, UID: svc.UID}
	if info, ok := genericapirequest.RequestInfoFrom(ctx); ok {
		details.Group = info.APIGroup
		details.Kind = info.Resource
	}
	status := &metav1.Status{Status: metav1.StatusSuccess, Details: details}
	return status, true, nil
}
func (rs *REST) releaseAllocatedResources(svc *api.Service) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if helper.IsServiceIPSet(svc) {
		rs.serviceIPs.Release(net.ParseIP(svc.Spec.ClusterIP))
	}
	for _, nodePort := range collectServiceNodePorts(svc) {
		err := rs.serviceNodePorts.Release(nodePort)
		if err != nil {
			utilruntime.HandleError(fmt.Errorf("Error releasing service %s node port %d: %v", svc.Name, nodePort, err))
		}
	}
	if apiservice.NeedsHealthCheck(svc) {
		nodePort := svc.Spec.HealthCheckNodePort
		if nodePort > 0 {
			err := rs.serviceNodePorts.Release(int(nodePort))
			if err != nil {
				utilruntime.HandleError(fmt.Errorf("Error releasing service %s health check node port %d: %v", svc.Name, nodePort, err))
			}
		}
	}
}
func externalTrafficPolicyUpdate(oldService, service *api.Service) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var neededExternalTraffic, needsExternalTraffic bool
	if oldService.Spec.Type == api.ServiceTypeNodePort || oldService.Spec.Type == api.ServiceTypeLoadBalancer {
		neededExternalTraffic = true
	}
	if service.Spec.Type == api.ServiceTypeNodePort || service.Spec.Type == api.ServiceTypeLoadBalancer {
		needsExternalTraffic = true
	}
	if neededExternalTraffic && !needsExternalTraffic {
		service.Spec.ExternalTrafficPolicy = api.ServiceExternalTrafficPolicyType("")
	}
}
func (rs *REST) healthCheckNodePortUpdate(oldService, service *api.Service, nodePortOp *portallocator.PortAllocationOperation) (bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	neededHealthCheckNodePort := apiservice.NeedsHealthCheck(oldService)
	oldHealthCheckNodePort := oldService.Spec.HealthCheckNodePort
	needsHealthCheckNodePort := apiservice.NeedsHealthCheck(service)
	newHealthCheckNodePort := service.Spec.HealthCheckNodePort
	switch {
	case !neededHealthCheckNodePort && needsHealthCheckNodePort:
		klog.Infof("Transition to LoadBalancer type service with ExternalTrafficPolicy=Local")
		if err := allocateHealthCheckNodePort(service, nodePortOp); err != nil {
			return false, errors.NewInternalError(err)
		}
	case neededHealthCheckNodePort && !needsHealthCheckNodePort:
		klog.Infof("Transition to non LoadBalancer type service or LoadBalancer type service with ExternalTrafficPolicy=Global")
		klog.V(4).Infof("Releasing healthCheckNodePort: %d", oldHealthCheckNodePort)
		nodePortOp.ReleaseDeferred(int(oldHealthCheckNodePort))
		service.Spec.HealthCheckNodePort = 0
	case neededHealthCheckNodePort && needsHealthCheckNodePort:
		if oldHealthCheckNodePort != newHealthCheckNodePort {
			klog.Warningf("Attempt to change value of health check node port DENIED")
			fldPath := field.NewPath("spec", "healthCheckNodePort")
			el := field.ErrorList{field.Invalid(fldPath, newHealthCheckNodePort, "cannot change healthCheckNodePort on loadBalancer service with externalTraffic=Local during update")}
			return false, errors.NewInvalid(api.Kind("Service"), service.Name, el)
		}
	}
	return true, nil
}
func (rs *REST) Update(ctx context.Context, name string, objInfo rest.UpdatedObjectInfo, createValidation rest.ValidateObjectFunc, updateValidation rest.ValidateObjectUpdateFunc, forceAllowCreate bool, options *metav1.UpdateOptions) (runtime.Object, bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	oldObj, err := rs.services.Get(ctx, name, &metav1.GetOptions{})
	if err != nil {
		return nil, false, err
	}
	oldService := oldObj.(*api.Service)
	obj, err := objInfo.UpdatedObject(ctx, oldService)
	if err != nil {
		return nil, false, err
	}
	service := obj.(*api.Service)
	if !rest.ValidNamespace(ctx, &service.ObjectMeta) {
		return nil, false, errors.NewConflict(api.Resource("services"), service.Namespace, fmt.Errorf("Service.Namespace does not match the provided context"))
	}
	if err := rest.BeforeUpdate(registry.Strategy, ctx, service, oldService); err != nil {
		return nil, false, err
	}
	releaseServiceIP := false
	defer func() {
		if releaseServiceIP {
			if helper.IsServiceIPSet(service) {
				rs.serviceIPs.Release(net.ParseIP(service.Spec.ClusterIP))
			}
		}
	}()
	nodePortOp := portallocator.StartOperation(rs.serviceNodePorts, dryrun.IsDryRun(options.DryRun))
	defer nodePortOp.Finish()
	if !dryrun.IsDryRun(options.DryRun) {
		if oldService.Spec.Type == api.ServiceTypeExternalName && service.Spec.Type != api.ServiceTypeExternalName {
			if releaseServiceIP, err = initClusterIP(service, rs.serviceIPs); err != nil {
				return nil, false, err
			}
		}
		if oldService.Spec.Type != api.ServiceTypeExternalName && service.Spec.Type == api.ServiceTypeExternalName {
			if helper.IsServiceIPSet(oldService) {
				rs.serviceIPs.Release(net.ParseIP(oldService.Spec.ClusterIP))
			}
		}
	}
	if (oldService.Spec.Type == api.ServiceTypeNodePort || oldService.Spec.Type == api.ServiceTypeLoadBalancer) && (service.Spec.Type == api.ServiceTypeExternalName || service.Spec.Type == api.ServiceTypeClusterIP) {
		releaseNodePorts(oldService, nodePortOp)
	}
	if service.Spec.Type == api.ServiceTypeNodePort || service.Spec.Type == api.ServiceTypeLoadBalancer {
		if err := updateNodePorts(oldService, service, nodePortOp); err != nil {
			return nil, false, err
		}
	}
	if service.Spec.Type != api.ServiceTypeLoadBalancer {
		service.Status.LoadBalancer = api.LoadBalancerStatus{}
	}
	success, err := rs.healthCheckNodePortUpdate(oldService, service, nodePortOp)
	if !success || err != nil {
		return nil, false, err
	}
	externalTrafficPolicyUpdate(oldService, service)
	if errs := validation.ValidateServiceExternalTrafficFieldsCombination(service); len(errs) > 0 {
		return nil, false, errors.NewInvalid(api.Kind("Service"), service.Name, errs)
	}
	out, created, err := rs.services.Update(ctx, service.Name, rest.DefaultUpdatedObjectInfo(service), createValidation, updateValidation, forceAllowCreate, options)
	if err == nil {
		el := nodePortOp.Commit()
		if el != nil {
			utilruntime.HandleError(fmt.Errorf("error(s) committing NodePorts changes: %v", el))
		}
		releaseServiceIP = false
	}
	return out, created, err
}

var _ = rest.Redirector(&REST{})

func (rs *REST) ResourceLocation(ctx context.Context, id string) (*url.URL, http.RoundTripper, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	svcScheme, svcName, portStr, valid := utilnet.SplitSchemeNamePort(id)
	if !valid {
		return nil, nil, errors.NewBadRequest(fmt.Sprintf("invalid service request %q", id))
	}
	if portNum, err := strconv.ParseInt(portStr, 10, 64); err == nil {
		obj, err := rs.services.Get(ctx, svcName, &metav1.GetOptions{})
		if err != nil {
			return nil, nil, err
		}
		svc := obj.(*api.Service)
		found := false
		for _, svcPort := range svc.Spec.Ports {
			if int64(svcPort.Port) == portNum {
				portStr = svcPort.Name
				found = true
				break
			}
		}
		if !found {
			return nil, nil, errors.NewServiceUnavailable(fmt.Sprintf("no service port %d found for service %q", portNum, svcName))
		}
	}
	obj, err := rs.endpoints.Get(ctx, svcName, &metav1.GetOptions{})
	if err != nil {
		return nil, nil, err
	}
	eps := obj.(*api.Endpoints)
	if len(eps.Subsets) == 0 {
		return nil, nil, errors.NewServiceUnavailable(fmt.Sprintf("no endpoints available for service %q", svcName))
	}
	ssSeed := rand.Intn(len(eps.Subsets))
	for ssi := 0; ssi < len(eps.Subsets); ssi++ {
		ss := &eps.Subsets[(ssSeed+ssi)%len(eps.Subsets)]
		if len(ss.Addresses) == 0 {
			continue
		}
		for i := range ss.Ports {
			if ss.Ports[i].Name == portStr {
				addrSeed := rand.Intn(len(ss.Addresses))
				for try := 0; try < len(ss.Addresses); try++ {
					addr := ss.Addresses[(addrSeed+try)%len(ss.Addresses)]
					if err := isValidAddress(ctx, &addr, rs.pods); err != nil {
						utilruntime.HandleError(fmt.Errorf("Address %v isn't valid (%v)", addr, err))
						continue
					}
					ip := addr.IP
					port := int(ss.Ports[i].Port)
					return &url.URL{Scheme: svcScheme, Host: net.JoinHostPort(ip, strconv.Itoa(port))}, rs.proxyTransport, nil
				}
				utilruntime.HandleError(fmt.Errorf("Failed to find a valid address, skipping subset: %v", ss))
			}
		}
	}
	return nil, nil, errors.NewServiceUnavailable(fmt.Sprintf("no endpoints available for service %q", id))
}
func (r *REST) ConvertToTable(ctx context.Context, object runtime.Object, tableOptions runtime.Object) (*metav1beta1.Table, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return r.services.ConvertToTable(ctx, object, tableOptions)
}
func isValidAddress(ctx context.Context, addr *api.EndpointAddress, pods rest.Getter) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if addr.TargetRef == nil {
		return fmt.Errorf("Address has no target ref, skipping: %v", addr)
	}
	if genericapirequest.NamespaceValue(ctx) != addr.TargetRef.Namespace {
		return fmt.Errorf("Address namespace doesn't match context namespace")
	}
	obj, err := pods.Get(ctx, addr.TargetRef.Name, &metav1.GetOptions{})
	if err != nil {
		return err
	}
	pod, ok := obj.(*api.Pod)
	if !ok {
		return fmt.Errorf("failed to cast to pod: %v", obj)
	}
	if pod == nil {
		return fmt.Errorf("pod is missing, skipping (%s/%s)", addr.TargetRef.Namespace, addr.TargetRef.Name)
	}
	if pod.Status.PodIP != addr.IP {
		return fmt.Errorf("pod ip doesn't match endpoint ip, skipping: %s vs %s (%s/%s)", pod.Status.PodIP, addr.IP, addr.TargetRef.Namespace, addr.TargetRef.Name)
	}
	return nil
}
func containsNumber(haystack []int, needle int) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}
	return false
}
func containsNodePort(serviceNodePorts []ServiceNodePort, serviceNodePort ServiceNodePort) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for _, snp := range serviceNodePorts {
		if snp == serviceNodePort {
			return true
		}
	}
	return false
}
func findRequestedNodePort(port int, servicePorts []api.ServicePort) int {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for i := range servicePorts {
		servicePort := servicePorts[i]
		if port == int(servicePort.Port) && servicePort.NodePort != 0 {
			return int(servicePort.NodePort)
		}
	}
	return 0
}
func allocateHealthCheckNodePort(service *api.Service, nodePortOp *portallocator.PortAllocationOperation) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	healthCheckNodePort := service.Spec.HealthCheckNodePort
	if healthCheckNodePort != 0 {
		err := nodePortOp.Allocate(int(healthCheckNodePort))
		if err != nil {
			return fmt.Errorf("failed to allocate requested HealthCheck NodePort %v: %v", healthCheckNodePort, err)
		}
		klog.V(4).Infof("Reserved user requested healthCheckNodePort: %d", healthCheckNodePort)
	} else {
		healthCheckNodePort, err := nodePortOp.AllocateNext()
		if err != nil {
			return fmt.Errorf("failed to allocate a HealthCheck NodePort %v: %v", healthCheckNodePort, err)
		}
		service.Spec.HealthCheckNodePort = int32(healthCheckNodePort)
		klog.V(4).Infof("Reserved allocated healthCheckNodePort: %d", healthCheckNodePort)
	}
	return nil
}
func initClusterIP(service *api.Service, serviceIPs ipallocator.Interface) (bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	switch {
	case service.Spec.ClusterIP == "":
		ip, err := serviceIPs.AllocateNext()
		if err != nil {
			return false, errors.NewInternalError(fmt.Errorf("failed to allocate a serviceIP: %v", err))
		}
		service.Spec.ClusterIP = ip.String()
		return true, nil
	case service.Spec.ClusterIP != api.ClusterIPNone && service.Spec.ClusterIP != "":
		if err := serviceIPs.Allocate(net.ParseIP(service.Spec.ClusterIP)); err != nil {
			el := field.ErrorList{field.Invalid(field.NewPath("spec", "clusterIP"), service.Spec.ClusterIP, err.Error())}
			return false, errors.NewInvalid(api.Kind("Service"), service.Name, el)
		}
		return true, nil
	}
	return false, nil
}
func initNodePorts(service *api.Service, nodePortOp *portallocator.PortAllocationOperation) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	svcPortToNodePort := map[int]int{}
	for i := range service.Spec.Ports {
		servicePort := &service.Spec.Ports[i]
		allocatedNodePort := svcPortToNodePort[int(servicePort.Port)]
		if allocatedNodePort == 0 {
			np := findRequestedNodePort(int(servicePort.Port), service.Spec.Ports)
			if np != 0 {
				err := nodePortOp.Allocate(np)
				if err != nil {
					el := field.ErrorList{field.Invalid(field.NewPath("spec", "ports").Index(i).Child("nodePort"), np, err.Error())}
					return errors.NewInvalid(api.Kind("Service"), service.Name, el)
				}
				servicePort.NodePort = int32(np)
				svcPortToNodePort[int(servicePort.Port)] = np
			} else {
				nodePort, err := nodePortOp.AllocateNext()
				if err != nil {
					return errors.NewInternalError(fmt.Errorf("failed to allocate a nodePort: %v", err))
				}
				servicePort.NodePort = int32(nodePort)
				svcPortToNodePort[int(servicePort.Port)] = nodePort
			}
		} else if int(servicePort.NodePort) != allocatedNodePort {
			if servicePort.NodePort == 0 {
				servicePort.NodePort = int32(allocatedNodePort)
			} else {
				err := nodePortOp.Allocate(int(servicePort.NodePort))
				if err != nil {
					el := field.ErrorList{field.Invalid(field.NewPath("spec", "ports").Index(i).Child("nodePort"), servicePort.NodePort, err.Error())}
					return errors.NewInvalid(api.Kind("Service"), service.Name, el)
				}
			}
		}
	}
	return nil
}
func updateNodePorts(oldService, newService *api.Service, nodePortOp *portallocator.PortAllocationOperation) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	oldNodePortsNumbers := collectServiceNodePorts(oldService)
	newNodePorts := []ServiceNodePort{}
	portAllocated := map[int]bool{}
	for i := range newService.Spec.Ports {
		servicePort := &newService.Spec.Ports[i]
		nodePort := ServiceNodePort{Protocol: servicePort.Protocol, NodePort: servicePort.NodePort}
		if nodePort.NodePort != 0 {
			if !containsNumber(oldNodePortsNumbers, int(nodePort.NodePort)) && !portAllocated[int(nodePort.NodePort)] {
				err := nodePortOp.Allocate(int(nodePort.NodePort))
				if err != nil {
					el := field.ErrorList{field.Invalid(field.NewPath("spec", "ports").Index(i).Child("nodePort"), nodePort.NodePort, err.Error())}
					return errors.NewInvalid(api.Kind("Service"), newService.Name, el)
				}
				portAllocated[int(nodePort.NodePort)] = true
			}
		} else {
			nodePortNumber, err := nodePortOp.AllocateNext()
			if err != nil {
				return errors.NewInternalError(fmt.Errorf("failed to allocate a nodePort: %v", err))
			}
			servicePort.NodePort = int32(nodePortNumber)
			nodePort.NodePort = servicePort.NodePort
		}
		if containsNodePort(newNodePorts, nodePort) {
			return fmt.Errorf("duplicate nodePort: %v", nodePort)
		}
		newNodePorts = append(newNodePorts, nodePort)
	}
	newNodePortsNumbers := collectServiceNodePorts(newService)
	for _, oldNodePortNumber := range oldNodePortsNumbers {
		if containsNumber(newNodePortsNumbers, oldNodePortNumber) {
			continue
		}
		nodePortOp.ReleaseDeferred(int(oldNodePortNumber))
	}
	return nil
}
func releaseNodePorts(service *api.Service, nodePortOp *portallocator.PortAllocationOperation) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	nodePorts := collectServiceNodePorts(service)
	for _, nodePort := range nodePorts {
		nodePortOp.ReleaseDeferred(nodePort)
	}
}
func collectServiceNodePorts(service *api.Service) []int {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	servicePorts := []int{}
	for i := range service.Spec.Ports {
		servicePort := &service.Spec.Ports[i]
		if servicePort.NodePort != 0 {
			servicePorts = append(servicePorts, int(servicePort.NodePort))
		}
	}
	return servicePorts
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
