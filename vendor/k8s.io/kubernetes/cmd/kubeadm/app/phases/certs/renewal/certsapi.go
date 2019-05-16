package renewal

import (
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"fmt"
	goformat "fmt"
	"github.com/pkg/errors"
	certsapi "k8s.io/api/certificates/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	certstype "k8s.io/client-go/kubernetes/typed/certificates/v1beta1"
	certutil "k8s.io/client-go/util/cert"
	csrutil "k8s.io/client-go/util/certificate/csr"
	goos "os"
	godefaultruntime "runtime"
	"time"
	gotime "time"
)

const certAPIPrefixName = "kubeadm-cert"

var watchTimeout = 5 * time.Minute

type CertsAPIRenewal struct {
	client certstype.CertificatesV1beta1Interface
}

func NewCertsAPIRenawal(client kubernetes.Interface) Interface {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &CertsAPIRenewal{client: client.CertificatesV1beta1()}
}
func (r *CertsAPIRenewal) Renew(cfg *certutil.Config) (*x509.Certificate, *rsa.PrivateKey, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	reqTmp := &x509.CertificateRequest{Subject: pkix.Name{CommonName: cfg.CommonName, Organization: cfg.Organization}, DNSNames: cfg.AltNames.DNSNames, IPAddresses: cfg.AltNames.IPs}
	key, err := certutil.NewPrivateKey()
	if err != nil {
		return nil, nil, errors.Wrap(err, "couldn't create new private key")
	}
	csr, err := certutil.MakeCSRFromTemplate(key, reqTmp)
	if err != nil {
		return nil, nil, errors.Wrap(err, "couldn't create certificate signing request")
	}
	usages := make([]certsapi.KeyUsage, len(cfg.Usages))
	for i, usage := range cfg.Usages {
		certsAPIUsage, ok := usageMap[usage]
		if !ok {
			return nil, nil, errors.Errorf("unknown key usage: %v", usage)
		}
		usages[i] = certsAPIUsage
	}
	k8sCSR := &certsapi.CertificateSigningRequest{ObjectMeta: metav1.ObjectMeta{GenerateName: fmt.Sprintf("%s-%s-", certAPIPrefixName, cfg.CommonName)}, Spec: certsapi.CertificateSigningRequestSpec{Request: csr, Usages: usages}}
	req, err := r.client.CertificateSigningRequests().Create(k8sCSR)
	if err != nil {
		return nil, nil, errors.Wrap(err, "couldn't create certificate signing request")
	}
	fmt.Printf("[certs] certificate request %q created\n", req.Name)
	certData, err := csrutil.WaitForCertificate(r.client.CertificateSigningRequests(), req, watchTimeout)
	if err != nil {
		return nil, nil, errors.Wrap(err, "certificate failed to appear")
	}
	cert, err := certutil.ParseCertsPEM(certData)
	if err != nil {
		return nil, nil, errors.Wrap(err, "couldn't parse issued certificate")
	}
	if len(cert) != 1 {
		return nil, nil, errors.Errorf("certificate request %q has %d certificates, wanted exactly 1", req.Name, len(cert))
	}
	return cert[0], key, nil
}

var usageMap = map[x509.ExtKeyUsage]certsapi.KeyUsage{x509.ExtKeyUsageAny: certsapi.UsageAny, x509.ExtKeyUsageServerAuth: certsapi.UsageServerAuth, x509.ExtKeyUsageClientAuth: certsapi.UsageClientAuth, x509.ExtKeyUsageCodeSigning: certsapi.UsageCodeSigning, x509.ExtKeyUsageEmailProtection: certsapi.UsageEmailProtection, x509.ExtKeyUsageIPSECEndSystem: certsapi.UsageIPsecEndSystem, x509.ExtKeyUsageIPSECTunnel: certsapi.UsageIPsecTunnel, x509.ExtKeyUsageIPSECUser: certsapi.UsageIPsecUser, x509.ExtKeyUsageTimeStamping: certsapi.UsageTimestamping, x509.ExtKeyUsageOCSPSigning: certsapi.UsageOCSPSigning, x509.ExtKeyUsageMicrosoftServerGatedCrypto: certsapi.UsageMicrosoftSGC, x509.ExtKeyUsageNetscapeServerGatedCrypto: certsapi.UsageNetscapSGC}

func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
