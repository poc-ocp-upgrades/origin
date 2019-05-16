package metrics

import (
	goformat "fmt"
	"github.com/prometheus/client_golang/prometheus"
	"k8s.io/apimachinery/pkg/labels"
	corelisters "k8s.io/client-go/listers/core/v1"
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/controller/volume/attachdetach/cache"
	"k8s.io/kubernetes/pkg/controller/volume/attachdetach/util"
	"k8s.io/kubernetes/pkg/volume"
	volumeutil "k8s.io/kubernetes/pkg/volume/util"
	goos "os"
	godefaultruntime "runtime"
	"sync"
	gotime "time"
)

const pluginNameNotAvailable = "N/A"

var (
	inUseVolumeMetricDesc     = prometheus.NewDesc(prometheus.BuildFQName("", "storage_count", "attachable_volumes_in_use"), "Measure number of volumes in use", []string{"node", "volume_plugin"}, nil)
	totalVolumesMetricDesc    = prometheus.NewDesc(prometheus.BuildFQName("", "attachdetach_controller", "total_volumes"), "Number of volumes in A/D Controller", []string{"plugin_name", "state"}, nil)
	forcedDetachMetricCounter = prometheus.NewCounter(prometheus.CounterOpts{Name: "attachdetach_controller_forced_detaches", Help: "Number of times the A/D Controller performed a forced detach"})
)
var registerMetrics sync.Once

func Register(pvcLister corelisters.PersistentVolumeClaimLister, pvLister corelisters.PersistentVolumeLister, podLister corelisters.PodLister, asw cache.ActualStateOfWorld, dsw cache.DesiredStateOfWorld, pluginMgr *volume.VolumePluginMgr) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	registerMetrics.Do(func() {
		prometheus.MustRegister(newAttachDetachStateCollector(pvcLister, podLister, pvLister, asw, dsw, pluginMgr))
		prometheus.MustRegister(forcedDetachMetricCounter)
	})
}

type attachDetachStateCollector struct {
	pvcLister       corelisters.PersistentVolumeClaimLister
	podLister       corelisters.PodLister
	pvLister        corelisters.PersistentVolumeLister
	asw             cache.ActualStateOfWorld
	dsw             cache.DesiredStateOfWorld
	volumePluginMgr *volume.VolumePluginMgr
}
type volumeCount map[string]map[string]int64

func (v volumeCount) add(typeKey, counterKey string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	count, ok := v[typeKey]
	if !ok {
		count = map[string]int64{}
	}
	count[counterKey]++
	v[typeKey] = count
}
func newAttachDetachStateCollector(pvcLister corelisters.PersistentVolumeClaimLister, podLister corelisters.PodLister, pvLister corelisters.PersistentVolumeLister, asw cache.ActualStateOfWorld, dsw cache.DesiredStateOfWorld, pluginMgr *volume.VolumePluginMgr) *attachDetachStateCollector {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &attachDetachStateCollector{pvcLister, podLister, pvLister, asw, dsw, pluginMgr}
}

var _ prometheus.Collector = &attachDetachStateCollector{}

func (collector *attachDetachStateCollector) Describe(ch chan<- *prometheus.Desc) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ch <- inUseVolumeMetricDesc
	ch <- totalVolumesMetricDesc
}
func (collector *attachDetachStateCollector) Collect(ch chan<- prometheus.Metric) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	nodeVolumeMap := collector.getVolumeInUseCount()
	for nodeName, pluginCount := range nodeVolumeMap {
		for pluginName, count := range pluginCount {
			metric, err := prometheus.NewConstMetric(inUseVolumeMetricDesc, prometheus.GaugeValue, float64(count), string(nodeName), pluginName)
			if err != nil {
				klog.Warningf("Failed to create metric : %v", err)
			}
			ch <- metric
		}
	}
	stateVolumeMap := collector.getTotalVolumesCount()
	for stateName, pluginCount := range stateVolumeMap {
		for pluginName, count := range pluginCount {
			metric, err := prometheus.NewConstMetric(totalVolumesMetricDesc, prometheus.GaugeValue, float64(count), pluginName, string(stateName))
			if err != nil {
				klog.Warningf("Failed to create metric : %v", err)
			}
			ch <- metric
		}
	}
}
func (collector *attachDetachStateCollector) getVolumeInUseCount() volumeCount {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	pods, err := collector.podLister.List(labels.Everything())
	if err != nil {
		klog.Errorf("Error getting pod list")
		return nil
	}
	nodeVolumeMap := make(volumeCount)
	for _, pod := range pods {
		if len(pod.Spec.Volumes) <= 0 {
			continue
		}
		if pod.Spec.NodeName == "" {
			continue
		}
		for _, podVolume := range pod.Spec.Volumes {
			volumeSpec, err := util.CreateVolumeSpec(podVolume, pod.Namespace, collector.pvcLister, collector.pvLister)
			if err != nil {
				continue
			}
			volumePlugin, err := collector.volumePluginMgr.FindPluginBySpec(volumeSpec)
			if err != nil {
				continue
			}
			pluginName := volumeutil.GetFullQualifiedPluginNameForVolume(volumePlugin.GetPluginName(), volumeSpec)
			nodeVolumeMap.add(pod.Spec.NodeName, pluginName)
		}
	}
	return nodeVolumeMap
}
func (collector *attachDetachStateCollector) getTotalVolumesCount() volumeCount {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	stateVolumeMap := make(volumeCount)
	for _, v := range collector.dsw.GetVolumesToAttach() {
		if plugin, err := collector.volumePluginMgr.FindPluginBySpec(v.VolumeSpec); err == nil {
			pluginName := pluginNameNotAvailable
			if plugin != nil {
				pluginName = volumeutil.GetFullQualifiedPluginNameForVolume(plugin.GetPluginName(), v.VolumeSpec)
			}
			stateVolumeMap.add("desired_state_of_world", pluginName)
		}
	}
	for _, v := range collector.asw.GetAttachedVolumes() {
		if plugin, err := collector.volumePluginMgr.FindPluginBySpec(v.VolumeSpec); err == nil {
			pluginName := pluginNameNotAvailable
			if plugin != nil {
				pluginName = volumeutil.GetFullQualifiedPluginNameForVolume(plugin.GetPluginName(), v.VolumeSpec)
			}
			stateVolumeMap.add("actual_state_of_world", pluginName)
		}
	}
	return stateVolumeMap
}
func RecordForcedDetachMetric() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	forcedDetachMetricCounter.Inc()
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
