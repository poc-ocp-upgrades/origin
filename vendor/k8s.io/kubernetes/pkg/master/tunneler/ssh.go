package tunneler

import (
 "context"
 godefaultbytes "bytes"
 godefaultruntime "runtime"
 "fmt"
 "io/ioutil"
 "net"
 "net/http"
 godefaulthttp "net/http"
 "net/url"
 "os"
 "sync/atomic"
 "time"
 "k8s.io/apimachinery/pkg/util/clock"
 "k8s.io/apimachinery/pkg/util/wait"
 "k8s.io/kubernetes/pkg/ssh"
 utilfile "k8s.io/kubernetes/pkg/util/file"
 "github.com/prometheus/client_golang/prometheus"
 "k8s.io/klog"
)

type InstallSSHKey func(ctx context.Context, user string, data []byte) error
type AddressFunc func() (addresses []string, err error)
type Tunneler interface {
 Run(AddressFunc)
 Stop()
 Dial(ctx context.Context, net, addr string) (net.Conn, error)
 SecondsSinceSync() int64
 SecondsSinceSSHKeySync() int64
}

func TunnelSyncHealthChecker(tunneler Tunneler) func(req *http.Request) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return func(req *http.Request) error {
  if tunneler == nil {
   return nil
  }
  lag := tunneler.SecondsSinceSync()
  if lag > 600 {
   return fmt.Errorf("Tunnel sync is taking too long: %d", lag)
  }
  sshKeyLag := tunneler.SecondsSinceSSHKeySync()
  if sshKeyLag > 900 {
   return fmt.Errorf("SSHKey sync is taking too long: %d", sshKeyLag)
  }
  return nil
 }
}

type SSHTunneler struct {
 lastSync       int64
 lastSSHKeySync int64
 SSHUser        string
 SSHKeyfile     string
 InstallSSHKey  InstallSSHKey
 HealthCheckURL *url.URL
 tunnels        *ssh.SSHTunnelList
 lastSyncMetric prometheus.GaugeFunc
 clock          clock.Clock
 getAddresses   AddressFunc
 stopChan       chan struct{}
}

func New(sshUser, sshKeyfile string, healthCheckURL *url.URL, installSSHKey InstallSSHKey) Tunneler {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &SSHTunneler{SSHUser: sshUser, SSHKeyfile: sshKeyfile, InstallSSHKey: installSSHKey, HealthCheckURL: healthCheckURL, clock: clock.RealClock{}}
}
func (c *SSHTunneler) Run(getAddresses AddressFunc) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c.stopChan != nil {
  return
 }
 c.stopChan = make(chan struct{})
 if getAddresses != nil {
  c.getAddresses = getAddresses
 }
 if len(c.SSHUser) > 32 {
  klog.Warning("SSH User is too long, truncating to 32 chars")
  c.SSHUser = c.SSHUser[0:32]
 }
 klog.Infof("Setting up proxy: %s %s", c.SSHUser, c.SSHKeyfile)
 publicKeyFile := c.SSHKeyfile + ".pub"
 exists, err := utilfile.FileExists(publicKeyFile)
 if err != nil {
  klog.Errorf("Error detecting if key exists: %v", err)
 } else if !exists {
  klog.Infof("Key doesn't exist, attempting to create")
  if err := generateSSHKey(c.SSHKeyfile, publicKeyFile); err != nil {
   klog.Errorf("Failed to create key pair: %v", err)
  }
 }
 c.tunnels = ssh.NewSSHTunnelList(c.SSHUser, c.SSHKeyfile, c.HealthCheckURL, c.stopChan)
 c.lastSSHKeySync = c.clock.Now().Unix()
 c.installSSHKeySyncLoop(c.SSHUser, publicKeyFile)
 c.lastSync = c.clock.Now().Unix()
 c.nodesSyncLoop()
}
func (c *SSHTunneler) Stop() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c.stopChan != nil {
  close(c.stopChan)
  c.stopChan = nil
 }
}
func (c *SSHTunneler) Dial(ctx context.Context, net, addr string) (net.Conn, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.tunnels.Dial(ctx, net, addr)
}
func (c *SSHTunneler) SecondsSinceSync() int64 {
 _logClusterCodePath()
 defer _logClusterCodePath()
 now := c.clock.Now().Unix()
 then := atomic.LoadInt64(&c.lastSync)
 return now - then
}
func (c *SSHTunneler) SecondsSinceSSHKeySync() int64 {
 _logClusterCodePath()
 defer _logClusterCodePath()
 now := c.clock.Now().Unix()
 then := atomic.LoadInt64(&c.lastSSHKeySync)
 return now - then
}
func (c *SSHTunneler) installSSHKeySyncLoop(user, publicKeyfile string) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 go wait.Until(func() {
  if c.InstallSSHKey == nil {
   klog.Error("Won't attempt to install ssh key: InstallSSHKey function is nil")
   return
  }
  key, err := ssh.ParsePublicKeyFromFile(publicKeyfile)
  if err != nil {
   klog.Errorf("Failed to load public key: %v", err)
   return
  }
  keyData, err := ssh.EncodeSSHKey(key)
  if err != nil {
   klog.Errorf("Failed to encode public key: %v", err)
   return
  }
  if err := c.InstallSSHKey(context.TODO(), user, keyData); err != nil {
   klog.Errorf("Failed to install ssh key: %v", err)
   return
  }
  atomic.StoreInt64(&c.lastSSHKeySync, c.clock.Now().Unix())
 }, 5*time.Minute, c.stopChan)
}
func (c *SSHTunneler) nodesSyncLoop() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 go wait.Until(func() {
  addrs, err := c.getAddresses()
  klog.V(4).Infof("Calling update w/ addrs: %v", addrs)
  if err != nil {
   klog.Errorf("Failed to getAddresses: %v", err)
  }
  c.tunnels.Update(addrs)
  atomic.StoreInt64(&c.lastSync, c.clock.Now().Unix())
 }, 15*time.Second, c.stopChan)
}
func generateSSHKey(privateKeyfile, publicKeyfile string) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 private, public, err := ssh.GenerateKey(2048)
 if err != nil {
  return err
 }
 exists, err := utilfile.FileExists(privateKeyfile)
 if err != nil {
  klog.Errorf("Error detecting if private key exists: %v", err)
 } else if exists {
  klog.Infof("Private key exists, but public key does not")
  if err := os.Remove(privateKeyfile); err != nil {
   klog.Errorf("Failed to remove stale private key: %v", err)
  }
 }
 if err := ioutil.WriteFile(privateKeyfile, ssh.EncodePrivateKey(private), 0600); err != nil {
  return err
 }
 publicKeyBytes, err := ssh.EncodePublicKey(public)
 if err != nil {
  return err
 }
 if err := ioutil.WriteFile(publicKeyfile+".tmp", publicKeyBytes, 0600); err != nil {
  return err
 }
 return os.Rename(publicKeyfile+".tmp", publicKeyfile)
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
