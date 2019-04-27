package session

import (
	"net/http"
	"github.com/gorilla/sessions"
	"k8s.io/klog"
)

type store struct {
	name	string
	store	sessions.Store
}

func NewStore(name string, secure bool, secrets ...[]byte) Store {
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
	cookie := sessions.NewCookieStore(secrets...)
	cookie.Options.MaxAge = 0
	cookie.Options.HttpOnly = true
	cookie.Options.Secure = secure
	return &store{name: name, store: cookie}
}
func (s *store) Get(r *http.Request) Values {
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
	session, err := s.store.New(r, s.name)
	if err != nil {
		klog.V(4).Infof("failed to decode secure session cookie %s: %v", s.name, err)
		return Values{}
	}
	return session.Values
}
func (s *store) Put(w http.ResponseWriter, v Values) error {
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
	r := &http.Request{}
	session, err := s.store.New(r, s.name)
	if err != nil {
		return err
	}
	session.Values = v
	return s.store.Save(r, w, session)
}
