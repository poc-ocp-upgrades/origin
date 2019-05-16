package pkiutil

import (
	"crypto"
	cryptorand "crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	goformat "fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"k8s.io/apimachinery/pkg/util/validation"
	certutil "k8s.io/client-go/util/cert"
	kubeadmapi "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm"
	kubeadmconstants "k8s.io/kubernetes/cmd/kubeadm/app/constants"
	kubeadmutil "k8s.io/kubernetes/cmd/kubeadm/app/util"
	"k8s.io/kubernetes/pkg/registry/core/service/ipallocator"
	"net"
	"os"
	goos "os"
	"path/filepath"
	godefaultruntime "runtime"
	"time"
	gotime "time"
)

func NewCertificateAuthority(config *certutil.Config) (*x509.Certificate, *rsa.PrivateKey, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	key, err := certutil.NewPrivateKey()
	if err != nil {
		return nil, nil, errors.Wrap(err, "unable to create private key")
	}
	cert, err := certutil.NewSelfSignedCACert(*config, key)
	if err != nil {
		return nil, nil, errors.Wrap(err, "unable to create self-signed certificate")
	}
	return cert, key, nil
}
func NewCertAndKey(caCert *x509.Certificate, caKey *rsa.PrivateKey, config *certutil.Config) (*x509.Certificate, *rsa.PrivateKey, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	key, err := certutil.NewPrivateKey()
	if err != nil {
		return nil, nil, errors.Wrap(err, "unable to create private key")
	}
	cert, err := certutil.NewSignedCert(*config, key, caCert, caKey)
	if err != nil {
		return nil, nil, errors.Wrap(err, "unable to sign certificate")
	}
	return cert, key, nil
}
func NewCSRAndKey(config *certutil.Config) (*x509.CertificateRequest, *rsa.PrivateKey, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	key, err := certutil.NewPrivateKey()
	if err != nil {
		return nil, nil, errors.Wrap(err, "unable to create private key")
	}
	csr, err := NewCSR(*config, key)
	if err != nil {
		return nil, nil, errors.Wrap(err, "unable to generate CSR")
	}
	return csr, key, nil
}
func HasServerAuth(cert *x509.Certificate) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for i := range cert.ExtKeyUsage {
		if cert.ExtKeyUsage[i] == x509.ExtKeyUsageServerAuth {
			return true
		}
	}
	return false
}
func WriteCertAndKey(pkiPath string, name string, cert *x509.Certificate, key *rsa.PrivateKey) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if err := WriteKey(pkiPath, name, key); err != nil {
		return errors.Wrap(err, "couldn't write key")
	}
	return WriteCert(pkiPath, name, cert)
}
func WriteCert(pkiPath, name string, cert *x509.Certificate) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if cert == nil {
		return errors.New("certificate cannot be nil when writing to file")
	}
	certificatePath := pathForCert(pkiPath, name)
	if err := certutil.WriteCert(certificatePath, certutil.EncodeCertPEM(cert)); err != nil {
		return errors.Wrapf(err, "unable to write certificate to file %s", certificatePath)
	}
	return nil
}
func WriteKey(pkiPath, name string, key *rsa.PrivateKey) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if key == nil {
		return errors.New("private key cannot be nil when writing to file")
	}
	privateKeyPath := pathForKey(pkiPath, name)
	if err := certutil.WriteKey(privateKeyPath, certutil.EncodePrivateKeyPEM(key)); err != nil {
		return errors.Wrapf(err, "unable to write private key to file %s", privateKeyPath)
	}
	return nil
}
func WriteCSR(csrDir, name string, csr *x509.CertificateRequest) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if csr == nil {
		return errors.New("certificate request cannot be nil when writing to file")
	}
	csrPath := pathForCSR(csrDir, name)
	if err := os.MkdirAll(filepath.Dir(csrPath), os.FileMode(0755)); err != nil {
		return errors.Wrapf(err, "failed to make directory %s", filepath.Dir(csrPath))
	}
	if err := ioutil.WriteFile(csrPath, EncodeCSRPEM(csr), os.FileMode(0644)); err != nil {
		return errors.Wrapf(err, "unable to write CSR to file %s", csrPath)
	}
	return nil
}
func WritePublicKey(pkiPath, name string, key *rsa.PublicKey) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if key == nil {
		return errors.New("public key cannot be nil when writing to file")
	}
	publicKeyBytes, err := certutil.EncodePublicKeyPEM(key)
	if err != nil {
		return err
	}
	publicKeyPath := pathForPublicKey(pkiPath, name)
	if err := certutil.WriteKey(publicKeyPath, publicKeyBytes); err != nil {
		return errors.Wrapf(err, "unable to write public key to file %s", publicKeyPath)
	}
	return nil
}
func CertOrKeyExist(pkiPath, name string) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	certificatePath, privateKeyPath := PathsForCertAndKey(pkiPath, name)
	_, certErr := os.Stat(certificatePath)
	_, keyErr := os.Stat(privateKeyPath)
	if os.IsNotExist(certErr) && os.IsNotExist(keyErr) {
		return false
	}
	return true
}
func CSROrKeyExist(csrDir, name string) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	csrPath := pathForCSR(csrDir, name)
	keyPath := pathForKey(csrDir, name)
	_, csrErr := os.Stat(csrPath)
	_, keyErr := os.Stat(keyPath)
	return !(os.IsNotExist(csrErr) && os.IsNotExist(keyErr))
}
func TryLoadCertAndKeyFromDisk(pkiPath, name string) (*x509.Certificate, *rsa.PrivateKey, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	cert, err := TryLoadCertFromDisk(pkiPath, name)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to load certificate")
	}
	key, err := TryLoadKeyFromDisk(pkiPath, name)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to load key")
	}
	return cert, key, nil
}
func TryLoadCertFromDisk(pkiPath, name string) (*x509.Certificate, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	certificatePath := pathForCert(pkiPath, name)
	certs, err := certutil.CertsFromFile(certificatePath)
	if err != nil {
		return nil, errors.Wrapf(err, "couldn't load the certificate file %s", certificatePath)
	}
	cert := certs[0]
	now := time.Now()
	if now.Before(cert.NotBefore) {
		return nil, errors.New("the certificate is not valid yet")
	}
	if now.After(cert.NotAfter) {
		return nil, errors.New("the certificate has expired")
	}
	return cert, nil
}
func TryLoadKeyFromDisk(pkiPath, name string) (*rsa.PrivateKey, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	privateKeyPath := pathForKey(pkiPath, name)
	privKey, err := certutil.PrivateKeyFromFile(privateKeyPath)
	if err != nil {
		return nil, errors.Wrapf(err, "couldn't load the private key file %s", privateKeyPath)
	}
	var key *rsa.PrivateKey
	switch k := privKey.(type) {
	case *rsa.PrivateKey:
		key = k
	default:
		return nil, errors.Errorf("the private key file %s isn't in RSA format", privateKeyPath)
	}
	return key, nil
}
func TryLoadCSRAndKeyFromDisk(pkiPath, name string) (*x509.CertificateRequest, *rsa.PrivateKey, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	csrPath := pathForCSR(pkiPath, name)
	csr, err := CertificateRequestFromFile(csrPath)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "couldn't load the certificate request %s", csrPath)
	}
	key, err := TryLoadKeyFromDisk(pkiPath, name)
	if err != nil {
		return nil, nil, errors.Wrap(err, "couldn't load key file")
	}
	return csr, key, nil
}
func TryLoadPrivatePublicKeyFromDisk(pkiPath, name string) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	privateKeyPath := pathForKey(pkiPath, name)
	privKey, err := certutil.PrivateKeyFromFile(privateKeyPath)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "couldn't load the private key file %s", privateKeyPath)
	}
	publicKeyPath := pathForPublicKey(pkiPath, name)
	pubKeys, err := certutil.PublicKeysFromFile(publicKeyPath)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "couldn't load the public key file %s", publicKeyPath)
	}
	k, ok := privKey.(*rsa.PrivateKey)
	if !ok {
		return nil, nil, errors.Errorf("the private key file %s isn't in RSA format", privateKeyPath)
	}
	p := pubKeys[0].(*rsa.PublicKey)
	return k, p, nil
}
func PathsForCertAndKey(pkiPath, name string) (string, string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return pathForCert(pkiPath, name), pathForKey(pkiPath, name)
}
func pathForCert(pkiPath, name string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return filepath.Join(pkiPath, fmt.Sprintf("%s.crt", name))
}
func pathForKey(pkiPath, name string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return filepath.Join(pkiPath, fmt.Sprintf("%s.key", name))
}
func pathForPublicKey(pkiPath, name string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return filepath.Join(pkiPath, fmt.Sprintf("%s.pub", name))
}
func pathForCSR(pkiPath, name string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return filepath.Join(pkiPath, fmt.Sprintf("%s.csr", name))
}
func GetAPIServerAltNames(cfg *kubeadmapi.InitConfiguration) (*certutil.AltNames, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	advertiseAddress := net.ParseIP(cfg.LocalAPIEndpoint.AdvertiseAddress)
	if advertiseAddress == nil {
		return nil, errors.Errorf("error parsing LocalAPIEndpoint AdvertiseAddress %v: is not a valid textual representation of an IP address", cfg.LocalAPIEndpoint.AdvertiseAddress)
	}
	_, svcSubnet, err := net.ParseCIDR(cfg.Networking.ServiceSubnet)
	if err != nil {
		return nil, errors.Wrapf(err, "error parsing CIDR %q", cfg.Networking.ServiceSubnet)
	}
	internalAPIServerVirtualIP, err := ipallocator.GetIndexedIP(svcSubnet, 1)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get first IP address from the given CIDR (%s)", svcSubnet.String())
	}
	altNames := &certutil.AltNames{DNSNames: []string{cfg.NodeRegistration.Name, "kubernetes", "kubernetes.default", "kubernetes.default.svc", fmt.Sprintf("kubernetes.default.svc.%s", cfg.Networking.DNSDomain)}, IPs: []net.IP{internalAPIServerVirtualIP, advertiseAddress}}
	if len(cfg.ControlPlaneEndpoint) > 0 {
		if host, _, err := kubeadmutil.ParseHostPort(cfg.ControlPlaneEndpoint); err == nil {
			if ip := net.ParseIP(host); ip != nil {
				altNames.IPs = append(altNames.IPs, ip)
			} else {
				altNames.DNSNames = append(altNames.DNSNames, host)
			}
		} else {
			return nil, errors.Wrapf(err, "error parsing cluster controlPlaneEndpoint %q", cfg.ControlPlaneEndpoint)
		}
	}
	appendSANsToAltNames(altNames, cfg.APIServer.CertSANs, kubeadmconstants.APIServerCertName)
	return altNames, nil
}
func GetEtcdAltNames(cfg *kubeadmapi.InitConfiguration) (*certutil.AltNames, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	advertiseAddress := net.ParseIP(cfg.LocalAPIEndpoint.AdvertiseAddress)
	if advertiseAddress == nil {
		return nil, errors.Errorf("error parsing LocalAPIEndpoint AdvertiseAddress %q: is not a valid textual representation of an IP address", cfg.LocalAPIEndpoint.AdvertiseAddress)
	}
	altNames := &certutil.AltNames{DNSNames: []string{cfg.NodeRegistration.Name, "localhost"}, IPs: []net.IP{advertiseAddress, net.IPv4(127, 0, 0, 1), net.IPv6loopback}}
	if cfg.Etcd.Local != nil {
		appendSANsToAltNames(altNames, cfg.Etcd.Local.ServerCertSANs, kubeadmconstants.EtcdServerCertName)
	}
	return altNames, nil
}
func GetEtcdPeerAltNames(cfg *kubeadmapi.InitConfiguration) (*certutil.AltNames, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	advertiseAddress := net.ParseIP(cfg.LocalAPIEndpoint.AdvertiseAddress)
	if advertiseAddress == nil {
		return nil, errors.Errorf("error parsing LocalAPIEndpoint AdvertiseAddress %v: is not a valid textual representation of an IP address", cfg.LocalAPIEndpoint.AdvertiseAddress)
	}
	altNames := &certutil.AltNames{DNSNames: []string{cfg.NodeRegistration.Name, "localhost"}, IPs: []net.IP{advertiseAddress, net.IPv4(127, 0, 0, 1), net.IPv6loopback}}
	if cfg.Etcd.Local != nil {
		appendSANsToAltNames(altNames, cfg.Etcd.Local.PeerCertSANs, kubeadmconstants.EtcdPeerCertName)
	}
	return altNames, nil
}
func appendSANsToAltNames(altNames *certutil.AltNames, SANs []string, certName string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for _, altname := range SANs {
		if ip := net.ParseIP(altname); ip != nil {
			altNames.IPs = append(altNames.IPs, ip)
		} else if len(validation.IsDNS1123Subdomain(altname)) == 0 {
			altNames.DNSNames = append(altNames.DNSNames, altname)
		} else {
			fmt.Printf("[certificates] WARNING: '%s' was not added to the '%s' SAN, because it is not a valid IP or RFC-1123 compliant DNS entry\n", altname, certName)
		}
	}
}
func EncodeCSRPEM(csr *x509.CertificateRequest) []byte {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	block := pem.Block{Type: certutil.CertificateRequestBlockType, Bytes: csr.Raw}
	return pem.EncodeToMemory(&block)
}
func parseCSRPEM(pemCSR []byte) (*x509.CertificateRequest, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	block, _ := pem.Decode(pemCSR)
	if block == nil {
		return nil, fmt.Errorf("data doesn't contain a valid certificate request")
	}
	if block.Type != certutil.CertificateRequestBlockType {
		var block *pem.Block
		return nil, fmt.Errorf("expected block type %q, but PEM had type %v", certutil.CertificateRequestBlockType, block.Type)
	}
	return x509.ParseCertificateRequest(block.Bytes)
}
func CertificateRequestFromFile(file string) (*x509.CertificateRequest, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	pemBlock, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read file")
	}
	csr, err := parseCSRPEM(pemBlock)
	if err != nil {
		return nil, fmt.Errorf("error reading certificate request file %s: %v", file, err)
	}
	return csr, nil
}
func NewCSR(cfg certutil.Config, key crypto.Signer) (*x509.CertificateRequest, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	template := &x509.CertificateRequest{Subject: pkix.Name{CommonName: cfg.CommonName, Organization: cfg.Organization}, DNSNames: cfg.AltNames.DNSNames, IPAddresses: cfg.AltNames.IPs}
	csrBytes, err := x509.CreateCertificateRequest(cryptorand.Reader, template, key)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create a CSR")
	}
	return x509.ParseCertificateRequest(csrBytes)
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
