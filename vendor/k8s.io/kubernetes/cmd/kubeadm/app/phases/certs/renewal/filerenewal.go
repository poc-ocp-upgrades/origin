package renewal

import (
	"crypto/rsa"
	"crypto/x509"
	certutil "k8s.io/client-go/util/cert"
	"k8s.io/kubernetes/cmd/kubeadm/app/util/pkiutil"
)

type FileRenewal struct {
	caCert *x509.Certificate
	caKey  *rsa.PrivateKey
}

func NewFileRenewal(caCert *x509.Certificate, caKey *rsa.PrivateKey) Interface {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &FileRenewal{caCert: caCert, caKey: caKey}
}
func (r *FileRenewal) Renew(cfg *certutil.Config) (*x509.Certificate, *rsa.PrivateKey, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return pkiutil.NewCertAndKey(r.caCert, r.caKey, cfg)
}
