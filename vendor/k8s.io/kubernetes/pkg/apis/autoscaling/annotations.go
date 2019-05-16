package autoscaling

import (
	goformat "fmt"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

const MetricSpecsAnnotation = "autoscaling.alpha.kubernetes.io/metrics"
const MetricStatusesAnnotation = "autoscaling.alpha.kubernetes.io/current-metrics"
const HorizontalPodAutoscalerConditionsAnnotation = "autoscaling.alpha.kubernetes.io/conditions"
const DefaultCPUUtilization = 80

func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
