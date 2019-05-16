package certificates

import certificates "k8s.io/api/certificates/v1beta1"

func IsCertificateRequestApproved(csr *certificates.CertificateSigningRequest) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	approved, denied := GetCertApprovalCondition(&csr.Status)
	return approved && !denied
}
func GetCertApprovalCondition(status *certificates.CertificateSigningRequestStatus) (approved bool, denied bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for _, c := range status.Conditions {
		if c.Type == certificates.CertificateApproved {
			approved = true
		}
		if c.Type == certificates.CertificateDenied {
			denied = true
		}
	}
	return
}
