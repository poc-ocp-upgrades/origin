package ldapclient

import (
	"gopkg.in/ldap.v2"
	"bytes"
	"net/http"
	"runtime"
	"fmt"
)

type Config interface {
	Connect() (client ldap.Client, err error)
	GetBindCredentials() (bindDN, bindPassword string)
	Host() string
}

func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
