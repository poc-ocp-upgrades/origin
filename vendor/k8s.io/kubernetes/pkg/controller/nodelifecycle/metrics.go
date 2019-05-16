package nodelifecycle

import (
	goformat "fmt"
	"github.com/prometheus/client_golang/prometheus"
	goos "os"
	godefaultruntime "runtime"
	"sync"
	gotime "time"
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
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	registerMetrics.Do(func() {
		prometheus.MustRegister(zoneHealth)
		prometheus.MustRegister(zoneSize)
		prometheus.MustRegister(unhealthyNodes)
		prometheus.MustRegister(evictionsNumber)
	})
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
