package healthcheck

import (
	"fmt"
	goformat "fmt"
	"github.com/renstrom/dedent"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/clock"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/tools/record"
	"k8s.io/klog"
	api "k8s.io/kubernetes/pkg/apis/core"
	"net"
	"net/http"
	goos "os"
	godefaultruntime "runtime"
	"strings"
	"sync"
	"sync/atomic"
	"time"
	gotime "time"
)

var nodeHealthzRetryInterval = 60 * time.Second

type Server interface {
	SyncServices(newServices map[types.NamespacedName]uint16) error
	SyncEndpoints(newEndpoints map[types.NamespacedName]int) error
}
type Listener interface {
	Listen(addr string) (net.Listener, error)
}
type HTTPServerFactory interface {
	New(addr string, handler http.Handler) HTTPServer
}
type HTTPServer interface {
	Serve(listener net.Listener) error
}

func NewServer(hostname string, recorder record.EventRecorder, listener Listener, httpServerFactory HTTPServerFactory) Server {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if listener == nil {
		listener = stdNetListener{}
	}
	if httpServerFactory == nil {
		httpServerFactory = stdHTTPServerFactory{}
	}
	return &server{hostname: hostname, recorder: recorder, listener: listener, httpFactory: httpServerFactory, services: map[types.NamespacedName]*hcInstance{}}
}

type stdNetListener struct{}

func (stdNetListener) Listen(addr string) (net.Listener, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return net.Listen("tcp", addr)
}

var _ Listener = stdNetListener{}

type stdHTTPServerFactory struct{}

func (stdHTTPServerFactory) New(addr string, handler http.Handler) HTTPServer {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &http.Server{Addr: addr, Handler: handler}
}

var _ HTTPServerFactory = stdHTTPServerFactory{}

type server struct {
	hostname    string
	recorder    record.EventRecorder
	listener    Listener
	httpFactory HTTPServerFactory
	lock        sync.Mutex
	services    map[types.NamespacedName]*hcInstance
}

func (hcs *server) SyncServices(newServices map[types.NamespacedName]uint16) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	hcs.lock.Lock()
	defer hcs.lock.Unlock()
	for nsn, svc := range hcs.services {
		if port, found := newServices[nsn]; !found || port != svc.port {
			klog.V(2).Infof("Closing healthcheck %q on port %d", nsn.String(), svc.port)
			if err := svc.listener.Close(); err != nil {
				klog.Errorf("Close(%v): %v", svc.listener.Addr(), err)
			}
			delete(hcs.services, nsn)
		}
	}
	for nsn, port := range newServices {
		if hcs.services[nsn] != nil {
			klog.V(3).Infof("Existing healthcheck %q on port %d", nsn.String(), port)
			continue
		}
		klog.V(2).Infof("Opening healthcheck %q on port %d", nsn.String(), port)
		svc := &hcInstance{port: port}
		addr := fmt.Sprintf(":%d", port)
		svc.server = hcs.httpFactory.New(addr, hcHandler{name: nsn, hcs: hcs})
		var err error
		svc.listener, err = hcs.listener.Listen(addr)
		if err != nil {
			msg := fmt.Sprintf("node %s failed to start healthcheck %q on port %d: %v", hcs.hostname, nsn.String(), port, err)
			if hcs.recorder != nil {
				hcs.recorder.Eventf(&v1.ObjectReference{Kind: "Service", Namespace: nsn.Namespace, Name: nsn.Name, UID: types.UID(nsn.String())}, api.EventTypeWarning, "FailedToStartServiceHealthcheck", msg)
			}
			klog.Error(msg)
			continue
		}
		hcs.services[nsn] = svc
		go func(nsn types.NamespacedName, svc *hcInstance) {
			klog.V(3).Infof("Starting goroutine for healthcheck %q on port %d", nsn.String(), svc.port)
			if err := svc.server.Serve(svc.listener); err != nil {
				klog.V(3).Infof("Healthcheck %q closed: %v", nsn.String(), err)
				return
			}
			klog.V(3).Infof("Healthcheck %q closed", nsn.String())
		}(nsn, svc)
	}
	return nil
}

type hcInstance struct {
	port      uint16
	listener  net.Listener
	server    HTTPServer
	endpoints int
}
type hcHandler struct {
	name types.NamespacedName
	hcs  *server
}

var _ http.Handler = hcHandler{}

func (h hcHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	h.hcs.lock.Lock()
	svc, ok := h.hcs.services[h.name]
	if !ok || svc == nil {
		h.hcs.lock.Unlock()
		klog.Errorf("Received request for closed healthcheck %q", h.name.String())
		return
	}
	count := svc.endpoints
	h.hcs.lock.Unlock()
	resp.Header().Set("Content-Type", "application/json")
	if count == 0 {
		resp.WriteHeader(http.StatusServiceUnavailable)
	} else {
		resp.WriteHeader(http.StatusOK)
	}
	fmt.Fprintf(resp, strings.Trim(dedent.Dedent(fmt.Sprintf(`
		{
			"service": {
				"namespace": %q,
				"name": %q
			},
			"localEndpoints": %d
		}
		`, h.name.Namespace, h.name.Name, count)), "\n"))
}
func (hcs *server) SyncEndpoints(newEndpoints map[types.NamespacedName]int) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	hcs.lock.Lock()
	defer hcs.lock.Unlock()
	for nsn, count := range newEndpoints {
		if hcs.services[nsn] == nil {
			klog.V(3).Infof("Not saving endpoints for unknown healthcheck %q", nsn.String())
			continue
		}
		klog.V(3).Infof("Reporting %d endpoints for healthcheck %q", count, nsn.String())
		hcs.services[nsn].endpoints = count
	}
	for nsn, hci := range hcs.services {
		if _, found := newEndpoints[nsn]; !found {
			hci.endpoints = 0
		}
	}
	return nil
}

type HealthzUpdater interface{ UpdateTimestamp() }
type HealthzServer struct {
	listener      Listener
	httpFactory   HTTPServerFactory
	clock         clock.Clock
	addr          string
	port          int32
	healthTimeout time.Duration
	recorder      record.EventRecorder
	nodeRef       *v1.ObjectReference
	lastUpdated   atomic.Value
}

func NewDefaultHealthzServer(addr string, healthTimeout time.Duration, recorder record.EventRecorder, nodeRef *v1.ObjectReference) *HealthzServer {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return newHealthzServer(nil, nil, nil, addr, healthTimeout, recorder, nodeRef)
}
func newHealthzServer(listener Listener, httpServerFactory HTTPServerFactory, c clock.Clock, addr string, healthTimeout time.Duration, recorder record.EventRecorder, nodeRef *v1.ObjectReference) *HealthzServer {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if listener == nil {
		listener = stdNetListener{}
	}
	if httpServerFactory == nil {
		httpServerFactory = stdHTTPServerFactory{}
	}
	if c == nil {
		c = clock.RealClock{}
	}
	return &HealthzServer{listener: listener, httpFactory: httpServerFactory, clock: c, addr: addr, healthTimeout: healthTimeout, recorder: recorder, nodeRef: nodeRef}
}
func (hs *HealthzServer) UpdateTimestamp() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	hs.lastUpdated.Store(hs.clock.Now())
}
func (hs *HealthzServer) Run() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	serveMux := http.NewServeMux()
	serveMux.Handle("/healthz", healthzHandler{hs: hs})
	server := hs.httpFactory.New(hs.addr, serveMux)
	go wait.Until(func() {
		klog.V(3).Infof("Starting goroutine for healthz on %s", hs.addr)
		listener, err := hs.listener.Listen(hs.addr)
		if err != nil {
			msg := fmt.Sprintf("Failed to start node healthz on %s: %v", hs.addr, err)
			if hs.recorder != nil {
				hs.recorder.Eventf(hs.nodeRef, api.EventTypeWarning, "FailedToStartNodeHealthcheck", msg)
			}
			klog.Error(msg)
			return
		}
		if err := server.Serve(listener); err != nil {
			klog.Errorf("Healthz closed with error: %v", err)
			return
		}
		klog.Error("Unexpected healthz closed.")
	}, nodeHealthzRetryInterval, wait.NeverStop)
}

type healthzHandler struct{ hs *HealthzServer }

func (h healthzHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	lastUpdated := time.Time{}
	if val := h.hs.lastUpdated.Load(); val != nil {
		lastUpdated = val.(time.Time)
	}
	currentTime := h.hs.clock.Now()
	resp.Header().Set("Content-Type", "application/json")
	if !lastUpdated.IsZero() && currentTime.After(lastUpdated.Add(h.hs.healthTimeout)) {
		resp.WriteHeader(http.StatusServiceUnavailable)
	} else {
		resp.WriteHeader(http.StatusOK)
	}
	fmt.Fprintf(resp, fmt.Sprintf(`{"lastUpdated": %q,"currentTime": %q}`, lastUpdated, currentTime))
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
