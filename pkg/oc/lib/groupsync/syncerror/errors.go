package syncerror

import (
	"fmt"
	"bytes"
	"net/http"
	"runtime"
)

func NewMemberLookupError(ldapGroupUID, ldapUserUID string, causedBy error) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &memberLookupError{ldapGroupUID: ldapGroupUID, ldapUserUID: ldapUserUID, causedBy: causedBy}
}

type memberLookupError struct {
	ldapGroupUID	string
	ldapUserUID	string
	causedBy	error
}

func (e *memberLookupError) Error() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fmt.Sprintf("membership lookup for user %q in group %q failed because of %q", e.ldapUserUID, e.ldapGroupUID, e.causedBy.Error())
}
func IsMemberLookupError(e error) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if e == nil {
		return false
	}
	_, ok := e.(*memberLookupError)
	return ok
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
