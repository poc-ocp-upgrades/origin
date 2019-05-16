package upgrade

import (
	"fmt"
	pkgerrors "github.com/pkg/errors"
	"io/ioutil"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/apimachinery/pkg/util/version"
	clientset "k8s.io/client-go/kubernetes"
	certutil "k8s.io/client-go/util/cert"
	kubeadmapi "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm"
	kubeadmapiv1beta1 "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/v1beta1"
	kubeadmconstants "k8s.io/kubernetes/cmd/kubeadm/app/constants"
	"k8s.io/kubernetes/cmd/kubeadm/app/phases/addons/dns"
	"k8s.io/kubernetes/cmd/kubeadm/app/phases/addons/proxy"
	"k8s.io/kubernetes/cmd/kubeadm/app/phases/bootstraptoken/clusterinfo"
	nodebootstraptoken "k8s.io/kubernetes/cmd/kubeadm/app/phases/bootstraptoken/node"
	certsphase "k8s.io/kubernetes/cmd/kubeadm/app/phases/certs"
	kubeletphase "k8s.io/kubernetes/cmd/kubeadm/app/phases/kubelet"
	patchnodephase "k8s.io/kubernetes/cmd/kubeadm/app/phases/patchnode"
	"k8s.io/kubernetes/cmd/kubeadm/app/phases/uploadconfig"
	"k8s.io/kubernetes/cmd/kubeadm/app/util/apiclient"
	dryrunutil "k8s.io/kubernetes/cmd/kubeadm/app/util/dryrun"
	"os"
	"path/filepath"
	"time"
)

var expiry = 180 * 24 * time.Hour

func PerformPostUpgradeTasks(client clientset.Interface, cfg *kubeadmapi.InitConfiguration, newK8sVer *version.Version, dryRun bool) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	errs := []error{}
	if err := uploadconfig.UploadConfiguration(cfg, client); err != nil {
		errs = append(errs, err)
	}
	if err := kubeletphase.CreateConfigMap(cfg, client); err != nil {
		errs = append(errs, pkgerrors.Wrap(err, "error creating kubelet configuration ConfigMap"))
	}
	if err := writeKubeletConfigFiles(client, cfg, newK8sVer, dryRun); err != nil {
		errs = append(errs, err)
	}
	if err := patchnodephase.AnnotateCRISocket(client, cfg.NodeRegistration.Name, cfg.NodeRegistration.CRISocket); err != nil {
		errs = append(errs, pkgerrors.Wrap(err, "error uploading crisocket"))
	}
	if err := nodebootstraptoken.AllowBootstrapTokensToPostCSRs(client); err != nil {
		errs = append(errs, err)
	}
	if err := nodebootstraptoken.AutoApproveNodeBootstrapTokens(client); err != nil {
		errs = append(errs, err)
	}
	if err := nodebootstraptoken.AutoApproveNodeCertificateRotation(client); err != nil {
		errs = append(errs, err)
	}
	if err := clusterinfo.CreateClusterInfoRBACRules(client); err != nil {
		errs = append(errs, err)
	}
	if err := BackupAPIServerCertIfNeeded(cfg, dryRun); err != nil {
		errs = append(errs, err)
	}
	if err := dns.EnsureDNSAddon(cfg, client); err != nil {
		errs = append(errs, err)
	}
	if err := removeOldDNSDeploymentIfAnotherDNSIsUsed(cfg, client, dryRun); err != nil {
		errs = append(errs, err)
	}
	if err := proxy.EnsureProxyAddon(cfg, client); err != nil {
		errs = append(errs, err)
	}
	return errors.NewAggregate(errs)
}
func removeOldDNSDeploymentIfAnotherDNSIsUsed(cfg *kubeadmapi.InitConfiguration, client clientset.Interface, dryRun bool) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return apiclient.TryRunCommand(func() error {
		installedDeploymentName := kubeadmconstants.KubeDNSDeploymentName
		deploymentToDelete := kubeadmconstants.CoreDNSDeploymentName
		if cfg.DNS.Type == kubeadmapi.CoreDNS {
			installedDeploymentName = kubeadmconstants.CoreDNSDeploymentName
			deploymentToDelete = kubeadmconstants.KubeDNSDeploymentName
		}
		if !dryRun {
			dnsDeployment, err := client.AppsV1().Deployments(metav1.NamespaceSystem).Get(installedDeploymentName, metav1.GetOptions{})
			if err != nil {
				return err
			}
			if dnsDeployment.Status.ReadyReplicas == 0 {
				return pkgerrors.New("the DNS deployment isn't ready yet")
			}
		}
		err := apiclient.DeleteDeploymentForeground(client, metav1.NamespaceSystem, deploymentToDelete)
		if err != nil && !apierrors.IsNotFound(err) {
			return err
		}
		return nil
	}, 10)
}
func BackupAPIServerCertIfNeeded(cfg *kubeadmapi.InitConfiguration, dryRun bool) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	certAndKeyDir := kubeadmapiv1beta1.DefaultCertificatesDir
	shouldBackup, err := shouldBackupAPIServerCertAndKey(certAndKeyDir)
	if err != nil {
		return pkgerrors.Wrap(err, "[postupgrade] WARNING: failed to determine to backup kube-apiserver cert and key")
	}
	if !shouldBackup {
		return nil
	}
	if dryRun {
		fmt.Println("[postupgrade] Would rotate the API server certificate and key.")
		return nil
	}
	if err := backupAPIServerCertAndKey(certAndKeyDir); err != nil {
		fmt.Printf("[postupgrade] WARNING: failed to backup kube-apiserver cert and key: %v", err)
	}
	return certsphase.CreateCertAndKeyFilesWithCA(&certsphase.KubeadmCertAPIServer, &certsphase.KubeadmCertRootCA, cfg)
}
func writeKubeletConfigFiles(client clientset.Interface, cfg *kubeadmapi.InitConfiguration, newK8sVer *version.Version, dryRun bool) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	kubeletDir, err := getKubeletDir(dryRun)
	if err != nil {
		return err
	}
	errs := []error{}
	if err := kubeletphase.DownloadConfig(client, newK8sVer, kubeletDir); err != nil {
		if !(apierrors.IsNotFound(err) && dryRun) {
			errs = append(errs, pkgerrors.Wrap(err, "error downloading kubelet configuration from the ConfigMap"))
		}
	}
	if dryRun {
		dryrunutil.PrintDryRunFile(kubeadmconstants.KubeletConfigurationFileName, kubeletDir, kubeadmconstants.KubeletRunDirectory, os.Stdout)
	}
	envFilePath := filepath.Join(kubeadmconstants.KubeletRunDirectory, kubeadmconstants.KubeletEnvFileName)
	if _, err := os.Stat(envFilePath); os.IsNotExist(err) {
		if err := kubeletphase.WriteKubeletDynamicEnvFile(cfg, false, kubeletDir); err != nil {
			errs = append(errs, pkgerrors.Wrap(err, "error writing a dynamic environment file for the kubelet"))
		}
		if dryRun {
			dryrunutil.PrintDryRunFile(kubeadmconstants.KubeletEnvFileName, kubeletDir, kubeadmconstants.KubeletRunDirectory, os.Stdout)
		}
	}
	return errors.NewAggregate(errs)
}
func getKubeletDir(dryRun bool) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if dryRun {
		return ioutil.TempDir("", "kubeadm-upgrade-dryrun")
	}
	return kubeadmconstants.KubeletRunDirectory, nil
}
func backupAPIServerCertAndKey(certAndKeyDir string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	subDir := filepath.Join(certAndKeyDir, "expired")
	if err := os.Mkdir(subDir, 0766); err != nil {
		return pkgerrors.Wrapf(err, "failed to created backup directory %s", subDir)
	}
	filesToMove := map[string]string{filepath.Join(certAndKeyDir, kubeadmconstants.APIServerCertName): filepath.Join(subDir, kubeadmconstants.APIServerCertName), filepath.Join(certAndKeyDir, kubeadmconstants.APIServerKeyName): filepath.Join(subDir, kubeadmconstants.APIServerKeyName)}
	return moveFiles(filesToMove)
}
func moveFiles(files map[string]string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	filesToRecover := map[string]string{}
	for from, to := range files {
		if err := os.Rename(from, to); err != nil {
			return rollbackFiles(filesToRecover, err)
		}
		filesToRecover[to] = from
	}
	return nil
}
func rollbackFiles(files map[string]string, originalErr error) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	errs := []error{originalErr}
	for from, to := range files {
		if err := os.Rename(from, to); err != nil {
			errs = append(errs, err)
		}
	}
	return pkgerrors.Errorf("couldn't move these files: %v. Got errors: %v", files, errors.NewAggregate(errs))
}
func shouldBackupAPIServerCertAndKey(certAndKeyDir string) (bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	apiServerCert := filepath.Join(certAndKeyDir, kubeadmconstants.APIServerCertName)
	certs, err := certutil.CertsFromFile(apiServerCert)
	if err != nil {
		return false, pkgerrors.Wrapf(err, "couldn't load the certificate file %s", apiServerCert)
	}
	if len(certs) == 0 {
		return false, pkgerrors.New("no certificate data found")
	}
	if time.Now().Sub(certs[0].NotBefore) > expiry {
		return true, nil
	}
	return false, nil
}
