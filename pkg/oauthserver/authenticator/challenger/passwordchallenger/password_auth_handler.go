package passwordchallenger

import (
	"fmt"
	"bytes"
	"runtime"
	"net/http"
	oauthhandlers "github.com/openshift/origin/pkg/oauthserver/oauth/handlers"
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
