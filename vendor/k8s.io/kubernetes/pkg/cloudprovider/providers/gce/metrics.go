package gce

import (
 "time"
 "github.com/prometheus/client_golang/prometheus"
)

const (
 computeV1Version    = "v1"
 computeAlphaVersion = "alpha"
 computeBetaVersion  = "beta"
)

type apiCallMetrics struct {
 latency *prometheus.HistogramVec
 errors  *prometheus.CounterVec
}

var (
 metricLabels = []string{"request", "region", "zone", "version"}
 apiMetrics   = registerAPIMetrics(metricLabels...)
)

type metricContext struct {
 start      time.Time
 attributes []string
}

const unusedMetricLabel = "<n/a>"

func (mc *metricContext) Observe(err error) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 apiMetrics.latency.WithLabelValues(mc.attributes...).Observe(time.Since(mc.start).Seconds())
 if err != nil {
  apiMetrics.errors.WithLabelValues(mc.attributes...).Inc()
 }
 return err
}
func newGenericMetricContext(prefix, request, region, zone, version string) *metricContext {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if len(zone) == 0 {
  zone = unusedMetricLabel
 }
 if len(region) == 0 {
  region = unusedMetricLabel
 }
 return &metricContext{start: time.Now(), attributes: []string{prefix + "_" + request, region, zone, version}}
}
func registerAPIMetrics(attributes ...string) *apiCallMetrics {
 _logClusterCodePath()
 defer _logClusterCodePath()
 metrics := &apiCallMetrics{latency: prometheus.NewHistogramVec(prometheus.HistogramOpts{Name: "cloudprovider_gce_api_request_duration_seconds", Help: "Latency of a GCE API call"}, attributes), errors: prometheus.NewCounterVec(prometheus.CounterOpts{Name: "cloudprovider_gce_api_request_errors", Help: "Number of errors for an API call"}, attributes)}
 prometheus.MustRegister(metrics.latency)
 prometheus.MustRegister(metrics.errors)
 return metrics
}
