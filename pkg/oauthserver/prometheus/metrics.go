package metrics

import (
	goformat "fmt"
	"github.com/prometheus/client_golang/prometheus"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

const (
	authSubsystem = "openshift_auth"
)
const (
	SuccessResult = "success"
	FailResult    = "failure"
	ErrorResult   = "error"
)

var (
	authPasswordTotal      = prometheus.NewCounterVec(prometheus.CounterOpts{Subsystem: authSubsystem, Name: "password_total", Help: "Counts total password authentication attempts"}, []string{})
	authFormCounter        = prometheus.NewCounterVec(prometheus.CounterOpts{Subsystem: authSubsystem, Name: "form_password_count", Help: "Counts form password authentication attempts"}, []string{})
	authFormCounterResult  = prometheus.NewCounterVec(prometheus.CounterOpts{Subsystem: authSubsystem, Name: "form_password_count_result", Help: "Counts form password authentication attempts by result"}, []string{"result"})
	authBasicCounter       = prometheus.NewCounterVec(prometheus.CounterOpts{Subsystem: authSubsystem, Name: "basic_password_count", Help: "Counts basic password authentication attempts"}, []string{})
	authBasicCounterResult = prometheus.NewCounterVec(prometheus.CounterOpts{Subsystem: authSubsystem, Name: "basic_password_count_result", Help: "Counts basic password authentication attempts by result"}, []string{"result"})
)

func init() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	prometheus.MustRegister(authPasswordTotal)
	prometheus.MustRegister(authFormCounter)
	prometheus.MustRegister(authFormCounterResult)
	prometheus.MustRegister(authBasicCounter)
	prometheus.MustRegister(authBasicCounterResult)
}
func RecordBasicPasswordAuth(result string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	authPasswordTotal.WithLabelValues().Inc()
	authBasicCounter.WithLabelValues().Inc()
	authBasicCounterResult.WithLabelValues(result).Inc()
}
func RecordFormPasswordAuth(result string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	authPasswordTotal.WithLabelValues().Inc()
	authFormCounter.WithLabelValues().Inc()
	authFormCounterResult.WithLabelValues(result).Inc()
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
