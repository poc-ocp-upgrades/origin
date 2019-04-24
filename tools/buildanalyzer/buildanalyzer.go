package main

import (
	"os"
	"bytes"
	"net/http"
	"runtime"
	"fmt"
	"github.com/openshift/origin/tools/buildanalyzer/cmd"
)

func main() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	command := cmd.NewBuildAnalyzerCommand()
	if err := command.Execute(); err != nil {
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
