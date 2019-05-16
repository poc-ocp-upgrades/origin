package session

import (
	"github.com/gorilla/sessions"
	"k8s.io/klog"
	"net/http"
)

type store struct {
	name  string
	store sessions.Store
}

func NewStore(name string, secure bool, secrets ...[]byte) Store {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	cookie := sessions.NewCookieStore(secrets...)
	cookie.Options.MaxAge = 0
	cookie.Options.HttpOnly = true
	cookie.Options.Secure = secure
	return &store{name: name, store: cookie}
}
func (s *store) Get(r *http.Request) Values {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	session, err := s.store.New(r, s.name)
	if err != nil {
		klog.V(4).Infof("failed to decode secure session cookie %s: %v", s.name, err)
		return Values{}
	}
	return session.Values
}
func (s *store) Put(w http.ResponseWriter, v Values) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	r := &http.Request{}
	session, err := s.store.New(r, s.name)
	if err != nil {
		return err
	}
	session.Values = v
	return s.store.Save(r, w, session)
}
