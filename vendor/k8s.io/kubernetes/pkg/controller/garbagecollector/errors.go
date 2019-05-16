package garbagecollector

import (
	"fmt"
)

type restMappingError struct {
	kind    string
	version string
}

func (r *restMappingError) Error() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	versionKind := fmt.Sprintf("%s/%s", r.version, r.kind)
	return fmt.Sprintf("unable to get REST mapping for %s.", versionKind)
}
func (r *restMappingError) Message() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	versionKind := fmt.Sprintf("%s/%s", r.version, r.kind)
	errMsg := fmt.Sprintf("unable to get REST mapping for %s. ", versionKind)
	errMsg += fmt.Sprintf(" If %s is an invalid resource, then you should manually remove ownerReferences that refer %s objects.", versionKind, versionKind)
	return errMsg
}
func newRESTMappingError(kind, version string) *restMappingError {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &restMappingError{kind: kind, version: version}
}
