package main

import (
	"bytes"
	"context"
	"fmt"
	clientv2 "github.com/coreos/etcd/client"
	"github.com/coreos/etcd/clientv3"
	"k8s.io/klog"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

type CombinedEtcdClient struct{ cfg *EtcdMigrateCfg }

func NewEtcdMigrateClient(cfg *EtcdMigrateCfg) (EtcdMigrateClient, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &CombinedEtcdClient{cfg}, nil
}
func (e *CombinedEtcdClient) Close() error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return nil
}
func (e *CombinedEtcdClient) SetEtcdVersionKeyValue(version *EtcdVersion) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return e.Put(version, "etcd_version", version.String())
}
func (e *CombinedEtcdClient) Put(version *EtcdVersion, key, value string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if version.Major == 2 {
		v2client, err := e.clientV2()
		if err != nil {
			return err
		}
		_, err = v2client.Set(context.Background(), key, value, nil)
		return err
	}
	v3client, err := e.clientV3()
	if err != nil {
		return err
	}
	defer v3client.Close()
	_, err = v3client.KV.Put(context.Background(), key, value)
	return err
}
func (e *CombinedEtcdClient) Get(version *EtcdVersion, key string) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if version.Major == 2 {
		v2client, err := e.clientV2()
		if err != nil {
			return "", err
		}
		resp, err := v2client.Get(context.Background(), key, nil)
		if err != nil {
			return "", err
		}
		return resp.Node.Value, nil
	}
	v3client, err := e.clientV3()
	if err != nil {
		return "", err
	}
	defer v3client.Close()
	resp, err := v3client.KV.Get(context.Background(), key)
	if err != nil {
		return "", err
	}
	kvs := resp.Kvs
	if len(kvs) != 1 {
		return "", fmt.Errorf("expected exactly one value for key %s but got %d", key, len(kvs))
	}
	return string(kvs[0].Value), nil
}
func (e *CombinedEtcdClient) clientV2() (clientv2.KeysAPI, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	v2client, err := clientv2.New(clientv2.Config{Endpoints: []string{e.endpoint()}})
	if err != nil {
		return nil, err
	}
	return clientv2.NewKeysAPI(v2client), nil
}
func (e *CombinedEtcdClient) clientV3() (*clientv3.Client, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return clientv3.New(clientv3.Config{Endpoints: []string{e.endpoint()}})
}
func (e *CombinedEtcdClient) Backup(version *EtcdVersion, backupDir string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if version.Major != 2 {
		return fmt.Errorf("etcd 2.x required but got version '%s'", version)
	}
	return e.runEtcdctlCommand(version, "--debug", "backup", "--data-dir", e.cfg.dataDirectory, "--backup-dir", backupDir)
}
func (e *CombinedEtcdClient) Snapshot(version *EtcdVersion, snapshotFile string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if version.Major != 3 {
		return fmt.Errorf("etcd 3.x required but got version '%s'", version)
	}
	return e.runEtcdctlCommand(version, "--endpoints", e.endpoint(), "snapshot", "save", snapshotFile)
}
func (e *CombinedEtcdClient) Restore(version *EtcdVersion, snapshotFile string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if version.Major != 3 {
		return fmt.Errorf("etcd 3.x required but got version '%s'", version)
	}
	return e.runEtcdctlCommand(version, "snapshot", "restore", snapshotFile, "--data-dir", e.cfg.dataDirectory, "--name", e.cfg.name, "--initial-advertise-peer-urls", e.cfg.peerAdvertiseUrls, "--initial-cluster", e.cfg.initialCluster)
}
func (e *CombinedEtcdClient) Migrate(version *EtcdVersion) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if version.Major != 3 {
		return fmt.Errorf("etcd 3.x required but got version '%s'", version)
	}
	return e.runEtcdctlCommand(version, "migrate", "--data-dir", e.cfg.dataDirectory)
}
func (e *CombinedEtcdClient) runEtcdctlCommand(version *EtcdVersion, args ...string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	etcdctlCmd := exec.Command(filepath.Join(e.cfg.binPath, fmt.Sprintf("etcdctl-%s", version)), args...)
	etcdctlCmd.Env = []string{fmt.Sprintf("ETCDCTL_API=%d", version.Major)}
	etcdctlCmd.Stdout = os.Stdout
	etcdctlCmd.Stderr = os.Stderr
	return etcdctlCmd.Run()
}
func (e *CombinedEtcdClient) AttachLease(leaseDuration time.Duration) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ttlKeysPrefix := e.cfg.ttlKeysDirectory
	if !strings.HasSuffix(ttlKeysPrefix, "/") {
		ttlKeysPrefix += "/"
	}
	ctx := context.Background()
	v3client, err := e.clientV3()
	if err != nil {
		return err
	}
	defer v3client.Close()
	objectsResp, err := v3client.KV.Get(ctx, ttlKeysPrefix, clientv3.WithPrefix())
	if err != nil {
		return fmt.Errorf("Error while getting objects to attach to the lease")
	}
	lease, err := v3client.Lease.Grant(ctx, int64(leaseDuration/time.Second))
	if err != nil {
		return fmt.Errorf("Error while creating lease: %v", err)
	}
	klog.Infof("Lease with TTL: %v created", lease.TTL)
	klog.Infof("Attaching lease to %d entries", len(objectsResp.Kvs))
	for _, kv := range objectsResp.Kvs {
		putResp, err := v3client.KV.Put(ctx, string(kv.Key), string(kv.Value), clientv3.WithLease(lease.ID), clientv3.WithPrevKV())
		if err != nil {
			klog.Errorf("Error while attaching lease to: %s", string(kv.Key))
		}
		if bytes.Compare(putResp.PrevKv.Value, kv.Value) != 0 {
			return fmt.Errorf("concurrent access to key detected when setting lease on %s, expected previous value of %s but got %s", kv.Key, kv.Value, putResp.PrevKv.Value)
		}
	}
	return nil
}
func (e *CombinedEtcdClient) endpoint() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return fmt.Sprintf("http://127.0.0.1:%d", e.cfg.port)
}
