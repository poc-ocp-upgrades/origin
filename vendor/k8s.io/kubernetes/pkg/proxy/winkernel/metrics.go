package winkernel

import (
	goformat "fmt"
	"github.com/prometheus/client_golang/prometheus"
	goos "os"
	godefaultruntime "runtime"
	"sync"
	"time"
	gotime "time"
)

const kubeProxySubsystem = "kubeproxy"

var (
	SyncProxyRulesLatency = prometheus.NewHistogram(prometheus.HistogramOpts{Subsystem: kubeProxySubsystem, Name: "sync_proxy_rules_latency_microseconds", Help: "SyncProxyRules latency", Buckets: prometheus.ExponentialBuckets(1000, 2, 15)})
)
var registerMetricsOnce sync.Once

func RegisterMetrics() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	registerMetricsOnce.Do(func() {
		prometheus.MustRegister(SyncProxyRulesLatency)
	})
}
func sinceInMicroseconds(start time.Time) float64 {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return float64(time.Since(start).Nanoseconds() / time.Microsecond.Nanoseconds())
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
