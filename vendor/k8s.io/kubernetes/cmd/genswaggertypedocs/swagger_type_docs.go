package main

import (
	"fmt"
	goformat "fmt"
	flag "github.com/spf13/pflag"
	"io"
	kruntime "k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog"
	"os"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

var (
	functionDest = flag.StringP("func-dest", "f", "-", "Output for swagger functions; '-' means stdout (default)")
	typeSrc      = flag.StringP("type-src", "s", "", "From where we are going to read the types")
	verify       = flag.BoolP("verify", "v", false, "Verifies if the given type-src file has documentation for every type")
)

func main() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	flag.Parse()
	if *typeSrc == "" {
		klog.Fatalf("Please define -s flag as it is the source file")
	}
	var funcOut io.Writer
	if *functionDest == "-" {
		funcOut = os.Stdout
	} else {
		file, err := os.Create(*functionDest)
		if err != nil {
			klog.Fatalf("Couldn't open %v: %v", *functionDest, err)
		}
		defer file.Close()
		funcOut = file
	}
	docsForTypes := kruntime.ParseDocumentationFrom(*typeSrc)
	if *verify == true {
		rc, err := kruntime.VerifySwaggerDocsExist(docsForTypes, funcOut)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error in verification process: %s\n", err)
		}
		os.Exit(rc)
	}
	if docsForTypes != nil && len(docsForTypes) > 0 {
		if err := kruntime.WriteSwaggerDocFunc(docsForTypes, funcOut); err != nil {
			fmt.Fprintf(os.Stderr, "Error when writing swagger documentation functions: %s\n", err)
			os.Exit(-1)
		}
	}
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
