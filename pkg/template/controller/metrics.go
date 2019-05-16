package controller

import (
	goformat "fmt"
	templatev1 "github.com/openshift/api/template/v1"
	"github.com/prometheus/client_golang/prometheus"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	goos "os"
	godefaultruntime "runtime"
	"time"
	gotime "time"
)

var templateInstanceCompleted = prometheus.NewCounterVec(prometheus.CounterOpts{Name: "openshift_template_instance_completed_total", Help: "Counts completed TemplateInstance objects by condition"}, []string{"condition"})

func newTemplateInstanceActiveAge() prometheus.Histogram {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return prometheus.NewHistogram(prometheus.HistogramOpts{Name: "openshift_template_instance_active_age_seconds", Help: "Shows the instantaneous age distribution of active TemplateInstance objects", Buckets: prometheus.LinearBuckets(600, 600, 7)})
}
func (c *TemplateInstanceController) Describe(ch chan<- *prometheus.Desc) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	templateInstanceActiveAge := newTemplateInstanceActiveAge()
	templateInstanceCompleted.Describe(ch)
	templateInstanceActiveAge.Describe(ch)
}
func (c *TemplateInstanceController) Collect(ch chan<- prometheus.Metric) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	templateInstanceCompleted.Collect(ch)
	now := c.clock.Now()
	templateInstances, err := c.lister.List(labels.Everything())
	if err != nil {
		utilruntime.HandleError(err)
		return
	}
	templateInstanceActiveAge := newTemplateInstanceActiveAge()
nextTemplateInstance:
	for _, templateInstance := range templateInstances {
		for _, cond := range templateInstance.Status.Conditions {
			if cond.Status == corev1.ConditionTrue && (cond.Type == templatev1.TemplateInstanceInstantiateFailure || cond.Type == templatev1.TemplateInstanceReady) {
				continue nextTemplateInstance
			}
		}
		templateInstanceActiveAge.Observe(float64(now.Sub(templateInstance.CreationTimestamp.Time) / time.Second))
	}
	templateInstanceActiveAge.Collect(ch)
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
