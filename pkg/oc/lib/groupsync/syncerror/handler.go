package syncerror

import (
	"fmt"
	"io"
	"github.com/openshift/origin/pkg/oauthserver/ldaputil"
)

type Handler interface {
	HandleError(err error) (handled bool, fatalError error)
}

func NewCompoundHandler(handlers ...Handler) Handler {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &compoundHandler{handlers: handlers}
}

type compoundHandler struct{ handlers []Handler }

func (h *compoundHandler) HandleError(err error) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, handler := range h.handlers {
		handled, handleErr := handler.HandleError(err)
		if handled || handleErr != nil {
			return handled, handleErr
		}
	}
	return false, nil
}
func NewMemberLookupOutOfBoundsSuppressor(err io.Writer) Handler {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &memberLookupOutOfBoundsSuppressor{err: err}
}

type memberLookupOutOfBoundsSuppressor struct{ err io.Writer }

func (h *memberLookupOutOfBoundsSuppressor) HandleError(err error) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	memberLookupError, isMemberLookupError := err.(*memberLookupError)
	if !isMemberLookupError {
		return false, nil
	}
	if ldaputil.IsQueryOutOfBoundsError(memberLookupError.causedBy) {
		fmt.Fprintf(h.err, "For group %q, ignoring member %q: %v\n", memberLookupError.ldapGroupUID, memberLookupError.ldapUserUID, memberLookupError.causedBy)
		return true, nil
	}
	return false, nil
}
func NewMemberLookupMemberNotFoundSuppressor(err io.Writer) Handler {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &memberLookupMemberNotFoundSuppressor{err: err}
}

type memberLookupMemberNotFoundSuppressor struct{ err io.Writer }

func (h *memberLookupMemberNotFoundSuppressor) HandleError(err error) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	memberLookupError, isMemberLookupError := err.(*memberLookupError)
	if !isMemberLookupError {
		return false, nil
	}
	if ldaputil.IsEntryNotFoundError(memberLookupError.causedBy) || ldaputil.IsNoSuchObjectError(memberLookupError.causedBy) {
		fmt.Fprintf(h.err, "For group %q, ignoring member %q: %v\n", memberLookupError.ldapGroupUID, memberLookupError.ldapUserUID, memberLookupError.causedBy)
		return true, nil
	}
	return false, nil
}
