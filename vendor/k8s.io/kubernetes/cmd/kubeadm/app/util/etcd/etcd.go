package etcd

import (
	"context"
	"crypto/tls"
	"fmt"
	goformat "fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/pkg/transport"
	"github.com/pkg/errors"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/klog"
	kubeadmapi "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm"
	"k8s.io/kubernetes/cmd/kubeadm/app/constants"
	"k8s.io/kubernetes/cmd/kubeadm/app/util/config"
	"k8s.io/kubernetes/cmd/kubeadm/app/util/staticpod"
	"net"
	"net/url"
	goos "os"
	"path/filepath"
	godefaultruntime "runtime"
	"strconv"
	"strings"
	"time"
	gotime "time"
)

type ClusterInterrogator interface {
	ClusterAvailable() (bool, error)
	GetClusterStatus() (map[string]*clientv3.StatusResponse, error)
	GetClusterVersions() (map[string]string, error)
	GetVersion() (string, error)
	HasTLS() bool
	WaitForClusterAvailable(delay time.Duration, retries int, retryInterval time.Duration) (bool, error)
	Sync() error
	AddMember(name string, peerAddrs string) ([]Member, error)
}
type Client struct {
	Endpoints []string
	TLS       *tls.Config
}

func (c Client) HasTLS() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return c.TLS != nil
}
func PodManifestsHaveTLS(ManifestDir string) (bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	etcdPodPath := constants.GetStaticPodFilepath(constants.Etcd, ManifestDir)
	etcdPod, err := staticpod.ReadStaticPodFromDisk(etcdPodPath)
	if err != nil {
		return false, errors.Wrap(err, "failed to check if etcd pod implements TLS")
	}
	tlsFlags := []string{"--cert-file=", "--key-file=", "--trusted-ca-file=", "--client-cert-auth=", "--peer-cert-file=", "--peer-key-file=", "--peer-trusted-ca-file=", "--peer-client-cert-auth="}
FlagLoop:
	for _, flag := range tlsFlags {
		for _, container := range etcdPod.Spec.Containers {
			for _, arg := range container.Command {
				if strings.Contains(arg, flag) {
					continue FlagLoop
				}
			}
		}
		return false, nil
	}
	return true, nil
}
func New(endpoints []string, ca, cert, key string) (*Client, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	client := Client{Endpoints: endpoints}
	if ca != "" || cert != "" || key != "" {
		tlsInfo := transport.TLSInfo{CertFile: cert, KeyFile: key, TrustedCAFile: ca}
		tlsConfig, err := tlsInfo.ClientConfig()
		if err != nil {
			return nil, err
		}
		client.TLS = tlsConfig
	}
	return &client, nil
}
func NewFromCluster(client clientset.Interface, certificatesDir string) (*Client, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	oldManifest := false
	klog.V(1).Infoln("checking etcd manifest")
	etcdManifestFile := constants.GetStaticPodFilepath(constants.Etcd, constants.GetStaticPodDirectory())
	etcdPod, err := staticpod.ReadStaticPodFromDisk(etcdManifestFile)
	if err == nil {
		etcdContainer := etcdPod.Spec.Containers[0]
		for _, arg := range etcdContainer.Command {
			if arg == "--listen-client-urls=https://127.0.0.1:2379" {
				klog.V(1).Infoln("etcd manifest created by kubeadm v1.12")
				oldManifest = true
			}
		}
		if oldManifest == true {
			endpoints := []string{fmt.Sprintf("localhost:%d", constants.EtcdListenClientPort)}
			etcdClient, err := New(endpoints, filepath.Join(certificatesDir, constants.EtcdCACertName), filepath.Join(certificatesDir, constants.EtcdHealthcheckClientCertName), filepath.Join(certificatesDir, constants.EtcdHealthcheckClientKeyName))
			if err != nil {
				return nil, errors.Wrapf(err, "error creating etcd client for %v endpoint", endpoints)
			}
			return etcdClient, nil
		}
	}
	clusterStatus, err := config.GetClusterStatus(client)
	if err != nil {
		return nil, err
	}
	endpoints := []string{}
	for _, e := range clusterStatus.APIEndpoints {
		endpoints = append(endpoints, GetClientURLByIP(e.AdvertiseAddress))
	}
	klog.V(1).Infof("etcd endpoints read from pods: %s", strings.Join(endpoints, ","))
	etcdClient, err := New(endpoints, filepath.Join(certificatesDir, constants.EtcdCACertName), filepath.Join(certificatesDir, constants.EtcdHealthcheckClientCertName), filepath.Join(certificatesDir, constants.EtcdHealthcheckClientKeyName))
	if err != nil {
		return nil, errors.Wrapf(err, "error creating etcd client for %v endpoints", endpoints)
	}
	err = etcdClient.Sync()
	if err != nil {
		return nil, errors.Wrap(err, "error syncing endpoints with etc")
	}
	klog.V(1).Infof("update etcd endpoints: %s", strings.Join(etcdClient.Endpoints, ","))
	return etcdClient, nil
}
func (c *Client) Sync() error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	cli, err := clientv3.New(clientv3.Config{Endpoints: c.Endpoints, DialTimeout: 20 * time.Second, TLS: c.TLS})
	if err != nil {
		return err
	}
	defer cli.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = cli.Sync(ctx)
	cancel()
	if err != nil {
		return err
	}
	klog.V(1).Infof("etcd endpoints read from etcd: %s", strings.Join(cli.Endpoints(), ","))
	c.Endpoints = cli.Endpoints()
	return nil
}

type Member struct {
	Name    string
	PeerURL string
}

func (c *Client) AddMember(name string, peerAddrs string) ([]Member, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	parsedPeerAddrs, err := url.Parse(peerAddrs)
	if err != nil {
		return nil, errors.Wrapf(err, "error parsing peer address %s", peerAddrs)
	}
	cli, err := clientv3.New(clientv3.Config{Endpoints: c.Endpoints, DialTimeout: 20 * time.Second, TLS: c.TLS})
	if err != nil {
		return nil, err
	}
	defer cli.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	resp, err := cli.MemberAdd(ctx, []string{peerAddrs})
	cancel()
	if err != nil {
		return nil, err
	}
	ret := []Member{}
	for _, m := range resp.Members {
		if m.Name == "" {
			ret = append(ret, Member{Name: name, PeerURL: m.PeerURLs[0]})
		} else {
			ret = append(ret, Member{Name: m.Name, PeerURL: m.PeerURLs[0]})
		}
	}
	c.Endpoints = append(c.Endpoints, GetClientURLByIP(parsedPeerAddrs.Hostname()))
	return ret, nil
}
func (c Client) GetVersion() (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var clusterVersion string
	versions, err := c.GetClusterVersions()
	if err != nil {
		return "", err
	}
	for _, v := range versions {
		if clusterVersion != "" && clusterVersion != v {
			return "", errors.Errorf("etcd cluster contains endpoints with mismatched versions: %v", versions)
		}
		clusterVersion = v
	}
	if clusterVersion == "" {
		return "", errors.New("could not determine cluster etcd version")
	}
	return clusterVersion, nil
}
func (c Client) GetClusterVersions() (map[string]string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	versions := make(map[string]string)
	statuses, err := c.GetClusterStatus()
	if err != nil {
		return versions, err
	}
	for ep, status := range statuses {
		versions[ep] = status.Version
	}
	return versions, nil
}
func (c Client) ClusterAvailable() (bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	_, err := c.GetClusterStatus()
	if err != nil {
		return false, err
	}
	return true, nil
}
func (c Client) GetClusterStatus() (map[string]*clientv3.StatusResponse, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	cli, err := clientv3.New(clientv3.Config{Endpoints: c.Endpoints, DialTimeout: 5 * time.Second, TLS: c.TLS})
	if err != nil {
		return nil, err
	}
	defer cli.Close()
	clusterStatus := make(map[string]*clientv3.StatusResponse)
	for _, ep := range c.Endpoints {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		resp, err := cli.Status(ctx, ep)
		cancel()
		if err != nil {
			return nil, err
		}
		clusterStatus[ep] = resp
	}
	return clusterStatus, nil
}
func (c Client) WaitForClusterAvailable(delay time.Duration, retries int, retryInterval time.Duration) (bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	fmt.Printf("[util/etcd] Waiting %v for initial delay\n", delay)
	time.Sleep(delay)
	for i := 0; i < retries; i++ {
		if i > 0 {
			fmt.Printf("[util/etcd] Waiting %v until next retry\n", retryInterval)
			time.Sleep(retryInterval)
		}
		klog.V(2).Infof("attempting to see if all cluster endpoints (%s) are available %d/%d", c.Endpoints, i+1, retries)
		resp, err := c.ClusterAvailable()
		if err != nil {
			switch err {
			case context.DeadlineExceeded:
				fmt.Println("[util/etcd] Attempt timed out")
			default:
				fmt.Printf("[util/etcd] Attempt failed with error: %v\n", err)
			}
			continue
		}
		return resp, nil
	}
	return false, errors.New("timeout waiting for etcd cluster to be available")
}
func CheckConfigurationIsHA(cfg *kubeadmapi.Etcd) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return cfg.External != nil && len(cfg.External.Endpoints) > 1
}
func GetClientURL(cfg *kubeadmapi.InitConfiguration) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return "https://" + net.JoinHostPort(cfg.LocalAPIEndpoint.AdvertiseAddress, strconv.Itoa(constants.EtcdListenClientPort))
}
func GetPeerURL(cfg *kubeadmapi.InitConfiguration) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return "https://" + net.JoinHostPort(cfg.LocalAPIEndpoint.AdvertiseAddress, strconv.Itoa(constants.EtcdListenPeerPort))
}
func GetClientURLByIP(ip string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return "https://" + net.JoinHostPort(ip, strconv.Itoa(constants.EtcdListenClientPort))
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
