package master

import (
	"fmt"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	utilnet "k8s.io/apimachinery/pkg/util/net"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	genericapiserver "k8s.io/apiserver/pkg/server"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/rest"
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/features"
	"k8s.io/kubernetes/pkg/master/reconcilers"
	"k8s.io/kubernetes/pkg/registry/core/rangeallocation"
	corerest "k8s.io/kubernetes/pkg/registry/core/rest"
	servicecontroller "k8s.io/kubernetes/pkg/registry/core/service/ipallocator/controller"
	portallocatorcontroller "k8s.io/kubernetes/pkg/registry/core/service/portallocator/controller"
	"k8s.io/kubernetes/pkg/util/async"
	"net"
	"net/http"
	"time"
)

const kubernetesServiceName = "kubernetes"

type Controller struct {
	ServiceClient             corev1client.ServicesGetter
	NamespaceClient           corev1client.NamespacesGetter
	EventClient               corev1client.EventsGetter
	healthClient              rest.Interface
	ServiceClusterIPRegistry  rangeallocation.RangeRegistry
	ServiceClusterIPInterval  time.Duration
	ServiceClusterIPRange     net.IPNet
	ServiceNodePortRegistry   rangeallocation.RangeRegistry
	ServiceNodePortInterval   time.Duration
	ServiceNodePortRange      utilnet.PortRange
	EndpointReconciler        reconcilers.EndpointReconciler
	EndpointInterval          time.Duration
	SystemNamespaces          []string
	SystemNamespacesInterval  time.Duration
	PublicIP                  net.IP
	ServiceIP                 net.IP
	ServicePort               int
	ExtraServicePorts         []corev1.ServicePort
	ExtraEndpointPorts        []corev1.EndpointPort
	PublicServicePort         int
	KubernetesServiceNodePort int
	runner                    *async.Runner
}

func (c *completedConfig) NewBootstrapController(legacyRESTStorage corerest.LegacyRESTStorage, serviceClient corev1client.ServicesGetter, nsClient corev1client.NamespacesGetter, eventClient corev1client.EventsGetter, healthClient rest.Interface) *Controller {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	_, publicServicePort, err := c.GenericConfig.SecureServing.HostPort()
	if err != nil {
		klog.Fatalf("failed to get listener address: %v", err)
	}
	systemNamespaces := []string{metav1.NamespaceSystem, metav1.NamespacePublic}
	if utilfeature.DefaultFeatureGate.Enabled(features.NodeLease) {
		systemNamespaces = append(systemNamespaces, corev1.NamespaceNodeLease)
	}
	return &Controller{ServiceClient: serviceClient, NamespaceClient: nsClient, EventClient: eventClient, healthClient: healthClient, EndpointReconciler: c.ExtraConfig.EndpointReconcilerConfig.Reconciler, EndpointInterval: c.ExtraConfig.EndpointReconcilerConfig.Interval, SystemNamespaces: systemNamespaces, SystemNamespacesInterval: 1 * time.Minute, ServiceClusterIPRegistry: legacyRESTStorage.ServiceClusterIPAllocator, ServiceClusterIPRange: c.ExtraConfig.ServiceIPRange, ServiceClusterIPInterval: 3 * time.Minute, ServiceNodePortRegistry: legacyRESTStorage.ServiceNodePortAllocator, ServiceNodePortRange: c.ExtraConfig.ServiceNodePortRange, ServiceNodePortInterval: 3 * time.Minute, PublicIP: c.GenericConfig.PublicAddress, ServiceIP: c.ExtraConfig.APIServerServiceIP, ServicePort: c.ExtraConfig.APIServerServicePort, ExtraServicePorts: c.ExtraConfig.ExtraServicePorts, ExtraEndpointPorts: c.ExtraConfig.ExtraEndpointPorts, PublicServicePort: publicServicePort, KubernetesServiceNodePort: c.ExtraConfig.KubernetesServiceNodePort}
}
func (c *Controller) PostStartHook(hookContext genericapiserver.PostStartHookContext) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	c.Start()
	return nil
}
func (c *Controller) PreShutdownHook() error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	c.Stop()
	return nil
}
func (c *Controller) Start() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if c.runner != nil {
		return
	}
	endpointPorts := createEndpointPortSpec(c.PublicServicePort, "https", c.ExtraEndpointPorts)
	if err := c.EndpointReconciler.RemoveEndpoints(kubernetesServiceName, c.PublicIP, endpointPorts); err != nil {
		klog.Errorf("Unable to remove old endpoints from kubernetes service: %v", err)
	}
	repairClusterIPs := servicecontroller.NewRepair(c.ServiceClusterIPInterval, c.ServiceClient, c.EventClient, &c.ServiceClusterIPRange, c.ServiceClusterIPRegistry)
	repairNodePorts := portallocatorcontroller.NewRepair(c.ServiceNodePortInterval, c.ServiceClient, c.EventClient, c.ServiceNodePortRange, c.ServiceNodePortRegistry)
	if err := repairClusterIPs.RunOnce(); err != nil {
		klog.Fatalf("Unable to perform initial IP allocation check: %v", err)
	}
	if err := repairNodePorts.RunOnce(); err != nil {
		klog.Fatalf("Unable to perform initial service nodePort check: %v", err)
	}
	c.runner = async.NewRunner(c.RunKubernetesNamespaces, c.RunKubernetesService, repairClusterIPs.RunUntil, repairNodePorts.RunUntil)
	c.runner.Start()
}
func (c *Controller) Stop() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if c.runner != nil {
		c.runner.Stop()
	}
	endpointPorts := createEndpointPortSpec(c.PublicServicePort, "https", c.ExtraEndpointPorts)
	finishedReconciling := make(chan struct{})
	go func() {
		defer close(finishedReconciling)
		klog.Infof("Shutting down kubernetes service endpoint reconciler")
		c.EndpointReconciler.StopReconciling()
		if err := c.EndpointReconciler.RemoveEndpoints(kubernetesServiceName, c.PublicIP, endpointPorts); err != nil {
			klog.Error(err)
		}
	}()
	select {
	case <-finishedReconciling:
	case <-time.After(2 * c.EndpointInterval):
		klog.Warning("RemoveEndpoints() timed out")
	}
}
func (c *Controller) RunKubernetesNamespaces(ch chan struct{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	wait.Until(func() {
		for _, ns := range c.SystemNamespaces {
			if err := createNamespaceIfNeeded(c.NamespaceClient, ns); err != nil {
				runtime.HandleError(fmt.Errorf("unable to create required kubernetes system namespace %s: %v", ns, err))
			}
		}
	}, c.SystemNamespacesInterval, ch)
}
func (c *Controller) RunKubernetesService(ch chan struct{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	wait.PollImmediateUntil(100*time.Millisecond, func() (bool, error) {
		var code int
		c.healthClient.Get().AbsPath("/healthz").Do().StatusCode(&code)
		return code == http.StatusOK, nil
	}, ch)
	wait.NonSlidingUntil(func() {
		if err := c.UpdateKubernetesService(false); err != nil {
			runtime.HandleError(fmt.Errorf("unable to sync kubernetes service: %v", err))
		}
	}, c.EndpointInterval, ch)
}
func (c *Controller) UpdateKubernetesService(reconcile bool) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if err := createNamespaceIfNeeded(c.NamespaceClient, metav1.NamespaceDefault); err != nil {
		return err
	}
	servicePorts, serviceType := createPortAndServiceSpec(c.ServicePort, c.PublicServicePort, c.KubernetesServiceNodePort, "https", c.ExtraServicePorts)
	if err := c.CreateOrUpdateMasterServiceIfNeeded(kubernetesServiceName, c.ServiceIP, servicePorts, serviceType, reconcile); err != nil {
		return err
	}
	endpointPorts := createEndpointPortSpec(c.PublicServicePort, "https", c.ExtraEndpointPorts)
	if err := c.EndpointReconciler.ReconcileEndpoints(kubernetesServiceName, c.PublicIP, endpointPorts, reconcile); err != nil {
		return err
	}
	return nil
}
func createPortAndServiceSpec(servicePort int, targetServicePort int, nodePort int, servicePortName string, extraServicePorts []corev1.ServicePort) ([]corev1.ServicePort, corev1.ServiceType) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	servicePorts := []corev1.ServicePort{{Protocol: corev1.ProtocolTCP, Port: int32(servicePort), Name: servicePortName, TargetPort: intstr.FromInt(targetServicePort)}}
	serviceType := corev1.ServiceTypeClusterIP
	if nodePort > 0 {
		servicePorts[0].NodePort = int32(nodePort)
		serviceType = corev1.ServiceTypeNodePort
	}
	if extraServicePorts != nil {
		servicePorts = append(servicePorts, extraServicePorts...)
	}
	return servicePorts, serviceType
}
func createEndpointPortSpec(endpointPort int, endpointPortName string, extraEndpointPorts []corev1.EndpointPort) []corev1.EndpointPort {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	endpointPorts := []corev1.EndpointPort{{Protocol: corev1.ProtocolTCP, Port: int32(endpointPort), Name: endpointPortName}}
	if extraEndpointPorts != nil {
		endpointPorts = append(endpointPorts, extraEndpointPorts...)
	}
	return endpointPorts
}
func (c *Controller) CreateOrUpdateMasterServiceIfNeeded(serviceName string, serviceIP net.IP, servicePorts []corev1.ServicePort, serviceType corev1.ServiceType, reconcile bool) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if s, err := c.ServiceClient.Services(metav1.NamespaceDefault).Get(serviceName, metav1.GetOptions{}); err == nil {
		if reconcile {
			if svc, updated := reconcilers.GetMasterServiceUpdateIfNeeded(s, servicePorts, serviceType); updated {
				klog.Warningf("Resetting master service %q to %#v", serviceName, svc)
				_, err := c.ServiceClient.Services(metav1.NamespaceDefault).Update(svc)
				return err
			}
		}
		return nil
	}
	svc := &corev1.Service{ObjectMeta: metav1.ObjectMeta{Name: serviceName, Namespace: metav1.NamespaceDefault, Labels: map[string]string{"provider": "kubernetes", "component": "apiserver"}}, Spec: corev1.ServiceSpec{Ports: servicePorts, Selector: nil, ClusterIP: serviceIP.String(), SessionAffinity: corev1.ServiceAffinityNone, Type: serviceType}}
	_, err := c.ServiceClient.Services(metav1.NamespaceDefault).Create(svc)
	if errors.IsAlreadyExists(err) {
		return c.CreateOrUpdateMasterServiceIfNeeded(serviceName, serviceIP, servicePorts, serviceType, reconcile)
	}
	return err
}
