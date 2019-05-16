package persistentvolume

import (
	"github.com/prometheus/client_golang/prometheus"
)

const VolumeSchedulerSubsystem = "scheduler_volume"

var (
	VolumeBindingRequestSchedulerBinderCache = prometheus.NewCounterVec(prometheus.CounterOpts{Subsystem: VolumeSchedulerSubsystem, Name: "binder_cache_requests_total", Help: "Total number for request volume binding cache"}, []string{"operation"})
	VolumeSchedulingStageLatency             = prometheus.NewHistogramVec(prometheus.HistogramOpts{Subsystem: VolumeSchedulerSubsystem, Name: "scheduling_duration_seconds", Help: "Volume scheduling stage latency", Buckets: prometheus.ExponentialBuckets(1000, 2, 15)}, []string{"operation"})
	VolumeSchedulingStageFailed              = prometheus.NewCounterVec(prometheus.CounterOpts{Subsystem: VolumeSchedulerSubsystem, Name: "scheduling_stage_error_total", Help: "Volume scheduling stage error count"}, []string{"operation"})
)

func RegisterVolumeSchedulingMetrics() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	prometheus.MustRegister(VolumeBindingRequestSchedulerBinderCache)
	prometheus.MustRegister(VolumeSchedulingStageLatency)
	prometheus.MustRegister(VolumeSchedulingStageFailed)
}
