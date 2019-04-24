package osintypes

import (
	"fmt"
	"bytes"
	"net/http"
	"runtime"
)

type InfoResponseData struct {
	Error			string	`json:"error"`
	ErrorDescription	string	`json:"error_description"`
	TokenType		string	`json:"token_type"`
	AccessToken		string	`json:"access_token"`
	RefreshToken		string	`json:"refresh_token"`
	Expiration		int32	`json:"expires_in"`
}

func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
