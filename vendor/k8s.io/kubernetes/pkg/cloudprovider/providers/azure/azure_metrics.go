package azure

import (
 "time"
 "github.com/prometheus/client_golang/prometheus"
)

type apiCallMetrics struct {
 latency *prometheus.HistogramVec
 errors  *prometheus.CounterVec
}

var (
 metricLabels = []string{"request", "resource_group", "subscription_id"}
 apiMetrics   = registerAPIMetrics(metricLabels...)
)

type metricContext struct {
 start      time.Time
 attributes []string
}

func newMetricContext(prefix, request, resourceGroup, subscriptionID string) *metricContext {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &metricContext{start: time.Now(), attributes: []string{prefix + "_" + request, resourceGroup, subscriptionID}}
}
func (mc *metricContext) Observe(err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 apiMetrics.latency.WithLabelValues(mc.attributes...).Observe(time.Since(mc.start).Seconds())
 if err != nil {
  apiMetrics.errors.WithLabelValues(mc.attributes...).Inc()
 }
}
func registerAPIMetrics(attributes ...string) *apiCallMetrics {
 _logClusterCodePath()
 defer _logClusterCodePath()
 metrics := &apiCallMetrics{latency: prometheus.NewHistogramVec(prometheus.HistogramOpts{Name: "cloudprovider_azure_api_request_duration_seconds", Help: "Latency of an Azure API call"}, attributes), errors: prometheus.NewCounterVec(prometheus.CounterOpts{Name: "cloudprovider_azure_api_request_errors", Help: "Number of errors for an Azure API call"}, attributes)}
 prometheus.MustRegister(metrics.latency)
 prometheus.MustRegister(metrics.errors)
 return metrics
}
