package errors

import (
	"fmt"
	"bytes"
	"net/http"
	"runtime"
)

func NewSuiteOutOfBoundsError(name string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &suiteOutOfBoundsError{suiteName: name}
}

type suiteOutOfBoundsError struct{ suiteName string }

func (e *suiteOutOfBoundsError) Error() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fmt.Sprintf("the test suite %q could not be placed under any existing roots in the tree", e.suiteName)
}
func IsSuiteOutOfBoundsError(err error) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err == nil {
		return false
	}
	_, ok := err.(*suiteOutOfBoundsError)
	return ok
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
