package teststorage

import (
	"errors"
	goformat "fmt"
	"github.com/RangelReale/osin"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

type Test struct {
	Clients       map[string]osin.Client
	AuthorizeData *osin.AuthorizeData
	Authorize     map[string]*osin.AuthorizeData
	AccessData    *osin.AccessData
	Access        map[string]*osin.AccessData
	Err           error
}

func New() *Test {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &Test{Clients: make(map[string]osin.Client), Authorize: make(map[string]*osin.AuthorizeData), Access: make(map[string]*osin.AccessData)}
}
func (t *Test) Clone() osin.Storage {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return t
}
func (t *Test) Close() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
}
func (t *Test) GetClient(id string) (osin.Client, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return t.Clients[id], t.Err
}
func (t *Test) SaveAuthorize(data *osin.AuthorizeData) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	t.AuthorizeData = data
	t.Authorize[data.Code] = data
	return t.Err
}
func (t *Test) LoadAuthorize(code string) (*osin.AuthorizeData, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return t.Authorize[code], t.Err
}
func (t *Test) RemoveAuthorize(code string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	delete(t.Authorize, code)
	return t.Err
}
func (t *Test) SaveAccess(data *osin.AccessData) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	t.AccessData = data
	t.Access[data.AccessToken] = data
	return t.Err
}
func (t *Test) LoadAccess(token string) (*osin.AccessData, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return t.Access[token], t.Err
}
func (t *Test) RemoveAccess(token string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	delete(t.Access, token)
	return t.Err
}
func (t *Test) LoadRefresh(token string) (*osin.AccessData, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for _, v := range t.Access {
		if v.RefreshToken == token {
			return v, t.Err
		}
	}
	return nil, errors.New("not found")
}
func (t *Test) RemoveRefresh(token string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	data, _ := t.LoadRefresh(token)
	if data != nil {
		delete(t.Access, data.AccessToken)
	}
	return t.Err
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
