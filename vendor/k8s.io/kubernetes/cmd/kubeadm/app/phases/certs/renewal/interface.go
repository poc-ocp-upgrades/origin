package renewal

import (
	"crypto/rsa"
	"crypto/x509"
	certutil "k8s.io/client-go/util/cert"
)

type Interface interface {
	Renew(*certutil.Config) (*x509.Certificate, *rsa.PrivateKey, error)
}
