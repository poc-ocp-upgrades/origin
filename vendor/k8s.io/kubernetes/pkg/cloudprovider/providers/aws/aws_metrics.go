package aws

import "github.com/prometheus/client_golang/prometheus"

var (
 awsAPIMetric          = prometheus.NewHistogramVec(prometheus.HistogramOpts{Name: "cloudprovider_aws_api_request_duration_seconds", Help: "Latency of AWS API calls"}, []string{"request"})
 awsAPIErrorMetric     = prometheus.NewCounterVec(prometheus.CounterOpts{Name: "cloudprovider_aws_api_request_errors", Help: "AWS API errors"}, []string{"request"})
 awsAPIThrottlesMetric = prometheus.NewCounterVec(prometheus.CounterOpts{Name: "cloudprovider_aws_api_throttled_requests_total", Help: "AWS API throttled requests"}, []string{"operation_name"})
)

func recordAWSMetric(actionName string, timeTaken float64, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err != nil {
  awsAPIErrorMetric.With(prometheus.Labels{"request": actionName}).Inc()
 } else {
  awsAPIMetric.With(prometheus.Labels{"request": actionName}).Observe(timeTaken)
 }
}
func recordAWSThrottlesMetric(operation string) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 awsAPIThrottlesMetric.With(prometheus.Labels{"operation_name": operation}).Inc()
}
func registerMetrics() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 prometheus.MustRegister(awsAPIMetric)
 prometheus.MustRegister(awsAPIErrorMetric)
 prometheus.MustRegister(awsAPIThrottlesMetric)
}
