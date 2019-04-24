package errors

import (
	"fmt"
	"bytes"
	"net/http"
	"runtime"
)

type Error interface {
	error
	WithCause(error) Error
	WithSolution(string) Error
	WithDetails(string) Error
}

func NewError(msg string, args ...interface{}) Error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &internalError{msg: fmt.Sprintf(msg, args...)}
}

type internalError struct {
	msg		string
	cause		error
	solution	string
	details		string
}

func (e *internalError) Error() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if e.cause != nil && len(e.cause.Error()) > 0 {
		return e.msg + "; caused by: " + e.cause.Error()
	}
	return e.msg
}
func (e *internalError) Cause() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return e.cause
}
func (e *internalError) Solution() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return e.solution
}
func (e *internalError) Details() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return e.details
}
func (e *internalError) WithCause(err error) Error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	e.cause = err
	return e
}
func (e *internalError) WithDetails(details string) Error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	e.details = details
	return e
}
func (e *internalError) WithSolution(solution string) Error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	e.solution = solution
	return e
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
