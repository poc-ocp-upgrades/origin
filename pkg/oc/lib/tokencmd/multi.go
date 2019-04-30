package tokencmd

import (
	"net/http"
	"k8s.io/klog"
	apierrs "k8s.io/apimachinery/pkg/api/errors"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
)

var _ = ChallengeHandler(&MultiHandler{})

type MultiHandler struct {
	handler			ChallengeHandler
	possibleHandlers	[]ChallengeHandler
	allHandlers		[]ChallengeHandler
}

func NewMultiHandler(handlers ...ChallengeHandler) ChallengeHandler {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &MultiHandler{possibleHandlers: handlers, allHandlers: handlers}
}
func (h *MultiHandler) CanHandle(headers http.Header) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if h.handler != nil {
		return h.handler.CanHandle(headers)
	}
	for _, handler := range h.possibleHandlers {
		if handler.CanHandle(headers) {
			return true
		}
	}
	return false
}
func (h *MultiHandler) HandleChallenge(requestURL string, headers http.Header) (http.Header, bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if h.handler != nil {
		return h.handler.HandleChallenge(requestURL, headers)
	}
	applicable := []ChallengeHandler{}
	for _, handler := range h.possibleHandlers {
		if handler.CanHandle(headers) {
			applicable = append(applicable, handler)
		}
	}
	h.possibleHandlers = applicable
	var (
		retryHeaders	http.Header
		retry		bool
		err		error
	)
	for i, handler := range h.possibleHandlers {
		retryHeaders, retry, err = handler.HandleChallenge(requestURL, headers)
		if err != nil {
			klog.V(5).Infof("handler[%d] error: %v", i, err)
		}
		if err == nil || i == len(h.possibleHandlers)-1 {
			h.handler = handler
			return retryHeaders, retry, err
		}
	}
	return nil, false, apierrs.NewUnauthorized("unhandled challenge")
}
func (h *MultiHandler) CompleteChallenge(requestURL string, headers http.Header) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if h.handler != nil {
		return h.handler.CompleteChallenge(requestURL, headers)
	}
	return nil
}
func (h *MultiHandler) Release() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var errs []error
	for _, handler := range h.allHandlers {
		if err := handler.Release(); err != nil {
			errs = append(errs, err)
		}
	}
	return utilerrors.NewAggregate(errs)
}
