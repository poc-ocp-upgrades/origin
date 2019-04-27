package unidler

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
	"k8s.io/klog"
	corev1 "k8s.io/api/core/v1"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/kubernetes/pkg/proxy"
	"k8s.io/kubernetes/pkg/proxy/userspace"
)

const (
	UDPBufferSize	= 4096
	NeedPodsReason	= "NeedPods"
)

var endpointDialTimeout = []time.Duration{250 * time.Millisecond, 500 * time.Millisecond, 1 * time.Second, 2 * time.Second}

type connectionList struct {
	conns		[]heldConn
	maxSize		int
	tickSize	time.Duration
	timeSinceStart	time.Duration
	timeout		time.Duration
	svcName		string
}
type heldConn struct {
	net.Conn
	connectedAt	time.Duration
}

func newConnectionList(maxSize int, tickSize time.Duration, timeout time.Duration, svcName string) *connectionList {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &connectionList{conns: []heldConn{}, maxSize: maxSize, tickSize: tickSize, timeSinceStart: 0, timeout: timeout, svcName: svcName}
}
func (l *connectionList) Add(conn net.Conn) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(l.conns) >= l.maxSize {
		utilruntime.HandleError(fmt.Errorf("max connections exceeded while waiting for idled service %s to awaken, dropping oldest", l.svcName))
		var oldConn net.Conn
		oldConn, l.conns = l.conns[0], l.conns[1:]
		oldConn.Close()
	}
	l.conns = append(l.conns, heldConn{conn, l.timeSinceStart})
}
func (l *connectionList) Tick() {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	l.timeSinceStart += l.tickSize
	l.cleanOldConnections()
}
func (l *connectionList) cleanOldConnections() {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	cleanInd := -1
	for i, conn := range l.conns {
		if l.timeSinceStart-conn.connectedAt < l.timeout {
			cleanInd = i
			break
		}
	}
	if cleanInd > 0 {
		oldConns := l.conns[:cleanInd]
		l.conns = l.conns[cleanInd:]
		utilruntime.HandleError(fmt.Errorf("timed out %v connections while waiting for idled service %s to awaken.", len(oldConns), l.svcName))
		for _, conn := range oldConns {
			conn.Close()
		}
	}
}
func (l *connectionList) GetConns() []net.Conn {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	conns := make([]net.Conn, len(l.conns))
	for i, conn := range l.conns {
		conns[i] = conn.Conn
	}
	return conns
}
func (l *connectionList) Len() int {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	return len(l.conns)
}
func (l *connectionList) Clear() {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, conn := range l.conns {
		conn.Close()
	}
	l.conns = []heldConn{}
}

var (
	MaxHeldConnections	= 16
	needPodsWaitTimeout	= 120 * time.Second
	needPodsTickLen		= 5 * time.Second
)

func newUnidlerSocket(protocol corev1.Protocol, ip net.IP, port int, signaler NeedPodsSignaler) (userspace.ProxySocket, error) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
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
		return &tcpUnidlerSocket{Listener: listener, port: port, signaler: signaler}, nil
	case "UDP":
		addr, err := net.ResolveUDPAddr("udp", net.JoinHostPort(host, strconv.Itoa(port)))
		if err != nil {
			return nil, err
		}
		conn, err := net.ListenUDP("udp", addr)
		if err != nil {
			return nil, err
		}
		return &udpUnidlerSocket{UDPConn: conn, port: port, signaler: signaler}, nil
	}
	return nil, fmt.Errorf("unknown protocol %q", protocol)
}

type tcpUnidlerSocket struct {
	net.Listener
	port		int
	signaler	NeedPodsSignaler
}

func (tcp *tcpUnidlerSocket) ListenPort() int {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	return tcp.port
}
func (tcp *tcpUnidlerSocket) waitForEndpoints(ch chan<- interface{}, service proxy.ServicePortName, loadBalancer userspace.LoadBalancer) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer close(ch)
	for {
		if loadBalancer.ServiceHasEndpoints(service) {
			return
		}
		time.Sleep(endpointDialTimeout[0])
	}
}
func (tcp *tcpUnidlerSocket) acceptConns(ch chan<- net.Conn, svcInfo *userspace.ServiceInfo) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer close(ch)
	for {
		inConn, err := tcp.Accept()
		if err != nil {
			if isTooManyFDsError(err) {
				panic("Accept failed: " + err.Error())
			}
			if isClosedError(err) {
				return
			}
			if !svcInfo.IsAlive() {
				return
			}
			utilruntime.HandleError(fmt.Errorf("Accept failed: %v", err))
			continue
		}
		ch <- inConn
	}
}
func (tcp *tcpUnidlerSocket) awaitAwakening(service proxy.ServicePortName, loadBalancer userspace.LoadBalancer, inConns <-chan net.Conn, endpointsAvail chan<- interface{}) (*connectionList, bool) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	sent_need_pods := false
	timeout_started := false
	ticker := time.NewTicker(needPodsTickLen)
	defer ticker.Stop()
	svcName := fmt.Sprintf("%s/%s:%s", service.Namespace, service.Name, service.Port)
	allConns := newConnectionList(MaxHeldConnections, needPodsTickLen, needPodsWaitTimeout, svcName)
	for {
		select {
		case inConn, ok := <-inConns:
			if !ok {
				return allConns, false
			}
			if !sent_need_pods && !loadBalancer.ServiceHasEndpoints(service) {
				klog.V(4).Infof("unidling TCP proxy sent unidle event to wake up service %s/%s:%s", service.Namespace, service.Name, service.Port)
				tcp.signaler.NeedPods(service.NamespacedName, service.Port)
				sent_need_pods = true
				timeout_started = true
			}
			if allConns.Len() == 0 {
				if !loadBalancer.ServiceHasEndpoints(service) {
					go tcp.waitForEndpoints(endpointsAvail, service, loadBalancer)
				}
			}
			allConns.Add(inConn)
			klog.V(4).Infof("unidling TCP proxy has accumulated %v connections while waiting for service %s/%s:%s to unidle", allConns.Len(), service.Namespace, service.Name, service.Port)
		case <-ticker.C:
			if !timeout_started {
				continue
			}
			allConns.Tick()
		}
	}
}
func (tcp *tcpUnidlerSocket) ProxyLoop(service proxy.ServicePortName, svcInfo *userspace.ServiceInfo, loadBalancer userspace.LoadBalancer) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	if !svcInfo.IsAlive() {
		return
	}
	inConns := make(chan net.Conn)
	go tcp.acceptConns(inConns, svcInfo)
	endpointsAvail := make(chan interface{})
	var allConns *connectionList
	for {
		klog.V(4).Infof("unidling TCP proxy start/reset for service %s/%s:%s", service.Namespace, service.Name, service.Port)
		var cont bool
		if allConns, cont = tcp.awaitAwakening(service, loadBalancer, inConns, endpointsAvail); !cont {
			break
		}
	}
	klog.V(4).Infof("unidling TCP proxy waiting for endpoints for service %s/%s:%s to become available with %v accumulated connections", service.Namespace, service.Name, service.Port, allConns.Len())
	select {
	case _, ok := <-endpointsAvail:
		if ok {
			close(endpointsAvail)
		}
	case <-time.NewTimer(needPodsWaitTimeout).C:
		if allConns.Len() > 0 {
			utilruntime.HandleError(fmt.Errorf("timed out %v TCP connections while waiting for idled service %s/%s:%s to awaken.", allConns.Len(), service.Namespace, service.Name, service.Port))
			allConns.Clear()
		}
		return
	}
	klog.V(4).Infof("unidling TCP proxy got endpoints for service %s/%s:%s, connecting %v accumulated connections", service.Namespace, service.Name, service.Port, allConns.Len())
	for _, inConn := range allConns.GetConns() {
		klog.V(3).Infof("Accepted TCP connection from %v to %v", inConn.RemoteAddr(), inConn.LocalAddr())
		outConn, err := userspace.TryConnectEndpoints(service, inConn.(*net.TCPConn).RemoteAddr(), "tcp", loadBalancer)
		if err != nil {
			utilruntime.HandleError(fmt.Errorf("Failed to connect to balancer: %v", err))
			inConn.Close()
			continue
		}
		go userspace.ProxyTCP(inConn.(*net.TCPConn), outConn.(*net.TCPConn))
	}
}

type udpUnidlerSocket struct {
	*net.UDPConn
	port		int
	signaler	NeedPodsSignaler
}

func (udp *udpUnidlerSocket) ListenPort() int {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	return udp.port
}
func (udp *udpUnidlerSocket) Addr() net.Addr {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	return udp.LocalAddr()
}
func (udp *udpUnidlerSocket) readFromSock(buffer []byte, svcInfo *userspace.ServiceInfo) bool {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	if !svcInfo.IsAlive() {
		return false
	}
	_, _, err := udp.ReadFrom(buffer)
	if err != nil {
		if e, ok := err.(net.Error); ok {
			if e.Temporary() {
				klog.V(1).Infof("ReadFrom had a temporary failure: %v", err)
				return true
			}
		}
		utilruntime.HandleError(fmt.Errorf("ReadFrom failed, exiting ProxyLoop: %v", err))
		return false
	}
	return true
}
func (udp *udpUnidlerSocket) sendWakeup(svcPortName proxy.ServicePortName, svcInfo *userspace.ServiceInfo) *time.Timer {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	timeoutTimer := time.NewTimer(needPodsWaitTimeout)
	klog.V(4).Infof("unidling proxy sent unidle event to wake up service %s/%s:%s", svcPortName.Namespace, svcPortName.Name, svcPortName.Port)
	udp.signaler.NeedPods(svcPortName.NamespacedName, svcPortName.Port)
	return timeoutTimer
}
func (udp *udpUnidlerSocket) ProxyLoop(svcPortName proxy.ServicePortName, svcInfo *userspace.ServiceInfo, loadBalancer userspace.LoadBalancer) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	var buffer [UDPBufferSize]byte
	klog.V(4).Infof("unidling proxy UDP proxy waiting for data for service %s/%s:%s", svcPortName.Namespace, svcPortName.Name, svcPortName.Port)
	if !udp.readFromSock(buffer[0:], svcInfo) {
		return
	}
	wakeupTimeoutTimer := udp.sendWakeup(svcPortName, svcInfo)
	for {
		if !udp.readFromSock(buffer[0:], svcInfo) {
			break
		}
		if active := wakeupTimeoutTimer.Reset(needPodsWaitTimeout); !active {
			wakeupTimeoutTimer = udp.sendWakeup(svcPortName, svcInfo)
		}
	}
}
func isTooManyFDsError(err error) bool {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	return strings.Contains(err.Error(), "too many open files")
}
func isClosedError(err error) bool {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	return strings.HasSuffix(err.Error(), "use of closed network connection")
}
