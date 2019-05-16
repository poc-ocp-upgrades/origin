package genutils

import (
	"fmt"
	goformat "fmt"
	"os"
	goos "os"
	"path/filepath"
	godefaultruntime "runtime"
	gotime "time"
)

func OutDir(path string) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	outDir, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}
	stat, err := os.Stat(outDir)
	if err != nil {
		return "", err
	}
	if !stat.IsDir() {
		return "", fmt.Errorf("output directory %s is not a directory", outDir)
	}
	outDir = outDir + "/"
	return outDir, nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
