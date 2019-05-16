package httpproxy

import (
	"bufio"
	"crypto/tls"
	"fmt"
	goformat "fmt"
	"io"
	"io/ioutil"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/third_party/forked/golang/netutil"
	restclient "k8s.io/client-go/rest"
	"k8s.io/klog"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	goos "os"
	godefaultruntime "runtime"
	"strings"
	"time"
	gotime "time"
)

type UpgradeAwareSingleHostReverseProxy struct {
	clientConfig *restclient.Config
	backendAddr  *url.URL
	transport    http.RoundTripper
	reverseProxy *httputil.ReverseProxy
}

func NewUpgradeAwareSingleHostReverseProxy(clientConfig *restclient.Config, backendAddr *url.URL) (*UpgradeAwareSingleHostReverseProxy, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	transport, err := restclient.TransportFor(clientConfig)
	if err != nil {
		return nil, err
	}
	reverseProxy := httputil.NewSingleHostReverseProxy(backendAddr)
	reverseProxy.FlushInterval = 200 * time.Millisecond
	p := &UpgradeAwareSingleHostReverseProxy{clientConfig: clientConfig, backendAddr: backendAddr, transport: transport, reverseProxy: reverseProxy}
	p.reverseProxy.Transport = p
	return p, nil
}
func (p *UpgradeAwareSingleHostReverseProxy) RoundTrip(req *http.Request) (*http.Response, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	resp, err := p.transport.RoundTrip(req)
	if err != nil {
		return resp, err
	}
	removeCORSHeaders(resp)
	removeChallengeHeaders(resp)
	if resp.StatusCode == http.StatusUnauthorized {
		utilruntime.HandleError(fmt.Errorf("got unauthorized error from backend for: %s %s", req.Method, req.URL))
		resp = &http.Response{StatusCode: http.StatusInternalServerError, Status: http.StatusText(http.StatusInternalServerError), Body: ioutil.NopCloser(strings.NewReader("Internal Server Error")), ContentLength: -1}
	}
	return resp, err
}
func singleJoiningSlash(a, b string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	aslash := strings.HasSuffix(a, "/")
	bslash := strings.HasPrefix(b, "/")
	switch {
	case aslash && bslash:
		return a + b[1:]
	case !aslash && !bslash:
		return a + "/" + b
	}
	return a + b
}
func (p *UpgradeAwareSingleHostReverseProxy) newProxyRequest(req *http.Request) (*http.Request, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	backendURL := *p.backendAddr
	backendURL.Path = singleJoiningSlash(backendURL.Path, req.URL.Path)
	backendURL.RawQuery = req.URL.RawQuery
	newReq, err := http.NewRequest(req.Method, backendURL.String(), req.Body)
	if err != nil {
		return nil, err
	}
	newReq.Header = req.Header
	removeAuthHeaders(newReq)
	return newReq, nil
}
func (p *UpgradeAwareSingleHostReverseProxy) isUpgradeRequest(req *http.Request) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for _, h := range req.Header[http.CanonicalHeaderKey("Connection")] {
		if strings.Contains(strings.ToLower(h), "upgrade") {
			return true
		}
	}
	return false
}
func (p *UpgradeAwareSingleHostReverseProxy) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	newReq, err := p.newProxyRequest(req)
	if err != nil {
		klog.Errorf("Error creating backend request: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !p.isUpgradeRequest(req) {
		p.reverseProxy.ServeHTTP(w, newReq)
		return
	}
	p.serveUpgrade(w, newReq)
}
func (p *UpgradeAwareSingleHostReverseProxy) dialBackend(req *http.Request) (net.Conn, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	dialAddr := netutil.CanonicalAddr(req.URL)
	switch p.backendAddr.Scheme {
	case "http":
		return net.Dial("tcp", dialAddr)
	case "https":
		tlsConfig, err := restclient.TLSConfigFor(p.clientConfig)
		if err != nil {
			return nil, err
		}
		tlsConn, err := tls.Dial("tcp", dialAddr, tlsConfig)
		if err != nil {
			return nil, err
		}
		hostToVerify, _, err := net.SplitHostPort(dialAddr)
		if err != nil {
			return nil, err
		}
		err = tlsConn.VerifyHostname(hostToVerify)
		return tlsConn, err
	default:
		return nil, fmt.Errorf("unknown scheme: %s", p.backendAddr.Scheme)
	}
}
func (p *UpgradeAwareSingleHostReverseProxy) serveUpgrade(w http.ResponseWriter, req *http.Request) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	backendConn, err := p.dialBackend(req)
	if err != nil {
		klog.Errorf("Error connecting to backend: %s", err)
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}
	defer backendConn.Close()
	addAuthHeaders(req, p.clientConfig)
	err = req.Write(backendConn)
	if err != nil {
		klog.Errorf("Error writing request to backend: %s", err)
		return
	}
	resp, err := http.ReadResponse(bufio.NewReader(backendConn), req)
	if err != nil {
		klog.Errorf("Error reading response from backend: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}
	if resp.StatusCode == http.StatusUnauthorized {
		klog.Errorf("Got unauthorized error from backend for: %s %s", req.Method, req.URL)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}
	requestHijackedConn, _, err := w.(http.Hijacker).Hijack()
	if err != nil {
		klog.Errorf("Error hijacking request connection: %s", err)
		return
	}
	defer requestHijackedConn.Close()
	removeCORSHeaders(resp)
	removeChallengeHeaders(resp)
	err = resp.Write(requestHijackedConn)
	if err != nil {
		klog.Errorf("Error writing backend response to client: %s", err)
		return
	}
	done := make(chan struct{}, 2)
	go func() {
		_, err := io.Copy(backendConn, requestHijackedConn)
		if err != nil && !strings.Contains(err.Error(), "use of closed network connection") {
			utilruntime.HandleError(fmt.Errorf("error proxying data from client to backend: %v", err))
		}
		done <- struct{}{}
	}()
	go func() {
		_, err := io.Copy(requestHijackedConn, backendConn)
		if err != nil && !strings.Contains(err.Error(), "use of closed network connection") {
			utilruntime.HandleError(fmt.Errorf("error proxying data from backend to client: %v", err))
		}
		done <- struct{}{}
	}()
	<-done
}
func removeAuthHeaders(req *http.Request) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	req.Header.Del("Authorization")
}
func removeChallengeHeaders(resp *http.Response) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	resp.Header.Del("WWW-Authenticate")
}
func removeCORSHeaders(resp *http.Response) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	resp.Header.Del("Access-Control-Allow-Credentials")
	resp.Header.Del("Access-Control-Allow-Headers")
	resp.Header.Del("Access-Control-Allow-Methods")
	resp.Header.Del("Access-Control-Allow-Origin")
}
func addAuthHeaders(req *http.Request, clientConfig *restclient.Config) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if clientConfig.BearerToken != "" {
		req.Header.Set("Authorization", "Bearer "+clientConfig.BearerToken)
	} else if clientConfig.Username != "" || clientConfig.Password != "" {
		req.SetBasicAuth(clientConfig.Username, clientConfig.Password)
	}
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
