package winuserspace

import (
 "fmt"
 "io"
 "net"
 "strconv"
 "strings"
 "sync"
 "sync/atomic"
 "time"
 "github.com/miekg/dns"
 "k8s.io/api/core/v1"
 "k8s.io/apimachinery/pkg/types"
 "k8s.io/apimachinery/pkg/util/runtime"
 "k8s.io/klog"
 "k8s.io/kubernetes/pkg/proxy"
 "k8s.io/kubernetes/pkg/util/ipconfig"
 "k8s.io/utils/exec"
)

const (
 clusterDomain                 = "cluster.local"
 serviceDomain                 = "svc." + clusterDomain
 namespaceServiceDomain        = "default." + serviceDomain
 dnsPortName                   = "dns"
 dnsTypeA               uint16 = 0x01
 dnsTypeAAAA            uint16 = 0x1c
 dnsClassInternet       uint16 = 0x01
)

type proxySocket interface {
 Addr() net.Addr
 Close() error
 ProxyLoop(service ServicePortPortalName, info *serviceInfo, proxier *Proxier)
 ListenPort() int
}

func newProxySocket(protocol v1.Protocol, ip net.IP, port int) (proxySocket, error) {
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

var endpointDialTimeout = []time.Duration{250 * time.Millisecond, 500 * time.Millisecond, 1 * time.Second, 2 * time.Second}

type tcpProxySocket struct {
 net.Listener
 port int
}

func (tcp *tcpProxySocket) ListenPort() int {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return tcp.port
}
func tryConnect(service ServicePortPortalName, srcAddr net.Addr, protocol string, proxier *Proxier) (out net.Conn, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 sessionAffinityReset := false
 for _, dialTimeout := range endpointDialTimeout {
  servicePortName := proxy.ServicePortName{NamespacedName: types.NamespacedName{Namespace: service.Namespace, Name: service.Name}, Port: service.Port}
  endpoint, err := proxier.loadBalancer.NextEndpoint(servicePortName, srcAddr, sessionAffinityReset)
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
func (tcp *tcpProxySocket) ProxyLoop(service ServicePortPortalName, myInfo *serviceInfo, proxier *Proxier) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for {
  if !myInfo.isAlive() {
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
   if !myInfo.isAlive() {
    return
   }
   klog.Errorf("Accept failed: %v", err)
   continue
  }
  klog.V(3).Infof("Accepted TCP connection from %v to %v", inConn.RemoteAddr(), inConn.LocalAddr())
  outConn, err := tryConnect(service, inConn.(*net.TCPConn).RemoteAddr(), "tcp", proxier)
  if err != nil {
   klog.Errorf("Failed to connect to balancer: %v", err)
   inConn.Close()
   continue
  }
  go proxyTCP(inConn.(*net.TCPConn), outConn.(*net.TCPConn))
 }
}
func proxyTCP(in, out *net.TCPConn) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var wg sync.WaitGroup
 wg.Add(2)
 klog.V(4).Infof("Creating proxy between %v <-> %v <-> %v <-> %v", in.RemoteAddr(), in.LocalAddr(), out.LocalAddr(), out.RemoteAddr())
 go copyBytes("from backend", in, out, &wg)
 go copyBytes("to backend", out, in, &wg)
 wg.Wait()
}
func copyBytes(direction string, dest, src *net.TCPConn, wg *sync.WaitGroup) {
 _logClusterCodePath()
 defer _logClusterCodePath()
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
 _logClusterCodePath()
 defer _logClusterCodePath()
 return udp.port
}
func (udp *udpProxySocket) Addr() net.Addr {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return udp.LocalAddr()
}

type clientCache struct {
 mu      sync.Mutex
 clients map[string]net.Conn
}

func newClientCache() *clientCache {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &clientCache{clients: map[string]net.Conn{}}
}

type dnsClientQuery struct {
 clientAddress string
 dnsQType      uint16
}
type dnsClientCache struct {
 mu      sync.Mutex
 clients map[dnsClientQuery]*dnsQueryState
}
type dnsQueryState struct {
 searchIndex int32
 msg         *dns.Msg
}

func newDNSClientCache() *dnsClientCache {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &dnsClientCache{clients: map[dnsClientQuery]*dnsQueryState{}}
}
func packetRequiresDNSSuffix(dnsType, dnsClass uint16) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return (dnsType == dnsTypeA || dnsType == dnsTypeAAAA) && dnsClass == dnsClassInternet
}
func isDNSService(portName string) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return portName == dnsPortName
}
func appendDNSSuffix(msg *dns.Msg, buffer []byte, length int, dnsSuffix string) (int, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if msg == nil || len(msg.Question) == 0 {
  return length, fmt.Errorf("DNS message parameter is invalid")
 }
 origName := msg.Question[0].Name
 if dnsSuffix != "" {
  msg.Question[0].Name += dnsSuffix + "."
 }
 mbuf, err := msg.PackBuffer(buffer)
 msg.Question[0].Name = origName
 if err != nil {
  klog.Warningf("Unable to pack DNS packet. Error is: %v", err)
  return length, err
 }
 if &buffer[0] != &mbuf[0] {
  return length, fmt.Errorf("Buffer is too small in packing DNS packet")
 }
 return len(mbuf), nil
}
func recoverDNSQuestion(origName string, msg *dns.Msg, buffer []byte, length int) (int, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if msg == nil || len(msg.Question) == 0 {
  return length, fmt.Errorf("DNS message parameter is invalid")
 }
 if origName == msg.Question[0].Name {
  return length, nil
 }
 msg.Question[0].Name = origName
 if len(msg.Answer) > 0 {
  msg.Answer[0].Header().Name = origName
 }
 mbuf, err := msg.PackBuffer(buffer)
 if err != nil {
  klog.Warningf("Unable to pack DNS packet. Error is: %v", err)
  return length, err
 }
 if &buffer[0] != &mbuf[0] {
  return length, fmt.Errorf("Buffer is too small in packing DNS packet")
 }
 return len(mbuf), nil
}
func processUnpackedDNSQueryPacket(dnsClients *dnsClientCache, msg *dns.Msg, host string, dnsQType uint16, buffer []byte, length int, dnsSearch []string) int {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if dnsSearch == nil || len(dnsSearch) == 0 {
  klog.V(1).Infof("DNS search list is not initialized and is empty.")
  return length
 }
 dnsClients.mu.Lock()
 state, found := dnsClients.clients[dnsClientQuery{host, dnsQType}]
 if !found {
  state = &dnsQueryState{0, msg}
  dnsClients.clients[dnsClientQuery{host, dnsQType}] = state
 }
 dnsClients.mu.Unlock()
 index := atomic.SwapInt32(&state.searchIndex, state.searchIndex+1)
 state.msg.MsgHdr.Id = msg.MsgHdr.Id
 if index < 0 || index >= int32(len(dnsSearch)) {
  klog.V(1).Infof("Search index %d is out of range.", index)
  return length
 }
 length, err := appendDNSSuffix(msg, buffer, length, dnsSearch[index])
 if err != nil {
  klog.Errorf("Append DNS suffix failed: %v", err)
 }
 return length
}
func processUnpackedDNSResponsePacket(svrConn net.Conn, dnsClients *dnsClientCache, msg *dns.Msg, rcode int, host string, dnsQType uint16, buffer []byte, length int, dnsSearch []string) (bool, int) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var drop bool
 var err error
 if dnsSearch == nil || len(dnsSearch) == 0 {
  klog.V(1).Infof("DNS search list is not initialized and is empty.")
  return drop, length
 }
 dnsClients.mu.Lock()
 state, found := dnsClients.clients[dnsClientQuery{host, dnsQType}]
 dnsClients.mu.Unlock()
 if found {
  index := atomic.SwapInt32(&state.searchIndex, state.searchIndex+1)
  if rcode != 0 && index >= 0 && index < int32(len(dnsSearch)) {
   drop = true
   length, err = appendDNSSuffix(state.msg, buffer, length, dnsSearch[index])
   if err != nil {
    klog.Errorf("Append DNS suffix failed: %v", err)
   }
   _, err = svrConn.Write(buffer[0:length])
   if err != nil {
    if !logTimeout(err) {
     klog.Errorf("Write failed: %v", err)
    }
   }
  } else {
   length, err = recoverDNSQuestion(state.msg.Question[0].Name, msg, buffer, length)
   if err != nil {
    klog.Errorf("Recover DNS question failed: %v", err)
   }
   dnsClients.mu.Lock()
   delete(dnsClients.clients, dnsClientQuery{host, dnsQType})
   dnsClients.mu.Unlock()
  }
 }
 return drop, length
}
func processDNSQueryPacket(dnsClients *dnsClientCache, cliAddr net.Addr, buffer []byte, length int, dnsSearch []string) (int, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 msg := &dns.Msg{}
 if err := msg.Unpack(buffer[:length]); err != nil {
  klog.Warningf("Unable to unpack DNS packet. Error is: %v", err)
  return length, err
 }
 if msg.MsgHdr.Response == true {
  return length, fmt.Errorf("DNS packet should be a query message")
 }
 if len(msg.Question) != 1 {
  klog.V(1).Infof("Number of entries in the question section of the DNS packet is: %d", len(msg.Question))
  klog.V(1).Infof("DNS suffix appending does not support more than one question.")
  return length, nil
 }
 if len(msg.Answer) != 0 || len(msg.Ns) != 0 || len(msg.Extra) != 0 {
  klog.V(1).Infof("DNS packet contains more than question section.")
  return length, nil
 }
 dnsQType := msg.Question[0].Qtype
 dnsQClass := msg.Question[0].Qclass
 if packetRequiresDNSSuffix(dnsQType, dnsQClass) {
  host, _, err := net.SplitHostPort(cliAddr.String())
  if err != nil {
   klog.V(1).Infof("Failed to get host from client address: %v", err)
   host = cliAddr.String()
  }
  length = processUnpackedDNSQueryPacket(dnsClients, msg, host, dnsQType, buffer, length, dnsSearch)
 }
 return length, nil
}
func processDNSResponsePacket(svrConn net.Conn, dnsClients *dnsClientCache, cliAddr net.Addr, buffer []byte, length int, dnsSearch []string) (bool, int, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var drop bool
 msg := &dns.Msg{}
 if err := msg.Unpack(buffer[:length]); err != nil {
  klog.Warningf("Unable to unpack DNS packet. Error is: %v", err)
  return drop, length, err
 }
 if msg.MsgHdr.Response == false {
  return drop, length, fmt.Errorf("DNS packet should be a response message")
 }
 if len(msg.Question) != 1 {
  klog.V(1).Infof("Number of entries in the response section of the DNS packet is: %d", len(msg.Answer))
  return drop, length, nil
 }
 dnsQType := msg.Question[0].Qtype
 dnsQClass := msg.Question[0].Qclass
 if packetRequiresDNSSuffix(dnsQType, dnsQClass) {
  host, _, err := net.SplitHostPort(cliAddr.String())
  if err != nil {
   klog.V(1).Infof("Failed to get host from client address: %v", err)
   host = cliAddr.String()
  }
  drop, length = processUnpackedDNSResponsePacket(svrConn, dnsClients, msg, msg.MsgHdr.Rcode, host, dnsQType, buffer, length, dnsSearch)
 }
 return drop, length, nil
}
func (udp *udpProxySocket) ProxyLoop(service ServicePortPortalName, myInfo *serviceInfo, proxier *Proxier) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var buffer [4096]byte
 var dnsSearch []string
 if isDNSService(service.Port) {
  dnsSearch = []string{"", namespaceServiceDomain, serviceDomain, clusterDomain}
  execer := exec.New()
  ipconfigInterface := ipconfig.New(execer)
  suffixList, err := ipconfigInterface.GetDNSSuffixSearchList()
  if err == nil {
   for _, suffix := range suffixList {
    dnsSearch = append(dnsSearch, suffix)
   }
  }
 }
 for {
  if !myInfo.isAlive() {
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
  if isDNSService(service.Port) {
   n, err = processDNSQueryPacket(myInfo.dnsClients, cliAddr, buffer[:], n, dnsSearch)
   if err != nil {
    klog.Errorf("Process DNS query packet failed: %v", err)
   }
  }
  svrConn, err := udp.getBackendConn(myInfo.activeClients, myInfo.dnsClients, cliAddr, proxier, service, myInfo.timeout, dnsSearch)
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
  err = svrConn.SetDeadline(time.Now().Add(myInfo.timeout))
  if err != nil {
   klog.Errorf("SetDeadline failed: %v", err)
   continue
  }
 }
}
func (udp *udpProxySocket) getBackendConn(activeClients *clientCache, dnsClients *dnsClientCache, cliAddr net.Addr, proxier *Proxier, service ServicePortPortalName, timeout time.Duration, dnsSearch []string) (net.Conn, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 activeClients.mu.Lock()
 defer activeClients.mu.Unlock()
 svrConn, found := activeClients.clients[cliAddr.String()]
 if !found {
  klog.V(3).Infof("New UDP connection from %s", cliAddr)
  var err error
  svrConn, err = tryConnect(service, cliAddr, "udp", proxier)
  if err != nil {
   return nil, err
  }
  if err = svrConn.SetDeadline(time.Now().Add(timeout)); err != nil {
   klog.Errorf("SetDeadline failed: %v", err)
   return nil, err
  }
  activeClients.clients[cliAddr.String()] = svrConn
  go func(cliAddr net.Addr, svrConn net.Conn, activeClients *clientCache, dnsClients *dnsClientCache, service ServicePortPortalName, timeout time.Duration, dnsSearch []string) {
   defer runtime.HandleCrash()
   udp.proxyClient(cliAddr, svrConn, activeClients, dnsClients, service, timeout, dnsSearch)
  }(cliAddr, svrConn, activeClients, dnsClients, service, timeout, dnsSearch)
 }
 return svrConn, nil
}
func (udp *udpProxySocket) proxyClient(cliAddr net.Addr, svrConn net.Conn, activeClients *clientCache, dnsClients *dnsClientCache, service ServicePortPortalName, timeout time.Duration, dnsSearch []string) {
 _logClusterCodePath()
 defer _logClusterCodePath()
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
  drop := false
  if isDNSService(service.Port) {
   drop, n, err = processDNSResponsePacket(svrConn, dnsClients, cliAddr, buffer[:], n, dnsSearch)
   if err != nil {
    klog.Errorf("Process DNS response packet failed: %v", err)
   }
  }
  if !drop {
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
 }
 activeClients.mu.Lock()
 delete(activeClients.clients, cliAddr.String())
 activeClients.mu.Unlock()
}
