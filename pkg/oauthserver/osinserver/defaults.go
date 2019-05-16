package osinserver

import (
	"fmt"
	goformat "fmt"
	"github.com/RangelReale/osin"
	"net/http"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

func NewDefaultServerConfig() *osin.ServerConfig {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	config := osin.NewServerConfig()
	config.AllowedAuthorizeTypes = osin.AllowedAuthorizeType{osin.CODE, osin.TOKEN}
	config.AllowedAccessTypes = osin.AllowedAccessType{osin.AUTHORIZATION_CODE, osin.REFRESH_TOKEN, osin.PASSWORD, osin.CLIENT_CREDENTIALS, osin.ASSERTION}
	config.AllowClientSecretInParams = true
	config.AllowGetAccessRequest = true
	config.RedirectUriSeparator = ","
	config.ErrorStatusCode = http.StatusBadRequest
	return config
}

type defaultErrorHandler struct{}

func NewDefaultErrorHandler() ErrorHandler {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return defaultErrorHandler{}
}
func (defaultErrorHandler) HandleError(err error, w http.ResponseWriter, req *http.Request) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, "Error: %s", err)
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
