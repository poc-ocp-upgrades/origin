package capabilities

import (
	goformat "fmt"
	goos "os"
	godefaultruntime "runtime"
	"sync"
	gotime "time"
)

type Capabilities struct {
	AllowPrivileged                        bool
	PrivilegedSources                      PrivilegedSources
	PerConnectionBandwidthLimitBytesPerSec int64
}
type PrivilegedSources struct {
	HostNetworkSources []string
	HostPIDSources     []string
	HostIPCSources     []string
}

var capInstance struct {
	once         sync.Once
	lock         sync.Mutex
	capabilities *Capabilities
}

func Initialize(c Capabilities) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	capInstance.once.Do(func() {
		capInstance.capabilities = &c
	})
}
func Setup(allowPrivileged bool, privilegedSources PrivilegedSources, perConnectionBytesPerSec int64) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	Initialize(Capabilities{AllowPrivileged: allowPrivileged, PrivilegedSources: privilegedSources, PerConnectionBandwidthLimitBytesPerSec: perConnectionBytesPerSec})
}
func SetForTests(c Capabilities) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	capInstance.lock.Lock()
	defer capInstance.lock.Unlock()
	capInstance.capabilities = &c
}
func Get() Capabilities {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	capInstance.lock.Lock()
	defer capInstance.lock.Unlock()
	if capInstance.capabilities == nil {
		Initialize(Capabilities{AllowPrivileged: false, PrivilegedSources: PrivilegedSources{HostNetworkSources: []string{}, HostPIDSources: []string{}, HostIPCSources: []string{}}})
	}
	return *capInstance.capabilities
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
