package ssh

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	goformat "fmt"
	"github.com/prometheus/client_golang/prometheus"
	"golang.org/x/crypto/ssh"
	"io"
	"io/ioutil"
	utilnet "k8s.io/apimachinery/pkg/util/net"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/klog"
	mathrand "math/rand"
	"net"
	"net/http"
	"net/url"
	"os"
	goos "os"
	godefaultruntime "runtime"
	"strings"
	"sync"
	"time"
	gotime "time"
)

var (
	tunnelOpenCounter     = prometheus.NewCounter(prometheus.CounterOpts{Name: "ssh_tunnel_open_count", Help: "Counter of ssh tunnel total open attempts"})
	tunnelOpenFailCounter = prometheus.NewCounter(prometheus.CounterOpts{Name: "ssh_tunnel_open_fail_count", Help: "Counter of ssh tunnel failed open attempts"})
)

func init() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	prometheus.MustRegister(tunnelOpenCounter)
	prometheus.MustRegister(tunnelOpenFailCounter)
}

type SSHTunnel struct {
	Config  *ssh.ClientConfig
	Host    string
	SSHPort string
	running bool
	sock    net.Listener
	client  *ssh.Client
}

func (s *SSHTunnel) copyBytes(out io.Writer, in io.Reader) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if _, err := io.Copy(out, in); err != nil {
		klog.Errorf("Error in SSH tunnel: %v", err)
	}
}
func NewSSHTunnel(user, keyfile, host string) (*SSHTunnel, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	signer, err := MakePrivateKeySignerFromFile(keyfile)
	if err != nil {
		return nil, err
	}
	return makeSSHTunnel(user, signer, host)
}
func NewSSHTunnelFromBytes(user string, privateKey []byte, host string) (*SSHTunnel, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	signer, err := MakePrivateKeySignerFromBytes(privateKey)
	if err != nil {
		return nil, err
	}
	return makeSSHTunnel(user, signer, host)
}
func makeSSHTunnel(user string, signer ssh.Signer, host string) (*SSHTunnel, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	config := ssh.ClientConfig{User: user, Auth: []ssh.AuthMethod{ssh.PublicKeys(signer)}, HostKeyCallback: ssh.InsecureIgnoreHostKey()}
	return &SSHTunnel{Config: &config, Host: host, SSHPort: "22"}, nil
}
func (s *SSHTunnel) Open() error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var err error
	s.client, err = realTimeoutDialer.Dial("tcp", net.JoinHostPort(s.Host, s.SSHPort), s.Config)
	tunnelOpenCounter.Inc()
	if err != nil {
		tunnelOpenFailCounter.Inc()
	}
	return err
}
func (s *SSHTunnel) Dial(ctx context.Context, network, address string) (net.Conn, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if s.client == nil {
		return nil, errors.New("tunnel is not opened.")
	}
	return s.client.Dial(network, address)
}
func (s *SSHTunnel) tunnel(conn net.Conn, remoteHost, remotePort string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if s.client == nil {
		return errors.New("tunnel is not opened.")
	}
	tunnel, err := s.client.Dial("tcp", net.JoinHostPort(remoteHost, remotePort))
	if err != nil {
		return err
	}
	go s.copyBytes(tunnel, conn)
	go s.copyBytes(conn, tunnel)
	return nil
}
func (s *SSHTunnel) Close() error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if s.client == nil {
		return errors.New("Cannot close tunnel. Tunnel was not opened.")
	}
	if err := s.client.Close(); err != nil {
		return err
	}
	return nil
}

type sshDialer interface {
	Dial(network, addr string, config *ssh.ClientConfig) (*ssh.Client, error)
}
type realSSHDialer struct{}

var _ sshDialer = &realSSHDialer{}

func (d *realSSHDialer) Dial(network, addr string, config *ssh.ClientConfig) (*ssh.Client, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	conn, err := net.DialTimeout(network, addr, config.Timeout)
	if err != nil {
		return nil, err
	}
	conn.SetReadDeadline(time.Now().Add(30 * time.Second))
	c, chans, reqs, err := ssh.NewClientConn(conn, addr, config)
	if err != nil {
		return nil, err
	}
	conn.SetReadDeadline(time.Time{})
	return ssh.NewClient(c, chans, reqs), nil
}

type timeoutDialer struct {
	dialer  sshDialer
	timeout time.Duration
}

const sshDialTimeout = 150 * time.Second

var realTimeoutDialer sshDialer = &timeoutDialer{&realSSHDialer{}, sshDialTimeout}

func (d *timeoutDialer) Dial(network, addr string, config *ssh.ClientConfig) (*ssh.Client, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	config.Timeout = d.timeout
	return d.dialer.Dial(network, addr, config)
}
func RunSSHCommand(cmd, user, host string, signer ssh.Signer) (string, string, int, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return runSSHCommand(realTimeoutDialer, cmd, user, host, signer, true)
}
func runSSHCommand(dialer sshDialer, cmd, user, host string, signer ssh.Signer, retry bool) (string, string, int, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if user == "" {
		user = os.Getenv("USER")
	}
	config := &ssh.ClientConfig{User: user, Auth: []ssh.AuthMethod{ssh.PublicKeys(signer)}, HostKeyCallback: ssh.InsecureIgnoreHostKey()}
	client, err := dialer.Dial("tcp", host, config)
	if err != nil && retry {
		err = wait.Poll(5*time.Second, 20*time.Second, func() (bool, error) {
			fmt.Printf("error dialing %s@%s: '%v', retrying\n", user, host, err)
			if client, err = dialer.Dial("tcp", host, config); err != nil {
				return false, err
			}
			return true, nil
		})
	}
	if err != nil {
		return "", "", 0, fmt.Errorf("error getting SSH client to %s@%s: '%v'", user, host, err)
	}
	session, err := client.NewSession()
	if err != nil {
		return "", "", 0, fmt.Errorf("error creating session to %s@%s: '%v'", user, host, err)
	}
	defer session.Close()
	code := 0
	var bout, berr bytes.Buffer
	session.Stdout, session.Stderr = &bout, &berr
	if err = session.Run(cmd); err != nil {
		if exiterr, ok := err.(*ssh.ExitError); ok {
			if code = exiterr.ExitStatus(); code != 0 {
				err = nil
			}
		} else {
			err = fmt.Errorf("failed running `%s` on %s@%s: '%v'", cmd, user, host, err)
		}
	}
	return bout.String(), berr.String(), code, err
}
func MakePrivateKeySignerFromFile(key string) (ssh.Signer, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	buffer, err := ioutil.ReadFile(key)
	if err != nil {
		return nil, fmt.Errorf("error reading SSH key %s: '%v'", key, err)
	}
	return MakePrivateKeySignerFromBytes(buffer)
}
func MakePrivateKeySignerFromBytes(buffer []byte) (ssh.Signer, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	signer, err := ssh.ParsePrivateKey(buffer)
	if err != nil {
		return nil, fmt.Errorf("error parsing SSH key: '%v'", err)
	}
	return signer, nil
}
func ParsePublicKeyFromFile(keyFile string) (*rsa.PublicKey, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	buffer, err := ioutil.ReadFile(keyFile)
	if err != nil {
		return nil, fmt.Errorf("error reading SSH key %s: '%v'", keyFile, err)
	}
	keyBlock, _ := pem.Decode(buffer)
	if keyBlock == nil {
		return nil, fmt.Errorf("error parsing SSH key %s: 'invalid PEM format'", keyFile)
	}
	key, err := x509.ParsePKIXPublicKey(keyBlock.Bytes)
	if err != nil {
		return nil, fmt.Errorf("error parsing SSH key %s: '%v'", keyFile, err)
	}
	rsaKey, ok := key.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("SSH key could not be parsed as rsa public key")
	}
	return rsaKey, nil
}

type tunnel interface {
	Open() error
	Close() error
	Dial(ctx context.Context, network, address string) (net.Conn, error)
}
type sshTunnelEntry struct {
	Address string
	Tunnel  tunnel
}
type sshTunnelCreator interface {
	NewSSHTunnel(user, keyFile, healthCheckURL string) (tunnel, error)
}
type realTunnelCreator struct{}

func (*realTunnelCreator) NewSSHTunnel(user, keyFile, healthCheckURL string) (tunnel, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return NewSSHTunnel(user, keyFile, healthCheckURL)
}

type SSHTunnelList struct {
	entries        []sshTunnelEntry
	adding         map[string]bool
	tunnelCreator  sshTunnelCreator
	tunnelsLock    sync.Mutex
	user           string
	keyfile        string
	healthCheckURL *url.URL
}

func NewSSHTunnelList(user, keyfile string, healthCheckURL *url.URL, stopChan chan struct{}) *SSHTunnelList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	l := &SSHTunnelList{adding: make(map[string]bool), tunnelCreator: &realTunnelCreator{}, user: user, keyfile: keyfile, healthCheckURL: healthCheckURL}
	healthCheckPoll := 1 * time.Minute
	go wait.Until(func() {
		l.tunnelsLock.Lock()
		defer l.tunnelsLock.Unlock()
		numTunnels := len(l.entries)
		for i, entry := range l.entries {
			delay := healthCheckPoll * time.Duration(i) / time.Duration(numTunnels)
			l.delayedHealthCheck(entry, delay)
		}
	}, healthCheckPoll, stopChan)
	return l
}
func (l *SSHTunnelList) delayedHealthCheck(e sshTunnelEntry, delay time.Duration) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	go func() {
		defer runtime.HandleCrash()
		time.Sleep(delay)
		if err := l.healthCheck(e); err != nil {
			klog.Errorf("Healthcheck failed for tunnel to %q: %v", e.Address, err)
			klog.Infof("Attempting once to re-establish tunnel to %q", e.Address)
			l.removeAndReAdd(e)
		}
	}()
}
func (l *SSHTunnelList) healthCheck(e sshTunnelEntry) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	transport := utilnet.SetTransportDefaults(&http.Transport{DialContext: e.Tunnel.Dial, TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, DisableKeepAlives: true})
	client := &http.Client{Transport: transport}
	resp, err := client.Get(l.healthCheckURL.String())
	if err != nil {
		return err
	}
	resp.Body.Close()
	return nil
}
func (l *SSHTunnelList) removeAndReAdd(e sshTunnelEntry) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	l.tunnelsLock.Lock()
	for i, entry := range l.entries {
		if entry.Tunnel == e.Tunnel {
			l.entries = append(l.entries[:i], l.entries[i+1:]...)
			l.adding[e.Address] = true
			break
		}
	}
	l.tunnelsLock.Unlock()
	if err := e.Tunnel.Close(); err != nil {
		klog.Infof("Failed to close removed tunnel: %v", err)
	}
	go l.createAndAddTunnel(e.Address)
}
func (l *SSHTunnelList) Dial(ctx context.Context, net, addr string) (net.Conn, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	start := time.Now()
	id := mathrand.Int63()
	klog.Infof("[%x: %v] Dialing...", id, addr)
	defer func() {
		klog.Infof("[%x: %v] Dialed in %v.", id, addr, time.Since(start))
	}()
	tunnel, err := l.pickTunnel(strings.Split(addr, ":")[0])
	if err != nil {
		return nil, err
	}
	return tunnel.Dial(ctx, net, addr)
}
func (l *SSHTunnelList) pickTunnel(addr string) (tunnel, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	l.tunnelsLock.Lock()
	defer l.tunnelsLock.Unlock()
	if len(l.entries) == 0 {
		return nil, fmt.Errorf("No SSH tunnels currently open. Were the targets able to accept an ssh-key for user %q?", l.user)
	}
	for _, entry := range l.entries {
		if entry.Address == addr {
			return entry.Tunnel, nil
		}
	}
	klog.Warningf("SSH tunnel not found for address %q, picking random node", addr)
	n := mathrand.Intn(len(l.entries))
	return l.entries[n].Tunnel, nil
}
func (l *SSHTunnelList) Update(addrs []string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	haveAddrsMap := make(map[string]bool)
	wantAddrsMap := make(map[string]bool)
	func() {
		l.tunnelsLock.Lock()
		defer l.tunnelsLock.Unlock()
		for i := range l.entries {
			haveAddrsMap[l.entries[i].Address] = true
		}
		for i := range addrs {
			if _, ok := haveAddrsMap[addrs[i]]; !ok {
				if _, ok := l.adding[addrs[i]]; !ok {
					l.adding[addrs[i]] = true
					addr := addrs[i]
					go func() {
						defer runtime.HandleCrash()
						l.createAndAddTunnel(addr)
					}()
				}
			}
			wantAddrsMap[addrs[i]] = true
		}
		var newEntries []sshTunnelEntry
		for i := range l.entries {
			if _, ok := wantAddrsMap[l.entries[i].Address]; !ok {
				tunnelEntry := l.entries[i]
				klog.Infof("Removing tunnel to deleted node at %q", tunnelEntry.Address)
				go func() {
					defer runtime.HandleCrash()
					if err := tunnelEntry.Tunnel.Close(); err != nil {
						klog.Errorf("Failed to close tunnel to %q: %v", tunnelEntry.Address, err)
					}
				}()
			} else {
				newEntries = append(newEntries, l.entries[i])
			}
		}
		l.entries = newEntries
	}()
}
func (l *SSHTunnelList) createAndAddTunnel(addr string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.Infof("Trying to add tunnel to %q", addr)
	tunnel, err := l.tunnelCreator.NewSSHTunnel(l.user, l.keyfile, addr)
	if err != nil {
		klog.Errorf("Failed to create tunnel for %q: %v", addr, err)
		return
	}
	if err := tunnel.Open(); err != nil {
		klog.Errorf("Failed to open tunnel to %q: %v", addr, err)
		l.tunnelsLock.Lock()
		delete(l.adding, addr)
		l.tunnelsLock.Unlock()
		return
	}
	l.tunnelsLock.Lock()
	l.entries = append(l.entries, sshTunnelEntry{addr, tunnel})
	delete(l.adding, addr)
	l.tunnelsLock.Unlock()
	klog.Infof("Successfully added tunnel for %q", addr)
}
func EncodePrivateKey(private *rsa.PrivateKey) []byte {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return pem.EncodeToMemory(&pem.Block{Bytes: x509.MarshalPKCS1PrivateKey(private), Type: "RSA PRIVATE KEY"})
}
func EncodePublicKey(public *rsa.PublicKey) ([]byte, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	publicBytes, err := x509.MarshalPKIXPublicKey(public)
	if err != nil {
		return nil, err
	}
	return pem.EncodeToMemory(&pem.Block{Bytes: publicBytes, Type: "PUBLIC KEY"}), nil
}
func EncodeSSHKey(public *rsa.PublicKey) ([]byte, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	publicKey, err := ssh.NewPublicKey(public)
	if err != nil {
		return nil, err
	}
	return ssh.MarshalAuthorizedKey(publicKey), nil
}
func GenerateKey(bits int) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	private, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}
	return private, &private.PublicKey, nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
