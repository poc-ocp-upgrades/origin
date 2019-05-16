package reconcilers

import (
	corev1 "k8s.io/api/core/v1"
	"net"
)

type noneEndpointReconciler struct{}

func NewNoneEndpointReconciler() EndpointReconciler {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &noneEndpointReconciler{}
}
func (r *noneEndpointReconciler) ReconcileEndpoints(serviceName string, ip net.IP, endpointPorts []corev1.EndpointPort, reconcilePorts bool) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return nil
}
func (r *noneEndpointReconciler) RemoveEndpoints(serviceName string, ip net.IP, endpointPorts []corev1.EndpointPort) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return nil
}
func (r *noneEndpointReconciler) StopReconciling() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
}
