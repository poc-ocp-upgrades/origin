package renewal

import (
	"crypto/x509"
	"github.com/pkg/errors"
	certutil "k8s.io/client-go/util/cert"
	"k8s.io/kubernetes/cmd/kubeadm/app/util/pkiutil"
)

func RenewExistingCert(certsDir, baseName string, impl Interface) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	certificatePath, _ := pkiutil.PathsForCertAndKey(certsDir, baseName)
	certs, err := certutil.CertsFromFile(certificatePath)
	if err != nil {
		return errors.Wrapf(err, "failed to load existing certificate %s", baseName)
	}
	if len(certs) != 1 {
		return errors.Errorf("wanted exactly one certificate, got %d", len(certs))
	}
	cfg := certToConfig(certs[0])
	newCert, newKey, err := impl.Renew(cfg)
	if err != nil {
		return errors.Wrapf(err, "failed to renew certificate %s", baseName)
	}
	if err := pkiutil.WriteCertAndKey(certsDir, baseName, newCert, newKey); err != nil {
		return errors.Wrapf(err, "failed to write new certificate %s", baseName)
	}
	return nil
}
func certToConfig(cert *x509.Certificate) *certutil.Config {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &certutil.Config{CommonName: cert.Subject.CommonName, Organization: cert.Subject.Organization, AltNames: certutil.AltNames{IPs: cert.IPAddresses, DNSNames: cert.DNSNames}, Usages: cert.ExtKeyUsage}
}
