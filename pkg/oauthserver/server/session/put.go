package session

import (
	"k8s.io/apiserver/pkg/authentication/user"
	"net/http"
	"time"
)

func putUser(store Store, w http.ResponseWriter, user user.Info, expiresIn time.Duration) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	values := Values{}
	values[userNameKey] = user.GetName()
	values[userUIDKey] = user.GetUID()
	var expires int64
	if expiresIn > 0 {
		expires = time.Now().Add(expiresIn).Unix()
	}
	values[expKey] = expires
	return store.Put(w, values)
}
