package integration

import (
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"github.com/gorilla/context"
	"k8s.io/klog"
	knet "k8s.io/apimachinery/pkg/util/net"
)

type User struct {
	ID		string
	Password	string
	Name		string
	Email		string
}
type BasicAuthChallenger struct {
	realm			string
	users			map[string]User
	authenticatedHandler	http.Handler
}

func NewBasicAuthChallenger(realm string, users []User, authenticatedHandler http.Handler) BasicAuthChallenger {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	userMap := make(map[string]User, len(users))
	for _, user := range users {
		userMap[user.ID] = user
	}
	return BasicAuthChallenger{realm, userMap, authenticatedHandler}
}
func (challenger BasicAuthChallenger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	klog.Infof("--- %v\n", r.URL)
	authHeader := r.Header["Authorization"]
	if len(authHeader) == 0 {
		challenger.Challenge(w)
		return
	}
	auth := strings.SplitN(authHeader[0], " ", 2)
	if len(auth) != 2 || auth[0] != "Basic" {
		Error("bad syntax", http.StatusBadRequest, w)
		return
	}
	payload, err := base64.StdEncoding.DecodeString(auth[1])
	if err != nil {
		Error("bad syntax", http.StatusBadRequest, w)
		return
	}
	pair := strings.SplitN(string(payload), ":", 2)
	if len(pair) != 2 {
		Error("bad syntax", http.StatusBadRequest, w)
		return
	}
	if !challenger.Validate(pair[0], pair[1]) {
		challenger.Challenge(w)
		return
	}
	context.Set(r, "username", challenger.users[pair[0]])
	challenger.authenticatedHandler.ServeHTTP(w, r)
}
func (challenger *BasicAuthChallenger) Challenge(w http.ResponseWriter) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	klog.Infof("Sending challenge\n")
	w.Header().Set("WWW-Authenticate", "Basic realm=\""+challenger.realm+"\"")
	Error("Authorization failed", http.StatusUnauthorized, w)
}
func Error(message string, status int, w http.ResponseWriter) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	klog.Infof("Writing error: %s\n", message)
	data := map[string]string{"error": message}
	json, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	io.WriteString(w, string(json))
}
func (challenger *BasicAuthChallenger) Validate(username, password string) bool {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	knownUser, exists := challenger.users[username]
	if !exists {
		return false
	}
	if knownUser.Password == password {
		klog.Infof("Validated user: %s\n", username)
		return true
	}
	klog.Infof("Rejected user: %s\n", username)
	return false
}

type identifyingHandler struct{}

func (handler *identifyingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	user := context.Get(r, "username")
	if user == nil {
		Error("No user found", http.StatusBadRequest, w)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	result, _ := json.Marshal(user)
	io.WriteString(w, string(result))
}
func NewIdentifyingHandler() http.Handler {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &identifyingHandler{}
}

type xRemoteUserProxyingHandler struct{ proxier *httputil.ReverseProxy }

func (handler *xRemoteUserProxyingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	user := context.Get(r, "username")
	if user == nil {
		Error("No user found", http.StatusBadRequest, w)
		return
	}
	r.Header.Add("X-Remote-User", user.(User).ID)
	handler.proxier.ServeHTTP(w, r)
}
func NewXRemoteUserProxyingHandler(rawURL string) http.Handler {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	parsedURL, _ := url.Parse(rawURL)
	proxier := httputil.NewSingleHostReverseProxy(parsedURL)
	proxier.Transport = insecureTransport()
	return &xRemoteUserProxyingHandler{proxier}
}
func insecureTransport() *http.Transport {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	return knet.SetTransportDefaults(&http.Transport{TLSClientConfig: &tls.Config{MinVersion: tls.VersionTLS10, InsecureSkipVerify: true}})
}
