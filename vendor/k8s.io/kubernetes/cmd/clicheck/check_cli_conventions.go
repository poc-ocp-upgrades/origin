package main

import (
	"fmt"
	goformat "fmt"
	"io/ioutil"
	"k8s.io/kubernetes/pkg/kubectl/cmd"
	cmdsanity "k8s.io/kubernetes/pkg/kubectl/cmd/util/sanity"
	"os"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

var (
	skip = []string{}
)

func main() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var errorCount int
	kubectl := cmd.NewKubectlCommand(os.Stdin, ioutil.Discard, ioutil.Discard)
	errors := cmdsanity.RunCmdChecks(kubectl, cmdsanity.AllCmdChecks, []string{})
	for _, err := range errors {
		errorCount++
		fmt.Fprintf(os.Stderr, "     %d. %s\n", errorCount, err)
	}
	errors = cmdsanity.RunGlobalChecks(cmdsanity.AllGlobalChecks)
	for _, err := range errors {
		errorCount++
		fmt.Fprintf(os.Stderr, "     %d. %s\n", errorCount, err)
	}
	if errorCount > 0 {
		fmt.Fprintf(os.Stdout, "Found %d errors.\n", errorCount)
		os.Exit(1)
	}
	fmt.Fprintln(os.Stdout, "Congrats, CLI looks good!")
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
