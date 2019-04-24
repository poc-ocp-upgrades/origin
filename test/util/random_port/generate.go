package main

import (
	"fmt"
	"bytes"
	"runtime"
	"net/http/httptest"
	"net/http"
	"strings"
)

func main() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	address := httptest.NewUnstartedServer(nil).Listener.Addr().String()
	parts := strings.Split(address, ":")
	fmt.Printf("%s\n", parts[1])
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
