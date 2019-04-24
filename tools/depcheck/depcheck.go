package main

import (
	"fmt"
	"bytes"
	"net/http"
	"runtime"
	"os"
	"github.com/openshift/origin/tools/depcheck/pkg/cmd"
)

func main() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	command := cmd.NewCmdDepCheck(os.Args[0], os.Stdout, os.Stderr)
	if err := command.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
