package testing

import (
	goformat "fmt"
	"net"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

type FakeNetwork struct {
	NetworkInterfaces []net.Interface
	Address           map[string][]net.Addr
}

func NewFakeNetwork() *FakeNetwork {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &FakeNetwork{NetworkInterfaces: make([]net.Interface, 0), Address: make(map[string][]net.Addr)}
}
func (f *FakeNetwork) AddInterfaceAddr(intf *net.Interface, addrs []net.Addr) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	f.NetworkInterfaces = append(f.NetworkInterfaces, *intf)
	f.Address[intf.Name] = addrs
}
func (f *FakeNetwork) Addrs(intf *net.Interface) ([]net.Addr, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return f.Address[intf.Name], nil
}
func (f *FakeNetwork) Interfaces() ([]net.Interface, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return f.NetworkInterfaces, nil
}

type AddrStruct struct{ Val string }

func (a AddrStruct) Network() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return a.Val
}
func (a AddrStruct) String() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return a.Val
}

var _ net.Addr = &AddrStruct{}

func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
