package errors

import (
	goformat "fmt"
	kapierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	goos "os"
	godefaultruntime "runtime"
	"strings"
	gotime "time"
)

func TolerateNotFoundError(err error) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if kapierrors.IsNotFound(err) {
		return nil
	}
	return err
}
func ErrorToSentence(err error) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	msg := err.Error()
	if len(msg) == 0 {
		return msg
	}
	msg = strings.ToUpper(msg)[:1] + msg[1:]
	if !strings.HasSuffix(msg, ".") {
		msg = msg + "."
	}
	return msg
}
func IsTimeoutErr(err error) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	e, ok := err.(*kapierrors.StatusError)
	if !ok {
		return false
	}
	return e.ErrStatus.Reason == metav1.StatusReasonTimeout
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
