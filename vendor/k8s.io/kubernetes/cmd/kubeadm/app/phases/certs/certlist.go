package certs

import (
	"crypto/rsa"
	"crypto/x509"
	goformat "fmt"
	"github.com/pkg/errors"
	certutil "k8s.io/client-go/util/cert"
	kubeadmapi "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm"
	kubeadmconstants "k8s.io/kubernetes/cmd/kubeadm/app/constants"
	"k8s.io/kubernetes/cmd/kubeadm/app/util/pkiutil"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

type configMutatorsFunc func(*kubeadmapi.InitConfiguration, *certutil.Config) error
type KubeadmCert struct {
	Name           string
	LongName       string
	BaseName       string
	CAName         string
	configMutators []configMutatorsFunc
	config         certutil.Config
}

func (k *KubeadmCert) GetConfig(ic *kubeadmapi.InitConfiguration) (*certutil.Config, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for _, f := range k.configMutators {
		if err := f(ic, &k.config); err != nil {
			return nil, err
		}
	}
	return &k.config, nil
}
func (k *KubeadmCert) CreateFromCA(ic *kubeadmapi.InitConfiguration, caCert *x509.Certificate, caKey *rsa.PrivateKey) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	cfg, err := k.GetConfig(ic)
	if err != nil {
		return errors.Wrapf(err, "couldn't create %q certificate", k.Name)
	}
	cert, key, err := pkiutil.NewCertAndKey(caCert, caKey, cfg)
	if err != nil {
		return err
	}
	err = writeCertificateFilesIfNotExist(ic.CertificatesDir, k.BaseName, caCert, cert, key)
	if err != nil {
		return errors.Wrapf(err, "failed to write certificate %q", k.Name)
	}
	return nil
}
func (k *KubeadmCert) CreateAsCA(ic *kubeadmapi.InitConfiguration) (*x509.Certificate, *rsa.PrivateKey, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	cfg, err := k.GetConfig(ic)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "couldn't get configuration for %q CA certificate", k.Name)
	}
	caCert, caKey, err := NewCACertAndKey(cfg)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "couldn't generate %q CA certificate", k.Name)
	}
	err = writeCertificateAuthorithyFilesIfNotExist(ic.CertificatesDir, k.BaseName, caCert, caKey)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "couldn't write out %q CA certificate", k.Name)
	}
	return caCert, caKey, nil
}

type CertificateTree map[*KubeadmCert]Certificates

func (t CertificateTree) CreateTree(ic *kubeadmapi.InitConfiguration) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for ca, leaves := range t {
		cfg, err := ca.GetConfig(ic)
		if err != nil {
			return err
		}
		var caKey *rsa.PrivateKey
		caCert, err := pkiutil.TryLoadCertFromDisk(ic.CertificatesDir, ca.BaseName)
		if err == nil {
			if !caCert.IsCA {
				return errors.Errorf("certificate %q is not a CA", ca.Name)
			}
			caKey, err = pkiutil.TryLoadKeyFromDisk(ic.CertificatesDir, ca.BaseName)
			if err != nil {
				for _, leaf := range leaves {
					cl := certKeyLocation{pkiDir: ic.CertificatesDir, baseName: leaf.BaseName, uxName: leaf.Name}
					if err := validateSignedCertWithCA(cl, caCert); err != nil {
						return errors.Wrapf(err, "could not load expected certificate %q or validate the existence of key %q for it", leaf.Name, ca.Name)
					}
				}
				continue
			}
		} else {
			caCert, caKey, err = NewCACertAndKey(cfg)
			if err != nil {
				return err
			}
			err = writeCertificateAuthorithyFilesIfNotExist(ic.CertificatesDir, ca.BaseName, caCert, caKey)
			if err != nil {
				return err
			}
		}
		for _, leaf := range leaves {
			if err := leaf.CreateFromCA(ic, caCert, caKey); err != nil {
				return err
			}
		}
	}
	return nil
}

type CertificateMap map[string]*KubeadmCert

func (m CertificateMap) CertTree() (CertificateTree, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	caMap := make(CertificateTree)
	for _, cert := range m {
		if cert.CAName == "" {
			if _, ok := caMap[cert]; !ok {
				caMap[cert] = []*KubeadmCert{}
			}
		} else {
			ca, ok := m[cert.CAName]
			if !ok {
				return nil, errors.Errorf("certificate %q references unknown CA %q", cert.Name, cert.CAName)
			}
			caMap[ca] = append(caMap[ca], cert)
		}
	}
	return caMap, nil
}

type Certificates []*KubeadmCert

func (c Certificates) AsMap() CertificateMap {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	certMap := make(map[string]*KubeadmCert)
	for _, cert := range c {
		certMap[cert.Name] = cert
	}
	return certMap
}
func GetDefaultCertList() Certificates {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return Certificates{&KubeadmCertRootCA, &KubeadmCertAPIServer, &KubeadmCertKubeletClient, &KubeadmCertFrontProxyCA, &KubeadmCertFrontProxyClient, &KubeadmCertEtcdCA, &KubeadmCertEtcdServer, &KubeadmCertEtcdPeer, &KubeadmCertEtcdHealthcheck, &KubeadmCertEtcdAPIClient}
}
func GetCertsWithoutEtcd() Certificates {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return Certificates{&KubeadmCertRootCA, &KubeadmCertAPIServer, &KubeadmCertKubeletClient, &KubeadmCertFrontProxyCA, &KubeadmCertFrontProxyClient}
}

var (
	KubeadmCertRootCA           = KubeadmCert{Name: "ca", LongName: "self-signed Kubernetes CA to provision identities for other Kubernetes components", BaseName: kubeadmconstants.CACertAndKeyBaseName, config: certutil.Config{CommonName: "kubernetes"}}
	KubeadmCertAPIServer        = KubeadmCert{Name: "apiserver", LongName: "certificate for serving the Kubernetes API", BaseName: kubeadmconstants.APIServerCertAndKeyBaseName, CAName: "ca", config: certutil.Config{CommonName: kubeadmconstants.APIServerCertCommonName, Usages: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth}}, configMutators: []configMutatorsFunc{makeAltNamesMutator(pkiutil.GetAPIServerAltNames)}}
	KubeadmCertKubeletClient    = KubeadmCert{Name: "apiserver-kubelet-client", LongName: "Client certificate for the API server to connect to kubelet", BaseName: kubeadmconstants.APIServerKubeletClientCertAndKeyBaseName, CAName: "ca", config: certutil.Config{CommonName: kubeadmconstants.APIServerKubeletClientCertCommonName, Organization: []string{kubeadmconstants.MastersGroup}, Usages: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth}}}
	KubeadmCertFrontProxyCA     = KubeadmCert{Name: "front-proxy-ca", LongName: "self-signed CA to provision identities for front proxy", BaseName: kubeadmconstants.FrontProxyCACertAndKeyBaseName, config: certutil.Config{CommonName: "front-proxy-ca"}}
	KubeadmCertFrontProxyClient = KubeadmCert{Name: "front-proxy-client", BaseName: kubeadmconstants.FrontProxyClientCertAndKeyBaseName, LongName: "client for the front proxy", CAName: "front-proxy-ca", config: certutil.Config{CommonName: kubeadmconstants.FrontProxyClientCertCommonName, Usages: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth}}}
	KubeadmCertEtcdCA           = KubeadmCert{Name: "etcd-ca", LongName: "self-signed CA to provision identities for etcd", BaseName: kubeadmconstants.EtcdCACertAndKeyBaseName, config: certutil.Config{CommonName: "etcd-ca"}}
	KubeadmCertEtcdServer       = KubeadmCert{Name: "etcd-server", LongName: "certificate for serving etcd", BaseName: kubeadmconstants.EtcdServerCertAndKeyBaseName, CAName: "etcd-ca", config: certutil.Config{Usages: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth}}, configMutators: []configMutatorsFunc{makeAltNamesMutator(pkiutil.GetEtcdAltNames), setCommonNameToNodeName()}}
	KubeadmCertEtcdPeer         = KubeadmCert{Name: "etcd-peer", LongName: "credentials for etcd nodes to communicate with each other", BaseName: kubeadmconstants.EtcdPeerCertAndKeyBaseName, CAName: "etcd-ca", config: certutil.Config{Usages: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth}}, configMutators: []configMutatorsFunc{makeAltNamesMutator(pkiutil.GetEtcdPeerAltNames), setCommonNameToNodeName()}}
	KubeadmCertEtcdHealthcheck  = KubeadmCert{Name: "etcd-healthcheck-client", LongName: "client certificate for liveness probes to healtcheck etcd", BaseName: kubeadmconstants.EtcdHealthcheckClientCertAndKeyBaseName, CAName: "etcd-ca", config: certutil.Config{CommonName: kubeadmconstants.EtcdHealthcheckClientCertCommonName, Organization: []string{kubeadmconstants.MastersGroup}, Usages: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth}}}
	KubeadmCertEtcdAPIClient    = KubeadmCert{Name: "apiserver-etcd-client", LongName: "client apiserver uses to access etcd", BaseName: kubeadmconstants.APIServerEtcdClientCertAndKeyBaseName, CAName: "etcd-ca", config: certutil.Config{CommonName: kubeadmconstants.APIServerEtcdClientCertCommonName, Organization: []string{kubeadmconstants.MastersGroup}, Usages: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth}}}
)

func makeAltNamesMutator(f func(*kubeadmapi.InitConfiguration) (*certutil.AltNames, error)) configMutatorsFunc {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return func(mc *kubeadmapi.InitConfiguration, cc *certutil.Config) error {
		altNames, err := f(mc)
		if err != nil {
			return err
		}
		cc.AltNames = *altNames
		return nil
	}
}
func setCommonNameToNodeName() configMutatorsFunc {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return func(mc *kubeadmapi.InitConfiguration, cc *certutil.Config) error {
		cc.CommonName = mc.NodeRegistration.Name
		return nil
	}
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
