package main

import (
	"fmt"
	"github.com/blang/semver"
	"k8s.io/klog"
	"os"
	"os/exec"
	"time"
)

type EtcdMigrateCfg struct {
	binPath           string
	name              string
	initialCluster    string
	port              uint64
	peerListenUrls    string
	peerAdvertiseUrls string
	etcdDataPrefix    string
	ttlKeysDirectory  string
	supportedVersions SupportedVersions
	dataDirectory     string
	etcdServerArgs    string
}
type EtcdMigrateClient interface {
	SetEtcdVersionKeyValue(version *EtcdVersion) error
	Get(version *EtcdVersion, key string) (string, error)
	Put(version *EtcdVersion, key, value string) error
	Backup(version *EtcdVersion, backupDir string) error
	Snapshot(version *EtcdVersion, snapshotFile string) error
	Restore(version *EtcdVersion, snapshotFile string) error
	Migrate(version *EtcdVersion) error
	AttachLease(leaseDuration time.Duration) error
	Close() error
}
type Migrator struct {
	cfg           *EtcdMigrateCfg
	dataDirectory *DataDirectory
	client        EtcdMigrateClient
}

func (m *Migrator) MigrateIfNeeded(target *EtcdVersionPair) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.Infof("Starting migration to %s", target)
	err := m.dataDirectory.Initialize(target)
	if err != nil {
		return fmt.Errorf("failed to initialize data directory %s: %v", m.dataDirectory.path, err)
	}
	var current *EtcdVersionPair
	vfExists, err := m.dataDirectory.versionFile.Exists()
	if err != nil {
		return err
	}
	if vfExists {
		current, err = m.dataDirectory.versionFile.Read()
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("existing data directory '%s' is missing version.txt file, unable to migrate", m.dataDirectory.path)
	}
	for {
		klog.Infof("Converging current version '%s' to target version '%s'", current, target)
		currentNextMinorVersion := &EtcdVersion{Version: semver.Version{Major: current.version.Major, Minor: current.version.Minor + 1}}
		switch {
		case current.version.MajorMinorEquals(target.version) || currentNextMinorVersion.MajorMinorEquals(target.version):
			klog.Infof("current version '%s' equals or is one minor version previous of target version '%s' - migration complete", current, target)
			err = m.dataDirectory.versionFile.Write(target)
			if err != nil {
				return fmt.Errorf("failed to write version.txt to '%s': %v", m.dataDirectory.path, err)
			}
			return nil
		case current.storageVersion == storageEtcd2 && target.storageVersion == storageEtcd3:
			klog.Infof("upgrading from etcd2 storage to etcd3 storage")
			current, err = m.etcd2ToEtcd3Upgrade(current, target)
		case current.version.Major == 3 && target.version.Major == 2:
			klog.Infof("downgrading from etcd 3.x to 2.x")
			current, err = m.rollbackToEtcd2(current, target)
		case current.version.Major == target.version.Major && current.version.Minor < target.version.Minor:
			stepVersion := m.cfg.supportedVersions.NextVersionPair(current)
			klog.Infof("upgrading etcd from %s to %s", current, stepVersion)
			current, err = m.minorVersionUpgrade(current, stepVersion)
		case current.version.Major == 3 && target.version.Major == 3 && current.version.Minor > target.version.Minor:
			klog.Infof("rolling etcd back from %s to %s", current, target)
			current, err = m.rollbackEtcd3MinorVersion(current, target)
		}
		if err != nil {
			return err
		}
	}
}
func (m *Migrator) backupEtcd2(current *EtcdVersion) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	backupDir := fmt.Sprintf("%s/%s", m.dataDirectory, "migration-backup")
	klog.Infof("Backup etcd before starting migration")
	err := os.Mkdir(backupDir, 0666)
	if err != nil {
		return fmt.Errorf("failed to create backup directory before starting migration: %v", err)
	}
	m.client.Backup(current, backupDir)
	klog.Infof("Backup done in %s", backupDir)
	return nil
}
func (m *Migrator) rollbackEtcd3MinorVersion(current *EtcdVersionPair, target *EtcdVersionPair) (*EtcdVersionPair, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if target.version.Minor != current.version.Minor-1 {
		return nil, fmt.Errorf("rollback from %s to %s not supported, only rollbacks to the previous minor version are supported", current.version, target.version)
	}
	klog.Infof("Performing etcd %s -> %s rollback", current.version, target.version)
	err := m.dataDirectory.Backup()
	if err != nil {
		return nil, err
	}
	snapshotFilename := fmt.Sprintf("%s.snapshot.db", m.dataDirectory.path)
	err = os.Remove(snapshotFilename)
	if err != nil && !os.IsNotExist(err) {
		return nil, fmt.Errorf("failed to clean snapshot file before rollback: %v", err)
	}
	runner := m.newServer()
	klog.Infof("Starting etcd version %s to capture rollback snapshot.", current.version)
	err = runner.Start(current.version)
	if err != nil {
		klog.Fatalf("Unable to automatically downgrade etcd: starting etcd version %s to capture rollback snapshot failed: %v", current.version, err)
		return nil, err
	}
	klog.Infof("Snapshotting etcd %s to %s", current.version, snapshotFilename)
	err = m.client.Snapshot(current.version, snapshotFilename)
	if err != nil {
		return nil, err
	}
	err = runner.Stop()
	if err != nil {
		return nil, err
	}
	klog.Infof("Backing up data before rolling back")
	backupDir := fmt.Sprintf("%s.bak", m.dataDirectory)
	err = os.RemoveAll(backupDir)
	if err != nil {
		return nil, err
	}
	origInfo, err := os.Stat(m.dataDirectory.path)
	if err != nil {
		return nil, err
	}
	err = exec.Command("mv", m.dataDirectory.path, backupDir).Run()
	if err != nil {
		return nil, err
	}
	klog.Infof("Restoring etcd %s from %s", target.version, snapshotFilename)
	err = m.client.Restore(target.version, snapshotFilename)
	if err != nil {
		return nil, err
	}
	err = os.Chmod(m.dataDirectory.path, origInfo.Mode())
	if err != nil {
		return nil, err
	}
	return target, nil
}
func (m *Migrator) rollbackToEtcd2(current *EtcdVersionPair, target *EtcdVersionPair) (*EtcdVersionPair, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if !(current.version.Major == 3 && current.version.Minor == 0 && target.version.Major == 2 && target.version.Minor == 2) {
		return nil, fmt.Errorf("etcd3 -> etcd2 downgrade is supported only between 3.0.x and 2.2.x, got current %s target %s", current, target)
	}
	klog.Infof("Backup and remove all existing v2 data")
	err := m.dataDirectory.Backup()
	if err != nil {
		return nil, err
	}
	err = RollbackV3ToV2(m.dataDirectory.path, time.Hour)
	if err != nil {
		return nil, fmt.Errorf("rollback to etcd 2.x failed: %v", err)
	}
	return target, nil
}
func (m *Migrator) etcd2ToEtcd3Upgrade(current *EtcdVersionPair, target *EtcdVersionPair) (*EtcdVersionPair, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if current.storageVersion != storageEtcd2 || target.version.Major != 3 || target.storageVersion != storageEtcd3 {
		return nil, fmt.Errorf("etcd2 to etcd3 upgrade is supported only for x.x.x/etcd2 to 3.0.x/etcd3, got current %s target %s", current, target)
	}
	runner := m.newServer()
	klog.Infof("Performing etcd2 -> etcd3 migration")
	err := m.client.Migrate(target.version)
	if err != nil {
		return nil, err
	}
	klog.Infof("Attaching leases to TTL entries")
	err = runner.Start(target.version)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = runner.Stop()
	}()
	err = m.client.AttachLease(1 * time.Hour)
	if err != nil {
		return nil, err
	}
	return target, err
}
func (m *Migrator) minorVersionUpgrade(current *EtcdVersionPair, target *EtcdVersionPair) (*EtcdVersionPair, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	runner := m.newServer()
	err := runner.Start(target.version)
	if err != nil {
		return nil, err
	}
	err = runner.Stop()
	return target, err
}
func (m *Migrator) newServer() *EtcdMigrateServer {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return NewEtcdMigrateServer(m.cfg, m.client)
}
