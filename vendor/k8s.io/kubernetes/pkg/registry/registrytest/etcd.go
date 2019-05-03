package registrytest

import (
 "testing"
 etcdtesting "k8s.io/apiserver/pkg/storage/etcd/testing"
 "k8s.io/apiserver/pkg/storage/storagebackend"
 "k8s.io/kubernetes/pkg/api/testapi"
)

func NewEtcdStorage(t *testing.T, group string) (*storagebackend.Config, *etcdtesting.EtcdTestServer) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 server, config := etcdtesting.NewUnsecuredEtcd3TestClientServer(t)
 config.Codec = testapi.Groups[group].StorageCodec()
 return config, server
}
