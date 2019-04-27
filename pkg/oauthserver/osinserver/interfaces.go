package osinserver

import (
	"net/http"
	"github.com/RangelReale/osin"
)

type AuthorizeHandler interface {
	HandleAuthorize(ar *osin.AuthorizeRequest, resp *osin.Response, w http.ResponseWriter) (handled bool, err error)
}
type AuthorizeHandlerFunc func(ar *osin.AuthorizeRequest, resp *osin.Response, w http.ResponseWriter) (bool, error)

func (f AuthorizeHandlerFunc) HandleAuthorize(ar *osin.AuthorizeRequest, resp *osin.Response, w http.ResponseWriter) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return f(ar, resp, w)
}

type AuthorizeHandlers []AuthorizeHandler

func (all AuthorizeHandlers) HandleAuthorize(ar *osin.AuthorizeRequest, resp *osin.Response, w http.ResponseWriter) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return f(ar, w)
}

type AccessHandlers []AccessHandler

func (all AccessHandlers) HandleAccess(ar *osin.AccessRequest, w http.ResponseWriter) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
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
