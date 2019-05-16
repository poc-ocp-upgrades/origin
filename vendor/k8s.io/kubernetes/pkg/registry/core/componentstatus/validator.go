package componentstatus

import (
	"crypto/tls"
	utilnet "k8s.io/apimachinery/pkg/util/net"
	"k8s.io/kubernetes/pkg/probe"
	httpprober "k8s.io/kubernetes/pkg/probe/http"
	"net/http"
	"sync"
	"time"
)

const (
	probeTimeOut = 20 * time.Second
)

type httpGet interface {
	Get(url string) (*http.Response, error)
}
type ValidatorFn func([]byte) error
type Server struct {
	Addr        string
	Port        int
	Path        string
	EnableHTTPS bool
	TLSConfig   *tls.Config
	Validate    ValidatorFn
	Prober      httpprober.Prober
	Once        sync.Once
}
type ServerStatus struct {
	Component  string       `json:"component,omitempty"`
	Health     string       `json:"health,omitempty"`
	HealthCode probe.Result `json:"healthCode,omitempty"`
	Msg        string       `json:"msg,omitempty"`
	Err        string       `json:"err,omitempty"`
}

func (server *Server) DoServerCheck() (probe.Result, string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	server.Once.Do(func() {
		if server.Prober != nil {
			return
		}
		server.Prober = httpprober.NewWithTLSConfig(server.TLSConfig)
	})
	scheme := "http"
	if server.EnableHTTPS {
		scheme = "https"
	}
	url := utilnet.FormatURL(scheme, server.Addr, server.Port, server.Path)
	result, data, err := server.Prober.Probe(url, nil, probeTimeOut)
	if err != nil {
		return probe.Unknown, "", err
	}
	if result == probe.Failure {
		return probe.Failure, string(data), err
	}
	if server.Validate != nil {
		if err := server.Validate([]byte(data)); err != nil {
			return probe.Failure, string(data), err
		}
	}
	return result, string(data), nil
}
