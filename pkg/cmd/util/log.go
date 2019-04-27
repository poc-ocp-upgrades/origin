package util

import (
	"io"
	"k8s.io/klog"
	"github.com/openshift/library-go/pkg/serviceability"
)

func NewGLogWriterV(level int) io.Writer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &gLogWriter{level: klog.Level(level)}
}

type gLogWriter struct{ level klog.Level }

func (w *gLogWriter) Write(p []byte) (n int, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if klog.V(w.level) {
		klog.InfoDepth(2, string(p))
	}
	return len(p), nil
}
func InitLogrus() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch {
	case bool(klog.V(4)):
		serviceability.InitLogrus("DEBUG")
	case bool(klog.V(2)):
		serviceability.InitLogrus("INFO")
	case bool(klog.V(0)):
		serviceability.InitLogrus("WARN")
	}
}
