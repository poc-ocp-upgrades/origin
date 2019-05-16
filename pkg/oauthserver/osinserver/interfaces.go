package osinserver

import (
	"github.com/RangelReale/osin"
	"net/http"
)

type AuthorizeHandler interface {
	HandleAuthorize(ar *osin.AuthorizeRequest, resp *osin.Response, w http.ResponseWriter) (handled bool, err error)
}
type AuthorizeHandlerFunc func(ar *osin.AuthorizeRequest, resp *osin.Response, w http.ResponseWriter) (bool, error)

func (f AuthorizeHandlerFunc) HandleAuthorize(ar *osin.AuthorizeRequest, resp *osin.Response, w http.ResponseWriter) (bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return f(ar, resp, w)
}

type AuthorizeHandlers []AuthorizeHandler

func (all AuthorizeHandlers) HandleAuthorize(ar *osin.AuthorizeRequest, resp *osin.Response, w http.ResponseWriter) (bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for _, h := range all {
		if handled, err := h.HandleAuthorize(ar, resp, w); handled || err != nil {
			return handled, err
		}
	}
	return false, nil
}

type AccessHandler interface {
	HandleAccess(ar *osin.AccessRequest, w http.ResponseWriter) error
}
type AccessHandlerFunc func(ar *osin.AccessRequest, w http.ResponseWriter) error

func (f AccessHandlerFunc) HandleAccess(ar *osin.AccessRequest, w http.ResponseWriter) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return f(ar, w)
}

type AccessHandlers []AccessHandler

func (all AccessHandlers) HandleAccess(ar *osin.AccessRequest, w http.ResponseWriter) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for _, h := range all {
		if err := h.HandleAccess(ar, w); err != nil {
			return err
		}
	}
	return nil
}

type ErrorHandler interface {
	HandleError(err error, w http.ResponseWriter, req *http.Request)
}
