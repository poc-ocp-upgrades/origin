package metrics

import (
	goformat "fmt"
	"github.com/prometheus/client_golang/prometheus"
	"k8s.io/api/core/v1"
	"k8s.io/klog"
	goos "os"
	godefaultruntime "runtime"
	"sync"
	gotime "time"
)

const (
	pvControllerSubsystem = "pv_collector"
	boundPVKey            = "bound_pv_count"
	unboundPVKey          = "unbound_pv_count"
	boundPVCKey           = "bound_pvc_count"
	unboundPVCKey         = "unbound_pvc_count"
	namespaceLabel        = "namespace"
	storageClassLabel     = "storage_class"
)

var registerMetrics sync.Once

type PVLister interface{ List() []interface{} }
type PVCLister interface{ List() []interface{} }

func Register(pvLister PVLister, pvcLister PVCLister) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	registerMetrics.Do(func() {
		prometheus.MustRegister(newPVAndPVCCountCollector(pvLister, pvcLister))
		prometheus.MustRegister(volumeOperationMetric)
		prometheus.MustRegister(volumeOperationErrorsMetric)
	})
}
func newPVAndPVCCountCollector(pvLister PVLister, pvcLister PVCLister) *pvAndPVCCountCollector {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &pvAndPVCCountCollector{pvLister, pvcLister}
}

type pvAndPVCCountCollector struct {
	pvLister  PVLister
	pvcLister PVCLister
}

var (
	boundPVCountDesc            = prometheus.NewDesc(prometheus.BuildFQName("", pvControllerSubsystem, boundPVKey), "Gauge measuring number of persistent volume currently bound", []string{storageClassLabel}, nil)
	unboundPVCountDesc          = prometheus.NewDesc(prometheus.BuildFQName("", pvControllerSubsystem, unboundPVKey), "Gauge measuring number of persistent volume currently unbound", []string{storageClassLabel}, nil)
	boundPVCCountDesc           = prometheus.NewDesc(prometheus.BuildFQName("", pvControllerSubsystem, boundPVCKey), "Gauge measuring number of persistent volume claim currently bound", []string{namespaceLabel}, nil)
	unboundPVCCountDesc         = prometheus.NewDesc(prometheus.BuildFQName("", pvControllerSubsystem, unboundPVCKey), "Gauge measuring number of persistent volume claim currently unbound", []string{namespaceLabel}, nil)
	volumeOperationMetric       = prometheus.NewHistogramVec(prometheus.HistogramOpts{Name: "volume_operation_total_seconds", Help: "Total volume operation time"}, []string{"plugin_name", "operation_name"})
	volumeOperationErrorsMetric = prometheus.NewCounterVec(prometheus.CounterOpts{Name: "volume_operation_total_errors", Help: "Total volume operation erros"}, []string{"plugin_name", "operation_name"})
)

func (collector *pvAndPVCCountCollector) Describe(ch chan<- *prometheus.Desc) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ch <- boundPVCountDesc
	ch <- unboundPVCountDesc
	ch <- boundPVCCountDesc
	ch <- unboundPVCCountDesc
}
func (collector *pvAndPVCCountCollector) Collect(ch chan<- prometheus.Metric) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	collector.pvCollect(ch)
	collector.pvcCollect(ch)
}
func (collector *pvAndPVCCountCollector) pvCollect(ch chan<- prometheus.Metric) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	boundNumberByStorageClass := make(map[string]int)
	unboundNumberByStorageClass := make(map[string]int)
	for _, pvObj := range collector.pvLister.List() {
		pv, ok := pvObj.(*v1.PersistentVolume)
		if !ok {
			continue
		}
		if pv.Status.Phase == v1.VolumeBound {
			boundNumberByStorageClass[pv.Spec.StorageClassName]++
		} else {
			unboundNumberByStorageClass[pv.Spec.StorageClassName]++
		}
	}
	for storageClassName, number := range boundNumberByStorageClass {
		metric, err := prometheus.NewConstMetric(boundPVCountDesc, prometheus.GaugeValue, float64(number), storageClassName)
		if err != nil {
			klog.Warningf("Create bound pv number metric failed: %v", err)
			continue
		}
		ch <- metric
	}
	for storageClassName, number := range unboundNumberByStorageClass {
		metric, err := prometheus.NewConstMetric(unboundPVCountDesc, prometheus.GaugeValue, float64(number), storageClassName)
		if err != nil {
			klog.Warningf("Create unbound pv number metric failed: %v", err)
			continue
		}
		ch <- metric
	}
}
func (collector *pvAndPVCCountCollector) pvcCollect(ch chan<- prometheus.Metric) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	boundNumberByNamespace := make(map[string]int)
	unboundNumberByNamespace := make(map[string]int)
	for _, pvcObj := range collector.pvcLister.List() {
		pvc, ok := pvcObj.(*v1.PersistentVolumeClaim)
		if !ok {
			continue
		}
		if pvc.Status.Phase == v1.ClaimBound {
			boundNumberByNamespace[pvc.Namespace]++
		} else {
			unboundNumberByNamespace[pvc.Namespace]++
		}
	}
	for namespace, number := range boundNumberByNamespace {
		metric, err := prometheus.NewConstMetric(boundPVCCountDesc, prometheus.GaugeValue, float64(number), namespace)
		if err != nil {
			klog.Warningf("Create bound pvc number metric failed: %v", err)
			continue
		}
		ch <- metric
	}
	for namespace, number := range unboundNumberByNamespace {
		metric, err := prometheus.NewConstMetric(unboundPVCCountDesc, prometheus.GaugeValue, float64(number), namespace)
		if err != nil {
			klog.Warningf("Create unbound pvc number metric failed: %v", err)
			continue
		}
		ch <- metric
	}
}
func RecordVolumeOperationMetric(pluginName, opName string, timeTaken float64, err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if pluginName == "" {
		pluginName = "N/A"
	}
	if err != nil {
		volumeOperationErrorsMetric.WithLabelValues(pluginName, opName).Inc()
		return
	}
	volumeOperationMetric.WithLabelValues(pluginName, opName).Observe(timeTaken)
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
