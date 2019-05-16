package etcdserver

import (
	goformat "fmt"
	"github.com/coreos/etcd/embed"
	"github.com/coreos/etcd/pkg/osutil"
	"github.com/coreos/etcd/pkg/types"
	"github.com/coreos/go-semver/semver"
	configapi "github.com/openshift/origin/pkg/cmd/server/apis/config"
	"k8s.io/klog"
	"net/url"
	goos "os"
	godefaultruntime "runtime"
	"strings"
	"time"
	gotime "time"
)

const defaultName = "openshift.local"

func RunEtcd(etcdServerConfig *configapi.EtcdConfig) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	cfg := embed.NewConfig()
	cfg.Debug = bool(klog.V(4))
	cfg.Name = defaultName
	cfg.Dir = etcdServerConfig.StorageDir
	clientTLS := true
	cfg.ClientTLSInfo.CAFile = etcdServerConfig.ServingInfo.ClientCA
	cfg.ClientTLSInfo.CertFile = etcdServerConfig.ServingInfo.ServerCert.CertFile
	cfg.ClientTLSInfo.KeyFile = etcdServerConfig.ServingInfo.ServerCert.KeyFile
	cfg.ClientTLSInfo.ClientCertAuth = len(cfg.ClientTLSInfo.CAFile) > 0
	u, err := types.NewURLs(addressToURLs(etcdServerConfig.ServingInfo.BindAddress, clientTLS))
	if err != nil {
		klog.Fatalf("Unable to build etcd peer URLs: %v", err)
	}
	cfg.LCUrls = []url.URL(u)
	peerTLS := true
	cfg.PeerTLSInfo.CAFile = etcdServerConfig.PeerServingInfo.ClientCA
	cfg.PeerTLSInfo.CertFile = etcdServerConfig.PeerServingInfo.ServerCert.CertFile
	cfg.PeerTLSInfo.KeyFile = etcdServerConfig.PeerServingInfo.ServerCert.KeyFile
	cfg.PeerTLSInfo.ClientCertAuth = len(cfg.PeerTLSInfo.CAFile) > 0
	u, err = types.NewURLs(addressToURLs(etcdServerConfig.PeerServingInfo.BindAddress, peerTLS))
	if err != nil {
		klog.Fatalf("Unable to build etcd peer URLs: %v", err)
	}
	cfg.LPUrls = []url.URL(u)
	u, err = types.NewURLs(addressToURLs(etcdServerConfig.Address, clientTLS))
	if err != nil {
		klog.Fatalf("Unable to build etcd announce client URLs: %v", err)
	}
	cfg.ACUrls = []url.URL(u)
	u, err = types.NewURLs(addressToURLs(etcdServerConfig.PeerAddress, peerTLS))
	if err != nil {
		klog.Fatalf("Unable to build etcd announce peer URLs: %v", err)
	}
	cfg.APUrls = []url.URL(u)
	cfg.InitialCluster = cfg.InitialClusterFromName(cfg.Name)
	osutil.HandleInterrupts()
	e, err := embed.StartEtcd(cfg)
	if err != nil {
		klog.Fatalf("Unable to start etcd: %v", err)
	}
	ready := make(chan struct{})
	go func() {
		defer e.Close()
		select {
		case <-e.Server.ReadyNotify():
			klog.V(4).Infof("Waiting for etcd to reach cluster version 3.0.0")
			for min := semver.Must(semver.NewVersion("3.0.0")); e.Server.ClusterVersion() == nil || e.Server.ClusterVersion().LessThan(*min); {
				time.Sleep(25 * time.Millisecond)
			}
			close(ready)
			klog.Infof("Started etcd at %s", etcdServerConfig.Address)
		case <-time.After(60 * time.Second):
			klog.Warning("etcd took too long to start, stopped")
			e.Server.Stop()
		}
		klog.Fatalf("etcd has returned an error: %v", <-e.Err())
	}()
	<-ready
}
func addressToURLs(addr string, isTLS bool) []string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	addrs := strings.Split(addr, ",")
	for i := range addrs {
		if strings.HasPrefix(addrs[i], "unix://") || strings.HasPrefix(addrs[i], "unixs://") {
			continue
		}
		if isTLS {
			addrs[i] = "https://" + addrs[i]
		} else {
			addrs[i] = "http://" + addrs[i]
		}
	}
	return addrs
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
