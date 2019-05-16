package flagtypes

import (
	"fmt"
	goformat "fmt"
	"net"
	"net/url"
	goos "os"
	godefaultruntime "runtime"
	"strconv"
	"strings"
	gotime "time"
)

var urlPrefixes = []string{"http://", "https://", "tcp://", "unix://"}

type Addr struct {
	DefaultScheme string
	DefaultPort   int
	AllowPrefix   bool
	Provided      bool
	Value         string
	URL           *url.URL
	Host          string
	IPv6Host      bool
	Port          int
}

func (a Addr) Default() Addr {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if err := a.Set(a.Value); err != nil {
		panic(err)
	}
	a.Provided = false
	return a
}
func (a *Addr) String() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if a.URL == nil {
		return a.Value
	}
	return a.URL.String()
}
func (a *Addr) HostPort(defaultPort int) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	port := a.Port
	if port == 0 {
		port = defaultPort
	}
	return net.JoinHostPort(a.Host, strconv.Itoa(port))
}
func (a *Addr) Set(value string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	scheme := a.DefaultScheme
	if len(scheme) == 0 {
		scheme = "tcp"
	}
	addr := &url.URL{Scheme: scheme}
	switch {
	case a.isURL(value):
		parsed, err := url.Parse(value)
		if err != nil {
			return fmt.Errorf("not a valid URL: %v", err)
		}
		if !a.AllowPrefix {
			parsed.Path = ""
		}
		parsed.RawQuery = ""
		parsed.Fragment = ""
		if parsed.Scheme != "unix" && strings.Contains(parsed.Host, ":") {
			host, port, err := net.SplitHostPort(parsed.Host)
			if err != nil {
				return fmt.Errorf("not a valid host:port: %v", err)
			}
			portNum, err := strconv.ParseUint(port, 10, 64)
			if err != nil {
				return fmt.Errorf("not a valid port: %v", err)
			}
			a.Host = host
			a.Port = int(portNum)
		} else {
			port := 0
			switch parsed.Scheme {
			case "http":
				port = 80
			case "https":
				port = 443
			case "unix":
				port = 0
			default:
				return fmt.Errorf("no port specified")
			}
			a.Host = parsed.Host
			a.Port = port
		}
		addr = parsed
	case isIPv6Host(value):
		a.Host = value
		a.Port = a.DefaultPort
	case strings.Contains(value, ":"):
		host, port, err := net.SplitHostPort(value)
		if err != nil {
			return fmt.Errorf("not a valid host:port: %v", err)
		}
		portNum, err := strconv.ParseUint(port, 10, 64)
		if err != nil {
			return fmt.Errorf("not a valid port: %v", err)
		}
		a.Host = host
		a.Port = int(portNum)
	default:
		port := a.DefaultPort
		if port == 0 {
			switch a.DefaultScheme {
			case "http":
				port = 80
			case "https":
				port = 443
			default:
				return fmt.Errorf("no port specified")
			}
		}
		a.Host = value
		a.Port = port
	}
	if a.Port > 0 {
		addr.Host = net.JoinHostPort(a.Host, strconv.FormatInt(int64(a.Port), 10))
	} else {
		addr.Host = a.Host
	}
	if value != a.Value {
		a.Provided = true
	}
	a.URL = addr
	a.IPv6Host = isIPv6Host(a.Host)
	a.Value = value
	return nil
}
func (a *Addr) Type() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return "string"
}
func (a *Addr) isURL(value string) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	prefixes := urlPrefixes
	if a.DefaultScheme != "" {
		prefixes = append(prefixes, fmt.Sprintf("%s://", a.DefaultScheme))
	}
	for _, p := range prefixes {
		if strings.HasPrefix(value, p) {
			return true
		}
	}
	return false
}
func isIPv6Host(value string) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if strings.HasPrefix(value, "[") {
		return false
	}
	return strings.Contains(value, "%") || strings.Count(value, ":") > 1
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
