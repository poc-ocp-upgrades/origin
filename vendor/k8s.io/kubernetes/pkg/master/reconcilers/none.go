package reconcilers

import (
 "net"
 corev1 "k8s.io/api/core/v1"
)

type noneEndpointReconciler struct{}

func NewNoneEndpointReconciler() EndpointReconciler {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &noneEndpointReconciler{}
}
func (r *noneEndpointReconciler) ReconcileEndpoints(serviceName string, ip net.IP, endpointPorts []corev1.EndpointPort, reconcilePorts bool) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return nil
}
func (r *noneEndpointReconciler) RemoveEndpoints(serviceName string, ip net.IP, endpointPorts []corev1.EndpointPort) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return nil
}
func (r *noneEndpointReconciler) StopReconciling() {
 _logClusterCodePath()
 defer _logClusterCodePath()
}
