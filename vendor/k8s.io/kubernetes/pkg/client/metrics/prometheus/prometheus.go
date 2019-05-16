package prometheus

import (
	goformat "fmt"
	"github.com/prometheus/client_golang/prometheus"
	"k8s.io/client-go/tools/metrics"
	"net/url"
	goos "os"
	godefaultruntime "runtime"
	"time"
	gotime "time"
)

var (
	requestLatency = prometheus.NewHistogramVec(prometheus.HistogramOpts{Name: "rest_client_request_latency_seconds", Help: "Request latency in seconds. Broken down by verb and URL.", Buckets: prometheus.ExponentialBuckets(0.001, 2, 10)}, []string{"verb", "url"})
	requestResult  = prometheus.NewCounterVec(prometheus.CounterOpts{Name: "rest_client_requests_total", Help: "Number of HTTP requests, partitioned by status code, method, and host."}, []string{"code", "method", "host"})
)

func init() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	prometheus.MustRegister(requestLatency)
	prometheus.MustRegister(requestResult)
	metrics.Register(&latencyAdapter{requestLatency}, &resultAdapter{requestResult})
}

type latencyAdapter struct{ m *prometheus.HistogramVec }

func (l *latencyAdapter) Observe(verb string, u url.URL, latency time.Duration) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	l.m.WithLabelValues(verb, u.String()).Observe(latency.Seconds())
}

type resultAdapter struct{ m *prometheus.CounterVec }

func (r *resultAdapter) Increment(code, method, host string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	r.m.WithLabelValues(code, method, host).Inc()
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
