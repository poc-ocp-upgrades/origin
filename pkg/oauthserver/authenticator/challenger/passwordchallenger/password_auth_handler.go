package passwordchallenger

import (
	"fmt"
	goformat "fmt"
	oauthhandlers "github.com/openshift/origin/pkg/oauthserver/oauth/handlers"
	"net/http"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

type basicPasswordAuthHandler struct{ realm string }

const CSRFTokenHeader = "X-CSRF-Token"

func NewBasicAuthChallenger(realm string) oauthhandlers.AuthenticationChallenger {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &basicPasswordAuthHandler{realm}
}
func (h *basicPasswordAuthHandler) AuthenticationChallenge(req *http.Request) (http.Header, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	headers := http.Header{}
	if len(req.Header.Get(CSRFTokenHeader)) == 0 {
		headers.Add("Warning", fmt.Sprintf(`%s %s "A non-empty %s header is required to receive basic-auth challenges"`, oauthhandlers.WarningHeaderMiscCode, oauthhandlers.WarningHeaderOpenShiftSource, CSRFTokenHeader))
	} else {
		headers.Add("WWW-Authenticate", fmt.Sprintf(`Basic realm="%s"`, h.realm))
	}
	return headers, nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
