package signer

import (
	"crypto"
	"crypto/x509"
	"fmt"
	goformat "fmt"
	"github.com/cloudflare/cfssl/config"
	"github.com/cloudflare/cfssl/helpers"
	"github.com/cloudflare/cfssl/signer"
	"github.com/cloudflare/cfssl/signer/local"
	"io/ioutil"
	capi "k8s.io/api/certificates/v1beta1"
	certificatesinformers "k8s.io/client-go/informers/certificates/v1beta1"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/kubernetes/pkg/controller/certificates"
	"os"
	goos "os"
	godefaultruntime "runtime"
	"time"
	gotime "time"
)

func NewCSRSigningController(client clientset.Interface, csrInformer certificatesinformers.CertificateSigningRequestInformer, caFile, caKeyFile string, certificateDuration time.Duration) (*certificates.CertificateController, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	signer, err := newCFSSLSigner(caFile, caKeyFile, client, certificateDuration)
	if err != nil {
		return nil, err
	}
	return certificates.NewCertificateController(client, csrInformer, signer.handle), nil
}

type cfsslSigner struct {
	ca                  *x509.Certificate
	priv                crypto.Signer
	sigAlgo             x509.SignatureAlgorithm
	client              clientset.Interface
	certificateDuration time.Duration
	nowFn               func() time.Time
}

func newCFSSLSigner(caFile, caKeyFile string, client clientset.Interface, certificateDuration time.Duration) (*cfsslSigner, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ca, err := ioutil.ReadFile(caFile)
	if err != nil {
		return nil, fmt.Errorf("error reading CA cert file %q: %v", caFile, err)
	}
	cakey, err := ioutil.ReadFile(caKeyFile)
	if err != nil {
		return nil, fmt.Errorf("error reading CA key file %q: %v", caKeyFile, err)
	}
	parsedCa, err := helpers.ParseCertificatePEM(ca)
	if err != nil {
		return nil, fmt.Errorf("error parsing CA cert file %q: %v", caFile, err)
	}
	strPassword := os.Getenv("CFSSL_CA_PK_PASSWORD")
	password := []byte(strPassword)
	if strPassword == "" {
		password = nil
	}
	priv, err := helpers.ParsePrivateKeyPEMWithPassword(cakey, password)
	if err != nil {
		return nil, fmt.Errorf("Malformed private key %v", err)
	}
	return &cfsslSigner{priv: priv, ca: parsedCa, sigAlgo: signer.DefaultSigAlgo(priv), client: client, certificateDuration: certificateDuration, nowFn: time.Now}, nil
}
func (s *cfsslSigner) handle(csr *capi.CertificateSigningRequest) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if !certificates.IsCertificateRequestApproved(csr) {
		return nil
	}
	csr, err := s.sign(csr)
	if err != nil {
		return fmt.Errorf("error auto signing csr: %v", err)
	}
	_, err = s.client.CertificatesV1beta1().CertificateSigningRequests().UpdateStatus(csr)
	if err != nil {
		return fmt.Errorf("error updating signature for csr: %v", err)
	}
	return nil
}
func (s *cfsslSigner) sign(csr *capi.CertificateSigningRequest) (*capi.CertificateSigningRequest, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var usages []string
	for _, usage := range csr.Spec.Usages {
		usages = append(usages, string(usage))
	}
	certExpiryDuration := s.certificateDuration
	durationUntilExpiry := s.ca.NotAfter.Sub(s.nowFn())
	if durationUntilExpiry <= 0 {
		return nil, fmt.Errorf("the signer has expired: %v", s.ca.NotAfter)
	}
	if durationUntilExpiry < certExpiryDuration {
		certExpiryDuration = durationUntilExpiry
	}
	policy := &config.Signing{Default: &config.SigningProfile{Usage: usages, Expiry: certExpiryDuration, ExpiryString: certExpiryDuration.String()}}
	cfs, err := local.NewSigner(s.priv, s.ca, s.sigAlgo, policy)
	if err != nil {
		return nil, err
	}
	csr.Status.Certificate, err = cfs.Sign(signer.SignRequest{Request: string(csr.Spec.Request)})
	if err != nil {
		return nil, err
	}
	return csr, nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
