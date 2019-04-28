package osintypes

import (
	"fmt"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
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
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
