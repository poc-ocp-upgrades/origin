package prometheus

import (
	"fmt"
	goformat "fmt"
	"github.com/openshift/origin/pkg/apps/util"
	"github.com/prometheus/client_golang/prometheus"
	"k8s.io/apimachinery/pkg/labels"
	kcorelisters "k8s.io/client-go/listers/core/v1"
	"k8s.io/klog"
	goos "os"
	godefaultruntime "runtime"
	"strings"
	"time"
	gotime "time"
)

const (
	completeRolloutCount         = "complete_rollouts_total"
	activeRolloutDurationSeconds = "active_rollouts_duration_seconds"
	lastFailedRolloutTime        = "last_failed_rollout_time"
	availablePhase               = "available"
	failedPhase                  = "failed"
	cancelledPhase               = "cancelled"
)

var (
	nameToQuery = func(name string) string {
		return strings.Join([]string{"openshift_apps_deploymentconfigs", name}, "_")
	}
	completeRolloutCountDesc         = prometheus.NewDesc(nameToQuery(completeRolloutCount), "Counts total complete rollouts", []string{"phase"}, nil)
	lastFailedRolloutTimeDesc        = prometheus.NewDesc(nameToQuery(lastFailedRolloutTime), "Tracks the time of last failure rollout per deployment config", []string{"namespace", "name", "latest_version"}, nil)
	activeRolloutDurationSecondsDesc = prometheus.NewDesc(nameToQuery(activeRolloutDurationSeconds), "Tracks the active rollout duration in seconds", []string{"namespace", "name", "phase", "latest_version"}, nil)
	apps                             = appsCollector{}
	registered                       = false
)

type appsCollector struct {
	lister kcorelisters.ReplicationControllerLister
	nowFn  func() time.Time
}

func InitializeMetricsCollector(rcLister kcorelisters.ReplicationControllerLister) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	apps.lister = rcLister
	apps.nowFn = time.Now
	if !registered {
		prometheus.MustRegister(&apps)
		registered = true
	}
	klog.V(4).Info("apps metrics registered with prometheus")
}
func (c *appsCollector) Describe(ch chan<- *prometheus.Desc) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ch <- completeRolloutCountDesc
	ch <- activeRolloutDurationSecondsDesc
}

type failedRollout struct {
	timestamp     float64
	latestVersion int64
}

func (c *appsCollector) Collect(ch chan<- prometheus.Metric) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	result, err := c.lister.List(labels.Everything())
	if err != nil {
		klog.V(4).Infof("Collecting metrics for apps failed: %v", err)
		return
	}
	var available, failed, cancelled float64
	latestFailedRollouts := map[string]failedRollout{}
	for _, d := range result {
		dcName := util.DeploymentConfigNameFor(d)
		if len(dcName) == 0 {
			continue
		}
		latestVersion := util.DeploymentVersionFor(d)
		key := d.Namespace + "/" + dcName
		if util.IsTerminatedDeployment(d) {
			if util.IsDeploymentCancelled(d) {
				cancelled++
				continue
			}
			if util.IsFailedDeployment(d) {
				failed++
				if r, exists := latestFailedRollouts[key]; exists && latestVersion <= r.latestVersion {
					continue
				}
				latestFailedRollouts[key] = failedRollout{timestamp: float64(d.CreationTimestamp.Unix()), latestVersion: latestVersion}
				continue
			}
			if util.IsCompleteDeployment(d) {
				if r, hasFailedRollout := latestFailedRollouts[key]; hasFailedRollout && r.latestVersion < latestVersion {
					delete(latestFailedRollouts, key)
				}
				available++
				continue
			}
		}
		phase := strings.ToLower(string(util.DeploymentStatusFor(d)))
		if len(phase) == 0 {
			phase = "unknown"
		}
		durationSeconds := c.nowFn().Unix() - d.CreationTimestamp.Unix()
		ch <- prometheus.MustNewConstMetric(activeRolloutDurationSecondsDesc, prometheus.CounterValue, float64(durationSeconds), []string{d.Namespace, dcName, phase, fmt.Sprintf("%d", latestVersion)}...)
	}
	for dc, r := range latestFailedRollouts {
		parts := strings.Split(dc, "/")
		ch <- prometheus.MustNewConstMetric(lastFailedRolloutTimeDesc, prometheus.GaugeValue, r.timestamp, []string{parts[0], parts[1], fmt.Sprintf("%d", r.latestVersion)}...)
	}
	ch <- prometheus.MustNewConstMetric(completeRolloutCountDesc, prometheus.GaugeValue, available, []string{availablePhase}...)
	ch <- prometheus.MustNewConstMetric(completeRolloutCountDesc, prometheus.GaugeValue, failed, []string{failedPhase}...)
	ch <- prometheus.MustNewConstMetric(completeRolloutCountDesc, prometheus.GaugeValue, cancelled, []string{cancelledPhase}...)
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
