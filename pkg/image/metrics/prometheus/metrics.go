package prometheus

import (
	goformat "fmt"
	"github.com/prometheus/client_golang/prometheus"
	"k8s.io/klog"
	goos "os"
	godefaultruntime "runtime"
	"sync"
	gotime "time"
)

const (
	separator          = "_"
	metricController   = "openshift_imagestreamcontroller"
	metricCount        = "count"
	metricSuccessCount = metricController + separator + "success" + separator + metricCount
	metricErrorCount   = metricController + separator + "error" + separator + metricCount
	labelScheduled     = "scheduled"
	labelRegistry      = "registry"
	labelReason        = "reason"
)

type ImportErrorInfo struct {
	Registry string
	Reason   string
}
type ImportSuccessCounts map[string]uint64
type ImportErrorCounts map[ImportErrorInfo]uint64
type QueuedImageStreamFetcher func() (ImportSuccessCounts, ImportErrorCounts, error)

var (
	successCountDesc    = prometheus.NewDesc(metricSuccessCount, "Counts successful image stream imports - both scheduled and not scheduled - per image registry", []string{labelScheduled, labelRegistry}, nil)
	errorCountDesc      = prometheus.NewDesc(metricErrorCount, "Counts number of failed image stream imports - both scheduled and not scheduled"+" - per image registry and failure reason", []string{labelScheduled, labelRegistry, labelReason}, nil)
	isc                 = importStatusCollector{}
	registerLock        = sync.Mutex{}
	collectorRegistered = false
)

type importStatusCollector struct {
	cbCollectISCounts        QueuedImageStreamFetcher
	cbCollectScheduledCounts QueuedImageStreamFetcher
}

func InitializeImportCollector(scheduled bool, cbCollectISCounts QueuedImageStreamFetcher) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	registerLock.Lock()
	defer registerLock.Unlock()
	if scheduled {
		isc.cbCollectScheduledCounts = cbCollectISCounts
	} else {
		isc.cbCollectISCounts = cbCollectISCounts
	}
	if collectorRegistered {
		return
	}
	if isc.cbCollectISCounts != nil && isc.cbCollectScheduledCounts != nil {
		prometheus.MustRegister(&isc)
		collectorRegistered = true
		klog.V(4).Info("Image import controller metrics registered with prometherus")
	}
}
func (isc *importStatusCollector) Describe(ch chan<- *prometheus.Desc) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ch <- successCountDesc
	ch <- errorCountDesc
}
func (isc *importStatusCollector) Collect(ch chan<- prometheus.Metric) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	successCounts, errorCounts, err := isc.cbCollectISCounts()
	if err != nil {
		klog.Errorf("Failed to collect image import metrics: %v", err)
		ch <- prometheus.NewInvalidMetric(successCountDesc, err)
	} else {
		pushSuccessCounts("false", successCounts, ch)
		pushErrorCounts("false", errorCounts, ch)
	}
	successCounts, errorCounts, err = isc.cbCollectScheduledCounts()
	if err != nil {
		klog.Errorf("Failed to collect scheduled image import metrics: %v", err)
		ch <- prometheus.NewInvalidMetric(errorCountDesc, err)
		return
	}
	pushSuccessCounts("true", successCounts, ch)
	pushErrorCounts("true", errorCounts, ch)
}
func pushSuccessCounts(scheduled string, counts ImportSuccessCounts, ch chan<- prometheus.Metric) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for registry, count := range counts {
		ch <- prometheus.MustNewConstMetric(successCountDesc, prometheus.CounterValue, float64(count), scheduled, registry)
	}
}
func pushErrorCounts(scheduled string, counts ImportErrorCounts, ch chan<- prometheus.Metric) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for info, count := range counts {
		ch <- prometheus.MustNewConstMetric(errorCountDesc, prometheus.CounterValue, float64(count), scheduled, info.Registry, info.Reason)
	}
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
