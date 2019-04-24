package test

import (
	"testing"
	"bytes"
	"net/http"
	"runtime"
	"fmt"
)

func TestFoo(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	t.Run("panic", func(t *testing.T) {
		panic("here")
	})
	t.Run("pass", func(t *testing.T) {
	})
	t.Run("skip", func(t *testing.T) {
		t.Skip("skipped")
	})
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
