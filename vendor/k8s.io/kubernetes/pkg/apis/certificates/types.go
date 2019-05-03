package certificates

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type CertificateSigningRequest struct {
 metav1.TypeMeta
 metav1.ObjectMeta
 Spec   CertificateSigningRequestSpec
 Status CertificateSigningRequestStatus
}
type CertificateSigningRequestSpec struct {
 Request  []byte
 Usages   []KeyUsage
 Username string
 UID      string
 Groups   []string
 Extra    map[string]ExtraValue
}
type ExtraValue []string
type CertificateSigningRequestStatus struct {
 Conditions  []CertificateSigningRequestCondition
 Certificate []byte
}
type RequestConditionType string

const (
 CertificateApproved RequestConditionType = "Approved"
 CertificateDenied   RequestConditionType = "Denied"
)

type CertificateSigningRequestCondition struct {
 Type           RequestConditionType
 Reason         string
 Message        string
 LastUpdateTime metav1.Time
}
type CertificateSigningRequestList struct {
 metav1.TypeMeta
 metav1.ListMeta
 Items []CertificateSigningRequest
}
type KeyUsage string

const (
 UsageSigning            KeyUsage = "signing"
 UsageDigitalSignature   KeyUsage = "digital signature"
 UsageContentCommittment KeyUsage = "content commitment"
 UsageKeyEncipherment    KeyUsage = "key encipherment"
 UsageKeyAgreement       KeyUsage = "key agreement"
 UsageDataEncipherment   KeyUsage = "data encipherment"
 UsageCertSign           KeyUsage = "cert sign"
 UsageCRLSign            KeyUsage = "crl sign"
 UsageEncipherOnly       KeyUsage = "encipher only"
 UsageDecipherOnly       KeyUsage = "decipher only"
 UsageAny                KeyUsage = "any"
 UsageServerAuth         KeyUsage = "server auth"
 UsageClientAuth         KeyUsage = "client auth"
 UsageCodeSigning        KeyUsage = "code signing"
 UsageEmailProtection    KeyUsage = "email protection"
 UsageSMIME              KeyUsage = "s/mime"
 UsageIPsecEndSystem     KeyUsage = "ipsec end system"
 UsageIPsecTunnel        KeyUsage = "ipsec tunnel"
 UsageIPsecUser          KeyUsage = "ipsec user"
 UsageTimestamping       KeyUsage = "timestamping"
 UsageOCSPSigning        KeyUsage = "ocsp signing"
 UsageMicrosoftSGC       KeyUsage = "microsoft sgc"
 UsageNetscapSGC         KeyUsage = "netscape sgc"
)
