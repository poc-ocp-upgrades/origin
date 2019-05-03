package openstack

import "github.com/prometheus/client_golang/prometheus"

const (
 openstackSubsystem         = "openstack"
 openstackOperationKey      = "cloudprovider_openstack_api_request_duration_seconds"
 openstackOperationErrorKey = "cloudprovider_openstack_api_request_errors"
)

var (
 openstackOperationsLatency = prometheus.NewHistogramVec(prometheus.HistogramOpts{Subsystem: openstackSubsystem, Name: openstackOperationKey, Help: "Latency of openstack api call"}, []string{"request"})
 openstackAPIRequestErrors  = prometheus.NewCounterVec(prometheus.CounterOpts{Subsystem: openstackSubsystem, Name: openstackOperationErrorKey, Help: "Cumulative number of openstack Api call errors"}, []string{"request"})
)

func registerMetrics() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 prometheus.MustRegister(openstackOperationsLatency)
 prometheus.MustRegister(openstackAPIRequestErrors)
}
