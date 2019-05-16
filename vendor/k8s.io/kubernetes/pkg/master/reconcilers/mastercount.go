package reconcilers

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/util/retry"
	"k8s.io/klog"
	endpointsv1 "k8s.io/kubernetes/pkg/api/v1/endpoints"
	"net"
	"sync"
)

type masterCountEndpointReconciler struct {
	masterCount           int
	endpointClient        corev1client.EndpointsGetter
	stopReconcilingCalled bool
	reconcilingLock       sync.Mutex
}

func NewMasterCountEndpointReconciler(masterCount int, endpointClient corev1client.EndpointsGetter) EndpointReconciler {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &masterCountEndpointReconciler{masterCount: masterCount, endpointClient: endpointClient}
}
func (r *masterCountEndpointReconciler) ReconcileEndpoints(serviceName string, ip net.IP, endpointPorts []corev1.EndpointPort, reconcilePorts bool) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	r.reconcilingLock.Lock()
	defer r.reconcilingLock.Unlock()
	if r.stopReconcilingCalled {
		return nil
	}
	e, err := r.endpointClient.Endpoints(metav1.NamespaceDefault).Get(serviceName, metav1.GetOptions{})
	if err != nil {
		e = &corev1.Endpoints{ObjectMeta: metav1.ObjectMeta{Name: serviceName, Namespace: metav1.NamespaceDefault}}
	}
	if errors.IsNotFound(err) {
		e.Subsets = []corev1.EndpointSubset{{Addresses: []corev1.EndpointAddress{{IP: ip.String()}}, Ports: endpointPorts}}
		_, err = r.endpointClient.Endpoints(metav1.NamespaceDefault).Create(e)
		return err
	}
	formatCorrect, ipCorrect, portsCorrect := checkEndpointSubsetFormat(e, ip.String(), endpointPorts, r.masterCount, reconcilePorts)
	if !formatCorrect {
		e.Subsets = []corev1.EndpointSubset{{Addresses: []corev1.EndpointAddress{{IP: ip.String()}}, Ports: endpointPorts}}
		klog.Warningf("Resetting endpoints for master service %q to %#v", serviceName, e)
		_, err = r.endpointClient.Endpoints(metav1.NamespaceDefault).Update(e)
		return err
	}
	if ipCorrect && portsCorrect {
		return nil
	}
	if !ipCorrect {
		e.Subsets[0].Addresses = append(e.Subsets[0].Addresses, corev1.EndpointAddress{IP: ip.String()})
		e.Subsets = endpointsv1.RepackSubsets(e.Subsets)
		if addrs := &e.Subsets[0].Addresses; len(*addrs) > r.masterCount {
			for i, addr := range *addrs {
				if addr.IP == ip.String() {
					for len(*addrs) > r.masterCount {
						remove := (i + 1) % len(*addrs)
						*addrs = append((*addrs)[:remove], (*addrs)[remove+1:]...)
					}
					break
				}
			}
		}
	}
	if !portsCorrect {
		e.Subsets[0].Ports = endpointPorts
	}
	klog.Warningf("Resetting endpoints for master service %q to %v", serviceName, e)
	_, err = r.endpointClient.Endpoints(metav1.NamespaceDefault).Update(e)
	return err
}
func (r *masterCountEndpointReconciler) RemoveEndpoints(serviceName string, ip net.IP, endpointPorts []corev1.EndpointPort) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	r.reconcilingLock.Lock()
	defer r.reconcilingLock.Unlock()
	e, err := r.endpointClient.Endpoints(metav1.NamespaceDefault).Get(serviceName, metav1.GetOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			return nil
		}
		return err
	}
	new := []corev1.EndpointAddress{}
	for _, addr := range e.Subsets[0].Addresses {
		if addr.IP != ip.String() {
			new = append(new, addr)
		}
	}
	e.Subsets[0].Addresses = new
	e.Subsets = endpointsv1.RepackSubsets(e.Subsets)
	err = retry.RetryOnConflict(retry.DefaultBackoff, func() error {
		_, err := r.endpointClient.Endpoints(metav1.NamespaceDefault).Update(e)
		return err
	})
	return err
}
func (r *masterCountEndpointReconciler) StopReconciling() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	r.reconcilingLock.Lock()
	defer r.reconcilingLock.Unlock()
	r.stopReconcilingCalled = true
}
func checkEndpointSubsetFormat(e *corev1.Endpoints, ip string, ports []corev1.EndpointPort, count int, reconcilePorts bool) (formatCorrect bool, ipCorrect bool, portsCorrect bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(e.Subsets) != 1 {
		return false, false, false
	}
	sub := &e.Subsets[0]
	portsCorrect = true
	if reconcilePorts {
		if len(sub.Ports) != len(ports) {
			portsCorrect = false
		}
		for i, port := range ports {
			if len(sub.Ports) <= i || port != sub.Ports[i] {
				portsCorrect = false
				break
			}
		}
	}
	for _, addr := range sub.Addresses {
		if addr.IP == ip {
			ipCorrect = len(sub.Addresses) <= count
			break
		}
	}
	return true, ipCorrect, portsCorrect
}
func GetMasterServiceUpdateIfNeeded(svc *corev1.Service, servicePorts []corev1.ServicePort, serviceType corev1.ServiceType) (s *corev1.Service, updated bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	formatCorrect := checkServiceFormat(svc, servicePorts, serviceType)
	if formatCorrect {
		return svc, false
	}
	svc.Spec.Ports = servicePorts
	svc.Spec.Type = serviceType
	return svc, true
}
func checkServiceFormat(s *corev1.Service, ports []corev1.ServicePort, serviceType corev1.ServiceType) (formatCorrect bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if s.Spec.Type != serviceType {
		return false
	}
	if len(ports) != len(s.Spec.Ports) {
		return false
	}
	for i, port := range ports {
		if port != s.Spec.Ports[i] {
			return false
		}
	}
	return true
}
