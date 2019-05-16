package glog

import (
	"fmt"
	goformat "fmt"
	"io"
	"k8s.io/klog"
	goos "os"
	godefaultruntime "runtime"
	"strings"
	gotime "time"
)

type Logger interface {
	Is(level int) bool
	V(level int) Logger
	Infof(format string, args ...interface{})
}

func ToFile(w io.Writer, level int) Logger {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return file{w, level}
}

var (
	None Logger = discard{}
	Log  Logger = glogger{}
)

type discard struct{}

func (discard) Is(level int) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return false
}
func (discard) V(level int) Logger {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return None
}
func (discard) Infof(_ string, _ ...interface{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
}

type glogger struct{}

func (glogger) Is(level int) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return bool(klog.V(klog.Level(level)))
}
func (glogger) V(level int) Logger {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return gverbose{klog.V(klog.Level(level))}
}
func (glogger) Infof(format string, args ...interface{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.InfoDepth(2, fmt.Sprintf(format, args...))
}

type gverbose struct{ klog.Verbose }

func (gverbose) Is(level int) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return bool(klog.V(klog.Level(level)))
}
func (gverbose) V(level int) Logger {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if klog.V(klog.Level(level)) {
		return Log
	}
	return None
}
func (g gverbose) Infof(format string, args ...interface{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if g.Verbose {
		klog.InfoDepth(2, fmt.Sprintf(format, args...))
	}
}

type file struct {
	w     io.Writer
	level int
}

func (f file) Is(level int) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return level <= f.level || bool(klog.V(klog.Level(level)))
}
func (f file) V(level int) Logger {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if !klog.V(klog.Level(level)) {
		return None
	}
	if level > f.level {
		return Log
	}
	return f
}
func (f file) Infof(format string, args ...interface{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	fmt.Fprintf(f.w, format, args...)
	if !strings.HasSuffix(format, "\n") {
		fmt.Fprintln(f.w)
	}
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
