package errorpage

import (
	goformat "fmt"
	"github.com/openshift/origin/pkg/oauthserver/userregistry/identitymapper"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

const (
	errorCodeClaim          = "mapping_claim_error"
	errorCodeLookup         = "mapping_lookup_error"
	errorCodeAuthentication = "authentication_error"
	errorCodeGrant          = "grant_error"
)

func AuthenticationErrorCode(err error) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	switch {
	case identitymapper.IsClaimError(err):
		return errorCodeClaim
	case identitymapper.IsLookupError(err):
		return errorCodeLookup
	default:
		return errorCodeAuthentication
	}
}
func AuthenticationErrorMessage(code string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	switch code {
	case errorCodeClaim:
		return "Could not create user."
	case errorCodeLookup:
		return "Could not find user."
	default:
		return "An authentication error occurred."
	}
}
func GrantErrorCode(err error) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return errorCodeGrant
}
func GrantErrorMessage(code string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return "A grant error occurred."
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
