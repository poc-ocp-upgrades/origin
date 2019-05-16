package ldaputil

import (
	"crypto/tls"
	"fmt"
	goformat "fmt"
	"github.com/openshift/origin/pkg/oauthserver/ldaputil/ldapclient"
	"gopkg.in/ldap.v2"
	"k8s.io/client-go/util/cert"
	"net"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

func NewLDAPClientConfig(URL, bindDN, bindPassword, CA string, insecure bool) (ldapclient.Config, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	url, err := ParseURL(URL)
	if err != nil {
		return nil, fmt.Errorf("Error parsing URL: %v", err)
	}
	tlsConfig := &tls.Config{}
	if len(CA) > 0 {
		roots, err := cert.NewPool(CA)
		if err != nil {
			return nil, fmt.Errorf("error loading cert pool from ca file %s: %v", CA, err)
		}
		tlsConfig.RootCAs = roots
	}
	return &ldapClientConfig{scheme: url.Scheme, host: url.Host, bindDN: bindDN, bindPassword: bindPassword, insecure: insecure, tlsConfig: tlsConfig}, nil
}

type ldapClientConfig struct {
	scheme       Scheme
	host         string
	bindDN       string
	bindPassword string
	insecure     bool
	tlsConfig    *tls.Config
}

var _ ldapclient.Config = &ldapClientConfig{}

func (l *ldapClientConfig) Connect() (ldap.Client, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	tlsConfig := l.tlsConfig
	if tlsConfig != nil && !tlsConfig.InsecureSkipVerify && len(tlsConfig.ServerName) == 0 {
		c := tlsConfig.Clone()
		if host, _, err := net.SplitHostPort(l.host); err == nil {
			c.ServerName = host
		} else {
			c.ServerName = l.host
		}
		tlsConfig = c
	}
	switch l.scheme {
	case SchemeLDAP:
		con, err := ldap.Dial("tcp", l.host)
		if err != nil {
			return nil, err
		}
		if l.insecure {
			return con, nil
		}
		if err := con.StartTLS(tlsConfig); err != nil {
			con.Close()
			return nil, err
		}
		return con, nil
	case SchemeLDAPS:
		return ldap.DialTLS("tcp", l.host, tlsConfig)
	default:
		return nil, fmt.Errorf("unsupported scheme %q", l.scheme)
	}
}
func (l *ldapClientConfig) GetBindCredentials() (string, string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return l.bindDN, l.bindPassword
}
func (l *ldapClientConfig) Host() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return l.host
}
func (l *ldapClientConfig) String() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return fmt.Sprintf("{Scheme: %v Host: %v BindDN: %v len(BbindPassword): %v Insecure: %v}", l.scheme, l.host, l.bindDN, len(l.bindPassword), l.insecure)
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
