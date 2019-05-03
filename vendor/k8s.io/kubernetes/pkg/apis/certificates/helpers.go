package certificates

import (
 "crypto/x509"
 "encoding/pem"
 "errors"
)

func ParseCSR(obj *CertificateSigningRequest) (*x509.CertificateRequest, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 pemBytes := obj.Spec.Request
 block, _ := pem.Decode(pemBytes)
 if block == nil || block.Type != "CERTIFICATE REQUEST" {
  return nil, errors.New("PEM block type must be CERTIFICATE REQUEST")
 }
 csr, err := x509.ParseCertificateRequest(block.Bytes)
 if err != nil {
  return nil, err
 }
 return csr, nil
}
