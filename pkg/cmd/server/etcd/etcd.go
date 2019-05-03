package etcd

import (
	"fmt"
	"github.com/coreos/etcd/clientv3"
	configapi "github.com/openshift/origin/pkg/cmd/server/apis/config"
	"golang.org/x/net/context"
	etcdutil "k8s.io/apiserver/pkg/storage/etcd/util"
	restclient "k8s.io/client-go/rest"
	"time"
)

func MakeEtcdClientV3Config(etcdClientInfo configapi.EtcdConnectionInfo) (*clientv3.Config, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	tlsConfig, err := restclient.TLSConfigFor(&restclient.Config{TLSClientConfig: restclient.TLSClientConfig{CertFile: etcdClientInfo.ClientCert.CertFile, KeyFile: etcdClientInfo.ClientCert.KeyFile, CAFile: etcdClientInfo.CA}})
	if err != nil {
		return nil, err
	}
	return &clientv3.Config{Endpoints: etcdClientInfo.URLs, DialTimeout: 30 * time.Second, TLS: tlsConfig}, nil
}
func MakeEtcdClientV3(etcdClientInfo configapi.EtcdConnectionInfo) (*clientv3.Client, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cfg, err := MakeEtcdClientV3Config(etcdClientInfo)
	if err != nil {
		return nil, err
	}
	return clientv3.New(*cfg)
}
func TestEtcdClientV3(etcdClient *clientv3.Client) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for i := 0; ; i++ {
		_, err := clientv3.NewKV(etcdClient).Get(context.Background(), "/", clientv3.WithLimit(1))
		if err == nil || etcdutil.IsEtcdNotFound(err) {
			break
		}
		if i > 100 {
			return fmt.Errorf("could not reach etcd(v3): %v", err)
		}
		time.Sleep(50 * time.Millisecond)
	}
	return nil
}
