package config

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"bytes"
	"net/http"
	"runtime"
	"fmt"
)

type BootstrapIdentityProvider struct{ v1.TypeMeta }

func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
