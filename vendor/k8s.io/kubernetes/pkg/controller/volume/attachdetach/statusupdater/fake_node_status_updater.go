package statusupdater

import (
	"fmt"
	goformat "fmt"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

func NewFakeNodeStatusUpdater(returnError bool) NodeStatusUpdater {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &fakeNodeStatusUpdater{returnError: returnError}
}

type fakeNodeStatusUpdater struct{ returnError bool }

func (fnsu *fakeNodeStatusUpdater) UpdateNodeStatuses() error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if fnsu.returnError {
		return fmt.Errorf("fake error on update node status")
	}
	return nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
