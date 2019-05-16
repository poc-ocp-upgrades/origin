package httprequest

import (
	"bitbucket.org/ww/goautoneg"
	goformat "fmt"
	"net"
	"net/http"
	goos "os"
	godefaultruntime "runtime"
	"strings"
	gotime "time"
)

func PrefersHTML(req *http.Request) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	accepts := goautoneg.ParseAccept(req.Header.Get("Accept"))
	acceptsHTML := false
	acceptsJSON := false
	for _, accept := range accepts {
		if accept.Type == "text" && accept.SubType == "html" {
			acceptsHTML = true
		} else if accept.Type == "application" && accept.SubType == "json" {
			acceptsJSON = true
		}
	}
	if acceptsHTML {
		return true
	}
	if acceptsJSON {
		return false
	}
	if strings.HasPrefix(req.UserAgent(), "Mozilla") {
		return true
	}
	return false
}
func SchemeHost(req *http.Request) (string, string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	forwarded := func(attr string) string {
		value := req.Header.Get("X-Forwarded-" + attr)
		value = strings.SplitN(value, ",", 2)[0]
		return strings.TrimSpace(value)
	}
	hasExplicitHost := func(h string) bool {
		_, _, err := net.SplitHostPort(h)
		return err == nil
	}
	forwardedHost := forwarded("Host")
	host := ""
	hostHadExplicitPort := false
	switch {
	case len(forwardedHost) > 0:
		host = forwardedHost
		hostHadExplicitPort = hasExplicitHost(host)
		if forwardedPort := forwarded("Port"); len(forwardedPort) > 0 {
			if h, _, err := net.SplitHostPort(forwardedHost); err == nil {
				host = net.JoinHostPort(h, forwardedPort)
			} else {
				host = net.JoinHostPort(forwardedHost, forwardedPort)
			}
		}
	case len(req.Host) > 0:
		host = req.Host
		hostHadExplicitPort = hasExplicitHost(host)
	case len(req.URL.Host) > 0:
		host = req.URL.Host
		hostHadExplicitPort = hasExplicitHost(host)
	}
	port := ""
	if _, p, err := net.SplitHostPort(host); err == nil {
		port = p
	}
	forwardedProto := forwarded("Proto")
	scheme := ""
	switch {
	case len(forwardedProto) > 0:
		scheme = forwardedProto
	case req.TLS != nil:
		scheme = "https"
	case len(req.URL.Scheme) > 0:
		scheme = req.URL.Scheme
	case port == "443":
		scheme = "https"
	default:
		scheme = "http"
	}
	if !hostHadExplicitPort {
		if (scheme == "https" && port == "443") || (scheme == "http" && port == "80") {
			if hostWithoutPort, _, err := net.SplitHostPort(host); err == nil {
				host = hostWithoutPort
			}
		}
	}
	return scheme, host
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
