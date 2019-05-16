package certs

import (
	"crypto/rsa"
	"crypto/x509"
	"fmt"
	"github.com/pkg/errors"
	certutil "k8s.io/client-go/util/cert"
	"k8s.io/klog"
	kubeadmapi "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm"
	kubeadmconstants "k8s.io/kubernetes/cmd/kubeadm/app/constants"
	"k8s.io/kubernetes/cmd/kubeadm/app/util/pkiutil"
	"os"
	"path/filepath"
)

func CreatePKIAssets(cfg *kubeadmapi.InitConfiguration) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(1).Infoln("creating PKI assets")
	var certList Certificates
	if cfg.Etcd.Local == nil {
		certList = GetCertsWithoutEtcd()
	} else {
		certList = GetDefaultCertList()
	}
	certTree, err := certList.AsMap().CertTree()
	if err != nil {
		return err
	}
	if err := certTree.CreateTree(cfg); err != nil {
		return errors.Wrap(err, "error creating PKI assets")
	}
	fmt.Printf("[certs] valid certificates and keys now exist in %q\n", cfg.CertificatesDir)
	if err := CreateServiceAccountKeyAndPublicKeyFiles(cfg); err != nil {
		return err
	}
	return nil
}
func CreateServiceAccountKeyAndPublicKeyFiles(cfg *kubeadmapi.InitConfiguration) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(1).Infoln("creating a new public/private key files for signing service account users")
	saSigningKey, err := NewServiceAccountSigningKey()
	if err != nil {
		return err
	}
	return writeKeyFilesIfNotExist(cfg.CertificatesDir, kubeadmconstants.ServiceAccountKeyBaseName, saSigningKey)
}
func NewServiceAccountSigningKey() (*rsa.PrivateKey, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	saSigningKey, err := certutil.NewPrivateKey()
	if err != nil {
		return nil, errors.Wrap(err, "failure while creating service account token signing key")
	}
	return saSigningKey, nil
}
func NewCACertAndKey(certSpec *certutil.Config) (*x509.Certificate, *rsa.PrivateKey, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	caCert, caKey, err := pkiutil.NewCertificateAuthority(certSpec)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failure while generating CA certificate and key")
	}
	return caCert, caKey, nil
}
func CreateCACertAndKeyFiles(certSpec *KubeadmCert, cfg *kubeadmapi.InitConfiguration) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if certSpec.CAName != "" {
		return errors.Errorf("this function should only be used for CAs, but cert %s has CA %s", certSpec.Name, certSpec.CAName)
	}
	klog.V(1).Infof("creating a new certificate authority for %s", certSpec.Name)
	certConfig, err := certSpec.GetConfig(cfg)
	if err != nil {
		return err
	}
	caCert, caKey, err := NewCACertAndKey(certConfig)
	if err != nil {
		return err
	}
	return writeCertificateAuthorithyFilesIfNotExist(cfg.CertificatesDir, certSpec.BaseName, caCert, caKey)
}
func NewCSR(certSpec *KubeadmCert, cfg *kubeadmapi.InitConfiguration) (*x509.CertificateRequest, *rsa.PrivateKey, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	certConfig, err := certSpec.GetConfig(cfg)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to retrieve cert configuration: %v", err)
	}
	return pkiutil.NewCSRAndKey(certConfig)
}
func CreateCSR(certSpec *KubeadmCert, cfg *kubeadmapi.InitConfiguration, path string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	csr, key, err := NewCSR(certSpec, cfg)
	if err != nil {
		return err
	}
	return writeCSRFilesIfNotExist(path, certSpec.BaseName, csr, key)
}
func CreateCertAndKeyFilesWithCA(certSpec *KubeadmCert, caCertSpec *KubeadmCert, cfg *kubeadmapi.InitConfiguration) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if certSpec.CAName != caCertSpec.Name {
		return errors.Errorf("expected CAname for %s to be %q, but was %s", certSpec.Name, certSpec.CAName, caCertSpec.Name)
	}
	caCert, caKey, err := LoadCertificateAuthority(cfg.CertificatesDir, caCertSpec.BaseName)
	if err != nil {
		return errors.Wrapf(err, "couldn't load CA certificate %s", caCertSpec.Name)
	}
	return certSpec.CreateFromCA(cfg, caCert, caKey)
}
func LoadCertificateAuthority(pkiDir string, baseName string) (*x509.Certificate, *rsa.PrivateKey, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if !pkiutil.CertOrKeyExist(pkiDir, baseName) {
		return nil, nil, errors.Errorf("couldn't load %s certificate authority from %s", baseName, pkiDir)
	}
	caCert, caKey, err := pkiutil.TryLoadCertAndKeyFromDisk(pkiDir, baseName)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "failure loading %s certificate authority", baseName)
	}
	if !caCert.IsCA {
		return nil, nil, errors.Errorf("%s certificate is not a certificate authority", baseName)
	}
	return caCert, caKey, nil
}
func writeCertificateAuthorithyFilesIfNotExist(pkiDir string, baseName string, caCert *x509.Certificate, caKey *rsa.PrivateKey) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if pkiutil.CertOrKeyExist(pkiDir, baseName) {
		caCert, _, err := pkiutil.TryLoadCertAndKeyFromDisk(pkiDir, baseName)
		if err != nil {
			return errors.Wrapf(err, "failure loading %s certificate", baseName)
		}
		if !caCert.IsCA {
			return errors.Errorf("certificate %s is not a CA", baseName)
		}
		fmt.Printf("[certs] Using the existing %q certificate and key\n", baseName)
	} else {
		fmt.Printf("[certs] Generating %q certificate and key\n", baseName)
		if err := pkiutil.WriteCertAndKey(pkiDir, baseName, caCert, caKey); err != nil {
			return errors.Wrapf(err, "failure while saving %s certificate and key", baseName)
		}
	}
	return nil
}
func writeCertificateFilesIfNotExist(pkiDir string, baseName string, signingCert *x509.Certificate, cert *x509.Certificate, key *rsa.PrivateKey) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if pkiutil.CertOrKeyExist(pkiDir, baseName) {
		signedCert, _, err := pkiutil.TryLoadCertAndKeyFromDisk(pkiDir, baseName)
		if err != nil {
			return errors.Wrapf(err, "failure loading %s certificate", baseName)
		}
		if err := signedCert.CheckSignatureFrom(signingCert); err != nil {
			return errors.Errorf("certificate %s is not signed by corresponding CA", baseName)
		}
		fmt.Printf("[certs] Using the existing %q certificate and key\n", baseName)
	} else {
		fmt.Printf("[certs] Generating %q certificate and key\n", baseName)
		if err := pkiutil.WriteCertAndKey(pkiDir, baseName, cert, key); err != nil {
			return errors.Wrapf(err, "failure while saving %s certificate and key", baseName)
		}
		if pkiutil.HasServerAuth(cert) {
			fmt.Printf("[certs] %s serving cert is signed for DNS names %v and IPs %v\n", baseName, cert.DNSNames, cert.IPAddresses)
		}
	}
	return nil
}
func writeKeyFilesIfNotExist(pkiDir string, baseName string, key *rsa.PrivateKey) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if pkiutil.CertOrKeyExist(pkiDir, baseName) {
		_, err := pkiutil.TryLoadKeyFromDisk(pkiDir, baseName)
		if err != nil {
			return errors.Wrapf(err, "%s key existed but it could not be loaded properly", baseName)
		}
		fmt.Printf("[certs] Using the existing %q key\n", baseName)
	} else {
		fmt.Printf("[certs] Generating %q key and public key\n", baseName)
		if err := pkiutil.WriteKey(pkiDir, baseName, key); err != nil {
			return errors.Wrapf(err, "failure while saving %s key", baseName)
		}
		if err := pkiutil.WritePublicKey(pkiDir, baseName, &key.PublicKey); err != nil {
			return errors.Wrapf(err, "failure while saving %s public key", baseName)
		}
	}
	return nil
}
func writeCSRFilesIfNotExist(csrDir string, baseName string, csr *x509.CertificateRequest, key *rsa.PrivateKey) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if pkiutil.CSROrKeyExist(csrDir, baseName) {
		_, _, err := pkiutil.TryLoadCSRAndKeyFromDisk(csrDir, baseName)
		if err != nil {
			return errors.Wrapf(err, "%s CSR existed but it could not be loaded properly", baseName)
		}
		fmt.Printf("[certs] Using the existing %q CSR\n", baseName)
	} else {
		fmt.Printf("[certs] Generating %q key and CSR\n", baseName)
		if err := pkiutil.WriteKey(csrDir, baseName, key); err != nil {
			return errors.Wrapf(err, "failure while saving %s key", baseName)
		}
		if err := pkiutil.WriteCSR(csrDir, baseName, csr); err != nil {
			return errors.Wrapf(err, "failure while saving %s CSR", baseName)
		}
	}
	return nil
}

type certKeyLocation struct {
	pkiDir     string
	caBaseName string
	baseName   string
	uxName     string
}

func SharedCertificateExists(cfg *kubeadmapi.InitConfiguration) (bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if err := validateCACertAndKey(certKeyLocation{cfg.CertificatesDir, kubeadmconstants.CACertAndKeyBaseName, "", "CA"}); err != nil {
		return false, err
	}
	if err := validatePrivatePublicKey(certKeyLocation{cfg.CertificatesDir, "", kubeadmconstants.ServiceAccountKeyBaseName, "service account"}); err != nil {
		return false, err
	}
	if err := validateCACertAndKey(certKeyLocation{cfg.CertificatesDir, kubeadmconstants.FrontProxyCACertAndKeyBaseName, "", "front-proxy CA"}); err != nil {
		return false, err
	}
	if cfg.Etcd.External == nil {
		if err := validateCACertAndKey(certKeyLocation{cfg.CertificatesDir, kubeadmconstants.EtcdCACertAndKeyBaseName, "", "etcd CA"}); err != nil {
			return false, err
		}
	}
	return true, nil
}
func UsingExternalCA(cfg *kubeadmapi.InitConfiguration) (bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if err := validateCACert(certKeyLocation{cfg.CertificatesDir, kubeadmconstants.CACertAndKeyBaseName, "", "CA"}); err != nil {
		return false, err
	}
	caKeyPath := filepath.Join(cfg.CertificatesDir, kubeadmconstants.CAKeyName)
	if _, err := os.Stat(caKeyPath); !os.IsNotExist(err) {
		return false, errors.Errorf("%s exists", kubeadmconstants.CAKeyName)
	}
	if err := validateSignedCert(certKeyLocation{cfg.CertificatesDir, kubeadmconstants.CACertAndKeyBaseName, kubeadmconstants.APIServerCertAndKeyBaseName, "API server"}); err != nil {
		return false, err
	}
	if err := validateSignedCert(certKeyLocation{cfg.CertificatesDir, kubeadmconstants.CACertAndKeyBaseName, kubeadmconstants.APIServerKubeletClientCertAndKeyBaseName, "API server kubelet client"}); err != nil {
		return false, err
	}
	if err := validatePrivatePublicKey(certKeyLocation{cfg.CertificatesDir, "", kubeadmconstants.ServiceAccountKeyBaseName, "service account"}); err != nil {
		return false, err
	}
	if err := validateCACert(certKeyLocation{cfg.CertificatesDir, kubeadmconstants.FrontProxyCACertAndKeyBaseName, "", "front-proxy CA"}); err != nil {
		return false, err
	}
	frontProxyCAKeyPath := filepath.Join(cfg.CertificatesDir, kubeadmconstants.FrontProxyCAKeyName)
	if _, err := os.Stat(frontProxyCAKeyPath); !os.IsNotExist(err) {
		return false, errors.Errorf("%s exists", kubeadmconstants.FrontProxyCAKeyName)
	}
	if err := validateSignedCert(certKeyLocation{cfg.CertificatesDir, kubeadmconstants.FrontProxyCACertAndKeyBaseName, kubeadmconstants.FrontProxyClientCertAndKeyBaseName, "front-proxy client"}); err != nil {
		return false, err
	}
	return true, nil
}
func validateCACert(l certKeyLocation) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	caCert, err := pkiutil.TryLoadCertFromDisk(l.pkiDir, l.caBaseName)
	if err != nil {
		return errors.Wrapf(err, "failure loading certificate for %s", l.uxName)
	}
	if !caCert.IsCA {
		return errors.Errorf("certificate %s is not a CA", l.uxName)
	}
	return nil
}
func validateCACertAndKey(l certKeyLocation) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if err := validateCACert(l); err != nil {
		return err
	}
	_, err := pkiutil.TryLoadKeyFromDisk(l.pkiDir, l.caBaseName)
	if err != nil {
		return errors.Wrapf(err, "failure loading key for %s", l.uxName)
	}
	return nil
}
func validateSignedCert(l certKeyLocation) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	caCert, err := pkiutil.TryLoadCertFromDisk(l.pkiDir, l.caBaseName)
	if err != nil {
		return errors.Wrapf(err, "failure loading certificate authority for %s", l.uxName)
	}
	return validateSignedCertWithCA(l, caCert)
}
func validateSignedCertWithCA(l certKeyLocation, caCert *x509.Certificate) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	signedCert, _, err := pkiutil.TryLoadCertAndKeyFromDisk(l.pkiDir, l.baseName)
	if err != nil {
		return errors.Wrapf(err, "failure loading certificate for %s", l.uxName)
	}
	if err := signedCert.CheckSignatureFrom(caCert); err != nil {
		return errors.Wrapf(err, "certificate %s is not signed by corresponding CA", l.uxName)
	}
	return nil
}
func validatePrivatePublicKey(l certKeyLocation) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	_, _, err := pkiutil.TryLoadPrivatePublicKeyFromDisk(l.pkiDir, l.baseName)
	if err != nil {
		return errors.Wrapf(err, "failure loading key for %s", l.uxName)
	}
	return nil
}
