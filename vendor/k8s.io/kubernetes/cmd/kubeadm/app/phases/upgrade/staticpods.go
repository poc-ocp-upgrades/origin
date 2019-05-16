package upgrade

import (
	"fmt"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/util/version"
	clientset "k8s.io/client-go/kubernetes"
	kubeadmapi "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm"
	"k8s.io/kubernetes/cmd/kubeadm/app/constants"
	certsphase "k8s.io/kubernetes/cmd/kubeadm/app/phases/certs"
	"k8s.io/kubernetes/cmd/kubeadm/app/phases/certs/renewal"
	controlplanephase "k8s.io/kubernetes/cmd/kubeadm/app/phases/controlplane"
	etcdphase "k8s.io/kubernetes/cmd/kubeadm/app/phases/etcd"
	"k8s.io/kubernetes/cmd/kubeadm/app/util"
	"k8s.io/kubernetes/cmd/kubeadm/app/util/apiclient"
	etcdutil "k8s.io/kubernetes/cmd/kubeadm/app/util/etcd"
	"k8s.io/kubernetes/cmd/kubeadm/app/util/staticpod"
	"os"
	"strings"
	"time"
)

const (
	UpgradeManifestTimeout = 5 * time.Minute
)

type StaticPodPathManager interface {
	MoveFile(oldPath, newPath string) error
	RealManifestPath(component string) string
	RealManifestDir() string
	TempManifestPath(component string) string
	TempManifestDir() string
	BackupManifestPath(component string) string
	BackupManifestDir() string
	BackupEtcdDir() string
	CleanupDirs() error
}
type KubeStaticPodPathManager struct {
	realManifestDir   string
	tempManifestDir   string
	backupManifestDir string
	backupEtcdDir     string
	keepManifestDir   bool
	keepEtcdDir       bool
}

func NewKubeStaticPodPathManager(realDir, tempDir, backupDir, backupEtcdDir string, keepManifestDir, keepEtcdDir bool) StaticPodPathManager {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &KubeStaticPodPathManager{realManifestDir: realDir, tempManifestDir: tempDir, backupManifestDir: backupDir, backupEtcdDir: backupEtcdDir, keepManifestDir: keepManifestDir, keepEtcdDir: keepEtcdDir}
}
func NewKubeStaticPodPathManagerUsingTempDirs(realManifestDir string, saveManifestsDir, saveEtcdDir bool) (StaticPodPathManager, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	upgradedManifestsDir, err := constants.CreateTempDirForKubeadm("kubeadm-upgraded-manifests")
	if err != nil {
		return nil, err
	}
	backupManifestsDir, err := constants.CreateTimestampDirForKubeadm("kubeadm-backup-manifests")
	if err != nil {
		return nil, err
	}
	backupEtcdDir, err := constants.CreateTimestampDirForKubeadm("kubeadm-backup-etcd")
	if err != nil {
		return nil, err
	}
	return NewKubeStaticPodPathManager(realManifestDir, upgradedManifestsDir, backupManifestsDir, backupEtcdDir, saveManifestsDir, saveEtcdDir), nil
}
func (spm *KubeStaticPodPathManager) MoveFile(oldPath, newPath string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return os.Rename(oldPath, newPath)
}
func (spm *KubeStaticPodPathManager) RealManifestPath(component string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return constants.GetStaticPodFilepath(component, spm.realManifestDir)
}
func (spm *KubeStaticPodPathManager) RealManifestDir() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return spm.realManifestDir
}
func (spm *KubeStaticPodPathManager) TempManifestPath(component string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return constants.GetStaticPodFilepath(component, spm.tempManifestDir)
}
func (spm *KubeStaticPodPathManager) TempManifestDir() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return spm.tempManifestDir
}
func (spm *KubeStaticPodPathManager) BackupManifestPath(component string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return constants.GetStaticPodFilepath(component, spm.backupManifestDir)
}
func (spm *KubeStaticPodPathManager) BackupManifestDir() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return spm.backupManifestDir
}
func (spm *KubeStaticPodPathManager) BackupEtcdDir() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return spm.backupEtcdDir
}
func (spm *KubeStaticPodPathManager) CleanupDirs() error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if err := os.RemoveAll(spm.TempManifestDir()); err != nil {
		return err
	}
	if !spm.keepManifestDir {
		if err := os.RemoveAll(spm.BackupManifestDir()); err != nil {
			return err
		}
	}
	if !spm.keepEtcdDir {
		if err := os.RemoveAll(spm.BackupEtcdDir()); err != nil {
			return err
		}
	}
	return nil
}
func upgradeComponent(component string, waiter apiclient.Waiter, pathMgr StaticPodPathManager, cfg *kubeadmapi.InitConfiguration, beforePodHash string, recoverManifests map[string]string, isTLSUpgrade bool) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	recoverEtcd := false
	waitForComponentRestart := true
	if component == constants.Etcd {
		recoverEtcd = true
	}
	if isTLSUpgrade {
		if component == constants.Etcd {
			waitForComponentRestart = false
		}
		if component == constants.KubeAPIServer {
			recoverEtcd = true
			fmt.Printf("[upgrade/staticpods] The %s manifest will be restored if component %q fails to upgrade\n", constants.Etcd, component)
		}
	}
	if err := renewCerts(cfg, component); err != nil {
		return errors.Wrapf(err, "failed to renew certificates for component %q", component)
	}
	currentManifestPath := pathMgr.RealManifestPath(component)
	newManifestPath := pathMgr.TempManifestPath(component)
	backupManifestPath := pathMgr.BackupManifestPath(component)
	recoverManifests[component] = backupManifestPath
	equal, err := staticpod.ManifestFilesAreEqual(currentManifestPath, newManifestPath)
	if err != nil {
		return err
	}
	if equal {
		fmt.Printf("[upgrade/staticpods] current and new manifests of %s are equal, skipping upgrade\n", component)
		return nil
	}
	if err := pathMgr.MoveFile(currentManifestPath, backupManifestPath); err != nil {
		return rollbackOldManifests(recoverManifests, err, pathMgr, recoverEtcd)
	}
	if err := pathMgr.MoveFile(newManifestPath, currentManifestPath); err != nil {
		return rollbackOldManifests(recoverManifests, err, pathMgr, recoverEtcd)
	}
	fmt.Printf("[upgrade/staticpods] Moved new manifest to %q and backed up old manifest to %q\n", currentManifestPath, backupManifestPath)
	if waitForComponentRestart {
		fmt.Println("[upgrade/staticpods] Waiting for the kubelet to restart the component")
		fmt.Printf("[upgrade/staticpods] This might take a minute or longer depending on the component/version gap (timeout %v)\n", UpgradeManifestTimeout)
		if err := waiter.WaitForStaticPodHashChange(cfg.NodeRegistration.Name, component, beforePodHash); err != nil {
			return rollbackOldManifests(recoverManifests, err, pathMgr, recoverEtcd)
		}
		if err := waiter.WaitForPodsWithLabel("component=" + component); err != nil {
			return rollbackOldManifests(recoverManifests, err, pathMgr, recoverEtcd)
		}
		fmt.Printf("[upgrade/staticpods] Component %q upgraded successfully!\n", component)
	} else {
		fmt.Printf("[upgrade/staticpods] Not waiting for pod-hash change for component %q\n", component)
	}
	return nil
}
func performEtcdStaticPodUpgrade(client clientset.Interface, waiter apiclient.Waiter, pathMgr StaticPodPathManager, cfg *kubeadmapi.InitConfiguration, recoverManifests map[string]string, isTLSUpgrade bool, oldEtcdClient, newEtcdClient etcdutil.ClusterInterrogator) (bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if cfg.Etcd.External != nil {
		return false, errors.New("external etcd detected, won't try to change any etcd state")
	}
	_, err := oldEtcdClient.GetClusterStatus()
	if err != nil {
		return true, errors.Wrap(err, "etcd cluster is not healthy")
	}
	backupEtcdDir := pathMgr.BackupEtcdDir()
	runningEtcdDir := cfg.Etcd.Local.DataDir
	if err := util.CopyDir(runningEtcdDir, backupEtcdDir); err != nil {
		return true, errors.Wrap(err, "failed to back up etcd data")
	}
	desiredEtcdVersion, err := constants.EtcdSupportedVersion(cfg.KubernetesVersion)
	if err != nil {
		return true, errors.Wrap(err, "failed to retrieve an etcd version for the target Kubernetes version")
	}
	currentEtcdVersions, err := oldEtcdClient.GetClusterVersions()
	if err != nil {
		return true, errors.Wrap(err, "failed to retrieve the current etcd version")
	}
	currentEtcdVersionStr, ok := currentEtcdVersions[etcdutil.GetClientURL(cfg)]
	if !ok {
		fmt.Println(currentEtcdVersions)
		return true, errors.Wrap(err, "failed to retrieve the current etcd version")
	}
	currentEtcdVersion, err := version.ParseSemantic(currentEtcdVersionStr)
	if err != nil {
		return true, errors.Wrapf(err, "failed to parse the current etcd version(%s)", currentEtcdVersionStr)
	}
	if desiredEtcdVersion.LessThan(currentEtcdVersion) {
		return false, errors.Errorf("the desired etcd version for this Kubernetes version %q is %q, but the current etcd version is %q. Won't downgrade etcd, instead just continue", cfg.KubernetesVersion, desiredEtcdVersion.String(), currentEtcdVersion.String())
	}
	if strings.Compare(desiredEtcdVersion.String(), currentEtcdVersion.String()) == 0 {
		return false, nil
	}
	beforeEtcdPodHash, err := waiter.WaitForStaticPodSingleHash(cfg.NodeRegistration.Name, constants.Etcd)
	if err != nil {
		return true, errors.Wrap(err, "failed to get etcd pod's hash")
	}
	if err := etcdphase.CreateLocalEtcdStaticPodManifestFile(pathMgr.TempManifestDir(), cfg); err != nil {
		return true, errors.Wrap(err, "error creating local etcd static pod manifest file")
	}
	noDelay := 0 * time.Second
	podRestartDelay := noDelay
	if isTLSUpgrade {
		podRestartDelay = 30 * time.Second
	}
	retries := 10
	retryInterval := 15 * time.Second
	if err := upgradeComponent(constants.Etcd, waiter, pathMgr, cfg, beforeEtcdPodHash, recoverManifests, isTLSUpgrade); err != nil {
		fmt.Printf("[upgrade/etcd] Failed to upgrade etcd: %v\n", err)
		fmt.Println("[upgrade/etcd] Waiting for previous etcd to become available")
		if _, err := oldEtcdClient.WaitForClusterAvailable(noDelay, retries, retryInterval); err != nil {
			fmt.Printf("[upgrade/etcd] Failed to healthcheck previous etcd: %v\n", err)
			fmt.Println("[upgrade/etcd] Rolling back etcd data")
			if err := rollbackEtcdData(cfg, pathMgr); err != nil {
				return true, errors.Errorf("fatal error rolling back local etcd cluster datadir: %v, the backup of etcd database is stored here:(%s)", err, backupEtcdDir)
			}
			fmt.Println("[upgrade/etcd] Etcd data rollback successful")
			fmt.Println("[upgrade/etcd] Waiting for previous etcd to become available")
			if _, err := oldEtcdClient.WaitForClusterAvailable(noDelay, retries, retryInterval); err != nil {
				fmt.Printf("[upgrade/etcd] Failed to healthcheck previous etcd: %v\n", err)
				return true, errors.Wrapf(err, "fatal error rolling back local etcd cluster manifest, the backup of etcd database is stored here:(%s)", backupEtcdDir)
			}
		}
		fmt.Println("[upgrade/etcd] Etcd was rolled back and is now available")
		return true, errors.Wrap(err, "fatal error when trying to upgrade the etcd cluster, rolled the state back to pre-upgrade state")
	}
	if newEtcdClient == nil {
		etcdClient, err := etcdutil.NewFromCluster(client, cfg.CertificatesDir)
		if err != nil {
			return true, errors.Wrap(err, "fatal error creating etcd client")
		}
		newEtcdClient = etcdClient
	}
	fmt.Println("[upgrade/etcd] Waiting for etcd to become available")
	if _, err = newEtcdClient.WaitForClusterAvailable(podRestartDelay, retries, retryInterval); err != nil {
		fmt.Printf("[upgrade/etcd] Failed to healthcheck etcd: %v\n", err)
		fmt.Println("[upgrade/etcd] Rolling back etcd data")
		if err := rollbackEtcdData(cfg, pathMgr); err != nil {
			return true, errors.Wrapf(err, "fatal error rolling back local etcd cluster datadir, the backup of etcd database is stored here:(%s)", backupEtcdDir)
		}
		fmt.Println("[upgrade/etcd] Etcd data rollback successful")
		fmt.Println("[upgrade/etcd] Rolling back etcd manifest")
		rollbackOldManifests(recoverManifests, err, pathMgr, true)
		fmt.Println("[upgrade/etcd] Waiting for previous etcd to become available")
		if _, err := oldEtcdClient.WaitForClusterAvailable(noDelay, retries, retryInterval); err != nil {
			fmt.Printf("[upgrade/etcd] Failed to healthcheck previous etcd: %v\n", err)
			return true, errors.Wrapf(err, "fatal error rolling back local etcd cluster manifest, the backup of etcd database is stored here:(%s)", backupEtcdDir)
		}
		fmt.Println("[upgrade/etcd] Etcd was rolled back and is now available")
		return true, errors.Wrap(err, "fatal error upgrading local etcd cluster, rolled the state back to pre-upgrade state")
	}
	return false, nil
}
func StaticPodControlPlane(client clientset.Interface, waiter apiclient.Waiter, pathMgr StaticPodPathManager, cfg *kubeadmapi.InitConfiguration, etcdUpgrade bool, oldEtcdClient, newEtcdClient etcdutil.ClusterInterrogator) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	recoverManifests := map[string]string{}
	var isTLSUpgrade bool
	var isExternalEtcd bool
	beforePodHashMap, err := waiter.WaitForStaticPodControlPlaneHashes(cfg.NodeRegistration.Name)
	if err != nil {
		return err
	}
	if oldEtcdClient == nil {
		if cfg.Etcd.External != nil {
			isExternalEtcd = true
			etcdClient, err := etcdutil.New(cfg.Etcd.External.Endpoints, cfg.Etcd.External.CAFile, cfg.Etcd.External.CertFile, cfg.Etcd.External.KeyFile)
			if err != nil {
				return errors.Wrap(err, "failed to create etcd client for external etcd")
			}
			oldEtcdClient = etcdClient
			if newEtcdClient == nil {
				newEtcdClient = etcdClient
			}
		} else {
			etcdClient, err := etcdutil.NewFromCluster(client, cfg.CertificatesDir)
			if err != nil {
				return errors.Wrap(err, "failed to create etcd client")
			}
			oldEtcdClient = etcdClient
		}
	}
	if !isExternalEtcd && etcdUpgrade {
		previousEtcdHasTLS := oldEtcdClient.HasTLS()
		isTLSUpgrade = !previousEtcdHasTLS
		if isTLSUpgrade {
			fmt.Printf("[upgrade/etcd] Upgrading to TLS for %s\n", constants.Etcd)
		}
		fatal, err := performEtcdStaticPodUpgrade(client, waiter, pathMgr, cfg, recoverManifests, isTLSUpgrade, oldEtcdClient, newEtcdClient)
		if err != nil {
			if fatal {
				return err
			}
			fmt.Printf("[upgrade/etcd] non fatal issue encountered during upgrade: %v\n", err)
		}
	}
	fmt.Printf("[upgrade/staticpods] Writing new Static Pod manifests to %q\n", pathMgr.TempManifestDir())
	err = controlplanephase.CreateInitStaticPodManifestFiles(pathMgr.TempManifestDir(), cfg)
	if err != nil {
		return errors.Wrap(err, "error creating init static pod manifest files")
	}
	for _, component := range constants.MasterComponents {
		if err = upgradeComponent(component, waiter, pathMgr, cfg, beforePodHashMap[component], recoverManifests, isTLSUpgrade); err != nil {
			return err
		}
	}
	return pathMgr.CleanupDirs()
}
func rollbackOldManifests(oldManifests map[string]string, origErr error, pathMgr StaticPodPathManager, restoreEtcd bool) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	errs := []error{origErr}
	for component, backupPath := range oldManifests {
		if component == constants.Etcd && !restoreEtcd {
			continue
		}
		realManifestPath := pathMgr.RealManifestPath(component)
		err := pathMgr.MoveFile(backupPath, realManifestPath)
		if err != nil {
			errs = append(errs, err)
		}
	}
	return errors.New("couldn't upgrade control plane. kubeadm has tried to recover everything into the earlier state. Errors faced")
}
func rollbackEtcdData(cfg *kubeadmapi.InitConfiguration, pathMgr StaticPodPathManager) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	backupEtcdDir := pathMgr.BackupEtcdDir()
	runningEtcdDir := cfg.Etcd.Local.DataDir
	if err := util.CopyDir(backupEtcdDir, runningEtcdDir); err != nil {
		return errors.Wrapf(err, "couldn't recover etcd database with error, the location of etcd backup: %s ", backupEtcdDir)
	}
	return nil
}
func renewCerts(cfg *kubeadmapi.InitConfiguration, component string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if cfg.Etcd.Local != nil {
		if component == constants.Etcd || component == constants.KubeAPIServer {
			caCert, caKey, err := certsphase.LoadCertificateAuthority(cfg.CertificatesDir, certsphase.KubeadmCertEtcdCA.BaseName)
			if err != nil {
				return errors.Wrapf(err, "failed to upgrade the %s CA certificate and key", constants.Etcd)
			}
			renewer := renewal.NewFileRenewal(caCert, caKey)
			if component == constants.Etcd {
				for _, cert := range []*certsphase.KubeadmCert{&certsphase.KubeadmCertEtcdServer, &certsphase.KubeadmCertEtcdPeer, &certsphase.KubeadmCertEtcdHealthcheck} {
					if err := renewal.RenewExistingCert(cfg.CertificatesDir, cert.BaseName, renewer); err != nil {
						return errors.Wrapf(err, "failed to renew %s certificate and key", cert.Name)
					}
				}
			}
			if component == constants.KubeAPIServer {
				cert := certsphase.KubeadmCertEtcdAPIClient
				if err := renewal.RenewExistingCert(cfg.CertificatesDir, cert.BaseName, renewer); err != nil {
					return errors.Wrapf(err, "failed to renew %s certificate and key", cert.Name)
				}
			}
		}
	}
	return nil
}
