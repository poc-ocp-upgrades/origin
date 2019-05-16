package userspace

import (
	"fmt"
	"io"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/proxy"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"
)

type ProxySocket interface {
	Addr() net.Addr
	Close() error
	ProxyLoop(service proxy.ServicePortName, info *ServiceInfo, loadBalancer LoadBalancer)
	ListenPort() int
}

func newProxySocket(protocol v1.Protocol, ip net.IP, port int) (ProxySocket, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	host := ""
	if ip != nil {
		host = ip.String()
	}
	switch strings.ToUpper(string(protocol)) {
	case "TCP":
		listener, err := net.Listen("tcp", net.JoinHostPort(host, strconv.Itoa(port)))
		if err != nil {
			return nil, err
		}
		return &tcpProxySocket{Listener: listener, port: port}, nil
	case "UDP":
		addr, err := net.ResolveUDPAddr("udp", net.JoinHostPort(host, strconv.Itoa(port)))
		if err != nil {
			return nil, err
		}
		conn, err := net.ListenUDP("udp", addr)
		if err != nil {
			return nil, err
		}
		return &udpProxySocket{UDPConn: conn, port: port}, nil
	case "SCTP":
		return nil, fmt.Errorf("SCTP is not supported for user space proxy")
	}
	return nil, fmt.Errorf("unknown protocol %q", protocol)
}

var EndpointDialTimeouts = []time.Duration{250 * time.Millisecond, 500 * time.Millisecond, 1 * time.Second, 2 * time.Second}

type tcpProxySocket struct {
	net.Listener
	port int
}

func (tcp *tcpProxySocket) ListenPort() int {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return tcp.port
}
func TryConnectEndpoints(service proxy.ServicePortName, srcAddr net.Addr, protocol string, loadBalancer LoadBalancer) (out net.Conn, err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	sessionAffinityReset := false
	for _, dialTimeout := range EndpointDialTimeouts {
		endpoint, err := loadBalancer.NextEndpoint(service, srcAddr, sessionAffinityReset)
		if err != nil {
			klog.Errorf("Couldn't find an endpoint for %s: %v", service, err)
			return nil, err
		}
		klog.V(3).Infof("Mapped service %q to endpoint %s", service, endpoint)
		outConn, err := net.DialTimeout(protocol, endpoint, dialTimeout)
		if err != nil {
			if isTooManyFDsError(err) {
				panic("Dial failed: " + err.Error())
			}
			klog.Errorf("Dial failed: %v", err)
			sessionAffinityReset = true
			continue
		}
		return outConn, nil
	}
	return nil, fmt.Errorf("failed to connect to an endpoint.")
}
func (tcp *tcpProxySocket) ProxyLoop(service proxy.ServicePortName, myInfo *ServiceInfo, loadBalancer LoadBalancer) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for {
		if !myInfo.IsAlive() {
			return
		}
		inConn, err := tcp.Accept()
		if err != nil {
			if isTooManyFDsError(err) {
				panic("Accept failed: " + err.Error())
			}
			if isClosedError(err) {
				return
			}
			if !myInfo.IsAlive() {
				return
			}
			klog.Errorf("Accept failed: %v", err)
			continue
		}
		klog.V(3).Infof("Accepted TCP connection from %v to %v", inConn.RemoteAddr(), inConn.LocalAddr())
		outConn, err := TryConnectEndpoints(service, inConn.(*net.TCPConn).RemoteAddr(), "tcp", loadBalancer)
		if err != nil {
			klog.Errorf("Failed to connect to balancer: %v", err)
			inConn.Close()
			continue
		}
		go ProxyTCP(inConn.(*net.TCPConn), outConn.(*net.TCPConn))
	}
}
func ProxyTCP(in, out *net.TCPConn) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var wg sync.WaitGroup
	wg.Add(2)
	klog.V(4).Infof("Creating proxy between %v <-> %v <-> %v <-> %v", in.RemoteAddr(), in.LocalAddr(), out.LocalAddr(), out.RemoteAddr())
	go copyBytes("from backend", in, out, &wg)
	go copyBytes("to backend", out, in, &wg)
	wg.Wait()
}
func copyBytes(direction string, dest, src *net.TCPConn, wg *sync.WaitGroup) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	defer wg.Done()
	klog.V(4).Infof("Copying %s: %s -> %s", direction, src.RemoteAddr(), dest.RemoteAddr())
	n, err := io.Copy(dest, src)
	if err != nil {
		if !isClosedError(err) {
			klog.Errorf("I/O error: %v", err)
		}
	}
	klog.V(4).Infof("Copied %d bytes %s: %s -> %s", n, direction, src.RemoteAddr(), dest.RemoteAddr())
	dest.Close()
	src.Close()
}

type udpProxySocket struct {
	*net.UDPConn
	port int
}

func (udp *udpProxySocket) ListenPort() int {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return udp.port
}
func (udp *udpProxySocket) Addr() net.Addr {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return udp.LocalAddr()
}

type ClientCache struct {
	Mu      sync.Mutex
	Clients map[string]net.Conn
}

func newClientCache() *ClientCache {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &ClientCache{Clients: map[string]net.Conn{}}
}
func (udp *udpProxySocket) ProxyLoop(service proxy.ServicePortName, myInfo *ServiceInfo, loadBalancer LoadBalancer) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var buffer [4096]byte
	for {
		if !myInfo.IsAlive() {
			break
		}
		n, cliAddr, err := udp.ReadFrom(buffer[0:])
		if err != nil {
			if e, ok := err.(net.Error); ok {
				if e.Temporary() {
					klog.V(1).Infof("ReadFrom had a temporary failure: %v", err)
					continue
				}
			}
			klog.Errorf("ReadFrom failed, exiting ProxyLoop: %v", err)
			break
		}
		svrConn, err := udp.getBackendConn(myInfo.ActiveClients, cliAddr, loadBalancer, service, myInfo.Timeout)
		if err != nil {
			continue
		}
		_, err = svrConn.Write(buffer[0:n])
		if err != nil {
			if !logTimeout(err) {
				klog.Errorf("Write failed: %v", err)
			}
			continue
		}
		err = svrConn.SetDeadline(time.Now().Add(myInfo.Timeout))
		if err != nil {
			klog.Errorf("SetDeadline failed: %v", err)
			continue
		}
	}
}
func (udp *udpProxySocket) getBackendConn(activeClients *ClientCache, cliAddr net.Addr, loadBalancer LoadBalancer, service proxy.ServicePortName, timeout time.Duration) (net.Conn, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	activeClients.Mu.Lock()
	defer activeClients.Mu.Unlock()
	svrConn, found := activeClients.Clients[cliAddr.String()]
	if !found {
		klog.V(3).Infof("New UDP connection from %s", cliAddr)
		var err error
		svrConn, err = TryConnectEndpoints(service, cliAddr, "udp", loadBalancer)
		if err != nil {
			return nil, err
		}
		if err = svrConn.SetDeadline(time.Now().Add(timeout)); err != nil {
			klog.Errorf("SetDeadline failed: %v", err)
			return nil, err
		}
		activeClients.Clients[cliAddr.String()] = svrConn
		go func(cliAddr net.Addr, svrConn net.Conn, activeClients *ClientCache, timeout time.Duration) {
			defer runtime.HandleCrash()
			udp.proxyClient(cliAddr, svrConn, activeClients, timeout)
		}(cliAddr, svrConn, activeClients, timeout)
	}
	return svrConn, nil
}
func (udp *udpProxySocket) proxyClient(cliAddr net.Addr, svrConn net.Conn, activeClients *ClientCache, timeout time.Duration) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	defer svrConn.Close()
	var buffer [4096]byte
	for {
		n, err := svrConn.Read(buffer[0:])
		if err != nil {
			if !logTimeout(err) {
				klog.Errorf("Read failed: %v", err)
			}
			break
		}
		err = svrConn.SetDeadline(time.Now().Add(timeout))
		if err != nil {
			klog.Errorf("SetDeadline failed: %v", err)
			break
		}
		n, err = udp.WriteTo(buffer[0:n], cliAddr)
		if err != nil {
			if !logTimeout(err) {
				klog.Errorf("WriteTo failed: %v", err)
			}
			break
		}
	}
	activeClients.Mu.Lock()
	delete(activeClients.Clients, cliAddr.String())
	activeClients.Mu.Unlock()
}
