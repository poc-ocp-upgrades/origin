package util

import (
	"crypto/tls"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog"
	"net"
	"net/http"
	"strings"
	"time"
)

func TryListen(network, hostPort string) (bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	l, err := net.Listen(network, hostPort)
	if err != nil {
		klog.V(5).Infof("Failure while checking listen on %s: %v", hostPort, err)
		return false, err
	}
	defer l.Close()
	return true, nil
}

type tcpKeepAliveListener struct{ *net.TCPListener }

func (ln tcpKeepAliveListener) Accept() (c net.Conn, err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	tc, err := ln.AcceptTCP()
	if err != nil {
		return
	}
	tc.SetKeepAlive(true)
	tc.SetKeepAlivePeriod(3 * time.Minute)
	return tc, nil
}
func ListenAndServeTLS(srv *http.Server, network string, certFile, keyFile string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	addr := srv.Addr
	if addr == "" {
		addr = ":https"
	}
	config := &tls.Config{}
	if srv.TLSConfig != nil {
		config = srv.TLSConfig
	}
	if config.NextProtos == nil {
		config.NextProtos = []string{"http/1.1"}
	}
	var err error
	config.Certificates = make([]tls.Certificate, 1)
	config.Certificates[0], err = tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return err
	}
	ln, err := net.Listen(network, addr)
	if err != nil {
		return err
	}
	tlsListener := tls.NewListener(tcpKeepAliveListener{ln.(*net.TCPListener)}, config)
	return srv.Serve(tlsListener)
}
func WaitForSuccessfulDial(https bool, network, address string, timeout, interval time.Duration, retries int) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var (
		conn net.Conn
		err  error
	)
	for i := 0; i <= retries; i++ {
		dialer := net.Dialer{Timeout: timeout}
		if https {
			conn, err = tls.DialWithDialer(&dialer, network, address, &tls.Config{InsecureSkipVerify: true})
		} else {
			conn, err = dialer.Dial(network, address)
		}
		if err != nil {
			klog.V(5).Infof("Got error %#v, trying again: %#v\n", err, address)
			time.Sleep(interval)
			continue
		}
		conn.Close()
		return nil
	}
	return err
}
func GetCertificateFunc(certs map[string]*tls.Certificate) func(*tls.ClientHelloInfo) (*tls.Certificate, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(certs) == 0 {
		return nil
	}
	return func(clientHello *tls.ClientHelloInfo) (*tls.Certificate, error) {
		if clientHello == nil {
			return nil, nil
		}
		name := clientHello.ServerName
		name = strings.ToLower(name)
		name = strings.TrimRight(name, ".")
		for _, candidate := range HostnameMatchSpecCandidates(name) {
			if cert, ok := certs[candidate]; ok {
				return cert, nil
			}
		}
		return nil, nil
	}
}
func HostnameMatchSpecCandidates(hostname string) []string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(hostname) == 0 {
		return nil
	}
	candidates := []string{hostname}
	labels := strings.Split(hostname, ".")
	for i := range labels {
		labels[i] = "*"
		candidates = append(candidates, strings.Join(labels, "."))
	}
	return candidates
}
func HostnameMatches(hostname string, matchSpec string) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return sets.NewString(HostnameMatchSpecCandidates(hostname)...).Has(matchSpec)
}
