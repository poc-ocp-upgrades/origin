package main

import (
	"fmt"
	goformat "fmt"
	"github.com/spf13/cobra/doc"
	"io/ioutil"
	"k8s.io/kubernetes/cmd/genutils"
	"k8s.io/kubernetes/pkg/kubectl/cmd"
	"os"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

func main() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	path := "docs/"
	if len(os.Args) == 2 {
		path = os.Args[1]
	} else if len(os.Args) > 2 {
		fmt.Fprintf(os.Stderr, "usage: %s [output directory]\n", os.Args[0])
		os.Exit(1)
	}
	outDir, err := genutils.OutDir(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to get output directory: %v\n", err)
		os.Exit(1)
	}
	os.Setenv("HOME", "/home/username")
	kubectl := cmd.NewKubectlCommand(os.Stdin, ioutil.Discard, ioutil.Discard)
	doc.GenMarkdownTree(kubectl, outDir)
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
