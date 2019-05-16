package config

import (
	"crypto/x509"
	"fmt"
	goformat "fmt"
	"k8s.io/apimachinery/third_party/forked/golang/netutil"
	x509request "k8s.io/apiserver/pkg/authentication/request/x509"
	"net"
	"net/url"
	goos "os"
	godefaultruntime "runtime"
	"strconv"
	"strings"
	gotime "time"
)

func GetClusterNicknameFromURL(apiServerLocation string) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	u, err := url.Parse(apiServerLocation)
	if err != nil {
		return "", err
	}
	hostPort := netutil.CanonicalAddr(u)
	return strings.Replace(hostPort, ".", "-", -1), nil
}
func GetUserNicknameFromCert(clusterNick string, chain ...*x509.Certificate) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	authResponse, _, err := x509request.CommonNameUserConversion(chain)
	if err != nil {
		return "", err
	}
	return authResponse.User.GetName() + "/" + clusterNick, nil
}
func GetContextNickname(namespace, clusterNick, userNick string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	tokens := strings.SplitN(userNick, "/", 2)
	return namespace + "/" + clusterNick + "/" + tokens[0]
}

var validURLSchemes = []string{"https://", "http://", "tcp://"}

func NormalizeServerURL(s string) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if !hasScheme(s) {
		s = validURLSchemes[0] + s
	}
	addr, err := url.Parse(s)
	if err != nil {
		return "", fmt.Errorf("Not a valid URL: %v.", err)
	}
	if strings.Contains(addr.Host, ":") {
		_, port, err := net.SplitHostPort(addr.Host)
		if err != nil {
			return "", fmt.Errorf("Not a valid host:port: %v.", err)
		}
		_, err = strconv.ParseUint(port, 10, 16)
		if err != nil {
			return "", fmt.Errorf("Not a valid port: %v. Port numbers must be between 0 and 65535.", port)
		}
	} else {
		port := 0
		switch addr.Scheme {
		case "http":
			port = 80
		case "https":
			port = 443
		default:
			return "", fmt.Errorf("No port specified.")
		}
		addr.Host = net.JoinHostPort(addr.Host, strconv.FormatInt(int64(port), 10))
	}
	if addr.Path == "/" {
		addr.Path = ""
	}
	return addr.String(), nil
}
func hasScheme(s string) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for _, p := range validURLSchemes {
		if strings.HasPrefix(s, p) {
			return true
		}
	}
	return false
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
