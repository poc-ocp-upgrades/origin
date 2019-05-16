package oauth

import (
	goformat "fmt"
	configv1 "github.com/openshift/api/config/v1"
	crvalidation "github.com/openshift/origin/pkg/admission/customresourcevalidation"
	"github.com/openshift/origin/pkg/cmd/server/apis/config/validation/common"
	kvalidation "k8s.io/apimachinery/pkg/util/validation"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"net"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

func isValidHostname(hostname string) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return len(kvalidation.IsDNS1123Subdomain(hostname)) == 0 || net.ParseIP(hostname) != nil
}
func ValidateRemoteConnectionInfo(remoteConnectionInfo configv1.OAuthRemoteConnectionInfo, fldPath *field.Path) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	allErrs := field.ErrorList{}
	if len(remoteConnectionInfo.URL) == 0 {
		allErrs = append(allErrs, field.Required(fldPath.Child("url"), ""))
	} else {
		_, urlErrs := common.ValidateSecureURL(remoteConnectionInfo.URL, fldPath.Child("url"))
		allErrs = append(allErrs, urlErrs...)
	}
	allErrs = append(allErrs, crvalidation.ValidateConfigMapReference(fldPath.Child("ca"), remoteConnectionInfo.CA, false)...)
	allErrs = append(allErrs, crvalidation.ValidateSecretReference(fldPath.Child("tlsClientCert"), remoteConnectionInfo.TLSClientCert, false)...)
	allErrs = append(allErrs, crvalidation.ValidateSecretReference(fldPath.Child("tlsClientKey"), remoteConnectionInfo.TLSClientKey, false)...)
	return allErrs
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
