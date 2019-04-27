package osinserver

import (
	"fmt"
	"net/http"
	"path"
	"github.com/RangelReale/osin"
	"k8s.io/klog"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"github.com/openshift/origin/pkg/oauth/urls"
	"github.com/openshift/origin/pkg/oauthserver"
)

type Server struct {
	config		*osin.ServerConfig
	server		*osin.Server
	authorize	AuthorizeHandler
	access		AccessHandler
	errorHandler	ErrorHandler
}
type Logger struct{}

func (l Logger) Printf(format string, v ...interface{}) {
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
	if klog.V(2) {
		klog.ErrorDepth(3, fmt.Sprintf("osin: "+format, v...))
	}
}
func New(config *osin.ServerConfig, storage osin.Storage, authorize AuthorizeHandler, access AccessHandler, errorHandler ErrorHandler) *Server {
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
	server := osin.NewServer(config, storage)
	server.AuthorizeTokenGen = TokenGen{}
	server.AccessTokenGen = TokenGen{}
	server.Logger = Logger{}
	return &Server{config: config, server: server, authorize: authorize, access: access, errorHandler: errorHandler}
}
func (s *Server) Install(mux oauthserver.Mux, paths ...string) {
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
	for _, prefix := range paths {
		mux.HandleFunc(path.Join(prefix, urls.AuthorizePath), s.handleAuthorize)
		mux.HandleFunc(path.Join(prefix, urls.TokenPath), s.handleToken)
		mux.HandleFunc(path.Join(prefix, urls.InfoPath), s.handleInfo)
	}
}
func (s *Server) handleAuthorize(w http.ResponseWriter, r *http.Request) {
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
	resp := s.server.NewResponse()
	defer resp.Close()
	if ar := s.server.HandleAuthorizeRequest(resp, r); ar != nil {
		if errorCode := r.FormValue("error"); len(errorCode) != 0 {
			resp.SetErrorUri(r.FormValue("error"), r.FormValue("error_description"), r.FormValue("error_uri"), r.FormValue("state"))
			resp.SetRedirect(ar.RedirectUri)
		} else {
			handled, err := s.authorize.HandleAuthorize(ar, resp, w)
			if err != nil {
				s.errorHandler.HandleError(err, w, r)
				return
			}
			if handled {
				return
			}
			s.server.FinishAuthorizeRequest(resp, r, ar)
		}
	}
	if resp.IsError && resp.InternalError != nil {
		utilruntime.HandleError(fmt.Errorf("internal error: %s", resp.InternalError))
	}
	osin.OutputJSON(resp, w, r)
}
func (s *Server) handleToken(w http.ResponseWriter, r *http.Request) {
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
	resp := s.server.NewResponse()
	defer resp.Close()
	if ar := s.server.HandleAccessRequest(resp, r); ar != nil {
		if err := s.access.HandleAccess(ar, w); err != nil {
			s.errorHandler.HandleError(err, w, r)
			return
		}
		s.server.FinishAccessRequest(resp, r, ar)
	}
	if resp.IsError && resp.InternalError != nil {
		utilruntime.HandleError(fmt.Errorf("internal error: %s", resp.InternalError))
	}
	osin.OutputJSON(resp, w, r)
}
func (s *Server) handleInfo(w http.ResponseWriter, r *http.Request) {
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
	resp := s.server.NewResponse()
	defer resp.Close()
	if ir := s.server.HandleInfoRequest(resp, r); ir != nil {
		s.server.FinishInfoRequest(resp, r, ir)
	}
	osin.OutputJSON(resp, w, r)
}
