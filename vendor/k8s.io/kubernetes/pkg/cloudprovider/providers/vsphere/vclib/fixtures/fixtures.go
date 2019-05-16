package fixtures

import (
	goformat "fmt"
	"os"
	goos "os"
	"path/filepath"
	"runtime"
	godefaultruntime "runtime"
	"strings"
	gotime "time"
)

var (
	CaCertPath      string
	ServerCertPath  string
	ServerKeyPath   string
	InvalidCertPath string
)

func init() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	_, thisFile, _, ok := runtime.Caller(0)
	if !ok {
		panic("Cannot get path to the fixtures")
	}
	fixturesDir := filepath.Dir(thisFile)
	cwd, err := os.Getwd()
	if err != nil {
		panic("Cannot get CWD: " + err.Error())
	}
	if !strings.HasPrefix(fixturesDir, cwd) {
		fixturesDir = filepath.Join(cwd, fixturesDir)
	}
	CaCertPath = filepath.Join(fixturesDir, "ca.pem")
	ServerCertPath = filepath.Join(fixturesDir, "server.pem")
	ServerKeyPath = filepath.Join(fixturesDir, "server.key")
	InvalidCertPath = filepath.Join(fixturesDir, "invalid.pem")
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
