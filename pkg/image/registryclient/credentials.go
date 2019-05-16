package registryclient

import (
	"github.com/docker/distribution/registry/client/auth"
	"net/url"
	"sync"
)

var (
	NoCredentials auth.CredentialStore = &noopCredentialStore{}
)

type RefreshTokenStore interface {
	RefreshToken(url *url.URL, service string) string
	SetRefreshToken(url *url.URL, service string, token string)
}

func NewRefreshTokenStore() RefreshTokenStore {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &refreshTokenStore{}
}

type refreshTokenKey struct {
	url     string
	service string
}
type refreshTokenStore struct {
	lock  sync.Mutex
	store map[refreshTokenKey]string
}

func (s *refreshTokenStore) RefreshToken(url *url.URL, service string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.store[refreshTokenKey{url: url.String(), service: service}]
}
func (s *refreshTokenStore) SetRefreshToken(url *url.URL, service string, token string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.store == nil {
		s.store = make(map[refreshTokenKey]string)
	}
	s.store[refreshTokenKey{url: url.String(), service: service}] = token
}

type noopCredentialStore struct{}

func (s *noopCredentialStore) Basic(url *url.URL) (string, string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return "", ""
}
func (s *noopCredentialStore) RefreshToken(url *url.URL, service string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return ""
}
func (s *noopCredentialStore) SetRefreshToken(url *url.URL, service string, token string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
}
func NewBasicCredentials() *BasicCredentials {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &BasicCredentials{refreshTokenStore: &refreshTokenStore{}}
}

type basicForURL struct {
	url                url.URL
	username, password string
}
type BasicCredentials struct {
	creds []basicForURL
	*refreshTokenStore
}

func (c *BasicCredentials) Add(url *url.URL, username, password string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	c.creds = append(c.creds, basicForURL{*url, username, password})
}
func (c *BasicCredentials) Basic(url *url.URL) (string, string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for _, cred := range c.creds {
		if len(cred.url.Host) != 0 && cred.url.Host != url.Host {
			continue
		}
		if len(cred.url.Path) != 0 && cred.url.Path != url.Path {
			continue
		}
		return cred.username, cred.password
	}
	return "", ""
}
