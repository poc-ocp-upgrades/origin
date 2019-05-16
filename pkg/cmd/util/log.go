package util

import (
	"github.com/openshift/library-go/pkg/serviceability"
	"io"
	"k8s.io/klog"
)

func NewGLogWriterV(level int) io.Writer {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &gLogWriter{level: klog.Level(level)}
}

type gLogWriter struct{ level klog.Level }

func (w *gLogWriter) Write(p []byte) (n int, err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if klog.V(w.level) {
		klog.InfoDepth(2, string(p))
	}
	return len(p), nil
}
func InitLogrus() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	switch {
	case bool(klog.V(4)):
		serviceability.InitLogrus("DEBUG")
	case bool(klog.V(2)):
		serviceability.InitLogrus("INFO")
	case bool(klog.V(0)):
		serviceability.InitLogrus("WARN")
	}
}
