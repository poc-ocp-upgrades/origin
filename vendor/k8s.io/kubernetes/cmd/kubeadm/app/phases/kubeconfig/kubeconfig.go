package kubeconfig

import (
	"bytes"
	"crypto/rsa"
	"crypto/x509"
	"fmt"
	goformat "fmt"
	"github.com/pkg/errors"
	"io"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
	certutil "k8s.io/client-go/util/cert"
	"k8s.io/klog"
	kubeadmapi "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm"
	kubeadmconstants "k8s.io/kubernetes/cmd/kubeadm/app/constants"
	kubeadmutil "k8s.io/kubernetes/cmd/kubeadm/app/util"
	kubeconfigutil "k8s.io/kubernetes/cmd/kubeadm/app/util/kubeconfig"
	"k8s.io/kubernetes/cmd/kubeadm/app/util/pkiutil"
	"os"
	goos "os"
	"path/filepath"
	godefaultruntime "runtime"
	gotime "time"
)

type clientCertAuth struct {
	CAKey         *rsa.PrivateKey
	Organizations []string
}
type tokenAuth struct{ Token string }
type kubeConfigSpec struct {
	CACert         *x509.Certificate
	APIServer      string
	ClientName     string
	TokenAuth      *tokenAuth
	ClientCertAuth *clientCertAuth
}

func CreateInitKubeConfigFiles(outDir string, cfg *kubeadmapi.InitConfiguration) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(1).Infoln("creating all kubeconfig files")
	return createKubeConfigFiles(outDir, cfg, kubeadmconstants.AdminKubeConfigFileName, kubeadmconstants.KubeletKubeConfigFileName, kubeadmconstants.ControllerManagerKubeConfigFileName, kubeadmconstants.SchedulerKubeConfigFileName)
}
func CreateJoinControlPlaneKubeConfigFiles(outDir string, cfg *kubeadmapi.InitConfiguration) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return createKubeConfigFiles(outDir, cfg, kubeadmconstants.AdminKubeConfigFileName, kubeadmconstants.ControllerManagerKubeConfigFileName, kubeadmconstants.SchedulerKubeConfigFileName)
}
func CreateKubeConfigFile(kubeConfigFileName string, outDir string, cfg *kubeadmapi.InitConfiguration) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(1).Infof("creating kubeconfig file for %s", kubeConfigFileName)
	return createKubeConfigFiles(outDir, cfg, kubeConfigFileName)
}
func createKubeConfigFiles(outDir string, cfg *kubeadmapi.InitConfiguration, kubeConfigFileNames ...string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	specs, err := getKubeConfigSpecs(cfg)
	if err != nil {
		return err
	}
	for _, kubeConfigFileName := range kubeConfigFileNames {
		spec, exists := specs[kubeConfigFileName]
		if !exists {
			return errors.Errorf("couldn't retrive KubeConfigSpec for %s", kubeConfigFileName)
		}
		config, err := buildKubeConfigFromSpec(spec, cfg.ClusterName)
		if err != nil {
			return err
		}
		if err = createKubeConfigFileIfNotExists(outDir, kubeConfigFileName, config); err != nil {
			return err
		}
	}
	return nil
}
func getKubeConfigSpecs(cfg *kubeadmapi.InitConfiguration) (map[string]*kubeConfigSpec, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	caCert, caKey, err := pkiutil.TryLoadCertAndKeyFromDisk(cfg.CertificatesDir, kubeadmconstants.CACertAndKeyBaseName)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't create a kubeconfig; the CA files couldn't be loaded")
	}
	masterEndpoint, err := kubeadmutil.GetMasterEndpoint(cfg)
	if err != nil {
		return nil, err
	}
	var kubeConfigSpec = map[string]*kubeConfigSpec{kubeadmconstants.AdminKubeConfigFileName: {CACert: caCert, APIServer: masterEndpoint, ClientName: "kubernetes-admin", ClientCertAuth: &clientCertAuth{CAKey: caKey, Organizations: []string{kubeadmconstants.MastersGroup}}}, kubeadmconstants.KubeletKubeConfigFileName: {CACert: caCert, APIServer: masterEndpoint, ClientName: fmt.Sprintf("%s%s", kubeadmconstants.NodesUserPrefix, cfg.NodeRegistration.Name), ClientCertAuth: &clientCertAuth{CAKey: caKey, Organizations: []string{kubeadmconstants.NodesGroup}}}, kubeadmconstants.ControllerManagerKubeConfigFileName: {CACert: caCert, APIServer: masterEndpoint, ClientName: kubeadmconstants.ControllerManagerUser, ClientCertAuth: &clientCertAuth{CAKey: caKey}}, kubeadmconstants.SchedulerKubeConfigFileName: {CACert: caCert, APIServer: masterEndpoint, ClientName: kubeadmconstants.SchedulerUser, ClientCertAuth: &clientCertAuth{CAKey: caKey}}}
	return kubeConfigSpec, nil
}
func buildKubeConfigFromSpec(spec *kubeConfigSpec, clustername string) (*clientcmdapi.Config, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if spec.TokenAuth != nil {
		return kubeconfigutil.CreateWithToken(spec.APIServer, clustername, spec.ClientName, certutil.EncodeCertPEM(spec.CACert), spec.TokenAuth.Token), nil
	}
	clientCertConfig := certutil.Config{CommonName: spec.ClientName, Organization: spec.ClientCertAuth.Organizations, Usages: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth}}
	clientCert, clientKey, err := pkiutil.NewCertAndKey(spec.CACert, spec.ClientCertAuth.CAKey, &clientCertConfig)
	if err != nil {
		return nil, errors.Wrapf(err, "failure while creating %s client certificate", spec.ClientName)
	}
	return kubeconfigutil.CreateWithCerts(spec.APIServer, clustername, spec.ClientName, certutil.EncodeCertPEM(spec.CACert), certutil.EncodePrivateKeyPEM(clientKey), certutil.EncodeCertPEM(clientCert)), nil
}
func validateKubeConfig(outDir, filename string, config *clientcmdapi.Config) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	kubeConfigFilePath := filepath.Join(outDir, filename)
	if _, err := os.Stat(kubeConfigFilePath); err != nil {
		return err
	}
	currentConfig, err := clientcmd.LoadFromFile(kubeConfigFilePath)
	if err != nil {
		return errors.Wrapf(err, "failed to load kubeconfig file %s that already exists on disk", kubeConfigFilePath)
	}
	expectedCtx := config.CurrentContext
	expectedCluster := config.Contexts[expectedCtx].Cluster
	currentCtx := currentConfig.CurrentContext
	currentCluster := currentConfig.Contexts[currentCtx].Cluster
	if !bytes.Equal(currentConfig.Clusters[currentCluster].CertificateAuthorityData, config.Clusters[expectedCluster].CertificateAuthorityData) {
		return errors.Errorf("a kubeconfig file %q exists already but has got the wrong CA cert", kubeConfigFilePath)
	}
	if currentConfig.Clusters[currentCluster].Server != config.Clusters[expectedCluster].Server {
		return errors.Errorf("a kubeconfig file %q exists already but has got the wrong API Server URL", kubeConfigFilePath)
	}
	return nil
}
func createKubeConfigFileIfNotExists(outDir, filename string, config *clientcmdapi.Config) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	kubeConfigFilePath := filepath.Join(outDir, filename)
	err := validateKubeConfig(outDir, filename, config)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
		fmt.Printf("[kubeconfig] Writing %q kubeconfig file\n", filename)
		err = kubeconfigutil.WriteToDisk(kubeConfigFilePath, config)
		if err != nil {
			return errors.Wrapf(err, "failed to save kubeconfig file %q on disk", kubeConfigFilePath)
		}
		return nil
	}
	fmt.Printf("[kubeconfig] Using existing up-to-date kubeconfig file: %q\n", kubeConfigFilePath)
	return nil
}
func WriteKubeConfigWithClientCert(out io.Writer, cfg *kubeadmapi.InitConfiguration, clientName string, organizations []string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	caCert, caKey, err := pkiutil.TryLoadCertAndKeyFromDisk(cfg.CertificatesDir, kubeadmconstants.CACertAndKeyBaseName)
	if err != nil {
		return errors.Wrap(err, "couldn't create a kubeconfig; the CA files couldn't be loaded")
	}
	masterEndpoint, err := kubeadmutil.GetMasterEndpoint(cfg)
	if err != nil {
		return err
	}
	spec := &kubeConfigSpec{ClientName: clientName, APIServer: masterEndpoint, CACert: caCert, ClientCertAuth: &clientCertAuth{CAKey: caKey, Organizations: organizations}}
	return writeKubeConfigFromSpec(out, spec, cfg.ClusterName)
}
func WriteKubeConfigWithToken(out io.Writer, cfg *kubeadmapi.InitConfiguration, clientName, token string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	caCert, _, err := pkiutil.TryLoadCertAndKeyFromDisk(cfg.CertificatesDir, kubeadmconstants.CACertAndKeyBaseName)
	if err != nil {
		return errors.Wrap(err, "couldn't create a kubeconfig; the CA files couldn't be loaded")
	}
	masterEndpoint, err := kubeadmutil.GetMasterEndpoint(cfg)
	if err != nil {
		return err
	}
	spec := &kubeConfigSpec{ClientName: clientName, APIServer: masterEndpoint, CACert: caCert, TokenAuth: &tokenAuth{Token: token}}
	return writeKubeConfigFromSpec(out, spec, cfg.ClusterName)
}
func writeKubeConfigFromSpec(out io.Writer, spec *kubeConfigSpec, clustername string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	config, err := buildKubeConfigFromSpec(spec, clustername)
	if err != nil {
		return err
	}
	configBytes, err := clientcmd.Write(*config)
	if err != nil {
		return errors.Wrap(err, "failure while serializing admin kubeconfig")
	}
	fmt.Fprintln(out, string(configBytes))
	return nil
}
func ValidateKubeconfigsForExternalCA(outDir string, cfg *kubeadmapi.InitConfiguration) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	kubeConfigFileNames := []string{kubeadmconstants.AdminKubeConfigFileName, kubeadmconstants.KubeletKubeConfigFileName, kubeadmconstants.ControllerManagerKubeConfigFileName, kubeadmconstants.SchedulerKubeConfigFileName}
	specs, err := getKubeConfigSpecs(cfg)
	if err != nil {
		return err
	}
	for _, kubeConfigFileName := range kubeConfigFileNames {
		spec, exists := specs[kubeConfigFileName]
		if !exists {
			return errors.Errorf("couldn't retrive KubeConfigSpec for %s", kubeConfigFileName)
		}
		kubeconfig, err := buildKubeConfigFromSpec(spec, cfg.ClusterName)
		if err != nil {
			return err
		}
		if err = validateKubeConfig(outDir, kubeConfigFileName, kubeconfig); err != nil {
			return err
		}
	}
	return nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
