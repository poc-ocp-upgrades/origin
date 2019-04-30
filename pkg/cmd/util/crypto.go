package util

import (
	"bytes"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"k8s.io/client-go/util/cert"
)

func CertPoolFromFile(filename string) (*x509.CertPool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pool := x509.NewCertPool()
	if len(filename) == 0 {
		return pool, nil
	}
	pemBlock, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	certs, err := cert.ParseCertsPEM(pemBlock)
	if err != nil {
		return nil, fmt.Errorf("Error reading %s: %s", filename, err)
	}
	for _, cert := range certs {
		pool.AddCert(cert)
	}
	return pool, nil
}
func CertificatesFromFile(file string) ([]*x509.Certificate, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(file) == 0 {
		return nil, nil
	}
	pemBlock, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	certs, err := cert.ParseCertsPEM(pemBlock)
	if err != nil {
		return nil, fmt.Errorf("Error reading %s: %s", file, err)
	}
	return certs, nil
}
func PrivateKeysFromPEM(pemCerts []byte) ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	buf := &bytes.Buffer{}
	for len(pemCerts) > 0 {
		var block *pem.Block
		block, pemCerts = pem.Decode(pemCerts)
		if block == nil {
			break
		}
		if len(block.Headers) != 0 {
			continue
		}
		switch block.Type {
		case "RSA PRIVATE KEY", "PRIVATE KEY", "ANY PRIVATE KEY", "DSA PRIVATE KEY", "ENCRYPTED PRIVATE KEY", "EC PRIVATE KEY":
			if err := pem.Encode(buf, block); err != nil {
				return nil, err
			}
		}
	}
	return buf.Bytes(), nil
}
